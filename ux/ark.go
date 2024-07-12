package ux

import (
	"fmt"

	"github.com/ddkwork/app/ms"
	"github.com/ddkwork/app/widget"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutArk() unison.Paneler {
	table, header := widget.NewTable(ms.NtApi{}, widget.TableContext[ms.NtApi]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[ms.NtApi]) (cells []widget.CellData) {
			KernelBase := fmt.Sprintf("%016X", node.Data.KernelBase)
			if node.Container() {
				KernelBase = node.Sum(KernelBase)
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
			sysCall := ms.NewSysCall(0)
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
}
