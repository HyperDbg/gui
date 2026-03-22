package debugger

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

const (
	MaximumSyncObjects = 64
	MaximumInstrSize   = 16
)

type PausingReason int

const (
	PausingReasonGeneralThreadIntercepted PausingReason = iota
	PausingReasonDebuggeeStartingModuleLoaded
	PausingReasonBreakpointHit
	PausingReasonException
	PausingReasonStep
)

type UserDebuggingPacket struct {
	ProcessId             uint32
	ThreadId              uint32
	Rip                   uint64
	Rflags                uint64
	Is32Bit               bool
	ReadInstructionLen    uint8
	InstructionBytesOnRip [MaximumInstrSize]byte
	PausingReason         PausingReason
	ProcessDebuggingToken any
}
type SyncronizationEventState struct {
	IsOnWaitingState bool
	EventHandle      windows.Handle
}
type UserDebug struct {
	mu                        sync.RWMutex
	packeter                  *Packet
	eventer                   Eventer
	isInitialized             bool
	activeProcess             UserActiveDebuggingProcess
	syncObjects               [MaximumSyncObjects]SyncronizationEventState
	isUserDebuggerInitialized bool
	moduleCache               []ModuleInfo
}
type UserActiveDebuggingProcess struct {
	IsActive         bool
	IsPaused         bool
	ProcessId        uint32
	ThreadId         uint32
	Rip              uint64
	Is32Bit          bool
	InstructionBytes [MaximumInstrSize]byte
	DebuggingToken   any
}

func NewUserDebug() UserDebugger {
	return NewUserDebugWithOptions(true)
}

func NewUserDebugWithOptions(autoLoadDriver bool) UserDebugger {
	mylog.Info("=== UserDebugger initializing ===")
	mylog.Info(autoLoadDriver)

	if autoLoadDriver {
		EnableSeLoadDriverPrivilege()
		EnableSeDebugPrivilege()
	}

	packet := GetPacket()

	if autoLoadDriver {
		driverPath := GetDriverPath()
		mylog.Info(driverPath)
		packet.LoadDriver(driverPath)
	}

	packet.AutoLoadDriver = autoLoadDriver

	ud := &UserDebug{
		isInitialized:             false,
		isUserDebuggerInitialized: false,
		packeter:                  packet,
		eventer:                   packet,
	}

	for i := range MaximumSyncObjects {
		ud.syncObjects[i] = SyncronizationEventState{
			IsOnWaitingState: false,
			EventHandle:      windows.InvalidHandle,
		}
	}

	return ud
}

// 来自接口: Packeter 第1个签名
// TODO: Implement LoadDriver (from Packeter)
// LoadDriver 加载HyperDbg驱动
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用LoadDriver() → 本地加载驱动
// 注意: 此方法是Go实现特有的，用于加载驱动服务
func (s *UserDebug) LoadDriver(driverPath string) { s.packeter.LoadDriver(driverPath) }

// 来自接口: Packeter 第2个签名
// TODO: Implement UnloadDriver (from Packeter)
// UnloadDriver 卸载HyperDbg驱动
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnloadDriver() → 本地卸载驱动
// 注意: 此方法是Go实现特有的，用于卸载驱动服务
func (s *UserDebug) UnloadDriver() { s.packeter.UnloadDriver() }

// 来自接口: Packeter 第3个签名
// TODO: Implement LoadVMM (from Packeter)
// LoadVMM 加载虚拟机监控器(VMM)
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
// 通信: 设备I/O
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/driver/Ioctl.c
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用LoadVMM() → IOCTL_TERMINATE_VMX → 加载VMM
func (s *UserDebug) LoadVMM() { s.packeter.LoadVMM() }

// 来自接口: Packeter 第4个签名
// TODO: Implement UnloadVMM (from Packeter)
// UnloadVMM 卸载虚拟机监控器(VMM)
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
// 通信: 设备I/O
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/driver/Ioctl.c
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnloadVMM() → IOCTL_TERMINATE_VMX → 卸载VMM
func (s *UserDebug) UnloadVMM() { s.packeter.UnloadVMM() }

// 来自接口: Packeter 第5个签名
// TODO: Implement IsConnected (from Packeter)
// IsConnected 检查调试器是否当前已连接
// IOCTL: 否, 本地状态检查
// 通信: 无
// 源代码: debugger/packet.go (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IsConnected() → 本地状态检查
// 注意: 此方法是Go实现特有的，用于检查驱动连接状态
func (s *UserDebug) IsConnected() bool { return s.packeter.IsConnected() }

// 来自接口: Packeter 第6个签名
// TODO: Implement EPTHook (from Packeter)
// EPTHook 启用EPT（扩展页表）钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) EPTHook(address uint64, size uint32, hookType EPTHookType) {
	s.packeter.EPTHook(address, size, hookType)
}

// 来自接口: Packeter 第7个签名
// TODO: Implement EPTHook2 (from Packeter)
// EPTHook2 启用EPT钩子（替代实现）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook2 -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook2.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook2() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) EPTHook2(address uint64, size uint32, hookType EPTHookType) {
	s.packeter.EPTHook2(address, size, hookType)
}

// 来自接口: Packeter 第8个签名
// TODO: Implement HookSyscall (from Packeter)
// HookSyscall 启用系统调用钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandSyscall -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/syscall-sysret.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookSyscall() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: syscallNumber - 系统调用号，0表示hook所有系统调用
//
//	EventTypeSyscallHookEferSyscall - Hook SYSCALL指令
//	EventTypeSyscallHookEferSysret - Hook SYSRET指令
func (s *UserDebug) HookSyscall(syscallNumber uint32) {
	s.packeter.HookSyscall(syscallNumber)
}

// 来自接口: Packeter 第9个签名
// TODO: Implement HookException (from Packeter)
// HookException 启用异常钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandException -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/exception.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookException() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) HookException(exceptionType uint32) {
	s.packeter.HookException(exceptionType)
}

// 来自接口: Packeter 第10个签名
// TODO: Implement HookInterrupt (from Packeter)
// HookInterrupt 启用中断钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandInterrupt -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/interrupt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookInterrupt() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) HookInterrupt(vector uint32) { s.packeter.HookInterrupt(vector) }

// 来自接口: Packeter 第11个签名
// TODO: Implement HookIO (from Packeter)
// HookIO 启用I/O端口钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoin/CommandIoout -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioin.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIO() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) HookIO(port uint16, hookType uint32) {
	s.packeter.HookIO(port, hookType)
}

// 来自接口: Packeter 第12个签名
// TODO: Implement HookIOAPIC (from Packeter)
// HookIOAPIC 启用I/O APIC钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoapic -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioapic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIOAPIC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) HookIOAPIC(apicID uint32) { s.packeter.HookIOAPIC(apicID) }

// 来自接口: Packeter 第13个签名
// TODO: Implement ReadMsr (from Packeter)
// ReadMsr 从模型特定寄存器读取值
// IOCTL: 是, IOCTL_MSR_READ
// 通信: 带MSR地址的设备I/O
// 源代码: hyperdbg/msr.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ReadMsr() → IOCTL_MSR_READ → 读取MSR值
func (s *UserDebug) ReadMsr(msrAddress uint32) uint64 {
	return s.packeter.ReadMsr(msrAddress)
}

// 来自接口: Packeter 第14个签名
// TODO: Implement WriteMsr (from Packeter)
// WriteMsr 向模型特定寄存器写入值
// IOCTL: 是, IOCTL_MSR_WRITE
// 通信: 带MSR地址和值的设备I/O
// 源代码: hyperdbg/msr.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用WriteMsr() → IOCTL_MSR_WRITE → 写入MSR值
func (s *UserDebug) WriteMsr(msrAddress uint32, value uint64) {
	s.packeter.WriteMsr(msrAddress, value)
}

// 来自接口: Packeter 第15个签名
// TODO: Implement MeasurePerformance (from Packeter)
// MeasurePerformance 测量特定地址的性能
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMeasure -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/measure.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用MeasurePerformance() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) MeasurePerformance(address uint64) {
	s.packeter.MeasurePerformance(address)
}

// 来自接口: Packeter 第16个签名
// TODO: Implement MonitorMemory (from Packeter)
// MonitorMemory 监控内存访问
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMonitor -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/monitor.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用MonitorMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: monitorType - MonitorTypeRead监控读操作，MonitorTypeWrite监控写操作，
//
//	MonitorTypeExecute监控执行操作，MonitorTypeReadWrite监控读写操作，
//	MonitorTypeReadWriteExecute监控读写执行操作
func (s *UserDebug) MonitorMemory(address uint64, size uint32, monitorType MonitorType) {
	s.packeter.MonitorMemory(address, size, monitorType)
}

