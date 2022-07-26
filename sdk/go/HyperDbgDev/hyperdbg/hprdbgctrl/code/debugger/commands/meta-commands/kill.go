package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\kill.cpp"
var killBuf string

type (
	Kill interface {
		//Fn() (ok bool)
	}
	kill struct{}
)

func Newkill() Kill { return &kill{} }
