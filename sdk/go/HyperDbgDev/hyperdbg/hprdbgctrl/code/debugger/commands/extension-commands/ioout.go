package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\ioout.cpp"
var iooutBuf string

type (
	Ioout interface {
		//Fn() (ok bool)
	}
	ioout struct{}
)

func Newioout() Ioout { return &ioout{} }
