package main

import (
	"github.com/ddkwork/app"
	"github.com/ddkwork/hyperdbgui/sdk"
	"github.com/ddkwork/hyperdbgui/ux"
	"github.com/richardwilkes/unison"
)

func main() {
	testSdkCommands()
	// ux.Run()
	// testDisassembly()
	// testParsePe()
	// testScript()
}

func testSdkCommands() {
	app.Run("commands", func(w *unison.Window) {
		w.Content().AddChild(sdk.LayoutCommands())
	})
}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		w.Content().AddChild(ux.LayoutDisassemblyTable("sdk/bin/hyperdbg-cli.exe"))
	})
}

func testParsePe() {
	app.Run("pe", func(w *unison.Window) {
		w.Content().AddChild(ux.LayoutPeView("sdk/bin/hyperlog.dll"))
	})
}

func testScript() {
	app.Run("script", func(w *unison.Window) {
		w.Content().AddChild(ux.LayoutScript())
	})
}
