package pdbex

import (
	"fmt"
	"sort"
	"strings"
)

type ExpansionType int

const (
	ExpansionNone ExpansionType = iota
	ExpansionInlineUnnamed
	ExpansionInlineAll
)

type Settings struct {
	MemberStructExpansion   ExpansionType
	CreatePaddingMembers    bool
	ShowOffsets             bool
	MicrosoftTypedefs       bool
	AllowBitFieldsInUnion   bool
	AllowAnonymousDataTypes bool
	UseStdInt               bool
	SymbolPrefix            string
	SymbolSuffix            string
	AnonymousUnionPrefix    string
	AnonymousStructPrefix   string
	PrintHeader             bool
	PrintDeclarations       bool
	PrintDefinitions        bool
	PrintFunctions          bool
	PrintPragmaPack         bool
	Sort                    bool
}

func DefaultSettings() *Settings {
	return &Settings{
		MemberStructExpansion:   ExpansionInlineUnnamed,
		CreatePaddingMembers:    true,
		ShowOffsets:             true,
		MicrosoftTypedefs:       true,
		AllowBitFieldsInUnion:   false,
		AllowAnonymousDataTypes: true,
		UseStdInt:               false,
		SymbolPrefix:            "",
		SymbolSuffix:            "",
		AnonymousUnionPrefix:    "",
		AnonymousStructPrefix:   "",
		PrintHeader:             true,
		PrintDeclarations:       true,
		PrintDefinitions:        true,
		PrintFunctions:          false,
		PrintPragmaPack:         true,
		Sort:                    false,
	}
}

type HeaderReconstructor struct {
	settings     *Settings
	output       strings.Builder
	visited      map[uint32]bool
	unnamedCount int
}

func NewHeaderReconstructor(settings *Settings) *HeaderReconstructor {
	if settings == nil {
		settings = DefaultSettings()
	}
	return &HeaderReconstructor{
		settings: settings,
		visited:  make(map[uint32]bool),
	}
}

func (r *HeaderReconstructor) ReconstructAll(pdb *PDB) string {
	r.output.Reset()
	r.visited = make(map[uint32]bool)

	if r.settings.PrintHeader {
		r.printHeader(pdb)
	}

	if r.settings.PrintDeclarations {
		r.printDeclarations(pdb)
	}

	if r.settings.PrintDefinitions {
		r.printDefinitions(pdb)
	}

	if r.settings.PrintFunctions {
		r.printFunctions(pdb)
	}

	return r.output.String()
}

func (r *HeaderReconstructor) ReconstructSymbol(pdb *PDB, name string) (string, error) {
	sym, ok := pdb.GetSymbolByName(name)
	if !ok {
		return "", fmt.Errorf("symbol not found: %s", name)
	}

	r.output.Reset()
	r.visited = make(map[uint32]bool)

	if r.settings.PrintHeader {
		r.printHeader(pdb)
	}

	if r.settings.MemberStructExpansion == ExpansionInlineAll {
		r.reconstructSymbol(sym)
	} else if r.settings.MemberStructExpansion == ExpansionInlineUnnamed {
		r.collectReferencedTypes(sym)
		r.printDefinitionsFromCollected(pdb)
		r.reconstructSymbol(sym)
	} else {
		r.collectReferencedTypes(sym)
		r.printDefinitionsFromCollected(pdb)
		r.reconstructSymbol(sym)
	}

	return r.output.String(), nil
}

func (r *HeaderReconstructor) printHeader(pdb *PDB) {
	r.output.WriteString("/*\n")
	r.output.WriteString(fmt.Sprintf(" * PDB file: %s\n", pdb.GetPath()))
	r.output.WriteString(fmt.Sprintf(" * Image architecture: %s (0x%04X)\n", pdb.GetArchitectureString(), pdb.GetMachineType()))
	r.output.WriteString(" *\n")
	r.output.WriteString(" * Dumped by pdbex Go implementation\n")
	r.output.WriteString(" */\n\n")
}

func (r *HeaderReconstructor) printDeclarations(pdb *PDB) {
	symbols := pdb.GetSortedSymbols()

	for _, sym := range symbols {
		if sym.IsUnnamed() {
			continue
		}

		switch sym.Tag {
		case SymTagUDT:
			r.output.WriteString(fmt.Sprintf("%s %s;\n", sym.GetUdtKindString(), r.getCorrectedName(sym)))
		case SymTagEnumType:
			r.output.WriteString(fmt.Sprintf("enum %s;\n", r.getCorrectedName(sym)))
		}
	}

	r.output.WriteString("\n")
}

