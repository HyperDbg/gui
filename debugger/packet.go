package debugger

import (
	"bytes"
	"encoding/binary"
	"sync"

	"github.com/ddkwork/HyperDbg/debugger/api"
	"github.com/ddkwork/HyperDbg/debugger/ark"
	"github.com/ddkwork/HyperDbg/debugger/bookmark"
	"github.com/ddkwork/HyperDbg/debugger/breakpoint"
	"github.com/ddkwork/HyperDbg/debugger/comment"
	"github.com/ddkwork/HyperDbg/debugger/disassembly"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/HyperDbg/debugger/exception"
	"github.com/ddkwork/HyperDbg/debugger/function"
	"github.com/ddkwork/HyperDbg/debugger/graph"
	"github.com/ddkwork/HyperDbg/debugger/handle"
	"github.com/ddkwork/HyperDbg/debugger/label"
	"github.com/ddkwork/HyperDbg/debugger/loop"
	"github.com/ddkwork/HyperDbg/debugger/memory"
	"github.com/ddkwork/HyperDbg/debugger/notes"
	"github.com/ddkwork/HyperDbg/debugger/peview"
	"github.com/ddkwork/HyperDbg/debugger/reference"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/ddkwork/HyperDbg/debugger/scylla"
	"github.com/ddkwork/HyperDbg/debugger/seh"
	"github.com/ddkwork/HyperDbg/debugger/source"
	"github.com/ddkwork/HyperDbg/debugger/stack"
	"github.com/ddkwork/HyperDbg/debugger/symbol"
	"github.com/ddkwork/HyperDbg/debugger/thread"
	"github.com/ddkwork/HyperDbg/debugger/trace"
	"github.com/ddkwork/HyperDbg/debugger/transparency"
	"github.com/ddkwork/HyperDbg/debugger/watch"
	"github.com/ddkwork/HyperDbg/debugger/xref"
	"github.com/ddkwork/HyperDbg/walker/hardware"
	type_ "github.com/ddkwork/HyperDbg/walker/types"
	"github.com/ddkwork/golibrary/std/mylog"
)

var autoLoadVMM bool

type Packet struct {
	mu             sync.RWMutex
	driver         driver.Api
	DriverProvider driver.Api
	eventLoop      *EventLoop
	isVmmLoaded    bool
	AutoLoadDriver bool

	State                             DebuggerState
	ActiveProcess                     ActiveDebuggingProcess
	BreakPrintingOutput               bool
	AutoUnpause                       bool
	IsConnectedToRemoteDebuggee       bool
	IsSerialConnectedToRemoteDebuggee bool
	IsInstrumentingInstructions       bool
	IgnorePauseRequests               bool
	EventChan                         chan *DebugEvent
	EventCallbacks                    []DebugEventCallback
	OnTitleUpdate                     func(title string)

	EventManager      *EventManager
	DebugEventHandler *DebugEventHandler
	UIComponents      []api.Interface

	RegistersMeta    api.Interface
	MemoryMeta       api.Interface
	BreakpointsMeta  api.Interface
	SymbolsMeta      api.Interface
	ThreadsMeta      api.Interface
	StackMeta        api.Interface
	DisassemblyMeta  api.Interface
	TraceMeta        api.Interface
	TransparencyMeta api.Interface
	HardwareMeta     api.Interface
	VmmMeta          api.Interface

	BookmarksMeta  api.Interface
	CommentsMeta   api.Interface
	LabelsMeta     api.Interface
	FunctionsMeta  api.Interface
	XrefsMeta      api.Interface
	TypesMeta      api.Interface
	WatchesMeta    api.Interface
	GraphsMeta     api.Interface
	ExceptionsMeta api.Interface
	NotesMeta      api.Interface
	SourceMeta     api.Interface
	ReferencesMeta api.Interface
	HandlesMeta    api.Interface
	ScriptMeta     api.Interface
	ArkMeta        api.Interface
	ScyllaMeta     api.Interface
	SehMeta        api.Interface
	LoopsMeta      api.Interface
	PeviewMeta     api.Interface
}

var (
	globalPacket *Packet
	packetOnce   sync.Once
)

func GetPacket() *Packet {
	packetOnce.Do(func() {
		mylog.Warning("⚠️  请确保在调用处添加: defer dbg.UnloadDriver()")
		mylog.Warning("⚠️  否则驱动不会正确卸载，可能导致系统不稳定！")

		globalPacket = &Packet{
			EventChan: make(chan *DebugEvent, 100),
		}

		globalPacket.RegistersMeta = register.New()
		globalPacket.MemoryMeta = memory.New()
		globalPacket.BreakpointsMeta = breakpoint.New()
		globalPacket.SymbolsMeta = symbol.New()
		globalPacket.ThreadsMeta = thread.New()
		globalPacket.StackMeta = stack.New()
		globalPacket.DisassemblyMeta = disassembly.New()
		globalPacket.TraceMeta = trace.New()
		globalPacket.TransparencyMeta = transparency.New()

		globalPacket.RegisterUIComponent(globalPacket.RegistersMeta)
		globalPacket.RegisterUIComponent(globalPacket.MemoryMeta)
		globalPacket.RegisterUIComponent(globalPacket.BreakpointsMeta)
		globalPacket.RegisterUIComponent(globalPacket.SymbolsMeta)
		globalPacket.RegisterUIComponent(globalPacket.ThreadsMeta)
		globalPacket.RegisterUIComponent(globalPacket.StackMeta)
		globalPacket.RegisterUIComponent(globalPacket.DisassemblyMeta)
		globalPacket.RegisterUIComponent(globalPacket.TraceMeta)
		globalPacket.RegisterUIComponent(globalPacket.TransparencyMeta)

		globalPacket.BookmarksMeta = bookmark.New()
		globalPacket.CommentsMeta = comment.New()
		globalPacket.LabelsMeta = label.New()
		globalPacket.FunctionsMeta = function.New()
		globalPacket.XrefsMeta = xref.New()
		globalPacket.WatchesMeta = watch.New()
		globalPacket.NotesMeta = notes.New()
		globalPacket.SourceMeta = source.New()
		globalPacket.ReferencesMeta = reference.New()
		globalPacket.HandlesMeta = handle.New()
		globalPacket.PeviewMeta = peview.New()
		globalPacket.ExceptionsMeta = exception.New()
		globalPacket.ArkMeta = ark.New()
		globalPacket.ScyllaMeta = scylla.New()
		globalPacket.GraphsMeta = graph.New()
		globalPacket.TypesMeta = type_.New()
		globalPacket.LoopsMeta = loop.New()
		globalPacket.SehMeta = seh.New()

		globalPacket.EventManager = NewEventManager(globalPacket.driver)
		globalPacket.eventLoop = NewEventLoop(globalPacket.EventManager, globalPacket.driver)
	})
	return globalPacket
}

