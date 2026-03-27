# HyperDbg API 设计文档

## 文件结构

```
HyperDbg_rust/
├── doc.go                 // 包文档说明
├── hyperdbg.go            // 导出所有公共类型和函数
├── config.go              // 常量 (ProtocolVersion, DefaultPort, DriverHTTPHost, DriverHTTPPort, HTTPTimeout)
├── types.go               // 类型别名 (ProcessId, ThreadId, Address, PhysicalAddress)
├── response.go            // Response[T] + EventCallback
│
├── packet.go              // Empty + ResponseType + Packet 结构体 + SendReceive + 所有请求方法
├── event_handlers.go      // EventHandlers 结构体 + Dispatch + OnBreakpoint/OnException 等注册方法
│
├── register.go            // RegisterState (共享结构体，被多个事件引用)
│
├── event_process.go       // ProcessInfo + ProcessEvent
├── event_thread.go        // ThreadInfo + ThreadEvent
├── event_module.go        // ModuleInfo + ModuleEvent
├── event_breakpoint.go    // BreakpointInfo + BreakpointType + BreakpointEvent
├── event_memory.go        // MemoryRegion + MemoryAccessType + MemoryAccessEvent
├── event_callstack.go     // CallStackFrame
├── event_symbol.go        // SymbolInfo
├── event_vmx.go           // VmxCapabilities + VmxExitReason + VmxExitEvent
├── event_debug_state.go   // DebugState
├── event_exception.go     // ExceptionCode + ExceptionEvent
├── event_syscall.go       // SyscallEvent
├── event_debug_print.go   // LogLevel + DebugPrintEvent
├── event_trap.go          // TrapEvent
├── event_hidden_hook.go   // HiddenHookEvent
├── event_cpuid.go         // CpuidEvent
├── event_tsc.go           // TscEvent
├── event_cr_access.go     // CrAccessEvent
├── event_dr_access.go     // DrAccessEvent
├── event_io_port.go       // IoPortEvent
├── event_msr.go           // MsrEvent
├── event_ept_violation.go // EptViolationEvent
│
├── cmd/
│   ├── hyperdbg/
│   │   └── main.go        // 主程序入口
│   ├── mcp/
│   │   └── mcp.go         // MCP 服务器
│   └── rustgen/
│       └── main.go        // Rust 代码生成器
│
├── walker/
│   ├── walker.go          // 代码分析器
│   └── walker_test.go     // 测试
│
├── rust-driver/           // Rust 驱动
│   ├── Cargo.toml         // 主驱动 crate 配置
│   ├── lib.rs             // 驱动入口 (DriverEntry)
│   ├── build.rs           // 构建脚本
│   ├── build.ps1          // PowerShell 构建脚本
│   │
│   ├── common/            // 通用模块（自动生成）
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── lib.rs
│   │       ├── types_gen/     // 生成的类型 (分拆)
│   │       │   ├── mod.rs         // 导出所有类型
│   │       │   ├── common.rs      // ProcessId, ThreadId, Address 等基础类型
│   │       │   ├── register.rs    // RegisterState
│   │       │   ├── response.rs    // Response[T], Empty
│   │       │   ├── event.rs       // DebuggerEvent, DebuggerEventType
│   │       │   ├── event_breakpoint.rs   // BreakpointType + BreakpointEvent
│   │       │   ├── event_exception.rs    // ExceptionCode + ExceptionEvent
│   │       │   ├── event_memory.rs       // MemoryAccessType + MemoryAccessEvent
│   │       │   ├── event_syscall.rs      // SyscallEvent
│   │       │   ├── event_process.rs      // ProcessInfo + ProcessEvent
│   │       │   ├── event_thread.rs       // ThreadInfo + ThreadEvent
│   │       │   ├── event_module.rs       // ModuleInfo + ModuleEvent
│   │       │   ├── event_debug_print.rs  // LogLevel + DebugPrintEvent
│   │       │   ├── event_vmx.rs          // VmxExitReason + VmxExitEvent
│   │       │   ├── event_trap.rs         // TrapEvent
│   │       │   ├── event_hidden_hook.rs  // HiddenHookEvent
│   │       │   ├── event_cpuid.rs        // CpuidEvent
│   │       │   ├── event_tsc.rs          // TscEvent
│   │       │   ├── event_cr_access.rs    // CrAccessEvent
│   │       │   ├── event_dr_access.rs    // DrAccessEvent
│   │       │   ├── event_io_port.rs      // IoPortEvent
│   │       │   ├── event_msr.rs          // MsrEvent
│   │       │   └── event_ept_violation.rs // EptViolationEvent
│   │       │
│   │       └── handlers_gen/   // 生成的处理器 (分拆)
│   │           ├── mod.rs         // 导出 + EventQueue + emit_*_event
│   │           ├── router.rs      // DebuggerApi trait + dispatch_api
│   │           └── emit.rs        // emit_*_event 函数
│   │
│   ├── net/               // 网络模块
│   │   ├── Cargo.toml
│   │   ├── README.md
│   │   ├── BSOD_FIX_REPORT.md
│   │   └── src/
│   │       ├── lib.rs
│   │       ├── http.rs
│   │       ├── json.rs
│   │       └── util.rs
│   │
│   ├── hyperhv/           // 虚拟化层
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── lib.rs
│   │       ├── vmm/           // VMX/EPT 核心
│   │       ├── hooks.rs       // Hook 机制
│   │       ├── events.rs      // 事件系统
│   │       ├── processor.rs   // 处理器管理
│   │       └── ...            // 30+ 其他模块
│   │
│   ├── hyperkd/           // 调试器核心
│   │   ├── Cargo.toml
│   │   └── src/
│   │       ├── lib.rs
│   │       ├── debugger.rs    // 调试器核心
│   │       ├── commands.rs    // 命令处理
│   │       ├── network.rs     // 网络通信
│   │       └── ...            // 15+ 其他模块
│   │
│   ├── kd/                // 内核调试器
│   │   ├── Cargo.toml
│   │   └── src/
│   │       └── lib.rs         // KD 上下文和断点管理
│   │
│   ├── pdbex/             // PDB 解析器
│   │   ├── Cargo.toml
│   │   └── src/
│   │       └── lib.rs         // 符号解析、类型重建
│   │
│   ├── disassembler/      // 反汇编器
│   │   ├── Cargo.toml
│   │   └── src/
│   │       └── lib.rs
│   │
│   ├── driver-framework/  // 驱动框架
│   │   ├── Cargo.toml
│   │   ├── out/                // WDK 绑定（自动生成）
│   │   │   ├── ntddk.rs        // 内核函数
│   │   │   ├── types.rs        // 类型定义
│   │   │   └── constants.rs    // 常量定义
│   │   └── src/
│   │       ├── lib.rs
│   │       ├── device.rs       // 设备管理
│   │       ├── ioctl.rs        // IOCTL 处理
│   │       ├── ffi.rs          // FFI 接口
│   │       ├── logger.rs       // 日志宏
│   │       └── utils.rs        // 工具函数
│   │
│   ├── logger/            // 日志库
│   │   ├── Cargo.toml
│   │   └── src/
│   │       └── lib.rs
│   │
│   ├── examples/          // 示例驱动
│   │   ├── netdemo/       // 网络驱动示例
│   │   ├── sysdemo/       // 系统驱动示例
│   │   ├── sample-wdm-driver/
│   │   ├── sample-kmdf-driver/
│   │   └── sample-umdf-driver/
│   │
│   └── doc/               // 文档
│       ├── RUST_INSTALL_README.md
│       ├── bug_fixes.md
│       ├── install-rust-mingw.ps1
│       └── setup_ewdk_env.ps1
│
├── rust-driver/IMPLEMENTATION_PROGRESS.md  // 实现进度
├── api_design.md          // 本文档
├── README.md              // 项目说明
├── go.mod                 // Go 模块
├── go.sum                 // Go 依赖
├── fmt.cmd                // 格式化脚本
└── .gitignore             // Git 忽略
```

