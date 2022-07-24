package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\unhide.cpp"
var unhideBuf string

type (
	Unhide interface {
		//Fn() (ok bool)
	}
	unhide struct{}
)

func Newunhide() Unhide { return &unhide{} }
