package ark

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
)

type Ark struct {
	Name string
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
			kinds := []string{"SSDT", "IDT", "GDT", "Driver", "Process", "Thread", "Module"}
			for _, kind := range kinds {
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