## C++ 源代码参考路径

**重要：以下路径为 C++ 原始实现的参考位置，用于验证 Rust 实现的正确性**

C++ 源代码位于压缩包中，需要先解压：
```
doc/cpp/HyperDbg.tar.zst
```

解压后的路径结构：
```
doc/cpp/HyperDbgUnified/HyperDbg/
├── hprdbghv/               # Hypervisor C++ 源码 (对应 hyperhv)
│   ├── code/
│   │   ├── vmm/
│   │   ├── hooks/
│   │   └── ...
│   └── header/
│
├── hyperkd/                # Debugger C++ 源码 (对应 hyperkd)
│   ├── code/
│   │   ├── debugger/
│   │   └── ...
│   └── header/
│
└── libhyperdbg/            # 用户模式库 (对应 Go Communicator)
    └── code/
```

## 通信架构

```
┌─────────────────────────────────────────────────────────────┐
│                        Go 客户端                             │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  packet.go                    event_handlers.go             │
│  ├─ SendReceive()             ├─ EventHandlers              │
│  ├─ Initialize()              ├─ Dispatch()                 │
│  ├─ SetBreakpoint()           └─ 处理驱动推送的事件          │
│  ├─ ReadMemory()                                           │
│  └─ ...                                                    │
│                                                             │
│     ↓ HTTP/JSON 请求            ↑ JSON 事件                  │
│     ↓ :50080                    ↑                           │
├─────────────────────────────────────────────────────────────┤
│                     Rust 内核驱动                            │
│  ├─ net/ (WSK HTTP Server :50080)                          │
│  │   ├─ HTTP 请求解析                                       │
│  │   └─ JSON 序列化/反序列化                                │
│  ├─ common/handlers_gen/                                    │
│  │   ├─ dispatch_api() 路由                                 │
│  │   └─ DebuggerApi trait                                   │
│  └─ emit_*_event() 推送事件                                 │
└─────────────────────────────────────────────────────────────┘
```

