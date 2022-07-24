package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\epthook2.cpp"
var epthook2Buf string

type (
	Epthook2 interface {
		//Fn() (ok bool)
	}
	epthook2 struct{}
)

func Newepthook2() Epthook2 { return &epthook2{} }
