package debugger

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\debugger.cpp"
var debuggerBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)
func New() Interface { return &object{} }


