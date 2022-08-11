package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\header\symbol-parser.h.back

type _SYMBOL_LOADED_MODULE_DETAILS struct {
	BaseAddress uint64    //col:9
	ModuleBase  uint64    //col:10
	ModuleName  [256]int8 //col:11
	PdbFilePath [260]int8 //col:12
}

