package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutLog() unison.Paneler {
	return widget.NewLogView()
}
