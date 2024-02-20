package Imports

import (
	"github.com/ddkwork/hyperdbgui/sdk/Headers"
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
