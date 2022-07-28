package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\msrread.cpp"
var msrreadBuf string

type (
	Msrread interface {
		//Fn() (ok bool)
	}
	msrread struct{}
)

func Newmsrread() Msrread { return &msrread{} }
