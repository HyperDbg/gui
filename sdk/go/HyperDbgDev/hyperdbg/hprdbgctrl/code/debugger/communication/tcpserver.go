package communication

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\tcpserver.cpp"
var tcpserverBuf string

type (
	Tcpserver interface {
		//Fn() (ok bool)
	}
	tcpserver struct{}
)

func Newtcpserver() Tcpserver { return &tcpserver{} }
