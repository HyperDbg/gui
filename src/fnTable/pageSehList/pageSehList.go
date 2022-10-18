package pageSehList

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/golibrary/src/fynelib/myTable"
	"github.com/fpabl0/sparky-go/swid"
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
	list, err := myTable.NewTable(d, myTable.WithSelectionHandler(func(id int, selected bool) {
		l := d.Lines()[id]
		mylog.Struct(l)
	}))
	if !mylog.Error(err) {
		return nil
	}
	filter := swid.NewTextFormField("filter", "")
	return container.NewBorder(list, filter, nil, nil)
}
