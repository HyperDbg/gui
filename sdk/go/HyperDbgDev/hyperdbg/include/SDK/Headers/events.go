package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Events.h"
var eventsBuf string

type (
	Events interface {
		//Fn() (ok bool)
	}
	events struct{}
)

func Newevents() Events { return &events{} }
