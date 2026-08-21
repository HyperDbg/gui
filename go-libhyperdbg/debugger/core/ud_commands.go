// Package core — ud_commands.go
//
// 用户态调试器命令（User-Debugger Commands）：通过 hyperdbgsdk.IoctlSendUserDebuggerCommands
// 向内核发送 ReadRegisters / RegularStep 指令。
//
// 对应 C++ libhyperdbg 的 UdCommandReadRegisters / UdCommandStepGeneralCommand 路径。
package core

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/debugger/misc"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
)

// udPacketSize 是 DEBUGGER_UD_COMMAND_PACKET 的字节大小。
const udPacketSize = unsafe.Sizeof(hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET{})

func minUint32(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}

// regDescSize 是 DEBUGGEE_REGISTER_READ_DESCRIPTION 的字节大小。
const regDescSize = unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_REGISTER_READ_DESCRIPTION{})

// guestRegsSize 是 GUEST_REGS 的字节大小。
const guestRegsSize = unsafe.Sizeof(hyperdbgsdk.GUEST_REGS{})

// extraRegsSize 是 GUEST_EXTRA_REGISTERS 的字节大小。
const extraRegsSize = unsafe.Sizeof(hyperdbgsdk.GUEST_EXTRA_REGISTERS{})

// DebuggeeShowAllRegisters 对应 C 宏 DEBUGGEE_SHOW_ALL_REGISTERS = 0xFFFFFFFF。
// SDK 已绑定: hyperdbgsdk.DebuggeeShowAllRegisters。

