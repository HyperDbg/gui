package astencoder

import (
	"encoding/binary"
	"testing"

	"github.com/hyperdbg/go-bridge/protocol"
)

// TestEncode_SimpleCallback verifies that a typical hook callback compiles
// to a valid binary AST payload that can be decoded back.
func TestEncode_SimpleCallback(t *testing.T) {
	src := `package hook
func hook(ctx *HookCtx) {
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("RAH ret=%x", ret)
	if ret < 0x10000000 {
		ctx.Break()
	}
}`
	payload, err := Encode(src)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if len(payload) == 0 {
		t.Fatal("empty payload")
	}
	// Decode and verify structure
	dec := protocol.NewDecoder(payload)
	root, err := dec.Decode()
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if root == nil {
		t.Fatal("decoded root is nil")
	}
	if root.Opcode != protocol.OpFuncLit {
		t.Errorf("root opcode = 0x%02x, want OpFuncLit", root.Opcode)
	}
	// FuncLit → Block → 3 statements (assign, call, if)
	if len(root.Children) != 1 {
		t.Fatalf("FuncLit children = %d, want 1 (body)", len(root.Children))
	}
	body := root.Children[0]
	if body.Opcode != protocol.OpBlockStmt {
		t.Errorf("body opcode = 0x%02x, want OpBlockStmt", body.Opcode)
	}
	if len(body.Children) != 3 {
		t.Errorf("body statements = %d, want 3", len(body.Children))
	}
	// Verify string table contains "ret" (the local variable).
	// Note: "ctx" is NOT in the table because ctx.Method() calls are encoded
	// directly as CallExpr with func_id, skipping the ctx selector.
	strs := dec.Strings()
	foundRet := false
	for _, s := range strs {
		if s == "ret" {
			foundRet = true
		}
	}
	if !foundRet {
		t.Error("string table missing 'ret'")
	}
}

// TestEncode_PrintfCall verifies that ctx.Printf is encoded as a CallExpr
// with the correct func_id.
func TestEncode_PrintfCall(t *testing.T) {
	src := `package hook
func hook(ctx *HookCtx) {
	ctx.Printf("test %x", 42)
}`
	payload, err := Encode(src)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	dec := protocol.NewDecoder(payload)
	root, _ := dec.Decode()
	body := root.Children[0] // BlockStmt
	call := body.Children[0] // first statement = CallExpr
	if call.Opcode != protocol.OpCallExpr {
		t.Errorf("expected OpCallExpr, got 0x%02x", call.Opcode)
	}
	if call.FuncID != protocol.FuncPrintf {
		t.Errorf("func_id = 0x%04x, want FuncPrintf", call.FuncID)
	}
	if len(call.Children) != 2 {
		t.Errorf("Printf args = %d, want 2 (fmt + value)", len(call.Children))
	}
}

// TestEncode_IfElse verifies if/else encoding.
func TestEncode_IfElse(t *testing.T) {
	src := `package hook
func hook(ctx *HookCtx) {
	x := ctx.GetIP()
	if x > 0x1000 {
		ctx.Break()
	} else {
		ctx.Continue()
	}
}`
	payload, err := Encode(src)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	dec := protocol.NewDecoder(payload)
	root, _ := dec.Decode()
	body := root.Children[0]
	// stmts: [assign, if]
	ifStmt := body.Children[1]
	if ifStmt.Opcode != protocol.OpIfStmt {
		t.Fatalf("expected OpIfStmt, got 0x%02x", ifStmt.Opcode)
	}
	if len(ifStmt.Children) != 3 {
		t.Fatalf("IfStmt children = %d, want 3 (cond, then, else)", len(ifStmt.Children))
	}
	if ifStmt.Children[2] == nil {
		t.Error("else branch is nil, expected a block")
	}
}

// TestEncode_ForLoop verifies three-clause for loop encoding.
func TestEncode_ForLoop(t *testing.T) {
	src := `package hook
func hook(ctx *HookCtx) {
	var i uint64 = 0
	for i = 0; i < 10; i = i + 1 {
		ctx.Printf("i=%d", i)
	}
}`
	payload, err := Encode(src)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	dec := protocol.NewDecoder(payload)
	root, _ := dec.Decode()
	body := root.Children[0]
	// stmts: [var_decl, for]
	forStmt := body.Children[1]
	if forStmt.Opcode != protocol.OpForStmt {
		t.Fatalf("expected OpForStmt, got 0x%02x", forStmt.Opcode)
	}
	if len(forStmt.Children) != 4 {
		t.Fatalf("ForStmt children = %d, want 4 (init, cond, post, body)", len(forStmt.Children))
	}
}

