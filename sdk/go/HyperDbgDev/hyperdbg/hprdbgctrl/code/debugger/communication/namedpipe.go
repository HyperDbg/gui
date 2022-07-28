package communication

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\namedpipe.cpp"
var namedpipeBuf string

type (
	Namedpipe interface {
		//Fn() (ok bool)
	}
	namedpipe struct{}
)

func Newnamedpipe() Namedpipe { return &namedpipe{} }