// ReadRegisters 读取当前暂停的调试目标的全部通用寄存器 + RIP + RFLAGS。
//
// 对应 C libhyperdbg 的 HyperDbgReadAllRegisters（r.cpp:175）：
//   - RegisterId = DEBUGGEE_SHOW_ALL_REGISTERS（0xFFFFFFFF）
//   - OptionalBuffer 布局：[DEBUGGEE_REGISTER_READ_DESCRIPTION][GUEST_REGS][GUEST_EXTRA_REGISTERS]
//   - 内核填充 GUEST_REGS 和 GUEST_EXTRA_REGISTERS 后返回
//
// 前提：调试目标必须处于暂停状态（StartProcess with CheckCallbackAtFirstInstruction
// 或 Pause 后）。C 代码在 r.cpp:216 检查 g_ActiveProcessDebuggingState.IsPaused。
//
// 返回 GUEST_REGS（16 个通用寄存器）+ RIP + RFLAGS（后两者来自 GUEST_EXTRA_REGISTERS）。
func (d *Debugger) ReadRegisters() (regs hyperdbgsdk.GUEST_REGS, rip uint64, rflags uint64, err error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.processToken == 0 {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: no process attached")
	}
	if d.pausedThreadId == 0 {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: no paused thread (wait for PAUSED or call Pause)")
	}

	// 构建 OptionalBuffer: [DEBUGGEE_REGISTER_READ_DESCRIPTION][GUEST_REGS][GUEST_EXTRA_REGISTERS]
	optBufSize := regDescSize + guestRegsSize + extraRegsSize
	optBuf := make([]byte, optBufSize)
	desc := (*hyperdbgsdk.DEBUGGEE_REGISTER_READ_DESCRIPTION)(unsafe.Pointer(&optBuf[0]))
	desc.RegisterId = hyperdbgsdk.DebuggeeShowAllRegisters

	// 构建 IOCTL packet
	pkt := hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET{
		UdAction: hyperdbgsdk.DEBUGGER_UD_COMMAND_ACTION{
			ActionType:     hyperdbgsdk.DebuggerUdCommandActionTypeReadRegisters,
			OptionalParam1: uint64(hyperdbgsdk.DebuggeeShowAllRegisters), // C: RegDes->RegisterId
		},
		ProcessDebuggingDetailToken: d.processToken,
		// 镜像 C++ r.cpp:221 → ud.cpp:1160-1166 UdSendReadRegisterToUserDebugger:
		// 传具体 active ThreadId + ApplyToAllPausedThreads=FALSE。
		// 不要用 ApplyToAllPausedThreads=true——调试目标常有多个暂停线程
		// （其它线程被 debug 拦截时也会发 PAUSED 包），内核遍历所有暂停
		// 线程后输出缓冲区留下的是"最后一个被处理线程"的寄存器，导致
		// 读到的 RIP/寄存器属于另一个线程。
		TargetThreadId:          d.pausedThreadId,
		ApplyToAllPausedThreads: false,
		WaitForEventCompletion:  true,
	}

	// 完整缓冲区: [DEBUGGER_UD_COMMAND_PACKET][OptionalBuffer]
	totalSize := udPacketSize + optBufSize
	buf := make([]byte, totalSize)
	pktBytes := (*[udPacketSize]byte)(unsafe.Pointer(&pkt))[:]
	copy(buf[0:], pktBytes)
	copy(buf[udPacketSize:], optBuf)

	if _, err = d.device.Ioctl(hyperdbgsdk.IoctlSendUserDebuggerCommands, buf, buf); err != nil {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: IOCTL failed: %w", err)
	}

	// 检查 packet Result
	outPkt := (*hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET)(unsafe.Pointer(&buf[0]))
	if outPkt.Result != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: kernel error (Result=0x%X)", outPkt.Result)
	}

	// 检查 DEBUGGEE_REGISTER_READ_DESCRIPTION.KernelStatus
	outDesc := (*hyperdbgsdk.DEBUGGEE_REGISTER_READ_DESCRIPTION)(unsafe.Pointer(&buf[udPacketSize]))
	if outDesc.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: kernel status error (0x%X)", outDesc.KernelStatus)
	}

	// GUEST_REGS 紧跟 DEBUGGEE_REGISTER_READ_DESCRIPTION 之后
	off := udPacketSize + regDescSize
	regs = *(*hyperdbgsdk.GUEST_REGS)(unsafe.Pointer(&buf[off]))

	// GUEST_EXTRA_REGISTERS 紧跟 GUEST_REGS 之后（含 RIP 和 RFLAGS）
	off += guestRegsSize
	extra := *(*hyperdbgsdk.GUEST_EXTRA_REGISTERS)(unsafe.Pointer(&buf[off]))

	// 镜像 C++ r.cpp:386 — 直接用 IOCTL 返回的 ExtraRegs.RIP / RFLAGS。
	// 不替换为 PAUSED 包中的 pausedRIP：C++ HyperDbgReadAllRegisters 也
	// 只读 GUEST_EXTRA_REGISTERS，不存在用 PAUSED 包覆盖 RIP 的逻辑。
	rip = extra.RIP
	rflags = extra.RFLAGS

	return regs, rip, rflags, nil
}

// maxInstructionLength 是 x86-64 单条指令的最大长度（15 字节）。
const maxInstructionLength = 15

// Step 执行单步步入。
//
// 镜像 C++ SteppingRegularStepIn()（steppings.cpp:80）→
// UdSendStepPacketToDebuggee(Token, ThreadId, STEP_IN)（ud.cpp:1255）。
// C++ 对 STEP_IN: IsCall=FALSE, CallLen=0。
// 内核 UdStepInstructions 对 STEP_IN 调 TracingRegularStepInInstruction()
// （Tracing.c:112）→ VmFuncSetRflagTrapFlag(TRUE) 设 TF → Continue →
// 下条指令触发 #DB → 线程重新暂停 → 内核发 DEBUGGEE_UD_PAUSED_PACKET。
func (d *Debugger) Step() error {
	return d.stepImpl(uint64(hyperdbgsdk.DebuggerRemoteSteppingRequestStepIn), false, 0)
}

