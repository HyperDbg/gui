# Rust 驱动实现进度

## 背景

经过 15 个空指针等基础问题导致的 BSOD 修复后，决定使用 Rust 重新实现 HyperDbg 驱动。Rust 的内存安全特性和编译时检查可以有效避免 C++ 中常见的空指针访问、缓冲区溢出等问题。

## 新架构设计

### 架构分层

```
┌─────────────────────────────────────────────────────────────┐
│                   Go 用户空间通信层                      │
│                   (packet.go + event_handlers.go)        │
└────────────────────┬────────────────────────────────────┘
                     │ HTTP/JSON (WSK Socket :50080)
┌────────────────────▼────────────────────────────────────┐
│              Rust 内核驱动层                            │
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

### 通信架构

- **WSK HTTP Server**：驱动监听 50080 端口，接收 HTTP/JSON 请求
- **协议层**：JSON 格式的请求/响应
- **API 调度器**：从 `handlers_gen/router.rs` 自动生成，根据 action 路由到对应方法

## 项目结构

```
rust-driver/
├── Cargo.toml              # 主驱动 crate 配置
├── lib.rs                  # 驱动入口 (DriverEntry)
├── build.rs                # 构建脚本
├── build.ps1               # PowerShell 构建脚本
│
├── common/                 # 通用模块（自动生成）
│   ├── src/
│   │   ├── types_gen/      # 生成的类型定义
│   │   │   ├── mod.rs
│   │   │   ├── common.rs       # 基础类型 (ProcessId, ThreadId, Address 等)
│   │   │   ├── register.rs    # RegisterState
│   │   │   ├── response.rs    # Response[T]
│   │   │   ├── event.rs       # DebuggerEvent
│   │   │   ├── event_*.rs     # 20+ 事件类型
│   │   │   └── mod.rs
│   │   │
│   │   ├── handlers_gen/  # 生成的处理器
│   │   │   ├── mod.rs         # 导出 + EventQueue + emit_*_event
│   │   │   ├── router.rs      # DebuggerApi trait + dispatch_api
│   │   │   └── emit.rs        # 事件发射函数
│   │   │
│   │   └── lib.rs
│   └── Cargo.toml
│
├── net/                    # 网络模块（WSK HTTP Server）
│   ├── src/
│   │   ├── lib.rs          # WSK Socket 服务器核心
│   │   ├── http.rs         # HTTP 请求处理
│   │   ├── json.rs         # JSON 编解码
│   │   └── util.rs         # Request 结构体和工具函数
│   ├── tests/              # 测试
│   ├── BSOD_FIX_REPORT.md
│   └── Cargo.toml
│
├── hyperhv/                # Hypervisor 层
│   ├── src/
│   │   ├── lib.rs          # #![no_std] 模块
│   │   ├── vmm/            # VMX/EPT 核心
│   │   │   ├── mod.rs
│   │   │   ├── vmx.rs
│   │   │   ├── vmcs.rs
│   │   │   ├── vmexit.rs
│   │   │   ├── ept.rs
│   │   │   └── sync.rs
│   │   ├── hooks.rs        # EPT Hook、Syscall Hook
│   │   ├── callbacks.rs    # 回调管理
│   │   ├── events.rs       # 事件系统
│   │   ├── processor.rs    # 处理器管理
│   │   ├── memory.rs       # 内存管理
│   │   ├── loader.rs       # 驱动加载
│   │   ├── apic.rs         # APIC 支持
│   │   ├── broadcast.rs    # 广播通信
│   │   ├── vmcall.rs       # VMCALL 处理
│   │   ├── interrupts.rs   # 中断处理
│   │   ├── idt.rs          # IDT 管理
│   │   ├── ept_vpid.rs     # EPT/VPID 支持
│   │   ├── smi.rs          # SMI 处理
│   │   ├── asm.rs          # 汇编代码
│   │   ├── globals.rs      # 全局变量
│   │   ├── scheduler.rs    # 调度器
│   │   ├── compatibility.rs # 兼容性检查
│   │   ├── switch_layout.rs # 栈切换
│   │   ├── halted_core.rs  # 停止核心
│   │   ├── tracing.rs      # 追踪
│   │   ├── meta_dispatch.rs # 元调度
│   │   ├── dirty_logging.rs # 脏页记录
│   │   ├── hyper_evade.rs  # 反检测
│   │   ├── transparency.rs # 透明性
│   │   ├── kernel_tests.rs # 内核测试
│   │   ├── lbr.rs          # LBR 支持
│   │   ├── debugger_asm.rs # 调试汇编
│   │   ├── serial_connection.rs # 串口连接
│   │   ├── termination.rs  # 终止处理
│   │   ├── communication.rs # 通信
│   │   ├── process.rs      # 进程管理
│   │   ├── thread.rs       # 线程管理
│   │   └── pci.rs          # PCI 支持
│   └── Cargo.toml
│
├── hyperkd/                # Debugger 层
│   ├── src/
│   │   ├── lib.rs          # #![no_std] 模块
│   │   ├── debugger.rs     # 调试器核心
│   │   ├── commands.rs     # 命令处理
│   │   ├── network.rs      # 网络通信
│   │   ├── user_debug.rs   # 用户态调试
│   │   ├── callstack.rs    # 调用栈
│   │   ├── allocations.rs  # 内存分配
│   │   ├── attaching.rs    # 进程附加
│   │   ├── dpc_routines.rs # DPC 例程
│   │   ├── driver.rs       # 驱动管理
│   │   ├── extension_commands.rs # 扩展命令
│   │   ├── ffi.rs          # FFI 接口
│   │   ├── halted_broadcast.rs # 停止广播
│   │   ├── process.rs      # 进程管理
│   │   ├── thread.rs       # 线程管理
│   │   ├── ud.rs           # 用户数据
│   │   └── user_access.rs  # 用户访问
│   └── Cargo.toml
│
├── kd/                     # 内核调试器
│   ├── src/
│   │   └── lib.rs          # KD 上下文和断点管理
│   └── Cargo.toml
│
├── pdbex/                  # PDB 解析器
│   ├── src/
│   │   └── lib.rs          # 符号解析、类型重建
│   └── Cargo.toml
│
├── disassembler/           # 反汇编器
│   ├── src/
│   │   └── lib.rs          # 使用 zydis-rs
│   └── Cargo.toml
│
├── driver-framework/       # 驱动框架库
│   ├── src/
│   │   ├── lib.rs
│   │   ├── device.rs       # 设备管理
│   │   ├── ioctl.rs        # IOCTL 处理
│   │   ├── ffi.rs          # FFI 接口
│   │   ├── logger.rs       # 日志宏
│   │   └── utils.rs        # 工具函数
│   ├── out/                # WDK 绑定（自动生成）
│   │   ├── ntddk.rs        # 内核函数
│   │   ├── types.rs        # 类型定义
│   │   └── constants.rs    # 常量定义
│   └── Cargo.toml
│
├── logger/                 # 日志库
│   ├── src/
│   │   └── lib.rs
│   └── Cargo.toml
│
├── examples/               # 示例驱动
│   ├── netdemo/            # 网络示例（当前测试用）
│   │   ├── netdemo.go      # Go 测试客户端
│   │   ├── src/lib.rs      # 驱动入口
│   │   └── build.ps1
│   ├── sysdemo/            # 系统驱动示例
│   ├── sample-wdm-driver/  # WDM 驱动示例
│   ├── sample-kmdf-driver/ # KMDF 驱动示例
│   └── sample-umdf-driver/ # UMDF 驱动示例
│
└── doc/                    # 文档
    ├── RUST_INSTALL_README.md
    ├── bug_fixes.md
    ├── install-rust-mingw.ps1
    └── setup_ewdk_env.ps1
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
| **common/types_gen** | `common/src/types_gen/` | ✅ 100% | 自动生成的类型定义 |
| **common/handlers_gen** | `common/src/handlers_gen/` | ✅ 100% | 自动生成的 API 调度器 |
| **net** | `net/src/` | ✅ 100% | WSK HTTP Server |
| **driver-framework** | `driver-framework/src/` | ✅ 100% | WDM/WDF 框架 |
| **hyperhv** | `hyperhv/src/` | 🔄 开发中 | Hypervisor 核心（已有框架代码） |
| **hyperkd** | `hyperkd/src/` | 🔄 开发中 | Debugger 核心（已有框架代码） |
| **kd** | `kd/src/` | 🔄 开发中 | 内核调试器（已有框架代码） |
| **pdbex** | `pdbex/src/` | 🔄 开发中 | PDB 解析器（已有框架代码） |
| **disassembler** | `disassembler/src/` | 🔄 开发中 | 反汇编器 |

