package pageCpu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/ImmediateData"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/dism"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/dump"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/reg"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/stack"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	split0 := container.NewVSplit(dism.New().CanvasObject(window), ImmediateData.New().CanvasObject(window))
	split0.SetOffset(0.8)
	hSplit := container.NewHSplit(split0, reg.New().CanvasObject(window))
	hSplit.Offset = 0.7

	split1 := container.NewHSplit(dump.New().CanvasObject(window), stack.New().CanvasObject(window))
	split1.Offset = 0.5

	cpu := container.NewVSplit(hSplit, split1)
	cpu.Offset = 0.7
	return cpu
}

func New() Interface { return &object{} }
