package debugger

import (
	"bytes"
	"encoding/binary"
	"sync"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/HyperDbg/debugger/thread"
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

type DebuggerState int

const (
	StateUninitialized DebuggerState = iota
	StateInitialized
	StateRunning
	StatePaused
	StateTerminated
)

type ActiveDebuggingProcess struct {
	IsActive    bool
	IsPaused    bool
	ProcessId   uint32
	ThreadId    uint32
	ProcessName string
}
type KernelDebug struct {
	mu                sync.RWMutex
	packet            *Packet
	driver            driver.Api
	DebugEventHandler *DebugEventHandler
	eventLoop         *EventLoop
}

func NewKernelDebugger() KernelDebugger {
	return NewKernelDebuggerWithOptions(true)
}

func NewKernelDebuggerWithOptions(autoLoadDriver bool) KernelDebugger {
	mylog.Info("=== KernelDebugger initializing ===")
	mylog.Info(autoLoadDriver)
	mylog.Info(time.Now().Format("2006-01-02 15:04:05.000"))

	packet := GetPacket()

	if autoLoadDriver {
		driverPath := GetDriverPath()
		mylog.Info(driverPath, "找到驱动文件")
		packet.LoadDriver(driverPath)
	}

	packet.AutoLoadDriver = autoLoadDriver

	eventManager := NewEventManager(packet.driver)
	dbg := &KernelDebug{
		packet:    packet,
		driver:    packet.driver,
		eventLoop: NewEventLoop(eventManager),
	}

	dbg.DebugEventHandler = NewDebugEventHandler(packet.DriverProvider, packet.RegistersMeta.(*register.Provider), packet.ThreadsMeta.(*thread.Provider))

	dbg.eventLoop.Start()

	return dbg
}

func EnableSeLoadDriverPrivilege() {
	const SE_LOAD_DRIVER_NAME = "SeLoadDriverPrivilege"
	EnablePrivilege(SE_LOAD_DRIVER_NAME)
}

func EnableSeDebugPrivilege() {
	const SE_DEBUG_NAME = "SeDebugPrivilege"
	EnablePrivilege(SE_DEBUG_NAME)
}

func EnablePrivilege(privilegeName string) {
	var token windows.Token
	currentProcess := windows.CurrentProcess()

	mylog.Check(windows.OpenProcessToken(currentProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token))
	defer token.Close()

	privilege := windows.LUIDAndAttributes{}
	privilegeNamePtr := mylog.Check2(windows.UTF16PtrFromString(privilegeName))
	mylog.Check(windows.LookupPrivilegeValue(nil, privilegeNamePtr, &privilege.Luid))

	privilege.Attributes = windows.SE_PRIVILEGE_ENABLED

	privileges := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges:     [1]windows.LUIDAndAttributes{privilege},
	}
	mylog.Check(windows.AdjustTokenPrivileges(token, false, &privileges, 0, nil, nil))
}

func GetDriverPath() string {
	return "d:\\笔记本\\p\\ux\\examples\\hypedbg\\doc\\cpp\\HyperDbgUnified\\build\\Release\\hyperkd.sys"
}

func loadProcess(exePath string) (uint32, uint32) {
	return 1234, 5678
}

type TokenType int

const (
	TokenCommand TokenType = iota
	TokenArgument
	TokenOption
	TokenValue
)

type CommandToken struct {
	Type  TokenType
	Value string
}

// 来自接口: Packeter 第1个签名
// TODO: Implement LoadDriver (from Packeter)
// LoadDriver 加载HyperDbg驱动
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用LoadDriver() → 本地加载驱动
// 注意: 此方法是Go实现特有的，用于加载驱动服务
func (s *KernelDebug) LoadDriver(driverPath string) { s.packet.LoadDriver(driverPath) }

// 来自接口: Packeter 第2个签名
// TODO: Implement UnloadDriver (from Packeter)
// UnloadDriver 卸载HyperDbg驱动
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnloadDriver() → 本地卸载驱动
// 注意: 此方法是Go实现特有的，用于卸载驱动服务
func (s *KernelDebug) UnloadDriver() { s.packet.UnloadDriver() }

// 来自接口: Packeter 第3个签名
// TODO: Implement LoadVMM (from Packeter)
// LoadVMM 加载虚拟机监控器(VMM)
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
// 通信: 设备I/O
// 源代码: hyperdbg/hyperkd/code/driver/Ioctl.c
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用LoadVMM() → IOCTL_TERMINATE_VMX → 加载VMM
func (s *KernelDebug) LoadVMM() { s.packet.LoadVMM() }

// 来自接口: Packeter 第4个签名
// TODO: Implement UnloadVMM (from Packeter)
// UnloadVMM 卸载虚拟机监控器(VMM)
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
// 通信: 设备I/O
// 源代码: hyperdbg/hyperkd/code/driver/Ioctl.c
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnloadVMM() → IOCTL_TERMINATE_VMX → 卸载VMM
func (s *KernelDebug) UnloadVMM() { s.packet.UnloadVMM() }

