package pageSymbol

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSymbol/exportApi"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSymbol/module"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSymbol/symbolChoose"
	"github.com/ddkwork/hyperdbgui/fnTable/pageSymbol/symbolLog"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	hSplit := container.NewHSplit(module.New().CanvasObject(window), exportApi.New().CanvasObject(window))
	hSplit1 := container.NewHSplit(symbolChoose.New().CanvasObject(window), symbolLog.New().CanvasObject(window))
	split := container.NewVSplit(hSplit, hSplit1)
	split.SetOffset(0.9)
	return split
}

func New() Interface {
	return &object{}
}
