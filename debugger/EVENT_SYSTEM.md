# HyperDbg 事件系统文档

## 概述

HyperDbg 事件系统是一个强大的、基于事件的调试框架，支持 30 种不同类型的事件，包括应用层调试事件和虚拟机监控器（VMM）级别的事件。该系统允许用户在特定条件发生时触发自定义动作，如中断到调试器、运行脚本或执行自定义代码。

## 事件类型总览

| 类别 | 事件类型 | 描述 | 常量值 |
|------|---------|------|---------|
| **应用层调试事件** | | | |
| | EventStep | 单步执行事件 | 0 |
| | EventBreakpoint | 断点事件 | 1 |
| | EventException | 异常事件 | 2 |
| | EventProcessCreated | 进程创建事件 | 3 |
| | EventProcessExited | 进程退出事件 | 4 |
| | EventThreadCreated | 线程创建事件 | 5 |
| | EventThreadExited | 线程退出事件 | 6 |
| | EventModuleLoaded | 模块加载事件 | 7 |
| | EventModuleUnloaded | 模块卸载事件 | 8 |
| **EPT 内存监控事件** | | | |
| | EptHookReadWriteAndExecute | 读写执行监控 | 9 |
| | EptHookReadWrite | 读写监控 | 10 |
| | EptHookReadAndExecute | 读执行监控 | 11 |
| | EptHookWriteAndExecute | 写执行监控 | 12 |
| | EptHookRead | 只读监控 | 13 |
| | EptHookWrite | 只写监控 | 14 |
| | EptHookExecute | 只执行监控 | 15 |
| **EPT Hook 事件** | | | |
| | EptHookExecDetours | EPT 执行 Hook（Detours） | 16 |
| | EptHookExecCC | EPT 执行 Hook（CC 指令） | 17 |
| **系统调用事件** | | | |
| | SyscallHookEferSyscall | SYSCALL 指令 Hook | 18 |
| | SyscallHookEferSysret | SYSRET 指令 Hook | 19 |
| **指令执行事件** | | | |
| | CpuidInstructionExecution | CPUID 指令执行 | 20 |
| | RdmsrInstructionExecution | RDMSR 指令执行 | 21 |
| | WrmsrInstructionExecution | WRMSR 指令执行 | 22 |
| | InInstructionExecution | IN 指令执行 | 23 |
| | OutInstructionExecution | OUT 指令执行 | 24 |
| | VmcallInstructionExecution | VMCALL 指令执行 | 25 |
| | XsetbvInstructionExecution | XSETBV 指令执行 | 29 |
| **中断/异常事件** | | | |
| | ExceptionOccurred | 异常发生事件 | 26 |
| | ExternalInterruptOccurred | 外部中断事件 | 27 |
| **调试寄存器事件** | | | |
| | DebugRegistersAccessed | 调试寄存器访问 | 28 |
| **定时和性能事件** | | | |
| | TscInstructionExecution | TSC 指令执行 | 30 |
| | PmcInstructionExecution | PMC 指令执行 | 31 |
| **控制寄存器事件** | | | |
| | ControlRegisterModified | 控制寄存器修改 | 32 |
| | ControlRegisterRead | 控制寄存器读取 | 33 |
| | ControlRegister3Modified | CR3 修改事件 | 34 |
| **执行陷阱事件** | | | |
| | TrapExecutionModeChanged | 执行模式改变陷阱 | 35 |
| | TrapExecutionInstructionTrace | 指令跟踪陷阱 | 36 |

## 事件动作类型

事件可以配置以下三种动作之一：

| 动作类型 | 常量值 | 描述 |
|---------|---------|------|
| ActionBreakToDebugger | 0 | 中断到调试器，暂停执行 |
| ActionRunScript | 1 | 运行预定义的脚本 |
| ActionRunCustomCode | 2 | 运行自定义代码 |

## 事件修改操作

支持对已注册的事件进行以下操作：

| 操作类型 | 常量值 | 描述 |
|---------|---------|------|
| ModifyQueryState | 0 | 查询事件状态（启用/禁用） |
| ModifyEnable | 1 | 启用事件 |
| ModifyDisable | 2 | 禁用事件 |
| ModifyClear | 3 | 清除（删除）事件 |

## 事件阶段

事件可以在不同的阶段触发：

