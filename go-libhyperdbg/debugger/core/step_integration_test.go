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
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ddkwork/hyperdbgsdk"
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
	// teardown 总顺序（LIFO，后注册先执行）：
	//   Continue+Disconnect → Terminate/Close proc → TERMINATE_VMX
	//   → pump.Stop → device.Close → UnloadDriver
	//
	// 关键约束（违反会 BSOD/驱动残留）：
	//   - Disconnect 必须在进程还活着、VMX 还在时发（内核 EPT/UD 清理
	//     要访问进程地址空间；晚发=对已死会话清理 → 状态污染，
	//     曾导致 csrss 0xEF CRITICAL_PROCESS_DIED BSOD）。
	//   - TERMINATE_VMX 必须在 pump.Stop 之前（DISALLOW_IOCTL 会阻止它）。
	//   - device/pump 句柄全部关闭后才能卸载驱动服务，否则 STOP_PENDING。
	t.Cleanup(func() { _ = dbg.UnloadDriver() }) // reg#1 → 最后执行

	if err := dbg.LoadDriver(driverPath); err != nil {
		t.Skipf("LoadDriver: %v (driver stuck? VT-x not available?)", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Skipf("InitVMM: %v (VT-x not available?)", err)
	}
	t.Logf("VMM loaded")
	t.Cleanup(func() { _ = dbg.device.Close() }) // reg#2 → 第6执行

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
	t.Cleanup(func() { pump.Stop() }) // reg#3 → 第5执行
	t.Cleanup(func() { // reg#4 → 第4执行：TERMINATE_VMX（须在 pump.Stop 前）
		if dbg.device != nil {
			_, _ = dbg.device.IoctlStruct(hyperdbgsdk.IoctlTerminateVmx, nil, nil, 0, 0)
		}
	})
	t.Logf("MessagePump started")

	// === 3. 启动调试进程 ===
	proc, err := dbg.StartProcess(exePath)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	t.Logf("process started: pid=%d tid=%d", proc.Pid, proc.Tid)

	// 不单独 detach/kill — UnloadVMM 内部已按 C++ ground truth 处理：
	// Continue+Detach（摘除监控）→ TERMINATE_VMX。这里的顺序是：
	//   Detach(含 Continue) → Terminate → Close → TERMINATE_VMX →
	//   pump.Stop → device.Close → UnloadDriver
	// Detach 必须最先：进程脱离 exec-trap 监控后 Terminate/VMXOFF 才安全，
	// 否则被拦截线程会让全核 DPC 永等 → 整机冻结。
	t.Cleanup(func() { _ = proc.Close() })                     // reg#5 → 第3执行
	t.Cleanup(func() { _ = proc.Terminate() })                 // reg#6 → 第2执行
	t.Cleanup(func() { _ = dbg.Detach() })                     // reg#7 → 最先执行

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

// TestStepModesAll 验证全部单步模式（GUI 显示三要素：寄存器 / 堆栈 / 反汇编）。
//
// 覆盖 OllyDbg 风格的单步家族（用户态调试）：
//   - 步入 Step Into      (F7)      → dbg.Step()
//   - 步过 Step Over      (F8)      → dbg.StepOver()（含步过 CALL 指令，内核 DR 断点路径）
//   - 执行到返回 Step Out (Ctrl+F9) → dbg.StepOut()（内核 gu：[rsp] 返回地址 DR 断点路径）
//
// "执行到用户空间"(Alt+F9) 为内核调试概念（从内核态回到用户态），用户态调试的
// 目标始终在用户空间执行，无对应语义（C++ 版同样未实现）。
//
// 每次单步后验证 GUI 显示所需的完整状态：
//   1. ReadRegisters 成功，RIP 与 PAUSED 包 pausedRIP 一致
//   2. 栈可读：RSP 处 32 字节读取成功
//   3. RIP 处反汇编可解析
//
// 运行（管理员 PowerShell）：
//
//	go test -v -count=1 -run TestStepModesAll ./debugger/core/
func TestStepModesAll(t *testing.T) {
	const driverPath = `C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys`

	system32, _ := windows.GetSystemDirectory()
	exePath := filepath.Join(system32, "notepad.exe")
	if _, err := os.Stat(exePath); err != nil {
		exePath = `C:\Users\Administrator\Desktop\calc32.exe`
		if _, err := os.Stat(exePath); err != nil {
			t.Skipf("debuggee not found")
		}
	}
	t.Logf("debuggee: %s", exePath)

	dbg := New()
	// teardown 总顺序（LIFO，后注册先执行）：
	//   Continue+Disconnect → Terminate/Close proc → TERMINATE_VMX
	//   → pump.Stop → device.Close → UnloadDriver
	// （约束同 TestStepSequence：Disconnect 须在进程活着时发，
	//   TERMINATE_VMX 须在 pump.Stop 前，句柄全关后才卸载驱动）
	t.Cleanup(func() { _ = dbg.UnloadDriver() }) // reg#1 → 最后执行

	if err := dbg.LoadDriver(driverPath); err != nil {
		t.Skipf("LoadDriver: %v", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Skipf("InitVMM: %v", err)
	}
	t.Cleanup(func() { _ = dbg.device.Close() }) // reg#2 → 第6执行

	pausedCh := make(chan struct{}, 32)
	dbg.OnPaused = func() {
		select {
		case pausedCh <- struct{}{}:
		default:
		}
	}

	pump, err := dbg.StartMessagePump()
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}
	t.Cleanup(func() { pump.Stop() }) // reg#3 → 第5执行
	t.Cleanup(func() { // reg#4 → 第4执行：TERMINATE_VMX（须在 pump.Stop 前）
		if dbg.device != nil {
			_, _ = dbg.device.IoctlStruct(hyperdbgsdk.IoctlTerminateVmx, nil, nil, 0, 0)
		}
	})

	proc, err := dbg.StartProcess(exePath)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	t.Logf("process started: pid=%d tid=%d", proc.Pid, proc.Tid)
	t.Cleanup(func() { _ = proc.Close() })                     // reg#5 → 第3执行
	t.Cleanup(func() { _ = proc.Terminate() })                 // reg#6 → 第2执行
	t.Cleanup(func() { _ = dbg.Detach() })                     // reg#7 → 最先执行

	// 等待初始 PAUSED
	select {
	case <-pausedCh:
	case <-time.After(10 * time.Second):
		_ = dbg.Pause()
		select {
		case <-pausedCh:
		case <-time.After(5 * time.Second):
			t.Fatalf("no PAUSED received — debuggee not paused")
		}
	}

	// ---- helpers ----

	// instrAt 返回 rip 处指令的 (文本, 长度)。
	instrAt := func(rip uint64) (string, uint32, bool) {
		b, _, err := readmem.ReadMemory(dbg.device, rip, dbg.processPid, 16,
			hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
		if err != nil || len(b) == 0 {
			return "", 0, false
		}
		dis := misc.NewDisassembler()
		r, err := dis.Disassemble(misc.ModeLong64, rip, b)
		if err != nil {
			return "", 0, false
		}
		return r.Text, uint32(r.Length), true
	}

	disasmAt := func(rip uint64) string {
		text, _, ok := instrAt(rip)
		if !ok {
			return "<unreadable>"
		}
		return text
	}

	// readStackQword 读 rsp+offset 处 8 字节（返回地址 / 栈帧数据）。
	readStackQword := func(rsp uint64, off uint64) (uint64, error) {
		b, _, err := readmem.ReadMemory(dbg.device, rsp+off, dbg.processPid, 8,
			hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
		if err != nil || len(b) != 8 {
			return 0, fmt.Errorf("read stack @0x%X: %v (len=%d)", rsp+off, err, len(b))
		}
		return binary.LittleEndian.Uint64(b), nil
	}

	// verifyGuiState 验证 GUI 显示三要素：寄存器 / 堆栈 / 反汇编。
	// 返回 (RIP, RSP, RFLAGS)。任何一项失败即 Fatal。
	verifyGuiState := func(label string) (uint64, uint64, uint64) {
		t.Helper()

		// 1. 寄存器
		regs, rip, rflags, err := dbg.ReadRegisters()
		if err != nil {
			t.Fatalf("%s: ReadRegisters: %v", label, err)
		}
		if rip == 0 {
			t.Fatalf("%s: RIP=0", label)
		}
		if rip != dbg.pausedRIP {
			t.Fatalf("%s: ReadRegisters RIP (0x%X) != PAUSED packet RIP (0x%X) — GUI 显示将不一致",
				label, rip, dbg.pausedRIP)
		}

		// 2. 堆栈：RSP 处 32 字节
		stack, _, err := readmem.ReadMemory(dbg.device, regs.Rsp, dbg.processPid, 32,
			hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
		if err != nil || len(stack) != 32 {
			t.Fatalf("%s: read 32 bytes at RSP=0x%X failed: %v (len=%d)", label, regs.Rsp, err, len(stack))
		}

		// 3. 反汇编：RIP 处指令可解析
		text, length, ok := instrAt(rip)
		if !ok || length == 0 {
			t.Fatalf("%s: disasm at RIP=0x%X failed", label, rip)
		}

		t.Logf("%s: RIP=0x%X (pausedRIP=0x%X) RSP=0x%X RFL=0x%X stack32=%s disasm=\"%s\" (len=%d)",
			label, rip, dbg.pausedRIP, regs.Rsp, rflags, hex.EncodeToString(stack), text, length)
		return rip, regs.Rsp, rflags
	}

	// stepUntilCall 单步直到当前指令为 CALL（不执行它），返回 (callAddr, callNext)。
	stepUntilCall := func(maxSteps int) (uint64, uint64) {
		t.Helper()
		for range maxSteps {
			rip, _, _ := verifyGuiState("scan")
			text, length, ok := instrAt(rip)
			if ok && strings.HasPrefix(text, "call") {
				return rip, rip + uint64(length)
			}
			if err := dbg.Step(); err != nil {
				t.Fatalf("scan Step failed: %v", err)
			}
		}
		return 0, 0
	}

	// ===== Phase 1: 步入 (Step Into, F7) x3 =====
	for i := range 3 {
		if err := dbg.Step(); err != nil {
			t.Fatalf("Step Into #%d failed: %v", i+1, err)
		}
		verifyGuiState(fmt.Sprintf("Step Into #%d", i+1))
	}
	t.Logf("=== Phase 1: Step Into x3 OK ===")

	// ===== Phase 2: 步过普通指令 (Step Over, F8) x3 =====
	for i := range 3 {
		ripBefore, _, _ := verifyGuiState(fmt.Sprintf("pre-StepOver #%d", i+1))
		if err := dbg.StepOver(); err != nil {
			t.Fatalf("Step Over #%d failed: %v (RIP was 0x%X, instr=%s)", i+1, err, ripBefore, disasmAt(ripBefore))
		}
		ripAfter, _, _ := verifyGuiState(fmt.Sprintf("Step Over #%d", i+1))
		// 普通指令步过 == 单步一条：RIP 前进
		if ripAfter == ripBefore {
			t.Errorf("Step Over #%d: RIP unchanged (0x%X)", i+1, ripAfter)
		}
	}
	t.Logf("=== Phase 2: Step Over (non-call) x3 OK ===")

	// ===== Phase 3: 步过 CALL 指令（内核 DR 断点路径） =====
	callAddr, callNext := stepUntilCall(60)
	if callAddr == 0 {
		t.Skipf("no CALL instruction found within 60 steps — cannot test step-over-call")
	}
	text, _, _ := instrAt(callAddr)
	t.Logf("found CALL at 0x%X (next=0x%X): %s", callAddr, callNext, text)

	_, rspBeforeCall, _ := verifyGuiState("pre-StepOver-Call")
	if err := dbg.StepOver(); err != nil {
		t.Fatalf("Step Over CALL failed: %v (call at 0x%X: %s)", err, callAddr, text)
	}
	ripAfterCall, rspAfterCall, _ := verifyGuiState("Step Over CALL")

	// 步过 CALL：应落在 callNext（call 指令的下一指令），且 RSP 不变（未进入被调函数）
	if ripAfterCall != callNext {
		t.Errorf("Step Over CALL: RIP=0x%X, want callNext=0x%X (skipped instruction?)", ripAfterCall, callNext)
	}
	if rspAfterCall != rspBeforeCall {
		t.Errorf("Step Over CALL: RSP changed 0x%X → 0x%X (entered callee?)", rspBeforeCall, rspAfterCall)
	}
	t.Logf("=== Phase 3: Step Over CALL OK (RIP=callNext=0x%X, RSP unchanged) ===", callNext)

	// ===== Phase 4: 步入 CALL 后执行到返回（内核 gu 路径） =====
	call2Addr, call2Next := stepUntilCall(60)
	if call2Addr == 0 {
		t.Skipf("no second CALL found within 60 steps — cannot test step-out")
	}
	text2, _, _ := instrAt(call2Addr)
	t.Logf("found second CALL at 0x%X (next=0x%X): %s", call2Addr, call2Next, text2)

	// 步入 CALL：执行 call 指令，进入被调函数
	if err := dbg.Step(); err != nil {
		t.Fatalf("Step Into CALL failed: %v", err)
	}
	ripInCallee, rspInCallee, _ := verifyGuiState("Step Into CALL")
	if ripInCallee == call2Next {
		t.Errorf("Step Into CALL: RIP=callNext=0x%X — call was not entered", ripInCallee)
	}
	// 步入 call 后 [rsp] == callNext（call 压入的返回地址）——步入成功的决定性证据
	pushedRet, err := readStackQword(rspInCallee, 0)
	if err != nil {
		t.Fatalf("read [rsp] after stepping into call: %v", err)
	}
	if pushedRet != call2Next {
		t.Errorf("Step Into CALL: [rsp]=0x%X, want return addr=0x%X", pushedRet, call2Next)
	}
	t.Logf("stepped into callee: RIP=0x%X, [rsp]=0x%X (== callNext ✓)", ripInCallee, pushedRet)

	// 执行到返回（Step Out / gu）
	if err := dbg.StepOut(); err != nil {
		t.Fatalf("Step Out failed: %v (was in callee at 0x%X, ret addr=0x%X)", err, ripInCallee, pushedRet)
	}
	ripAfterOut, rspAfterOut, _ := verifyGuiState("Step Out")

	// 返回断言：RIP == 之前压栈的返回地址；RSP 回升 ≥ 8（弹出了返回地址）
	if ripAfterOut != pushedRet {
		t.Errorf("Step Out: RIP=0x%X, want return addr=0x%X", ripAfterOut, pushedRet)
	}
	if rspAfterOut < rspInCallee+8 {
		t.Errorf("Step Out: RSP 0x%X not unwound above callee frame (was 0x%X, want ≥ 0x%X)",
			rspAfterOut, rspInCallee, rspInCallee+8)
	}
	t.Logf("=== Phase 4: Step Into CALL + Step Out OK (RIP==retaddr=0x%X, RSP unwound) ===", pushedRet)

	// ===== Phase 5: Step Out 之后继续单步（GUI 连续操作） =====
	if err := dbg.Step(); err != nil {
		t.Fatalf("Step after Step Out failed: %v", err)
	}
	verifyGuiState("Step after Step Out")

	t.Logf("=== TestStepModesAll PASSED: Step Into / Step Over (plain + CALL) / Step Out all verified (regs+stack+disasm) ===")
}