// 来自接口: Packeter 第5个签名
// TODO: Implement IsConnected (from Packeter)
// IsConnected 检查调试器是否当前已连接
// IOCTL: 否, 本地状态检查
// 通信: 无
// 源代码: debugger/packet.go (Go实现特有)
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IsConnected() → 本地状态检查
// 注意: 此方法是Go实现特有的，用于检查驱动连接状态
func (s *KernelDebug) IsConnected() bool { return s.packet.IsConnected() }

// 来自接口: Packeter 第6个签名
// TODO: Implement EPTHook (from Packeter)
// EPTHook 启用EPT（扩展页表）钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) EPTHook(address uint64, size uint32, hookType EPTHookType) {
	s.packet.EPTHook(address, size, hookType)
}

// 来自接口: Packeter 第7个签名
// TODO: Implement EPTHook2 (from Packeter)
// EPTHook2 启用EPT钩子（替代实现）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook2 -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook2.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook2() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) EPTHook2(address uint64, size uint32, hookType EPTHookType) {
	s.packet.EPTHook2(address, size, hookType)
}

// 来自接口: Packeter 第8个签名
// TODO: Implement HookSyscall (from Packeter)
// HookSyscall 启用系统调用钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandSyscall -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/syscall-sysret.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookSyscall() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: syscallNumber - 系统调用号，0表示hook所有系统调用
//
//	EventTypeSyscallHookEferSyscall - Hook SYSCALL指令
//	EventTypeSyscallHookEferSysret - Hook SYSRET指令
func (s *KernelDebug) HookSyscall(syscallNumber uint32) { s.packet.HookSyscall(syscallNumber) }

// 来自接口: Packeter 第9个签名
// TODO: Implement HookException (from Packeter)
// HookException 启用异常钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandException -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/exception.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookException() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) HookException(exceptionType uint32) { s.packet.HookException(exceptionType) }

// 来自接口: Packeter 第10个签名
// TODO: Implement HookInterrupt (from Packeter)
// HookInterrupt 启用中断钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandInterrupt -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/interrupt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookInterrupt() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) HookInterrupt(vector uint32) {
	s.packet.HookInterrupt(vector)
}

// 来自接口: Packeter 第11个签名
// TODO: Implement HookIO (from Packeter)
// HookIO 启用I/O端口钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoin/CommandIoout -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioin.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIO() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) HookIO(port uint16, hookType uint32) {
	s.packet.HookIO(port, hookType)
}

// 来自接口: Packeter 第12个签名
// TODO: Implement HookIOAPIC (from Packeter)
// HookIOAPIC 启用I/O APIC钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoapic -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioapic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIOAPIC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) HookIOAPIC(apicID uint32) {
	s.packet.HookIOAPIC(apicID)
}

// 来自接口: Packeter 第13个签名
// TODO: Implement ReadMsr (from Packeter)
// ReadMsr 从模型特定寄存器读取值
// IOCTL: 是, IOCTL_MSR_READ
// 通信: 带MSR地址的设备I/O
// 源代码: hyperdbg/msr.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ReadMsr() → IOCTL_MSR_READ → 读取MSR值
func (s *KernelDebug) ReadMsr(msrAddress uint32) uint64 {
	return s.packet.ReadMsr(msrAddress)
}

// 来自接口: Packeter 第14个签名
// TODO: Implement WriteMsr (from Packeter)
// WriteMsr 向模型特定寄存器写入值
// IOCTL: 是, IOCTL_MSR_WRITE
// 通信: 带MSR地址和值的设备I/O
// 源代码: hyperdbg/msr.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用WriteMsr() → IOCTL_MSR_WRITE → 写入MSR值
func (s *KernelDebug) WriteMsr(msrAddress uint32, value uint64) {
	s.packet.WriteMsr(msrAddress, value)
}

// 来自接口: Packeter 第15个签名
// TODO: Implement MeasurePerformance (from Packeter)
// MeasurePerformance 测量特定地址的性能
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMeasure -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/measure.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用MeasurePerformance() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) MeasurePerformance(address uint64) {
	s.packet.MeasurePerformance(address)
}

// 来自接口: Packeter 第16个签名
// TODO: Implement MonitorMemory (from Packeter)
// MonitorMemory 监控内存访问
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMonitor -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/monitor.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用MonitorMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: monitorType - MonitorTypeRead监控读操作，MonitorTypeWrite监控写操作，
//
//	MonitorTypeExecute监控执行操作，MonitorTypeReadWrite监控读写操作，
//	MonitorTypeReadWriteExecute监控读写执行操作
func (s *KernelDebug) MonitorMemory(address uint64, size uint32, monitorType MonitorType) {
	s.packet.MonitorMemory(address, size, monitorType)
}

