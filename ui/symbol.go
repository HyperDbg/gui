package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func NewSymbol() ux.Widget {
	t := ux.NewTreeTable(Symbol{})
	t.TableContext = ux.TableContext[Symbol]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Symbol]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Symbol]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Symbol], rows []ux.CellData) Symbol {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: fmt.Sprintf("%016X", n.Data.BaseAddress)},
			//	{Text: n.Data.Module},
			//	{Text: n.Data.Level},
			//	{Text: n.Data.Path},
			//	{Text: fmt.Sprintf("%016X", n.Data.Address)},
			// }
			return ux.UnmarshalRow[Symbol](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for range 100 {
				ts := Symbol{
					BaseAddress: 0x00007FF884ED0000,
					Module:      "ntdll.dll",
					Level:       "系统模块",
					Path:        "C:\\Windows\\System32\\ntdll.dll",
					Address:     0, // 状态=Unloaded
				}
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "Symbol",
		IsDocument: false,
	}
	return t
}

type Symbol struct {
	BaseAddress int
	Module      string
	Level       string
	Path        string
	Address     int
}
