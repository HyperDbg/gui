package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Connection.h"
var connectionBuf string

type (
	Connection interface {
		//Fn() (ok bool)
	}
	connection struct{}
)

func Newconnection() Connection { return &connection{} }
