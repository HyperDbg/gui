# HyperDbg Go 项目架构设计

## 设计原则

**单一职责原则**：每个包只负责一个领域的功能，通过分层调用实现解耦。

**依赖注入原则**：所有业务包通过构造函数接收 driver 指针，实现松耦合。

**事件驱动架构**：通过事件机制实现模块间通信，支持异步处理。

**Provider 模式**：所有业务包实现 `api.Interface` 接口，提供统一的访问方式。

## 目录结构

```
hypedbg/
├── debugger/              # 调试器核心包
│   ├── debugger.go        # 调试器主入口（初始化、协调）
│   ├── steppings.go       # 单步执行
│   ├── packet.go          # IRP 发送层
│   ├── ioctl_codes.go     # IOCTL 代码定义
│   ├── structures.go      # 数据结构定义
│   ├── event.types.go     # 事件类型定义
│   ├── event_manager.go   # 事件管理器
│   ├── event_handler.go   # 事件处理器
│   ├── event_loop.go      # 事件循环
│   ├── user_debugger.go   # 用户态调试器
│   ├── vmm.go             # VMM 管理
│   ├── api/               # 基础接口层
│   │   └── interface.go   # 统一接口定义（包括 gio Widget 接口）
│   │
│   ├── driver/            # 驱动管理层
│   │   ├── handle.go      # 驱动句柄管理
│   │   ├── provider.go    # 驱动提供者
│   │   └── config.go      # 配置管理
│   │
│   ├── breakpoint/        # 断点管理
│   │   ├── breakpoint.go  # 断点操作
│   │   └── control.go    # 断点控制
│   │
│   ├── memory/            # 内存操作
│   │   └── memory.go     # 调用 packet 发送 IRP
│   │
│   ├── thread/            # 线程操作
│   │   └── thread.go     # 调用 packet 发送 IRP
│   │
│   ├── process/           # 进程操作
│   │   └── process.go    # 调用 packet 发送 IRP
│   │
│   ├── register/          # 寄存器操作
│   │   └── register.go   # 调用 packet 发送 IRP
│   │
│   ├── disassembly/       # 反汇编
│   │   └── disassembly.go
│   │
│   ├── stack/             # 堆栈跟踪
│   │   └── stack.go
│   │
│   ├── symbol/            # 符号表
│   │   └── symbol.go
│   │
│   ├── trace/             # 跟踪
│   │   └── trace.go
│   │
│   ├── transparency/      # 透明调试器
│   │   └── transparency.go
│   │
│   ├── hardware/          # 硬件包
│   │   ├── cpu.go        # CPU/VMX 检测
│   │   └── cpuid_amd64.s # CPUID 汇编实现
│   │
│   ├── commands/          # 命令处理
│   │   ├── debugging/    # 调试命令（断点、单步等）
│   │   ├── extension/    # 扩展命令
│   │   ├── meta/         # 元命令（连接、加载等）
│   │   └── events.go     # 事件命令
│   │
│   ├── script/            # 脚本引擎
│   │   └── script_engine.go
│   │
│   ├── mcp/               # AI 调试服务器
│   │   ├── server.go     # MCP 服务器
│   │   ├── handlers.go   # AI 处理器
│   │   └── types.go      # MCP 类型定义
│   │
│   ├── bookmark/          # 书签管理
│   ├── comment/           # 注释管理
│   ├── label/             # 标签管理
│   ├── function/          # 函数管理
│   ├── xref/              # 交叉引用
│   ├── watch/             # 监视点
│   ├── notes/             # 笔记
│   ├── source/            # 源代码
│   ├── reference/         # 引用
│   ├── handle/            # 句柄管理
│   ├── peview/            # PE 查看
│   ├── exception/         # 异常处理
│   ├── ark/               # ARK 工具
│   ├── scylla/            # Scylla 工具
│   ├── graph/             # 图形
│   ├── type_/             # 类型
│   ├── loop/              # 循环
│   ├── seh/               # SEH 处理
│   ├── remote/            # 远程调试
│   └── log/               # 日志
│       └── logger.go      # 日志初始化
│
├── ui/                   # UI 界面
│   ├── ui.go             # UI 主入口
│   └── pages/            # UI 页面
│       ├── cpu.go
│       ├── events.go
│       └── log.go
│
├── doc/                  # 文档
│   ├── ARCHITECTURE.md   # 架构文档
│   ├── README.md         # 项目说明
│   └── ...               # 其他文档
│
├── main.go               # 程序入口
└── go.mod                # Go 模块定义

```

## 核心架构设计

### 架构概览

