package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\e.cpp"
var eBuf string

type (
	E interface {
		//Fn() (ok bool)
	}
	e struct{}
)

func Newe() E { return &e{} }
