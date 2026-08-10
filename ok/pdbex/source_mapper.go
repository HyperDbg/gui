package pdbex

import (
	"fmt"
	"path/filepath"
	"strings"
)

type SourceLocation struct {
	File      string
	Line      uint32
	Column    uint32
	EndLine   uint32
	EndColumn uint32
}

type SourceMapper struct {
	pdb         *PDB
	sourceBase  string
	outputBase  string
	locationMap map[uint64]*SourceLocation
	fileMap     map[string][]uint64
}

func NewSourceMapper(pdb *PDB) *SourceMapper {
	return &SourceMapper{
		pdb:         pdb,
		locationMap: make(map[uint64]*SourceLocation),
		fileMap:     make(map[string][]uint64),
	}
}

func (s *SourceMapper) SetSourceBase(path string) {
	s.sourceBase = path
}

func (s *SourceMapper) SetOutputBase(path string) {
	s.outputBase = path
}

func (s *SourceMapper) BuildLocationMap() error {
	if s.pdb == nil {
		return fmt.Errorf("PDB is nil")
	}

	functions := s.pdb.GetAllFunctions()
	for name, fn := range functions {
		if fn.SourceFile != "" {
			s.locationMap[fn.Address] = &SourceLocation{
				File: fn.SourceFile,
				Line: fn.LineNumber,
			}

			s.fileMap[fn.SourceFile] = append(s.fileMap[fn.SourceFile], fn.Address)
		}

		_ = name
	}

	return nil
}

func (s *SourceMapper) GetSourceLocation(address uint64) (*SourceLocation, bool) {
	loc, ok := s.locationMap[address]
	return loc, ok
}

func (s *SourceMapper) GetFunctionsInFile(file string) []uint64 {
	return s.fileMap[file]
}

func (s *SourceMapper) MapSourcePath(sourcePath string) string {
	if s.sourceBase == "" || s.outputBase == "" {
		return sourcePath
	}

	rel, err := filepath.Rel(s.sourceBase, sourcePath)
	if err != nil {
		return sourcePath
	}

	return filepath.Join(s.outputBase, rel)
}

func (s *SourceMapper) GenerateSourceComment(name string, address uint64) string {
	loc, ok := s.GetSourceLocation(address)
	if !ok {
		return ""
	}

	return fmt.Sprintf("// Source: %s:%d", loc.File, loc.Line)
}

func (s *SourceMapper) GenerateFunctionComment(fn *FunctionInfo) string {
	if fn == nil {
		return ""
	}

	var parts []string

	if fn.SourceFile != "" {
		parts = append(parts, fmt.Sprintf("Source: %s", fn.SourceFile))
	}
	if fn.LineNumber > 0 {
		parts = append(parts, fmt.Sprintf("Line: %d", fn.LineNumber))
	}
	if fn.Address > 0 {
		parts = append(parts, fmt.Sprintf("RVA: 0x%X", fn.Address))
	}
	if fn.Size > 0 {
		parts = append(parts, fmt.Sprintf("Size: %d", fn.Size))
	}

	if len(parts) == 0 {
		return ""
	}

	return fmt.Sprintf("// %s", strings.Join(parts, ", "))
}

type SourceFile struct {
	Path      string
	Functions []*FunctionInfo
	Types     []*Symbol
}

func (s *SourceMapper) GetSourceFiles() []*SourceFile {
	files := make(map[string]*SourceFile)

	for path, addresses := range s.fileMap {
		sf := &SourceFile{
			Path: path,
		}

		for _, addr := range addresses {
			if name, ok := s.pdb.GetFunctionByOffset(addr); ok {
				if fn, ok := s.pdb.GetFunctionInfo(name); ok {
					sf.Functions = append(sf.Functions, fn)
				}
			}
		}

		files[path] = sf
	}

	result := make([]*SourceFile, 0, len(files))
	for _, sf := range files {
		result = append(result, sf)
	}

	return result
}

type AnnotatedSymbol struct {
	Symbol     *Symbol
	SourceFile string
	SourceLine uint32
	Comment    string
}

type AnnotatedFunction struct {
	Function   *FunctionInfo
	SourceFile string
	SourceLine uint32
	Comment    string
}

type SourceAnnotator struct {
	mapper *SourceMapper
}

