package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript() unison.Paneler {
	path := "send_version.scala"
	// path = "plugin/scriptGen/cpuid.ds"
	return widget.NewCodeEditor(path)
}
