package scriptengine

import (
	"fmt"
	"sync"

	"github.com/ddkwork/sdk"
)

type ScriptEngineWrapper struct {
	engine              *ScriptEngine
	symbolTable         *SymbolTable
	textMessageCallback func(message string)
	mu                  sync.RWMutex
}

func NewScriptEngineWrapper() *ScriptEngineWrapper {
	return &ScriptEngineWrapper{
		engine:      NewScriptEngine(),
		symbolTable: NewSymbolTable(),
	}
}

func (sew *ScriptEngineWrapper) ConvertNameToAddress(functionOrVariableName string) (uint64, bool) {
	return sew.symbolTable.ConvertNameToAddress(functionOrVariableName)
}

func (sew *ScriptEngineWrapper) LoadFileSymbol(baseAddress uint64, pdbFileName string, customModuleName string) uint32 {
	return sew.symbolTable.LoadFileSymbol(baseAddress, pdbFileName, customModuleName)
}

func (sew *ScriptEngineWrapper) SetTextMessageCallback(handler func(message string)) {
	sew.mu.Lock()
	defer sew.mu.Unlock()
	sew.textMessageCallback = handler
}

func (sew *ScriptEngineWrapper) UnloadAllSymbols() uint32 {
	return sew.symbolTable.UnloadAllSymbols()
}

func (sew *ScriptEngineWrapper) UnloadModuleSymbol(moduleName string) uint32 {
	return sew.symbolTable.UnloadModuleSymbol(moduleName)
}

func (sew *ScriptEngineWrapper) GetEngine() *ScriptEngine {
	return sew.engine
}

func (sew *ScriptEngineWrapper) GetSymbolTable() *SymbolTable {
	return sew.symbolTable
}

func (sew *ScriptEngineWrapper) EvalExpression(expr string) (uint64, bool) {
	return sew.engine.EvalSingleExpression(expr)
}

func (sew *ScriptEngineWrapper) ExecuteScript(scriptContent string) error {
	return nil
}

func (sew *ScriptEngineWrapper) CallFunction(functionName string, args []uint64) (uint64, error) {
	return 0, nil
}

// ScriptEngineConvertFileToPdbPathWrapper converts a local file path to PDB path
func (sew *ScriptEngineWrapper) ConvertFileToPdbPath(localFilePath string) (string, bool) {
	// Extract PDB path from PE file debug directory
	// This requires PE parsing - returning the .pdb filename convention
	return sew.symbolTable.ConvertFileToPdbPath(localFilePath)
}

// ScriptEngineSymbolInitLoadWrapper initializes and loads symbols
func (sew *ScriptEngineWrapper) SymbolInitLoad(downloadIfAvailable bool, symbolPath string, isSilentLoad bool) bool {
	return sew.symbolTable.SymbolInitLoad(downloadIfAvailable, symbolPath, isSilentLoad)
}

// ScriptEngineShowDataBasedOnSymbolTypesWrapper shows data based on symbol types
func (sew *ScriptEngineWrapper) ShowDataBasedOnSymbolTypes(typeName string, address uint64, isStruct bool, bufferAddress []byte, additionalParameters string) bool {
	return sew.symbolTable.ShowDataBasedOnSymbolTypes(typeName, address, isStruct, bufferAddress, additionalParameters)
}

// ScriptEngineSymbolAbortLoadingWrapper aborts symbol loading
func (sew *ScriptEngineWrapper) SymbolAbortLoading() {
	sew.symbolTable.AbortLoading()
}

// ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper converts file to PDB details
func (sew *ScriptEngineWrapper) ConvertFileToPdbFileAndGuidAndAgeDetails(localFilePath string, is32BitModule bool) (pdbFilePath string, guidAndAgeDetails string, ok bool) {
	return sew.symbolTable.ConvertFileToPdbFileAndGuidAndAgeDetails(localFilePath, is32BitModule)
}

// ScriptEngineParseWrapper parses a script expression
func (sew *ScriptEngineWrapper) Parse(expr string, showErrorMessageIfAny bool) *ParsedCodeBuffer {
	codeBuffer := sew.engine.ParseScript(expr, !showErrorMessageIfAny)
	if codeBuffer == nil {
		if showErrorMessageIfAny && sew.textMessageCallback != nil {
			sew.textMessageCallback(fmt.Sprintf("err, parse error in: %s\n", expr))
		}
		return nil
	}
	return codeBuffer
}

// PrintSymbolBufferWrapper prints the symbol buffer (for debugging)
func (sew *ScriptEngineWrapper) PrintSymbolBuffer(codeBuffer *ParsedCodeBuffer) {
	if codeBuffer == nil {
		return
	}
	// Print debug info about the parsed code buffer
	if sew.textMessageCallback != nil {
		sew.textMessageCallback(fmt.Sprintf("CodeBuffer: Address=0x%x Length=%d Pointer=%d\n",
			codeBuffer.Address, codeBuffer.Length, codeBuffer.Pointer))
	}
}

