package ux

import (
	"fmt"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutBreak() unison.Paneler {
	table, header := widget.NewTable(Break{}, widget.TableContext[Break]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Break]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum("todo")
			}
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Type)},
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ModuleLabelAndEx},
				{Text: node.Data.Status},
				{Text: node.Data.Disassembly},
				{Text: fmt.Sprintf("%d", node.Data.Hit)},
				{Text: node.Data.Summary},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Break]) {
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
				root.AddChildByData(ts)
			}
		},
		JsonName:   "break",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
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
