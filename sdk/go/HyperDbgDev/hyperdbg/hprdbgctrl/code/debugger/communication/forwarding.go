package communication

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\communication\\forwarding.cpp"
var forwardingBuf string

type (
	Forwarding interface {
		//Fn() (ok bool)
	}
	forwarding struct{}
)

func Newforwarding() Forwarding { return &forwarding{} }
