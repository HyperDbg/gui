// Package core implements the central Debugger state machine that ties together
// driver loading, IOCTL communication, event registration and hook management.
//
// The Debugger struct is the single owner of all mutable state; there are no
// package-level globals (see API design spec). It is safe for concurrent use
// from a single goroutine; the higher-level api.Debugger wraps it with a mutex
// for multi-goroutine access (GUI/MCP).
package core

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
	"unsafe"

	astencoder "github.com/hyperdbg/go-bridge/ast"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
	"github.com/hyperdbg/go-libhyperdbg/types"
)

// DebuggerState tracks the high-level state machine.
type DebuggerState int

const (
	StateDisconnected DebuggerState = iota
	StateConnected
	StateVmmLoaded
	StateProcessRunning
	StateProcessPaused
)

// Debugger is the core debugger instance. It owns the driver service, the
// device handle, and the registered hook table.
type Debugger struct {
	mu     sync.Mutex
	state  DebuggerState
	driver *driverloader.Driver
	device *comm.Device

	// nextTag is the monotonically increasing event tag. Tags identify
	// event+action pairs in the driver.
	nextTag uint64

	// logFile is the currently open log file (nil if none).
	logFile WriteCloser

	// processToken is the kernel-side debugging handle returned by the
	// ATTACH IOCTL (see attachProcess). It is 0 when no debuggee is
	// attached. Continue/Pause/Command IOCTLs require this token; it is
	// cleared by UnloadVMM/Close.
	processToken uint64

	// processPid is the PID of the debuggee process (set by StartProcess).
	// Used as the Pid field in DEBUGGER_READ_MEMORY / EDIT_MEMORY / etc.
	// The kernel's MemoryManagerReadProcessMemoryNormal calls
	// PsLookupProcessByProcessId(Pid, ...) — passing 0 fails with
	// DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER (0xC000003C) because
	// there is no process with PID 0 that can be looked up that way.
	// This mirrors the C++ side: g_ActiveProcessDebuggingState.ProcessId.
	processPid uint32

	// pausedRIP/pausedRFLAGS 由 MessagePump 从 DEBUGGEE_UD_PAUSED_PACKET 中提取。
	// ReadRegisters 返回这些值（因为 IOCTL 只返回部分 GUEST_REGS，RIP 不在其中）。
	pausedRIP      uint64
	pausedRFLAGS   uint64
	pausedThreadId uint32 // 从 DEBUGGEE_UD_PAUSED_PACKET 获取，用于 Step 的 TargetThreadId

	// OnPaused is called (from the MessagePump goroutine) when a
	// DEBUGGEE_UD_PAUSED_PACKET is received — i.e. the debuggee has
	// paused (breakpoint hit, single-step complete, manual Pause, OEP,
	// etc.). UI layers set this to auto-refresh registers/disasm/stack
	// without waiting for the user to click a button. Set to nil to
	// disable. Called outside d.mu; the callback must not call core
	// methods that could deadlock (use a goroutine or channel).
	OnPaused func()

	// pauseEvent is signaled (non-blocking) by MessagePump when a
	// DEBUGGEE_UD_PAUSED_PACKET is received. Step/TraceInto/StepOver
	// use this to wait for the step to complete after sending the IOCTL
	// with WaitForEventCompletion=false (matching the C libhyperdbg
	// approach, which uses DbgWaitForUserResponse instead of the
	// kernel's WaitForEventCompletion flag that can deadlock).
	pauseEvent chan struct{}
}

// WriteCloser is the minimal interface for log output (os.File implements it).
type WriteCloser interface {
	Write([]byte) (int, error)
	Close() error
}

// New creates a Debugger. It does not connect or load any driver; call
// Connect + LoadVMM explicitly.
func New() *Debugger {
	return &Debugger{state: StateDisconnected}
}

// State returns the current debugger state (thread-safe).
func (d *Debugger) State() DebuggerState {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.state
}

// Connect opens the HyperDbg device. For "local" target this opens
// \\.\HyperDbgDebuggerDevice. The device must already be created by a loaded
// VMM driver; call LoadVMM first if the driver is not yet running.
func (d *Debugger) Connect(ctx context.Context, target string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state != StateDisconnected {
		return fmt.Errorf("Connect: already connected (state=%d)", d.state)
	}
	dev, err := comm.Open(ctx, comm.DeviceName)
	if err != nil {
		return fmt.Errorf("Connect(%q): %w", target, err)
	}
	d.device = dev
	d.state = StateConnected
	return nil
}

