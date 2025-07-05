package clang

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/valyala/fastjson"
)

type refTracker struct {
	refs map[string][]*NodeDecl
	defs map[string]Node
}

type valueUnmarshaller interface {
	UnmarshalValue(rt *refTracker, data *fastjson.Value) error
}
type Location struct {
	Offset   int    `json:"offset,omitempty"`
	Line     int    `json:"line,omitempty"`
	Col      int    `json:"col,omitempty"`
	TokenLen int    `json:"tokLen,omitempty"`
	File     string `json:"file,omitempty"`
}

func parseLocation(n *fastjson.Value) *Location {
	if n == nil {
		return nil
	}
	return &Location{
		Offset:   n.GetInt("offset"),
		Line:     n.GetInt("line"),
		Col:      n.GetInt("col"),
		TokenLen: n.GetInt("tokLen"),
		File:     string(n.GetStringBytes("file")),
	}
}

type SourceRange struct {
	Begin        *Location `json:"begin,omitempty"`
	End          *Location `json:"end,omitempty"`
	SpellingLoc  *Location `json:"spellingLoc,omitempty"`
	ExpansionLoc *Location `json:"expansionLoc,omitempty"`
}

func parseSourceRange(n *fastjson.Value) *SourceRange {
	if n == nil {
		return nil
	}
	return &SourceRange{
		Begin:        parseLocation(n.Get("begin")),
		End:          parseLocation(n.Get("end")),
		SpellingLoc:  parseLocation(n.Get("spellingLoc")),
		ExpansionLoc: parseLocation(n.Get("expansionLoc")),
	}
}

type Node interface {
	Base() *baseNode
	ID() string
	Kind() string
	Location() *Location
	SrcRange() *SourceRange
	Children() []Node
	At(i int) Node
	Comment() string
}

func First[T Node](n Node) (res T) {
	if n == nil {
		return
	}
	for _, c := range n.Children() {
		if res, ok := c.(T); ok {
			return res
		}
	}
	return
}

func All[T Node](n Node) (res []T) {
	if n == nil {
		return
	}
	for _, c := range n.Children() {
		if c, ok := c.(T); ok {
			res = append(res, c)
		}
	}
	return
}

func Visit[T Node](n Node, f func(T) bool) {
	if n == nil {
		return
	}
	for _, c := range n.Children() {
		if c, ok := c.(T); ok {
			if f(c) {
				Visit(c, f)
			}
		}
	}
}

type baseNode struct {
	NodeID   string       `json:"id"`
	NodeKind string       `json:"kind"`
	Loc      *Location    `json:"loc,omitempty"`
	Range    *SourceRange `json:"range,omitempty"`
	Inner    []Node       `json:"inner,omitempty"`
}

func (n *baseNode) Comment() string {
	builder := &strings.Builder{}
	for _, c := range n.Inner {
		if c, ok := c.(CommentNode); ok {
			c.AppendTo(builder)
		}
	}
	return builder.String()
}

func (n *baseNode) unmarshal(rt *refTracker, data *fastjson.Value) (err error) {
	n.NodeID = string(data.GetStringBytes("id"))
	n.NodeKind = string(data.GetStringBytes("kind"))
	n.Loc = parseLocation(data.Get("loc"))
	n.Range = parseSourceRange(data.Get("range"))
	arr := data.GetArray("inner")
	n.Inner = make([]Node, len(arr))
	for i, v := range arr {
		node, e := unmarshalNode(rt, v)
		if e != nil {
			mylog.CheckIgnore(e)
			return nil
			return e
		}
		n.Inner[i] = node
	}
	return nil
}

func (n *baseNode) ID() string {
	return n.NodeID
}

func (n *baseNode) Kind() string {
	return n.NodeKind
}

func (n *baseNode) Location() *Location {
	return n.Loc
}

func (n *baseNode) Children() []Node {
	return n.Inner
}

func (n *baseNode) SrcRange() *SourceRange {
	return n.Range
}

func (n *baseNode) At(i int) Node {
	if i < 0 || i >= len(n.Inner) {
		return nil
	}
	return n.Inner[i]
}

