package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/ddkwork/golibrary/src/fynelib/canvasobjectapi"
	"github.com/ddkwork/golibrary/src/fynelib/fyneTheme"
	"github.com/ddkwork/hyperdbgui/fnTable"
	"github.com/ddkwork/hyperdbgui/meau"
	"github.com/ddkwork/hyperdbgui/toolbar"
	"github.com/fpabl0/sparky-go/swid"
)

func main() {
	a := app.NewWithID("org.hyperdbg")
	a.SetIcon(fyne.NewStaticResource("ico1", ico1))
	fyneTheme.Dark()
	w := a.NewWindow("Hyper Debugger")
	w.SetMaster()
	w.SetPadded(false)
	h := New()
	w.SetMainMenu(h.MainMenu())
	w.CenterOnScreen()
	w.SetContent(h.CanvasObject(w))
	w.ShowAndRun()
}

var (
	//go:embed 1786555.png
	ico1 []byte

	//go:embed stock-illustration-debugger.jpg
	ico2 []byte
)

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
	//todo
	// ImmediateData   window and register window redesign
	// hide all table header,need api set
	command := swid.NewSelectEntryFormField("command", "", []string{"default", "script"})
	command.Hint = "Hyper Debugger is running ..."
	return container.NewBorder(toolbar.New().CanvasObject(nil), command, nil, nil, fnTable.New().CanvasObject(window))
}
