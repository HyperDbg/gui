package pdbex

import (
	"fmt"
	"sort"

	"github.com/saferwall/pe"
)

type PEFile struct {
	path      string
	peFile    *pe.File
	imageBase uint64
	sections  []Section
	exports   []Export
	pdb       *PDB
}

type Section struct {
	Name            string
	VirtualAddress  uint32
	VirtualSize     uint32
	RawOffset       uint32
	RawSize         uint32
	Characteristics uint32
}

type Export struct {
	Name    string
	Address uint32
	Ordinal uint16
}

func NewPEFile(path string) (*PEFile, error) {
	peFile, err := pe.New(path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open PE file: %w", err)
	}

	if err := peFile.Parse(); err != nil {
		peFile.Close()
		return nil, fmt.Errorf("failed to parse PE file: %w", err)
	}

	pf := &PEFile{
		path:   path,
		peFile: peFile,
	}

	pf.parseFromSaferwall()

	return pf, nil
}

func (pf *PEFile) parseFromSaferwall() {
	if pf.peFile.Is64 {
		if oh64, ok := pf.peFile.NtHeader.OptionalHeader.(pe.ImageOptionalHeader64); ok {
			pf.imageBase = oh64.ImageBase
		}
	} else {
		if oh32, ok := pf.peFile.NtHeader.OptionalHeader.(pe.ImageOptionalHeader32); ok {
			pf.imageBase = uint64(oh32.ImageBase)
		}
	}

	pf.sections = make([]Section, 0, len(pf.peFile.Sections))
	for _, sec := range pf.peFile.Sections {
		pf.sections = append(pf.sections, Section{
			Name:            sec.String(),
			VirtualAddress:  sec.Header.VirtualAddress,
			VirtualSize:     sec.Header.VirtualSize,
			RawOffset:       sec.Header.PointerToRawData,
			RawSize:         sec.Header.SizeOfRawData,
			Characteristics: sec.Header.Characteristics,
		})
	}

	pf.exports = make([]Export, 0, len(pf.peFile.Export.Functions))
	for _, exp := range pf.peFile.Export.Functions {
		pf.exports = append(pf.exports, Export{
			Name:    exp.Name,
			Address: exp.FunctionRVA,
			Ordinal: uint16(exp.Ordinal),
		})
	}
}

func (pe *PEFile) Close() error {
	if pe.peFile != nil {
		return pe.peFile.Close()
	}
	return nil
}

func (pe *PEFile) GetImageBase() uint64 {
	return pe.imageBase
}

func (pe *PEFile) GetSections() []Section {
	return pe.sections
}

func (pe *PEFile) GetExports() []Export {
	return pe.exports
}

func (pe *PEFile) RvaToVa(rva uint32) uint64 {
	return pe.imageBase + uint64(rva)
}

func (pe *PEFile) VaToRva(va uint64) uint32 {
	if va < pe.imageBase {
		return 0
	}
	return uint32(va - pe.imageBase)
}

func (pe *PEFile) RvaToOffset(rva uint32) uint32 {
	return pe.peFile.GetOffsetFromRva(rva)
}

func (pe *PEFile) GetExportByName(name string) *Export {
	for i := range pe.exports {
		if pe.exports[i].Name == name {
			return &pe.exports[i]
		}
	}
	return nil
}

func (pe *PEFile) GetExportByOrdinal(ordinal uint16) *Export {
	for i := range pe.exports {
		if pe.exports[i].Ordinal == ordinal {
			return &pe.exports[i]
		}
	}
	return nil
}

type PEIntegration struct {
	pdb *PDB
	pe  *PEFile
}

func NewPEIntegration(pdbPath, pePath string) (*PEIntegration, error) {
	pdb := NewPDB()
	if err := pdb.Open(pdbPath); err != nil {
		return nil, fmt.Errorf("failed to open PDB: %w", err)
	}

	pe, err := NewPEFile(pePath)
	if err != nil {
		pdb.Close()
		return nil, fmt.Errorf("failed to open PE: %w", err)
	}

	return &PEIntegration{
		pdb: pdb,
		pe:  pe,
	}, nil
}

func (p *PEIntegration) Close() {
	if p.pdb != nil {
		p.pdb.Close()
	}
	if p.pe != nil {
		p.pe.Close()
	}
}

func (p *PEIntegration) GetPDB() *PDB {
	return p.pdb
}

func (p *PEIntegration) GetPE() *PEFile {
	return p.pe
}

func (p *PEIntegration) GetFunctionNameByRVA(rva uint32) (string, bool) {
	return p.pdb.GetFunctionByOffset(uint64(rva))
}

func (p *PEIntegration) GetFunctionNameByVA(va uint64) (string, bool) {
	rva := p.pe.VaToRva(va)
	if rva == 0 {
		return "", false
	}
	return p.GetFunctionNameByRVA(rva)
}

func (p *PEIntegration) GetFunctionNameByFileOffset(offset uint64) (string, bool) {
	for _, sec := range p.pe.sections {
		if offset >= uint64(sec.RawOffset) && offset < uint64(sec.RawOffset+sec.RawSize) {
			rva := uint32(offset-uint64(sec.RawOffset)) + sec.VirtualAddress
			return p.GetFunctionNameByRVA(rva)
		}
	}
	return "", false
}

func (p *PEIntegration) GetAllFunctionsWithAddresses() []FunctionWithAddress {
	functions := p.pdb.GetAllFunctions()
	result := make([]FunctionWithAddress, 0, len(functions))

	for name, fn := range functions {
		result = append(result, FunctionWithAddress{
			Name:   name,
			RVA:    fn.Address,
			VA:     p.pe.RvaToVa(uint32(fn.Address)),
			Size:   fn.Size,
			Source: fn.SourceFile,
			Line:   fn.LineNumber,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].RVA < result[j].RVA
	})

	return result
}

type FunctionWithAddress struct {
	Name   string
	RVA    uint64
	VA     uint64
	Size   uint32
	Source string
	Line   uint32
}

func (f *FunctionWithAddress) String() string {
	return fmt.Sprintf("0x%08X: %s (size: %d)", f.RVA, f.Name, f.Size)
}

func (p *PEIntegration) FindFunctionContainingRVA(rva uint32) *FunctionWithAddress {
	functions := p.GetAllFunctionsWithAddresses()

	idx := sort.Search(len(functions), func(i int) bool {
		return functions[i].RVA > uint64(rva)
	})

	if idx == 0 {
		return nil
	}

	fn := &functions[idx-1]
	if fn.RVA+uint64(fn.Size) > uint64(rva) {
		return fn
	}

	return nil
}

func (p *PEIntegration) GetSourceLocation(rva uint32) (string, uint32, bool) {
	fn := p.FindFunctionContainingRVA(rva)
	if fn != nil && fn.Source != "" {
		return fn.Source, fn.Line, true
	}
	return "", 0, false
}
