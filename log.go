package main

import (
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutLog() unison.Paneler {
	return unison.NewPanel()
	return widget.NewCodeView("log.log")
}