// LoadVMM installs and starts the VMM driver (hyperhv/hyperkd), connects
// to the device, and initializes the VMM. driverPath is the absolute path
// to the .sys file.
//
// If the Debugger is Disconnected, this method will:
//  1. Install+start the driver service (driverloader.Load)
//  2. Open the device (comm.Open) — the device is created by DriverEntry
//  3. Send IOCTL_INIT_VMM
//
// If the Debugger is already Connected (device open but VMM not yet
// initialized), only steps 1 and 3 are performed.
func (d *Debugger) LoadVMM(ctx context.Context, driverPath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Step 1: Install and start the driver service.
	d.driver = driverloader.NewDriver(driverPath)
	if err := d.driver.Load(ctx); err != nil {
		return fmt.Errorf("LoadVMM: %w", err)
	}

	// Step 2: Open the device if not already connected.
	if d.state == StateDisconnected {
		dev, err := comm.Open(ctx, comm.DeviceName)
		if err != nil {
			// Driver loaded but device not yet created? Retry with delay.
			for i := 0; i < 5; i++ {
				time.Sleep(500 * time.Millisecond)
				dev, err = comm.Open(ctx, comm.DeviceName)
				if err == nil {
					break
				}
			}
			if err != nil {
				return fmt.Errorf("LoadVMM: open device after driver load: %w", err)
			}
		}
		d.device = dev
		d.state = StateConnected
	}

	if d.state != StateConnected && d.state < StateVmmLoaded {
		return fmt.Errorf("LoadVMM: unexpected state %d", d.state)
	}

	// Step 3: Initialize the VMM.
	// IOCTL_INIT_VMM requires a DEBUGGER_INIT_VMM_PACKET (METHOD_BUFFERED
	// alias, input = output). The driver reads KernelStatus from the input
	// buffer and writes the result back. Sending nil causes STATUS_INVALID_PARAMETER.
	var vmmPacket uint32 // DEBUGGER_INIT_VMM_PACKET.KernelStatus
	vmmSize := uint32(unsafe.Sizeof(vmmPacket))
	if _, err := d.device.IoctlStruct(ctx, comm.IOCTL_CODE_INIT_VMM,
		unsafe.Pointer(&vmmPacket), unsafe.Pointer(&vmmPacket), vmmSize, vmmSize); err != nil {
		return fmt.Errorf("LoadVMM: IOCTL_INIT_VMM failed: %w", err)
	}
	const debuggerOperationWasSuccessful uint32 = 0xFFFFFFFF
	if vmmPacket != debuggerOperationWasSuccessful {
		return fmt.Errorf("LoadVMM: VMM init failed (KernelStatus=0x%08X)", vmmPacket)
	}
	d.state = StateVmmLoaded
	return nil
}

