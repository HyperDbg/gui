package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\thread.cpp"
var threadBuf string

type (
	Thread interface {
		//Fn() (ok bool)
	}
	thread struct{}
)

func Newthread() Thread { return &thread{} }
