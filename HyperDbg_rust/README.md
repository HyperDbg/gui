# Go Communicator

## 概述

Go Communicator 是 HyperDbg Rust 驱动的用户空间通信层，负责与 Rust 内核驱动进行网络通信，并通过 MCP (Model Context Protocol) 暴露调试器功能。

## 架构

```
┌─────────────────────────────────────────────────────────────┐
│                   MCP 客户端                              │
│              (网络调试、远程调用)                        │
└────────────────────┬────────────────────────────────────┘
                     │ MCP 协议
┌────────────────────▼────────────────────────────────────┐
│              Go Communicator                            │
├──────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐               │
│  │   Debugger   │  │   MCP Server │               │
│  │  (调试器接口) │  │  (MCP 服务)  │               │
│  └──────┬───────┘  └──────┬───────┘               │
│         │                  │                          │
│  ┌──────▼──────────────────▼──────┐               │
│  │        Protocol               │               │
│  │   (消息协议定义)              │               │
│  └──────────────────────────────┘               │
│  ┌──────────────────────────────┐               │
│  │          Server             │               │
│  │   (网络服务器)              │               │
│  └──────────────────────────────┘               │
└────────────────────┬────────────────────────────────────┘
                     │ WSK Socket
┌────────────────────▼────────────────────────────────────┐
│              Rust 内核驱动                            │
│           (hyperhv + hyperkd)                         │
└──────────────────────────────────────────────────────────┘
```

## 目录结构

```
go-communicator/
├── cmd/
│   └── generate/
│       └── mcp_generator.go      # MCP 生成器
├── debugger/
│   ├── debugger.go               # 调试器实现
│   └── interfaces.go            # 调试器接口定义
├── mcp/
│   ├── communicator_mcp.go       # MCP 服务器（自动生成）
│   └── helper.go                 # 辅助函数（自动生成）
├── protocol/
│   └── protocol.go              # 协议定义
├── server/
│   └── server.go                # 网络服务器
├── cmd/
│   └── main.go                  # 主程序
└── go.mod                       # Go 模块定义
```

## 核心组件

### 1. Debugger 接口

`debugger/interfaces.go` 定义了调试器的核心接口：

```go
type Debugger interface {
    Start(port int) error
    Stop()
    WaitForDriver(timeout time.Duration) error
    Initialize() error
    Terminate() error
    AttachProcess(processID uint32) error
    DetachProcess() error
    SetBreakpoint(address uint64, bpType BreakpointType) error
    RemoveBreakpoint(breakpointID uint64) error
    Continue() error
    Pause() error
    StepInto() error
    StepOver() error
    StepOut() error
    ReadMemory(address uint64, size uint32) ([]byte, error)
    WriteMemory(address uint64, data []byte) error
    ReadRegisters() (*RegisterState, error)
    WriteRegisters(regs *RegisterState) error
    GetState() DebugState
    WaitForEvent(timeout time.Duration) *Message
    ExecuteScript(script string) (string, error)
    ExecuteScriptWithContext(ctx context.Context, script string) (string, error)
}
```

### 2. 脚本引擎

`debugger/script.go` 实现了完整的脚本引擎，支持以下功能：

#### 脚本命令

| 命令 | 说明 | 示例 |
|------|------|------|
| **bp** | 设置断点 | `bp 0x140001000` |
| **bc** | 清除断点 | `bc 1` |
| **bl** | 列出断点 | `bl` |
| **g** | 继续执行 | `g` |
| **p** | 单步跳过 | `p` |
| **t** | 单步进入 | `t` |
| **gu** | 单步跳出 | `gu` |
| **r** | 显示/设置寄存器 | `r` 或 `r rax 0x1234` |
| **db/dw/dd/dq** | 转储内存 | `db 0x140001000 0x100` |
| **eb/ew/ed/eq** | 编辑内存 | `eb 0x140001000 0x90` |
| **k/kv** | 显示调用栈 | `k` 或 `kv` |
| **lm** | 列出模块 | `lm` |
| **process** | 显示/附加进程 | `process 1234` |
| **thread** | 显示/设置线程 | `thread 5678` |
| **attach** | 附加到进程 | `attach 1234` |
| **detach** | 分离进程 | `detach` |
| **eval** | 计算表达式 | `eval 0x140001000 + 0x10` |
| **set** | 设置变量 | `set addr 0x140001000` |
| **get** | 获取变量 | `get addr` |
| **clear** | 清除变量 | `clear addr` |
| **sleep** | 睡眠 | `sleep 1000` |
| **loop** | 循环执行 | `loop 10 bp $addr` |
| **if** | 条件执行 | `if $count > 0 echo "Continue"` |
| **echo** | 输出文本 | `echo "Hello"` |
| **log** | 记录日志 | `log info "Message"` |

