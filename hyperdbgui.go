package hyperdbgui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
	"github.com/ddkwork/hyperdbgui/src/meau"
	"github.com/ddkwork/hyperdbgui/src/toolbar"
	"github.com/ddkwork/librarygo/src/fynelib/canvasobjectapi"
	"github.com/fpabl0/sparky-go/swid"
)

//todo driver load fix
type (
	Interface interface {
		canvasobjectapi.Interface //ui show and event finish
		MainMenu() *fyne.MainMenu
	}
	object struct {
		mainMenu *fyne.MainMenu
	}
)

func New() Interface {
	return &object{
		mainMenu: fyne.NewMainMenu(),
	}
}
func (o *object) MainMenu() *fyne.MainMenu { return o.mainMenu }

func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	topMeau := meau.New()
	Items := make([]*fyne.Menu, 0)
	Items = append(Items,
		topMeau.File(),
		topMeau.View(),
		topMeau.Debug(),
		topMeau.Trace(),
		topMeau.Plugin(),
		topMeau.Favor(),
		topMeau.Option(),
		topMeau.Help(),
	)
	o.mainMenu.Items = Items
	//todo need fix all table item layout for page view
	// ImmediateData   window and register window redesign
	// hide all table header,need api set
	//box := container.NewVBox(toolbar.New().CanvasObject(nil), fnTable.New().CanvasObject(nil))
	//command := swid.NewSelectEntryFormField("command", "", []string{"default", "script"})
	//command.Hint = "Hyper Debugger is running ..."
	//return container.NewBorder(box, command, nil, nil)

	command := swid.NewSelectEntryFormField("command", "", []string{"default", "script"})
	command.Hint = "Hyper Debugger is running ..."
	return container.NewBorder(toolbar.New().CanvasObject(nil), command, nil, nil, pageCpu.New().CanvasObject(window))
}
