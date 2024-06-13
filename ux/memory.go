package ux

import (
	"fmt"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

type Memory struct {
	Address           int `format:"%016X"`
	Size              int
	Level             string
	PageInfo          string
	AssignmentType    string
	PageProtection    string
	InitialProtection string
}

func LayoutMemory(parent unison.Paneler) unison.Paneler {

	table, header := widget.NewTable(Memory{}, widget.TableContext[Memory]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Memory]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%d", node.Data.Address)},
				{Text: fmt.Sprintf("%d", node.Data.Size)},
				{Text: node.Data.Level},
				{Text: node.Data.PageInfo},
				{Text: node.Data.AssignmentType},
				{Text: node.Data.PageProtection},
				{Text: node.Data.InitialProtection},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Memory]) {
			for i := range 100 {
				ts := Memory{
					Address:           0x000001027D020000 + i,
					Size:              0o000000000011000,
					Level:             "用户模块",
					PageInfo:          "\\Device\\HarddiskVolume3\\Windows\\System32\\C_1252.NLS",
					AssignmentType:    "MAP",
					PageProtection:    "-R---",
					InitialProtection: "-R---",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}
