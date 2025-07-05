package gengo

import (
	"fmt"
	"go/token"
	"maps"
	"strconv"
	"strings"

	"github.com/dave/dst"
	"github.com/ddkwork/golibrary/std/mylog"
)

type TypeRef struct {
	Identifier
	Decl *dst.GenDecl
}

type TypeClass int

const (
	TcEnum TypeClass = iota
	TcStruct
	TcUnion
	TcTypedef
	TcCount
)

// Builtin types.
const (
	BuiltinVoid       = RelaxedIdentifier("void")
	BuiltinInt8       = RelaxedIdentifier("int8")
	BuiltinInt16      = RelaxedIdentifier("int16")
	BuiltinInt32      = RelaxedIdentifier("int32")
	BuiltinInt64      = RelaxedIdentifier("int64")
	BuiltinUint8      = RelaxedIdentifier("uint8")
	BuiltinUint16     = RelaxedIdentifier("uint16")
	BuiltinUint32     = RelaxedIdentifier("uint32")
	BuiltinUint64     = RelaxedIdentifier("uint64")
	BuiltinUint       = RelaxedIdentifier("uint")
	BuiltinInt        = RelaxedIdentifier("int")
	BuiltinFloat32    = RelaxedIdentifier("float32")
	BuiltinFloat64    = RelaxedIdentifier("float64")
	BuiltinComplex64  = RelaxedIdentifier("complex64")
	BuiltinComplex128 = RelaxedIdentifier("complex128")
	BuiltinBool       = RelaxedIdentifier("bool")
	BuiltinByte       = RelaxedIdentifier("byte")
	BuiltinRune       = RelaxedIdentifier("rune")
	BuiltinUintptr    = RelaxedIdentifier("uintptr")
	BuiltinAny        = RelaxedIdentifier("any")
)

// Reserved identifiers.
var reservedIdentifiers = map[string]struct{}{
	"break":       {},
	"default":     {},
	"func":        {},
	"interface":   {},
	"select":      {},
	"case":        {},
	"defer":       {},
	"go":          {},
	"map":         {},
	"struct":      {},
	"chan":        {},
	"else":        {},
	"goto":        {},
	"package":     {},
	"switch":      {},
	"const":       {},
	"fallthrough": {},
	"if":          {},
	"range":       {},
	"type":        {},
	"continue":    {},
	"for":         {},
	"import":      {},
	"return":      {},
	"var":         {},
	"error":       {},
	"string":      {},
	"true":        {},
	"false":       {},
	"iota":        {},
	"nil":         {},
	"append":      {},
	"cap":         {},
	"close":       {},
	"complex":     {},
	"copy":        {},
	"delete":      {},
	"imag":        {},
	"len":         {},
	"make":        {},
	"new":         {},
	"panic":       {},
	"print":       {},
	"println":     {},
	"real":        {},
	"recover":     {},
}

// Default mappings of builtin C types.
var defaultBuiltinMap = map[string]Identifier{
	"void":                 BuiltinVoid,
	"char":                 BuiltinByte,
	"short":                BuiltinInt16,
	"int":                  BuiltinInt32,
	"long":                 BuiltinInt64,
	"long long":            BuiltinInt64,
	"signed char":          BuiltinInt8,
	"unsigned char":        BuiltinUint8,
	"signed short":         BuiltinInt16,
	"unsigned short":       BuiltinUint16,
	"signed int":           BuiltinInt32,
	"unsigned int":         BuiltinUint32,
	"signed long":          BuiltinInt64,
	"unsigned long":        BuiltinUint64,
	"signed long long":     BuiltinInt64,
	"unsigned long long":   BuiltinUint64,
	"signed":               BuiltinInt32,
	"unsigned":             BuiltinUint32,
	"float":                BuiltinFloat32,
	"double":               BuiltinFloat64,
	"long double":          BuiltinFloat64,
	"size_t":               BuiltinUint,
	"uint8_t":              BuiltinUint8,
	"uint16_t":             BuiltinUint16,
	"uint32_t":             BuiltinUint32,
	"uint64_t":             BuiltinUint64,
	"int8_t":               BuiltinInt8,
	"int16_t":              BuiltinInt16,
	"int32_t":              BuiltinInt32,
	"int64_t":              BuiltinInt64,
	"uintptr_t":            BuiltinUintptr,
	"intmax_t":             BuiltinInt64,
	"uintmax_t":            BuiltinUint64,
	"intptr_t":             BuiltinInt,
	"bool":                 BuiltinBool,
	"_Bool":                BuiltinBool,
	"char16_t":             BuiltinInt16,
	"char32_t":             BuiltinInt32,
	"float _Complex":       BuiltinComplex64,
	"double _Complex":      BuiltinComplex128,
	"long double _Complex": BuiltinComplex128,
}

