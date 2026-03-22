package main

import (
	"testing"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/fakeError"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestNotepadDebugging(t *testing.T) {
	dbg := debugger.NewUserDebug()
	defer dbg.UnloadDriver()

	mylog.Call(func() {
		dbg.StartProcess("c:\\windows\\system32\\notepad.exe")
		dbg.SetBreakpoint(0x401000)
		dbg.StepInto()
		dbg.ReadMemory(0x401000, 16)
		dbg.WriteMemory(0x401000, []byte{0x90, 0x90})
		// todo killprocess test
		// dbg.KillProcess()
	})

	t.Log("驱动加载/卸载测试通过")
}

func TestFixError(t *testing.T) {
	mylog.Call(func() {
		fakeError.Walk(".", false)
		stream.Fmt(".")
		stream.Fix(".")
	})
}
