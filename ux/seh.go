package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"github.com/richardwilkes/unison"
)

func LayoutSeh(parent unison.Paneler) unison.Paneler {
	SehTable(parent)
}

type Seh struct {
	Address                   int `format:"%016X"`
	ExceptionHandlingRoutines string
	Label                     string
	Notes                     string
}

func SehTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*Seh, 100)
	for i := range breaks {
		ts := &Seh{
			Address:                   0x00000002429BF7B8 + i,
			ExceptionHandlingRoutines: "xxx",
			Label:                     "www",
			Notes:                     "this is notes",
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetReadOnly(true)
	tv.SetSlice(&breaks)
	return tv
}
