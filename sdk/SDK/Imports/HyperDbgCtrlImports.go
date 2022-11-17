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
		HyperDbgShowSignature() void
		HyperDbgSetTextMessageCallback(Callback handler) void
		HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int
		HyperDbgContinuePreviousCommand() bool
		HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool
	}
)
