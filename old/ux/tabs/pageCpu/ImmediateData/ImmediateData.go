package ImmediateData

import (
	"fmt"
	"github.com/richardwilkes/toolbox/log/jot"

	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/behavior"
	"strings"
)

type (
	object struct{}
)

func New() *object {
	return &object{}
}
func (o *object) CanvasObject() *unison.Panel {
	panel := createListPanel()
	//panel.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
	//	pref.Width = 220
	//	pref.Height = 100
	//	return min, pref, unison.MaxSize(max)
	//})
	return panel
}

func createListPanel() *unison.Panel {
	lst := unison.NewList[string]()
	lst.Append(
		"Three with some long text to make it interesting",
		"Four",
		"Five",
	)
	lst.NewSelectionCallback = func() {
		var buffer strings.Builder
		buffer.WriteString("Selection changed in the list. Now:")
		index := -1
		first := true
		for {
			index = lst.Selection.NextSet(index + 1)
			if index == -1 {
				break
			}
			if first {
				first = false
			} else {
				buffer.WriteString(",")
			}
			fmt.Fprintf(&buffer, " %d", index)
		}
		jot.Info(buffer.String())
	}
	lst.DoubleClickCallback = func() {
		jot.Info("Double-clicked on the list")
	}
	_, prefSize, _ := lst.Sizes(unison.Size{})
	lst.SetFrameRect(unison.Rect{Size: prefSize})
	scroller := unison.NewScrollPanel()
	scroller.SetBorder(unison.NewLineBorder(unison.ControlEdgeColor, 0, unison.NewUniformInsets(1), false))
	scroller.SetContent(lst, behavior.Fill, behavior.Fill)
	scroller.SetLayoutData(&unison.FlexLayoutData{
		HSpan:  1,
		VSpan:  1,
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	return scroller.AsPanel()
}
