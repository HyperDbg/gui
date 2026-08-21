// Package api — commands_extension.go
//
// 对应 debugger/commands/extension 包的 35 个 extension 命令（"!" 前缀）的 typed API。
//
// 命令对照表（与 extension.go RegisterAll 顺序一致）：
//
//	!apic      → Apic             (本文件, stub)
//	!cpuid     → CpuidHook        (本文件, stub)
//	!crwrite   → CrwriteHook      (本文件, stub)
//	!dr        → DrHook           (本文件, stub)
//	!epthook   → EptHook/...      (api.debugger.go, 已有)
//	!epthook2  → EptHook2         (本文件, stub)
//	!exception → ExceptionHook    (本文件, stub)
//	!hide      → Hide             (本文件, stub)
//	!idt       → Idt              (本文件, stub)
//	!interrupt → InterruptHook    (本文件, stub)
//	!ioapic    → Ioapic           (本文件, stub)
//	!ioin      → IoInHook         (本文件, stub)
//	!ioout     → IoOutHook        (本文件, stub)
//	!lbr       → Lbr              (本文件, stub)
//	!lbrdump   → LbrDump          (本文件, stub)
//	!measure   → Measure          (本文件, stub)
//	!mode      → ModeHook         (本文件, stub)
//	!monitor   → MonitorReadForProcess (本文件, 已实装) + MonitorWrite/MonitorExec (stub)
//	!msrread   → MsrReadHook      (本文件, stub)
//	!msrwrite  → MsrWriteHook     (本文件, stub)
//	!pa2va     → Pa2Va            (本文件, stub)
//	!pcicam    → PciCam           (本文件, stub)
//	!pcitree   → PciTree          (本文件, stub)
//	!pmc       → Pmc              (本文件, stub)
//	!pt        → Pt               (本文件, stub)
//	!pte       → Pte              (本文件, stub)
//	!rev       → Rev              (本文件, stub)
//	!smi       → Smi              (本文件, stub)
//	!syscall   → SyscallHook      (本文件, stub)
//	!sysret    → SysretHook       (本文件, stub)
//	!trace     → Trace            (本文件, stub)
//	!track     → Track            (本文件, stub)
//	!tsc       → Tsc              (本文件, stub)
//	!unhide    → Unhide           (本文件, stub)
//	!va2pa     → Va2Pa            (本文件, stub)
//	!vmcall    → VmcallHook       (本文件, stub)
//	!xsetbv    → XsetbvHook       (本文件, stub)
//
// hook 类命令（CpuidHook/SysretHook/...）的签名约定：
//
//	func (d *Debugger) XxxHook(callbackSrc string) (uint64, error)
//	  - callbackSrc 是 Go AST hook 源码（与 EptHook 相同格式）
//	  - 返回 hook tag，用于后续 Unhook
//
// 事件管理通用方法（对应 events 命令的子命令 c/d/e）：
//
//	ClearEvent / DisableEvent / EnableEvent — 在本文件末尾
package api

import (
	"fmt"

	"github.com/ddkwork/hyperdbgsdk"
)


// ============================================================
// A) 已实装命令的 typed API
// ============================================================

// MonitorReadForProcess 对应 '!monitor r <addr> <size> <pid>'：监控指定进程
// 对 [addrStart, addrEnd) 的读取访问。callbackSrc 是 Go AST hook 源码。
// 返回 hook tag。
//
// 这是少数几个 core 层已实装的 ! 命令之一（core.MonitorReadForProcess）。
// MonitorWrite 和 MonitorExec 见下方。
func (d *Debugger) MonitorReadForProcess(addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.MonitorReadForProcess(addrStart, addrEnd, pid, callbackSrc)
}

// ============================================================
// B) hook 类命令 — 委托 core 层 registerHookEvent
// ============================================================

// CpuidHook 对应 '!cpuid'：挂钩所有 CPUID 指令。
// 每次目标执行 CPUID 时触发 callbackSrc 编译的 hook。
// （如需只挂钩特定 EAX，使用 core.CpuidHook 的 hasEax 重载。）
func (d *Debugger) CpuidHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.CpuidHook(false, 0, callbackSrc)
}

