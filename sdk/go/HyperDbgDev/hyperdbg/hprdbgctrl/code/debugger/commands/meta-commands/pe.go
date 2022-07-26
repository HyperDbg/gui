package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\pe.cpp"
var peBuf string

type (
	Pe interface {
		//Fn() (ok bool)
	}
	pe struct{}
)

func Newpe() Pe { return &pe{} }
