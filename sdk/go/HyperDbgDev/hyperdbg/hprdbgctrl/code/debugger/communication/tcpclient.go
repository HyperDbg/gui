package tcpclient

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\tcpclient.cpp"
var tcpclientBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)
func New() Interface { return &object{} }


