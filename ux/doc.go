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
		//完成度
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
			//分类
			//单个调试器支持x64和x32应用程序
			//pe解析器
			//scylla
			//packer检测，完成themida
			//ark 完成度 ssdt
			//硬件信息hook
			//unicorn binee
			//ghidra 反编译引擎集成
			//脱壳脚本 尚未实现编辑功能
			//脱壳插件 go plugin？ tmd
			//树形反汇编，目前全部调试器都是格式化为线性反汇编
			//hex编辑器 未完成
			//VMware bypass
			//绑定引擎 80%
			//注册快捷菜单 bug
			//解析堆栈参数  未完成
			//命令api绑定 80%
			//

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
