package main

import (
	"syscall"
	"unsafe"
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
	HyperDbgVmxSupportDetection()
	HyperDbgReadVendorString()

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
	_HyperDbgVmxSupportDetection.Call()
}

func HyperDbgReadVendorString() {
	//void HyperDbgReadVendorString(char *)
	vendorString := "YourVendorString"
	// 调用HyperDbgReadVendorString函数并传递字符串参数
	ret, _, _ := _HyperDbgReadVendorString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(vendorString))))

	// 检查返回值
	if ret != 0 {
		// 处理错误
	}
}

func HyperDbgLoadVmm() {
	_HyperDbgLoadVmm.Call()
}

func HyperDbgUnloadVmm() {
	_HyperDbgUnloadVmm.Call()
}

func HyperDbgInstallVmmDriver() {
	_HyperDbgInstallVmmDriver.Call()
}

func HyperDbgUninstallVmmDriver() {
	_HyperDbgUninstallVmmDriver.Call()
}

func HyperDbgStopVmmDriver() {
	_HyperDbgStopVmmDriver.Call()
}

func HyperDbgInterpreter() {
	_HyperDbgInterpreter.Call()
}

func HyperDbgShowSignature() {
	_HyperDbgShowSignature.Call()
}

func HyperDbgSetTextMessageCallback() {
	_HyperDbgSetTextMessageCallback.Call()
}

func HyperDbgScriptReadFileAndExecuteCommandline() {
	_HyperDbgScriptReadFileAndExecuteCommandline.Call()
}

func HyperDbgContinuePreviousCommand() {
	_HyperDbgContinuePreviousCommand.Call()
}

func HyperDbgCheckMultilineCommand() {
	_HyperDbgCheckMultilineCommand.Call()
}
