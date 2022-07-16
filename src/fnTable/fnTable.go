package fnTable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
)

type (
	Interface interface {
		canvasobjectapi.Interface
	}
	object struct{}
)

func New() Interface { return &object{} }

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	return container.NewAppTabs(
		container.NewTabItemWithIcon("cpu", theme.ConfirmIcon(), pageCpu.New().CanvasObject(window)),
		container.NewTabItemWithIcon("log", theme.ConfirmIcon(), widget.NewButton("log", nil)),
		container.NewTabItemWithIcon("notes", theme.ConfirmIcon(), widget.NewButton("notes", nil)),
		container.NewTabItemWithIcon("breaks", theme.ConfirmIcon(), widget.NewButton("breaks", nil)),
		container.NewTabItemWithIcon("memory", theme.ConfirmIcon(), widget.NewButton("memory", nil)),
		container.NewTabItemWithIcon("stack", theme.ConfirmIcon(), widget.NewButton("stack", nil)),
		container.NewTabItemWithIcon("sehList", theme.ConfirmIcon(), widget.NewButton("sehList", nil)),
		container.NewTabItemWithIcon("script", theme.ConfirmIcon(), widget.NewButton("script", nil)),
		container.NewTabItemWithIcon("symbols", theme.ConfirmIcon(), widget.NewButton("symbols", nil)),
		container.NewTabItemWithIcon("source", theme.ConfirmIcon(), widget.NewButton("source", nil)),
		container.NewTabItemWithIcon("xFrom", theme.ConfirmIcon(), widget.NewButton("xFrom", nil)),
		container.NewTabItemWithIcon("thead", theme.ConfirmIcon(), widget.NewButton("thead", nil)),
		container.NewTabItemWithIcon("handle", theme.ConfirmIcon(), widget.NewButton("handle", nil)),
		container.NewTabItemWithIcon("trance", theme.ConfirmIcon(), widget.NewButton("trance", nil)),
	)
}
