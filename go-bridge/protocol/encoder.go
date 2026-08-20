package protocol

import (
	"encoding/binary"
	"fmt"
)

// Encoder serialises a Node tree into the wire format described in
// docs/go-subset-spec.md. The output is a complete payload (ast_size +
// nodes + string table) ready to be sent to the kernel driver.
//
// Usage:
//
//	enc := NewEncoder()
//	enc.AddString("ctx") // pre-register names you reference
//	enc.Encode(rootNode)
//	payload := enc.Bytes()
type Encoder struct {
	// nodeBuf accumulates the serialised AST node stream.
	nodeBuf []byte
	// strs is the string table. Maps string → index.
	strs   []string
	strMap map[string]uint16
}

// NewEncoder creates an empty Encoder.
func NewEncoder() *Encoder {
	return &Encoder{strMap: make(map[string]uint16)}
}

// AddString registers s in the string table (if not already present) and
// returns its index. Identifiers, field names, and string literals all share
// this table.
func (e *Encoder) AddString(s string) uint16 {
	if id, ok := e.strMap[s]; ok {
		return id
	}
	id := uint16(len(e.strs))
	e.strs = append(e.strs, s)
	e.strMap[s] = id
	return id
}

// Encode serialises the root node and returns the complete payload bytes.
// The Encoder is reset after Encode, so it can be reused.
//
// Payload layout (must match the kernel C decoder in ast_decode.c::AstDecode,
// which starts reading the root opcode at position 0 — NO ast_size prefix):
//
//	[ast nodes][4 bytes str_count][strings...]
func (e *Encoder) Encode(root *Node) ([]byte, error) {
	e.nodeBuf = e.nodeBuf[:0]
	if err := e.encodeNode(root); err != nil {
		return nil, err
	}
	out := make([]byte, 0, len(e.nodeBuf)+4+e.stringTableSize())
	out = append(out, e.nodeBuf...)
	out = appendU32LE(out, uint32(len(e.strs)))
	for _, s := range e.strs {
		out = appendU16LE(out, uint16(len(s)))
		out = append(out, s...)
	}
	return out, nil
}

// Bytes is an alias for Encode kept for ergonomics when the node was already
// streamed via the low-level write methods. Most callers should use Encode.
func (e *Encoder) Bytes() []byte {
	out := make([]byte, 0, len(e.nodeBuf)+4+e.stringTableSize())
	out = append(out, e.nodeBuf...)
	out = appendU32LE(out, uint32(len(e.strs)))
	for _, s := range e.strs {
		out = appendU16LE(out, uint16(len(s)))
		out = append(out, s...)
	}
	return out
}

func (e *Encoder) stringTableSize() int {
	n := 4 // count
	for _, s := range e.strs {
		n += 2 + len(s)
	}
	return n
}

