# sysDemo Rust Driver

这是 `d:\ux\examples\hypedbg\debugger\driver\sysDemo` 的 Rust 实现版本。

## 致谢

本项目参考了 [erebus](https://github.com/ovxrfl0w/erebus) 项目的架构设计，感谢原作者提供的优秀 Rust 驱动开发示例。

## 功能

这个驱动实现了一个简单的 IOCTL 通信接口，支持以下功能：

- **IOCTL_SEND_DATA (0x800)**: 从用户空间接收数据并存储在驱动中
- **IOCTL_RECEIVE_DATA (0x801)**: 从驱动读取数据，并在前面添加 "received data by user " 前缀

## 项目结构

```
sysdemo/
├── Cargo.toml          # Rust 项目配置
├── build.ps1          # 构建脚本
└── src/
    └── lib.rs         # 驱动实现
```

## 与 C++ 版本的差异

| 特性 | C++ 版本 | Rust 版本 |
|------|----------|-----------|
| 内存管理 | 手动管理 | 所有权系统 + Arc/Mutex |
| 错误处理 | 返回码 | Result 类型 |
| 并发 | 自旋锁 | Atomic 类型 |
| 字符串处理 | UNICODE_STRING | String |
| 设备扩展 | 结构体 | 结构体 + 方法 |

## 构建要求

- Rust 工具链 (x86_64-pc-windows-msvc)
- EWDK (Enterprise Windows Driver Kit)
- PowerShell 5.0+

## 构建步骤

1. **安装依赖**:
   ```powershell
   # 确保 Rust 已安装
   rustup default stable-x86_64-pc-windows-msvc
   ```

2. **构建驱动**:
   ```powershell
   cd d:\ux\examples\hypedbg\rust-driver\sysdemo
   .\build.ps1
   ```

3. **输出位置**:
   ```
   .\target\debug\sysdemo.sys
   ```

## 使用方法

### 加载驱动

```powershell
# 创建服务
sc create sysdemo type= kernel binPath= "C:\path\to\sysdemo.sys"

# 启动驱动
sc start sysdemo
```

### 测试驱动

```c
#include <windows.h>
#include <stdio.h>

#define IOCTL_SEND_DATA    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)
#define IOCTL_RECEIVE_DATA CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)

int main() {
    HANDLE hDevice = CreateFile(
        L"\\\\.\\sysDemo",
        GENERIC_READ | GENERIC_WRITE,
        0,
        NULL,
        OPEN_EXISTING,
        FILE_ATTRIBUTE_NORMAL,
        NULL
    );

    if (hDevice == INVALID_HANDLE_VALUE) {
        printf("Failed to open device: %d\n", GetLastError());
        return 1;
    }

    // 发送数据
    char sendData[] = "Hello from user mode!";
    DWORD bytesReturned;
    BOOL result = DeviceIoControl(
        hDevice,
        IOCTL_SEND_DATA,
        sendData,
        sizeof(sendData),
        NULL,
        0,
        &bytesReturned,
        NULL
    );

    if (!result) {
        printf("Failed to send data: %d\n", GetLastError());
        CloseHandle(hDevice);
        return 1;
    }

    printf("Sent %d bytes\n", bytesReturned);

    // 接收数据
    char receiveBuffer[1024];
    result = DeviceIoControl(
        hDevice,
        IOCTL_RECEIVE_DATA,
        NULL,
        0,
        receiveBuffer,
        sizeof(receiveBuffer),
        &bytesReturned,
        NULL
    );

    if (!result) {
        printf("Failed to receive data: %d\n", GetLastError());
        CloseHandle(hDevice);
        return 1;
    }

    printf("Received: %.*s\n", bytesReturned, receiveBuffer);
    CloseHandle(hDevice);
    return 0;
}
```

### 卸载驱动

```powershell
# 停止驱动
sc stop sysdemo

# 删除服务
sc delete sysdemo
```

## 技术细节

### 设备扩展

```rust
#[repr(C)]
pub struct DeviceExtension {
    data_buffer: [u8; MAX_DATA_SIZE],
    data_length: AtomicU32,
}
```

- `data_buffer`: 存储用户发送的数据 (最大 4096 字节)
- `data_length`: 使用原子操作存储数据长度，保证线程安全

### IOCTL 处理

- **IOCTL_SEND_DATA**: 从用户空间复制数据到设备扩展的缓冲区
- **IOCTL_RECEIVE_DATA**: 从设备扩展读取数据，添加前缀后返回给用户

### IRP 处理

- **IRP_MJ_CREATE**: 处理设备打开
- **IRP_MJ_CLOSE**: 处理设备关闭
- **IRP_MJ_DEVICE_CONTROL**: 处理 IOCTL 请求

## 调试

驱动使用 `wdk` crate 的日志功能。可以在 DebugView 中查看驱动输出：

```
DriverEntry - Device created successfully
IOCTL_SEND_DATA received
Received data from user sended: 20 bytes
IOCTL_RECEIVE_DATA received
Returning data to user: 38 bytes
DriverUnload
```

## 优势

相比 C++ 版本，Rust 版本具有以下优势：

1. **内存安全**: 编译时检查防止空指针、缓冲区溢出等问题
2. **线程安全**: 原子类型和所有权系统保证数据竞争安全
3. **错误处理**: Result 类型强制处理所有可能的错误情况
4. **零成本抽象**: Rust 的高级特性不会带来运行时开销
5. **现代工具链**: Cargo 提供依赖管理和构建系统

## 注意事项

1. 驱动需要在测试环境或虚拟机中运行
2. 需要管理员权限加载驱动
3. 首次加载可能需要禁用驱动签名强制
4. 生产环境需要代码签名证书

## 许可证

与原 HyperDbg 项目保持一致
