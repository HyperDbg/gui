package misc

import "testing"

// TestDisassemble_NOP verifies a single-byte NOP (0x90) decodes to "nop".
func TestDisassemble_NOP(t *testing.T) {
	d := NewDisassembler()
	r, err := d.Disassemble(ModeLong64, 0, []byte{0x90})
	if err != nil {
		t.Fatalf("Disassemble failed: %v", err)
	}
	if r.Length != 1 {
		t.Errorf("length = %d, want 1", r.Length)
	}
	if r.Text != "nop" {
		t.Errorf("text = %q, want \"nop\"", r.Text)
	}
}

// TestDisassemble_MOV verifies a 3-byte MOV decodes correctly.
func TestDisassemble_MOV(t *testing.T) {
	d := NewDisassembler()
	// 48 89 D8 = mov rax, rbx
	r, err := d.Disassemble(ModeLong64, 0, []byte{0x48, 0x89, 0xD8})
	if err != nil {
		t.Fatalf("Disassemble failed: %v", err)
	}
	if r.Length != 3 {
		t.Errorf("length = %d, want 3", r.Length)
	}
	if r.Text != "mov rax, rbx" {
		t.Errorf("text = %q, want \"mov rax, rbx\"", r.Text)
	}
}

// TestDisassembleRange verifies multi-instruction decoding.
func TestDisassembleRange(t *testing.T) {
	d := NewDisassembler()
	// nop; nop; ret
	code := []byte{0x90, 0x90, 0xC3}
	results, err := d.DisassembleRange(ModeLong64, 0x1000, code)
	if err != nil {
		t.Fatalf("DisassembleRange failed: %v", err)
	}
	if len(results) != 3 {
		t.Fatalf("got %d results, want 3", len(results))
	}
	if results[0].Text != "nop" || results[1].Text != "nop" || results[2].Text != "ret" {
		t.Errorf("texts = %q, %q, %q; want nop, nop, ret",
			results[0].Text, results[1].Text, results[2].Text)
	}
	if results[2].Runtime != 0x1002 {
		t.Errorf("third instr runtime = %#x, want 0x1002", results[2].Runtime)
	}
}
