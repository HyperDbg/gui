package interp

import (
	"testing"
)

// TestInterpBasicEval verifies that a bare arithmetic expression is evaluated
// and its value returned. This is the "REPL expression" path: no package
// declaration, no imports, just an expression yaegi wraps in a synthetic
// program.
func TestInterpBasicEval(t *testing.T) {
	i := NewInterpreter(nil)
	v, err := i.Eval("1+1")
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
	v, err := i.Eval("hyperdbg.Version()")
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
