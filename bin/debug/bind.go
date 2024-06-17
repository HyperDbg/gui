package main

import (
	"syscall"

	"github.com/ddkwork/golibrary/mylog"
)

func main() {
	/*
		.connect local  is HyperDbgInstallVmmDriver ???
		load vmm
	*/

	dllPath := "HPRDBGCTRL.dll"
	dll := syscall.MustLoadDLL(dllPath)

	// Support Detection
	_HyperDbgVmxSupportDetection := dll.MustFindProc("HyperDbgVmxSupportDetection")
	_HyperDbgReadVendorString := dll.MustFindProc("HyperDbgReadVendorString")

	// VMM Module
	_HyperDbgLoadVmm := dll.MustFindProc("HyperDbgLoadVmm")
	_HyperDbgUnloadVmm := dll.MustFindProc("HyperDbgUnloadVmm")
	_HyperDbgInstallVmmDriver := dll.MustFindProc("HyperDbgInstallVmmDriver")
	_HyperDbgUninstallVmmDriver := dll.MustFindProc("HyperDbgUninstallVmmDriver")
	_HyperDbgStopVmmDriver := dll.MustFindProc("HyperDbgStopVmmDriver")

	// General imports
	_HyperDbgInterpreter := dll.MustFindProc("HyperDbgInterpreter")
	_HyperDbgShowSignature := dll.MustFindProc("HyperDbgShowSignature")
	_HyperDbgSetTextMessageCallback := dll.MustFindProc("HyperDbgSetTextMessageCallback")
	_HyperDbgScriptReadFileAndExecuteCommandline := dll.MustFindProc("HyperDbgScriptReadFileAndExecuteCommandline")
	_HyperDbgContinuePreviousCommand := dll.MustFindProc("HyperDbgContinuePreviousCommand")
	_HyperDbgCheckMultilineCommand := dll.MustFindProc("HyperDbgCheckMultilineCommand")

	mylog.Check3(_HyperDbgVmxSupportDetection.Call())
	mylog.Check3(_HyperDbgReadVendorString.Call())

	return

	mylog.Check3(_HyperDbgLoadVmm.Call())
	mylog.Check3(_HyperDbgUnloadVmm.Call())
	mylog.Check3(_HyperDbgInstallVmmDriver.Call())
	mylog.Check3(_HyperDbgUninstallVmmDriver.Call())
	mylog.Check3(_HyperDbgStopVmmDriver.Call())

	mylog.Check3(_HyperDbgInterpreter.Call())
	mylog.Check3(_HyperDbgShowSignature.Call())
	mylog.Check3(_HyperDbgSetTextMessageCallback.Call())
	mylog.Check3(_HyperDbgScriptReadFileAndExecuteCommandline.Call())
	mylog.Check3(_HyperDbgContinuePreviousCommand.Call())
	mylog.Check3(_HyperDbgCheckMultilineCommand.Call())
}