### types_gen 模块（已自动生成）

| 文件 | 内容 | 进度 |
|------|------|------|
| `common.rs` | ProcessId, ThreadId, Address, PhysicalAddress | ✅ |
| `register.rs` | RegisterState | ✅ |
| `response.rs` | Response[T], Empty | ✅ |
| `event.rs` | DebuggerEvent, DebuggerEventType | ✅ |
| `event_breakpoint.rs` | BreakpointType, BreakpointEvent | ✅ |
| `event_exception.rs` | ExceptionCode, ExceptionEvent | ✅ |
| `event_memory.rs` | MemoryAccessType, MemoryAccessEvent | ✅ |
| `event_syscall.rs` | SyscallEvent | ✅ |
| `event_process.rs` | ProcessInfo, ProcessEvent | ✅ |
| `event_thread.rs` | ThreadInfo, ThreadEvent | ✅ |
| `event_module.rs` | ModuleInfo, ModuleEvent | ✅ |
| `event_debug_print.rs` | LogLevel, DebugPrintEvent | ✅ |
| `event_vmx.rs` | VmxExitReason, VmxExitEvent | ✅ |
| `event_trap.rs` | TrapEvent | ✅ |
| `event_hidden_hook.rs` | HiddenHookEvent | ✅ |
| `event_cpuid.rs` | CpuidEvent | ✅ |
| `event_tsc.rs` | TscEvent | ✅ |
| `event_cr_access.rs` | CrAccessEvent | ✅ |
| `event_dr_access.rs` | DrAccessEvent | ✅ |
| `event_io_port.rs` | IoPortEvent | ✅ |
| `event_msr.rs` | MsrEvent | ✅ |
| `event_ept_violation.rs` | EptViolationEvent | ✅ |

