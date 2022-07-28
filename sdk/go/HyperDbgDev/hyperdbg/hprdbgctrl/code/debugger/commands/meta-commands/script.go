package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\script.cpp"
var scriptBuf string

type (
	Script interface {
		//Fn() (ok bool)
	}
	script struct{}
)

func Newscript() Script { return &script{} }
