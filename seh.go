package main

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/states"
)

func pageCallSeh(parent *gi.Frame) {
	SehTable(parent)
}

type Seh struct {
	Address                   int
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
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&breaks)
	return tv
}