func NewPacket(dh *driver.Provider) Packeter {
	p := &Packet{
		driver: dh,
	}

	p.RegistersMeta = register.New()
	p.MemoryMeta = memory.New()
	p.BreakpointsMeta = breakpoint.New()
	p.SymbolsMeta = symbol.New()
	p.ThreadsMeta = thread.New()
	p.StackMeta = stack.New()
	p.DisassemblyMeta = disassembly.New()
	p.TraceMeta = trace.New()
	p.TransparencyMeta = transparency.New()

	p.RegisterUIComponent(p.RegistersMeta)
	p.RegisterUIComponent(p.MemoryMeta)
	p.RegisterUIComponent(p.BreakpointsMeta)
	p.RegisterUIComponent(p.SymbolsMeta)
	p.RegisterUIComponent(p.ThreadsMeta)
	p.RegisterUIComponent(p.StackMeta)
	p.RegisterUIComponent(p.DisassemblyMeta)
	p.RegisterUIComponent(p.TraceMeta)
	p.RegisterUIComponent(p.TransparencyMeta)

	p.BookmarksMeta = bookmark.New()
	p.CommentsMeta = comment.New()
	p.LabelsMeta = label.New()
	p.FunctionsMeta = function.New()
	p.XrefsMeta = xref.New()
	p.WatchesMeta = watch.New()
	p.NotesMeta = notes.New()
	p.SourceMeta = source.New()
	p.ReferencesMeta = reference.New()
	p.HandlesMeta = handle.New()
	p.PeviewMeta = peview.New()
	p.ExceptionsMeta = exception.New()
	p.ArkMeta = ark.New()
	p.ScyllaMeta = scylla.New()
	p.GraphsMeta = graph.New()
	p.TypesMeta = type_.New()
	p.LoopsMeta = loop.New()
	p.SehMeta = seh.New()

	p.EventManager = NewEventManager(p.driver)
	p.eventLoop = NewEventLoop(p.EventManager, p.driver)

	return p
}

// Packet 返回底层的 Packet 实例
func (p *Packet) Packet() *Packet {
	return p
}

func (p *Packet) RegisterUIComponent(component api.Interface) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.UIComponents = append(p.UIComponents, component)
}

// 来自接口: Packeter 第1个签名
func (p *Packet) LoadDriver(driverPath string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	mylog.Info("=== 正在加载驱动 ===")

	mylog.Info("检查 CPU 厂商...")
	hardware.CheckCpuVendor()

	mylog.Info("检查 VMX 支持...")
	if !hardware.VmxSupportDetection() {
		mylog.Check("VMX 不支持")
	}

	if p.DriverProvider == nil {
		p.DriverProvider = driver.New(driverPath, "KernelDebugger", "\\\\.\\HyperDbgDebuggerDevice")
	}

	mylog.Info("安装驱动...")
	p.DriverProvider.Install()

	mylog.Info("启动驱动（VMM 将自动初始化）...")
	p.DriverProvider.Start()

	p.driver = p.DriverProvider
	p.isVmmLoaded = true

	mylog.Info("✓ 驱动加载成功")
	mylog.Info("=== 驱动加载完成 ===")
}

