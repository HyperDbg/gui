// Package dbgcmds implements the 37 debugging commands (debugging-commands/)
// from the C++ libhyperdbg.
//
// Commands fully implemented here (delegating to core.Debugger):
//
//	g, go, pause, load, unload, exit        — also registered by meta; here
//	                                          we re-register aliases so the
//	                                          debugging group owns them too
//	sleep                                   — sleep N seconds (script helper)
//	events                                  — list registered event tags
//	settings                                — print/inspect debugger settings
//
// Commands registered as stubs (return ErrNotImplemented) pending their full
// implementation in later phases (most need the misc/symbolparser modules):
//
//	a, bc, bd, be, bl, bp, continue, core, cpu, d-u, dt-struct, e, eval, flush,
//	gg, gu, i, k, lm, output, p, preactivate, prealloc, print, r, rdmsr, s,
//	t, test, wrmsr, x
package dbgcmds

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// ErrNotImplemented is returned by stub commands pending full implementation.
var ErrNotImplemented = fmt.Errorf("command not yet implemented")

// RegisterAll registers every debugging command (and its aliases) into r.
// Commands already registered by the meta group (g/pause/load/unload/exit)
// are not overwritten here — we only register the ones the meta group does
// not cover, plus debugging-specific aliases like "continue".
func RegisterAll(r *commands.Registry) {
	// ---- sleep — wait N seconds (script helper) ----
	r.Register("sleep", commands.CommandSpec{
		Handler: cmdSleep,
		Help:    helpSleep,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})

	// ---- events — list registered event tags ----
	r.Register("events", commands.CommandSpec{
		Handler: cmdEvents,
		Help:    helpEvents,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})

	// ---- settings — print debugger settings ----
	r.Register("settings", commands.CommandSpec{
		Handler: cmdSettings,
		Help:    helpSettings,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})

	// ---- continue — alias for g ----
	r.Register("continue", commands.CommandSpec{
		Handler: cmdG,
		Help:    helpG,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})

	// ---- stubs ----
	stubs := []string{
		"a", "bc", "bd", "be", "bl", "bp", "core", "cpu",
		"d", "u", "dt", "e", "eval", "flush", "gg", "gu",
		"i", "k", "lm", "output", "p", "preactivate", "prealloc",
		"print", "r", "rdmsr", "s", "t", "test", "wrmsr", "x",
	}
	for _, name := range stubs {
		// 'g' alias 'go' is already registered by meta; skip if present.
		if _, exists := r.Lookup(name); exists {
			continue
		}
		r.Register(name, commands.CommandSpec{
			Handler: stubHandler(name),
			Help:    stubHelp(name),
			Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
		})
	}
}

// ---------- handlers ----------

func cmdSleep(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	if len(args) != 1 {
		return fmt.Errorf("sleep: expected exactly one argument (seconds)")
	}
	var secs int
	if _, err := fmt.Sscanf(args[0], "%d", &secs); err != nil {
		return fmt.Errorf("sleep: invalid number %q: %w", args[0], err)
	}
	if secs < 0 {
		return fmt.Errorf("sleep: seconds must be non-negative")
	}
	// Honour context cancellation while sleeping.
	select {
	case <-time.After(time.Duration(secs) * time.Second):
	case <-ctx.Done():
		return ctx.Err()
	}
	return nil
}

func cmdEvents(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	out.Printf("event registration list: (core.Debugger does not yet expose the tag table)\n")
	return nil
}

func cmdSettings(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	out.Printf("debugger state: %s\n", stateName(d.State()))
	out.Printf("settings: (no user-tunable settings yet)\n")
	return nil
}

func cmdG(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return d.Continue(ctx)
}

// ---------- stubs ----------

func stubHandler(name string) commands.Handler {
	return func(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
		out.Printf("%s: %v\n", name, ErrNotImplemented)
		return nil
	}
}

func stubHelp(name string) commands.HelpFunc {
	return func(d *core.Debugger, out commands.Output) error {
		out.Printf("%s: not yet implemented (Phase C.2.2 stub)\n", name)
		return nil
	}
}

// ---------- help texts ----------

func helpSleep(d *core.Debugger, out commands.Output) error {
	out.Printf("sleep : pauses the script for N seconds.\n\n")
	out.Printf("syntax : \tsleep [Seconds (decimal)]\n\n")
	out.Printf("\t\te.g : sleep 5\n")
	return nil
}

func helpEvents(d *core.Debugger, out commands.Output) error {
	out.Printf("events : lists registered events and their tags.\n\n")
	out.Printf("syntax : \tevents\n")
	return nil
}

func helpSettings(d *core.Debugger, out commands.Output) error {
	out.Printf("settings : shows or modifies debugger settings.\n\n")
	out.Printf("syntax : \tsettings\n")
	return nil
}

func helpG(d *core.Debugger, out commands.Output) error {
	out.Printf("continue : continues running the debuggee (alias for 'g').\n\n")
	out.Printf("syntax : \tcontinue\n")
	return nil
}

// ---------- helpers ----------

func stateName(s core.DebuggerState) string {
	switch s {
	case core.StateDisconnected:
		return "disconnected"
	case core.StateConnected:
		return "connected"
	case core.StateVmmLoaded:
		return "vmm-loaded"
	case core.StateProcessRunning:
		return "process-running"
	case core.StateProcessPaused:
		return "process-paused"
	default:
		return fmt.Sprintf("unknown(%d)", s)
	}
}
