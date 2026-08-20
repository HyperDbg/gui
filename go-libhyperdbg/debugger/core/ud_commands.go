// Package core — ud_commands.go
//
// 用户态调试器命令（User-Debugger Commands）：通过 IOCTL_CODE_SEND_USER_DEBUGGER_COMMANDS
// 向内核发送 ReadRegisters / RegularStep 指令。
//
// 对应 C++ libhyperdbg 的 UdCommandReadRegisters / UdCommandStepGeneralCommand 路径。
package core

import (
	"encoding/binary"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
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

// DEBUGGEE_SHOW_ALL_REGISTERS 对应 C 宏 DEBUGGEE_SHOW_ALL_REGISTERS = 0xFFFFFFFF。
// 设置为 RegisterId 时，内核一次性返回全部通用寄存器 + 扩展寄存器。
const debuggeeshAllRegisters uint32 = 0xFFFFFFFF

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

	// 构建 OptionalBuffer: [DEBUGGEE_REGISTER_READ_DESCRIPTION][GUEST_REGS][GUEST_EXTRA_REGISTERS]
	optBufSize := regDescSize + guestRegsSize + extraRegsSize
	optBuf := make([]byte, optBufSize)
	desc := (*hyperdbgsdk.DEBUGGEE_REGISTER_READ_DESCRIPTION)(unsafe.Pointer(&optBuf[0]))
	desc.RegisterId = debuggeeshAllRegisters

	// 构建 IOCTL packet
	pkt := hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET{
		UdAction: hyperdbgsdk.DEBUGGER_UD_COMMAND_ACTION{
			ActionType:     hyperdbgsdk.DebuggerUdCommandActionTypeReadRegisters,
			OptionalParam1: uint64(debuggeeshAllRegisters), // C: RegDes->RegisterId
		},
		ProcessDebuggingDetailToken: d.processToken,
		// ApplyToAllPausedThreads=true：不指定具体线程，让内核对所有暂停线程
		// 应用命令。C CLI 传具体 ThreadId，但我们 StartProcess 后只有主线程暂停。
		ApplyToAllPausedThreads: true,
		WaitForEventCompletion:  true,
	}

	// 完整缓冲区: [DEBUGGER_UD_COMMAND_PACKET][OptionalBuffer]
	totalSize := udPacketSize + optBufSize
	buf := make([]byte, totalSize)
	pktBytes := (*[udPacketSize]byte)(unsafe.Pointer(&pkt))[:]
	copy(buf[0:], pktBytes)
	copy(buf[udPacketSize:], optBuf)

	if _, err = d.device.Ioctl(comm.IOCTL_CODE_SEND_USER_DEBUGGER_COMMANDS, buf, buf); err != nil {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: IOCTL failed: %w", err)
	}

	// 检查 packet Result
	outPkt := (*hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET)(unsafe.Pointer(&buf[0]))
	if outPkt.Result != DebuggerOperationWasSuccessful {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: kernel error (Result=0x%X)", outPkt.Result)
	}

	// 检查 DEBUGGEE_REGISTER_READ_DESCRIPTION.KernelStatus
	outDesc := (*hyperdbgsdk.DEBUGGEE_REGISTER_READ_DESCRIPTION)(unsafe.Pointer(&buf[udPacketSize]))
	if outDesc.KernelStatus != DebuggerOperationWasSuccessful {
		return regs, 0, 0, fmt.Errorf("ReadRegisters: kernel status error (0x%X)", outDesc.KernelStatus)
	}

	// GUEST_REGS 紧跟 DEBUGGEE_REGISTER_READ_DESCRIPTION 之后
	off := udPacketSize + regDescSize
	regs = *(*hyperdbgsdk.GUEST_REGS)(unsafe.Pointer(&buf[off]))

	// GUEST_EXTRA_REGISTERS 紧跟 GUEST_REGS 之后（含 RIP 和 RFLAGS）
	off += guestRegsSize
	extra := *(*hyperdbgsdk.GUEST_EXTRA_REGISTERS)(unsafe.Pointer(&buf[off]))
	rip = extra.RIP
	rflags = extra.RFLAGS

	return regs, rip, rflags, nil
}

