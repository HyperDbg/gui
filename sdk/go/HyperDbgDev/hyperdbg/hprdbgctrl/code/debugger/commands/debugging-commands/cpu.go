package debuggingcommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\debugging-commands\\cpu.cpp"
var cpuBuf string

type (
	Cpu interface {
		//Fn() (ok bool)
	}
	cpu struct{}
)

func Newcpu() Cpu { return &cpu{} }
