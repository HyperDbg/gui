package main

import (
	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	// testLayoutCpu()
	// return
	Run()
	return
	// testDisassembly()
	// return
	// testParsePe()
	// return
	testScript()
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