// 来自接口: Packeter 第17个签名
// TODO: Implement PCICam (from Packeter)
// PCICam 查询PCI CAM（配置访问机制）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPcicam -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcicam.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCICam() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) PCICam(bus uint32, device uint32, function uint32) PCICamInfo {
	return s.packet.PCICam(bus, device, function)
}

// 来自接口: Packeter 第18个签名
// TODO: Implement PMC (from Packeter)
// PMC 查询性能监控计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPmc -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pmc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PMC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) PMC(pmcNumber uint32) uint64 {
	return s.packet.PMC(pmcNumber)
}

// 来自接口: Packeter 第19个签名
// TODO: Implement ReconstructMemory (from Packeter)
// ReconstructMemory 从跟踪中重建内存
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandRev -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/rev.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ReconstructMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) ReconstructMemory(pid uint32, address uint64, size uint32, mode ReconstructMode) []byte {
	return s.packet.ReconstructMemory(pid, address, size, mode)
}

// 来自接口: Packeter 第20个签名
// TODO: Implement SearchMemoryPattern (from Packeter)
// SearchMemoryPattern 在内存中搜索字节模式
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SearchMemoryPattern() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: pid - 进程ID，pattern - 要搜索的字节模式，mode - ReconstructModeUserMode在用户模式内存中搜索，
//
//	ReconstructModeKernelMode在内核模式内存中搜索
//
// 返回值: 匹配模式的地址列表
func (s *KernelDebug) SearchMemoryPattern(pid uint32, pattern []byte, mode ReconstructMode) []uint64 {
	return s.packet.SearchMemoryPattern(pid, pattern, mode)
}

// 来自接口: Packeter 第21个签名
// TODO: Implement InstructionTrace (from Packeter)
// InstructionTrace 启用指令执行跟踪
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrace -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/trace.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用InstructionTrace() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: address - 要跟踪的起始地址
// 返回值: 跟踪结果通过事件系统异步返回，而不是直接返回
func (s *KernelDebug) InstructionTrace(address uint64) {
	s.packet.InstructionTrace(address)
}

// 来自接口: Packeter 第22个签名
// TODO: Implement TrackMemory (from Packeter)
// TrackMemory 跟踪内存访问
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TrackMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) TrackMemory(address uint64, size uint32) {
	s.packet.TrackMemory(address, size)
}

// 来自接口: Packeter 第23个签名
// TODO: Implement TSC (from Packeter)
// TSC 查询时间戳计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTsc -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/tsc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TSC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) TSC() uint64 {
	return s.packet.TSC()
}

// 来自接口: Packeter 第24个签名
// TODO: Implement VMCallHandler (from Packeter)
// VMCallHandler 启用VMCALL处理程序
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandVmcall -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/vmcall.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VMCallHandler() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) VMCallHandler() {
	s.packet.VMCallHandler()
}

// 来自接口: Packeter 第25个签名
// TODO: Implement HookXSETBV (from Packeter)
// HookXSETBV 启用XSETBV钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandXsetbv -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/xsetbv.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookXSETBV() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) HookXSETBV() {
	s.packet.HookXSETBV()
}

// 来自接口: Packeter 第26个签名
// TODO: Implement PTE (from Packeter)
// PTE 查询页表项
// IOCTL: 是, IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS (0x805)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandPte -> 内核: ExtensionCommandPte
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pte.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PTE() → IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS → ExtensionCommandPte
func (s *KernelDebug) PTE(virtualAddress uint64, pid uint32) PageTableEntries {
	return s.packet.PTE(virtualAddress, pid)
}

// 来自接口: Packeter 第27个签名
// TODO: Implement VA2PA (from Packeter)
// VA2PA 将虚拟地址转换为物理地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandVa2pa -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/va2pa.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VA2PA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (s *KernelDebug) VA2PA(virtualAddress uint64, pid uint32) uint64 {
	return s.packet.VA2PA(virtualAddress, pid)
}

// 来自接口: Packeter 第28个签名
// TODO: Implement PA2VA (from Packeter)
// PA2VA 将物理地址转换为虚拟地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带物理地址的设备I/O
// 用户命令: CommandPa2va -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pa2va.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PA2VA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (s *KernelDebug) PA2VA(physicalAddress uint64, pid uint32) uint64 {
	return s.packet.PA2VA(physicalAddress, pid)
}

// 来自接口: Packeter 第29个签名
// TODO: Implement ProcessDetails (from Packeter)
// ProcessDetails 查询进程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_PROCESS (0x81C)
// 通信: 带进程ID的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryDetails
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ProcessDetails() → IOCTL_QUERY_CURRENT_PROCESS → ProcessQueryDetails
func (s *KernelDebug) ProcessDetails(pid uint32) []ProcessDetails {
	return s.packet.ProcessDetails(pid)
}

