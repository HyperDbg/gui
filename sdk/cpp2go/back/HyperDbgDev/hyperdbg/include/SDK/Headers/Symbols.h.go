package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Symbols.h.back

type MODULE_SYMBOL_DETAIL struct{
IsSymbolDetailsFound bool
IsLocalSymbolPath bool
IsSymbolPDBAvaliable bool
IsUserMode bool
BaseAddress uint64
FilePath[MAX_PATH] char
ModuleSymbolPath[MAX_PATH] char
ModuleSymbolGuidAndAge[MAXIMUM_GUID_AND_AGE_SIZE] char
}


type USERMODE_LOADED_MODULE_SYMBOLS struct{
BaseAddress uint64
Entrypoint uint64
FilePath[MAX_PATH] wchar_t
}


type USERMODE_LOADED_MODULE_DETAILS struct{
ProcessId uint32
OnlyCountModules bool
ModulesCount uint32
Result uint32
}


type DEBUGGER_UPDATE_SYMBOL_TABLE struct{
TotalSymbols uint32
CurrentSymbolIndex uint32
SymbolDetailPacket MODULE_SYMBOL_DETAIL
}


type DEBUGGEE_SYMBOL_UPDATE_RESULT struct{
KernelStatus uint64
}




