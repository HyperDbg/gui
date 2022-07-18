#pragma once
typedef struct _MODULE_SYMBOL_DETAIL {
    BOOLEAN IsSymbolDetailsFound; // TRUE if the details of symbols found, FALSE if not found
    BOOLEAN IsLocalSymbolPath;    // TRUE if the ModuleSymbolPath is a real path
    BOOLEAN IsSymbolPDBAvaliable; // TRUE if the module's pdb is avilable(if exists in the sympath)
    BOOLEAN IsUserMode;           // TRUE if the module is a user-mode module
    UINT64  BaseAddress;
    char    FilePath[MAX_PATH];
    char    ModuleSymbolPath[MAX_PATH];
    char    ModuleSymbolGuidAndAge[MAXIMUM_GUID_AND_AGE_SIZE];
} MODULE_SYMBOL_DETAIL, *PMODULE_SYMBOL_DETAIL;
typedef struct _USERMODE_LOADED_MODULE_SYMBOLS {
    UINT64  BaseAddress;
    UINT64  Entrypoint;
    wchar_t FilePath[MAX_PATH];
} USERMODE_LOADED_MODULE_SYMBOLS, *PUSERMODE_LOADED_MODULE_SYMBOLS;
typedef struct _USERMODE_LOADED_MODULE_DETAILS {
    UINT32  ProcessId;
    BOOLEAN OnlyCountModules;
    UINT32  ModulesCount;
    UINT32  Result;
} USERMODE_LOADED_MODULE_DETAILS, *PUSERMODE_LOADED_MODULE_DETAILS;
typedef VOID (*SymbolMapCallback)(UINT64 Address, char * ModuleName, char * ObjectName, unsigned int ObjectSize);
typedef struct _DEBUGGER_UPDATE_SYMBOL_TABLE {
    UINT32               TotalSymbols;
    UINT32               CurrentSymbolIndex;
    MODULE_SYMBOL_DETAIL SymbolDetailPacket;
} DEBUGGER_UPDATE_SYMBOL_TABLE, *PDEBUGGER_UPDATE_SYMBOL_TABLE;
static_assert(sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger than DEBUGGER_UPDATE_SYMBOL_TABLE (MODULE_SYMBOL_DETAIL)");
typedef struct _DEBUGGEE_SYMBOL_UPDATE_RESULT {
    UINT64 KernelStatus; // Kerenl put the status in this field
} DEBUGGEE_SYMBOL_UPDATE_RESULT, *PDEBUGGEE_SYMBOL_UPDATE_RESULT;