// StepOver 执行单步步过。
//
// 镜像 C++ UdSendStepPacketToDebuggee(StepType=DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER)
// （ud.cpp:1255）。C++ 用 HyperDbgCheckWhetherTheCurrentInstructionIsCall 判断：
//   - 是 CALL: IsCall=TRUE, CallLen=指令长度 → 内核 UdRegularStepOver 在
//     LastRip+CallLen 设硬件 DR 断点（Ud.c:227）
//   - 非 CALL: IsCall=FALSE, CallLen=0 → 内核走 TracingRegularStepInInstruction (TF)
//
// 反汇编失败（zydis DLL 未加载、内存不可读等）时退化为 STEP_IN，
// 不返回错误，避免 UI 步过按钮在边界场景下整体不可用。
func (d *Debugger) StepOver() error {
	isCall, callLen, err := d.detectCallAtPausedRip()
	if err != nil {
		isCall, callLen = false, 0
	}
	return d.stepImpl(uint64(hyperdbgsdk.DebuggerRemoteSteppingRequestStepOver), isCall, callLen)
}

// stepImpl 是 Step / StepOver / StepOut 的共享实现。
//
// stepType: StepIn / StepOver / StepOverForGu，对应 C++ steppings.cpp 的三个函数。
// isCall / callLen: 仅 StepOver 且当前指令为 CALL 时有意义（C++ 仅对 STEP_OVER 检查 IsCall）。
//
// 流程（镜像 C++ UdSendStepPacketToDebuggee + DbgWaitForUserResponse）：
//  1. 排空 pauseEvent 旧信号
//  2. 发 RegularStep IOCTL（WaitForEventCompletion=false，立即返回）
//  3. 等 pauseEvent（MessagePump 收到 PAUSED 包后信号）
//  4. 超时→pauseProcess 强制暂停→再等一次
//
// 内核 UdStepInstructions 路径（Ud.c:303）：
//   - STEP_IN: TracingRegularStepInInstruction (设 TF)
//   - STEP_OVER + IsCall: UdRegularStepOver (设 DR 断点 at LastRip+CallLen)
//   - STEP_OVER + 非 CALL: TracingRegularStepInInstruction (设 TF)
//   - STEP_OVER_FOR_GU: UdRegularStepOverForGu
//
// WaitForEventCompletion=false 与 C++ UdSendCommand(...,FALSE,...) 一致：
// IOCTL 立即返回，随后由 pauseEvent 等待 PAUSED 包（对应 DbgWaitForUserResponse）。
func (d *Debugger) stepImpl(stepType uint64, isCall bool, callLen uint32) error {
	if d.processToken == 0 {
		return fmt.Errorf("Step: no process attached")
	}

	d.mu.Lock()
	// 排空 pauseEvent 旧信号，确保等到的是本次 step 的 PAUSED。
	if d.pauseEvent != nil {
		select {
		case <-d.pauseEvent:
		default:
		}
	}
	pkt := hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET{
		UdAction: hyperdbgsdk.DEBUGGER_UD_COMMAND_ACTION{
			ActionType:     hyperdbgsdk.DebuggerUdCommandActionTypeRegularStep,
			OptionalParam1: stepType,
			OptionalParam2: boolToUint64(isCall),
			OptionalParam3: uint64(callLen),
		},
		ProcessDebuggingDetailToken: d.processToken,
		// 使用从 DEBUGGEE_UD_PAUSED_PACKET 获取的 ThreadId 指定单线程，
		// 与 C libhyperdbg 的 g_ActiveProcessDebuggingState.ThreadId 一致。
		TargetThreadId:          d.pausedThreadId,
		ApplyToAllPausedThreads: false,
		// WaitForEventCompletion=false: 内核设 RFLAGS.TF 并恢复线程后
		// 立即返回 IOCTL。step 完成时 #DB 触发→线程重新暂停→内核通过
		// IRP 通道发 DEBUGGEE_UD_PAUSED_PACKET→MessagePump 信号 pauseEvent。
		// 这与 C++ UdSendStepPacketToDebuggee + DbgWaitForUserResponse 一致。
		WaitForEventCompletion: false,
	}
	dev := d.device
	pe := d.pauseEvent
	d.state = StateProcessRunning
	d.mu.Unlock()

	// 用同一个缓冲区做输入和输出（METHOD_BUFFERED）
	buf := (*[udPacketSize]byte)(unsafe.Pointer(&pkt))[:]

	if _, err := dev.Ioctl(hyperdbgsdk.IoctlSendUserDebuggerCommands, buf, buf); err != nil {
		d.mu.Lock()
		d.state = StateProcessPaused
		d.mu.Unlock()
		return fmt.Errorf("Step: IOCTL failed: %w", err)
	}

	// 读取 Result
	outPkt := (*hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET)(unsafe.Pointer(&buf[0]))
	if outPkt.Result != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		d.mu.Lock()
		d.state = StateProcessPaused
		d.mu.Unlock()
		return fmt.Errorf("Step: kernel error (Result=0x%X)", outPkt.Result)
	}

	// 无 MessagePump（pe==nil）: 调用方自行轮询，直接返回。
	if pe == nil {
		return nil
	}

	return d.waitForPaused(dev, pe, 5*time.Second)
}

