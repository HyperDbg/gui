package pageCpu

import (
	"github.com/ddkwork/GolandProjects/ui/widget"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/dism"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/dump"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/reg"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/stack"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/side"
)

type (
	object struct{}
)

func (o *object) CanvasObject() *unison.Dock {
	dismapp := widget.NewAppTabs("", "", nil, nil)
	regAPP := widget.NewAppTabs("", "", nil, nil)
	regAPP.SetSizer(func(_ unison.Size) (min, pref, max unison.Size) {
		pref.Width = 190
		pref.Height = 190
		return min, pref, pref
	})

	dumpAPP := widget.NewAppTabs("", "", nil, nil)
	stackAPP := widget.NewAppTabs("", "", nil, nil)

	dock := unison.NewDock()
	dock.DockTo(dismapp, nil, side.Left)

	dismapp.AddChild(dism.New().CanvasObject())
	dumpAPP.AddChild(dump.New().CanvasObject())
	stackAPP.AddChild(stack.New().CanvasObject())
	regAPP.AddChild(reg.New().CanvasObject())

	//第一页是个复杂的布局
	//反汇编在左边
	dock.DockTo(dumpAPP, unison.Ancestor[*unison.DockContainer](dismapp), side.Bottom) //dump基于反汇编下边
	dock.DockTo(stackAPP, unison.Ancestor[*unison.DockContainer](dumpAPP), side.Right) //堆栈基于dump右边
	dock.DockTo(regAPP, unison.Ancestor[*unison.DockContainer](dismapp), side.Right)   //寄存器基于反汇编右边

	return dock
}

func New() *object { return &object{} }
