package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ddkwork/hyperdbgui/src/fnTable/pageCpu"
)

func test() {
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