// TestEncode_RejectsGoroutine verifies that unsupported constructs are
// rejected before encoding.
func TestEncode_RejectsGoroutine(t *testing.T) {
	src := `package hook
func hook(ctx *HookCtx) {
	go ctx.Break()
}`
	_, err := Encode(src)
	if err == nil {
		t.Fatal("expected error for goroutine, got nil")
	}
}

// TestEncode_BinaryOps verifies all binary operators map correctly.
func TestEncode_BinaryOps(t *testing.T) {
	cases := []struct {
		expr string
		op   byte
	}{
		{"1 + 2", protocol.BinOpAdd},
		{"1 - 2", protocol.BinOpSub},
		{"1 * 2", protocol.BinOpMul},
		{"1 / 2", protocol.BinOpQuo},
		{"1 % 2", protocol.BinOpRem},
		{"1 & 2", protocol.BinOpAnd},
		{"1 | 2", protocol.BinOpOr},
		{"1 ^ 2", protocol.BinOpXor},
		{"1 << 2", protocol.BinOpShl},
		{"1 >> 2", protocol.BinOpShr},
		{"1 == 2", protocol.BinOpEql},
		{"1 != 2", protocol.BinOpNeq},
		{"1 < 2", protocol.BinOpLss},
		{"1 > 2", protocol.BinOpGtr},
		{"1 <= 2", protocol.BinOpLeq},
		{"1 >= 2", protocol.BinOpGeq},
	}
	for _, c := range cases {
		src := "package hook\nfunc hook(ctx *HookCtx) { _ = " + c.expr + " }"
		payload, err := Encode(src)
		if err != nil {
			t.Errorf("Encode(%q) failed: %v", c.expr, err)
			continue
		}
		dec := protocol.NewDecoder(payload)
		root, _ := dec.Decode()
		body := root.Children[0]      // BlockStmt
		assign := body.Children[0]    // _ = expr
		binExpr := assign.Children[1] // RHS = BinaryExpr
		if binExpr.Opcode != protocol.OpBinaryExpr {
			t.Errorf("expected OpBinaryExpr, got 0x%02x", binExpr.Opcode)
		}
		if binExpr.Op != c.op {
			t.Errorf("op = 0x%02x, want 0x%02x", binExpr.Op, c.op)
		}
	}
}

// TestEncode_HexLiteralValue is a regression test for a bug where hex
// literals (0x...) were silently parsed as 0 by fmt.Sscanf("%d"), which
// reads only the leading "0" as a decimal digit and stops at "x" without
// error. This broke every hex literal in hook scripts:
//   - `& 0xFFFFFFFF` masked everything to 0
//   - API address constants like 0x76F5FD50 printed as 0
//
// The fix uses strconv.ParseUint(s, 0, 64) which auto-detects the base.
func TestEncode_HexLiteralValue(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want uint64
	}{
		{"hex 0x76F5FD50", `0x76F5FD50`, 0x76F5FD50},
		{"hex 0xFFFFFFFF", `0xFFFFFFFF`, 0xFFFFFFFF},
		{"hex 0xabcdef", `0xabcdef`, 0xABCDEF},
		{"decimal 123", `123`, 123},
		{"zero", `0`, 0},
		{"octal 0o17", `0o17`, 0o17},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			src := "package hook\nfunc hook(ctx *HookCtx) {\n\tctx.Printf(\"%x\", " + tc.src + ")\n}\n"
			payload, err := Encode(src)
			if err != nil {
				t.Fatalf("Encode failed: %v", err)
			}
			dec := protocol.NewDecoder(payload)
			root, err := dec.Decode()
			if err != nil {
				t.Fatalf("Decode failed: %v", err)
			}
			// FuncLit -> Block -> CallExpr -> [fmtStr, value]
			call := root.Children[0].Children[0]
			if call.Opcode != protocol.OpCallExpr {
				t.Fatalf("expected OpCallExpr, got 0x%02x", call.Opcode)
			}
			if len(call.Children) != 2 {
				t.Fatalf("Printf args = %d, want 2", len(call.Children))
			}
			lit := call.Children[1]
			if lit.Opcode != protocol.OpLiteral {
				t.Fatalf("expected OpLiteral, got 0x%02x", lit.Opcode)
			}
			if lit.Kind != protocol.KindUint64 {
				t.Fatalf("literal kind = 0x%02x, want KindUint64", lit.Kind)
			}
			if len(lit.Value) != 8 {
				t.Fatalf("literal value len = %d, want 8", len(lit.Value))
			}
			got := binary.LittleEndian.Uint64(lit.Value)
			if got != tc.want {
				t.Fatalf("literal value = 0x%x, want 0x%x", got, tc.want)
			}
		})
	}
}
