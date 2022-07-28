package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\help.h"
var helpBuf string

type (
	Help interface {
		//Fn() (ok bool)
	}
	help struct{}
)

func Newhelp() Help { return &help{} }
