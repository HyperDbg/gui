package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\restart.cpp"
var restartBuf string

type (
	Restart interface {
		//Fn() (ok bool)
	}
	restart struct{}
)

func Newrestart() Restart { return &restart{} }
