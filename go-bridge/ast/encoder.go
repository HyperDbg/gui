// Package astencoder converts a Go source callback (a FuncLit body) into the
// binary AST wire format defined by go-bridge/protocol. It is the user-mode
// compiler counterpart to the kernel C decoder/interpreter.
//
// Pipeline:
//
//	src string
//	  → go/parser.ParseFile
//	  → go-bridge/subset.Validator (reject unsupported constructs)
//	  → go-bridge/ast.Encode (go/ast → protocol.Node tree → []byte)
//	  → []byte payload sent to driver via IOCTL
//
// The encoder handles the subset of Go that the kernel interpreter can
// execute: see docs/go-subset-spec.md for the full definition.
package astencoder

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"

	"github.com/hyperdbg/go-bridge/protocol"
	"github.com/hyperdbg/go-bridge/subset"
)

// Encode parses src, validates it against the Go subset, and returns the
// binary AST payload. src must be a Go source file containing at least one
// FuncLit (the hook callback) or a function body.
//
// Example src:
//
//	package hook
//	func(ctx *HookCtx) {
//	    ret := ctx.StackReadQword(0) & 0xFFFFFFFF
//	    ctx.Printf("RAH ret=%x", ret)
//	    if ret < 0x10000000 {
//	        ctx.Break()
//	    }
//	}
func Encode(src string) ([]byte, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hook.go", src, parser.AllErrors)
	if err != nil {
		return nil, protocol.NewValidationError("parse", err.Error())
	}
	// 1. Validate against the subset.
	v := subset.New()
	if err := v.ValidateFile(f); err != nil {
		return nil, err
	}
	// 2. Find the hook callback (first FuncLit or FuncDecl).
	cb := findCallback(f)
	if cb == nil {
		return nil, protocol.NewValidationError("encode", "no hook callback (FuncLit or FuncDecl) found in source")
	}
	// 3. Convert go/ast → protocol.Node tree.
	conv := newConverter()
	root, err := conv.convertFuncLit(cb)
	if err != nil {
		return nil, err
	}
	// 4. Serialise to wire format.
	enc := protocol.NewEncoder()
	// Pre-register all strings collected during conversion.
	for _, s := range conv.strings {
		enc.AddString(s)
	}
	return enc.Encode(root)
}

// EncodeFuncLit encodes a FuncLit directly (used by tests that construct
// go/ast nodes manually).
func EncodeFuncLit(fl *ast.FuncLit) ([]byte, error) {
	conv := newConverter()
	root, err := conv.convertFuncLit(fl)
	if err != nil {
		return nil, err
	}
	enc := protocol.NewEncoder()
	for _, s := range conv.strings {
		enc.AddString(s)
	}
	return enc.Encode(root)
}

// findCallback returns the first FuncLit or FuncDecl body in the file. When a
// FuncDecl is found, its body is wrapped in a synthetic FuncLit so the rest of
// the encoder treats it uniformly.
func findCallback(f *ast.File) *ast.FuncLit {
	// Prefer a FuncLit (anonymous closure assigned to a variable).
	var lit *ast.FuncLit
	ast.Inspect(f, func(n ast.Node) bool {
		if lit != nil {
			return false
		}
		if fl, ok := n.(*ast.FuncLit); ok {
			lit = fl
			return false
		}
		return true
	})
	if lit != nil {
		return lit
	}
	// Fall back to the first top-level FuncDecl.
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok && fd.Body != nil {
			return &ast.FuncLit{Type: fd.Type, Body: fd.Body}
		}
	}
	return nil
}

// converter walks go/ast and builds a protocol.Node tree, collecting strings
// for the string table.
type converter struct {
	strings   []string
	stringIDs map[string]uint16
}

func newConverter() *converter {
	return &converter{stringIDs: make(map[string]uint16)}
}

