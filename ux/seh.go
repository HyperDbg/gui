package ux

import (
	"fmt"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutSeh() unison.Paneler {
	table, header := widget.NewTable(Seh{}, widget.TableContext[Seh]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Seh]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum("todo")
			}
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ExceptionHandlingRoutines},
				{Text: node.Data.Label},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow: func(node *widget.Node[Seh], values []string) {
			mylog.Todo("UnmarshalRow")
		},
		SelectionChangedCallback: func(root *widget.Node[Seh]) {
			mylog.Todo("SelectionChangedCallback")
		},
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
	return widget.NewTableScrollPanel(table, header)
}

type Seh struct {
	Address                   int
	ExceptionHandlingRoutines string
	Label                     string
	Notes                     string
}