// CrwriteHook 对应 '!crwrite <cr>'：挂钩控制寄存器写入。
// cr 指定要监控的寄存器（0/4）。
func (d *Debugger) CrwriteHook(cr uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.CrwriteHook(cr, 0, callbackSrc)
}

// DrHook 对应 '!dr'：挂钩调试寄存器访问。
func (d *Debugger) DrHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.DrHook(callbackSrc)
}

// EptHook2 对应 '!epthook2 <addr>'：EPT hook 的变体（detour 模式）。
// 与 EptHook 的区别在于 hook2 不触发 #VMEXIT，性能更好但有限制。
// 本签名不传 pid，故应用于所有进程；用户态地址需用 core.EptHook2
// 显式传 pid。
func (d *Debugger) EptHook2(hookAddress uint64, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EptHook2(hookAddress, hyperdbgsdk.DebuggerEventApplyToAllProcesses, callbackSrc)
}

// ExceptionHook 对应 '!exception <vector>'：挂钩异常（#PF/#GP/#UD 等）。
// vector 是异常向量号（0-31）。
func (d *Debugger) ExceptionHook(vector uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.ExceptionHook(vector, callbackSrc)
}

// InterruptHook 对应 '!interrupt <vector>'：挂钩硬件中断。
// vector 是中断向量号（32-255）。
func (d *Debugger) InterruptHook(vector uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.InterruptHook(vector, callbackSrc)
}

// IoInHook 对应 '!ioin <port>'：挂钩 I/O 端口读（IN 指令）。
func (d *Debugger) IoInHook(port uint16, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.IoInHook(port, callbackSrc)
}

// IoOutHook 对应 '!ioout <port>'：挂钩 I/O 端口写（OUT 指令）。
func (d *Debugger) IoOutHook(port uint16, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.IoOutHook(port, callbackSrc)
}

// ModeHook 对应 '!mode'：挂钩执行模式切换（user↔kernel）。
// 本签名不传 mode/pid，默认拦截 user+kernel 两种模式切换、应用于所有进程。
// 如需精细控制，使用 core.ModeHook。
func (d *Debugger) ModeHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.ModeHook(hyperdbgsdk.DebuggerEventModeTypeUserModeAndKernelMode, hyperdbgsdk.DebuggerEventApplyToAllProcesses, callbackSrc)
}

// MonitorWrite 对应 '!monitor w <addr> <size> <pid>'：监控内存写入。
func (d *Debugger) MonitorWrite(addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.MonitorWrite(addrStart, addrEnd, pid, callbackSrc)
}

// MonitorExec 对应 '!monitor e <addr> <pid>'：监控内存执行。
// 本签名只传单个 addr，core 层将其作为单字节范围 [addr, addr] 注册。
func (d *Debugger) MonitorExec(addr uint64, pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.MonitorExec(addr, addr, pid, callbackSrc)
}

// MsrReadHook 对应 '!msrread <msr>'：挂钩 MSR 读取（RDMSR 指令）。
// msr=0 表示挂钩所有 MSR；非零值只挂钩指定 MSR。
func (d *Debugger) MsrReadHook(msr uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.MsrReadHook(msr, callbackSrc)
}

// MsrWriteHook 对应 '!msrwrite <msr>'：挂钩 MSR 写入（WRMSR 指令）。
func (d *Debugger) MsrWriteHook(msr uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.MsrWriteHook(msr, callbackSrc)
}

// SyscallHook 对应 '!syscall'：挂钩所有 SYSCALL 指令入口。
//
// 内核驱动通过 SYSCALL_HOOK_EFER_SYSCALL（设置 EFER.SCE + MSR LSTAR）实现。
// callbackSrc 是 Go AST hook 源码，返回 hook tag。
// 如需只挂钩特定 syscall 号，使用 core.SyscallHook 传 syscallNumber。
func (d *Debugger) SyscallHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.SyscallHook(0xFFFFFFFF, callbackSrc)
}