```
┌─────────────────────────────────────────────────────────────────┐
│                        UI 层 (ui/)                             │
│  接收用户操作，触发调试器行为                                   │
└────────────────────────────┬────────────────────────────────┘
                         │ 调用
┌────────────────────────────▼────────────────────────────────┐
│              调试器主入口 (debugger/debugger.go)          │
│  - 初始化驱动                                                │
│  - 创建所有业务包（Provider 模式）                             │
│  - 创建事件系统（EventManager、EventHandler、EventLoop）        │
│  - 协调各组件                                                │
└────────────────────────────┬────────────────────────────────┘
                         │ 依赖注入
┌────────────────────────────▼────────────────────────────────┐
│              业务层 (debugger/)                           │
│  所有业务包实现 api.Interface 接口                          │
│  - memory/, thread/, process/, breakpoint/, register/        │
│  - disassembly/, stack/, symbol/, trace/                     │
│  - transparency/, hardware/, vmm/                            │
│  - bookmark/, comment/, label/, function/, xref/             │
│  - watch/, notes/, source/, reference/, handle/              │
│  - peview/, exception/, ark/, scylla/, graph/              │
│  - type_/, loop/, seh/, remote/                             │
└────────────────────────────┬────────────────────────────────┘
                         │ 调用
┌────────────────────────────▼────────────────────────────────┐
│            事件系统 (debugger/event_*.go)                │
│  - EventManager: 事件管理器                                  │
│  - EventHandler: 事件处理器                                   │
│  - EventLoop: 事件循环                                       │
│  - DebugEventHandler: 调试事件处理器（处理驱动返回）          │
└────────────────────────────┬────────────────────────────────┘
                         │ 发送
┌────────────────────────────▼────────────────────────────────┐
│            IRP 发送层 (debugger/packet.go)               │
│  - 封装所有 IRP 发送逻辑                                     │
│  - 提供统一的发送接口                                        │
│  - 处理驱动返回数据                                          │
└────────────────────────────┬────────────────────────────────┘
                         │ 调用
┌────────────────────────────▼────────────────────────────────┐
│            驱动管理层 (debugger/driver/)                 │
│  - DriverHandle: 驱动句柄管理                               │
│  - DriverProvider: 驱动提供者（加载/卸载）                    │
│  - DriverConfig: 配置管理                                    │
└────────────────────────────┬────────────────────────────────┘
                         │ 操作
┌────────────────────────────▼────────────────────────────────┐
│              Windows 驱动 (hyperkd.sys)               │
└─────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│              命令层 (debugger/commands/)                 │
│  - 解析用户命令，调用业务层方法                                 │
│  - 不直接发送 IRP，所有操作通过业务层完成                         │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│              脚本层 (debugger/script/)                    │
│  - 执行脚本，调用业务层方法                                     │
│  - 不直接发送 IRP，所有操作通过业务层完成                         │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│              AI 调试服务器 (debugger/mcp/)                │
│  - 提供 AI 调试接口，调用业务层方法                              │
│  - 不直接发送 IRP，所有操作通过业务层完成                         │
└─────────────────────────────────────────────────────────────────┘
```

### 初始化流程

```
main.go
  │
  ├─> ui.Run(func(dbg *debugger.HyperDbg) { ... })
  │       │
  │       └─> debugger.NewHyperDbg()
  │               │
  │               ├─> 1. 初始化驱动
  │               │   ├─> driver.NewDriverHandle()
  │               │   ├─> driver.NewDriverProvider()
  │               │   └─> loader.Install()
  │               │
  │               ├─> 2. 创建 Packet（IRP 发送层）
  │               │   └─> debugger.NewPacket(driverHandle)
  │               │
  │               ├─> 3. 创建 VMM Provider
  │               │   └─> VmmProviderNew(packet)
  │               │       └─> vmmProvider.Initialize()
  │               │
  │               ├─> 4. 创建所有业务包（Provider 模式）
  │               │   ├─> register.New()
  │               │   ├─> memory.New()
  │               │   ├─> breakpoint.New()
  │               │   ├─> symbol.New()
  │               │   ├─> thread.New()
  │               │   ├─> stack.New()
  │               │   ├─> disassembly.New()
  │               │   ├─> trace.New()
  │               │   ├─> process.New()
  │               │   ├─> transparency.New()
  │               │   ├─> bookmark.New()
  │               │   ├─> comment.New()
  │               │   ├─> label.New()
  │               │   ├─> function.New()
  │               │   ├─> xref.New()
  │               │   ├─> watch.New()
  │               │   ├─> notes.New()
  │               │   ├─> source.New()
  │               │   ├─> reference.New()
  │               │   ├─> handle.New()
  │               │   ├─> peview.New()
  │               │   ├─> exception.New()
  │               │   ├─> ark.New()
  │               │   ├─> scylla.New()
  │               │   ├─> graph.New()
  │               │   ├─> type_.New()
  │               │   ├─> loop.New()
  │               │   └─> seh.New()
  │               │
  │               ├─> 5. 创建事件系统
  │               │   ├─> NewEventManager(driverHandle)
  │               │   ├─> NewEventHandler(driverHandle)
  │               │   ├─> NewDebugEventHandler(driverHandle, regMgr, procMgr, threadMgr)
  │               │   └─> NewEventLoop(eventManager)
  │               │
  │               ├─> 6. 注册所有事件处理器
  │               │   └─> eventLoop.RegisterHandler(...)
  │               │
  │               └─> 7. 启动事件循环
  │                   └─> eventLoop.Start()
  │
  └─> 启动 UI
```

## 各层职责

