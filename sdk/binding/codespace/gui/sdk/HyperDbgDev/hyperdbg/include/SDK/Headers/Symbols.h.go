package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Symbols.h.back

type _MODULE_SYMBOL_DETAIL struct {
	IsSymbolDetailsFound   bool                            //col:13
	IsLocalSymbolPath      bool                            //col:14
	IsSymbolPDBAvaliable   bool                            //col:15
	IsUserMode             bool                            //col:16
	BaseAddress            uint64                          //col:17
	FilePath               [260]int8                       //col:18
	ModuleSymbolPath       [260]int8                       //col:19
	ModuleSymbolGuidAndAge [MAXIMUM_GUID_AND_AGE_SIZE]int8 //col:20
}

type _USERMODE_LOADED_MODULE_SYMBOLS struct {
	BaseAddress uint64       //col:19
	Entrypoint  uint64       //col:20
	FilePath    [260]*uint32 //col:21
}

type _USERMODE_LOADED_MODULE_DETAILS struct {
	ProcessId        uint32 //col:26
	OnlyCountModules bool   //col:27
	ModulesCount     uint32 //col:28
	Result           uint32 //col:29
}

type _DEBUGGER_UPDATE_SYMBOL_TABLE struct {
	TotalSymbols       uint32               //col:32
	CurrentSymbolIndex uint32               //col:33
	SymbolDetailPacket MODULE_SYMBOL_DETAIL //col:34
}

type _DEBUGGEE_SYMBOL_UPDATE_RESULT struct {
	KernelStatus uint64 //col:36
}

