package debugger

import "github.com/ddkwork/HyperDbg/debugger/register"

// Eventer 定义事件操作的接口
type Eventer interface {
	// ==================== 事件循环管理 ====================

	// StartEventLoop 启动事件循环
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/event_loop.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用StartEventLoop() → 启动事件循环线程 → 开始处理事件
	// 注意: 此方法是Go实现特有的，用于管理事件循环和事件处理器
	StartEventLoop()

	// StopEventLoop 停止事件循环
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/event_loop.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用StopEventLoop() → 停止事件循环线程 → 停止处理事件
	// 注意: 此方法是Go实现特有的，用于管理事件循环和事件处理器
	StopEventLoop()

	// IsEventLoopRunning 检查事件循环是否在运行
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/event_loop.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用IsEventLoopRunning() → 返回事件循环运行状态
	// 注意: 此方法是Go实现特有的，用于管理事件循环和事件处理器
	IsEventLoopRunning() bool

	// RegisterEventHandler 注册事件处理器
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/event_loop.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用RegisterEventHandler() → 将处理器添加到事件类型映射中
	// 注意: 此方法是Go实现特有的，用于注册事件处理器
	RegisterEventHandler(eventType EventType, handler func(*DebugEvent))

	// UnregisterEventHandler 取消注册事件处理器
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/event_loop.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用UnregisterEventHandler() → 从事件类型映射中移除处理器
	// 注意: 此方法是Go实现特有的，用于取消注册事件处理器
	UnregisterEventHandler(eventType EventType, handler func(*DebugEvent))

	// RegisterEvent 向调试器注册新事件
	// IOCTL: 是, IOCTL_REGISTER_EVENT (0x800) 或 IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件缓冲区的设备I/O
	// 用户命令: !mode命令 -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/mode.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用RegisterEvent() → IOCTL_REGISTER_EVENT/IOCTL_DEBUGGER_REGISTER_EVENT → 注册事件
	RegisterEvent(event *Event)

	// ModifyEvent 修改已注册的事件
	// IOCTL: 是, IOCTL_DEBUGGER_MODIFY_EVENTS (0x80C)
	// 通信: 带事件标签的设备I/O
	// 用户命令: events命令 -> 内核: CommandEventsModifyAndQueryEvents
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/events.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用ModifyEvent() → IOCTL_DEBUGGER_MODIFY_EVENTS → 修改事件
	ModifyEvent(eventTag uint64, modifications *EventModifications)

	// RemoveEvent 通过标签移除已注册的事件
	// IOCTL: 是, IOCTL_DEBUGGER_MODIFY_EVENTS (0x80C)
	// 通信: 带事件标签的设备I/O
	// 用户命令: events命令 -> 内核: CommandEventsModifyAndQueryEvents
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/events.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用RemoveEvent() → IOCTL_DEBUGGER_MODIFY_EVENTS → 移除事件
	RemoveEvent(eventTag uint64)

	// Events 检索所有已注册的事件
	// IOCTL: 是, IOCTL_DEBUGGER_MODIFY_EVENTS (0x80C)
	// 通信: 带输出缓冲区的设备I/O
	// 用户命令: events命令 -> 内核: CommandEventsModifyAndQueryEvents
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/events.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Events() → IOCTL_DEBUGGER_MODIFY_EVENTS → 获取事件列表
	Events() []Event

	// SetEventMode 设置事件触发模式
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandMode -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/mode.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用SetEventMode() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	// 参数说明: mode - EventModePre表示在操作执行前触发，EventModePost表示在操作执行后触发
	SetEventMode(mode EventMode)
}

