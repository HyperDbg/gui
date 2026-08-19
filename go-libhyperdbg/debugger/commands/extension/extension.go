// Package extcmds implements the 35 extension commands (extension-commands/)
// from the C++ libhyperdbg. These are the "!" prefixed commands.
//
// Most extension commands directly drive IOCTLs against the VMM driver
// (e.g. !epthook, !monitor, !cpuid, !msrread). They are registered as stubs
// here; Phase C.2.3 fills them in once the comm/types IOCTL packing helpers
// are factored out of core.Debugger.EptHook.
//
// Commands that overlap with the Go AST hook path (!epthook, !epthook2) will
// be implemented first because find-oep.go and the parity tests need them.
package extcmds

import (
	"context"
	"fmt"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// ErrNotImplemented is returned by stub commands.
var ErrNotImplemented = fmt.Errorf("command not yet implemented")

// RegisterAll registers all 35 extension commands as stubs. The "!epthook"
// family is wired to the real EptHook path once the symbol resolver lands.
func RegisterAll(r *commands.Registry) {
	stubs := []string{
		"!apic", "!cpuid", "!crwrite", "!dr", "!epthook", "!epthook2",
		"!exception", "!hide", "!idt", "!interrupt", "!ioapic",
		"!ioin", "!ioout", "!lbr", "!lbrdump", "!measure", "!mode",
		"!monitor", "!msrread", "!msrwrite", "!pa2va", "!pcicam",
		"!pcitree", "!pmc", "!pt", "!pte", "!rev", "!smi",
		"!syscall", "!sysret", "!trace", "!track", "!tsc", "!unhide",
		"!va2pa", "!vmcall", "!xsetbv",
	}
	for _, name := range stubs {
		if _, exists := r.Lookup(name); exists {
			continue
		}
		r.Register(name, commands.CommandSpec{
			Handler: stubHandler(name),
			Help:    stubHelp(name),
			Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdExtension},
		})
	}
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
		out.Printf("%s: not yet implemented (Phase C.2.3 stub)\n", name)
		return nil
	}
}
