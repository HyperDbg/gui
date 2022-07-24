package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\install.h"
var installBuf string

type (
	Install interface {
		//Fn() (ok bool)
	}
	install struct{}
)

func Newinstall() Install { return &install{} }