// ScriptAutomaticStatementsTestWrapper tests script engine statements
func (sew *ScriptEngineWrapper) AutomaticStatementsTest(expr string, expectationValue uint64, exceptError bool) bool {
	// Reset global result
	sew.engine.currentExprEvalResult = 0
	sew.engine.currentExprEvalHasError = false

	// Run the test parser
	sew.TestParser(expr)

	// Check results
	if sew.engine.currentExprEvalHasError && exceptError {
		return true
	}
	if expectationValue == sew.engine.currentExprEvalResult {
		return true
	}
	return false
}

// AllocateStructForCasting allocates a test structure for casting tests
func (sew *ScriptEngineWrapper) AllocateStructForCasting() any {
	// Allocate test structures for script engine testing
	type UnicodeString struct {
		Length        uint16
		MaximumLength uint16
		Buffer        uint64
	}

	type StupidStruct1 struct {
		Flag32      uint32
		Pad0        [4]byte // padding for alignment
		Flag64      uint64
		Context     uint64
		StringValue uint64
	}

	type StupidStruct2 struct {
		Sina32        uint32
		Pad0          [4]byte
		Sina64        uint64
		AghaaSina     uint64
		UnicodeStr    uint64
		StupidStruct1 uint64
	}

	return &StupidStruct2{}
}

// ScriptEngineWrapperTestParser tests the script parser with fake registers
func (sew *ScriptEngineWrapper) TestParser(expr string) {
	// Create fake guest registers for testing
	regs := &sdk.GUEST_REGS{
		Rax: 0x1,
		Rcx: 0x2,
		Rdx: 0x3,
		Rbx: 0x4,
		Rsp: 0x5,
		Rbp: 0x6,
		Rsi: 0x7,
		Rdi: 0x8,
		R8:  0x9,
		R9:  0xa,
		R10: 0xb,
		R11: 0xc,
		R12: 0xd,
		R13: 0xe,
		R14: 0xf,
		R15: 0x10,
	}

	sew.engine.EvalWithRegisters(regs, expr)
}

// ScriptEngineWrapperTestParserForHwdbg tests parser for hwdbg
func (sew *ScriptEngineWrapper) TestParserForHwdbg(expr string) {
	sew.engine.EvalWithRegisters(nil, expr)
}

// ScriptEngineWrapperGetHead returns the head address of the code buffer
func (sew *ScriptEngineWrapper) GetHead(codeBuffer *ParsedCodeBuffer) uint64 {
	if codeBuffer == nil {
		return 0
	}
	return codeBuffer.Address
}

// ScriptEngineWrapperGetSize returns the size of the code buffer
func (sew *ScriptEngineWrapper) GetSize(codeBuffer *ParsedCodeBuffer) uint32 {
	if codeBuffer == nil {
		return 0
	}
	return codeBuffer.Length
}

// ScriptEngineWrapperGetPointer returns the pointer of the code buffer
func (sew *ScriptEngineWrapper) GetPointer(codeBuffer *ParsedCodeBuffer) uint32 {
	if codeBuffer == nil {
		return 0
	}
	return codeBuffer.Pointer
}

// ScriptEngineWrapperRemoveSymbolBuffer removes the symbol buffer
func (sew *ScriptEngineWrapper) RemoveSymbolBuffer(codeBuffer *ParsedCodeBuffer) {
	if codeBuffer == nil {
		return
	}
	sew.engine.RemoveSymbolBuffer(codeBuffer.CodeBuffer)
}

// ScriptEngineSearchSymbolForMaskWrapper searches symbols matching a mask
func (sew *ScriptEngineWrapper) SearchSymbolForMask(searchMask string) uint32 {
	return sew.symbolTable.SearchSymbolForMask(searchMask)
}

// ScriptEngineGetFieldOffsetWrapper gets the offset of a field in a type
func (sew *ScriptEngineWrapper) GetFieldOffset(typeName, fieldName string) (uint32, bool) {
	return sew.symbolTable.GetFieldOffset(typeName, fieldName)
}

// ScriptEngineGetDataTypeSizeWrapper gets the size of a data type
func (sew *ScriptEngineWrapper) GetDataTypeSize(typeName string) (uint64, bool) {
	return sew.symbolTable.GetDataTypeSize(typeName)
}

// ScriptEngineCreateSymbolTableForDisassemblerWrapper creates symbol table for disassembler
func (sew *ScriptEngineWrapper) CreateSymbolTableForDisassembler(callback func(uint64, string, string, uint32)) bool {
	return sew.symbolTable.CreateSymbolTableForDisassembler(callback)
}

// ScriptEngineEvalUInt64StyleExpressionWrapper evaluates expression in VMI mode
func (sew *ScriptEngineWrapper) EvalUInt64StyleExpression(expr string, hasError *bool) uint64 {
	// In VMI mode, all registers are zero
	*hasError = false
	result, err := sew.engine.EvalSingleExpression(expr)
	if err == false {
		*hasError = true
		return 0
	}
	return result
}
