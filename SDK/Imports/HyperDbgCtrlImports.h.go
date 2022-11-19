package Imports

import (
	"github.com/ddkwork/hyperdbgui/SDK/Headers"
	"unsafe"
)

type (
	VmmModule interface {
		HyperDbgLoadVmm() int
		HyperDbgUnloadVmm() int
		HyperDbgInstallVmmDriver() int
		HyperDbgUninstallVmmDriver() int
		HyperDbgStopVmmDriver() int
	}
	General interface {
		HyperDbgInterpreter(Command *string) int
		HyperDbgShowSignature() Headers.VOID
		HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID
		HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int
		HyperDbgContinuePreviousCommand() bool
		HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool
	}
	Ctrl interface {
		VmmModule
		General
	}
	ctrl struct{}
)

func NewCtrl() Ctrl { return &ctrl{} }

func (c *ctrl) HyperDbgShowSignature() Headers.VOID {
	valu := Call(api.Proc(HyperDbgShowSignature))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

func (c *ctrl) HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID {
	valu := Call(api.Proc(HyperDbgSetTextMessageCallback), uintptr(unsafe.Pointer(&handler)))
	DecodeErrorCode(valu)
	return Headers.VOID(unsafe.Pointer(valu))
}

// todo gen it and test on windows
func (c *ctrl) HyperDbgLoadVmm() int {
	valu := Call(api.Proc(HyperDbgLoadVmm))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgUnloadVmm() int {
	valu := Call(api.Proc(HyperDbgUnloadVmm))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgInstallVmmDriver() int {
	valu := Call(api.Proc(HyperDbgInstallVmmDriver))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgUninstallVmmDriver() int {
	valu := Call(api.Proc(HyperDbgUninstallVmmDriver))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgStopVmmDriver() int {
	valu := Call(api.Proc(HyperDbgStopVmmDriver))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgInterpreter(Command *string) int {
	valu := Call(api.Proc(HyperDbgInterpreter), uintptr(unsafe.Pointer(Command)))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	valu := Call(api.Proc(HyperDbgScriptReadFileAndExecuteCommandline), uintptr(argc), uintptr(unsafe.Pointer(&argv)))
	DecodeErrorCode(valu)
	return int(valu)
}

func (c *ctrl) HyperDbgContinuePreviousCommand() bool {
	api.Proc(HyperDbgContinuePreviousCommand).Call()
	valu := Call(api.Proc(HyperDbgContinuePreviousCommand))
	//DecodeErrorCode(valu) //this is not status code ?
	return valu == 0
}

func (c *ctrl) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	valu := Call(api.Proc(HyperDbgCheckMultilineCommand), uintptr(unsafe.Pointer(&CurrentCommand)), uintptr(Bool2UintPtr(Reset)))
	//DecodeErrorCode(valu)
	return valu == 0
}
