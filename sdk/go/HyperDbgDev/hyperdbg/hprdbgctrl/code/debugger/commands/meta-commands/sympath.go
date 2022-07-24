package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\sympath.cpp"
var sympathBuf string

type (
	Sympath interface {
		//Fn() (ok bool)
	}
	sympath struct{}
)

func Newsympath() Sympath { return &sympath{} }
