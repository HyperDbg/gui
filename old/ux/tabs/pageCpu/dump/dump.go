package dump

import (
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageLog"
)

type (
	object struct{}
)

func (o *object) CanvasObject() *unison.Panel {
	panel := pageLog.New().CanvasObject()
	panel.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		pref.Width = 90
		pref.Height = 90
		return min, pref, unison.MaxSize(max)
	})
	return panel
}

func New() *object { return &object{} }
