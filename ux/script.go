package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript(parent unison.Paneler) unison.Paneler {
	path := "ux/send_version.scala"
	codeEditor := widget.NewCodeEditor(parent, path)
	return codeEditor
}
