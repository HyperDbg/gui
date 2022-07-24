package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\exit.cpp"
var exitBuf string

type (
	Exit interface {
		//Fn() (ok bool)
	}
	exit struct{}
)

func Newexit() Exit { return &exit{} }
