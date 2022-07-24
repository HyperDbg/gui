package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\unload.cpp"
var unloadBuf string

type (
	Unload interface {
		//Fn() (ok bool)
	}
	unload struct{}
)

func Newunload() Unload { return &unload{} }
