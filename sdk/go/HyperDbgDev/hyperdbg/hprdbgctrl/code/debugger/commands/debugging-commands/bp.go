package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\bp.cpp"
var bpBuf string

type (
	Bp interface {
		//Fn() (ok bool)
	}
	bp struct{}
)

func Newbp() Bp { return &bp{} }
