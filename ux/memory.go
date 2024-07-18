package ux

import (
	"fmt"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

type Memory struct {
	Address           int
	Size              int
	Level             string
	PageInfo          string
	AssignmentType    string
	PageProtection    string
	InitialProtection string
}

func LayoutMemory() unison.Paneler {
	table, header := widget.NewTable(Memory{}, widget.TableContext[Memory]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Memory]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum()
			}
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
		UnmarshalRow: func(node *widget.Node[Memory], values []string) {
			mylog.Todo("UnmarshalRow")
		},
		SelectionChangedCallback: func(root *widget.Node[Memory]) {
			mylog.Todo("SelectionChangedCallback")
		},
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
		JsonName:   "memory",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
}
