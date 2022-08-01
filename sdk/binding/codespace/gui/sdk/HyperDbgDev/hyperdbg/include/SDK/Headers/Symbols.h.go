package Headers

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Symbols.h.back

type MODULE_SYMBOL_DETAIL struct {
	IsSymbolDetailsFound   bool                            //col:3
	IsLocalSymbolPath      bool                            //col:4
	IsSymbolPDBAvaliable   bool                            //col:5
	IsUserMode             bool                            //col:6
	BaseAddress            uint64                          //col:7
	FilePath               [MAX_PATH]int8                  //col:8
	ModuleSymbolPath       [MAX_PATH]int8                  //col:9
	ModuleSymbolGuidAndAge [MAXIMUM_GUID_AND_AGE_SIZE]int8 //col:10
}

type USERMODE_LOADED_MODULE_SYMBOLS struct {
	BaseAddress uint64            //col:14
	Entrypoint  uint64            //col:15
	FilePath    [MAX_PATH]*uint32 //col:16
}

type USERMODE_LOADED_MODULE_DETAILS struct {
	ProcessId        uint32 //col:20
	OnlyCountModules bool   //col:21
	ModulesCount     uint32 //col:22
	Result           uint32 //col:23
}

type DEBUGGER_UPDATE_SYMBOL_TABLE struct {
	TotalSymbols       uint32               //col:27
	CurrentSymbolIndex uint32               //col:28
	SymbolDetailPacket MODULE_SYMBOL_DETAIL //col:29
}

type DEBUGGEE_SYMBOL_UPDATE_RESULT struct {
	KernelStatus uint64 //col:33
}

