package main

import (
	"github.com/ddkwork/golibrary/mylog"
	"syscall"
)

func main() {
	mylog.Todo("merge bin into sdk dir")

	/*
		.connect local  is HyperDbgInstallVmmDriver ???
		load vmm
	*/

	dllPath := "../bin/debug/HPRDBGCTRL.dll"
	dll := syscall.MustLoadDLL(dllPath)

	//HyperDbgInstallVmmDriver
	//HyperDbgUninstallVmmDriver
	//HyperDbgStopVmmDriver
	//HyperDbgInterpreter

	HyperDbgLoadVmmProc := dll.MustFindProc("HyperDbgLoadVmm")
	HyperDbgUnloadVmmProc := dll.MustFindProc("HyperDbgUnloadVmm")
	HyperDbgStartVmmProc := dll.MustFindProc("HyperDbgStartVmm")
	HyperDbgStopVmmProc := dll.MustFindProc("HyperDbgStopVmm")
	HyperDbgInterpreterProc := dll.MustFindProc("HyperDbgInterpreter")

	value, statusCode := mylog.Check3(HyperDbgLoadVmmProc.Call(a...))

	//sdk := sdk.New()
	//sdk.Ctrl().HyperDbgInstallVmmDriver() // TODO decode error codes
	//sdk.Ctrl().HyperDbgLoadVmm()
	//sdk.Ctrl().HyperDbgUnloadVmm()

	// sdk.Script().PrintSymbol()
	// sdk.Script().ScriptEngineParse()
	// sdk.Symbol().SymbolInitLoad()
	select {}
}
