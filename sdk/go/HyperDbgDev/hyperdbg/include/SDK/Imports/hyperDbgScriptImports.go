package Imports

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Imports\\HyperDbgScriptImports.h"
var hyperDbgScriptImportsBuf string

type (
	HyperDbgScriptImports interface {
		//Fn() (ok bool)
	}
	hyperDbgScriptImports struct{}
)

func NewhyperDbgScriptImports() HyperDbgScriptImports { return &hyperDbgScriptImports{} }
