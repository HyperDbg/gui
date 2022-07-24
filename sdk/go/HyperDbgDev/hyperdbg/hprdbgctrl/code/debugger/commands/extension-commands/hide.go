package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\hide.cpp"
var hideBuf string

type (
	Hide interface {
		//Fn() (ok bool)
	}
	hide struct{}
)

func Newhide() Hide { return &hide{} }
