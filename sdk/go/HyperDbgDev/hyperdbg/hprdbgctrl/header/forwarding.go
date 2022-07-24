package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\forwarding.h"
var forwardingBuf string

type (
	Forwarding interface {
		//Fn() (ok bool)
	}
	forwarding struct{}
)

func Newforwarding() Forwarding { return &forwarding{} }
