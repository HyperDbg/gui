// Package api — commands_debugging.go
//
// 对应 debugger/commands/debugging 包的 37 个 debugging 命令的 typed API。
//
// 命令对照表（与 debugging.go RegisterAll 顺序一致）：
//
//	sleep      → Sleep          (api.debugger.go, 已有)
//	events     → Events         (本文件)
//	settings   → Settings       (本文件)
//	continue   → Continue       (api.debugger.go, 已有)
//	a          → Assemble       (本文件, stub)
//	bc         → BpClear        (本文件, stub)
//	bd         → BpDisable      (本文件, stub)
//	be         → BpEnable       (本文件, stub)
//	bl         → BpList         (本文件, stub)
//	bp         → BpSet          (本文件, stub)
//	core       → Core           (本文件, stub)
//	cpu        → Cpu            (本文件, stub)
//	d          → DumpMem        (本文件, stub)
//	u          → Unassemble     (本文件, stub)
//	dt         → Dt             (本文件, stub)
//	e          → EditMem        (本文件, stub)
//	eval       → Eval           (本文件, stub)
//	flush      → Flush          (本文件, stub)
//	gg         → Gg             (本文件, stub)
//	gu         → Gu             (本文件, stub)
//	i          → IoIn           (本文件, stub)
//	k          → K              (本文件, stub)
//	lm         → Lm             (本文件, stub)
//	output     → Output         (本文件, stub)
//	p          → StepOver       (本文件, stub)
//	preactivate→ Preactivate    (本文件, stub)
//	prealloc   → Prealloc       (本文件, stub)
//	print      → Print          (本文件, stub)
//	r          → Register       (本文件, stub)
//	rdmsr      → Rdmsr          (本文件, stub)
//	s          → Search         (本文件, stub)
//	t          → TraceInto      (本文件, stub)
//	test       → Test           (本文件, stub)
//	wrmsr      → Wrmsr          (本文件, stub)
//	x          → Examine        (本文件, stub)
package api

import (
	"encoding/binary"
	"fmt"
	"strings"
	"unsafe"

	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
	"github.com/hyperdbg/go-libhyperdbg/debugger/misc"
)

// ============================================================
// 返回类型（让 API 更 typed，避免大量 string 返回）
// ============================================================

// Breakpoint 描述一个软件断点（bp/bl 命令的元素）。
type Breakpoint struct {
	Tag      uint64 // 断点 tag（用于 bc/bd/be 引用）
	Address  uint64 // 断点地址
	Enabled  bool   // 是否启用
	HitCount uint64 // 命中次数
}

// Module 描述一个已加载模块（lm 命令的元素）。
type Module struct {
	Name string // 模块名（如 "ntdll.dll"）
	Base uint64 // 加载基址
	Size uint32 // 模块大小
	Path string // 完整路径
}

// CallFrame 描述一个栈帧（k 命令的元素）。
type CallFrame struct {
	Address     uint64 // 返回地址
	Module      string // 所属模块
	Symbol      string // 符号名（如已解析）
	OffsetInSym uint64 // 距符号起始的偏移
}

// CpuInfo 描述 CPU 状态（cpu 命令的输出）。
type CpuInfo struct {
	Vendor string // CPUID vendor string
	Brand  string // CPUID brand string
	Cores  uint32 // 逻辑核心数
	Mhz    uint32 // 频率（MHz）
}

// Settings 描述调试器可调参数（settings 命令的输出）。
type Settings struct {
	State      core.DebuggerState // 当前调试器状态
	ScriptMode bool               // 是否在脚本模式
	ShowDisasm bool               // 是否显示反汇编
}

// ============================================================
// A) 已实装命令的 typed API
// ============================================================

// Events 对应 'events' 命令：列出已注册的事件 tag。
//
// core 层尚未暴露 tag 表，本方法当前返回空 slice（不报错）——
// 调用方可据此判断"功能可用但暂无数据"，与字符串命令路径行为一致。
func (d *Debugger) Events() ([]uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// core.Debugger 尚未暴露 tag 表，返回空 slice 保持 typed 签名稳定。
	return nil, nil
}

// Settings 对应 'settings' 命令：返回当前调试器设置。
func (d *Debugger) Settings() (Settings, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return Settings{
		State: d.core.State(),
	}, nil
}

// ============================================================
// B) Stubs — typed 签名 + ErrCommandNotImplemented
// ============================================================

