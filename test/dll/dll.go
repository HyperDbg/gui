package main

import (
	"syscall"

	"github.com/ddkwork/golibrary/mylog"
)

func main() {
	mylog.Todo("merge bin into sdk dir")

	/*
		.connect local  is HyperDbgInstallVmmDriver ???
		load vmm
	*/

	dllPath := "../../bin/debug/HPRDBGCTRL.dll" // not working
	dllPath = "HPRDBGCTRL.dll"
	dll := syscall.MustLoadDLL(dllPath)

	// HyperDbgInstallVmmDriver
	// HyperDbgUninstallVmmDriver
	// HyperDbgStopVmmDriver
	// HyperDbgInterpreter

	HyperDbgLoadVmmProc := dll.MustFindProc("HyperDbgLoadVmm")
	HyperDbgUnloadVmmProc := dll.MustFindProc("HyperDbgUnloadVmm")
	// HyperDbgStartVmmProc := dll.MustFindProc("HyperDbgStartVmm")
	// HyperDbgStopVmmProc := dll.MustFindProc("HyperDbgStopVmm")
	// HyperDbgInterpreterProc := dll.MustFindProc("HyperDbgInterpreter")

	mylog.Check3(HyperDbgLoadVmmProc.Call())
	mylog.Check3(HyperDbgUnloadVmmProc.Call())

	// sdk := sdk.New()
	// sdk.Ctrl().HyperDbgInstallVmmDriver() // TODO decode error codes
	// sdk.Ctrl().HyperDbgLoadVmm()
	// sdk.Ctrl().HyperDbgUnloadVmm()

	// sdk.Script().PrintSymbol()
	// sdk.Script().ScriptEngineParse()
	// sdk.Symbol().SymbolInitLoad()
	select {}
}
