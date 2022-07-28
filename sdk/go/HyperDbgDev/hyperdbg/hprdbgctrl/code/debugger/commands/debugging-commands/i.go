package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\i.cpp"
var iBuf string

type (
	I interface {
		//Fn() (ok bool)
	}
	i struct{}
)

func Newi() I { return &i{} }
