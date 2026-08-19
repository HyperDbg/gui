# HyperDbg UI更新机制设计文档

## 概述

HyperDbg采用事件驱动的UI更新机制，通过统一的接口管理所有调试器组件的UI页面更新。该设计确保了调试事件、停止调试、重载等情况下能够高效地刷新所有UI页面。

## 核心架构

### 1. UI组件接口

所有调试器组件都实现了统一的`api.Interface`接口：

```go
type Interface interface {
    Layout() layout.Widget
    Update() error
    Clear()
    Self() any
}
```

- `Layout()`: 返回组件的UI布局
- `Update()`: 更新组件的数据和UI显示
- `Clear()`: 清空组件的数据
- `Self()`: 返回组件的具体类型实例

### 2. 调试器组件管理

调试器核心(`debugger.Debugger`)维护所有UI组件的引用：

```go
type Debugger struct {
    breakpoints api.Interface
    memory      api.Interface
    registers   api.Interface
    symbols     api.Interface
    threads     api.Interface
    // ... 其他组件

    uiComponents []api.Interface
}
```

在`New()`函数中，所有组件被注册到`uiComponents`列表：

```go
func New() *Debugger {
    dbg := &Debugger{
        // ... 初始化各组件
    }

    dbg.RegisterUIComponent(dbg.breakpoints)
    dbg.RegisterUIComponent(dbg.memory)
    dbg.RegisterUIComponent(dbg.registers)
    // ... 注册所有组件

    return dbg
}
```

## 事件驱动的UI更新机制

### 1. 调试事件通道

调试器通过`eventChan`通道向UI发送调试事件：

```go
type Debugger struct {
    eventChan chan *windows.DebugEvent
}

func (d *Debugger) GetEventChan() <-chan *windows.DebugEvent {
    return d.eventChan
}
```

### 2. UI事件监听

UI层在启动时启动一个goroutine监听调试事件：

```go
func Run() {
    dbg := debugger.New()
    
    go func() {
        for range dbg.GetEventChan() {
            dbg.UpdateAllPages()
        }
    }()
    
    // ... UI初始化
}
```

### 3. 批量更新所有页面

`UpdateAllPages()`方法遍历所有注册的UI组件并调用它们的`Update()`方法：

```go
func (d *Debugger) UpdateAllPages() error {
    for _, component := range d.uiComponents {
        if err := component.Update(); err != nil {
            fmt.Printf("更新UI组件失败: %v\n", err)
        }
    }
    return nil
}
```

## UI更新时机

### 1. 调试事件暂停时

当调试器接收到需要暂停的调试事件时（如断点、单步异常等），自动更新所有UI页面：

```go
if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
    // 处理异常...
    autoContinue = false
}

if !autoContinue {
    d.state = StatePaused
    d.UpdateAllPages()
    
    // 等待用户操作
}
```

### 2. 进程退出时

当被调试进程退出时，清空所有UI页面：

```go
if event.DebugEventCode() == windows.EXIT_PROCESS_DEBUG_EVENT {
    d.state = StateStopped
    d.ClearAllPages()
    autoContinue = false
}
```

### 3. 附加进程时

附加到现有进程后，启动事件循环：

```go
func (d *Debugger) Attach(pid uint32) error {
    handle, err := windows.DebugActiveProcess(pid)
    if err != nil {
        return err
    }

    d.memory.Self().(*memory.Manager).SetHandle(handle)
    d.symbols.Self().(*symbol.Manager).Initialize(handle)

    d.startEventLoop()

    d.processHandle = handle
    d.processId = pid
    d.state = StatePaused

    return nil
}
```

### 4. 分离进程时

从进程分离时，清空所有UI页面：

```go
func (d *Debugger) Detach() error {
    if d.processHandle == 0 {
        return nil
    }

    d.stopEventLoop()

    err := windows.DebugActiveProcessStop(d.processId)
    if err != nil {
        return err
    }

    d.processHandle = 0
    d.processId = 0
    d.state = StateStopped

    d.ClearAllPages()

    return nil
}
```

### 5. 终止进程时

终止被调试进程时，清空所有UI页面：

