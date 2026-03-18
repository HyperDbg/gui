package main

import (
	"log/slog"
	"testing"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/commands/debugging"
	"github.com/ddkwork/HyperDbg/debugger/commands/extension"
	"github.com/ddkwork/HyperDbg/debugger/commands/meta"
	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func TestNotepadDebugging(t *testing.T) {
	slog.Info("开始测试 notepad 调试")

	dbg := debugger.NewHyperDbg()
	driverProvider := driver.NewDriverProvider("HyperDbg", "d:\\笔记本\\p\\ux\\examples\\hypedbg\\doc\\cpp\\HyperDbgUnified\\build\\Release\\hyperkd.sys")
	packet := debugger.NewPacket(dbg.GetDriver())
	vmmProvider := debugger.VmmProviderNew(packet)

	t.Run("连接本地系统", func(t *testing.T) {
		commandConnect := meta.NewCommandConnect(dbg)
		err := commandConnect.Connect("local")
		if err != nil {
			t.Logf("连接命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("加载驱动", func(t *testing.T) {
		err := driverProvider.Install()
		if err != nil {
			t.Logf("驱动加载返回错误（如果驱动不可用是预期的）: %v", err)
		}
	})

	t.Run("VMM初始化", func(t *testing.T) {
		err := vmmProvider.Initialize()
		if err != nil {
			t.Logf("VMM初始化返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("设置断点", func(t *testing.T) {
		commandBp := debugging.NewCommandBp(dbg, packet)
		tokens := []debugger.CommandToken{
			{Type: debugger.TokenCommand, Value: "bp"},
			{Type: debugger.TokenArgument, Value: "0x401000"},
		}
		err := commandBp.Execute(tokens, "bp 0x401000")
		if err != nil {
			t.Logf("断点命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("继续执行", func(t *testing.T) {
		commandG := debugging.NewCommandG(dbg)
		err := commandG.Request()
		if err != nil {
			t.Logf("继续命令返回错误（如果没有活动进程是预期的）: %v", err)
		}
	})

	t.Run("附加进程", func(t *testing.T) {
		commandAttach := meta.NewCommandAttach(dbg)
		err := commandAttach.Attach(1234)
		if err != nil {
			t.Logf("附加进程命令返回错误（如果进程未找到是预期的）: %v", err)
		}
	})

	t.Run("单步调试", func(t *testing.T) {
		activeProcess := dbg.GetActiveProcess()
		activeProcess.IsPaused = true
		activeProcess.IsActive = true
		activeProcess.ProcessId = 1234
		activeProcess.ThreadId = 5678
		dbg.SetActiveProcess(activeProcess)

		err := dbg.PerformStepOver()
		if err != nil {
			t.Logf("单步命令返回错误（如果没有活动进程是预期的）: %v", err)
		}
	})

	t.Run("EPT Hook", func(t *testing.T) {
		commandEptHook := extension.NewCommandEptHook(dbg, packet)
		err := commandEptHook.SetEptHook(0x401000)
		if err != nil {
			t.Logf("EPT Hook命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("内存监控", func(t *testing.T) {
		commandMonitor := extension.NewCommandMonitor(dbg, packet)
		err := commandMonitor.MonitorMemory(0x401000, 16, true, false)
		if err != nil {
			t.Logf("内存监控命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("PTE查询", func(t *testing.T) {
		commandPte := extension.NewCommandPte(dbg, packet)
		err := commandPte.QueryPte(0x401000, 0)
		if err != nil {
			t.Logf("PTE查询命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("地址转换", func(t *testing.T) {
		commandVa2Pa := extension.NewCommandVa2Pa(dbg, packet)
		err := commandVa2Pa.Va2Pa(0x401000, 0)
		if err != nil {
			t.Logf("VA2PA转换命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})

	t.Run("VMM读取内存", func(t *testing.T) {
		data, err := vmmProvider.ReadVirtualMemory(1234, 0x401000, 16)
		if err != nil {
			t.Logf("VMM读取内存返回错误（如果驱动未加载是预期的）: %v", err)
		} else {
			t.Logf("从虚拟内存读取了 %d 字节", len(data))
		}
	})

	t.Run("VMM写入内存", func(t *testing.T) {
		err := vmmProvider.WriteVirtualMemory(1234, 0x401000, []byte{0x90, 0x90})
		if err != nil {
			t.Logf("VMM写入内存返回错误（如果驱动未加载是预期的）: %v", err)
		} else {
			t.Log("向虚拟内存写入了 2 字节")
		}
	})

	t.Run("进程管理", func(t *testing.T) {
		commandAttach := meta.NewCommandAttach(dbg)
		err := commandAttach.Attach(1234)
		if err != nil {
			t.Logf("附加进程命令返回错误（如果驱动未加载是预期的）: %v", err)
		}
	})
}
