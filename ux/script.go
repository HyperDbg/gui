package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript() unison.Paneler {
	path := "user-mode-memory-allocations.ds"
	path = "plugin/scriptGen/cpuid.ds"
	return widget.NewDsScriptView(path)
}
