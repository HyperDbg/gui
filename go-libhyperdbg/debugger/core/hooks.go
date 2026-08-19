// Package core — hooks.go
//
// Real implementations of the HyperDbg "!" extension hook commands.
//
// All hook commands share the same IOCTL flow (mirrors the C++ pattern in
// libhyperdbg/code/debugger/commands/extension-commands/*.cpp):
//
//  1. Build DEBUGGER_GENERAL_EVENT_DETAIL with the command-specific
//     EventType (e.g. SYSCALL_HOOK_EFER_SYSRET) and Options.OptionalParamN.
//  2. Send via IOCTL_CODE_DEBUGGER_REGISTER_EVENT (SendEventToKernel in C++).
//  3. Build DEBUGGER_GENERAL_ACTION with ActionType=RunScript and the
//     compiled Go AST bytes appended; send via
//     IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT (RegisterActionToEvent in C++).
//
// This file factorises that flow into registerHookEvent and exposes one
// public method per command. The api layer (api/commands_extension.go)
// delegates to these methods.
//
// References:
//   - C++ pattern: cpuid.cpp / syscall-sysret.cpp / crwrite.cpp / ...
//   - Existing Go reference: EptHook in debugger.go
package core

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	astencoder "github.com/hyperdbg/go-bridge/ast"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
)

// Common constants for event registration (mirror Constants.h).
const (
	// debuggerEventApplyToAllProcesses is DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES
	// (0xFFFFFFFF) — used as ProcessId to apply an event to every process.
	debuggerEventApplyToAllProcesses uint32 = 0xFFFFFFFF

	// debuggerEventSyscallAllSysretsOrSyscalls is
	// DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS (0xFFFFFFFF) — used as
	// OptionalParam1 for !syscall/!sysret to mean "all syscalls/sysrets".
	debuggerEventSyscallAllSysretsOrSyscalls uint64 = 0xFFFFFFFF

	// hookFlagGoAst is the ScriptBufferPointer flag that tells the kernel
	// the script buffer is a Go AST wire-format payload (go-bridge/ast),
	// not a HyperDbg script-engine string or raw custom code.
	hookFlagGoAst uint32 = 0x02
)

// registerHookEvent is the shared backend for every "!xxx hook" command.
// It compiles callbackSrc to a Go AST, registers an event of the given
// type with the given per-process id and options, then attaches the AST
// as a RunScript action. Returns the event tag on success.
//
// Mirrors the C++ InterpretGeneralEventAndActionsFields +
// SendEventToKernel + RegisterActionToEvent sequence.
//
// pid == debuggerEventApplyToAllProcesses means "all processes"; a real
// pid is required for user-mode hooks (EptHook2/Monitor*/ModeHook) so the
// kernel can resolve the virtual address through that process's CR3.
func (d *Debugger) registerHookEvent(
	ctx context.Context,
	eventType hyperdbgsdk.VMM_EVENT_TYPE_ENUM,
	pid uint32,
	options hyperdbgsdk.DEBUGGER_EVENT_OPTIONS,
	callbackSrc string,
) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("registerHookEvent: VMM not loaded")
	}

	// 1. Compile the Go callback source to binary AST.
	astBytes, err := astencoder.Encode(callbackSrc)
	if err != nil {
		return 0, fmt.Errorf("registerHookEvent: callback compile failed: %w", err)
	}

	// 2. Register the event (IOCTL_CODE_DEBUGGER_REGISTER_EVENT).
	tag := d.nextTag
	d.nextTag++

	event := hyperdbgsdk.DEBUGGER_GENERAL_EVENT_DETAIL{
		ProcessId: pid,
		IsEnabled: true,
		EventType: eventType,
		Tag:       tag,
		Options:   options,
	}
	eventBuf, err := structToBytes(&event)
	if err != nil {
		return 0, err
	}
	var result hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	resultBuf, err := structToBytes(&result)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_REGISTER_EVENT,
		eventBuf, resultBuf); err != nil {
		return 0, fmt.Errorf("registerHookEvent: REGISTER_EVENT IOCTL failed: %w", err)
	}
	if err := bytesToStruct(resultBuf, &result); err != nil {
		return 0, err
	}
	if !result.IsSuccessful {
		return 0, fmt.Errorf("registerHookEvent: event registration failed, kernel error=0x%08X", result.Error)
	}

	// 3. Attach the Go AST script as a RunScript action
	//    (IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT). The action struct is
	//    followed immediately by the AST bytes in the same buffer.
	action := hyperdbgsdk.DEBUGGER_GENERAL_ACTION{
		EventTag:            tag,
		ActionType:          hyperdbgsdk.RunScript,
		ScriptBufferSize:    uint32(len(astBytes)),
		ScriptBufferPointer: hookFlagGoAst,
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

	var actionResult hyperdbgsdk.DEBUGGER_EVENT_AND_ACTION_RESULT
	actionResultBuf, err := structToBytes(&actionResult)
	if err != nil {
		return 0, err
	}
	if _, err := d.device.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT,
		buf, actionResultBuf); err != nil {
		return 0, fmt.Errorf("registerHookEvent: ADD_ACTION IOCTL failed: %w", err)
	}
	if err := bytesToStruct(actionResultBuf, &actionResult); err != nil {
		return 0, err
	}
	if !actionResult.IsSuccessful {
		return 0, fmt.Errorf("registerHookEvent: action registration failed, kernel error=0x%08X", actionResult.Error)
	}
	return tag, nil
}

