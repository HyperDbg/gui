package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\p.cpp"
var pBuf string

type (
	P interface {
		//Fn() (ok bool)
	}
	p struct{}
)

func Newp() P { return &p{} }
