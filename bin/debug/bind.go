package main

import (
	"syscall"
	"unsafe"

	"github.com/ddkwork/golibrary/mylog"
)

var (
	dllPath = "HPRDBGCTRL.dll"
	dll     = syscall.MustLoadDLL(dllPath)

	// Support Detection
	_HyperDbgVmxSupportDetection = dll.MustFindProc("HyperDbgVmxSupportDetection")
	_HyperDbgReadVendorString    = dll.MustFindProc("HyperDbgReadVendorString")

	// VMM Module
	_HyperDbgLoadVmm            = dll.MustFindProc("HyperDbgLoadVmm")
	_HyperDbgUnloadVmm          = dll.MustFindProc("HyperDbgUnloadVmm")
	_HyperDbgInstallVmmDriver   = dll.MustFindProc("HyperDbgInstallVmmDriver")
	_HyperDbgUninstallVmmDriver = dll.MustFindProc("HyperDbgUninstallVmmDriver")
	_HyperDbgStopVmmDriver      = dll.MustFindProc("HyperDbgStopVmmDriver")

	// General imports
	_HyperDbgInterpreter                         = dll.MustFindProc("HyperDbgInterpreter")
	_HyperDbgShowSignature                       = dll.MustFindProc("HyperDbgShowSignature")
	_HyperDbgSetTextMessageCallback              = dll.MustFindProc("HyperDbgSetTextMessageCallback")
	_HyperDbgScriptReadFileAndExecuteCommandline = dll.MustFindProc("HyperDbgScriptReadFileAndExecuteCommandline")
	_HyperDbgContinuePreviousCommand             = dll.MustFindProc("HyperDbgContinuePreviousCommand")
	_HyperDbgCheckMultilineCommand               = dll.MustFindProc("HyperDbgCheckMultilineCommand")
)

func main() {
	/*
		.connect local  is HyperDbgInstallVmmDriver ???
		load vmm
	*/

	//todo please tell me the all function status and fix all todo in this file

	HyperDbgVmxSupportDetection()
	HyperDbgReadVendorString()

	//todo HyperDbgInstallVmmDriver first?

	return

	HyperDbgLoadVmm()
	HyperDbgUnloadVmm()
	HyperDbgInstallVmmDriver()
	HyperDbgUninstallVmmDriver()
	HyperDbgStopVmmDriver()

	HyperDbgInterpreter()
	HyperDbgShowSignature()
	HyperDbgSetTextMessageCallback()
	HyperDbgScriptReadFileAndExecuteCommandline()
	HyperDbgContinuePreviousCommand()
	HyperDbgCheckMultilineCommand()
}

func HyperDbgVmxSupportDetection() {
	r1, r2 := mylog.Check3(_HyperDbgVmxSupportDetection.Call())
	mylog.Trace("r1", r1) //todo return 1 is true or false ?
	mylog.Trace("r2", r2) //return 8389720? what meaning?
}

func HyperDbgReadVendorString() {
	// void HyperDbgReadVendorString(char *)
	vendorString := "GenericIntel" // todo what arg need put ?
	mylog.Check3(_HyperDbgReadVendorString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(vendorString)))))
}

func HyperDbgLoadVmm() {
	mylog.Check3(_HyperDbgLoadVmm.Call())
}

func HyperDbgUnloadVmm() {
	mylog.Check3(_HyperDbgUnloadVmm.Call())
}

func HyperDbgInstallVmmDriver() {
	mylog.Check3(_HyperDbgInstallVmmDriver.Call())
}

func HyperDbgUninstallVmmDriver() {
	mylog.Check3(_HyperDbgUninstallVmmDriver.Call())
}

func HyperDbgStopVmmDriver() {
	mylog.Check3(_HyperDbgStopVmmDriver.Call())
}

func HyperDbgInterpreter() {
	mylog.Check3(_HyperDbgInterpreter.Call())
}

func HyperDbgShowSignature() {
	mylog.Check3(_HyperDbgShowSignature.Call())
}

func HyperDbgSetTextMessageCallback() {
	mylog.Check3(_HyperDbgSetTextMessageCallback.Call())
}

func HyperDbgScriptReadFileAndExecuteCommandline() {
	mylog.Check3(_HyperDbgCheckMultilineCommand.Call())
}

func HyperDbgContinuePreviousCommand() {
	mylog.Check3(_HyperDbgCheckMultilineCommand.Call())
}

func HyperDbgCheckMultilineCommand() {
	mylog.Check3(_HyperDbgCheckMultilineCommand.Call())
}
