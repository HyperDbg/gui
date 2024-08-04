package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/unison"
)

func LayoutScript() unison.Paneler {
	path := "plugin/scriptGen/user-mode-memory-allocations.ds"
	path = "plugin/scriptGen/cpuid.ds"
	path = "plugin/scriptGen/1.carbon"
	return widget.NewDsScriptView(path)
}