## 事件处理流程

```
驱动 emit_*_event()
        ↓
    EventQueue (handlers_gen/mod.rs)
        ↓
客户端轮询获取事件 (poll_events)
        ↓
    Packet 接收
        ↓
    EventHandlers.Dispatch()
        ↓
    对应 handler 执行
```

## 三位一体设计

每个事件文件包含：枚举 + 结构体 + 事件结构体

示例 `event_exception.go`:
```go
type ExceptionCode uint32

const (
    ExceptionCodeDivideError ExceptionCode = 0
    ExceptionCodeDebug ExceptionCode = 1
    // ...
)

type ExceptionEvent struct {
    Header        EventHeader
    ExceptionCode ExceptionCode
    Address       Address
    ErrorCode     uint64
    Registers     RegisterState
}
```

## 核心类型定义

### config.go
```go
const (
    ProtocolVersion = 1
    DefaultPort     = 9527
    DriverHTTPHost  = "127.0.0.1"
    DriverHTTPPort  = 50080
    HTTPTimeout     = 10 * time.Second
)
```

### types.go
```go
type ProcessId = uint32
type ThreadId = uint32
type Address = uint64
type PhysicalAddress = uint64
```

### response.go
```go
type Response[T ResponseType] struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    T      `json:"data,omitempty"`
}

type EventCallback func(event any)
```

### packet.go
```go
type Empty = struct{}

type ResponseType interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string | ~bool |
    []byte | *RegisterState | []ProcessInfo | []ThreadInfo | []ModuleInfo | []BreakpointInfo |
    Empty
}

func SendReceive[T ResponseType](p *Packet, jsonData []byte) *Response[T]
```

### event_handlers.go
```go
type DebuggerEventType int

const (
    DebuggerEventBreakpoint DebuggerEventType = iota
    DebuggerEventException
    DebuggerEventMemoryAccess
    DebuggerEventSyscall
    DebuggerEventProcessCreate
    DebuggerEventProcessExit
    DebuggerEventThreadCreate
    DebuggerEventThreadExit
    DebuggerEventModuleLoad
    DebuggerEventModuleUnload
    DebuggerEventDebugPrint
    DebuggerEventVmxExit
    DebuggerEventTrap
    DebuggerEventHiddenHook
    DebuggerEventCpuid
    DebuggerEventTsc
    DebuggerEventCrAccess
    DebuggerEventDrAccess
    DebuggerEventIoPort
    DebuggerEventMsr
    DebuggerEventEptViolation
)

type EventHandlers struct {
    Breakpoint    func(BreakpointEvent) error
    Exception     func(ExceptionEvent) error
    MemoryAccess  func(MemoryAccessEvent) error
    Syscall       func(SyscallEvent) error
    ProcessCreate func(ProcessEvent) error
    ProcessExit   func(ProcessEvent) error
    ThreadCreate  func(ThreadEvent) error
    ThreadExit    func(ThreadEvent) error
    ModuleLoad    func(ModuleEvent) error
    ModuleUnload  func(ModuleEvent) error
    DebugPrint    func(DebugPrintEvent) error
    VmxExit       func(VmxExitEvent) error
    Trap          func(TrapEvent) error
    HiddenHook    func(HiddenHookEvent) error
    Cpuid         func(CpuidEvent) error
    Tsc           func(TscEvent) error
    CrAccess      func(CrAccessEvent) error
    DrAccess      func(DrAccessEvent) error
    IoPort        func(IoPortEvent) error
    Msr           func(MsrEvent) error
    EptViolation  func(EptViolationEvent) error
}

func (h *EventHandlers) Dispatch(event DebuggerEvent) error
```

## API 方法总览

| 方法 | Action | Go形参 | 发送JSON字段 | 返回数据 | 返回类型 |
|------|--------|--------|-------------|---------|---------|
| Initialize | initialize | 无 | action | - | Empty |
| Terminate | terminate | 无 | action | - | Empty |
| AttachProcess | attach_process | processID uint32 | action, process_id | - | Empty |
| DetachProcess | detach_process | 无 | action | - | Empty |
| SetBreakpoint | set_breakpoint | address uint64, bpType BreakpointType | action, address, type | - | Empty |
| RemoveBreakpoint | remove_breakpoint | breakpointID uint64 | action, breakpoint_id | - | Empty |
| Continue | continue | 无 | action | - | Empty |
| Pause | pause | 无 | action | - | Empty |
| StepInto | step_into | 无 | action | - | Empty |
| StepOver | step_over | 无 | action | - | Empty |
| StepOut | step_out | 无 | action | - | Empty |
| ReadMemory | read_memory | address uint64, size uint32 | action, address, size | []byte | []byte |
| WriteMemory | write_memory | address uint64, data []byte | action, address, data | - | Empty |
| ReadRegisters | read_registers | 无 | action | *RegisterState | *RegisterState |
| WriteRegisters | write_registers | regs *RegisterState | action, data | - | Empty |
| GetProcessList | get_process_list | 无 | action | []ProcessInfo | []ProcessInfo |
| GetThreadList | get_thread_list | processID uint32 | action, process_id | []ThreadInfo | []ThreadInfo |
| GetModuleList | get_module_list | processID uint32 | action, process_id | []ModuleInfo | []ModuleInfo |

