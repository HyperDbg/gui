package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	HyperDbgCtrlImports interface {
		// VMM Module
		HyperDbgLoadVmm() int
		HyperDbgUnloadVmm() int
		HyperDbgInstallVmmDriver() int
		HyperDbgUninstallVmmDriver() int
		HyperDbgStopVmmDriver() int

		// General imports
		HyperDbgInterpreter(Command *string) int
		HyperDbgShowSignature() Headers.VOID
		HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID
		HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int
		HyperDbgContinuePreviousCommand() bool
		HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool
	}
	hyperDbgCtrlImports struct {
	}
)

func (h *hyperDbgCtrlImports) HyperDbgShowSignature() Headers.VOID {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID {
	//TODO implement me
	panic("implement me")
}

// todo test on windows
func (h *hyperDbgCtrlImports) HyperDbgLoadVmm() int {
	call, u, err := HyperDbgLoadVmmProc.Call()
	if err != nil {
		return 0
	}
}

func (h *hyperDbgCtrlImports) HyperDbgUnloadVmm() int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgInstallVmmDriver() int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgUninstallVmmDriver() int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgStopVmmDriver() int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgInterpreter(Command *string) int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgContinuePreviousCommand() bool {
	//TODO implement me
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	//TODO implement me
	panic("implement me")
}

func newHyperDbgCtrlImports() HyperDbgCtrlImports {
	return &hyperDbgCtrlImports{}
}
