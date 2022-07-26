package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\pause.cpp"
var pauseBuf string

type (
	Pause interface {
		//Fn() (ok bool)
	}
	pause struct{}
)

func Newpause() Pause { return &pause{} }