func (r *HeaderReconstructor) printDefinitions(pdb *PDB) {
	if r.settings.UseStdInt {
		r.output.WriteString("#include <stdint.h>\n\n")
	}

	if r.settings.PrintPragmaPack {
		r.output.WriteString("#include <pshpack1.h>\n\n")
	}

	symbols := pdb.GetSortedSymbols()
	for _, sym := range symbols {
		if sym.IsUnnamed() && r.settings.MemberStructExpansion == ExpansionInlineUnnamed {
			continue
		}
		r.reconstructSymbol(sym)
	}

	if r.settings.PrintPragmaPack {
		r.output.WriteString("\n#include <poppack.h>\n")
	}
}

func (r *HeaderReconstructor) printDefinitionsFromCollected(pdb *PDB) {
	if r.settings.UseStdInt {
		r.output.WriteString("#include <stdint.h>\n\n")
	}

	if r.settings.PrintPragmaPack {
		r.output.WriteString("#include <pshpack1.h>\n\n")
	}

	for typeId := range r.visited {
		if sym, ok := pdb.GetSymbolByTypeId(typeId); ok {
			if sym.IsUnnamed() && r.settings.MemberStructExpansion == ExpansionInlineUnnamed {
				continue
			}
			r.reconstructSymbol(sym)
		}
	}

	if r.settings.PrintPragmaPack {
		r.output.WriteString("\n#include <poppack.h>\n")
	}
}

func (r *HeaderReconstructor) printFunctions(pdb *PDB) {
	r.output.WriteString("/*\n")
	functions := pdb.GetAllFunctions()

	var names []string
	for name := range functions {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		r.output.WriteString(name)
		r.output.WriteString("\n")
	}
	r.output.WriteString("*/\n")
}

func (r *HeaderReconstructor) reconstructSymbol(sym *Symbol) {
	if sym == nil {
		return
	}

	if r.visited[sym.TypeId] {
		return
	}
	r.visited[sym.TypeId] = true

	switch sym.Tag {
	case SymTagUDT:
		r.reconstructUdt(sym)
	case SymTagEnumType:
		r.reconstructEnum(sym)
	case SymTagTypedef:
		r.reconstructTypedef(sym)
	}
}

func (r *HeaderReconstructor) reconstructUdt(sym *Symbol) {
	if sym.Udt == nil {
		return
	}

	kind := "struct"
	switch sym.Udt.Kind {
	case UdtClass:
		kind = "class"
	case UdtUnion:
		kind = "union"
	}

	name := r.getCorrectedName(sym)
	if sym.IsUnnamed() {
		name = r.getUnnamedSymbolName(sym)
	}

	r.output.WriteString(fmt.Sprintf("%s %s\n{\n", kind, name))

	for _, field := range sym.Udt.Fields {
		r.reconstructUdtField(field)
	}

	r.output.WriteString("}")

	if r.settings.MicrosoftTypedefs && !sym.IsUnnamed() {
		r.output.WriteString(fmt.Sprintf(" %s, *P%s", name, name))
	}

	r.output.WriteString(";\n\n")
}

func (r *HeaderReconstructor) reconstructUdtField(field *SymbolUdtField) {
	if field == nil {
		return
	}

	if r.settings.ShowOffsets {
		r.output.WriteString(fmt.Sprintf("  /* 0x%04X */ ", field.Offset))
	} else {
		r.output.WriteString("  ")
	}

	typeStr := r.getTypeString(field.Type)
	r.output.WriteString(typeStr)

	if field.Bits > 0 {
		r.output.WriteString(fmt.Sprintf(" : %d", field.Bits))
	}

	r.output.WriteString(fmt.Sprintf(" %s;\n", field.Name))
}

func (r *HeaderReconstructor) reconstructEnum(sym *Symbol) {
	if sym.Enum == nil {
		return
	}

	name := r.getCorrectedName(sym)
	if sym.IsUnnamed() {
		name = r.getUnnamedSymbolName(sym)
	}

	r.output.WriteString(fmt.Sprintf("enum %s\n{\n", name))

	for i, field := range sym.Enum.Fields {
		r.output.WriteString(fmt.Sprintf("  %s = %v", field.Name, field.Value))
		if i < len(sym.Enum.Fields)-1 {
			r.output.WriteString(",")
		}
		r.output.WriteString("\n")
	}

	r.output.WriteString("}")

	if r.settings.MicrosoftTypedefs && !sym.IsUnnamed() {
		r.output.WriteString(fmt.Sprintf(" %s, *P%s", name, name))
	}

	r.output.WriteString(";\n\n")
}

