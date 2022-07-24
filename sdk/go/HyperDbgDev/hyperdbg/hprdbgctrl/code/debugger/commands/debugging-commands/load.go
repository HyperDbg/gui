package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\load.cpp"
var loadBuf string

type (
	Load interface {
		//Fn() (ok bool)
	}
	load struct{}
)

func Newload() Load { return &load{} }
