package hyper

type Window struct {
	Title     string
	ToolBar   ToolBar
	SideBar   SideBar
	Pages     []Page
	StatusBar StatusBar
}

func (w *Window) Show() error {
	return nil
}
