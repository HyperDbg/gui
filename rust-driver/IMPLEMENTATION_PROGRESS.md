# Rust 驱动实现进度

## 背景

经过 15 个空指针等基础问题导致的 BSOD 修复后，决定使用 Rust 重新实现 HyperDbg 驱动。Rust 的内存安全特性和编译时检查可以有效避免 C++ 中常见的空指针访问、缓冲区溢出等问题。

## 编译状态

**当前状态**: ✅ 编译通过，无警告

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd
cargo check
# Finished `dev` profile [unoptimized + debuginfo] target(s) in 0.12s
```

## 新架构设计

### 架构分层

```
┌─────────────────────────────────────────────────────────────┐
│                   Go 用户空间通信层                          │
│                   (packet.go + event_handlers.go)            │
└────────────────────┬────────────────────────────────────────┘
                     │ HTTP/JSON (WSK Socket :50080)
┌────────────────────▼────────────────────────────────────────┐
│              Rust 内核驱动层                                 │
├──────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐       │
│  │   hyperhv    │  │   hyperkd    │  │   pdbex      │       │
│  │ (Hypervisor) │  │  (Debugger)  │  │ (PDB Parser) │       │
│  └──────────────┘  └──────────────┘  └──────────────┘       │
│         │                 │                  │               │
│  ┌──────▼──────────────────▼──────────────────▼─────────┐   │
│  │        generated (types + handlers + hooks)           │   │
│  │   - event_*.rs    : 事件类型定义 (22个文件)            │   │
│  │   - router.rs     : API 路由器                        │   │
│  │   - emit.rs       : 事件发射器                        │   │
│  │   - hook_db.rs    : Hook 数据库                       │   │
│  └───────────────────────────────────────────────────────┘   │
│  ┌───────────────────────────────────────────────────────┐   │
│  │                 net (WSK HTTP Server)                 │   │
│  │   - HTTP 服务器监听 50080 端口                         │   │
│  │   - JSON 序列化/反序列化                               │   │
│  └───────────────────────────────────────────────────────┘   │
└──────────────────────────────────────────────────────────────┘
```

## 项目结构

```
rust-driver/
└── kd/                         # 主驱动 crate
    ├── Cargo.toml
    ├── build.ps1               # 构建脚本 (EWDK)
    ├── build.rs                # 构建脚本 (Rust)
    └── src/
        ├── lib.rs              # 驱动入口
        ├── kd.rs               # 内核调试器
        ├── ud.rs               # 用户调试器
        ├── hyperdbg_api.rs     # API 实现
        │
        ├── generated/          # 自动生成目录
        │   ├── mod.rs
        │   ├── common.rs       # 基础类型
        │   ├── register.rs     # 寄存器状态
        │   ├── response.rs     # 响应类型
        │   ├── event.rs        # 事件枚举
        │   ├── event_*.rs      # 事件类型 (22个文件)
        │   ├── router.rs       # API 路由器
        │   ├── emit.rs         # 事件发射器
        │   ├── ioctl.rs        # IOCTL 代码
        │   ├── hook_db.rs      # Hook 数据库
        │   ├── hook_types.rs   # Hook 类型
        │   ├── hook_args.rs    # Hook 参数
        │   ├── ept_hook.rs     # EPT Hook 数据库
        │   ├── inline_hook.rs  # Inline Hook 数据库
        │   ├── ntddk.rs        # WDK 函数导出 (1251个)
        │   ├── types.rs        # WDK 类型 (7317个)
        │   └── constants.rs    # WDK 常量 (5348个)
        │
        ├── go_script/          # Go 脚本引擎
        │   ├── parser/         # Go 解析器 (10个文件)
        │   ├── analyzer.rs
        │   ├── executor.rs
        │   └── generator.rs
        │
        ├── hyperkd/            # 调试器核心
        │   ├── hyperhv/        # Hypervisor 实现
        │   │   ├── vmm/        # VMX/EPT 核心
        │   │   ├── hooks/      # Hook 机制
        │   │   ├── memory/     # 内存管理
        │   │   ├── processor/  # 处理器管理
        │   │   └── interface/  # 接口层
        │   ├── debugger.rs
        │   ├── commands.rs
        │   └── network.rs
        │
        ├── net/                # WSK HTTP Server
        │   ├── http.rs
        │   ├── json.rs
        │   └── util.rs
        │
        ├── framework/          # 驱动框架
        │   ├── device.rs
        │   ├── ioctl.rs
        │   └── utils.rs
        │
        ├── logger/             # 日志模块
        │   └── debug_logger.rs
        │
        ├── disassembler/       # 反汇编器
        └── pdbex/              # PDB 解析器
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

## 实现进度

### 核心模块

| 模块 | 路径 | 进度 | 说明 |
|------|------|------|------|
| **generated** | `kd/src/generated/` | ✅ 100% | 自动生成模块 |
| **go_script** | `kd/src/go_script/` | ✅ 100% | Go 脚本引擎 |
| **net** | `kd/src/net/` | ✅ 100% | WSK HTTP Server |
| **framework** | `kd/src/framework/` | ✅ 100% | 驱动框架 |
| **logger** | `kd/src/logger/` | ✅ 100% | 日志模块 |
| **hyperhv** | `kd/src/hyperkd/hyperhv/` | 🔄 开发中 | Hypervisor 核心 |
| **hyperkd** | `kd/src/hyperkd/` | 🔄 开发中 | 调试器核心 |
| **disassembler** | `kd/src/disassembler/` | 🔄 开发中 | 反汇编器 |
| **pdbex** | `kd/src/pdbex/` | 🔄 开发中 | PDB 解析器 |

