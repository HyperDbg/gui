package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\debug.cpp"
var debugBuf string

type (
	Debug interface {
		//Fn() (ok bool)
	}
	debug struct{}
)

func Newdebug() Debug { return &debug{} }
