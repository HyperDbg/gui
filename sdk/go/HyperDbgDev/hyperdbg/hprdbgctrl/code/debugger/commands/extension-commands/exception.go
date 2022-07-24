package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\exception.cpp"
var exceptionBuf string

type (
	Exception interface {
		//Fn() (ok bool)
	}
	exception struct{}
)

func Newexception() Exception { return &exception{} }
