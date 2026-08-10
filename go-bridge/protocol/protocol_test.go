package protocol

import (
	"bytes"
	"testing"
)

// TestEncodeDecodeRoundTrip verifies that every node type survives an
// Encode → Decode round-trip with byte-for-byte equality.
func TestEncodeDecodeRoundTrip(t *testing.T) {
	// Build a representative tree exercising most opcodes.
	ctxID := uint16(0) // placeholder; real IDs assigned by encoder
	_ = ctxID
	root := NewFuncLit(
		NewBlock(
			// ret := ctx.StackReadQword(0) & 0xFFFFFFFF
			NewAssign(AssignAssign,
				NewIdent(1), // "ret"
				NewBinary(BinOpAnd,
					NewCall(FuncStackReadQword, NewLiteralUint32(0)),
					NewLiteralUint64(0xFFFFFFFF),
				),
			),
			// ctx.Printf("RAH ret=%x", ret)
			NewCall(FuncPrintf,
				NewLiteralString("RAH ret=%x"),
				NewIdent(1), // "ret"
			),
			// if ret < 0x10000000 { ctx.Break() }
			NewIf(
				NewBinary(BinOpLss, NewIdent(1), NewLiteralUint64(0x10000000)),
				NewBlock(NewCall(FuncBreak)),
				nil, // no else
			),
		),
	)

	enc := NewEncoder()
	// Pre-register strings: 0=ctx, 1=ret
	enc.AddString("ctx")
	enc.AddString("ret")
	payload, err := enc.Encode(root)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}

	dec := NewDecoder(payload)
	got, err := dec.Decode()
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}

	// Verify string table
	strs := dec.Strings()
	if len(strs) < 2 {
		t.Fatalf("expected at least 2 strings, got %d", len(strs))
	}
	if strs[0] != "ctx" || strs[1] != "ret" {
		t.Errorf("string table = %v, want [ctx ret ...]", strs)
	}

	// Verify the tree structure matches
	if !nodesEqual(root, got) {
		t.Errorf("round-trip mismatch:\n--- want ---\n%s\n--- got ---\n%s", root.String(), got.String())
	}
}

// TestEncodeDecode_AllNodeTypes verifies each opcode individually.
func TestEncodeDecode_AllNodeTypes(t *testing.T) {
	cases := []struct {
		name string
		node *Node
	}{
		{"Nil", nil},
		{"LiteralBool", NewLiteralBool(true)},
		{"LiteralUint8", NewLiteral(KindUint8, []byte{42})},
		{"LiteralUint32", NewLiteralUint32(0xDEADBEEF)},
		{"LiteralUint64", NewLiteralUint64(0x0123456789ABCDEF)},
		{"LiteralString", NewLiteralString("hello")},
		{"Ident", NewIdent(5)},
		{"BinaryExpr", NewBinary(BinOpAdd, NewLiteralUint64(1), NewLiteralUint64(2))},
		{"UnaryExpr", NewUnary(UnOpNot, NewLiteralBool(false))},
		{"CallExpr", NewCall(FuncGetPid)},
		{"SelectorExpr", NewSelector(NewIdent(0), 3)},
		{"AssignStmt", NewAssign(AssignAdd, NewIdent(1), NewLiteralUint64(1))},
		{"IfStmt", NewIf(NewLiteralBool(true), NewBlock(), NewBlock())},
		{"ForStmt", NewFor(nil, NewLiteralBool(false), nil, NewBlock())},
		{"BlockStmt", NewBlock(NewCall(FuncBreak), NewCall(FuncContinue))},
		{"ReturnStmt", NewReturn(NewLiteralUint64(0))},
		{"FuncLit", NewFuncLit(NewBlock())},
		{"DeclStmt", NewDecl(KindUint64, 2, NewLiteralUint64(0))},
		{"IndexExpr", &Node{Opcode: OpIndexExpr, Children: []*Node{NewIdent(0), NewLiteralUint32(0)}}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			enc := NewEncoder()
			payload, err := enc.Encode(c.node)
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}
			dec := NewDecoder(payload)
			got, err := dec.Decode()
			if err != nil {
				t.Fatalf("Decode failed: %v", err)
			}
			if !nodesEqual(c.node, got) {
				t.Errorf("mismatch:\n--- want ---\n%s\n--- got ---\n%s", c.node.String(), got.String())
			}
		})
	}
}

// TestDecode_TruncatedBuffer verifies malformed buffers produce DecodeError,
// not panics.
func TestDecode_TruncatedBuffer(t *testing.T) {
	cases := [][]byte{
		{},                // too short for root opcode
		{0xFF},            // unknown opcode 0xFF
		{OpLiteral},       // literal without kind
		{OpLiteral, 0x01}, // uint8 literal without value byte
	}
	for i, buf := range cases {
		dec := NewDecoder(buf)
		_, err := dec.Decode()
		if err == nil {
			t.Errorf("case %d: expected error for truncated buffer %v", i, buf)
		}
	}
}

// TestWhitelistFuncs verifies all whitelist methods are registered.
func TestWhitelistFuncs(t *testing.T) {
	expected := []string{
		"StackReadQword", "StackReadDword", "Reg", "SetReg",
		"ReadMem", "ReadMemQword", "Printf", "Break",
		"Continue", "GetPid", "GetTid", "GetIP",
	}
	for _, name := range expected {
		if _, ok := WhitelistFuncs[name]; !ok {
			t.Errorf("whitelist method %q not in WhitelistFuncs map", name)
		}
	}
	if len(WhitelistFuncs) != len(expected) {
		t.Errorf("WhitelistFuncs has %d entries, expected %d", len(WhitelistFuncs), len(expected))
	}
}

// nodesEqual compares two Node trees structurally.
func nodesEqual(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Opcode != b.Opcode || a.Kind != b.Kind || a.Op != b.Op ||
		a.FuncID != b.FuncID || a.NameID != b.NameID || a.FieldID != b.FieldID {
		return false
	}
	if !bytes.Equal(a.Value, b.Value) {
		return false
	}
	if len(a.Children) != len(b.Children) {
		return false
	}
	for i := range a.Children {
		if !nodesEqual(a.Children[i], b.Children[i]) {
			return false
		}
	}
	return true
}
