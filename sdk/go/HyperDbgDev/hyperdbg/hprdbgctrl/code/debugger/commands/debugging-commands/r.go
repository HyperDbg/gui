package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\r.cpp"
var rBuf string

type (
	R interface {
		//Fn() (ok bool)
	}
	r struct{}
)

func Newr() R { return &r{} }