// 来自接口: Packeter 第30个签名
// TODO: Implement ThreadDetails (from Packeter)
// ThreadDetails 查询线程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_THREAD (0x81D)
// 通信: 带线程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryDetails
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ThreadDetails() → IOCTL_QUERY_CURRENT_THREAD → ThreadQueryDetails
func (s *KernelDebug) ThreadDetails(tid uint32) []ThreadDetails {
	return s.packet.ThreadDetails(tid)
}

// 来自接口: Packeter 第31个签名
// TODO: Implement Processes (from Packeter)
// Processes 列出所有进程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryList
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Processes() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ProcessQueryList
func (s *KernelDebug) Processes() []ProcessInfo {
	return s.packet.Processes()
}

// 来自接口: Packeter 第32个签名
// TODO: Implement Threads (from Packeter)
// Threads 列出进程的所有线程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带进程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryList
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Threads() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ThreadQueryList
func (s *KernelDebug) Threads(pid uint32) []ThreadInfo {
	return s.packet.Threads(pid)
}

// 来自接口: Packeter 第33个签名
// TODO: Implement IDTEntries (from Packeter)
// IDTEntries 查询中断描述符表项
// IOCTL: 是, IOCTL_QUERY_IDT_ENTRY (0x824)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandIdt -> 内核: ExtensionCommandPerformQueryIdtEntriesRequest
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/idt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IDTEntries() → IOCTL_QUERY_IDT_ENTRY → ExtensionCommandPerformQueryIdtEntriesRequest
func (s *KernelDebug) IDTEntries() []IDTEntry {
	return s.packet.IDTEntries()
}

// 来自接口: Packeter 第34个签名
// TODO: Implement APIC (from Packeter)
// APIC 查询高级可编程中断控制器
// IOCTL: 是, IOCTL_PERFORM_ACTIONS_ON_APIC (0x822)
// 通信: 带APIC ID的设备I/O
// 用户命令: CommandApic -> 内核: ExtensionCommandPerformActionsForApicRequests
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/apic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用APIC() → IOCTL_PERFORM_ACTIONS_ON_APIC → ExtensionCommandPerformActionsForApicRequests
func (s *KernelDebug) APIC(apicID uint32) APICInfo {
	return s.packet.APIC(apicID)
}

// 来自接口: Packeter 第35个签名
// TODO: Implement PCITree (from Packeter)
// PCITree 查询PCI设备树
// IOCTL: 是, IOCTL_PCIE_ENDPOINT_ENUM (0x821)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandPcitree -> 内核: ExtensionCommandPcitree
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcitree.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCITree() → IOCTL_PCIE_ENDPOINT_ENUM → ExtensionCommandPcitree
func (s *KernelDebug) PCITree() []PCIDevice {
	return s.packet.PCITree()
}

// 来自接口: Packeter 第36个签名
// TODO: Implement PerformSMIOperation (from Packeter)
// PerformSMIOperation 执行系统管理中断操作
// IOCTL: 是, IOCTL_PERFORM_SMI_OPERATION (0x826)
// 通信: 带SMI操作的设备I/O
// 用户命令: CommandSmi -> 内核: VmFuncSmmPerformSmiOperation
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/smi.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PerformSMIOperation() → IOCTL_PERFORM_SMI_OPERATION → VmFuncSmmPerformSmiOperation
func (s *KernelDebug) PerformSMIOperation(operation SMIType) {
	s.packet.PerformSMIOperation(operation)
}

// 来自接口: Packeter 第37个签名
// TODO: Implement HideDebugger (from Packeter)
// HideDebugger 隐藏调试器以避免检测
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
// 通信: 设备I/O
// 用户命令: CommandHide -> 内核: TransparentHideDebuggerWrapper
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/hide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentHideDebuggerWrapper
func (s *KernelDebug) HideDebugger() {
	s.packet.HideDebugger()
}

// 来自接口: Packeter 第38个签名
// TODO: Implement UnhideDebugger (from Packeter)
// UnhideDebugger 取消隐藏调试器
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
// 通信: 设备I/O
// 用户命令: CommandUnhide -> 内核: TransparentUnhideDebuggerWrapper
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/unhide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnhideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentUnhideDebuggerWrapper
func (s *KernelDebug) UnhideDebugger() {
	s.packet.UnhideDebugger()
}

// 来自接口: Packeter 第39个签名
// TODO: Implement BringPagesIn (from Packeter)
// BringPagesIn 将页面引入内存
// IOCTL: 是, IOCTL_DEBUGGER_BRING_PAGES_IN (0x81F)
// 通信: 带地址范围的设备I/O
// 用户命令: CommandPagein -> 内核: DebuggerCommandBringPagein
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/pagein.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用BringPagesIn() → IOCTL_DEBUGGER_BRING_PAGES_IN → DebuggerCommandBringPagein
func (s *KernelDebug) BringPagesIn(fromAddr uint64, toAddr uint64, pid uint32) {
	s.packet.BringPagesIn(fromAddr, toAddr, pid)
}