### handlers_gen 模块（已自动生成）

| 文件 | 内容 | 进度 |
|------|------|------|
| `router.rs` | DebuggerApi trait + dispatch_api | ✅ 100% |
| `emit.rs` | emit_*_event 函数 | ✅ 100% |
| `mod.rs` | 模块导出 | ✅ 100% |

### net 模块（WSK HTTP Server）

| 功能 | 状态 | 说明 |
|------|------|------|
| WSK Socket 监听 | ✅ 100% | 监听 50080 端口 |
| HTTP 请求解析 | ✅ 100% | POST /api, /ping, /status |
| JSON 序列化 | ✅ 100% | Response::WriteJSON |
| JSON 反序列化 | ✅ 100% | Request::ReadJSON |
| Request.URL() | ✅ 100% | 动态构建完整 URL |

### hyperhv 模块（Hypervisor 层）

| 功能 | 状态 | C++ 参考路径 |
|------|------|-------------|
| VMX 操作 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/code/vmm/` |
| EPT 管理 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/code/vmm/` |
| VM Exit 处理 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/code/vmm/` |
| Hook 机制 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/code/hooks/` |
| 事件系统 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/` |
| 处理器管理 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/` |
| APIC 支持 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/` |
| 广播通信 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/` |
| SMI 处理 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv/` |

**当前 hyperhv 模块文件列表：**
- `vmm/` - VMX/EPT 核心 (vmx.rs, vmcs.rs, vmexit.rs, ept.rs, sync.rs)
- `hooks.rs` - Hook 机制
- `events.rs` - 事件系统
- `processor.rs` - 处理器管理
- `memory.rs` - 内存管理
- `loader.rs` - 驱动加载
- `apic.rs` - APIC 支持
- `broadcast.rs` - 广播通信
- `vmcall.rs` - VMCALL 处理
- `interrupts.rs` - 中断处理
- `idt.rs` - IDT 管理
- `ept_vpid.rs` - EPT/VPID 支持
- `smi.rs` - SMI 处理
- `asm.rs` - 汇编代码
- `globals.rs` - 全局变量
- `scheduler.rs` - 调度器
- `compatibility.rs` - 兼容性检查
- `switch_layout.rs` - 栈切换
- `halted_core.rs` - 停止核心
- `tracing.rs` - 追踪
- `meta_dispatch.rs` - 元调度
- `dirty_logging.rs` - 脏页记录
- `hyper_evade.rs` - 反检测
- `transparency.rs` - 透明性
- `kernel_tests.rs` - 内核测试
- `lbr.rs` - LBR 支持
- `debugger_asm.rs` - 调试汇编
- `serial_connection.rs` - 串口连接
- `termination.rs` - 终止处理
- `communication.rs` - 通信
- `process.rs` - 进程管理
- `thread.rs` - 线程管理
- `pci.rs` - PCI 支持
- `callbacks.rs` - 回调管理

### hyperkd 模块（Debugger 层）

| 功能 | 状态 | C++ 参考路径 |
|------|------|-------------|
| 调试器核心 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/code/debugger/` |
| 命令处理 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/` |
| 网络通信 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/` |
| 用户态调试 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/` |
| 调用栈 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/` |
| 进程附加 | 🔄 开发中 | `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd/` |

