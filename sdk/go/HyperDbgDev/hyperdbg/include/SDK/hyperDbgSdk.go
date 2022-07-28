package SDK

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\HyperDbgSdk.h"
var hyperDbgSdkBuf string

type (
	HyperDbgSdk interface {
		//Fn() (ok bool)
	}
	hyperDbgSdk struct{}
)

func NewhyperDbgSdk() HyperDbgSdk { return &hyperDbgSdk{} }
