// Package commands implements the HyperDbg command set (97 commands across
// meta/debugging/extension/hwdbg groups) and the registry that dispatches
// command strings to their handlers.
//
// Design (per the API spec):
//   - Commands are methods on *core.Debugger (or free functions taking one),
//     not package-level globals.
//   - Output goes through the Output interface so CLI/GUI/MCP can all consume
//     the same command implementation.
//   - Registry is owned by the caller (api.Debugger typically constructs one);
//     no package-level mutable state.
//   - Aliases register the same CommandSpec under multiple names (e.g.
//     "g"/"go", ".connect"/"connect").
package commands

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// CommandType groups commands for help/parity purposes (matches the C++
// DEBUGGER_COMMAND_* attributes in interpreter.cpp).
type CommandType int

const (
	CmdMeta CommandType = iota
	CmdDebugging
	CmdExtension
	CmdHwdbg
)

// CommandAttributes mirrors COMMAND_PROPERTIES attributes from the C++
// interpreter. Visible=false hides the command from `help` output (used by
// internal/legacy aliases).
type CommandAttributes struct {
	Visible bool
	Type    CommandType
}

// Output abstracts command output. CLI passes os.Stdout, GUI passes a widget
// writer, MCP passes a JSON channel. Implementations must be goroutine-safe
// if the Registry is used concurrently.
type Output interface {
	Write(p []byte) (int, error)
	Printf(format string, args ...any) error
}

// Handler is the signature every command handler implements. args is the
// command line split on whitespace with the command name removed (argv[1:]).
type Handler func(ctx context.Context, d *core.Debugger, args []string, out Output) error

// HelpFunc prints the command's help text to out.
type HelpFunc func(d *core.Debugger, out Output) error

// CommandSpec mirrors COMMAND_PROPERTIES from the C++ interpreter.
type CommandSpec struct {
	Handler Handler
	Help    HelpFunc
	Attrs   CommandAttributes
}

// Registry holds the registered commands. The zero value is not usable; use
// NewRegistry.
type Registry struct {
	cmds   map[string]CommandSpec
	output Output
}

// NewRegistry creates an empty Registry with the given output sink.
func NewRegistry(out Output) *Registry {
	return &Registry{
		cmds:   make(map[string]CommandSpec),
		output: out,
	}
}

// Register adds a command under the given name. If the name already exists it
// is overwritten (matching the C++ behaviour where later registrations win).
func (r *Registry) Register(name string, spec CommandSpec) {
	r.cmds[name] = spec
}

// RegisterAlias registers an alias that points to the same spec as target.
// If target is not registered, the alias is silently dropped (matching the
// C++ behaviour where aliases are only added after the canonical name).
func (r *Registry) RegisterAlias(alias, target string) {
	spec, ok := r.cmds[target]
	if !ok {
		return
	}
	r.cmds[alias] = spec
}

// Lookup returns the CommandSpec for the given name and whether it exists.
func (r *Registry) Lookup(name string) (CommandSpec, bool) {
	spec, ok := r.cmds[name]
	return spec, ok
}

// Names returns all registered command names. The order is unspecified; callers
// that need a stable order should sort the result.
func (r *Registry) Names() []string {
	names := make([]string, 0, len(r.cmds))
	for n := range r.cmds {
		names = append(names, n)
	}
	return names
}

// Exec parses cmdLine into a command name + args, looks up the handler, and
// invokes it. An empty line is a no-op. Unknown commands return an error
// matching the C++ "err_cmd_not_found" behaviour.
//
// The command name is matched case-sensitively (HyperDbg commands are
// case-sensitive). Commands prefixed with '.' (e.g. ".connect") and '!'
// (e.g. "!epthook") are meta/extension commands and are matched verbatim.
func (r *Registry) Exec(ctx context.Context, d *core.Debugger, cmdLine string) error {
	cmdLine = strings.TrimSpace(cmdLine)
	if cmdLine == "" {
		return nil
	}
	fields := splitFields(cmdLine)
	name := fields[0]
	args := fields[1:]

	spec, ok := r.cmds[name]
	if !ok {
		return fmt.Errorf("err_cmd_not_found (%q)", name)
	}
	if spec.Handler == nil {
		return fmt.Errorf("err_cmd_no_handler (%q)", name)
	}
	return spec.Handler(ctx, d, args, r.output)
}

// splitFields splits cmdLine on runs of whitespace, respecting double-quoted
// substrings (quotes are stripped). This matches the C++ CommandTokenize
// behaviour used by the interpreter.
func splitFields(s string) []string {
	var fields []string
	var b strings.Builder
	inQuote := false
	for _, r := range s {
		switch {
		case r == '"':
			inQuote = !inQuote
		case (r == ' ' || r == '\t') && !inQuote:
			if b.Len() > 0 {
				fields = append(fields, b.String())
				b.Reset()
			}
		default:
			b.WriteRune(r)
		}
	}
	if b.Len() > 0 {
		fields = append(fields, b.String())
	}
	return fields
}

// NewHelpHandler returns a Handler that lists visible commands (no args) or
// prints the help text for a specific command (one arg). It needs the Registry
// reference so it can enumerate names — handlers normally do not have it, so
// this factory exists specifically for the help command.
func (r *Registry) NewHelpHandler() Handler {
	return func(ctx context.Context, d *core.Debugger, args []string, out Output) error {
		if len(args) == 0 {
			// List all visible commands, grouped by type.
			byType := map[CommandType][]string{}
			for name, spec := range r.cmds {
				if !spec.Attrs.Visible {
					continue
				}
				byType[spec.Attrs.Type] = append(byType[spec.Attrs.Type], name)
			}
			out.Printf("HyperDbg commands:\n\n")
			typeNames := map[CommandType]string{
				CmdMeta:      "Meta",
				CmdDebugging: "Debugging",
				CmdExtension: "Extension",
				CmdHwdbg:     "Hwdbg",
			}
			for _, t := range []CommandType{CmdMeta, CmdDebugging, CmdExtension, CmdHwdbg} {
				names := byType[t]
				if len(names) == 0 {
					continue
				}
				sort.Strings(names)
				out.Printf("  [%s]\n", typeNames[t])
				for _, n := range names {
					out.Printf("    %s\n", n)
				}
				out.Printf("\n")
			}
			out.Printf("Type 'help <command>' for detailed help.\n")
			return nil
		}
		// Help for a specific command.
		name := args[0]
		spec, ok := r.cmds[name]
		if !ok {
			return fmt.Errorf("err_cmd_not_found (%q)", name)
		}
		if spec.Help == nil {
			out.Printf("%s: no help available\n", name)
			return nil
		}
		return spec.Help(d, out)
	}
}
