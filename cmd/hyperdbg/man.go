package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/ddkwork/hyperdbgui"
	"github.com/ddkwork/librarygo/src/fynelib/fyneTheme"
)

//go:embed 1786555.png
var ico1 []byte

//go:embed stock-illustration-debugger.jpg
var ico2 []byte

//go:generate  go build .
func main() {
	a := app.NewWithID("org.hyperdbg")
	a.SetIcon(fyne.NewStaticResource("ico1", ico1))
	fyneTheme.New().SetDarkTheme(a)
	w := a.NewWindow("Hyper Debugger")
	w.SetMaster()
	w.FullScreen()
	h := hyperdbgui.New()
	w.SetMainMenu(h.MainMenu())
	w.FullScreen()
	w.CenterOnScreen()
	w.SetContent(h.CanvasObject(w))
	w.ShowAndRun()
}
