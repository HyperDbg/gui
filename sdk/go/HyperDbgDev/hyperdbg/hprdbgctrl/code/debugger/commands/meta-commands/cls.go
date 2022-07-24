package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\cls.cpp"
var clsBuf string

type (
	Cls interface {
		//Fn() (ok bool)
	}
	cls struct{}
)

func Newcls() Cls { return &cls{} }
