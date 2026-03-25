# Rust Windows 驱动开发指南

## ⚠️ 重要警告

**严禁直接使用 `cargo build` 命令！**

- 本项目使用 **EWDK (Enterprise Windows Driver Kit)** 进行编译
- **永远不会安装 VS2022**
- 只能使用 PowerShell 构建脚本进行编译
- 当用户要求执行某个 `.ps1` 脚本时，**必须直接执行该脚本**，不要尝试手动运行 cargo 命令

## 构建环境

### 使用通用构建脚本

项目使用 EWDK (Enterprise Windows Driver Kit) 进行编译，**不要直接使用 `cargo build`**。

```powershell
# 使用通用构建脚本编译所有驱动项目
.\rust-driver\examples\build-all.ps1
```

该脚本会自动：
1. 挂载 EWDK ISO
2. 设置 WDK 构建环境
3. 编译所有驱动项目
4. 重命名 DLL 为 SYS

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

### 使用方法

在 `Cargo.toml` 中添加依赖：

```toml
[dependencies]
driver-framework = { path = "../driver-framework" }
```

在驱动代码中使用：

```rust
#![no_std]

extern crate alloc;
extern crate wdk_panic;

use driver_framework::{
    log, LogLevel,
    utf16, ToU16Vec, ToUnicodeString,
    DriverConfig, create_device_with_config, cleanup_device,
    ctl_code, default_create_close, get_ioctl_params, complete_request,
};
use wdk_alloc::WdkAllocator;
use wdk_sys::{NTSTATUS, PCUNICODE_STRING, PDRIVER_OBJECT, STATUS_SUCCESS};

#[global_allocator]
static GLOBAL_ALLOCATOR: WdkAllocator = WdkAllocator;

#[unsafe(export_name = "DriverEntry")]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    log!(LogLevel::Info, "Driver loading...");
    
    let config = DriverConfig {
        nt_device_name: "\\Device\\MyDriver",
        dos_device_name: "\\??\\MyDriver",
        ..Default::default()
    };
    
    let status = create_device_with_config(driver, &config);
    if !wdk::nt_success(status) {
        log!(LogLevel::Error, "Failed to create device: {:x}", status);
        return status;
    }
    
    log!(LogLevel::Success, "Driver loaded successfully");
    STATUS_SUCCESS
}
```

### IOCTL 定义

```rust
use driver_framework::{ctl_code, FILE_DEVICE_UNKNOWN, METHOD_BUFFERED, FILE_ANY_ACCESS};

const IOCTL_MY_COMMAND: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
```

### IOCTL 处理

```rust
use driver_framework::{get_ioctl_params, complete_request, read_input_buffer, write_output_buffer};

pub unsafe extern "C" fn handle_ioctl(_device: *mut DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    if let Some((code, input, input_len, output, output_len)) = get_ioctl_params(p_irp) {
        match code {
            IOCTL_MY_COMMAND => {
                if let Some(data) = read_input_buffer(input, input_len) {
                    // 处理数据
                    let response = b"OK";
                    let written = write_output_buffer(output, response);
                    complete_request(p_irp, STATUS_SUCCESS, written as u64);
                    return STATUS_SUCCESS;
                }
            }
            _ => {}
        }
    }
    complete_request(p_irp, STATUS_INVALID_DEVICE_REQUEST, 0);
    STATUS_INVALID_DEVICE_REQUEST
}
```

## 项目结构

### 已验证可编译的项目

1. **driver-framework** - `rust-driver/driver-framework/`
   - 通用驱动框架库
   - 所有驱动项目的基础依赖

2. **sysdemo** - `rust-driver/examples/sysdemo/`
   - 模块化 WDM 驱动示例
   - 支持 IOCTL 通信
   - 测试脚本: `src/m.go`

3. **erebus-main** - `rust-driver/erebus-main/`
   - 完整的驱动框架
   - 包含 km (内核模式) 和 um (用户模式)
   - shared 模块定义共享类型

### 待修复的项目

