package misc

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\misc\\callstack.cpp"
var callstackBuf string

type (
	Callstack interface {
		//Fn() (ok bool)
	}
	callstack struct{}
)

func Newcallstack() Callstack { return &callstack{} }
