package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\communication.h"
var communicationBuf string

type (
	Communication interface {
		//Fn() (ok bool)
	}
	communication struct{}
)

func Newcommunication() Communication { return &communication{} }
