package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
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
		content.SetLayout(&unison.FlexLayout{Columns: 1})
		t := newToolbar()
		toolBar := widget.NewToolBar(t.Elems()...)
		content.AddChild(toolBar)

		//dock := unison.NewDock()
		//dock.AsPanel().SetLayoutData(&unison.FlexLayoutData{
		//	HSpan:  1,
		//	VSpan:  1,
		//	HAlign: align.Fill,
		//	VAlign: align.Fill,
		//	HGrab:  true,
		//	VGrab:  true,
		//})

		fileName := "hyperdbg-cli.exe"
		asm := LayoutDisassemblyTable(fileName)
		TopHSplit := widget.NewHSplit(
			widget.NewTab("cpu with fast call", "todo fast call layout", true, asm),
			widget.NewTab("reg", "todo reg", true, widget.NewButton("todo reg")),
			0.3)
		content.AddChild(asm)
		return
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
