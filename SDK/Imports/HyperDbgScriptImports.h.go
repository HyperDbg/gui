package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	HyperDbgScriptImports interface {
		ScriptEngineParse(str *int8) PSYMBOL_BUFFER
		PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID
		PrintSymbol(PSYMBOL Symbol) Headers.VOID
		RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer) Headers.VOID

		// pdb parser
		ScriptEngineSetTextMessageCallback(PVOID Handler) Headers.VOID
		ScriptEngineSymbolAbortLoading() Headers.VOID
		ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64
		ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32
		ScriptEngineUnloadAllSymbols() uint32
		ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32
		ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32
		ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool
		ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool
		ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool
		ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool
		ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool
		ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool
		ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool
	}
	hyperDbgScriptImports struct {
	}
)

// todo goto windows write code
func (h *hyperDbgScriptImports) ScriptEngineParse(str *int8) interface{} {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) PrintSymbolBuffer(SymbolBuffer interface{}) Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) PrintSymbol(PSYMBOL interface{}) Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) RemoveSymbolBuffer(PSYMBOL_BUFFER interface{}) Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineSetTextMessageCallback(PVOID interface{}) Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineSymbolAbortLoading() Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineUnloadAllSymbols() uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32 {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgScriptImports) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	//TODO implement me
	panic("implement me")
}

func newHyperDbgScriptImports() HyperDbgScriptImports {
	return &hyperDbgScriptImports{}
}
