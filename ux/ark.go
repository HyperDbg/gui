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

type Ark struct{}

type (
	NtTable     struct{}
	Win32kTable struct{}
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
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ExceptionHandlingRoutines},
				{Text: node.Data.Label},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[ms.NtApi]) {
			for i := range 100 {
				ts := ms.NtApi{
					Address:                   0x00000002429BF7B8 + i,
					ExceptionHandlingRoutines: "xxx",
					Label:                     "www",
					Notes:                     "this is notes",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "ms.NtApi",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}
