package header

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\header\\list.h"
var listBuf string

type (
	List interface {
		//Fn() (ok bool)
	}
	list struct{}
)

func Newlist() List { return &list{} }
