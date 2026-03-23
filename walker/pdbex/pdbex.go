package pdbex

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type PDB struct {
	path            string
	machineType     uint16
	language        uint32
	symbols         map[uint32]*Symbol
	symbolNames     map[string]*Symbol
	functions       map[string]*FunctionInfo
	functionsByAddr map[uint64]*FunctionInfo
	dia             *diaSession
	mu              sync.RWMutex
}

func NewPDB() *PDB {
	return &PDB{
		symbols:         make(map[uint32]*Symbol),
		symbolNames:     make(map[string]*Symbol),
		functions:       make(map[string]*FunctionInfo),
		functionsByAddr: make(map[uint64]*FunctionInfo),
	}
}

func (p *PDB) Open(path string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("PDB file not found: %s", absPath)
	}

	p.path = absPath
	p.symbols = make(map[uint32]*Symbol)
	p.symbolNames = make(map[string]*Symbol)
	p.functions = make(map[string]*FunctionInfo)
	p.functionsByAddr = make(map[uint64]*FunctionInfo)

	return p.parsePDB()
}

func (p *PDB) parsePDB() error {
	return p.parseWithDIA()
}

func (p *PDB) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.dia != nil {
		p.dia.close()
		p.dia = nil
	}

	p.symbols = nil
	p.symbolNames = nil
	p.functions = nil
	p.functionsByAddr = nil
}

func (p *PDB) IsOpened() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.symbols != nil
}

func (p *PDB) GetPath() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.path
}

func (p *PDB) GetMachineType() uint16 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.machineType
}

func (p *PDB) GetLanguage() uint32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.language
}

func (p *PDB) GetSymbolByName(name string) (*Symbol, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	sym, ok := p.symbolNames[name]
	return sym, ok
}

func (p *PDB) GetSymbolByTypeId(typeId uint32) (*Symbol, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	sym, ok := p.symbols[typeId]
	return sym, ok
}

func (p *PDB) GetAllSymbols() map[uint32]*Symbol {
	p.mu.RLock()
	defer p.mu.RUnlock()

	result := make(map[uint32]*Symbol, len(p.symbols))
	for k, v := range p.symbols {
		result[k] = v
	}
	return result
}

func (p *PDB) GetAllNamedSymbols() map[string]*Symbol {
	p.mu.RLock()
	defer p.mu.RUnlock()

	result := make(map[string]*Symbol, len(p.symbolNames))
	for k, v := range p.symbolNames {
		result[k] = v
	}
	return result
}

func (p *PDB) GetAllFunctions() map[string]*FunctionInfo {
	p.mu.RLock()
	defer p.mu.RUnlock()

	result := make(map[string]*FunctionInfo, len(p.functions))
	for k, v := range p.functions {
		result[k] = v
	}
	return result
}

func (p *PDB) GetFunctionByOffset(offset uint64) (string, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if fn, ok := p.functionsByAddr[offset]; ok {
		return fn.Name, true
	}

	var candidates []struct {
		addr uint64
		name string
	}
	for addr, fn := range p.functionsByAddr {
		if addr <= offset {
			candidates = append(candidates, struct {
				addr uint64
				name string
			}{addr, fn.Name})
		}
	}

	if len(candidates) == 0 {
		return "", false
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].addr > candidates[j].addr
	})

	return candidates[0].name, true
}

func (p *PDB) GetFunctionInfo(name string) (*FunctionInfo, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	fn, ok := p.functions[name]
	return fn, ok
}

func (p *PDB) GetSortedSymbols() []*Symbol {
	p.mu.RLock()
	defer p.mu.RUnlock()

	symbols := make([]*Symbol, 0, len(p.symbols))
	for _, sym := range p.symbols {
		if sym.Name != "" && !sym.IsUnnamed() {
			symbols = append(symbols, sym)
		}
	}

	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i].Name < symbols[j].Name
	})

	return symbols
}

