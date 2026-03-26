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

## 📚 参考项目 (已验证可用)

### WSK 网络编程参考

| 项目 | 路径 | 说明 |
|------|------|------|
| **KSOCKET** | `todo/参考价值/KSOCKET-master/` | Berkeley Socket API 风格的 WSK 封装，支持 TCP/UDP |
| **afdmjhk** | `todo/参考价值/afdmjhk-master/WSK/` | WSK 示例代码，包含 echoserv 服务器 |
| **DriverUtil** | `todo/参考价值/代码+课件/DriverUtil/` | EPT Hook 和 WSK Socket 实现 |
| **Windows-driver-samples** | `todo/参考价值/Windows-driver-samples-main/` | 微软官方驱动示例 |

### KSOCKET 关键实现 (Berkeley Socket API 风格)

```c
// 来自 KSOCKET-master/src/ksocket/ksocket.c

// 异步上下文结构 - 核心！
typedef struct _KSOCKET_ASYNC_CONTEXT
{
  KEVENT CompletionEvent;  // 完成事件
  PIRP Irp;                // IRP 指针
} KSOCKET_ASYNC_CONTEXT;

// 完成回调 - 必须返回 STATUS_MORE_PROCESSING_REQUIRED
NTSTATUS NTAPI KspAsyncContextCompletionRoutine(
  _In_ PDEVICE_OBJECT DeviceObject,
  _In_ PIRP Irp,
  _In_ PKEVENT CompletionEvent
  )
{
  KeSetEvent(CompletionEvent, IO_NO_INCREMENT, FALSE);
  return STATUS_MORE_PROCESSING_REQUIRED;  // 关键！
}

// 初始化异步上下文
NTSTATUS NTAPI KspAsyncContextAllocate(PKSOCKET_ASYNC_CONTEXT AsyncContext)
{
  KeInitializeEvent(&AsyncContext->CompletionEvent, SynchronizationEvent, FALSE);
  AsyncContext->Irp = IoAllocateIrp(1, FALSE);
  
  IoSetCompletionRoutine(
    AsyncContext->Irp,
    &KspAsyncContextCompletionRoutine,
    &AsyncContext->CompletionEvent,
    TRUE, TRUE, TRUE  // InvokeOnSuccess, InvokeOnError, InvokeOnCancel
  );
  return STATUS_SUCCESS;
}

// 等待完成
NTSTATUS NTAPI KspAsyncContextWaitForCompletion(
  PKSOCKET_ASYNC_CONTEXT AsyncContext,
  PNTSTATUS Status
)
{
  if (*Status == STATUS_PENDING) {
    KeWaitForSingleObject(&AsyncContext->CompletionEvent, Executive, KernelMode, FALSE, NULL);
    *Status = AsyncContext->Irp->IoStatus.Status;
  }
  return *Status;
}

// 重置异步上下文 (用于重用 IRP)
VOID NTAPI KspAsyncContextReset(PKSOCKET_ASYNC_CONTEXT AsyncContext)
{
  KeResetEvent(&AsyncContext->CompletionEvent);
  IoReuseIrp(AsyncContext->Irp, STATUS_UNSUCCESSFUL);
  IoSetCompletionRoutine(AsyncContext->Irp, &KspAsyncContextCompletionRoutine,
    &AsyncContext->CompletionEvent, TRUE, TRUE, TRUE);
}
```

### simplewsk 关键实现 (afdmjhk)