#### 脚本特性

1. **变量支持**
   ```go
   set addr 0x140001000
   bp $addr
   db $addr 0x100
   ```

2. **循环执行**
   ```go
   loop 10 echo "Iteration"
   ```

3. **条件执行**
   ```go
   if $count > 0 echo "Continue"
   ```

4. **注释**
   ```go
   # This is a comment
   // This is also a comment
   bp 0x140001000  # Set breakpoint
   ```

5. **多行脚本**
   ```go
   set target 0x140001000
   bp $target
   g
   sleep 1000
   bc 1
   ```

#### 使用脚本引擎

```go
package main

import (
    "hyperdbg-communicator/debugger"
)

func main() {
    dbg := debugger.NewDebugger()
    
    // 执行单行脚本
    result, err := dbg.ExecuteScript("bp 0x140001000")
    if err != nil {
        panic(err)
    }
    println(result)
    
    // 执行多行脚本
    script := `
        set target 0x140001000
        bp $target
        g
        sleep 1000
        bc 1
    `
    result, err = dbg.ExecuteScript(script)
    if err != nil {
        panic(err)
    }
    println(result)
}
```

### 3. MCP 生成器

`cmd/generate/mcp_generator.go` 自动从 Go 接口生成 MCP 服务器代码：

- 解析 `debugger/interfaces.go` 中的接口定义
- 生成 MCP 工具定义
- 生成工具处理函数
- 支持类型转换和错误处理

**使用方法：**

```bash
cd go-communicator
go run cmd/generate/mcp_generator.go
```

这将生成：
- `mcp/communicator_mcp.go` - MCP 服务器实现
- `mcp/helper.go` - 辅助函数

### 3. 网络服务器

`server/server.go` 实现与 Rust 驱动的网络通信：

- 使用 TCP Socket 连接 Rust 驱动
- 处理消息序列化/反序列化
- 管理驱动连接状态
- 支持事件回调

### 4. 协议定义

`protocol/protocol.go` 定义与 Rust 驱动通信的消息格式：

- 消息类型
- 事件类型
- 调试状态
- 寄存器状态
- 断点类型

## MCP 集成

### MCP 工具列表

生成的 MCP 服务器将暴露以下工具：

1. **Start** - 启动调试器服务器
2. **Stop** - 停止调试器服务器
3. **WaitForDriver** - 等待驱动连接
4. **Initialize** - 初始化调试器
5. **Terminate** - 终止调试器
6. **AttachProcess** - 附加到进程
7. **DetachProcess** - 分离进程
8. **SetBreakpoint** - 设置断点
9. **RemoveBreakpoint** - 移除断点
10. **Continue** - 继续执行
11. **Pause** - 暂停执行
12. **StepInto** - 单步进入
13. **StepOver** - 单步跳过
14. **StepOut** - 单步跳出
15. **ReadMemory** - 读取内存
16. **WriteMemory** - 写入内存
17. **ReadRegisters** - 读取寄存器
18. **WriteRegisters** - 写入寄存器
19. **GetState** - 获取调试状态
20. **WaitForEvent** - 等待事件
21. **ExecuteScript** - 执行脚本
22. **ExecuteScriptWithContext** - 在上下文中执行脚本

### 使用 MCP 客户端

