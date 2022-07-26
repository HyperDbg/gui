package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\common.h"
var commonBuf string

type (
	Common interface {
		//Fn() (ok bool)
	}
	common struct{}
)

func Newcommon() Common { return &common{} }
