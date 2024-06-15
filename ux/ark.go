package ux

import (
	"fmt"

	"github.com/ddkwork/app/ms"
	"github.com/ddkwork/app/ms/driverTool/environment"
	"github.com/ddkwork/app/ms/hook/winver"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func arkTodo() {
	// todo merge plugin/symbol into here
	// hook random for aes key
	// add hardware info show and hook panel
	// vstart panel for quick start cracking tookBox ?
	// crypt tool merge here?
	ms.DecodeTableByDll()
	println(winver.WindowVersion())
	ms.MiGetPteAddress()
	ms.DecodeTableByDll()
	ms.DecodeTableByDisassembly()
	ms.NtDeviceIoControlFile()
	// IopXxxControlFile()
	widget.NewExplorer(nil, ".")
	environment.New()

	// taskexplorer todo call here
	mylog.Todo("implement registry editor")
}

// LayoutArk ark panel is show
// 1 nt and win32k table,finished
// 2 file explorer,finished
// 3 registry editor
// etc.
func LayoutArk(parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(ms.NtApi{}, widget.TableContext[ms.NtApi]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[ms.NtApi]) (cells []widget.CellData) {
			KernelBase := fmt.Sprintf("%016X", node.Data.KernelBase)
			if node.Container() {
				KernelBase = node.Type
			}
			return []widget.CellData{
				{Text: KernelBase},
				{Text: fmt.Sprintf("%016X", node.Data.ArgValue)},
				{Text: node.Data.Name},
				{Text: fmt.Sprintf("%04d / %016X", node.Data.Index, node.Data.Index)},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[ms.NtApi]) {
			sysCall := ms.NewSysCall(0)
			sysCall.KeServiceDescriptorTable = ms.DecodeNtApi("C:\\Windows\\System32\\ntdll.dll")
			sysCall.KeServiceDescriptorTableShadow = ms.DecodeNtApi("C:\\Windows\\System32\\win32u.dll")
			NtTableContainer := widget.NewContainerNode("NtTable", ms.NtApi{})
			Win32kTableContainer := widget.NewContainerNode("Win32kTable", ms.NtApi{})
			for _, api := range sysCall.KeServiceDescriptorTable {
				NtTableContainer.AddChildByData(api)
			}
			for _, api := range sysCall.KeServiceDescriptorTableShadow {
				Win32kTableContainer.AddChildByData(api)
			}
			root.AddChild(NtTableContainer)
			root.AddChild(Win32kTableContainer)
		},
		JsonName:   "ms.NtApi",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}
