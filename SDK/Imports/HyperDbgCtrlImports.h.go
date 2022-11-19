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
	api.Proc(HyperDbgShowSignature).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID {
	api.Proc(HyperDbgSetTextMessageCallback).Call()

}

// todo gen it and test on windows
func (h *hyperDbgCtrlImports) HyperDbgLoadVmm() int {
	call, u, err := api.Proc(HyperDbgLoadVmm).Call()
	if err != nil {
		return 0
	}
}

func (h *hyperDbgCtrlImports) HyperDbgUnloadVmm() int {
	api.Proc(HyperDbgUnloadVmm).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgInstallVmmDriver() int {
	api.Proc(HyperDbgInstallVmmDriver).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgUninstallVmmDriver() int {
	api.Proc(HyperDbgUninstallVmmDriver).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgStopVmmDriver() int {
	api.Proc(HyperDbgStopVmmDriver).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgInterpreter(Command *string) int {
	api.Proc(HyperDbgInterpreter).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	api.Proc(HyperDbgScriptReadFileAndExecuteCommandline).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgContinuePreviousCommand() bool {
	api.Proc(HyperDbgContinuePreviousCommand).Call()
	panic("implement me")
}

func (h *hyperDbgCtrlImports) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	api.Proc(HyperDbgCheckMultilineCommand).Call()
	panic("implement me")
}

func newHyperDbgCtrlImports() HyperDbgCtrlImports {
	return &hyperDbgCtrlImports{}
}
