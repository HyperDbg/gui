package include

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\ZycoreExportConfig.h"
var zycoreExportConfigBuf string

type (
	ZycoreExportConfig interface {
		//Fn() (ok bool)
	}
	zycoreExportConfig struct{}
)

func NewzycoreExportConfig() ZycoreExportConfig { return &zycoreExportConfig{} }