// UnloadVMM sends IOCTL_TERMINATE_VMX and stops+removes the driver service.
//
// If a debuggee is attached, the cleanup sequence is:
//  1. Continue (resume any paused threads) — AttachingPerformDetach rejects
//     detach with DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS
//     while threads are paused, and a failed detach leaves the PEB monitor
//     EPT hook in g_EptState->HookedPagesList. TERMINATE_VMX then crashes
//     in EptHookUnHookAll → EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList
//     because that hook has IsHiddenBreakpoint == FALSE but no !epthook2
//     was ever registered (g_EptHook2sDetourListHead is uninitialized).
//  2. Detach (best-effort, removes the PEB monitor hook).
//  3. TERMINATE_VMX.
//  4. Unload the driver service.
//
// Errors from the continue/detach path are swallowed because the unload
// itself is more important: a stale token after UnloadVMM is unrecoverable.
func (d *Debugger) UnloadVMM(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	// Step 1: Clear all registered events. If any EPT hooks are still
	// active when TERMINATE_VMX is called, the EPT hook cleanup code
	// (EptHookUnHookAll) may crash or leave VMX in a bad state, which
	// prevents the driver from being loaded again without a reboot.
	if d.device != nil && d.state >= StateVmmLoaded {
		clearAll := types.DEBUGGER_MODIFY_EVENTS{
			Tag:          0,
			TypeOfAction: types.DebuggerModifyEventsClear,
		}
		clearBuf, err := structToBytes(&clearAll)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] MODIFY_EVENTS Clear structToBytes failed: %v\n", err)
		} else {
			var dummy [256]byte
			n, ioctlErr := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_MODIFY_EVENTS, clearBuf, dummy[:])
			if ioctlErr != nil {
				fmt.Fprintf(os.Stderr, "[UnloadVMM] MODIFY_EVENTS Clear failed: %v (n=%d)\n", ioctlErr, n)
			} else {
				fmt.Fprintf(os.Stderr, "[UnloadVMM] MODIFY_EVENTS Clear OK\n")
			}
		}
	}

	// Step 2: Resume + detach attached process.
	// 关键：detach 必须成功才能继续 TERMINATE_VMX，否则被拦截的
	// csrss/系统线程状态不一致 → 堆元数据损坏 → CRITICAL_PROCESS_DIED BSOD。
	// 重试多次，仍失败则跳过 TERMINATE_VMX（驱动残留好过蓝屏）。
	detachOk := true
	if d.device != nil && d.processToken != 0 {
		fmt.Fprintf(os.Stderr, "[UnloadVMM] Continuing process (token=0x%X)...\n", d.processToken)
		if err := continueProcess(ctx, d.device, d.processToken); err != nil {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] Continue failed: %v\n", err)
		}
		const maxDetachRetries = 5
		for attempt := 1; attempt <= maxDetachRetries; attempt++ {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] Detaching process (token=0x%X) attempt %d/%d...\n",
				d.processToken, attempt, maxDetachRetries)
			if err := detachProcess(ctx, d.device, d.processToken); err != nil {
				fmt.Fprintf(os.Stderr, "[UnloadVMM] Detach failed: %v\n", err)
				detachOk = false
				if attempt < maxDetachRetries {
					d.mu.Unlock()
					time.Sleep(time.Duration(attempt) * time.Second)
					d.mu.Lock()
				}
			} else {
				fmt.Fprintf(os.Stderr, "[UnloadVMM] Detach OK\n")
				detachOk = true
				break
			}
		}
		d.processToken = 0
		d.processPid = 0
	}

	if !detachOk {
		fmt.Fprintf(os.Stderr, "[UnloadVMM] Detach failed after retries — skipping TERMINATE_VMX to avoid BSOD (csrss heap corruption). Driver will residue; reboot to clean up.\n")
		d.state = StateConnected
		return fmt.Errorf("UnloadVMM: detach failed, skipped TERMINATE_VMX to avoid BSOD")
	}

	// Step 3: Terminate VMX with retry.
	// If TERMINATE_VMX fails (e.g. a core is stuck in VMX root),
	// VT-x remains locked and the driver cannot be reloaded without reboot.
	// We retry up to 3 times with increasing delays.
	if d.device != nil && d.state >= StateVmmLoaded {
		const maxRetries = 3
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] TERMINATE_VMX attempt %d/%d...\n", attempt, maxRetries)
			n, ioctlErr := d.device.Ioctl(ctx, comm.IOCTL_CODE_TERMINATE_VMX, nil, nil)
			if ioctlErr == nil {
				fmt.Fprintf(os.Stderr, "[UnloadVMM] TERMINATE_VMX OK\n")
				break
			}
			fmt.Fprintf(os.Stderr, "[UnloadVMM] TERMINATE_VMX failed: %v (n=%d)\n", ioctlErr, n)
			if attempt < maxRetries {
				delay := time.Duration(attempt*2) * time.Second
				fmt.Fprintf(os.Stderr, "[UnloadVMM] Waiting %v before retry...\n", delay)
				d.mu.Unlock()
				time.Sleep(delay)
				d.mu.Lock()
			}
		}
		// Give the kernel time to complete EPT hook cleanup and VMXOFF
		// on all cores (asynchronous VMX root → non-root transitions).
		d.mu.Unlock()
		time.Sleep(3 * time.Second)
		d.mu.Lock()
	}

	// Step 4: Unload the driver service.
	if d.driver != nil {
		fmt.Fprintf(os.Stderr, "[UnloadVMM] Unloading driver service...\n")
		if err := d.driver.Unload(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] Driver unload failed: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "[UnloadVMM] Driver unloaded OK\n")
		}
	}
	d.state = StateConnected
	return nil
}

// Close releases all resources (device handle, driver service, log file).
// If a debuggee is attached, the same continue→detach sequence as UnloadVMM
// is performed first to avoid the BSOD described in UnloadVMM.
func (d *Debugger) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken != 0 && d.device != nil {
		_ = continueProcess(context.Background(), d.device, d.processToken)
		_ = detachProcess(context.Background(), d.device, d.processToken)
		d.processToken = 0
		d.processPid = 0
	}
	if d.logFile != nil {
		_ = d.logFile.Close()
		d.logFile = nil
	}
	if d.device != nil {
		_ = d.device.Close()
		d.device = nil
	}
	d.state = StateDisconnected
	return nil
}

// LogOpen opens a file for log output. Printf-style log lines from hooks are
// written here.
func (d *Debugger) LogOpen(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Use os.Create through a wrapper to avoid importing os in the core.
	f, err := openFileForLog(path)
	if err != nil {
		return fmt.Errorf("LogOpen(%q): %w", path, err)
	}
	d.logFile = f
	return nil
}

// LogClose closes the log file.
func (d *Debugger) LogClose() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.logFile == nil {
		return nil
	}
	err := d.logFile.Close()
	d.logFile = nil
	return err
}

// WriteLog writes raw bytes to the log file (if open).
func (d *Debugger) WriteLog(p []byte) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.logFile == nil {
		return 0, fmt.Errorf("no log file open")
	}
	return d.logFile.Write(p)
}

