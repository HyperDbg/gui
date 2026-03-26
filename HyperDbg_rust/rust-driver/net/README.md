# WSK 模块文档

## 概述

WSK (Windows Sockets Kernel) 模块是 HyperDbg 项目中用于内核级网络通信的核心库。该模块封装了 Windows 内核网络接口，提供了一套完整的 TCP/IP 通信功能，可在多个驱动程序中共享使用。

## 实现过程

### 1. 参考项目分析

本实现参考了 `windows-kernel-rs-main` 项目（位于 `d:\ux\examples\hypedbg\rust-driver\examples\netdemo\src\windows-kernel-rs-main`），该项目的 WSK 实现展示了在 Rust 中使用 WSK API 的正确方式。

### 2. 关键技术决策

#### 2.1 主动轮询模式 vs 事件回调模式

**问题**：最初实现使用 WSK 事件回调模式（`WSK_SET_STATIC_EVENT_CALLBACKS`），但 `accept_event` 回调从未被调用，导致无法接受客户端连接。

**解决方案**：参考 `windows-kernel-rs-main` 项目，改用主动轮询模式。通过 `op_accept` 函数主动调用 `WskAccept` 来接受连接，配合 `accept_completion` 回调处理异步结果。

**参考代码**：`windows-kernel/src/sync/wsk.rs` 中的主动轮询实现。

#### 2.2 结构体字段对齐

**问题**：`WSK_PROVIDER_DISPATCH` 结构体最初缺少 `Version` 和 `Reserved` 字段，导致所有函数指针偏移错误 4 字节，触发 BSOD (Bugcheck 0xFC: ATTEMPTED_EXECUTE_OF_NOEXECUTE_MEMORY)。

**解决方案**：严格按照 WDK 头文件定义结构体，添加 `Version: u16` 和 `Reserved: u16` 字段。

**参考**：`d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1534-1545`

### 3. 架构设计

#### 3.1 模块结构

```
wsk/
├── src/
│   ├── lib.rs          # 核心 WSK 实现
│   └── logger.rs       # 日志宏
└── Cargo.toml
```

#### 3.2 核心组件

1. **WSK 注册和初始化**
   - `WskRegister`: 注册 WSK 客户端
   - `WskCaptureProviderNPI`: 获取 WSK 提供者接口
   - `WskReleaseProviderNPI`: 释放 WSK 提供者接口
   - `WskDeregister`: 注销 WSK 客户端

2. **Socket 管理**
   - `Socket`: WSK socket 对象
   - `SocketContext`: Socket 上下文，包含工作队列和操作上下文
   - `SocketOpContext`: 单个操作的上下文（IRP、缓冲区等）

3. **工作队列**
   - `WorkQueue`: 基于 `SLIST_HEADER` 的无锁队列
   - `worker_thread`: 工作线程，处理队列中的操作
   - `enqueue_op`: 将操作入队

4. **操作处理**
   - `op_start_listen`: 启动监听
   - `op_accept`: 主动接受连接
   - `op_receive`: 接收数据
   - `op_send`: 发送数据
   - `op_disconnect`: 断开连接
   - `op_close`: 关闭 socket
   - `op_free`: 释放资源

5. **完成回调**
   - `accept_completion`: 连接接受完成
   - `receive_completion`: 接收完成
   - `send_completion`: 发送完成
   - `disconnect_completion`: 断开完成
   - `sync_completion_routine`: 同步完成

## 暴露的接口

### 公共常量

```rust
pub const AF_INET6: u16 = 23;
pub const SOCK_STREAM: u16 = 1;
pub const IPPROTO_TCP: u32 = 6;
pub const IPPROTO_IPV6: u32 = 41;
pub const IPV6_V6ONLY: u32 = 27;
pub const WSK_SET_STATIC_EVENT_CALLBACKS: u32 = 7;
pub const WSK_EVENT_ACCEPT: u32 = 0x00000080;
pub const WSK_FLAG_LISTEN_SOCKET: u32 = 0x00000001;
pub const WSK_INFINITE_WAIT: ULONG = 0xFFFFFFFF;
pub const WSK_NO_WAIT: ULONG = 0;
pub const WSK_CAPTURE_TIMEOUT_MS: i64 = 1000;
pub const KERNEL_MODE: KPROCESSOR_MODE = 0;
pub const WskSetOption: i32 = 0;
pub const WskGetOption: i32 = 1;
pub const POOL_TAG: u32 = 0x736B7377;
pub const DATA_BUFFER_LENGTH: usize = 2048;
pub const OP_COUNT: usize = 2;
```

### 公共结构体

#### WSK 接口结构体

```rust
pub struct ClientDispatch;
pub struct ClientNpi;
pub struct ProviderNpi;
pub struct ProviderDispatch;
pub struct Registration;
```

#### Socket 相关结构体

```rust
pub struct Socket;
pub struct WskBuf;
pub struct EventCallbackControl;
pub struct ProviderBasicDispatch;
pub struct ProviderListenDispatch;
pub struct ProviderConnectionDispatch;
pub struct ClientListenDispatch;
pub struct ClientConnectionDispatch;
pub struct SockaddrIn6;
pub type SOCKADDR = SockaddrIn6;
```

#### 工作队列结构体