func (r *HeaderReconstructor) reconstructTypedef(sym *Symbol) {
	if sym.Typedef == nil {
		return
	}

	typeStr := r.getTypeString(sym.Typedef.Type)
	name := r.getCorrectedName(sym)

	r.output.WriteString(fmt.Sprintf("typedef %s %s;\n\n", typeStr, name))
}

func (r *HeaderReconstructor) getTypeString(sym *Symbol) string {
	if sym == nil {
		return "void"
	}

	switch sym.Tag {
	case SymTagBaseType:
		return GetBasicTypeString(sym.BaseType, sym.Size, r.settings.UseStdInt)
	case SymTagPointerType:
		if sym.Pointer != nil {
			baseType := r.getTypeString(sym.Pointer.Type)
			if sym.Pointer.IsReference {
				return baseType + "&"
			}
			return baseType + "*"
		}
		return "void*"
	case SymTagArrayType:
		if sym.Array != nil {
			elemType := r.getTypeString(sym.Array.ElementType)
			return fmt.Sprintf("%s[%d]", elemType, sym.Array.ElementCount)
		}
		return "void[]"
	case SymTagUDT:
		if sym.IsUnnamed() && r.settings.MemberStructExpansion == ExpansionInlineUnnamed {
			return r.getUnnamedSymbolName(sym)
		}
		return r.getCorrectedName(sym)
	case SymTagEnumType:
		return r.getCorrectedName(sym)
	case SymTagTypedef:
		return r.getCorrectedName(sym)
	case SymTagFunctionType:
		return "void*"
	default:
		if sym.Name != "" {
			return sym.Name
		}
		return "void"
	}
}

func (r *HeaderReconstructor) getCorrectedName(sym *Symbol) string {
	if sym == nil {
		return ""
	}

	name := sym.Name
	if name == "" {
		return r.getUnnamedSymbolName(sym)
	}

	if r.settings.SymbolPrefix != "" {
		name = r.settings.SymbolPrefix + name
	}
	if r.settings.SymbolSuffix != "" {
		name = name + r.settings.SymbolSuffix
	}

	return name
}

func (r *HeaderReconstructor) getUnnamedSymbolName(sym *Symbol) string {
	if sym == nil {
		return ""
	}

	r.unnamedCount++
	prefix := r.settings.AnonymousStructPrefix
	if sym.Udt != nil && sym.Udt.Kind == UdtUnion {
		prefix = r.settings.AnonymousUnionPrefix
	}

	if prefix == "" {
		prefix = "TAG_UNNAMED_"
	}

	return fmt.Sprintf("%s%d", prefix, r.unnamedCount)
}

func (r *HeaderReconstructor) collectReferencedTypes(sym *Symbol) {
	if sym == nil || r.visited[sym.TypeId] {
		return
	}

	switch sym.Tag {
	case SymTagUDT:
		if sym.Udt != nil {
			for _, field := range sym.Udt.Fields {
				if field.Type != nil && !field.Type.IsUnnamed() {
					r.collectReferencedTypes(field.Type)
					r.visited[field.Type.TypeId] = true
				}
			}
		}
	case SymTagPointerType:
		if sym.Pointer != nil && sym.Pointer.Type != nil {
			r.collectReferencedTypes(sym.Pointer.Type)
		}
	case SymTagArrayType:
		if sym.Array != nil && sym.Array.ElementType != nil {
			r.collectReferencedTypes(sym.Array.ElementType)
		}
	case SymTagTypedef:
		if sym.Typedef != nil && sym.Typedef.Type != nil {
			r.collectReferencedTypes(sym.Typedef.Type)
		}
	}
}

func (r *HeaderReconstructor) ReconstructToFiles(pdb *PDB, outputDir string) error {
	symbols := pdb.GetSortedSymbols()

	for _, sym := range symbols {
		if sym.IsUnnamed() {
			continue
		}

		r.output.Reset()
		r.visited = make(map[uint32]bool)

		if r.settings.PrintHeader {
			r.printHeader(pdb)
		}

		r.reconstructSymbol(sym)

		filename := fmt.Sprintf("%s/%s.h", outputDir, sym.Name)
		content := r.output.String()

		if err := writeFile(filename, content); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(filename, content string) error {
	return nil
}