func (c *converter) addString(s string) uint16 {
	if id, ok := c.stringIDs[s]; ok {
		return id
	}
	id := uint16(len(c.strings))
	c.strings = append(c.strings, s)
	c.stringIDs[s] = id
	return id
}

func (c *converter) convertFuncLit(fl *ast.FuncLit) (*protocol.Node, error) {
	body, err := c.convertBlockStmt(fl.Body)
	if err != nil {
		return nil, err
	}
	return protocol.NewFuncLit(body), nil
}

func (c *converter) convertBlockStmt(b *ast.BlockStmt) (*protocol.Node, error) {
	stmts := make([]*protocol.Node, 0, len(b.List))
	for _, s := range b.List {
		node, err := c.convertStmt(s)
		if err != nil {
			return nil, err
		}
		if node != nil {
			stmts = append(stmts, node)
		}
	}
	return protocol.NewBlock(stmts...), nil
}

func (c *converter) convertStmt(s ast.Stmt) (*protocol.Node, error) {
	switch n := s.(type) {
	case nil:
		return nil, nil
	case *ast.ExprStmt:
		return c.convertExpr(n.X)
	case *ast.AssignStmt:
		return c.convertAssignStmt(n)
	case *ast.DeclStmt:
		return c.convertDeclStmt(n)
	case *ast.IfStmt:
		return c.convertIfStmt(n)
	case *ast.ForStmt:
		return c.convertForStmt(n)
	case *ast.BlockStmt:
		return c.convertBlockStmt(n)
	case *ast.ReturnStmt:
		return c.convertReturnStmt(n)
	case *ast.IncDecStmt:
		return c.convertIncDecStmt(n)
	case *ast.BranchStmt:
		// break / continue / goto — encoded as ReturnStmt sentinel
		// For now, encode break/continue as a CallExpr with a sentinel func_id.
		// The interpreter treats FuncBreak/FuncContinue specially only inside
		// loops. This is a simplification for Phase B; full break/continue
		// support arrives in Phase C.5.
		return nil, nil // TODO: implement break/continue encoding in Phase C.5
	}
	return nil, protocol.NewValidationError("encode", fmt.Sprintf("unsupported statement: %T", s))
}

func (c *converter) convertAssignStmt(n *ast.AssignStmt) (*protocol.Node, error) {
	if len(n.Lhs) != 1 || len(n.Rhs) != 1 {
		return nil, protocol.NewValidationError("encode", "multi-assign not supported")
	}
	// Short variable declaration (:=) → encode as DeclStmt.
	if n.Tok == token.DEFINE {
		ident, ok := n.Lhs[0].(*ast.Ident)
		if !ok {
			return nil, protocol.NewValidationError("encode", ":= target must be an identifier")
		}
		nameID := c.addString(ident.Name)
		init, err := c.convertExpr(n.Rhs[0])
		if err != nil {
			return nil, err
		}
		return protocol.NewDecl(protocol.KindUint64, nameID, init), nil
	}
	op := assignOpFromToken(n.Tok)
	if op == 0 {
		return nil, protocol.NewValidationError("encode", "unsupported assign token: "+n.Tok.String())
	}
	lhs, err := c.convertExpr(n.Lhs[0])
	if err != nil {
		return nil, err
	}
	rhs, err := c.convertExpr(n.Rhs[0])
	if err != nil {
		return nil, err
	}
	return protocol.NewAssign(op, lhs, rhs), nil
}

func (c *converter) convertDeclStmt(n *ast.DeclStmt) (*protocol.Node, error) {
	gd, ok := n.Decl.(*ast.GenDecl)
	if !ok {
		return nil, protocol.NewValidationError("encode", "unsupported decl type")
	}
	for _, spec := range gd.Specs {
		if vs, ok := spec.(*ast.ValueSpec); ok {
			if len(vs.Names) != 1 {
				return nil, protocol.NewValidationError("encode", "multi-name var decl not supported")
			}
			nameID := c.addString(vs.Names[0].Name)
			kind := kindFromType(vs.Type)
			var init *protocol.Node
			if len(vs.Values) > 0 {
				var err error
				init, err = c.convertExpr(vs.Values[0])
				if err != nil {
					return nil, err
				}
			}
			return protocol.NewDecl(kind, nameID, init), nil
		}
	}
	return nil, protocol.NewValidationError("encode", "no value spec in decl")
}

