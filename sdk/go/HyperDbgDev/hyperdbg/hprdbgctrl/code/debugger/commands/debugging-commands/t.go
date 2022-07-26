package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\t.cpp"
var tBuf string

type (
	T interface {
		//Fn() (ok bool)
	}
	t struct{}
)

func Newt() T { return &t{} }
