package Imports

import (
	"unsafe"

	"github.com/ddkwork/hyperdbgui/sdk/Headers"
)

func (c *ctrl) HyperDbgShowSignature() Headers.VOID {
	value := Call(api.Proc(HyperDbgShowSignature))
	DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (c *ctrl) HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID {
	value := Call(api.Proc(HyperDbgSetTextMessageCallback), uintptr(unsafe.Pointer(&handler)))
	DecodeErrorCode(value)
	return Headers.VOID(unsafe.Pointer(value))
}

func (c *ctrl) HyperDbgLoadVmm() int {
	value := Call(api.Proc(HyperDbgLoadVmm))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgUnloadVmm() int {
	value := Call(api.Proc(HyperDbgUnloadVmm))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgInstallVmmDriver() int {
	value := Call(api.Proc(HyperDbgInstallVmmDriver))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgUninstallVmmDriver() int {
	value := Call(api.Proc(HyperDbgUninstallVmmDriver))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgStopVmmDriver() int {
	value := Call(api.Proc(HyperDbgStopVmmDriver))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgInterpreter(Command *string) int {
	value := Call(api.Proc(HyperDbgInterpreter), uintptr(unsafe.Pointer(Command)))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	value := Call(api.Proc(HyperDbgScriptReadFileAndExecuteCommandline), uintptr(argc), uintptr(unsafe.Pointer(&argv)))
	DecodeErrorCode(value)
	return int(value)
}

func (c *ctrl) HyperDbgContinuePreviousCommand() bool {
	value := Call(api.Proc(HyperDbgContinuePreviousCommand))
	// DecodeErrorCode(value) //this is not status code ?
	return value == 0
}

func (c *ctrl) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	value := Call(api.Proc(HyperDbgCheckMultilineCommand), uintptr(unsafe.Pointer(&CurrentCommand)), Bool2UintPtr(Reset))
	// DecodeErrorCode(value)
	return value == 0
}
