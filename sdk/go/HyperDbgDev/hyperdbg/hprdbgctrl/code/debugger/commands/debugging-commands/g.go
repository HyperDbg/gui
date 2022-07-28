package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\g.cpp"
var gBuf string

type (
	G interface {
		//Fn() (ok bool)
	}
	g struct{}
)

func Newg() G { return &g{} }