以下项目代码存在但可能无法编译：
- `hyperhv` - Hypervisor 核心
- `hyperkd` - 内核调试器
- `protocol` - 通信协议
- `wsk` - WSK 网络套接字
- `disassembler` - 反汇编器
- `pdbex` - PDB 解析
- `kd` - 内核调试支持

## 驱动开发模式

### sysdemo 模块化架构

```
sysdemo/src/
├── lib.rs        # 驱动入口点 (DriverEntry)
├── constants.rs  # 设备名称、符号链接等常量
├── device.rs     # 设备创建、清理
├── dispatch.rs   # IRP 派发函数、IOCTL 处理
├── ffi.rs        # Windows 内核 FFI 绑定
├── logger.rs     # 内核日志宏
└── utils.rs      # UTF-16 字符串转换等工具
```

### 关键技术点

#### 1. UTF-16 字符串

Windows 内核使用 `UNICODE_STRING`，必须使用 UTF-16 编码：

```rust
// 使用 utf16! 宏创建 UTF-16 字符串
utf16!(DEVICE_NAME, "\\Device\\sysDemo");
utf16!(SYMBOLIC_LINK_NAME, "\\??\\sysDemo");

// 初始化 UNICODE_STRING
RtlInitUnicodeString(&mut unicode_str, utf16_ptr);
```

#### 2. IOCTL 控制码

```rust
macro_rules! ctl_code {
    ($device_type:expr, $function:expr, $method:expr, $access:expr) => {
        (($device_type) << 16) | (($access) << 14) | (($function) << 2) | ($method)
    };
}

const IOCTL_SEND_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);
const IOCTL_RECEIVE_DATA: u32 = ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);
```

#### 3. 设备扩展

```rust
#[repr(C)]
pub struct DeviceExtension {
    pub buffer: [u8; BUFFER_SIZE],
    pub buffer_len: usize,
}

// 在 IoCreateDevice 时设置
(*device_object).DeviceExtension = extension_ptr;
```

#### 4. 日志宏

```rust
#[macro_export]
macro_rules! log {
    ($level:expr, $($arg:tt)*) => {
        unsafe {
            $crate::logger::log_inner($level, format_args!($($arg)*));
        }
    };
}

// 使用
log!(LogLevel::Info, "Driver loaded!");
log!(LogLevel::Error, "IOCTL failed: {:#x}", status);
```

## Go 用户模式测试

### 驱动通信

```go
p := driver.New(driverPath, serviceName, deviceName)
p.Install()
p.Start()

// IOCTL 通信
p.Send(bytes.NewBufferString("hello"), IOCTL_SEND_DATA)
receive := p.Receive(IOCTL_RECEIVE_DATA)

p.Stop()
p.Uninstall()
```

### provider.go 日志

已完善的日志输出：
- 驱动安装成功
- 驱动启动成功
- 驱动停止成功
- 驱动卸载成功

## 网络通信 (WSK)

### protocol 模块

定义调试器通信协议：
- `types.rs` - 基本类型 (RegisterState, ProcessInfo 等)
- `events.rs` - 事件类型
- 常量: `PROTOCOL_VERSION = 1`, `DEFAULT_PORT = 9527`

### wsk 模块

WSK (Winsock Kernel) 网络套接字封装：
- `TcpStream` - TCP 客户端
- `TcpListener` - TCP 服务端
- `UdpSocket` - UDP 套接字
- `WskProvider` - WSK 提供者注册

### go-communicator

Go 端通信实现：
- `protocol/protocol.go` - 协议定义
- `debugger/debugger.go` - 调试器接口
- `mcp/communicator_mcp.go` - MCP 通信

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

## 文件排除

`.gitignore` 配置：
```
**/target/
rust-driver/todo/
```

## 下一步计划

1. 创建网络通信 demo 驱动
2. 修复 `protocol` 模块编译
3. 修复 `wsk` 模块编译
4. 实现 WSK 实际调用
5. 集成到 hyperhv/hyperkd
