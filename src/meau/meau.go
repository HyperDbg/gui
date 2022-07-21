package meau

import "fyne.io/fyne/v2"

type (
	Interface interface {
		//canvasobjectapi.Interface
		//Fn() (ok bool)
		File() *fyne.Menu
		View() *fyne.Menu
		Debug() *fyne.Menu
		Trace() *fyne.Menu
		Plugin() *fyne.Menu
		Favor() *fyne.Menu
		Option() *fyne.Menu
		Help() *fyne.Menu
	}
	object struct {
		file    *fyne.Menu
		view    *fyne.Menu
		debug   *fyne.Menu
		trace   *fyne.Menu
		plugins *fyne.Menu
		favour  *fyne.Menu
		options *fyne.Menu
		help    *fyne.Menu
	}
)

func (o *object) File() *fyne.Menu {
	o.file.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("open", func() {}),
		fyne.NewMenuItem("attach", func() {}),
		fyne.NewMenuItem("database", func() {}),
		fyne.NewMenuItem("path", func() {}),
		fyne.NewMenuItem("command", func() {}),
		fyne.NewMenuItem("exit", func() { fyne.CurrentApp().Quit() }),
	}
	return o.file
}

func (o *object) View() *fyne.Menu {
	o.view.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.view
}

func (o *object) Debug() *fyne.Menu {
	o.debug.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.debug
}

func (o *object) Trace() *fyne.Menu {
	o.trace.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.trace
}

func (o *object) Plugin() *fyne.Menu {
	o.plugins.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.plugins
}

func (o *object) Favor() *fyne.Menu {
	o.favour.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.favour
}

func (o *object) Option() *fyne.Menu {
	o.options.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.options
}

func (o *object) Help() *fyne.Menu {
	o.help.Items = []*fyne.MenuItem{
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
		fyne.NewMenuItem("todo", func() {}),
	}
	return o.help
}

func New() Interface {
	return &object{
		file:    fyne.NewMenu("file"),
		view:    fyne.NewMenu("view"),
		debug:   fyne.NewMenu("debug"),
		trace:   fyne.NewMenu("trace"),
		plugins: fyne.NewMenu("plugins"),
		favour:  fyne.NewMenu("favour"),
		options: fyne.NewMenu("options"),
		help:    fyne.NewMenu("help"),
	}
}
