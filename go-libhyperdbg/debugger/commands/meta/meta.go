// Package metacmds implements the 23 meta commands (meta-commands/) from the
// C++ libhyperdbg. Each command is a free function taking *core.Debugger so
// the commands package stays independent of the api layer.
//
// Commands implemented with real behaviour:
//
//	.connect / connect       — open the local HyperDbg device
//	load                     — install + start the VMM driver
//	unload                   — stop + remove the VMM driver (debugging group, also here)
//	g / go                   — continue execution
//	pause                    — halt execution
//	.logopen / logopen       — open log file
//	.logclose / logclose     — close log file
//	.start path              — launch debuggee (start command, meta group)
//	status                   — print debugger state
//	cls / clear              — clear screen (CLI side; emits ANSI clear)
//	help / .help / !help     — list commands
//	exit / .exit             — request shutdown (returns ErrExit)
//
// Commands registered as stubs (return ErrNotImplemented) pending their full
// implementation in later phases:
//
//	attach, debug, detach, disconnect, dump, formats, kill, listen, pagein,
//	pe, process, restart, script, switch, sym, sympath, thread
package metacmds

import (
	"context"
	"errors"
	"fmt"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// ErrExit signals the CLI loop to terminate. Returned by the exit command.
var ErrExit = errors.New("exit requested")

// ErrNotImplemented is returned by stub commands pending full implementation.
var ErrNotImplemented = errors.New("command not yet implemented")

// RegisterAll registers every meta command (and its aliases) into r.
// The order of registration does not matter; aliases are resolved by
// RegisterAlias after the canonical name is in place.
func RegisterAll(r *commands.Registry) {
	// ---- .connect / connect ----
	r.Register(".connect", commands.CommandSpec{
		Handler: cmdConnect,
		Help:    helpConnect,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias("connect", ".connect")

	// ---- load ----
	r.Register("load", commands.CommandSpec{
		Handler: cmdLoad,
		Help:    helpLoad,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})

	// ---- unload (debugging group in C++ but used here for symmetry) ----
	r.Register("unload", commands.CommandSpec{
		Handler: cmdUnload,
		Help:    helpUnload,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})

	// ---- g / go ----
	r.Register("g", commands.CommandSpec{
		Handler: cmdG,
		Help:    helpG,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})
	r.RegisterAlias("go", "g")

	// ---- pause ----
	r.Register("pause", commands.CommandSpec{
		Handler: cmdPause,
		Help:    helpPause,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdDebugging},
	})

	// ---- .logopen / logopen ----
	r.Register(".logopen", commands.CommandSpec{
		Handler: cmdLogopen,
		Help:    helpLogopen,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias("logopen", ".logopen")

	// ---- .logclose / logclose ----
	r.Register(".logclose", commands.CommandSpec{
		Handler: cmdLogclose,
		Help:    helpLogclose,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias("logclose", ".logclose")

	// ---- .start path (launch debuggee) ----
	r.Register(".start", commands.CommandSpec{
		Handler: cmdStart,
		Help:    helpStart,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})

	// ---- status ----
	r.Register("status", commands.CommandSpec{
		Handler: cmdStatus,
		Help:    helpStatus,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})

	// ---- cls / clear ----
	r.Register("cls", commands.CommandSpec{
		Handler: cmdCls,
		Help:    helpCls,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias("clear", "cls")

	// ---- help / .help / !help ----
	// The help handler needs to enumerate all registered commands, so it is
	// built by the Registry itself (NewHelpHandler) rather than implemented
	// here.
	r.Register("help", commands.CommandSpec{
		Handler: r.NewHelpHandler(),
		Help:    helpHelp,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias(".help", "help")
	r.RegisterAlias("!help", "help")
	r.RegisterAlias(".hh", "help")

	// ---- exit / .exit ----
	r.Register("exit", commands.CommandSpec{
		Handler: cmdExit,
		Help:    helpExit,
		Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
	})
	r.RegisterAlias(".exit", "exit")
	r.RegisterAlias("quit", "exit")

	// ---- stubs (visible so users can discover them; behaviour pending) ----
	stubs := []string{
		"attach", "debug", "detach", "disconnect", "dump", "formats",
		"kill", "listen", "pagein", "pe", "process", "restart",
		"script", "switch", "sym", "sympath", "thread",
	}
	for _, name := range stubs {
		r.Register(name, commands.CommandSpec{
			Handler: stubHandler(name),
			Help:    stubHelp(name),
			Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdMeta},
		})
	}
}

// ---------- handlers ----------

func cmdConnect(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	if len(args) == 0 || args[0] == "local" {
		return d.Connect(ctx, "local")
	}
	// Remote: .connect <ip> <port> — not yet supported (no TCP transport).
	if len(args) >= 2 {
		return fmt.Errorf(".connect: remote debugging not yet implemented (ip=%s port=%s)", args[0], args[1])
	}
	return fmt.Errorf(".connect: expected 'local' or '<ip> <port>'")
}

func cmdLoad(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	if len(args) != 1 {
		return fmt.Errorf("load: expected exactly one argument (driver path), got %d", len(args))
	}
	return d.LoadVMM(ctx, args[0])
}

func cmdUnload(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return d.UnloadVMM(ctx)
}

func cmdG(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return d.Continue(ctx)
}

func cmdPause(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return d.Pause(ctx)
}

func cmdLogopen(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	if len(args) != 1 {
		return fmt.Errorf(".logopen: expected exactly one argument (file path)")
	}
	return d.LogOpen(args[0])
}

func cmdLogclose(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return d.LogClose()
}

func cmdStart(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	// ".start path <exe>" — args[0] == "path", args[1] == exe path
	if len(args) >= 2 && args[0] == "path" {
		proc, err := d.StartProcess(ctx, args[1])
		if err != nil {
			return err
		}
		out.Printf("started pid=%d\n", proc.Pid)
		return nil
	}
	return fmt.Errorf(".start: expected 'path <exe>'")
}

func cmdStatus(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	out.Printf("debugger state: %s\n", stateName(d.State()))
	return nil
}

func cmdCls(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	// ANSI clear-screen + cursor home. The CLI passes a real terminal; GUI/MCP
	// implementations can interpret or ignore this.
	out.Write([]byte("\x1b[2J\x1b[H"))
	return nil
}

func cmdExit(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
	return ErrExit
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
		out.Printf("%s: not yet implemented (Phase C.2.1 stub)\n", name)
		return nil
	}
}

// ---------- help texts ----------

func helpConnect(d *core.Debugger, out commands.Output) error {
	out.Printf(".connect : connects to a remote or local machine to start debugging.\n\n")
	out.Printf("syntax : \t.connect [local]\n")
	out.Printf("syntax : \t.connect [Ip (string)] [Port (decimal)]\n\n")
	out.Printf("\t\te.g : .connect local\n")
	out.Printf("\t\te.g : .connect 192.168.1.5 50000\n")
	return nil
}

func helpLoad(d *core.Debugger, out commands.Output) error {
	out.Printf("load : installs and starts the VMM driver.\n\n")
	out.Printf("syntax : \tload [DriverPath (string)]\n\n")
	out.Printf("\t\te.g : load vmm\n")
	out.Printf("\t\te.g : load Debug\\hyperkd.sys\n")
	return nil
}

func helpUnload(d *core.Debugger, out commands.Output) error {
	out.Printf("unload : stops and removes the VMM driver.\n\n")
	out.Printf("syntax : \tunload vmm\n")
	return nil
}

func helpG(d *core.Debugger, out commands.Output) error {
	out.Printf("g : continues running the debuggee.\n\n")
	out.Printf("syntax : \tg\n")
	return nil
}

func helpPause(d *core.Debugger, out commands.Output) error {
	out.Printf("pause : halts the debuggee.\n\n")
	out.Printf("syntax : \tpause\n")
	return nil
}

func helpLogopen(d *core.Debugger, out commands.Output) error {
	out.Printf(".logopen : saves commands and results in a file.\n\n")
	out.Printf("syntax : \t.logopen [FilePath (string)]\n\n")
	out.Printf("\t\te.g : .logopen log.txt\n")
	return nil
}

func helpLogclose(d *core.Debugger, out commands.Output) error {
	out.Printf(".logclose : closes the open log file.\n\n")
	out.Printf("syntax : \t.logclose\n")
	return nil
}

func helpStart(d *core.Debugger, out commands.Output) error {
	out.Printf(".start : launches a process for debugging.\n\n")
	out.Printf("syntax : \t.start path [FilePath (string)]\n\n")
	out.Printf("\t\te.g : .start path notepad.exe\n")
	return nil
}

func helpStatus(d *core.Debugger, out commands.Output) error {
	out.Printf("status : prints the current debugger state.\n\n")
	out.Printf("syntax : \tstatus\n")
	return nil
}

func helpCls(d *core.Debugger, out commands.Output) error {
	out.Printf("cls : clears the screen.\n\n")
	out.Printf("syntax : \tcls\n")
	return nil
}

func helpHelp(d *core.Debugger, out commands.Output) error {
	out.Printf("help : lists commands or shows help for a specific command.\n\n")
	out.Printf("syntax : \thelp [Command (string)]\n")
	return nil
}

func helpExit(d *core.Debugger, out commands.Output) error {
	out.Printf("exit : exits the HyperDbg CLI.\n\n")
	out.Printf("syntax : \texit\n")
	return nil
}

// ---------- helpers ----------

// stateName returns a human-readable name for a debugger state.
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
