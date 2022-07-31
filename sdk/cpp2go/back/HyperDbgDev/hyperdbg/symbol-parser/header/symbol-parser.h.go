package header
//back\HyperDbgDev\hyperdbg\symbol-parser\header\symbol-parser.h.back

const(
DoNotShowDetailedResult = TRUE //col:18
)

type (
SymbolParser interface{
__declspec()(ok bool)//col:67
}
)

func NewSymbolParser() { return & symbolParser{} }

func (s *symbolParser)__declspec()(ok bool){//col:67
/*__declspec(dllimport) int pdbex_export(int argc, char ** argv, bool is_struct, void * buffer_address);
__declspec(dllimport) void pdbex_set_logging_method_export(PVOID handler);
__declspec(dllexport) VOID SymSetTextMessageCallback(PVOID handler);
__declspec(dllexport) UINT32 SymLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName);
__declspec(dllexport) UINT32 SymUnloadAllSymbols();
__declspec(dllexport) UINT32 SymUnloadModuleSymbol(char * ModuleName);
__declspec(dllexport) UINT32 SymSearchSymbolForMask(const char * SearchMask);
__declspec(dllexport) BOOLEAN SymGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);
__declspec(dllexport) BOOLEAN SymGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize);
__declspec(dllexport) BOOLEAN SymCreateSymbolTableForDisassembler(void * CallbackFunction);
__declspec(dllexport) UINT64 SymConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound);
__declspec(dllexport) BOOLEAN SymConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath);
__declspec(dllexport) BOOLEAN SymConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails);
__declspec(dllexport) BOOLEAN SymbolInitLoad(PVOID BufferToStoreDetails, UINT32 StoredLength, BOOLEAN DownloadIfAvailable, const char * SymbolPath, BOOLEAN IsSilentLoad);
__declspec(dllexport) BOOLEAN SymShowDataBasedOnSymbolTypes(const char * TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress, const char * AdditionalParameters);
__declspec(dllexport) VOID SymbolAbortLoading();
__declspec(dllexport) BOOLEAN SymQuerySizeof(_In_ const char * StructNameOrTypeName, _Out_ UINT32 * SizeOfField);
__declspec(dllexport) BOOLEAN SymCastingQueryForFiledsAndTypes(_In_ const char * StructName, _In_ const char * FiledOfStructName, _Out_ PBOOLEAN IsStructNamePointerOrNot, _Out_ PBOOLEAN IsFiledOfStructNamePointerOrNot, _Out_ char ** NewStructOrTypeName, _Out_ UINT32 * OffsetOfFieldFromTop, _Out_ UINT32 * SizeOfField);
}*/
return true
}



