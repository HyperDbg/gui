package dism

import (
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/ImmediateData"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
)

type (
	object struct{}
)

func (o *object) CanvasObject() *unison.Panel {
	d := pageLog.New().CanvasObject()
	d.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		pref.Width = 90
		pref.Height = 290
		return min, pref, unison.MaxSize(max)
	})
	flowdataWindow := ImmediateData.New().CanvasObject()
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 1})
	panel.AddChild(d)
	panel.AddChild(flowdataWindow)
	return panel
}

func New() *object {
	return &object{}
}
