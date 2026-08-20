// Package core — epthook_realaddr_test.go
//
// Integration test for EPT hook registration with REAL user-mode addresses.
//
// This test directly verifies the fix for the 0xC0000005
// (DEBUGGER_ERROR_INVALID_ADDRESS) regression: when ProcessId is set to
// DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES (0xFFFFFFFF), the kernel's
// ValidateEventEptHookHiddenBreakpointAndInlineHooks replaces it with
// PsGetCurrentProcessId() (the Go test process) and validates the hook
// address against THAT process's address space. System DLLs (kernelbase,
// ntdll) are loaded in every user-mode process, so a real export address
// must be accepted.
//
// To run from an elevated shell:
//
//	go test -v -count=1 -run TestEptHook_RealAddress ./debugger/core/
package core

import (
	"os"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/driverloader"
)

// TestEptHook_RealAddress verifies that EptHook succeeds when given a real
// kernelbase!SetEvent address resolved via GetProcAddress. This is the exact
// scenario that was failing with 0xC0000005 before the ProcessId fix.
func TestEptHook_RealAddress(t *testing.T) {
	driverPath := findHyperkdDriver(t)
	if driverPath == "" {
		t.Skip("hyperkd.sys not found in build output; build the VMM driver first")
	}
	if !isAdmin() {
		t.Skip("not running as administrator; driver load + device open require elevation")
	}
	t.Logf("using driver: %s", driverPath)

	// Resolve real addresses via GetProcAddress. kernelbase!SetEvent is the
	// actual implementation (kernel32!SetEvent is a forwarder stub).
	kb := syscall.NewLazyDLL("kernelbase.dll")
	setEventProc := kb.NewProc("SetEvent")
	virtualAllocProc := kb.NewProc("VirtualAlloc")
	ntdll := syscall.NewLazyDLL("ntdll.dll")
	rtlAllocateHeapProc := ntdll.NewProc("RtlAllocateHeap")

	setEventAddr := uint64(setEventProc.Addr())
	virtualAllocAddr := uint64(virtualAllocProc.Addr())
	rtlAllocateHeapAddr := uint64(rtlAllocateHeapProc.Addr())

	t.Logf("resolved addresses:")
	t.Logf("  kernelbase!SetEvent       = 0x%X", setEventAddr)
	t.Logf("  kernelbase!VirtualAlloc   = 0x%X", virtualAllocAddr)
	t.Logf("  ntdll!RtlAllocateHeap     = 0x%X", rtlAllocateHeapAddr)

	if setEventAddr == 0 {
		t.Fatal("failed to resolve kernelbase!SetEvent")
	}

	d := driverloader.NewDriver(driverPath)

	// Best-effort cleanup of any stale service from a prior run.
	_ = d.Unload()
	if exists, _ := d.Exists(); exists {
		time.Sleep(500 * time.Millisecond)
		_ = d.Unload()
	}

	if err := d.Load(); err != nil {
		t.Fatalf("driverloader.Load failed: %v", err)
	}
	time.Sleep(2 * time.Second) // let DriverEntry create the device
	t.Cleanup(func() { _ = d.Unload() })

	// Open the device handle (with retries).
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

	// Best-effort VMX teardown before unloading the driver.
	t.Cleanup(func() {
		// Clear all events first to avoid EPT hook cleanup crashes.
		clearAll := hyperdbgsdk.DEBUGGER_MODIFY_EVENTS{
			TypeOfAction: hyperdbgsdk.DebuggerModifyEventsClear,
		}
		clearBuf := byteslice.FromStruct(&clearAll)
		var dummy [256]byte
		_, _ = dev.Ioctl(comm.IOCTL_CODE_DEBUGGER_MODIFY_EVENTS, clearBuf, dummy[:])
		_, _ = dev.IoctlStruct(
			comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		time.Sleep(2 * time.Second)
	})

	// Step 1: Initialise the VMM.
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

	// Step 2: Use core.Debugger.EptHookForProcess to register hooks scoped
	// to THIS test process only. Using EptHook (global, ProcessId=0xFFFFFFFF)
	// with ctx.Break() on system-critical APIs like SetEvent/VirtualAlloc/
	// RtlAllocateHeap is DANGEROUS: every process on the system triggers the
	// hook, and ctx.Break() in VMX-root mode on a system process (csrss,
	// winlogon) causes a deadlock or BSOD. EptHookForProcess limits the hook
	// to the test process, which doesn't call these APIs during the test.
	dbg := New()
	dbg.state = StateVmmLoaded // bypass LoadVMM (driver already loaded above)
	dbg.device = dev

	// Use the test process's own PID so the hook only fires for this process.
	pid := uint32(os.Getpid())
	t.Logf("using per-process hook (pid=%d) to avoid global system impact", pid)

	// Harmless callback: just reads the return address and logs it.
	// NOT ctx.Break() — that would pause the caller, which is fatal if a
	// system process somehow triggers it.
	callbackSrc := `package hook
func hook(ctx *HookCtx) {
	ctx.Printf("hook fired\n")
}`

	// Test 1: kernelbase!SetEvent (the primary failing case).
	// Note: core.Debugger's nextTag starts at 0, so Tag=0 is a valid first
	// hook id (the driver accepts it). Success is indicated by err == nil.
	tagSE, err := dbg.EptHookForProcess(setEventAddr, pid, callbackSrc)
	if err != nil {
		t.Errorf("EptHookForProcess(kernelbase!SetEvent @ 0x%X) failed: %v", setEventAddr, err)
	} else {
		t.Logf("PASS: EptHookForProcess(kernelbase!SetEvent @ 0x%X) → Tag=%d", setEventAddr, tagSE)
	}

	// Test 2: ntdll!RtlAllocateHeap (the OEP detection hook).
	if rtlAllocateHeapAddr != 0 {
		tagRAH, err := dbg.EptHookForProcess(rtlAllocateHeapAddr, pid, callbackSrc)
		if err != nil {
			t.Errorf("EptHookForProcess(ntdll!RtlAllocateHeap @ 0x%X) failed: %v", rtlAllocateHeapAddr, err)
		} else {
			t.Logf("PASS: EptHookForProcess(ntdll!RtlAllocateHeap @ 0x%X) → Tag=%d", rtlAllocateHeapAddr, tagRAH)
		}
	}

	// Test 3: kernelbase!VirtualAlloc (the OEP detection hook).
	if virtualAllocAddr != 0 {
		tagVA, err := dbg.EptHookForProcess(virtualAllocAddr, pid, callbackSrc)
		if err != nil {
			t.Errorf("EptHookForProcess(kernelbase!VirtualAlloc @ 0x%X) failed: %v", virtualAllocAddr, err)
		} else {
			t.Logf("PASS: EptHookForProcess(kernelbase!VirtualAlloc @ 0x%X) → Tag=%d", virtualAllocAddr, tagVA)
		}
	}
}