## format 标签规范

| format 值 | 说明 | Go 格式化 | Rust 解析 |
|-----------|------|-----------|-----------|
| `hex` | 十六进制字符串 | `fmt.Sprintf("0x%x", v)` | parse with 0x prefix |
| `dec` | 十进制字符串 | `fmt.Sprintf("%d", v)` | parse as decimal |
| `int` | 直接整数 | 直接传值 | 直接解析 |

## 事件类型映射

| DebuggerEventType | 事件结构体 | 文件 |
|-------------------|-----------|------|
| DebuggerEventBreakpoint | BreakpointEvent | event_breakpoint.go |
| DebuggerEventException | ExceptionEvent | event_exception.go |
| DebuggerEventMemoryAccess | MemoryAccessEvent | event_memory.go |
| DebuggerEventSyscall | SyscallEvent | event_syscall.go |
| DebuggerEventProcessCreate | ProcessEvent | event_process.go |
| DebuggerEventProcessExit | ProcessEvent | event_process.go |
| DebuggerEventThreadCreate | ThreadEvent | event_thread.go |
| DebuggerEventThreadExit | ThreadEvent | event_thread.go |
| DebuggerEventModuleLoad | ModuleEvent | event_module.go |
| DebuggerEventModuleUnload | ModuleEvent | event_module.go |
| DebuggerEventDebugPrint | DebugPrintEvent | event_debug_print.go |
| DebuggerEventVmxExit | VmxExitEvent | event_vmx.go |
| DebuggerEventTrap | TrapEvent | event_trap.go |
| DebuggerEventHiddenHook | HiddenHookEvent | event_hidden_hook.go |
| DebuggerEventCpuid | CpuidEvent | event_cpuid.go |
| DebuggerEventTsc | TscEvent | event_tsc.go |
| DebuggerEventCrAccess | CrAccessEvent | event_cr_access.go |
| DebuggerEventDrAccess | DrAccessEvent | event_dr_access.go |
| DebuggerEventIoPort | IoPortEvent | event_io_port.go |
| DebuggerEventMsr | MsrEvent | event_msr.go |
| DebuggerEventEptViolation | EptViolationEvent | event_ept_violation.go |

## 各文件详细内容

### doc.go
```go
// Package hyperdbgrust provides a Go client for HyperDbg debugger.
//
// This package communicates with the HyperDbg Rust driver via HTTP/JSON
// and provides event handling for debugger events.
package hyperdbgrust
```

### hyperdbg.go
```go
package hyperdbgrust

// 导出所有公共类型
type (
    // 基础类型
    ProcessId       = ProcessId
    ThreadId        = ThreadId
    Address         = Address
    PhysicalAddress = PhysicalAddress
    
    // 响应类型
    Response[T ResponseType] = Response[T]
    Empty                    = Empty
    
    // 事件类型
    DebuggerEventType = DebuggerEventType
    DebuggerEvent     = DebuggerEvent
    EventHandlers     = EventHandlers
    
    // 所有事件结构体
    BreakpointEvent    = BreakpointEvent
    ExceptionEvent     = ExceptionEvent
    MemoryAccessEvent  = MemoryAccessEvent
    SyscallEvent       = SyscallEvent
    ProcessEvent       = ProcessEvent
    ThreadEvent        = ThreadEvent
    ModuleEvent        = ModuleEvent
    DebugPrintEvent    = DebugPrintEvent
    VmxExitEvent       = VmxExitEvent
    TrapEvent          = TrapEvent
    HiddenHookEvent    = HiddenHookEvent
    CpuidEvent         = CpuidEvent
    TscEvent           = TscEvent
    CrAccessEvent      = CrAccessEvent
    DrAccessEvent      = DrAccessEvent
    IoPortEvent        = IoPortEvent
    MsrEvent           = MsrEvent
    EptViolationEvent  = EptViolationEvent
    
    // 信息结构体
    ProcessInfo  = ProcessInfo
    ThreadInfo   = ThreadInfo
    ModuleInfo   = ModuleInfo
    RegisterState = RegisterState
)

// 导出公共函数
var (
    NewPacket = NewPacket
)
```

