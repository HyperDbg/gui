package main

import (
	"github.com/ddkwork/HyperDbg/sdk"
	"github.com/ddkwork/HyperDbg/ux"
	"github.com/ddkwork/app"
	"github.com/ddkwork/unison"
)

func main() {
	// testSdkCommands()
	sdk.TargetFilePath = "testdata/asm.exe"
	ux.Run()
	// testDisassembly()
	// testParsePe()
	// testScript()
}

func testSdkCommands() {
	app.Run("commands", func(w *unison.Window) {
		// w.Content().AddChild(sdk.LayoutCommands())
	})
}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		// w.Content().AddChild(ux.LayoutDisassemblyTable("C:\\Users\\Admin\\Desktop\\tutorial1.exe"))//32 bit not work
		ux.TargetExePath = "sdk/bin/hyperdbg-cli.exe"
		w.Content().AddChild(ux.LayoutDisassemblyTable())
	})
}

func testParsePe() {
	app.Run("pe", func(w *unison.Window) {
		ux.TargetExePath = "sdk/bin/hyperlog.dll"
		w.Content().AddChild(ux.LayoutPeView())
	})
}

func testScript() {
	app.Run("script", func(w *unison.Window) {
		w.Content().AddChild(ux.LayoutScript())
	})
}
