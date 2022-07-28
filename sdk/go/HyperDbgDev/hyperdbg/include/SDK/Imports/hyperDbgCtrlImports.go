package Imports

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Imports\\HyperDbgCtrlImports.h"
var hyperDbgCtrlImportsBuf string

type (
	HyperDbgCtrlImports interface {
		//Fn() (ok bool)
	}
	hyperDbgCtrlImports struct{}
)

func NewhyperDbgCtrlImports() HyperDbgCtrlImports { return &hyperDbgCtrlImports{} }