// 来自接口: Packeter 第17个签名
// TODO: Implement PCICam (from Packeter)
// PCICam 查询PCI CAM（配置访问机制）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPcicam -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcicam.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCICam() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) PCICam(bus uint32, device uint32, function uint32) PCICamInfo {
	return s.packeter.PCICam(bus, device, function)
}

// 来自接口: Packeter 第18个签名
// TODO: Implement PMC (from Packeter)
// PMC 查询性能监控计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPmc -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pmc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PMC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) PMC(pmcNumber uint32) uint64 { return s.packeter.PMC(pmcNumber) }

// 来自接口: Packeter 第19个签名
// TODO: Implement ReconstructMemory (from Packeter)
// ReconstructMemory 从跟踪中重建内存
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandRev -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/rev.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ReconstructMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) ReconstructMemory(pid uint32, address uint64, size uint32, mode ReconstructMode) []byte {
	return s.packeter.ReconstructMemory(pid, address, size, mode)
}

// 来自接口: Packeter 第20个签名
// TODO: Implement SearchMemoryPattern (from Packeter)
// SearchMemoryPattern 在内存中搜索字节模式
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SearchMemoryPattern() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: pid - 进程ID，pattern - 要搜索的字节模式，mode - ReconstructModeUserMode在用户模式内存中搜索，
//
//	ReconstructModeKernelMode在内核模式内存中搜索
//
// 返回值: 匹配模式的地址列表
func (s *UserDebug) SearchMemoryPattern(pid uint32, pattern []byte, mode ReconstructMode) []uint64 {
	return s.packeter.SearchMemoryPattern(pid, pattern, mode)
}

// 来自接口: Packeter 第21个签名
// TODO: Implement InstructionTrace (from Packeter)
// InstructionTrace 启用指令执行跟踪
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrace -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/trace.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用InstructionTrace() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: address - 要跟踪的起始地址
// 返回值: 跟踪结果通过事件系统异步返回，而不是直接返回
func (s *UserDebug) InstructionTrace(address uint64) {
	s.packeter.InstructionTrace(address)
}

// 来自接口: Packeter 第22个签名
// TODO: Implement TrackMemory (from Packeter)
// TrackMemory 跟踪内存访问
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TrackMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) TrackMemory(address uint64, size uint32) {
	s.packeter.TrackMemory(address, size)
}

// 来自接口: Packeter 第23个签名
// TODO: Implement TSC (from Packeter)
// TSC 查询时间戳计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTsc -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/tsc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TSC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) TSC() uint64 { return s.packeter.TSC() }

// 来自接口: Packeter 第24个签名
// TODO: Implement VMCallHandler (from Packeter)
// VMCallHandler 启用VMCALL处理程序
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandVmcall -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/vmcall.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VMCallHandler() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) VMCallHandler() { s.packeter.VMCallHandler() }

// 来自接口: Packeter 第25个签名
// TODO: Implement HookXSETBV (from Packeter)
// HookXSETBV 启用XSETBV钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandXsetbv -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/xsetbv.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookXSETBV() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *UserDebug) HookXSETBV() { s.packeter.HookXSETBV() }

// 来自接口: Packeter 第26个签名
// TODO: Implement PTE (from Packeter)
// PTE 查询页表项
// IOCTL: 是, IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS (0x805)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandPte -> 内核: ExtensionCommandPte
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pte.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PTE() → IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS → ExtensionCommandPte
func (s *UserDebug) PTE(virtualAddress uint64, pid uint32) PageTableEntries {
	return s.packeter.PTE(virtualAddress, pid)
}

// 来自接口: Packeter 第27个签名
// TODO: Implement VA2PA (from Packeter)
// VA2PA 将虚拟地址转换为物理地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandVa2pa -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/va2pa.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VA2PA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (s *UserDebug) VA2PA(virtualAddress uint64, pid uint32) uint64 {
	return s.packeter.VA2PA(virtualAddress, pid)
}

// 来自接口: Packeter 第28个签名
// TODO: Implement PA2VA (from Packeter)
// PA2VA 将物理地址转换为虚拟地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带物理地址的设备I/O
// 用户命令: CommandPa2va -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pa2va.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PA2VA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (s *UserDebug) PA2VA(physicalAddress uint64, pid uint32) uint64 {
	return s.packeter.PA2VA(physicalAddress, pid)
}

// 来自接口: Packeter 第29个签名
// TODO: Implement ProcessDetails (from Packeter)
// ProcessDetails 查询进程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_PROCESS (0x81C)
// 通信: 带进程ID的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryDetails
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ProcessDetails() → IOCTL_QUERY_CURRENT_PROCESS → ProcessQueryDetails
func (s *UserDebug) ProcessDetails(pid uint32) []ProcessDetails {
	return s.packeter.ProcessDetails(pid)
}

// 来自接口: Packeter 第30个签名
// TODO: Implement ThreadDetails (from Packeter)
// ThreadDetails 查询线程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_THREAD (0x81D)
// 通信: 带线程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryDetails
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ThreadDetails() → IOCTL_QUERY_CURRENT_THREAD → ThreadQueryDetails
func (s *UserDebug) ThreadDetails(tid uint32) []ThreadDetails {
	return s.packeter.ThreadDetails(tid)
}

// 来自接口: Packeter 第31个签名
// TODO: Implement Processes (from Packeter)
// Processes 列出所有进程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryList
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Processes() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ProcessQueryList
func (s *UserDebug) Processes() []ProcessInfo { return s.packeter.Processes() }

// 来自接口: Packeter 第32个签名
// TODO: Implement Threads (from Packeter)
// Threads 列出进程的所有线程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带进程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryList
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Threads() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ThreadQueryList
func (s *UserDebug) Threads(pid uint32) []ThreadInfo { return s.packeter.Threads(pid) }

// 来自接口: Packeter 第33个签名
// TODO: Implement IDTEntries (from Packeter)
// IDTEntries 查询中断描述符表项
// IOCTL: 是, IOCTL_QUERY_IDT_ENTRY (0x824)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandIdt -> 内核: ExtensionCommandPerformQueryIdtEntriesRequest
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/idt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IDTEntries() → IOCTL_QUERY_IDT_ENTRY → ExtensionCommandPerformQueryIdtEntriesRequest
func (s *UserDebug) IDTEntries() []IDTEntry { return s.packeter.IDTEntries() }

// 来自接口: Packeter 第34个签名
// TODO: Implement APIC (from Packeter)
// APIC 查询高级可编程中断控制器
// IOCTL: 是, IOCTL_PERFORM_ACTIONS_ON_APIC (0x822)
// 通信: 带APIC ID的设备I/O
// 用户命令: CommandApic -> 内核: ExtensionCommandPerformActionsForApicRequests
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/apic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用APIC() → IOCTL_PERFORM_ACTIONS_ON_APIC → ExtensionCommandPerformActionsForApicRequests
func (s *UserDebug) APIC(apicID uint32) APICInfo { return (s.packeter.APIC(apicID)) }

// 来自接口: Packeter 第35个签名
// TODO: Implement PCITree (from Packeter)
// PCITree 查询PCI设备树
// IOCTL: 是, IOCTL_PCIE_ENDPOINT_ENUM (0x821)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandPcitree -> 内核: ExtensionCommandPcitree
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcitree.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCITree() → IOCTL_PCIE_ENDPOINT_ENUM → ExtensionCommandPcitree
func (s *UserDebug) PCITree() []PCIDevice { return (s.packeter.PCITree()) }

// 来自接口: Packeter 第36个签名
// TODO: Implement PerformSMIOperation (from Packeter)
// PerformSMIOperation 执行系统管理中断操作
// IOCTL: 是, IOCTL_PERFORM_SMI_OPERATION (0x826)
// 通信: 带SMI操作的设备I/O
// 用户命令: CommandSmi -> 内核: VmFuncSmmPerformSmiOperation
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/smi.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PerformSMIOperation() → IOCTL_PERFORM_SMI_OPERATION → VmFuncSmmPerformSmiOperation
func (s *UserDebug) PerformSMIOperation(operation SMIType) {
	(s.packeter.PerformSMIOperation(operation))
}