### debugger/debugger.go - 调试器主入口
- **职责**：
  - 初始化驱动（加载、卸载、句柄管理）
  - 创建所有业务包（Provider 模式）
  - 创建事件系统（EventManager、EventHandler、EventLoop）
  - 协调各组件
  - 管理调试器状态
- **依赖**：所有业务包、driver/、packet、事件系统

### debugger/packet.go - IRP 发送层
- **职责**：
  - 封装所有 IRP 发送逻辑
  - 提供统一的发送接口（SendBuffer、ReceiveBuffer）
  - 处理驱动返回数据
  - 定义 IOCTL 代码
- **依赖**：debugger/driver/

### debugger/driver/ - 驱动管理层
- **职责**：
  - 管理驱动句柄（DriverHandle）
  - 驱动加载/卸载（DriverProvider）
  - 配置管理（DriverConfig）
- **依赖**：Windows API

### debugger/event_manager.go - 事件管理器
- **职责**：
  - 管理事件注册
  - 触发事件
  - 协调事件处理
- **依赖**：debugger/driver/

### debugger/event_handler.go - 事件处理器
- **职责**：
  - 处理基本事件（断点、单步、继续等）
  - 发送 IRP 到驱动
  - 处理驱动返回
- **依赖**：debugger/driver/

### debugger/event_loop.go - 事件循环
- **职责**：
  - 监听事件通道
  - 分发事件到注册的处理器
  - 异步处理事件
- **依赖**：debugger/event_manager.go

### debugger/event_handler.go - DebugEventHandler
- **职责**：
  - 处理调试事件（37 种事件类型）
  - 解码驱动返回数据
  - 更新寄存器、进程、线程等状态
- **依赖**：debugger/driver/、register/、process/、thread/

### debugger/api/interface.go - 基础接口层
- **职责**：
  - 定义统一接口（api.Interface）
  - 包含 gio Widget 接口
  - 所有业务包都实现此接口
- **依赖**：Gio UI 框架

### debugger/ 下的业务包
所有业务包遵循相同模式：
- **职责**：封装业务逻辑，通过 packet 发送 IRP
- **依赖**：debugger/packet.go
- **设计模式**：Provider 模式、实现 api.Interface

