// Package core — attach.go
//
// Implements the user-mode debuggee attach/continue/pause IOCTLs that drive
// the DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS request structure through
// IOCTL_CODE_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS.
//
// Mirrors the C libhyperdbg functions:
//   - UdAttachToProcess     (libhyperdbg/code/debugger/user-level/ud.cpp:380)
//   - UdContinueProcess     (libhyperdbg/code/debugger/user-level/ud.cpp:941)
//   - UdPauseProcess        (libhyperdbg/code/debugger/user-level/ud.cpp:873)
//
// All three IOCTLs reuse the same DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
// packet: the caller fills the Action field (and ProcessId/ThreadId for
// attach), the kernel reads it, performs the action, and writes the result
// (Token on attach, Result on continue/pause) back into the same buffer
// (METHOD_BUFFERED aliases input and output).
package core

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
)

// attachDetachRequestSize is the C ABI size of
// DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS. The kernel expects both the
// input and output buffer to be exactly this size; a smaller buffer is
// rejected with STATUS_BUFFER_TOO_SMALL.
const attachDetachRequestSize = unsafe.Sizeof(hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{})

// attachProcess sends the attach IOCTL to the kernel. It is the low-level
// helper under StartProcess: the caller has already created the debuggee
// (suspended, via createSuspendedProcess) and supplies its pid/tid plus
// whether the kernel should pause at the very first instruction.
//
// On success the kernel returns the process debugging Token in the same
// packet (Result is also set to DEBUGGER_OPERATION_WAS_SUCCESSFUL). The
// Token is required for all subsequent Continue/Pause/Command IOCTLs.
//
// Mirrors UdAttachToProcess in ud.cpp:380-480.
func attachProcess(dev *comm.Device, pid uint32, tid uint32, checkCallbackAtFirstInstruction bool) (uint64, error) {
	if dev == nil {
		return 0, fmt.Errorf("attachProcess: nil device (not connected?)")
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		IsStartingNewProcess:            true,
		ProcessId:                       pid,
		ThreadId:                        tid,
		CheckCallbackAtFirstInstruction: checkCallbackAtFirstInstruction,
		Action:                          hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionAttach,
	}
	// METHOD_BUFFERED: input and output alias the same kernel SystemBuffer,
	// so we pass the same slice for both and the driver writes the response
	// back into it.
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := dev.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess,
		buf, buf); err != nil {
		return 0, fmt.Errorf("attachProcess: ATTACH IOCTL failed: %w", err)
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return 0, fmt.Errorf("attachProcess: kernel rejected attach (Result=0x%016X, see DEBUGGER_ERROR_* in SDK)", pkt.Result)
	}
	if pkt.Token == 0 {
		return 0, fmt.Errorf("attachProcess: kernel returned Token=0 (attach did not register a debugging session)")
	}
	return pkt.Token, nil
}

// continueProcess sends the continue IOCTL for the given debugging token.
// The debuggee resumes execution and runs until a pause is requested or a
// registered event (hook/breakpoint) fires.
//
// Mirrors UdContinueProcess in ud.cpp:941-1000.
func continueProcess(dev *comm.Device, token uint64) error {
	if dev == nil {
		return fmt.Errorf("continueProcess: nil device")
	}
	if token == 0 {
		return fmt.Errorf("continueProcess: token is 0 (no process attached?)")
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		Action: hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionContinueProcess,
		Token:  token,
	}
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := dev.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess,
		buf, buf); err != nil {
		return fmt.Errorf("continueProcess: CONTINUE IOCTL failed: %w", err)
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("continueProcess: kernel rejected continue (Result=0x%016X)", pkt.Result)
	}
	return nil
}

// ErrAlreadyPaused is returned by PauseProcess when the kernel reports
// DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS (0xC0000031). This is
// not a real error: it means the debuggee is already in the intercepting
// phase (typically because the PEB read/write monitor EPT hook installed
// by AttachingPerformAttachToProcess fired and paused the process). The
// C libhyperdbg's CommandPauseRequest silently ignores this case
// (pause.cpp:55 — the if condition simply evaluates to false and the
// "please keep interacting" message is skipped), so Go callers should
// treat errors.Is(err, ErrAlreadyPaused) as a success condition.
var ErrAlreadyPaused = errors.New("process already in intercepting (paused) phase")

// pauseProcess sends the pause IOCTL for the given debugging token. The
// kernel arranges for the debuggee to halt at the next instruction and emit
// a DEBUGGEE_UD_PAUSED_PACKET (which the user-level listener dispatches).
//
// Returns ErrAlreadyPaused (non-fatal) if the debuggee is already paused —
// see the doc comment on ErrAlreadyPaused for why this happens and why the
// C libhyperdbg treats it as success.
//
// Mirrors UdPauseProcess in ud.cpp:873-932.
func pauseProcess(dev *comm.Device, token uint64) error {
	if dev == nil {
		return fmt.Errorf("pauseProcess: nil device")
	}
	if token == 0 {
		return fmt.Errorf("pauseProcess: token is 0 (no process attached?)")
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		Action: hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionPauseProcess,
		Token:  token,
	}
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := dev.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess,
		buf, buf); err != nil {
		return fmt.Errorf("pauseProcess: PAUSE IOCTL failed: %w", err)
	}
	// DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS = 0xC0000031 per
	// HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h:338. The kernel sets
	// this when IsOnThreadInterceptingPhase is already TRUE — i.e. the
	// debuggee is already paused (typically by the PEB monitor hook). The
	// C CommandPauseRequest treats this as a no-op, so we surface it as a
	// sentinel error that callers can handle with errors.Is.
	const debuggerErrorUnableToPauseTheProcessThreads uint64 = 0xC0000031
	if pkt.Result == debuggerErrorUnableToPauseTheProcessThreads {
		return ErrAlreadyPaused
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("pauseProcess: kernel rejected pause (Result=0x%016X)", pkt.Result)
	}
	return nil
}

// detachProcess sends the detach IOCTL for the given process. The kernel
// looks up the debug session by ProcessId (not Token) — see
// AttachingPerformDetach in Attaching.c:1224. The caller must have already
// continued the process (C++ UdDetachProcess calls UdContinueProcess first).
//
// Mirrors UdDetachProcess in ud.cpp:795-864.
func detachProcess(dev *comm.Device, pid uint32) error {
	if dev == nil {
		return fmt.Errorf("detachProcess: nil device")
	}
	if pid == 0 {
		return nil
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		Action:    hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionDetach,
		ProcessId: pid,
	}
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := dev.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess,
		buf, buf); err != nil {
		return fmt.Errorf("detachProcess: DETACH IOCTL failed: %w", err)
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("detachProcess: kernel rejected detach (Result=0x%016X)", pkt.Result)
	}
	return nil
}
