package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func NewBreak() ux.Widget {
	t := ux.NewTreeTable(Break{})
	t.TableContext = ux.TableContext[Break]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[Break]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[Break]) (cells []ux.CellData) {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: fmt.Sprintf("%016X", n.Data.Type)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Address)},
			//	{Text: n.Data.ModuleLabelAndEx},
			//	{Text: n.Data.Status},
			//	{Text: n.Data.Disassembly},
			//	{Text: fmt.Sprintf("%d", n.Data.Hit)},
			//	{Text: n.Data.Summary},
			// }
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[Break], rows []ux.CellData) Break {
			return ux.UnmarshalRow[Break](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for i := range 100 {
				ts := Break{
					Type:             1,
					Address:          0x00007FF838E51030 + i,
					ModuleLabelAndEx: "<x64dbg.exe.OptionalHeader.AddressOfEntryPoint>",
					Status:           "一次性",
					Disassembly:      "sub rsp,28",
					Hit:              0,
					Summary:          "入口断点",
				}
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "break",
		IsDocument: false,
	}

	return t
}

type Break struct {
	Type             int
	Address          int
	ModuleLabelAndEx string
	Status           string
	Disassembly      string // Assemble Disassembly
	Hit              int
	Summary          string
}
