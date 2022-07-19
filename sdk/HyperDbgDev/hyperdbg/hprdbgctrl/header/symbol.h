#pragma once
typedef struct _LOCAL_FUNCTION_DESCRIPTION {
    std::string ObjectName;
    UINT32      ObjectSize;
} LOCAL_FUNCTION_DESCRIPTION, *PLOCAL_FUNCTION_DESCRIPTION;
#define PDBEX_DEFAULT_CONFIGURATION "-j- -k- -e n -i"
VOID
SymbolBuildAndShowSymbolTable();
BOOLEAN
SymbolShowFunctionNameBasedOnAddress(UINT64 Address, PUINT64 UsedBaseAddress);
BOOLEAN
SymbolLoadOrDownloadSymbols(BOOLEAN IsDownload, BOOLEAN SilentLoad);
BOOLEAN
SymbolConvertNameOrExprToAddress(const string & TextToConvert, PUINT64 Result);
BOOLEAN
SymbolDeleteSymTable();
BOOLEAN
SymbolBuildSymbolTable(PMODULE_SYMBOL_DETAIL * BufferToStoreDetails,
                       PUINT32                 StoredLength,
                       UINT32                  UserProcessId,
                       BOOLEAN                 SendOverSerial);
BOOLEAN
SymbolBuildAndUpdateSymbolTable(PMODULE_SYMBOL_DETAIL SymbolDetail);
VOID
SymbolInitialReload();
BOOLEAN
SymbolLocalReload(UINT32 UserProcessId);
VOID
SymbolPrepareDebuggerWithSymbolInfo(UINT32 UserProcessId);
BOOLEAN
SymbolReloadSymbolTableInDebuggerMode(UINT32 ProcessId);
