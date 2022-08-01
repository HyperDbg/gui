package header
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\header\symbol-parser.h.back

const(
DoNotShowDetailedResult = TRUE //col:1
)

type SYMBOL_LOADED_MODULE_DETAILS struct{
BaseAddress uint64
ModuleBase uint64
ModuleName[_MAX_FNAME] char
PdbFilePath[MAX_PATH] char
}