// EptHook registers an EPT execution hook (detours-style) at the given
// address with a Go callback. The callback source is compiled to the binary
// AST wire format (go-bridge/ast) and sent to the driver as the script buffer
// of a RunScript action.
//
// Returns the event tag (hook ID) on success.
func (d *Debugger) EptHook(ctx context.Context, hookAddress uint64, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("EptHook: VMM not loaded")
	}

	// 1. Encode the Go callback to binary AST.
	astBytes, err := astencoder.Encode(callbackSrc)
	if err != nil {
		return 0, fmt.Errorf("EptHook: callback compile failed: %w", err)
	}

	// 2. Register the event (IOCTL_DEBUGGER_REGISTER_EVENT).
	tag := d.nextTag
	d.nextTag++

	event := types.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: 0xFFFFFFFF, // DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES
		IsEnabled: true,
		EventType: types.HiddenHookExecCc, // !epthook = EPT exec CC (safe for CET shadow stack)
		Tag:       tag,
		Options: types.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: hookAddress, // target address
		},
	}

	eventBuf, err := structToBytes(&event)
	if err != nil {
		return 0, err
	}
	var result types.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf, err := structToBytes(&result)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("EptHook: REGISTER_EVENT IOCTL failed: %w", err)
	}
	// Re-read result from the output buffer.
	if err := bytesToStruct(resultBuf, &result); err != nil {
		return 0, err
	}
	if !result.IsSuccessful {
		return 0, fmt.Errorf("EptHook: event registration failed, error=%d", result.Error)
	}

	// 3. Add action with the Go AST script (IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT).
	// The action struct is followed immediately by the script bytes in the
	// same buffer.
	action := types.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          types.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes, err := structToBytes(&action)
	if err != nil {
		return 0, err
	}
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult types.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf, err := structToBytes(&actionResult)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("EptHook: ADD_ACTION IOCTL failed: %w", err)
	}
	if err := bytesToStruct(actionResultBuf, &actionResult); err != nil {
		return 0, err
	}
	if !actionResult.IsSuccessful {
		return 0, fmt.Errorf("EptHook: action registration failed, error=%d", actionResult.Error)
	}

	return tag, nil
}

// EptHookForProcess registers an EPT execution hook for a specific process.
// It is identical to EptHook except that ProcessId is set to pid instead of
// 0 (all processes). This is required for hooking user-mode addresses: the
// kernel's ValidateEventEptHookHiddenBreakpointAndInlineHooks calls
// VirtualAddressToPhysicalAddressByProcessId(addr, pid) to validate the hook
// target, and pid=0 defaults to the System process (PID 4) whose address
// space does not contain user-mode DLLs like ntdll.dll.
//
// Returns the event tag (hook ID) on success.
func (d *Debugger) EptHookForProcess(ctx context.Context, hookAddress uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("EptHookForProcess: VMM not loaded")
	}

	astBytes, err := astencoder.Encode(callbackSrc)
	if err != nil {
		return 0, fmt.Errorf("EptHookForProcess: callback compile failed: %w", err)
	}

	tag := d.nextTag
	d.nextTag++

	event := types.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: pid,
		IsEnabled: true,
		EventType: types.HiddenHookExecCc,
		Tag:       tag,
		Options: types.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: hookAddress,
		},
	}

	eventBuf, err := structToBytes(&event)
	if err != nil {
		return 0, err
	}
	var result types.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf, err := structToBytes(&result)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("EptHookForProcess: REGISTER_EVENT IOCTL failed: %w", err)
	}
	if err := bytesToStruct(resultBuf, &result); err != nil {
		return 0, err
	}
	if !result.IsSuccessful {
		return 0, fmt.Errorf("EptHookForProcess: event registration failed, error=%d", result.Error)
	}

	action := types.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          types.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes, err := structToBytes(&action)
	if err != nil {
		return 0, err
	}
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult types.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf, err := structToBytes(&actionResult)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("EptHookForProcess: ADD_ACTION IOCTL failed: %w", err)
	}
	if err := bytesToStruct(actionResultBuf, &actionResult); err != nil {
		return 0, err
	}
	if !actionResult.IsSuccessful {
		return 0, fmt.Errorf("EptHookForProcess: action registration failed, error=%d", actionResult.Error)
	}

	return tag, nil
}