| 阶段类型 | 常量值 | 描述 |
|---------|---------|------|
| EventStagePre | 0 | 前置事件：在操作执行前触发 |
| EventStagePost | 1 | 后置事件：在操作执行后触发 |

## 事件选项

每个事件可以配置最多 6 个可选参数：

```go
type EventOptions struct {
    OptionalParam1 uint64  // 可选参数 1
    OptionalParam2 uint64  // 可选参数 2
    OptionalParam3 uint64  // 可选参数 3
    OptionalParam4 uint64  // 可选参数 4
    OptionalParam5 uint64  // 可选参数 5
    OptionalParam6 uint64  // 可选参数 6
}
```

## 事件结构

完整的调试事件结构包含以下字段：

```go
type DebugEvent struct {
    Type              DebugEventType      // 事件类型
    ProcessId         uint32            // 进程 ID
    ThreadId          uint32            // 线程 ID
    CoreId            uint32            // CPU 核心 ID
    EIP               uint64            // 指令指针
    RAX               uint64            // 寄存器 RAX
    RBX               uint64            // 寄存器 RBX
    RCX               uint64            // 寄存器 RCX
    RDX               uint64            // 寄存器 RDX
    RSI               uint64            // 寄存器 RSI
    RDI               uint64            // 寄存器 RDI
    RBP               uint64            // 寄存器 RBP
    RSP               uint64            // 寄存器 RSP
    R8                uint64            // 寄存器 R8
    R9                uint64            // 寄存器 R9
    R10               uint64            // 寄存器 R10
    R11               uint64            // 寄存器 R11
    R12               uint64            // 寄存器 R12
    R13               uint64            // 寄存器 R13
    R14               uint64            // 寄存器 R14
    R15               uint64            // 寄存器 R15
    RFLAGS            uint32            // 标志寄存器
    Timestamp         int64             // 时间戳
    IsEnabled         bool              // 是否启用
    EnableShortCircuiting bool         // 是否启用短路事件
    EventStage       EventCallingStage // 事件阶段
    HasCustomOutput  bool              // 是否有自定义输出
    Tag              uint64            // 事件标签
    Options          EventOptions       // 事件选项
    CountOfActions   uint32            // 动作数量
    CreationTime     int64             // 创建时间
}
```

## 事件管理器

事件管理器（`EventManager`）提供以下功能：

### 核心功能

- **注册事件**：`RegisterEvent(event *DebugEvent) (uint64, error)`
  - 注册新事件并返回事件标签（Tag）
  - 自动分配标签和创建时间

- **修改事件**：`ModifyEvent(tag uint64, modifyType EventModifyType) (bool, error)`
  - 修改单个事件的状态（启用/禁用/清除/查询）

- **批量修改**：`ModifyAllEvents(modifyType EventModifyType) error`
  - 批量修改所有事件的状态

- **查询事件**：`QueryEventState(tag uint64) (bool, error)`
  - 查询指定事件的状态

- **获取事件**：`GetEvent(tag uint64) (*DebugEvent, error)`
  - 获取指定事件的详细信息

- **获取所有事件**：`GetEvents() []*DebugEvent`
  - 获取所有已注册的事件列表

- **清除所有事件**：`ClearAllEvents()`
  - 清除所有已注册的事件

### 回调机制

- **添加回调**：`AddCallback(callback DebugEventCallback)`
  - 注册事件回调函数

- **移除回调**：`RemoveCallback(callback DebugEventCallback)`
  - 移除事件回调函数

- **触发事件**：`TriggerEvent(event *DebugEvent)`
  - 手动触发事件并通知所有回调

### 事件通道

- **事件通道**：`EventChannel() <-chan *DebugEvent`
  - 提供事件通道用于异步处理事件

## 与驱动通信

### IOCTL 接口

HyperDbg 通过 Windows IOCTL 接口与内核驱动通信，主要使用以下 IOCTL 命令：

| IOCTL 命令 | 功能 |
|-------------|------|
| IOCTL_DEBUGGER_REGISTER_EVENT | 注册新事件到内核 |
| IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT | 为事件添加动作 |
| IOCTL_DEBUGGER_MODIFY_EVENTS | 修改事件状态（启用/禁用/清除） |

### 通信流程

