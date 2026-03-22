package main

import (
	"testing"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

func TestNotepadDebugging(t *testing.T) {
	dbg := debugger.NewUserDebug()
	defer dbg.UnloadDriver()

	mylog.Call(func() {
		dbg.StartProcess("c:\\windows\\system32\\notepad.exe")
		activeProcess := dbg.GetActiveDebuggingProcess()
		entryPoint := activeProcess.Rip
		if entryPoint == 0 {
			t.Error("无法获取入口点地址")
			return
		}
		mylog.Hex(entryPoint)
		dbg.SetBreakpoint(entryPoint)
		dbg.StepInto()
		dbg.ReadMemory(entryPoint, 16)
		dbg.WriteMemory(entryPoint, []byte{0x90, 0x90})
		dbg.KillProcess(activeProcess.ProcessId)
	})

	t.Log("驱动加载/卸载测试通过")
}
