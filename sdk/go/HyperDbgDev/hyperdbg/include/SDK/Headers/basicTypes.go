package Headers

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\BasicTypes.h"
var basicTypesBuf string

type (
	BasicTypes interface {
		//Fn() (ok bool)
	}
	basicTypes struct{}
)

func NewbasicTypes() BasicTypes { return &basicTypes{} }

const ()
