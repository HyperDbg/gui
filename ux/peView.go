package ux

import (
	"fmt"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
	"ms/xed"
)

type PeView struct {
	Name    string
	Offset  uint64
	Address uint64
	Size    uint64
	Is64Bit bool
	//overlay *pe.Overlay todo
}

func layoutPeView(name string, parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(PeView{}, widget.TableContext[PeView]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[PeView]) (cells []widget.CellData) {
			return []widget.CellData{
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: node.Data.ExceptionHandlingRoutines},
				{Text: node.Data.Label},
				{Text: node.Data.Notes},
			}
		},
		UnmarshalRow:             nil,
		SelectionChangedCallback: nil,
		SetRootRowsCallBack: func(root *widget.Node[PeView]) {
			const (
				Section = "Section"
				Import  = "Import"
				Export  = "Export"
			)
			pe := xed.ParserPe(name)
			containerNodes := widget.NewContainerNodes[PeView]([]string{Section, Import, Export})
			for _, node := range containerNodes {
				switch node.Type {
				case Section:
					for i, section := range pe.Sections {
						node.AddChildByData(PeView{})
					}
				case Import:
					for i, i2 := range pe.Imports {
						node.AddChildByData(PeView{})
					}
				case Export:
					for i, function := range pe.Export.Functions {
						node.AddChildByData(PeView{})
					}
				}
				root.AddChild(node)
			}
		},
		JsonName:   "PeView",
		IsDocument: false,
	})
	return widget.NewTableScrollPanel(parent, table, header)
}
