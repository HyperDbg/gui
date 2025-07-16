package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func NewSeh() ux.Widget {
	t := ux.NewTreeTable(Seh{})
	t.TableContext = ux.TableContext[Seh]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Seh]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Seh]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Seh], rows []ux.CellData) Seh {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: fmt.Sprintf("%016X", n.Data.Address)},
			//	{Text: n.Data.ExceptionHandlingRoutines},
			//	{Text: n.Data.Label},
			//	{Text: n.Data.Notes},
			// }
			return ux.UnmarshalRow[Seh](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for i := range 100 {
				ts := Seh{
					Address:                   0x00000002429BF7B8 + i,
					ExceptionHandlingRoutines: "xxx",
					Label:                     "www",
					Notes:                     "this is notes",
				}
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "Seh",
		IsDocument: false,
	}
	return t
}

type Seh struct {
	Address                   int
	ExceptionHandlingRoutines string
	Label                     string
	Notes                     string
}
