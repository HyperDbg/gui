# Rust Windows 驱动开发指南

## ⚠️ 重要警告

**严禁直接使用 `cargo build` 命令！**

- 本项目使用 **EWDK (Enterprise Windows Driver Kit)** 进行编译
- **永远不会安装 VS2022**
- 只能使用 PowerShell 构建脚本进行编译
- 当用户要求执行某个 `.ps1` 脚本时，**必须直接执行该脚本**，不要尝试手动运行 cargo 命令

## ⛔ 严禁查看未验证模块代码

**执行任务时，严禁查看以下未验证模块的代码：**
- `hyperhv` - Hypervisor 核心
- `hyperkd` - 内核调试器
- `disassembler` - 反汇编器
- `pdbex` - PDB 解析
- `kd` - 内核调试支持

**只能参考已验证可编译的项目：**
- `driver-framework` - 通用驱动框架库
- `sysdemo` - 模块化 WDM 驱动示例
- `netdemo` - 网络通信测试驱动
- `protocol` - 调试器通信协议定义
- `wsk` - WSK 网络封装

**WSK 实现参考：**
- `todo/代码+课件/DriverUtil/SysTemp17/WskSocket.cpp` - C++ WSK 完整实现
- `todo/代码+课件/DriverUtil/SysTemp17/WskSocket.h` - C++ WSK 接口定义

## 🔴 强制阅读 WDK-SYS 绑定文件

**写任何模块之前，必须先阅读以下文件，确认函数签名和类型定义：**

```
d:\ux\examples\hypedbg\rust-driver\driver-framework\out\ntddk.rs    # 所有内核函数
d:\ux\examples\hypedbg\rust-driver\driver-framework\out\types.rs    # 所有类型定义
d:\ux\examples\hypedbg\rust-driver\driver-framework\out\constants.rs # 所有常量
```

**这些是微软的绑定，虽然是每次编译动态绑定，但是没有这些文件你就会胡乱发明创造，写出来的代码基本都是 BSOD！**

## WDK-SYS 关键函数签名

### IRP 相关函数 (ntddk.rs)

```rust
pub fn IoAllocateIrp(StackSize: CCHAR, ChargeQuota: BOOLEAN) -> PIRP;
pub fn IoFreeIrp(Irp: PIRP);
pub fn IoAllocateMdl(
    VirtualAddress: PVOID,
    Length: ULONG,
    SecondaryBuffer: BOOLEAN,
    ChargeQuota: BOOLEAN,
    Irp: PVOID,
) -> PMDL;
pub fn IoFreeMdl(Mdl: PMDL);
```

### 事件相关函数 (ntddk.rs)

```rust
pub fn KeInitializeEvent(Event: PRKEVENT, Type: EVENT_TYPE, State: BOOLEAN);
pub fn KeSetEvent(Event: PRKEVENT, Increment: KPRIORITY, Wait: BOOLEAN) -> LONG;
pub fn KeWaitForSingleObject(
    Object: PVOID,
    WaitReason: KWAIT_REASON,
    WaitMode: KPROCESSOR_MODE,
    Alertable: BOOLEAN,
    Timeout: PLARGE_INTEGER,
) -> NTSTATUS;
```

### 内存相关函数 (ntddk.rs)

```rust
pub fn MmProbeAndLockPages(
    MemoryDescriptorList: PMDL,
    AccessMode: KPROCESSOR_MODE,
    Operation: LOCK_OPERATION,
);
pub fn MmUnlockPages(MemoryDescriptorList: PMDL);
```

### 设备相关函数 (ntddk.rs)

```rust
pub fn IoCreateDevice(
    DriverObject: PDRIVER_OBJECT,
    DeviceExtensionSize: ULONG,
    DeviceName: PUNICODE_STRING,
    DeviceType: DEVICE_TYPE,
    DeviceCharacteristics: ULONG,
    Exclusive: BOOLEAN,
    DeviceObject: *mut PDEVICE_OBJECT,
) -> NTSTATUS;
pub fn IoDeleteDevice(DeviceObject: PDEVICE_OBJECT);
pub fn IoCreateSymbolicLink(SymbolicLinkName: PUNICODE_STRING, DeviceName: PUNICODE_STRING) -> NTSTATUS;
pub fn IoDeleteSymbolicLink(SymbolicLinkName: PUNICODE_STRING) -> NTSTATUS;
```

## WDK-SYS 关键类型定义

### 基本类型 (types.rs)

