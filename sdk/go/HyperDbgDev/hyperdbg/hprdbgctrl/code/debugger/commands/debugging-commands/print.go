package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\print.cpp"
var printBuf string

type (
	Print interface {
		//Fn() (ok bool)
	}
	print struct{}
)

func Newprint() Print { return &print{} }
