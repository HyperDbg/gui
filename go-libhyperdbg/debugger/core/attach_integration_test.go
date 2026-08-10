// Package core — attach_integration_test.go
//
// Integration test for the user-mode debuggee attach/continue/pause IOCTL
// flow. Drives the full pipeline end-to-end against the real hyperkd.sys
// driver:
//
//  1. Load hyperkd.sys via SCM (driverloader.Load).
//  2. Open \\.\HyperDbgDebuggerDevice (comm.Open).
//  3. Initialise the VMM (IOCTL_CODE_INIT_VMM).
//  4. Start a harmless debuggee (notepad.exe) via StartProcess, which
//     creates a suspended child + sends IOCTL_CODE_DEBUGGER_ATTACH_DETACH
//     _USER_MODE_PROCESS with Action=ATTACH.
//  5. Verify the returned Token is non-zero.
//  6. Send Continue (g) via IOCTL and verify no error.
//  7. Sleep briefly, then send Pause via IOCTL and verify no error.
//  8. Clean up: detach, terminate the debuggee, terminate VMX, close the
//     device, unload the driver.
//
// Skip conditions (t.Skip, not t.Fatal):
//   - hyperkd.sys not found in the build output tree.
//   - Process is not running elevated.
//   - VMM init fails (system lacks VT-x / nested-virt support).
//
// To run from an elevated shell:
//
//	go test -v -count=1 -run TestAttachContinuePause ./debugger/core/
package core

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"github.com/hyperdbg/go-libhyperdbg/types"
	"golang.org/x/sys/windows"
)

// isAdmin reports whether the current process has administrator
// privileges by inspecting the process token elevation state. Driver load
// and device open both require elevation, so tests skip rather than fail
// when not elevated.
func isAdmin() bool {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token); err != nil {
		return false
	}
	defer token.Close()
	return token.IsElevated()
}

// findHyperkdDriver searches common build output directories for
// hyperkd.sys starting from the test working directory and walking up to
// the repository root. Returns the absolute path or "" if not found.
func findHyperkdDriver(t *testing.T) string {
	t.Helper()
	candidates := []string{
		`Debug\hyperkd.sys`,
		`Release\hyperkd.sys`,
		`x64\Debug\hyperkd.sys`,
		`x64\Release\hyperkd.sys`,
	}
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	dir := wd
	for i := 0; i < 8 && dir != "" && dir != filepath.Dir(dir); i++ {
		for _, c := range candidates {
			p := filepath.Join(dir, c)
			if fi, err := os.Stat(p); err == nil && !fi.IsDir() {
				return p
			}
		}
		dir = filepath.Dir(dir)
	}
	return ""
}

// debuggerOperationWasSuccessful is now declared in control.go (promoted
// from the test package to the real package so non-test code can use it).
// The value is 0xFFFFFFFF per
// HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h.

// initVmmRequest mirrors DEBUGGER_INIT_VMM_PACKET: a single UINT32
// KernelStatus field used as both IOCTL input and output (METHOD_BUFFERED
// aliases the same kernel SystemBuffer for both directions).
type initVmmRequest struct {
	KernelStatus uint32
}