// MonitorReadForProcess registers a hidden read-hook on the virtual address
// range [addrStart, addrEnd) for a specific process. This is the Go
// equivalent of the HyperDbg !monitor r command: the EPT read-access
// violation fires whenever the target process reads any byte in the range,
// and the callback runs in VMX-root mode.
//
// The event type is HiddenHookRead (HIDDEN_HOOK_READ in the C SDK);
// OptionalParam1 = start address, OptionalParam2 = end address.
// The kernel validates both addresses via
// VirtualAddressToPhysicalAddressByProcessId, so the target process must
// already be running (pid > 0).
//
// Returns the event tag on success.
func (d *Debugger) MonitorReadForProcess(ctx context.Context, addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("MonitorReadForProcess: VMM not loaded")
	}

	astBytes, err := astencoder.Encode(callbackSrc)
	if err != nil {
		return 0, fmt.Errorf("MonitorReadForProcess: callback compile failed: %w", err)
	}

	tag := d.nextTag
	d.nextTag++

	event := types.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: pid,
		IsEnabled: true,
		EventType: types.HiddenHookRead,
		Tag:       tag,
		Options: types.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: addrStart,
			OptionalParam2: addrEnd,
		},
	}

	eventBuf, err := structToBytes(&event)
	if err != nil {
		return 0, err
	}
	var result types.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf, err := structToBytes(&result)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("MonitorReadForProcess: REGISTER_EVENT IOCTL failed: %w", err)
	}
	if err := bytesToStruct(resultBuf, &result); err != nil {
		return 0, err
	}
	if !result.IsSuccessful {
		return 0, fmt.Errorf("MonitorReadForProcess: event registration failed, error=%d", result.Error)
	}

	action := types.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          types.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes, err := structToBytes(&action)
	if err != nil {
		return 0, err
	}
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult types.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf, err := structToBytes(&actionResult)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("MonitorReadForProcess: ADD_ACTION IOCTL failed: %w", err)
	}
	if err := bytesToStruct(actionResultBuf, &actionResult); err != nil {
		return 0, err
	}
	if !actionResult.IsSuccessful {
		return 0, fmt.Errorf("MonitorReadForProcess: action registration failed, error=%d", actionResult.Error)
	}

	return tag, nil
}

// MessagePump owns a goroutine that receives kernel→user messages (e.g. LogInfo
// output produced by hook callbacks via ctx.Printf) by repeatedly issuing
// IOCTL_REGISTER_EVENT with Type=IRP_BASED on a DEDICATED device handle.
//
// The dedicated handle is required because the IRP stays pending in the kernel
// until a message arrives — using the main IOCTL handle would block all other
// commands (Continue/Pause/EptHook/…).
//
// Each received packet is laid out as DEBUGGEE_MESSAGE_PACKET:
//
//	UINT32 OperationCode;   // 1=LogInfo, 2=LogWarning, 3=LogError, 4=LogNonImmediate, …
//	CHAR  Message[4096];    // NUL-terminated body
//
// The pump writes Message (trailing NULs stripped, '\n' appended) to the
// Debugger's open log file. Messages with the END_OF_IRPS OperationCode (set by
// the kernel when the driver is being unloaded) cause the loop to exit.
//
// Lifecycle:
//
//	mp, _ := dbg.StartMessagePump(ctx)   // spawns the goroutine
//	...                                  // hooks fire, log file fills
//	mp.Stop()                            // signal + wait for goroutine
//
// Stop sends IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL on the MAIN
// device handle (mirrors the C++ HyperDbgClose path in libhyperdbg.cpp:534).
// The kernel then completes the pending IRP with
// OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS, the goroutine sees that opCode and
// exits, and its deferred Close releases the dedicated handle. Stop MUST be
// called BEFORE UnloadVMM (which stops the driver service and makes the main
// device handle unusable).
type MessagePump struct {
	// mainDev is the Debugger's main IOCTL handle, used only to send the
	// END_OF_IRPS signal in Stop. The pump never sends other IOCTLs on it.
	mainDev *comm.Device
	// dev is the dedicated handle the pump blocks on.
	dev  *comm.Device
	stop chan struct{}
	done chan struct{}
}

// StartMessagePump spawns a goroutine that drains kernel log messages to the
// open log file (see LogOpen). It must be called AFTER LogOpen and AFTER
// LoadVMM; the caller must invoke (*MessagePump).Stop BEFORE UnloadVMM so the
// dedicated device handle is released cleanly and the main handle is still
// usable for the END_OF_IRPS signal.
func (d *Debugger) StartMessagePump(ctx context.Context) (*MessagePump, error) {
	d.mu.Lock()
	if d.state < StateVmmLoaded {
		d.mu.Unlock()
		return nil, fmt.Errorf("StartMessagePump: VMM not loaded")
	}
	if d.logFile == nil {
		d.mu.Unlock()
		return nil, fmt.Errorf("StartMessagePump: no log file open (call LogOpen first)")
	}
	if d.pauseEvent == nil {
		d.pauseEvent = make(chan struct{}, 1)
	}
	mainDev := d.device
	d.mu.Unlock()

	// Open a dedicated device handle so the pending IRP does not block the
	// main IOCTL handle used by Continue/Pause/EptHook/… (mirrors the C++
	// ReadIrpBasedBuffer opening a second \\.\HyperDbgDebuggerDevice).
	dev, err := comm.Open(ctx, comm.DeviceName)
	if err != nil {
		return nil, fmt.Errorf("StartMessagePump: open dedicated device: %w", err)
	}

	mp := &MessagePump{
		mainDev: mainDev,
		dev:     dev,
		stop:    make(chan struct{}),
		done:    make(chan struct{}),
	}
	go mp.run(ctx, d)
	return mp, nil
}