```c
// 来自 afdmjhk-master/WSK/Samples/echoserv/simplewsk.c

// 初始化 WSK 数据
NTSTATUS InitWskData(PIRP* pIrp, PKEVENT CompletionEvent)
{
  *pIrp = IoAllocateIrp(1, FALSE);
  KeInitializeEvent(CompletionEvent, SynchronizationEvent, FALSE);
  IoSetCompletionRoutine(*pIrp, CompletionRoutine, CompletionEvent, TRUE, TRUE, TRUE);
  return STATUS_SUCCESS;
}

// 初始化 WSK 缓冲区
NTSTATUS InitWskBuffer(PVOID Buffer, ULONG BufferSize, PWSK_BUF WskBuffer)
{
  WskBuffer->Offset = 0;
  WskBuffer->Length = BufferSize;
  WskBuffer->Mdl = IoAllocateMdl(Buffer, BufferSize, FALSE, FALSE, NULL);
  
  __try {
    MmProbeAndLockPages(WskBuffer->Mdl, KernelMode, IoWriteAccess);
  }
  __except(EXCEPTION_EXECUTE_HANDLER) {
    IoFreeMdl(WskBuffer->Mdl);
    return STATUS_ACCESS_VIOLATION;
  }
  return STATUS_SUCCESS;
}

// 释放 WSK 缓冲区
VOID FreeWskBuffer(PWSK_BUF WskBuffer)
{
  MmUnlockPages(WskBuffer->Mdl);
  IoFreeMdl(WskBuffer->Mdl);
}

// 创建 Socket
PWSK_SOCKET CreateSocket(ADDRESS_FAMILY AddressFamily, USHORT SocketType, ULONG Protocol, ULONG Flags)
{
  KEVENT CompletionEvent;
  PIRP Irp;
  PWSK_SOCKET WskSocket;
  
  InitWskData(&Irp, &CompletionEvent);
  
  Status = g_WskProvider.Dispatch->WskSocket(
    g_WskProvider.Client, AddressFamily, SocketType, Protocol, Flags,
    NULL, NULL, NULL, NULL, NULL, Irp);
    
  if (Status == STATUS_PENDING) {
    KeWaitForSingleObject(&CompletionEvent, Executive, KernelMode, FALSE, NULL);
    Status = Irp->IoStatus.Status;
  }
  
  WskSocket = NT_SUCCESS(Status) ? (PWSK_SOCKET)Irp->IoStatus.Information : NULL;
  IoFreeIrp(Irp);
  return WskSocket;
}

// TCP Server 示例 (echoserv.c)
static VOID NTAPI ServerThread(PVOID p)
{
  SOCKADDR_IN LocalAddress = {0}, RemoteAddress = {0};
  PWSK_SOCKET Socket;
  
  // 创建监听 Socket
  g_ServerSocket = CreateSocket(AF_INET, SOCK_STREAM, IPPROTO_TCP, WSK_FLAG_LISTEN_SOCKET);
  
  LocalAddress.sin_family = AF_INET;
  LocalAddress.sin_addr.s_addr = INADDR_ANY;
  LocalAddress.sin_port = HTONS(SERVER_PORT);
  
  // 绑定
  Bind(g_ServerSocket, (PSOCKADDR)&LocalAddress);
  
  // 接受连接循环
  while ((Socket = Accept(g_ServerSocket, (PSOCKADDR)&LocalAddress, (PSOCKADDR)&RemoteAddress)) != NULL)
  {
    // 为每个客户端创建线程
    PsCreateSystemThread(&hThread, THREAD_ALL_ACCESS, NULL, NULL, NULL, ClientThread, Socket);
    InterlockedIncrement(&g_ClientsCount);
    ZwClose(hThread);
  }
}
```

### WSK Socket 类型标志

```c
// 创建 Socket 时必须指定类型
#define WSK_FLAG_BASIC_SOCKET        0x00000000  // 基本套接字
#define WSK_FLAG_LISTEN_SOCKET       0x00000001  // 监听套接字 (TCP Server)
#define WSK_FLAG_DATAGRAM_SOCKET     0x00000002  // 数据报套接字 (UDP)
#define WSK_FLAG_CONNECTION_SOCKET   0x00000004  // 连接套接字 (TCP Client)
```

### WSK Dispatch 表结构

```c
// 不同类型的 Socket 有不同的 Dispatch 表
typedef struct _WSK_PROVIDER_BASIC_DISPATCH {
  WSK_CLOSE_SOCKET WskCloseSocket;
  WSK_CONTROL_SOCKET WskControlSocket;
} WSK_PROVIDER_BASIC_DISPATCH;

typedef struct _WSK_PROVIDER_CONNECTION_DISPATCH {
  WSK_CLOSE_SOCKET WskCloseSocket;
  WSK_BIND WskBind;
  WSK_CONNECT WskConnect;
  WSK_GET_LOCAL_ADDRESS WskGetLocalAddress;
  WSK_GET_REMOTE_ADDRESS WskGetRemoteAddress;
  WSK_SEND WskSend;
  WSK_RECEIVE WskReceive;
  WSK_DISCONNECT WskDisconnect;
  WSK_RELEASE WskRelease;
  WSK_CONTROL_SOCKET WskControlSocket;
} WSK_PROVIDER_CONNECTION_DISPATCH;

typedef struct _WSK_PROVIDER_LISTEN_DISPATCH {
  WSK_CLOSE_SOCKET WskCloseSocket;
  WSK_BIND WskBind;
  WSK_ACCEPT WskAccept;
  WSK_INSPECT WskInspect;
  WSK_CONTROL_SOCKET WskControlSocket;
} WSK_PROVIDER_LISTEN_DISPATCH;
```

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