```rust
pub type NTSTATUS = LONG;
pub type PVOID = *mut ::core::ffi::c_void;
pub type PIRP = *mut IRP;
pub type PMDL = *mut _MDL;
pub type PRKEVENT = *mut _KEVENT;
pub type BOOLEAN = u8;  // 实际是 ::core::ffi::c_uchar
pub type CCHAR = ::core::ffi::c_char;  // 有符号 8 位
pub type UCHAR = ::core::ffi::c_uchar;  // 无符号 8 位
pub type ULONG = ::core::ffi::c_ulong;  // 无符号 32 位
pub type LONG = ::core::ffi::c_long;    // 有符号 32 位
pub type ULONG_PTR = ::core::ffi::c_ulonglong;  // 64 位指针大小
pub type SIZE_T = ULONG_PTR;
pub type KPROCESSOR_MODE = CCHAR;  // 内核模式 = 0, 用户模式 = 1
```

### 指针类型 (types.rs)

```rust
pub type PDRIVER_OBJECT = *mut _DRIVER_OBJECT;
pub type PDEVICE_OBJECT = *mut _DEVICE_OBJECT;
pub type PKTHREAD = *mut _KTHREAD;
pub type PEPROCESS = *mut _KPROCESS;
pub type PKIRQL = *mut KIRQL;
```

### 枚举类型 (types.rs) - 重要！

**枚举类型是 `i32` 类型别名，必须通过模块路径访问常量！**

```rust
// 错误用法 ❌
use wdk_sys::{EVENT_TYPE, KWAIT_REASON, LOCK_OPERATION};
EVENT_TYPE::SynchronizationEvent  // 编译错误！EVENT_TYPE 是 i32

// 正确用法 ✅
use wdk_sys::_EVENT_TYPE;         // 导入模块
_EVENT_TYPE::SynchronizationEvent  // 正确！

use wdk_sys::_KWAIT_REASON;
_KWAIT_REASON::Executive           // 正确！

use wdk_sys::_LOCK_OPERATION;
_LOCK_OPERATION::IoReadAccess      // 正确！
```

**枚举模块定义：**
```rust
pub mod _EVENT_TYPE {
    pub type Type = ::core::ffi::c_int;
    pub const NotificationEvent: Type = 0;
    pub const SynchronizationEvent: Type = 1;
}

pub mod _KWAIT_REASON {
    pub type Type = ::core::ffi::c_int;
    pub const Executive: Type = 0;
    pub const FreePage: Type = 1;
    pub const PageIn: Type = 2;
    // ... 更多值
}

pub mod _LOCK_OPERATION {
    pub type Type = ::core::ffi::c_int;
    pub const IoReadAccess: Type = 0;
    pub const IoWriteAccess: Type = 1;
    pub const IoModifyAccess: Type = 2;
}

pub mod _MODE {
    pub type Type = ::core::ffi::c_int;
    pub const KernelMode: Type = 0;
    pub const UserMode: Type = 1;
}
```

## WDK-SYS 常量

### 状态码 (wdk-sys 内部定义)

```rust
// 从 wdk-sys 导入
use wdk_sys::{STATUS_SUCCESS, STATUS_PENDING, STATUS_UNSUCCESSFUL};

// STATUS_SUCCESS = 0
// STATUS_PENDING = 0x00000103
```

## IRP 处理说明

**driver-framework (IOCTL 用)：**
- `get_ioctl_params()` - 获取 IOCTL 参数
- `complete_request()` - 完成 IRP 请求
- `read_input_buffer()` / `write_output_buffer()` - 读写缓冲区

**WSK IRP 操作 (参考 DriverUtil)：**
- WSK 需要独立的 IRP 处理，不能使用 driver-framework
- WSK 使用 `IoAllocateIrp` / `IoFreeIrp` 分配/释放 IRP
- WSK 使用 `IoAllocateMdl` / `IoFreeMdl` 分配/释放 MDL
- WSK 使用 `KeInitializeEvent` / `KeWaitForSingleObject` 等待操作完成

## WSK 实现模式

参考 DriverUtil 的实现，WSK 操作遵循以下模式：

```rust
// 1. 分配 IRP
let irp = IoAllocateIrp(1, false as BOOLEAN);

// 2. 初始化事件
KeInitializeEvent(event as PRKEVENT, SYNCHRONIZATION_EVENT, false as BOOLEAN);

// 3. 设置完成回调 (手动实现 IoSetCompletionRoutine 宏)
io_set_completion_routine(irp, Some(completion_routine), event as PVOID, true, true, true);

// 4. 调用 WSK 函数
let status = wsk_function(..., irp);

// 5. 等待完成
if status == STATUS_PENDING {
    KeWaitForSingleObject(event as PVOID, EXECUTIVE, KERNEL_MODE, false as BOOLEAN, ptr::null());
    status = (*irp).IoStatus.__bindgen_anon_1.Status;
}

// 6. 获取结果并释放 IRP
let result = (*irp).IoStatus.Information;
IoFreeIrp(irp);
```

