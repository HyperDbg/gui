package core

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\core\\debugger.cpp"
var debuggerBuf string

type (
	Debugger interface {
		//Fn() (ok bool)
	}
	debugger struct{}
)

func Newdebugger() Debugger { return &debugger{} }