// 来自接口: Packeter 第2个签名
// UnloadDriver 卸载驱动
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802) 和 IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL (0x801)
// 通信: 设备I/O
// 源代码: hyperkd/code/driver/Ioctl.c
//
// ┌──────────────────────────────────────────────────────────────────────────────┐
// │  ⚠️  【警告】禁止更改 IOCTL 发送顺序！                                         │
// │                                                                              │
// │  此函数的 IOCTL 发送顺序经过严格验证，错误的顺序会导致 BSOD。                    │
// │  任何修改都必须经过充分测试，否则可能导致系统崩溃。                              │
// └──────────────────────────────────────────────────────────────────────────────┘
//
// 【IOCTL 发送顺序说明】
//
// 正确顺序（当前实现）:
//
//  1. IOCTL_TERMINATE_VMX - 终止 VMX，使 CPU 退出虚拟化模式
//  2. IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL - 禁止后续 IOCTL
//
// 错误顺序（曾导致 BSOD）:
//
//  1. IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL - 设置 g_AllowIOCTLFromUsermode = FALSE
//  2. IOCTL_TERMINATE_VMX - 被 Ioctl.c 第55行的 if (g_AllowIOCTLFromUsermode) 跳过！
//
// 【BSOD 原因分析】
//
//   - Ioctl.c 中 IOCTL 处理逻辑: if (g_AllowIOCTLFromUsermode) { switch(...) { ... } }
//   - 如果先发送 IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL，会设置 g_AllowIOCTLFromUsermode = FALSE
//   - 后续的 IOCTL_TERMINATE_VMX 会被 if 条件跳过，导致 VMX 未被终止
//   - 驱动卸载后，CPU 仍处于 VMX root 模式，VM-exit handler 地址已无效
//   - 当发生 VM-exit 时，CPU 跳转到已卸载的 AsmVmexitHandler (偏移 0x3b6a0)，触发 BSOD
//
// 【为什么普通调试手段无法解决】
//
//   - 驱动已卸载，代码地址无效，无法设置断点
//   - 崩溃发生在内核态，用户态调试器无法捕获
//   - 需要通过 minidump + PDB 符号文件进行事后分析
//   - 必须使用 WinDbg kd.exe 分析 dump 文件定位崩溃函数
//
// 【分析工具】
//
//	debugger/driver/analyze_bsod.ps1 - BSOD dump 分析脚本
//	用法: .\analyze_bsod.ps1 -PdbPath "build\Release" -DumpFile "C:\Windows\Minidump\xxx.dmp"
//	该脚本通过 PDB 符号文件定位崩溃函数，是解决此类问题的关键工具
//
// 【崩溃位置】
//
//	hyperkd!AsmVmexitHandler (AsmVmexitHandler.asm, 偏移 0x3b6a0)
func (p *Packet) UnloadDriver() {
	p.mu.Lock()
	defer p.mu.Unlock()

	mylog.Info("=== 正在卸载驱动 ===")

	//
	// 【关键】必须先发送 IOCTL_TERMINATE_VMX 终止 VMX
	// 然后再发送 IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL 禁止后续 IOCTL
	// 顺序错误会导致 VMX 未终止，驱动卸载后触发 BSOD
	//
	if p.driver != nil && p.driver.IsConnected() {
		mylog.Info("发送终止 VMX IOCTL...")
		emptyBuffer := bytes.NewBuffer(make([]byte, 8))
		p.driver.SendReceive(emptyBuffer, IoctlTerminateVmx)

		mylog.Info("发送 IRP Pending 完成 IOCTL...")
		emptyBuffer = bytes.NewBuffer(make([]byte, 8))
		p.driver.SendReceive(emptyBuffer, IoctlReturnIrpPendingPacketsAndDisallowIoctl)
	}

	if p.eventLoop != nil {
		mylog.Info("停止事件循环...")
		p.eventLoop.Stop()
	}

	if p.DriverProvider != nil {
		mylog.Info("停止驱动服务...")
		p.DriverProvider.Stop()

		mylog.Info("卸载驱动...")
		p.DriverProvider.Uninstall()
	}

	if p.driver != nil {
		mylog.Info("关闭驱动句柄...")
		p.driver.Close()
		p.driver = nil
	}

	p.isVmmLoaded = false
	mylog.Info("✓ 驱动卸载成功")
	mylog.Info("=== 驱动卸载完成 ===")
}

// 来自接口: Packeter 第3个签名
// LoadVMM 加载虚拟机监控器
// IOCTL: 否, VMM在驱动打开时自动加载
// 通信: 无
// 源代码: hyperkd/code/driver/Driver.c:DrvCreate -> LoaderInitVmmAndDebugger
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: VMM在DrvCreate时自动初始化，此方法仅检查状态
func (p *Packet) LoadVMM() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.driver == nil {
		mylog.Check("驱动未加载")
	}

	if p.isVmmLoaded {
		return
	}

	mylog.Check("VMM 未初始化，请先调用 LoadDriver()")
}

// 来自接口: Packeter 第4个签名
// UnloadVMM 卸载虚拟机监控器
// IOCTL: 是, IOCTL_TERMINATE_VMX (0x802) 和 IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL (0x801)
// 通信: 设备I/O
// 源代码: hyperkd/code/driver/Ioctl.c:IOCTL_TERMINATE_VMX -> DebuggerUninitialize + VmFuncUninitVmm
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: VMM卸载由UnloadDriver()统一处理，此方法仅检查状态
func (p *Packet) UnloadVMM() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !p.isVmmLoaded {
		return
	}

	mylog.Check("VMM 仍在运行中，请先调用 UnloadDriver()")
}

// 来自接口: Packeter 第5个签名
// IsConnected 检查调试器是否当前已连接
// IOCTL: 否, 本地状态检查
// 通信: 无
// 源代码: hyperdbg/debugger.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IsConnected() → 本地状态检查
func (p *Packet) IsConnected() bool {
	return p.driver != nil && p.driver.IsConnected()
}

// 来自接口: Packeter 第6个签名
// EPTHook 启用EPT（扩展页表）钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) EPTHook(address uint64, size uint32, hookType EPTHookType) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerEpthookRequest{
		Address:  address,
		Size:     size,
		HookType: hookType,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第7个签名
// EPTHook2 启用EPT钩子（替代实现）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandEpthook2 -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/epthook2.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EPTHook2() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) EPTHook2(address uint64, size uint32, hookType EPTHookType) {
	p.EPTHook(address, size, hookType)
}

// 来自接口: Packeter 第8个签名
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
func (p *Packet) HookSyscall(syscallNumber uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerSyscallHookRequest{
		SyscallNumber: syscallNumber,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第9个签名
// HookException 启用异常钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandException -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/exception.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookException() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) HookException(exceptionType uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerExceptionHookRequest{
		ExceptionType: exceptionType,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第10个签名
// HookInterrupt 启用中断钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandInterrupt -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/interrupt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookInterrupt() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) HookInterrupt(vector uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerInterruptHookRequest{
		Vector: vector,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第11个签名
// HookIO 启用I/O端口钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoin/CommandIoout -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioin.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIO() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) HookIO(port uint16, hookType uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerIoHookRequest{
		Port:     port,
		HookType: hookType,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第12个签名
// HookIOAPIC 启用I/O APIC钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandIoapic -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/ioapic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookIOAPIC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) HookIOAPIC(apicID uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerApicRequest{
		ApicID: apicID,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第13个签名
func (p *Packet) ReadMsr(msr uint32) uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return 0
	}

	req := DebuggerReadOrWriteMsr{
		Msr:        msr,
		ActionType: DebuggerMsrActionRead,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerReadOrWriteMsr)

	var result DebuggerReadOrWriteMsr
	binary.Read(response, binary.LittleEndian, &result)
	return result.Value
}