```go
func (d *Debugger) TerminateProcess(exitCode uint32) error {
    if d.processHandle == 0 {
        return nil
    }

    err := windows.TerminateProcess(d.processHandle, exitCode)
    if err != nil {
        return err
    }

    windows.CloseHandle(d.processHandle)

    d.stopEventLoop()

    d.processHandle = 0
    d.processId = 0
    d.state = StateStopped

    d.ClearAllPages()

    return nil
}
```

## 特殊组件的更新逻辑

某些组件需要特殊的数据处理逻辑：

### 1. 反汇编组件

反汇编组件需要从内存读取指令数据并反汇编：

```go
if component == d.disassembly {
    if d.processHandle != 0 {
        var buffer []byte
        var rip uint64

        allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
        if len(allThreads) > 0 {
            thread := allThreads[0]
            threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, thread.Id)
            if err == nil {
                defer windows.CloseHandle(threadHandle)
                regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
                if err == nil {
                    rip = regCtx.RIP
                    buffer, err = d.memory.Self().(*memory.Manager).ReadMemory(rip, 4096)
                }
            }
        }

        if buffer != nil && rip != 0 {
            d.disassembly.Self().(*disassembly.Disassembler).PopulateTable(
                d.disassembly.Self().(*disassembly.Disassembler).GetTable(),
                d.baseAddress,
                d.entryPoint,
                d.exePath,
                buffer,
                rip,
            )
        }
    }
}
```

### 2. 堆栈组件

堆栈组件需要遍历调用栈：

```go
if component == d.stack {
    if d.processHandle != 0 {
        allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
        if len(allThreads) > 0 {
            thread := allThreads[0]
            threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, thread.Id)
            if err == nil {
                defer windows.CloseHandle(threadHandle)
                regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
                if err == nil {
                    frames, err := d.stack.Self().(*stack.Manager).WalkStack(
                        threadHandle,
                        regCtx,
                        d.memory.Self().(*memory.Manager).ReadMemory,
                        d.symbols.Self().(*symbol.Manager).GetSymbolFromAddress,
                        d.symbols.Self().(*symbol.Manager).GetModuleByAddress,
                    )
                    if err == nil {
                        table := d.stack.Self().(*stack.Manager).GetTable()
                        table.Root().SetChildren(nil)
                        for _, frame := range frames {
                            table.Root().AddChild(table.NewNode(*frame))
                        }
                    }
                }
            }
        }
    }
}
```

### 3. SEH组件

SEH组件需要扫描异常处理链：

```go
if component == d.seh {
    if d.processHandle != 0 {
        allThreads := d.threads.Self().(*thread.Manager).GetAllThreads()
        if len(allThreads) > 0 {
            thread := allThreads[0]
            threadHandle, err := windows.OpenThread(windows.THREAD_GET_CONTEXT, false, thread.Id)
            if err == nil {
                defer windows.CloseHandle(threadHandle)
                regCtx, err := d.registers.Self().(*register.Manager).GetThreadContext(threadHandle)
                if err == nil {
                    handlers, err := d.seh.Self().(*seh.Manager).ScanSEH(
                        regCtx,
                        d.memory.Self().(*memory.Manager).ReadMemory,
                    )
                    if err == nil {
                        table := d.seh.Self().(*seh.Manager).GetTable()
                        table.Root().SetChildren(nil)
                        for _, handler := range handlers {
                            table.Root().AddChild(table.NewNode(*handler))
                        }
                    }
                }
            }
        }
    }
}
```

## UI页面布局

UI采用Tab页面布局，每个调试器组件对应一个Tab页面：

```go
type TabPageType int

const (
    CpuType TabPageType = iota
    PeViewType
    LogType
    NotesType
    BreaksType
    MemoryType
    SehType
    ScriptType
    SymbolType
    SourceType
    ReferencesType
    ThreadType
    HandleType
    TraceType
    ArkType
    ScyllaType
    LabelsType
    CommentsType
    FunctionsType
    XrefsType
    TypesType
    WatchesType
    GraphsType
    ExceptionsType
    BookmarksType
    LoopsType
)
```

注意：某些组件（如寄存器、堆栈、IMM）已经集成在CPU页面中，不作为单独的Tab页面显示。

