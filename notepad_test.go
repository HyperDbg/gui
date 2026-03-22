package main

import (
	"testing"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

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
		mylog.Info("入口点地址", entryPoint)
		dbg.SetBreakpoint(entryPoint)
		dbg.StepInto()
		dbg.ReadMemory(entryPoint, 16)
		dbg.WriteMemory(entryPoint, []byte{0x90, 0x90})
		dbg.KillProcess()
	})

	t.Log("驱动加载/卸载测试通过")
}

func TestFixError(t *testing.T) {
	mylog.Call(func() {
		// fakeError.Walk(".", false)
		stream.Fmt(".")
		stream.Fix(".")
	})
}