func (c *converter) convertIfStmt(n *ast.IfStmt) (*protocol.Node, error) {
	cond, err := c.convertExpr(n.Cond)
	if err != nil {
		return nil, err
	}
	thenNode, err := c.convertBlockStmt(n.Body)
	if err != nil {
		return nil, err
	}
	var elseNode *protocol.Node
	if n.Else != nil {
		if eb, ok := n.Else.(*ast.BlockStmt); ok {
			elseNode, err = c.convertBlockStmt(eb)
		} else {
			elseNode, err = c.convertStmt(n.Else)
		}
		if err != nil {
			return nil, err
		}
	}
	return protocol.NewIf(cond, thenNode, elseNode), nil
}

func (c *converter) convertForStmt(n *ast.ForStmt) (*protocol.Node, error) {
	var initN, condN, postN, bodyN *protocol.Node
	var err error
	if n.Init != nil {
		initN, err = c.convertStmt(n.Init)
		if err != nil {
			return nil, err
		}
	}
	if n.Cond != nil {
		condN, err = c.convertExpr(n.Cond)
		if err != nil {
			return nil, err
		}
	}
	if n.Post != nil {
		postN, err = c.convertStmt(n.Post)
		if err != nil {
			return nil, err
		}
	}
	if n.Body != nil {
		bodyN, err = c.convertBlockStmt(n.Body)
		if err != nil {
			return nil, err
		}
	}
	return protocol.NewFor(initN, condN, postN, bodyN), nil
}

func (c *converter) convertReturnStmt(n *ast.ReturnStmt) (*protocol.Node, error) {
	vals := make([]*protocol.Node, 0, len(n.Results))
	for _, r := range n.Results {
		node, err := c.convertExpr(r)
		if err != nil {
			return nil, err
		}
		vals = append(vals, node)
	}
	return protocol.NewReturn(vals...), nil
}

func (c *converter) convertIncDecStmt(n *ast.IncDecStmt) (*protocol.Node, error) {
	// i++ → i = i + 1 ; i-- → i = i - 1
	x, err := c.convertExpr(n.X)
	if err != nil {
		return nil, err
	}
	var oneNode *protocol.Node
	// Determine the literal kind from the expression; default to uint64.
	oneNode = protocol.NewLiteralUint64(1)
	var op byte
	if n.Tok == token.INC {
		op = protocol.BinOpAdd
	} else {
		op = protocol.BinOpSub
	}
	addExpr := protocol.NewBinary(op, x, oneNode)
	lhs, err := c.convertExpr(n.X)
	if err != nil {
		return nil, err
	}
	return protocol.NewAssign(protocol.AssignAssign, lhs, addExpr), nil
}

func (c *converter) convertExpr(e ast.Expr) (*protocol.Node, error) {
	switch n := e.(type) {
	case nil:
		return nil, nil
	case *ast.BasicLit:
		return c.convertBasicLit(n)
	case *ast.Ident:
		nameID := c.addString(n.Name)
		return protocol.NewIdent(nameID), nil
	case *ast.BinaryExpr:
		return c.convertBinaryExpr(n)
	case *ast.UnaryExpr:
		return c.convertUnaryExpr(n)
	case *ast.ParenExpr:
		return c.convertExpr(n.X)
	case *ast.CallExpr:
		return c.convertCallExpr(n)
	case *ast.SelectorExpr:
		return c.convertSelectorExpr(n)
	case *ast.IndexExpr:
		x, err := c.convertExpr(n.X)
		if err != nil {
			return nil, err
		}
		idx, err := c.convertExpr(n.Index)
		if err != nil {
			return nil, err
		}
		return &protocol.Node{Opcode: protocol.OpIndexExpr, Children: []*protocol.Node{x, idx}}, nil
	case *ast.CompositeLit:
		return c.convertCompositeLit(n)
	case *ast.FuncLit:
		return c.convertFuncLit(n)
	}
	return nil, protocol.NewValidationError("encode", fmt.Sprintf("unsupported expression: %T", e))
}

