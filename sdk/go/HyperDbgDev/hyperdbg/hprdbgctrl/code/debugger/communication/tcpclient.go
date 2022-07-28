package communication

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\tcpclient.cpp"
var tcpclientBuf string

type (
	Tcpclient interface {
		//Fn() (ok bool)
	}
	tcpclient struct{}
)

func Newtcpclient() Tcpclient { return &tcpclient{} }