1. **注册事件流程**：
   ```
   用户空间 → EventManager.RegisterEvent()
            → 构造事件结构
            → IOCTL_DEBUGGER_REGISTER_EVENT
            → 内核驱动处理
            → 返回事件标签
   ```

2. **修改事件流程**：
   ```
   用户空间 → EventManager.ModifyEvent()
            → 构造修改请求
            → IOCTL_DEBUGGER_MODIFY_EVENTS
            → 内核驱动处理
            → 返回操作结果
   ```

3. **事件触发流程**：
   ```
   内核驱动 → 检测到事件
            → 通知用户空间
            → EventManager.TriggerEvent()
            → 执行回调函数
            → 发送到事件通道
   ```

### 驱动通信示例

```go
// 注册事件到内核
func (em *EventManager) sendEventToKernel(event *DebugEvent) error {
    // 构造事件缓冲区
    buffer := em.constructEventBuffer(event)
    
    // 通过驱动句柄发送
    _, err := em.driverHandle.SendBuffer(buffer)
    if err != nil {
        return fmt.Errorf("failed to send event to kernel: %v", err)
    }
    
    return nil
}

// 修改事件状态
func (em *EventManager) modifyEventInKernel(tag uint64, modifyType EventModifyType) error {
    // 构造修改请求
    request := &ModifyEventRequest{
        Tag:         tag,
        TypeOfAction: modifyType,
    }
    
    // 通过驱动句柄发送
    _, err := em.driverHandle.SendBuffer(request)
    if err != nil {
        return fmt.Errorf("failed to modify event in kernel: %v", err)
    }
    
    return nil
}
```

## Events 命令

HyperDbg 提供 `events` 命令来管理事件：

### 命令语法

```bash
events                              # 显示所有事件
events e [tag|all]                 # 启用事件
events d [tag|all]                 # 禁用事件
events c [tag|all]                 # 清除事件
events sc [on|off]                 # 设置短路事件状态
```

### 使用示例

```bash
# 显示所有事件
events

# 启用标签为 1000 的事件
events e 1000

# 禁用标签为 1000 的事件
events d 1000

# 清除标签为 1000 的事件
events c 1000

# 启用所有事件
events e all

# 禁用所有事件
events d all

# 清除所有事件
events c all

# 启用短路事件
events sc on

# 禁用短路事件
events sc off
```

## 脚本引擎

HyperDbg 提供内置脚本引擎，支持在事件触发时执行脚本。

### 支持的脚本命令

| 命令 | 描述 | 示例 |
|------|------|------|
| print | 打印消息或寄存器值 | `print "Hello"` 或 `print RAX` |
| break | 中断到调试器 | `break` |
| continue | 继续执行 | `continue` |

### 脚本示例

```go
// 示例脚本：打印寄存器信息并中断
script := `
    print "Exception occurred at:"
    print RIP
    print "RAX ="
    print RAX
    print "RBX ="
    print RBX
    break
`

// 编译并执行脚本
engine := NewSimpleScriptEngine()
compiled, err := engine.Compile(script)
if err != nil {
    log.Fatal(err)
}

// 在事件触发时执行
err = compiled.Execute(event)
```

## 使用示例

### 示例 1：注册断点事件

```go
// 创建断点事件
event := &DebugEvent{
    Type:      EventBreakpoint,
    ProcessId: 1234,
    ThreadId:  5678,
    EIP:       0x1000,
    IsEnabled:  true,
}

// 注册事件
tag, err := eventManager.RegisterEvent(event)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Breakpoint event registered with tag: 0x%x\n", tag)
```

### 示例 2：注册 EPT Hook 事件

```go
// 创建 EPT Hook 事件
event := &DebugEvent{
    Type:      EptHookReadWriteAndExecute,
    ProcessId: 0xffffffff,  // 所有进程
    CoreId:    0xffffffff,    // 所有核心
    Options: EventOptions{
        OptionalParam1: 0x1000,  // 起始地址
        OptionalParam2: 0x2000,  // 结束地址
    },
    IsEnabled: true,
}

// 注册事件
tag, err := eventManager.RegisterEvent(event)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("EPT Hook event registered with tag: 0x%x\n", tag)
```

### 示例 3：注册系统调用 Hook 事件

