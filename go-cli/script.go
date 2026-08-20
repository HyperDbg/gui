// script.go implements the script-execution mode for the Go CLI. It is
// the counterpart of HyperDbgScriptReadFileAndExecuteCommand in the C++
// libhyperdbg: a file of HyperDbg CLI commands (one per line) is read
// and each non-empty line is fed to api.Debugger.Exec.
//
// Two entry points exist:
//
//   - Run(.ds):    line-by-line command execution. The file extension is
//                  not enforced; ".ds" is the conventional name in the
//                  HyperDbg docs but any text file works.
//   - RunGoScript: placeholder for executing .go scripts via yaegi
//                  (Phase C.4.3). It returns ErrYaegiNotIntegrated until
//                  the interpreter is wired up.
//
// After a successful .ds run the runner prints "Script finished. Press
// Enter to exit..." and blocks for a single line of input, mirroring the
// post-script getchar() pause in hyperdbg-cli.cpp so the user can read
// the output before the terminal closes.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/hyperdbg/go-libhyperdbg/api"
	metacmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/meta"
)

// ErrYaegiNotIntegrated is returned by RunGoScript until the yaegi
// interpreter is wired up in Phase C.4.3. Callers can errors.Is this
// value to distinguish "feature missing" from a real failure.
var ErrYaegiNotIntegrated = errors.New("yaegi Go script execution not yet implemented (Phase C.4.3)")

// ScriptRunner executes HyperDbg command scripts. It holds a reference to
// the api.Debugger so callers (main) don't have to thread it through
// alongside the script path.
type ScriptRunner struct {
	dbg *api.Debugger
}

// NewScriptRunner constructs a ScriptRunner bound to dbg.
func NewScriptRunner(dbg *api.Debugger) *ScriptRunner {
	return &ScriptRunner{dbg: dbg}
}

// Run reads scriptPath line by line and dispatches each non-empty,
// non-comment line to dbg.Exec. An `exit`/`quit` command aborts the run
// early (matching the C++ behaviour where CommandScriptRunCommand calls
// exit(0) when the interpreter returns 1).
//
// Lines beginning with "//" or "#" are treated as comments and skipped,
// which is a small extension over the C++ version (the C++ parser does
// not support comments in command scripts). Multi-line commands (the
// `{ ... }` blocks the C++ interpreter supports) are not handled here —
// the Go command registry is single-line.
//
// SIGINT (Ctrl+C) aborts the run gracefully: a non-blocking check on
// sigCh between lines breaks the loop and skips the post-script pause.
// Any other error (file open, read failure) is returned to the caller.
func (r *ScriptRunner) Run(scriptPath string) error {
	f, err := os.Open(scriptPath)
	if err != nil {
		return fmt.Errorf("open script %q: %w", scriptPath, err)
	}
	defer f.Close()

	// Allow Ctrl+C to abort the script. We register SIGINT here (not in
	// main) so the REPL can install its own per-command handler without
	// conflicting with us.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	fmt.Fprintf(os.Stdout, "Running script: %s\n", scriptPath)

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 4096), 1<<20)
	lineNo := 0
	aborted := false

	for sc.Scan() {
		lineNo++
		// Honor Ctrl+C between lines.
		select {
		case <-sigCh:
			aborted = true
			break
		default:
		}
		if aborted {
			break
		}

		raw := sc.Text()
		line := strings.TrimSpace(raw)
		if line == "" || strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}

		// Echo the command so the user sees what is running, the same way
		// CommandScriptRunCommand prints the line via ShowMessages.
		fmt.Fprintf(os.Stdout, "hyperdbg> %s\n", line)

		if err := r.dbg.Exec(line); err != nil {
			if errors.Is(err, metacmds.ErrExit) {
				fmt.Fprintf(os.Stdout, "[exit requested at line %d]\n", lineNo)
				break
			}
			fmt.Fprintf(os.Stderr, "err (line %d): %v\n", lineNo, err)
		}
	}
	if !aborted {
		if err := sc.Err(); err != nil {
			return fmt.Errorf("read script %q: %w", scriptPath, err)
		}
	} else {
		fmt.Fprint(os.Stderr, "\n[script aborted by Ctrl+C]\n")
	}

	// Post-script pause: let the user read the output before the program
	// exits (mirrors the C++ hyperdbg-cli getchar() pattern). Skipped on
	// abort so Ctrl+C exits immediately.
	if !aborted {
		r.waitEnter()
	}
	return nil
}

// RunGoScript is the placeholder entry point for executing .go scripts
// through yaegi (Phase C.4.3). It currently always returns
// ErrYaegiNotIntegrated; once yaegi is wired up it will parse the file,
// export the api package symbols, and run main() in the script.
func (r *ScriptRunner) RunGoScript(scriptPath string) error {
	return fmt.Errorf("RunGoScript(%q): %w", scriptPath, ErrYaegiNotIntegrated)
}

// waitEnter blocks until the user presses Enter (or sends EOF). It is
// the Go equivalent of the C++ `getchar()` call that keeps the console
// window open after a script finishes. Errors from stdin are swallowed
// — the pause is best-effort and must never mask a real script error.
func (r *ScriptRunner) waitEnter() {
	fmt.Fprint(os.Stdout, "Script finished. Press Enter to exit...")
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64), 1<<10)
	_ = sc.Scan()
	fmt.Fprintln(os.Stdout)
}
