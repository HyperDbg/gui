// Package hwdbgcmds implements the 2 hardware-debugging commands
// (hwdbg-commands/) from the C++ libhyperdbg: hw and hw_clk.
//
// These target the HyperDbg hardware-debugging device (Phase C.2.4). They are
// registered as stubs until the hwdbg module (go-libhyperdbg/hwdbg/) lands.
package hwdbgcmds

import (
	"context"
	"fmt"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// ErrNotImplemented is returned by stub commands.
var ErrNotImplemented = fmt.Errorf("command not yet implemented")

// RegisterAll registers hw and hw_clk as stubs.
func RegisterAll(r *commands.Registry) {
	for _, name := range []string{"hw", "hw_clk"} {
		if _, exists := r.Lookup(name); exists {
			continue
		}
		r.Register(name, commands.CommandSpec{
			Handler: stubHandler(name),
			Help:    stubHelp(name),
			Attrs:   commands.CommandAttributes{Visible: true, Type: commands.CmdHwdbg},
		})
	}
}

func stubHandler(name string) commands.Handler {
	return func(ctx context.Context, d *core.Debugger, args []string, out commands.Output) error {
		out.Printf("%s: %v\n", name, ErrNotImplemented)
		return nil
	}
}

func stubHelp(name string) commands.HelpFunc {
	return func(d *core.Debugger, out commands.Output) error {
		out.Printf("%s: not yet implemented (Phase C.2.4 stub)\n", name)
		return nil
	}
}
