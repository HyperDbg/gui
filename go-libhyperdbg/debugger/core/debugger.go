// Package core implements the central Debugger state machine that ties together
// driver loading, IOCTL communication, event registration and hook management.
//
// The Debugger struct is the single owner of all mutable state; there are no
// package-level globals (see API design spec). It is safe for concurrent use
// from a single goroutine; the higher-level api.Debugger wraps it with a mutex
// for multi-goroutine access (GUI/MCP).
package core

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/hyperdbgsdk"
	astencoder "github.com/hyperdbg/go-bridge/ast"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
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

// Kernel IRP operation codes for the message pump. Mirrors
// OPERATION_* constants in hyperdbg/include/SDK/headers/Constants.h.
const (
	operationHypervisorDriverEndOfIrps            uint32 = 14 | (1 << 31)
	operationNotificationFromUserDebuggerPause    uint32 = 16 | (1 << 31)
)

// Debugger is the core debugger instance. It owns the driver service, the
// device handle, and the registered hook table.
type Debugger struct {
	mu     sync.Mutex
	state  DebuggerState
	driver *driverloader.Driver
	device *comm.Device

	// connected mirrors C++ g_IsConnectedToHyperDbgLocally: Connect() only
	// sets this flag, it does NOT open the device. The device is opened by
	// InitVMM (matching the C++ `load vmm` command, where
	// HyperDbgCreateHandleFromKdModule runs inside HyperDbgLoadVmmModule).
	// target records the connect argument ("local" or future remote addr).
	connected bool
	target    string

	// nextTag is the monotonically increasing event tag. Tags identify
	// event+action pairs in the driver.
	nextTag uint64

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

	// OnPacket is the single callback for every packet the kernel pushes
	// through the IRP-based channel (see MessagePump). It receives the raw
	// operation code and the message body (bytes after the 4-byte opCode
	// prefix). The PAUSED packet is handled internally by the pump (it
	// updates pausedRIP/pausedThreadId and signals pauseEvent + OnPaused)
	// and is also forwarded here so consumers can react (e.g. refresh UI).
	// All other packets (log messages, command output text, …) are
	// forwarded unchanged. Set to nil to drop kernel output.
	OnPacket func(opCode uint32, payload []byte)

	// pauseEvent is signaled (non-blocking) by MessagePump when a
	// DEBUGGEE_UD_PAUSED_PACKET is received. Step/TraceInto/StepOver
	// use this to wait for the step to complete after sending the IOCTL
	// with WaitForEventCompletion=false (matching the C libhyperdbg
	// approach, which uses DbgWaitForUserResponse instead of the
	// kernel's WaitForEventCompletion flag that can deadlock).
	pauseEvent chan struct{}
}

// New creates a Debugger. It does not connect or load any driver; call
// Connect (flag) + LoadDriver + InitVMM explicitly.
func New() *Debugger {
	return &Debugger{state: StateDisconnected}
}

// State returns the current debugger state (thread-safe).
func (d *Debugger) State() DebuggerState {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.state
}

// Connect marks this Debugger as connected to the given target. For "local"
// this mirrors C++ ConnectLocalDebugger (connect.cpp): it only sets the
// connected flag — it does NOT open the device. The device is opened later
// by InitVMM, matching the C++ `load vmm` command flow where
// HyperDbgCreateHandleFromKdModule runs inside HyperDbgLoadVmmModule (so
// `.connect local` then `load vmm` is the C++ pattern). This avoids the
// previous duplicate where both Connect and InitVMM opened the device.
func (d *Debugger) Connect(target string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.connected {
		return fmt.Errorf("Connect: already connected to %q", d.target)
	}
	d.target = target
	d.connected = true
	return nil
}

// LoadDriver installs and starts the kernel driver service.
func (d *Debugger) LoadDriver(driverPath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.driver = driverloader.NewDriver(driverPath)
	if err := d.driver.Load(); err != nil {
		return fmt.Errorf("LoadDriver: %w", err)
	}
	return nil
}

