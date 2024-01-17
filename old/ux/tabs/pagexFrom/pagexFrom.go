package pagexFrom

import (
	"fyne.io/fyne/v2"
)

type (
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	//TODO implement me
	panic("implement me")
}

func New() *object { return &object{} }
