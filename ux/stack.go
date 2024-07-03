package ux

import (
	"fmt"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

func LayoutStack() unison.Paneler {
	table, header := widget.NewTable(CallStack{}, widget.TableContext[CallStack]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[CallStack]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.ThreadId)},
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("%016X", node.Data.ReturnTo)},
				{Text: fmt.Sprintf("%016X", node.Data.ReturnFrom)},
				{Text: fmt.Sprintf("%d", node.Data.Size)},
				{Text: node.Data.Level},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[CallStack]) {
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
				root.AddChildByData(ts)
			}
		},
		JsonName:   "stack",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(table, header)
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
