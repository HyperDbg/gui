# HyperKD - 内核调试器

HyperDbg 的内核调试器模块，提供调试功能和命令处理。

## 目录结构

```
hyperkd/
├── hyperhv/               # Hypervisor 实现
│   └── ...               # (见 hyperhv/README.md)
│
├── debugger.rs            # 调试器核心
├── commands.rs            # 命令处理
├── network.rs             # 网络通信
├── user_debug.rs          # 用户态调试
├── callstack.rs           # 调用栈追踪
├── attaching.rs           # 进程附加
├── process.rs             # 进程管理
├── thread.rs              # 线程管理
├── allocations.rs         # 内存分配
├── dpc_routines.rs        # DPC 例程
├── extension_commands.rs  # 扩展命令
├── ffi.rs                 # FFI 接口
├── halted_broadcast.rs    # 停止广播
├── serial_connection.rs   # 串口连接
├── user_access.rs         # 用户访问
├── ud.rs                  # 用户调试器
└── mod.rs                 # 模块入口
```

## 核心功能

### 调试器核心 (debugger.rs)

- 调试会话管理
- 断点管理
- 单步执行
- 内存读写

### 命令处理 (commands.rs)

- 命令解析
- 命令分发
- 结果返回

### 网络通信 (network.rs)

- HTTP 请求处理
- JSON 序列化
- 事件通知

### 用户态调试 (user_debug.rs)

- 进程附加/分离
- 用户态断点
- 异常处理

### 调用栈追踪 (callstack.rs)

- 栈回溯
- 符号解析
- 帧信息

## 状态

| 功能 | 状态 | 说明 |
|------|------|------|
| 调试器核心 | 🔄 | 开发中 |
| 命令处理 | 🔄 | 开发中 |
| 网络通信 | ✅ | 已完成 |
| 用户态调试 | 🔄 | 开发中 |
| 调用栈追踪 | 🔄 | 开发中 |

## 相关文档

- [hyperhv/README.md](hyperhv/README.md) - Hypervisor 实现
- [SKILL.md](../../../.trae/skills/hyperdbg-rust-driver/SKILL.md) - 开发技能指南
- [C++ 源码](../../../doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd) - 原始实现
