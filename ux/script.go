package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
)

func pageScript(parent *gi.Frame) {
	ScriptTable(parent)
}

type Script struct {
	Line    int
	Contest string
	Info    string
}

func ScriptTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*Script, 100)
	for i := range breaks {
		ts := &Script{
			Line:    i,
			Contest: "run",
			Info:    "run code",
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetReadOnly(true)
	tv.SetSlice(&breaks)
	return tv
}
