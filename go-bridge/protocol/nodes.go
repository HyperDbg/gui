package protocol

import "fmt"

// Node is the in-memory representation of an AST node. The encoder builds this
// tree (from go/ast) and then serializes it; tests can also construct nodes
// directly to verify the decoder.
type Node struct {
	Opcode     byte
	Kind       byte   // Literal kind (OpLiteral only)
	Op         byte   // binary/unary/assign operator
	FuncID     uint16 // whitelist func_id (OpCallExpr only)
	NameID     uint16 // string-table index (OpIdent, OpDeclStmt, OpFuncDecl)
	FieldID    uint16 // string-table index (OpSelectorExpr)
	IsVariadic bool   // OpCallExpr: last arg is variadic (Printf)
	// Value holds the literal bytes for OpLiteral (KindString includes length
	// prefix only in the wire format; here Value is the raw string bytes).
	Value []byte
	// Children holds sub-nodes. Meaning depends on Opcode:
	//   OpBinaryExpr: [LHS, RHS]
	//   OpUnaryExpr:  [Operand]
	//   OpCallExpr:   [arg0, arg1, ...]
	//   OpSelectorExpr: [X]
	//   OpAssignStmt: [LHS, RHS]
	//   OpIfStmt:     [Cond, Then, Else] (Else may be nil)
	//   OpForStmt:    [Init, Cond, Post, Body] (Init/Post may be nil)
	//   OpBlockStmt:  [stmt0, stmt1, ...]
	//   OpReturnStmt: [val0, val1, ...]
	//   OpFuncDecl/OpFuncLit: [ParamType0, ..., RetType0, ..., Body]
	//   OpDeclStmt:   [InitExpr]
	//   OpCompositeLit: [elem0, elem1, ...]
	//   OpIndexExpr:  [X, Index]
	Children []*Node
}

// NewNode creates a Node with the given opcode and optional children.
func NewNode(opcode byte, children ...*Node) *Node {
	return &Node{Opcode: opcode, Children: children}
}

// NewLiteral creates a Literal node from a Go value of a supported kind.
func NewLiteral(kind byte, value []byte) *Node {
	return &Node{Opcode: OpLiteral, Kind: kind, Value: value}
}

// NewLiteralBool creates a bool literal.
func NewLiteralBool(b bool) *Node {
	v := byte(0)
	if b {
		v = 1
	}
	return NewLiteral(KindBool, []byte{v})
}

// NewLiteralUint64 creates a uint64 literal (8 bytes, little-endian).
func NewLiteralUint64(u uint64) *Node {
	v := make([]byte, 8)
	for i := range 8 {
		v[i] = byte(u >> (8 * i))
	}
	return NewLiteral(KindUint64, v)
}

// NewLiteralUint32 creates a uint32 literal (4 bytes, little-endian).
func NewLiteralUint32(u uint32) *Node {
	v := make([]byte, 4)
	for i := range 4 {
		v[i] = byte(u >> (8 * i))
	}
	return NewLiteral(KindUint32, v)
}

// NewLiteralString creates a string literal. The Value field stores the raw
// bytes; the wire format prepends a 2-byte length (handled by the encoder).
func NewLiteralString(s string) *Node {
	return NewLiteral(KindString, []byte(s))
}

// NewBinary creates a BinaryExpr node.
func NewBinary(op byte, lhs, rhs *Node) *Node {
	return &Node{Opcode: OpBinaryExpr, Op: op, Children: []*Node{lhs, rhs}}
}

// NewUnary creates a UnaryExpr node.
func NewUnary(op byte, operand *Node) *Node {
	return &Node{Opcode: OpUnaryExpr, Op: op, Children: []*Node{operand}}
}

// NewCall creates a CallExpr node for a whitelist function.
func NewCall(funcID uint16, args ...*Node) *Node {
	return &Node{Opcode: OpCallExpr, FuncID: funcID, Children: args}
}

// NewIdent creates an Ident node referencing string-table index nameID.
func NewIdent(nameID uint16) *Node {
	return &Node{Opcode: OpIdent, NameID: nameID}
}

// NewSelector creates a SelectorExpr node (X.FieldID).
func NewSelector(x *Node, fieldID uint16) *Node {
	return &Node{Opcode: OpSelectorExpr, FieldID: fieldID, Children: []*Node{x}}
}

// NewAssign creates an AssignStmt node.
func NewAssign(op byte, lhs, rhs *Node) *Node {
	return &Node{Opcode: OpAssignStmt, Op: op, Children: []*Node{lhs, rhs}}
}

// NewIf creates an IfStmt node. elseNode may be nil.
func NewIf(cond, thenNode, elseNode *Node) *Node {
	return &Node{Opcode: OpIfStmt, Children: []*Node{cond, thenNode, elseNode}}
}

// NewFor creates a ForStmt node. init and post may be nil.
func NewFor(init, cond, post, body *Node) *Node {
	return &Node{Opcode: OpForStmt, Children: []*Node{init, cond, post, body}}
}

// NewBlock creates a BlockStmt node.
func NewBlock(stmts ...*Node) *Node {
	return &Node{Opcode: OpBlockStmt, Children: stmts}
}

// NewReturn creates a ReturnStmt node.
func NewReturn(vals ...*Node) *Node {
	return &Node{Opcode: OpReturnStmt, Children: vals}
}

