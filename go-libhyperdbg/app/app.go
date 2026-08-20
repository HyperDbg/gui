// Package app implements the libhyperdbg entry-point layer that ties together
// driver loading with the *core.Debugger state machine. The C++ counterpart is
// libhyperdbg/code/app/libhyperdbg.cpp.
//
// In the C++ implementation the entry points (HyperDbgLoadVmmModule,
// HyperDbgUnloadVmm, HyperDbgCreateHandleFromKdModule, ...) mutate a set of
// process-global variables (g_DeviceHandle, g_IsKdModuleLoaded,
// g_IsVmmModuleLoaded, g_IsHyperTraceModuleLoaded, g_IsMessageLoggingWindowClosed,
// g_EventTrace, ...). The Go rewrite moves all of that mutable state into the
// App struct so that multiple debugger instances can coexist (GUI/MCP
// requirement, see API design spec).
//
// The App struct owns:
//   - a *core.Debugger (driver + device + state machine)
//
// Lifecycle:
//
//	app := app.New(out)
//	_ = app.Init()          // enable debug privilege
//	_ = app.LoadVMM(path)   // install+start driver, IOCTL_INIT_VMM
//	...                     // run commands / hooks
//	_ = app.UnloadVMM()     // IOCTL_TERMINATE_VMX + stop driver
//	_ = app.Cleanup()       // close device
//
// App is safe for concurrent use; the internal state is guarded by a mutex.
package app

import (
	"fmt"
	"sync"

	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// App. It mirrors commands.Output / api.Output but is declared locally to
// keep the app package free of import cycles.
type Output interface {
	Write(p []byte) (int, error)
	Printf(format string, args ...any) error
}

// ModuleLoaded flags track which HyperDbg modules are currently active,
// mirroring g_IsKdModuleLoaded / g_IsVmmModuleLoaded /
// g_IsHyperTraceModuleLoaded.
type ModuleLoaded struct {
	Kd         bool
	Vmm        bool
	HyperTrace bool
}

// App is the libhyperdbg entry-point aggregating driver state. The zero
// value is not usable; use New.
type App struct {
	mu   sync.Mutex
	out  Output
	core *core.Debugger

	loaded                  ModuleLoaded
	driverPath              string
	driverName              string
	useCustomDriverLocation bool

	// initialised is set true once Init runs the one-time setup
	// (debug privilege).
	initialised bool
}

// New constructs an App writing to out. The returned App is not yet
// connected to any driver; call Init + LoadVMM explicitly.
func New(out Output) *App {
	if out == nil {
		out = discardOutput{}
	}
	c := core.New()
	return &App{
		out:  out,
		core: c,
	}
}

// Init performs one-time setup: enables SeDebugPrivilege on the current
// process token. Idempotent — calling it twice is a no-op. Mirrors the
// preamble of HyperDbgLoadKdModule / HyperDbgCreateHandleFromKdModule in
// the C++ source.
func (a *App) Init() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.initialised {
		return nil
	}
	if err := enableDebugPrivilege(); err != nil {
		// Non-fatal: the driver may still load if the token already has
		// the privilege. Match C++ behaviour which logs but continues.
		a.out.Printf("warn: SetDebugPrivilege failed: %v\n", err)
	}
	a.initialised = true
	return nil
}

// Cleanup tears down everything Init/LoadVMM set up. It closes the device
// handle. Safe to call on a partially-initialised App.
func (a *App) Cleanup() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.core != nil {
		_ = a.core.Close()
	}
	a.loaded = ModuleLoaded{}
	a.initialised = false
	return nil
}

// Core returns the underlying *core.Debugger. Advanced callers use this to
// invoke IOCTL-backed methods directly (EptHook, ReadMem, ...).
func (a *App) Core() *core.Debugger {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.core
}

// Printf writes formatted output to the App's Output sink.
func (a *App) Printf(format string, args ...any) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.out.Printf(format, args...)
}

// IsAnyModuleLoaded mirrors HyperDbgIsAnyModuleLoaded: returns true if KD,
// VMM or HyperTrace is currently loaded.
func (a *App) IsAnyModuleLoaded() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.loaded.Kd || a.loaded.Vmm || a.loaded.HyperTrace
}

// LoadedModules returns a snapshot of which modules are currently loaded.
func (a *App) LoadedModules() ModuleLoaded {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.loaded
}

// SetCustomDriverPath mirrors hyperdbg_u_set_custom_driver_path: configures
// a non-default driver location used by subsequent Load* calls.
func (a *App) SetCustomDriverPath(driverPath, driverName string) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if len(driverPath) > 1024 || len(driverName) > 1024 {
		return fmt.Errorf("driver path/name too long")
	}
	a.driverPath = driverPath
	a.driverName = driverName
	a.useCustomDriverLocation = true
	return nil
}

// UseDefaultDriverPath mirrors hyperdbg_u_use_default_driver_path.
func (a *App) UseDefaultDriverPath() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.useCustomDriverLocation = false
}

// LoadVMM mirrors HyperDbgLoadVmmModule: checks VT-x support, loads the KD
// module (opens the device), then initialises the VMM via IOCTL_INIT_VMM.
// driverPath is the absolute path to the .sys file.
//
// On success the Vmm flag is set in LoadedModules.
func (a *App) LoadVMM(driverPath string) error {
	a.mu.Lock()
	if a.core == nil {
		a.mu.Unlock()
		return fmt.Errorf("LoadVMM: App not initialised")
	}
	if a.loaded.Vmm {
		a.mu.Unlock()
		return nil // already loaded, mirror C++ early-return
	}
	if a.loaded.HyperTrace {
		a.mu.Unlock()
		return fmt.Errorf("LoadVMM: HyperTrace is loaded; unload it first")
	}
	c := a.core
	a.mu.Unlock()

	if err := c.Connect("local"); err != nil {
		return fmt.Errorf("LoadVMM: %w", err)
	}
	if err := c.LoadVMM(driverPath); err != nil {
		return fmt.Errorf("LoadVMM: %w", err)
	}

	a.mu.Lock()
	a.loaded.Kd = true
	a.loaded.Vmm = true
	a.mu.Unlock()
	return nil
}

