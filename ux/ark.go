package ux

import (
	"fmt"
	"github.com/ddkwork/app/ms"
	"github.com/ddkwork/app/ms/driverTool/environment"
	"github.com/ddkwork/app/ms/hook/winver"
	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

type Ark struct{}

type (
	NtTable     struct{}
	Win32kTable struct{}
)

// LayoutArk ark panel is show
// 1 nt and win32k table,finished
// 2 file explorer,finished
// 3 registry editor
// etc.
func LayoutArk(parent unison.Paneler) unison.Paneler {
	//todo merge branch/gui/plugin/symbol into here
	ms.DecodeTableByDll()
	println(winver.WindowVersion())
	ms.MiGetPteAddress()
	ms.DecodeTableByDll()
	ms.DecodeTableByDisassembly()
	ms.NtDeviceIoControlFile()
	//IopXxxControlFile()
	widget.NewExplorer(parent, ".")
	//taskexplorer todo call here
	environment.New()
	mylog.Todo("implement registry editor")

	table, header := widget.NewTable(Seh{}, widget.TableContext[Seh]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[Seh]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ExceptionHandlingRoutines},
				{Text: node.Data.Label},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[Seh]) {
			for i := range 100 {
				ts := Seh{
					Address:                   0x00000002429BF7B8 + i,
					ExceptionHandlingRoutines: "xxx",
					Label:                     "www",
					Notes:                     "this is notes",
				}
				root.AddChildByData(ts)
			}
		},
		JsonName:   "seh",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}
