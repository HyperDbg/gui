package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\listen.cpp"
var listenBuf string

type (
	Listen interface {
		//Fn() (ok bool)
	}
	listen struct{}
)

func Newlisten() Listen { return &listen{} }
