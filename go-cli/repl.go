// repl.go implements the interactive read-eval-print loop for the Go CLI.
// It is the counterpart of the `while (!exit_from_debugger)` block in
// hyperdbg-cli.cpp: it shows the "hyperdbg> " prompt, reads a line, and
// dispatches it through api.Debugger.Exec.
//
// Beyond the C++ version it adds:
//   - persistent command history (~/.hyperdbg/history)
//   - "!n" expansion to re-run the nth history entry (bash-like)
//   - "history" built-in to list past commands
//   - Ctrl+C interrupts the running command without tearing down the REPL
//   - Ctrl+D (EOF) exits cleanly
//
// All command logic stays in go-libhyperdbg; the REPL only routes strings
// to dbg.Exec and reacts to the meta commands' ErrExit sentinel.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/hyperdbg/go-libhyperdbg/api"
	metacmds "github.com/hyperdbg/go-libhyperdbg/debugger/commands/meta"
)

// Repl encapsulates the interactive REPL state. It owns a *api.Debugger
// (passed in from main), a History buffer, and a handful of options. The
// zero value is not usable — use NewRepl.
type Repl struct {
	dbg         *api.Debugger
	history     *History
	prompt      string
	historyPath string
	scriptMode  bool // empty line exits (post-script behaviour, mirrors C++)
}

// ReplOption configures a Repl at construction time.
type ReplOption func(*Repl)

// WithPrompt overrides the default "hyperdbg> " prompt.
func WithPrompt(p string) ReplOption {
	return func(r *Repl) { r.prompt = p }
}

// WithHistoryPath enables persistent history at the given file path. The
// file is loaded on Run entry and saved on Run exit (best effort).
func WithHistoryPath(p string) ReplOption {
	return func(r *Repl) { r.historyPath = p }
}

// WithHistorySize sets the in-memory entry cap. Defaults to
// defaultHistoryMax when unset.
func WithHistorySize(n int) ReplOption {
	return func(r *Repl) { r.history = NewHistory(n) }
}

// WithScriptMode enables the C++ "script_mode" semantics: an empty input
// line exits the REPL. Used when the REPL is entered after a --script run
// so the user can quickly quit by pressing Enter.
func WithScriptMode(b bool) ReplOption {
	return func(r *Repl) { r.scriptMode = b }
}

// NewRepl constructs a Repl bound to dbg. Pass options to enable history
// persistence, change the prompt, or enable script-mode semantics.
func NewRepl(dbg *api.Debugger, opts ...ReplOption) *Repl {
	r := &Repl{
		dbg:     dbg,
		history: NewHistory(defaultHistoryMax),
		prompt:  "hyperdbg> ",
	}
	for _, opt := range opts {
		opt(r)
	}
	if r.history == nil {
		r.history = NewHistory(defaultHistoryMax)
	}
	return r
}

// Run enters the REPL loop. It returns nil on a clean exit (exit/quit
// command, Ctrl+D/EOF). Unexpected I/O errors are returned for the caller
// to report.
//
// SIGINT (Ctrl+C) is intercepted per-iteration: if a command is running,
// its child context is cancelled so dbg.Exec returns promptly; if the
// REPL is idle at the prompt, the signal is swallowed and a new prompt
// is shown.
func (r *Repl) Run() error {
	// Load persisted history (best effort — a missing file is fine).
	if r.historyPath != "" {
		_ = r.history.Load(r.historyPath)
	}
	defer func() {
		if r.historyPath != "" {
			_ = r.history.Save(r.historyPath)
		}
	}()

	// Intercept Ctrl+C so we can interrupt the current command without
	// killing the REPL. We register on both os.Interrupt and syscall.SIGINT
	// for cross-platform robustness; on Windows os.Interrupt maps to
	// CTRL_C_EVENT/CTRL_BREAK_EVENT.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT)
	defer signal.Stop(sigCh)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 4096), 1<<20)

	for {
		// Drain any SIGINT that arrived while idle so it does not
		// accidentally interrupt the next command.
		select {
		case <-sigCh:
			fmt.Fprint(os.Stderr, "^C\n")
		default:
		}

		fmt.Fprint(os.Stdout, r.prompt)
		if !scanner.Scan() {
			// EOF (Ctrl+D / Ctrl+Z+Enter on Windows) or read error.
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("repl input: %w", err)
			}
			fmt.Fprintln(os.Stdout)
			return nil
		}
		line := strings.TrimSpace(scanner.Text())

		// Empty line: in script mode this means exit (C++ behaviour);
		// otherwise just re-show the prompt.
		if line == "" {
			if r.scriptMode {
				return nil
			}
			continue
		}

		// "!n" — re-execute the nth history entry (1-indexed, bash-like).
		if strings.HasPrefix(line, "!") && line != "!" {
			if expanded, ok := r.expandHistory(line[1:]); ok {
				fmt.Fprintf(os.Stdout, "%s\n", expanded)
				line = expanded
			} else {
				continue
			}
		}

		// "history" — local built-in listing past commands. We handle it
		// here rather than registering a command so the CLI keeps full
		// ownership of the REPL's history surface.
		if line == "history" {
			r.history.Add(line)
			for i, e := range r.history.List() {
				fmt.Fprintf(os.Stdout, "%4d  %s\n", i+1, e)
			}
			continue
		}

		// Record the command before execution so it is persisted even
		// if dbg.Exec hangs or panics.
		r.history.Add(line)

		if err := r.execInterruptible(sigCh, line); err != nil {
			if errors.Is(err, metacmds.ErrExit) {
				fmt.Fprintln(os.Stdout, "bye.")
				return nil
			}
			fmt.Fprintf(os.Stderr, "err: %v\n", err)
		}
	}
}

// execInterruptible runs dbg.Exec in a goroutine. Because Exec is a
// synchronous IOCTL it cannot be interrupted mid-call; on SIGINT we wait
// for the in-flight command to finish, then report ^C and keep the REPL
// alive.
func (r *Repl) execInterruptible(sigCh <-chan os.Signal, line string) error {
	done := make(chan error, 1)
	go func() {
		done <- r.dbg.Exec(line)
	}()

	select {
	case err := <-done:
		return err
	case <-sigCh:
		// Exec is uninterruptible; wait for it to finish so we don't
		// leak the goroutine or interleave its output with the next prompt.
		err := <-done
		if err != nil {
			return err
		}
		fmt.Fprint(os.Stderr, "^C (command interrupted)\n")
		return nil
	}
}

// expandHistory resolves an "!n" token to the nth history entry. n must
// be a positive base-10 integer; otherwise the call prints a diagnostic
// and returns ok=false.
func (r *Repl) expandHistory(nStr string) (string, bool) {
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		fmt.Fprintf(os.Stderr, "err: bad history index %q (expected !N with N>=1)\n", nStr)
		return "", false
	}
	cmd, ok := r.history.Get(n)
	if !ok {
		fmt.Fprintf(os.Stderr, "err: no history entry %d (have %d)\n", n, r.history.Len())
		return "", false
	}
	return cmd, true
}
