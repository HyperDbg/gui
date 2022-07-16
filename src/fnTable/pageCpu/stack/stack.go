package stack

import (
	"fyne.io/fyne/v2"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/librarygo/src/fynelib/myTable"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
)

type (
	Interface interface{ canvasobjectapi.Interface }
	object    struct{}
)

func New() Interface {
	return &object{}
}

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	d := newDisassemblyObject()
	d.SetLines(nil)
	list, err := myTable.NewDVList(d, myTable.WithSelectionHandler(func(id int, selected bool) {
		l := d.Lines()[id]
		mylog.Struct(l)
	}))
	if !mycheck.Error(err) {
		return nil
	}
	return list
}
