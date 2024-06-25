package main

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript(parent unison.Paneler) unison.Paneler {
	path := "ux/send_version.scala"
	// path = "ux/plugin/scriptGen/cpuid.ds"
	return widget.NewCodeEditor(parent, path)
}
