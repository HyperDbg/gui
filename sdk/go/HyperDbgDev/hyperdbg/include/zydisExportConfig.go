package include

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\ZydisExportConfig.h"
var zydisExportConfigBuf string

type (
	ZydisExportConfig interface {
		//Fn() (ok bool)
	}
	zydisExportConfig struct{}
)

func NewzydisExportConfig() ZydisExportConfig { return &zydisExportConfig{} }
