package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\namedpipe.h"
var namedpipeBuf string

type (
	Namedpipe interface {
		//Fn() (ok bool)
	}
	namedpipe struct{}
)

func Newnamedpipe() Namedpipe { return &namedpipe{} }
