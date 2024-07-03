package main

import (
	"path/filepath"

	"github.com/ddkwork/app"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/richardwilkes/unison"
)

func main() {
	// testSdk()
	// return
	Run()
	return
	// testDisassembly()
	// return
	// testParsePe()
	// return
	testScript()
}

func testSdk() {
	mylog.Call(func() {
		absPath := mylog.Check2(filepath.Abs("hyperkd.sys"))
		SetCustomDriverPath(stringToBytePointer(absPath), stringToBytePointer(stream.BaseName(absPath)))
		VmxSupportDetection()
		mylog.Trace("InstallVmmDriver", InstallVmmDriver())
		ConnectLocalDebugger()
		mylog.Trace("LoadVmm", LoadVmm())

		// todo:
		// start debugger
		// read memory(todo bug read address buffer for disassembly,see cpu.go LayoutDisassemblyTable ,oep buf seems error)
		// read registers
		// set breakpoints
		// step over/into/out
		// continue
		// read stack
		// stop debugger

		mylog.Trace("UnloadVmm", UnloadVmm())
		mylog.Trace("StopVmmDriver", StopVmmDriver())
		mylog.Trace("UninstallVmmDriver", UninstallVmmDriver())
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
