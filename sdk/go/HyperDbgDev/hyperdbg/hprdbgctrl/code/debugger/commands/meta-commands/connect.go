package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\connect.cpp"
var connectBuf string

type (
	Connect interface {
		//Fn() (ok bool)
	}
	connect struct{}
)

func Newconnect() Connect { return &connect{} }
