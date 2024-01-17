package symbolLog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("log ...")
	return entry
}

func New() Interface {
	return &object{}
}
