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

	dllPath := "HPRDBGCTRL.dll"
	dll := syscall.MustLoadDLL(dllPath)

	// Support Detection
	HyperDbgVmxSupportDetection := dll.MustFindProc("HyperDbgVmxSupportDetection")
	HyperDbgReadVendorString := dll.MustFindProc("HyperDbgReadVendorString")

	// VMM Module
	HyperDbgLoadVmm := dll.MustFindProc("HyperDbgLoadVmm")
	HyperDbgUnloadVmm := dll.MustFindProc("HyperDbgUnloadVmm")
	HyperDbgInstallVmmDriver := dll.MustFindProc("HyperDbgInstallVmmDriver")
	HyperDbgUninstallVmmDriver := dll.MustFindProc("HyperDbgUninstallVmmDriver")
	HyperDbgStopVmmDriver := dll.MustFindProc("HyperDbgStopVmmDriver")

	// General imports
	HyperDbgInterpreter := dll.MustFindProc("HyperDbgInterpreter")
	HyperDbgShowSignature := dll.MustFindProc("HyperDbgShowSignature")
	HyperDbgSetTextMessageCallback := dll.MustFindProc("HyperDbgSetTextMessageCallback")
	HyperDbgScriptReadFileAndExecuteCommandline := dll.MustFindProc("HyperDbgScriptReadFileAndExecuteCommandline")
	HyperDbgContinuePreviousCommand := dll.MustFindProc("HyperDbgContinuePreviousCommand")
	HyperDbgCheckMultilineCommand := dll.MustFindProc("HyperDbgCheckMultilineCommand")

	mylog.Check3(HyperDbgVmxSupportDetection.Call())
	mylog.Check3(HyperDbgReadVendorString.Call())

	return

	mylog.Check3(HyperDbgLoadVmm.Call())
	mylog.Check3(HyperDbgUnloadVmm.Call())
	mylog.Check3(HyperDbgInstallVmmDriver.Call())
	mylog.Check3(HyperDbgUninstallVmmDriver.Call())
	mylog.Check3(HyperDbgStopVmmDriver.Call())

	mylog.Check3(HyperDbgInterpreter.Call())
	mylog.Check3(HyperDbgShowSignature.Call())
	mylog.Check3(HyperDbgSetTextMessageCallback.Call())
	mylog.Check3(HyperDbgScriptReadFileAndExecuteCommandline.Call())
	mylog.Check3(HyperDbgContinuePreviousCommand.Call())
	mylog.Check3(HyperDbgCheckMultilineCommand.Call())

	select {}
}
