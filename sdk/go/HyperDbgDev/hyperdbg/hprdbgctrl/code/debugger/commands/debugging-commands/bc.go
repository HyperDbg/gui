package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\bc.cpp"
var bcBuf string

type (
	Bc interface {
		//Fn() (ok bool)
	}
	bc struct{}
)

func Newbc() Bc { return &bc{} }
