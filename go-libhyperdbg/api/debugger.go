// Package api is the top-level HyperDbg debugger API. It wraps the core
// Debugger with an Option-based constructor and an Output interface so that
// CLI, GUI (future), and MCP (future) can all consume the same API without
// REPL dependencies or global state.
//
// Usage:
//
//	dbg, _ := api.New(api.WithOutput(os.Stdout))
//	dbg.Connect(ctx, "local")
//	dbg.LoadVMM(ctx, `Debug\hyperkd.sys`)
//	hookID, _ := dbg.EptHook(ctx, 0x00c12000, `func hook(ctx *HookCtx) { ctx.Break() }`)
package api

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	dbgcmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/debugging"
	extcmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/extension"
	hwdbgcmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/hwdbg"
	metacmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/meta"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
	"github.com/hyperdbg/go-libhyperdbg/symbolparser"
)

// Debugger is the public API wrapper around core.Debugger. It adds mutex
// protection for concurrent access (GUI/MCP need this), an Output interface,
// and convenience methods.
type Debugger struct {
	mu       sync.Mutex
	core     *core.Debugger
	output   Output
	commands *commands.Registry
	symbols  symbolparser.Resolver
}

// Output abstracts command output. CLI passes os.Stdout, GUI passes a widget
// writer, MCP passes a JSON channel writer.
type Output interface {
	io.Writer
	Printf(format string, args ...any) error
}

// stdoutOutput wraps an io.Writer (typically os.Stdout) to satisfy Output.
type stdoutOutput struct{ w io.Writer }

func (s *stdoutOutput) Write(p []byte) (int, error) { return s.w.Write(p) }
func (s *stdoutOutput) Printf(format string, args ...any) error {
	_, err := fmt.Fprintf(s.w, format, args...)
	return err
}

// Option configures a Debugger at construction time.
type Option func(*Debugger) error

// WithOutput sets the output destination for command results.
func WithOutput(w io.Writer) Option {
	return func(d *Debugger) error {
		d.output = &stdoutOutput{w: w}
		return nil
	}
}

// WithSymbolResolver injects a symbol resolver (e.g. symbolparser.New()).
// When set, EptHookSymbol can resolve "module!symbol" strings to addresses
// before registering the hook. Without a resolver, only EptHook (which takes
// a raw address) is usable.
func WithSymbolResolver(r symbolparser.Resolver) Option {
	return func(d *Debugger) error {
		d.symbols = r
		return nil
	}
}

// New creates a Debugger with the given options. The command registry is
// initialised after options are applied, so WithOutput must be passed here
// rather than changed later (the registry captures the output sink).
func New(opts ...Option) (*Debugger, error) {
	d := &Debugger{
		core:   core.New(),
		output: &stdoutOutput{w: io.Discard},
	}
	for _, opt := range opts {
		if err := opt(d); err != nil {
			return nil, err
		}
	}
	d.commands = commands.NewRegistry(d.output)
	metacmds.RegisterAll(d.commands)
	dbgcmds.RegisterAll(d.commands)
	extcmds.RegisterAll(d.commands)
	hwdbgcmds.RegisterAll(d.commands)
	return d, nil
}

// Connect opens the HyperDbg device for the given target ("local" for local
// debugging).
func (d *Debugger) Connect(ctx context.Context, target string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Connect(ctx, target)
}

// LoadVMM installs and starts the VMM driver.
func (d *Debugger) LoadVMM(ctx context.Context, driverPath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.LoadVMM(ctx, driverPath)
}

// UnloadVMM terminates the VMM and removes the driver service.
func (d *Debugger) UnloadVMM(ctx context.Context) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.UnloadVMM(ctx)
}

// Close releases all resources.
func (d *Debugger) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Close()
}

// EptHook registers an EPT execution hook at hookAddress with a Go callback.
// The callback source is compiled to binary AST and sent to the driver.
// Returns the hook ID (event tag).
func (d *Debugger) EptHook(ctx context.Context, hookAddress uint64, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EptHook(ctx, hookAddress, callbackSrc)
}

// EptHookForProcess registers an EPT execution hook at hookAddress for a
// specific process (pid). This is required for WOW64 target processes whose
// DLL addresses are not valid in the debugger process's address space.
func (d *Debugger) EptHookForProcess(ctx context.Context, hookAddress uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EptHookForProcess(ctx, hookAddress, pid, callbackSrc)
}

