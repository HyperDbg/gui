package main

import (
	"fmt"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutSymbol(parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(Symbol{}, widget.TableContext[Symbol]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Symbol]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.BaseAddress)},
				{Text: node.Data.Module},
				{Text: node.Data.Level},
				{Text: node.Data.Path},
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Symbol]) {
			for range 100 {
				ts := Symbol{
					BaseAddress: 0x00007FF884ED0000,
					Module:      "ntdll.dll",
					Level:       "系统模块",
					Path:        "C:\\Windows\\System32\\ntdll.dll",
					Address:     0, // 状态=Unloaded
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "symbols",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}

type Symbol struct {
	BaseAddress int
	Module      string
	Level       string
	Path        string
	Address     int
}