func (c *converter) convertBasicLit(n *ast.BasicLit) (*protocol.Node, error) {
	switch n.Kind {
	case token.INT:
		// Parse with base 0: auto-detects 0x (hex), 0o (octal), 0b (binary),
		// and plain decimal. Strconv also strips Go's underscore separators.
		// NB: a previous implementation used fmt.Sscanf("%d"), which silently
		// parsed "0x76F5FD50" as 0 (reading only the leading "0" as a decimal
		// digit and stopping at "x" without error). This broke every hex
		// literal in hook scripts (e.g. `& 0xFFFFFFFF` masked to 0, and
		// API address constants like 0x76F5FD50 printed as 0).
		v, err := strconv.ParseUint(n.Value, 0, 64)
		if err != nil {
			return nil, protocol.NewValidationError("encode", "cannot parse int literal: "+n.Value)
		}
		return protocol.NewLiteralUint64(v), nil
	case token.STRING:
		s := n.Value
		if len(s) >= 2 {
			s = s[1 : len(s)-1] // strip quotes
		}
		return protocol.NewLiteralString(s), nil
	case token.CHAR:
		// Rune literal — encode as uint64.
		s := n.Value
		if len(s) >= 2 {
			s = s[1 : len(s)-1]
		}
		var r uint64
		fmt.Sscanf(s, "%c", &r)
		return protocol.NewLiteralUint64(r), nil
	case token.FLOAT:
		return nil, protocol.NewValidationError("encode", "float literals not supported in Go subset")
	}
	return nil, protocol.NewValidationError("encode", "unsupported literal kind: "+n.Kind.String())
}

func (c *converter) convertBinaryExpr(n *ast.BinaryExpr) (*protocol.Node, error) {
	op := binaryOpFromToken(n.Op)
	if op == 0 {
		return nil, protocol.NewValidationError("encode", "unsupported binary operator: "+n.Op.String())
	}
	lhs, err := c.convertExpr(n.X)
	if err != nil {
		return nil, err
	}
	rhs, err := c.convertExpr(n.Y)
	if err != nil {
		return nil, err
	}
	return protocol.NewBinary(op, lhs, rhs), nil
}

func (c *converter) convertUnaryExpr(n *ast.UnaryExpr) (*protocol.Node, error) {
	op := unaryOpFromToken(n.Op)
	if op == 0 {
		return nil, protocol.NewValidationError("encode", "unsupported unary operator: "+n.Op.String())
	}
	operand, err := c.convertExpr(n.X)
	if err != nil {
		return nil, err
	}
	return protocol.NewUnary(op, operand), nil
}

func (c *converter) convertCallExpr(n *ast.CallExpr) (*protocol.Node, error) {
	// Only ctx.Method(...) calls are supported (SelectorExpr with X=ctx).
	sel, ok := n.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, protocol.NewValidationError("encode", "only ctx.Method() calls are supported")
	}
	// Resolve the method name to a func_id.
	funcID, ok := protocol.WhitelistFuncs[sel.Sel.Name]
	if !ok {
		return nil, protocol.NewValidationError("encode", "unknown whitelist method: "+sel.Sel.Name)
	}
	args := make([]*protocol.Node, 0, len(n.Args))
	for _, a := range n.Args {
		node, err := c.convertExpr(a)
		if err != nil {
			return nil, err
		}
		args = append(args, node)
	}
	return protocol.NewCall(funcID, args...), nil
}

