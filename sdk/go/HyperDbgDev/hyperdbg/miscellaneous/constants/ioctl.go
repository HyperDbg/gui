package constants

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\miscellaneous\\constants\\ioctl.txt"
var ioctlBuf string

type (
	Ioctl interface {
		//Fn() (ok bool)
	}
	ioctl struct{}
)

func Newioctl() Ioctl { return &ioctl{} }
