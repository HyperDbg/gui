package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\status.cpp"
var statusBuf string

type (
	Status interface {
		//Fn() (ok bool)
	}
	status struct{}
)

func Newstatus() Status { return &status{} }