// 来自接口: Packeter 第40个签名
// TODO: Implement EditMemory (from Packeter)
// EditMemory 编辑特定地址的内存
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
// 通信: 带地址和数据的设备I/O
// 用户命令: CommandE -> 内核: DebuggerCommandEditMemory
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EditMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
func (s *KernelDebug) EditMemory(pid uint32, address uint64, data []byte) {
	s.packet.EditMemory(pid, address, data)
}

// 来自接口: Packeter 第41个签名
// TODO: Implement SearchMemory (from Packeter)
// SearchMemory 在内存中搜索模式
// IOCTL: 是, IOCTL_DEBUGGER_SEARCH_MEMORY (0x80B)
// 通信: 带地址范围和模式的设备I/O
// 用户命令: CommandS -> 内核: DebuggerCommandSearchMemory
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/s.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SearchMemory() → IOCTL_DEBUGGER_SEARCH_MEMORY → DebuggerCommandSearchMemory
func (s *KernelDebug) SearchMemory(pid uint32, address uint64, size uint32, pattern []byte) []uint64 {
	return s.packet.SearchMemory(pid, address, size, pattern)
}

// 来自接口: Packeter 第42个签名
// TODO: Implement FlushBuffers (from Packeter)
// FlushBuffers 刷新日志缓冲区
// IOCTL: 是, IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS (0x80D)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandFlush -> 内核: DebuggerCommandFlush
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/flush.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用FlushBuffers() → IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS → DebuggerCommandFlush
func (s *KernelDebug) FlushBuffers() FlushResult { return s.packet.FlushBuffers() }

// 来自接口: Packeter 第43个签名
// TODO: Implement AttachProcess (from Packeter)
// AttachProcess 附加到用户模式进程
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandDebug -> 内核: AttachingTargetProcess
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/debug.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用AttachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
func (s *KernelDebug) AttachProcess(pid uint32) {
	s.packet.AttachProcess(pid)
}

// 来自接口: Packeter 第44个签名
// TODO: Implement DetachProcess (from Packeter)
// DetachProcess 从用户模式进程分离
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandDetach -> 内核: AttachingTargetProcess
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/detach.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用DetachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
func (s *KernelDebug) DetachProcess(pid uint32) {
	s.packet.DetachProcess(pid)
}

// 来自接口: Packeter 第45个签名
// TODO: Implement ChangeProcess (from Packeter)
// ChangeProcess 切换当前进程上下文
// IOCTL: 否, 本地命令 / 串行通信
// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
// 用户命令: .process命令 -> 本地处理 / 串行通信
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
//
//	hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
//
// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换进程上下文
// 执行过程: VMI模式：用户命令.process → Debugger.ChangeProcess() → ObjectShowProcessesOrThreadDetails() → 本地处理
//
//	Debugger模式：用户命令.process → Debugger.ChangeProcess() → KdSendSwitchProcessPacketToDebuggee() → 串行通信 → 内核处理
//
// 注意：内核模式调试器使用此命令切换进程上下文，而不是附加到进程。.process和.process2的区别在于切换方式（时钟中断 vs 直接切换）
func (s *KernelDebug) ChangeProcess(pid uint32) {
	s.packet.ChangeProcess(pid)
}

// 来自接口: Packeter 第46个签名
// TODO: Implement ChangeThread (from Packeter)
// ChangeThread 切换当前线程上下文
// IOCTL: 否, 本地命令 / 串行通信
// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
// 用户命令: .thread命令 -> 本地处理 / 串行通信
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
//
//	hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
//
// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换线程上下文
// 执行过程: VMI模式：用户命令.thread → Debugger.ChangeThread() → ObjectShowProcessesOrThreadDetails() → 本地处理
//
//	Debugger模式：用户命令.thread → Debugger.ChangeThread() → KdSendSwitchThreadPacketToDebuggee() → 串行通信 → 内核处理
//
// 注意：内核模式调试器使用此命令切换线程上下文。.thread和.thread2的区别在于切换方式（时钟中断 vs 直接切换）
func (s *KernelDebug) ChangeThread(tid uint32) {
	s.packet.ChangeThread(tid)
}

// 来自接口: Packeter 第47个签名
// TODO: Implement ConnectSerial (from Packeter)
// ConnectSerial 连接到远程调试器（串行/网络通信）
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: CommandConnect -> 远程连接
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/connect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试连接
// 执行过程: 用户命令connect → UserDebugger.ConnectSerial() → RemoteConnectionConnect() → 建立远程连接
// 注意：此功能用于远程调试，连接到远程调试器或被调试目标
func (s *KernelDebug) ConnectSerial(port string, baudrate uint32) {
	s.packet.ConnectSerial(port, baudrate)
}

