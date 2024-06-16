package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript(parent unison.Paneler) unison.Paneler {
	path := "ux/send_version.scala"
	return widget.NewCodeEditor(parent, path)
}
