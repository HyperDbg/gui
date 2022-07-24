package common

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\common\\list.cpp"
var listBuf string

type (
	List interface {
		//Fn() (ok bool)
	}
	list struct{}
)

func Newlist() List { return &list{} }
