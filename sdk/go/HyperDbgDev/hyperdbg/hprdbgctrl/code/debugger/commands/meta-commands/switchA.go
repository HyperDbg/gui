package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\switch.cpp"
var switchABuf string

type (
	SwitchA interface {
		//Fn() (ok bool)
	}
	switchA struct{}
)

func NewswitchA() SwitchA { return &switchA{} }
