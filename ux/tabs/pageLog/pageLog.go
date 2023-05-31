package pageLog

import (
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/google/uuid"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		unison.TableRowData[*object]
	}
	object struct {
		table    *unison.Table[*object]
		parent   *object
		id       uuid.UUID
		children []*object
		//checkbox     *unison.CheckBox
		container bool
		open      bool
		//doubleHeight bool
		line
	}
)

func New() Interface                             { return &object{} }
func (o *object) IsOpen() bool                   { return o.open }
func (o *object) SetOpen(open bool)              { o.open = open }
func (o *object) UUID() uuid.UUID                { return o.id }
func (o *object) Parent() *object                { return o.parent }
func (o *object) SetParent(parent *object)       { o.parent = parent }
func (o *object) CanHaveChildren() bool          { return o.container }
func (o *object) Children() []*object            { return o.children }
func (o *object) SetChildren(children []*object) { o.children = children }

func (o *object) CloneForTarget(target unison.Paneler, newParent *object) *object {
	table, ok := target.(*unison.Table[*object])
	if !ok {
		jot.Fatal(1, "invalid target")
	}
	clone := *o
	clone.table = table
	clone.parent = newParent
	clone.id = uuid.New()
	return &clone
}

func (o *object) CellDataForSort(col int) string {
	switch col {
	case 0:
		return o.address
	case 1:
		return o.data
	case 2:
		return o.dism
	case 3:
		return o.notes
	default:
		return ""
	}
}

func (o *object) ColumnCell(row, col int, foreground, background unison.Ink, selected, indirectlySelected, focused bool) unison.Paneler {
	switch col {
	case 0:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.address, foreground, unison.LabelFont, width)
		//if o.doubleHeight {
		//	addWrappedText(wrapper, o.address, foreground,
		//		unison.LabelFont.Face().Font(unison.LabelFont.Size()-1), width)
		//}
		wrapper.UpdateTooltipCallback = func(where unison.Point, suggestedAvoidInRoot unison.Rect) unison.Rect {
			wrapper.Tooltip = unison.NewTooltipWithText("A tooltip for the cell")
			return wrapper.RectToRoot(wrapper.ContentRect(true))
		}
		return wrapper
	case 1:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.data, foreground, unison.LabelFont, width)
		//if o.doubleHeight {
		//	addWrappedText(wrapper, o.data, foreground,
		//		unison.LabelFont.Face().Font(unison.LabelFont.Size()-1), width)
		//}
		wrapper.UpdateTooltipCallback = func(where unison.Point, suggestedAvoidInRoot unison.Rect) unison.Rect {
			wrapper.Tooltip = unison.NewTooltipWithText("A tooltip for the cell")
			return wrapper.RectToRoot(wrapper.ContentRect(true))
		}
		return wrapper
	case 2:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.dism, foreground, unison.LabelFont, width)
		return wrapper
	case 3:
		wrapper := unison.NewPanel()
		wrapper.SetLayout(&unison.FlexLayout{Columns: 1})
		width := o.table.CellWidth(row, col)
		addWrappedText(wrapper, o.notes, foreground, unison.LabelFont, width)
		return wrapper
	default:
		jot.Errorf("column index out of range (0-2): %o", col)
		return unison.NewLabel()
	}
}

func addWrappedText(parent *unison.Panel, text string, ink unison.Ink, font unison.Font, width float32) {
	decoration := &unison.TextDecoration{Font: font}
	var lines []*unison.Text
	if width > 0 {
		lines = unison.NewTextWrappedLines(text, decoration, width)
	} else {
		lines = unison.NewTextLines(text, decoration)
	}
	for _, line := range lines {
		label := unison.NewLabel()
		label.Text = line.String()
		label.Font = font
		label.LabelTheme.OnBackgroundInk = ink
		parent.AddChild(label)
	}
}