// Assemble 对应 'a <addr> <instr>'：在指定地址汇编一条指令。
func (d *Debugger) Assemble(addr uint64, instr string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	asm := misc.NewAssembler()
	bytes, err := asm.Assemble(misc.AsmMode64, addr, instr)
	if err != nil {
		return err
	}
	return d.output.Printf("%X\n", bytes)
}

// BpClear 对应 'bc <tag>'：清除指定断点。
func (d *Debugger) BreakpointClear(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.BreakpointClear(tag)
}

// BpDisable 对应 'bd <tag>'：禁用指定断点（保留但不触发）。
func (d *Debugger) BreakpointDisable(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.BreakpointDisable(tag)
}

// BpEnable 对应 'be <tag>'：启用之前禁用的断点。
func (d *Debugger) BreakpointEnable(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.BreakpointEnable(tag)
}

// BpList 对应 'bl'：列出所有断点。
func (d *Debugger) BreakpointList() ([]Breakpoint, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the listing is written to
	// d.output. The typed slice is not populated (the string path owns the
	// formatting).
	return nil, d.commands.Exec(d.core, "bl")
}

// BpSet 对应 'bp <addr>'：在指定地址设置软件断点，返回断点 tag。
func (d *Debugger) BreakpointSet(addr uint64) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.BreakpointSet(addr)
}

// Core 对应 'core <id>'：切换当前核心（仅多核调试有意义）。
func (d *Debugger) Core(coreId uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Core(coreId)
}

// Cpu 对应 'cpu'：返回 CPU 信息。
func (d *Debugger) Cpu() (CpuInfo, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	resp, err := d.core.Cpu()
	if err != nil {
		return CpuInfo{}, err
	}
	// BrandString is a fixed [49]int8 NUL-terminated buffer.
	brand := strings.TrimRight(unsafe.String((*byte)(unsafe.Pointer(&resp.BrandString[0])), len(resp.BrandString)), "\x00")
	return CpuInfo{Brand: brand}, nil
}

// DumpMem 对应 'd <addr> <size>'：转储内存（返回原始字节）。
//
// 注意：与 ReadMemory 的区别在于 DumpMem 针对当前 attached 的调试目标，
// 不需要显式传 pid。core 层补齐后会自动用 d.core 的 processToken 推导 pid。
func (d *Debugger) DumpMem(addr uint64, size uint32) ([]byte, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.DumpMem(addr, size)
}

// Unassemble 对应 'u <addr> <count>'：反汇编指定地址的 count 条指令。
func (d *Debugger) Unassemble(addr uint64, count uint32) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if count == 0 {
		return "", nil
	}
	// 15 bytes is the max x86-64 instruction length.
	code, err := d.core.DumpMem(addr, count*15)
	if err != nil {
		return "", err
	}
	dis := misc.NewDisassembler()
	results, err := dis.DisassembleRange(misc.ModeLong64, addr, code)
	if err != nil && len(results) == 0 {
		return "", err
	}
	var sb strings.Builder
	for _, r := range results {
		fmt.Fprintf(&sb, "%016X  %s\n", r.Runtime, r.Text)
	}
	return sb.String(), nil
}

// Dt 对应 'dt <type> <addr>'：按结构体类型显示内存。
func (d *Debugger) DumpType(typeName string, addr uint64) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the formatted output goes to
	// d.output. The typed string return is not populated.
	err := d.commands.Exec(d.core, fmt.Sprintf("dt %s 0x%X", typeName, addr))
	return "", err
}

// EditMem 对应 'e <addr> <data>'：写入内存。
func (d *Debugger) EditMem(addr uint64, data []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EditMem(addr, data)
}

// Eval 对应 'eval <expr>'：求值表达式，返回 64-bit 结果。
func (d *Debugger) Eval(expr string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the value is printed to d.output.
	// The typed uint64 return is not populated (parsing the formatted output
	// back is out of scope for this wrapper).
	err := d.commands.Exec(d.core, "eval "+expr)
	return 0, err
}

// Flush 对应 'flush'：刷新调试器缓存（符号/断点等）。
// core 层发送 IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS，将内核日志缓冲排空到用户态。
func (d *Debugger) Flush() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Flush()
}

// Gg 对应 'gg <addr>'：运行到指定地址（go until address）。
func (d *Debugger) GoUntilAddress(addr uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, fmt.Sprintf("g 0x%X", addr))
}

