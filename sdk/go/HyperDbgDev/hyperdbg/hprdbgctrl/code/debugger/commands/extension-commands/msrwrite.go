package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\msrwrite.cpp"
var msrwriteBuf string

type (
	Msrwrite interface {
		//Fn() (ok bool)
	}
	msrwrite struct{}
)

func Newmsrwrite() Msrwrite { return &msrwrite{} }
