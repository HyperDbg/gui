package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\help.cpp"
var helpBuf string

type (
	Help interface {
		//Fn() (ok bool)
	}
	help struct{}
)

func Newhelp() Help { return &help{} }
