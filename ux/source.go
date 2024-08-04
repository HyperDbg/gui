package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/unison"
)

func LayoutSource() unison.Paneler {
	return unison.NewPanel()
	return widget.NewCodeView("log.log")
}