// run is the IRP-reading loop. It mirrors ReadIrpBasedBuffer in
// go-libhyperdbg/app/packets.go but writes packets straight to the log file
// instead of routing them through the Messaging dispatcher.
func (mp *MessagePump) run(ctx context.Context, d *Debugger) {
	defer close(mp.done)
	defer mp.dev.Close()

	// REGISTER_NOTIFY_BUFFER{Type: IRP_BASED, HEvent: NULL}.
	var reg types.REGISTER_NOTIFY_BUFFER
	reg.Type = types.IrpBased
	regBuf := (*[unsafe.Sizeof(reg)]byte)(unsafe.Pointer(&reg))[:]

	// UsermodeBufferSize = sizeof(UINT32) + PacketChunkSize + 1 = 4101.
	// See hyperdbg/include/SDK/headers/Constants.h.
	out := make([]byte, 4+4096+1)

	// operationHypervisorDriverEndOfIrps — the kernel completes the pending
	// IRP with this code when Stop signals it via
	// IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL (see Ioctl.c:337).
	const operationHypervisorDriverEndOfIrps uint32 = 14 | (1 << 31)
	// operationNotificationFromUserDebuggerPause — the kernel sends this
	// when the user-mode debugger pauses the debuggee (Step/Pause/bp hit).
	// The body is a DEBUGGEE_UD_PAUSED_PACKET. Value mirrors
	// OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE in
	// hyperdbg/include/SDK/headers/Constants.h (16 | mandatory bit).
	const operationNotificationFromUserDebuggerPause uint32 = 16 | (1 << 31)

	for {
		select {
		case <-mp.stop:
			return
		default:
		}

		// Zero the buffer (mirrors ZeroMemory in the C++ loop).
		for i := range out {
			out[i] = 0
		}

		// Synchronous IOCTL — blocks until the kernel has a packet to
		// deliver. Unblocked by either (a) a real message arriving or
		// (b) Stop sending IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL
		// on the main handle, which causes the kernel to complete this IRP
		// with OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS.
		n, err := mp.dev.Ioctl(ctx, comm.IOCTL_CODE_REGISTER_EVENT, regBuf, out)
		if err != nil {
			// IRP cancelled (Stop closed mp.dev) or driver gone — exit if
			// stop was signaled, otherwise continue (mirrors C++ behaviour
			// where a flush command can transiently fail the IRP).
			select {
			case <-mp.stop:
				return
			default:
				continue
			}
		}
		if n < 4 {
			continue
		}

		opCode := binary.LittleEndian.Uint32(out[:4])
		if opCode == operationHypervisorDriverEndOfIrps {
			return
		}

		// Message body: bytes [4:n].
		msg := out[4:n]

		// 严格按 opCode 判断是否为 DEBUGGEE_UD_PAUSED_PACKET。
		// 之前用 "Rip 在用户空间范围" 启发式判断会被 LogInfo 等文本
		// 消息误判（前 8 字节恰好是用户态地址时），导致 pauseEvent
		// 被错误信号、Step 在错误时刻返回 → 后续单步超时。
		if opCode == operationNotificationFromUserDebuggerPause {
			pausedSize := unsafe.Sizeof(types.DEBUGGEE_UD_PAUSED_PACKET{})
			if uint32(len(msg)) >= uint32(pausedSize) {
				paused := (*types.DEBUGGEE_UD_PAUSED_PACKET)(unsafe.Pointer(&msg[0]))
				d.mu.Lock()
				d.pausedRIP = paused.Rip
				d.pausedRFLAGS = paused.Rflags
				d.pausedThreadId = paused.ThreadId
				cb := d.OnPaused
				// Signal pauseEvent (non-blocking: if channel is full,
				// drain first so the latest pause is what Step waits on).
				if d.pauseEvent != nil {
					select {
					case <-d.pauseEvent:
					default:
					}
					d.pauseEvent <- struct{}{}
				}
				d.mu.Unlock()
				// 在 mu 外部调用回调，避免回调中的 core 方法死锁。
				// 回调运行在 pump goroutine 中，应只做异步刷新（如
				// cpuPage.Refresh() 本身就是 go refreshInternal()）。
				if cb != nil {
					cb()
				}
			}
			// PAUSED 包是二进制结构，不写入文本日志
			continue
		}

		// Strip trailing NULs for text log output
		for i, b := range msg {
			if b == 0 {
				msg = msg[:i]
				break
			}
		}
		if len(msg) == 0 {
			continue
		}

		// Append '\n' so each kernel message is on its own line (matches
		// the format expected by ParseAPILog which uses bufio.Scanner).
		_, _ = d.WriteLog(append(msg, '\n'))
	}
}

