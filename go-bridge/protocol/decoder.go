package protocol

import (
	"encoding/binary"
	"fmt"
)

// Decoder reads a wire-format AST buffer and reconstructs the Node tree.
//
// It is the Go mirror of the kernel C decoder (ast_decode.c). Keeping both in
// sync is how we verify the protocol — the Go test suite round-trips every
// node type through Encoder → Decoder and asserts equality.
type Decoder struct {
	buf  []byte
	pos  int
	strs []string // string table, indexed by name_id/field_id
}

// NewDecoder creates a Decoder for the given wire-format buffer. The buffer
// must be a complete payload (AST nodes + string table) as produced by
// Encoder.Bytes().
func NewDecoder(buf []byte) *Decoder {
	return &Decoder{buf: buf}
}

// Decode parses the entire buffer and returns the root Node.
//
// Layout (matches the kernel C decoder ast_decode.c::AstDecode, which starts
// reading the root opcode at position 0 — NO ast_size prefix):
//
//	[ast nodes][4 bytes str_count][strings...]
func (d *Decoder) Decode() (*Node, error) {
	if len(d.buf) < 1 {
		return nil, NewDecodeError(uint32(d.pos), "buffer too short for root opcode")
	}
	node, err := d.decodeNode()
	if err != nil {
		return nil, err
	}

	// Parse string table.
	if d.pos+4 > len(d.buf) {
		return nil, NewDecodeError(uint32(d.pos), "buffer too short for str_count")
	}
	strCount := binary.LittleEndian.Uint32(d.buf[d.pos : d.pos+4])
	d.pos += 4
	d.strs = make([]string, strCount)
	for i := uint32(0); i < strCount; i++ {
		s, err := d.decodeString()
		if err != nil {
			return nil, err
		}
		d.strs[i] = s
	}
	return node, nil
}

// Strings returns the decoded string table. Valid after Decode().
func (d *Decoder) Strings() []string { return d.strs }

// StringAt returns the string at the given index, or "" if out of range.
func (d *Decoder) StringAt(id uint16) string {
	if int(id) < len(d.strs) {
		return d.strs[id]
	}
	return ""
}

func (d *Decoder) decodeString() (string, error) {
	if d.pos+2 > len(d.buf) {
		return "", NewDecodeError(uint32(d.pos), "buffer too short for string length")
	}
	n := binary.LittleEndian.Uint16(d.buf[d.pos : d.pos+2])
	d.pos += 2
	if int(n) > len(d.buf)-d.pos {
		return "", NewDecodeError(uint32(d.pos), "string length exceeds buffer")
	}
	s := string(d.buf[d.pos : d.pos+int(n)])
	d.pos += int(n)
	return s, nil
}

func (d *Decoder) readByte() (byte, error) {
	if d.pos >= len(d.buf) {
		return 0, NewDecodeError(uint32(d.pos), "unexpected end of buffer")
	}
	b := d.buf[d.pos]
	d.pos++
	return b, nil
}

func (d *Decoder) readU16() (uint16, error) {
	if d.pos+2 > len(d.buf) {
		return 0, NewDecodeError(uint32(d.pos), "unexpected end of buffer (u16)")
	}
	v := binary.LittleEndian.Uint16(d.buf[d.pos : d.pos+2])
	d.pos += 2
	return v, nil
}

func (d *Decoder) readU32() (uint32, error) {
	if d.pos+4 > len(d.buf) {
		return 0, NewDecodeError(uint32(d.pos), "unexpected end of buffer (u32)")
	}
	v := binary.LittleEndian.Uint32(d.buf[d.pos : d.pos+4])
	d.pos += 4
	return v, nil
}

func (d *Decoder) readBytes(n int) ([]byte, error) {
	if n < 0 || d.pos+n > len(d.buf) {
		return nil, NewDecodeError(uint32(d.pos), fmt.Sprintf("cannot read %d bytes", n))
	}
	b := d.buf[d.pos : d.pos+n]
	d.pos += n
	return b, nil
}

