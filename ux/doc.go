package ux

import (
	"fmt"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutDoc() unison.Paneler {
	type doc struct { // 功能分类和描述，实现与否，gif demo链接到别的仓库? 4个字段，同步更新readme
		Function    string
		Description string
		Implement   bool
		DemoLink    string
	}
	table, header := widget.NewTable(doc{}, widget.TableContext[doc]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[doc]) (cells []widget.CellData) {
			if node.Container() {
				node.Sum("todo")
			}
			return []widget.CellData{
				{Text: node.Data.Function},
				{Text: node.Data.Description},
				{Text: fmt.Sprintf("%t", node.Data.Implement)},
				{Text: node.Data.DemoLink},
			}
		},
		UnmarshalRow: func(node *widget.Node[doc], values []string) {
			mylog.Todo("UnmarshalRow")
		},
		SelectionChangedCallback: func(root *widget.Node[doc]) {
			mylog.Todo("SelectionChangedCallback")
		},
		SetRootRowsCallBack: func(root *widget.Node[doc]) {
			for range 100 {
				ts := doc{
					Function:    "",
					Description: "",
					Implement:   false,
					DemoLink:    "",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "doc",
		IsDocument: false, //todo
	})
	return widget.NewTableScrollPanel(table, header)
}