参考 KSOCKET 和 simplewsk 的实现，WSK 操作遵循以下模式：

### 核心 Rust 实现模式

```rust
// 异步上下文结构
struct Context {
    event: KEVENT,  // 使用 KEVENT 类型 (24字节)，不能用 [u8; 32]
    irp: PIRP,
}

// 完成回调 - 必须返回 STATUS_MORE_PROCESSING_REQUIRED
// 注意：第一个参数必须是 PDEVICE_OBJECT，不是 PVOID！
unsafe extern "C" fn wsk_completion_routine(
    _device: PDEVICE_OBJECT,  // 必须是 PDEVICE_OBJECT 类型
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    if !context.is_null() {
        KeSetEvent(context as PRKEVENT, 0, false as BOOLEAN);
    }
    STATUS_MORE_PROCESSING_REQUIRED  // 关键！不是 STATUS_SUCCESS
}

// 初始化异步上下文
impl Context {
    fn new() -> Result<Self, SocketError> {
        unsafe {
            let irp = IoAllocateIrp(1, false as BOOLEAN);
            if irp.is_null() {
                return Err(SocketError::IrpAllocationFailed);
            }

            let mut ctx = Context {
                event: core::mem::zeroed(),  // KEVENT 初始化
                irp,
            };

            KeInitializeEvent(
                &raw mut ctx.event as PRKEVENT, 
                _EVENT_TYPE::SynchronizationEvent, 
                false as BOOLEAN
            );
            
            // 设置完成回调
            io_set_completion_routine(
                ctx.irp,
                Some(wsk_completion_routine),
                &raw mut ctx.event as PVOID,
                true, true, true  // InvokeOnSuccess, InvokeOnError, InvokeOnCancel
            );

            Ok(ctx)
        }
    }

    // 等待完成
    fn wait_for_completion(&mut self) -> NTSTATUS {
        unsafe {
            let _ = KeWaitForSingleObject(
                &raw mut self.event as PVOID,
                _KWAIT_REASON::Executive,
                KERNEL_MODE,
                false as BOOLEAN,
                ptr::null::<LARGE_INTEGER>() as PLARGE_INTEGER,
            );
            (*self.irp).IoStatus.__bindgen_anon_1.Status
        }
    }

    fn get_information(&self) -> usize {
        unsafe { (*self.irp).IoStatus.Information as usize }
    }

    fn free(self) {
        unsafe {
            if !self.irp.is_null() {
                IoFreeIrp(self.irp);
            }
        }
    }
}

// 设置完成回调的辅助函数
unsafe fn io_set_completion_routine(
    irp: PIRP,
    routine: Option<unsafe extern "C" fn(PDEVICE_OBJECT, PIRP, PVOID) -> NTSTATUS>,
    context: PVOID,
    invoke_on_success: bool,
    invoke_on_error: bool,
    invoke_on_cancel: bool,
) {
    let stack = io_get_next_irp_stack_location(irp);
    
    (*stack).CompletionRoutine = routine;
    (*stack).Context = context;
    
    let mut control: u8 = 0;
    if invoke_on_success { control |= SL_INVOKE_ON_SUCCESS as u8; }
    if invoke_on_error { control |= SL_INVOKE_ON_ERROR as u8; }
    if invoke_on_cancel { control |= SL_INVOKE_ON_CANCEL as u8; }
    (*stack).Control = control;
}

// 手动实现 IoGetNextIrpStackLocation (wdk-sys 没有提供)
#[inline(always)]
unsafe fn io_get_next_irp_stack_location(irp: PIRP) -> PIO_STACK_LOCATION {
    let irp_ref = &*irp;
    let current_location = irp_ref.CurrentLocation;
    let stack_count = irp_ref.StackCount;
    
    let stack_size = core::mem::size_of::<wdk_sys::_IO_STACK_LOCATION>();
    let irp_size = core::mem::size_of::<IRP>();
    
    let irp_addr = irp as usize;
    let stack_array_start = irp_addr + irp_size;
    
    let index = (stack_count - current_location + 1) as usize;
    let stack_addr = stack_array_start + (index * stack_size);
    
    stack_addr as PIO_STACK_LOCATION
}
```

### WSK 初始化流程

