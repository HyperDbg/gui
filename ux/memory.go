package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
)

func pageMemory(parent *gi.Frame) {
	MemoryTable(parent)
}

type Memory struct {
	Address           int `format:"%016X"`
	Size              int
	Level             string
	PageInfo          string
	AssignmentType    string
	PageProtection    string
	InitialProtection string
}

func MemoryTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*Memory, 100)
	for i := range breaks {
		ts := &Memory{
			Address:           0x000001027D020000 + i,
			Size:              0o000000000011000,
			Level:             "用户模块",
			PageInfo:          "\\Device\\HarddiskVolume3\\Windows\\System32\\C_1252.NLS",
			AssignmentType:    "MAP",
			PageProtection:    "-R---",
			InitialProtection: "-R---",
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetReadOnly(true)
	tv.SetSlice(&breaks)
	return tv
}
