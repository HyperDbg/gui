package header
//back\HyperDbgDev\hyperdbg\script-engine\header\script-engine.h.back

type     SCRIPT_ENGINE_ERROR_FREE uint32
const(
    SCRIPT_ENGINE_ERROR_FREE SCRIPT_ENGINE_ERROR_TYPE = 1  //col:55
    SCRIPT_ENGINE_ERROR_SYNTAX SCRIPT_ENGINE_ERROR_TYPE = 2  //col:56
    SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN SCRIPT_ENGINE_ERROR_TYPE = 3  //col:57
    SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE SCRIPT_ENGINE_ERROR_TYPE = 4  //col:58
    SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE SCRIPT_ENGINE_ERROR_TYPE = 5  //col:59
    SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL SCRIPT_ENGINE_ERROR_TYPE = 6  //col:60
)



type (
ScriptEngine interface{
__declspec()(ok bool)//col:61
}
)

func NewScriptEngine() { return & scriptEngine{} }

func (s *scriptEngine)__declspec()(ok bool){//col:61
/*__declspec(dllimport) BOOLEAN
    ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);
__declspec(dllimport) BOOLEAN
    ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize);
__declspec(dllexport) UINT64
    ScriptEngineConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound);
__declspec(dllexport) UINT32
    ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName);
__declspec(dllexport) UINT32
    ScriptEngineUnloadAllSymbols();
__declspec(dllexport) UINT32
    ScriptEngineUnloadModuleSymbol(char * ModuleName);
__declspec(dllexport) UINT32
    ScriptEngineSearchSymbolForMask(const char * SearchMask);
__declspec(dllexport) BOOLEAN
    ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);
__declspec(dllexport) BOOLEAN
    ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize);
__declspec(dllexport) BOOLEAN
    ScriptEngineCreateSymbolTableForDisassembler(void * CallbackFunction);
__declspec(dllexport) BOOLEAN
    ScriptEngineConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath);
__declspec(dllexport) BOOLEAN
    ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails);
__declspec(dllexport) BOOLEAN
    ScriptEngineSymbolInitLoad(PVOID BufferToStoreDetails, UINT32 StoredLength, BOOLEAN DownloadIfAvailable, const char * SymbolPath, BOOLEAN IsSilentLoad);
__declspec(dllexport) BOOLEAN
    ScriptEngineShowDataBasedOnSymbolTypes(const char * TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress, const char * AdditionalParameters);
__declspec(dllexport) VOID
    ScriptEngineSymbolAbortLoading();
__declspec(dllexport) VOID
    ScriptEngineSetTextMessageCallback(PVOID Handler);
typedef enum _SCRIPT_ENGINE_ERROR_TYPE
{
    SCRIPT_ENGINE_ERROR_FREE,
    SCRIPT_ENGINE_ERROR_SYNTAX,
    SCRIPT_ENGINE_ERROR_UNKOWN_TOKEN,
    SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE,
    SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE,
    SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL
} SCRIPT_ENGINE_ERROR_TYPE,*/
return true
}