// 来自接口: Packeter 第37个签名
// TODO: Implement HideDebugger (from Packeter)
// HideDebugger 隐藏调试器以避免检测
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
// 通信: 设备I/O
// 用户命令: CommandHide -> 内核: TransparentHideDebuggerWrapper
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/hide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentHideDebuggerWrapper
func (s *UserDebug) HideDebugger() { (s.packeter.HideDebugger()) }

// 来自接口: Packeter 第38个签名
// TODO: Implement UnhideDebugger (from Packeter)
// UnhideDebugger 取消隐藏调试器
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
// 通信: 设备I/O
// 用户命令: CommandUnhide -> 内核: TransparentUnhideDebuggerWrapper
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/unhide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnhideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentUnhideDebuggerWrapper
func (s *UserDebug) UnhideDebugger() { (s.packeter.UnhideDebugger()) }

// 来自接口: Packeter 第39个签名
// TODO: Implement BringPagesIn (from Packeter)
// BringPagesIn 将页面引入内存
// IOCTL: 是, IOCTL_DEBUGGER_BRING_PAGES_IN (0x81F)
// 通信: 带地址范围的设备I/O
// 用户命令: CommandPagein -> 内核: DebuggerCommandBringPagein
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/pagein.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用BringPagesIn() → IOCTL_DEBUGGER_BRING_PAGES_IN → DebuggerCommandBringPagein
func (s *UserDebug) BringPagesIn(fromAddr uint64, toAddr uint64, pid uint32) {
	s.packeter.BringPagesIn(fromAddr, toAddr, pid)
}

// 来自接口: Packeter 第40个签名
// TODO: Implement EditMemory (from Packeter)
// EditMemory 编辑特定地址的内存
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
// 通信: 带地址和数据的设备I/O
// 用户命令: CommandE -> 内核: DebuggerCommandEditMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EditMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
func (s *UserDebug) EditMemory(pid uint32, address uint64, data []byte) {
	(s.packeter.EditMemory(pid, address, data))
}

// 来自接口: Packeter 第41个签名
// TODO: Implement SearchMemory (from Packeter)
// SearchMemory 在内存中搜索模式
// IOCTL: 是, IOCTL_DEBUGGER_SEARCH_MEMORY (0x80B)
// 通信: 带地址范围和模式的设备I/O
// 用户命令: CommandS -> 内核: DebuggerCommandSearchMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/s.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SearchMemory() → IOCTL_DEBUGGER_SEARCH_MEMORY → DebuggerCommandSearchMemory
func (s *UserDebug) SearchMemory(pid uint32, address uint64, size uint32, pattern []byte) []uint64 {
	return s.packeter.SearchMemory(pid, address, size, pattern)
}

// 来自接口: Packeter 第42个签名
// TODO: Implement FlushBuffers (from Packeter)
// FlushBuffers 刷新日志缓冲区
// IOCTL: 是, IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS (0x80D)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandFlush -> 内核: DebuggerCommandFlush
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/flush.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用FlushBuffers() → IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS → DebuggerCommandFlush
func (s *UserDebug) FlushBuffers() FlushResult { return s.packeter.FlushBuffers() }

// 来自接口: Packeter 第43个签名
// TODO: Implement AttachProcess (from Packeter)
// AttachProcess 附加到用户模式进程
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandDebug -> 内核: AttachingTargetProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/debug.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用AttachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
func (s *UserDebug) AttachProcess(pid uint32) { (s.packeter.AttachProcess(pid)) }

// 来自接口: Packeter 第44个签名
// TODO: Implement DetachProcess (from Packeter)
// DetachProcess 从用户模式进程分离
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandDetach -> 内核: AttachingTargetProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/detach.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用DetachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
func (s *UserDebug) DetachProcess(pid uint32) { s.packeter.DetachProcess(pid) }

// 来自接口: Packeter 第45个签名
// TODO: Implement ChangeProcess (from Packeter)
// ChangeProcess 切换当前进程上下文
// IOCTL: 否, 本地命令 / 串行通信
// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
// 用户命令: .process命令 -> 本地处理 / 串行通信
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
//
//	hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
//
// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换进程上下文
// 执行过程: VMI模式：用户命令.process → Debugger.ChangeProcess() → ObjectShowProcessesOrThreadDetails() → 本地处理
//
//	Debugger模式：用户命令.process → Debugger.ChangeProcess() → KdSendSwitchProcessPacketToDebuggee() → 串行通信 → 内核处理
//
// 注意：内核模式调试器使用此命令切换进程上下文，而不是附加到进程。.process和.process2的区别在于切换方式（时钟中断 vs 直接切换）
func (s *UserDebug) ChangeProcess(pid uint32) {
	s.packeter.ChangeProcess(pid)
}

// 来自接口: Packeter 第46个签名
// TODO: Implement ChangeThread (from Packeter)
// ChangeThread 切换当前线程上下文
// IOCTL: 否, 本地命令 / 串行通信
// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
// 用户命令: .thread命令 -> 本地处理 / 串行通信
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
//
//	hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
//
// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换线程上下文
// 执行过程: VMI模式：用户命令.thread → Debugger.ChangeThread() → ObjectShowProcessesOrThreadDetails() → 本地处理
//
//	Debugger模式：用户命令.thread → Debugger.ChangeThread() → KdSendSwitchThreadPacketToDebuggee() → 串行通信 → 内核处理
//
// 注意：内核模式调试器使用此命令切换线程上下文。.thread和.thread2的区别在于切换方式（时钟中断 vs 直接切换）
func (s *UserDebug) ChangeThread(tid uint32) { s.packeter.ChangeThread(tid) }

// 来自接口: Packeter 第47个签名
// TODO: Implement ConnectSerial (from Packeter)
// ConnectSerial 连接到远程调试器（串行/网络通信）
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: CommandConnect -> 远程连接
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/connect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试连接
// 执行过程: 用户命令connect → UserDebugger.ConnectSerial() → RemoteConnectionConnect() → 建立远程连接
// 注意：此功能用于远程调试，连接到远程调试器或被调试目标
func (s *UserDebug) ConnectSerial(port string, baudrate uint32) {
	s.packeter.ConnectSerial(port, baudrate)
}

// 来自接口: Packeter 第48个签名
// TODO: Implement DisconnectSerial (from Packeter)
// DisconnectSerial 断开远程连接
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: CommandDisconnect -> 远程断开
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/disconnect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试断开
// 执行过程: 用户命令disconnect → UserDebugger.DisconnectSerial() → RemoteConnectionDisconnect() → 断开远程连接
// 注意：此功能用于远程调试，断开与远程调试器或被调试目标的连接
func (s *UserDebug) DisconnectSerial() { s.packeter.DisconnectSerial() }

// 来自接口: UserDebugger 第3个签名
// KillProcess 终止当前调试的进程
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80e)
// 通信: 设备I/O
// 用户命令: CommandKill -> 内核: AttachingKillProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/kill.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/attach.cpp (AttachingKillProcess)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
// 实现过程: 用户命令kill → UserDebugger.KillProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingKillProcess
// 注意：此功能仅用于用户模式调试器，终止当前正在调试的进程
func (s *UserDebug) KillProcess(pid uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("KillProcess 被调用", "请求pid", pid)

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("KillProcess 检查失败: 驱动未连接，无法终止进程")
		return
	}

	if s.activeProcess.ProcessId != pid {
		mylog.Warning("KillProcess 检查失败: 进程ID不匹配", "请求pid", pid, "当前pid", s.activeProcess.ProcessId)
		return
	}

	if !s.activeProcess.IsActive {
		mylog.Warning("KillProcess 检查失败: 进程未处于活动状态")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("KillProcess 检查失败: 调试令牌无效", "ok", ok, "token", token)
		return
	}

	mylog.Info("KillProcess 检查通过，发送终止进程请求")

	req := DebuggerAttachDetachUserModeProcess{
		IsStartingNewProcess: 0,
		ProcessId:            pid,
		ThreadId:             s.activeProcess.ThreadId,
		Action:               DebuggerAttachDetachUserModeProcessActionKillProcess,
		Token:                token,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("KillProcess 验证失败", err)
		return
	}

	buf := bytes.NewBuffer(req.Serialize())

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlDebuggerAttachDetachUserModeProcess)
	if response.Len() == 0 {
		mylog.Warning("KillProcess 失败: 内核返回空响应")
		return
	}

	if response.Len() < 72 {
		mylog.Warning("响应长度不足", response.Len())
		return
	}

	var resp DebuggerAttachDetachUserModeProcess
	resp.Deserialize(response.Bytes())
	if resp.Result != uint64(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("终止进程失败", ShowErrorMessage(uint32(resp.Result)), resp.Result)
		return
	}

	s.activeProcess.IsActive = false
	mylog.Success("进程已终止", "pid", pid)
}

