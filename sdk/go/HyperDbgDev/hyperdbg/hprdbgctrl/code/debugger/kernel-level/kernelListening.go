package kernellevel

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\kernel-level\\kernel-listening.cpp"
var kernelListeningBuf string

type (
	KernelListening interface {
		//Fn() (ok bool)
	}
	kernelListening struct{}
)

func NewkernelListening() KernelListening { return &kernelListening{} }