// SysretHook 对应 '!sysret'：挂钩所有 SYSRET 指令（系统调用返回）。
//
// 内核驱动通过 SYSCALL_HOOK_EFER_SYSRET（与 SyscallHook 共用 EFER.SCE 配置）实现。
// 用户场景：用于在系统调用返回到用户态时触发回调，比 EptHook Nt* 包装函数
// 更底层、更难被反调试检测。
//
// 注意：此变体将 ProcessId 设为 ALL_PROCESSES，会对系统中每个进程的每个
// SYSRET 都触发 Go 回调（即使回调内立刻按 pid 过滤返回）。系统级 syscall
// 流量 5000+/sec 会导致 VMX-root 过载，可能使关键系统进程（svchost 等）
// 触发 __fastfail → CRITICAL_PROCESS_DIED 0xEF BSOD。如只关心单个进程，
// 必须使用 SysretHookForProcess。
func (d *Debugger) SysretHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.SysretHook(0xFFFFFFFF, callbackSrc)
}

// SysretHookForProcess 对应 '!sysret'，但将 hook 限定在单个进程上。
//
// 与 SysretHook 的区别：把真实 pid 传给内核，内核在 VmmCallbackTriggerEvents
// 中按 pid 过滤（Debugger.c:1215），对非目标进程的 SYSRET 直接跳过 Go 回调，
// 不进入 AST 解释器。EFER hook 仍对全系统 syscall/sysret 触发 VM exit（硬件
// 级无法避免），但跳过解释器执行可大幅降低 VMX-root 时间，避免 BSOD。
//
// 推荐在所有"只关心单个目标进程"的场景使用本方法替代 SysretHook。
func (d *Debugger) SysretHookForProcess(pid uint32, callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.SysretHookForProcess(pid, 0xFFFFFFFF, callbackSrc)
}

// VmcallHook 对应 '!vmcall'：挂钩 VMCALL 指令（hypercall）。
func (d *Debugger) VmcallHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.VmcallHook(callbackSrc)
}

// XsetbvHook 对应 '!xsetbv'：挂钩 XSETBV 指令（设置 XCR0）。
func (d *Debugger) XsetbvHook(callbackSrc string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.XsetbvHook(callbackSrc)
}

// ============================================================
// C) Stubs — 控制/查询类命令
// ============================================================

// Apic 对应 '!apic'：显示本地 APIC 状态。
func (d *Debugger) Apic() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.core.Apic()
	return err
}

// Hide 对应 '!hide'：将 VMM 隐藏（让目标检测不到 hypervisor）。
// 用于绕过基于 CPUID hypervisor bit 的反调试检测。
// 本签名不传 pid，core 层传 0（内核解释为当前进程）；如需指定进程
// 使用 core.Hide。
func (d *Debugger) Hide() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Hide(0)
}

// Idt 对应 '!idt <vector>'：读取并返回指定向量的 IDT 表项地址。
// core 层读取全部 256 个表项后取 [vector]。
func (d *Debugger) InterruptDescriptorTable(vector uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if vector > 255 {
		return 0, fmt.Errorf("Idt: vector %d out of range (0-255)", vector)
	}
	pkt, err := d.core.InterruptDescriptorTable()
	if err != nil {
		return 0, err
	}
	return pkt.IdtEntry[vector], nil
}

// Ioapic 对应 '!ioapic'：显示 I/O APIC 状态。
func (d *Debugger) IoApic() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.core.IoApic()
	return err
}

// Lbr 对应 '!lbr'：启用 LBR（Last Branch Record）分支记录。
func (d *Debugger) LastBranchRecord() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.LastBranchRecord()
}

// LbrDump 对应 '!lbrdump'：转储 LBR 记录的分支。
func (d *Debugger) LbrDump() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.LbrDump()
}

// Measure 对应 '!measure'：测量 VM exit 开销（用于反调试 timing 检测评估）。
func (d *Debugger) Measure() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Measure()
}

