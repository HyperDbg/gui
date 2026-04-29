# HyperDbg Debugger - Go Implementation

这是 HyperDbg 调试器的 Go 语言实现，完整复刻了 C++ 版本的 ring3 功能。

## 项目结构

```
hypedbg/
├── debugger/              # 调试器核心
│   ├── interfaces.go      # 核心接口定义（Eventer、DebuggerPublic、KernelDebugger、UserDebugger）
│   ├── packet.go         # IRP 发送层，实现 DebuggerPublic 接口
│   ├── ioctl_codes.go    # IOCTL 代码定义
│   ├── structures.go     # 数据结构定义
│   ├── event.types.go    # 事件类型定义
│   ├── event_manager.go  # 事件管理器
│   ├── event_handler.go  # 事件处理器
│   ├── event_loop.go     # 事件循环
│   ├── kernel_debugger.go # 内核模式调试器（匿名嵌入 Packet）
│   ├── user_debugger.go  # 用户模式调试器（匿名嵌入 Packet）
│   ├── steppings.go     # 单步执行
│   ├── api/             # API 接口定义
│   ├── breakpoint/       # 断点管理
│   ├── commands/        # 调试命令
│   ├── driver/          # 驱动通信
│   ├── hardware/        # 硬件检测
│   ├── memory/          # 内存操作
│   ├── register/        # 寄存器操作
│   ├── script/          # 脚本引擎
│   ├── symbol/          # 符号表
│   ├── mcp/             # AI 调试服务器（自动生成）
│   ├── cmd/             # 代码生成器
│   │   └── generate/    # MCP 服务器生成器
│   └── ...             # 其他模块
├── ui/                  # 用户界面
│   ├── pages/            # UI 页面
│   ├── asserts/          # 资源文件
│   └── ui.go            # UI 主入口
├── doc/                 # 文档
├── main.go              # 程序入口
└── go.mod              # Go 模块定义
```

## 功能特性

### 1. 核心接口 (debugger/interfaces.go)
- **Eventer 接口** - 事件操作接口
  - `StartEventLoop()` - 启动事件循环
  - `StopEventLoop()` - 停止事件循环
  - `RegisterEvent()` - 注册事件
  - `ModifyEvent()` - 修改事件
  - `RemoveEvent()` - 移除事件
  - `Events()` - 查询事件

- **DebuggerPublic 接口** - 调试器公共接口（由 Packet 实现）
  - 驱动管理（LoadDriver、UnloadDriver）
  - VMM 管理（LoadVMM、UnloadVMM）
  - 事件管理（继承 Eventer）
  - 内存操作（ReadMemory、WriteMemory）
  - 断点管理（EPTHook、HookSyscall）
  - 进程/线程管理（Processes、Threads）
  - 硬件操作（ReadMSR、WriteMSR）

- **KernelDebugger 接口** - 内核模式调试器接口
  - 继承 DebuggerPublic
  - 内核特定方法（SetBreakpoint、ReadMemory、WriteMemory）
  - CPUID、CRWrite、DR 等内核操作

- **UserDebugger 接口** - 用户模式调试器接口
  - 继承 DebuggerPublic
  - 用户特定方法（AttachProcess、KillProcess、StartProcess）
  - 断点管理、单步调试

### 2. IRP 发送层 (debugger/packet.go)
- **Packet 结构** - 实现 DebuggerPublic 接口
  - 驱动加载/卸载
  - VMM 加载/卸载
  - 所有 IRP 发送逻辑
  - 事件循环管理
  - 内存读写
  - 断点管理
  - 进程/线程操作

### 3. 调试器实现 (debugger/)
- **内核模式调试器** ([kernel_debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\kernel_debugger.go))
  - 匿名嵌入 Packet
  - 实现 KernelDebugger 接口
  - 内核模式特定方法

- **用户模式调试器** ([user_debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\user_debugger.go))
  - 匿名嵌入 Packet
  - 实现 UserDebugger 接口
  - 用户模式特定方法

### 4. 驱动管理 (debugger/driver/)
- **驱动提供者** ([provider.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\provider.go))
  - `Load()` - 加载驱动
  - `Unload()` - 卸载驱动
  - `Manage()` - 统一的驱动管理接口

- **驱动句柄** ([handle.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\handle.go))
  - Windows 设备驱动通信
  - IOCTL 接口支持
  - 缓冲区发送和接收

- **驱动配置** ([config.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\config.go))
  - 配置管理
  - 设备名称配置

### 5. 事件系统 (debugger/event_*.go)
- **事件管理器** ([event_manager.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_manager.go))
  - 事件注册和查询
  - 事件修改和删除
  - 事件动作管理

- **事件处理器** ([event_handler.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_handler.go))
  - 37 种事件类型处理
  - 事件回调机制
  - 事件分发

- **事件循环** ([event_loop.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_loop.go))
  - 事件监听
  - 事件分发
  - 异步处理

### 6. 单步调试 (debugger/steppings.go)
- **单步调试** ([steppings.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\steppings.go))
  - `StepInto()` - 单步进入
  - `StepOver()` - 单步跳过
  - `StepOut()` - 单步跳出
  - 指令跟踪
  - 用户模式和内核模式支持