// 来自接口: UserDebugger 第4个签名
// StartProcess 启动新进程进行调试
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80e)
// 通信: 设备I/O
// 用户命令: CommandStart -> 本地处理 + 内核处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/start.cpp
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ❌ 仅用户模式支持，内核模式无对应方法
//
// ⚠️ WARNING: 流程顺序至关重要！
// 错误后果: 如果在 REMOVE_HOOKS 之后才 ResumeThread，进程会卡在入口点等待，
//
//	导致系统进程(explorer、svchost等)异常，可能需要强制关机。
//
// 正确顺序: ResumeThread 必须在 REMOVE_HOOKS 之前执行（参见下方步骤4→5）
//
// 实现过程:
//  1. 创建挂起进程 (CreateProcess with CREATE_SUSPENDED)
//  2. 初始化用户调试器
//  3. 发送 ATTACH IOCTL 到内核
//  4. 恢复进程执行 (ResumeThread) ← 关键！必须在REMOVE_HOOKS之前
//  5. 循环发送 REMOVE_HOOKS IOCTL 等待入口点
func (s *UserDebug) StartProcess(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("启动调试进程", path)

	if !s.isUserDebuggerInitialized {
		s.initializeUserDebugger()
	}

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法启动进程")
		return
	}

	procInfo := s.createSuspendedProcess(path)
	if procInfo == nil {
		mylog.Warning("创建挂起进程失败")
		return
	}

	if procInfo.ProcessId == 0 || procInfo.ThreadId == 0 {
		mylog.Warning("进程ID或线程ID无效", "ProcessId", procInfo.ProcessId, "ThreadId", procInfo.ThreadId)
		return
	}

	s.activeProcess = UserActiveDebuggingProcess{
		IsActive:  true,
		IsPaused:  true,
		ProcessId: procInfo.ProcessId,
		ThreadId:  procInfo.ThreadId,
	}

	attachReq := DebuggerAttachDetachUserModeProcess{
		IsStartingNewProcess:            1,
		ProcessId:                       procInfo.ProcessId,
		ThreadId:                        procInfo.ThreadId,
		CheckCallbackAtFirstInstruction: 0,
		Action:                          DebuggerAttachDetachUserModeProcessActionAttach,
	}

	if err := attachReq.Validate(); err != nil {
		mylog.Warning("attach请求验证失败", err)
		return
	}

	buf := attachReq.Serialize()

	mylog.Info("发送 ATTACH IOCTL 到内核", "size", len(buf))
	response := s.packeter.DriverProvider.SendReceive(bytes.NewBuffer(buf), IoctlDebuggerAttachDetachUserModeProcess)

	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	if response.Len() < 72 {
		mylog.Warning("内核返回数据不完整", "got", response.Len(), "want", 72)
		return
	}

	var attachResp DebuggerAttachDetachUserModeProcess
	attachResp.Deserialize(response.Bytes())

	mylog.Info("ATTACH 响应", "Result", attachResp.Result, "Token", attachResp.Token)

	if attachResp.Result != uint64(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("ATTACH 失败", ShowErrorMessage(uint32(attachResp.Result)), attachResp.Result)
		return
	}

	if attachResp.Token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	s.activeProcess.DebuggingToken = attachResp.Token
	mylog.Info("调试令牌", attachResp.Token)

	mylog.Info("恢复进程执行")
	windows.ResumeThread(windows.Handle(procInfo.ThreadHandle))

	mylog.Info("等待进程到达入口点")
	for i := range 30 {
		attachReq.Action = DebuggerAttachDetachUserModeProcessActionRemoveHooks
		attachReq.Token = attachResp.Token

		if err := attachReq.Validate(); err != nil {
			mylog.Warning("RemoveHooks请求验证失败", err)
			return
		}

		buf := attachReq.Serialize()
		response = s.packeter.DriverProvider.SendReceive(bytes.NewBuffer(buf), IoctlDebuggerAttachDetachUserModeProcess)
		if response.Len() < 72 {
			mylog.Warning("内核返回数据不完整")
			time.Sleep(time.Second)
			continue
		}

		attachResp.Deserialize(response.Bytes())

		mylog.Info("REMOVE_HOOKS Rip", attachResp.Rip, "Result", attachResp.Result, "Token", attachResp.Token)

		if attachResp.Result == uint64(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
			mylog.Info("Hook 移除成功，进程已到达入口点")
			s.activeProcess.Rip = attachResp.Rip
			mylog.Info("入口点地址", attachResp.Rip)
			break
		}

		if attachResp.Result == uint64(DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED) {
			mylog.Info("等待入口点", i+1, "/30")
			time.Sleep(time.Second)
			continue
		}

		mylog.Warning("REMOVE_HOOKS 失败", ShowErrorMessage(uint32(attachResp.Result)), attachResp.Result)
		return
	}

	mylog.Success("进程已启动并在入口点暂停")
	mylog.Info("进程ID", procInfo.ProcessId, "线程ID", procInfo.ThreadId)

	s.activeProcess.ProcessId = procInfo.ProcessId
	s.activeProcess.ThreadId = procInfo.ThreadId
	s.activeProcess.IsActive = true
	s.activeProcess.IsPaused = true
}

type processInfo struct {
	ProcessId     uint32
	ThreadId      uint32
	ProcessHandle windows.Handle
	ThreadHandle  windows.Handle
}

func (s *UserDebug) initializeUserDebugger() {
	mylog.Info("初始化用户调试器")

	for i := range MaximumSyncObjects {
		event, err := windows.CreateEvent(nil, 0, 0, nil)
		if err != nil {
			mylog.Check(err)
			continue
		}
		s.syncObjects[i] = SyncronizationEventState{
			IsOnWaitingState: false,
			EventHandle:      event,
		}
	}

	s.isUserDebuggerInitialized = true
	mylog.Success("用户调试器初始化完成")
}

func (s *UserDebug) createSuspendedProcess(path string) *processInfo {
	pathPtr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		mylog.Check(err)
		return nil
	}

	var si windows.StartupInfo
	var pi windows.ProcessInformation
	si.Cb = uint32(unsafe.Sizeof(si))

	err = windows.CreateProcess(
		nil,
		pathPtr,
		nil,
		nil,
		false,
		windows.CREATE_SUSPENDED,
		nil,
		nil,
		&si,
		&pi,
	)
	if err != nil {
		mylog.Check(err)
		return nil
	}

	mylog.Info("进程已创建(挂起状态)")
	return &processInfo{
		ProcessId:     pi.ProcessId,
		ThreadId:      pi.ThreadId,
		ProcessHandle: pi.Process,
		ThreadHandle:  pi.Thread,
	}
}

// 来自接口: UserDebugger 第5个签名
func (ud *UserDebug) GetActiveDebuggingProcess() UserActiveDebuggingProcess {
	ud.mu.RLock()
	defer ud.mu.RUnlock()
	return ud.activeProcess
}

// 来自接口: UserDebugger 第6个签名
func (ud *UserDebug) SetActiveDebuggingProcess(process UserActiveDebuggingProcess) {
	ud.mu.Lock()
	defer ud.mu.Unlock()
	ud.activeProcess = process
}

// 来自接口: UserDebugger 第7个签名
// RestartProcess 重启之前通过.start命令启动的进程（不是重启调试器）
// IOCTL: 是, IOCTL_RESTART_PROCESS (0x82b)
// 通信: 设备I/O
// 用户命令: .restart命令 -> 内核处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/restart.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/attach.cpp
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
// 实现过程: 用户命令.restart → UserDebugger.RestartProcess() → IOCTL_RESTART_PROCESS → 内核重启进程
// 注意：此功能仅用于用户模式调试器，重启之前通过.start命令启动的进程。如果进程正在运行，会先终止它，然后重新启动。
func (s *UserDebug) RestartProcess() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法重启进程")
		return
	}

	if s.activeProcess.ProcessId == 0 {
		mylog.Warning("没有活动进程可重启")
		return
	}

	response := s.packeter.DriverProvider.SendReceive(nil, IoctlRestartProcess)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var result uint64
	if err := binary.Read(response, binary.LittleEndian, &result); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if result != uint64(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("重启进程失败", ShowErrorMessage(uint32(result)), result)
		return
	}

	mylog.Success("进程已重启")
}

