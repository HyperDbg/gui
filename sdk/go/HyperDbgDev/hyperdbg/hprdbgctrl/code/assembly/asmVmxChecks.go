package assembly

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\assembly\\asm-vmx-checks.asm"
var asmVmxChecksBuf string

type (
	AsmVmxChecks interface {
		//Fn() (ok bool)
	}
	asmVmxChecks struct{}
)

func NewasmVmxChecks() AsmVmxChecks { return &asmVmxChecks{} }