### generated 模块（自动生成）

| 类别 | 文件数 | 说明 |
|------|--------|------|
| 事件类型 | 22 | event_*.rs |
| Hook 相关 | 5 | hook_db.rs, hook_types.rs, hook_args.rs, ept_hook.rs, inline_hook.rs |
| WDK 绑定 | 3 | ntddk.rs (1251函数), types.rs (7317类型), constants.rs (5348常量) |
| API 处理 | 3 | router.rs, emit.rs, ioctl.rs |
| 基础类型 | 3 | common.rs, register.rs, response.rs |

### go_script 模块

| 文件 | 内容 | 进度 |
|------|------|------|
| `parser/` | Go 解析器 (10个文件) | ✅ 100% |
| `analyzer.rs` | AST 分析器 | ✅ 100% |
| `executor.rs` | 脚本执行器 | ✅ 100% |
| `generator.rs` | Rust 代码生成器 | ✅ 100% |

### net 模块（WSK HTTP Server）

| 功能 | 状态 |
|------|------|
| WSK Socket 监听 | ✅ 100% |
| HTTP 请求解析 | ✅ 100% |
| JSON 序列化 | ✅ 100% |
| JSON 反序列化 | ✅ 100% |

### hyperhv 模块（Hypervisor 层）

| 子模块 | 状态 | 文件数 |
|--------|------|--------|
| `vmm/ept/` | 🔄 | 3 |
| `vmm/vmx/` | 🔄 | 8 |
| `hooks/` | 🔄 | 5 |
| `memory/` | 🔄 | 6 |
| `processor/` | 🔄 | 5 |
| `interface/` | 🔄 | 10 |
| `broadcast/` | 🔄 | 4 |
| `features/` | 🔄 | 7 |
| `devices/` | 🔄 | 3 |
| `common/` | 🔄 | 3 |
| `assembly/` | 🔄 | 3 |
| `globals/` | 🔄 | 2 |
| `components/` | 🔄 | 2 |
| `disassembler/` | 🔄 | 1 |
| `mmio/` | 🔄 | 1 |

### hyperkd 模块（Debugger 层）

| 文件 | 内容 | 状态 |
|------|------|------|
| `debugger.rs` | 调试器核心 | 🔄 |
| `commands.rs` | 命令处理 | 🔄 |
| `network.rs` | 网络通信 | ✅ |
| `user_debug.rs` | 用户态调试 | 🔄 |
| `callstack.rs` | 调用栈 | 🔄 |
| `attaching.rs` | 进程附加 | 🔄 |
| `process.rs` | 进程管理 | 🔄 |
| `thread.rs` | 线程管理 | 🔄 |
| `allocations.rs` | 内存分配 | 🔄 |
| `dpc_routines.rs` | DPC 例程 | 🔄 |
| `extension_commands.rs` | 扩展命令 | 🔄 |
| `ffi.rs` | FFI 接口 | 🔄 |
| `halted_broadcast.rs` | 停止广播 | 🔄 |
| `serial_connection.rs` | 串口连接 | 🔄 |
| `user_access.rs` | 用户访问 | 🔄 |

## 代码生成器

### 入口命令

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\cmd\rustgen
go run .
```

### 生成器功能

| 生成器 | 文件 | 功能 |
|--------|------|------|
| Rust 类型生成 | rustgen.go | 从 Go 代码生成 Rust 类型 |
| WDK 绑定验证 | wdkgen.go | 验证 WDK 函数签名一致性 |
| MCP 服务器生成 | mcpgen.go | 从接口生成 MCP 服务器 |

### 生成器输出

```
[1/3] Generating Rust types...
  - Output: rust-driver/kd/src/generated/
  - 生成 27 个文件

[2/3] Generating WDK bindings...
  - Output: rust-driver/kd/src/generated/
  - Loaded WDK bindings: 1251 functions, 7317 types, 5348 constants

[3/3] Generating MCP server...
  - Server: cmd/mcp/mcp.go
```

## 下一步计划

### 主要任务

1. **完善 hyperhv 模块** - Hypervisor 核心实现
2. **完善 hyperkd 模块** - 调试器核心实现
3. **添加测试用例** - 单元测试和集成测试
4. **文档完善** - API 文档和使用指南

## 里程碑

### 2025-03-31: 编译警告修复完成

- 修复所有编译警告
- 添加必要的 `#![allow(...)]` 属性
- 更新代码生成器以自动添加警告抑制

### 2025-03-31: 文档更新

- 更新 SKILL.md
- 更新 README.md
- 更新 IMPLEMENTATION_STATUS.md

### 2025-03-27: hyperhv/hyperkd 模块完善

- hyperhv 模块包含完整的框架代码（60+ 文件）
- hyperkd 模块包含完整的框架代码（15+ 文件）
- 主驱动 lib.rs 实现完整的 HTTP 服务器入口

### 2024-03-27: WSK HTTP Server 完成

- 驱动监听 50080 端口
- HTTP 请求解析
- JSON 序列化/反序列化
- 多客户端连接

---

## 相关文档

- [SKILL.md](../../.trae/skills/hyperdbg-rust-driver/SKILL.md) - 开发技能指南
- [README.md](../README.md) - 项目概述
- [api_design.md](../api_design.md) - API 设计详细文档
- [kd/IMPLEMENTATION_STATUS.md](kd/IMPLEMENTATION_STATUS.md) - 实现状态报告