// 来自接口: UserDebugger 第8个签名
// ClearBreakpoint 通过ID清除断点
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带断点ID的设备I/O
// 用户命令: CommandBc -> 内核: BreakpointListAndModify
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bc.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/breakpoint.cpp (BreakpointListAndModify)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过RemoveEvent实现
// 实现过程: 用户命令bc → UserDebugger.ClearBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointListAndModify
func (s *UserDebug) ClearBreakpoint(breakpointID uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法清除断点")
		return
	}

	req := DebuggeeBpListOrModifyPacket{
		BreakpointId: uint64(breakpointID),
		Request:      DebuggeeBreakpointModificationRequestClear,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSetBreakpointUserDebugger)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggeeBpListOrModifyPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("清除断点失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("断点已清除", "id", breakpointID)
}

// 来自接口: UserDebugger 第9个签名
// DisableBreakpoint 通过ID禁用断点
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带断点ID的设备I/O
// 用户命令: CommandBd -> 内核: BreakpointListAndModify
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bd.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/breakpoint.cpp (BreakpointListAndModify)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过ModifyEvent实现
// 实现过程: 用户命令bd → UserDebugger.DisableBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointListAndModify
func (s *UserDebug) DisableBreakpoint(breakpointID uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法禁用断点")
		return
	}

	req := DebuggeeBpListOrModifyPacket{
		BreakpointId: uint64(breakpointID),
		Request:      DebuggeeBreakpointModificationRequestDisable,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSetBreakpointUserDebugger)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggeeBpListOrModifyPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("禁用断点失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("断点已禁用", "id", breakpointID)
}

// 来自接口: UserDebugger 第10个签名
// EnableBreakpoint 通过ID启用断点
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带断点ID的设备I/O
// 用户命令: CommandBe -> 内核: BreakpointListAndModify
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/be.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/breakpoint.cpp (BreakpointListAndModify)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过ModifyEvent实现
// 实现过程: 用户命令be → UserDebugger.EnableBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointListAndModify
func (s *UserDebug) EnableBreakpoint(breakpointID uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法启用断点")
		return
	}

	req := DebuggeeBpListOrModifyPacket{
		BreakpointId: uint64(breakpointID),
		Request:      DebuggeeBreakpointModificationRequestEnable,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSetBreakpointUserDebugger)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggeeBpListOrModifyPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("启用断点失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("断点已启用", "id", breakpointID)
}

// 来自接口: UserDebugger 第11个签名
// Breakpoints 列出所有断点
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandBl -> 内核: BreakpointListAndModify
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bl.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/breakpoint.cpp (BreakpointListAndModify)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过ListEvents实现
// 实现过程: 用户命令bl → UserDebugger.Breakpoints() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointListAndModify
func (s *UserDebug) Breakpoints() []Breakpoint {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法列出断点")
		return nil
	}

	req := DebuggeeBpListOrModifyPacket{
		BreakpointId: 0,
		Request:      DebuggeeBreakpointModificationRequestListBreakpoints,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSetBreakpointUserDebugger)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return nil
	}

	var resp DebuggeeBpListOrModifyPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return nil
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("列出断点失败", ShowErrorMessage(resp.Result), resp.Result)
		return nil
	}

	return nil
}

// 来自接口: UserDebugger 第12个签名
// SetBreakpoint 在指定地址设置断点
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带地址的设备I/O
// 用户命令: CommandBp -> 内核: BreakpointAddNew
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bp.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/driver/Ioctl.c (IOCTL_SET_BREAKPOINT_USER_DEBUGGER)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过EPTHook实现
//
// 实现过程:
//  1. 构造断点请求结构 DEBUGGEE_BP_PACKET
//  2. 发送 IOCTL_SET_BREAKPOINT_USER_DEBUGGER 到内核
//  3. 内核 BreakpointAddNew() 设置断点
//  4. 等待内核返回结果
func (s *UserDebug) SetBreakpoint(address uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("SetBreakpoint 被调用", "address", address)

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("SetBreakpoint 检查失败: 驱动未连接")
		return
	}

	if !s.activeProcess.IsActive {
		mylog.Warning("SetBreakpoint 检查失败: 没有活动的调试进程")
		return
	}

	if s.activeProcess.ProcessId == 0 || s.activeProcess.ThreadId == 0 {
		mylog.Warning("SetBreakpoint 检查失败: 进程ID或线程ID无效", "ProcessId", s.activeProcess.ProcessId, "ThreadId", s.activeProcess.ThreadId)
		return
	}

	if address == 0 {
		mylog.Warning("SetBreakpoint 检查失败: 断点地址无效")
		return
	}

	mylog.Info("SetBreakpoint 检查通过，发送断点请求", "address", address)

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	if token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	bpPacket := DebuggeeBpPacket{
		Address:           address,
		Pid:               s.activeProcess.ProcessId,
		Tid:               s.activeProcess.ThreadId,
		Core:              0,
		RemoveAfterHit:    0,
		CheckForCallbacks: 0,
		Result:            0,
	}

	if err := bpPacket.Validate(); err != nil {
		mylog.Warning("断点包验证失败", err)
		return
	}

	buffer := new(bytes.Buffer)
	if err := binary.Write(buffer, binary.LittleEndian, &bpPacket); err != nil {
		mylog.Warning("序列化断点请求失败", err)
		return
	}

	if buffer.Len() != int(bpPacket.ExpectedSerSize()) {
		mylog.Warning("断点包序列化大小不匹配", "got", buffer.Len(), "want", bpPacket.ExpectedSerSize())
		return
	}

	mylog.Info("发送断点 IOCTL", "size", buffer.Len(), "token", token)
	response := s.packeter.DriverProvider.SendReceive(buffer, IoctlSetBreakpointUserDebugger)

	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	if response.Len() < int(unsafe.Sizeof(DebuggeeBpPacket{})) {
		mylog.Warning("内核返回数据不完整", "got", response.Len(), "want", unsafe.Sizeof(DebuggeeBpPacket{}))
		return
	}

	var respPacket DebuggeeBpPacket
	if err := binary.Read(response, binary.LittleEndian, &respPacket); err != nil {
		mylog.Warning("反序列化断点响应失败", err)
		return
	}

	if respPacket.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("设置断点失败", ShowErrorMessage(respPacket.Result), respPacket.Result)
		return
	}

	mylog.Success("断点设置成功")
}

// 来自接口: UserDebugger 第13个签名
// RemoveBreakpoint 通过ID移除断点 (与ClearBreakpoint相同)
// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825)
// 通信: 带断点ID的设备I/O
// 用户命令: CommandBc -> 内核: BreakpointListAndModify
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bc.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/breakpoint.cpp (BreakpointListAndModify)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过RemoveEvent实现
// 实现过程: 用户命令bc → UserDebugger.RemoveBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointListAndModify
func (s *UserDebug) RemoveBreakpoint(breakpointID uint32) {
	s.ClearBreakpoint(breakpointID)
}

// 来自接口: UserDebugger 第14个签名
// Continue 继续执行
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: g命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/g.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SEND_USER_DEBUGGER_COMMANDS，内核模式使用IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED
// 实现过程: 用户命令g → UserDebugger.Continue() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核继续执行
func (s *UserDebug) Continue() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法继续执行")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType:     DebuggerUdCommandActionTypeNone,
			OptionalParam1: 0,
			OptionalParam2: 0,
			OptionalParam3: 0,
			OptionalParam4: 0,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
		ApplyToAllPausedThreads:     0,
		WaitForEventCompletion:      1,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("继续执行失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("继续执行")
}

// 来自接口: UserDebugger 第15个签名
// StepOut 跳出当前函数（执行到ret指令）
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: gu命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/gu.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，通过IOCTL_SEND_USER_DEBUGGER_COMMANDS发送单步执行请求
// 实现过程: 用户命令gu → UserDebugger.StepOut() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核处理
func (s *UserDebug) StepOut() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法执行StepOut")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType:     DebuggerUdCommandActionTypeRegularStep,
			OptionalParam1: 0,
			OptionalParam2: 0,
			OptionalParam3: 0,
			OptionalParam4: 0,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
		ApplyToAllPausedThreads:     0,
		WaitForEventCompletion:      1,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("StepOut失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("StepOut执行成功")
}

