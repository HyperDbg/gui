package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\be.cpp"
var beBuf string

type (
	Be interface {
		//Fn() (ok bool)
	}
	be struct{}
)

func Newbe() Be { return &be{} }