func (e *Encoder) encodeNode(n *Node) error {
	if n == nil {
		e.nodeBuf = append(e.nodeBuf, OpNil)
		return nil
	}
	switch n.Opcode {
	case OpLiteral:
		return e.encodeLiteral(n)
	case OpIdent:
		e.nodeBuf = append(e.nodeBuf, OpIdent)
		e.nodeBuf = appendU16LE(e.nodeBuf, n.NameID)
	case OpBinaryExpr:
		e.nodeBuf = append(e.nodeBuf, OpBinaryExpr, n.Op)
		if err := e.encodeNode(n.Children[0]); err != nil {
			return err
		}
		return e.encodeNode(n.Children[1])
	case OpUnaryExpr:
		e.nodeBuf = append(e.nodeBuf, OpUnaryExpr, n.Op)
		return e.encodeNode(n.Children[0])
	case OpCallExpr:
		if len(n.Children) > MaxNArgs {
			return NewEncodeError("encode", fmt.Sprintf("too many call args: %d", len(n.Children)))
		}
		e.nodeBuf = append(e.nodeBuf, OpCallExpr)
		e.nodeBuf = appendU16LE(e.nodeBuf, n.FuncID)
		e.nodeBuf = append(e.nodeBuf, byte(len(n.Children)))
		for _, c := range n.Children {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpSelectorExpr:
		e.nodeBuf = append(e.nodeBuf, OpSelectorExpr)
		if err := e.encodeNode(n.Children[0]); err != nil {
			return err
		}
		e.nodeBuf = appendU16LE(e.nodeBuf, n.FieldID)
	case OpAssignStmt:
		e.nodeBuf = append(e.nodeBuf, OpAssignStmt, n.Op)
		if err := e.encodeNode(n.Children[0]); err != nil {
			return err
		}
		return e.encodeNode(n.Children[1])
	case OpIfStmt:
		e.nodeBuf = append(e.nodeBuf, OpIfStmt)
		if err := e.encodeNode(n.Children[0]); err != nil { // cond
			return err
		}
		if err := e.encodeNode(n.Children[1]); err != nil { // then
			return err
		}
		if len(n.Children) > 2 && n.Children[2] != nil {
			return e.encodeNode(n.Children[2]) // else
		}
		e.nodeBuf = append(e.nodeBuf, OpNil) // absent else
	case OpForStmt:
		e.nodeBuf = append(e.nodeBuf, OpForStmt)
		for i := range 4 {
			var c *Node
			if i < len(n.Children) {
				c = n.Children[i]
			}
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpBlockStmt:
		if len(n.Children) > MaxNStatements {
			return NewEncodeError("encode", fmt.Sprintf("too many statements: %d", len(n.Children)))
		}
		e.nodeBuf = append(e.nodeBuf, OpBlockStmt)
		e.nodeBuf = appendU16LE(e.nodeBuf, uint16(len(n.Children)))
		for _, c := range n.Children {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpReturnStmt:
		e.nodeBuf = append(e.nodeBuf, OpReturnStmt, byte(len(n.Children)))
		for _, c := range n.Children {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpFuncDecl:
		e.nodeBuf = append(e.nodeBuf, OpFuncDecl)
		e.nodeBuf = appendU16LE(e.nodeBuf, n.NameID)
		nargs := byte(len(n.Children) - 1)      // last is body
		e.nodeBuf = append(e.nodeBuf, nargs, 0) // nargs, nrets=0 (simplified)
		for _, c := range n.Children {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpFuncLit:
		nargs := byte(0)
		nrets := byte(0)
		// Simplified: children = [body]; args/rets inferred by interpreter.
		e.nodeBuf = append(e.nodeBuf, OpFuncLit, nargs, nrets)
		for _, c := range n.Children {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpDeclStmt:
		// Kernel decoder (ast_decode.c) reads NameId (U16) THEN Kind (U8).
		e.nodeBuf = append(e.nodeBuf, OpDeclStmt)
		e.nodeBuf = appendU16LE(e.nodeBuf, n.NameID)
		e.nodeBuf = append(e.nodeBuf, n.Kind)
		if len(n.Children) > 0 {
			return e.encodeNode(n.Children[0])
		}
		e.nodeBuf = append(e.nodeBuf, OpNil)
	case OpArrayType:
		e.nodeBuf = append(e.nodeBuf, OpArrayType)
		if len(n.Value) >= 5 {
			e.nodeBuf = append(e.nodeBuf, n.Value[:5]...)
		}
	case OpCompositeLit:
		e.nodeBuf = append(e.nodeBuf, OpCompositeLit)
		if err := e.encodeNode(n.Children[0]); err != nil { // type
			return err
		}
		nels := len(n.Children) - 1
		e.nodeBuf = appendU16LE(e.nodeBuf, uint16(nels))
		for _, c := range n.Children[1:] {
			if err := e.encodeNode(c); err != nil {
				return err
			}
		}
	case OpIndexExpr:
		e.nodeBuf = append(e.nodeBuf, OpIndexExpr)
		if err := e.encodeNode(n.Children[0]); err != nil {
			return err
		}
		return e.encodeNode(n.Children[1])
	case OpNil:
		e.nodeBuf = append(e.nodeBuf, OpNil)
	default:
		return NewEncodeError("encode", fmt.Sprintf("unknown opcode 0x%02x", n.Opcode))
	}
	return nil
}

func (e *Encoder) encodeLiteral(n *Node) error {
	e.nodeBuf = append(e.nodeBuf, OpLiteral, n.Kind)
	switch n.Kind {
	case KindUint8, KindInt8, KindBool:
		if len(n.Value) != 1 {
			return NewEncodeError("encode", fmt.Sprintf("bool/u8 literal needs 1 byte, got %d", len(n.Value)))
		}
		e.nodeBuf = append(e.nodeBuf, n.Value...)
	case KindUint16, KindInt16:
		if len(n.Value) != 2 {
			return NewEncodeError("encode", fmt.Sprintf("u16 literal needs 2 bytes, got %d", len(n.Value)))
		}
		e.nodeBuf = append(e.nodeBuf, n.Value...)
	case KindUint32, KindInt32:
		if len(n.Value) != 4 {
			return NewEncodeError("encode", fmt.Sprintf("u32 literal needs 4 bytes, got %d", len(n.Value)))
		}
		e.nodeBuf = append(e.nodeBuf, n.Value...)
	case KindUint64, KindInt64:
		if len(n.Value) != 8 {
			return NewEncodeError("encode", fmt.Sprintf("u64 literal needs 8 bytes, got %d", len(n.Value)))
		}
		e.nodeBuf = append(e.nodeBuf, n.Value...)
	case KindString:
		if len(n.Value) > MaxStringLength {
			return NewEncodeError("encode", fmt.Sprintf("string too long: %d", len(n.Value)))
		}
		e.nodeBuf = appendU16LE(e.nodeBuf, uint16(len(n.Value)))
		e.nodeBuf = append(e.nodeBuf, n.Value...)
	default:
		return NewEncodeError("encode", fmt.Sprintf("unknown literal kind 0x%02x", n.Kind))
	}
	return nil
}

func appendU16LE(b []byte, v uint16) []byte {
	return append(b, byte(v), byte(v>>8))
}

func appendU32LE(b []byte, v uint32) []byte {
	return append(b, byte(v), byte(v>>8), byte(v>>16), byte(v>>24))
}

// quickEncode is a convenience that creates an Encoder, encodes root, and
// returns the payload. Used by tests and one-shot callers.
func quickEncode(root *Node) ([]byte, error) {
	enc := NewEncoder()
	return enc.Encode(root)
}

var _ = binary.LittleEndian // keep import for reference; append helpers are hand-rolled for clarity