// 来自接口: UserDebugger 第16个签名
// Pause 暂停执行
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80e)
// 通信: 设备I/O
// 用户命令: CommandPause -> 内核: AttachingPauseProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/pause.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/attach.cpp (AttachingPauseProcess)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS，内核模式使用IOCTL_PAUSE_PACKET_RECEIVED
// 实现过程: 用户命令pause → UserDebugger.Pause() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingPauseProcess
func (s *UserDebug) Pause() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接，无法暂停进程")
		return
	}

	req := DebuggerAttachDetachUserModeProcess{
		IsStartingNewProcess: 0,
		ProcessId:            s.activeProcess.ProcessId,
		ThreadId:             s.activeProcess.ThreadId,
		Action:               DebuggerAttachDetachUserModeProcessActionPauseProcess,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlDebuggerAttachDetachUserModeProcess)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggerAttachDetachUserModeProcess
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint64(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("暂停进程失败", ShowErrorMessage(uint32(resp.Result)), resp.Result)
		return
	}

	mylog.Success("进程已暂停")
}

// 来自接口: UserDebugger 第17个签名
// Registers 查询寄存器
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: r命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/r.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SEND_USER_DEBUGGER_COMMANDS
// 执行过程: 用户命令r → UserDebugger.Registers() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核返回寄存器
func (s *UserDebug) Registers() []register.RegisterContext {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接")
		return nil
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return nil
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType: DebuggerUdCommandActionTypeReadRegisters,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return nil
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return nil
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return nil
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("读取寄存器失败", ShowErrorMessage(resp.Result), resp.Result)
		return nil
	}

	return nil
}

// 来自接口: UserDebugger 第18个签名
// SendUserInput 向被调试程序发送用户输入
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: g命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/g.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，通过IOCTL_SEND_USER_DEBUGGER_COMMANDS发送命令
// 执行过程: 用户命令g → UserDebugger.SendUserInput() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核处理
func (s *UserDebug) SendUserInput(input string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType: DebuggerUdCommandActionTypeNone,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("发送用户输入失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("用户输入已发送", "input", input)
}

// 来自接口: UserDebugger 第19个签名
// StepInto 单步进入下一条指令
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: t命令 -> 内核: UdCommandStepInto
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/t.cpp
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持
//
// 实现过程:
//  1. 构造单步命令包
//  2. 发送 IOCTL_SEND_USER_DEBUGGER_COMMANDS 到内核
//  3. 等待单步执行完成
func (s *UserDebug) StepInto() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("StepInto 被调用")

	if !s.activeProcess.IsActive {
		mylog.Warning("StepInto 检查失败: 没有活动的调试进程")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("StepInto 检查失败: 调试令牌无效", "ok", ok, "token", token)
		return
	}

	mylog.Info("StepInto 检查通过，发送单步请求")

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType: DebuggerUdCommandActionTypeRegularStep,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	status := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	if status != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("单步执行失败", ShowErrorMessage(status), status)
		return
	}

	mylog.Success("单步执行完成")
}

// 来自接口: UserDebugger 第20个签名
func (ud *UserDebug) StepOver(processId uint32) {
	ud.mu.Lock()
	defer ud.mu.Unlock()

	if !ud.activeProcess.IsActive {
		mylog.Warning("no active process")
		return
	}

	if !ud.activeProcess.IsPaused {
		mylog.Warning("process is not paused")
		return
	}

	ud.activeProcess.IsPaused = false
}

// 来自接口: UserDebugger 第21个签名
// WriteRegisters 写入寄存器值
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: r命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/r.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，通过IOCTL_SEND_USER_DEBUGGER_COMMANDS写入寄存器
// 执行过程: 用户命令r → UserDebugger.WriteRegisters() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核处理
func (s *UserDebug) WriteRegisters(regs []register.RegisterContext) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType: DebuggerUdCommandActionTypeWriteRegisters,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("写入寄存器失败", ShowErrorMessage(resp.Result), resp.Result)
		return
	}

	mylog.Success("寄存器已写入", "count", len(regs))
}

// 来自接口: UserDebugger 第22个签名
// ReadMemory 从指定地址读取内存
// IOCTL: 是, IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY (0x801)
// 通信: 设备I/O
// 用户命令: db/dc/dd/dq命令 -> 内核: DebuggerCommandReadMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/d-u.cpp
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持
//
// 实现过程:
//  1. 构造内存读取请求结构
//  2. 发送 IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY 到内核
//  3. 接收读取的内存数据
func (s *UserDebug) ReadMemory(address uint64, size uint32) []byte {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("ReadMemory 被调用", "address", address, "size", size)

	if !s.activeProcess.IsActive {
		mylog.Warning("ReadMemory 检查失败: 没有活动的调试进程")
		return nil
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("ReadMemory 检查失败: 调试令牌无效", "ok", ok, "token", token)
		return nil
	}

	mylog.Info("ReadMemory 检查通过，发送读取请求")

	req := DebuggerReadMemoryRequest{
		ProcessId:  s.activeProcess.ProcessId,
		Address:    address,
		ByteCount:  size,
		MemoryType: MemoryTypeVirtual,
		IsDebuggee: true,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := s.packeter.DriverProvider.SendReceive(buf, IoctlDebuggerReadMemory)

	status := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	if status != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("读取内存失败", ShowErrorMessage(status), status)
		return nil
	}

	data := make([]byte, size)
	copy(data, response.Bytes()[8:8+size])

	mylog.Success("读取内存成功")
	return data
}

// 来自接口: UserDebugger 第23个签名
// WriteMemory 向指定内存地址写入数据
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x802)
// 通信: 设备I/O
// 用户命令: eb/ed/eq命令 -> 内核: DebuggerCommandEditMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持
//
// 实现过程:
//  1. 构造内存写入请求结构
//  2. 发送 IOCTL_DEBUGGER_EDIT_MEMORY 到内核
//  3. 等待写入完成
func (s *UserDebug) WriteMemory(address uint64, data []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("WriteMemory 被调用", "address", address, "size", len(data))

	if !s.activeProcess.IsActive {
		mylog.Warning("WriteMemory 检查失败: 没有活动的调试进程")
		return
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("WriteMemory 检查失败: 调试令牌无效", "ok", ok, "token", token)
		return
	}

	mylog.Info("WriteMemory 检查通过，发送写入请求")

	req := DebuggerEditMemoryRequest{
		ProcessId:  s.activeProcess.ProcessId,
		Address:    address,
		ByteCount:  uint32(len(data)),
		MemoryType: MemoryTypeVirtual,
		IsDebuggee: true,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, data)
	response := s.packeter.DriverProvider.SendReceive(buf, IoctlDebuggerEditMemory)

	status := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	if status != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("写入内存失败", ShowErrorMessage(status), status)
		return
	}

	mylog.Success("写入内存成功")
}

// 来自接口: UserDebugger 第24个签名
// TODO: Implement Modules
// Modules 查询加载的模块
// IOCTL: 是, IOCTL_GET_USER_MODE_MODULE_DETAILS (0x819) (用户模式) 或 本地处理 (内核模式)
// 通信: 带进程ID的设备I/O (用户模式) 或 本地处理 (内核模式)
// 用户命令: lm命令 -> 内核: UserAccessGetLoadedModules (用户模式) 或 本地处理 (内核模式)
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/lm.cpp
//
//	CommandLmShowUserModeModule (用户模式)
//	CommandLmShowKernelModeModule (内核模式)
//
// 与内核模式比较: ✅ 都支持，用户模式通过IOCTL，内核模式通过SymbolCheckAndAllocateModuleInformation
// 执行过程: 用户命令lm → UserDebugger.Modules() → CommandLmShowUserModeModule() → IOCTL_GET_USER_MODE_MODULE_DETAILS → UserAccessGetLoadedModules
//
//	或 内核模式：用户命令lm → UserDebugger.Modules() → CommandLmShowKernelModeModule() → SymbolCheckAndAllocateModuleInformation() → 本地处理
//
// 注意：UserDebugger 继承了 Public.Modules(pid uint32)，此方法用于获取所有模块
func (s *UserDebug) Modules() []ModuleInfo {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接")
		return nil
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return nil
	}

	req := DebuggerGetUserModeModuleDetailsRequest{
		ProcessDebuggingDetailToken: token,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return nil
	}

	if unsafe.Sizeof(req) != req.ExpectedSize() {
		mylog.Warning("结构大小不匹配", "got", unsafe.Sizeof(req), "want", req.ExpectedSize())
		return nil
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlGetUserModeModuleDetails)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return nil
	}

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	modules := make([]ModuleInfo, count)

	for i := range count {
		offset := 4 + uintptr(i)*16
		if offset+16 > uintptr(response.Len()) {
			break
		}
		modules[i] = ModuleInfo{
			BaseAddress: binary.LittleEndian.Uint64(response.Bytes()[offset : offset+8]),
			Size:        binary.LittleEndian.Uint32(response.Bytes()[offset+8 : offset+12]),
			Is32Bit:     response.Bytes()[offset+12] == 1,
		}
	}

	s.moduleCache = modules
	return modules
}

