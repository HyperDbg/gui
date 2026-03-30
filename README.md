# HyperDbg Rust

HyperDbg 调试器的 Rust 实现，包含内核驱动和代码生成工具。

## 项目概述

本项目是 HyperDbg 调试器的 Rust 移植版本，使用 Rust 的内存安全特性来增强内核驱动的可靠性。

## 目录结构

```
HyperDbg_rust/
├── cmd/                        # 命令行工具
│   ├── mcp/                    # MCP 服务器 (自动生成)
│   │   └── mcp.go
│   └── rustgen/                # Rust 代码生成器
│       ├── main.go             # 入口
│       ├── rustgen.go          # 类型生成
│       ├── wdkgen.go           # WDK 绑定
│       └── mcpgen.go           # MCP 生成
│
├── debugger/                   # Go 类型定义
│   ├── interfaces.go           # 接口定义
│   ├── types.go                # 类型定义
│   ├── packet.go               # 通信协议
│   └── event_*.go              # 事件类型
│
└── rust-driver/                # Rust 驱动
    └── kd/                     # 主驱动 crate
        ├── Cargo.toml
        ├── build.ps1           # 构建脚本
        └── src/
            ├── lib.rs          # 入口
            ├── generated/      # 自动生成
            ├── go_script/      # Go 脚本引擎
            ├── hyperkd/        # 调试器
            ├── net/            # HTTP 服务器
            └── framework/      # 驱动框架
```

## 核心功能

### 1. Hypervisor (hyperhv)

- VMX 操作 (VM Entry/Exit)
- EPT 内存虚拟化
- Hook 机制 (EPT Hook, Inline Hook)
- 多处理器支持

### 2. Debugger (hyperkd)

- 内核调试
- 用户态调试
- 内存读写
- 断点管理
- 调用栈追踪

### 3. 通信层

- WSK HTTP Server (端口 50080)
- JSON 序列化/反序列化
- 事件通知机制

### 4. Go 脚本引擎

- Go 语法解析器
- Hook 脚本执行
- 运行时代码生成

## 构建要求

| 组件 | 版本要求 |
|------|---------|
| Rust | 1.85+ |
| Go | 1.21+ |
| WDK | Windows 11 SDK |
| EWDK | 最新版本 |

## 快速开始

### 1. 生成代码

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\cmd\rustgen
go run .
```

### 2. 检查编译

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd
cargo check
```

### 3. 构建驱动

```powershell
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd
powershell -ExecutionPolicy Bypass -File build.ps1
```

## 架构图

```
┌─────────────────────────────────────────────────────────────────┐
│                      用户空间 (Go)                               │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐             │
│  │   CLI UI    │  │  MCP Server │  │   Client    │             │
│  └──────┬──────┘  └──────┬──────┘  └──────┬──────┘             │
│         │                │                │                     │
│         └────────────────┼────────────────┘                     │
│                          │ HTTP/JSON                            │
└──────────────────────────┼──────────────────────────────────────┘
                           │ :50080
┌──────────────────────────┼──────────────────────────────────────┐
│                      内核空间 (Rust)                             │
│  ┌───────────────────────▼───────────────────────────────────┐  │
│  │                 net (WSK HTTP Server)                     │  │
│  └───────────────────────┬───────────────────────────────────┘  │
│                          │                                      │
│  ┌───────────────────────▼───────────────────────────────────┐  │
│  │              generated (类型 + 处理器)                     │  │
│  │  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐       │  │
│  │  │ 事件类型    │  │ API 路由器  │  │ Hook 数据库 │       │  │
│  │  └─────────────┘  └─────────────┘  └─────────────┘       │  │
│  └───────────────────────┬───────────────────────────────────┘  │
│                          │                                      │
│  ┌───────────────────────▼───────────────────────────────────┐  │
│  │                    hyperkd (调试器)                       │  │
│  │  ┌─────────────────────────────────────────────────────┐ │  │
│  │  │                 hyperhv (Hypervisor)                │ │  │
│  │  │  ┌──────────┐  ┌──────────┐  ┌──────────┐          │ │  │
│  │  │  │   VMX    │  │   EPT    │  │  Hooks   │          │ │  │
│  │  │  └──────────┘  └──────────┘  └──────────┘          │ │  │
│  │  └─────────────────────────────────────────────────────┘ │  │
│  └───────────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────────┘
```

## 文档

- [SKILL.md](.trae/skills/hyperdbg-rust-driver/SKILL.md) - 开发技能指南
- [rust-driver/README.md](rust-driver/README.md) - 驱动说明
- [rust-driver/kd/IMPLEMENTATION_STATUS.md](rust-driver/kd/IMPLEMENTATION_STATUS.md) - 实现状态

## 许可证

GPL-3.0
