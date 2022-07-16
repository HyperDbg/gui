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
		container.NewTabItemWithIcon("log", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("notes", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("breaks", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("memory", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("stack", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("sehList", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("script", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("symbols", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("source", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("xFrom", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("thead", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("handle", theme.ConfirmIcon(), widget.NewButton("", nil)),
		container.NewTabItemWithIcon("trance", theme.ConfirmIcon(), widget.NewButton("", nil)),
	)
}
