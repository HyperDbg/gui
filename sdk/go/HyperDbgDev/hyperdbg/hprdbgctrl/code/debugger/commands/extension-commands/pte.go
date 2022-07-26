package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\pte.cpp"
var pteBuf string

type (
	Pte interface {
		//Fn() (ok bool)
	}
	pte struct{}
)

func Newpte() Pte { return &pte{} }