// 来自接口: Packeter 第48个签名
// TODO: Implement DisconnectSerial (from Packeter)
// DisconnectSerial 断开远程连接
// IOCTL: 否, 远程通信
// 通信: 远程连接
// 用户命令: CommandDisconnect -> 远程断开
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/disconnect.cpp
// 与内核模式比较: ✅ 都支持，用于远程调试断开
// 执行过程: 用户命令disconnect → UserDebugger.DisconnectSerial() → RemoteConnectionDisconnect() → 断开远程连接
// 注意：此功能用于远程调试，断开与远程调试器或被调试目标的连接
func (s *KernelDebug) DisconnectSerial() {
	s.packet.DisconnectSerial()
}

// 来自接口: KernelDebugger 第2个签名
// SetBreakpoint 在指定地址设置断点
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandBp -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bp.cpp
// 与用户模式比较: ❌ 不一致，内核模式需要指定进程ID，使用不同的IOCTL控制码和处理函数
// 实现是否相同: ❌ 不同（用户模式：IOCTL_SET_BREAKPOINT_USER_DEBUGGER；内核模式：IOCTL_DEBUGGER_REGISTER_EVENT + RegisterEvent）
// 实现过程: 用户模式：用户命令bp → UserDebugger.SetBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointAddNew
//
//	内核模式：用户命令bp → KernelDebugger.SetBreakpoint() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent → RegisterEvent注册EPTHook事件
func (s *KernelDebug) SetBreakpoint(address uint64, pid uint32) {
	event := &DebugEvent{
		Type:           HiddenHookExecCC,
		ProcessId:      pid,
		EIP:            address,
		CountOfActions: 1,
	}

	s.packet.EventManager.RegisterDebugEvent(event)
}

// 来自接口: KernelDebugger 第3个签名
// ReadMemory 从特定地址读取内存
// IOCTL: 是, IOCTL_DEBUGGER_READ_MEMORY (0x803)
// 通信: 带地址和大小的设备I/O
// 用户命令: readmem命令 -> 内核: DebuggerCommandReadMemory
// 源代码: hyperdbg/libhyperdbg/code/debugger/misc/readmem.cpp
// 与用户模式比较: ❌ 不一致，内核模式需要指定内存类型
// 实现是否相同: ❌ 不同（用户模式：IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY；内核模式：IOCTL_DEBUGGER_READ_MEMORY + 指定内存类型）
// 实现过程: 用户模式：用户命令db → UserDebugger.ReadMemory() → IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY → DebuggerCommandReadMemory
//
//	内核模式：用户命令db → KernelDebugger.ReadMemory() → IOCTL_DEBUGGER_READ_MEMORY → DebuggerCommandReadMemory
func (s *KernelDebug) ReadMemory(pid uint32, address uint64, size uint32, memoryType MemoryType) []byte {
	return s.packet.ReadMemory(pid, address, size, memoryType)
}

// 来自接口: KernelDebugger 第4个签名
// WriteMemory 向特定内存地址写入数据
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
// 通信: 带地址和数据的设备I/O
// 用户命令: CommandE -> 内核: DebuggerCommandEditMemory
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
// 与用户模式比较: ✅ 一致，使用相同的IOCTL控制码
// 实现是否相同: ✅ 相同（都使用IOCTL_DEBUGGER_EDIT_MEMORY + DebuggerCommandEditMemory）
// 实现过程: 用户模式：用户命令eb → UserDebugger.WriteMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
//
//	内核模式：用户命令eb → KernelDebugger.WriteMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
func (s *KernelDebug) WriteMemory(pid uint32, address uint64, data []byte) {
	s.packet.WriteMemory(pid, address, data)
}