// 来自接口: Packeter 第14个签名
func (p *Packet) WriteMsr(msr uint32, value uint64) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerReadOrWriteMsr{
		Msr:        msr,
		ActionType: DebuggerMsrActionWrite,
		Value:      value,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerReadOrWriteMsr)
}

// 来自接口: Packeter 第15个签名
// MeasurePerformance 测量特定地址的性能
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandMeasure -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/measure.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用MeasurePerformance() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) MeasurePerformance(address uint64) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerMeasureRequest{
		Address: address,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第16个签名
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
func (p *Packet) MonitorMemory(address uint64, size uint32, monitorType MonitorType) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerMonitorRequest{
		Address:     address,
		Size:        size,
		MonitorType: monitorType,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第17个签名
// PCICam 查询PCI CAM（配置访问机制）
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPcicam -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pcicam.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCICam() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) PCICam(bus, device, function uint32) PCICamInfo {
	var result PCICamInfo
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return result
	}

	req := DebuggerPciCamRequest{
		Bus:      bus,
		Device:   device,
		Function: function,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)

	result.Bus = bus
	result.Device = device
	result.Function = function
	result.VendorID = binary.LittleEndian.Uint16(response.Bytes()[0:2])
	result.DeviceID = binary.LittleEndian.Uint16(response.Bytes()[2:4])
	result.Data = binary.LittleEndian.Uint32(response.Bytes()[4:8])

	return result
}

// 来自接口: Packeter 第18个签名
// PMC 查询性能监控计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandPmc -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pmc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PMC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) PMC(pmcNumber uint32) uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return 0
	}

	req := DebuggerPmcRequest{
		PmcNumber: pmcNumber,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)

	return binary.LittleEndian.Uint64(response.Bytes()[0:8])
}

// 来自接口: Packeter 第19个签名
// ReconstructMemory 从跟踪中重建内存
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandRev -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/rev.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ReconstructMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) ReconstructMemory(pid uint32, address uint64, size uint32, mode ReconstructMode) []byte {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggerReconstructMemoryRequest{
		ProcessId: pid,
		Size:      size,
		Mode:      mode,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, address)
	response := p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)

	return response.Bytes()
}

// 来自接口: Packeter 第20个签名
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
func (p *Packet) SearchMemoryPattern(pid uint32, pattern []byte, mode ReconstructMode) []uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggerSearchMemoryPatternRequest{
		ProcessId:   pid,
		PatternSize: uint32(len(pattern)),
		Mode:        mode,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, pattern)
	response := p.driver.SendReceive(buf, IoctlDebuggerSearchMemory)

	resultCount := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	results := make([]uint64, resultCount)
	for i := range resultCount {
		results[i] = binary.LittleEndian.Uint64(response.Bytes()[4+i*8 : 12+i*8])
	}

	return results
}

// 来自接口: Packeter 第21个签名
// InstructionTrace 启用指令执行跟踪
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrace -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/trace.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用InstructionTrace() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
// 参数说明: address - 要跟踪的起始地址
// 返回值: 跟踪结果通过事件系统异步返回，而不是直接返回
func (p *Packet) InstructionTrace(address uint64) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerInstructionTraceRequest{
		Address: address,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第22个签名
// TrackMemory 跟踪内存访问
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTrack -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/track.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TrackMemory() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) TrackMemory(address uint64, size uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerTrackMemoryRequest{
		Address: address,
		Size:    size,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第23个签名
// TSC 查询时间戳计数器
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandTsc -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/tsc.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用TSC() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) TSC() uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return 0
	}

	response := p.driver.SendReceive(nil, IoctlDebuggerRegisterEvent)

	return binary.LittleEndian.Uint64(response.Bytes()[0:8])
}

// 来自接口: Packeter 第24个签名
// VMCallHandler 启用VMCALL处理程序
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandVmcall -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/vmcall.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VMCallHandler() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) VMCallHandler() {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	p.driver.SendReceive(nil, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第25个签名
// HookXSETBV 启用XSETBV钩子
// IOCTL: 是, IOCTL_DEBUGGER_REGISTER_EVENT (0x806)
// 通信: 带事件结构的设备I/O
// 用户命令: CommandXsetbv -> 内核: DebuggerParseEvent
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/xsetbv.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HookXSETBV() → IOCTL_DEBUGGER_REGISTER_EVENT → DebuggerParseEvent
func (p *Packet) HookXSETBV() {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	p.driver.SendReceive(nil, IoctlDebuggerRegisterEvent)
}

// 来自接口: Packeter 第26个签名
// PTE 查询页表项
// IOCTL: 是, IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS (0x805)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandPte -> 内核: ExtensionCommandPte
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pte.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PTE() → IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS → ExtensionCommandPte
func (p *Packet) PTE(virtualAddress uint64, pid uint32) PageTableEntries {
	var result PageTableEntries
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return result
	}

	req := DebuggerPteRequest{
		VirtualAddress: virtualAddress,
		ProcessId:      pid,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerReadPageTableEntriesDetails)

	result.VirtualAddress = virtualAddress
	result.ProcessId = pid
	result.Pml4eVirtualAddress = binary.LittleEndian.Uint64(response.Bytes()[0:8])
	result.Pml4eValue = binary.LittleEndian.Uint64(response.Bytes()[8:16])
	result.PdpteVirtualAddress = binary.LittleEndian.Uint64(response.Bytes()[16:24])
	result.PdpteValue = binary.LittleEndian.Uint64(response.Bytes()[24:32])
	result.PdeVirtualAddress = binary.LittleEndian.Uint64(response.Bytes()[32:40])
	result.PdeValue = binary.LittleEndian.Uint64(response.Bytes()[40:48])
	result.PteVirtualAddress = binary.LittleEndian.Uint64(response.Bytes()[48:56])
	result.PteValue = binary.LittleEndian.Uint64(response.Bytes()[56:64])

	return result
}

// 来自接口: Packeter 第27个签名
// VA2PA 将虚拟地址转换为物理地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带虚拟地址的设备I/O
// 用户命令: CommandVa2pa -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/va2pa.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用VA2PA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (p *Packet) VA2PA(virtualAddress uint64, pid uint32) uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return 0
	}

	req := DebuggerVa2paAndPa2vaCommands{
		VirtualAddress:     virtualAddress,
		ProcessId:          pid,
		IsVirtual2Physical: true,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerVa2paAndPa2vaCommands)

	var result DebuggerVa2paAndPa2vaCommands
	binary.Read(response, binary.LittleEndian, &result)
	return result.PhysicalAddress
}