// Stop signals the pump goroutine to exit and waits for it to return. It sends
// IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL on the MAIN device handle
// (mirrors the C++ HyperDbgClose path), which causes the kernel to complete the
// pump's pending IRP with OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS. The goroutine
// then exits and its deferred Close releases the dedicated handle.
//
// Stop MUST be called BEFORE UnloadVMM (UnloadVMM stops the driver service and
// makes the main handle unusable for the END_OF_IRPS signal). Safe to call
// multiple times — subsequent calls are no-ops.
func (mp *MessagePump) Stop() {
	select {
	case <-mp.stop:
		// Already stopped.
		return
	default:
		close(mp.stop)
	}

	// Signal the kernel to complete the pending IRP with END_OF_IRPS. This
	// is the only reliable way to unblock the synchronous DeviceIoControl
	// in the pump goroutine — closing mp.dev from another goroutine while
	// DeviceIoControl is blocked on it can deadlock (the driver's IRP_MJ_CLOSE
	// dispatcher waits for pending IRPs to drain). Sending the IOCTL on the
	// MAIN handle avoids that race because the kernel completes the pump's
	// IRP *before* returning from this IOCTL.
	if mp.mainDev != nil {
		_, _ = mp.mainDev.Ioctl(context.Background(),
			comm.IOCTL_CODE_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL, nil, nil)
	}

	// Wait for the goroutine to exit (it received END_OF_IRPS and returned,
	// or it hit the stop channel check). Bounded wait to avoid hanging the
	// test if the kernel never completes the IRP.
	select {
	case <-mp.done:
	case <-time.After(5 * time.Second):
		// Last-resort: force-close the dedicated handle in a goroutine so
		// Stop() does not block. DeviceIoControl is synchronous and has no
		// cancel routine, so Close() can deadlock in the driver's
		// IRP_MJ_CLOSE (which waits for the pending IRP). Running it
		// asynchronously lets Stop() return so the test fails fast instead
		// of hanging for the full test timeout. UnloadVMM may subsequently
		// fail because the dev handle is still open; that is acceptable
		// since the pump is wedged anyway. Close is safe to call twice —
		// the second call is a no-op (handle already 0).
		go mp.dev.Close()
	}
}

// ReadMemory reads size bytes from the target process at the given virtual address.
// This is a convenience wrapper around readmem.ReadMemory.
func (d *Debugger) ReadMemory(ctx context.Context, addr uint64, pid uint32, size uint32) ([]byte, types.DEBUGGER_READ_MEMORY_ADDRESS_MODE, error) {
	d.mu.Lock()
	dev := d.device
	d.mu.Unlock()
	if dev == nil {
		return nil, 0, fmt.Errorf("ReadMemory: not connected")
	}
	return readmem.ReadMemory(ctx, dev, addr, pid, size,
		types.DebuggerReadVirtualAddress, types.ReadFromKernel, false)
}

// Continue resumes execution of the debugged process by sending the
// CONTINUE_PROCESS IOCTL with the stored process token. The debuggee runs
// until a pause is requested (Pause) or a registered event fires.
//
// Mirrors the C libhyperdbg 'g' command path: UdContinueProcess(token).
func (d *Debugger) Continue(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken == 0 {
		return fmt.Errorf("Continue: no process attached (call StartProcess first)")
	}
	if err := continueProcess(ctx, d.device, d.processToken); err != nil {
		return fmt.Errorf("Continue: %w", err)
	}
	d.state = StateProcessRunning
	return nil
}

// Pause halts the debugged process by sending the PAUSE_PROCESS IOCTL with
// the stored process token. The kernel arranges for the debuggee to halt at
// the next instruction; the user-level listener receives the
// DEBUGGEE_UD_PAUSED_PACKET asynchronously.
//
// Returns ErrAlreadyPaused (non-fatal) if the debuggee is already in the
// intercepting phase — see attach.ErrAlreadyPaused for details. Callers that
// only care about "the process is or will be paused" can ignore
// errors.Is(err, ErrAlreadyPaused).
//
// Mirrors the C libhyperdbg 'pause' command path: UdPauseProcess(token)
// followed by CommandPauseRequest's silent handling of FALSE returns
// (pause.cpp:55).
func (d *Debugger) Pause(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken == 0 {
		return fmt.Errorf("Pause: no process attached (call StartProcess first)")
	}
	if err := pauseProcess(ctx, d.device, d.processToken); err != nil {
		if errors.Is(err, ErrAlreadyPaused) {
			// Already paused — not an error, matching C pause.cpp:55.
			d.state = StateProcessPaused
			return err
		}
		return fmt.Errorf("Pause: %w", err)
	}
	d.state = StateProcessPaused
	return nil
}

