package ark

import (
	"fmt"

	"github.com/ddkwork/app/ms/packer"

	"github.com/ddkwork/HyperDbg/sdk"
	"github.com/ddkwork/app"
	"github.com/ddkwork/app/ms"
	"github.com/ddkwork/app/ms/driverTool/environment"
	"github.com/ddkwork/app/ms/hook/winver"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream/orderedmap"
	"github.com/ddkwork/unison"
)

func main() {
	app.Run("ark tool", func(w *unison.Window) {
		content := w.Content()
		panel := widget.NewPanel()
		panel.AddChild(Layout())
		scrollPanelFill := widget.NewScrollPanelFill(panel)
		content.AddChild(scrollPanelFill)
	})
}

func arkTodo() {
	ms.DecodeTableByDll()
	println(winver.WindowVersion())
	ms.MiGetPteAddress()
	ms.DecodeTableByDll()
	ms.DecodeTableByDisassembly()
	ms.NtDeviceIoControlFile()
	// IopXxxControlFile()
	widget.NewExplorer(".")
	environment.New()
}

func Layout() *unison.Panel {
	type ark struct{ Name ArksKind }
	table, header := widget.NewTable(ark{}, widget.TableContext[ark]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[ark]) (cells []widget.CellData) {
			name := node.Data.Name.String()
			if node.Container() {
				name = node.Sum()
			}
			return []widget.CellData{{Text: name}}
		},
		UnmarshalRow: func(node *widget.Node[ark], values []string) {
			mylog.Todo("unmarshal row")
		},
		SelectionChangedCallback: func(root *widget.Node[ark]) {
			mylog.Todo("selection changed callback")
		},
		SetRootRowsCallBack: func(root *widget.Node[ark]) {
			for _, kind := range InvalidArksKind.Kinds() {
				root.AddChildByData(ark{kind})
			}
		},
		JsonName:   "ark",
		IsDocument: false,
	})

	splitPanel := widget.NewPanel()
	widget.SetScrollLayout(splitPanel, 2)

	left := widget.NewTableScrollPanel(table, header)
	layouts := orderedmap.New(InvalidArksKind, func() unison.Paneler { return widget.NewPanel() })
	layouts.Set(KernelTablesKind, func() unison.Paneler {
		table, header := widget.NewTable(ms.NtApi{}, widget.TableContext[ms.NtApi]{
			ContextMenuItems: nil,
			MarshalRow: func(node *widget.Node[ms.NtApi]) (cells []widget.CellData) {
				KernelBase := fmt.Sprintf("%016X", node.Data.KernelBase)
				if node.Container() {
					KernelBase = node.Sum()
				}
				return []widget.CellData{
					{Text: KernelBase},
					{Text: fmt.Sprintf("%016X", node.Data.ArgValue)},
					{Text: node.Data.Name},
					{Text: fmt.Sprintf("%04d / %08X", node.Data.Index, node.Data.Index)},
				}
			},
			UnmarshalRow: func(node *widget.Node[ms.NtApi], values []string) {
				mylog.Todo("UnmarshalRow")
			},
			SelectionChangedCallback: func(root *widget.Node[ms.NtApi]) {
				mylog.Todo("SelectionChangedCallback")
			},
			SetRootRowsCallBack: func(root *widget.Node[ms.NtApi]) {
				sysCall := ms.NewSysCall(int64(sdk.GetKernelBase()))
				sysCall.KeServiceDescriptorTable = ms.DecodeNtApi("C:\\Windows\\System32\\ntdll.dll")
				sysCall.KeServiceDescriptorTableShadow = ms.DecodeNtApi("C:\\Windows\\System32\\win32u.dll")
				NtTableContainer := widget.NewContainerNode("NtTable", ms.NtApi{})
				Win32kTableContainer := widget.NewContainerNode("Win32kTable", ms.NtApi{})
				for _, api := range sysCall.KeServiceDescriptorTable {
					api.Index++
					NtTableContainer.AddChildByData(api)
				}
				for _, api := range sysCall.KeServiceDescriptorTableShadow {
					api.Index++
					Win32kTableContainer.AddChildByData(api)
				}
				root.AddChild(NtTableContainer)
				root.AddChild(Win32kTableContainer)
			},
			JsonName:   "ms.NtApi",
			IsDocument: false,
		})
		return widget.NewTableScrollPanel(table, header)
	})

	right := widget.NewPanel()
	right.AddChild(mylog.Check2Bool(layouts.Get(KernelTablesKind))()) // todo make a welcoming page
	splitPanel.AddChild(left)
	splitPanel.AddChild(right)

	// todo get and set inputted ctx,not clean it every time
	table.SelectionChangedCallback = func() {
		for i, n := range table.SelectedRows(false) {
			if i > 1 {
				break
			}
			switch n.Data.Name {
			case KernelTablesKind:
				right.RemoveAllChildren()
				paneler := mylog.Check2Bool(layouts.Get(KernelTablesKind))()
				right.AddChild(paneler)
				splitPanel.AddChild(right)

			case PackerKind:
				mylog.Todo("packer")
				// pe包似乎有个检测壳的，看看全不全
				// https://git.homegu.com/inc0d3/malware/blob/master/tools%2Funpacker%2Fthemida-2.x%2FThemida%20-%20Winlicense%20Ultra%20Unpacker%201.4.txt
				// 实现themida，vmprotect脱壳插件
				// todo dump overlay data
				// 拖放文件进调试器之后自动识别vmp tmd wl等壳并更新壳名称到调试器的主窗口标题
				// binee python 那个
				// more
				packer.CheckPacker("c:\\windows\\system32\\kernelbase.dll")
			default:
			}
		}
	}
	return splitPanel.AsPanel()
}