// EptHookSymbol registers an EPT hook at the address of the given symbol
// string (e.g. "ntdll!RtlAllocateHeap"). Requires a SymbolResolver to have
// been injected via WithSymbolResolver; otherwise it returns an error.
func (d *Debugger) EptHookSymbol(ctx context.Context, symbol string, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	resolver := d.symbols
	d.mu.Unlock()
	if resolver == nil {
		return 0, fmt.Errorf("EptHookSymbol(%q): no symbol resolver injected (use WithSymbolResolver)", symbol)
	}
	addr, err := resolver.FromName(ctx, symbol)
	if err != nil {
		return 0, fmt.Errorf("EptHookSymbol(%q): resolve failed: %w", symbol, err)
	}
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EptHook(ctx, addr, callbackSrc)
}

// SymbolResolver returns the injected symbol resolver (nil if none). GUI/MCP
// layers use it to resolve addresses for display.
func (d *Debugger) SymbolResolver() symbolparser.Resolver {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.symbols
}

// StartProcess launches a process for debugging. The VMM (already loaded)
// intercepts the child via the debug-port callback. Returns a Process
// handle the caller owns (call proc.Close when done).
func (d *Debugger) StartProcess(ctx context.Context, exePath string) (core.Process, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.StartProcess(ctx, exePath)
}

// Continue resumes the debugged process.
func (d *Debugger) Continue(ctx context.Context) error {
	// 不持 api.mu —— core 层管理锁，等 PAUSED 期间允许查询并发
	return d.core.Continue(ctx)
}

// Pause halts the debugged process.
func (d *Debugger) Pause(ctx context.Context) error {
	return d.core.Pause(ctx)
}

// LogOpen opens a file for log output.
func (d *Debugger) LogOpen(path string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.LogOpen(path)
}

// StartMessagePump spawns a goroutine that drains kernel log messages (produced
// by hook callbacks via ctx.Printf → LogInfo) to the open log file. It must be
// called after LogOpen and after LoadVMM. The returned pump MUST be stopped
// (pump.Stop) before UnloadVMM so the dedicated device handle is released
// cleanly and the goroutine exits.
//
// The pump opens a SECOND device handle so the pending IRP does not block the
// main IOCTL handle used by Continue/Pause/EptHook/… — this matches the C++
// ReadIrpBasedBuffer pattern.
func (d *Debugger) StartMessagePump(ctx context.Context) (*core.MessagePump, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.StartMessagePump(ctx)
}

// SetOnPaused registers a callback that fires (from the MessagePump goroutine)
// whenever the debuggee pauses: breakpoint hit, single-step complete, manual
// Pause, OEP break, etc. UI layers use this to auto-refresh registers/disasm/
// stack without waiting for a button press. Pass nil to unregister.
//
// The callback runs in the pump goroutine and must not call core methods that
// acquire d.mu (risk of deadlock). Use a goroutine or channel to defer work.
func (d *Debugger) SetOnPaused(fn func()) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.core.OnPaused = fn
}

// ReadMemory reads size bytes from the target process at the given virtual
// address. It triggers page-fault mapping for demand-paged DLLs, which is
// required before EptHookForProcess can validate the address (MmGetPhysicalAddress
// returns 0 for not-present pages).
func (d *Debugger) ReadMemory(ctx context.Context, addr uint64, pid uint32, size uint32) ([]byte, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	data, _, err := d.core.ReadMemory(ctx, addr, pid, size)
	return data, err
}

// LogClose closes the log file.
func (d *Debugger) LogClose() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.LogClose()
}

// Sleep waits for the given duration. This is a convenience for scripts that
// need to let the target process run before pausing.
func (d *Debugger) Sleep(dur time.Duration) {
	time.Sleep(dur)
}

// Printf writes formatted output to the debugger's Output.
func (d *Debugger) Printf(format string, args ...any) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.output.Printf(format, args...)
}

// Exec parses and runs a command line (e.g. "g", "load vmm", ".logopen x.txt").
// It is the single entry point the CLI REPL and MCP tools use to drive the
// debugger without calling individual methods. Returns meta.ErrExit when the
// user runs `exit` (the CLI loop checks for this).
func (d *Debugger) Exec(ctx context.Context, cmdLine string) error {
	// Exec does not take the api mutex itself: command handlers acquire the
	// core mutex as needed (via core.Debugger methods), and the registry's
	// output sink is goroutine-safe. This avoids deadlocks when a handler
	// calls back into the api layer.
	return d.commands.Exec(ctx, d.core, cmdLine)
}

// Commands returns the underlying command registry. GUI/MCP layers use it to
// list commands, build help UIs, or dispatch directly without going through
// Exec's string parser.
func (d *Debugger) Commands() *commands.Registry {
	return d.commands
}