// InitVMM opens the device and sends IOCTL_INIT_VMM. This matches the C++
// `load vmm` command flow: HyperDbgCreateHandleFromKdModule (open device) +
// HyperDbgInitVmmModule (IOCTL_INIT_VMM). It opens the device only if not
// already open (StateDisconnected), so it's safe to call after Connect
// (which no longer opens the device) or directly after LoadDriver.
func (d *Debugger) InitVMM() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state == StateDisconnected {
		dev, err := comm.Open(comm.DeviceName)
		if err != nil {
			return fmt.Errorf("InitVMM: open device: %w", err)
		}
		d.device = dev
		d.state = StateConnected
	}
	var vmmPacket uint32
	vmmSize := uint32(unsafe.Sizeof(vmmPacket))
	if _, err := d.device.IoctlStruct(comm.IOCTL_CODE_INIT_VMM,
		unsafe.Pointer(&vmmPacket), unsafe.Pointer(&vmmPacket), vmmSize, vmmSize); err != nil {
		return fmt.Errorf("InitVMM: IOCTL_INIT_VMM failed: %w", err)
	}
	const debuggerOperationWasSuccessful uint32 = 0xFFFFFFFF
	if vmmPacket != debuggerOperationWasSuccessful {
		return fmt.Errorf("InitVMM: VMM init failed (KernelStatus=0x%08X)", vmmPacket)
	}
	d.state = StateVmmLoaded
	return nil
}

// UnloadVMM detaches the debuggee, sends IOCTL_TERMINATE_VMX and closes the
// device handle. This matches the C++ `unload vmm` command flow (unload.cpp):
// HyperDbgUnloadVmm (UdUninitializeUserDebugger + IOCTL_TERMINATE_VMX) +
// HyperDbgUnloadKd (close device). The device close lives here, not in
// UnloadDriver, so a mid-session UnloadVMM (without UnloadDriver) leaves the
// driver service intact for a later InitVMM restart.
func (d *Debugger) UnloadVMM() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.device != nil && d.processToken != 0 {
		_ = continueProcess(d.device, d.processToken)
		_ = detachProcess(d.device, d.processPid)
		d.processToken = 0
		d.processPid = 0
	}
	if d.device != nil && d.state >= StateVmmLoaded {
		_, _ = d.device.Ioctl(comm.IOCTL_CODE_TERMINATE_VMX, nil, nil)
	}
	if d.device != nil {
		_ = d.device.Close()
		d.device = nil
	}
	d.state = StateDisconnected
	return nil
}

// UnloadDriver stops and removes the driver service.
func (d *Debugger) UnloadDriver() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.driver != nil {
		if err := d.driver.Unload(); err != nil {
			return fmt.Errorf("UnloadDriver: %w", err)
		}
	}
	return nil
}

