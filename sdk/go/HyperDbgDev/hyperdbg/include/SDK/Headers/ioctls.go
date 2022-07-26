package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Ioctls.h"
var ioctlsBuf string

type (
	Ioctls interface {
		//Fn() (ok bool)
	}
	ioctls struct{}
)

func Newioctls() Ioctls { return &ioctls{} }

const ()
