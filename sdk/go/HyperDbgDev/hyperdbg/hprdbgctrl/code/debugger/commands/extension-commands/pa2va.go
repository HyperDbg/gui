package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\pa2va.cpp"
var pa2vaBuf string

type (
	Pa2va interface {
		//Fn() (ok bool)
	}
	pa2va struct{}
)

func Newpa2va() Pa2va { return &pa2va{} }
