package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\sym.cpp"
var symBuf string

type (
	Sym interface {
		//Fn() (ok bool)
	}
	sym struct{}
)

func Newsym() Sym { return &sym{} }
