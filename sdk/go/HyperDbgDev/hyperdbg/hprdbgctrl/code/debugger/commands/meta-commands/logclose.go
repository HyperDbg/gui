package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\logclose.cpp"
var logcloseBuf string

type (
	Logclose interface {
		//Fn() (ok bool)
	}
	logclose struct{}
)

func Newlogclose() Logclose { return &logclose{} }
