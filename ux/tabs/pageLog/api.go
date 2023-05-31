package pageLog

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/richardwilkes/unison"
)

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 1})

	// Create the table
	table := unison.NewTable[*object](&unison.SimpleTableModel[*object]{})
	table.HierarchyColumnID = 1
	table.Columns = make([]unison.ColumnInfo, 4)
	for i := range table.Columns {
		table.Columns[i].ID = i
		table.Columns[i].Minimum = 20
		table.Columns[i].Maximum = 10000
	}
	_, checkColSize, _ := unison.NewCheckBox().Sizes(unison.Size{})
	table.Columns[0].Minimum = checkColSize.Width
	table.Columns[0].Maximum = checkColSize.Width

	d := newDisassemblyObject()
	buf := []byte{
		0x55,             // pushq %rbp
		0x48, 0x89, 0xe5, // movq %rsp, %rbp
		0x48, 0x83, 0xe4, 0xe0, // andq $-32, %rsp
		0x48, 0x81, 0xec, 0xa0, 0x00, 0x00, 0x00, // subq $160, %rsp
		0xc5, 0xf9, 0xef, 0xc0, // vpxor   %xmm0, %xmm0, %xmm0
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x60, // vmovdqa %ymm0, 96(%rsp)
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x40, // vmovdqa %ymm0, 64(%rsp)
		0xc5, 0xfd, 0x7f, 0x44, 0x24, 0x20, // vmovdqa %ymm0, 32(%rsp)
		0xc5, 0xfd, 0x7f, 0x04, 0x24, // vmovdqa %ymm0, (%rsp)
		0x4c, 0x8d, 0x05, 0x00, 0x00, 0x00, 0x00, // leaq (%rip), %r8
	}
	d.Decode(buf)

	rows := make([]*object, len(d.lines))
	for i := range rows {
		l := d.lines[i]
		l.notes = fmt.Sprint(i) + "   this is a comment"
		row := &object{
			table:     table,
			parent:    nil,
			id:        uuid.New(),
			children:  nil,
			container: false,
			open:      false,
			line:      l,
		}
		if i%10 == 3 {
			row.children = make([]*object, 5)
			for j := range row.children {
				child := &object{
					table:     table,
					parent:    row,
					id:        uuid.New(),
					children:  nil,
					container: false,
					open:      false,
					line:      l,
				}
				row.children[j] = child
				if j < 2 {
					child.children = make([]*object, 2)
					for k := range child.children {
						child.children[k] = &object{
							table:     table,
							parent:    child,
							id:        uuid.New(),
							children:  nil,
							container: false,
							open:      false,
							line:      l,
						}
					}
				}
			}
		}
		rows[i] = row
	}
	table.SetRootRows(rows)
	table.SizeColumnsToFit(true)
	table.InstallDragSupport(nil, "object", "Row", "Rows")
	unison.InstallDropSupport[*object, any](table, "object",
		func(from, to *unison.Table[*object]) bool { return from == to }, nil, nil)

	header := unison.NewTableHeader[*object](table,
		unison.NewTableColumnHeader[*object]("address", "地址"),
		unison.NewTableColumnHeader[*object]("data", "数据"),
		unison.NewTableColumnHeader[*object]("dism", "反汇编"),
		unison.NewTableColumnHeader[*object]("notes", "笔记"),
	)
	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
	})
	panel.AddChild(header)

	// Create a scroll panel and place a table panel inside it
	scrollArea := unison.NewScrollPanel()
	scrollArea.SetContent(table, unison.FillBehavior, unison.FillBehavior)
	scrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: unison.FillAlignment,
		VAlign: unison.FillAlignment,
		HGrab:  true,
		VGrab:  true,
	})
	panel.AddChild(scrollArea)
	return panel
}
