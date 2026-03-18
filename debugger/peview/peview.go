package peview

import (
	"iter"

	"gioui.org/layout"
	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/ddk/xed"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/ux/widget/treetable"
)

type PeView struct {
	Name    string
	Offset  uint32
	Address uint32
	Size    uint32
	Is64Bit bool
}

type Manager struct {
	table         *treetable.TreeTable[PeView]
	TargetExePath string
	populateFunc  func()
}

func New() api.Interface {
	m := &Manager{}
	m.initTable()
	return m
}

func (m *Manager) initTable() {
	m.table = treetable.NewTreeTable[PeView]()
	m.populateFunc = func() {
		if m.TargetExePath == "" {
			return
		}
		pe := xed.ParserPe(m.TargetExePath)

		sectionNode := m.table.NewContainerNode("Section", PeView{Name: "Section"})
		for _, section := range pe.Sections {
			sectionNode.AddChild(m.table.NewNode(PeView{
				Name:    section.String(),
				Offset:  section.Header.PointerToRawData,
				Address: section.Header.VirtualAddress,
				Size:    section.Header.VirtualSize,
				Is64Bit: pe.Is64,
			}))
		}
		m.table.Root().AddChild(sectionNode)

		importNode := m.table.NewContainerNode("Import", PeView{Name: "Import"})
		for _, v := range pe.Imports {
			importNode.AddChild(m.table.NewNode(PeView{
				Name:    v.Name,
				Offset:  v.Offset,
				Size:    0,
				Is64Bit: pe.Is64,
			}))
		}
		m.table.Root().AddChild(importNode)

		exportNode := m.table.NewContainerNode("Export", PeView{Name: "Export"})
		for _, function := range pe.Export.Functions {
			exportNode.AddChild(m.table.NewNode(PeView{
				Name:    function.Name,
				Offset:  function.FunctionRVA,
				Address: 0,
				Size:    0,
				Is64Bit: pe.Is64,
			}))
		}
		m.table.Root().AddChild(exportNode)
	}

	m.table.AirTable.TableContext = treetable.TableContext{
		CustomContextMenuItems: func(gtx layout.Context, n *treetable.Node) iter.Seq[*menu.MenuItem] {
			return func(yield func(*menu.MenuItem) bool) {}
		},
		RowSelectedCallback:    func() {},
		RowDoubleClickCallback: func() {},
		SetRootRowsCallBack:    m.populateFunc,
		JsonName:               "PeView",
	}
}

func (m *Manager) Layout() layout.Widget {
	return m.table.AirTable.Layout
}

func (m *Manager) Clear() {
	m.table.Root().SetChildren(nil)
}

func (m *Manager) Update() error {
	m.populateFunc()
	m.table.AirTable.Refresh()
	return nil
}

func (m *Manager) Self() any {
	return m
}

func (m *Manager) SetTargetPath(path string) {
	m.TargetExePath = path
}
