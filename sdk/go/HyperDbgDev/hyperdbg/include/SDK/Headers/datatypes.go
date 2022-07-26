package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Datatypes.h"
var datatypesBuf string

type (
	Datatypes interface {
		//Fn() (ok bool)
	}
	datatypes struct{}
)

func Newdatatypes() Datatypes { return &datatypes{} }
