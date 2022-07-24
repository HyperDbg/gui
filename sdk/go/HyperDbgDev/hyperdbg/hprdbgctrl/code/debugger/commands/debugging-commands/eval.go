package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\eval.cpp"
var evalBuf string

type (
	Eval interface {
		//Fn() (ok bool)
	}
	eval struct{}
)

func Neweval() Eval { return &eval{} }