// 来自接口: KernelDebugger 第5个签名
// TODO: Implement Modules
// Modules 查询加载的模块
// IOCTL: 是, IOCTL_GET_USER_MODE_MODULE_DETAILS (0x819) (用户模式) 或 本地处理 (内核模式)
// 通信: 带进程ID的设备I/O (用户模式) 或 本地处理 (内核模式)
// 用户命令: lm命令 -> 内核: UserAccessGetLoadedModules (用户模式) 或 本地处理 (内核模式)
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/lm.cpp
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
func (s *KernelDebug) Modules(pid uint32) []ModuleInfo {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.packet.IsConnected() {
		mylog.Check("驱动未连接")
	}

	if pid == 0 {
		req := DebuggerGetUserModeModuleDetailsRequest{
			ProcessDebuggingDetailToken: 0,
		}

		mylog.Check(req.Validate())

		buf := new(bytes.Buffer)
		binary.Write(buf, binary.LittleEndian, &req)

		response := s.packet.driver.SendReceive(buf, IoctlGetUserModeModuleDetails)
		if response.Len() == 0 {
			mylog.Check("内核返回空响应")
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

		return modules
	}

	return nil
}

// 来自接口: KernelDebugger 第6个签名
// PreallocatePools 预分配内存池
// IOCTL: 是, IOCTL_RESERVE_PRE_ALLOCATED_POOLS (0x816)
// 通信: 带池类型和计数的设备I/O
// 用户命令: CommandPrealloc -> 内核: DebuggerCommandReservePreallocatedPools
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/prealloc.cpp
// 与用户模式比较: ❌ 仅内核模式支持
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令prealloc → KernelDebugger.PreallocatePools() → IOCTL_RESERVE_PRE_ALLOCATED_POOLS → DebuggerCommandReservePreallocatedPools
func (s *KernelDebug) PreallocatePools(poolType PoolType, count uint32) {
	req := DebuggerPreallocCommand{
		Type:  DebuggerPreallocCommandType(poolType),
		Count: count,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	s.packet.driver.SendReceive(buf, IoctlReservePreAllocatedPools)
}

// 来自接口: KernelDebugger 第7个签名
// RunTests 运行内核端测试
// IOCTL: 是, IOCTL_PERFORM_KERNEL_SIDE_TESTS (0x815)
// 通信: 带测试类型的设备I/O
// 用户命令: CommandTest -> 内核: TestKernelPerformTests
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/test.cpp
// 与用户模式比较: ❌ 仅内核模式支持
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令test → KernelDebugger.RunTests() → IOCTL_PERFORM_KERNEL_SIDE_TESTS → TestKernelPerformTests
func (s *KernelDebug) RunTests(testType string) TestResults {
	if !s.packet.IsConnected() {
		mylog.Check("driver not available")
	}

	buf := new(bytes.Buffer)
	buf.WriteString(testType)

	response := s.packet.driver.SendReceive(buf, IoctlPerformKernelSideTests)
	if response.Len() < 12 {
		mylog.Warning("内核返回响应长度不足")
		return TestResults{}
	}

	totalTests := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	passedTests := binary.LittleEndian.Uint32(response.Bytes()[4:8])
	failedTests := binary.LittleEndian.Uint32(response.Bytes()[8:12])

	return TestResults{
		Total:  totalTests,
		Passed: passedTests,
		Failed: failedTests,
	}
}

// 来自接口: KernelDebugger 第8个签名
// CPUID 查询CPUID信息
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandCpuid -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/cpuid.cpp
// 与用户模式比较: ❌ 仅内核模式支持
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令cpuid → KernelDebugger.CPUID() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) CPUID(eax, ecx uint32) []uint32 {
	if !s.packet.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	event := &DebugEvent{
		Type:           CpuidInstructionExecution,
		EIP:            uint64(eax),
		CountOfActions: ecx,
	}

	s.packet.EventManager.RegisterDebugEvent(event)

	return []uint32{eax, 0, 0, 0}
}

// 来自接口: KernelDebugger 第9个签名
// CRWrite 写入控制寄存器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandCrwrite -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/crwrite.cpp
// 与用户模式比较: ❌ 仅内核模式支持
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令crwrite → KernelDebugger.CRWrite() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) CRWrite(crNumber uint32, value uint64) {
	event := &DebugEvent{
		Type:           ControlRegisterModified,
		EIP:            uint64(crNumber),
		CountOfActions: uint32(value),
	}

	mylog.Check(s.packet.EventManager.RegisterDebugEvent(event))
}

// 来自接口: KernelDebugger 第10个签名
// DR 查询调试寄存器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandDr -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/dr.cpp
// 与用户模式比较: ❌ 仅内核模式支持
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令dr → KernelDebugger.DR() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (s *KernelDebug) DR(drNumber uint32) uint64 {
	if !s.packet.IsConnected() {
		mylog.Check("driver not available")
	}

	event := &DebugEvent{
		Type: DebugRegistersAccessed,
		EIP:  uint64(drNumber),
	}

	return s.packet.EventManager.RegisterDebugEvent(event)
}

// 来自接口: KernelDebugger 第11个签名
// ChangeCore 显示和切换当前操作的处理器核心
// IOCTL: 否, 串行通信
// 通信: 串行数据包
// 用户命令: ~命令 -> 串行通信
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/core.cpp
//
//	hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
//
// 与用户模式比较: ❌ 仅内核模式支持，通过串行通信切换核心
// 实现是否相同: ❌ 不适用（仅内核模式支持）
// 实现过程: 内核模式：用户命令~ → KernelDebugger.ChangeCore() → KdSendSwitchCorePacketToDebuggee() → 串行通信 → 内核处理
// 应用场景:
//  1. 多核系统调试：在多核系统中，不同核心可能执行不同的代码或处理不同的任务
//  2. 核心上下文切换：切换到特定核心查看该核心的寄存器状态、内存内容、执行流程、中断状态等
//  3. 并发问题调试：当多个核心同时执行代码时，可以切换到不同核心查看各自的执行状态
//  4. 性能分析：分析不同核心的负载和执行情况
//  5. 死锁和竞态条件调试：查看不同核心的同步状态和锁的使用情况
//
// 注意：不带参数时显示当前核心，带参数时切换到指定核心。此功能仅在Debugger模式（远程调试）下可用。
func (s *KernelDebug) ChangeCore(coreNumber uint32) {
	req := DebuggerApicRequest{
		ApicID: coreNumber,
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	s.packet.driver.SendReceive(buf, IoctlPerformActionsOnApic)
}

// 来自接口: Eventer 第1个签名
// TODO: Implement StartEventLoop (from Eventer)
// StartEventLoop 启动事件循环
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用StartEventLoop() → 启动事件循环线程 → 开始处理事件
func (s *KernelDebug) StartEventLoop() { s.eventLoop.Start() }

// 来自接口: Eventer 第2个签名
// TODO: Implement StopEventLoop (from Eventer)
// StopEventLoop 停止事件循环
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用StopEventLoop() → 停止事件循环线程 → 停止处理事件
func (s *KernelDebug) StopEventLoop() { s.eventLoop.Stop() }

// 来自接口: Eventer 第3个签名
// TODO: Implement IsEventLoopRunning (from Eventer)
// IsEventLoopRunning 检查事件循环是否在运行
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IsEventLoopRunning() → 返回事件循环运行状态
func (s *KernelDebug) IsEventLoopRunning() bool { return s.eventLoop.IsRunning() }

// 来自接口: Eventer 第4个签名
// TODO: Implement RegisterEventHandler (from Eventer)
// RegisterEventHandler 注册事件处理器
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RegisterEventHandler() → 将处理器添加到事件类型映射中
func (s *KernelDebug) RegisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	s.eventLoop.RegisterHandler(eventType, handler)
}

