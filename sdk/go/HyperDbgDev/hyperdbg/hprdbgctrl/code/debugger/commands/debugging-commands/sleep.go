package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\sleep.cpp"
var sleepBuf string

type (
	Sleep interface {
		//Fn() (ok bool)
	}
	sleep struct{}
)

func Newsleep() Sleep { return &sleep{} }
