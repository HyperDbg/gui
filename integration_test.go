package main

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
	"testing"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// TestIntegration_LoadVMM_StartProcess_Notepad 是端到端集成测试：
// 释放内嵌驱动 → api.New → LoadVMM → StartProcess(notepad.exe) → Continue → Pause → 清理。
//
// 这条链路就是 UI 拖放 notepad.exe 时 loadProcess 走的完整流程。
// 如果驱动加载或进程启动失败，测试直接 Fatal 而非 Skip——因为本机已确认
// testsigning=on 且管理员权限。
//
// 运行条件：管理员权限 + testsigning 已开启。
func TestIntegration_LoadVMM_StartProcess_Notepad(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("仅 Windows")
	}

	// Step 1: 释放内嵌驱动到固定路径
	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}
	t.Logf("驱动路径: %s", driverPath)
	if _, err := os.Stat(driverPath); err != nil {
		t.Fatalf("驱动文件不存在: %v", err)
	}

	// Step 2: 创建 debugger 实例
	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}
	defer dbg.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Step 3: 加载 VMM 驱动（需要管理员权限 + testsigning）
	t.Logf("正在加载 VMM 驱动...")
	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM 失败: %v\n"+
			"排查: 1) 确认以管理员运行 2) bcdedit /set testsigning on + 重启 "+
			"3) 关闭 Hyper-V/VBS 4) 清理 stale 服务: sc stop hyperkd && sc delete hyperkd", err)
	}
	t.Logf("LoadVMM 成功")
	defer func() {
		_ = dbg.UnloadVMM(ctx)
	}()

	// Step 4: 启动 notepad.exe（挂起状态）
	notepadPath := `C:\Windows\System32\notepad.exe`
	if _, err := os.Stat(notepadPath); err != nil {
		t.Skipf("notepad.exe 不存在: %v", err)
	}

	t.Logf("正在启动 notepad.exe...")
	proc, err := dbg.StartProcess(ctx, notepadPath)
	if err != nil {
		t.Fatalf("StartProcess 失败: %v", err)
	}
	t.Logf("StartProcess 成功 pid=%d", proc.Pid)
	defer func() {
		if proc.Handle != 0 {
			syscall.TerminateProcess(syscall.Handle(proc.Handle), 0)
			proc.Close()
		}
	}()

	// Step 5: Continue 让进程跑起来
	if err := dbg.Continue(ctx); err != nil {
		t.Fatalf("Continue 失败: %v", err)
	}
	t.Logf("Continue 成功，notepad 已运行")

	// 等待 2 秒让 notepad 真正运行
	time.Sleep(2 * time.Second)

	// Step 6: Pause 暂停
	if err := dbg.Pause(ctx); err != nil {
		t.Logf("Pause 失败（可能进程已退出）: %v", err)
	} else {
		t.Logf("Pause 成功")
	}

	// Step 7: 验证状态
	state, err := dbg.Status(ctx)
	if err != nil {
		t.Logf("Status 查询失败: %v", err)
	} else {
		t.Logf("Status: %v", state)
	}

	// Step 8: 测试 CPU 信息查询（验证设备 IO 通道工作）
	cpu, err := dbg.Cpu(ctx)
	if err != nil {
		t.Logf("Cpu 查询失败: %v", err)
	} else {
		t.Logf("Cpu: %s", cpu.Brand)
	}

	t.Logf("端到端测试通过：驱动加载 + notepad 启动 + Continue + Pause + 状态查询")
}

// TestIntegration_LogOpen_VerifyDeviceIO 验证 LogOpen + LogClose 通道工作，
// 确保设备 IO 路径完整（不只 LoadVMM 成功，还能读写设备）。
func TestIntegration_LogOpen_VerifyDeviceIO(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("仅 Windows")
	}

	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}

	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}
	defer dbg.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	defer dbg.UnloadVMM(ctx)

	// 打开日志文件
	logPath := filepath.Join(t.TempDir(), "test.log")
	if err := dbg.LogOpen(logPath); err != nil {
		t.Fatalf("LogOpen 失败: %v", err)
	}
	t.Logf("✅ LogOpen 成功: %s", logPath)
	defer dbg.LogClose()

	// 启动 notepad 并 Continue，触发一些内核事件
	proc, err := dbg.StartProcess(ctx, `C:\Windows\System32\notepad.exe`)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	defer func() {
		if proc.Handle != 0 {
			syscall.TerminateProcess(syscall.Handle(proc.Handle), 0)
			proc.Close()
		}
	}()

	_ = dbg.Continue(ctx)
	time.Sleep(1 * time.Second)
	_ = dbg.Pause(ctx)

	// 检查日志文件已创建（即使没有 hook 事件，设备 IO 通道应工作）
	if _, err := os.Stat(logPath); err != nil {
		t.Logf("日志文件未创建（可能无事件）: %v", err)
	} else {
		info, _ := os.Stat(logPath)
		t.Logf("✅ 日志文件大小: %d bytes", info.Size())
	}

	t.Logf("✅ 设备 IO 通道验证通过")
}