```rust
// 1. 定义 WSK 客户端 Dispatch
static mut WSK_CLIENT_DISPATCH: ClientDispatch = ClientDispatch {
    Version: 0x0100,  // WSK 版本 1.0
    Reserved: 0,
    WskClientEvent: None,
};

// 2. 注册 WSK 客户端
let wsk_client = ClientNpi {
    ClientContext: ptr::null_mut(),
    Dispatch: &raw const WSK_CLIENT_DISPATCH,
};
WskRegister(&wsk_client, &raw mut WSK_REGISTRATION);

// 3. 获取 Provider (使用 WSK_NO_WAIT 或 WSK_INFINITE_WAIT)
let status = WskCaptureProviderNPI(&raw mut WSK_REGISTRATION, WSK_NO_WAIT, &raw mut provider);

// 4. 创建 Socket (指定类型标志)
let wsk_socket_fn = (*provider.Dispatch).WskSocket;
wsk_socket(
    provider.Client, 
    AF_INET,           // AddressFamily
    SOCK_STREAM,       // SocketType
    IPPROTO_TCP,       // Protocol
    WSK_FLAG_LISTEN_SOCKET,  // Flags: 监听套接字
    null, null, null, null, null,
    irp
);
```

### WSK 数据收发

```rust
// 发送数据
pub fn send(&mut self, buf: &[u8]) -> Result<usize, SocketError> {
    unsafe {
        let mut ctx = Context::new()?;
        
        // 分配 MDL
        let mdl = IoAllocateMdl(
            buf.as_ptr() as PVOID, 
            buf.len() as ULONG, 
            false, false, null
        );

        if mdl.is_null() {
            ctx.free();
            return Err(SocketError::MdlAllocationFailed);
        }

        // 锁定页面 (发送用 IoWriteAccess)
        MmProbeAndLockPages(mdl, KERNEL_MODE, _LOCK_OPERATION::IoWriteAccess);

        let mut wsk_buf = Buffer { 
            Mdl: mdl, 
            Offset: 0, 
            Length: buf.len() as u64 
        };

        let dispatch = (*self.socket).Dispatch as *const ConnectionDispatch;
        let wsk_send_fn = (*dispatch).WskSend;
        let status = wsk_send(self.socket, &mut wsk_buf, 0, ctx.irp);

        let final_status = if status == STATUS_PENDING {
            ctx.wait_for_completion()
        } else {
            status
        };

        // 释放资源
        MmUnlockPages(mdl);
        IoFreeMdl(mdl);
        
        let bytes_sent = if final_status == STATUS_SUCCESS {
            ctx.get_information()
        } else {
            0
        };

        ctx.free();

        if final_status == STATUS_SUCCESS {
            Ok(bytes_sent)
        } else {
            Err(SocketError::SendFailed)
        }
    }
}

// 接收数据 (使用 IoReadAccess)
pub fn recv(&mut self, buf: &mut [u8]) -> Result<usize, SocketError> {
    // ... 类似 send，但使用 IoReadAccess
    MmProbeAndLockPages(mdl, KERNEL_MODE, _LOCK_OPERATION::IoReadAccess);
    // ...
}
```

### WSK 关键类型和常量

```rust
// KEVENT 大小为 24 字节，不能使用 [u8; 32]
use wdk_sys::KEVENT;

// 完成回调签名 - 第一个参数必须是 PDEVICE_OBJECT
unsafe extern "C" fn completion_routine(
    _device: PDEVICE_OBJECT,  // 不是 PVOID！
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    if !context.is_null() {
        KeSetEvent(context as PRKEVENT, 0, false as BOOLEAN);
    }
    STATUS_MORE_PROCESSING_REQUIRED  // 重要！不是 STATUS_SUCCESS
}

// SL_INVOKE_* 常量值 (来自 wdk-sys)
use wdk_sys::{SL_INVOKE_ON_SUCCESS, SL_INVOKE_ON_ERROR, SL_INVOKE_ON_CANCEL};

// Socket 类型标志
const WSK_FLAG_LISTEN_SOCKET: u32 = 0x00000001;     // TCP Server
const WSK_FLAG_DATAGRAM_SOCKET: u32 = 0x00000002;   // UDP
const WSK_FLAG_CONNECTION_SOCKET: u32 = 0x00000004; // TCP Client

// 地址族
const AF_INET: u16 = 2;      // IPv4
const AF_INET6: u16 = 23;    // IPv6

// Socket 类型
const SOCK_STREAM: u16 = 1;  // TCP
const SOCK_DGRAM: u16 = 2;   // UDP

// 协议
const IPPROTO_TCP: u32 = 6;
const IPPROTO_UDP: u32 = 17;
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
│  │  HTTP Client    │─────── HTTP POST ─────────►│  HTTP Server    │ │
│  │  (连接 50080)   │                              │  (监听 50080)   │ │
│  │                 │                              │                 │ │
│  │  - 发送 HTTP    │◄────── HTTP 200 JSON ──────│  - 解析 HTTP     │ │
│  │  - 解析 JSON    │                              │  - 执行调试     │ │
│  │  - 接收结果     │                              │  - 返回 JSON     │ │
│  └────────┬────────┘                              └────────┬────────┘ │
│           │                                                │          │
│           │                                                │          │
│           └──────────────────┐ ┌──────────────────────────┘          │
│                              │ │                                      │
│                              │ │   HTTP over TCP Socket               │
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

### 为什么选择 JSON + HTTP 架构？

| 对比项 | IOCTL + 结构体 | JSON + HTTP |
|--------|----------------|-------------|
| 内存对齐 | 必须严格对齐，容易出错 | 无需对齐，JSON 是文本协议 |
| 调试 | 困难，二进制数据难读 | 方便，直接看到明文 HTTP |
| 跨语言 | Go/Rust 结构体需一一对应 | 只需共享 JSON Schema |
| 依赖 | 简单，WDK 内置 | 需要 WSK |
| 灵活性 | 结构体固定，增减字段麻烦 | JSON 可扩展性好 |
| 远程调试 | 不支持 | 支持，Go 和驱动可不在同一机器 |
| 标准化 | 自定义协议 | HTTP 标准协议，易于测试 |

### 通信流程

```
1. Rust 驱动加载后，启动 HTTP Server 监听 50080
2. Go 作为 HTTP Client 发送 POST 请求到 127.0.0.1:50080
3. Go 发送 HTTP POST 请求:
   POST /api HTTP/1.1
   Content-Type: application/json
   
   {"action": "read_memory", "target": "0x1000", "size": 64}
