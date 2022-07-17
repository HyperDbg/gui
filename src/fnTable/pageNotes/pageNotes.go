package pageNotes

import (
	"fyne.io/fyne/v2"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/librarygo/src/fynelib/notes"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		//Fn() (ok bool)
	}
	object struct{}
)

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	list := &notes.Notelist{Pref: fyne.CurrentApp().Preferences()}
	list.Load()
	notesUI := &notes.Ui{Notes: list}

	//window.SetContent(notesUI.LoadUI())
	notesUI.RegisterKeys(window)
	return notesUI.LoadUI()
}

func New() Interface { return &object{} }