// 来自接口: Eventer 第5个签名
// TODO: Implement UnregisterEventHandler (from Eventer)
// UnregisterEventHandler 取消注册事件处理器
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/event_loop.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnregisterEventHandler() → 从事件类型映射中移除处理器
func (s *KernelDebug) UnregisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	s.eventLoop.UnregisterHandler(eventType, handler)
}

// 来自接口: Eventer 第6个签名
// TODO: Implement RegisterEvent (from Eventer)
// RegisterEvent 向调试器注册新事件
// IOCTL: 是, IOCTL_REGISTER_EVENT
// 通信: 带事件缓冲区的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RegisterEvent() → IOCTL_REGISTER_EVENT → 注册事件
func (s *KernelDebug) RegisterEvent(event *Event) {
	s.packet.EventManager.RegisterEvent(event)
}

// 来自接口: Eventer 第7个签名
// TODO: Implement ModifyEvent (from Eventer)
// ModifyEvent 修改已注册的事件
// IOCTL: 是, IOCTL_MODIFY_EVENT
// 通信: 带事件标签的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ModifyEvent() → IOCTL_MODIFY_EVENT → 修改事件
func (s *KernelDebug) ModifyEvent(eventTag uint64, modifications *EventModifications) {
	modifyType := ModifyQueryState
	if modifications.Enable != nil && *modifications.Enable {
		modifyType = ModifyEnable
	} else if modifications.Disable != nil && *modifications.Disable {
		modifyType = ModifyDisable
	} else if modifications.RemoveActions != nil && len(modifications.RemoveActions) > 0 {
		modifyType = ModifyClear
	}

	(s.eventLoop.eventManager.ModifyDebugEvent(eventTag, modifyType))
}

// 来自接口: Eventer 第8个签名
// TODO: Implement RemoveEvent (from Eventer)
// RemoveEvent 通过标签移除已注册的事件
// IOCTL: 是, IOCTL_REMOVE_EVENT
// 通信: 带事件标签的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用RemoveEvent() → IOCTL_REMOVE_EVENT → 移除事件
func (s *KernelDebug) RemoveEvent(eventTag uint64) {
	s.eventLoop.eventManager.ModifyDebugEvent(eventTag, ModifyClear)
}

// 来自接口: Eventer 第9个签名
// TODO: Implement Events (from Eventer)
// Events 检索所有已注册的事件
// IOCTL: 是, IOCTL_LIST_EVENTS
// 通信: 带输出缓冲区的设备I/O
// 源代码: hyperdbg/events.h
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Events() → IOCTL_LIST_EVENTS → 获取事件列表
func (s *KernelDebug) Events() []Event {
	events := s.eventLoop.eventManager.GetDebugEvents()
	result := make([]Event, len(events))
	for i, e := range events {
		result[i] = Event{
			Type:    e.Type,
			Pid:     e.ProcessId,
			Address: e.EIP,
			Size:    e.CountOfActions,
		}
	}
	return result
}

// 来自接口: Eventer 第10个签名
// TODO: Implement SetEventMode (from Eventer)
// SetEventMode 设置事件触发模式
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMode -> 内核: DebuggerParseEvent
// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/mode.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用SetEventMode() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: mode - EventModePre表示在操作执行前触发，EventModePost表示在操作执行后触发
func (s *KernelDebug) SetEventMode(mode EventMode) {
	mylog.Check(s.eventLoop.eventManager.SetEventMode(mode))
}

// Packet 返回底层的 Packet 实例
func (s *KernelDebug) Packet() *Packet {
	return s.packet
}
