package control

import (
	"github.com/ddkwork/golibrary/mylog"
)

// Control is the controller all using golang
type Control struct{}

func NewControl() *Control {
	return &Control{}
}

func main() {
	/*
		.connect local  is HyperDbgInstallVmmDriver ???
		load vmm
	*/

	// todo please tell me the all function setups and fix all todo in this file

	mylog.Check(HyperDbgVmxSupportDetection())
	HyperDbgReadVendorString()

	// todo HyperDbgInstallVmmDriver first?
	HyperDbgShowSignature()
	HyperDbgLoadVmm()
	HyperDbgUnloadVmm()

	return
	HyperDbgInstallVmmDriver()
	HyperDbgUninstallVmmDriver()
	HyperDbgStopVmmDriver()

	HyperDbgInterpreter()

	HyperDbgSetTextMessageCallback()
	HyperDbgScriptReadFileAndExecuteCommandline()
	HyperDbgContinuePreviousCommand()
	HyperDbgCheckMultilineCommand()
}

func HyperDbgVmxSupportDetection() bool {
	return false
}

func HyperDbgReadVendorString() {
}

func HyperDbgLoadVmm() {
}

func HyperDbgUnloadVmm() {
}

func HyperDbgInstallVmmDriver() {
}

func HyperDbgUninstallVmmDriver() {
}

func HyperDbgStopVmmDriver() {
}

func HyperDbgInterpreter() { // todo set arg
}

func HyperDbgShowSignature() {
}

func HyperDbgSetTextMessageCallback() {
}

func HyperDbgScriptReadFileAndExecuteCommandline() { // todo set arg
}

func HyperDbgContinuePreviousCommand() {
}

func HyperDbgCheckMultilineCommand() { // todo set arg
}