// 来自接口: Packeter 第28个签名
// PA2VA 将物理地址转换为虚拟地址
// IOCTL: 是, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS (0x809)
// 通信: 带物理地址的设备I/O
// 用户命令: CommandPa2va -> 内核: ExtensionCommandVa2paAndPa2va
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pa2va.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PA2VA() → IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS → ExtensionCommandVa2paAndPa2va
func (p *Packet) PA2VA(physicalAddress uint64, pid uint32) uint64 {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return 0
	}

	req := DebuggerVa2paAndPa2vaCommands{
		PhysicalAddress:    physicalAddress,
		ProcessId:          pid,
		IsVirtual2Physical: false,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerVa2paAndPa2vaCommands)

	var result DebuggerVa2paAndPa2vaCommands
	binary.Read(response, binary.LittleEndian, &result)
	return result.VirtualAddress
}

// 来自接口: Packeter 第29个签名
// ProcessDetails 查询进程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_PROCESS (0x81C)
// 通信: 带进程ID的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryDetails
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ProcessDetails() → IOCTL_QUERY_CURRENT_PROCESS → ProcessQueryDetails
func (p *Packet) ProcessDetails(pid uint32) []ProcessDetails {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggeeDetailsAndSwitchProcessPacket{
		ActionType: DebuggeeDetailsAndSwitchProcessGetProcessDetails,
		ProcessId:  pid,
	}

	// 注意: 不能使用 binary.Write 序列化此结构体
	// binary.Write 存在 bug，无法正确序列化嵌套结构体 DebuggeeProcessListNeededDetails
	// 导致只写入 57 字节而不是预期的 72 字节
	// 内核驱动检查 InputBufferLength >= SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET (72字节)
	// 如果缓冲区太小会返回 STATUS_INVALID_PARAMETER 错误
	// 因此必须手动序列化确保发送完整的 72 字节缓冲区
	buf := make([]byte, 72)
	binary.LittleEndian.PutUint32(buf[0:], req.ActionType)
	binary.LittleEndian.PutUint32(buf[4:], req.ProcessId)
	binary.LittleEndian.PutUint64(buf[8:], req.Process)
	if req.IsSwitchByClkIntr {
		buf[16] = 1
	}
	copy(buf[17:], req.ProcessName[:])
	binary.LittleEndian.PutUint64(buf[40:], req.ProcessListSymDetails.PsActiveProcessHead)
	binary.LittleEndian.PutUint32(buf[48:], req.ProcessListSymDetails.ImageFileNameOffset)
	binary.LittleEndian.PutUint32(buf[52:], req.ProcessListSymDetails.UniquePidOffset)
	binary.LittleEndian.PutUint32(buf[56:], req.ProcessListSymDetails.ActiveProcessLinksOffset)
	binary.LittleEndian.PutUint32(buf[64:], req.Result)

	response := p.driver.SendReceive(bytes.NewBuffer(buf), IoctlQueryCurrentProcess)

	if response.Len() == 0 {
		mylog.Warning("内核返回空响应")
		return nil
	}

	if response.Len() < 72 {
		mylog.Warning("ProcessDetails 响应长度不足", response.Len())
		return nil
	}

	var resp DebuggeeDetailsAndSwitchProcessPacket
	err := binary.Read(bytes.NewReader(response.Bytes()), binary.LittleEndian, &resp)
	if err != nil {
		mylog.Warning("解析 ProcessDetails 响应失败", err)
		return nil
	}

	details := []ProcessDetails{
		{
			ProcessId:    resp.ProcessId,
			BaseAddress:  0,
			EntryPoint:   0,
			Is32Bit:      false,
			IsDebugged:   true,
			KernelStatus: resp.Result,
		},
	}

	mylog.Info("ProcessDetails ProcessId", resp.ProcessId, "Result", resp.Result)

	return details
}

// 来自接口: Packeter 第30个签名
// ThreadDetails 查询线程的详细信息
// IOCTL: 是, IOCTL_QUERY_CURRENT_THREAD (0x81D)
// 通信: 带线程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryDetails
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ThreadDetails() → IOCTL_QUERY_CURRENT_THREAD → ThreadQueryDetails
func (p *Packet) ThreadDetails(tid uint32) []ThreadDetails {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggerQueryThreadRequest{
		ThreadId: tid,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlQueryCurrentThread)

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	details := make([]ThreadDetails, count)
	for i := range count {
		offset := 4 + i*64
		details[i] = ThreadDetails{
			ThreadId:     binary.LittleEndian.Uint32(response.Bytes()[offset : offset+4]),
			ProcessId:    binary.LittleEndian.Uint32(response.Bytes()[offset+4 : offset+8]),
			StartAddress: binary.LittleEndian.Uint64(response.Bytes()[offset+8 : offset+16]),
			IsRunning:    binary.LittleEndian.Uint32(response.Bytes()[offset+16:offset+20]) == 1,
			IsSuspended:  binary.LittleEndian.Uint32(response.Bytes()[offset+20:offset+24]) == 1,
			KernelStatus: binary.LittleEndian.Uint32(response.Bytes()[offset+24 : offset+28]),
		}
	}

	return details
}

