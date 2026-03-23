package pdbex

type SymTag int

const (
	SymTagNull SymTag = iota
	SymTagExe
	SymTagCompiland
	SymTagCompilandDetails
	SymTagCompilandEnv
	SymTagFunction
	SymTagBlock
	SymTagData
	SymTagAnnotation
	SymTagLabel
	SymTagPublicSymbol
	SymTagUDT
	SymTagEnumType
	SymTagFunctionType
	SymTagPointerType
	SymTagArrayType
	SymTagBaseType
	SymTagTypedef
	SymTagBaseClass
	SymTagFriend
	SymTagFunctionArgType
	SymTagFuncDebugStart
	SymTagFuncDebugEnd
	SymTagUsingNamespace
	SymTagVTableShape
	SymTagVTable
	SymTagCustom
	SymTagThunk
	SymTagCustomType
	SymTagManagedType
	SymTagDimension
)

func (s SymTag) String() string {
	names := map[SymTag]string{
		SymTagNull:             "Null",
		SymTagExe:              "Exe",
		SymTagCompiland:        "Compiland",
		SymTagCompilandDetails: "CompilandDetails",
		SymTagCompilandEnv:     "CompilandEnv",
		SymTagFunction:         "Function",
		SymTagBlock:            "Block",
		SymTagData:             "Data",
		SymTagAnnotation:       "Annotation",
		SymTagLabel:            "Label",
		SymTagPublicSymbol:     "PublicSymbol",
		SymTagUDT:              "UDT",
		SymTagEnumType:         "Enum",
		SymTagFunctionType:     "FunctionType",
		SymTagPointerType:      "PointerType",
		SymTagArrayType:        "ArrayType",
		SymTagBaseType:         "BaseType",
		SymTagTypedef:          "Typedef",
		SymTagBaseClass:        "BaseClass",
		SymTagFriend:           "Friend",
		SymTagFunctionArgType:  "FunctionArgType",
		SymTagFuncDebugStart:   "FuncDebugStart",
		SymTagFuncDebugEnd:     "FuncDebugEnd",
		SymTagUsingNamespace:   "UsingNamespace",
		SymTagVTableShape:      "VTableShape",
		SymTagVTable:           "VTable",
		SymTagCustom:           "Custom",
		SymTagThunk:            "Thunk",
		SymTagCustomType:       "CustomType",
		SymTagManagedType:      "ManagedType",
		SymTagDimension:        "Dimension",
	}
	if name, ok := names[s]; ok {
		return name
	}
	return "Unknown"
}

type UdtKind int

const (
	UdtStruct UdtKind = iota
	UdtClass
	UdtUnion
)

func (u UdtKind) String() string {
	switch u {
	case UdtStruct:
		return "struct"
	case UdtClass:
		return "class"
	case UdtUnion:
		return "union"
	default:
		return "unknown"
	}
}

type BasicType int

const (
	btNoType BasicType = iota
	btVoid
	btChar
	btWChar
	btInt
	btUInt
	btFloat
	btBCD
	btBool
	btLong
	btULong
	btCurrency
	btDate
	btVariant
	btComplex
	btBitfield
	btBSTR
	btHresult
	btChar16
	btChar32
)

func (b BasicType) String() string {
	names := map[BasicType]string{
		btNoType:   "btNoType",
		btVoid:     "void",
		btChar:     "char",
		btWChar:    "wchar_t",
		btInt:      "int",
		btUInt:     "unsigned int",
		btFloat:    "float",
		btBCD:      "BCD",
		btBool:     "bool",
		btLong:     "long",
		btULong:    "unsigned long",
		btCurrency: "CURRENCY",
		btDate:     "DATE",
		btVariant:  "VARIANT",
		btComplex:  "complex",
		btBitfield: "bitfield",
		btBSTR:     "BSTR",
		btHresult:  "HRESULT",
		btChar16:   "char16_t",
		btChar32:   "char32_t",
	}
	if name, ok := names[b]; ok {
		return name
	}
	return "unknown"
}

type DataKind int

const (
	DataKindUnknown DataKind = iota
	DataKindLocal
	DataKindStaticLocal
	DataKindParam
	DataKindObjectPtr
	DataKindFileStatic
	DataKindGlobal
	DataKindMember
	DataKindStaticMember
	DataKindConstant
)

func (d DataKind) String() string {
	names := map[DataKind]string{
		DataKindUnknown:      "Unknown",
		DataKindLocal:        "Local",
		DataKindStaticLocal:  "StaticLocal",
		DataKindParam:        "Param",
		DataKindObjectPtr:    "ObjectPtr",
		DataKindFileStatic:   "FileStatic",
		DataKindGlobal:       "Global",
		DataKindMember:       "Member",
		DataKindStaticMember: "StaticMember",
		DataKindConstant:     "Constant",
	}
	if name, ok := names[d]; ok {
		return name
	}
	return "Unknown"
}

type CallingConvention int

const (
	CallConvNearC CallingConvention = iota
	CallConvFarC
	CallConvNearPascal
	CallConvFarPascal
	CallConvNearFast
	CallConvFarFast
	CallConvSkipped
	CallConvNearStd
	CallConvFarStd
	CallConvNearSys
	CallConvFarSys
	CallConvThisCall
	CallConvMipsCall
	CallConvGeneric
	CallConvAlphaCall
	CallConvPpcCall
	CallConvSHCall
	CallConvArmCall
	CallConvAM33Call
	CallConvTriCall
	CallConvSH5Call
	CallConvM32RCall
	CallConvClrCall
	CallConvInline
	CallConvNearVector
)

