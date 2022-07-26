package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\cpuid.cpp"
var cpuidBuf string

type (
	Cpuid interface {
		//Fn() (ok bool)
	}
	cpuid struct{}
)

func Newcpuid() Cpuid { return &cpuid{} }
