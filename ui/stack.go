package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func NewStack() ux.Widget {
	t := ux.NewTreeTable(CallStack{})
	t.TableContext = ux.TableContext[CallStack]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[CallStack]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[CallStack]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[CallStack], rows []ux.CellData) CallStack {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: fmt.Sprintf("%016X", n.Data.ThreadId)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Address)},
			//	{Text: fmt.Sprintf("%016X", n.Data.ReturnTo)},
			//	{Text: fmt.Sprintf("%016X", n.Data.ReturnFrom)},
			//	{Text: fmt.Sprintf("%d", n.Data.Size)},
			//	{Text: n.Data.Level},
			//	{Text: n.Data.Notes},
			// }

			return ux.UnmarshalRow[CallStack](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for i := range 100 {
				ts := CallStack{
					ThreadId:   0,
					Address:    0x00000002429BF7B8 + i,
					ReturnTo:   0x00007FF884F43EB3,
					ReturnFrom: 0x00007FF884F9A184,
					Size:       280,
					Level:      "系统模块",
					Notes:      "ntdll.EtwLogTraceEvent+19F74",
				}
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "stack",
		IsDocument: false,
	}

	return t
}

type CallStack struct {
	ThreadId   int
	Address    int
	ReturnTo   int
	ReturnFrom int
	Size       int
	Level      string
	Notes      string
}