// 来自接口: Packeter 第31个签名
// Processes 列出所有进程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: process命令 -> 内核: ProcessQueryList
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Processes() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ProcessQueryList
func (p *Packet) Processes() []ProcessInfo {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	response := p.driver.SendReceive(nil, IoctlGetListOfThreadsAndProcesses)

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	processes := make([]ProcessInfo, count)
	for i := range count {
		offset := 4 + i*128
		processes[i] = ProcessInfo{
			ProcessId:   binary.LittleEndian.Uint32(response.Bytes()[offset : offset+4]),
			BaseAddress: binary.LittleEndian.Uint64(response.Bytes()[offset+8 : offset+16]),
		}
	}

	return processes
}

// 来自接口: Packeter 第32个签名
// Threads 列出进程的所有线程
// IOCTL: 是, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES (0x81B)
// 通信: 带进程ID的设备I/O
// 用户命令: thread命令 -> 内核: ThreadQueryList
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用Threads() → IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES → ThreadQueryList
func (p *Packet) Threads(pid uint32) []ThreadInfo {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggerQueryThreadsRequest{
		ProcessId: pid,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlGetListOfThreadsAndProcesses)

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	threads := make([]ThreadInfo, count)
	for i := range count {
		offset := 4 + i*64
		threads[i] = ThreadInfo{
			ThreadId:     binary.LittleEndian.Uint32(response.Bytes()[offset : offset+4]),
			ProcessId:    binary.LittleEndian.Uint32(response.Bytes()[offset+4 : offset+8]),
			StartAddress: binary.LittleEndian.Uint64(response.Bytes()[offset+8 : offset+16]),
		}
	}

	return threads
}

// 来自接口: Packeter 第33个签名
// IDTEntries 查询中断描述符表项
// IOCTL: 是, IOCTL_QUERY_IDT_ENTRY (0x824)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandIdt -> 内核: ExtensionCommandPerformQueryIdtEntriesRequest
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/idt.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用IDTEntries() → IOCTL_QUERY_IDT_ENTRY → ExtensionCommandPerformQueryIdtEntriesRequest
func (p *Packet) IDTEntries() []IDTEntry {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	response := p.driver.SendReceive(nil, IoctlQueryIdtEntry)

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	entries := make([]IDTEntry, count)
	for i := range count {
		offset := 4 + i*16
		entries[i] = IDTEntry{
			Offset:   binary.LittleEndian.Uint64(response.Bytes()[offset : offset+8]),
			Selector: binary.LittleEndian.Uint16(response.Bytes()[offset+8 : offset+10]),
			Type:     response.Bytes()[offset+10],
			IST:      response.Bytes()[offset+11],
		}
	}

	return entries
}

// 来自接口: Packeter 第34个签名
// APIC 查询高级可编程中断控制器
// IOCTL: 是, IOCTL_PERFORM_ACTIONS_ON_APIC (0x822)
// 通信: 带APIC ID的设备I/O
// 用户命令: CommandApic -> 内核: ExtensionCommandPerformActionsOnApic
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/apic.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用APIC() → IOCTL_PERFORM_ACTIONS_ON_APIC → ExtensionCommandPerformActionsOnApic
func (p *Packet) APIC(apicID uint32) APICInfo {
	var result APICInfo
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return result
	}

	req := DebuggerApicRequest{
		ApicID: apicID,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlPerformActionsOnApic)

	result.APICID = binary.LittleEndian.Uint32(response.Bytes()[0:4])
	result.Version = binary.LittleEndian.Uint32(response.Bytes()[4:8])
	result.EOI = binary.LittleEndian.Uint32(response.Bytes()[8:12])
	result.TPR = binary.LittleEndian.Uint32(response.Bytes()[12:16])

	return result
}

// 来自接口: Packeter 第35个签名
// PCITree 查询PCI设备树
// IOCTL: 是, IOCTL_PCIE_ENDPOINT_ENUM (0x823)
// 通信: 带输出缓冲区的设备I/O
// 用户命令: CommandPci -> 内核: ExtensionCommandPcieEndpointEnum
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/pci.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PCITree() → IOCTL_PCIE_ENDPOINT_ENUM → ExtensionCommandPcieEndpointEnum
func (p *Packet) PCITree() []PCIDevice {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	response := p.driver.SendReceive(nil, IoctlPcieEndpointEnum)

	count := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	devices := make([]PCIDevice, count)
	for i := range count {
		offset := 4 + i*32
		devices[i] = PCIDevice{
			Bus:        binary.LittleEndian.Uint32(response.Bytes()[offset : offset+4]),
			Device:     binary.LittleEndian.Uint32(response.Bytes()[offset+4 : offset+8]),
			Function:   binary.LittleEndian.Uint32(response.Bytes()[offset+8 : offset+12]),
			VendorID:   binary.LittleEndian.Uint16(response.Bytes()[offset+12 : offset+14]),
			DeviceID:   binary.LittleEndian.Uint16(response.Bytes()[offset+14 : offset+16]),
			ClassCode:  binary.LittleEndian.Uint32(response.Bytes()[offset+16 : offset+20]),
			RevisionID: response.Bytes()[offset+20],
		}
	}

	return devices
}