func (p *PDB) GetSymbolsByTag(tag SymTag) []*Symbol {
	p.mu.RLock()
	defer p.mu.RUnlock()

	var result []*Symbol
	for _, sym := range p.symbols {
		if sym.Tag == tag {
			result = append(result, sym)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}

func (p *PDB) GetUDTs() []*Symbol {
	return p.GetSymbolsByTag(SymTagUDT)
}

func (p *PDB) GetEnums() []*Symbol {
	return p.GetSymbolsByTag(SymTagEnumType)
}

func (p *PDB) GetTypedefs() []*Symbol {
	return p.GetSymbolsByTag(SymTagTypedef)
}

func (p *PDB) registerSymbol(sym *Symbol) {
	if sym == nil {
		return
	}

	p.symbols[sym.TypeId] = sym

	if sym.Name != "" {
		p.symbolNames[sym.Name] = sym
	}
}

func (p *PDB) registerFunction(fn *FunctionInfo) {
	if fn == nil || fn.Name == "" {
		return
	}

	p.functions[fn.Name] = fn
	if fn.Address > 0 {
		p.functionsByAddr[fn.Address] = fn
	}
}

func (p *PDB) GetArchitectureString() string {
	switch p.machineType {
	case 0x014c:
		return "i386"
	case 0x8664:
		return "AMD64"
	case 0x0200:
		return "IA64"
	case 0x01c0:
		return "ArmNT"
	case 0xaa64:
		return "ARM64"
	default:
		return "Unknown"
	}
}

func (p *PDB) GetMachineTypeString() string {
	return p.GetArchitectureString()
}

func (p *PDB) GetFunctionByTypeId(typeId uint32) (*Symbol, bool) {
	return p.GetSymbolByTypeId(typeId)
}

func (p *PDB) GetFunctionBySymbolName(name string) (*FunctionInfo, bool) {
	return p.GetFunctionInfo(name)
}

func (p *PDB) DumpSymbol(name string) (string, error) {
	sym, ok := p.GetSymbolByName(name)
	if !ok {
		return "", fmt.Errorf("symbol not found: %s", name)
	}

	var sb strings.Builder
	p.dumpSymbolRecursive(&sb, sym, 0)
	return sb.String(), nil
}

func (p *PDB) dumpSymbolRecursive(sb *strings.Builder, sym *Symbol, indent int) {
	prefix := strings.Repeat("  ", indent)

	sb.WriteString(prefix)
	sb.WriteString(fmt.Sprintf("Symbol: %s (Tag: %s, Size: %d)\n", sym.Name, sym.Tag, sym.Size))

	switch sym.Tag {
	case SymTagUDT:
		if sym.Udt != nil {
			sb.WriteString(prefix)
			sb.WriteString(fmt.Sprintf("  Kind: %s, Fields: %d\n", sym.Udt.Kind, sym.Udt.FieldCount))
			for _, field := range sym.Udt.Fields {
				sb.WriteString(prefix)
				sb.WriteString(fmt.Sprintf("    [+0x%04X] %s", field.Offset, field.Name))
				if field.Type != nil {
					sb.WriteString(fmt.Sprintf(": %s", field.Type.Name))
				}
				if field.Bits > 0 {
					sb.WriteString(fmt.Sprintf(" : %d", field.Bits))
				}
				sb.WriteString("\n")
			}
		}
	case SymTagEnumType:
		if sym.Enum != nil {
			sb.WriteString(prefix)
			sb.WriteString(fmt.Sprintf("  Fields: %d\n", sym.Enum.FieldCount))
			for _, field := range sym.Enum.Fields {
				sb.WriteString(prefix)
				sb.WriteString(fmt.Sprintf("    %s = %v\n", field.Name, field.Value))
			}
		}
	case SymTagPointerType:
		if sym.Pointer != nil {
			sb.WriteString(prefix)
			ref := ""
			if sym.Pointer.IsReference {
				ref = " &"
			}
			sb.WriteString(fmt.Sprintf("  Pointer%s to: %s\n", ref, sym.Pointer.Type.Name))
		}
	case SymTagArrayType:
		if sym.Array != nil {
			sb.WriteString(prefix)
			sb.WriteString(fmt.Sprintf("  Array[%d] of: %s\n", sym.Array.ElementCount, sym.Array.ElementType.Name))
		}
	}
}

func (p *PDB) FindSourceLineByRVA(rva uint32) (fileName string, lineNumber uint32, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.dia == nil {
		return "", 0, false
	}

	return p.dia.FindSourceLineByRVA(rva)
}

func (p *PDB) FindSourceLineByVA(va uint64) (fileName string, lineNumber uint32, ok bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.dia == nil {
		return "", 0, false
	}

	return p.dia.FindSourceLineByVA(va)
}

type SourceLocationInfo struct {
	FileName   string
	LineNumber uint32
	Function   string
	Address    uint64
}

func (p *PDB) FindSourceLocationByRVA(rva uint32) (*SourceLocationInfo, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.dia == nil {
		return nil, false
	}

	info := &SourceLocationInfo{
		Address: uint64(rva),
	}

	fileName, lineNumber, ok := p.dia.FindSourceLineByRVA(rva)
	if ok {
		info.FileName = fileName
		info.LineNumber = lineNumber
	}

	if fn, ok := p.functionsByAddr[uint64(rva)]; ok {
		info.Function = fn.Name
	} else {
		if name, ok := p.GetFunctionByOffset(uint64(rva)); ok {
			info.Function = name
		}
	}

	return info, true
}
