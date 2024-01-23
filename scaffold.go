package hyper

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Scaffold struct {
	ToolBar   ToolBar
	SideBar   SideBar
	Pages     []Page
	StatusBar StatusBar
	fyne.Window
}

func (s *Scaffold) Show() error {
	s.SetContent(container.NewVBox(
		widget.NewLabel("Welcome to HyperGDB!"),
	))

	s.Resize(fyne.NewSize(700, 500))
	s.CenterOnScreen()
	s.Window.Show()
	return nil
}