// Packeter 定义HyperDbg调试器操作的公共接口，由packet实现，需要检查已经存在的多余方法，应该在别处实现
type Packeter interface {
	// LoadDriver 加载HyperDbg驱动
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用LoadDriver() → 本地加载驱动
	// 注意: 此方法是Go实现特有的，用于加载驱动服务
	LoadDriver(driverPath string)

	// UnloadDriver 卸载HyperDbg驱动
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用UnloadDriver() → 本地卸载驱动
	// 注意: 此方法是Go实现特有的，用于卸载驱动服务
	UnloadDriver()

	// LoadVMM 加载虚拟机监控器(VMM)
	// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
	// 通信: 设备I/O
	// 用户命令: 内核: DebuggerUninitialize
	// 源代码: hyperdbg/hyperkd/code/driver/Ioctl.c
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用LoadVMM() → IOCTL_TERMINATE_VMX → 加载VMM
	LoadVMM()

	// UnloadVMM 卸载虚拟机监控器(VMM)
	// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802)
	// 通信: 设备I/O
	// 用户命令: 内核: DebuggerUninitialize
	// 源代码: hyperdbg/hyperkd/code/driver/Ioctl.c
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用UnloadVMM() → IOCTL_TERMINATE_VMX → 卸载VMM
	UnloadVMM()

	// IsConnected 检查调试器是否当前已连接
	// IOCTL: 否, 本地状态检查
	// 通信: 无
	// 源代码: debugger/packet.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用IsConnected() → 本地状态检查
	// 注意: 此方法是Go实现特有的，用于检查驱动连接状态
	IsConnected() bool

	// ==================== Hook 相关 ====================

	// EPTHook 启用EPT（扩展页表）钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandEpthook -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用EPTHook() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	EPTHook(address uint64, size uint32, hookType EPTHookType)

	// EPTHook2 启用EPT钩子（替代实现）
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandEpthook2 -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook2.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用EPTHook2() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	EPTHook2(address uint64, size uint32, hookType EPTHookType)

	// HookSyscall 启用系统调用钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandSyscall -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/syscall-sysret.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookSyscall() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	// 参数说明: syscallNumber - 系统调用号，0表示hook所有系统调用
	//           EventTypeSyscallHookEferSyscall - Hook SYSCALL指令
	//           EventTypeSyscallHookEferSysret - Hook SYSRET指令
	HookSyscall(syscallNumber uint32)

	// HookException 启用异常钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandException -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/exception.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookException() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	HookException(exceptionType uint32)

	// HookInterrupt 启用中断钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandInterrupt -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/interrupt.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookInterrupt() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	HookInterrupt(vector uint32)

	// HookIO 启用I/O端口钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandIoin/CommandIoout -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioin.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookIO() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	HookIO(port uint16, hookType uint32)

	// HookIOAPIC 启用I/O APIC钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandIoapic -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioapic.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookIOAPIC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	HookIOAPIC(apicID uint32)

	// ReadMsr 从模型特定寄存器读取值
	// IOCTL: 是, IOCTL_DEBUGGER_READ_OR_WRITE_MSR (0x804)
	// 通信: 带MSR地址的设备I/O
	// 用户命令: rdmsr命令 -> 内核: DeviceIoControl
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/rdmsr.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用ReadMsr() → IOCTL_DEBUGGER_READ_OR_WRITE_MSR → 读取MSR值
	ReadMsr(msrAddress uint32) uint64

	// WriteMsr 向模型特定寄存器写入值
	// IOCTL: 是, IOCTL_DEBUGGER_READ_OR_WRITE_MSR (0x804)
	// 通信: 带MSR地址和值的设备I/O
	// 用户命令: wrmsr命令 -> 内核: DeviceIoControl
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/wrmsr.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用WriteMsr() → IOCTL_DEBUGGER_READ_OR_WRITE_MSR → 写入MSR值
	WriteMsr(msrAddress uint32, value uint64)

	// MeasurePerformance 测量特定地址的性能
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandMeasure -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/measure.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用MeasurePerformance() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	MeasurePerformance(address uint64)

	// MonitorMemory 监控内存访问
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandMonitor -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/monitor.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用MonitorMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	// 参数说明: monitorType - MonitorTypeRead监控读操作，MonitorTypeWrite监控写操作，
	//           MonitorTypeExecute监控执行操作，MonitorTypeReadWrite监控读写操作，
	//           MonitorTypeReadWriteExecute监控读写执行操作
	MonitorMemory(address uint64, size uint32, monitorType MonitorType)

	// PCICam 查询PCI CAM（配置访问机制）
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandPcicam -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcicam.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PCICam() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	PCICam(bus, device, function uint32) PCICamInfo

	// PMC 查询性能监控计数器
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandPmc -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pmc.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PMC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	PMC(pmcNumber uint32) uint64

	// ReconstructMemory 从跟踪中重建内存
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandRev -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/rev.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用ReconstructMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	ReconstructMemory(pid uint32, address uint64, size uint32, mode ReconstructMode) []byte

	// SearchMemoryPattern 在内存中搜索字节模式
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用SearchMemoryPattern() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	// 参数说明: pid - 进程ID，pattern - 要搜索的字节模式，mode - ReconstructModeUserMode在用户模式内存中搜索，
	//           ReconstructModeKernelMode在内核模式内存中搜索
	// 返回值: 匹配模式的地址列表
	SearchMemoryPattern(pid uint32, pattern []byte, mode ReconstructMode) []uint64

	// InstructionTrace 启用指令执行跟踪
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandTrace -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/trace.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用InstructionTrace() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	// 参数说明: address - 要跟踪的起始地址
	// 返回值: 跟踪结果通过事件系统异步返回，而不是直接返回
	InstructionTrace(address uint64)

	// TrackMemory 跟踪内存访问
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用TrackMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	TrackMemory(address uint64, size uint32)

	// TSC 查询时间戳计数器
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandTsc -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/tsc.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用TSC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	TSC() uint64

	// VMCallHandler 启用VMCALL处理程序
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandVmcall -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/vmcall.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用VMCallHandler() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	VMCallHandler()

	// HookXSETBV 启用XSETBV钩子
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandXsetbv -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/xsetbv.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HookXSETBV() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	HookXSETBV()

	// PTE 查询页表项
	// IOCTL: 是, IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS (0x805)
	// 通信: 带虚拟地址的设备I/O
	// 用户命令: CommandPte -> 内核: ExtensionCommandPte
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pte.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PTE() → IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS → ExtensionCommandPte
	PTE(virtualAddress uint64, pid uint32) PageTableEntries

	// VA2PA 将虚拟地址转换为物理地址
	// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
	// 通信: 带虚拟地址的设备I/O
	// 用户命令: CommandVa2pa -> 内核: ExtensionCommandVa2paAndPa2va
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/va2pa.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用VA2PA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
	VA2PA(virtualAddress uint64, pid uint32) uint64

	// PA2VA 将物理地址转换为虚拟地址
	// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
	// 通信: 带物理地址的设备I/O
	// 用户命令: CommandPa2va -> 内核: ExtensionCommandVa2paAndPa2va
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pa2va.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PA2VA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
	PA2VA(physicalAddress uint64, pid uint32) uint64

	// ProcessDetails 查询进程的详细信息
	// IOCTL: 是, IOCTL_QUERY_CURRENT_PROCESS (0x81C)
	// 通信: 带进程ID的设备I/O
	// 用户命令: process命令 -> 内核: ProcessQueryDetails
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用ProcessDetails() → IOCTL_QUERY_CURRENT_PROCESS → ProcessQueryDetails
	ProcessDetails(pid uint32) []ProcessDetails

	// ThreadDetails 查询线程的详细信息
	// IOCTL: 是, IOCTL_QUERY_CURRENT_THREAD (0x81D)
	// 通信: 带线程ID的设备I/O
	// 用户命令: thread命令 -> 内核: ThreadQueryDetails
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用ThreadDetails() → IOCTL_QUERY_CURRENT_THREAD → ThreadQueryDetails
	ThreadDetails(tid uint32) []ThreadDetails

	// Processes 列出所有进程
	// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
	// 通信: 带输出缓冲区的设备I/O
	// 用户命令: process命令 -> 内核: ProcessQueryList
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Processes() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ProcessQueryList
	Processes() []ProcessInfo

	// Threads 列出进程的所有线程
	// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
	// 通信: 带进程ID的设备I/O
	// 用户命令: thread命令 -> 内核: ThreadQueryList
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Threads() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ThreadQueryList
	Threads(pid uint32) []ThreadInfo

	// IDTEntries 查询中断描述符表项
	// IOCTL: 是, IOCTL_QUERY_IDT_ENTRY (0x824)
	// 通信: 带输出缓冲区的设备I/O
	// 用户命令: CommandIdt -> 内核: ExtensionCommandPerformQueryIdtEntriesRequest
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/idt.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用IDTEntries() → IOCTL_QUERY_IDT_ENTRY → ExtensionCommandPerformQueryIdtEntriesRequest
	IDTEntries() []IDTEntry

	// APIC 查询高级可编程中断控制器
	// IOCTL: 是, IOCTL_PERFORM_ACTIONS_ON_APIC (0x822)
	// 通信: 带APIC ID的设备I/O
	// 用户命令: CommandApic -> 内核: ExtensionCommandPerformActionsForApicRequests
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/apic.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用APIC() → IOCTL_PERFORM_ACTIONS_ON_APIC → ExtensionCommandPerformActionsForApicRequests
	APIC(apicID uint32) APICInfo

	// PCITree 查询PCI设备树
	// IOCTL: 是, IOCTL_PCIE_ENDPOINT_ENUM (0x821)
	// 通信: 带输出缓冲区的设备I/O
	// 用户命令: CommandPcitree -> 内核: ExtensionCommandPcitree
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcitree.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PCITree() → IOCTL_PCIE_ENDPOINT_ENUM → ExtensionCommandPcitree
	PCITree() []PCIDevice

	// PerformSMIOperation 执行系统管理中断操作
	// IOCTL: 是, IOCTL_PERFORM_SMI_OPERATION (0x826)
	// 通信: 带SMI操作的设备I/O
	// 用户命令: CommandSmi -> 内核: VmFuncSmmPerformSmiOperation
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/smi.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用PerformSMIOperation() → IOCTL_PERFORM_SMI_OPERATION → VmFuncSmmPerformSmiOperation
	PerformSMIOperation(operation SMIType)

	// HideDebugger 隐藏调试器以避免检测
	// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
	// 通信: 设备I/O
	// 用户命令: CommandHide -> 内核: TransparentHideDebuggerWrapper
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/hide.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用HideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentHideDebuggerWrapper
	HideDebugger()

	// UnhideDebugger 取消隐藏调试器
	// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x808)
	// 通信: 设备I/O
	// 用户命令: CommandUnhide -> 内核: TransparentUnhideDebuggerWrapper
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/unhide.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用UnhideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → TransparentUnhideDebuggerWrapper
	UnhideDebugger()

	// BringPagesIn 将页面引入内存
	// IOCTL: 是, IOCTL_DEBUGGER_BRING_PAGES_IN (0x81F)
	// 通信: 带地址范围的设备I/O
	// 用户命令: CommandPagein -> 内核: DebuggerCommandBringPagein
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/pagein.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用BringPagesIn() → IOCTL_DEBUGGER_BRING_PAGES_IN → DebuggerCommandBringPagein
	BringPagesIn(fromAddr, toAddr uint64, pid uint32)

	// EditMemory 编辑特定地址的内存
	// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
	// 通信: 带地址和数据的设备I/O
	// 用户命令: CommandE -> 内核: DebuggerCommandEditMemory
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用EditMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
	EditMemory(pid uint32, address uint64, data []byte)

	// SearchMemory 在内存中搜索模式
	// IOCTL: 是, IOCTL_DEBUGGER_SEARCH_MEMORY (0x80B)
	// 通信: 带地址范围和模式的设备I/O
	// 用户命令: CommandS -> 内核: DebuggerCommandSearchMemory
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/s.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用SearchMemory() → IOCTL_DEBUGGER_SEARCH_MEMORY → DebuggerCommandSearchMemory
	SearchMemory(pid uint32, address uint64, size uint32, pattern []byte) []uint64

	// FlushBuffers 刷新日志缓冲区
	// IOCTL: 是, IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS (0x80D)
	// 通信: 带输出缓冲区的设备I/O
	// 用户命令: CommandFlush -> 内核: DebuggerCommandFlush
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/flush.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用FlushBuffers() → IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS → DebuggerCommandFlush
	FlushBuffers() FlushResult

	// AttachProcess 附加到用户模式进程
	// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
	// 通信: 带进程ID的设备I/O
	// 用户命令: CommandDebug -> 内核: AttachingTargetProcess
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/debug.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用AttachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
	AttachProcess(pid uint32)

	// DetachProcess 从用户模式进程分离
	// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
	// 通信: 带进程ID的设备I/O
	// 用户命令: CommandDetach -> 内核: AttachingTargetProcess
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/detach.cpp
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用DetachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → AttachingTargetProcess
	DetachProcess(pid uint32)

	// ChangeProcess 切换当前进程上下文
	// IOCTL: 否, 本地命令 / 串行通信
	// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
	// 用户命令: .process命令 -> 本地处理 / 串行通信
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
	//         hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
	// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换进程上下文
	// 执行过程: VMI模式：用户命令.process → Debugger.ChangeProcess() → ObjectShowProcessesOrThreadDetails() → 本地处理
	//           Debugger模式：用户命令.process → Debugger.ChangeProcess() → KdSendSwitchProcessPacketToDebuggee() → 串行通信 → 内核处理
	// 注意：内核模式调试器使用此命令切换进程上下文，而不是附加到进程。.process和.process2的区别在于切换方式（时钟中断 vs 直接切换）
	ChangeProcess(pid uint32)

	// ChangeThread 切换当前线程上下文
	// IOCTL: 否, 本地命令 / 串行通信
	// 通信: 本地处理 (VMI模式) / 串行数据包 (Debugger模式)
	// 用户命令: .thread命令 -> 本地处理 / 串行通信
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
	//         hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
	// 与用户模式/内核模式比较: ✅ 都支持，VMI模式本地处理，Debugger模式通过串行通信切换线程上下文
	// 执行过程: VMI模式：用户命令.thread → Debugger.ChangeThread() → ObjectShowProcessesOrThreadDetails() → 本地处理
	//           Debugger模式：用户命令.thread → Debugger.ChangeThread() → KdSendSwitchThreadPacketToDebuggee() → 串行通信 → 内核处理
	// 注意：内核模式调试器使用此命令切换线程上下文。.thread和.thread2的区别在于切换方式（时钟中断 vs 直接切换）
	ChangeThread(tid uint32)

	// ConnectSerial 连接到远程调试器（串行/网络通信）
	// IOCTL: 否, 远程通信
	// 通信: 远程连接
	// 用户命令: CommandConnect -> 远程连接
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/connect.cpp
	// 与内核模式比较: ✅ 都支持，用于远程调试连接
	// 执行过程: 用户命令connect → UserDebugger.ConnectSerial() → RemoteConnectionConnect() → 建立远程连接
	// 注意：此功能用于远程调试，连接到远程调试器或被调试目标
	ConnectSerial(port string, baudrate uint32)

	// DisconnectSerial 断开远程连接
	// IOCTL: 否, 远程通信
	// 通信: 远程连接
	// 用户命令: CommandDisconnect -> 远程断开
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/disconnect.cpp
	// 与内核模式比较: ✅ 都支持，用于远程调试断开
	// 执行过程: 用户命令disconnect → UserDebugger.DisconnectSerial() → RemoteConnectionDisconnect() → 断开远程连接
	// 注意：此功能用于远程调试，断开与远程调试器或被调试目标的连接
	DisconnectSerial()

	// Packet 返回底层的 Packet 实例
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/packet.go (Go实现特有)
	// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Packet() → 返回底层的 *Packet 实例
	// 注意: 此方法是Go实现特有的，用于访问底层的 Packet 实例
	Packet() *Packet
}