### register.go
```go
package hyperdbgrust

type RegisterState struct {
    RAX    uint64 `json:"RAX"`
    RBX    uint64 `json:"RBX"`
    RCX    uint64 `json:"RCX"`
    RDX    uint64 `json:"RDX"`
    RSI    uint64 `json:"RSI"`
    RDI    uint64 `json:"RDI"`
    RBP    uint64 `json:"RBP"`
    RSP    uint64 `json:"RSP"`
    R8     uint64 `json:"R8"`
    R9     uint64 `json:"R9"`
    R10    uint64 `json:"R10"`
    R11    uint64 `json:"R11"`
    R12    uint64 `json:"R12"`
    R13    uint64 `json:"R13"`
    R14    uint64 `json:"R14"`
    R15    uint64 `json:"R15"`
    RIP    uint64 `json:"RIP"`
    RFLAGS uint64 `json:"RFLAGS"`
    CR0    uint64 `json:"CR0"`
    CR2    uint64 `json:"CR2"`
    CR3    uint64 `json:"CR3"`
    CR4    uint64 `json:"CR4"`
    DR0    uint64 `json:"DR0"`
    DR1    uint64 `json:"DR1"`
    DR2    uint64 `json:"DR2"`
    DR3    uint64 `json:"DR3"`
    DR6    uint64 `json:"DR6"`
    DR7    uint64 `json:"DR7"`
    GDTR   uint64 `json:"GDTR"`
    GSBase uint64 `json:"GSBase"`
    FSBase uint64 `json:"FSBase"`
}
```

### event_process.go
```go
package hyperdbgrust

type ProcessInfo struct {
    ProcessID   uint32 `json:"process_id"`
    ImageName   string `json:"image_name"`
    BaseAddress string `json:"base_address"`
    Size        uint64 `json:"size"`
    CR3         uint64 `json:"cr3"`
}

type ProcessEvent struct {
    Header          EventHeader `json:"header"`
    ProcessInfo     ProcessInfo `json:"process_info"`
    ParentProcessID ProcessId   `json:"parent_process_id"`
    IsCreate        bool        `json:"is_create"`
}
```

### event_thread.go
```go
package hyperdbgrust

type ThreadInfo struct {
    ThreadID     uint32 `json:"thread_id"`
    ProcessID    uint32 `json:"process_id"`
    StartAddress string `json:"start_address"`
    TEB          string `json:"teb"`
    State        uint32 `json:"state"`
}

type ThreadEvent struct {
    Header     EventHeader `json:"header"`
    ThreadInfo ThreadInfo  `json:"thread_info"`
    IsCreate   bool        `json:"is_create"`
}
```

### event_module.go
```go
package hyperdbgrust

type ModuleInfo struct {
    BaseAddress string `json:"base_address"`
    Size        uint64 `json:"size"`
    Name        string `json:"name"`
    Path        string `json:"path"`
}

type ModuleEvent struct {
    Header     EventHeader `json:"header"`
    ModuleInfo ModuleInfo  `json:"module_info"`
    IsLoad     bool        `json:"is_load"`
}
```

### event_breakpoint.go
```go
package hyperdbgrust

type BreakpointType uint32

const (
    BreakpointSoftware BreakpointType = 0
    BreakpointHardware BreakpointType = 1
    BreakpointHidden   BreakpointType = 2
    BreakpointEpt      BreakpointType = 3
)

type BreakpointInfo struct {
    ID        uint64 `json:"id"`
    Address   string `json:"address"`
    Type      int    `json:"type"`
    ProcessID uint32 `json:"process_id"`
    Enabled   bool   `json:"enabled"`
    HitCount  uint64 `json:"hit_count"`
}

type BreakpointEvent struct {
    Header       EventHeader   `json:"header"`
    Address      Address       `json:"address"`
    BreakpointID uint64        `json:"breakpoint_id"`
    Registers    RegisterState `json:"registers"`
}
```

### event_memory.go
```go
package hyperdbgrust

type MemoryAccessType uint32

const (
    MemoryAccessTypeRead    MemoryAccessType = 0
    MemoryAccessTypeWrite   MemoryAccessType = 1
    MemoryAccessTypeExecute MemoryAccessType = 2
)

type MemoryRegion struct {
    BaseAddress Address `json:"base_address"`
    Size        uint64  `json:"size"`
    Protection  uint32  `json:"protection"`
    State       uint32  `json:"state"`
    Type        uint32  `json:"type_"`
}

type MemoryAccessEvent struct {
    Header          EventHeader      `json:"header"`
    VirtualAddress  Address          `json:"virtual_address"`
    PhysicalAddress PhysicalAddress  `json:"physical_address"`
    AccessType      MemoryAccessType `json:"access_type"`
    Size            uint32           `json:"size"`
    Value           uint64           `json:"value"`
}
```

### event_callstack.go
```go
package hyperdbgrust

type CallStackFrame struct {
    InstructionPointer Address `json:"instruction_pointer"`
    ReturnAddress      Address `json:"return_address"`
    StackPointer       Address `json:"stack_pointer"`
    FramePointer       Address `json:"frame_pointer"`
    ModuleName         string  `json:"module_name"`
    SymbolName         string  `json:"symbol_name,omitempty"`
}
```

