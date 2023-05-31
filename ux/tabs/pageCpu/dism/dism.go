package dism

import (
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/ImmediateData"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	d := pageLog.New().CanvasObject(window)
	//todo d.setsize
	flowdataWindow := ImmediateData.New().CanvasObject(window)
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 1})
	panel.AddChild(d)
	panel.AddChild(flowdataWindow)
	panel.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		pref.Width = 220
		pref.Height = 220
		return min, pref, unison.MaxSize(max)
	})
	return panel
}

func New() Interface {
	return &object{}
}
