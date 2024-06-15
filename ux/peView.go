package ux

import (
	"fmt"
	"ms/xed"
	"strings"

	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
)

type PeView struct {
	Name    string
	Offset  uint32
	Address uint32
	Size    uint32
	Is64Bit bool
	// overlay *pe.Overlay todo
}

func LayoutPeView(fileName string, parent unison.Paneler) unison.Paneler {
	table, header := widget.NewTable(PeView{}, widget.TableContext[PeView]{
		ContextMenuItems: nil,
		MarshalRow: func(node *widget.Node[PeView]) (cells []widget.CellData) {
			if node.Container() {
				node.Data.Name = node.Type
			}
			return []widget.CellData{
				{Text: node.Data.Name},
				{Text: fmt.Sprintf("%08X", node.Data.Offset)},
				{Text: fmt.Sprintf("%08X", node.Data.Size)},
				{Text: fmt.Sprintf("%016X", node.Data.Address)},
				{Text: fmt.Sprintf("%t", node.Data.Is64Bit)},
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
			if fileName == "" {
				return
			}
			pe := xed.ParserPe(fileName)
			containerNodes := widget.NewContainerNodes[PeView]([]string{Section, Import, Export})
			for _, node := range containerNodes {
				t := node.Type
				t = strings.TrimSuffix(t, widget.ContainerKeyPostfix)
				switch t {
				case Section:
					for _, section := range pe.Sections {
						node.AddChildByData(PeView{
							Name:    section.String(),
							Offset:  section.Header.PointerToRawData,
							Address: section.Header.VirtualAddress,
							Size:    section.Header.VirtualSize,
							Is64Bit: pe.Is64,
						})
					}
				case Import:
					for _, v := range pe.Imports {
						node.AddChildByData(PeView{
							Name:   v.Name,
							Offset: v.Offset,
							// Address: v.,
							Size:    0,
							Is64Bit: pe.Is64,
						})
					}
				case Export:
					for _, function := range pe.Export.Functions {
						node.AddChildByData(PeView{
							Name:    function.Name,
							Offset:  function.FunctionRVA,
							Address: 0,
							Size:    0,
							Is64Bit: pe.Is64,
						})
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
