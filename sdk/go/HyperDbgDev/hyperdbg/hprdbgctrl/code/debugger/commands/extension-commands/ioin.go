package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\ioin.cpp"
var ioinBuf string

type (
	Ioin interface {
		//Fn() (ok bool)
	}
	ioin struct{}
)

func Newioin() Ioin { return &ioin{} }