### event_symbol.go
```go
package hyperdbgrust

type SymbolInfo struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
    Size    uint64  `json:"size"`
    Module  string  `json:"module"`
}
```

### event_vmx.go
```go
package hyperdbgrust

type VmxExitReason uint32

const (
    VmxExitReasonExceptionNmi          VmxExitReason = 0
    VmxExitReasonExternalInterrupt     VmxExitReason = 1
    VmxExitReasonTripleFault           VmxExitReason = 2
    // ... 完整列表见 types.go
)

type VmxCapabilities struct {
    VMXSupported            bool   `json:"vmx_supported"`
    EPTSupported            bool   `json:"ept_supported"`
    VPIDSupported           bool   `json:"vpid_supported"`
    MSRBitmapSupported      bool   `json:"msr_bitmap_supported"`
    IOBitmapSupported       bool   `json:"io_bitmap_supported"`
    MaxPhysicalAddressWidth uint8  `json:"max_physical_address_width"`
    ProcessorCount          uint32 `json:"processor_count"`
}

type VmxExitEvent struct {
    Header            EventHeader   `json:"header"`
    ExitReason        VmxExitReason `json:"exit_reason"`
    ExitQualification uint64        `json:"exit_qualification"`
    GuestRIP          Address       `json:"guest_rip"`
    GuestRSP          Address       `json:"guest_rsp"`
    InstructionLength uint32        `json:"instruction_length"`
}
```

### event_debug_state.go
```go
package hyperdbgrust

type DebugState uint32

const (
    StateRunning    DebugState = 0
    StatePaused     DebugState = 1
    StateStepping   DebugState = 2
    StateTerminated DebugState = 3
)
```

### event_exception.go
```go
package hyperdbgrust

type ExceptionCode uint32

const (
    ExceptionCodeDivideError        ExceptionCode = 0
    ExceptionCodeDebug              ExceptionCode = 1
    ExceptionCodeNmi                ExceptionCode = 2
    ExceptionCodeBreakpoint         ExceptionCode = 3
    ExceptionCodeOverflow           ExceptionCode = 4
    ExceptionCodeBound              ExceptionCode = 5
    ExceptionCodeInvalidOpcode      ExceptionCode = 6
    ExceptionCodeNoMath             ExceptionCode = 7
    ExceptionCodeDoubleFault        ExceptionCode = 8
    ExceptionCodeCoprocessorSegment ExceptionCode = 9
    ExceptionCodeInvalidTss         ExceptionCode = 10
    ExceptionCodeSegmentNotPresent  ExceptionCode = 11
    ExceptionCodeStackSegment       ExceptionCode = 12
    ExceptionCodeGeneralProtection  ExceptionCode = 13
    ExceptionCodePageFault          ExceptionCode = 14
    ExceptionCodeFloatingPoint      ExceptionCode = 16
    ExceptionCodeAlignmentCheck     ExceptionCode = 17
    ExceptionCodeMachineCheck       ExceptionCode = 18
    ExceptionCodeSimdFloatingPoint  ExceptionCode = 19
)

type ExceptionEvent struct {
    Header        EventHeader   `json:"header"`
    ExceptionCode ExceptionCode `json:"exception_code"`
    Address       Address       `json:"address"`
    ErrorCode     uint64        `json:"error_code"`
    Registers     RegisterState `json:"registers"`
}
```

### event_syscall.go
```go
package hyperdbgrust

type SyscallEvent struct {
    Header        EventHeader `json:"header"`
    SyscallNumber uint32      `json:"syscall_number"`
    Arguments     [8]uint64   `json:"arguments"`
    ReturnValue   uint64      `json:"return_value"`
    IsEntry       bool        `json:"is_entry"`
}
```

### event_debug_print.go
```go
package hyperdbgrust

type LogLevel uint32

const (
    LogLevelTrace LogLevel = 0
    LogLevelDebug LogLevel = 1
    LogLevelInfo  LogLevel = 2
    LogLevelWarn  LogLevel = 3
    LogLevelError LogLevel = 4
)

type DebugPrintEvent struct {
    Header  EventHeader `json:"header"`
    Message string      `json:"message"`
    Level   LogLevel    `json:"level"`
}
```

### event_trap.go
```go
package hyperdbgrust

type TrapEvent struct {
    Header     EventHeader `json:"header"`
    TrapNumber uint32      `json:"trap_number"`
    ErrorCode  uint64      `json:"error_code"`
    Address    Address     `json:"address"`
}
```

### event_hidden_hook.go
```go
package hyperdbgrust

type HiddenHookEvent struct {
    Header      EventHeader      `json:"header"`
    HookAddress Address          `json:"hook_address"`
    HookType    MemoryAccessType `json:"hook_type"`
    Data        []byte           `json:"data"`
}
```