4. Rust 驱动解析 HTTP 请求和 JSON，执行对应操作
5. Rust 驱动返回 HTTP 响应:
   HTTP/1.1 200 OK
   Content-Type: application/json
   Content-Length: 33
   
   {"success": true, "message": "ok"}
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
- **网络功能**：驱动作为 HTTP Server 监听 Go 客户端，使用 JSON 通信

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
   - 使用 HTTP Server + JSON 协议

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
│   └── netdemo/               # 网络测试驱动 (HTTP Server + JSON)
|
└── erebus-main/               # 完整驱动框架
    ├── km/                    # 内核模式
    └── um/                    # 用户模式
```

## 下一步计划

```
1. [x] 实现 WSK HTTP Server 功能 -> 驱动监听 50080 ✅ 已完成 (2024-03-26)
2. [x] 实现 JSON 解析 -> 解析收到的 JSON 数据 ✅ 已完成 (2024-03-26)
3. [x] 实现命令处理 -> 执行调试命令 ✅ 已完成 (2024-03-26)
4. [x] 实现响应发送 -> 返回 JSON 结果 ✅ 已完成 (2024-03-26)
```

---

## ✅ 任务 1 完成记录: WSK HTTP Server

### 完成日期
2024-03-26

### 验证结果
- ✅ 驱动加载成功，无 BSOD
- ✅ 监听 socket 创建成功
- ✅ 端口绑定成功 (50080)
- ✅ 主动轮询接受连接成功
- ✅ 多客户端连接处理正常
- ✅ HTTP 请求处理正常
- ✅ JSON 响应正常

### 关键修复记录
详见: `rust-driver/examples/netdemo/BSOD_FIX_REPORT.md`

| # | 问题 | 修复 |
|---|------|------|
| 9 | WSK_PROVIDER_DISPATCH 缺少 Version/Reserved 字段 | 参考 WDK wsk.h 修正结构体定义 |
| 10 | accept_event 事件回调从未被调用 | 改用主动 WskAccept 轮询 |
| 11 | CLIENT_LISTEN_DISPATCH 设置方式错误 | 运行时设置函数指针 |
| 12 | 端口字节序错误 | 使用 to_be() 转换为网络字节序 |
| 13 | Handler 值传递导致响应长度为 0 | 改为指针传递 `*mut ResponseWriter` |
| 14 | WriteJSON 缺少 HTTP 响应头 | 添加 HTTP/1.1 200 OK 头部 |

### 回退参考
如遇网络通信问题，请检查:
1. `rust-driver/net/src/lib.rs` - WSK 核心实现
2. `rust-driver/examples/netdemo/src/lib.rs` - 驱动入口和 handler
3. `rust-driver/examples/netdemo/BSOD_FIX_REPORT.md` - 完整修复记录

### 测试命令
```powershell
# 构建驱动
cd d:\ux\examples\hypedbg\rust-driver\examples\netdemo
.\build.ps1

# 运行测试客户端
go run m.go
```
