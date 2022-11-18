package Imports

type (
	HyperDbgScriptImports interface {
		 //PSYMBOL_BUFFER ScriptEngineParse(char * str);
 //void PrintSymbolBuffer(const PSYMBOL_BUFFER SymbolBuffer);
 // PrintSymbol(PSYMBOL Symbol)VOID
 // RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer)VOID


// pdb parser
ScriptEngineSetTextMessageCallback(PVOID Handler)VOID
ScriptEngineSymbolAbortLoading()VOID
//ScriptEngineConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound)uont64
//ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName)UINT32
ScriptEngineUnloadAllSymbols()uint32
//ScriptEngineUnloadModuleSymbol(char * ModuleName)uint32
//ScriptEngineSearchSymbolForMask(const char * SearchMask)uint32
ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)bool
ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize)bool
ScriptEngineCreateSymbolTableForDisassembler(void * CallbackFunction)bool
ScriptEngineConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath)bool
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails)bool
ScriptEngineSymbolInitLoad(PVOID BufferToStoreDetails, UINT32 StoredLength, BOOLEAN DownloadIfAvailable, const char * SymbolPath, BOOLEAN IsSilentLoad)bool
ScriptEngineShowDataBasedOnSymbolTypes(const char * TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress, const char * AdditionalParameters)bool
}
)