// DEBUGGER_REMOTE_STEPPING_REQUEST 枚举值（RequestStructures.h:1044）。
// OptionalParam1 传给 UdStepInstructions 决定步入/步过。
const (
	stepRequestStepIn   uint64 = 0 // DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN
	stepRequestStepOver uint64 = 3 // DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER
)

// maxInstructionLength 是 x86-64 单条指令的最大长度（15 字节）。
const maxInstructionLength = 15

// Step 执行单步步入（RegularStep, STEP_IN）。内核设置 RFLAGS.TF 后让进程
// 执行一条指令，#DB 触发后再次暂停并发 DEBUGGEE_UD_PAUSED_PACKET。
//
// 与 C libhyperdbg 一致，使用 WaitForEventCompletion=false：IOCTL 立即返回，
// 然后通过 MessagePump 的 DEBUGGEE_UD_PAUSED_PACKET 等待单步完成。
//
// 超时恢复：若等待 PAUSED 包超时（典型场景：Step 4 走 AttachingHandleEntrypointInterception
// 路径不发 PAUSED，详见 Attaching.c:440），则调用 pauseProcess 强制重新暂停，
// 再等一次 PAUSED。这样即使某一步丢失 PAUSED 也能恢复，避免测试卡死。
func (d *Debugger) Step() error {
	return d.stepImpl(stepRequestStepIn, false, 0)
}

// StepOver 执行单步步过（RegularStep, STEP_OVER）。若当前指令是 CALL，
// 内核在下一指令设硬件断点；否则等同 StepIn。
//
// 对应 C++ UdSendStepPacketToDebuggee(StepType=STEP_OVER)（ud.cpp:1254）：
// 先反汇编当前指令判断是否为 CALL，若是则传 IsCall=true + CallLength，
// 否则 IsCall=false（内核走 TracingRegularStepInInstruction 路径）。
func (d *Debugger) StepOver() error {
	isCall, callLen, err := d.detectCallAtPausedRip()
	if err != nil {
		// 反汇编失败（zydis DLL 未加载、内存不可读等）时退化为 STEP_IN。
		// 不返回错误，避免 UI 步过按钮在边界场景下整体不可用。
		isCall, callLen = false, 0
	}
	return d.stepImpl(stepRequestStepOver, isCall, callLen)
}