// Provider is a type provider allowing hooks for library specific type conversions.
type Provider interface {
	// NameField converts a field name to a Go compatible name.
	NameField(name string, recordName string) string
	// NameType converts a type name to a Go compatible name.
	NameType(name string) string
	// NameValue converts a value name to a Go compatible name.
	NameValue(name string) string
	// NameFunc converts a function name to a Go compatible name.
	NameFunc(name string) string
	// NameArg converts an argument name to a Go compatible name.
	NameArg(name string, argType, funcName string) string
	// NameGetter takes a transformed field name and returns the getter name.
	NameGetter(name string) string
	// NameSetter takes a transformed field name and returns the setter name.
	NameSetter(name string) string

	// ForceSynthetic returns true if the type should be syntethic (accessors instead of direct fields).
	ForceSynthetic(name string) bool
	// ConvertQualType converts a qualified type to a Go expression.
	ConvertQualType(q string) dst.Expr
	// ConvertTypeExpr converts a type node to a Go expression.
	ConvertTypeExpr(n clang.TypeNode) dst.Expr

	// AddType adds a type to the provider.
	AddType(tc TypeClass, name string, decl dst.Expr) TypeRef
	// RemapType remaps a given qualified typename to a type reference.
	RemapType(name string, tr TypeRef)
	// FindType resolves a type reference from a Go name.
	FindType(name string) (TypeRef, bool)

	// InferMethod returns the receiver type given the method name.
	InferMethod(name string) (rcv string, newName string)
}

// Normalize anonymous type names.
func normalizeAnonName(name string) string {
	if strings.IndexByte(name, ')') != -1 {
		// struct ZydisFormatter_::(anonymous at ../rsrc/Zydis.h:11222:5)
		tag, rest, _ := strings.Cut(name, " ")
		_, pos, _ := strings.Cut(rest, " at ")
		return tag + "@" + pos
	}
	return name
}

type BaseProviderOption func(*BaseProvider)

func WithRemovePrefix(prefixes ...string) BaseProviderOption {
	return func(p *BaseProvider) {
		p.RemovedPrefixes = append(p.RemovedPrefixes, prefixes...)
	}
}

func WithInferredMethods(rules []MethodInferenceRule) BaseProviderOption {
	return func(p *BaseProvider) {
		p.InferredMethods = append(p.InferredMethods, rules...)
	}
}

func WithForcedSynthetic(names ...string) BaseProviderOption {
	return func(p *BaseProvider) {
		for _, name := range names {
			p.ForcedSynthetic[name] = struct{}{}
		}
	}
}

type MethodInferenceRule struct {
	Name     string
	Receiver string
}

type BaseProvider struct {
	Types           map[string]TypeRef
	Builtins        map[string]Identifier
	RemovedPrefixes []string
	InferredMethods []MethodInferenceRule
	ForcedSynthetic map[string]struct{}
}

func NewBaseProvider(opt ...BaseProviderOption) *BaseProvider {
	dc := &BaseProvider{
		Types:           map[string]TypeRef{},
		Builtins:        map[string]Identifier{},
		InferredMethods: []MethodInferenceRule{},
		ForcedSynthetic: map[string]struct{}{},
	}
	maps.Copy(dc.Builtins, defaultBuiltinMap)
	for _, o := range opt {
		o(dc)
	}
	return dc
}

func (p *BaseProvider) removePrefixes(name string) string {
	for _, prefix := range p.RemovedPrefixes {
		if len(name) > len(prefix) && strings.EqualFold(name[:len(prefix)], prefix) {
			return name[len(prefix):]
		}
	}
	return name
}

func (p *BaseProvider) NameGetter(name string) string {
	if name == "" {
		name = "GetMust"
		mylog.Warning("Empty field name, using default getter name", name)
	}
	return name
}

func (p *BaseProvider) NameSetter(name string) string {
	return "Set" + name
}

func (p *BaseProvider) NameField(name string, recordName string) string {
	name = p.removePrefixes(name)
	return ConvertCase(name, ConventionPascalCase)
}

func (p *BaseProvider) NameFunc(name string) string {
	name = p.removePrefixes(name)
	return ConvertCase(name, ConventionPascalCase)
}

func (p *BaseProvider) NameType(name string) string {
	name = p.removePrefixes(name)
	return ConvertCase(name, ConventionPascalCase)
}

func (p *BaseProvider) NameValue(name string) string {
	name = p.removePrefixes(name)
	return ConvertCase(name, ConventionConstCase)
}

func (p *BaseProvider) NameArg(name string, argType, funcName string) string {
	if _, ok := reservedIdentifiers[name]; ok {
		return "_" + name
	}
	return name
}

func (p *BaseProvider) ForceSynthetic(name string) bool {
	if _, ok := p.ForcedSynthetic[name]; ok {
		return true
	}
	return false
}