```rust
pub struct WorkQueue {
    pub Head: _SLIST_HEADER,
    pub Event: KEVENT,
    pub Stop: BOOLEAN,
    pub Thread: PVOID,
}

pub struct SocketOpContext {
    pub QueueEntry: _SLIST_ENTRY,
    pub OpHandler: Option<OpHandlerFn>,
    pub SocketContext: *mut SocketContext,
    pub Irp: PIRP,
    pub DataBuffer: PVOID,
    pub DataMdl: PMDL,
    pub BufferLength: usize,
    pub DataLength: usize,
}

pub struct SocketContext {
    pub Socket: *mut Socket,
    pub WorkQueue: *mut WorkQueue,
    pub Closing: BOOLEAN,
    pub Disconnecting: BOOLEAN,
    pub StopListening: BOOLEAN,
    pub OpContext: [SocketOpContext; OP_COUNT],
}
```

### 公共函数

#### WSK 外部函数

```rust
extern "system" {
    pub fn WskRegister(
        WskClientNpi: *const ClientNpi,
        WskRegistration: *mut Registration,
    ) -> NTSTATUS;

    pub fn WskCaptureProviderNPI(
        WskRegistration: *mut Registration,
        WaitTimeout: ULONG,
        WskProviderNpi: *mut ProviderNpi,
    ) -> NTSTATUS;

    pub fn WskReleaseProviderNPI(
        WskRegistration: *mut Registration,
    );

    pub fn WskDeregister(
        WskRegistration: *mut Registration,
    );
    
    pub static NPI_WSK_INTERFACE_ID: [u8; 16];
}
```

#### 辅助函数

```rust
pub unsafe fn set_completion_routine(
    irp: PIRP,
    routine: PIO_COMPLETION_ROUTINE,
    context: PVOID
);

pub unsafe fn allocate_socket_context(
    work_queue: *mut WorkQueue,
    buffer_length: ULONG,
) -> *mut SocketContext;

pub unsafe fn free_socket_context(
    socket_context: *mut SocketContext
);

pub unsafe fn enqueue_op(
    op_ctx: *mut SocketOpContext,
    handler: OpHandlerFn
);
```

### 公共静态变量

```rust
pub static mut WSK_REGISTRATION: Registration;
pub static mut WSK_CLIENT_NPI: ClientNpi;
pub static mut WORK_QUEUE: WorkQueue;
pub static mut LISTENING_SOCKET_CONTEXT: *mut SocketContext;
pub static mut LISTENING_ADDRESS: SockaddrIn6;
```

### 日志宏

```rust
log!(format_args!);
log_success!(format_args!);
log_error!(format_args!);
log_info!(format_args!);
log_warn!(format_args!);
```

## 使用示例

### 在驱动中使用 WSK 模块

1. 在 `Cargo.toml` 中添加依赖：

```toml
[dependencies]
wsk = { workspace = true }
```

2. 在驱动代码中导入：

```rust
use wsk::*;
```

3. 实现 DriverEntry：

```rust
#[no_mangle]
pub extern "system" fn DriverEntry(
    driver: *mut DRIVER_OBJECT,
    _registry_path: *const UNICODE_STRING,
) -> NTSTATUS {
    log_info!("DriverEntry: 开始");
    
    let socket_context = unsafe { 
        allocate_socket_context(&mut WORK_QUEUE as *mut _, DATA_BUFFER_LENGTH as ULONG) 
    };
    if socket_context.is_null() {
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    
    unsafe { LISTENING_SOCKET_CONTEXT = socket_context };
    unsafe { WSK_CLIENT_NPI.Dispatch = &WSK_CLIENT_DISPATCH };
    
    let status = unsafe { WskRegister(&WSK_CLIENT_NPI, &mut WSK_REGISTRATION) };
    if status != STATUS_SUCCESS {
        unsafe { ExFreePool(socket_context as PVOID) };
        return status;
    }
    
    let status = unsafe { start_work_queue(&mut WORK_QUEUE) };
    if status != STATUS_SUCCESS {
        unsafe { 
            WskDeregister(&mut WSK_REGISTRATION);
            ExFreePool(socket_context as PVOID);
        };
        return status;
    }
    
    unsafe { enqueue_op(&mut (*socket_context).OpContext[0], op_start_listen) };
    unsafe { (*driver).DriverUnload = Some(driver_unload) };
    
    STATUS_SUCCESS
}
```

## 依赖关系

### 外部依赖

- `wdk-alloc`: Windows 内核内存分配器
- `wdk-sys`: Windows 内核系统类型和函数
- `wdk`: Windows 驱动开发工具包
- `wdk-panic`: 内核 panic 处理

### 内部依赖

- `logger`: 日志宏模块

## 已知问题和限制

1. **仅支持 IPv6 TCP**: 当前实现仅支持 IPv6 TCP 流式 socket
2. **单一监听端口**: 默认监听端口 0x479c (18364)
3. **固定缓冲区大小**: 数据缓冲区大小固定为 2048 字节
4. **同步操作**: 某些操作使用同步等待，可能影响性能

## 未来改进方向

1. 支持更多协议族（IPv4、UDP）
2. 可配置的监听端口和缓冲区大小
3. 完全异步操作，避免阻塞等待
4. 更完善的错误处理和恢复机制
5. 连接池管理
6. 性能优化和资源限制

## 参考资料

1. Windows Driver Kit (WDK) 文档
2. `d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h`
3. `windows-kernel-rs-main` 项目
4. Microsoft WSK 示例代码 `wsksmple.c`

## 版本历史

- v0.1.0: 初始版本，支持基本的 IPv6 TCP 通信
  - 实现主动轮询模式接受连接
  - 修复结构体对齐问题
  - 提供完整的 send/receive/disconnect 流程
