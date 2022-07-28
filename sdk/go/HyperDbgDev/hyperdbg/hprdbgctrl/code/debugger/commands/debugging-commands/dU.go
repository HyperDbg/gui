package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\d-u.cpp"
var dUBuf string

type (
	DU interface {
		//Fn() (ok bool)
	}
	dU struct{}
)

func NewdU() DU { return &dU{} }
