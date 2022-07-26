package communication

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\remote-connection.cpp"
var remoteConnectionBuf string

type (
	RemoteConnection interface {
		//Fn() (ok bool)
	}
	remoteConnection struct{}
)

func NewremoteConnection() RemoteConnection { return &remoteConnection{} }