// stepImpl 是 Step / StepOver 的共享实现。
//
// stepType: stepRequestStepIn 或 stepRequestStepOver。
// isCall / callLen: 仅 stepOver 时有意义，由 detectCallAtPausedRip 计算。
func (d *Debugger) stepImpl(stepType uint64, isCall bool, callLen uint32) error {
	if d.processToken == 0 {
		return fmt.Errorf("Step: no process attached")
	}

	d.mu.Lock()
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
		// WaitForEventCompletion=true: 内核在 VMX-root #DB handler 中
		// SynchronizationSetEvent 后才返回 IOCTL（Ud.c:698）。Go 用独立
		// 设备句柄做 MessagePump，不会死锁。这是比 C++ 的 false +
		// DbgWaitForUserResponse 更简洁的方案 — 不依赖 MessagePump 的
		// IRP 时序，从根本上消除了 PAUSED 包丢包竞态。
		WaitForEventCompletion: true,
	}
	dev := d.device
	d.mu.Unlock()

	// 用同一个缓冲区做输入和输出（METHOD_BUFFERED）
	buf := (*[udPacketSize]byte)(unsafe.Pointer(&pkt))[:]

	if _, err := dev.Ioctl(comm.IOCTL_CODE_SEND_USER_DEBUGGER_COMMANDS, buf, buf); err != nil {
		return fmt.Errorf("Step: IOCTL failed: %w", err)
	}

	// 读取 Result
	outPkt := (*hyperdbgsdk.DEBUGGER_UD_COMMAND_PACKET)(unsafe.Pointer(&buf[0]))
	if outPkt.Result != DebuggerOperationWasSuccessful {
		return fmt.Errorf("Step: kernel error (Result=0x%X)", outPkt.Result)
	}

	// WaitForEventCompletion=true 时 IOCTL 返回即表示 step 完成，线程已暂停。
	// PAUSED 包通过 MessagePump 异步到达，更新 pausedRIP（用于下一次
	// detectCallAtPausedRip / OnPaused 回调刷新 UI）。
	d.mu.Lock()
	d.state = StateProcessPaused
	d.mu.Unlock()
	return nil
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
// 实现：读 [RSP] 得到返回地址 → 在返回地址设临时断点（RemoveAfterHit=true）
// → Continue → 等断点命中的 PAUSED 包。断点命中后内核自动移除。
//
// 对应 OllyDbg 的 Ctrl+F9 / x64dbg 的 "执行到返回"。
func (d *Debugger) StepOut() error {
	if d.processToken == 0 {
		return fmt.Errorf("StepOut: no process attached")
	}

	// 1. 读寄存器拿 RSP
	regs, _, _, err := d.ReadRegisters()
	if err != nil {
		return fmt.Errorf("StepOut: read registers: %w", err)
	}
	rsp := regs.Rsp
	if rsp == 0 {
		return fmt.Errorf("StepOut: RSP=0")
	}

	// 2. 读 [RSP] 8 字节 = 返回地址
	d.mu.Lock()
	dev := d.device
	pid := d.processPid
	pe := d.pauseEvent
	d.mu.Unlock()
	retAddr, _, err := readmem.ReadMemory(dev, rsp, pid, 8,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
	if err != nil || len(retAddr) < 8 {
		return fmt.Errorf("StepOut: read return address: %w", err)
	}
	returnAddress := binary.LittleEndian.Uint64(retAddr[:8])
	if returnAddress == 0 {
		return fmt.Errorf("StepOut: return address is 0 (not in a call frame?)")
	}

	// 3. 在返回地址设临时断点（RemoveAfterHit=true）
	tag, err := d.bpSetTemporary(returnAddress)
	if err != nil {
		return fmt.Errorf("StepOut: set temp breakpoint at 0x%X: %w", returnAddress, err)
	}
	// 保险：即使命中后内核自动移除，defer 也尝试 clear（已移除则 no-op）
	defer func() { _ = d.BreakpointClear(tag) }()

	// 4. Drain stale pauseEvent, Continue, wait for PAUSED
	d.mu.Lock()
	if d.pauseEvent != nil {
		select {
		case <-d.pauseEvent:
		default:
		}
	}
	d.mu.Unlock()

	if err := continueProcess(dev, d.processToken); err != nil {
		return fmt.Errorf("StepOut: continue: %w", err)
	}
	d.mu.Lock()
	d.state = StateProcessRunning
	d.mu.Unlock()

	if pe == nil {
		return nil // no MessagePump — caller polls
	}

	select {
	case <-pe:
		d.mu.Lock()
		d.state = StateProcessPaused
		d.mu.Unlock()
		return nil
	case <-time.After(5 * time.Second):
		// 超时恢复：强制 Pause
		if perr := pauseProcess(dev, d.processToken); perr != nil && perr != ErrAlreadyPaused {
			return fmt.Errorf("StepOut: timeout, recovery Pause failed: %w", perr)
		}
		select {
		case <-pe:
			d.mu.Lock()
			d.state = StateProcessPaused
			d.mu.Unlock()
			return nil
		case <-time.After(5 * time.Second):
			return fmt.Errorf("StepOut: timeout waiting for return breakpoint")
		}
	}
}

// bpSetTemporary 在 addr 设一个命中后自动移除的软件断点，返回 tag。
func (d *Debugger) bpSetTemporary(addr uint64) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("bpSetTemporary: VMM not loaded")
	}
	pkt := hyperdbgsdk.DEBUGGEE_BP_PACKET{
		Address:           addr,
		Pid:               0,
		Tid:               0,
		Core:              0xFFFFFFFF,
		RemoveAfterHit:    true, // 关键：命中后内核自动移除
		CheckForCallbacks: true,
	}
	reqBuf := byteslice.FromStruct(&pkt)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_SET_BREAKPOINT_USER_DEBUGGER, reqBuf, reqBuf); err != nil {
		return 0, fmt.Errorf("bpSetTemporary: IOCTL failed: %w", err)
	}
	pkt = *byteslice.ToStruct[hyperdbgsdk.DEBUGGEE_BP_PACKET](reqBuf)
	return uint64(pkt.Result), nil
}
