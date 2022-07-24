package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\crwrite.cpp"
var crwriteBuf string

type (
	Crwrite interface {
		//Fn() (ok bool)
	}
	crwrite struct{}
)

func Newcrwrite() Crwrite { return &crwrite{} }