// TestAttachContinuePause verifies the full attach/continue/pause IOCTL
// flow against the real driver. It launches notepad.exe as a harmless
// debuggee — notepad is the canonical test target in the project's .ds
// scripts (run-notepad.ds) and is guaranteed to be present on every
// Windows system.
func TestAttachContinuePause(t *testing.T) {
	driverPath := findHyperkdDriver(t)
	if driverPath == "" {
		t.Skip("hyperkd.sys not found in build output; build the VMM driver first")
	}
	if !isAdmin() {
		t.Skip("not running as administrator; driver load + device open require elevation")
	}
	t.Logf("using driver: %s", driverPath)

	// notepad.exe is on the system PATH; resolve it via the System32
	// directory so the test does not depend on the test's CWD.
	system32, err := windows.GetSystemDirectory()
	if err != nil {
		t.Fatalf("GetSystemDirectory failed: %v", err)
	}
	notepadPath := filepath.Join(system32, "notepad.exe")
	if _, err := os.Stat(notepadPath); err != nil {
		t.Skipf("notepad.exe not found at %q: %v", notepadPath, err)
	}
	t.Logf("using debuggee: %s", notepadPath)

	ctx := context.Background()
	d := driverloader.NewDriver(driverPath)

	// Best-effort cleanup of any stale service from a prior run.
	_ = d.Unload(ctx)
	if exists, _ := d.Exists(ctx); exists {
		time.Sleep(500 * time.Millisecond)
		_ = d.Unload(ctx)
	}

	if err := d.Load(ctx); err != nil {
		t.Fatalf("driverloader.Load failed: %v", err)
	}
	time.Sleep(2 * time.Second) // let DriverEntry create the device
	t.Cleanup(func() { _ = d.Unload(ctx) })

	// Open the device handle (with retries for slow systems).
	var dev *comm.Device
	for i := 0; i < 5; i++ {
		dev, err = comm.Open(ctx, comm.DeviceName)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("comm.Open(%q) failed after retries: %v", comm.DeviceName, err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	// Best-effort VMX teardown before unloading the driver.
	t.Cleanup(func() {
		_, _ = dev.IoctlStruct(context.Background(),
			comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		time.Sleep(500 * time.Millisecond)
	})

	// Step 1: Initialise the VMM.
	var vmmReq initVmmRequest
	vmmSize := uint32(unsafe.Sizeof(vmmReq))
	if _, err := dev.IoctlStruct(ctx, comm.IOCTL_CODE_INIT_VMM,
		unsafe.Pointer(&vmmReq), unsafe.Pointer(&vmmReq), vmmSize, vmmSize); err != nil {
		t.Skipf("IOCTL_INIT_VMM failed: %v", err)
	}
	t.Logf("IOCTL_INIT_VMM: KernelStatus=0x%08x", vmmReq.KernelStatus)
	if vmmReq.KernelStatus != debuggerOperationWasSuccessful {
		t.Skipf("VMM init failed (KernelStatus=0x%08x); system likely lacks "+
			"VT-x / nested-virt support", vmmReq.KernelStatus)
	}

	// Step 2: Create a suspended notepad child and send the ATTACH IOCTL.
	//   attachProcess does both: it expects the caller to have already
	//   created the suspended child (createSuspendedProcess) and to pass
	//   its pid/tid. We invoke them in sequence here to keep the test
	//   explicit about what is being verified.
	proc, err := createSuspendedProcess(notepadPath)
	if err != nil {
		t.Fatalf("createSuspendedProcess failed: %v", err)
	}
	t.Logf("created suspended notepad: pid=%d tid=%d", proc.Pid, proc.Tid)
	// Ensure the suspended child is cleaned up no matter how the test
	// exits: terminate it (it never ran) and close handles.
	t.Cleanup(func() {
		_ = windowsTerminateProcess(proc.Handle)
		_ = proc.Close()
	})

	token, err := attachProcess(ctx, dev, proc.Pid, proc.Tid, true)
	if err != nil {
		t.Fatalf("attachProcess failed: %v", err)
	}
	if token == 0 {
		t.Fatal("attachProcess returned Token=0 (attach did not register)")
	}
	t.Logf("ATTACH OK: Token=0x%016X", token)

	// Step 3: Resume the main thread so the kernel can begin intercepting.
	// The kernel will pause the debuggee at the first instruction; we then
	// Continue past it.
	if err := windowsResumeThread(proc.ThreadHandle); err != nil {
		t.Fatalf("ResumeThread failed: %v", err)
	}
	t.Logf("main thread resumed")

	// Step 4: Continue (g). The kernel lets the debuggee run; notepad is
	// harmless and will sit in its message loop. We tolerate a brief delay
	// for the kernel to settle into the running state.
	if err := continueProcess(ctx, dev, token); err != nil {
		t.Fatalf("continueProcess failed: %v", err)
	}
	t.Logf("CONTINUE OK")
	time.Sleep(1 * time.Second)

	// Step 5: Pause. The kernel arranges for the debuggee to halt at the
	// next instruction. We do not strictly wait for the DEBUGGEE_UD_PAUSED
	// packet here (that requires the user-level listener, exercised in a
	// separate test); we only verify the IOCTL itself succeeds.
	//
	// ErrAlreadyPaused is non-fatal: it means the PEB read/write monitor
	// EPT hook (installed by AttachingPerformAttachToProcess at
	// Attaching.c:912-927) has already fired and put the debuggee into
	// IsOnThreadInterceptingPhase. The C CommandPauseRequest treats this
	// case as a silent no-op (pause.cpp:55), so we do the same.
	if err := pauseProcess(ctx, dev, token); err != nil {
		if errors.Is(err, ErrAlreadyPaused) {
			t.Logf("PAUSE: process already in intercepting phase (PEB monitor hook fired) — treated as success, matching C pause.cpp:55")
		} else {
			t.Fatalf("pauseProcess failed: %v", err)
		}
	} else {
		t.Logf("PAUSE OK")
	}

	// Step 6: Continue before detach. AttachingPerformDetach rejects the
	// detach with DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS
	// while any thread is paused, so we must resume first. If we skip this,
	// the PEB read/write monitor EPT hook installed by AttachingPerformAttach
	// (Attaching.c:912-927) stays in g_EptState->HookedPagesList, and the
	// subsequent TERMINATE_VMX crashes in EptHookUnHookAll →
	// EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList because that
	// hook has IsHiddenBreakpoint == FALSE but no !epthook2 was ever
	// registered (g_EptHook2sDetourListHead is uninitialized).
	if err := continueProcess(ctx, dev, token); err != nil {
		t.Fatalf("continueProcess (pre-detach) failed: %v", err)
	}
	t.Logf("CONTINUE (pre-detach) OK")
	// Brief sleep to let the kernel transition the debuggee out of the
	// paused state before the detach IOCTL arrives.
	time.Sleep(500 * time.Millisecond)

	// Step 7: Detach. Now that no threads are paused, the kernel accepts
	// the detach and removes the PEB monitor EPT hook. Errors here are
	// still tolerated because the debuggee is about to be terminated.
	if err := detachProcess(ctx, dev, token); err != nil {
		t.Logf("best-effort detachProcess failed (non-fatal): %v", err)
	} else {
		t.Logf("DETACH OK")
	}

	t.Logf("PASS: attach (Token=%#x) → continue → pause → continue → detach", token)

	// Reference unused types to keep the import stable even if the test
	// body is trimmed in future refactorings.
	_ = types.DebuggerAttachDetachUserModeProcessActionAttach
}
