// Package protocol defines the AST binary protocol shared between the Go
// user-mode compiler (go-bridge/ast) and the kernel C interpreter
// (hyperkd/code/go-interp). Both sides MUST agree byte-for-byte with these
// constants.
//
// See docs/go-subset-spec.md for the full specification.
package protocol

// Node opcodes — first byte of every encoded node.
const (
	OpNil          byte = 0x00 // empty node (absent Else)
	OpLiteral      byte = 0x01 // 1 byte kind + N bytes value
	OpIdent        byte = 0x02 // 2 bytes name_id
	OpBinaryExpr   byte = 0x03 // 1 byte op + LHS + RHS
	OpUnaryExpr    byte = 0x04 // 1 byte op + Operand
	OpCallExpr     byte = 0x05 // 2 bytes func_id + 1 byte nargs + N children
	OpSelectorExpr byte = 0x06 // X + 2 bytes field_id
	OpAssignStmt   byte = 0x07 // 1 byte op + LHS + RHS
	OpIfStmt       byte = 0x08 // Cond + Then + Else
	OpForStmt      byte = 0x09 // Init + Cond + Post + Body
	OpBlockStmt    byte = 0x0A // 2 bytes nstmts + N statements
	OpReturnStmt   byte = 0x0B // 1 byte nvals + N children
	OpFuncDecl     byte = 0x0C // 2 bytes name_id + 1 byte nargs + N args + 1 byte nrets + N rets + Body
	OpFuncLit      byte = 0x0D // 1 byte nargs + N args + 1 byte nrets + N rets + Body
	OpDeclStmt     byte = 0x0E // 2 bytes name_id + 1 byte kind + Expr
	OpArrayType    byte = 0x0F // 4 bytes len + 1 byte elem_kind
	OpCompositeLit byte = 0x10 // Type + 2 bytes nels + N elements
	OpIndexExpr    byte = 0x11 // X + Index
)

// Literal kinds — second byte of an OpLiteral node.
const (
	KindUint8  byte = 0x01 // 1 byte value
	KindUint16 byte = 0x02 // 2 bytes value (LE)
	KindUint32 byte = 0x03 // 4 bytes value (LE)
	KindUint64 byte = 0x04 // 8 bytes value (LE)
	KindInt8   byte = 0x05 // 1 byte value
	KindInt16  byte = 0x06 // 2 bytes value (LE)
	KindInt32  byte = 0x07 // 4 bytes value (LE)
	KindInt64  byte = 0x08 // 8 bytes value (LE)
	KindBool   byte = 0x09 // 1 byte (0=false, 1=true)
	KindString byte = 0x0A // 2 bytes len + N bytes UTF-8
)

// Binary operators — second byte of an OpBinaryExpr node.
const (
	BinOpAdd    byte = 0x01 // +
	BinOpSub    byte = 0x02 // -
	BinOpMul    byte = 0x03 // *
	BinOpQuo    byte = 0x04 // /
	BinOpRem    byte = 0x05 // %
	BinOpAnd    byte = 0x06 // &
	BinOpOr     byte = 0x07 // |
	BinOpXor    byte = 0x08 // ^
	BinOpShl    byte = 0x09 // <<
	BinOpShr    byte = 0x0A // >>
	BinOpEql    byte = 0x0B // ==
	BinOpNeq    byte = 0x0C // !=
	BinOpLss    byte = 0x0D // <
	BinOpGtr    byte = 0x0E // >
	BinOpLeq    byte = 0x0F // <=
	BinOpGeq    byte = 0x10 // >=
	BinOpLAnd   byte = 0x11 // &&
	BinOpLOr    byte = 0x12 // ||
	BinOpAndNot byte = 0x13 // &^
)

// Unary operators — second byte of an OpUnaryExpr node.
const (
	UnOpNeg byte = 0x01 // - (negate)
	UnOpNot byte = 0x02 // ! (logical not)
	UnOpXor byte = 0x03 // ^ (bitwise not)
)

// Assignment operators — second byte of an OpAssignStmt node.
const (
	AssignAssign byte = 0x01 // =
	AssignAdd    byte = 0x02 // +=
	AssignSub    byte = 0x03 // -=
	AssignMul    byte = 0x04 // *=
	AssignQuo    byte = 0x05 // /=
	AssignRem    byte = 0x06 // %=
	AssignAnd    byte = 0x07 // &=
	AssignOr     byte = 0x08 // |=
	AssignXor    byte = 0x09 // ^=
	AssignShl    byte = 0x0A // <<=
	AssignShr    byte = 0x0B // >>=
	AssignAndNot byte = 0x0C // &^=
)

// Whitelist function IDs — the 2-byte func_id in OpCallExpr when the callee is
// a SelectorExpr ctx.Method. See docs/go-subset-spec.md §3.
const (
	FuncStackReadQword   uint16 = 0x0001
	FuncStackReadDword   uint16 = 0x0002
	FuncReg              uint16 = 0x0003
	FuncSetReg           uint16 = 0x0004
	FuncReadMem          uint16 = 0x0005
	FuncReadMemQword     uint16 = 0x0006
	FuncPrintf           uint16 = 0x0007
	FuncBreak            uint16 = 0x0008
	FuncContinue         uint16 = 0x0009
	FuncGetPid           uint16 = 0x000A
	FuncGetTid           uint16 = 0x000B
	FuncGetIP            uint16 = 0x000C
	FuncSetCtxVar        uint16 = 0x000D
	FuncGetCtxVar        uint16 = 0x000E
	FuncSetMemWriteBp    uint16 = 0x000F
	FuncReadInstrOperand uint16 = 0x0010
)

// WhitelistFuncs maps ctx method names to their func_id. Used by the encoder
// to resolve SelectorExpr field names.
var WhitelistFuncs = map[string]uint16{
	"StackReadQword":   FuncStackReadQword,
	"StackReadDword":   FuncStackReadDword,
	"Reg":              FuncReg,
	"SetReg":           FuncSetReg,
	"ReadMem":          FuncReadMem,
	"ReadMemQword":     FuncReadMemQword,
	"Printf":           FuncPrintf,
	"Break":            FuncBreak,
	"Continue":         FuncContinue,
	"GetPid":           FuncGetPid,
	"GetTid":           FuncGetTid,
	"GetIP":            FuncGetIP,
	"SetCtxVar":        FuncSetCtxVar,
	"GetCtxVar":        FuncGetCtxVar,
	"SetMemWriteBp":    FuncSetMemWriteBp,
	"ReadInstrOperand": FuncReadInstrOperand,
}

// CycleBudget is the maximum number of AST nodes the kernel interpreter will
// execute per hook callback before aborting with Continue. See spec §4.
const CycleBudget = 1_000_000

// MaxNStatements is the upper bound on statements in a BlockStmt (2-byte
// count field → 65535, but we cap lower for safety).
const MaxNStatements = 65535

// MaxNArgs is the upper bound on function arguments/parameters.
const MaxNArgs = 255

// MaxNElems is the upper bound on composite literal elements.
const MaxNElems = 65535

// MaxStringLength is the upper bound on a single string in the string table.
const MaxStringLength = 65535
