package driverloader

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\driver-loader\\install.cpp"
var installBuf string

type (
	Install interface {
		//Fn() (ok bool)
	}
	install struct{}
)

func Newinstall() Install { return &install{} }
