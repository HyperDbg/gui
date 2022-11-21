package Imports

import (
	"github.com/ddkwork/hyperdbgui/SDK/Headers"
)

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
		//ScriptEngineParse(str *int8) PSYMBOL_BUFFER
		//PrintSymbolBuffer(SymbolBuffer PSYMBOL_BUFFER) Headers.VOID
		//PrintSymbol(PSYMBOL Symbol) Headers.VOID
		//RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer) Headers.VOID
		pdbParser
	}
	script struct{}
)

func NewScript() Script { return &script{} }
