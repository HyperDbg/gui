package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\RequestStructures.h"
var requestStructuresBuf string

type (
	RequestStructures interface {
		//Fn() (ok bool)
	}
	requestStructures struct{}
)

func NewrequestStructures() RequestStructures { return &requestStructures{} }
