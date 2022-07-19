package main

import (
	_ "embed"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
	"github.com/ddkwork/librarygo/src/fynelib/fyneTheme"
)

//go:embed 1786555.png
var ico1 []byte

//go:embed stock-illustration-debugger.jpg
var ico2 []byte

//go:generate  go build .
func main() {
	//main2()
	//return
	a := app.NewWithID("org.hyperdbg")
	a.SetIcon(fyne.NewStaticResource("ico1", ico1))
	fyneTheme.New().SetDarkTheme(a)
	w := a.NewWindow("Hyper Debugger")
	w.SetMaster()
	//w.FullScreen()
	h := hyperdbgui.New()
	w.SetMainMenu(h.MainMenu())
	w.CenterOnScreen()
	w.SetContent(h.CanvasObject(w))
	w.ShowAndRun()
}

func main2() {
	a := app.New()
	w := a.NewWindow("TEST")
	tab := container.NewDocTabs()
	tab.Append(&container.TabItem{
		Text: "Test",
		Content: container.NewMax(
			pageCpu.New().CanvasObject(w),
			//makeTableTab(w),
		),
	})
	tab.Append(&container.TabItem{
		Text: "Test2",
		Content: container.NewMax(
			pageCpu.New().CanvasObject(w),
			//makeTableTab(w),
		),
	})

	w.SetContent(tab)

	w.Resize(fyne.NewSize(800, 500))
	w.CenterOnScreen()
	w.ShowAndRun()
}

func makeTableTab(_ fyne.Window) fyne.CanvasObject {
	t := widget.NewTable(
		func() (int, int) { return 500, 150 },
		func() fyne.CanvasObject {
			return widget.NewLabel("Cell 000, 000")
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			switch id.Col {
			case 0:
				label.SetText(fmt.Sprintf("%d", id.Row+1))
			case 1:
				label.SetText("A longer cell")
			default:
				label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
			}
		})
	t.SetColumnWidth(0, 34)
	t.SetColumnWidth(1, 102)
	return t
}