### event_cpuid.go
```go
package hyperdbgrust

type CpuidEvent struct {
    Header  EventHeader `json:"header"`
    Leaf    uint32      `json:"leaf"`
    SubLeaf uint32      `json:"sub_leaf"`
    EAX     uint32      `json:"eax"`
    EBX     uint32      `json:"ebx"`
    ECX     uint32      `json:"ecx"`
    EDX     uint32      `json:"edx"`
}
```

### event_tsc.go
```go
package hyperdbgrust

type TscEvent struct {
    Header    EventHeader `json:"header"`
    TscValue  uint64      `json:"tsc_value"`
    RdtscExit bool        `json:"rdtsc_exit"`
}
```

### event_cr_access.go
```go
package hyperdbgrust

type CrAccessEvent struct {
    Header   EventHeader `json:"header"`
    CrNumber uint32      `json:"cr_number"`
    IsWrite  bool        `json:"is_write"`
    OldValue uint64      `json:"old_value"`
    NewValue uint64      `json:"new_value"`
}
```

### event_dr_access.go
```go
package hyperdbgrust

type DrAccessEvent struct {
    Header   EventHeader `json:"header"`
    DrNumber uint32      `json:"dr_number"`
    IsWrite  bool        `json:"is_write"`
    Value    uint64      `json:"value"`
}
```

### event_io_port.go
```go
package hyperdbgrust

type IoPortEvent struct {
    Header  EventHeader `json:"header"`
    Port    uint16      `json:"port"`
    Size    uint32      `json:"size"`
    IsWrite bool        `json:"is_write"`
    Value   uint32      `json:"value"`
}
```

### event_msr.go
```go
package hyperdbgrust

type MsrEvent struct {
    Header  EventHeader `json:"header"`
    Msr     uint32      `json:"msr"`
    IsWrite bool        `json:"is_write"`
    Value   uint64      `json:"value"`
}
```

### event_ept_violation.go
```go
package hyperdbgrust

type EptViolationEvent struct {
    Header               EventHeader     `json:"header"`
    GuestPhysicalAddress PhysicalAddress `json:"guest_physical_address"`
    GuestVirtualAddress  Address         `json:"guest_virtual_address"`
    Read                 bool            `json:"read"`
    Write                bool            `json:"write"`
    Execute              bool            `json:"execute"`
    Readable             bool            `json:"readable"`
    Writable             bool            `json:"writable"`
    Executable           bool            `json:"executable"`
}
```

## 事件轮询机制

### packet.go 事件轮询方法
```go
// StartEventLoop 启动事件轮询 goroutine
func (p *Packet) StartEventLoop() {
    ticker := time.NewTicker(100 * time.Millisecond)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            events := p.pollEvents()
            for _, event := range events {
                p.handlers.Dispatch(event)
            }
        case <-p.stopChan:
            return
        }
    }
}

// pollEvents 从驱动获取事件
func (p *Packet) pollEvents() []DebuggerEvent {
    data, _ := json.Marshal(map[string]string{"action": "poll_events"})
    resp := SendReceive[[]DebuggerEvent](p, data)
    if resp == nil || !resp.Success {
        return nil
    }
    return resp.Data
}

// StopEventLoop 停止事件轮询
func (p *Packet) StopEventLoop() {
    close(p.stopChan)
}

// OnBreakpoint 注册断点事件处理器
func (p *Packet) OnBreakpoint(handler func(BreakpointEvent) error) {
    p.handlers.Breakpoint = handler
}

// OnException 注册异常事件处理器
func (p *Packet) OnException(handler func(ExceptionEvent) error) {
    p.handlers.Exception = handler
}

// ... 其他 On* 方法
```

### cmd/hyperdbg/main.go 命令行入口
```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
    
    hyperdbgrust "github.com/ddkwork/HyperDbg_rust"
)

func main() {
    debugger := hyperdbgrust.NewPacket()
    
    // 注册事件处理器
    debugger.OnBreakpoint(func(e hyperdbgrust.BreakpointEvent) error {
        fmt.Printf("[断点] 地址: 0x%x, 进程: %d, 线程: %d\n", 
            e.Address, e.Header.ProcessID, e.Header.ThreadID)
        return nil
    })
    
    debugger.OnException(func(e hyperdbgrust.ExceptionEvent) error {
        fmt.Printf("[异常] 代码: %d, 地址: 0x%x\n", e.ExceptionCode, e.Address)
        return nil
    })
    
    // 启动事件轮询
    go debugger.StartEventLoop()
    defer debugger.StopEventLoop()
    
    // 等待信号
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    <-sigChan
}
```

## Go-Rust 类型同步机制

### ⚠️ 重要：修改数据结构时必须更新生成器

**当修改任何共享数据结构时，必须运行生成器以确保 Go 和 Rust 端的一致性！**

**必须运行生成器的情况：**

| 修改类型 | 需要运行的命令 | 说明 |
|----------|---------------|------|
| 修改 `types.go` | `cd cmd/rustgen && go run main.go` | 更新 Rust 类型定义和 API 处理器 |
| 修改 `interfaces.go` | `cd cmd/rustgen && go run main.go` | 自动扫描接口更新 API 处理器 |
| 修改 `packet.go` | `cd cmd/rustgen && go run main.go` | 自动扫描 action 字符串 |