// 来自接口: Packeter 第36个签名
// PerformSMIOperation 执行系统管理模式(SMI)操作
// IOCTL: 是, IOCTL_PERFORM_SMI_OPERATION (0x825)
// 通信: 带SMI类型的设备I/O
// 用户命令: CommandSmi -> 内核: ExtensionCommandPerformSmiOperation
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/smi.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用PerformSMIOperation() → IOCTL_PERFORM_SMI_OPERATION → ExtensionCommandPerformSmiOperation
func (p *Packet) PerformSMIOperation(operation SMIType) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerSmiRequest{
		Operation: operation,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlPerformSmiOperation)
}

// 来自接口: Packeter 第37个签名
// HideDebugger 隐藏调试器
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x826)
// 通信: 带隐藏标志的设备I/O
// 用户命令: CommandHide -> 内核: ExtensionCommandHideAndUnhideToTransparentTheDebugger
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/hide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用HideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → ExtensionCommandHideAndUnhideToTransparentTheDebugger
func (p *Packet) HideDebugger() {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerHideRequest{
		IsHide: true,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerHideAndUnhideToTransparentTheDebugger)
}

// 来自接口: Packeter 第38个签名
// UnhideDebugger 取消隐藏调试器
// IOCTL: 是, IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER (0x826)
// 通信: 带取消隐藏标志的设备I/O
// 用户命令: CommandUnhide -> 内核: ExtensionCommandHideAndUnhideToTransparentTheDebugger
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/hide.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UnhideDebugger() → IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER → ExtensionCommandHideAndUnhideToTransparentTheDebugger
func (p *Packet) UnhideDebugger() {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerHideRequest{
		IsHide: false,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerHideAndUnhideToTransparentTheDebugger)
}

// 来自接口: Packeter 第39个签名
// BringPagesIn 将页面调入内存
// IOCTL: 是, IOCTL_DEBUGGER_BRING_PAGES_IN (0x827)
// 通信: 带地址范围和进程ID的设备I/O
// 用户命令: CommandBring -> 内核: ExtensionCommandBringPagesIn
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/bring.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用BringPagesIn() → IOCTL_DEBUGGER_BRING_PAGES_IN → ExtensionCommandBringPagesIn
func (p *Packet) BringPagesIn(fromAddr, toAddr uint64, pid uint32) {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	req := DebuggerBringPagesInRequest{
		FromAddress: fromAddr,
		ToAddress:   toAddr,
		ProcessId:   pid,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlDebuggerBringPagesIn)
}

// 来自接口: Packeter 第40个签名
// EditMemory 编辑内存
// IOCTL: 是, IOCTL_DEBUGGER_EDIT_MEMORY (0x80A)
// 通信: 带进程ID、地址和数据的设备I/O
// 用户命令: CommandDbgWrite -> 内核: DebuggerCommandEditMemory
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/dbgwrite.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用EditMemory() → IOCTL_DEBUGGER_EDIT_MEMORY → DebuggerCommandEditMemory
func (p *Packet) EditMemory(pid uint32, address uint64, data []byte) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerEditMemoryRequest{
		ProcessId:      pid,
		Address:        address,
		ByteCount:      uint32(len(data)),
		MemoryType:     MemoryTypeVirtual,
		IsDebuggee:     true,
		IsPhysical:     false,
		Is32BitProcess: false,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, data)
	p.driver.SendReceive(buf, IoctlDebuggerEditMemory)
}

func (p *Packet) ReadMemory(pid uint32, address uint64, size uint32, memoryType MemoryType) []byte {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return nil
	}

	req := DebuggerReadMemoryRequest{
		ProcessId:      pid,
		Address:        address,
		ByteCount:      size,
		MemoryType:     memoryType,
		IsDebuggee:     true,
		IsPhysical:     memoryType == MemoryTypePhysical,
		Is32BitProcess: false,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerReadMemory)
	return response.Bytes()
}

func (p *Packet) WriteMemory(pid uint32, address uint64, data []byte) {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	req := DebuggerEditMemoryRequest{
		ProcessId:      pid,
		Address:        address,
		ByteCount:      uint32(len(data)),
		MemoryType:     MemoryTypeVirtual,
		IsDebuggee:     true,
		IsPhysical:     false,
		Is32BitProcess: false,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, data)
	p.driver.SendReceive(buf, IoctlDebuggerEditMemory)
}

func (p *Packet) RegisterEvent(event *Event) {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	p.EventManager.RegisterEvent(event)
}