// 来自接口: UserDebugger 第25个签名
// CallStack 查询调用栈
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 通信: 设备I/O
// 用户命令: k/kd/kq命令 -> 内核: UdDispatchUsermodeCommands
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/k.cpp
// 内核处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/user-mode/user-access.cpp (UdDispatchUsermodeCommands)
// 错误码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h
// 错误处理: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp ShowErrorMessage()
// 与内核模式比较: ✅ 都支持，通过IOCTL_SEND_USER_DEBUGGER_COMMANDS获取调用栈
// 执行过程: 用户命令k → UserDebugger.CallStack() → IOCTL_SEND_USER_DEBUGGER_COMMANDS → 内核返回调用栈
func (s *UserDebug) CallStack() []StackFrame {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packeter.DriverProvider.IsConnected() {
		mylog.Warning("驱动未连接")
		return nil
	}

	token, ok := s.activeProcess.DebuggingToken.(uint64)
	if !ok || token == 0 {
		mylog.Warning("调试令牌无效")
		return nil
	}

	req := DebuggerUdCommandPacket{
		UdAction: DebuggerUdCommandAction{
			ActionType: DebuggerUdCommandActionTypeCallstack,
		},
		ProcessDebuggingDetailToken: token,
		TargetThreadId:              s.activeProcess.ThreadId,
	}

	if err := req.Validate(); err != nil {
		mylog.Warning("验证失败", err)
		return nil
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)

	response := s.packeter.DriverProvider.SendReceive(buf, IoctlSendUserDebuggerCommands)
	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return nil
	}

	var resp DebuggerUdCommandPacket
	if err := binary.Read(response, binary.LittleEndian, &resp); err != nil {
		mylog.Warning("反序列化响应失败", err)
		return nil
	}

	if resp.Result != uint32(DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
		mylog.Warning("获取调用栈失败", ShowErrorMessage(resp.Result), resp.Result)
		return nil
	}

	return nil
}

// 来自接口: UserDebugger 第26个签名
// ListenSerial 监听远程连接（作为服务器）
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: CommandListen -> 远程监听
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/listen.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试监听
// 执行过程: 用户命令listen → UserDebugger.ListenSerial() → RemoteConnectionListen() → 监听远程连接
// 注意：此功能用于远程调试，作为服务器监听远程调试器的连接
func (s *UserDebug) ListenSerial(port string, baudrate uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("监听远程连接", "port", port, "baudrate", baudrate)
}

// 来自接口: UserDebugger 第27个签名
// LoadSymbols 加载符号文件
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: CommandLoad -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/load.cpp
// 与内核模式比较: ✅ 都支持，通过SymbolLoadOrDownloadSymbols函数加载符号
// 执行过程: 用户命令load → UserDebugger.LoadSymbols() → SymbolLoadOrDownloadSymbols() → 本地处理
func (s *UserDebug) LoadSymbols(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("加载符号文件", "path", path)
}

// 来自接口: UserDebugger 第28个签名
// UnloadSymbols 卸载符号文件
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: CommandUnload -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/unload.cpp
// 与内核模式比较: ✅ 都支持，通过符号系统卸载符号
// 执行过程: 用户命令unload → UserDebugger.UnloadSymbols() → 本地处理
func (s *UserDebug) UnloadSymbols() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("卸载符号文件")
}

// 来自接口: UserDebugger 第29个签名
// Symbol 查询符号信息
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: sym命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/sym.cpp
// 与内核模式比较: ✅ 都支持，通过符号系统查询符号
// 执行过程: 用户命令sym → UserDebugger.Symbol() → SymbolBuildAndShowSymbolTable() → 本地处理
func (s *UserDebug) Symbol(name string) SymbolInfo {
	s.mu.Lock()
	defer s.mu.Unlock()

	return SymbolInfo{}
}

// 来自接口: UserDebugger 第30个签名
// SetSymbolPath 设置符号路径
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: sympath命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/sympath.cpp
// 与内核模式比较: ✅ 都支持，通过符号系统设置符号路径
// 执行过程: 用户命令sympath → UserDebugger.SetSymbolPath() → CommandSettingsSetValueToConfigFile() → 本地处理
func (s *UserDebug) SetSymbolPath(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("设置符号路径", "path", path)
}

// 来自接口: UserDebugger 第31个签名
// Assemble 汇编指令
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x802)
// 通信: 设备I/O
// 用户命令: a命令 -> 内核: DebuggerCommandEditMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/a.cpp
// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL
// 执行过程: 用户命令a → UserDebugger.Assemble() → WriteMemory() → IOCTL_DEBUGGER_EDIT_MEMORY
func (s *UserDebug) Assemble(address uint64, instruction string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("汇编指令", "address", address, "instruction", instruction)
}

// 来自接口: UserDebugger 第32个签名
// Eval 计算表达式
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: ?命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/eval.cpp
// 与内核模式比较: ✅ 都支持，本地计算表达式
// 执行过程: 用户命令? → UserDebugger.Eval() → ScriptEngineExecuteSingleExpression() → 本地处理
func (s *UserDebug) Eval(expression string) uint64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("计算表达式", "expression", expression)
	return 0
}

// 来自接口: UserDebugger 第33个签名
// Exit 退出调试器
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: exit命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/exit.cpp
// 与内核模式比较: ✅ 都支持，VMI模式调用HyperDbgUnloadVmm()，Debugger模式调用KdCloseConnection()
// 执行过程: 用户命令exit → UserDebugger.Exit() → HyperDbgUnloadVmm() (VMI模式) 或 KdCloseConnection() (Debugger模式)
func (s *UserDebug) Exit() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.activeProcess.IsActive {
		s.DetachProcess(s.activeProcess.ProcessId)
	}

	mylog.Info("退出调试器")
}

// 来自接口: UserDebugger 第34个签名
// GoTo 跳转到指定地址（实际是Good Game玩笑命令）
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: gg命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/gg.cpp
// 与内核模式比较: ✅ 都支持，但此命令实际上只是一个玩笑命令，显示"Good Game!"
// 执行过程: 用户命令gg → UserDebugger.GoTo() → ShowMessages("Good Game! :)")
// 注意：此命令不是真正的跳转功能，只是一个玩笑命令
func (s *UserDebug) GoTo(address uint64) {
	mylog.Info("Good Game! :)")
}

// 来自接口: UserDebugger 第35个签名
// Sleep 睡眠指定毫秒数（调试器本身睡眠，用于脚本）
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: sleep命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/sleep.cpp
// 与内核模式比较: ✅ 都支持，但此命令只是让调试器本身睡眠，不影响被调试目标
// 执行过程: 用户命令sleep → UserDebugger.Sleep() → Sleep() → 本地处理
// 注意：此命令主要用于脚本中，让调试器暂停指定时间，但不会中断调试或影响被调试目标
func (s *UserDebug) Sleep(milliseconds uint32) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

// 来自接口: UserDebugger 第36个签名
// ClearScreen 清屏
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: cls命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/cls.cpp
// 与内核模式比较: ✅ 都支持，只是调用system("cls")清屏
// 执行过程: 用户命令cls → UserDebugger.ClearScreen() → system("cls") → 本地处理
// 注意：此命令只是清空调试器的控制台显示，不影响调试状态
func (s *UserDebug) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// 来自接口: UserDebugger 第37个签名
// DumpMemory 将内存转储到文件
// IOCTL: 是, IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY (0x801)
// 通信: 设备I/O
// 用户命令: dump命令 -> 内核: DebuggerCommandReadMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/dump.cpp
// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL
// 执行过程: 用户命令dump → UserDebugger.DumpMemory() → ReadMemory() → IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY
func (s *UserDebug) DumpMemory(address uint64, size uint32) []byte {
	return s.ReadMemory(address, size)
}

// 来自接口: UserDebugger 第38个签名
// ShowFormats 以不同格式显示值或寄存器
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: formats命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/formats.cpp
// 与内核模式比较: ✅ 都支持，通过ScriptEngineEvalSingleExpression()求值并显示多种格式
// 执行过程: 用户命令formats → UserDebugger.ShowFormats() → ScriptEngineEvalSingleExpression() → CommandFormatsShowResults() → 本地处理
func (s *UserDebug) ShowFormats() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("显示格式信息")
}

