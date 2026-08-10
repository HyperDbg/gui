package driverloader

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"golang.org/x/sys/windows"
)

// TestNewDriver_NameDerivation verifies the service name is derived from the
// file basename (without extension), matching the C ManageDriver convention.
func TestNewDriver_NameDerivation(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{`D:\HyperDbg\Debug\hyperhv.sys`, "hyperhv"},
		{`C:\drivers\hyperkd.sys`, "hyperkd"},
		{`hyperhv.sys`, "hyperhv"},
		{`/tmp/my_driver.SYS`, "my_driver"}, // case-insensitive ext strip
		{`nodriver`, "nodriver"},            // no extension
	}
	for _, c := range cases {
		d := NewDriver(c.path)
		if d.Name != c.want {
			t.Errorf("NewDriver(%q).Name = %q, want %q", c.path, d.Name, c.want)
		}
		if d.Path != c.path {
			t.Errorf("NewDriver(%q).Path = %q, want %q", c.path, d.Path, c.path)
		}
	}
}

// TestDriver_IdempotentRemoveOnNonExistent verifies Remove succeeds (is
// idempotent) when the service was never installed. This needs no admin
// rights on most configurations because OpenService on a missing service
// returns ERROR_SERVICE_DOES_NOT_EXIST which we treat as success.
//
// NOTE: OpenSCManager with SC_MANAGER_ALL_ACCESS may itself require admin
// rights; if so the test is skipped rather than failed.
func TestDriver_IdempotentRemoveOnNonExistent(t *testing.T) {
	t.Skip("requires SCM access; enable manually with admin rights")
	d := NewDriver(filepath.Join(os.TempDir(), "hyperdbg_nonexistent_driver.sys"))
	if err := d.Remove(context.Background()); err != nil {
		t.Fatalf("Remove on non-existent service should be idempotent: %v", err)
	}
}

// findVmmDriver searches common build output directories for the HyperDbg
// driver .sys file and returns its absolute path, or "" if not found.
//
// In this project's CMake configuration hyperhv is compiled as a static
// library (Hyperhv.lib) and linked into hyperkd.sys, so the loadable driver
// is hyperkd.sys. We search for both names for robustness.
func findVmmDriver(t *testing.T) string {
	t.Helper()
	candidates := []string{
		`Debug\hyperhv.sys`,
		`Debug\hyperkd.sys`,
		`Release\hyperhv.sys`,
		`Release\hyperkd.sys`,
		`x64\Debug\hyperhv.sys`,
		`x64\Debug\hyperkd.sys`,
		`x64\Release\hyperhv.sys`,
		`x64\Release\hyperkd.sys`,
	}
	// Walk up from the test working directory (module root) to find a Debug
	// folder containing the .sys.
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	dir := wd
	for i := 0; i < 6 && dir != "" && dir != filepath.Dir(dir); i++ {
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

// isAdmin reports whether the current process has administrator privileges
// by inspecting the process token elevation state via the Windows API.
func isAdmin() bool {
	var token windows.Token
	if err := windows.OpenProcessToken(windows.CurrentProcess(), windows.TOKEN_QUERY, &token); err != nil {
		return false
	}
	defer token.Close()
	return token.IsElevated()
}

// TestLoadUnloadVMM is the Phase A acceptance integration test.
//
// Flow: load hyperhv.sys → (device open + IOCTL_INIT_VMM is exercised in the
// comm package test) → stop+remove. This test is skipped unless the driver
// binary is discoverable AND the process is elevated.
//
// To run: build hyperhv.sys, run `go test -run TestLoadUnloadVMM` from an
// elevated shell.
func TestLoadUnloadVMM(t *testing.T) {
	driverPath := findVmmDriver(t)
	if driverPath == "" {
		t.Skip("hyperhv.sys/hyperkd.sys not found in build output; build the VMM driver first")
	}
	if !isAdmin() {
		t.Skip("not running as administrator; driver load requires elevation")
	}

	ctx := context.Background()
	d := NewDriver(driverPath)

	// Best-effort cleanup of any stale service from a prior run. A previous
	// test (e.g. TestGoInterpInKernel) may have left the service in a
	// "marked for delete" state because VMM teardown is asynchronous; retry
	// Unload after a short wait so ControlService(STOP) succeeds.
	_ = d.Unload(ctx)
	if exists, _ := d.Exists(ctx); exists {
		time.Sleep(500 * time.Millisecond)
		_ = d.Unload(ctx)
	}

	if err := d.Load(ctx); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload(ctx) })

	// Verify the service is registered.
	exists, err := d.Exists(ctx)
	if err != nil {
		t.Fatalf("Exists failed: %v", err)
	}
	if !exists {
		t.Fatal("service not registered after Load")
	}

	// Unload and verify removal.
	if err := d.Unload(ctx); err != nil {
		t.Fatalf("Unload failed: %v", err)
	}
	exists, err = d.Exists(ctx)
	if err != nil {
		t.Fatalf("Exists failed after Unload: %v", err)
	}
	if exists {
		t.Fatal("service still registered after Unload")
	}
}

