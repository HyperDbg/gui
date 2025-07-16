package ui

import (
	"iter"
	"strings"

	"gioui.org/layout"
	"github.com/ddkwork/ddk/xed"

	"github.com/ddkwork/ux"
)

type PeView struct {
	Name    string
	Offset  uint32
	Address uint32
	Size    uint32
	Is64Bit bool
	// overlay *pe.Overlay todo
}

func NewPeView() ux.Widget {
	t := ux.NewTreeTable(PeView{})
	t.TableContext = ux.TableContext[PeView]{
		CustomContextMenuItems: func(gtx layout.Context, n *ux.Node[PeView]) iter.Seq[ux.ContextMenuItem] {
			return func(yield func(ux.ContextMenuItem) bool) {
			}
		},
		MarshalRowCells: func(n *ux.Node[PeView]) (cells []ux.CellData) {
			if n.Container() {
				n.Data.Name = n.SumChildren()
			}
			return ux.MarshalRow(n.Data, func(key string, field any) (value string) {
				return ""
			})
		},
		UnmarshalRowCells: func(n *ux.Node[PeView], rows []ux.CellData) PeView {
			if n.Container() {
				n.Data.Name = n.SumChildren()
			}
			// return []ux.CellData{
			//	{Text: n.Data.Name},
			//	{Text: fmt.Sprintf("%08X", n.Data.Offset)},
			//	{Text: fmt.Sprintf("%08X", n.Data.Size)},
			//	{Text: fmt.Sprintf("%016X", n.Data.Address)},
			//	{Text: fmt.Sprintf("%t", n.Data.Is64Bit)},
			// }
			return ux.UnmarshalRow[PeView](rows, func(key, value string) (field any) {
				return nil
			})
		},
		RowSelectedCallback: func() {
		},
		RowDoubleClickCallback: func() {
		},
		SetRootRowsCallBack: func() {
			const (
				Section = "Section"
				Import  = "Import"
				Export  = "Export"
			)
			if TargetExePath == "" {
				return
			}
			pe := xed.ParserPe(TargetExePath)
			containerNodes := ux.NewContainerNodes[PeView]([]string{Section, Import, Export})
			for _, node := range containerNodes {
				typ := node.Type
				typ = strings.TrimSuffix(typ, ux.ContainerKeyPostfix)
				switch typ {
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
				t.Root.AddChild(node)
			}
		},
		JsonName:   "PeView",
		IsDocument: false,
	}
	return t
}
