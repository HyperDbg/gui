package transparency

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\transparency\\gaussian-rng.cpp"
var gaussianRngBuf string

type (
	GaussianRng interface {
		//Fn() (ok bool)
	}
	gaussianRng struct{}
)

func NewgaussianRng() GaussianRng { return &gaussianRng{} }
