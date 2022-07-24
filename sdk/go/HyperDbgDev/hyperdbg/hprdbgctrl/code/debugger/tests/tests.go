package tests

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\tests\\tests.cpp"
var testsBuf string

type (
	Tests interface {
		//Fn() (ok bool)
	}
	tests struct{}
)

func Newtests() Tests { return &tests{} }