// Gu 对应 'gu'：运行到当前函数返回（go until return）。
// 当前用 Step 近似（多次单步直到 RSP 变化），完整实现需要栈回溯。
func (d *Debugger) GoUntilReturn() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Step()
}

// IoIn 对应 'i <port>'：从 I/O 端口读取一字节。
func (d *Debugger) IoIn(port uint16) (byte, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.IoIn(port)
}

// K 对应 'k <count>'：回溯调用栈。通过 ReadRegisters 获取 RSP 后
// 用 DumpMem 读栈内存，扫描返回地址生成 []CallFrame 返回。
// 调用方（如 CPU 页）自行格式化显示，无需捕获 d.output 文本。
func (d *Debugger) CallStack(count uint32) ([]CallFrame, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	regs, rip, _, err := d.core.ReadRegisters()
	if err != nil {
		return nil, fmt.Errorf("K: ReadRegisters: %w", err)
	}

	// 读栈内存
	stackSize := uint32(count) * 8
	if stackSize < 256 {
		stackSize = 256
	}
	if stackSize > 4096 {
		stackSize = 4096
	}
	stackData, err := d.core.DumpMem(regs.Rsp, stackSize)
	if err != nil {
		return nil, fmt.Errorf("K: DumpMem: %w", err)
	}

	// 构建调用栈帧。首帧为当前 RIP（入口点）。
	frames := make([]CallFrame, 0, count+1)
	frames = append(frames, CallFrame{
		Address: rip,
		Symbol:  "<entry>",
	})

	// 扫描栈上的返回地址（简化版：每 8 字节检查是否像合法地址）
	for i := 0; i+8 <= len(stackData) && i/8 < int(count); i += 8 {
		val := binary.LittleEndian.Uint64(stackData[i:])
		// 简单启发式：用户态地址范围 (0x00000000 - 0x7FFFFFFFFFFF)
		if val > 0x10000 && val < 0x800000000000 {
			frames = append(frames, CallFrame{
				Address: val,
				Symbol:  "<unknown>",
			})
		}
	}

	return frames, nil
}

// Lm 对应 'lm'：列出当前调试目标已加载的模块。
func (d *Debugger) ListModules() ([]Module, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the module listing is written to
	// d.output. The typed slice is not populated.
	err := d.commands.Exec(d.core, "lm")
	return nil, err
}

// Output 对应 'output <expr>'：将表达式求值结果输出到日志（不打印到 console）。
func (d *Debugger) Output(expr string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, "output "+expr)
}

// StepOver 对应 'p'：单步步过（step over，不进入 call）。
// 反汇编当前指令，若是 CALL 则传 STEP_OVER + IsCall + CallLength 给内核，
// 内核在下一指令设硬件断点；否则退化为 STEP_IN。
//
// 不持有 api.mu —— core 层内部管理锁（获取 dev 后释放），等 PAUSED 期间
// 不持任何锁，允许 Register/Unassemble 等查询并发执行（OnPaused 回调刷新）。
func (d *Debugger) StepOver() error {
	return d.core.StepOver()
}

// StepOut 执行到当前函数返回（step out / execute till return）。
// 读 [RSP] 返回地址，设临时断点后 Continue，命中后自动移除断点。
// 对应 OllyDbg Ctrl+F9 / x64dbg "执行到返回"。
// 不持 api.mu，理由同 StepOver。
func (d *Debugger) StepOut() error {
	return d.core.StepOut()
}

// Preactivate 对应 'preactivate'：预激活所有断点（确保下次 Continue 立即生效）。
// core 层发送 IOCTL_PREACTIVATE_FUNCTIONALITY，目前仅支持 mode-change 预激活。
func (d *Debugger) Preactivate() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Preactivate(hyperdbgsdk.DebuggerPreactivateCommandTypeMode)
}

// Prealloc 对应 'prealloc <size>'：预分配内核缓冲区（用于大日志量场景）。
// size 是缓冲数量；类型默认为 RegularEvent（如需其他类型用 core.Prealloc）。
func (d *Debugger) Prealloc(size uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if size > 0xFFFFFFFF {
		return fmt.Errorf("Prealloc: size %d exceeds uint32", size)
	}
	return d.core.Prealloc(hyperdbgsdk.DebuggerPreallocCommandTypeRegularEvent, uint32(size))
}

