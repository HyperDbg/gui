package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ddkwork/golibrary/src/fynelib/fyneTheme"
	"github.com/ddkwork/hyperdbgui"
)

func main() {
	a := app.NewWithID("org.hyperdbg")
	a.SetIcon(fyne.NewStaticResource("ico1", ico1))
	fyneTheme.Dark()
	w := a.NewWindow("Hyper Debugger")
	w.SetMaster()
	w.SetPadded(false)
	w.SetFullScreen(true)
	h := hyperdbgui.New()
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
