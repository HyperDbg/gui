package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\pmc.cpp"
var pmcBuf string

type (
	Pmc interface {
		//Fn() (ok bool)
	}
	pmc struct{}
)

func Newpmc() Pmc { return &pmc{} }
