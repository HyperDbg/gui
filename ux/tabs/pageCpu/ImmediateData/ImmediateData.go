package ImmediateData

import (
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	panel := pageLog.New().CanvasObject(window)
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