## Toolbar按钮回调函数

UI层通过toolbar按钮提供调试控制功能，每个按钮都有对应的回调函数：

### 已实现的按钮功能

```go
func toolbarButtons(m *safemap.M[string, []byte], dbg *debugger.Debugger) iter.Seq[*TipIconButton] {
    return func(yield func(*TipIconButton) bool) {
        yield(NewTooltipButton(m.GetMust("close.png"), "close", func() {
            if dbg.GetProcessHandle() != 0 {
                dbg.Detach()
            }
        }))
        yield(NewTooltipButton(m.GetMust("run.png"), "run", func() {
            dbg.Continue()
        }))
        yield(NewTooltipButton(m.GetMust("stepin.png"), "stepin", func() {
            dbg.StepInto()
        }))
        yield(NewTooltipButton(m.GetMust("stepover.png"), "stepover", func() {
            dbg.StepOver()
        }))
        // ... 其他按钮
    }
}
```

### 按钮功能说明

| 按钮 | 功能 | 调用方法 | 状态 |
|------|------|----------|------|
| open | 打开文件 | - | 待实现 |
| restart | 重新启动调试 | - | 待实现 |
| close | 关闭/分离进程 | `dbg.Detach()` | ✅ 已实现 |
| run | 继续执行 | `dbg.Continue()` | ✅ 已实现 |
| runthread | 运行线程 | - | 待实现 |
| pause | 暂停执行 | - | 待实现 |
| stepin | 单步进入 | `dbg.StepInto()` | ✅ 已实现 |
| stepover | 单步跳过 | `dbg.StepOver()` | ✅ 已实现 |
| trin | 追踪进入 | - | 待实现 |
| trover | 追踪跳过 | - | 待实现 |
| tillret | 运行到返回 | - | 待实现 |
| tilluser | 运行到用户代码 | - | 待实现 |
| log | 日志 | - | 待实现 |
| modules | 模块 | - | 待实现 |
| windows | 窗口 | - | 待实现 |
| threads | 线程 | - | 待实现 |
| cpu | CPU | - | 待实现 |
| search | 搜索 | - | 待实现 |
| trace | 追踪 | - | 待实现 |
| bpoints | 断点 | - | 待实现 |
| bpmem | 内存断点 | - | 待实现 |
| bphard | 硬件断点 | - | 待实现 |
| options | 选项 | - | 待实现 |
| scylla | Scylla | - | 待实现 |
| about | 关于 | - | 待实现 |
| settings | 设置 | - | 待实现 |

### 按钮回调函数实现

UI层在初始化时创建所有toolbar按钮并设置回调函数：

```go
func Run() {
    p := panel.New()
    hPanel := panel.NewHPanel()
    p.AddChild(hPanel)

    dbg := debugger.New()
    NewToolbar(hPanel, dbg)

    go func() {
        for range dbg.GetEventChan() {
            dbg.UpdateAllPages()
        }
    }()

    // ... UI初始化
}

func NewToolbar(hpanel *panel.Panel, dbg *debugger.Debugger) {
    m := stream.ReadEmbedFileMap(bar, "asserts/bar")
    for tipBtn := range toolbarButtons(m, dbg) {
        action := appbar.Action{
            Widget:    tipBtn.Layout,
            Update:    func(gtx layout.Context) {},
            AlignLeft: true,
        }
        myAppBar.AddAction(action)
    }
    // ... 添加其他action
}
```

## 设计优势

1. **统一接口**: 所有组件实现相同的接口，便于管理和扩展
2. **事件驱动**: 基于调试事件自动更新UI，无需手动刷新
3. **批量更新**: 一次调用更新所有页面，提高效率
4. **解耦设计**: 调试器核心与UI层通过通道通信，降低耦合
5. **灵活扩展**: 新增组件只需实现接口并注册即可

## 总结

HyperDbg的UI更新机制通过统一的接口和事件驱动的方式，实现了高效、可靠的UI页面更新。调试器核心负责处理调试事件并触发更新，UI层负责监听事件并刷新显示，两者通过通道通信，实现了良好的解耦和扩展性。
