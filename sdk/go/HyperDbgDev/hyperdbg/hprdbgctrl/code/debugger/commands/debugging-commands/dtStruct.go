package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\dt-struct.cpp"
var dtStructBuf string

type (
	DtStruct interface {
		//Fn() (ok bool)
	}
	dtStruct struct{}
)

func NewdtStruct() DtStruct { return &dtStruct{} }
