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
	call, u, err := api.Proc(HyperDbgShowSignature).Call()
	if err != nil {
		return nil
	}

}

func (c *ctrl) HyperDbgSetTextMessageCallback(handler Headers.CallBack) Headers.VOID {
	api.Proc(HyperDbgSetTextMessageCallback).Call()

}

// todo gen it and test on windows
func (c *ctrl) HyperDbgLoadVmm() int {
	call, u, err := api.Proc(HyperDbgLoadVmm).Call() //todo cut long name as HyperDbg+ api name
	if err != nil {
		return 0
	}
}

func (c *ctrl) HyperDbgUnloadVmm() int {
	api.Proc(HyperDbgUnloadVmm).Call()

}

func (c *ctrl) HyperDbgInstallVmmDriver() int {
	api.Proc(HyperDbgInstallVmmDriver).Call()

}

func (c *ctrl) HyperDbgUninstallVmmDriver() int {
	api.Proc(HyperDbgUninstallVmmDriver).Call()

}

func (c *ctrl) HyperDbgStopVmmDriver() int {
	api.Proc(HyperDbgStopVmmDriver).Call()

}

func (c *ctrl) HyperDbgInterpreter(Command *string) int {
	api.Proc(HyperDbgInterpreter).Call()

}

func (c *ctrl) HyperDbgScriptReadFileAndExecuteCommandline(argc int, argv []string) int {
	api.Proc(HyperDbgScriptReadFileAndExecuteCommandline).Call()

}

func (c *ctrl) HyperDbgContinuePreviousCommand() bool {
	api.Proc(HyperDbgContinuePreviousCommand).Call()

}

func (c *ctrl) HyperDbgCheckMultilineCommand(CurrentCommand string, Reset bool) bool {
	api.Proc(HyperDbgCheckMultilineCommand).Call()

}
