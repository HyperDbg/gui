package symbolLog

import (
	"fyne.io/fyne/v2/widget"
)
import (
	"fyne.io/fyne/v2"
)

type (
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	entry := widget.NewMultiLineEntry()
	entry.SetPlaceHolder("log ...")
	return entry
}

func New() *object {
	return &object{}
}
