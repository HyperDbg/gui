package hyper

var application *Application

func init() {
	application = &Application{
		Window{
			Title: "Hyper Debugger",
		},
	}
}

func Get() *Application {
	return application
}

// TODO
//  1. Handle signals
//  2. Add plugin loader/manager

type Application struct {
	Window Window
}

func (a *Application) Run() error {
	return a.Window.Show()
}
