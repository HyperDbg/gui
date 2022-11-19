package Imports

import "github.com/ddkwork/hyperdbgui/SDK/Headers"

type (
	VmmModule interface {
		HyperDbgLoadVmm() int
		HyperDbgUnloadVmm() int
		HyperDbgInstallVmmDriver() int
		HyperDbgUninstallVmmDriver() int
		HyperDbgStopVmmDriver() int
	}
	GeneralImports interface {
		HyperDbgInterpreter(Command *string) int
		HyperDbgShowSignature() Headers.VOID
		HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID
		HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int
		HyperDbgContinuePreviousCommand() bool
		HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool
	}
	HyperDbgCtrlImports interface {
		VmmModule
		GeneralImports
	}
	hyperDbgCtrlImports struct {
	}
)

func (h *hyperDbgCtrlImports) HyperDbgShowSignature() Headers.VOID {
	call, u, err := api.Proc(HyperDbgShowSignature).Call()
	if err != nil {
		return nil
	}

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

}

func (h *hyperDbgCtrlImports) HyperDbgInstallVmmDriver() int {
	api.Proc(HyperDbgInstallVmmDriver).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgUninstallVmmDriver() int {
	api.Proc(HyperDbgUninstallVmmDriver).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgStopVmmDriver() int {
	api.Proc(HyperDbgStopVmmDriver).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgInterpreter(Command *string) int {
	api.Proc(HyperDbgInterpreter).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	api.Proc(HyperDbgScriptReadFileAndExecuteCommandline).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgContinuePreviousCommand() bool {
	api.Proc(HyperDbgContinuePreviousCommand).Call()

}

func (h *hyperDbgCtrlImports) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	api.Proc(HyperDbgCheckMultilineCommand).Call()

}

func newHyperDbgCtrlImports() HyperDbgCtrlImports {
	return &hyperDbgCtrlImports{}
}