**当前 hyperkd 模块文件列表：**
- `debugger.rs` - 调试器核心
- `commands.rs` - 命令处理
- `network.rs` - 网络通信
- `user_debug.rs` - 用户态调试
- `callstack.rs` - 调用栈
- `allocations.rs` - 内存分配
- `attaching.rs` - 进程附加
- `dpc_routines.rs` - DPC 例程
- `driver.rs` - 驱动管理
- `extension_commands.rs` - 扩展命令
- `ffi.rs` - FFI 接口
- `halted_broadcast.rs` - 停止广播
- `process.rs` - 进程管理
- `thread.rs` - 线程管理
- `ud.rs` - 用户数据
- `user_access.rs` - 用户访问

### kd 模块（内核调试器）

| 功能 | 状态 | 说明 |
|------|------|------|
| KD 上下文 | 🔄 开发中 | KdContext 结构体 |
| 断点管理 | 🔄 开发中 | Breakpoint 结构体 |
| 状态管理 | 🔄 开发中 | KdState 枚举 |

### pdbex 模块（PDB 解析器）

| 功能 | 状态 | 说明 |
|------|------|------|
| 符号类型 | 🔄 开发中 | SymTag 枚举 |
| UDT 类型 | 🔄 开发中 | UdtKind 枚举 |
| 基础类型 | 🔄 开发中 | BasicType 枚举 |

## 主驱动入口 (lib.rs)

主驱动入口文件 `lib.rs` 负责初始化 HTTP 服务器和事件队列：

```rust
#![no_std]
#![allow(non_snake_case)]

extern crate alloc;
extern crate wdk_panic;

use net::{ResponseWriter, Request, Server};
use types_gen::*;
use handlers_gen::{DebuggerApi, dispatch_api, NoOpDebugger, EventQueue};

static mut GLOBAL_SERVER: *mut Server = core::ptr::null_mut();
static mut EVENT_QUEUE: *mut EventQueue = core::ptr::null_mut();

#[no_mangle]
pub unsafe extern "system" fn DriverEntry(
    driver_object: *mut DRIVER_OBJECT,
    _registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    // 初始化事件队列
    let queue = alloc::boxed::Box::new(EventQueue::new(1000));
    EVENT_QUEUE = alloc::boxed::Box::into_raw(queue);
    
    // 创建 HTTP 服务器
    let server = Server::NewServer();
    GLOBAL_SERVER = server;
    
    // 注册 API 处理器
    (*server).Handle("/api", api_handler);
    
    // 启动监听
    (*server).ListenAndServe(":50080")
}
```

## API 调度流程

```
Go 客户端
    │
    │ HTTP POST /api/initialize
    │ { "action": "initialize", ... }
    ▼
rust-driver/net (WSK HTTP Server)
    │
    │ readRequest() 解析 HTTP 请求
    ▼
Request { Method, URL, Host, Port, Body, ... }
    │
    │ ServeHTTP() 路由
    ▼
Router.match_pattern() 匹配路由
    │
    ▼
dispatch_api(api, w, r)
    │
    │ 根据 action 调用
    ▼
DebuggerApi trait 方法
    │
    │ initialize(&mut self, w, r)
    ▼
Result<Response, Error> → JSON
```

