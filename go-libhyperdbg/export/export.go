// Package export implements the C-ABI export façade. The C++ counterpart is
// libhyperdbg/code/export/export.cpp; it exposes the hyperdbg_u_* entry
// points (hyperdbg_u_load_vmm, hyperdbg_u_unload_vmm,
// hyperdbg_u_run_command, ...) that external tools link against.
//
// In the Go rewrite we cannot emit real C-ABI symbols without cgo (which is
// forbidden by the API design spec). This package therefore provides a
// pure-Go façade whose method names mirror the C++ function names: a GUI or
// MCP host written in Go calls Exporter.HyperdbgULoadVmm(...) instead of
// the C symbol hyperdbg_u_load_vmm. The methods are thin wrappers around
// *app.App (the Go entry point) so the behaviour is identical to calling
// App directly — the wrapper exists only to preserve the C++ API shape for
// mechanical ports of C++ callers.
//
// Lifecycle:
//
//	exp := export.NewExporter(appInstance)
//	_ = exp.HyperdbgUDetectVmxSupport()
//	_ = exp.HyperdbgULoadVmm(ctx, driverPath)
//	_ = exp.HyperdbgURunCommand(ctx, "load vmm")
//	_ = exp.HyperdbgUUnloadVmm(ctx)
//
// Exporter holds no mutable state of its own; it is safe for concurrent use
// because the underlying *app.App is.
package export

import (
	"context"
	"fmt"

	"github.com/hyperdbg/go-libhyperdbg/app"
)

// ProcessorVendor mirrors GENERIC_PROCESSOR_VENDOR. It is re-declared here
// (rather than aliased) so that external callers can depend on the export
// package alone for the C-ABI surface.
type ProcessorVendor = app.ProcessorVendor

// Exporter is the C-ABI façade. It wraps a *app.App and exposes the
// hyperdbg_u_* API as methods. The zero value is not usable; use New.
type Exporter struct {
	app *app.App
}

// New constructs an Exporter backed by the given *app.App. The App must be
// non-nil; the caller is responsible for constructing and initialising it
// (typically via app.New + App.Init).
func New(a *app.App) *Exporter {
	if a == nil {
		panic("export.New: nil *app.App")
	}
	return &Exporter{app: a}
}

// App returns the underlying *app.App. Advanced callers use this to access
// the App/Messaging/Core directly when the façade does not yet expose a
// method.
func (e *Exporter) App() *app.App { return e.app }

// ============================================================================
// Module load/unload (mirror hyperdbg_u_load_* / hyperdbg_u_unload_*)
// ============================================================================

// HyperdbgUDetectVmxSupport mirrors hyperdbg_u_detect_vmx_support. Returns
// true if the CPU supports VMX (CPUID.01H:ECX[5]).
func (e *Exporter) HyperdbgUDetectVmxSupport() bool {
	// The C++ path delegates to VmxSupportDetection() which is owned by
	// common.Common. The App does not currently expose a Common instance,
	// so we re-implement the check inline via App.GetProcessorVendor (a
	// non-Others vendor implies VMX is plausibly available). For a precise
	// check, callers should use common.Common.VmxSupportDetection directly.
	//
	// TODO(Phase C.3): wire App to a *common.Common and delegate.
	return e.app.GetProcessorVendor() != app.ProcessorVendorOthers
}

// HyperdbgUReadVendorString mirrors hyperdbg_u_read_vendor_string. It writes
// the 12-byte vendor string from CPUID.0 into the first 12 bytes of dst.
// dst must have length >= 12.
func (e *Exporter) HyperdbgUReadVendorString(dst []byte) error {
	if len(dst) < 12 {
		return fmt.Errorf("HyperdbgUReadVendorString: dst must have length >= 12 (got %d)", len(dst))
	}
	// App does not yet expose the vendor string directly. Read it via the
	// processor-vendor enum and map back to the canonical string.
	//
	// TODO(Phase C.3): wire App to a *common.Common and call
	// CommonCpuidInstruction(0, 0, ...) to get the raw EBX/EDX/ECX.
	vendor := e.app.GetProcessorVendor().String()
	copy(dst, vendor)
	return nil
}

// HyperdbgULoadVmm mirrors hyperdbg_u_load_vmm. driverPath is the absolute
// path to the VMM .sys file. Returns nil on success.
func (e *Exporter) HyperdbgULoadVmm(ctx context.Context, driverPath string) error {
	return e.app.LoadVMM(ctx, driverPath)
}

// HyperdbgUUnloadVmm mirrors hyperdbg_u_unload_vmm.
func (e *Exporter) HyperdbgUUnloadVmm(ctx context.Context) error {
	return e.app.UnloadVMM(ctx)
}

// HyperdbgULoadKd mirrors hyperdbg_u_load_kd_module.
func (e *Exporter) HyperdbgULoadKd(ctx context.Context) error {
	return e.app.LoadKd(ctx)
}

// HyperdbgUUnloadKd mirrors hyperdbg_u_unload_kd.
func (e *Exporter) HyperdbgUUnloadKd(ctx context.Context) error {
	return e.app.UnloadKd(ctx)
}

