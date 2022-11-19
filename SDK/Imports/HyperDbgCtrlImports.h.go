package Imports

type (
	HyperDbgCtrlImports interface {
		// VMM Module
		HyperDbgLoadVmm() int
		HyperDbgUnloadVmm() int
		HyperDbgInstallVmmDriver() int
		HyperDbgUninstallVmmDriver() int
		HyperDbgStopVmmDriver() int

		// General imports
		HyperDbgInterpreter(Command string) int //todo *string ?
		//HyperDbgShowSignature() void
		//HyperDbgSetTextMessageCallback(Callback handler) void
		HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int
		HyperDbgContinuePreviousCommand() bool
		HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool
	}
	hyperDbgCtrlImports struct {
	}
)

// todo come to windows init dll and write call dll method
func (h *hyperDbgCtrlImports) HyperDbgLoadVmm() int {
	//TODO implement me
	panic("implement me")
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

func (h *hyperDbgCtrlImports) HyperDbgInterpreter(Command string) int {
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