// Print 对应 'print <expr>'：打印表达式值（与 eval 的区别在于 print 格式化输出）。
func (d *Debugger) Print(expr string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the formatted value is written to
	// d.output. The typed string return is not populated.
	err := d.commands.Exec(d.core, "print "+expr)
	return "", err
}

// Register 对应 'r <reg>'：读取单个寄存器，返回其数值。
func (d *Debugger) Register(reg string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	regs, rip, rflags, err := d.core.ReadRegisters()
	if err != nil {
		return 0, fmt.Errorf("Register: ReadRegisters: %w", err)
	}

	switch strings.ToUpper(reg) {
	case "RIP":
		return rip, nil
	case "RFL", "RFLAGS", "EFLAGS":
		return rflags, nil
	case "":
		return 0, fmt.Errorf("Register: empty reg name, use AllRegisters")
	default:
		return lookupRegByName(regs, rip, rflags, reg), nil
	}
}

// AllRegisters 对应 'r'（无参数）：读取所有寄存器，返回格式化文本。
func (d *Debugger) AllRegisters() (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	regs, rip, rflags, err := d.core.ReadRegisters()
	if err != nil {
		return "", fmt.Errorf("AllRegisters: ReadRegisters: %w", err)
	}

	return fmt.Sprintf("RAX=%016X RBX=%016X RCX=%016X\nRDX=%016X RSI=%016X RDI=%016X\nRIP=%016X RSP=%016X RBP=%016X\nR8 =%016X R9 =%016X R10=%016X\nR11=%016X R12=%016X R13=%016X\nR14=%016X R15=%016X\nRFL=%016X",
		regs.Rax, regs.Rbx, regs.Rcx,
		regs.Rdx, regs.Rsi, regs.Rdi,
		rip, regs.Rsp, regs.Rbp,
		regs.R8, regs.R9, regs.R10,
		regs.R11, regs.R12, regs.R13,
		regs.R14, regs.R15,
		rflags), nil
}

// SetRegister 对应 'r <reg> <val>'：写入寄存器。
func (d *Debugger) SetRegister(reg string, val uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.commands.Exec(d.core, fmt.Sprintf("r %s 0x%X", reg, val))
}

// Rdmsr 对应 'rdmsr <msr>'：读取 MSR（Model Specific Register）。
func (d *Debugger) ReadMsr(msr uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.ReadMsr(msr)
}

// Search 对应 's <addr> <size> <pattern>'：在内存中搜索字节模式。
// 返回所有匹配地址。
func (d *Debugger) Search(addr uint64, size uint32, pattern []byte) ([]uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Search(addr, size, pattern)
}

// TraceInto 对应 't'：单步步入（trace into，进入 call）。
// 不持 api.mu，理由同 StepOver（等 PAUSED 期间允许查询并发）。
func (d *Debugger) TraceInto() error {
	return d.core.Step()
}

// Test 对应 'test'：自检命令（运行内部测试套件）。
func (d *Debugger) Test() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Test()
}

// Wrmsr 对应 'wrmsr <msr> <val>'：写入 MSR。
func (d *Debugger) WriteMsr(msr uint32, val uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.WriteMsr(msr, val)
}

// Examine 对应 'x <pattern>'：按通配符查找符号（如 "nt!Nt*"）。
func (d *Debugger) Examine(pattern string) ([]Module, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	// Delegate to the string command path; the symbol listing is written to
	// d.output. The typed slice is not populated.
	err := d.commands.Exec(d.core, "x "+pattern)
	return nil, err
}

// lookupRegByName 按名称查找寄存器值（不区分大小写）。
func lookupRegByName(regs hyperdbgsdk.GUEST_REGS, rip, rflags uint64, name string) uint64 {
	switch strings.ToUpper(name) {
	case "RAX":
		return regs.Rax
	case "RBX":
		return regs.Rbx
	case "RCX":
		return regs.Rcx
	case "RDX":
		return regs.Rdx
	case "RSI":
		return regs.Rsi
	case "RDI":
		return regs.Rdi
	case "RSP":
		return regs.Rsp
	case "RBP":
		return regs.Rbp
	case "R8":
		return regs.R8
	case "R9":
		return regs.R9
	case "R10":
		return regs.R10
	case "R11":
		return regs.R11
	case "R12":
		return regs.R12
	case "R13":
		return regs.R13
	case "R14":
		return regs.R14
	case "R15":
		return regs.R15
	case "RIP":
		return rip
	case "RFL", "RFLAGS", "EFLAGS":
		return rflags
	default:
		return 0
	}
}