// KernelDebugger 定义内核模式调试操作的接口
type KernelDebugger interface {
	Packeter

	// SetBreakpoint 在指定地址设置断点
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandBp -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bp.cpp
	// 与用户模式比较: ❌ 不一致，内核模式需要指定进程ID，使用不同的IOCTL控制码和处理函数
	// 实现是否相同: ❌ 不同（用户模式：IOCTL_SET_BREAKPOINT_USER_DEBUGGER；内核模式：IOCTL_DEBUGGER_REGISTER_EVENT + RegisterEvent）
	// 实现过程: 用户模式：用户命令bp → UserDebugger.SetBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointAddNew
	//           内核模式：用户命令bp → KernelDebugger.SetBreakpoint() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent → RegisterEvent注册EPTHook事件
	SetBreakpoint(address uint64, pid uint32)

	// ReadMemory 从特定地址读取内存
	// IOCTL: 是, IOCTL_DEBUGGER_READ_MEMORY (0x803)
	// 通信: 带地址和大小的设备I/O
	// 用户命令: readmem命令 -> 内核: DebuggerCommandReadMemory
	// 源代码: hyperdbg/libhyperdbg/code/debugger/misc/readmem.cpp
	// 与用户模式比较: ❌ 不一致，内核模式需要指定内存类型
	// 实现是否相同: ❌ 不同（用户模式：IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY；内核模式：IOCTL_DEBUGGER_READ_MEMORY + 指定内存类型）
	// 实现过程: 用户模式：用户命令db → UserDebugger.ReadMemory() → IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY → DebuggerCommandReadMemory
	//           内核模式：用户命令db → KernelDebugger.ReadMemory() → IOCTL_DEBUGGER_READ_MEMORY → DebuggerCommandReadMemory
	ReadMemory(pid uint32, address uint64, size uint32, memoryType MemoryType) []byte

	// WriteMemory 向特定内存地址写入数据
	// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
	// 通信: 带地址和数据的设备I/O
	// 用户命令: CommandE -> 内核: DebuggerCommandEditMemory
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
	// 与用户模式比较: ✅ 一致，使用相同的IOCTL控制码
	// 实现是否相同: ✅ 相同（都使用IOCTL_DEBUGGER_EDIT_MEMORY + DebuggerCommandEditMemory）
	// 实现过程: 用户模式：用户命令eb → UserDebugger.WriteMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
	//           内核模式：用户命令eb → KernelDebugger.WriteMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
	WriteMemory(pid uint32, address uint64, data []byte)

	// Modules 查询加载的模块
	// IOCTL: 是, IOCTL_GET_USER_MODE_MODULE_DETAILS (0x819) (用户模式) 或 本地处理 (内核模式)
	// 通信: 带进程ID的设备I/O (用户模式) 或 本地处理 (内核模式)
	// 用户命令: lm命令 -> 内核: UserAccessGetLoadedModules (用户模式) 或 本地处理 (内核模式)
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/lm.cpp
	//         CommandLmShowUserModeModule (用户模式)
	//         CommandLmShowKernelModeModule (内核模式)
	// 与内核模式比较: ✅ 都支持，用户模式通过IOCTL，内核模式通过SymbolCheckAndAllocateModuleInformation
	// 执行过程: 用户命令lm → UserDebugger.Modules() → CommandLmShowUserModeModule() → IOCTL_GET_USER_MODE_MODULE_DETAILS → UserAccessGetLoadedModules
	//           或 内核模式：用户命令lm → UserDebugger.Modules() → CommandLmShowKernelModeModule() → SymbolCheckAndAllocateModuleInformation() → 本地处理
	// 注意：UserDebugger 继承了 Public.Modules(pid uint32)，此方法用于获取所有模块
	Modules(pid uint32) []ModuleInfo

	// PreallocatePools 预分配内存池
	// IOCTL: 是, IOCTL_RESERVE_PRE_ALLOCATED_POOLS (0x816)
	// 通信: 带池类型和计数的设备I/O
	// 用户命令: CommandPrealloc -> 内核: DebuggerCommandReservePreallocatedPools
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/prealloc.cpp
	// 与用户模式比较: ❌ 仅内核模式支持
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令prealloc → KernelDebugger.PreallocatePools() → IOCTL_RESERVE_PRE_ALLOCATED_POOLS → DebuggerCommandReservePreallocatedPools
	PreallocatePools(poolType PoolType, count uint32)

	// RunTests 运行内核端测试
	// IOCTL: 是, IOCTL_PERFORM_KERNEL_SIDE_TESTS (0x815)
	// 通信: 带测试类型的设备I/O
	// 用户命令: CommandTest -> 内核: TestKernelPerformTests
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/test.cpp
	// 与用户模式比较: ❌ 仅内核模式支持
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令test → KernelDebugger.RunTests() → IOCTL_PERFORM_KERNEL_SIDE_TESTS → TestKernelPerformTests
	RunTests(testType string) TestResults

	// CPUID 查询CPUID信息
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandCpuid -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/cpuid.cpp
	// 与用户模式比较: ❌ 仅内核模式支持
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令cpuid → KernelDebugger.CPUID() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	CPUID(eax, ecx uint32) []uint32

	// CRWrite 写入控制寄存器
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandCrwrite -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/crwrite.cpp
	// 与用户模式比较: ❌ 仅内核模式支持
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令crwrite → KernelDebugger.CRWrite() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	CRWrite(crNumber uint32, value uint64)

	// DR 查询调试寄存器
	// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带事件结构的设备I/O
	// 用户命令: CommandDr -> 内核: DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/dr.cpp
	// 与用户模式比较: ❌ 仅内核模式支持
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令dr → KernelDebugger.DR() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
	DR(drNumber uint32) uint64

	// ChangeCore 显示和切换当前操作的处理器核心
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: ~命令 -> 串行通信
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/core.cpp
	//         hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp
	// 与用户模式比较: ❌ 仅内核模式支持，通过串行通信切换核心
	// 实现是否相同: ❌ 不适用（仅内核模式支持）
	// 实现过程: 内核模式：用户命令~ → KernelDebugger.ChangeCore() → KdSendSwitchCorePacketToDebuggee() → 串行通信 → 内核处理
	// 应用场景:
	//   1. 多核系统调试：在多核系统中，不同核心可能执行不同的代码或处理不同的任务
	//   2. 核心上下文切换：切换到特定核心查看该核心的寄存器状态、内存内容、执行流程、中断状态等
	//   3. 并发问题调试：当多个核心同时执行代码时，可以切换到不同核心查看各自的执行状态
	//   4. 性能分析：分析不同核心的负载和执行情况
	//   5. 死锁和竞态条件调试：查看不同核心的同步状态和锁的使用情况
	// 注意：不带参数时显示当前核心，带参数时切换到指定核心。此功能仅在Debugger模式（远程调试）下可用。
	ChangeCore(coreID uint32)

	// Packet 返回底层的 Packet 实例
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/kernel_debugger.go (Go实现特有)
	// 与用户模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Packet() → 返回底层的 *Packet 实例
	// 注意: 此方法是Go实现特有的，用于访问底层的 Packet 实例
	Packet() *Packet

	Eventer
}

