package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\logopen.cpp"
var logopenBuf string

type (
	Logopen interface {
		//Fn() (ok bool)
	}
	logopen struct{}
)

func Newlogopen() Logopen { return &logopen{} }