func (c *converter) convertSelectorExpr(n *ast.SelectorExpr) (*protocol.Node, error) {
	x, err := c.convertExpr(n.X)
	if err != nil {
		return nil, err
	}
	fieldID := c.addString(n.Sel.Name)
	return protocol.NewSelector(x, fieldID), nil
}

func (c *converter) convertCompositeLit(n *ast.CompositeLit) (*protocol.Node, error) {
	typ, err := c.convertExpr(n.Type)
	if err != nil {
		return nil, err
	}
	els := make([]*protocol.Node, 0, len(n.Elts)+1)
	els = append(els, typ)
	for _, e := range n.Elts {
		node, err := c.convertExpr(e)
		if err != nil {
			return nil, err
		}
		els = append(els, node)
	}
	return &protocol.Node{Opcode: protocol.OpCompositeLit, Children: els}, nil
}

// --- Token → opcode mapping helpers ---

func binaryOpFromToken(t token.Token) byte {
	switch t {
	case token.ADD:
		return protocol.BinOpAdd
	case token.SUB:
		return protocol.BinOpSub
	case token.MUL:
		return protocol.BinOpMul
	case token.QUO:
		return protocol.BinOpQuo
	case token.REM:
		return protocol.BinOpRem
	case token.AND:
		return protocol.BinOpAnd
	case token.OR:
		return protocol.BinOpOr
	case token.XOR:
		return protocol.BinOpXor
	case token.SHL:
		return protocol.BinOpShl
	case token.SHR:
		return protocol.BinOpShr
	case token.EQL:
		return protocol.BinOpEql
	case token.NEQ:
		return protocol.BinOpNeq
	case token.LSS:
		return protocol.BinOpLss
	case token.GTR:
		return protocol.BinOpGtr
	case token.LEQ:
		return protocol.BinOpLeq
	case token.GEQ:
		return protocol.BinOpGeq
	case token.LAND:
		return protocol.BinOpLAnd
	case token.LOR:
		return protocol.BinOpLOr
	case token.AND_NOT:
		return protocol.BinOpAndNot
	}
	return 0
}

func unaryOpFromToken(t token.Token) byte {
	switch t {
	case token.SUB:
		return protocol.UnOpNeg
	case token.NOT:
		return protocol.UnOpNot
	case token.XOR:
		return protocol.UnOpXor
	}
	return 0
}

func assignOpFromToken(t token.Token) byte {
	switch t {
	case token.ASSIGN:
		return protocol.AssignAssign
	case token.ADD_ASSIGN:
		return protocol.AssignAdd
	case token.SUB_ASSIGN:
		return protocol.AssignSub
	case token.MUL_ASSIGN:
		return protocol.AssignMul
	case token.QUO_ASSIGN:
		return protocol.AssignQuo
	case token.REM_ASSIGN:
		return protocol.AssignRem
	case token.AND_ASSIGN:
		return protocol.AssignAnd
	case token.OR_ASSIGN:
		return protocol.AssignOr
	case token.XOR_ASSIGN:
		return protocol.AssignXor
	case token.SHL_ASSIGN:
		return protocol.AssignShl
	case token.SHR_ASSIGN:
		return protocol.AssignShr
	case token.AND_NOT_ASSIGN:
		return protocol.AssignAndNot
	}
	return 0
}

func kindFromType(t ast.Expr) byte {
	if ident, ok := t.(*ast.Ident); ok {
		switch ident.Name {
		case "uint8":
			return protocol.KindUint8
		case "uint16":
			return protocol.KindUint16
		case "uint32":
			return protocol.KindUint32
		case "uint64":
			return protocol.KindUint64
		case "int8":
			return protocol.KindInt8
		case "int16":
			return protocol.KindInt16
		case "int32":
			return protocol.KindInt32
		case "int64":
			return protocol.KindInt64
		case "bool":
			return protocol.KindBool
		case "string":
			return protocol.KindString
		}
	}
	return protocol.KindUint64 // default
}