### 自动生成器

项目使用 `cmd/rustgen` 生成器自动同步 Go 和 Rust 之间的类型定义，确保 JSON 通信的一致性。

**生成的文件：**
- `rust-driver/common/src/types_gen/*.rs` - 从 `types.go` 和 `event_*.go` 自动生成的类型定义
- `rust-driver/common/src/handlers_gen/router.rs` - 自动生成的 `DebuggerApi` trait 和 `dispatch_api` 调度器
- `rust-driver/common/src/handlers_gen/emit.rs` - 自动生成的 `emit_*_event` 函数

### 同步流程（自动扫描 AST）

```
┌─────────────────┐     rustgen      ┌─────────────────────────────┐
│   types.go      │ ───────────────► │   common/types_gen/*.rs     │
│  (Go 类型定义)   │                  │  (Rust 类型定义)             │
└─────────────────┘                  └─────────────────────────────┘
                                             
┌─────────────────┐     rustgen      ┌─────────────────────────────┐
│ interfaces.go   │ ───────────────► │  handlers_gen/router.rs     │
│ (接口方法签名)   │   (自动扫描AST)   │  (API 处理器)                │
└─────────────────┘                  └─────────────────────────────┘
        │
        │  ┌─────────────────┐
        └─►│   packet.go     │
           │ (action 字符串)  │
           │  (自动提取)      │
           └─────────────────┘
```

### 何时需要运行生成器

**必须运行生成器的情况：**

1. **修改 `types.go` 后**
   ```bash
   cd cmd/rustgen && go run main.go
   ```
   - 添加/修改/删除结构体
   - 添加/修改/删除枚举
   - 修改字段类型或 JSON 标签

2. **修改 `interfaces.go` 后**
   - 生成器会自动扫描 `Debugger` 接口的方法签名
   - 自动提取返回类型和参数类型
   - 重新运行生成器即可

3. **修改 `packet.go` 后**
   - 生成器会自动扫描 `json.Marshal` 调用中的 action 字符串
   - 无需手动维护方法列表
   - 重新运行生成器即可

### 生成的 Rust 代码

**types_gen/register.rs 示例：**
```rust
#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct RegisterState {
    #[serde(rename = "RAX")]
    pub rax: u64,
    #[serde(rename = "RBX")]
    pub rbx: u64,
    // ...
}
```

**handlers_gen/router.rs 示例：**
```rust
pub trait DebuggerApi {
    fn initialize(&mut self, req: &Request) -> Result<Empty, String>;
    fn read_memory(&mut self, req: &Request) -> Result<Vec<u8>, String>;
    // ...
}

pub fn dispatch_api<T: DebuggerApi>(api: &mut T, body: &[u8]) -> Vec<u8> {
    // 自动路由到对应方法
}
```

**handlers_gen/emit.rs 示例：**
```rust
pub fn emit_breakpoint_event(queue: &mut EventQueue, event: BreakpointEvent) {
    queue.push(DebuggerEvent::Breakpoint(event));
}
```

### 开发指南

#### 添加新的调试器方法

1. 在 `types.go` 中添加必要的类型定义（如果需要）
2. 在 `interfaces.go` 的 `Debugger` 接口中添加新方法
3. 在 `packet.go` 中实现该方法（使用 `json.Marshal` 发送 action）
4. 运行生成器：
   ```bash
   cd cmd/rustgen && go run main.go
   ```
5. 在 Rust 驱动中实现 `DebuggerApi` trait 的新方法

#### 修改现有类型

1. 修改 `types.go` 中的类型定义
2. 运行生成器更新 Rust 类型
3. 确保 Go 端和 Rust 端的 JSON 标签一致
4. 更新测试用例

## 模块状态

| 模块 | 状态 | 说明 |
|------|------|------|
| **common/types_gen** | ✅ 完成 | 自动生成的类型定义 |
| **common/handlers_gen** | ✅ 完成 | 自动生成的 API 调度器 |
| **net** | ✅ 完成 | WSK HTTP Server |
| **driver-framework** | ✅ 完成 | WDM/WDF 框架 |
| **hyperhv** | 🔄 开发中 | Hypervisor 核心（30+ 模块文件） |
| **hyperkd** | 🔄 开发中 | Debugger 核心（15+ 模块文件） |
| **kd** | 🔄 开发中 | 内核调试器 |
| **pdbex** | 🔄 开发中 | PDB 解析器 |
| **disassembler** | 🔄 开发中 | 反汇编器 |

## 相关文档

- [IMPLEMENTATION_PROGRESS.md](rust-driver/IMPLEMENTATION_PROGRESS.md) - 实现进度
- [README.md](README.md) - 项目说明
- [.trae/skills/rust-driver-development/SKILL.md](.trae/skills/rust-driver-development/SKILL.md) - 开发技能
