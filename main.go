package main

import (
	"time"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/app"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/side"
)

func main() {
	testLayoutCpu()
	return
	// Run()
	// return
	// testDisassembly()
	// return
	// testParsePe()
	// return
	testScript()
}

func testLayoutCpu() {
	app.Run("cpu", func(w *unison.Window) {
		content := w.Content()
		content.SetLayout(&unison.FlexLayout{
			Columns:      1,
			HSpacing:     0,
			VSpacing:     0,
			HAlign:       align.Fill,
			VAlign:       align.Fill,
			EqualColumns: false,
		})
		t := newToolbar()
		toolBar := widget.NewToolBar(t.Elems()...)
		content.AddChild(toolBar)

		fileName := "hyperdbg-cli.exe"
		asm := LayoutDisassemblyTable(fileName)

		mylog.Todo("NewTab has some bug")

		left := widget.NewTab("cpu with fast call", "todo fast call layout", true, asm)
		right := widget.NewTab("reg", "todo reg", true, widget.NewButton("todo reg"))

		dock := unison.NewDock()
		dock.AsPanel().SetLayoutData(&unison.FlexLayoutData{
			HSpan:  1,
			VSpan:  1,
			HAlign: align.Fill,
			VAlign: align.Fill,
			HGrab:  true,
			VGrab:  true,
		})
		display := unison.PrimaryDisplay()
		r := display.Usable.Size
		pos := float32(0)

		dock.DockTo(left, nil, side.Left)
		// LeftContainer := widget.NewDockContainer(left)
		LeftContainer := unison.Ancestor[*unison.DockContainer](left)
		dock.DockTo(right, LeftContainer, side.Right)
		RightContainer := widget.NewDockContainer(right)
		scale := float32(0.3)
		pos = (r.Width/(display.ScaleX*1000)/2 + scale) * 1000
		RightContainer.Stack(right, -1)
		dock.RootDockLayout().SetDividerPosition(pos)
		unison.InvokeTaskAfter(func() { dock.RootDockLayout().SetDividerPosition(pos) }, time.Millisecond)

		LeftContainer.SetCurrentDockable(left)   // left not work
		RightContainer.SetCurrentDockable(right) // right not work
		content.AddChild(dock)

		return
		// content.AddChild(left)
		// return
		TopHSplit := widget.NewHSplit(
			left,
			widget.NewTab("reg", "todo reg", true, widget.NewButton("todo reg")),
			0.3)
		// content.AddChild(asm)
		// return
		content.AddChild(TopHSplit) // todo bug

		return
		cpu := LayoutCpu(fileName)
		content.AddChild(cpu)
	})
}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		w.Content().AddChild(LayoutDisassemblyTable("hyperdbg-cli.exe"))
	})
}

func testParsePe() {
	app.Run("pe", func(w *unison.Window) {
		w.Content().AddChild(LayoutPeView("hyperlog.dll"))
	})
}

func testScript() {
	app.Run("script", func(w *unison.Window) {
		w.Content().AddChild(LayoutScript())
	})
}
