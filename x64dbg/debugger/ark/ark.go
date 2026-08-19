package ark

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/ddk"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type Ark struct {
	Name ArksType
}

type Manager struct {
	table *treetable.TreeTable[Ark]
}

func New() api.Interface {
	m := &Manager{}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[Ark]()
	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			for _, kind := range KernelTablesType.EnumTypes() {
				m.table.Root().AddChild(m.table.NewNode(Ark{Name: kind}))
			}
		},
		JsonName: "ark",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Update() error {
	return nil
}

func (m *Manager) Self() any {
	return m
}

func arkTodo() {
	ddk.DecodeTableByDll()
	ddk.MiGetPteAddress()
	ddk.DecodeTableByDll()
	ddk.DecodeTableByDisassembly()
	ddk.NtDeviceIoControlFile()
}

func NewNtApiTable() *treetable.TreeTable[ddk.NtApi] {
	t := treetable.NewTreeTable[ddk.NtApi]()
	t.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack: func() {
			sysCall := ddk.NewSysCall(int64(0))
			sysCall.KeServiceDescriptorTable = ddk.DecodeNtApi("C:\\Windows\\System32\\ntdll.dll")
			sysCall.KeServiceDescriptorTableShadow = ddk.DecodeNtApi("C:\\Windows\\System32\\win32u.dll")

			ntTableContainer := t.NewContainerNode("NtTable", ddk.NtApi{Name: "NtTable"})
			win32kTableContainer := t.NewContainerNode("Win32kTable", ddk.NtApi{Name: "Win32kTable"})

			for _, api := range sysCall.KeServiceDescriptorTable {
				api.Index++
				ntTableContainer.AddChild(t.NewNode(api))
			}
			for _, api := range sysCall.KeServiceDescriptorTableShadow {
				api.Index++
				win32kTableContainer.AddChild(t.NewNode(api))
			}
			t.Root().AddChild(ntTableContainer)
			t.Root().AddChild(win32kTableContainer)
		},
		JsonName: "NtApi",
	}
	return t
}
