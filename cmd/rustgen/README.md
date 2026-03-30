# Rust 代码生成器

从 Go 类型定义自动生成 Rust 代码的工具。

## 功能

1. **类型生成** - 将 Go 结构体转换为 Rust 结构体
2. **处理器生成** - 生成 API 请求处理器
3. **WDK 绑定验证** - 验证 WDK 函数签名一致性
4. **MCP 服务器生成** - 从接口定义生成 MCP 服务器

## 使用方法

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\cmd\rustgen
go run .
```

## 输出目录

```
rust-driver/kd/src/generated/
├── mod.rs           # 模块入口
├── common.rs        # 基础类型
├── register.rs      # 寄存器状态
├── response.rs      # 响应类型
├── event.rs         # 事件枚举
├── event_*.rs       # 事件类型 (22个文件)
├── router.rs        # API 路由器
├── emit.rs          # 事件发射器
├── ioctl.rs         # IOCTL 代码
├── hook_db.rs       # Hook 数据库
├── hook_types.rs    # Hook 类型
├── hook_args.rs     # Hook 参数
├── ept_hook.rs      # EPT Hook 数据库
├── inline_hook.rs   # Inline Hook 数据库
├── ntddk.rs         # WDK 函数导出
├── types.rs         # WDK 类型
└── constants.rs     # WDK 常量
```

## 源文件

| 文件 | 用途 |
|------|------|
| `debugger/types.go` | 类型定义 |
| `debugger/interfaces.go` | 接口定义 |
| `debugger/packet.go` | 通信协议 |
| `debugger/event_*.go` | 事件类型 |

## 生成器文件

| 文件 | 功能 |
|------|------|
| `main.go` | 入口点 |
| `rustgen.go` | Rust 类型生成 |
| `wdkgen.go` | WDK 绑定验证 |
| `mcpgen.go` | MCP 服务器生成 |

## 类型映射

| Go 类型 | Rust 类型 |
|---------|----------|
| `uint8` | `u8` |
| `uint16` | `u16` |
| `uint32` | `u32` |
| `uint64` | `u64` |
| `int32` | `i32` |
| `int64` | `i64` |
| `string` | `alloc::string::String` |
| `[]T` | `alloc::vec::Vec<T>` |
| `map[K]V` | `alloc::collections::BTreeMap<K, V>` |

## 生成的代码特性

- `#![allow(non_snake_case)]` - 允许非蛇形命名
- `#![allow(dead_code)]` - 允许未使用代码
- `#![allow(unused_imports)]` - 允许未使用导入
- `#![allow(clashing_extern_declarations)]` - 允许外部声明冲突

## WDK 绑定统计

| 类别 | 数量 |
|------|------|
| 函数 | 1251 |
| 类型 | 7317 |
| 常量 | 5348 |

## 相关文档

- [SKILL.md](../../.trae/skills/hyperdbg-rust-driver/SKILL.md) - 开发技能指南
- [IMPLEMENTATION_STATUS.md](../../rust-driver/kd/IMPLEMENTATION_STATUS.md) - 实现状态
