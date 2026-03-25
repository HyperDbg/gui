# Rust Windows 驱动开发指南

## ⚠️ 重要警告

**严禁直接使用 `cargo build` 命令！**

- 本项目使用 **EWDK (Enterprise Windows Driver Kit)** 进行编译
- **永远不会安装 VS2022**
- 只能使用 PowerShell 构建脚本进行编译
- 当用户要求执行某个 `.ps1` 脚本时，**必须直接执行该脚本**，不要尝试手动运行 cargo 命令

## 架构设计

### 调试器网络通信架构

```
┌─────────────────────────────────────────────────────────────────────┐
│                        HyperDbg 调试系统                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ┌─────────────────┐                              ┌─────────────────┐ │
│  │   Go 用户态      │                              │   Rust 驱动      │ │
│  │                 │                              │                 │ │
│  │  TCP Server     │◄─────── JSON 协议 ─────────►│  TCP Client     │ │
│  │  (监听 9527)    │                              │  (主动连接)      │ │
│  │                 │                              │                 │ │
│  │  - 解析 JSON    │                              │  - 解析 JSON     │ │
│  │  - 发送命令     │                              │  - 执行调试     │ │
│  │  - 接收结果     │                              │  - 返回结果     │ │
│  └────────┬────────┘                              └────────┬────────┘ │
│           │                                                │          │
│           │                                                │          │
│           └──────────────────┐ ┌──────────────────────────┘          │
│                              │ │                                      │
│                              │ │   TCP Socket 连接                     │
└──────────────────────────────┼─┼──────────────────────────────────────┘
                               │ │
                               │ ▼
                    ┌──────────────────────┐
                    │    Windows 内核        │
                    │                       │
                    │  WSK (Winsock Kernel) │
                    │  - TcpClientConnect   │
                    │  - TcpSend / TcpRecv  │
                    │                       │
                    └──────────────────────┘
```

### 为什么选择 JSON + TCP Client 架构？

| 对比项 | IOCTL + 结构体 | JSON + TCP Client |
|--------|----------------|-------------------|
| 内存对齐 | 必须严格对齐，容易出错 | 无需对齐，JSON 是文本协议 |
| 调试 | 困难，二进制数据难读 | 方便，直接看到明文 |
| 跨语言 | Go/Rust 结构体需一一对应 | 只需共享 JSON Schema |
| 依赖 | 简单，WDK 内置 | 需要 WSK 或 smoltcp |
| 灵活性 | 结构体固定，增减字段麻烦 | JSON 可扩展性好 |
| 远程调试 | 不支持 | 支持，Go 和驱动可不在同一机器 |

### 通信流程

```
1. Go 启动 TCP Server 监听 9527
2. Rust 驱动加载后，主动连接 Go Server
3. Go 发送 JSON 命令:
   {
     "action": "read_memory",
     "target": 0xFFFFF80000000000,
     "size": 8
   }
4. Rust 驱动解析 JSON，执行对应操作
5. Rust 驱动返回 JSON 结果:
   {
     "status": "ok",
     "data": "deadbeef12345678"
   }
```

## 构建环境

### 使用通用构建脚本

项目使用 EWDK (Enterprise Windows Driver Kit) 进行编译，**不要直接使用 `cargo build`**。

```powershell
# 使用通用构建脚本编译所有驱动项目
.\rust-driver\examples\build-all.ps1
```

### EWDK 配置

在 `build-all.ps1` 中配置：
```powershell
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
```

## driver-framework 通用框架

### 位置

`rust-driver/driver-framework/` - 所有驱动项目共用的通用框架

### 结构

```
driver-framework/src/
├── lib.rs        # 模块导出
├── logger.rs     # 内核日志宏 (log!)
├── utils.rs      # UTF-16 转换工具 (utf16!, ToU16Vec, ToUnicodeString)
├── ffi.rs        # Windows 内核 FFI 辅助函数
├── device.rs     # 设备创建/清理通用逻辑
└── ioctl.rs      # IOCTL 辅助宏和函数
```

## wsk 模块 (WSK 网络封装)

### 位置
`rust-driver/wsk/` - Winsock Kernel 网络套接字封装