// NewFuncLit creates a FuncLit node. The Children are [Body] (a BlockStmt);
// parameter/return types are inferred by the interpreter.
func NewFuncLit(body *Node) *Node {
	return &Node{Opcode: OpFuncLit, Children: []*Node{body}}
}

// NewDecl creates a DeclStmt node (var declaration).
func NewDecl(kind byte, nameID uint16, init *Node) *Node {
	return &Node{Opcode: OpDeclStmt, Kind: kind, NameID: nameID, Children: []*Node{init}}
}

// NewBreak creates a return that signals break. We encode break/continue as a
// CallExpr with a sentinel func_id (the interpreter treats them specially).
// TODO: This is a placeholder; the current spec does not have break/continue
// opcodes. For now, for-loops with break/continue use the ReturnStmt path.
// This will be refined in Phase C.5 when for-loops are fully supported.

// String returns a human-readable debug representation of the node tree.
func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return nodeString(n, 0)
}

func nodeString(n *Node, indent int) string {
	if n == nil {
		return "nil"
	}
	pad := ""
	for range indent {
		pad += "  "
	}
	s := pad + opcodeName(n.Opcode)
	switch n.Opcode {
	case OpLiteral:
		s += "(" + kindName(n.Kind) + ")"
		if n.Kind == KindString {
			s += " " + fmt.Sprintf("%q", string(n.Value))
		} else {
			s += " " + fmt.Sprintf("%x", n.Value)
		}
	case OpBinaryExpr, OpUnaryExpr, OpAssignStmt:
		s += "(" + opName(n.Opcode, n.Op) + ")"
	case OpCallExpr:
		s += fmt.Sprintf("(func_id=0x%04x, nargs=%d)", n.FuncID, len(n.Children))
	case OpIdent:
		s += fmt.Sprintf("(name_id=%d)", n.NameID)
	case OpSelectorExpr:
		s += fmt.Sprintf("(field_id=%d)", n.FieldID)
	case OpDeclStmt:
		s += fmt.Sprintf("(kind=%s, name_id=%d)", kindName(n.Kind), n.NameID)
	}
	for _, c := range n.Children {
		s += "\n" + nodeString(c, indent+1)
	}
	return s
}

func opcodeName(op byte) string {
	switch op {
	case OpNil:
		return "Nil"
	case OpLiteral:
		return "Literal"
	case OpIdent:
		return "Ident"
	case OpBinaryExpr:
		return "BinaryExpr"
	case OpUnaryExpr:
		return "UnaryExpr"
	case OpCallExpr:
		return "CallExpr"
	case OpSelectorExpr:
		return "SelectorExpr"
	case OpAssignStmt:
		return "AssignStmt"
	case OpIfStmt:
		return "IfStmt"
	case OpForStmt:
		return "ForStmt"
	case OpBlockStmt:
		return "BlockStmt"
	case OpReturnStmt:
		return "ReturnStmt"
	case OpFuncDecl:
		return "FuncDecl"
	case OpFuncLit:
		return "FuncLit"
	case OpDeclStmt:
		return "DeclStmt"
	case OpArrayType:
		return "ArrayType"
	case OpCompositeLit:
		return "CompositeLit"
	case OpIndexExpr:
		return "IndexExpr"
	}
	return fmt.Sprintf("Unknown(0x%02x)", op)
}

func kindName(k byte) string {
	switch k {
	case KindUint8:
		return "uint8"
	case KindUint16:
		return "uint16"
	case KindUint32:
		return "uint32"
	case KindUint64:
		return "uint64"
	case KindInt8:
		return "int8"
	case KindInt16:
		return "int16"
	case KindInt32:
		return "int32"
	case KindInt64:
		return "int64"
	case KindBool:
		return "bool"
	case KindString:
		return "string"
	}
	return fmt.Sprintf("kind(0x%02x)", k)
}

func opName(opcode, op byte) string {
	switch opcode {
	case OpBinaryExpr:
		switch op {
		case BinOpAdd:
			return "+"
		case BinOpSub:
			return "-"
		case BinOpMul:
			return "*"
		case BinOpQuo:
			return "/"
		case BinOpRem:
			return "%"
		case BinOpAnd:
			return "&"
		case BinOpOr:
			return "|"
		case BinOpXor:
			return "^"
		case BinOpShl:
			return "<<"
		case BinOpShr:
			return ">>"
		case BinOpEql:
			return "=="
		case BinOpNeq:
			return "!="
		case BinOpLss:
			return "<"
		case BinOpGtr:
			return ">"
		case BinOpLeq:
			return "<="
		case BinOpGeq:
			return ">="
		case BinOpLAnd:
			return "&&"
		case BinOpLOr:
			return "||"
		case BinOpAndNot:
			return "&^"
		}
	case OpUnaryExpr:
		switch op {
		case UnOpNeg:
			return "-"
		case UnOpNot:
			return "!"
		case UnOpXor:
			return "^"
		}
	case OpAssignStmt:
		switch op {
		case AssignAssign:
			return "="
		case AssignAdd:
			return "+="
		case AssignSub:
			return "-="
		case AssignMul:
			return "*="
		case AssignQuo:
			return "/="
		case AssignRem:
			return "%="
		case AssignAnd:
			return "&="
		case AssignOr:
			return "|="
		case AssignXor:
			return "^="
		case AssignShl:
			return "<<="
		case AssignShr:
			return ">>="
		case AssignAndNot:
			return "&^="
		}
	}
	return fmt.Sprintf("op(0x%02x)", op)
}
