package main

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/states"
)

func pageCallStack(parent *gi.Frame) {
	CallStackTable(parent)
}

type CallStack struct {
	ThreadId   int
	Address    int `format:"%016X"`
	ReturnTo   int `format:"%016X"`
	ReturnFrom int `format:"%016X"`
	Size       int
	Level      string
	Notes      string
}

func CallStackTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*CallStack, 100)
	for i := range breaks {
		ts := &CallStack{
			ThreadId:   0,
			Address:    0x00000002429BF7B8 + i,
			ReturnTo:   0x00007FF884F43EB3,
			ReturnFrom: 0x00007FF884F9A184,
			Size:       280,
			Level:      "系统模块",
			Notes:      "ntdll.EtwLogTraceEvent+19F74",
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&breaks)
	return tv
}
