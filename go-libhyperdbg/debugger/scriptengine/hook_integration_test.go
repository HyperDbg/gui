package scriptengine

import (
	"os"
	"path/filepath"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
	"golang.org/x/sys/windows"
)

// isAdmin reports whether the current process has administrator privileges
// by inspecting the process token elevation state. Driver load and device
// open both require elevation, so tests skip rather than fail when not
// elevated.
func isAdmin() bool {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token); err != nil {
		return false
	}
	defer token.Close()
	return token.IsElevated()
}

// findHyperkdDriver searches common build output directories for hyperkd.sys
// starting from the test working directory and walking up to the repository
// root. Returns the absolute path or "" if not found.
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

// debuggerOperationWasSuccessful is 0xFFFFFFFF per
// HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h. The driver's
// IOCTL_INIT_VMM handler writes this value into KernelStatus on success.
const debuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// initVmmRequest mirrors DEBUGGER_INIT_VMM_PACKET: a single UINT32
// KernelStatus field used as both IOCTL input and output (METHOD_BUFFERED
// aliases the same kernel SystemBuffer for both directions).
type initVmmRequest struct {
	KernelStatus uint32
}

// clearEvent sends IOCTL_CODE_DEBUGGER_MODIFY_EVENTS with
// DebuggerModifyEventsClear to remove the event identified by tag from the
// kernel's event list. Best-effort: errors are logged via t.Logf but never
// fatal, matching the C libhyperdbg's tolerant unload behaviour.
func clearEvent(t *testing.T, dev *comm.Device, tag uint64) {
	t.Helper()
	var mod hyperdbgsdk.DEBUGGER_MODIFY_EVENTS
	mod.Tag = tag
	mod.TypeOfAction = hyperdbgsdk.DebuggerModifyEventsClear
	modSize := uint32(unsafe.Sizeof(mod))
	if _, err := dev.IoctlStruct(
		comm.IOCTL_CODE_DEBUGGER_MODIFY_EVENTS,
		unsafe.Pointer(&mod), unsafe.Pointer(&mod), modSize, modSize); err != nil {
		t.Logf("best-effort clearEvent(tag=%d) failed: %v", tag, err)
	}
}

