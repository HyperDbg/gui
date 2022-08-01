package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\header\symbol-parser.h.back

const (
	DoNotShowDetailedResult = TRUE //col:1
)

type SYMBOL_LOADED_MODULE_DETAILS struct {
	BaseAddress uint64           //col:3
	ModuleBase  uint64           //col:4
	ModuleName  [_MAX_FNAME]int8 //col:5
	PdbFilePath [MAX_PATH]int8   //col:6
}
