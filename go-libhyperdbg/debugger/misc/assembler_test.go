package misc

import "testing"

// TestAssemble_NOP verifies "nop" assembles to 0x90.
func TestAssemble_NOP(t *testing.T) {
	a := NewAssembler()
	out, err := a.Assemble(AsmMode64, 0, "nop")
	if err != nil {
		t.Fatalf("Assemble failed: %v", err)
	}
	if len(out) != 1 || out[0] != 0x90 {
		t.Errorf("got %x, want [90]", out)
	}
}

// TestAssemble_RET verifies "ret" assembles to 0xC3.
func TestAssemble_RET(t *testing.T) {
	a := NewAssembler()
	out, err := a.Assemble(AsmMode64, 0, "ret")
	if err != nil {
		t.Fatalf("Assemble failed: %v", err)
	}
	if len(out) != 1 || out[0] != 0xC3 {
		t.Errorf("got %x, want [c3]", out)
	}
}

// TestAssemble_MOV verifies "mov rax, rbx" assembles to 48 89 D8.
func TestAssemble_MOV(t *testing.T) {
	a := NewAssembler()
	out, err := a.Assemble(AsmMode64, 0, "mov rax, rbx")
	if err != nil {
		t.Fatalf("Assemble failed: %v", err)
	}
	want := []byte{0x48, 0x89, 0xD8}
	if len(out) != len(want) {
		t.Fatalf("got %x (len %d), want %x", out, len(out), want)
	}
	for i, b := range want {
		if out[i] != b {
			t.Errorf("byte %d: got %02x, want %02x", i, out[i], b)
		}
	}
}
