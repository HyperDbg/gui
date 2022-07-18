package symbolChoose

import (
	"fyne.io/fyne/v2"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	//TODO implement me
	panic("implement me")
}

func New() Interface {
	return &object{}
}
