package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\epthook.cpp"
var epthookBuf string

type (
	Epthook interface {
		//Fn() (ok bool)
	}
	epthook struct{}
)

func Newepthook() Epthook { return &epthook{} }