**关键函数：**
- `WskRegister` / `WskDeregister` - 注册/注销 WSK (需手动声明，不在 wdk-sys 中)
- `WskCaptureProviderNPI` / `WskReleaseProviderNPI` - 获取/释放 Provider (需手动声明)
- `WskSocket` - 创建 Socket
- `WskBind` - 绑定地址
- `WskListen` - 监听连接
- `WskAccept` - 接受连接
- `WskConnect` - 连接服务器
- `WskSend` / `WskReceive` - 发送/接收数据
- `WskCloseSocket` - 关闭 Socket

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
│  │  TCP Client     │─────── JSON 协议 ─────────►│  TCP Server     │ │
│  │  (连接 50080)   │                              │  (监听 50080)   │ │
│  │                 │                              │                 │ │
│  │  - 解析 JSON    │◄────── JSON 响应 ──────────│  - 解析 JSON     │ │
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
                    │  - TcpListen         │
                    │  - TcpAccept         │
                    │  - TcpSend / TcpRecv │
                    │                       │
                    └──────────────────────┘
```

### 为什么选择 JSON + TCP Server 架构？

| 对比项 | IOCTL + 结构体 | JSON + TCP Server |
|--------|----------------|-------------------|
| 内存对齐 | 必须严格对齐，容易出错 | 无需对齐，JSON 是文本协议 |
| 调试 | 困难，二进制数据难读 | 方便，直接看到明文 |
| 跨语言 | Go/Rust 结构体需一一对应 | 只需共享 JSON Schema |
| 依赖 | 简单，WDK 内置 | 需要 WSK |
| 灵活性 | 结构体固定，增减字段麻烦 | JSON 可扩展性好 |
| 远程调试 | 不支持 | 支持，Go 和驱动可不在同一机器 |

### 通信流程

```
1. Rust 驱动加载后，启动 TCP Server 监听 50080
2. Go 作为 TCP Client 连接到 127.0.0.1:50080
3. Go 发送 JSON 命令:
   {
     "action": "read_memory",
     "target": "0x1000",
     "size": 64
   }
4. Rust 驱动解析 JSON，执行对应操作
5. Rust 驱动返回 JSON 响应:
   {
     "success": true,
     "message": "ok"
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
pub struct Listener {
    socket: *mut Socket,
}

pub struct StreamSocket {
    socket: *mut Socket,
}

// 类型别名，保持兼容性
pub type WskListener = Listener;
pub type WskSocket = StreamSocket;
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
- **网络功能**：驱动作为 TCP Server 监听 Go 客户端，使用 JSON 通信

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
   - 使用 TCP Server + JSON 协议

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

### 4. 驱动签名问题
**重要**：当构建脚本输出"请等待驱动签名后再进行测试..."时，表示驱动已编译完成，但需要等待用户签名。此时**不要立即运行测试**，必须等待用户签名完成后再进行测试。

### 5. 类型不匹配
**原因**：没有查看 wdk-sys 绑定文件，自己发明类型
**解决**：强制阅读 `driver-framework/out/ntddk.rs` 和 `types.rs`，使用正确的类型

## 项目结构

```
rust-driver/
|
├── driver-framework/          # 通用驱动框架库
│   ├── out/                   # ⚠️ WDK-SYS 绑定文件 (必须阅读!)
│   │   ├── ntddk.rs           # 所有内核函数签名
│   │   ├── types.rs           # 所有类型定义
│   │   └── constants.rs       # 所有常量
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
│       ├── socket.rs          # Listener, StreamSocket 实现
│       ├── buffer.rs          # 数据缓冲区
│       ├── provider.rs        # WSK Provider
│       └── error.rs           # 错误类型
|
├── protocol/                  # 通信协议定义
│   ├── types.rs               # 基本类型
│   └── events.rs              # 事件类型
|
├── examples/
│   ├── sysdemo/               # 驱动示例 (IOCTL)
│   └── netdemo/               # 网络测试驱动 (TCP Server + JSON)
|
└── erebus-main/               # 完整驱动框架
    ├── km/                    # 内核模式
    └── um/                    # 用户模式
```

## 下一步计划

```
1. [ ] 实现 WSK TCP Server 功能 -> 驱动监听 50080
2. [ ] 实现 JSON 解析 -> 解析收到的 JSON 数据
3. [ ] 实现命令处理 -> 执行调试命令
4. [ ] 实现响应发送 -> 返回 JSON 结果
```
