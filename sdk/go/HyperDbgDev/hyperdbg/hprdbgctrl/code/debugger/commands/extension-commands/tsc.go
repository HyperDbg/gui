package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\tsc.cpp"
var tscBuf string

type (
	Tsc interface {
		//Fn() (ok bool)
	}
	tsc struct{}
)

func Newtsc() Tsc { return &tsc{} }
