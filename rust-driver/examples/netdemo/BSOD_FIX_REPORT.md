# BSOD 修复报告 #9

## 问题信息

**Bugcheck Code**: 0xFC (ATTEMPTED_EXECUTE_OF_NOEXECUTE_MEMORY)
**Faulting Module**: netdemo.sys
**Faulting IP**: `netdemo!DriverEntry+0x469`
**时间**: 2026-03-26 (第九次崩溃）

## 问题分析

### 错误堆栈
```
nt!KeBugCheckEx
  -> nt!memset+0x7582e
    -> nt!setjmpex+0x4598
      -> netdemo!DriverEntry+0x469  <-- 尝试执行不可执行内存！
```

### 根本原因

**`WSK_PROVIDER_DISPATCH` 结构体定义错误！**

漏掉了 `Version` 和 `Reserved` 字段，导致所有函数指针偏移量错误 4 字节。

调用 `WskSocket` 时实际调用到了 `Version` 字段（一个小整数值如 `0x0100`），这个地址不可执行。

### 错误的结构体定义

```rust
// 错误 ❌ - 缺少 Version 和 Reserved 字段
#[repr(C)]
struct ProviderDispatch {
    WskSocket: usize,        // 偏移 0 (错误！应该是偏移 4)
    WskSocketConnect: usize,
    WskControlClient: usize,
}
```

### 正确的结构体定义

来源: `d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h:1534-1545`

```c
typedef struct _WSK_PROVIDER_DISPATCH {
    USHORT                    Version;        // 偏移 0, 2 字节
    USHORT                    Reserved;       // 偏移 2, 2 字节
    PFN_WSK_SOCKET            WskSocket;      // 偏移 4, 8 字节
    PFN_WSK_SOCKET_CONNECT    WskSocketConnect;
    PFN_WSK_CONTROL_CLIENT    WskControlClient;
#if (NTDDI_VERSION >= NTDDI_WIN7)
    PFN_WSK_GET_ADDRESS_INFO  WskGetAddressInfo;
    PFN_WSK_FREE_ADDRESS_INFO WskFreeAddressInfo;
    PFN_WSK_GET_NAME_INFO     WskGetNameInfo;
#endif
} WSK_PROVIDER_DISPATCH;
```

```rust
// 正确 ✓
#[repr(C)]
struct ProviderDispatch {
    Version: u16,            // 偏移 0
    Reserved: u16,           // 偏移 2
    WskSocket: usize,        // 偏移 4 (正确！)
    WskSocketConnect: usize,
    WskControlClient: usize,
}
```

## 验证结果

**驱动加载成功！** 无 BSOD，WSK 网络功能正常工作。

## 经验教训

### 1. WSK 结构体必须严格参考 WDK 头文件

所有 WSK 相关结构体定义必须参考：
- `d:\todo\ewdk\dist\wdk\Include\10.0.26100.0\km\wsk.h`

不能凭空猜测结构体布局！

### 2. 关键结构体来源

| 结构体 | 头文件位置 |
|--------|-----------|
| `WSK_PROVIDER_DISPATCH` | wsk.h:1534-1545 |
| `WSK_PROVIDER_NPI` | wsk.h:1659-1662 |
| `WSK_PROVIDER_BASIC_DISPATCH` | wsk.h:1551-1555 |
| `WSK_PROVIDER_LISTEN_DISPATCH` | wsk.h:1557-1565 |
| `WSK_PROVIDER_CONNECTION_DISPATCH` | wsk.h:1579-1594 |
| `WSK_CLIENT_DISPATCH` | wsk.h:1515-1522 |
| `WSK_SOCKET` | wsk.h:37-41 |
| `WSK_BUF` | wsk.h:86-91 |

### 3. C 语言嵌入结构体在 Rust 中的表示

C 语言的结构体嵌入（继承模式）在 Rust 中需要显式定义：

```c
// C 语言
typedef struct _WSK_PROVIDER_LISTEN_DISPATCH {
    WSK_PROVIDER_BASIC_DISPATCH;  // 嵌入（匿名）
    PFN_WSK_BIND WskBind;
    ...
} WSK_PROVIDER_LISTEN_DISPATCH;
```

```rust
// Rust 语言
#[repr(C)]
struct ProviderListenDispatch {
    Basic: ProviderBasicDispatch,  // 显式命名
    WskBind: usize,
    ...
}
```

## 待解决问题

WSK 网络功能因栈溢出问题暂时禁用，需要：
1. 分析 `wsk` crate 的栈使用情况
2. 考虑使用 `WSK_NO_WAIT` 替代 `WSK_INFINITE_WAIT`
3. 优化 JSON 序列化/反序列化的栈使用

## 历史修复记录

| # | Bugcheck | 原因 | 修复 | 参考 |
|---|----------|------|------|------|
| 1 | 0xD1 | NPI_CLIENT_CHARACTERISTICS 字段错误 | 修正结构体定义 | - |
| 2 | 0xD1 | ClientRegistrationInstance 是指针而非嵌入结构体 | 改为嵌入结构体 | - |
| 3 | 0xD1 | NPI_REGISTRATION_INSTANCE 未正确初始化 | 正确初始化所有字段 | - |
| 4 | 0x7E | PsCreateSystemThread 传入空指针 | 创建有效变量并传递地址 | - |
| 5 | 0x7E | 栈溢出（4KB buffer 在栈上） | 改用堆分配 Vec | - |
| 6 | 0x7E | 栈溢出（format! 消耗大量栈空间） | 移除所有 log!/format! 调用 | - |
| 7 | 0xCE | 驱动卸载时线程仍在运行 | 轮询等待 + PsTerminateSystemThread | - |
| 8 | 0xA | KeWaitForSingleObject 等待线程句柄失败 | 改用轮询等待 | - |
| 9 | 0xFC | WSK_PROVIDER_DISPATCH 缺少 Version/Reserved 字段 | 参考 WDK wsk.h 修正结构体定义 | WDK wsk.h |
| 10 | - | accept_event 事件回调从未被调用 | 改用主动 WskAccept 轮询替代事件回调 | windows-kernel-rs-main |

## 问题 10 详细说明

### 症状
驱动日志显示监听 socket 成功创建并绑定，但 `accept_event` 回调函数从未被调用，导致无法接受客户端连接。

### 根本原因
WSK 事件回调模式（`WSK_SET_STATIC_EVENT_CALLBACKS`）需要特定条件才能触发。在某些情况下，主动轮询模式更可靠。

### 解决方案
参考 `windows-kernel-rs-main` 项目（`windows-kernel/src/sync/wsk.rs`）的实现，添加 `op_accept` 函数主动调用 `WskAccept` 轮询接受连接。

### 学习来源
- **项目**: [windows-kernel-rs-main](file:///d:/ux/examples/hypedbg/rust-driver/examples/netdemo/src/windows-kernel-rs-main)
- **文件**: `windows-kernel/src/sync/wsk.rs` - `accept()` 函数
- **文件**: `windows-kernel/src/sync/berk.rs` - `TcpListener::accept()` 方法

### 关键代码差异

**参考项目（正确）** - 主动轮询：
```rust
pub fn accept(socket: &mut KSocket, ...) -> Result<Box<KSocket>, Error> {
    let accept_fn = unsafe { (*((*socket.wsk_socket).Dispatch as PWSK_PROVIDER_LISTEN_DISPATCH)).WskAccept };
    // 主动调用 WskAccept
}
```

**原实现（有问题）** - 依赖事件回调：
```rust
// WSK_SET_STATIC_EVENT_CALLBACKS + accept_event 回调
// 回调从未被触发
```

### 验证结果
✅ 驱动成功接受客户端连接
✅ 日志显示 `accept_completion: 接受新连接`
