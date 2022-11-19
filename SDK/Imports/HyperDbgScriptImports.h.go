package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	pdbParser interface {
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
	Script interface {
		ScriptEngineParse(str *int8) PSYMBOL_BUFFER
		PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID
		PrintSymbol(PSYMBOL Symbol) Headers.VOID
		RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer) Headers.VOID
		pdbParser
	}
	script struct{}
)

func NewScript() Script { return &script{} }

func (s *script) ScriptEngineParse(str *int8) interface{} {
	api.Proc(ScriptEngineParse).Call()
}

func (s *script) PrintSymbolBuffer(SymbolBuffer interface{}) Headers.VOID {
	api.Proc(PrintSymbolBuffer).Call()
}

func (s *script) PrintSymbol(PSYMBOL interface{}) Headers.VOID {
	api.Proc(PrintSymbol).Call()
}

func (s *script) RemoveSymbolBuffer(PSYMBOL_BUFFER interface{}) Headers.VOID {
	api.Proc(RemoveSymbolBuffer).Call()
}

func (s *script) ScriptEngineSetTextMessageCallback(Handler Headers.PVOID) Headers.VOID {
	api.Proc(ScriptEngineSetTextMessageCallback).Call()
}

func (s *script) ScriptEngineSymbolAbortLoading() Headers.VOID {
	api.Proc(ScriptEngineSymbolAbortLoading).Call()
}

func (s *script) ScriptEngineConvertNameToAddress(FunctionOrVariableName *int8, WasFound bool) uint64 {
	api.Proc(ScriptEngineConvertNameToAddress).Call()
}

func (s *script) ScriptEngineLoadFileSymbol(BaseAddress uint64, PdbFileName *int8) uint32 {
	api.Proc(ScriptEngineLoadFileSymbol).Call()
}

func (s *script) ScriptEngineUnloadAllSymbols() uint32 {
	api.Proc(ScriptEngineUnloadAllSymbols).Call()
}

func (s *script) ScriptEngineUnloadModuleSymbol(ModuleName *int8) uint32 {
	api.Proc(ScriptEngineUnloadModuleSymbol).Call()
}

func (s *script) ScriptEngineSearchSymbolForMask(SearchMask *int8) uint32 {
	api.Proc(ScriptEngineSearchSymbolForMask).Call()
}

func (s *script) ScriptEngineGetFieldOffset(TypeName, FieldName *int8, FieldOffset *uint32) bool {
	api.Proc(ScriptEngineGetFieldOffset).Call()
}

func (s *script) ScriptEngineGetDataTypeSize(TypeName *int8, TypeSize *uint64) bool {
	api.Proc(ScriptEngineGetDataTypeSize).Call()
}

func (s *script) ScriptEngineCreateSymbolTableForDisassembler(CallbackFunction Headers.PVOID) bool {
	api.Proc(ScriptEngineCreateSymbolTableForDisassembler).Call()
}

func (s *script) ScriptEngineConvertFileToPdbPath(LocalFilePath, ResultPath *int8) bool {
	api.Proc(ScriptEngineConvertFileToPdbPath).Call()
}

func (s *script) ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(LocalFilePath, PdbFilePath, GuidAndAgeDetails *int8) bool {
	api.Proc(ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails).Call()
}

func (s *script) ScriptEngineSymbolInitLoad(BufferToStoreDetails Headers.PVOID, StoredLength uint32, DownloadIfAvailable bool, SymbolPath *int8, IsSilentLoad bool) bool {
	api.Proc(ScriptEngineSymbolInitLoad).Call()
}

func (s *script) ScriptEngineShowDataBasedOnSymbolTypes(TypeName *int8, Address uint64, IsStruct bool, BufferAddress Headers.PVOID, AdditionalParameters *int8) bool {
	api.Proc(ScriptEngineShowDataBasedOnSymbolTypes).Call()
}
