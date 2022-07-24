package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\process.cpp"
var processBuf string

type (
	Process interface {
		//Fn() (ok bool)
	}
	process struct{}
)

func Newprocess() Process { return &process{} }
