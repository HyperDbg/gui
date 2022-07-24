package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\x.cpp"
var xBuf string

type (
	X interface {
		//Fn() (ok bool)
	}
	x struct{}
)

func Newx() X { return &x{} }
