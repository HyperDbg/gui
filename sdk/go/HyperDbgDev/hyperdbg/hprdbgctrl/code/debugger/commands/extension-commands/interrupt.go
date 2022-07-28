package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\interrupt.cpp"
var interruptBuf string

type (
	Interrupt interface {
		//Fn() (ok bool)
	}
	interrupt struct{}
)

func Newinterrupt() Interrupt { return &interrupt{} }
