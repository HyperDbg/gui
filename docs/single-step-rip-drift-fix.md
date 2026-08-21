# TestStepSequence 单步 RIP 漂移问题修复报告

- **日期**：2026-08-21
- **测试**：`go-libhyperdbg/debugger/core/step_integration_test.go` → `TestStepSequence`
- **结果**：FAIL（超时 10.87s）→ **PASS（1.30s）**

---

## 1. 问题现象

对 notepad.exe（64 位）attach 后连续单步（`Step`），出现以下症状：

| 症状 | 表现 |
|------|------|
| RIP 漂移 | `ReadRegisters` 返回的 RIP 与 PAUSED 包的 `pausedRIP` 不一致，每步超前 2~3 字节 |
| 增量错误 | RIP 增量与实际指令长度不符（记录到 +4/+4/+5，实际指令为 2/5/3 字节） |
| 进程崩溃 | 第 4 步后 debuggee 在残缺指令上执行，随后内存读取失败 |
| Step 超时 | 超时后 recovery Pause 也无法收到新 PAUSED 包，测试失败 |

崩溃点反汇编证据（关键线索）：

```
Step #4: RIP=0x7FFA6D94D619
bytes   = 48 8b 04 25 60 00 00 00   → mov rax, [0x60]   ; 访问绝对地址 0x60 → #PF
真实指令流 D618 = 65 48 8b 04 25 60 00 00 00          → mov rax, gs:[0x60] ; 带 GS 前缀，合法
```

CPU 真实执行的字节恰好从 `ReadRegisters` 返回的 RIP 开始，且落在真指令流的**前缀中间**——说明 guest RIP 被逐 步推进到了指令边界之外，最终执行残缺指令导致崩溃。

## 2. 根因分析

### 2.1 单步机制链路

```
Go: Step() 
  → IOCTL: DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP (STEP_IN)
  → 内核 UdStepInstructions (Ud.c:303)
  → TracingRegularStepInInstruction (Tracing.c:112)
  → VmFuncSetRflagTrapFlag(TRUE)        ; 设置 RFLAGS.TF
  → AttachingConfigureInterceptingThreads(token, FALSE) ; 放行线程
  → guest 执行 1 条指令
  → #DB (trap 型异常) → VM-exit (EXCEPTION_OR_NMI)
```

### 2.2 双重推进 bug

1. **TF 触发的 #DB 是 trap 类异常**：VM-exit 发生在指令边界之后，`VMCS_GUEST_RIP` **已经指向下一条指令**（Intel SDM trap 语义）。

2. **VM-exit 处理器默认再推一次 RIP**：

   ```c
   // Vmexit.c:58 — 进入 VM-exit handler 时默认置位
   VCpu->IncrementRip = TRUE;

   // Vmexit.c:348 — VM-exit 尾部统一推进
   if (!VCpu->VmxoffState.IsVmxoffExecuted && VCpu->IncrementRip)
       HvResumeToNextInstruction();   // RIP += VMCS_VMEXIT_INSTRUCTION_LENGTH
   ```

3. **#DB 分支缺少 RIP 抑制**。对比同文件两个分支：

   ```c
   // IdtEmulation.c — #BP（软件断点，fault 型）分支：有抑制
   if (!DebuggingCallbackHandleBreakpointException(VCpu->CoreId)) {
       HvSuppressRipIncrement(VCpu);      // ← L370 有
       EventInjectBreakpoint();
   }

   // IdtEmulation.c — #DB 分支：无抑制（bug 所在）
   else if (!DebuggingCallbackHandleDebugBreakpointException(VCpu->CoreId)) {
       EventInjectInterruptOrException(InterruptExit);
   }
   // ← 调试器已处理（返回 TRUE）时什么都不做，IncrementRip 保持 TRUE
   ```

4. **结果**：`guest RIP = 下一条指令 + 被单步指令长度`，每步多推一个指令长度，逐步漂移进指令内部。

> 注：上表"每步 +4/+4/+5"是按 VMCS_VMEXIT_INSTRUCTION_LENGTH（不同 exit 类型的语义差异）累加的表现，与"每步多推一次"一致。

### 2.3 辅助根因（Go 侧，本会话早前已修）

多线程 debuggee 下还有三个用户态侧问题，会放大漂移并污染测试观察：

| # | 文件 | 问题 | 修复 |
|---|------|------|------|
| 1 | `ud_commands.go` ReadRegisters | `ApplyToAllPausedThreads=true` 导致内核遍历所有暂停线程，输出缓冲区留下**最后一个线程**的寄存器 | 改为 `TargetThreadId: d.pausedThreadId, ApplyToAllPausedThreads: false`（镜像 C++ r.cpp:221） |
| 2 | `debugger.go` MessagePump | 任意线程的 PAUSED 包无条件覆盖 `pausedRIP/pausedThreadId`，active 线程被切换 | 仅当 `pausedThreadId==0`（首次）或同线程时才更新（镜像 C++ user-listening.cpp:74） |
| 3 | `debugger.go` Continue | 恢复运行后 `pausedThreadId` 未清除，新 PAUSED 包被过滤 | Continue 成功后置 `pausedThreadId=0`（镜像 C++ g.cpp:75 IsPaused=FALSE） |

## 3. 修复内容

### 3.1 内核修复（核心，单处最小改动）

**文件**：`HyperDbg/hyperdbg/hyperhv/code/vmm/vmx/IdtEmulation.c`
**分支**：`EXCEPTION_VECTOR_DEBUG_BREAKPOINT`（#DB）

