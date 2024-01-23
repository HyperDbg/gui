package hyper

import (
	"fmt"

	"fyne.io/fyne/v2"
	fyne_app "fyne.io/fyne/v2/app"
)

var app *Application

func init() {
	app = &Application{
		gui: fyne_app.New(),
	}

	app.Scaffold = Scaffold{
		ToolBar:   ToolBar{},
		SideBar:   SideBar{},
		Pages:     []Page{},
		StatusBar: StatusBar{},
		Window:    app.gui.NewWindow("Hyper Debugger"),
	}
}

func Get() *Application {
	return app
}

// TODO
//  1. Handle signals
//  2. Add plugin loader/manager

type Application struct {
	Scaffold Scaffold
	gui      fyne.App
}

func (a *Application) Run() error {
	if err := a.Scaffold.Show(); err != nil {
		return fmt.Errorf("show window: %v", err)
	}

	a.gui.Run()
	return nil
}
