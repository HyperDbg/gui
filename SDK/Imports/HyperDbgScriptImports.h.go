package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	HyperDbgScriptImports interface {
		ScriptEngineParse(str *int8) PSYMBOL_BUFFER
		PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID
		PrintSymbol(PSYMBOL Symbol) Headers.VOID
		RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer) Headers.VOID

		// pdb parser
		ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID
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

func (h *hyperDbgScriptImports) ScriptEngineParse(str *int8) interface{} {
	api.Proc(ScriptEngineParse).Call()
}

func (h *hyperDbgScriptImports) PrintSymbolBuffer(SymbolBuffer interface{}) Headers.VOID {
	api.Proc(PrintSymbolBuffer).Call()
}

func (h *hyperDbgScriptImports) PrintSymbol(PSYMBOL interface{}) Headers.VOID {
	api.Proc(PrintSymbol).Call()
}

func (h *hyperDbgScriptImports) RemoveSymbolBuffer(PSYMBOL_BUFFER interface{}) Headers.VOID {
	api.Proc(RemoveSymbolBuffer).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	api.Proc(ScriptEngineSetTextMessageCallback).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineSymbolAbortLoading() Headers.VOID {
	api.Proc(ScriptEngineSymbolAbortLoading).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64 {
	api.Proc(ScriptEngineConvertNameToAddress).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32 {
	api.Proc(ScriptEngineLoadFileSymbol).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineUnloadAllSymbols() uint32 {
	api.Proc(ScriptEngineUnloadAllSymbols).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32 {
	api.Proc(ScriptEngineUnloadModuleSymbol).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32 {
	api.Proc(ScriptEngineSearchSymbolForMask).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	api.Proc(ScriptEngineGetFieldOffset).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	api.Proc(ScriptEngineGetDataTypeSize).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	api.Proc(ScriptEngineCreateSymbolTableForDisassembler).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	api.Proc(ScriptEngineConvertFileToPdbPath).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	api.Proc(ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	api.Proc(ScriptEngineSymbolInitLoad).Call()
}

func (h *hyperDbgScriptImports) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	api.Proc(ScriptEngineShowDataBasedOnSymbolTypes).Call()
}

func newHyperDbgScriptImports() HyperDbgScriptImports {
	return &hyperDbgScriptImports{}
}