// ----------------------------------------------------------------
// Hook commands — each mirrors a C++ !xxx command.
// ----------------------------------------------------------------

// CpuidHook registers a CPUID-instruction execution hook.
//
// If hasEax is true, only CPUIDs with the given EAX index trigger the
// callback; otherwise every CPUID triggers it.
//
// C++: cpuid.cpp — EventType=CPUID_INSTRUCTION_EXECUTION,
// OptionalParam1=(bool)HasEax, OptionalParam2=EaxValue.
func (d *Debugger) CpuidHook(ctx context.Context, hasEax bool, eax uint64, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: 0}
	if hasEax {
		opts.OptionalParam1 = 1
		opts.OptionalParam2 = eax
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.CpuidInstructionExecution,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// CrwriteHook registers a control-register write hook.
//
// cr must be 0 (CR0) or 4 (CR4) — the only registers the VMX exit
// "control-register access" can qualify. mask filters which bit changes
// trigger the event (0 = any write).
//
// C++: crwrite.cpp — EventType=CONTROL_REGISTER_MODIFIED,
// OptionalParam1=TargetRegister, OptionalParam2=MaskRegister.
func (d *Debugger) CrwriteHook(ctx context.Context, cr uint32, mask uint64, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: uint64(cr),
		OptionalParam2: mask,
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.ControlRegisterModified,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// DrHook registers a debug-register access hook (MOV to/from DRn).
//
// C++: dr.cpp — EventType=DEBUG_REGISTERS_ACCESSED, no optional params.
func (d *Debugger) DrHook(ctx context.Context, callbackSrc string) (uint64, error) {
	return d.registerHookEvent(ctx, hyperdbgsdk.DebugRegistersAccessed,
		debuggerEventApplyToAllProcesses, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{}, callbackSrc)
}

// EptHook2 registers an EPT execution hook in detours mode (!epthook2).
//
// Unlike EptHook (which uses an INT3/cc hidden hook), epthook2 places a
// detour so no #VMEXIT is taken on hit — faster but with extra constraints
// (see hyperdbg.com/docs). pid must be a real process id for user-mode
// targets so the kernel can validate the address.
//
// C++: epthook2.cpp — EventType=HIDDEN_HOOK_EXEC_DETOURS,
// OptionalParam1=hookAddress.
func (d *Debugger) EptHook2(ctx context.Context, hookAddress uint64, pid uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: hookAddress}
	return d.registerHookEvent(ctx, hyperdbgsdk.HiddenHookExecDetours,
		pid, opts, callbackSrc)
}

// ExceptionHook registers an exception hook for the given vector (0-31).
//
// C++: exception.cpp — EventType=EXCEPTION_OCCURRED,
// OptionalParam1=vector.
func (d *Debugger) ExceptionHook(ctx context.Context, vector uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(vector)}
	return d.registerHookEvent(ctx, hyperdbgsdk.ExceptionOccurred,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// InterruptHook registers a hardware-interrupt hook for the given vector
// (32-255). The kernel rejects vector==0 because trapping every interrupt
// makes the system unresponsive.
//
// C++: interrupt.cpp — EventType=EXTERNAL_INTERRUPT_OCCURRED,
// OptionalParam1=vector.
func (d *Debugger) InterruptHook(ctx context.Context, vector uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(vector)}
	return d.registerHookEvent(ctx, hyperdbgsdk.ExternalInterruptOccurred,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// IoInHook registers an IN-instruction hook for the given I/O port.
//
// C++: ioin.cpp — EventType=IN_INSTRUCTION_EXECUTION,
// OptionalParam1=port.
func (d *Debugger) IoInHook(ctx context.Context, port uint16, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(port)}
	return d.registerHookEvent(ctx, hyperdbgsdk.InInstructionExecution,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// IoOutHook registers an OUT-instruction hook for the given I/O port.
//
// C++: ioout.cpp — EventType=OUT_INSTRUCTION_EXECUTION,
// OptionalParam1=port.
func (d *Debugger) IoOutHook(ctx context.Context, port uint16, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(port)}
	return d.registerHookEvent(ctx, hyperdbgsdk.OutInstructionExecution,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// ModeHook registers an execution-mode-change hook (user↔kernel).
//
// mode is one of DEBUGGER_EVENT_MODE_TYPE_*; use
// DebuggerEventModeTypeUserModeAndKernelMode to trap both directions.
// Unlike most hooks, the kernel requires a specific pid here (pid==all
// is rejected by the C++ command path), so callers targeting user mode
// must pass a real pid.
//
// C++: mode.cpp — EventType=TRAP_EXECUTION_MODE_CHANGED,
// OptionalParam1=TargetInterceptionMode.
func (d *Debugger) ModeHook(ctx context.Context, mode hyperdbgsdk.DEBUGGER_EVENT_MODE_TYPE, pid uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(mode)}
	return d.registerHookEvent(ctx, hyperdbgsdk.TrapExecutionModeChanged,
		pid, opts, callbackSrc)
}

// MonitorWrite registers a hidden EPT write-hook on [addrStart, addrEnd]
// for the given process. Every write to any byte in the range fires the
// callback in VMX-root.
//
// C++: monitor.cpp (attribute 'w') — EventType=HIDDEN_HOOK_WRITE,
// OptionalParam1=start, OptionalParam2=end, OptionalParam3=HookMemoryType.
func (d *Debugger) MonitorWrite(ctx context.Context, addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: addrStart,
		OptionalParam2: addrEnd,
		OptionalParam3: 0, // HookMemoryType: 0 = default
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.HiddenHookWrite,
		pid, opts, callbackSrc)
}

// MonitorExec registers a hidden EPT execute-hook on [addrStart, addrEnd]
// for the given process. For a single-instruction hook pass addrEnd=addrStart.
//
// C++: monitor.cpp (attribute 'x') — EventType=HIDDEN_HOOK_EXECUTE,
// OptionalParam1=start, OptionalParam2=end, OptionalParam3=HookMemoryType.
func (d *Debugger) MonitorExec(ctx context.Context, addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: addrStart,
		OptionalParam2: addrEnd,
		OptionalParam3: 0,
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.HiddenHookExecute,
		pid, opts, callbackSrc)
}

// MsrReadHook registers an RDMSR hook. msr==0 means "all MSRs".
//
// C++: msrread.cpp — EventType=RDMSR_INSTRUCTION_EXECUTION,
// OptionalParam1=msr (0 = all).
func (d *Debugger) MsrReadHook(ctx context.Context, msr uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(msr)}
	return d.registerHookEvent(ctx, hyperdbgsdk.RdmsrInstructionExecution,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// MsrWriteHook registers a WRMSR hook. msr==0 means "all MSRs".
//
// C++: msrwrite.cpp — EventType=WRMSR_INSTRUCTION_EXECUTION,
// OptionalParam1=msr (0 = all).
func (d *Debugger) MsrWriteHook(ctx context.Context, msr uint32, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{OptionalParam1: uint64(msr)}
	return d.registerHookEvent(ctx, hyperdbgsdk.WrmsrInstructionExecution,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// SyscallHook registers a SYSCALL-instruction hook. syscallNumber selects
// a specific syscall; pass 0xFFFFFFFF to hook every syscall.
//
// The kernel implements this by setting EFER.SCE and intercepting the
// SYSCALL MSR target (LSTAR). The "safe access memory" mode
// (DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY) is used so the
// callback can safely read user memory from VMX-root.
//
// C++: syscall-sysret.cpp — EventType=SYSCALL_HOOK_EFER_SYSCALL,
// OptionalParam1=syscallNumber, OptionalParam2=SafeAccessMemory.
func (d *Debugger) SyscallHook(ctx context.Context, syscallNumber uint64, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: syscallNumber,
		OptionalParam2: uint64(hyperdbgsdk.DebuggerEventSyscallSysretSafeAccessMemory),
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.SyscallHookEferSyscall,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// SysretHook registers a SYSRET-instruction hook. syscallNumber selects a
// specific returning syscall; pass 0xFFFFFFFF to hook every SYSRET.
//
// The kernel implements this via SYSCALL_HOOK_EFER_SYSRET (shares the
// EFER.SCE configuration with SyscallHook). Useful for catching the
// user-mode return path of a syscall — harder to detect than EptHook on
// the Nt* wrapper because no user-mode instruction is patched.
//
// C++: syscall-sysret.cpp — EventType=SYSCALL_HOOK_EFER_SYSRET,
// OptionalParam1=syscallNumber, OptionalParam2=SafeAccessMemory.
func (d *Debugger) SysretHook(ctx context.Context, syscallNumber uint64, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: syscallNumber,
		OptionalParam2: uint64(hyperdbgsdk.DebuggerEventSyscallSysretSafeAccessMemory),
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.SyscallHookEferSysret,
		debuggerEventApplyToAllProcesses, opts, callbackSrc)
}

// SysretHookForProcess registers a SYSRET-instruction hook scoped to a
// single process. Unlike SysretHook (which sets ProcessId=ALL_PROCESSES),
// this passes the real pid so the kernel skips the event for every other
// process BEFORE invoking the Go callback.
//
// This is critical: the EFER hook still causes a VM exit on every
// syscall/sysret system-wide (hardware-level unavoidable), but only the
// target process's exits run the (expensive) AST interpreter. With
// ALL_PROCESSES, system-wide syscall traffic (5000+/sec) drives 10000+
// VM exits/sec and the interpreter runs for every one of them, overloading
// VMX-root and crashing critical system processes (observed: svchost.exe
// __fastfail → CRITICAL_PROCESS_DIED 0xEF).
//
// Use this instead of SysretHook whenever you only care about one process.
//
// syscallNumber has the same meaning as in SysretHook (pass 0xFFFFFFFF for
// all sysrets; the kernel cannot in general know the syscall number at
// SYSRET time, so this filter is best-effort).
func (d *Debugger) SysretHookForProcess(ctx context.Context, pid uint32, syscallNumber uint64, callbackSrc string) (uint64, error) {
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: syscallNumber,
		OptionalParam2: uint64(hyperdbgsdk.DebuggerEventSyscallSysretSafeAccessMemory),
	}
	return d.registerHookEvent(ctx, hyperdbgsdk.SyscallHookEferSysret,
		pid, opts, callbackSrc)
}

// VmcallHook registers a VMCALL-instruction hook (hypercall).
//
// C++: vmcall.cpp — EventType=VMCALL_INSTRUCTION_EXECUTION, no params.
func (d *Debugger) VmcallHook(ctx context.Context, callbackSrc string) (uint64, error) {
	return d.registerHookEvent(ctx, hyperdbgsdk.VmcallInstructionExecution,
		debuggerEventApplyToAllProcesses, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{}, callbackSrc)
}

// XsetbvHook registers an XSETBV-instruction hook (writes XCR0).
//
// C++: xsetbv.cpp — EventType=XSETBV_INSTRUCTION_EXECUTION, no params.
func (d *Debugger) XsetbvHook(ctx context.Context, callbackSrc string) (uint64, error) {
	return d.registerHookEvent(ctx, hyperdbgsdk.XsetbvInstructionExecution,
		debuggerEventApplyToAllProcesses, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{}, callbackSrc)
}

// TscHook registers an RDTSC/RDTSCP-instruction hook.
//
// C++: tsc.cpp — EventType=TSC_INSTRUCTION_EXECUTION, no params.
func (d *Debugger) TscHook(ctx context.Context, callbackSrc string) (uint64, error) {
	return d.registerHookEvent(ctx, hyperdbgsdk.TscInstructionExecution,
		debuggerEventApplyToAllProcesses, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{}, callbackSrc)
}

// PmcHook registers an RDPMC-instruction hook.
//
// C++: pmc.cpp — EventType=PMC_INSTRUCTION_EXECUTION, no params.
func (d *Debugger) PmcHook(ctx context.Context, callbackSrc string) (uint64, error) {
	return d.registerHookEvent(ctx, hyperdbgsdk.PmcInstructionExecution,
		debuggerEventApplyToAllProcesses, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{}, callbackSrc)
}
