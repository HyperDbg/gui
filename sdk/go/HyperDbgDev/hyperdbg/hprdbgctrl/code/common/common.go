package common

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\common\\common.cpp"
var commonBuf string

type (
	Common interface {
		//Fn() (ok bool)
	}
	common struct{}
)

func Newcommon() Common { return &common{} }