func (d *Decoder) decodeNode() (*Node, error) {
	op, err := d.readByte()
	if err != nil {
		return nil, err
	}
	if op == OpNil {
		return nil, nil // nil node (absent Else)
	}
	node := &Node{Opcode: op}
	switch op {
	case OpLiteral:
		kind, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Kind = kind
		val, err := d.decodeLiteralValue(kind)
		if err != nil {
			return nil, err
		}
		node.Value = val
	case OpIdent:
		id, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.NameID = id
	case OpBinaryExpr:
		o, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Op = o
		lhs, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		rhs, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{lhs, rhs}
	case OpUnaryExpr:
		o, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Op = o
		operand, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{operand}
	case OpCallExpr:
		fid, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.FuncID = fid
		nargs, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Children = make([]*Node, nargs)
		for i := 0; i < int(nargs); i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[i] = c
		}
	case OpSelectorExpr:
		x, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		fid, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.FieldID = fid
		node.Children = []*Node{x}
	case OpAssignStmt:
		o, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Op = o
		lhs, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		rhs, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{lhs, rhs}
	case OpIfStmt:
		cond, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		thenNode, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		elseNode, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{cond, thenNode, elseNode}
	case OpForStmt:
		init, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		cond, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		post, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		body, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{init, cond, post, body}
	case OpBlockStmt:
		n, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.Children = make([]*Node, n)
		for i := 0; i < int(n); i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[i] = c
		}
	case OpReturnStmt:
		n, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Children = make([]*Node, n)
		for i := 0; i < int(n); i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[i] = c
		}
	case OpFuncDecl:
		nameID, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.NameID = nameID
		nargs, err := d.readByte()
		if err != nil {
			return nil, err
		}
		nrets, err := d.readByte()
		if err != nil {
			return nil, err
		}
		// args + rets + body
		total := int(nargs) + int(nrets) + 1
		node.Children = make([]*Node, total)
		for i := 0; i < total; i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[i] = c
		}
	case OpFuncLit:
		nargs, err := d.readByte()
		if err != nil {
			return nil, err
		}
		nrets, err := d.readByte()
		if err != nil {
			return nil, err
		}
		total := int(nargs) + int(nrets) + 1
		node.Children = make([]*Node, total)
		for i := 0; i < total; i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[i] = c
		}
	case OpDeclStmt:
		// Kernel decoder (ast_decode.c) reads NameId (U16) THEN Kind (U8).
		nameID, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.NameID = nameID
		kind, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Kind = kind
		expr, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{expr}
	case OpArrayType:
		// 4 bytes len + 1 byte elem_kind
		length, err := d.readU32()
		if err != nil {
			return nil, err
		}
		elemKind, err := d.readByte()
		if err != nil {
			return nil, err
		}
		node.Value = make([]byte, 5)
		binary.LittleEndian.PutUint32(node.Value[:4], length)
		node.Value[4] = elemKind
	case OpCompositeLit:
		// Type + 2 bytes nels + N elements
		typ, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		nels, err := d.readU16()
		if err != nil {
			return nil, err
		}
		node.Children = make([]*Node, 1+int(nels))
		node.Children[0] = typ
		for i := 0; i < int(nels); i++ {
			c, err := d.decodeNode()
			if err != nil {
				return nil, err
			}
			node.Children[1+i] = c
		}
	case OpIndexExpr:
		x, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		idx, err := d.decodeNode()
		if err != nil {
			return nil, err
		}
		node.Children = []*Node{x, idx}
	default:
		return nil, NewDecodeError(uint32(d.pos-1), fmt.Sprintf("unknown opcode 0x%02x", op))
	}
	return node, nil
}

func (d *Decoder) decodeLiteralValue(kind byte) ([]byte, error) {
	switch kind {
	case KindUint8, KindInt8, KindBool:
		return d.readBytes(1)
	case KindUint16, KindInt16:
		return d.readBytes(2)
	case KindUint32, KindInt32:
		return d.readBytes(4)
	case KindUint64, KindInt64:
		return d.readBytes(8)
	case KindString:
		n, err := d.readU16()
		if err != nil {
			return nil, err
		}
		return d.readBytes(int(n))
	}
	return nil, NewDecodeError(uint32(d.pos), fmt.Sprintf("unknown literal kind 0x%02x", kind))
}
