package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\disconnect.cpp"
var disconnectBuf string

type (
	Disconnect interface {
		//Fn() (ok bool)
	}
	disconnect struct{}
)

func Newdisconnect() Disconnect { return &disconnect{} }
