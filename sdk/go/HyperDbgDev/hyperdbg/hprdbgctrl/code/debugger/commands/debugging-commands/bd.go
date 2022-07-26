package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\bd.cpp"
var bdBuf string

type (
	Bd interface {
		//Fn() (ok bool)
	}
	bd struct{}
)

func Newbd() Bd { return &bd{} }
