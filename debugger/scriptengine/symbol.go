package scriptengine

import (
	"fmt"
	"sync"

	"github.com/ddkwork/HyperDbg/debugger/common"
)

type SymbolDetail struct {
	Address    uint64
	ModuleName string
	ObjectName string
	ObjectSize uint32
}

type SymbolTable struct {
	symbols      []SymbolDetail
	currentIndex uint32
	mu           sync.RWMutex
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make([]SymbolDetail, 0, 1000),
	}
}

func (st *SymbolTable) ConvertNameToAddress(functionOrVariableName string) (uint64, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	for _, symbol := range st.symbols {
		if symbol.ObjectName == functionOrVariableName {
			return symbol.Address, true
		}
	}

	return 0, false
}

func (st *SymbolTable) LoadFileSymbol(baseAddress uint64, pdbFileName string, customModuleName string) uint32 {
	return 0
}

func (st *SymbolTable) UnloadAllSymbols() uint32 {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.symbols = make([]SymbolDetail, 0, 1000)
	st.currentIndex = 0

	return 0
}

func (st *SymbolTable) UnloadModuleSymbol(moduleName string) uint32 {
	st.mu.Lock()
	defer st.mu.Unlock()

	var newSymbols []SymbolDetail
	for _, symbol := range st.symbols {
		if symbol.ModuleName != moduleName {
			newSymbols = append(newSymbols, symbol)
		}
	}
	st.symbols = newSymbols

	return 0
}

func (st *SymbolTable) AddSymbol(address uint64, moduleName string, objectName string, objectSize uint32) {
	st.mu.Lock()
	defer st.mu.Unlock()

	symbol := SymbolDetail{
		Address:    address,
		ModuleName: moduleName,
		ObjectName: objectName,
		ObjectSize: objectSize,
	}

	st.symbols = append(st.symbols, symbol)
	st.currentIndex++
}

func (st *SymbolTable) GetSymbolByAddress(address uint64) (*SymbolDetail, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	for i := range st.symbols {
		if st.symbols[i].Address == address {
			return &st.symbols[i], true
		}
	}

	return nil, false
}

func (st *SymbolTable) GetSymbolCount() uint32 {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.currentIndex
}

func (st *SymbolTable) FindNearestSymbol(address uint64) (*SymbolDetail, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()

	var nearest *SymbolDetail
	var minDistance uint64 = ^uint64(0)

	for i := range st.symbols {
		if address >= st.symbols[i].Address {
			distance := address - st.symbols[i].Address
			if distance < minDistance {
				minDistance = distance
				nearest = &st.symbols[i]
			}
		}
	}

	if nearest != nil {
		return nearest, true
	}

	return nil, false
}

func (se *ScriptEngine) ShowFunctionNameBasedOnAddress(address uint64) (bool, uint64) {
	fmt.Printf("0x%x", address)
	return true, 0
}

// ConvertFileToPdbPath converts a local file path to PDB path
func (st *SymbolTable) ConvertFileToPdbPath(localFilePath string) (string, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	// This would parse the PE file's debug directory to find the PDB path
	// For now, return a placeholder
	return "", false
}

// SymbolInitLoad initializes and loads symbols
func (st *SymbolTable) SymbolInitLoad(downloadIfAvailable bool, symbolPath string, isSilentLoad bool) bool {
	st.mu.Lock()
	defer st.mu.Unlock()
	// Symbol loading implementation placeholder
	return true
}

// ShowDataBasedOnSymbolTypes shows data based on symbol types
func (st *SymbolTable) ShowDataBasedOnSymbolTypes(typeName string, address uint64, isStruct bool, bufferAddress []byte, additionalParameters string) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()
	// Display data at address according to the type definition
	return false
}

// AbortLoading aborts the current symbol loading process
func (st *SymbolTable) AbortLoading() {
	st.mu.Lock()
	defer st.mu.Unlock()
	// Signal to abort any ongoing symbol loading
}

// ConvertFileToPdbFileAndGuidAndAgeDetails extracts PDB file and GUID/age from PE
func (st *SymbolTable) ConvertFileToPdbFileAndGuidAndAgeDetails(localFilePath string, is32BitModule bool) (pdbFilePath string, guidAndAgeDetails string, ok bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	// This would parse the PE file to extract PDB info from the debug directory
	// For now, return not found
	return "", "", false
}

// SearchSymbolForMask searches symbols matching a wildcard mask
func (st *SymbolTable) SearchSymbolForMask(searchMask string) uint32 {
	st.mu.RLock()
	defer st.mu.RUnlock()

	var count uint32
	for _, symbol := range st.symbols {
		if matchWildcard(searchMask, symbol.ObjectName) {
			fmt.Printf("%s!%s @ %x\n", symbol.ModuleName, symbol.ObjectName, symbol.Address)
			count++
		}
	}
	return count
}

// GetFieldOffset gets the offset of a field within a type
func (st *SymbolTable) GetFieldOffset(typeName, fieldName string) (uint32, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	// Would need type information from PDB symbols
	return 0, false
}

// GetDataTypeSize gets the size of a data type
func (st *SymbolTable) GetDataTypeSize(typeName string) (uint64, bool) {
	st.mu.RLock()
	defer st.mu.RUnlock()
	// Would need type information from PDB symbols
	return 0, false
}

// CreateSymbolTableForDisassembler creates symbol table for disassembler use
func (st *SymbolTable) CreateSymbolTableForDisassembler(callback func(uint64, string, string, uint32)) bool {
	st.mu.RLock()
	defer st.mu.RUnlock()

	if callback == nil {
		return false
	}

	for _, symbol := range st.symbols {
		callback(symbol.Address, symbol.ModuleName, symbol.ObjectName, symbol.ObjectSize)
	}
	return true
}

// matchWildcard performs simple wildcard matching
func matchWildcard(pattern, text string) bool {
	if pattern == "" {
		return text == ""
	}
	if pattern == "*" {
		return true
	}

	pi := 0
	ti := 0
	starIdx := -1
	matchIdx := 0

	for ti < len(text) {
		if pi < len(pattern) && (pattern[pi] == text[ti] || pattern[pi] == '?') {
			pi++
			ti++
		} else if pi < len(pattern) && pattern[pi] == '*' {
			starIdx = pi
			matchIdx = ti
			pi++
		} else if starIdx != -1 {
			pi = starIdx + 1
			matchIdx++
			ti = matchIdx
		} else {
			return false
		}
	}

	for pi < len(pattern) && pattern[pi] == '*' {
		pi++
	}

	return pi == len(pattern)
}

// SymbolConvertNameOrExprToAddress converts symbol name or expression to address
func (st *SymbolTable) ConvertNameOrExprToAddress(textToConvert string, evalExpr func(string) (uint64, bool)) (uint64, bool) {
	// First try to convert as a number
	result := common.ConvertStringToUInt64(textToConvert)
	if result.Ok {
		return result.Value, true
	}

	// Check for symbol object names
	addr, found := st.ConvertNameToAddress(textToConvert)
	if found {
		return addr, true
	}

	// Try as expression
	if evalExpr != nil {
		val, hasError := evalExpr(textToConvert)
		if !hasError {
			return val, true
		}
	}

	return 0, false
}

// DeleteSymTable deletes and frees the symbol table
func (st *SymbolTable) DeleteSymTable() bool {
	st.mu.Lock()
	defer st.mu.Unlock()

	st.symbols = make([]SymbolDetail, 0, 1000)
	st.currentIndex = 0
	return true
}