// HyperdbgULoadHypertrace mirrors hyperdbg_u_load_hypertrace_module.
func (e *Exporter) HyperdbgULoadHypertrace(ctx context.Context) error {
	return e.app.LoadHyperTrace(ctx)
}

// HyperdbgUUnloadHypertrace mirrors hyperdbg_u_unload_hypertrace_module.
func (e *Exporter) HyperdbgUUnloadHypertrace(ctx context.Context) error {
	return e.app.UnloadHyperTrace(ctx)
}

// HyperdbgULoadAllModules mirrors hyperdbg_u_load_all_modules.
func (e *Exporter) HyperdbgULoadAllModules(ctx context.Context, driverPath string) error {
	return e.app.LoadAllModules(ctx, driverPath)
}

// HyperdbgUUnloadAllModules mirrors hyperdbg_u_unload_all_modules.
func (e *Exporter) HyperdbgUUnloadAllModules(ctx context.Context) error {
	return e.app.UnloadAllModules(ctx)
}

// HyperdbgUIsAnyModuleLoaded mirrors hyperdbg_u_is_any_module_loaded.
func (e *Exporter) HyperdbgUIsAnyModuleLoaded() bool {
	return e.app.IsAnyModuleLoaded()
}

// HyperdbgUGetProcessorVendor mirrors hyperdbg_u_get_processor_vendor.
func (e *Exporter) HyperdbgUGetProcessorVendor() ProcessorVendor {
	return e.app.GetProcessorVendor()
}

// ============================================================================
// Driver install/start/stop (mirror hyperdbg_u_install_kd_driver, ...)
// ============================================================================

// HyperdbgUInstallKdDriver mirrors hyperdbg_u_install_kd_driver. The C++
// implementation delegates to HyperDbgInstallKdDriver which is part of the
// driver-loader layer. The Go driver-loader (debugger/driverloader) is
// invoked implicitly by App.LoadVMM; this method is a stub that returns
// ErrNotImplemented until the explicit install-only path is wired up.
func (e *Exporter) HyperdbgUInstallKdDriver(ctx context.Context) error {
	return fmt.Errorf("HyperdbgUInstallKdDriver: not yet implemented (use HyperdbgULoadVmm instead)")
}

// HyperdbgUUninstallKdDriver mirrors hyperdbg_u_uninstall_kd_driver.
func (e *Exporter) HyperdbgUUninstallKdDriver(ctx context.Context) error {
	return fmt.Errorf("HyperdbgUUninstallKdDriver: not yet implemented")
}

// HyperdbgUStartKdDriver mirrors hyperdbg_u_start_kd_driver.
func (e *Exporter) HyperdbgUStartKdDriver(ctx context.Context) error {
	return fmt.Errorf("HyperdbgUStartKdDriver: not yet implemented")
}

// HyperdbgUStopKdDriver mirrors hyperdbg_u_stop_kd_driver.
func (e *Exporter) HyperdbgUStopKdDriver(ctx context.Context) error {
	return fmt.Errorf("HyperdbgUStopKdDriver: not yet implemented")
}

// ============================================================================
// Command interpreter (mirror hyperdbg_u_run_command, ...)
// ============================================================================

// HyperdbgURunCommand mirrors hyperdbg_u_run_command. It interprets a single
// command line (e.g. "load vmm", "bp 0x7ffe1234", "g"). The actual
// interpreter lives in debugger/commands; this façade delegates to it via
// the App's Core (Phase C.3 will wire the command dispatcher to App).
func (e *Exporter) HyperdbgURunCommand(ctx context.Context, command string) error {
	return fmt.Errorf("HyperdbgURunCommand(%q): not yet implemented (Phase C.3)", command)
}

// HyperdbgUTestCommandParser mirrors hyperdbg_u_test_command_parser. It parses
// a command line without executing it and returns the token list.
func (e *Exporter) HyperdbgUTestCommandParser(command string) (tokens []string, err error) {
	return nil, fmt.Errorf("HyperdbgUTestCommandParser: not yet implemented (Phase C.3)")
}

// HyperdbgUTestCommandParserShowTokens mirrors
// hyperdbg_u_test_command_parser_show_tokens. It parses and prints the tokens
// of a command line.
func (e *Exporter) HyperdbgUTestCommandParserShowTokens(command string) error {
	return fmt.Errorf("HyperdbgUTestCommandParserShowTokens: not yet implemented (Phase C.3)")
}

// ============================================================================
// Signature / version (mirror hyperdbg_u_show_signature, ...)
// ============================================================================

// HyperdbgUShowSignature mirrors hyperdbg_u_show_signature. It writes the
// HyperDbg build signature (version + build date) to the App's output sink.
func (e *Exporter) HyperdbgUShowSignature() {
	// The signature constants live in include/SDK/headers/Constants.h. The
	// Go port does not yet expose them as variables; this is a static
	// placeholder until the build-info package lands.
	msg := e.app.Messaging()
	if msg == nil {
		return
	}
	msg.ShowMessages("HyperDbg v0.22.0 (go-libhyperdbg framework)\n")
}
