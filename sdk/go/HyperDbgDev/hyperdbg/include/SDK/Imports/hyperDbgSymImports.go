package Imports

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Imports\\HyperDbgSymImports.h"
var hyperDbgSymImportsBuf string

type (
	HyperDbgSymImports interface {
		//Fn() (ok bool)
	}
	hyperDbgSymImports struct{}
)

func NewhyperDbgSymImports() HyperDbgSymImports { return &hyperDbgSymImports{} }
