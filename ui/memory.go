package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
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

func NewMemory() ux.Widget {
	t := ux.NewTreeTable(Memory{})
	t.TableContext = ux.TableContext[Memory]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Memory]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Memory]) (cells []ux.CellData) {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: fmt.Sprintf("%d", n.Data.Address)},
			//	{Text: fmt.Sprintf("%d", n.Data.Size)},
			//	{Text: n.Data.Level},
			//	{Text: n.Data.PageInfo},
			//	{Text: n.Data.AssignmentType},
			//	{Text: n.Data.PageProtection},
			//	{Text: n.Data.InitialProtection},
			// }
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Memory], rows []ux.CellData) Memory {
			return ux.UnmarshalRow[Memory](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
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
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "memory",
		IsDocument: false,
	}
	return t
}
