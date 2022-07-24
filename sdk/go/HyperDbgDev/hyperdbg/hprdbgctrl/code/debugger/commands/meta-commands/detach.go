package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\detach.cpp"
var detachBuf string

type (
	Detach interface {
		//Fn() (ok bool)
	}
	detach struct{}
)

func Newdetach() Detach { return &detach{} }
