package ux

import (
	"fmt"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutScript(parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(Script{}, widget.TableContext[Script]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Script]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%d", node.Data.Line)},
				{Text: node.Data.Contest},
				{Text: node.Data.Info},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Script]) {
			for i := range 100 {
				ts := Script{
					Line:    i,
					Contest: "run",
					Info:    "run code",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}

type Script struct {
	Line    int
	Contest string
	Info    string
}
