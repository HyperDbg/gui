package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
)

func pageSymbol(parent *gi.Frame) {
	SymbolTable(parent) // todo 左右都是标，适合树形表格
}

type Symbol struct {
	BaseAddress int `format:"%016X"`
	Module      string
	Level       string
	Path        string
	Address     int `format:"%016X"`
}

func SymbolTable(frame *gi.Frame) *giv.TableView {
	breaks := make([]*Symbol, 100)
	for i := range breaks {
		ts := &Symbol{
			BaseAddress: 0x00007FF884ED0000,
			Module:      "ntdll.dll",
			Level:       "系统模块",
			Path:        "C:\\Windows\\System32\\ntdll.dll",
			Address:     0, // 状态=Unloaded
		}
		breaks[i] = ts
	}
	tv := giv.NewTableView(frame, "tv")
	tv.SetReadOnly(true)
	tv.SetSlice(&breaks)
	return tv
}