// EptHook registers an EPT execution hook (detours-style) at the given
// address with a Go callback. The callback source is compiled to the binary
// AST wire format (go-bridge/ast) and sent to the driver as the script buffer
// of a RunScript action.
//
// Returns the event tag (hook ID) on success.
func (d *Debugger) EptHook(hookAddress uint64, callbackSrc string) (uint64, error) {
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

	event := hyperdbgsdk.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: 0xFFFFFFFF, // DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES
		IsEnabled: true,
		EventType: hyperdbgsdk.HiddenHookExecCc, // !epthook = EPT exec CC (safe for CET shadow stack)
		Tag:       tag,
		Options: hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: hookAddress, // target address
		},
	}

	eventBuf := byteslice.FromStruct(&event)
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf := byteslice.FromStruct(&result)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("EptHook: REGISTER_EVENT IOCTL failed: %w", err)
	}
	// Re-read result from the output buffer.
	result = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](resultBuf)
	if !result.IsSuccessful {
		return 0, fmt.Errorf("EptHook: event registration failed, error=%d", result.Error)
	}

	// 3. Add action with the Go AST script (IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT).
	// The action struct is followed immediately by the script bytes in the
	// same buffer.
	action := hyperdbgsdk.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          hyperdbgsdk.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes := byteslice.FromStruct(&action)
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf := byteslice.FromStruct(&actionResult)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("EptHook: ADD_ACTION IOCTL failed: %w", err)
	}
	actionResult = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](actionResultBuf)
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
func (d *Debugger) EptHookForProcess(hookAddress uint64, pid uint32, callbackSrc string) (uint64, error) {
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

	event := hyperdbgsdk.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: pid,
		IsEnabled: true,
		EventType: hyperdbgsdk.HiddenHookExecCc,
		Tag:       tag,
		Options: hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: hookAddress,
		},
	}

	eventBuf := byteslice.FromStruct(&event)
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf := byteslice.FromStruct(&result)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("EptHookForProcess: REGISTER_EVENT IOCTL failed: %w", err)
	}
	result = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](resultBuf)
	if !result.IsSuccessful {
		return 0, fmt.Errorf("EptHookForProcess: event registration failed, error=%d", result.Error)
	}

	action := hyperdbgsdk.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          hyperdbgsdk.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes := byteslice.FromStruct(&action)
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf := byteslice.FromStruct(&actionResult)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("EptHookForProcess: ADD_ACTION IOCTL failed: %w", err)
	}
	actionResult = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](actionResultBuf)
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
func (d *Debugger) MonitorReadForProcess(addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
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

	event := hyperdbgsdk.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: pid,
		IsEnabled: true,
		EventType: hyperdbgsdk.HiddenHookRead,
		Tag:       tag,
		Options: hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
			OptionalParam1: addrStart,
			OptionalParam2: addrEnd,
		},
	}

	eventBuf := byteslice.FromStruct(&event)
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf := byteslice.FromStruct(&result)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("MonitorReadForProcess: REGISTER_EVENT IOCTL failed: %w", err)
	}
	result = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](resultBuf)
	if !result.IsSuccessful {
		return 0, fmt.Errorf("MonitorReadForProcess: event registration failed, error=%d", result.Error)
	}

	action := hyperdbgsdk.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          hyperdbgsdk.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: 0x02, // HOOK_FLAG_GO_AST
	}
	actionSize := unsafe.Sizeof(action)
	totalSize := uint32(actionSize) + uint32(len(astBytes))
	buf := make([]byte, totalSize)
	actionBytes := byteslice.FromStruct(&action)
	copy(buf, actionBytes)
	copy(buf[actionSize:], astBytes)

	var actionResult hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf := byteslice.FromStruct(&actionResult)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("MonitorReadForProcess: ADD_ACTION IOCTL failed: %w", err)
	}
	actionResult = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT](actionResultBuf)
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
//	mp, _ := dbg.StartMessagePump()   // spawns the goroutine
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

// StartMessagePump spawns a goroutine that drains kernel packets via the
// IRP-based channel, forwarding each one to OnPacket. It must be called
// AFTER LoadVMM; the caller must invoke (*MessagePump).Stop BEFORE
// UnloadVMM so the dedicated device handle is released cleanly and the
// main handle is still usable for the END_OF_IRPS signal.
func (d *Debugger) StartMessagePump() (*MessagePump, error) {
	d.mu.Lock()
	if d.state < StateVmmLoaded {
		d.mu.Unlock()
		return nil, fmt.Errorf("StartMessagePump: VMM not loaded")
	}
	if d.pauseEvent == nil {
		d.pauseEvent = make(chan struct{}, 1)
	}
	mainDev := d.device
	d.mu.Unlock()

	// Open a dedicated device handle so the pending IRP does not block the
	// main IOCTL handle used by Continue/Pause/EptHook/… (mirrors the C++
	// ReadIrpBasedBuffer opening a second \\.\HyperDbgDebuggerDevice).
	dev, err := comm.Open(comm.DeviceName)
	if err != nil {
		return nil, fmt.Errorf("StartMessagePump: open dedicated device: %w", err)
	}

	mp := &MessagePump{
		mainDev: mainDev,
		dev:     dev,
		stop:    make(chan struct{}),
		done:    make(chan struct{}),
	}
	go mp.run(d)
	return mp, nil
}