### 7. 断点管理 (debugger/breakpoint/)
- **断点控制** ([control.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\breakpoint\control.go))
  - CTRL+C/CTRL+BREAK 处理
  - 暂停/继续控制
  - 内核调试命令发送

### 8. MCP 服务器 (debugger/mcp/)
- **内核模式 MCP 服务器** ([kernel_mode_mcp.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\mcp\kernel_mode_mcp.go))
  - 从 KernelDebugger 接口自动生成
  - 提供 AI 调试接口
  - 实现内核模式所有方法

- **用户模式 MCP 服务器** ([user_mode_mcp.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\mcp\user_mode_mcp.go))
  - 从 UserDebugger 接口自动生成
  - 提供 AI 调试接口
  - 实现用户模式所有方法

### 9. MCP 生成器 (debugger/cmd/generate/)
- **MCP 服务器生成器** ([mcp_generator.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\cmd\generate\mcp_generator.go))
  - 从接口定义自动生成 MCP 服务器代码
  - 解析 Go 接口 AST
  - 生成符合 MCP 协议的工具定义
  - 支持递归提取继承接口的方法
  - 使用方法：`go run debugger/cmd/generate/mcp_generator.go`

### 10. 事件循环和UI刷新 (debugger/packet.go)
- **Packet 事件循环** ([packet.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\packet.go))
  - 管理所有业务包（register、memory、breakpoint、thread、process 等）
  - 协调事件循环和 UI 更新
  - 在事件触发时自动刷新所有 UI 页面
  - 注册事件处理器，处理 37 种事件类型
  - 提供统一的 UI 组件注册和更新机制

### 11. 业务包 (debugger/ 下的各个包)
所有业务包都实现 `api.Interface` 接口，提供统一的 UI 更新机制：
- **register/**：寄存器操作和显示
- **memory/**：内存读写和显示
- **breakpoint/**：断点管理
- **thread/**：线程操作和显示
- **process/**：进程操作和显示
- **disassembly/**：反汇编
- **stack/**：堆栈跟踪
- **symbol/**：符号表
- **trace/**：跟踪
- **transparency/**：透明调试器
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

## 架构设计

### 接口继承关系

```
Eventer (事件操作接口)
    ↓
DebuggerPublic (调试器公共接口，由 Packet 实现)
    ↓
    ├── KernelDebugger (内核模式调试器接口)
    │       └─> KernelDebugger 结构体（匿名嵌入 Packet）
    │
    └── UserDebugger (用户模式调试器接口)
            └─> UserDebugger 结构体（匿名嵌入 Packet）
```

### 核心设计原则

1. **接口驱动设计**：所有调试器操作都通过接口定义，确保类型安全和可测试性
2. **匿名嵌入**：UserDebugger 和 KernelDebugger 通过匿名嵌入 Packet 继承所有公共方法
3. **自动生成**：MCP 服务器代码从接口定义自动生成，减少手动维护
4. **分层架构**：清晰的分层设计，UI → 调试器接口 → Packet → 驱动

## 使用示例

### 内核模式调试

```go
// 创建内核模式调试器
packet := debugger.NewPacket(driverHandle)
kernelDbg := debugger.NewKernelDebugger(packet)

// 加载驱动和 VMM
kernelDbg.LoadDriver("hyperkd.sys")
kernelDbg.LoadVMM()

// 设置断点
kernelDbg.SetBreakpoint(0x140001000, 1234)

// 读取内存
data, err := kernelDbg.ReadMemory(1234, 0x140001000, 0x100, debugger.MemoryTypeVirtual)
```

### 用户模式调试

```go
// 创建用户模式调试器
packet := debugger.NewPacket(driverHandle)
userDbg := debugger.NewUserDebugger(packet)

// 加载驱动和 VMM
userDbg.LoadDriver("hyperkd.sys")
userDbg.LoadVMM()

// 附加到进程
userDbg.AttachProcess(1234)

// 设置断点
userDbg.SetBreakpoint(0x140001000)

// 单步调试
userDbg.StepInto()
```

### MCP 服务器使用

```go
// 创建内核模式 MCP 服务器
kernelMCP := mcp.NewKernelModeMCPServer(kernelDbg)
kernelMCP.RegisterTools(mcpServer)

// 创建用户模式 MCP 服务器
userMCP := mcp.NewUserModeMCPServer(userDbg)
userMCP.RegisterTools(mcpServer)
```

## 开发指南

### 生成 MCP 服务器

当修改 `interfaces.go` 中的接口定义后，需要重新生成 MCP 服务器代码：

```bash
go run debugger/cmd/generate/mcp_generator.go
```

这将自动生成：
- `debugger/mcp/kernel_mode_mcp.go` - 内核模式 MCP 服务器
- `debugger/mcp/user_mode_mcp.go` - 用户模式 MCP 服务器

## 文档

- [架构设计](doc/ARCHITECTURE.md) - 详细的架构设计文档
- [内部机制分析](doc/HyperDbg内部机制分析.md) - HyperDbg 内部机制分析
- [驱动通信](doc/HYPERDBG_CPP_DRIVER_COMMUNICATION.md) - C++ 驱动通信协议

## 许可证

本项目遵循 HyperDbg 原项目的许可证。