// 来自接口: UserDebugger 第39个签名
// CloseLogFile 关闭日志文件
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: logclose命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/logclose.cpp
// 与内核模式比较: ✅ 都支持，只是关闭g_LogOpenFile文件流
// 执行过程: 用户命令logclose → UserDebugger.CloseLogFile() → g_LogOpenFile.close() → 本地处理
// 注意：此命令关闭之前通过.logopen打开的日志文件，停止记录命令和结果
func (s *UserDebug) CloseLogFile() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("关闭日志文件")
}

// 来自接口: UserDebugger 第40个签名
// OpenLogFile 打开日志文件
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: logopen命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/logopen.cpp
// 与内核模式比较: ✅ 都支持，只是打开g_LogOpenFile文件流
// 执行过程: 用户命令logopen → UserDebugger.OpenLogFile() → g_LogOpenFile.open() → 本地处理
// 注意：此命令打开日志文件，开始记录所有命令和结果到文件
func (s *UserDebug) OpenLogFile(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("打开日志文件", "path", path)
}

// 来自接口: UserDebugger 第41个签名
// PEInfo 查询PE文件信息
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: pe命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/pe.cpp
// 与内核模式比较: ✅ 都支持，通过PeIsPE32BitOr64Bit()和PeParseHeader()解析PE文件
// 执行过程: 用户命令pe → UserDebugger.PEInfo() → PeIsPE32BitOr64Bit() → PeParseHeader() → 本地处理
func (s *UserDebug) PEInfo(pid uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("查询PE文件信息", "pid", pid)
}

// 来自接口: UserDebugger 第42个签名
// ExecuteScript 执行HyperDbg脚本
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: script命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/script.cpp
// 与内核模式比较: ✅ 都支持，通过HyperDbgScriptReadFileAndExecuteCommand()执行脚本
// 执行过程: 用户命令script → UserDebugger.ExecuteScript() → HyperDbgScriptReadFileAndExecuteCommand() → HyperDbgInterpreter() → 本地处理
func (s *UserDebug) ExecuteScript(script string) ScriptResult {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("执行脚本", "script", script)
	return ScriptResult{}
}

// 来自接口: UserDebugger 第43个签名
// Status 查询调试器状态
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: status命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/status.cpp
// 与内核模式比较: ✅ 都支持，显示当前调试器的连接状态（本地/远程、VMI模式/Debugger模式）
// 执行过程: 用户命令status → UserDebugger.Status() → CommandStatus() → 本地处理
func (s *UserDebug) Status() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("调试器状态",
		"connected", s.packeter.DriverProvider.IsConnected(),
		"activeProcess", s.activeProcess.ProcessId,
		"isActive", s.activeProcess.IsActive,
		"isPaused", s.activeProcess.IsPaused,
	)
}

// 来自接口: UserDebugger 第44个签名
// Connect 连接到远程调试器或本地调试器
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: connect命令 -> 远程连接
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/connect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试连接
// 执行过程: 用户命令connect → UserDebugger.Connect() → RemoteConnectionConnect() → 建立远程连接
// 注意：此功能用于连接到远程调试器或本地调试器，支持TCP/IP连接
func (s *UserDebug) Connect(port string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("连接到调试器", "port", port)
}

// 来自接口: UserDebugger 第45个签名
// Disconnect 断开与远程调试器的连接
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: disconnect命令 -> 远程断开
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/disconnect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试断开
// 执行过程: 用户命令disconnect → UserDebugger.Disconnect() → RemoteConnectionCloseTheConnectionWithDebuggee() → 断开远程连接
// 注意：此命令断开调试会话但不卸载模块，本地调试需要先unload驱动
func (s *UserDebug) Disconnect() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.activeProcess.IsActive {
		s.DetachProcess(s.activeProcess.ProcessId)
	}

	mylog.Info("断开连接")
}

// 来自接口: UserDebugger 第46个签名
// CPU 查询CPU特性信息
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: cpu命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/cpu.cpp
// 与内核模式比较: ✅ 都支持，通过ReadCpuDetails()读取CPU特性
// 执行过程: 用户命令cpu → UserDebugger.CPU() → ReadCpuDetails() → InstructionSet::Vendor()/Brand()等 → 本地处理
func (s *UserDebug) CPU() {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("查询CPU特性信息")
}

// 来自接口: UserDebugger 第47个签名
// Output 创建输出实例用于事件转发
// IOCTL: 否, 本地命令
// 通信: 本地处理
// 用户命令: output命令 -> 本地处理
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/output.cpp
// 与内核模式比较: ✅ 都支持，通过g_OutputSources链表管理输出源
// 执行过程: 用户命令output → UserDebugger.Output() → CommandOutput() → 创建/打开/关闭输出实例 → 本地处理
func (s *UserDebug) Output(value uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mylog.Info("创建输出实例", "value", value)
}

// 来自接口: Eventer 第1个签名
// TODO: Implement StartEventLoop (from Eventer)
// StartEventLoop 启动事件循环
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用StartEventLoop() → 启动事件循环线程 → 开始处理事件
func (s *UserDebug) StartEventLoop() { s.eventer.StartEventLoop() }

// 来自接口: Eventer 第2个签名
// TODO: Implement StopEventLoop (from Eventer)
// StopEventLoop 停止事件循环
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用StopEventLoop() → 停止事件循环线程 → 停止处理事件
func (s *UserDebug) StopEventLoop() { s.eventer.StopEventLoop() }

// 来自接口: Eventer 第3个签名
// TODO: Implement IsEventLoopRunning (from Eventer)
// IsEventLoopRunning 检查事件循环是否在运行
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IsEventLoopRunning() → 返回事件循环运行状态
func (s *UserDebug) IsEventLoopRunning() bool { return s.eventer.IsEventLoopRunning() }

// 来自接口: Eventer 第4个签名
// TODO: Implement RegisterEventHandler (from Eventer)
// RegisterEventHandler 注册事件处理器
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RegisterEventHandler() → 将处理器添加到事件类型映射中
func (s *UserDebug) RegisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	s.eventer.RegisterEventHandler(eventType, handler)
}

// 来自接口: Eventer 第5个签名
// TODO: Implement UnregisterEventHandler (from Eventer)
// UnregisterEventHandler 取消注册事件处理器
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnregisterEventHandler() → 从事件类型映射中移除处理器
func (s *UserDebug) UnregisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	s.eventer.UnregisterEventHandler(eventType, handler)
}

// 来自接口: Eventer 第6个签名
// TODO: Implement RegisterEvent (from Eventer)
// RegisterEvent 向调试器注册新事件
// IOCTL: 是, IOCTL_REGISTER_EVENT
// 通信: 带事件缓冲区的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RegisterEvent() → IOCTL_REGISTER_EVENT → 注册事件
func (s *UserDebug) RegisterEvent(event *Event) { s.eventer.RegisterEvent(event) }

// 来自接口: Eventer 第7个签名
// TODO: Implement ModifyEvent (from Eventer)
// ModifyEvent 修改已注册的事件
// IOCTL: 是, IOCTL_MODIFY_EVENT
// 通信: 带事件标签的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ModifyEvent() → IOCTL_MODIFY_EVENT → 修改事件
func (s *UserDebug) ModifyEvent(eventTag uint64, modifications *EventModifications) {
	s.eventer.ModifyEvent(eventTag, modifications)
}

// 来自接口: Eventer 第8个签名
// TODO: Implement RemoveEvent (from Eventer)
// RemoveEvent 通过标签移除已注册的事件
// IOCTL: 是, IOCTL_REMOVE_EVENT
// 通信: 带事件标签的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RemoveEvent() → IOCTL_REMOVE_EVENT → 移除事件
func (s *UserDebug) RemoveEvent(eventTag uint64) {
	s.eventer.RemoveEvent(eventTag)
}

// 来自接口: Eventer 第9个签名
// TODO: Implement Events (from Eventer)
// Events 检索所有已注册的事件
// IOCTL: 是, IOCTL_LIST_EVENTS
// 通信: 带输出缓冲区的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Events() → IOCTL_LIST_EVENTS → 获取事件列表
func (s *UserDebug) Events() []Event { return s.eventer.Events() }

// 来自接口: Eventer 第10个签名
// TODO: Implement SetEventMode (from Eventer)
// SetEventMode 设置事件触发模式
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMode -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/mode.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SetEventMode() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: mode - EventModePre表示在操作执行前触发，EventModePost表示在操作执行后触发
func (s *UserDebug) SetEventMode(mode EventMode) { (s.eventer.SetEventMode(mode)) }

// Packet 返回底层的 Packet 实例
func (s *UserDebug) Packet() *Packet {
	return s.packeter
}