// Pa2Va 对应 '!pa2va <pa> <pid>'：物理地址转虚拟地址。
func (d *Debugger) PhysicalToVirtual(pa uint64, pid uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.PhysicalToVirtual(pa, pid)
}

// PciCam 对应 '!pcicam'：通过 PCI CAM（Configuration Access Method）读取配置。
func (d *Debugger) PciCam(bus, device, function, offset uint32) (uint32, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.PciCam(bus, device, function, offset)
}

// PciTree 对应 '!pcitree'：打印 PCI 设备树。
func (d *Debugger) PciTree() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.core.PciTree()
	return err
}

// Pmc 对应 '!pmc'：读取 PMC（Performance Monitoring Counter）。
func (d *Debugger) PerfCounter(counter uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.PerfCounter(counter)
}

// Pt 对应 '!pt'：Intel PT（Processor Trace）配置。
func (d *Debugger) ProcessorTrace() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.ProcessorTrace()
}

// Pte 对应 '!pte <va>'：读取指定虚拟地址的页表项链（PML4E/PDPTE/PDE/PTE）。
// 返回 PTE 的值（最末一级页表项）。pid 默认为 0（当前进程）。
func (d *Debugger) PageTableEntry(va uint64) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	details, err := d.core.PageTableEntry(va, 0)
	if err != nil {
		return 0, err
	}
	// PteValue is the last-level PTE value (see DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS).
	return details.PteValue, nil
}

// Rev 对应 '!rev'：触发 reversing-machine 内存重建（pattern/reconstruct）。
// 本签名不传 pid/mode/type，core 层使用默认值；返回 0（该 IOCTL 无数值返回）。
// 如需精细控制，使用 core.Rev。
func (d *Debugger) Revision() (uint32, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	err := d.core.Revision(0,
		hyperdbgsdk.ReversingMachineReconstructMemoryModeUnknown,
		hyperdbgsdk.ReversingMachineReconstructMemoryTypeUnknown)
	return 0, err
}

// Smi 对应 '!smi'：触发 SMI（System Management Interrupt）。
func (d *Debugger) SmiInterrupt() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	_, err := d.core.SmiInterrupt()
	return err
}

// Trace 对应 '!trace'：启用 Intel PT 追踪（与 Pt 的区别在于 Trace 是控制开关）。
func (d *Debugger) Trace() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Trace()
}

// Track 对应 '!track'：追踪内存访问模式。
func (d *Debugger) Track() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Track()
}

// Tsc 对应 '!tsc'：读取 TSC（Time Stamp Counter）。
func (d *Debugger) TimeStampCounter() (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.TimeStampCounter()
}

// Unhide 对应 '!unhide'：取消 !hide 的隐藏（让 hypervisor 重新可见）。
func (d *Debugger) Unhide() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.Unhide()
}

// Va2Pa 对应 '!va2pa <va> <pid>'：虚拟地址转物理地址。
func (d *Debugger) VirtualToPhysical(va uint64, pid uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.VirtualToPhysical(va, pid)
}

// ============================================================
// D) 事件管理通用方法（对应 events 命令的子命令）
// ============================================================
//
// HyperDbg 的 events 命令支持子命令：c(d)/e(nable)/d(isable)。
// 这些是事件管理的通用 API，对所有 hook 类型（EptHook/SysretHook/...）都适用。
// 放在本文件因为 hook 是主要使用场景；调用方通常在 extension 命令后用这些方法管理。

// ClearEvent 对应 'events c <tag>'：清除（移除）指定 tag 的事件。
// 对所有 hook 类型通用（EptHook/SyscallHook/MonitorRead/...）。
func (d *Debugger) ClearEvent(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.ClearEvent(tag)
}

// DisableEvent 对应 'events d <tag>'：临时禁用指定 tag 的事件（保留配置）。
func (d *Debugger) DisableEvent(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.DisableEvent(tag)
}

// EnableEvent 对应 'events e <tag>'：重新启用被 DisableEvent 禁用的事件。
func (d *Debugger) EnableEvent(tag uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.core.EnableEvent(tag)
}