func (n *baseNode) Base() *baseNode {
	return n
}

type UnknownNode struct {
	baseNode
	Raw json.RawMessage
}

func (n *UnknownNode) UnmarshalValue(rt *refTracker, data *fastjson.Value) error {
	n.Raw = data.MarshalTo(nil)
	return n.baseNode.unmarshal(rt, data)
}

type Type struct {
	QualType          string  `json:"qualType"`
	TypeAliasDeclId   string  `json:"typeAliasDeclId,omitempty"`
	DesugaredQualType *string `json:"desugaredQualType,omitempty"`
}

func (n *Type) UnmarshalValue(data *fastjson.Value) (err error) {
	n.QualType = string(data.GetStringBytes("qualType"))
	if v := data.Get("desugaredQualType"); v != nil {
		s := string(v.GetStringBytes())
		n.DesugaredQualType = &s
	}
	if v := data.Get("typeAliasDeclId"); v != nil {
		n.TypeAliasDeclId = string(v.GetStringBytes())
	}
	return nil
}

type NodeDecl struct {
	ID   string `json:"id,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	Type *Type  `json:"type,omitempty"`
	Ref  Node   `json:"-"`
}

func (n *NodeDecl) UnmarshalValue(rt *refTracker, data *fastjson.Value) (err error) {
	n.Name = string(data.GetStringBytes("name"))
	n.Kind = string(data.GetStringBytes("kind"))
	n.ID = string(data.GetStringBytes("id"))
	if v := data.Get("type"); v != nil {
		n.Type = &Type{}
		mylog.Check(n.Type.UnmarshalValue(v))
	}
	rt.refs[n.ID] = append(rt.refs[n.ID], n)
	if n.Name == "" {
		mylog.Trace(n.Kind, n.Name) // name是空的，需要cc现代解析器，生成的枚举没有名称，语法错误，调整原始头文件的名称位置麻烦且不通用
	}
	return
}

type TranslationUnitDecl struct {
	baseNode
}
type TypeVisibilityAttr struct {
	baseNode
	Implicit bool `json:"implicit"`
}
type StaticAssertDecl struct {
	baseNode
}
type MaxFieldAlignmentAttr struct {
	baseNode
	Implicit bool `json:"implicit"`
}
type CompoundStmt struct {
	baseNode
}
type ReturnStmt struct {
	baseNode
}
type IfStmt struct {
	baseNode
}

//
// Value nodes.
//

type ValueNode interface {
	Node
	Category() string
	Type() Type
}

type ConstValueNode interface {
	ValueNode
	Value() string
}

type baseValueNode struct {
	baseNode
	ValueCategory string `json:"valueCategory"`
	ValueType     Type   `json:"type"`
}

func (n *baseValueNode) Type() Type {
	return n.ValueType
}

func (n *baseValueNode) Category() string {
	return n.ValueCategory
}

type IntegerLiteral struct {
	ConstantExpr
}
type StringLiteral struct {
	ConstantExpr
}
type ConstantExpr struct {
	baseValueNode
	ExpressionValue string `json:"value"`
}

func (l *ConstantExpr) Value() string { return l.ExpressionValue }

type ImplicitCastExpr struct {
	baseValueNode
	CastKind             string `json:"castKind"`
	IsPartOfExplicitCast bool   `json:"isPartOfExplicitCast"`
}
type BinaryOperator struct {
	baseValueNode
	Opcode string `json:"opcode"`
}
type UnaryExprOrTypeTraitExpr struct {
	baseValueNode
	Name    string `json:"name"`
	ArgType Type   `json:"argType"`
}
type CStyleCastExpr struct {
	baseValueNode
	CastKind string `json:"castKind"`
}
type UnaryOperator struct {
	baseValueNode
	CanOverflow bool   `json:"canOverflow"`
	IsPostfix   bool   `json:"isPostfix"`
	Opcode      string `json:"opcode"`
}
type ParenExpr struct {
	baseValueNode
}
type DeclRefExpr struct {
	baseValueNode
	ReferencedDecl NodeDecl `json:"referencedDecl"`
}
type ConditionalOperator struct {
	baseValueNode
}

//
// Comment nodes.
//

type CommentNode interface {
	Node
	AppendTo(wr *strings.Builder)
}

type baseCommentNode struct {
	baseNode
}

func (n *baseCommentNode) AppendTo(wr *strings.Builder) {
	for _, c := range n.Inner {
		if c, ok := c.(CommentNode); ok {
			c.AppendTo(wr)
		}
	}
}

type FullComment struct {
	baseCommentNode
}
type ParagraphComment struct {
	baseCommentNode
}
type TextComment struct {
	baseCommentNode
	Text string `json:"text"`
}

func (n *TextComment) AppendTo(wr *strings.Builder) {
	t := strings.TrimSpace(n.Text)
	wr.WriteString(t)
	if !strings.HasSuffix(t, "\n") {
		wr.WriteString("\n")
	}
}

type ParamCommandComment struct {
	baseCommentNode
	Direction string `json:"direction"`
	Param     string `json:"param"`
	ParamIdx  int    `json:"paramIdx"`
}

func (n *ParamCommandComment) AppendTo(wr *strings.Builder) {
	wr.WriteString("@param ")
	wr.WriteString(n.Param)
	wr.WriteString(" ")
	n.baseCommentNode.AppendTo(wr)
}

type BlockCommandComment struct {
	baseCommentNode
	Name string `json:"name"`
}

func (n *BlockCommandComment) AppendTo(wr *strings.Builder) {
	wr.WriteString("@")
	wr.WriteString(n.Name)
	wr.WriteString(" ")
	n.baseCommentNode.AppendTo(wr)
}

type InlineCommandComment struct {
	baseCommentNode
	Name       string   `json:"name"`
	RenderKind string   `json:"renderKind"`
	Args       []string `json:"args"`
}

func (n *InlineCommandComment) AppendTo(wr *strings.Builder) {
	wr.WriteString("@")
	wr.WriteString(n.Name)
	for _, a := range n.Args {
		wr.WriteString(" ")
		wr.WriteString(a)
	}
	wr.WriteString("\n")
	n.baseCommentNode.AppendTo(wr)
}

//
// Type nodes.
//

type TypeNode interface {
	Node
	QualifiedType() string
	TypeAliasDeclID() string
	DesugaredQualifiedType() string
}
type baseTypeNode struct {
	baseNode
	Type Type `json:"type"`
}

func (n *baseTypeNode) QualifiedType() string {
	return n.Type.QualType
}

func (n *baseTypeNode) TypeAliasDeclID() string {
	return n.Type.TypeAliasDeclId
}

func (n *baseTypeNode) DesugaredQualifiedType() string {
	s := n.Type.DesugaredQualType
	if s == nil {
		return ""
	}
	return *s
}

type BuiltinType struct {
	baseTypeNode
}
type RecordType struct {
	baseTypeNode
	Decl NodeDecl `json:"decl"`
}
type PointerType struct {
	baseTypeNode
}
type QualType struct {
	baseTypeNode
	Qualifiers string `json:"qualifiers"`
}
type ElaboratedType struct {
	baseTypeNode
	OwnedTagDecl NodeDecl `json:"ownedTagDecl"`
}
type TypedefType struct {
	baseTypeNode
	Decl NodeDecl `json:"decl"`
}
type EnumType struct {
	baseTypeNode
	Decl NodeDecl `json:"decl"`
}
type ParenType struct {
	baseTypeNode
}
type FunctionProtoType struct {
	baseTypeNode
	Cc string `json:"cc"`
}

//
// Decl nodes.
//

type DeclNode interface {
	Node
	DeclName() string
}
type baseDeclNode struct {
	baseNode
	Name string `json:"name"`
}

func (n *baseDeclNode) DeclName() string {
	return n.Name
}

type EnumDecl struct {
	baseDeclNode
}
type EnumConstantDecl struct {
	baseDeclNode
	Type         Type `json:"type"`
	IsReferenced bool `json:"isReferenced"`
}
type RecordDecl struct {
	baseDeclNode
	IsImplicit         bool   `json:"isImplicit"`
	TagUsed            string `json:"tagUsed"`
	CompleteDefinition bool   `json:"completeDefinition"`
	ParentDeclContext  string `json:"parentDeclContext"`
	PreviousDecl       string `json:"previousDecl"`
}
type TypedefDecl struct {
	baseDeclNode
	IsImplicit   bool `json:"isImplicit"`
	Type         Type `json:"type"`
	IsReferenced bool `json:"isReferenced"`
}
type FunctionDecl struct {
	baseDeclNode
	MangledName  string `json:"mangledName"`
	Type         Type   `json:"type"`
	StorageClass string `json:"storageClass"`
	Inline       bool   `json:"inline"`
}
type ParmVarDecl struct {
	baseDeclNode
	Type   Type `json:"type"`
	IsUsed bool `json:"isUsed"`
}
type FieldDecl struct {
	baseDeclNode
	Type       Type `json:"type"`
	IsImplicit bool `json:"isImplicit"`
}
type IndirectFieldDecl struct {
	baseDeclNode
	IsImplicit bool `json:"isImplicit"`
}

func makeUnmarshallerFor(ty reflect.Type) func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
	if reflect.PointerTo(ty).Implements(reflect.TypeFor[valueUnmarshaller]()) {
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			return reflect.NewAt(ty, out).Interface().(valueUnmarshaller).UnmarshalValue(rt, val)
		}
	}
	if ty == reflect.TypeFor[baseNode]() {
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			return (*baseNode)(out).unmarshal(rt, val)
		}
	}
	if ty == reflect.TypeFor[Node]() {
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			node := mylog.Check2(unmarshalNode(rt, val))

			*(*Node)(out) = node
			return nil
		}
	}

	switch ty.Kind() {
	case reflect.String:
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			*(*string)(out) = string(val.GetStringBytes())
			return nil
		}
	case reflect.Bool:
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			*(*bool)(out) = val.GetBool()
			return nil
		}
	case reflect.Int:
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			*(*int)(out) = val.GetInt()
			return nil
		}
	case reflect.Slice:
		elemTy := ty.Elem()
		elemUnmarshal := makeUnmarshallerFor(elemTy)
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			arr := val.GetArray()
			slice := reflect.MakeSlice(ty, len(arr), len(arr))
			for i, v := range arr {
				elemUnmarshal(rt, v, unsafe.Pointer(slice.Index(i).UnsafeAddr()))
			}
			reflect.NewAt(slice.Type(), out).Elem().Set(slice)
			return nil
		}
	case reflect.Pointer:
		elemTy := ty.Elem()
		elemUnmarshal := makeUnmarshallerFor(elemTy)
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			if val == nil {
				*(*unsafe.Pointer)(out) = nil
				return nil
			}
			ptr := reflect.New(elemTy)
			elemUnmarshal(rt, val, unsafe.Pointer(ptr.Pointer()))
			*(*unsafe.Pointer)(out) = unsafe.Pointer(ptr.Pointer())
			return nil
		}
	case reflect.Struct:
		var unmarshalSteps []func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error
		for i := range ty.NumField() {
			field := ty.Field(i)
			if field.Anonymous {
				fieldUnmarshal := makeUnmarshallerFor(field.Type)
				unmarshalSteps = append(unmarshalSteps, func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
					return fieldUnmarshal(rt, val, out)
				})
				continue
			}
			fieldUnmarshal := makeUnmarshallerFor(field.Type)
			fieldName := field.Tag.Get("json")
			fieldName, _, _ = strings.Cut(fieldName, ",")
			if fieldName == "" || fieldName == "-" {
				continue
			}

			unmarshalSteps = append(unmarshalSteps, func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
				f := val.Get(fieldName)
				if f == nil {
					return nil
				}
				return fieldUnmarshal(rt, f, unsafe.Pointer(uintptr(out)+field.Offset))
			})
		}
		return func(rt *refTracker, val *fastjson.Value, out unsafe.Pointer) error {
			for _, f := range unmarshalSteps {
				mylog.Check(f(rt, val, out))
			}
			return nil
		}
	}
	panic("unsupported type " + ty.String())
}

func RegisterNode[T Node]() {
	ty := reflect.TypeFor[T]()
	if ty.Kind() != reflect.Pointer {
		panic("expected pointer type")
	}

	ty = ty.Elem()
	unmarshaller := makeUnmarshallerFor(ty)
	NodeKinds[ty.Name()] = func(rt *refTracker, val *fastjson.Value) (Node, error) {
		ptr := reflect.New(ty)
		mylog.Check(unmarshaller(rt, val, unsafe.Pointer(ptr.Pointer())))

		node := ptr.Interface().(Node)
		rt.defs[node.ID()] = node
		return node, nil
	}
}

// Unmarshaller.
func unmarshalNode(rt *refTracker, data *fastjson.Value) (Node, error) {
	if data == nil {
		return nil, nil
	}
	if data.Get("kind") == nil {
		return nil, errors.New("expected node.kind")
	}
	ty := NodeKinds[string(data.GetStringBytes("kind"))]
	if ty == nil {
		res := &UnknownNode{}
		return res, res.UnmarshalValue(rt, data)
	} else {
		return ty(rt, data)
	}
}

func ParseAST(data []byte) (node Node, err error) {
	iter := mylog.Check2(fastjson.ParseBytes(data))

	rt := &refTracker{
		refs: map[string][]*NodeDecl{},
		defs: map[string]Node{},
	}
	node = mylog.Check2(unmarshalNode(rt, iter))
	for id, def := range rt.defs {
		for _, ref := range rt.refs[id] {
			ref.Ref = def
		}
	}
	return
}

// Printers
func Fprint(w io.Writer, n Node) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(n)
}

func Print(n Node) error {
	return Fprint(io.Writer(os.Stdout), n)
}

var NodeKinds = map[string]func(rt *refTracker, val *fastjson.Value) (Node, error){}

func init() {
	RegisterNode[*TranslationUnitDecl]()
	RegisterNode[*RecordDecl]()
	RegisterNode[*TypeVisibilityAttr]()
	RegisterNode[*TypedefDecl]()
	RegisterNode[*BuiltinType]()
	RegisterNode[*RecordType]()
	RegisterNode[*PointerType]()
	RegisterNode[*StaticAssertDecl]()
	RegisterNode[*ImplicitCastExpr]()
	RegisterNode[*BinaryOperator]()
	RegisterNode[*UnaryExprOrTypeTraitExpr]()
	RegisterNode[*IntegerLiteral]()
	RegisterNode[*StringLiteral]()
	RegisterNode[*CStyleCastExpr]()
	RegisterNode[*UnaryOperator]()
	RegisterNode[*ParenExpr]()
	RegisterNode[*FullComment]()
	RegisterNode[*ParagraphComment]()
	RegisterNode[*TextComment]()
	RegisterNode[*QualType]()
	RegisterNode[*ElaboratedType]()
	RegisterNode[*TypedefType]()
	RegisterNode[*EnumDecl]()
	RegisterNode[*EnumConstantDecl]()
	RegisterNode[*ConstantExpr]()
	RegisterNode[*DeclRefExpr]()
	RegisterNode[*EnumType]()
	RegisterNode[*FunctionDecl]()
	RegisterNode[*ParmVarDecl]()
	RegisterNode[*ParamCommandComment]()
	RegisterNode[*BlockCommandComment]()
	RegisterNode[*MaxFieldAlignmentAttr]()
	RegisterNode[*FieldDecl]()
	RegisterNode[*IndirectFieldDecl]()
	RegisterNode[*InlineCommandComment]()
	RegisterNode[*ParenType]()
	RegisterNode[*FunctionProtoType]()
	RegisterNode[*CompoundStmt]()
	RegisterNode[*ReturnStmt]()
	RegisterNode[*ConditionalOperator]()
	RegisterNode[*IfStmt]()
}