### 文件结构
```
wsk/
├── Cargo.toml
└── src/
    ├── lib.rs          # 模块导出
    ├── socket.rs       # WskSocket, WskListener 实现
    ├── buffer.rs       # 数据缓冲区
    ├── provider.rs     # WSK Provider (stub)
    └── error.rs        # 错误类型
```

### 核心类型

```rust
// Socket 类型
pub struct WskSocket {
    socket: *mut u8,
    dispatch: *const WSK_PROVIDER_DISPATCH_STRUCT,
    socket_type: SocketType,
    protocol: Protocol,
    connected: bool,
    bound: bool,
    recv_buffer: Buffer,
    send_buffer: Buffer,
}

// 监听器 (驱动作为 TCP Client 不需要这个)
pub struct WskListener {
    socket: WskSocket,
    backlog: u32,
}
```

### 依赖配置

```toml
[dependencies]
serde = { version = "1.0", default-features = false, features = ["derive"] }
serde_json = { version = "1.0", default-features = false, features = ["alloc"] }
```

### JSON 解析示例

```rust
use serde::{Deserialize, Serialize};

#[derive(Deserialize, Serialize)]
pub struct Command {
    pub action: String,
    pub target: u64,
    pub value: u64,
}

fn parse_json(data: &[u8]) -> Option<Command> {
    serde_json::from_slice(data).ok()
}
```

## netdemo 网络测试驱动

### 项目位置
`rust-driver/examples/netdemo/`

### 文件结构
```
netdemo/
├── Cargo.toml     # 依赖配置
├── build.ps1      # 构建脚本
└── src/
    ├── lib.rs     # 驱动主代码
    └── m.go       # Go 测试程序
```

### 当前状态
- **IOCTL 功能**：已实现，用于本地通信测试
- **网络功能**：驱动作为 TCP Client 连接 Go 服务端，使用 JSON 通信

## 项目结构

### 已验证可编译的项目

1. **driver-framework** - `rust-driver/driver-framework/`
   - 通用驱动框架库
   - 所有驱动项目的基础依赖

2. **sysdemo** - `rust-driver/examples/sysdemo/`
   - 模块化 WDM 驱动示例
   - 支持 IOCTL 通信

3. **netdemo** - `rust-driver/examples/netdemo/`
   - 网络通信测试驱动
   - 使用 TCP Client + JSON 协议

4. **protocol** - `rust-driver/protocol/`
   - 调试器通信协议定义 (JSON Schema)
   - `types.rs` - 基本类型
   - `events.rs` - 事件类型
   - 常量: `PROTOCOL_VERSION = 1`, `DEFAULT_PORT = 9527`

### 不可用的模块（未实现）

- `hyperhv` - Hypervisor 核心
- `hyperkd` - 内核调试器
- `disassembler` - 反汇编器
- `pdbex` - PDB 解析
- `kd` - 内核调试支持

## 常见错误

### 1. "The parameter is incorrect"
原因：`UNICODE_STRING` 初始化错误
解决：使用 `RtlInitUnicodeString` 正确初始化

### 2. BSOD (蓝屏)
原因：内存访问错误、类型不匹配
解决：检查指针操作、确保 UTF-16 字符串正确

### 3. 链接错误
原因：未使用 EWDK 环境
解决：使用 `build-all.ps1` 构建

### 调试器网络通信架构

```
.==========================================================================.
||                        HyperDbg 调试系统                               ||
||======================================================================||
||                                                                      ||
||    Go 用户态                        Rust 驱动                        ||
||  +----------------+              +----------------+                  ||
||  |                |              |                |                  ||
||  |  TCP Server    |<====JSON====>|  TCP Client   |                  ||
||  |  (监听 9527)   |              |  (主动连接)    |                  ||
||  |                |              |                |                  ||
||  |  - 解析 JSON   |              |  - 解析 JSON   |                  ||
||  |  - 发送命令    |              |  - 执行调试    |                  ||
||  |  - 接收结果    |              |  - 返回结果    |                  ||
||  +-------+--------+              +--------+-------+                  ||
||          |                                  |                         ||
||          |            TCP Socket            |                         ||
||          +===============| |===============+                         ||
||                          | |                                         ||
||                          ▼ ▼                                         ||
||               +----------------------+                               ||
||               |   Windows 内核        |                               ||
||               |                      |                               ||
||               |  WSK (Winsock Kernel) |                               ||
||               |  - TcpClientConnect  |                               ||
||               |  - TcpSend / TcpRecv |                               ||
||               +----------------------+                               ||
||==========================================================================|
`````

