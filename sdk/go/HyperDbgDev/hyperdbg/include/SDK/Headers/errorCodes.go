package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\ErrorCodes.h"
var errorCodesBuf string

type (
	ErrorCodes interface {
		//Fn() (ok bool)
	}
	errorCodes struct{}
)

func NewerrorCodes() ErrorCodes { return &errorCodes{} }

const ()