// UserDebugger 定义用户模式调试操作的接口
type UserDebugger interface {
	Packeter
	// ==================== 进程管理相关 ====================

	// AttachProcess 附加到进程（VMI模式用户调试器）
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: .attach命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/attach.cpp
	// 与内核模式比较: ❌ 仅VMI模式用户调试器支持，内核模式使用.process命令切换进程上下文
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 用户命令.attach → UserDebugger.AttachProcess() → UdAttachToProcess() → 本地处理
	// 注意：此功能目前处于beta版本，主要用于VMI模式。内核模式调试器使用.process命令切换进程上下文，而不是附加到进程
	//       此函数与Debugger接口中的AttachProcess()不同，后者通过IOCTL附加到用户模式进程
	AttachProcess(pid uint32)

	// KillProcess 终止当前调试的进程
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: CommandKill -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/kill.cpp
	// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 用户命令kill → UserDebugger.KillProcess() → UdKillProcess() → 本地处理
	// 注意：此功能仅用于用户模式调试器，终止当前正在调试的进程
	KillProcess(pid uint32)

	// StartProcess 启动新进程
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: CommandStart -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/start.cpp
	// 与内核模式比较: ❌ 仅用户模式支持，内核模式无对应方法
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 用户命令start → UserDebugger.StartProcess() → 本地处理
	StartProcess(path string)

	// GetActiveDebuggingProcess 获取当前活动的调试进程
	// IOCTL: 否, 本地状态
	// 通信: 本地处理
	// 用户命令: 内部使用
	// 源代码: hyperdbg/libhyperdbg/code/debugger/user-level/ud.cpp
	// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 本地返回 activeProcess 状态
	// 注意：此功能仅用于用户模式调试器，获取当前正在调试的进程状态
	GetActiveDebuggingProcess() UserActiveDebuggingProcess

	// SetActiveDebuggingProcess 设置当前活动的调试进程
	// IOCTL: 否, 本地状态
	// 通信: 本地处理
	// 用户命令: 内部使用
	// 源代码: hyperdbg/libhyperdbg/code/debugger/user-level/ud.cpp
	// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 本地设置 activeProcess 状态
	// 注意：此功能仅用于用户模式调试器，设置当前正在调试的进程状态
	SetActiveDebuggingProcess(process UserActiveDebuggingProcess)

	// RestartProcess 重启之前通过.start命令启动的进程（不是重启调试器）
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: .restart命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/restart.cpp
	// 与内核模式比较: ❌ 仅用户模式调试器支持，内核模式无对应方法
	// 实现是否相同: ❌ 不适用（仅用户模式支持）
	// 实现过程: 用户命令.restart → UserDebugger.RestartProcess() → UdKillProcess() → UdAttachToProcess() → 本地处理
	// 注意：此功能仅用于用户模式调试器，重启之前通过.start命令启动的进程。如果进程正在运行，会先终止它，然后重新启动。
	//       内核模式调试器不需要此功能，因为内核调试器持续运行，不需要重启被调试的"目标"（整个操作系统）。
	RestartProcess()

	// ClearBreakpoint 通过ID清除断点
	// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825) / IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带断点ID的设备I/O
	// 用户命令: CommandBc -> 内核: BreakpointAddNew / DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bc.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过EPTHook实现
	// 实现是否相同: ❌ 不同（用户模式：IOCTL_SET_BREAKPOINT_USER_DEBUGGER + DeviceIoControl；内核模式：RemoveEvent移除EPTHook事件）
	// 实现过程: 用户模式：用户命令bc → UserDebugger.ClearBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointAddNew
	//           内核模式：用户命令bc → UserDebugger.ClearBreakpoint() → RemoveEvent() → 移除EPTHook事件
	ClearBreakpoint(breakpointID uint32)

	// DisableBreakpoint 通过ID禁用断点
	// IOCTL: 否, 串行通信 / IOCTL_DEBUGGER_MODIFY_EVENTS (0x80C)
	// 通信: 串行数据包 / 带事件标签的设备I/O
	// 用户命令: CommandBd -> 内核: 串行数据包处理 / DebuggerParseEventsModification
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bd.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用串行通信，内核模式通过ModifyEvent禁用EPTHook事件
	// 实现是否相同: ❌ 不同（用户模式：串行通信；内核模式：ModifyEvent禁用EPTHook事件）
	// 实现过程: 用户模式：用户命令bd → UserDebugger.DisableBreakpoint() → 串行通信 → 内核处理
	//           内核模式：用户命令bd → UserDebugger.DisableBreakpoint() → ModifyEvent() → 禁用EPTHook事件
	DisableBreakpoint(breakpointID uint32)

	// EnableBreakpoint 通过ID启用断点
	// IOCTL: 否, 串行通信 / IOCTL_DEBUGGER_MODIFY_EVENTS (0x80C)
	// 通信: 串行数据包 / 带事件标签的设备I/O
	// 用户命令: CommandBe -> 内核: 串行数据包处理 / DebuggerParseEventsModification
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/be.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用串行通信，内核模式通过ModifyEvent启用EPTHook事件
	// 实现是否相同: ❌ 不同（用户模式：串行通信；内核模式：ModifyEvent启用EPTHook事件）
	// 实现过程: 用户模式：用户命令be → UserDebugger.EnableBreakpoint() → 串行通信 → 内核处理
	//           内核模式：用户命令be → UserDebugger.EnableBreakpoint() → ModifyEvent() → 启用EPTHook事件
	EnableBreakpoint(breakpointID uint32)

	// Breakpoints 列出所有断点
	// IOCTL: 否, 串行通信 / IOCTL_LIST_EVENTS (0x800)
	// 通信: 串行数据包 / 带输出缓冲区的设备I/O
	// 用户命令: CommandBl -> 内核: 串行数据包处理 / LogRegisterIrpBasedNotification
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bl.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用串行通信，内核模式通过ListEvents列出EPTHook事件
	// 实现是否相同: ❌ 不同（用户模式：串行通信；内核模式：ListEvents列出EPTHook事件）
	// 实现过程: 用户模式：用户命令bl → UserDebugger.Breakpoints() → 串行通信 → 内核处理
	//           内核模式：用户命令bl → UserDebugger.Breakpoints() → ListEvents() → 列出EPTHook事件
	Breakpoints() []Breakpoint

	// SetBreakpoint 在指定地址设置断点
	// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825) / IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带地址的设备I/O
	// 用户命令: CommandBp -> 内核: BreakpointAddNew / DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bp.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过EPTHook实现
	// 实现是否相同: ❌ 不同（用户模式：IOCTL_SET_BREAKPOINT_USER_DEBUGGER + DeviceIoControl；内核模式：RegisterEvent注册EPTHook事件）
	// 实现过程: 用户模式：用户命令bp → UserDebugger.SetBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointAddNew
	//           内核模式：用户命令bp → UserDebugger.SetBreakpoint() → RegisterEvent() → 注册EPTHook事件
	SetBreakpoint(address uint64)

	// RemoveBreakpoint 通过ID移除断点
	// IOCTL: 是, IOCTL_SET_BREAKPOINT_USER_DEBUGGER (0x825) / IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
	// 通信: 带断点ID的设备I/O
	// 用户命令: CommandBc -> 内核: BreakpointAddNew / DebuggerParseEvent
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/bc.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用IOCTL_SET_BREAKPOINT_USER_DEBUGGER，内核模式通过EPTHook实现
	// 实现是否相同: ❌ 不同（用户模式：IOCTL_SET_BREAKPOINT_USER_DEBUGGER + DeviceIoControl；内核模式：RemoveEvent移除EPTHook事件）
	// 实现过程: 用户模式：用户命令bc → UserDebugger.RemoveBreakpoint() → IOCTL_SET_BREAKPOINT_USER_DEBUGGER → BreakpointAddNew
	//           内核模式：用户命令bc → UserDebugger.RemoveBreakpoint() → RemoveEvent() → 移除EPTHook事件
	RemoveBreakpoint(breakpointID uint32)

	// Continue 继续执行
	// IOCTL: 否, 串行通信 / IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED (0x812)
	// 通信: 串行数据包 / 设备I/O
	// 用户命令: CommandContinue -> 内核: 串行数据包处理 / DebuggerCommandSignalExecutionState
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/continue.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用串行通信，内核模式使用IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED
	// 实现是否相同: ❌ 不同（用户模式：串行通信；内核模式：IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED）
	// 实现过程: 用户模式：用户命令continue → UserDebugger.Continue() → 串行通信 → 内核处理
	//           内核模式：用户命令continue → KernelDebugger.Continue() → IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED → DebuggerCommandSignalExecutionState
	Continue()

	// StepOut 跳出当前函数（执行到ret指令）
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: gu命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/gu.cpp
	//         hyperdbg/libhyperdbg/code/debugger/core/steppings.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送单步执行请求
	// 实现是否相同: ✅ 相同（都使用串行通信 + SteppingStepOverForGu()）
	// 实现过程: 用户模式：用户命令gu → UserDebugger.StepOut() → SteppingStepOverForGu() → 串行通信 → 内核处理
	//           内核模式：用户命令gu → KernelDebugger.StepOut() → SteppingStepOverForGu() → 串行通信 → 内核处理
	StepOut()

	// Pause 暂停执行
	// IOCTL: 否, 串行通信 / IOCTL_PAUSE_PACKET_RECEIVED (0x811)
	// 通信: 串行数据包 / 设备I/O
	// 用户命令: CommandPause -> 内核: 串行数据包处理 / KdHaltSystem
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/pause.cpp
	// 与内核模式比较: ✅ 都支持，用户模式使用串行通信，内核模式使用IOCTL_PAUSE_PACKET_RECEIVED
	// 实现是否相同: ❌ 不同（用户模式：UdPauseProcess()；内核模式：IOCTL_PAUSE_PACKET_RECEIVED）
	// 实现过程: 用户模式：用户命令pause → UserDebugger.Pause() → UdPauseProcess() → 本地处理
	//           内核模式：用户命令pause → KernelDebugger.Pause() → IOCTL_PAUSE_PACKET_RECEIVED → KdHaltSystem
	Pause()

	// Registers 查询寄存器
	// IOCTL: 否, 本地命令 / IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER (0x814)
	// 通信: 本地处理 / 设备I/O
	// 用户命令: r命令 -> 内核: DebuggerCommandSendGeneralBufferToDebugger
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/r.cpp
	// 与内核模式比较: ✅ 都支持，用户模式本地处理，内核模式通过IOCTL发送寄存器数据
	// 执行过程: 用户命令r → UserDebugger.Registers() → 本地处理
	//           内核模式通过IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER获取寄存器
	Registers() []register.RegisterContext

	// SendUserInput 向被调试程序发送用户输入（'g'命令继续执行）
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: g命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/g.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送命令
	// 执行过程: 用户命令g → UserDebugger.SendUserInput() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收'g'命令继续执行
	SendUserInput(input string)

	// StepInto 单步进入下一条指令
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: t命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/t.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送单步命令
	// 执行过程: 用户命令t → UserDebugger.StepInto() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收单步命令
	StepInto()

	// StepOver 单步跳过下一条指令
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: p命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/p.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送单步命令
	// 执行过程: 用户命令p → UserDebugger.StepOver() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收单步命令
	StepOver(processId uint32)

	// WriteRegisters 写入寄存器值
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: r命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/r.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信写入寄存器
	// 执行过程: 用户命令r → UserDebugger.WriteRegisters() → HyperDbgReadTargetRegister() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收寄存器写入请求
	WriteRegisters(regs []register.RegisterContext)

	// ReadMemory 从指定地址读取内存
	// IOCTL: 是, IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY (0x801) (VMI模式) 或 串行通信 (Debugger模式)
	// 通信: 带内存读取结构的设备I/O 或 串行数据包
	// 用户命令: db/dc/dd/dq命令 -> 内核: DebuggerCommandReadMemory 或 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/d-u.cpp
	//         hyperdbg/libhyperdbg/code/debugger/misc/readmem.cpp
	// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL，Debugger模式通过串行通信
	// 执行过程: 用户命令db → UserDebugger.ReadMemory() → HyperDbgReadMemory() → IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY → DebuggerCommandReadMemory
	//           或 Debugger模式：用户命令db → UserDebugger.ReadMemory() → HyperDbgReadMemory() → 串行通信 → 内核处理
	ReadMemory(address uint64, size uint32) []byte

	// WriteMemory 向指定内存地址写入数据
	// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x802) (VMI模式) 或 串行通信 (Debugger模式)
	// 通信: 带内存写入结构的设备I/O 或 串行数据包
	// 用户命令: eb/ed/eq命令 -> 内核: DebuggerCommandEditMemory 或 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp
	//         hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/e.cpp (WriteMemoryContent函数)
	// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL，Debugger模式通过串行通信
	// 执行过程: 用户命令eb → UserDebugger.WriteMemory() → WriteMemoryContent() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
	//           或 Debugger模式：用户命令eb → UserDebugger.WriteMemory() → WriteMemoryContent() → 串行通信 → 内核处理
	WriteMemory(address uint64, data []byte)

	// Modules 查询加载的模块
	// IOCTL: 是, IOCTL_GET_USER_MODE_MODULE_DETAILS (0x819) (用户模式) 或 本地处理 (内核模式)
	// 通信: 带进程ID的设备I/O (用户模式) 或 本地处理 (内核模式)
	// 用户命令: lm命令 -> 内核: UserAccessGetLoadedModules (用户模式) 或 本地处理 (内核模式)
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/lm.cpp
	//         CommandLmShowUserModeModule (用户模式)
	//         CommandLmShowKernelModeModule (内核模式)
	// 与内核模式比较: ✅ 都支持，用户模式通过IOCTL，内核模式通过SymbolCheckAndAllocateModuleInformation
	// 执行过程: 用户命令lm → UserDebugger.Modules() → CommandLmShowUserModeModule() → IOCTL_GET_USER_MODE_MODULE_DETAILS → UserAccessGetLoadedModules
	//           或 内核模式：用户命令lm → UserDebugger.Modules() → CommandLmShowKernelModeModule() → SymbolCheckAndAllocateModuleInformation() → 本地处理
	// 注意：UserDebugger 继承了 Public.Modules(pid uint32)，此方法用于获取所有模块
	Modules() []ModuleInfo

	// CallStack 查询调用栈
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: k/kd/kq命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/k.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送调用栈查询请求
	// 执行过程: 用户命令k → UserDebugger.CallStack() → KdSendCallStackPacketToDebuggee() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收调用栈查询请求并返回结果
	CallStack() []StackFrame

	// ListenSerial 监听远程连接（作为服务器）
	// IOCTL: 否, 远程通信
	// 通信: 远程连接
	// 用户命令: CommandListen -> 远程监听
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/listen.cpp
	// 与内核模式比较: ✅ 都支持，用于远程调试监听
	// 执行过程: 用户命令listen → UserDebugger.ListenSerial() → RemoteConnectionListen() → 监听远程连接
	// 注意：此功能用于远程调试，作为服务器监听远程调试器的连接
	ListenSerial(port string, baudrate uint32)

	// LoadSymbols 加载符号文件
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: CommandLoad -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/load.cpp
	//         hyperdbg/libhyperdbg/code/debugger/script-engine/symbol.cpp
	// 与内核模式比较: ✅ 都支持，通过SymbolLoadOrDownloadSymbols函数加载符号
	// 执行过程: 用户命令load → UserDebugger.LoadSymbols() → SymbolLoadOrDownloadSymbols() → 本地处理
	//           内核模式同样支持符号加载，通过相同的符号系统
	LoadSymbols(path string)

	// UnloadSymbols 卸载符号文件
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: CommandUnload -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/unload.cpp
	//         hyperdbg/libhyperdbg/code/debugger/script-engine/symbol.cpp
	// 与内核模式比较: ✅ 都支持，通过符号系统卸载符号
	// 执行过程: 用户命令unload → UserDebugger.UnloadSymbols() → 本地处理
	//           内核模式同样支持符号卸载
	UnloadSymbols()

	// Symbol 查询符号信息
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: sym命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/sym.cpp
	//         hyperdbg/libhyperdbg/code/debugger/script-engine/symbol.cpp
	// 与内核模式比较: ✅ 都支持，通过符号系统查询符号
	// 执行过程: 用户命令sym → UserDebugger.Symbol() → SymbolBuildAndShowSymbolTable() → 本地处理
	//           内核模式同样支持符号查询
	Symbol(name string) SymbolInfo

	// SetSymbolPath 设置符号路径
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: sympath命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/sympath.cpp
	// 与内核模式比较: ✅ 都支持，通过符号系统设置符号路径
	// 执行过程: 用户命令sympath → UserDebugger.SetSymbolPath() → CommandSettingsSetValueToConfigFile() → 本地处理
	//           内核模式同样支持符号路径设置
	SetSymbolPath(path string)

	// Assemble 汇编指令
	// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x802) (VMI模式) 或 串行通信 (Debugger模式)
	// 通信: 带内存写入结构的设备I/O 或 串行数据包
	// 用户命令: a命令 -> 内核: DebuggerCommandEditMemory 或 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/a.cpp
	//         hyperdbg/libhyperdbg/code/debugger/misc/assembler.cpp
	// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL，Debugger模式通过串行通信
	// 执行过程: 用户命令a → UserDebugger.Assemble() → ScriptEngineAssembleWrapper() → WriteMemoryContent() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
	//           或 Debugger模式：用户命令a → UserDebugger.Assemble() → ScriptEngineAssembleWrapper() → WriteMemoryContent() → 串行通信 → 内核处理
	Assemble(address uint64, instruction string)

	// Eval 计算表达式
	// IOCTL: 否, 串行通信
	// 通信: 串行数据包
	// 用户命令: ?命令 -> 内核: 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/eval.cpp
	//         hyperdbg/libhyperdbg/code/debugger/script-engine/script-engine.cpp
	// 与内核模式比较: ✅ 都支持，通过串行通信发送表达式求值请求
	// 执行过程: 用户命令? → UserDebugger.Eval() → ScriptEngineExecuteSingleExpression() → 串行通信 → 内核处理
	//           内核模式通过串行协议接收表达式求值请求并返回结果
	Eval(expression string) uint64
	// Exit 退出调试器
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: exit命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/exit.cpp
	// 与内核模式比较: ✅ 都支持，VMI模式调用HyperDbgUnloadVmm()，Debugger模式调用KdCloseConnection()
	// 执行过程: 用户命令exit → UserDebugger.Exit() → HyperDbgUnloadVmm() (VMI模式) 或 KdCloseConnection() (Debugger模式)
	//           内核模式通过KdCloseConnection()关闭连接
	Exit()

	// GoTo 跳转到指定地址（实际是Good Game玩笑命令）
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: gg命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/gg.cpp
	// 与内核模式比较: ✅ 都支持，但此命令实际上只是一个玩笑命令，显示"Good Game!"
	// 执行过程: 用户命令gg → UserDebugger.GoTo() → ShowMessages("Good Game! :)")
	// 注意：此命令不是真正的跳转功能，只是一个玩笑命令
	GoTo(address uint64)

	// Sleep 睡眠指定毫秒数（调试器本身睡眠，用于脚本）
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: sleep命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/sleep.cpp
	// 与内核模式比较: ✅ 都支持，但此命令只是让调试器本身睡眠，不影响被调试目标
	// 执行过程: 用户命令sleep → UserDebugger.Sleep() → Sleep() → 本地处理
	// 注意：此命令主要用于脚本中，让调试器暂停指定时间，但不会中断调试或影响被调试目标
	Sleep(milliseconds uint32)

	// ClearScreen 清屏
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: cls命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/cls.cpp
	// 与内核模式比较: ✅ 都支持，只是调用system("cls")清屏
	// 执行过程: 用户命令cls → UserDebugger.ClearScreen() → system("cls") → 本地处理
	// 注意：此命令只是清空调试器的控制台显示，不影响调试状态
	ClearScreen()

	// DumpMemory 将内存转储到文件
	// IOCTL: 是, IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY (0x801) (VMI模式) 或 串行通信 (Debugger模式)
	// 通信: 带内存读取结构的设备I/O 或 串行数据包
	// 用户命令: dump命令 -> 内核: DebuggerCommandReadMemory 或 串行数据包处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/dump.cpp
	//         hyperdbg/libhyperdbg/code/debugger/misc/readmem.cpp
	// 与内核模式比较: ✅ 都支持，VMI模式通过IOCTL，Debugger模式通过串行通信
	// 执行过程: 用户命令dump → UserDebugger.DumpMemory() → HyperDbgShowMemoryOrDisassemble() → HyperDbgReadMemory() → IOCTL_DEBUGGER_READ_VIRTUAL_MEMORY → DebuggerCommandReadMemory
	//           或 Debugger模式：用户命令dump → UserDebugger.DumpMemory() → HyperDbgShowMemoryOrDisassemble() → HyperDbgReadMemory() → 串行通信 → 内核处理
	DumpMemory(address uint64, size uint32) []byte

	// ShowFormats 以不同格式显示值或寄存器
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: formats命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/formats.cpp
	//         hyperdbg/libhyperdbg/code/debugger/script-engine/script-engine.cpp
	// 与内核模式比较: ✅ 都支持，通过ScriptEngineEvalSingleExpression()求值并显示多种格式
	// 执行过程: 用户命令formats → UserDebugger.ShowFormats() → ScriptEngineEvalSingleExpression() → CommandFormatsShowResults() → 本地处理
	//           内核模式同样支持，显示十六进制、十进制、八进制、二进制、字符、时间、浮点数等多种格式
	ShowFormats()

	// CloseLogFile 关闭日志文件
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: logclose命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/logclose.cpp
	// 与内核模式比较: ✅ 都支持，只是关闭g_LogOpenFile文件流
	// 执行过程: 用户命令logclose → UserDebugger.CloseLogFile() → g_LogOpenFile.close() → 本地处理
	// 注意：此命令关闭之前通过.logopen打开的日志文件，停止记录命令和结果
	CloseLogFile()

	// OpenLogFile 打开日志文件
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: logopen命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/logopen.cpp
	// 与内核模式比较: ✅ 都支持，只是打开g_LogOpenFile文件流
	// 执行过程: 用户命令logopen → UserDebugger.OpenLogFile() → g_LogOpenFile.open() → 本地处理
	// 注意：此命令打开日志文件，开始记录所有命令和结果到文件
	OpenLogFile(path string)

	// PEInfo 查询PE文件信息
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: pe命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/pe.cpp
	//         hyperdbg/libhyperdbg/code/debugger/user-level/pe-parser.cpp
	// 与内核模式比较: ✅ 都支持，通过PeIsPE32BitOr64Bit()和PeParseHeader()解析PE文件
	// 执行过程: 用户命令pe → UserDebugger.PEInfo() → PeIsPE32BitOr64Bit() → PeParseHeader() → 本地处理
	//           内核模式同样支持PE文件解析，显示PE头信息和节区内容
	PEInfo(pid uint32)

	// ExecuteScript 执行HyperDbg脚本
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: script命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/script.cpp
	// 与内核模式比较: ✅ 都支持，通过HyperDbgScriptReadFileAndExecuteCommand()执行脚本
	// 执行过程: 用户命令script → UserDebugger.ExecuteScript() → HyperDbgScriptReadFileAndExecuteCommand() → HyperDbgInterpreter() → 本地处理
	//           内核模式同样支持脚本执行，可以批量执行调试命令
	ExecuteScript(script string) ScriptResult

	// Status 查询调试器状态
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: status命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/status.cpp
	// 与内核模式比较: ✅ 都支持，显示当前调试器的连接状态（本地/远程、VMI模式/Debugger模式）
	// 执行过程: 用户命令status → UserDebugger.Status() → CommandStatus() → 本地处理
	//           内核模式同样支持状态查询，显示调试器模式、连接信息等
	Status()

	// Connect 连接到远程调试器或本地调试器
	// IOCTL: 否, 远程通信
	// 通信: 远程连接
	// 用户命令: connect命令 -> 远程连接
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/connect.cpp
	// 与内核模式比较: ✅ 都支持，用于远程调试连接
	// 执行过程: 用户命令connect → UserDebugger.Connect() → RemoteConnectionConnect() → 建立远程连接
	//           或 ConnectLocalDebugger() → 本地VMI模式连接
	// 注意：此功能用于连接到远程调试器或本地调试器，支持TCP/IP连接
	Connect(port string)

	// Disconnect 断开与远程调试器的连接
	// IOCTL: 否, 远程通信
	// 通信: 远程连接
	// 用户命令: disconnect命令 -> 远程断开
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/disconnect.cpp
	// 与内核模式比较: ✅ 都支持，用于远程调试断开
	// 执行过程: 用户命令disconnect → UserDebugger.Disconnect() → RemoteConnectionCloseTheConnectionWithDebuggee() → 断开远程连接
	// 注意：此命令断开调试会话但不卸载模块，本地调试需要先unload驱动
	Disconnect()

	// CPU 查询CPU特性信息
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: cpu命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/cpu.cpp
	// 与内核模式比较: ✅ 都支持，通过ReadCpuDetails()读取CPU特性
	// 执行过程: 用户命令cpu → UserDebugger.CPU() → ReadCpuDetails() → InstructionSet::Vendor()/Brand()等 → 本地处理
	//           内核模式同样支持，显示CPU厂商、品牌、支持的指令集（SSE、AVX、AVX2等）等信息
	CPU()

	// Output 创建输出实例用于事件转发
	// IOCTL: 否, 本地命令
	// 通信: 本地处理
	// 用户命令: output命令 -> 本地处理
	// 源代码: hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/output.cpp
	//         hyperdbg/libhyperdbg/code/debugger/communication/forwarding.cpp
	// 与内核模式比较: ✅ 都支持，通过g_OutputSources链表管理输出源
	// 执行过程: 用户命令output → UserDebugger.Output() → CommandOutput() → 创建/打开/关闭输出实例 → 本地处理
	//           内核模式同样支持，可以创建文件、命名管道、TCP连接、模块等输出源，用于事件转发
	Output(value uint64)

	// Packet 返回底层的 Packet 实例
	// IOCTL: 否, 本地操作
	// 通信: 无
	// 源代码: debugger/user_debugger.go (Go实现特有)
	// 与内核模式比较: ✅ 一致，两种模式都支持
	// 执行过程: 调用Packet() → 返回底层的 *Packet 实例
	// 注意: 此方法是Go实现特有的，用于访问底层的 Packet 实例
	Packet() *Packet

	// QueryProcessStatus 查询进程状态
	// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80E)
	// 通信: 设备I/O
	// 执行过程: 调用QueryProcessStatus() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → 返回暂停状态和RIP
	QueryProcessStatus() (isPaused bool, rip uint64)

	// WaitForBreakpoint 等待断点命中
	// IOCTL: 否, 本地操作（轮询QueryProcessStatus）
	// 通信: 无
	// 执行过程: 调用WaitForBreakpoint() → 轮询QueryProcessStatus() → 等待进程暂停
	WaitForBreakpoint(timeoutMs int) bool

	Eventer
}