// StartProcess launches a process for debugging and attaches the VMM to it.
//
// The flow mirrors UdAttachToProcess (ud.cpp:380-480):
//  1. Create the child in a suspended state (CREATE_SUSPENDED |
//     CREATE_NEW_CONSOLE) so no user code runs before the attach completes.
//  2. Send IOCTL_CODE_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS with
//     Action=ATTACH and the child's pid/tid. The kernel registers the
//     debug session and returns a Token.
//  3. Store the Token for subsequent Continue/Pause IOCTLs.
//  4. Resume the main thread so the kernel can begin intercepting debug
//     events. The kernel pauses the debuggee at the first instruction (or
//     the entry point, depending on CheckCallbackAtFirstInstruction) and
//     waits for a Continue (g) before letting it run.
//
// The caller owns the returned Process and must Close it when done
// (typically via defer).
func (d *Debugger) StartProcess(ctx context.Context, exePath string) (Process, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return Process{}, fmt.Errorf("StartProcess: VMM not loaded")
	}

	// 1. Create the child in a suspended state.
	proc, err := createSuspendedProcess(exePath)
	if err != nil {
		return Process{}, fmt.Errorf("StartProcess: %w", err)
	}

	// 2. Attach the VMM to the suspended child. CheckCallbackAtFirstInstruction
	//    is true so the kernel pauses at the very first instruction; the user
	//    then calls Continue to reach the entry point and continue execution.
	//    This matches the .ds script flow: .start pauses, 'g' runs to entry
	//    point, second 'g' continues execution.
	token, err := attachProcess(ctx, d.device, proc.Pid, proc.Tid, true)
	if err != nil {
		// Best-effort cleanup of the suspended child on attach failure: kill
		// the process (it never ran) and close handles so we don't leak.
		_ = windowsTerminateProcess(proc.Handle)
		_ = proc.Close()
		return Process{}, fmt.Errorf("StartProcess: %w", err)
	}

	// 3. Store the token for Continue/Pause.
	d.processToken = token
	d.processPid = proc.Pid

	// 4. Resume the main thread. The kernel's attach callback has already
	//    registered the debug session, so the first instruction the thread
	//    executes will be intercepted and the debuggee will be reported as
	//    paused. The user calls Continue to let it actually run.
	if err := windowsResumeThread(proc.ThreadHandle); err != nil {
		// ResumeThread failed: the kernel attach succeeded but we cannot
		// resume the main thread. Detach and clean up so the user can retry.
		_ = detachProcess(ctx, d.device, token)
		d.processToken = 0
		d.processPid = 0
		_ = windowsTerminateProcess(proc.Handle)
		_ = proc.Close()
		return Process{}, fmt.Errorf("StartProcess: ResumeThread failed: %w", err)
	}

	d.state = StateProcessPaused
	return proc, nil
}

// Process represents a launched debuggee process.
type Process struct {
	// Handle is the Win32 process handle (owned by Process.Close).
	Handle uintptr
	// ThreadHandle is the Win32 main-thread handle (owned by Process.Close).
	// Required so StartProcess can ResumeThread after the attach IOCTL
	// succeeds; the user never needs to touch it directly.
	ThreadHandle uintptr
	Pid          uint32
	Tid          uint32
}

// binary.LittleEndian is referenced to keep the import for future use.
var _ = binary.LittleEndian

// The layout matches the C ABI because the types package preserves field
// alignment.
func structToBytes(v any) ([]byte, error) {
	// Use encoding/binary for a safer path when possible; for arbitrary structs
	// we fall back to unsafe.
	switch s := v.(type) {
	case *types.DEBUGGER_GENERAL_EVENT_DETAIL:
		var buf [unsafe.Sizeof(*s)]byte
		copy(buf[:], (*[unsafe.Sizeof(*s)]byte)(unsafe.Pointer(s))[:])
		return buf[:], nil
	case *types.DEBUGGER_GENERAL_ACTION:
		var buf [unsafe.Sizeof(*s)]byte
		copy(buf[:], (*[unsafe.Sizeof(*s)]byte)(unsafe.Pointer(s))[:])
		return buf[:], nil
	case *types.DEBUGGER_EVENT_AND_ACTION_RESULT:
		var buf [unsafe.Sizeof(*s)]byte
		copy(buf[:], (*[unsafe.Sizeof(*s)]byte)(unsafe.Pointer(s))[:])
		return buf[:], nil
	case *types.DEBUGGER_MODIFY_EVENTS:
		var buf [unsafe.Sizeof(*s)]byte
		copy(buf[:], (*[unsafe.Sizeof(*s)]byte)(unsafe.Pointer(s))[:])
		return buf[:], nil
	}
	return nil, fmt.Errorf("structToBytes: unsupported type %T", v)
}

// bytesToStruct deserialises a byte slice back into a pointer-to-struct.
func bytesToStruct(buf []byte, v any) error {
	switch s := v.(type) {
	case *types.DEBUGGER_EVENT_AND_ACTION_RESULT:
		size := unsafe.Sizeof(*s)
		if uint32(len(buf)) < uint32(size) {
			return fmt.Errorf("bytesToStruct: buffer too small (%d < %d)", len(buf), size)
		}
		copy((*[unsafe.Sizeof(*s)]byte)(unsafe.Pointer(s))[:], buf[:size])
		return nil
	}
	return fmt.Errorf("bytesToStruct: unsupported type %T", v)
}

// openFileForLog is set by the os import in file_log.go to avoid importing
// "os" in this file (keeps the core importable in constrained environments).
var openFileForLog func(path string) (WriteCloser, error)

// binary.LittleEndian is referenced to keep the import for future use.
var _ = binary.LittleEndian
