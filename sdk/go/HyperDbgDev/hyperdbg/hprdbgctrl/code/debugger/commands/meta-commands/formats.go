package metacommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\meta-commands\\formats.cpp"
var formatsBuf string

type (
	Formats interface {
		//Fn() (ok bool)
	}
	formats struct{}
)

func Newformats() Formats { return &formats{} }
