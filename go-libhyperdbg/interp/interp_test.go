package interp

import (
	"context"
	"errors"
	"testing"
	"time"
)

// TestInterpBasicEval verifies that a bare arithmetic expression is evaluated
// and its value returned. This is the "REPL expression" path: no package
// declaration, no imports, just an expression yaegi wraps in a synthetic
// program.
func TestInterpBasicEval(t *testing.T) {
	i := NewInterpreter(nil)
	v, err := i.Eval(context.Background(), "1+1")
	if err != nil {
		t.Fatalf("Eval(\"1+1\"): %v", err)
	}
	got, ok := v.(int)
	if !ok {
		t.Fatalf("Eval(\"1+1\") returned %T (%v), want int", v, v)
	}
	if got != 2 {
		t.Errorf("Eval(\"1+1\") = %d, want 2", got)
	}
}

// TestInterpHyperdbgSymbol verifies that the hyperdbg package is auto-imported
// (via ImportUsed in NewInterpreter) and that hyperdbg.Version() returns the
// framework version string. This exercises the full symbol-registration path:
// registerHyperdbgSymbols -> yi.Use -> ImportUsed -> bare-identifier access.
func TestInterpHyperdbgSymbol(t *testing.T) {
	i := NewInterpreter(nil)
	v, err := i.Eval(context.Background(), "hyperdbg.Version()")
	if err != nil {
		t.Fatalf("Eval(\"hyperdbg.Version()\"): %v", err)
	}
	got, ok := v.(string)
	if !ok {
		t.Fatalf("hyperdbg.Version() returned %T (%v), want string", v, v)
	}
	if got != Version {
		t.Errorf("hyperdbg.Version() = %q, want %q", got, Version)
	}
}

// TestInterpContextCancel verifies that cancelling the Eval context promptly
// aborts an infinite interpreted loop and returns context.Canceled. The loop
// is wrapped in an immediately-invoked func literal, matching the pattern
// yaegi's own TestEvalWithContext uses (a bare `for {}` at top level is not a
// valid Go expression-statement).
//
// A 5s WithTimeout safety net guards against a regression where stop() fails
// to interrupt the loop: the test would otherwise hang until the Go test
// timeout. The manual cancel fires at 200ms, well before the safety net.
func TestInterpContextCancel(t *testing.T) {
	i := NewInterpreter(nil)

	// 5s safety net so the test fails fast instead of hanging.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Manual cancel after 200ms. ctx.Err() after this returns
	// context.Canceled (manual cancel wins over the deadline).
	go func() {
		time.Sleep(200 * time.Millisecond)
		cancel()
	}()

	start := time.Now()
	_, err := i.Eval(ctx, `(func() { for {} })()`)
	elapsed := time.Since(start)

	if !errors.Is(err, context.Canceled) {
		t.Errorf("Eval error = %v, want context.Canceled", err)
	}
	// The manual cancel fires at 200ms; allow generous slack for scheduling
	// and yaegi's stop() propagation, but well under the 5s safety net.
	if elapsed > 3*time.Second {
		t.Errorf("Eval did not return promptly after cancel: %v", elapsed)
	}
}
