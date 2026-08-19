// Package interp wraps the yaegi Go interpreter to evaluate user-supplied Go
// source against the HyperDbg API. It is the user-mode counterpart of the
// kernel C interpreter: yaegi runs top-level script logic (REPL expressions,
// .go files), while the driver-side C interpreter runs hook callbacks in
// VMX-root.
//
// Design (per the API spec):
//   - No global state: each Interpreter owns its own yaegi instance and is
//     bound to a single *api.Debugger.
//   - Context propagation: Eval/EvalFile respect ctx cancellation (yaegi's
//     EvalWithContext/EvalPathWithContext runs the evaluation in a goroutine
//     and interrupts it via interp.stop() on ctx.Done). The current ctx is
//     also surfaced to registered HyperDbg symbols so API calls honour it.
//   - Panic isolation: yaegi's EvalWithContext recovers panics from user code
//     and returns them as errors (the SEH-equivalent required by the spec).
//   - Concurrency-safe: a mutex serialises Eval/EvalFile/Use calls because
//     yaegi's Interpreter is single-goroutine.
//   - No cgo: yaegi is pure Go.
package interp

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"sync"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// Version is the go-libhyperdbg framework version exposed to scripts via
// hyperdbg.Version(). It mirrors the signature string in
// go-libhyperdbg/export/export.go.
const Version = "0.22.0"

// hyperdbgPkgPath is the symbol-table key under which HyperDbg API symbols
// are registered with yaegi. The format is "importPath/packageName"
// (path.Dir/path.Base), so scripts import it as `import "hyperdbg"` and
// reference symbols as `hyperdbg.Connect`, `hyperdbg.Version`, etc. In REPL
// mode (after ImportUsed) the bare `hyperdbg` identifier works without an
// explicit import.
const hyperdbgPkgPath = "hyperdbg/hyperdbg"

// Interpreter wraps a yaegi.Interpreter with HyperDbg API symbols bound to a
// specific *api.Debugger instance. It is safe for concurrent use: mu
// serialises Eval/EvalFile/Use because yaegi's Interpreter is
// single-goroutine, and ctxMu protects ctx which is read by registered
// symbol closures running inside the interpreter's goroutine.
type Interpreter struct {
	// mu serialises Eval/EvalFile/Use. yaegi's EvalWithContext takes its own
	// internal mutex, but concurrent Evals on the same Interpreter are not
	// supported, so we add our own.
	mu sync.Mutex

	yi  *interp.Interpreter
	dbg *api.Debugger

	// ctxMu protects ctx. Symbol closures call currentCtx() to read it; Eval
	// writes it before invoking yaegi. An RWMutex is used because reads (one
	// per API call in user code) outnumber writes (one per Eval).
	ctxMu sync.RWMutex
	ctx   context.Context
}

// NewInterpreter creates a yaegi interpreter pre-loaded with the Go standard
// library and the HyperDbg API. The HyperDbg symbols (hyperdbg.Connect,
// hyperdbg.LoadVMM, hyperdbg.EptHook, etc.) are bound to dbg and auto-imported
// via ImportUsed so REPL expressions can reference them without an explicit
// import statement (e.g. `hyperdbg.Version()` works directly).
//
// If dbg is nil, the HyperDbg symbols are still registered but calling any
// that need a debugger returns errNoDebugger. This supports dry-running
// scripts (and the unit tests) without a live debugger.
func NewInterpreter(dbg *api.Debugger) *Interpreter {
	yi := interp.New(interp.Options{
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
	// stdlib.Symbols registers every Go standard-library package so scripts
	// can use fmt, strings, time, etc. Errors here are fatal in practice
	// (only happens on malformed symbol tables); we ignore them so a bad
	// stdlib registration doesn't panic the caller.
	_ = yi.Use(stdlib.Symbols)

	i := &Interpreter{yi: yi, dbg: dbg}
	registerHyperdbgSymbols(yi, i)

	// Auto-import all registered packages (stdlib + hyperdbg) into the
	// universe scope so REPL expressions like "hyperdbg.Version()" or
	// "fmt.Println(1)" work without an explicit import block. Must be called
	// once, before any Eval.
	yi.ImportUsed()
	return i
}

// currentCtx returns the context associated with the in-flight Eval call, or
// context.Background() if no Eval is running. Symbol closures call this to
// propagate cancellation to api.Debugger methods. Safe to call from any
// goroutine.
func (i *Interpreter) currentCtx() context.Context {
	i.ctxMu.RLock()
	defer i.ctxMu.RUnlock()
	if i.ctx != nil {
		return i.ctx
	}
	return context.Background()
}

// setCtx stores the evaluation context. Called under i.mu.
func (i *Interpreter) setCtx(ctx context.Context) {
	i.ctxMu.Lock()
	i.ctx = ctx
	i.ctxMu.Unlock()
}

// Eval interprets a Go source string. The ctx is respected: if cancelled,
// Eval returns ctx.Err() promptly (yaegi's EvalWithContext runs the evaluation
// in a goroutine and interrupts it via interp.stop() on ctx.Done, which
// aborts interpreted loops like `(func(){ for {} })()`).
//
// Panics from user code are recovered by yaegi and returned as errors of type
// interp.Panic (the SEH-equivalent required by the spec); Eval unwraps them
// to a plain error so callers don't depend on yaegi types.
//
// The returned interface{} is the reflect.Value.Interface() of yaegi's result:
// for expression evaluations like "1+1" it is the expression value (int(2)),
// for statements it is nil.
func (i *Interpreter) Eval(ctx context.Context, src string) (interface{}, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.setCtx(ctx)

	v, err := i.yi.EvalWithContext(ctx, src)
	if err != nil {
		return nil, unwrapPanic(err)
	}
	if !v.IsValid() {
		return nil, nil
	}
	return v.Interface(), nil
}

// EvalFile reads and evaluates a Go script file. Context and panic handling
// are identical to Eval. The path must be readable.
func (i *Interpreter) EvalFile(ctx context.Context, path string) (interface{}, error) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.setCtx(ctx)

	v, err := i.yi.EvalPathWithContext(ctx, path)
	if err != nil {
		return nil, unwrapPanic(err)
	}
	if !v.IsValid() {
		return nil, nil
	}
	return v.Interface(), nil
}

// Use registers additional symbols with the underlying yaegi interpreter,
// extending the set available to scripts beyond stdlib and hyperdbg. pkg may
// be either:
//   - interp.Exports (map[string]map[string]reflect.Value): registered as-is,
//     path is ignored.
//   - map[string]reflect.Value (symbol name -> value): registered under the
//     given path (e.g. "myproj/myproj").
//
// Example:
//
//	i.Use(map[string]reflect.Value{
//	    "Greet": reflect.ValueOf(func(name string) string { return "hi " + name }),
//	}, "myapp/myapp")
func (i *Interpreter) Use(pkg interface{}, path string) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	switch s := pkg.(type) {
	case interp.Exports:
		return i.yi.Use(s)
	case map[string]reflect.Value:
		return i.yi.Use(interp.Exports{path: s})
	default:
		return fmt.Errorf("interp.Use: unsupported pkg type %T (want interp.Exports or map[string]reflect.Value)", pkg)
	}
}

// Debugger returns the *api.Debugger bound to this interpreter. Scripts don't
// need this (they use the hyperdbg.* symbols), but GUI/MCP layers may use it
// to access the underlying debugger directly.
func (i *Interpreter) Debugger() *api.Debugger {
	return i.dbg
}
