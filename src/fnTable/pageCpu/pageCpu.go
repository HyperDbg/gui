package pageCpu

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/dism"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/dump"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/reg"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu/stack"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	d := dism.New().CanvasObject(window) //table
	r := reg.New().CanvasObject(window)
	dump := dump.New().CanvasObject(window) //table
	s := stack.New().CanvasObject(window)   //table
	hSplit := container.NewHSplit(d, r)
	hSplit.Offset = 0.7
	newHSplit := container.NewHSplit(dump, s)
	newHSplit.Offset = 0.5
	cpu := container.NewAdaptiveGrid(1, hSplit, newHSplit)
	//cpu := container.NewVSplit(hSplit, newHSplit)
	//cpu.Offset = 0.7
	return cpu
}

func New() Interface { return &object{} }
