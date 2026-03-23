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
		dbg.Continue()
		dbg.StepInto()
		dbg.ReadMemory(entryPoint, 16)
		dbg.WriteMemory(entryPoint, []byte{0x90, 0x90})
		dbg.KillProcess(activeProcess.ProcessId)
	})

	t.Log("驱动加载/卸载测试通过")
}

func TestMultipleDriverInitialization(t *testing.T) {
	t.Log("=== 多次驱动初始化测试 ===")
	t.Log("此测试验证驱动能否多次成功加载而不BSOD")

	iterations := 5
	for i := 0; i < iterations; i++ {
		t.Logf("=== 第 %d 次初始化 ===", i+1)

		dbg := debugger.NewUserDebug()

		if !dbg.IsConnected() {
			t.Errorf("第 %d 次初始化失败: 驱动未连接", i+1)
			continue
		}

		t.Logf("第 %d 次初始化成功: 驱动已连接", i+1)

		t.Logf("注意: 第 %d 次不卸载驱动，继续下一次初始化", i+1)
	}

	t.Log("=== 多次初始化测试完成 ===")
	t.Logf("共完成 %d 次初始化", iterations)
}