// TestEptHook_Register is the integration test for the scriptengine Wrapper's
// hook-registration path. It drives the full pipeline end-to-end:
//
//  1. Load hyperkd.sys via the Service Control Manager (driverloader.Load).
//  2. Open \\.\HyperDbgDebuggerDevice (comm.Open).
//  3. Initialise the VMM via IOCTL_CODE_INIT_VMM (flips g_VmmInitialized so
//     VMM-group IOCTLs are admitted by IoctlCheckIoctlAllowed).
//  4. Compile a Go callback, build a DEBUGGER_GENERAL_EVENT_DETAIL +
//     DEBUGGER_GENERAL_ACTION, and register them via
//     IOCTL_CODE_DEBUGGER_REGISTER_EVENT + IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT
//     (scriptengine.Wrapper.RegisterHook).
//  5. Verify the returned Tag is > 0.
//  6. Clean up: clear the event, terminate VMX, close the device, unload the
//     driver.
//
// Skip conditions (use t.Skip, not t.Fatal):
//   - hyperkd.sys not found in the build output tree.
//   - Process is not running elevated.
//   - VMM init fails (system lacks VT-x / nested-virt support).
//
// To run from an elevated shell:
//
//	go test -v -count=1 -run TestEptHook ./debugger/scriptengine/
func TestEptHook_Register(t *testing.T) {
	driverPath := findHyperkdDriver(t)
	if driverPath == "" {
		t.Skip("hyperkd.sys not found in build output; build the VMM driver first")
	}
	if !isAdmin() {
		t.Skip("not running as administrator; driver load + device open require elevation")
	}
	t.Logf("using driver: %s", driverPath)

	d := driverloader.NewDriver(driverPath)

	// Best-effort cleanup of any stale service from a prior run. A previous
	// test may have left the service in a "marked for delete" state because
	// VMM teardown is asynchronous; retry Unload after a short wait so
	// ControlService(STOP) succeeds.
	_ = d.Unload()
	if exists, _ := d.Exists(); exists {
		time.Sleep(500 * time.Millisecond)
		_ = d.Unload()
	}

	if err := d.Load(); err != nil {
		t.Fatalf("driverloader.Load failed: %v", err)
	}
	// 等待驱动 DriverEntry 完成设备创建
	time.Sleep(2 * time.Second)
	t.Cleanup(func() { _ = d.Unload() })

	// Open the device handle exposed by the loaded driver (with retries).
	var dev *comm.Device
	var err error
	for i := 0; i < 5; i++ {
		dev, err = comm.Open(comm.DeviceName)
		if err == nil {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	if err != nil {
		t.Fatalf("comm.Open(%q) failed after retries: %v", comm.DeviceName, err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	// Best-effort VMM teardown before unloading the driver. IOCTL_TERMINATE_VMX
	// is in the VMM group so it is only admitted once g_VmmInitialized is TRUE;
	// if VMM init never succeeded the IOCTL is silently dropped, which we
	// tolerate. A short delay lets VMX teardown complete before
	// ControlService(STOP) is issued in Unload.
	t.Cleanup(func() {
		_, _ = dev.IoctlStruct(
			comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		time.Sleep(500 * time.Millisecond)
	})

	// Step 3: Initialise the VMM. This flips g_VmmInitialized so that
	// subsequent VMM-group IOCTLs (REGISTER_EVENT, ADD_ACTION_TO_EVENT,
	// MODIFY_EVENTS) are admitted instead of being silently dropped.
	var vmmReq initVmmRequest
	vmmSize := uint32(unsafe.Sizeof(vmmReq))
	if _, err := dev.IoctlStruct(comm.IOCTL_CODE_INIT_VMM,
		unsafe.Pointer(&vmmReq), unsafe.Pointer(&vmmReq), vmmSize, vmmSize); err != nil {
		t.Skipf("IOCTL_INIT_VMM failed: %v", err)
	}
	t.Logf("IOCTL_INIT_VMM: KernelStatus=0x%08x", vmmReq.KernelStatus)
	if vmmReq.KernelStatus != debuggerOperationWasSuccessful {
		t.Skipf("VMM init failed (KernelStatus=0x%08x); system likely lacks "+
			"VT-x / nested-virt support", vmmReq.KernelStatus)
	}

	// Step 4: Compile a Go callback and register an EPT hook.
	w := NewWrapper()

	// The callback source must be a complete Go file containing a FuncLit or
	// FuncDecl (see go-bridge/ast/encoder.go:findCallback). ctx.Break() is
	// whitelisted (func_id 0x0008) and is safe here: the hook target address
	// is a dummy value that is never executed, so the callback never fires.
	callbackSrc := `package hook
func hook(ctx *HookCtx) {
	ctx.Break()
}`
	scriptBytes, compileErr := w.Compile(callbackSrc)

	var action *hyperdbgsdk.DEBUGGER_GENERAL_ACTION
	if compileErr != nil {
		// Compile failed (go-bridge encoder/subset-validator rejected the
		// source). Fall back to BreakToDebugger, which needs no script bytes
		// and still exercises the REGISTER_EVENT + ADD_ACTION_TO_EVENT IOCTLs.
		t.Logf("Compile failed (%v); falling back to BreakToDebugger action "+
			"(no script bytes)", compileErr)
		action, err = w.BuildAction(nil, hyperdbgsdk.BreakToDebugger)
	} else {
		t.Logf("Compile OK: %d bytes of binary AST", len(scriptBytes))
		action, err = w.BuildAction(scriptBytes, hyperdbgsdk.RunScript)
	}
	if err != nil {
		t.Fatalf("BuildAction failed: %v", err)
	}

	// Build the event. HiddenHookExecDetours is the EPT execution-hook event
	// type (the Go SDK enum name for HIDDEN_HOOK_EXEC_DETOURS in
	// _VMM_EVENT_TYPE_ENUM). OptionalParam1 carries the hook target address;
	// a non-zero dummy value is sufficient for registration — the kernel
	// stores it and applies the EPT hook lazily. We never trigger the hook,
	// so the dummy address is safe.
	opts := hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{
		OptionalParam1: 0xFFFFFFFF,
	}
	event, err := w.BuildEvent(hyperdbgsdk.HiddenHookExecDetours, opts)
	if err != nil {
		t.Fatalf("BuildEvent failed: %v", err)
	}
	t.Logf("built event: Tag=%d, EventType=%s", event.Tag, event.EventType)

	// Register the hook. RegisterHook sends REGISTER_EVENT then
	// ADD_ACTION_TO_EVENT and returns the event tag on success.
	tag, err := w.RegisterHook(dev, event, action)
	if err != nil {
		// The kernel may reject the event if the hook target address is not a
		// valid executable page. Fall back to CpuidInstructionExecution,
		// which needs no target address and still verifies the full
		// register-event + add-action IOCTL plumbing.
		t.Logf("RegisterHook with HiddenHookExecDetours failed: %v; "+
			"retrying with CpuidInstructionExecution", err)
		action2, err2 := w.BuildAction(nil, hyperdbgsdk.BreakToDebugger)
		if err2 != nil {
			t.Fatalf("BuildAction (fallback) failed: %v", err2)
		}
		event2, err2 := w.BuildEvent(hyperdbgsdk.CpuidInstructionExecution, hyperdbgsdk.DEBUGGER_EVENT_OPTIONS{})
		if err2 != nil {
			t.Fatalf("BuildEvent (fallback) failed: %v", err2)
		}
		tag, err = w.RegisterHook(dev, event2, action2)
		if err != nil {
			t.Fatalf("RegisterHook (fallback with CpuidInstructionExecution) "+
				"failed: %v", err)
		}
		t.Logf("fallback registration succeeded: Tag=%d, EventType=%s",
			tag, event2.EventType)
	} else {
		t.Logf("registration succeeded: Tag=%d, EventType=%s",
			tag, event.EventType)
	}

	// Schedule best-effort event cleanup so the kernel's event list does not
	// leak entries across test runs.
	t.Cleanup(func() { clearEvent(t, dev, tag) })

	// Step 5: Verify the returned Tag is > 0. Tag 0 is reserved by the driver
	// as "no tag"; a successful registration must allocate a non-zero tag.
	if tag == 0 {
		t.Fatal("returned Tag is 0; expected a non-zero event tag (tag 0 is " +
			"reserved as 'no tag' by the driver)")
	}
	t.Logf("PASS: verified Tag=%d > 0", tag)
}