#### 核心业务包
- **register/**：寄存器操作
- **memory/**：内存读写
- **breakpoint/**：断点管理
- **thread/**：线程操作
- **process/**：进程操作
- **disassembly/**：反汇编
- **stack/**：堆栈跟踪
- **symbol/**：符号表
- **trace/**：跟踪
- **transparency/**：透明调试器
- **hardware/**：硬件操作
- **vmm/**：VMM 管理

#### 辅助业务包
- **bookmark/**：书签管理
- **comment/**：注释管理
- **label/**：标签管理
- **function/**：函数管理
- **xref/**：交叉引用
- **watch/**：监视点
- **notes/**：笔记
- **source/**：源代码
- **reference/**：引用
- **handle/**：句柄管理
- **peview/**：PE 查看
- **exception/**：异常处理
- **ark/**：ARK 工具
- **scylla/**：Scylla 工具
- **graph/**：图形
- **type_/**：类型
- **loop/**：循环
- **seh/**：SEH 处理
- **remote/**：远程调试

### debugger/commands/ - 命令处理层
- **职责**：
  - 解析用户命令
  - 调用业务层方法
  - 不直接发送 IRP
- **依赖**：debugger/ 下所有业务包
- **设计原则**：只调用业务层方法，不绕过业务层直接发包

### debugger/script/ - 脚本引擎层
- **职责**：
  - 执行脚本
  - 调用业务层方法
  - 不直接发送 IRP
- **依赖**：debugger/ 下所有业务包
- **设计原则**：只调用业务层方法，不绕过业务层直接发包

### debugger/mcp/ - AI 调试服务器
- **职责**：
  - 提供 AI 调试接口
  - 调用业务层方法
  - 不直接发送 IRP
- **依赖**：debugger/ 下所有业务包
- **设计原则**：只调用业务层方法，不绕过业务层直接发包

### ui/ - UI 层
- **职责**：
  - 接收用户操作
  - 触发调试器行为
  - 渲染 UI 界面
- **依赖**：debugger/、Gio UI 框架

### main.go - 程序入口
- **职责**：
  - 启动 UI
  - 传递回调函数
- **依赖**：ui/、debugger/

## 依赖关系图

```
main.go
  │
  └─> ui.Run(func(dbg *debugger.HyperDbg) { ... })
          │
          └─> debugger.NewHyperDbg()
                  │
                  ├─> 1. 初始化驱动
                  │   ├─> driver.NewDriverHandle()
                  │   └─> driver.NewDriverProvider()
                  │
                  ├─> 2. 创建 Packet
                  │   └─> debugger.NewPacket(driverHandle)
                  │
                  ├─> 3. 创建 VMM Provider
                  │   └─> VmmProviderNew(packet)
                  │
                  ├─> 4. 创建所有业务包
                  │   ├─> register.New()
                  │   ├─> memory.New()
                  │   ├─> breakpoint.New()
                  │   ├─> symbol.New()
                  │   ├─> thread.New()
                  │   ├─> stack.New()
                  │   ├─> disassembly.New()
                  │   ├─> trace.New()
                  │   ├─> process.New()
                  │   ├─> transparency.New()
                  │   ├─> bookmark.New()
                  │   ├─> comment.New()
                  │   ├─> label.New()
                  │   ├─> function.New()
                  │   ├─> xref.New()
                  │   ├─> watch.New()
                  │   ├─> notes.New()
                  │   ├─> source.New()
                  │   ├─> reference.New()
                  │   ├─> handle.New()
                  │   ├─> peview.New()
                  │   ├─> exception.New()
                  │   ├─> ark.New()
                  │   ├─> scylla.New()
                  │   ├─> graph.New()
                  │   ├─> type_.New()
                  │   ├─> loop.New()
                  │   └─> seh.New()
                  │
                  ├─> 5. 创建事件系统
                  │   ├─> NewEventManager(driverHandle)
                  │   ├─> NewEventHandler(driverHandle)
                  │   ├─> NewDebugEventHandler(driverHandle, regMgr, procMgr, threadMgr)
                  │   └─> NewEventLoop(eventManager)
                  │
                  ├─> 6. 注册所有事件处理器
                  │   └─> eventLoop.RegisterHandler(...)
                  │
                  └─> 7. 启动事件循环
                      └─> eventLoop.Start()

业务包依赖关系：
  所有业务包 → debugger/packet.go → debugger/driver/ → Windows 驱动

命令层依赖关系：
  commands/ → 所有业务包

脚本层依赖关系：
  script/ → 所有业务包

AI 调试服务器依赖关系：
  mcp/ → 所有业务包

UI 层依赖关系：
  ui/ → debugger/
```

## 调用示例

### 1. main.go 启动流程

```go
func main() {
    driverOnly := flag.Bool("driver-only", false, "Only load driver, don't start UI")
    flag.Parse()

    ui.SetEnableConsoleLog(true)

    if *driverOnly {
        ui.RunDriverOnly()
    } else {
        ui.Run(func(dbg *debugger.HyperDbg) {
            // UI 启动后的回调
        })
    }
}
```

### 2. debugger.NewHyperDbg() 初始化流程

```go
func NewHyperDbg() *HyperDbg {
    // 1. 初始化驱动
    driverConfig := &driver.DriverConfig{
        DeviceName:    "\\\\.\\HyperDbgDebuggerDevice",
        ServiceName:   "HyperDbg",
        DriverFileName: "hyperkd.sys",
    }

    loader := driver.NewDriverProvider("HyperDbg", driverPath)
    driverHandle, err := driver.NewDriverHandle(driverConfig)

    // 2. 创建 Packet
    packet := NewPacket(driverHandle)

    // 3. 创建 VMM Provider
    vmmProvider := VmmProviderNew(packet)
    vmmProvider.Initialize()

    // 4. 创建所有业务包
    regMgr := register.New()
    memMgr := memory.New()
    bpMgr := breakpoint.New()
    symMgr := symbol.New()
    threadMgr := thread.New()
    stackMgr := stack.New()
    disasmMgr := disassembly.New()
    traceMgr := trace.New()
    processMgr := process.New()
    transparencyMgr := transparency.New()
    bookmarkMgr := bookmark.New()
    commentMgr := comment.New()
    labelMgr := label.New()
    functionMgr := function.New()
    xrefMgr := xref.New()
    watchMgr := watch.New()
    notesMgr := notes.New()
    sourceMgr := source.New()
    referenceMgr := reference.New()
    handleMgr := handle.New()
    peviewMgr := peview.New()
    exceptionMgr := exception.New()
    arkMgr := ark.New()
    scyllaMgr := scylla.New()
    graphMgr := graph.New()
    typeMgr := type_.New()
    loopMgr := loop.New()
    sehMgr := seh.New()

    // 5. 创建事件系统
    eventManager := NewEventManager(driverHandle)
    eventHandler := NewEventHandler(driverHandle)
    debugEventHandler := NewDebugEventHandler(driverHandle, regMgr.(*register.Provider), processMgr.(*process.Provider), threadMgr.(*thread.Provider))
    eventLoop := NewEventLoop(eventManager)

    // 6. 注册所有事件处理器
    eventLoop.RegisterHandler(EventStep, func(event *DebugEvent) {
        debugEventHandler.Handle(event)
    })
    eventLoop.RegisterHandler(EventBreakpoint, func(event *DebugEvent) {
        debugEventHandler.Handle(event)
    })
    // ... 注册其他事件

    // 7. 启动事件循环
    eventLoop.Start()

    // 8. 创建调试器实例
    dbg := &HyperDbg{
        Driver:              driverHandle,
        DriverLoader:        loader,
        EventManager:        eventManager,
        EventHandler:        eventHandler,
        DebugEventHandler:   debugEventHandler,
        EventLoop:           eventLoop,
        Registers:           regMgr,
        Memory:              memMgr,
        Breakpoints:         bpMgr,
        Symbols:             symMgr,
        Threads:             threadMgr,
        Stack:               stackMgr,
        Disassembly:         disasmMgr,
        Trace:               traceMgr,
        Process:             processMgr,
        Transparency:        transparencyMgr,
        Bookmarks:           bookmarkMgr,
        Comments:            commentMgr,
        Labels:              labelMgr,
        Functions:           functionMgr,
        Xrefs:               xrefMgr,
        Types:               typeMgr,
        Watches:             watchMgr,
        Graphs:              graphMgr,
        Exceptions:          exceptionMgr,
        Notes:               notesMgr,
        Source:              sourceMgr,
        References:          referenceMgr,
        Handles:             handleMgr,
        Ark:                 arkMgr,
        Scylla:              scyllaMgr,
        Seh:                 sehMgr,
        Loops:               loopMgr,
        PeView:              peviewMgr,
    }

    return dbg
}
```

### 3. 业务包实现（Provider 模式）

```go
// debugger/memory/memory.go
type Provider struct {
    packet *Packet
}

func New() *Provider {
    return &Provider{}
}

func (p *Provider) SetPacket(packet *Packet) {
    p.packet = packet
}

func (p *Provider) ReadMemory(pid uint32, address uint64, size uint32) ([]byte, error) {
    req := &DebuggerReadMemory{
        Pid:        pid,
        Address:    address,
        Size:       size,
        MemoryType: DEBUGGER_READ_MEMORY_TYPE_READ_FROM_MEMORY,
    }

    buffer := p.constructReadMemoryBuffer(req)

    err := p.packet.SendBuffer(buffer, IoctlDebuggerReadMemory)
    if err != nil {
        return nil, err
    }

    response, err := p.packet.ReceiveBuffer()
    if err != nil {
        return nil, err
    }

    return p.decodeReadMemoryResponse(response), nil
}

func (p *Provider) Layout() layout.Widget {
    return func(gtx layout.Context) layout.Dimensions {
        // UI 渲染逻辑
        return layout.Dimensions{Size: gtx.Constraints.Min}
    }
}

func (p *Provider) Update() error {
    // 更新逻辑
    return nil
}

func (p *Provider) Clear() {
    // 清理逻辑
}

func (p *Provider) Self() any {
    return p
}
```

### 4. 事件处理器实现

```go
// debugger/event_handler.go
type EventHandler struct {
    driver *driver.DriverHandle
}

func NewEventHandler(driverHandle *driver.DriverHandle) *EventHandler {
    return &EventHandler{
        driver: driverHandle,
    }
}

func (h *EventHandler) HandleSetBreakpoint(event *Event) error {
    if h.driver == nil || !h.driver.IsConnected() {
        return fmt.Errorf("driver not available")
    }

    buffer := h.constructBreakpointBuffer(event)
    err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
    if err != nil {
        slog.Error("Handler: Set breakpoint failed", "error", err)
        return err
    }

    slog.Info("Handler: Set breakpoint success", "pid", event.Pid, "address", event.Address)
    return nil
}
```

### 5. 调试事件处理器实现

```go
// debugger/event_handler.go
type DebugEventHandler struct {
    driver     *driver.DriverHandle
    regMgr     *register.Provider
    procMgr    *process.Provider
    threadMgr  *thread.Provider
}

func NewDebugEventHandler(driverHandle *driver.DriverHandle, regMgr *register.Provider, procMgr *process.Provider, threadMgr *thread.Provider) *DebugEventHandler {
    return &DebugEventHandler{
        driver:    driverHandle,
        regMgr:    regMgr,
        procMgr:   procMgr,
        threadMgr: threadMgr,
    }
}

func (h *DebugEventHandler) HandleStep(event *DebugEvent) error {
    if h.driver == nil || !h.driver.IsConnected() {
        return fmt.Errorf("driver not available")
    }

    buffer := h.constructDebugEventBuffer(event)

    err := h.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)
    if err != nil {
        return fmt.Errorf("failed to send step hook request: %v", err)
    }

    response, err := h.driver.ReceiveBuffer()
    if err != nil {
        return fmt.Errorf("failed to receive step hook response: %v", err)
    }

    h.decodeStepResponse(response)
    return nil
}

func (h *DebugEventHandler) decodeStepResponse(buffer []byte) {
    if len(buffer) < 48 {
        return
    }

    pausedPacket := &DebuggeeKdPausedPacket{
        Rip:                  binary.LittleEndian.Uint64(buffer[0:8]),
        IsProcessorOn32BitMode: buffer[8] != 0,
        IgnoreDisassembling:    buffer[9] != 0,
        PausingReason:         binary.LittleEndian.Uint32(buffer[12:16]),
        CurrentCore:           binary.LittleEndian.Uint32(buffer[16:20]),
        EventTag:              binary.LittleEndian.Uint64(buffer[20:28]),
        EventCallingStage:     binary.LittleEndian.Uint32(buffer[28:32]),
        Rflags:                binary.LittleEndian.Uint64(buffer[32:40]),
    }

    if len(buffer) >= 48 {
        copy(pausedPacket.InstructionBytesOnRip[:], buffer[40:55])
        pausedPacket.ReadInstructionLen = binary.LittleEndian.Uint16(buffer[55:57])
    }

    if h.regMgr != nil {
        regCtx := h.regMgr.GetRegisterContext()
        regCtx.RIP = pausedPacket.Rip
        regCtx.RFLAGS = uint32(pausedPacket.Rflags)
        h.regMgr.SetRegisterContext(regCtx)
    }
}
```

### 6. 事件循环实现

```go
// debugger/event_loop.go
type EventLoop struct {
    eventManager *EventManager
    handlers     map[EventType]func(*DebugEvent)
    stopChan     chan struct{}
}

func NewEventLoop(eventManager *EventManager) *EventLoop {
    return &EventLoop{
        eventManager: eventManager,
        handlers:     make(map[EventType]func(*DebugEvent)),
        stopChan:     make(chan struct{}),
    }
}

func (el *EventLoop) RegisterHandler(eventType EventType, handler func(*DebugEvent)) {
    el.handlers[eventType] = handler
}

func (el *EventLoop) Start() {
    go func() {
        for {
            select {
            case event := <-el.eventManager.EventChan:
                if handler, ok := el.handlers[event.Type]; ok {
                    handler(event)
                }
            case <-el.stopChan:
                return
            }
        }
    }()
}

func (el *EventLoop) Stop() {
    close(el.stopChan)
}
```

### 7. commands/ 调用业务层方法

```go
// debugger/commands/debugging/breakpoint.go
func HandleSetBreakpoint(dbg *debugger.HyperDbg, args []string) error {
    if len(args) < 2 {
        return fmt.Errorf("usage: bp <pid> <address>")
    }

    pid, err := strconv.ParseUint(args[0], 10, 32)
    if err != nil {
        return fmt.Errorf("invalid pid: %v", err)
    }

    address, err := strconv.ParseUint(args[1], 16, 64)
    if err != nil {
        return fmt.Errorf("invalid address: %v", err)
    }

    // 调用业务层方法（不直接发送 IRP）
    bpMgr := dbg.Breakpoints.(*breakpoint.Provider)
    err = bpMgr.SetBreakpoint(uint32(pid), address)
    if err != nil {
        return fmt.Errorf("failed to set breakpoint: %v", err)
    }

    fmt.Printf("Breakpoint set at 0x%x for process %d\n", address, pid)
    return nil
}
```

### 8. script/ 调用业务层方法

```go
// debugger/script/script_engine.go
func (e *ScriptEngine) ExecuteScript(script string) error {
    stmts, err := parseScript(script)
    if err != nil {
        return err
    }

    for _, stmt := range stmts {
        switch stmt.Type {
        case "read_memory":
            memMgr := e.dbg.Memory.(*memory.Provider)
            data, err := memMgr.ReadMemory(stmt.Pid, stmt.Address, stmt.Size)
            if err != nil {
                return err
            }
            fmt.Printf("Memory: %x\n", data)

        case "list_threads":
            threadMgr := e.dbg.Threads.(*thread.Provider)
            threads, err := threadMgr.ListThreads(stmt.Pid)
            if err != nil {
                return err
            }
            fmt.Printf("Threads: %v\n", threads)

        case "attach_process":
            procMgr := e.dbg.Process.(*process.Provider)
            err := procMgr.AttachProcess(stmt.Pid)
            if err != nil {
                return err
            }
        }
    }

    return nil
}
```

### 9. mcp/ AI 调试服务器调用业务层方法

```go
// debugger/mcp/handlers.go
func (h *MCPHandlers) HandleReadMemory(request *MCPRequest) (*MCPResponse, error) {
    pid := request.Params["pid"].(uint32)
    address := request.Params["address"].(uint64)
    size := request.Params["size"].(uint32)

    // 调用业务层方法（不直接发送 IRP）
    memMgr := h.dbg.Memory.(*memory.Provider)
    data, err := memMgr.ReadMemory(pid, address, size)
    if err != nil {
        return nil, err
    }

    return &MCPResponse{
        Data: data,
    }, nil
}
```

## 事件系统

### 事件类型

HyperDbg 支持 37 种调试事件类型：

1. **基本事件**
   - EventStep: 单步执行
   - EventBreakpoint: 断点命中
   - EventException: 异常发生
   - EventContinue: 继续执行

2. **进程事件**
   - EventProcessCreated: 进程创建
   - EventProcessExited: 进程退出

3. **线程事件**
   - EventThreadCreated: 线程创建
   - EventThreadExited: 线程退出

4. **模块事件**
   - EventModuleLoaded: 模块加载
   - EventModuleUnloaded: 模块卸载

5. **EPT Hook 事件**
   - EptHookReadWriteAndExecute: 读写执行
   - EptHookReadWrite: 读写
   - EptHookReadAndExecute: 读执行
   - EptHookWriteAndExecute: 写执行
   - EptHookRead: 读
   - EptHookWrite: 写
   - EptHookExecute: 执行
   - EptHookExecDetours: 执行 Detours
   - EptHookExecCC: 执行 CC

6. **系统调用 Hook 事件**
   - SyscallHookEferSyscall: SYSCALL
   - SyscallHookEferSysret: SYSRET

7. **指令执行事件**
   - CpuidInstructionExecution: CPUID
   - RdmsrInstructionExecution: RDMSR
   - WrmsrInstructionExecution: WRMSR
   - InInstructionExecution: IN
   - OutInstructionExecution: OUT
   - TscInstructionExecution: TSC
   - PmcInstructionExecution: PMC
   - VmcallInstructionExecution: VMCALL
   - XsetbvInstructionExecution: XSETBV

8. **控制寄存器事件**
   - ControlRegisterModified: 控制寄存器修改
   - ControlRegisterRead: 控制寄存器读取
   - ControlRegister3Modified: CR3 修改

9. **其他事件**
   - ExceptionOccurred: 异常发生
   - ExternalInterruptOccurred: 外部中断
   - DebugRegistersAccessed: 调试寄存器访问
   - TrapExecutionModeChanged: Trap 执行模式改变
   - TrapExecutionInstructionTrace: Trap 指令跟踪

### 事件处理流程

```
1. 事件触发
   ↓
2. EventManager 接收事件
   ↓
3. EventLoop 监听事件通道
   ↓
4. EventLoop 分发事件到注册的处理器
   ↓
5. DebugEventHandler 处理事件
   ↓
6. 发送 IRP 到驱动
   ↓
7. 接收驱动返回
   ↓
8. 解码返回数据
   ↓
9. 更新状态（寄存器、进程、线程等）
   ↓
10. 触发回调函数
```

### 事件响应解码

不同事件类型返回不同的数据结构：

1. **DebuggeeKdPausedPacket**（单步、断点）
   - RIP、RFLAGS、指令字节等

2. **DebuggeeDetailsAndSwitchProcessPacket**（进程创建/退出）
   - 进程 ID、进程名称、进程地址等

3. **DebuggeeDetailsAndSwitchThreadPacket**（线程创建/退出）
   - 线程 ID、进程 ID、线程地址等

4. **EventAndActionResult**（大多数事件）
   - 是否成功、错误代码

5. **RegisterContext**（寄存器事件）
   - 所有寄存器值

## 设计模式

### 1. Provider 模式

所有业务包都实现 Provider 模式：

```go
type Provider struct {
    packet *Packet
}

func New() *Provider {
    return &Provider{}
}

func (p *Provider) SetPacket(packet *Packet) {
    p.packet = packet
}
```

### 2. 依赖注入

通过构造函数注入依赖：

```go
func NewDebugEventHandler(driverHandle *driver.DriverHandle, regMgr *register.Provider, procMgr *process.Provider, threadMgr *thread.Provider) *DebugEventHandler {
    return &DebugEventHandler{
        driver:    driverHandle,
        regMgr:    regMgr,
        procMgr:   procMgr,
        threadMgr: threadMgr,
    }
}
```

### 3. 事件驱动

通过事件机制实现模块间通信：

```go
eventLoop.RegisterHandler(EventStep, func(event *DebugEvent) {
    debugEventHandler.Handle(event)
})
```

### 4. 接口统一

所有业务包实现 `api.Interface` 接口：

```go
type Interface interface {
    Layout() layout.Widget
    Update() error
    Clear()
    Self() any
}
```

## UI 更新机制

### 概述

HyperDbg 采用事件驱动的 UI 更新机制，通过统一的接口管理所有调试器组件的 UI 页面更新。该设计确保了调试事件、停止调试、重载等情况下能够高效地刷新所有 UI 页面。

### UI 组件接口

所有调试器组件都实现了统一的 `api.Interface` 接口：

```go
type Interface interface {
    Layout() layout.Widget
    Update() error
    Clear()
    Self() any
}
```

- `Layout()`: 返回组件的 UI 布局
- `Update()`: 更新组件的数据和 UI 显示
- `Clear()`: 清空组件的数据
- `Self()`: 返回组件的具体类型实例

### 调试器组件管理

调试器核心 (`debugger.HyperDbg`) 维护所有 UI 组件的引用：

```go
type HyperDbg struct {
    Registers    api.Interface
    Memory       api.Interface
    Breakpoints  api.Interface
    Symbols      api.Interface
    Threads      api.Interface
    Stack        api.Interface
    Disassembly  api.Interface
    Trace        api.Interface
    Process      api.Interface
    Transparency api.Interface
    // ... 其他组件

    UIComponents []api.Interface
}
```

在 `NewHyperDbg()` 函数中，所有组件被注册到 `UIComponents` 列表：

```go
func NewHyperDbg() *HyperDbg {
    dbg := &HyperDbg{
        // ... 初始化各组件
    }

    dbg.RegisterUIComponent(regMgr)
    dbg.RegisterUIComponent(memMgr)
    dbg.RegisterUIComponent(bpMgr)
    // ... 注册所有组件

    return dbg
}
```

### 事件驱动的 UI 更新机制

#### 1. 调试事件通道

调试器通过 `EventChan` 通道向 UI 发送调试事件：

```go
type HyperDbg struct {
    EventChan chan *DebugEvent
}

func (h *HyperDbg) GetEventChan() <-chan *DebugEvent {
    return h.EventChan
}
```

#### 2. UI 事件监听

UI 层在启动时启动一个 goroutine 监听调试事件：

```go
func Run(callback func(dbg *debugger.HyperDbg)) {
    dbg := debugger.NewHyperDbg()
    
    go func() {
        for range dbg.GetEventChan() {
            dbg.UpdateAllPages()
        }
    }()
    
    callback(dbg)
    
    // ... UI 初始化
}
```

#### 3. 批量更新所有页面

`UpdateAllPages()` 方法遍历所有注册的 UI 组件并调用它们的 `Update()` 方法：

```go
func (h *HyperDbg) UpdateAllPages() error {
    for _, component := range h.UIComponents {
        if err := component.Update(); err != nil {
            slog.Error("更新UI组件失败", "error", err)
        }
    }
    return nil
}
```

### UI 更新时机

#### 1. 调试事件暂停时

当调试器接收到需要暂停的调试事件时（如断点、单步异常等），自动更新所有 UI 页面：

```go
if event.Type == EventBreakpoint || event.Type == EventStep {
    // 处理事件...
    h.SetState(StatePaused)
    h.UpdateAllPages()
}
```

#### 2. 进程退出时

当被调试进程退出时，清空所有 UI 页面：

```go
if event.Type == EventProcessExited {
    h.SetState(StateTerminated)
    h.ClearAllPages()
}
```

#### 3. 附加进程时

附加到现有进程后，启动事件循环：

```go
func (h *HyperDbg) AttachProcess(pid uint32) error {
    // 附加进程逻辑...
    
    h.SetState(StatePaused)
    h.UpdateAllPages()
    
    return nil
}
```

#### 4. 分离进程时

从进程分离时，清空所有 UI 页面：

```go
func (h *HyperDbg) DetachProcess() error {
    // 分离进程逻辑...
    
    h.SetState(StateInitialized)
    h.ClearAllPages()
    
    return nil
}
```

### UI 页面布局

UI 采用 Tab 页面布局，每个调试器组件对应一个 Tab 页面：

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

注意：某些组件（如寄存器、堆栈）已经集成在 CPU 页面中，不作为单独的 Tab 页面显示。

### Toolbar 按钮回调函数

UI 层通过 toolbar 按钮提供调试控制功能，每个按钮都有对应的回调函数：

#### 已实现的按钮功能

| 按钮 | 功能 | 调用方法 | 状态 |
|------|------|----------|------|
| close | 关闭/分离进程 | `dbg.DetachProcess()` | ✅ 已实现 |
| run | 继续执行 | `dbg.Continue()` | ✅ 已实现 |
| stepin | 单步进入 | `dbg.StepInto()` | ✅ 已实现 |
| stepover | 单步跳过 | `dbg.StepOver()` | ✅ 已实现 |
| 其他 | 其他功能 | - | 待实现 |

#### 按钮回调函数实现

UI 层在初始化时创建所有 toolbar 按钮并设置回调函数：

```go
func NewToolbar(dbg *debugger.HyperDbg) {
    // 创建按钮并设置回调
    closeButton := NewTooltipButton("close.png", "close", func() {
        if dbg.GetProcessHandle() != 0 {
            dbg.DetachProcess()
        }
    })
    
    runButton := NewTooltipButton("run.png", "run", func() {
        dbg.Continue()
    })
    
    stepInButton := NewTooltipButton("stepin.png", "stepin", func() {
        dbg.StepInto()
    })
    
    stepOverButton := NewTooltipButton("stepover.png", "stepover", func() {
        dbg.StepOver()
    })
    
    // ... 其他按钮
}
```

### 设计优势

1. **统一接口**: 所有组件实现相同的接口，便于管理和扩展
2. **事件驱动**: 基于调试事件自动更新 UI，无需手动刷新
3. **批量更新**: 一次调用更新所有页面，提高效率
4. **解耦设计**: 调试器核心与 UI 层通过通道通信，降低耦合
5. **灵活扩展**: 新增组件只需实现接口并注册即可

## 扩展指南

### 添加新的业务包

1. 在 `debugger/` 下创建新目录
2. 实现 `api.Interface` 接口
3. 在 `debugger.go` 中初始化
4. 在 UI 中注册

### 添加新的事件类型

1. 在 `event.types.go` 中定义事件类型
2. 在 `event_handler.go` 中实现处理函数
3. 在 `debugger.go` 中注册事件处理器
4. 在 `event_loop.go` 中注册处理器

### 添加新的命令

1. 在 `commands/` 下创建新文件
2. 实现命令处理函数
3. 在命令解释器中注册

## 总结

HyperDbg Go 项目采用分层架构，通过依赖注入、事件驱动和 Provider 模式实现松耦合。所有业务包都实现统一的接口，通过 IRP 发送层与驱动通信。事件系统支持 37 种调试事件类型，提供完整的调试功能。UI 更新机制采用事件驱动的方式，实现了高效、可靠的 UI 页面更新。
