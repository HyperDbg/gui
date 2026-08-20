package interp

import (
	"fmt"
	"reflect"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
	"github.com/traefik/yaegi/interp"
)

// errNoDebugger is returned by HyperDbg symbol closures when the Interpreter
// was constructed without a bound *api.Debugger (dbg == nil). This lets
// scripts parse and type-check without a live debugger, while still failing
// clearly at runtime if a debugger-backed symbol is invoked.
var errNoDebugger = fmt.Errorf("hyperdbg: no debugger bound to interpreter")

// errNotImpl is returned by symbols whose api.Debugger counterpart is not yet
// implemented (RemoveHook/ReadMem/Events). They are registered so scripts
// parse and type-check; calling them surfaces a clear error rather than a
// "symbol not found" parse failure.
func errNotImpl(name string) error {
	return fmt.Errorf("hyperdbg.%s: not yet implemented", name)
}

// unwrapPanic converts a yaegi interp.Panic into a plain error so callers
// don't depend on yaegi types. Other errors (parse errors, etc.) are returned
// unchanged.
func unwrapPanic(err error) error {
	if p, ok := err.(interp.Panic); ok {
		return fmt.Errorf("interp panic: %v", p.Value)
	}
	return err
}

// registerHyperdbgSymbols injects the HyperDbg API into the yaegi interpreter
// under the hyperdbgPkgPath ("hyperdbg/hyperdbg") package path. Each exported
// symbol is a closure that captures the *api.Debugger (for the actual call).
//
// Exported symbols (matching the task spec):
//   - Version() string
//   - Connect(target string) error
//   - LoadVMM(driverPath string) error
//   - UnloadVMM() error
//   - StartProcess(exePath string) (core.Process, error)
//   - Continue() error
//   - Pause() error
//   - EptHook(hookAddress uint64, callbackSrc string) (uint64, error)
//   - RemoveHook(hookID uint64) error         // stub: not yet on api.Debugger
//   - ReadMem(addr uint64, buf []byte) error   // stub: not yet on api.Debugger
//   - Events() interface{}                     // stub: not yet on api.Debugger
//   - Exec(cmdLine string) error
//   - Sleep(dur time.Duration)
func registerHyperdbgSymbols(yi *interp.Interpreter, i *Interpreter) {
	dbg := i.dbg
	symbols := interp.Exports{
		hyperdbgPkgPath: map[string]reflect.Value{
			// Version returns the go-libhyperdbg framework version. Always
			// available (no debugger required) so scripts can feature-detect.
			"Version": reflect.ValueOf(func() string { return Version }),

			// Connect opens the HyperDbg device for the given target
			// ("local" for local debugging).
			"Connect": reflect.ValueOf(func(target string) error {
				if dbg == nil {
					return errNoDebugger
				}
				return dbg.Connect(target)
			}),

			// LoadVMM installs and starts the VMM driver from driverPath.
		"LoadVMM": reflect.ValueOf(func(driverPath string) error {
			if dbg == nil {
				return errNoDebugger
			}
			if err := dbg.LoadDriver(driverPath); err != nil {
				return err
			}
			return dbg.InitVMM()
		}),

			// UnloadVMM terminates the VMM and removes the driver service.
			"UnloadVMM": reflect.ValueOf(func() error {
				if dbg == nil {
					return errNoDebugger
				}
				return dbg.UnloadVMM()
			}),

			// StartProcess launches a debuggee process and returns its
			// handle/PID. The concrete core.Process type is returned so
			// scripts can access proc.Pid and proc.Handle directly.
			"StartProcess": reflect.ValueOf(func(exePath string) (core.Process, error) {
				if dbg == nil {
					return core.Process{}, errNoDebugger
				}
				return dbg.StartProcess(exePath)
			}),

			// Continue resumes the debugged process.
			"Continue": reflect.ValueOf(func() error {
				if dbg == nil {
					return errNoDebugger
				}
				return dbg.Continue()
			}),

			// Pause halts the debugged process.
			"Pause": reflect.ValueOf(func() error {
				if dbg == nil {
					return errNoDebugger
				}
				return dbg.Pause()
			}),

			// EptHook registers an EPT execution hook at hookAddress with a
			// Go callback (compiled to binary AST and sent to the driver).
			// Returns the hook ID (event tag).
			"EptHook": reflect.ValueOf(func(hookAddress uint64, callbackSrc string) (uint64, error) {
				if dbg == nil {
					return 0, errNoDebugger
				}
				return dbg.EptHook(hookAddress, callbackSrc)
			}),

			// RemoveHook removes a previously registered hook by ID. Stub:
			// api.Debugger.RemoveHook is not yet implemented.
			"RemoveHook": reflect.ValueOf(func(hookID uint64) error {
				return errNotImpl("RemoveHook")
			}),

			// ReadMem reads len(buf) bytes from addr into buf. Stub:
			// api.Debugger.ReadMem is not yet implemented.
			"ReadMem": reflect.ValueOf(func(addr uint64, buf []byte) error {
				return errNotImpl("ReadMem")
			}),

			// Events returns the hook-event channel (nil for now). Stub:
			// api.Debugger.Events is not yet implemented; returns nil so
			// scripts can range over it without panicking (range on a nil
			// channel blocks forever, which is the correct "no events"
			// behaviour).
			"Events": reflect.ValueOf(func() interface{} {
				return nil
			}),

			// Exec parses and runs a HyperDbg command line (e.g. "lm",
			// "load vmm", ".logopen x.txt"). This is the single entry point
			// the CLI REPL and MCP tools use to drive the debugger.
			"Exec": reflect.ValueOf(func(cmdLine string) error {
				if dbg == nil {
					return errNoDebugger
				}
				return dbg.Exec(cmdLine)
			}),

			// Sleep waits for the given duration. Convenience for scripts
			// that need to let the target run before pausing. Falls back to
			// time.Sleep when no debugger is bound so tests can use it.
			"Sleep": reflect.ValueOf(func(dur time.Duration) {
				if dbg != nil {
					dbg.Sleep(dur)
					return
				}
				time.Sleep(dur)
			}),
		},
	}
	_ = yi.Use(symbols)
}