// TestLoad_MissingFile verifies Load fails fast when the driver file does not
// exist, and that the error mentions the path.
func TestLoad_MissingFile(t *testing.T) {
	d := NewDriver(filepath.Join(os.TempDir(), "definitely_not_here_12345.sys"))
	err := d.Load(context.Background())
	if err == nil {
		t.Fatal("Load should fail for a missing driver file")
	}
	if !strings.Contains(err.Error(), "driver file not accessible") {
		t.Errorf("expected error about inaccessible driver file, got: %v", err)
	}
}

// debuggerPerformKernelTests mirrors the C DEBUGGER_PERFORM_KERNEL_TESTS
// struct (RequestStructures.h): a single UINT32 KernelStatus field used as
// both the IOCTL input and output payload (METHOD_BUFFERED aliases the buffer).
//
// DEBUGGER_INIT_VMM_PACKET has the identical layout, so the same Go type is
// reused for both IOCTL_INIT_VMM and IOCTL_PERFORM_KERNEL_SIDE_TESTS.
type debuggerPerformKernelTests struct {
	KernelStatus uint32
}

// DEBUGGER_OPERATION_WAS_SUCCESSFUL is 0xFFFFFFFF per
// HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h.
const debuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// sendKernelStatusIoctl sends an IOCTL whose payload is a single UINT32
// KernelStatus field used as both input and output (METHOD_BUFFERED aliases
// the same kernel SystemBuffer for both directions). It returns the
// KernelStatus the driver wrote back, plus the number of bytes the driver
// returned and any Win32-level error from DeviceIoControl.
//
// Note: the driver's DrvAdjustStatusAndSetOutputSize always completes the IRP
// with STATUS_SUCCESS (overriding per-handler failure codes), so a non-error
// return here does NOT imply the operation succeeded — callers must check the
// returned KernelStatus against DEBUGGER_OPERATION_WAS_SUCCESSFUL.
func sendKernelStatusIoctl(t *testing.T, dev *comm.Device, code uint32, label string) (status uint32, bytes uint32) {
	t.Helper()
	var req debuggerPerformKernelTests
	reqSize := uint32(unsafe.Sizeof(req))
	n, err := dev.IoctlStruct(context.Background(), code,
		unsafe.Pointer(&req), unsafe.Pointer(&req), reqSize, reqSize)
	if err != nil {
		t.Fatalf("%s IOCTL failed: %v (bytes=%d)", label, err, n)
	}
	t.Logf("%s: KernelStatus=0x%08x, bytes=%d", label, req.KernelStatus, n)
	return req.KernelStatus, n
}

