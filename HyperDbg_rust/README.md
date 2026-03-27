# HyperDbg Rust - Go Communicator

## 概述

HyperDbg Rust 是 HyperDbg 调试器的 Rust 重实现项目，包含 Go 用户空间通信层和 Rust 内核驱动。Go Communicator 负责与 Rust 内核驱动进行网络通信，并通过 MCP (Model Context Protocol) 暴露调试器功能。

## 架构

```
┌─────────────────────────────────────────────────────┐
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
                     │ HTTP/JSON (WSK Socket :50080)
┌────────────────────▼────────────────────────────────────┐
│              Rust 内核驱动                            │
├──────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │   hyperhv    │  │   hyperkd    │  │   pdbex      │  │
│  │ (Hypervisor) │  │  (Debugger)  │  │ (PDB Parser) │  │
│  └──────────────┘  └──────────────┘  └──────────────┘  │
│         │                 │                  │            │
│  ┌──────▼──────────────────▼──────────────────▼──────┐  │
│  │        common (types_gen + handlers_gen)          │  │
│  │   - types_gen/    : 生成的类型定义                 │  │
│  │   - handlers_gen/ : 生成的 API 调度器              │  │
│  └────────────────────────────────────────────────────┘  │
│  ┌────────────────────────────────────────────────────┐  │
│  │                 net (WSK HTTP Server)             │  │
│  │   - HTTP 服务器监听 50080 端口                    │  │
│  │   - JSON 序列化/反序列化                          │  │
│  │   - Request/Response 结构体                       │  │
│  └────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
```

## 目录结构

```
HyperDbg_rust/
├── doc.go                 // 包文档说明
├── hyperdbg.go            // 导出所有公共类型和函数
├── config.go              // 常量定义
├── types.go               // 类型别名
├── response.go            // Response[T] + EventCallback
├── packet.go              // Packet 结构体 + SendReceive
├── event_handlers.go      // EventHandlers 结构体 + Dispatch + OnBreakpoint/OnException 等注册方法
├── register.go            // RegisterState (共享结构体，被多个事件引用)
├── event_*.go             // 事件类型定义
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
│   ├── kd/                // 主驱动 crate (整合所有模块)
│   │   ├── Cargo.toml
│   │   ├── build.ps1
│   │   ├── build.rs
│   │   └── src/
│   │       ├── lib.rs             // 驱动入口点
│   │       ├── kd.rs              // 内核调试器
│   │       ├── ud.rs              // 用户调试器
│   │       ├── common/            // 内部通用模块 (types_gen + handlers_gen)
│   │       ├── logger/            // 内部日志模块
│   │       ├── net/               // 内部网络模块 (WSK HTTP Server)
│   │       ├── framework/         // 内部驱动框架
│   │       ├── disassembler/      // 反汇编器
│   │       ├── pdbex/             // PDB 解析
│   │       └── hyperkd/           // 调试器核心
│   │           └── hyperhv/       // Hypervisor 实现
│   │
│   ├── driver-framework/  // 驱动框架库 (独立 crate，用于 examples)
│   ├── logger/            // 日志模块 (独立 crate，用于 examples)
│   ├── common/            // 通用模块 (独立 crate，用于 examples)
│   ├── net/               // 网络模块 (独立 crate，用于 examples)
│   │
│   └── examples/          // 示例驱动
│       ├── sysdemo/       // WDM 驱动示例
│       └── netdemo/       // 网络驱动示例
│
├── rust-driver/IMPLEMENTATION_PROGRESS.md  // 实现进度
├── api_design.md          // API 设计文档
├── README.md              // 本文档
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

### 4. 网络服务器

`server/server.go` 实现与 Rust 驱动的网络通信：

- 使用 HTTP/JSON 协议与 Rust 驱动通信
- 驱动监听 50080 端口
- 处理消息序列化/反序列化
- 管理驱动连接状态
- 支持事件回调

### 5. 协议定义

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

## 与 Rust 驱动的通信

### 通信流程

1. **启动阶段**
   - Rust 驱动启动 WSK HTTP Server，监听 50080 端口
   - Go Communicator 连接到驱动
   - 建立连接后进行握手

2. **命令执行**
   - MCP 客户端调用工具
   - Go Communicator 转换为 JSON 请求
   - 通过 HTTP POST 发送到 Rust 驱动
   - Rust 驱动执行命令并返回 JSON 结果
   - Go Communicator 解析结果并返回给 MCP 客户端

3. **事件通知**
   - Rust 驱动生成事件（断点、异常等）
   - 通过 `emit_*_event` 函数推送到 EventQueue
   - Go Communicator 轮询获取事件
   - MCP 客户端接收事件通知

### HTTP API 端点

| 端点 | 方法 | 说明 |
|------|------|------|
| `/api` | POST | API 请求处理 |
| `/ping` | GET | 健康检查 |
| `/status` | GET | 驱动状态 |

### JSON 请求格式

```json
{
    "action": "initialize",
    "process_id": 1234,
    ...
}
```

### JSON 响应格式

```json
{
    "success": true,
    "message": "",
    "data": { ... }
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
| 修改 `interfaces.go` (MCP) | `cd cmd/mcp && go run main.go` | 更新 MCP 服务器代码 |

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

4. **修改 `interfaces.go` (MCP) 后**
   - 生成器会自动扫描 MCP 相关的接口定义
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
- [api_design.md](api_design.md) - API 设计详细文档
- [.trae/skills/rust-driver-development/SKILL.md](.trae/skills/rust-driver-development/SKILL.md) - 开发技能

## 依赖

- `github.com/modelcontextprotocol/go-sdk/mcp` - MCP SDK
