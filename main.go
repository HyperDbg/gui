package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/hyperdbgui/ux/asserts"
	"github.com/ddkwork/hyperdbgui/ux/menus"
	"github.com/ddkwork/hyperdbgui/ux/tabs"
	"github.com/ddkwork/hyperdbgui/ux/toolbar"
	"github.com/richardwilkes/unison"
)

func main() {
	//NoGui()
	//select {}
	unison.AttachConsole()
	unison.Start(unison.StartupFinishedCallback(func() {
		w, err := unison.NewWindow(fmt.Sprintf("Hyper Debugger"))
		if err != nil {
			return
		}
		w.MinMaxContentSizeCallback = func() (min, max unison.Size) {
			return unison.NewSize(1000, 600), unison.NewSize(10000, 1280)
		}
		image, err := unison.NewImageFromBytes(asserts.Ico1, 0.5)
		if !mylog.Error(err) {
			return
		}
		w.SetTitleIcons([]*unison.Image{image})
		CanvasObject(w)
		w.Pack()
		rect := w.FrameRect()
		rect.Point = unison.PrimaryDisplay().Usable.Point
		w.SetFrameRect(rect) //	w.SetContent(h.CanvasObject(w))
		w.ToFront()          //	w.ShowAndRun()
	}))
}

func CanvasObject(w *unison.Window) (ok bool) {
	menus.New(w)
	content := w.Content()
	content.SetLayout(&unison.FlexLayout{Columns: 1})
	content.AddChild(toolbar.New().CanvasObject(w))
	content.AddChild(tabs.New().CanvasObject(w))
	//n := "Thaumatology - RPM Advantage Modifiers.adm"
	//n = "Template Toolkit 2 - Races Advantage Modifiers.adm"
	//noteTableDockableFromFile, err := ux.NewTraitModifierTableDockableFromFile(n)
	//if !mylog.Error(err) {
	//	return
	//}
	//var scrollArea = unison.NewScrollPanel()
	//scrollArea.SetBorder(
	//	unison.NewEmptyBorder(unison.Insets{
	//		Top:    0,
	//		Left:   0,
	//		Bottom: 600,
	//		Right:  0,
	//	}),
	//)
	//scrollArea.Sync() //todo bug: add scrollArea late the Table is not show
	//scrollArea.MarkForLayoutAndRedraw()
	//scrollArea.SetLayout(&unison.FlexLayout{
	//	Columns:  1,
	//	VSpacing: unison.StdVSpacing,
	//})
	//table.ColumnSizes[0].Minimum = 20
	//	o.mitmMock()
	//const topLevelRowsToMake = 10
	//table.HierarchyColumnIndex = 1
	//scrollArea.SetContent(noteTableDockableFromFile, unison.FillBehavior, unison.FillBehavior)
	//scrollArea.SetLayoutData(&unison.FlexLayoutData{
	//	HAlign: unison.FillAlignment,
	//	VAlign: unison.FillAlignment,
	//	HGrab:  true,
	//	VGrab:  true,
	//})
	//content.AddChild(noteTableDockableFromFile)
	//content.AddChild(bodyView.CreateBodyView())
	return true
}
