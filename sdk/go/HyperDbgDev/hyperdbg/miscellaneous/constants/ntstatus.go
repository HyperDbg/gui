package constants

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\miscellaneous\\constants\\ntstatus.txt"
var ntstatusBuf string

type (
	Ntstatus interface {
		//Fn() (ok bool)
	}
	ntstatus struct{}
)

func Newntstatus() Ntstatus { return &ntstatus{} }