func NewSourceAnnotator(mapper *SourceMapper) *SourceAnnotator {
	return &SourceAnnotator{
		mapper: mapper,
	}
}

func (a *SourceAnnotator) AnnotateFunction(fn *FunctionInfo) *AnnotatedFunction {
	if fn == nil {
		return nil
	}

	annotated := &AnnotatedFunction{
		Function:   fn,
		SourceFile: fn.SourceFile,
		SourceLine: fn.LineNumber,
		Comment:    a.mapper.GenerateFunctionComment(fn),
	}

	return annotated
}

func (a *SourceAnnotator) AnnotateFunctions(functions map[string]*FunctionInfo) []*AnnotatedFunction {
	result := make([]*AnnotatedFunction, 0, len(functions))

	for _, fn := range functions {
		if annotated := a.AnnotateFunction(fn); annotated != nil {
			result = append(result, annotated)
		}
	}

	return result
}

func (a *SourceAnnotator) GenerateAnnotatedHeader(pdb *PDB, settings *Settings) string {
	var sb strings.Builder

	if settings.PrintHeader {
		sb.WriteString(fmt.Sprintf("/*\n"))
		sb.WriteString(fmt.Sprintf(" * PDB file: %s\n", pdb.GetPath()))
		sb.WriteString(fmt.Sprintf(" * Architecture: %s\n", pdb.GetArchitectureString()))
		sb.WriteString(" *\n")
		sb.WriteString(" * Generated with source annotations\n")
		sb.WriteString(" */\n\n")
	}

	if settings.UseStdInt {
		sb.WriteString("#include <stdint.h>\n\n")
	}

	reconstructor := NewHeaderReconstructor(settings)

	udts := pdb.GetUDTs()
	for _, udt := range udts {
		if udt.IsUnnamed() && settings.MemberStructExpansion == ExpansionInlineUnnamed {
			continue
		}

		sb.WriteString(fmt.Sprintf("// UDT: %s (Size: %d)\n", udt.Name, udt.Size))
		header, err := reconstructor.ReconstructSymbol(pdb, udt.Name)
		if err == nil {
			sb.WriteString(header)
			sb.WriteString("\n")
		}
	}

	enums := pdb.GetEnums()
	for _, enum := range enums {
		if enum.IsUnnamed() {
			continue
		}

		sb.WriteString(fmt.Sprintf("// Enum: %s\n", enum.Name))
		header, err := reconstructor.ReconstructSymbol(pdb, enum.Name)
		if err == nil {
			sb.WriteString(header)
			sb.WriteString("\n")
		}
	}

	return sb.String()
}

func (a *SourceAnnotator) GenerateFunctionList(pdb *PDB) string {
	var sb strings.Builder

	functions := pdb.GetAllFunctions()

	sb.WriteString("/*\n")
	sb.WriteString(" * Function List with Source Locations\n")
	sb.WriteString(" */\n\n")

	for name, fn := range functions {
		comment := a.mapper.GenerateFunctionComment(fn)
		if comment != "" {
			sb.WriteString(comment)
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("%s()\n", name))
		sb.WriteString("\n")
	}

	return sb.String()
}

type SourceIndex struct {
	FunctionsByFile map[string][]*FunctionInfo
	TypesByFile     map[string][]*Symbol
	AllFiles        []string
}

func BuildSourceIndex(pdb *PDB) *SourceIndex {
	index := &SourceIndex{
		FunctionsByFile: make(map[string][]*FunctionInfo),
		TypesByFile:     make(map[string][]*Symbol),
	}

	fileSet := make(map[string]bool)

	functions := pdb.GetAllFunctions()
	for _, fn := range functions {
		if fn.SourceFile != "" {
			index.FunctionsByFile[fn.SourceFile] = append(index.FunctionsByFile[fn.SourceFile], fn)
			fileSet[fn.SourceFile] = true
		}
	}

	for file := range fileSet {
		index.AllFiles = append(index.AllFiles, file)
	}

	return index
}

func (idx *SourceIndex) GetFunctionsForFile(file string) []*FunctionInfo {
	return idx.FunctionsByFile[file]
}

func (idx *SourceIndex) GetFileCount() int {
	return len(idx.AllFiles)
}

func (idx *SourceIndex) GetFunctionCount() int {
	count := 0
	for _, fns := range idx.FunctionsByFile {
		count += len(fns)
	}
	return count
}
