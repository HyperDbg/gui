package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\prealloc.cpp"
var preallocBuf string

type (
	Prealloc interface {
		//Fn() (ok bool)
	}
	prealloc struct{}
)

func Newprealloc() Prealloc { return &prealloc{} }
