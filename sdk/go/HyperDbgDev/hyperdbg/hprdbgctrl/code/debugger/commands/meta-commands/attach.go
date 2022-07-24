package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\attach.cpp"
var attachBuf string

type (
	Attach interface {
		//Fn() (ok bool)
	}
	attach struct{}
)

func Newattach() Attach { return &attach{} }
