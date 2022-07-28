package hprdbgctrl

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\pch.h"
var pchBuf string

type (
	Pch interface {
		//Fn() (ok bool)
	}
	pch struct{}
)

func Newpch() Pch { return &pch{} }
