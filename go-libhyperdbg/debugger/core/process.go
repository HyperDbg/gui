// Process launching for the debuggee.
//
// Mirrors libhyperdbg/code/debugger/user-level/ud.cpp:UdCreateSuspendedProcess
// for the CREATE_SUSPENDED + attach flow that lets the VMM intercept the
// child before its first instruction. Pure syscalls via
// golang.org/x/sys/windows (no cgo).
package core

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// DEBUG_ONLY_THIS_PROCESS is the CreateProcess flag that tells the VMM to
// debug only the immediate child (not its descendants). Value matches the
// Win32 SDK constant.
const DEBUG_ONLY_THIS_PROCESS = 0x00000002

// CREATE_NEW_CONSOLE gives the debuggee its own console so its stdout does
// not pollute the debugger's output.
const CREATE_NEW_CONSOLE = 0x00000010

// CREATE_SUSPENDED creates the process in a suspended state. The VMM attach
// IOCTL is issued against this suspended process; only after the attach
// succeeds does the user call Continue (g) to resume the main thread.
const CREATE_SUSPENDED = 0x00000004

// createSuspendedProcess launches exePath in a suspended state using
// CreateProcessW with CREATE_SUSPENDED | CREATE_NEW_CONSOLE. The HyperDbg
// VMM (already loaded) does not intercept the child yet — that happens
// only after attachProcess sends the ATTACH IOCTL with this pid/tid.
//
// The caller owns both the process and thread handles and must CloseHandle
// them (typically via Process.Close after the thread is resumed and the
// attach IOCTL completes).
//
// Mirrors UdCreateSuspendedProcess in ud.cpp:355-367.
func createSuspendedProcess(exePath string) (Process, error) {
	cmdLine, err := windows.UTF16PtrFromString(exePath)
	if err != nil {
		return Process{}, fmt.Errorf("createSuspendedProcess: bad path: %w", err)
	}

	var si windows.StartupInfo
	si.Cb = uint32(unsafe.Sizeof(si))
	var pi windows.ProcessInformation

	// CREATE_SUSPENDED | CREATE_NEW_CONSOLE mirrors the C libhyperdbg
	// combination. CREATE_NEW_CONSOLE gives the debuggee its own console so
	// its stdout/stderr do not pollute the debugger's output; CREATE_SUSPENDED
	// pauses the main thread at the first instruction so the VMM attach IOCTL
	// can register the debug session before any user code runs.
	//
	// Note: DEBUG_ONLY_THIS_PROCESS is intentionally NOT set here. The C
	// libhyperdbg's UdCreateSuspendedProcess uses CREATE_SUSPENDED |
	// CREATE_NEW_CONSOLE only; the debug-port interception is established by
	// the subsequent ATTACH IOCTL, not by CreateProcess flags.
	flags := uint32(CREATE_SUSPENDED | CREATE_NEW_CONSOLE)
	if err := windows.CreateProcess(
		nil,      // application name (NULL = parse from cmd line)
		cmdLine,  // command line (mutable in-place)
		nil, nil, // process/thread attrs (not inherited)
		false, // inherit handles
		flags, // creation flags
		nil,   // environment (NULL = inherit)
		nil,   // current directory (NULL = inherit)
		&si, &pi); err != nil {
		return Process{}, fmt.Errorf("createSuspendedProcess: CreateProcess failed: %w", err)
	}

	// Keep both handles: the process handle is returned in Process.Handle,
	// and the thread handle is returned in Process.ThreadHandle so the
	// caller can ResumeThread on it after the attach IOCTL succeeds.
	return Process{
		Handle:       uintptr(pi.Process),
		ThreadHandle: uintptr(pi.Thread),
		Pid:          pi.ProcessId,
		Tid:          pi.ThreadId,
	}, nil
}

// Close releases the process and thread handles. Safe to call multiple
// times.
func (p *Process) Close() error {
	if p.Handle != 0 {
		h := windows.Handle(p.Handle)
		p.Handle = 0
		_ = windows.CloseHandle(h)
	}
	if p.ThreadHandle != 0 {
		h := windows.Handle(p.ThreadHandle)
		p.ThreadHandle = 0
		_ = windows.CloseHandle(h)
	}
	return nil
}

// Terminate force-kills the debuggee process via TerminateProcess(exit=1).
// It is a best-effort cleanup step used on debugger shutdown: after
// UnloadVMM has detached the kernel debug session, the debuggee may still
// be alive and must be killed so it does not outlive the driver (a running
// debuggee whose VMM hooks were just torn down can fault). Safe to call on
// an already-closed Process (returns nil).
func (p *Process) Terminate() error {
	return windowsTerminateProcess(p.Handle)
}

// windowsResumeThread is a thin wrapper around windows.ResumeThread so the
// caller can use a uintptr thread handle (the type stored in Process) without
// repeating the cast. Returns an error if the resume fails.
//
// The returned previous suspend count is discarded; callers only need to
// know whether the resume succeeded.
func windowsResumeThread(threadHandle uintptr) error {
	if threadHandle == 0 {
		return fmt.Errorf("ResumeThread: handle is 0")
	}
	if _, err := windows.ResumeThread(windows.Handle(threadHandle)); err != nil {
		return fmt.Errorf("ResumeThread failed: %w", err)
	}
	return nil
}

// windowsTerminateProcess is a thin wrapper around windows.TerminateProcess
// used for best-effort cleanup when an attach or resume fails after
// CreateProcess succeeded. The exit code is forced to 1.
func windowsTerminateProcess(processHandle uintptr) error {
	if processHandle == 0 {
		return nil
	}
	return windows.TerminateProcess(windows.Handle(processHandle), 1)
}
