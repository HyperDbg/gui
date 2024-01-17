package main

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/states"
)

func pageBreak(parent *gi.Frame) {
	breakTable(parent)
}

type Break struct {
	Type             int
	Address          int
	ModuleLabelAndEx string
	Status           string
	Disassembly      string //Assemble Disassembly
	Hit              int
	Summary          string
}

func breakTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*Break, 100)
	for i := range breaks {
		ts := &Break{
			Type:             1,
			Address:          0x00007FF838E51030 + i,
			ModuleLabelAndEx: "<x64dbg.exe.OptionalHeader.AddressOfEntryPoint>",
			Status:           "一次性",
			Disassembly:      "sub rsp,28",
			Hit:              0,
			Summary:          "入口断点",
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetState(true, states.ReadOnly)
	tv.SetSlice(&breaks)
	return tv
}
