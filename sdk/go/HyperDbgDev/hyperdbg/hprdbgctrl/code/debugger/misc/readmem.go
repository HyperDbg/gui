package misc

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\misc\\readmem.cpp"
var readmemBuf string

type (
	Readmem interface {
		//Fn() (ok bool)
	}
	readmem struct{}
)

func Newreadmem() Readmem { return &readmem{} }
