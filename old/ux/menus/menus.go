package menus

import (
	"github.com/ddkwork/hyperdbgui/ux/about"
)

func New(wnd *unison.Window) {
	widget.NewMenus(wnd, func(m unison.Menu) {
		unison.InsertStdMenus(m, about.ShowAboutWindow, nil, nil)
		fileMenu := m.Menu(unison.FileMenuID)
		f := fileMenu.Factory() //todo
		file := f.NewMenu(0, "file", nil)
		view := f.NewMenu(0, "view", nil)
		debug := f.NewMenu(0, "debug", nil)
		trace := f.NewMenu(0, "trace", nil)
		plugins := f.NewMenu(0, "plugins", nil)
		favour := f.NewMenu(0, "favour", nil)
		options := f.NewMenu(0, "options", nil)
		help := f.NewMenu(0, "help", nil)
		fileMenu.InsertMenu(-1, file)
		fileMenu.InsertMenu(-1, view)
		fileMenu.InsertMenu(-1, debug)
		fileMenu.InsertMenu(-1, trace)
		fileMenu.InsertMenu(-1, plugins)
		fileMenu.InsertMenu(-1, favour)
		fileMenu.InsertMenu(-1, options)
		fileMenu.InsertMenu(-1, help)

		//newMenu.InsertItem(-1, f.NewItem(-1, "file", unison.KeyBinding{
		//	KeyCode:   unison.Key0,
		//	Modifiers: 0,
		//}, func(item unison.MenuItem) bool { return true }, func(item unison.MenuItem) {}))

		//fileMenu.InsertItem(1, OpenAction.NewMenuItem(f))
	})
}