// 来自接口: Packeter 第41个签名
func (p *Packet) SearchMemory(pid uint32, address uint64, size uint32, pattern []byte) []uint64 {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	req := DebuggerSearchMemoryRequest{
		ProcessId:   pid,
		Address:     address,
		Size:        size,
		PatternSize: uint32(len(pattern)),
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	binary.Write(buf, binary.LittleEndian, pattern)
	response := p.driver.SendReceive(buf, IoctlDebuggerSearchMemory)

	resultCount := binary.LittleEndian.Uint32(response.Bytes()[0:4])
	results := make([]uint64, resultCount)
	for i := range resultCount {
		results[i] = binary.LittleEndian.Uint64(response.Bytes()[4+i*8 : 12+i*8])
	}

	return results
}

// 来自接口: Packeter 第42个签名
func (p *Packet) FlushBuffers() FlushResult {
	var result FlushResult
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	response := p.driver.SendReceive(nil, IoctlDebuggerFlushLoggingBuffers)

	result.KernelStatus = binary.LittleEndian.Uint32(response.Bytes()[0:4])
	result.CountOfMessagesThatSetAsReadFromVmxRoot = binary.LittleEndian.Uint32(response.Bytes()[4:8])
	result.CountOfMessagesThatSetAsReadFromVmxNonRoot = binary.LittleEndian.Uint32(response.Bytes()[8:12])

	return result
}

// 来自接口: Packeter 第43个签名
func (p *Packet) AttachProcess(pid uint32) {
	if !p.IsConnected() {
		mylog.Warning("driver not available")
		return
	}

	req := DebuggerAttachDetachUserModeProcessRequest{
		ProcessId: pid,
		Action:    DebuggerAttachDetachUserModeProcessActionAttach,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	response := p.driver.SendReceive(buf, IoctlDebuggerAttachDetachUserModeProcess)

	var result DebuggerAttachDetachUserModeProcessResponse
	binary.Read(response, binary.LittleEndian, &result)
	if result.KernelStatus != 0 {
		mylog.Warning(result.KernelStatus)
		return
	}
}

// 来自接口: Packeter 第44个签名
// DetachProcess 分离进程
// IOCTL: 是, IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x808)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandDetach -> 内核: DebuggerCommandAttachDetachUserModeProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/meta-commands/detach.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用DetachProcess() → IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS → DebuggerCommandAttachDetachUserModeProcess
func (p *Packet) DetachProcess(pid uint32) {
	return
}

// 来自接口: Packeter 第45个签名
// ChangeProcess 切换进程上下文
// IOCTL: 是, IOCTL_DEBUGGER_SWITCH_PROCESS (0x828)
// 通信: 带进程ID的设备I/O
// 用户命令: CommandProcess -> 内核: ExtensionCommandSwitchProcess
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/process.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ChangeProcess() → IOCTL_DEBUGGER_SWITCH_PROCESS → ExtensionCommandSwitchProcess
func (p *Packet) ChangeProcess(pid uint32) {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	mylog.Info(pid)

	req := DebuggerSwitchProcessRequest{
		ProcessId: pid,
	}

	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlSwitchProcess)
}

// 来自接口: Packeter 第46个签名
// ChangeThread 切换线程上下文
// IOCTL: 是, IOCTL_DEBUGGER_SWITCH_THREAD (0x829)
// 通信: 带线程ID的设备I/O
// 用户命令: CommandThread -> 内核: ExtensionCommandSwitchThread
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/extension-commands/thread.cpp
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用ChangeThread() → IOCTL_DEBUGGER_SWITCH_THREAD → ExtensionCommandSwitchThread
func (p *Packet) ChangeThread(tid uint32) {
	if !p.IsConnected() {
		mylog.Check("driver not available")
	}

	mylog.Info(tid)

	req := DebuggerSwitchThreadRequest{
		ThreadId: tid,
	}
	mylog.Check(req.Validate())

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, &req)
	p.driver.SendReceive(buf, IoctlSwitchThread)
}

// 来自接口: Packeter 第47个签名
func (p *Packet) ConnectSerial(port string, baudrate uint32) {
	mylog.Info(port, baudrate)

	p.mu.Lock()
	defer p.mu.Unlock()

	p.IsSerialConnectedToRemoteDebuggee = true
	p.IsConnectedToRemoteDebuggee = true
}

// 来自接口: Packeter 第48个签名
func (p *Packet) DisconnectSerial() {
	mylog.Info("断开串口连接")

	p.mu.Lock()
	defer p.mu.Unlock()

	p.IsSerialConnectedToRemoteDebuggee = false
	p.IsConnectedToRemoteDebuggee = false
}

// UpdateAllPages 更新所有UI组件
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/debugger.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用UpdateAllPages() → 遍历UIComponents列表 → 调用每个组件的Update()方法
func (p *Packet) UpdateAllPages() {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, component := range p.UIComponents {
		mylog.Check(component.Update())
	}
}

// GetEventChan 获取事件通道
// IOCTL: 否, 本地操作
// 通信: 无
// 源代码: hyperdbg/debugger.go
// 与用户模式/内核模式比较: ✅ 一致，两种模式都支持
// 执行过程: 调用GetEventChan() → 返回EventChan通道
func (p *Packet) GetEventChan() <-chan *DebugEvent {
	return p.EventChan
}

func (p *Packet) StartEventLoop() {
	if p.eventLoop == nil {
		p.eventLoop = NewEventLoop(p.EventManager, p.driver)
	}
	p.eventLoop.Start()
}

func (p *Packet) StopEventLoop() {
	if p.eventLoop == nil {
		return
	}
	p.eventLoop.Stop()
}

func (p *Packet) IsEventLoopRunning() bool {
	if p.eventLoop == nil {
		return false
	}
	return p.eventLoop.IsRunning()
}

func (p *Packet) RegisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	if p.eventLoop == nil {
		p.eventLoop = NewEventLoop(p.EventManager, p.driver)
	}
	p.eventLoop.RegisterHandler(eventType, handler)
}

func (p *Packet) UnregisterEventHandler(eventType EventType, handler func(*DebugEvent)) {
	if p.eventLoop == nil {
		return
	}
	p.eventLoop.UnregisterHandler(eventType, handler)
}

func (p *Packet) Events() []Event {
	if p.EventManager == nil {
		return nil
	}
	debugEvents := p.EventManager.GetDebugEvents()
	events := make([]Event, 0, len(debugEvents))
	for _, de := range debugEvents {
		events = append(events, Event{
			Type: de.Type,
			Pid:  de.ProcessId,
		})
	}
	return events
}

func (p *Packet) RemoveEvent(eventTag uint64) {
	if p.EventManager == nil {
		return
	}
	p.EventManager.ModifyDebugEvent(eventTag, ModifyClear)
}

func (p *Packet) ModifyEvent(eventTag uint64, modifications *EventModifications) {
	if p.EventManager == nil {
		return
	}
	if modifications.Enable != nil && *modifications.Enable {
		p.EventManager.ModifyDebugEvent(eventTag, ModifyEnable)
	}
	if modifications.Disable != nil && *modifications.Disable {
		p.EventManager.ModifyDebugEvent(eventTag, ModifyDisable)
	}
}

func (p *Packet) SetEventMode(mode EventMode) {
	if p.EventManager == nil {
		return
	}
	p.EventManager.SetEventMode(mode)
}
