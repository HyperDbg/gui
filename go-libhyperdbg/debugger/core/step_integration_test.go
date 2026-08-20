// Package core — step_integration_test.go
//
// 集成测试：验证连续单步执行（Step/StepOver）的正确性。
//
// 日志显示第一次 TraceInto 成功后，后续连续单步超时并最终返回
// 0xC0000059 (DEBUGGER_ERROR_UNABLE_TO_APPLY_COMMAND_TO_THE_TARGET_THREAD)。
// 本测试用真实驱动复现该场景，定位根因。
//
// 运行（管理员 PowerShell）：
//
//	go test -v -count=1 -run TestStepSequence ./debugger/core/
package core

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/misc"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
	"golang.org/x/sys/windows"
)

// TestStepSequence 验证连续单步执行的完整流程：
//  1. 加载驱动 + VMM + MessagePump
//  2. StartProcess（calc32/notepad）
//  3. 等待初始 PAUSED（通过 OnPaused 回调）
//  4. 连续执行 N 次 Step，每次验证 RIP 变化 + 寄存器可读
//  5. 测试 Continue → Pause 循环
func TestStepSequence(t *testing.T) {
	const driverPath = `C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys`

	// 选择调试目标：优先 64 位 notepad.exe（system32），回退 calc32.exe。
	// 用 64 位目标避免 WOW64 模式转换指令导致 step #DB 不触发的问题。
	system32, _ := windows.GetSystemDirectory()
	exePath := filepath.Join(system32, "notepad.exe")
	if _, err := os.Stat(exePath); err != nil {
		exePath = `C:\Users\Administrator\Desktop\calc32.exe`
		if _, err := os.Stat(exePath); err != nil {
			t.Skipf("debuggee not found")
		}
	}
	t.Logf("driver: %s", driverPath)
	t.Logf("debuggee: %s", exePath)

	// === 1. 创建 Debugger 并初始化 ===
	dbg := New()
	t.Cleanup(func() { _ = dbg.UnloadVMM(); _ = dbg.UnloadDriver() })

	if err := dbg.LoadDriver(driverPath); err != nil {
		t.Skipf("LoadDriver: %v (driver stuck? VT-x not available?)", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Skipf("InitVMM: %v (VT-x not available?)", err)
	}
	t.Logf("VMM loaded")

	// 用 OnPaused 回调追踪暂停事件
	pausedCh := make(chan struct{}, 16)
	dbg.OnPaused = func() {
		select {
		case pausedCh <- struct{}{}:
		default:
		}
	}

	// 启动 MessagePump
	pump, err := dbg.StartMessagePump()
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}
	// TERMINATE_VMX 必须在 pump.Stop() 之前执行：
	// pump.Stop() 发 DISALLOW_IOCTL，会阻止后续 TERMINATE_VMX，
	// 导致 VMX 未清理 → 驱动卸载卡死在 STOP_PENDING。
	t.Cleanup(func() {
		if dbg.device != nil {
			_, _ = dbg.device.IoctlStruct(comm.IOCTL_CODE_TERMINATE_VMX, nil, nil, 0, 0)
		}
		pump.Stop()
	})
	t.Logf("MessagePump started")

	// === 3. 启动调试进程 ===
	proc, err := dbg.StartProcess(exePath)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	t.Logf("process started: pid=%d tid=%d", proc.Pid, proc.Tid)

	// cleanup 顺序（LIFO，后注册先执行）：
	// 只 terminate + close，不 detach — detach 在 step 超时/异常后
	// 可能触发 EPT hook 清理访问已释放页面 → 0x50 BSOD。
	// terminate 后进程退出，内核自动清理调试会话。
	t.Cleanup(func() {
		_ = proc.Close()
	})
	t.Cleanup(func() {
		_ = proc.Terminate()
	})

	// === 4. 等待初始 PAUSED ===
	t.Logf("waiting for initial PAUSED...")
	select {
	case <-pausedCh:
		t.Logf("initial PAUSED received, RIP=0x%X", dbg.pausedRIP)
	case <-time.After(10 * time.Second):
		// 可能已经处于 intercepting phase，尝试 Pause
		t.Logf("no initial PAUSED in 10s, trying Pause...")
		_ = dbg.Pause()
		select {
		case <-pausedCh:
			t.Logf("PAUSED after manual Pause, RIP=0x%X", dbg.pausedRIP)
		case <-time.After(5 * time.Second):
			t.Fatalf("no PAUSED received — debuggee not paused")
		}
	}

	// 读初始寄存器
	regs, rip, rflags, err := dbg.ReadRegisters()
	if err != nil {
		t.Fatalf("initial ReadRegisters: %v", err)
	}
	t.Logf("initial regs: RIP=0x%X RSP=0x%X RAX=0x%X RFL=0x%X",
		rip, regs.Rsp, regs.Rax, rflags)
	prevRIP := rip

	// helper: dump bytes at rip (diagnostics)
	dumpBytes := func(rip uint64) string {
		b, _, err := readmem.ReadMemory(dbg.device, rip, dbg.processPid, 16,
			hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
		if err != nil || len(b) == 0 {
			return "<read failed>"
		}
		return hex.EncodeToString(b)
	}

	// helper: disassemble instruction at rip
	disasmAt := func(rip uint64) string {
		b, _, err := readmem.ReadMemory(dbg.device, rip, dbg.processPid, 16,
			hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
		if err != nil || len(b) == 0 {
			return "<read failed>"
		}
		dis := misc.NewDisassembler()
		r, err := dis.Disassemble(misc.ModeLong64, rip, b)
		if err != nil {
			return fmt.Sprintf("<disasm err: %v, bytes=%s>", err, hex.EncodeToString(b[:min(8, len(b))]))
		}
		return fmt.Sprintf("%s (len=%d)", r.Text, r.Length)
	}

	// helper: read debuggee memory via Windows ReadProcessMemory
	// (bypasses HyperDbg's read path — useful to verify HyperDbg isn't
	// substituting/hiding bytes via EPT split or breakpoint invisibility).
	windowsReadBytes := func(rip uint64) string {
		h, err := windows.OpenProcess(windows.PROCESS_VM_READ|windows.PROCESS_QUERY_INFORMATION, false, dbg.processPid)
		if err != nil {
			return fmt.Sprintf("<openproc: %v>", err)
		}
		defer windows.CloseHandle(h)
		var n uintptr
		buf := make([]byte, 16)
		var old uint32
		// virtualprotect first to ensure read access (executable pages may be read-only)
		_ = windows.VirtualProtectEx(h, uintptr(rip), uintptr(len(buf)), windows.PAGE_EXECUTE_READ, &old)
		if err := windows.ReadProcessMemory(h, uintptr(rip), &buf[0], uintptr(len(buf)), &n); err != nil {
			return fmt.Sprintf("<rpm: %v>", err)
		}
		return hex.EncodeToString(buf[:n])
	}

	// === 5. 连续单步 5 次 ===
	const stepCount = 5
	for i := range stepCount {
		// drain 旧的 pausedCh 信号
		select {
		case <-pausedCh:
		default:
		}

		// 打印 rip 处的指令字节，便于定位失败点
		t.Logf("Step #%d: RIP=0x%X hyperdbg_bytes=%s windows_bytes=%s hyperdbg_disasm=%s",
			i+1, prevRIP, dumpBytes(prevRIP), windowsReadBytes(prevRIP), disasmAt(prevRIP))

		err := dbg.Step()
		if err != nil {
			t.Fatalf("Step #%d failed: %v\n  prevRIP=0x%X disasm=%s bytes=%s", i+1, err, prevRIP, disasmAt(prevRIP), dumpBytes(prevRIP))
		}

		// 读寄存器验证 RIP 变化
		regs2, rip2, rflags2, err := dbg.ReadRegisters()
		if err != nil {
			t.Fatalf("Step #%d ReadRegisters: %v", i+1, err)
		}

		// 同时打印 pausedRIP（来自 PAUSED 包）与 ReadRegisters 的 RIP 对比
		t.Logf("Step #%d: RIP 0x%X → 0x%X (delta=0x%X), pausedRIP=0x%X, RSP=0x%X, RFL=0x%X, next_disasm=%s",
			i+1, prevRIP, rip2, rip2-prevRIP, dbg.pausedRIP, regs2.Rsp, rflags2, disasmAt(rip2))

		if rip2 == prevRIP {
			t.Errorf("Step #%d: RIP did not change (0x%X)", i+1, rip2)
		}
		prevRIP = rip2
	}
	t.Logf("=== %d consecutive steps all succeeded ===", stepCount)

	// === 6. 测试 Continue → Pause 循环 ===
	if err := dbg.Continue(); err != nil {
		t.Fatalf("Continue: %v", err)
	}
	t.Logf("Continue OK, process running")

	time.Sleep(500 * time.Millisecond)

	// drain
	select {
	case <-pausedCh:
	default:
	}

	err = dbg.Pause()
	if err != nil && err.Error() != "" {
		// ErrAlreadyPaused is acceptable
		t.Logf("Pause returned: %v (may be already-paused)", err)
	}

	// 等 PAUSED
	select {
	case <-pausedCh:
		t.Logf("PAUSED after Continue→Pause")
	case <-time.After(8 * time.Second):
		t.Errorf("no PAUSED after Continue→Pause within 8s")
	}

	// 读寄存器确认
	regs3, rip3, _, err := dbg.ReadRegisters()
	if err != nil {
		t.Fatalf("post-Continue ReadRegisters: %v", err)
	}
	t.Logf("after Continue→Pause: RIP=0x%X RSP=0x%X", rip3, regs3.Rsp)

	// === 7. 单步一次确认 Continue→Pause 后仍能单步 ===
	select {
	case <-pausedCh:
	default:
	}
	if err := dbg.Step(); err != nil {
		t.Errorf("Step after Continue→Pause failed: %v", err)
	} else {
		_, rip4, _, _ := dbg.ReadRegisters()
		t.Logf("Step after Continue→Pause: RIP=0x%X", rip4)
	}

	t.Logf("=== TestStepSequence PASSED ===")
}