```go
package main

import (
    "context"
    "hyperdbg-communicator/mcp"
    "hyperdbg-communicator/debugger"
)

func main() {
    // 创建调试器实例
    dbg := debugger.NewDebugger()
    
    // 创建 MCP 服务器
    mcpServer := mcp.NewCommunicatorMCPServer(dbg)
    
    // 注册 MCP 工具
    mcpServer.RegisterTools(mcp.DefaultServer)
    
    // 启动调试器
    ctx := context.Background()
    mcpServer.HandleStart(ctx, mcp.CallToolRequest{}, mcp.CommunicatorMCPServerStartParams{
        Port: 8080,
    })
    
    // 等待驱动连接
    mcpServer.HandleWaitForDriver(ctx, mcp.CallToolRequest{}, mcp.CommunicatorMCPServerWaitForDriverParams{
        Timeout: 30 * time.Second,
    })
    
    // 初始化调试器
    mcpServer.HandleInitialize(ctx, mcp.CallToolRequest{}, mcp.CommunicatorMCPServerInitializeParams{})
}
```

## 与 Rust 驱动的通信

### 通信流程

1. **启动阶段**
   - Go Communicator 启动网络服务器
   - Rust 驱动连接到服务器
   - 建立连接后进行握手

2. **命令执行**
   - MCP 客户端调用工具
   - Go Communicator 转换为协议消息
   - 通过网络发送到 Rust 驱动
   - Rust 驱动执行命令并返回结果
   - Go Communicator 解析结果并返回给 MCP 客户端

3. **事件通知**
   - Rust 驱动生成事件（断点、异常等）
   - 通过网络发送到 Go Communicator
   - Go Communicator 触发回调
   - MCP 客户端接收事件通知

### 消息格式

所有消息遵循以下格式：

```go
type Message struct {
    Header  MessageHeader
    Payload []byte
}

type MessageHeader struct {
    Type    uint32
    Length  uint32
    DriverID uint64
    RequestID uint64
}
```

## 功能实现策略

### 不在 Rust 驱动中实现的功能

以下功能通过 Go 方法实现，通过 MCP 暴露：

| 功能 | 实现方式 | 说明 |
|------|----------|------|
| **调试器命令** | Go 方法 + MCP | 所有调试器命令作为 Go 方法暴露 |
| **脚本引擎** | Go 方法 + MCP | 脚本解析和执行在 Go 层实现 |
| **反汇编** | Rust 驱动 + Go 方法 | Rust 驱动提供反汇编能力，Go 方法暴露接口 |
| **符号解析** | Rust 驱动 + Go 方法 | Rust 驱动提供 PDB 解析，Go 方法暴露接口 |

### Rust 驱动提供的能力

Rust 驱动提供底层能力，通过 Go 方法暴露：

- **VMX/EPT 操作** - Hypervisor 层功能
- **Hook 机制** - EPT Hook、Syscall Hook、Inline Hook
- **内存操作** - 读取/写入内存
- **寄存器操作** - 读取/写入寄存器
- **事件生成** - 断点、异常等事件

### Go 层提供的能力

Go 层提供高级功能，通过 MCP 暴露：

- **调试器命令** - 所有调试器命令
- **脚本引擎** - 脚本解析和执行
- **反汇编接口** - 封装 Rust 驱动的反汇编能力
- **符号解析接口** - 封装 Rust 驱动的 PDB 解析能力
- **网络调试** - 通过 MCP 支持远程调试

## 开发指南

### 添加新的调试器方法

1. 在 `debugger/interfaces.go` 中添加接口方法
2. 在 `debugger/debugger.go` 中实现方法
3. 运行 MCP 生成器：

```bash
go run cmd/generate/mcp_generator.go
```

4. 新的 MCP 工具将自动生成

### 添加新的消息类型

1. 在 `protocol/protocol.go` 中定义消息类型
2. 在 `server/server.go` 中添加消息处理逻辑
3. 在 `debugger/debugger.go` 中添加相应方法

### 测试

```bash
# 启动调试器
go run cmd/main.go

# 测试 MCP 连接
# 使用 MCP 客户端连接到服务器
```

## 依赖

- `github.com/modelcontextprotocol/go-sdk/mcp` - MCP SDK
- `github.com/ddkwork/golibrary/std/mylog` - 日志库

## 许可证

与 HyperDbg 项目保持一致
