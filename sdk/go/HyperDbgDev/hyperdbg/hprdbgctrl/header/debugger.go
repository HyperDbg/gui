package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\debugger.h"
var debuggerBuf string

type (
	Debugger interface {
		//Fn() (ok bool)
	}
	debugger struct{}
)

func Newdebugger() Debugger { return &debugger{} }