// waitForPaused 等 pauseEvent 信号（MessagePump 收到 PAUSED 包后触发），
// 超时后调 pauseProcess 强制暂停再等一次。Step/StepOut 共用。
func (d *Debugger) waitForPaused(dev *comm.Device, pe chan struct{}, timeout time.Duration) error {
	select {
	case <-pe:
		d.mu.Lock()
		d.state = StateProcessPaused
		d.mu.Unlock()
		return nil
	case <-time.After(timeout):
		// 超时恢复：#DB / 断点未触发，强制 Pause 重新暂停线程。
		if perr := pauseProcess(dev, d.processToken); perr != nil && perr != ErrAlreadyPaused {
			d.mu.Lock()
			d.state = StateProcessPaused
			d.mu.Unlock()
			return fmt.Errorf("Step: timeout, recovery Pause failed: %w", perr)
		}
		select {
		case <-pe:
			d.mu.Lock()
			d.state = StateProcessPaused
			d.mu.Unlock()
			return nil
		case <-time.After(5 * time.Second):
			d.mu.Lock()
			d.state = StateProcessPaused
			d.mu.Unlock()
			return fmt.Errorf("Step: timeout waiting for PAUSED after recovery Pause")
		}
	}
}

// detectCallAtPausedRip 反汇编当前暂停 RIP 处的指令，判断是否为 CALL。
// 返回 (isCall, callLength, err)。非 CALL 指令返回 (false, 0, nil)。
//
// 对应 C++ HyperDbgCheckWhetherTheCurrentInstructionIsCall（disassembler.cpp:755）。
func (d *Debugger) detectCallAtPausedRip() (bool, uint32, error) {
	d.mu.Lock()
	rip := d.pausedRIP
	pid := d.processPid
	dev := d.device
	d.mu.Unlock()
	if rip == 0 || dev == nil {
		return false, 0, fmt.Errorf("detectCallAtPausedRip: no paused RIP")
	}

	// 读取 15 字节（x86-64 最大指令长度）
	code, _, err := readmem.ReadMemory(dev, rip, pid, maxInstructionLength,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
	if err != nil || len(code) == 0 {
		return false, 0, fmt.Errorf("detectCallAtPausedRip: read memory: %w", err)
	}

	dis := misc.NewDisassembler()
	r, err := dis.Disassemble(misc.ModeLong64, rip, code)
	if err != nil {
		return false, 0, fmt.Errorf("detectCallAtPausedRip: disasm: %w", err)
	}
	if strings.HasPrefix(r.Text, "call") {
		return true, uint32(r.Length), nil
	}
	return false, 0, nil
}

// boolToUint64 将 bool 转为 UINT64（0/1），用于 OptionalParam2。
func boolToUint64(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

// StepOut 执行到当前函数返回（Execute till Return / Step Out）。
//
// 镜像 C++ SteppingStepOverForGu(FALSE)（steppings.cpp:155）→
// UdSendStepPacketToDebuggee(Token, ThreadId, STEP_OVER_FOR_GU)（ud.cpp:1255）。
// C++ 对 STEP_OVER_FOR_GU: IsCall=FALSE, CallLen=0（只有 STEP_OVER 才检查 IsCall）。
// 内核 UdStepInstructions 对 STEP_OVER_FOR_GU 走 UdRegularStepOverForGu 路径。
func (d *Debugger) StepOut() error {
	return d.stepImpl(uint64(hyperdbgsdk.DebuggerRemoteSteppingRequestStepOverForGu), false, 0)
}
