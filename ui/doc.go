package ui

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ux"
)

func LayoutDoc() ux.Widget {
	type doc struct { // 功能分类和描述，实现与否，gif demo链接到别的仓库? 4个字段，同步更新readme
		Function    string
		Description string
		Implement   bool
		DemoLink    string
		// 完成度
	}

	t := ux.NewTreeTable(doc{})
	t.TableContext = ux.TableContext[doc]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[doc]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[doc]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[doc], rows []ux.CellData) doc {
			if n.Container() {
				n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: n.Data.Function},
			//	{Text: n.Data.Description},
			//	{Text: fmt.Sprintf("%t", n.Data.Implement)},
			//	{Text: n.Data.DemoLink},
			// }
			return ux.UnmarshalRow[doc](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			// 分类
			// 内存读写
			// 寄存器读写
			// 单个调试器支持x64和x32应用程序
			// pe解析器
			// scylla
			// packer检测，完成themida
			// ark 完成度 ssdt
			// 硬件信息hook
			// unicorn binee
			// ghidra 反编译引擎集成
			// 脱壳脚本 尚未实现编辑功能
			// 脱壳插件 go plugin？ tmd
			// 树形反汇编，目前全部调试器都是格式化为线性反汇编
			// hex编辑器 未完成
			// VMware bypass
			// 绑定引擎 80%
			// 注册快捷菜单 bug
			// 解析堆栈参数  未完成
			// 命令api绑定 80%
			for range 100 {
				ts := doc{
					Function:    "",
					Description: "",
					Implement:   false,
					DemoLink:    "",
				}
				t.Root.AddChildByData(ts)
			}
		},
		JsonName:   "doc",
		IsDocument: false,
	}

	return t
}
