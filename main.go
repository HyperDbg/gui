package main

import (
	"path/filepath"

	"github.com/ddkwork/HyperDbg/sdk"
	"github.com/ddkwork/HyperDbg/ux"
	"github.com/ddkwork/app"
	"github.com/ddkwork/golibrary/stream"
	"github.com/richardwilkes/unison"
)

func main() {
	// generateRegistryFile()
	// testSdkCommands()
	//ux.Run()
	// testDisassembly()
	testParsePe()
	// testScript()
}

func generateRegistryFile() {
	path := stream.RunDir()
	path += string(filepath.Separator)
	g := stream.NewGeneratedFile()
	g.P("Windows Registry Editor Version 5.00")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\HyperDbg]")
	g.P("@=\"Run HyperDbg Here\"")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\Directory\\Background\\shell\\HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe --cd=\\\"%V\\\"\"")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\Directory\\shell\\HyperDbg]")
	g.P("@=\"Run HyperDbg Here\"")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\Directory\\shell\\HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe --cd=%V\"")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\*\\shell\\Open with HyperDbg]")
	g.P("")

	g.P("[HKEY_CLASSES_ROOT\\*\\shell\\Open with HyperDbg\\command]")
	g.P("@=\"", path, "HyperDbg.exe \\\"%1\\\"\"")
	g.P("")
	stream.WriteTruncate("open.reg", g.Buffer)
	// reg import open.reg
	// reg export HKEY_CLASSES_ROOT HKCR_backup.reg
	stream.RunCommand("reg import open.reg")
}

func testSdkCommands() {
	app.Run("commands", func(w *unison.Window) {
		w.Content().AddChild(sdk.LayoutCommands())
	})
}

func testDisassembly() {
	app.Run("asm", func(w *unison.Window) {
		w.Content().AddChild(ux.LayoutDisassemblyTable("C:\\Users\\Admin\\Desktop\\tutorial1.exe"))
		// w.Content().AddChild(ux.LayoutDisassemblyTable("sdk/bin/hyperdbg-cli.exe"))
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