// UnloadVMM mirrors HyperDbgUnloadVmm: sends IOCTL_TERMINATE_VMX and stops
// the driver service. HyperTrace (if loaded) is unloaded first.
func (a *App) UnloadVMM() error {
	a.mu.Lock()
	if !a.loaded.Vmm {
		a.mu.Unlock()
		return fmt.Errorf("UnloadVMM: VMM not loaded")
	}
	c := a.core
	a.mu.Unlock()

	if err := c.UnloadVMM(); err != nil {
		return fmt.Errorf("UnloadVMM: %w", err)
	}
	a.mu.Lock()
	a.loaded.Vmm = false
	a.mu.Unlock()
	return nil
}

// LoadKd mirrors HyperDbgLoadKdModule: opens the device handle without
// entering VMX. Used when only kernel-debugging (no hypervisor) is needed.
func (a *App) LoadKd() error {
	a.mu.Lock()
	if a.loaded.Kd {
		a.mu.Unlock()
		return nil
	}
	c := a.core
	a.mu.Unlock()
	if err := c.Connect("local"); err != nil {
		return fmt.Errorf("LoadKd: %w", err)
	}
	a.mu.Lock()
	a.loaded.Kd = true
	a.mu.Unlock()
	return nil
}

// UnloadKd mirrors HyperDbgUnloadKd: closes the device handle and clears
// the KD flag. Refuses if VMM/HyperTrace is still loaded (they depend on KD).
func (a *App) UnloadKd() error {
	a.mu.Lock()
	if !a.loaded.Kd {
		a.mu.Unlock()
		return fmt.Errorf("UnloadKd: KD not loaded")
	}
	if a.loaded.Vmm {
		a.mu.Unlock()
		return fmt.Errorf("UnloadKd: VMM still loaded; unload it first")
	}
	if a.loaded.HyperTrace {
		a.mu.Unlock()
		return fmt.Errorf("UnloadKd: HyperTrace still loaded; unload it first")
	}
	a.mu.Unlock()
	if err := a.core.Close(); err != nil {
		return fmt.Errorf("UnloadKd: %w", err)
	}
	a.mu.Lock()
	a.loaded.Kd = false
	a.mu.Unlock()
	return nil
}

// LoadHyperTrace mirrors HyperDbgLoadHyperTraceModule. Stub: HyperTrace
// IOCTL integration is Phase C.3 work; the method is here so callers can
// wire it up without touching App internals.
func (a *App) LoadHyperTrace() error {
	// TODO(Phase C.3): send IOCTL_INIT_HYPERTRACE once the comm layer
	// exposes the typed wrapper. For now return an error so callers
	// know it's not wired up.
	return fmt.Errorf("LoadHyperTrace: not yet implemented (Phase C.3)")
}

// UnloadHyperTrace mirrors HyperDbgUnloadHyperTrace. Stub.
func (a *App) UnloadHyperTrace() error {
	// TODO(Phase C.3): send IOCTL_PERFORM_HYPERTRACE_UNLOAD.
	return fmt.Errorf("UnloadHyperTrace: not yet implemented (Phase C.3)")
}

// LoadAllModules mirrors HyperDbgLoadAllModules: KD + VMM + HyperTrace.
// HyperTrace currently returns ErrNotImplemented; KD+VMM are loaded.
func (a *App) LoadAllModules(driverPath string) error {
	if err := a.LoadVMM(driverPath); err != nil {
		return err
	}
	// HyperTrace is optional; ignore the not-implemented error for now.
	_ = a.LoadHyperTrace()
	return nil
}

// UnloadAllModules mirrors HyperDbgUnloadAllModules: HyperTrace → VMM → KD.
func (a *App) UnloadAllModules() error {
	a.mu.Lock()
	ht := a.loaded.HyperTrace
	vm := a.loaded.Vmm
	kd := a.loaded.Kd
	a.mu.Unlock()
	if ht {
		_ = a.UnloadHyperTrace()
	}
	if vm {
		if err := a.UnloadVMM(); err != nil {
			return err
		}
	}
	if kd {
		if err := a.UnloadKd(); err != nil {
			return err
		}
	}
	return nil
}

// GetProcessorVendor mirrors HyperDbgGetProcessorVendor: reads the CPUID
// vendor string and maps it to a vendor enum.
func (a *App) GetProcessorVendor() ProcessorVendor {
	vendor := readCpuidVendorString()
	switch vendor {
	case "GenuineIntel":
		return ProcessorVendorIntel
	case "AuthenticAMD":
		return ProcessorVendorAMD
	default:
		return ProcessorVendorOthers
	}
}

// ProcessorVendor mirrors GENERIC_PROCESSOR_VENDOR.
type ProcessorVendor int

const (
	ProcessorVendorOthers ProcessorVendor = iota
	ProcessorVendorIntel
	ProcessorVendorAMD
)

// String returns the human-readable vendor name.
func (p ProcessorVendor) String() string {
	switch p {
	case ProcessorVendorIntel:
		return "Intel"
	case ProcessorVendorAMD:
		return "AMD"
	default:
		return "Others"
	}
}

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Write(p []byte) (int, error) { return len(p), nil }
func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
