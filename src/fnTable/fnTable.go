package fnTable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
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
	return container.NewAppTabs(
		container.NewTabItemWithIcon("cpu", nil, pageCpu.New().CanvasObject(window)),
		container.NewTabItemWithIcon("log", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("notes", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("breaks", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("memory", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("stack", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("sehList", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("script", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("symbols", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("source", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("xFrom", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("thead", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("handle", nil, widget.NewButton("", nil)),
		container.NewTabItemWithIcon("trance", nil, widget.NewButton("", nil)),
	)
}

func New() Interface { return &object{} }
