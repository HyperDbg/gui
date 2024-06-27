package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/app/widget"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/align"
)

func main() {
	// testLayoutCpu()
	// return
	// Run()
	// return
	testDisassembly()
	return
	testParsePe()
	testScript()
}

func testLayoutCpu() {
	app.Run("cpu", func(w *unison.Window) {
		parent := w.Content()

		t := newToolbar()
		toolBar := widget.NewToolBar(t.Elems()...)
		parent.AddChild(toolBar)

		dock := unison.NewDock()
		dock.AsPanel().SetLayoutData(&unison.FlexLayoutData{
			HSpan:  1,
			VSpan:  1,
			HAlign: align.Fill,
			VAlign: align.Fill,
			HGrab:  true,
			VGrab:  true,
		})

		path := "hyperdbg-cli.exe"
		cpu := LayoutCpu(path)
		parent.AddChild(cpu)
	})
}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		w.Content().AddChild(LayoutDisassemblyTable("hyperdbg-cli.exe"))
	})
}

func testParsePe() {
	app.Run("pe", func(w *unison.Window) {
		LayoutPeView("HPRDBGCTRL.dll")
	})
}

func testScript() {
	app.Run("script", func(w *unison.Window) {
		LayoutScript()
	})
}
