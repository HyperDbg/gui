package pageCpu

import (
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/dism"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/dump"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/reg"
	"github.com/ddkwork/hyperdbgui/ux/tabs/pageCpu/stack"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func (o *object) CanvasObject(window *unison.Window) *unison.Panel {
	//没有拆分布局的锅
	d := dism.New().CanvasObject(window)
	reg := reg.New().CanvasObject(window)
	dump := dump.New().CanvasObject(window)
	stack := stack.New().CanvasObject(window)

	panel12 := unison.NewPanel()
	panel12.SetLayout(&unison.FlexLayout{Columns: 2})
	panel12.AddChild(d)
	//panel12.AddChild(reg)
	reg = reg

	panel34 := unison.NewPanel()
	panel34.SetLayout(&unison.FlexLayout{Columns: 2})
	panel34.AddChild(dump)
	panel34.AddChild(stack)

	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 1})
	panel.AddChild(panel12)
	panel.AddChild(panel34)
	return panel
}

//func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
//	split0 := container.NewVSplit(dism.New().CanvasObject(window), ImmediateData.New().CanvasObject(window))
//	split0.SetOffset(0.8)
//	hSplit := container.NewHSplit(split0, reg.New().CanvasObject(window))
//	hSplit.Offset = 0.7
//
//	split1 := container.NewHSplit(dump.New().CanvasObject(window), stack.New().CanvasObject(window))
//	split1.Offset = 0.5
//
//	cpu := container.NewVSplit(hSplit, split1)
//	cpu.Offset = 0.7
//	return cpu
//}

func New() Interface { return &object{} }
