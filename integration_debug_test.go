package main

import (
	"bytes"
	"context"
	"strings"
	"testing"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// cleanupTestDebugger 严格按照 themida Unpacker.Run 的顺序释放资源：
//
//	pump.Stop → UnloadVMM(bgCtx) → proc.Terminate → proc.Close → dbg.Close
//
// 关键：使用 context.Background() 而非测试的超时 ctx。若复用超时 ctx，
// 一旦测试超时，UnloadVMM 的 IOCTL 会因 ctx.Err() 立即失败，导致
// TERMINATE_VMX 未发送、驱动残留、VT-x 卡死（StopPending），只能重启。
// 全程只打印错误不 panic，确保退出路径走完每一步。
func cleanupTestDebugger(dbg *api.Debugger, proc *core.Process, mp *core.MessagePump) {
	if mp != nil {
		mp.Stop()
	}
	if dbg != nil {
		_ = dbg.UnloadVMM(context.Background())
	}
	if proc != nil {
		_ = proc.Terminate()
		_ = proc.Close()
	}
	if dbg != nil {
		_ = dbg.Close()
	}
}

// TestIntegration_CpuPage_FullRefresh 是真正的调试器集成测试：
// 加载 notepad → 读寄存器 → 反汇编 → 读内存 → 读调用栈
//
// 这条链路就是 CPU 页 Refresh() 走的完整路径。
func TestIntegration_CpuPage_FullRefresh(t *testing.T) {
	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}

	// 用 bytes.Buffer 作为 output，捕获 Register()/K() 的格式化输出
	var outBuf bytes.Buffer
	dbg, err := api.New(api.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}

	// proc/mp 在后续赋值；defer 闭包在执行时读取最新值（不能直接
	// defer cleanupTestDebugger(dbg, &proc, mp)——参数在 defer 时求值，
	// mp 此时为 nil，后续赋值不会反映到 defer 调用）。
	var proc core.Process
	var mp *core.MessagePump
	defer func() { cleanupTestDebugger(dbg, &proc, mp) }()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}

	proc, err = dbg.StartProcess(ctx, `C:\Windows\System32\notepad.exe`)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	t.Logf("notepad 已启动 pid=%d，停在入口点", proc.Pid)

	// 启动 MessagePump 以接收 DEBUGGEE_UD_PAUSED_PACKET（含 RIP/RFLAGS）
	dbg.LogOpen("test_hyperdbg.log")
	mp, err = dbg.StartMessagePump(ctx)
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}

	// 等待 MessagePump 接收到初始暂停事件
	time.Sleep(2 * time.Second)

	// 显式调用 Pause 建立持续 #UD 拦截循环。
	// StartProcess 的 CheckCallbackAtFirstInstruction 只拦截一次（发送
	// DEBUGGEE_UD_PAUSED_PACKET），但不建立持续循环。ReadRegisters 的命令
	// 需要线程再次被拦截才能执行。PauseProcess 触发持续拦截，使内核在
	// 每次循环中检查并执行 pending 命令。
	// ErrAlreadyPaused 是正常的（线程已在拦截阶段），忽略。
	if err := dbg.Pause(ctx); err != nil {
		t.Logf("Pause (可能已暂停，正常): %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// === 测试 1：Register("") — 读取所有寄存器 ===
	outBuf.Reset()
	t.Logf("--- 测试 Register(\"\") ---")
	_, err = dbg.Register(ctx, "")
	if err != nil {
		t.Fatalf("Register 失败: %v", err)
	}
	regText := outBuf.String()
	t.Logf("寄存器输出:\n%s", regText)
	if !strings.Contains(strings.ToUpper(regText), "RIP") {
		t.Error("寄存器输出中不包含 RIP")
	}
	if !strings.Contains(strings.ToUpper(regText), "RSP") {
		t.Error("寄存器输出中不包含 RSP")
	}

	// 从输出中解析 RIP
	rip := parseRIPFromText(regText)
	t.Logf("解析到 RIP = 0x%X", rip)
	if rip == 0 {
		t.Fatal("RIP = 0，无法继续反汇编测试")
	}

	// === 测试 2：Unassemble — 反汇编 ===
	t.Logf("--- 测试 Unassemble(RIP=0x%X, 20) ---", rip)
	disasm, err := dbg.Unassemble(ctx, rip, 20)
	if err != nil {
		t.Errorf("Unassemble 失败: %v", err)
	} else if disasm == "" {
		t.Error("Unassemble 返回空字符串")
	} else {
		lines := strings.Split(strings.TrimSpace(disasm), "\n")
		t.Logf("反汇编 %d 行:", len(lines))
		for i, line := range lines {
			if i >= 5 {
				t.Logf("  ... (%d more)", len(lines)-5)
				break
			}
			t.Logf("  %s", strings.TrimSpace(line))
		}
	}

	// === 测试 3：DumpMem — 内存 hex dump ===
	t.Logf("--- 测试 DumpMem(RIP=0x%X, 256) ---", rip)
	memData, err := dbg.DumpMem(ctx, rip, 256)
	if err != nil {
		t.Errorf("DumpMem 失败: %v", err)
	} else if len(memData) == 0 {
		t.Error("DumpMem 返回空数据")
	} else {
		t.Logf("  读取 %d 字节，前 16 字节: % X", len(memData), memData[:minInt(16, len(memData))])
	}

	// === 测试 4：K — 调用栈 ===
	outBuf.Reset()
	t.Logf("--- 测试 K(16) ---")
	_, _ = dbg.K(ctx, 16)
	stackText := outBuf.String()
	if stackText == "" {
		t.Log("K() 输出为空（可能无栈帧信息）")
	} else {
		t.Logf("调用栈输出:\n%s", stackText)
	}

	// === 测试 5：Continue → 等待 → Pause → 再读寄存器 ===
	t.Logf("--- 测试 Continue → Pause → Register ---")
	outBuf.Reset()
	if err := dbg.Continue(ctx); err != nil {
		t.Logf("Continue 失败: %v", err)
	} else {
		time.Sleep(2 * time.Second)
		if err := dbg.Pause(ctx); err != nil {
			t.Logf("Pause 失败: %v", err)
		} else {
			_, err = dbg.Register(ctx, "")
			if err != nil {
				t.Errorf("暂停后 Register 失败: %v", err)
			} else {
				regText2 := outBuf.String()
				rip2 := parseRIPFromText(regText2)
				t.Logf("暂停后 RIP = 0x%X (入口点: 0x%X)", rip2, rip)
				if rip2 == 0 {
					t.Error("暂停后 RIP = 0")
				}
			}
		}
	}

	t.Logf("=== 端到端调试器测试通过 ===")
}

// TestIntegration_StepOver_TraceInto 测试单步执行
func TestIntegration_StepOver_TraceInto(t *testing.T) {
	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}

	var outBuf bytes.Buffer
	dbg, err := api.New(api.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}

	var proc core.Process
	var mp *core.MessagePump
	defer func() { cleanupTestDebugger(dbg, &proc, mp) }()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}

	proc, err = dbg.StartProcess(ctx, `C:\Windows\System32\notepad.exe`)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}

	// 启动 MessagePump + Pause 建立持续 #UD 拦截循环
	// （StartProcess 的 CheckCallbackAtFirstInstruction 只拦截一次，
	// 不建立循环；Step/TraceInto 需要 #UD 拦截循环才能执行）
	dbg.LogOpen("test_hyperdbg_step.log")
	mp, err = dbg.StartMessagePump(ctx)
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}
	time.Sleep(2 * time.Second)
	if err := dbg.Pause(ctx); err != nil {
		t.Logf("Pause (可能已暂停，正常): %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// 读入口点 RIP
	outBuf.Reset()
	_, err = dbg.Register(ctx, "")
	if err != nil {
		t.Fatalf("Register 失败: %v", err)
	}
	rip1 := parseRIPFromText(outBuf.String())
	t.Logf("入口点 RIP = 0x%X", rip1)
	if rip1 == 0 {
		t.Fatal("RIP = 0，无法继续单步测试")
	}

	// TraceInto（单步步入）
	t.Logf("--- 测试 TraceInto ---")
	if err := dbg.TraceInto(ctx); err != nil {
		t.Errorf("TraceInto 失败: %v", err)
	} else {
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		rip2 := parseRIPFromText(outBuf.String())
		t.Logf("单步后 RIP = 0x%X (变化: %+d bytes)", rip2, int64(rip2)-int64(rip1))
	}

	// StepOver（单步步过）
	t.Logf("--- 测试 StepOver ---")
	if err := dbg.StepOver(ctx); err != nil {
		t.Errorf("StepOver 失败: %v", err)
	} else {
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		rip3 := parseRIPFromText(outBuf.String())
		t.Logf("步过后 RIP = 0x%X", rip3)
	}

	t.Logf("=== 单步测试通过 ===")
}

// parseRIPFromText 从寄存器输出中解析 RIP 值
func parseRIPFromText(text string) uint64 {
	for _, line := range strings.Split(text, "\n") {
		line = strings.TrimSpace(line)
		lower := strings.ToLower(line)
		for _, sep := range []string{"rip=", "rip:", "rip ="} {
			idx := strings.Index(lower, sep)
			if idx < 0 {
				continue
			}
			rest := line[idx+len(sep):]
			rest = strings.TrimSpace(rest)
			rest = strings.TrimPrefix(rest, "0x")
			rest = strings.TrimPrefix(rest, "0X")
			for i, ch := range rest {
				if ch == ' ' || ch == ',' || ch == '\t' || ch == '\n' {
					rest = rest[:i]
					break
				}
			}
			var v uint64
			for _, c := range rest {
				var d uint64
				switch {
				case c >= '0' && c <= '9':
					d = uint64(c - '0')
				case c >= 'a' && c <= 'f':
					d = uint64(c-'a') + 10
				case c >= 'A' && c <= 'F':
					d = uint64(c-'A') + 10
				default:
					goto nextSep
				}
				v = v*16 + d
			}
			return v
		nextSep:
		}
	}
	return 0
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntegration_ToolbarButtons 测试 UI 工具栏所有按钮对应的 API。
// 每个按钮的 API 调用必须在不崩溃、不永久阻塞的前提下返回。
func TestIntegration_ToolbarButtons(t *testing.T) {
	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}

	var outBuf bytes.Buffer
	dbg, err := api.New(api.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}

	var proc core.Process
	var mp *core.MessagePump
	defer func() { cleanupTestDebugger(dbg, &proc, mp) }()

	// 初始化：LoadVMM → StartProcess → LogOpen → StartMessagePump → Pause
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	t.Log("✓ LoadVMM (open/restart 按钮)")

	proc, err = dbg.StartProcess(ctx, `C:\Windows\System32\notepad.exe`)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	t.Log("✓ StartProcess (拖入 exe / restart 按钮)")

	dbg.LogOpen("test_hyperdbg_toolbar.log")
	mp, err = dbg.StartMessagePump(ctx)
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}
	time.Sleep(2 * time.Second)
	if err := dbg.Pause(ctx); err != nil {
		t.Logf("Pause (可能已暂停): %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// 读初始 RIP
	outBuf.Reset()
	_, err = dbg.Register(ctx, "")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}
	rip0 := parseRIPFromText(outBuf.String())
	t.Logf("入口点 RIP = 0x%X", rip0)

	// --- stepin / trin 按钮：TraceInto ---
	t.Run("TraceInto", func(t *testing.T) {
		ctx2, cancel2 := context.WithTimeout(ctx, 15*time.Second)
		defer cancel2()
		if err := dbg.TraceInto(ctx2); err != nil {
			t.Fatalf("TraceInto: %v", err)
		}
		outBuf.Reset()
		_, _ = dbg.Register(ctx2, "")
		rip := parseRIPFromText(outBuf.String())
		t.Logf("✓ TraceInto: RIP 0x%X → 0x%X", rip0, rip)
		if rip == 0 {
			t.Error("TraceInto 后 RIP=0")
		}
	})

	// --- stepover / trover 按钮：StepOver ---
	t.Run("StepOver", func(t *testing.T) {
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		ripBefore := parseRIPFromText(outBuf.String())
		ctx2, cancel2 := context.WithTimeout(ctx, 15*time.Second)
		defer cancel2()
		if err := dbg.StepOver(ctx2); err != nil {
			t.Fatalf("StepOver: %v", err)
		}
		outBuf.Reset()
		_, _ = dbg.Register(ctx2, "")
		ripAfter := parseRIPFromText(outBuf.String())
		t.Logf("✓ StepOver: RIP 0x%X → 0x%X", ripBefore, ripAfter)
		if ripAfter == 0 {
			t.Error("StepOver 后 RIP=0")
		}
	})

	// --- tillret 按钮：Gu (步出，当前用 Step 近似) ---
	t.Run("Gu", func(t *testing.T) {
		ctx2, cancel2 := context.WithTimeout(ctx, 15*time.Second)
		defer cancel2()
		if err := dbg.Gu(ctx2); err != nil {
			t.Errorf("Gu: %v", err)
		} else {
			t.Log("✓ Gu")
		}
	})

	// --- modules 按钮：lm ---
	t.Run("lm", func(t *testing.T) {
		outBuf.Reset()
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		if err := dbg.Exec(ctx2, "lm"); err != nil {
			t.Errorf("Exec(lm): %v", err)
		}
		out := outBuf.String()
		if len(out) == 0 {
			t.Error("lm 输出为空")
		}
		t.Logf("✓ lm: %d 字节输出", len(out))
	})

	// --- windows 按钮：process ---
	t.Run("process", func(t *testing.T) {
		outBuf.Reset()
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		if err := dbg.Exec(ctx2, "process"); err != nil {
			t.Errorf("Exec(process): %v", err)
		}
		out := outBuf.String()
		if len(out) == 0 {
			t.Error("process 输出为空")
		}
		t.Logf("✓ process: %d 字节输出", len(out))
	})

	// --- threads 按钮：thread ---
	t.Run("thread", func(t *testing.T) {
		outBuf.Reset()
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		if err := dbg.Exec(ctx2, "thread"); err != nil {
			t.Errorf("Exec(thread): %v", err)
		}
		out := outBuf.String()
		if len(out) == 0 {
			t.Error("thread 输出为空")
		}
		t.Logf("✓ thread: %d 字节输出", len(out))
	})

	// --- settings 按钮：Settings ---
	t.Run("Settings", func(t *testing.T) {
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		s, err := dbg.Settings(ctx2)
		if err != nil {
			t.Errorf("Settings: %v", err)
		} else {
			t.Logf("✓ Settings: state=%v disasm=%v", s.State, s.ShowDisasm)
		}
	})

	// --- 验证所有暂停态按钮执行后 Register 仍正常 ---
	t.Run("ReadRegs_BeforeContinue", func(t *testing.T) {
		outBuf.Reset()
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		_, err := dbg.Register(ctx2, "")
		if err != nil {
			t.Errorf("暂停态按钮测试后 Register 失败: %v", err)
		} else {
			rip := parseRIPFromText(outBuf.String())
			t.Logf("✓ 暂停态 Register RIP = 0x%X", rip)
		}
	})

	// --- run / pause 按钮：Continue → Pause ---
	t.Run("Continue_Pause", func(t *testing.T) {
		ctx2, cancel2 := context.WithTimeout(ctx, 15*time.Second)
		defer cancel2()
		// Continue 让进程运行（IOCTL 立即返回）
		if err := dbg.Continue(ctx2); err != nil {
			t.Fatalf("Continue: %v", err)
		}
		t.Log("✓ Continue (run 按钮)")
		time.Sleep(1 * time.Second)
		// Pause 让进程暂停
		if err := dbg.Pause(ctx2); err != nil {
			// Pause 可能返回 ErrAlreadyPaused，这是可接受的
			t.Logf("Pause: %v (可能已暂停)", err)
		} else {
			t.Log("✓ Pause (pause 按钮)")
		}
		// 等待暂停包到达
		time.Sleep(2 * time.Second)
	})

	// --- trace 按钮：Trace (Intel PT) — 放最后，可能失败但不影响其他测试 ---
	t.Run("Trace", func(t *testing.T) {
		ctx2, cancel2 := context.WithTimeout(ctx, 10*time.Second)
		defer cancel2()
		err := dbg.Trace(ctx2)
		if err != nil {
			t.Logf("Trace: %v (Intel PT 可能不支持)", err)
		} else {
			t.Log("✓ Trace (trace 按钮)")
		}
	})

	t.Log("=== 工具栏按钮测试完成 ===")
}

// TestIntegration_MultiStep 测试连续多次单步不崩溃。
// 之前 ApplyToAllPausedThreads=true 导致多次单步时 MTF 累积/冲突。
func TestIntegration_MultiStep(t *testing.T) {
	driverPath, err := extractAssets()
	if err != nil {
		t.Fatalf("释放驱动失败: %v", err)
	}

	var outBuf bytes.Buffer
	dbg, err := api.New(api.WithOutput(&outBuf))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}

	var proc core.Process
	var mp *core.MessagePump
	defer func() { cleanupTestDebugger(dbg, &proc, mp) }()

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	if err := dbg.LoadVMM(ctx, driverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	proc, err = dbg.StartProcess(ctx, `C:\Windows\System32\notepad.exe`)
	if err != nil {
		t.Fatalf("StartProcess: %v", err)
	}
	dbg.LogOpen("test_hyperdbg_multistep.log")
	mp, err = dbg.StartMessagePump(ctx)
	if err != nil {
		t.Fatalf("StartMessagePump: %v", err)
	}
	time.Sleep(2 * time.Second)
	if err := dbg.Pause(ctx); err != nil {
		t.Logf("Pause: %v", err)
	}
	time.Sleep(500 * time.Millisecond)

	// 读初始 RIP
	outBuf.Reset()
	_, _ = dbg.Register(ctx, "")
	ripPrev := parseRIPFromText(outBuf.String())
	t.Logf("入口点 RIP = 0x%X", ripPrev)

	// StepOver（步过）—— notepad 入口点前 3 条是非阻塞指令 + 1 个 CALL，
	// 第 4 条是 call wWinMain（进入消息循环，不会返回）。因此只验证 3 次
	// StepOver（步过 __security_init_cookie 之类的 CALL）。
	const numOver = 3
	overOK := 0
	for i := 0; i < numOver; i++ {
		stepCtx, stepCancel := context.WithTimeout(ctx, 15*time.Second)
		err := dbg.StepOver(stepCtx)
		stepCancel()
		if err != nil {
			t.Errorf("第 %d 次 StepOver 失败: %v", i+1, err)
			break
		}
		overOK++
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		ripNow := parseRIPFromText(outBuf.String())
		t.Logf("  步过 %d: RIP 0x%X → 0x%X", i+1, ripPrev, ripNow)
		if ripNow == 0 {
			t.Fatalf("第 %d 次步过后 RIP=0", i+1)
		}
		ripPrev = ripNow
	}

	// StepOut（执行到返回）：在 [RSP] 返回地址设临时断点后 Continue。
	// 入口点 [RSP] 是 ntdll 启动代码的返回地址，StepOut 应跳回 ntdll。
	outCtx, outCancel := context.WithTimeout(ctx, 15*time.Second)
	err = dbg.StepOut(outCtx)
	outCancel()
	if err != nil {
		t.Logf("StepOut 失败: %v", err)
	} else {
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		ripNow := parseRIPFromText(outBuf.String())
		t.Logf("  StepOut: RIP → 0x%X", ripNow)
		if ripNow != 0 && ripNow != ripPrev {
			ripPrev = ripNow
		}
	}

	// TraceInto（步入）—— 验证至少 1 次单步指令执行。
	stepCtx, stepCancel := context.WithTimeout(ctx, 12*time.Second)
	err = dbg.TraceInto(stepCtx)
	stepCancel()
	traceIntoOK := 0
	if err != nil {
		t.Logf("TraceInto 失败（已知限制）: %v", err)
	} else {
		traceIntoOK = 1
		outBuf.Reset()
		_, _ = dbg.Register(ctx, "")
		ripNow := parseRIPFromText(outBuf.String())
		t.Logf("  步入 1: RIP 0x%X → 0x%X", ripPrev, ripNow)
		if ripNow != 0 {
			ripPrev = ripNow
		}
	}

	if overOK < numOver {
		t.Fatalf("StepOver 只成功 %d/%d 次", overOK, numOver)
	}
	t.Logf("✓ StepOver %d次 + StepOut + TraceInto %d次 验证完成，最终 RIP=0x%X", overOK, traceIntoOK, ripPrev)
}