func (p *BaseProvider) ConvertQualType(q string) dst.Expr {
	// Dumb qualifiers.
	q = strings.TrimSpace(q)
	q = strings.ReplaceAll(q, "const ", "")
	q = strings.ReplaceAll(q, "volatile ", "")
	q = strings.ReplaceAll(q, "restrict ", "")
	q = strings.TrimSpace(q)

	// Pointer type.
	if q, ok := strings.CutSuffix(q, "*"); ok {
		res := p.ConvertQualType(q)
		if ident, ok := res.(*dst.Ident); ok && (ident.Name == "any" || ident.Name == "void") {
			return &dst.SelectorExpr{
				X:   dst.NewIdent("unsafe"),
				Sel: dst.NewIdent("Pointer"),
			}
		}
		return &dst.StarExpr{
			X: res,
		}
	}

	// Array type.
	if q, ok := strings.CutSuffix(q, "]"); ok {
		lastBracker := strings.LastIndex(q, "[")
		if lastBracker == -1 {
			fmt.Printf("[WARN] Invalid array type: %s\n", q)
			return BuiltinAny.Ref()
		}
		before := q[:lastBracker]
		after := q[lastBracker+1:]
		n := mylog.Check2(strconv.Atoi(after))

		return &dst.ArrayType{
			Len: dst.NewIdent(strconv.Itoa(n)),
			Elt: p.ConvertQualType(before),
		}
	}

	// Builtin types.
	if b, ok := p.Builtins[q]; ok {
		return b.Ref()
	}

	// Normalize anonymous type names.
	q = normalizeAnonName(q)

	// Named types.
	if td, ok := p.Types[q]; ok {
		return td.Ref()
	}

	// Unknown type.
	mylog.Trace("Unknown type", q)
	return BuiltinAny.Ref()
}

func (p *BaseProvider) ConvertTypeExpr(n clang.TypeNode) dst.Expr {
	switch n := n.(type) {
	case *clang.BuiltinType:
		return p.ConvertQualType(n.Type.QualType)
	case *clang.PointerType:
		var innerExpr dst.Expr
		inner := clang.First[clang.TypeNode](n)
		if inner != nil {
			innerExpr = p.ConvertTypeExpr(inner)
			if innerExpr != nil {
				if ident, ok := innerExpr.(*dst.Ident); ok {
					if ident.Name == "void" || ident.Name == "any" {
						innerExpr = nil
					}
				}
			}
		}

		if innerExpr == nil {
			return &dst.SelectorExpr{
				X:   dst.NewIdent("unsafe"),
				Sel: dst.NewIdent("Pointer"),
			}
		}
		return &dst.StarExpr{
			X: innerExpr,
		}
	case *clang.QualType:
		var innerExpr dst.Expr
		inner := clang.First[clang.TypeNode](n)
		if inner != nil {
			innerExpr = p.ConvertTypeExpr(inner)
		}
		if innerExpr == nil {
			return p.ConvertQualType(n.Type.QualType)
		}
		return innerExpr
	case *clang.RecordType:
		return p.ConvertQualType(n.Type.QualType)
	case *clang.EnumType:
		return p.ConvertQualType(n.Type.QualType)
	case *clang.TypedefType:
		return p.ConvertQualType(n.Type.QualType)
	case *clang.ElaboratedType:
		return p.ConvertQualType(n.Type.QualType)
	default:
		// case *clang.FunctionProtoType:
		// case *clang.ParenType:
		return nil
	}
}

func (p *BaseProvider) AddType(tc TypeClass, name string, decl dst.Expr) TypeRef {
	ident := &TrackedIdentifier{Name: name}
	gen := &dst.GenDecl{
		Tok: token.TYPE,
		Specs: []dst.Spec{
			&dst.TypeSpec{
				Name:   ident.Ref(),
				Assign: tc == TcTypedef,
				Type:   decl,
			},
		},
	}

	tr := TypeRef{Identifier: ident, Decl: gen}
	switch tc {
	case TcEnum:
		p.Types["enum "+name] = tr
	case TcStruct:
		p.Types["struct "+name] = tr
	case TcUnion:
		p.Types["union "+name] = tr
	case TcTypedef:
		p.Types[name] = tr
	default:
		panic("invalid type class")
	}
	return tr
}

func (p *BaseProvider) RemapType(name string, tr TypeRef) {
	name = normalizeAnonName(name)
	p.Types[name] = tr
}

func (p *BaseProvider) FindType(name string) (TypeRef, bool) {
	for _, v := range p.Types {
		if name == v.String() {
			return v, true
		}
	}
	return TypeRef{}, false
}

func (p *BaseProvider) InferMethod(name string) (rcv string, newName string) {
	for _, rule := range p.InferredMethods {
		if strings.HasPrefix(name, rule.Name) {
			return rule.Receiver, strings.TrimPrefix(name, rule.Name)
		}
	}
	return "", name
}