// run is the IRP-reading loop. It mirrors ReadIrpBasedBuffer in
// go-libhyperdbg/app/packets.go but writes packets straight to the log file
// instead of routing them through the Messaging dispatcher.
func (mp *MessagePump) run(d *Debugger) {
	defer close(mp.done)
	defer mp.dev.Close()

	// REGISTER_NOTIFY_BUFFER{Type: IRP_BASED, HEvent: NULL}.
	var reg hyperdbgsdk.REGISTER_NOTIFY_BUFFER
	reg.Type = hyperdbgsdk.IrpBased
	regBuf := (*[unsafe.Sizeof(reg)]byte)(unsafe.Pointer(&reg))[:]

	// UsermodeBufferSize = sizeof(UINT32) + PacketChunkSize + 1 = 4101.
	// See hyperdbg/include/SDK/headers/Constants.h.
	out := make([]byte, 4+4096+1)

	// operationHypervisorDriverEndOfIrps — the kernel completes the pending
	// IRP with this code when Stop signals it via
	// IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL (see Ioctl.c:337).
	// operationNotificationFromUserDebuggerPause — the kernel sends this
	// when the user-mode debugger pauses the debuggee (Step/Pause/bp hit).
	// The body is a DEBUGGEE_UD_PAUSED_PACKET. Value mirrors
	// OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE in
	// hyperdbg/include/SDK/headers/Constants.h (16 | mandatory bit).

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
		n, err := mp.dev.Ioctl(comm.IOCTL_CODE_REGISTER_EVENT, regBuf, out)
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

		// PAUSED packets are handled internally first (they update the
		// paused register/thread state and signal pauseEvent + OnPaused
		// so Step/TraceInto can complete). The packet is then forwarded
		// to OnPacket like every other packet so consumers can react
		// (e.g. refresh the UI).
		// 严格按 opCode 判断：之前用 "Rip 在用户空间范围" 启发式判断
		// 会被 LogInfo 等文本消息误判（前 8 字节恰好是用户态地址），
		// 导致 pauseEvent 被错误信号、Step 在错误时刻返回 → 单步超时。
		if opCode == operationNotificationFromUserDebuggerPause {
			pausedSize := unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_UD_PAUSED_PACKET{})
			if uint32(len(msg)) >= uint32(pausedSize) {
				paused := (*hyperdbgsdk.DEBUGGEE_UD_PAUSED_PACKET)(unsafe.Pointer(&msg[0]))
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
				if cb != nil {
					cb()
				}
			}
			// 转发 PAUSED 包给 OnPacket 消费者（即使 PAUSED 也转发，
			// 让 UI 等消费者知道暂停发生）。PAUSED 包是二进制结构，
			// 消费者应按 opCode 区分，不要当文本写日志。
			if cb := d.OnPacket; cb != nil {
				cb(opCode, msg)
			}
			continue
		}

		// 其他包统一转发给 OnPacket 回调（文本日志、命令输出等）。
		// 消费者自行处理 NUL 终止和换行。
		if cb := d.OnPacket; cb != nil {
			cb(opCode, msg)
		}
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
		_, _ = mp.mainDev.Ioctl(
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
func (d *Debugger) ReadMemory(addr uint64, pid uint32, size uint32) ([]byte, hyperdbgsdk.DEBUGGER_READ_MEMORY_ADDRESS_MODE, error) {
	d.mu.Lock()
	dev := d.device
	d.mu.Unlock()
	if dev == nil {
		return nil, 0, fmt.Errorf("ReadMemory: not connected")
	}
	return readmem.ReadMemory(dev, addr, pid, size,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
}

// Continue resumes execution of the debugged process by sending the
// CONTINUE_PROCESS IOCTL with the stored process token. The debuggee runs
// until a pause is requested (Pause) or a registered event fires.
//
// Mirrors the C libhyperdbg 'g' command path: UdContinueProcess(token).
func (d *Debugger) Continue() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken == 0 {
		return fmt.Errorf("Continue: no process attached (call StartProcess first)")
	}
	if err := continueProcess(d.device, d.processToken); err != nil {
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
func (d *Debugger) Pause() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken == 0 {
		return fmt.Errorf("Pause: no process attached (call StartProcess first)")
	}
	if err := pauseProcess(d.device, d.processToken); err != nil {
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
func (d *Debugger) StartProcess(exePath string) (Process, error) {
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
	token, err := attachProcess(d.device, proc.Pid, proc.Tid, true)
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
		_ = detachProcess(d.device, proc.Pid)
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
