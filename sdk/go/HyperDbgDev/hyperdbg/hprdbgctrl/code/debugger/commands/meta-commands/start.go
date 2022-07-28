package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\start.cpp"
var startBuf string

type (
	Start interface {
		//Fn() (ok bool)
	}
	start struct{}
)

func Newstart() Start { return &start{} }