## 生成的代码

### DebuggerApi trait (handlers_gen/router.rs)

```rust
pub trait DebuggerApi {
    fn initialize(&mut self, req: &Request) -> Result<Empty, String>;
    fn terminate(&mut self, req: &Request) -> Result<Empty, String>;
    fn attach_process(&mut self, req: &Request) -> Result<Empty, String>;
    fn read_memory(&mut self, req: &Request) -> Result<Vec<u8>, String>;
    // ... 更多方法
}
```

### dispatch_api 函数 (handlers_gen/router.rs)

```rust
pub fn dispatch_api<T: DebuggerApi>(api: &mut T, body: &[u8]) -> Vec<u8> {
    let req: Result<Request, _> = serde_json::from_slice(body);
    match req {
        Ok(req) => {
            match req.action.as_str() {
                "initialize" => api.initialize(&req),
                "terminate" => api.terminate(&req),
                "attach_process" => api.attach_process(&req),
                // ... 更多匹配
                _ => error_response("unknown action"),
            }
        }
        Err(e) => error_response(&e.to_string()),
    }
}
```

## 下一步计划

### 主要任务

#### 1. hyperhv 模块 - 完善 Hypervisor 实现

**路径：** `rust-driver/hyperhv`

**C++ 参考：** `doc/cpp/HyperDbgUnified/HyperDbg/hprdbghv`

- 完善 VMX/EPT 核心实现
- 完善 Hook 机制
- 完善事件系统

#### 2. hyperkd 模块 - 完善 Debugger 实现

**路径：** `rust-driver/hyperkd`

**C++ 参考：** `doc/cpp/HyperDbgUnified/HyperDbg/hyperkd`

- 完善调试器核心
- 完善命令处理
- 完善网络通信

#### 3. kd 模块 - 完善内核调试器

**路径：** `rust-driver/kd`

- 完善断点管理
- 完善状态管理

#### 4. pdbex 模块 - 完善 PDB 解析器

**路径：** `rust-driver/pdbex`

- 完善符号解析
- 完善类型重建

### 最终目标

更新 `cmd/rustgen/main.go` 生成器，自动生成完整的 API 调度代码，与 hyperhv/hyperkd/kd/pdbex 对接。

## 关键差异

### 架构差异

| 方面 | C++ 原始实现 | Rust 新实现 |
|------|--------------|-------------|
| **通信方式** | DeviceIoControl + 串口 | WSK Socket HTTP |
| **命令处理** | IOCTL 命令 | HTTP/JSON + Rust 方法 |
| **类型生成** | 手动维护 | 自动生成 (rustgen) |

### 生成器流程

```
cmd/rustgen/main.go
    │
    │ 扫描 Go 类型定义
    ▼
types.go (Go)
    │
    │ 自动生成
    ▼
common/src/types_gen/*.rs (Rust)
    │
    │ 自动生成
    ▼
common/src/handlers_gen/router.rs (Rust)
    │
    ▼
hyperhv/hyperkd 实现 DebuggerApi trait
```

## 里程碑

### 2024-03-26: WSK HTTP Server 完成

- 驱动监听 50080 端口
- HTTP 请求解析
- JSON 序列化/反序列化
- 多客户端连接

### 2024-03-27: Request 结构体重构

- 添加 Host 和 Port 字段
- 实现 URL() 方法
- 动态构建完整 URL

### 2024-03-27: 模块结构完善

- 添加 kd 模块（内核调试器）
- 添加 pdbex 模块（PDB 解析器）
- 添加 logger 模块
- 完善项目结构

### 2025-03-27: hyperhv/hyperkd 模块完善

- hyperhv 模块包含完整的框架代码（30+ 文件）
- hyperkd 模块包含完整的框架代码（15+ 文件）
- 主驱动 lib.rs 实现完整的 HTTP 服务器入口

---

## 相关文档

- [api_design.md](../api_design.md) - API 设计详细文档
- [README.md](../README.md) - Go Communicator 文档
- [.trae/skills/rust-driver-development/SKILL.md](../../.trae/skills/rust-driver-development/SKILL.md) - 开发技能
