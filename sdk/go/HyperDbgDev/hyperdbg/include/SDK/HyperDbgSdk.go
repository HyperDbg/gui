package HyperDbgSdk

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\HyperDbgSdk.h"
var HyperDbgSdkBuf string

type (
	Interface interface {
		//Fn() (ok bool)
	}
	object struct{}
)

func New() Interface { return &object{} }
