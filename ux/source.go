package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutSource(parent unison.Paneler) unison.Paneler {
	return widget.NewCodeView("log.log")
}
