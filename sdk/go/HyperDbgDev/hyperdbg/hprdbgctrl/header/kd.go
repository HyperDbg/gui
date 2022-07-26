package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\kd.h"
var kdBuf string

type (
	Kd interface {
		//Fn() (ok bool)
	}
	kd struct{}
)

func Newkd() Kd { return &kd{} }