// TestGoInterpInKernel is the Phase B end-to-end acceptance test for the
// in-kernel Go-subset interpreter.
//
// Flow: load hyperhv/hyperkd → open \\.\HyperDbgDebuggerDevice →
// IOCTL_INIT_VMM (to flip g_VmmInitialized so VMM-group IOCTLs are admitted
// by IoctlCheckIoctlAllowed) → IOCTL_PERFORM_KERNEL_SIDE_TESTS (whose handler
// TestKernelPerformTests now also drives GoInterpRunHookCallback on a minimal
// `func(ctx){}` AST) → verify KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL
// → IOCTL_TERMINATE_VMX (best-effort teardown).
//
// This is the first test that actually exercises the go-interp C code linked
// into hyperkd.sys from user-mode Go.
//
// To run: build hyperkd.sys, then run from an elevated shell:
//
//	go test -v -count=1 -run TestGoInterpInKernel ./debugger/driverloader/
func TestGoInterpInKernel(t *testing.T) {
	driverPath := findVmmDriver(t)
	if driverPath == "" {
		t.Skip("hyperhv.sys/hyperkd.sys not found in build output; build the VMM driver first")
	}
	if !isAdmin() {
		t.Skip("not running as administrator; driver load requires elevation")
	}

	ctx := context.Background()
	d := NewDriver(driverPath)

	// Best-effort cleanup of any stale service from a prior run.
	_ = d.Unload(ctx)

	if err := d.Load(ctx); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	t.Cleanup(func() { _ = d.Unload(ctx) })

	// Open the \\.\HyperDbgDebuggerDevice handle exposed by the loaded driver.
	dev, err := comm.Open(ctx, comm.DeviceName)
	if err != nil {
		t.Fatalf("comm.Open(%q) failed: %v", comm.DeviceName, err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	// Best-effort VMM teardown before unloading the driver. IOCTL_TERMINATE_VMX
	// is in the VMM group so it is only admitted once g_VmmInitialized is TRUE;
	// if VMM init never succeeded the IOCTL is silently dropped (STATUS_SUCCESS
	// + 0 bytes), which we tolerate — the driver unload path also cleans up.
	//
	// A short delay is needed after terminate so the driver's VMX teardown
	// completes before ControlService(STOP) is issued in Unload; otherwise
	// Stop can fail and the service entry lingers as "marked for delete",
	// breaking subsequent tests that expect a clean slate.
	t.Cleanup(func() {
		_, _ = dev.IoctlStruct(context.Background(),
			comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		time.Sleep(500 * time.Millisecond)
	})

	// Initialize the VMM. This flips g_VmmInitialized so that subsequent
	// VMM-group IOCTLs (including IOCTL_PERFORM_KERNEL_SIDE_TESTS) are
	// admitted by IoctlCheckIoctlAllowed instead of being silently dropped.
	//
	// VMM init enables VMX on every core; it may fail on systems without VT-x
	// or inside a non-nested VM. In that case we skip rather than fail.
	initStatus, _ := sendKernelStatusIoctl(t, dev,
		comm.IOCTL_CODE_INIT_VMM, "IOCTL_INIT_VMM")
	if initStatus != debuggerOperationWasSuccessful {
		t.Skipf("VMM init failed (KernelStatus=0x%08x); go-interp IOCTL would be "+
			"silently dropped. System likely lacks VT-x / nested-virt support",
			initStatus)
	}

	// Now exercise the go-interp path: TestKernelPerformTests runs
	// GoInterpRunHookCallback on a hard-coded `func(ctx){}` AST and sets
	// KernelStatus to DEBUGGER_OPERATION_WAS_SUCCESSFUL on a Continue action.
	testStatus, n := sendKernelStatusIoctl(t, dev,
		comm.IOCTL_CODE_PERFORM_KERNEL_SIDE_TESTS, "IOCTL_PERFORM_KERNEL_SIDE_TESTS")
	if n == 0 {
		t.Fatalf("IOCTL_PERFORM_KERNEL_SIDE_TESTS returned 0 bytes; " +
			"IoctlCheckIoctlAllowed dropped it despite VMM init succeeding")
	}
	if testStatus != debuggerOperationWasSuccessful {
		t.Fatalf("KernelStatus = 0x%08x, want 0x%08x (DEBUGGER_OPERATION_WAS_SUCCESSFUL); "+
			"go-interp self-check failed inside the driver",
			testStatus, debuggerOperationWasSuccessful)
	}
	t.Logf("go-interp ran `func(ctx){}` successfully inside hyperkd: "+
		"KernelStatus=0x%08x", testStatus)
}
