package ux

import (
	"fmt"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutSeh(parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(Seh{}, widget.TableContext[Seh]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Seh]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ExceptionHandlingRoutines},
				{Text: node.Data.Label},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Seh]) {
			for i := range 100 {
				ts := Seh{
					Address:                   0x00000002429BF7B8 + i,
					ExceptionHandlingRoutines: "xxx",
					Label:                     "www",
					Notes:                     "this is notes",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "seh",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}

type Seh struct {
	Address                   int `format:"%016X"`
	ExceptionHandlingRoutines string
	Label                     string
	Notes                     string
}
