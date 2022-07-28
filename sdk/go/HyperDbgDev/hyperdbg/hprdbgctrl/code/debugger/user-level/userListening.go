package userlevel

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\user-level\\user-listening.cpp"
var userListeningBuf string

type (
	UserListening interface {
		//Fn() (ok bool)
	}
	userListening struct{}
)

func NewuserListening() UserListening { return &userListening{} }
