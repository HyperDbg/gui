package ark

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ddk"
	"github.com/ddkwork/ddk/winver"
	"github.com/ddkwork/ddk/worker/environment"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ddkwork/ux"
)

// func main() {
//	app.Run("ark tool", func(w *unison.Window) {
//		content := w.Content()
//		panel := ux.NewPanel()
//		panel.AddChild(Layout())
//		scrollPanelFill := ux.NewScrollPanelFill(panel)
//		content.AddChild(scrollPanelFill)
//	})
// }

func arkTodo() {
	ddk.DecodeTableByDll()
	println(winver.WindowVersion())
	ddk.MiGetPteAddress()
	ddk.DecodeTableByDll()
	ddk.DecodeTableByDisassembly()
	ddk.NtDeviceIoControlFile()
	// IopXxxControlFile()
	// ux.NewExplorer(".")
	environment.New()
}

func New() ux.Widget {
	type ark struct{ Name ArksType }
	t := ux.NewTreeTable(ark{})
	t.TableContext = ux.TableContext[ark]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[ark]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[ark]) (cells []ux.CellData) {
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[ark], rows []ux.CellData) ark {
			name := n.Data.Name.String()
			if n.Container() {
				name = n.SumChildren()
			}
			name = name // todo
			// return []ux.CellData{{Text: name}}
			return ux.UnmarshalRow[ark](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			for _, kind := range KernelTablesType.EnumTypes() {
				t.Root.AddChildByData(ark{kind})
			}
		},
		JsonName:   "ark",
		IsDocument: false,
	}

	// splitPanel := ux.NewPanel()//todo
	// ux.SetScrollLayout(splitPanel, 2)
	//
	// left := ux.NewTableScrollPanel(table, header)
	layouts := new(safemap.M[ArksType, func() ux.Widget])
	layouts.Set(KernelTablesType, func() ux.Widget {
		t := ux.NewTreeTable(ddk.NtApi{})
		t.TableContext = ux.TableContext[ddk.NtApi]{
			CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[ddk.NtApi]) iter.Seq[ux.ContextMenuItem] {
				return func(yield func(ux.ContextMenuItem) bool) {
				}
			},
			MarshalRowCells: func(n *ux.Node[ddk.NtApi]) (cells []ux.CellData) {
				KernelBase := stream.FormatIntegerHex(n.Data.KernelBase)
				if n.Container() {
					KernelBase = n.SumChildren()
				}
				KernelBase = KernelBase // todo
				// return []ux.CellData{
				//	{Text: KernelBase},
				//	{Text: stream.FormatIntegerHex(n.Data.ArgValue)},
				//	{Text: n.Data.Name},
				//	{Text: stream.FormatInteger(n.Data.Index)},
				// }
				return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
					return ""
				})
			},
			UnmarshalRowCells: func(n *ux.Node[ddk.NtApi], rows []ux.CellData) ddk.NtApi {
				return ux.UnmarshalRow[ddk.NtApi](rows, func(key, value string) (field any) {
					return nil
				})
			},
			RowSelectedCallback: func() {
				switch t.SelectedNode.Data.Name {
				// case KernelTablesType:
				// right.RemoveAllChildren()
				// paneler := (layouts.GetMust(KernelTablesType))()
				// right.AddChild(paneler)
				// splitPanel.AddChild(right)

				// case PackerType:
				// mylog.Todo("packer")
				// pe包似乎有个检测壳的，看看全不全
				// https://github.com/inc0d3/malware/blob/master/tools%2Funpacker%2Fthemida-2.x%2FThemida%20-%20Winlicense%20Ultra%20Unpacker%201.4.txt
				// 实现themida，vmprotect脱壳插件
				// todo dump overlay data
				// 拖放文件进调试器之后自动识别vmp tmd wl等壳并更新壳名称到调试器的主窗口标题
				// binee python 那个
				// more
				// packer.CheckPacker("c:\\windows\\system32\\kernelbase.dll")
				default:
				}
			},
			RowDoubleClickCallback: func() {
			},
			SetRootRowsCallBack: func() {
				//sysCall := ddk.NewSysCall(int64(sdk.GetKernelBase()))
				sysCall := ddk.NewSysCall(int64(0)) //todo
				sysCall.KeServiceDescriptorTable = ddk.DecodeNtApi("C:\\Windows\\System32\\ntdll.dll")
				sysCall.KeServiceDescriptorTableShadow = ddk.DecodeNtApi("C:\\Windows\\System32\\win32u.dll")
				NtTableContainer := ux.NewContainerNode("NtTable", ddk.NtApi{})
				Win32kTableContainer := ux.NewContainerNode("Win32kTable", ddk.NtApi{})
				for _, api := range sysCall.KeServiceDescriptorTable {
					api.Index++
					NtTableContainer.AddChildByData(api)
				}
				for _, api := range sysCall.KeServiceDescriptorTableShadow {
					api.Index++
					Win32kTableContainer.AddChildByData(api)
				}
				t.Root.AddChild(NtTableContainer)
				t.Root.AddChild(Win32kTableContainer)
			},
			JsonName:   "NtApi",
			IsDocument: false,
		}
		return t
	})

	// right := ux.NewPanel()
	// right.AddChild((layouts.GetMust(KernelTablesType))()) // todo make a welcoming page
	// splitPanel.AddChild(left)
	// splitPanel.AddChild(right)

	return t
}
