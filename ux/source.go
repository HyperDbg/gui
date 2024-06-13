package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutSource(parent unison.Paneler) unison.Paneler {
	return unison.NewPanel()
	return widget.NewCodeView("log.log")
}
