package imm

import (
	"strconv"

	"gioui.org/layout"
	"gioui.org/widget"

	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/menu"
	"github.com/ddkwork/x64dbg/debugger/api"
)

type Imm struct {
	*layout.List
	rows           []ImmData
	clickables     []*widget.Clickable
	OnFollowInDump func(address uint64)
}

type ImmData struct {
	Reg     string
	Address uint64
	Mem     string
}

type FastCall struct {
	Register string
	Address  int
	ImmData  string
}

func New() api.Interface {
	return &Imm{List: &layout.List{Axis: layout.Vertical}}
}

func (m *Imm) UpdateRows(rows []ImmData) {
	m.rows = rows
	m.clickables = make([]*widget.Clickable, len(rows))
	for i := range m.clickables {
		m.clickables[i] = &widget.Clickable{}
	}
}

func (m *Imm) Update() error {
	return nil
}

func (m *Imm) Clear() {
	m.rows = nil
	m.clickables = nil
}

func (m *Imm) Self() any {
	return m
}

func (m *Imm) Layout() layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		if m.rows == nil || len(m.rows) == 0 {
			return layout.Dimensions{}
		}
		return m.List.Layout(gtx, len(m.rows), func(gtx layout.Context, i int) layout.Dimensions {
			return m.clickables[i].Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Horizontal,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:      layout.Horizontal,
							Spacing:   5,
							Alignment: 0,
							WeightSum: 0,
						}.Layout(gtx,
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return wdk.LabelS(gtx, "Reg:"+m.rows[i].Reg)
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								return wdk.LabelS(gtx, "  Addr:"+strconv.FormatUint(m.rows[i].Address, 16))
							}),
							layout.Rigid(func(gtx layout.Context) layout.Dimensions {
								if m.rows[i].Mem != "" {
									return wdk.LabelS(gtx, "  Mem:"+m.rows[i].Mem)
								}
								return layout.Dimensions{}
							}),
						)
					}))
			})
		})
	}
}

func (m *Imm) GetContextMenuItems() []*menu.MenuItem {
	var items []*menu.MenuItem
	for _, row := range m.rows {
		if row.Address != 0 {
			addr := row.Address
			items = append(items, &menu.MenuItem{
				Label: "Follow " + row.Reg + " in Dump (0x" + strconv.FormatUint(addr, 16) + ")",
				Do: func() {
					if m.OnFollowInDump != nil {
						m.OnFollowInDump(addr)
					}
				},
			})
		}
	}
	return items
}

func (m *Imm) GetSelectedAddress(index int) uint64 {
	if index >= 0 && index < len(m.rows) {
		return m.rows[index].Address
	}
	return 0
}