### 为什么选择 JSON + TCP Client 架构？

```
+-------------------+        +-------------------+
|   IOCTL + 结构体   |        |  JSON + TCP Client |
+===================+        +===================+
| [X] 内存对齐       |        | [ ] 内存对齐       |
|     必须严格对齐    |        |     无需对齐       |
+-------------------+        +-------------------+
| [X] 调试困难       |        | [√] 调试方便       |
|     二进制数据难读  |        |     直接看明文     |
+-------------------+        +-------------------+
| [X] 跨语言复杂     |        | [√] 跨语言简单     |
|     结构体需对应    |        |     只需 JSON     |
+-------------------+        +-------------------+
| [√] 依赖简单       |        | [X] 需要 WSK      |
|     WDK 内置      |        |     或 smoltcp    |
+-------------------+        +-------------------+
| [X] 灵活性差       |        | [√] 灵活性好      |
|     结构体需固定    |        |     JSON 可扩展   |
+-------------------+        +-------------------+
| [X] 不支持远程调试  |        | [√] 支持远程调试   |
+-------------------+        +-------------------+
```

### 通信流程

```
┌──────────────────────────────────────────────────────────────────┐
│ 1. Go 启动 TCP Server 监听 9527                                  │
│                                                                  │
│ 2. Rust 驱动加载后，主动连接 Go Server                           │
│                                                                  │
│ 3. Go 发送 JSON 命令:                                           │
│    ┌─────────────────────────────────────────────────────────┐   │
│    │ {                                                         │   │
│    │   "action": "read_memory",                               │   │
│    │   "target": "0xFFFFF80000000000",                        │   │
│    │   "size": 8                                              │   │
│    │ }                                                         │   │
│    └─────────────────────────────────────────────────────────┘   │
│                                                                  │
│ 4. Rust 驱动解析 JSON，执行对应操作                              │
│                                                                  │
│ 5. Rust 驱动返回 JSON 结果:                                     │
│    ┌─────────────────────────────────────────────────────────┐   │
│    │ {                                                         │   │
│    │   "status": "ok",                                        │   │
│    │   "data": "deadbeef12345678"                            │   │
│    │ }                                                         │   │
│    └─────────────────────────────────────────────────────────┘   │
└──────────────────────────────────────────────────────────────────┘
```

## 项目结构

```
rust-driver/
|
├── driver-framework/          # 通用驱动框架库
│   └── src/
│       ├── lib.rs             # 模块导出
│       ├── logger.rs          # 日志宏 (log!)
│       ├── utils.rs           # UTF-16 转换工具
│       ├── ffi.rs             # Windows FFI
│       ├── device.rs          # 设备创建/清理
│       └── ioctl.rs           # IOCTL 辅助
|
├── wsk/                       # WSK 网络封装
│   ├── Cargo.toml
│   └── src/
│       ├── lib.rs             # 模块导出
│       ├── socket.rs          # WskSocket 实现
│       ├── buffer.rs          # 数据缓冲区
│       ├── provider.rs         # WSK Provider
│       └── error.rs           # 错误类型
|
├── protocol/                  # 通信协议定义
│   ├── types.rs               # 基本类型
│   └── events.rs              # 事件类型
|
├── examples/
│   ├── sysdemo/               # 驱动示例 (IOCTL)
│   └── netdemo/               # 网络测试驱动 (TCP Client + JSON)
|
└── erebus-main/               # 完整驱动框架
    ├── km/                    # 内核模式
    └── um/                    # 用户模式
```

## 下一步计划

```
1. [ ] 实现 wsk stub              -> 让 netdemo 编译通过
2. [ ] 实现 TCP Client 连接        -> 驱动主动连接 Go 服务端
3. [ ] 集成 serde_json            -> 驱动端 JSON 解析
4. [ ] Go 端 TCP Server           -> 发送 JSON 命令测试
```