func (c CallingConvention) String() string {
	names := map[CallingConvention]string{
		CallConvNearC:      "__cdecl",
		CallConvFarC:       "__cdecl (far)",
		CallConvNearPascal: "__pascal",
		CallConvFarPascal:  "__pascal (far)",
		CallConvNearFast:   "__fastcall",
		CallConvFarFast:    "__fastcall (far)",
		CallConvNearStd:    "__stdcall",
		CallConvFarStd:     "__stdcall (far)",
		CallConvNearSys:    "__syscall",
		CallConvFarSys:     "__syscall (far)",
		CallConvThisCall:   "__thiscall",
		CallConvMipsCall:   "__mipscall",
		CallConvGeneric:    "__generic",
		CallConvAlphaCall:  "__alphacall",
		CallConvPpcCall:    "__ppccall",
		CallConvSHCall:     "__shcall",
		CallConvArmCall:    "__armcall",
		CallConvAM33Call:   "__am33call",
		CallConvTriCall:    "__tricall",
		CallConvSH5Call:    "__sh5call",
		CallConvM32RCall:   "__m32rcall",
		CallConvClrCall:    "__clrcall",
		CallConvInline:     "__inline",
		CallConvNearVector: "__vectorcall",
	}
	if name, ok := names[c]; ok {
		return name
	}
	return "unknown"
}

type SymbolEnumField struct {
	Name   string
	Value  interface{}
	Parent *Symbol
}

type SymbolEnum struct {
	FieldCount uint32
	Fields     []*SymbolEnumField
}

type SymbolUdtField struct {
	Name        string
	Type        *Symbol
	Offset      uint32
	Bits        uint32
	BitPosition uint32
	Parent      *Symbol
}

type SymbolUdt struct {
	Kind       UdtKind
	FieldCount uint32
	Fields     []*SymbolUdtField
}

type SymbolPointer struct {
	Type        *Symbol
	IsReference bool
}

type SymbolArray struct {
	ElementType  *Symbol
	ElementCount uint32
}

type SymbolFunction struct {
	ReturnType        *Symbol
	CallingConvention CallingConvention
	ArgumentCount     uint32
	Arguments         []*Symbol
}

type SymbolFunctionArg struct {
	Type *Symbol
}

type SymbolTypedef struct {
	Type *Symbol
}

type Symbol struct {
	Tag        SymTag
	DataKind   DataKind
	BaseType   BasicType
	TypeId     uint32
	Size       uint32
	IsConst    bool
	IsVolatile bool
	Name       string

	Enum        *SymbolEnum
	Typedef     *SymbolTypedef
	Pointer     *SymbolPointer
	Array       *SymbolArray
	Function    *SymbolFunction
	FunctionArg *SymbolFunctionArg
	Udt         *SymbolUdt
}

func (s *Symbol) IsUnnamed() bool {
	if s.Name == "" {
		return true
	}
	if len(s.Name) >= 11 && s.Name[:11] == "<anonymous-" {
		return true
	}
	if len(s.Name) >= 9 && s.Name[:9] == "<unnamed-" {
		return true
	}
	if len(s.Name) >= 9 && s.Name[:9] == "__unnamed" {
		return true
	}
	return false
}

func (s *Symbol) GetUdtKindString() string {
	if s.Udt != nil {
		return s.Udt.Kind.String()
	}
	return "unknown"
}

func GetBasicTypeString(baseType BasicType, size uint32, useStdInt bool) string {
	if useStdInt {
		switch baseType {
		case btVoid:
			return "void"
		case btChar:
			return "char"
		case btWChar:
			return "wchar_t"
		case btInt:
			switch size {
			case 1:
				return "int8_t"
			case 2:
				return "int16_t"
			case 4:
				return "int32_t"
			case 8:
				return "int64_t"
			}
		case btUInt:
			switch size {
			case 1:
				return "uint8_t"
			case 2:
				return "uint16_t"
			case 4:
				return "uint32_t"
			case 8:
				return "uint64_t"
			}
		case btFloat:
			switch size {
			case 4:
				return "float"
			case 8:
				return "double"
			}
		case btBool:
			return "bool"
		}
	} else {
		switch baseType {
		case btVoid:
			return "void"
		case btChar:
			return "char"
		case btWChar:
			return "wchar_t"
		case btInt:
			switch size {
			case 1:
				return "char"
			case 2:
				return "short"
			case 4:
				return "int"
			case 8:
				return "__int64"
			}
		case btUInt:
			switch size {
			case 1:
				return "unsigned char"
			case 2:
				return "unsigned short"
			case 4:
				return "unsigned int"
			case 8:
				return "unsigned __int64"
			}
		case btFloat:
			switch size {
			case 4:
				return "float"
			case 8:
				return "double"
			}
		case btBool:
			return "bool"
		}
	}
	return "unknown"
}

type FunctionInfo struct {
	Name       string
	Address    uint64
	Size       uint32
	SourceFile string
	LineNumber uint32
}
