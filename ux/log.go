package ux

import (
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutLog() unison.Paneler {
	return widget.NewLogView(mylog.Body())
}
