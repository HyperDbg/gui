package main

import (
	"github.com/ddkwork/hyperdbgui/old/ux/asserts"
	//"github.com/ddkwork/GolandProjects/ui/widget"
	//"github.com/ddkwork/hyperdbgui/ux/asserts"
	//"github.com/ddkwork/hyperdbgui/ux/tabs"
	//"github.com/ddkwork/hyperdbgui/ux/toolbar"
	"github.com/ddkwork/hyperdbgui/old/ux/tabs"
)

//go:generate  go build .

func main() {
	//menus.New(w)
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{Columns: 1})
	panel.AddChild(toolbar.New().CanvasObject())
	panel.AddChild(tabs.New().CanvasObject())

	widget.NewWindowWithRun(func(context widget.WindowContext) widget.WindowContext {
		return widget.WindowContext{
			Title:      "Hyper Debugger",
			FullScreen: true,
			ImageBuf:   asserts.Ico1,
			Panel:      panel,
		}
	})
}