```c
else if (!DebuggingCallbackHandleDebugBreakpointException(VCpu->CoreId))
{
    //
    // It's not because of thread change detection, so re-inject it
    //
    EventInjectInterruptOrException(InterruptExit);
}
else if (VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION(VCpu->ExitQualification))
{
    //
    // The #DB was handled by the (user/kernel) debugger and it was a
    // single-step trap (RFLAGS.TF). A trap-class exception reports the
    // guest RIP *after* the stepped instruction, so the default RIP
    // increment at the end of the VM-exit handler must be suppressed;
    // otherwise the guest resumes in the middle of an instruction and
    // the debuggee gradually derails (each step is advanced by an
    // extra instruction length). DR0-3 breakpoints are fault-class
    // (RIP points to the breakpoint instruction itself) and keep the
    // original behavior.
    //
    HvSuppressRipIncrement(VCpu);
}
```

**设计要点**：

- **判定条件用 exit qualification 的 BS 位**（bit 14 `SingleInstruction`，ia32.h 已有现成宏 `VMX_EXIT_QUALIFICATION_DEBUG_EXCEPTION_SINGLE_INSTRUCTION`），即 RFLAGS.TF 单步触发的 #DB——trap 型，RIP 已在指令后。
- **仅在调试器已接管（前一分支返回 TRUE）时抑制**；re-inject 路径（交给 guest 自身调试器）保持原行为。
- **DR0-3 硬件断点不走此分支**：硬件断点是 fault 型（RIP 指向断点指令本身），本就需要推进/保持原逻辑，不受影响。
- **syscall-callback 的 TF 路径不受影响**：该路径在前面的 `if` 分支已提前返回。

### 3.2 Go 侧配套修复（本会话早前完成）

见 2.3 节表格，共 3 处，均已合入并随本次验证通过。

## 4. 验证

### 4.1 构建与部署

```powershell
# 重新编译（IdtEmulation.c 属 hyperhv，统一链接进 hyperkd.sys）
cmake --build Debug --config Debug
# → Successfully signed: D:/ux/HyperDbgUnified/Debug/hyperkd.sys

# 部署到测试加载路径
Copy-Item Debug\hyperkd.sys C:\Users\Administrator\AppData\Local\hyperdbg\hyperkd.sys
```

### 4.2 测试结果对比

**修复前**（Step #4 失败，10.87s）：

```
Step #1: RIP 0x...D60C → 0x...D610 (delta=0x4)   ; jnz 实际 2 字节 → 漂移 +2
Step #2: RIP 0x...D610 → 0x...D614 (delta=0x4)   ; add 实际 2 字节 → 漂移 +2
Step #3: RIP 0x...D614 → 0x...D619 (delta=0x5)   ; cmp 实际 2 字节 → 漂移 +3
Step #4: failed — RIP 落在 65 48 8b... 前缀中间，执行 mov rax,[0x60] → #PF → 超时
```

**修复后**（全部通过，1.30s）：

```
Step #1: 0x...D60C → 0x...D60E (delta=0x2)   ; jnz  (len=2)  ✓ 精确匹配
Step #2: 0x...D60E → 0x...D613 (delta=0x5)   ; mov  (len=5)  ✓ 精确匹配
Step #3: 0x...D613 → 0x...D616 (delta=0x3)   ; cmp  (len=3)  ✓ 精确匹配
Step #4: 0x...D616 → 0x...D62A (delta=0x14)  ; jz taken 跳转 ✓ 分支语义正确
Step #5: 0x...D62A → 0x...D62F (delta=0x5)   ; mov  (len=5)  ✓ 精确匹配

=== 5 consecutive steps all succeeded ===
Continue OK → Pause → PAUSED after Continue→Pause
Step after Continue→Pause: RIP=0x1400019C0 → 0x1400019C4 (delta=0x4, 指令 4 字节 ✓)
=== TestStepSequence PASSED ===
```

关键验证点：

| 检查项 | 修复前 | 修复后 |
|--------|--------|--------|
| RIP 增量 == 被单步指令长度 | ✗（+2~+3 漂移） | ✓ 全部精确匹配 |
| `ReadRegisters` RIP == PAUSED 包 `pausedRIP` | ✗ 不一致 | ✓ 完全一致 |
| 连续 5 步单步 | ✗ 第 4 步崩溃超时 | ✓ 全部通过 |
| Continue → Pause → 再单步 | 未到达 | ✓ 正常（RIP 0x1400019C0 属 notepad 主模块，说明 Continue 后程序真实运行） |
| 条件跳转（jz taken +0x14） | 未验证到 | ✓ 语义正确（不是简单 RIP+n） |

## 5. 涉及文件

| 文件 | 改动 | 类型 |
|------|------|------|
| `HyperDbg/hyperdbg/hyperhv/code/vmm/vmx/IdtEmulation.c` | #DB 分支新增 BS 位判断 + `HvSuppressRipIncrement` | **内核（核心修复）** |
| `go-libhyperdbg/debugger/core/ud_commands.go` | ReadRegisters 指定 TargetThreadId | Go（配套） |
| `go-libhyperdbg/debugger/core/debugger.go` | PAUSED 包线程过滤；Continue 重置 active 线程 | Go（配套） |

## 6. 遗留说明

- 测试机曾出现一次 BSOD（0x3B，FLTMGR/WdFilter，Edge 更新进程触发），与单步修复无关（栈中无 hyperhv/hyperkd 帧），未处理。
- 修复仅覆盖 **TF 单步**路径；`i` 命令（instrumentation step-in，MTF 机制）与 `p`（step-over，DR 断点机制）走独立路径，本次未改动，建议后续补充对应集成测试。