```go
// 创建系统调用 Hook 事件
event := &DebugEvent{
    Type:      SyscallHookEferSyscall,
    ProcessId: 1234,
    EventStage: EventStagePre,  // 在 SYSCALL 执行前触发
    Options: EventOptions{
        OptionalParam1: 0x100,  // 系统调用号
    },
    IsEnabled: true,
}

// 注册事件
tag, err := eventManager.RegisterEvent(event)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Syscall hook event registered with tag: 0x%x\n", tag)
```

### 示例 4：使用回调处理事件

```go
// 注册事件回调
eventManager.AddCallback(func(event *DebugEvent) {
    fmt.Printf("Event triggered: Type=%d, ProcessId=%d, ThreadId=%d\n",
        event.Type, event.ProcessId, event.ThreadId)
    
    // 根据事件类型执行不同操作
    switch event.Type {
    case EventBreakpoint:
        fmt.Printf("Breakpoint hit at 0x%x\n", event.EIP)
    case EventException:
        fmt.Printf("Exception occurred at 0x%x\n", event.EIP)
    case EptHookReadWriteAndExecute:
        fmt.Printf("EPT Hook triggered at 0x%x\n", event.EIP)
    }
})

// 从事件通道异步处理事件
go func() {
    for event := range eventManager.EventChannel() {
        fmt.Printf("Async event: %+v\n", event)
    }
}()
```

### 示例 5：修改事件状态

```go
// 禁用事件
enabled, err := eventManager.ModifyEvent(tag, ModifyDisable)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Event disabled (was enabled: %v)\n", enabled)

// 启用事件
enabled, err = eventManager.ModifyEvent(tag, ModifyEnable)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Event enabled\n")

// 查询事件状态
enabled, err = eventManager.QueryEventState(tag)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Event state: %v\n", enabled)

// 清除事件
_, err = eventManager.ModifyEvent(tag, ModifyClear)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Event cleared\n")
```

## 高级功能

### 短路事件

短路事件允许在事件触发时快速处理，避免不必要的性能开销：

```go
event := &DebugEvent{
    Type:                  EventBreakpoint,
    EnableShortCircuiting: true,  // 启用短路
    ProcessId:             1234,
    IsEnabled:             true,
}
```

### 多核支持

事件可以指定在特定的 CPU 核心上触发：

```go
event := &DebugEvent{
    Type:      EptHookReadWriteAndExecute,
    CoreId:    2,  // 仅在核心 2 上触发
    ProcessId: 1234,
    IsEnabled:  true,
}
```

### 多进程支持

事件可以指定在特定的进程上触发：

```go
event := &DebugEvent{
    Type:      EptHookReadWriteAndExecute,
    ProcessId: 1234,  // 仅在进程 1234 上触发
    IsEnabled:  true,
}
```

### 自定义输出

事件可以配置自定义输出源：

```go
event := &DebugEvent{
    Type:             EventBreakpoint,
    HasCustomOutput:  true,  // 启用自定义输出
    ProcessId:        1234,
    IsEnabled:        true,
}
```

## 性能考虑

1. **事件过滤**：使用 `ProcessId` 和 `CoreId` 字段来限制事件触发范围，减少不必要的处理。

2. **短路事件**：对于高频事件，启用短路事件以提高性能。

3. **异步处理**：使用事件通道进行异步事件处理，避免阻塞主线程。

4. **批量操作**：使用 `ModifyAllEvents` 进行批量操作，减少驱动通信次数。

## 故障排除

### 常见问题

1. **事件未触发**
   - 检查事件是否已启用（`IsEnabled = true`）
   - 确认驱动已正确加载
   - 验证 `ProcessId` 和 `CoreId` 设置正确

2. **性能问题**
   - 减少事件数量
   - 使用短路事件
   - 限制事件触发范围

3. **驱动通信失败**
   - 检查驱动是否已加载
   - 验证驱动句柄是否有效
   - 检查权限是否足够

## 总结

HyperDbg 事件系统提供了一个强大而灵活的调试框架，支持从应用层到 VMM 层面的各种事件。通过合理使用事件类型、动作和选项，用户可以实现复杂的调试逻辑和自动化脚本。

## 相关文件

- [event.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\core\event.go)：事件类型定义和结构
- [event_manager.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\core\event_manager.go)：事件管理器实现
- [events.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\events.go)：Events 命令处理器
- [script_engine.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\core\script_engine.go)：脚本引擎实现
- [loader.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\loader.go)：驱动加载和通信