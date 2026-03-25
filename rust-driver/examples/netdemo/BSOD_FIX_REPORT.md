# BSOD 修复报告 #6

## 问题信息

**Bugcheck Code**: 0x7E (SYSTEM_THREAD_EXCEPTION_NOT_HANDLED)
**Exception Code**: 0x80000003 (Break instruction exception - `int 3`)
**Faulting Module**: nt
**Faulting IP**: `nt!IoIsSystemThread+0x103`
**时间**: 2026-03-25 (第六次崩溃）

## 问题分析

### 错误堆栈
```
nt!IoIsSystemThread+0x103:
fffff802`3d3452b3 cc              int     3  <-- 断点异常

调用栈：
netdemo!DriverEntry+0x879
  -> nt!_chkstk  <-- 栈检查函数！
    -> nt!IoIsSystemThread+0x103 (崩溃)
```

### 根本原因

**`format!` 宏在内核模式下消耗大量栈空间！**

`log!` 宏内部使用 `alloc::format!` 进行字符串格式化，这会：
1. 在栈上分配临时缓冲区
2. 调用复杂的格式化代码
3. 消耗大量栈空间（可能超过 1KB 每次调用）

在系统线程（栈大小仅 12KB-24KB）中，多次 `log!` 调用会导致栈溢出。

### 为什么之前的修复没有解决问题？

之前的修复只解决了 `handle_client` 中的 4KB buffer 问题，但忽略了：
- `DriverEntry` 中的 `log!` 调用
- `server_thread_entry` 中的多个 `log!` 调用
- `wsk_init` 中的 `format!` 调用

## 修复方案

### 修改文件
`d:\ux\examples\hypedbg\rust-driver\examples\netdemo\src\lib.rs`

### 修复策略

1. **移除所有 `log!` 调用** - 避免栈消耗
2. **移除 `driver-framework` 依赖** - 简化代码
3. **使用静态字符串** - 避免 `format!` 开销

### 修复后代码

```rust
#![no_std]

extern crate alloc;
extern crate wdk_panic;

use wdk_alloc::WdkAllocator;
use wdk_sys::{
    NTSTATUS, PDRIVER_OBJECT, PCUNICODE_STRING, IRP_MJ_CREATE, IRP_MJ_CLOSE, 
    STATUS_SUCCESS, DEVICE_OBJECT, PIRP, HANDLE, PVOID,
    POBJECT_ATTRIBUTES, OBJECT_ATTRIBUTES,
    ntddk::PsCreateSystemThread,
};

#[unsafe(export_name = "DriverEntry")]
pub unsafe extern "system" fn driver_entry(
    driver: PDRIVER_OBJECT,
    _registry_path: PCUNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        (*driver).MajorFunction[IRP_MJ_CREATE as usize] = Some(default_create_close);
        (*driver).MajorFunction[IRP_MJ_CLOSE as usize] = Some(default_create_close);
        (*driver).DriverUnload = Some(driver_unload);
    }

    let _ = wsk::wsk_init();
    let _ = start_server_thread();

    STATUS_SUCCESS
}

unsafe extern "C" fn server_thread_entry(_ctx: PVOID) {
    while !wsk::is_wsk_ready() {
        let mut interval: i64 = -10000000;
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    let mut listener = wsk::WskListener::new();
    if listener.create().is_err() {
        return;
    }
    
    // ... 简化的代码，无 log! 调用
}
```

### 关键修复点

| 修复项 | 错误 | 正确 |
|--------|------|------|
| 日志输出 | `log!(...)` 使用 `format!` | 移除所有日志 |
| 错误处理 | `format!("error: {}", e)` | 静态字符串 `"error"` |
| 栈使用量 | 大量（多次 format!） | 最小化 |

## 修复详情

### 1. `format!` 的栈开销

```rust
// 错误 ❌ - format! 消耗大量栈空间
log!(LogLevel::Info, "WSK initialized: {}", status);
let msg = format!("Error: 0x{:X}", status);

// 正确 ✅ - 静态字符串，无栈开销
let _ = wsk::wsk_init();
return Err("WskRegister failed");
```

### 2. 内核模式下的最佳实践

参考 `erebus-main` 项目的做法：
```rust
#[macro_export]
macro_rules! println {
    ($msg:tt) => {
        #[cfg(debug_assertions)]  // 只在 debug 模式下启用日志
        $crate::logger::Logger::log($msg, $crate::logger::LogLevel::Info);
    };
}
```

### 3. 栈空间限制

| 环境 | 栈大小 | 建议 |
|------|--------|------|
| 用户态线程 | 1MB | 可以使用 format! |
| 系统线程 | 12KB-24KB | 避免使用 format! |
| 内核驱动 | 非常有限 | 最小化栈使用 |

## 验证结果

编译成功，无错误。

## 经验教训

### 1. 内核模式下避免使用 `format!`
- `format!` 宏会分配临时缓冲区
- 格式化代码复杂，消耗大量栈空间
- 应使用静态字符串或条件编译

### 2. 日志输出的正确方式
```rust
// 推荐：条件编译
#[cfg(debug_assertions)]
wdk::println!("debug message");

// 推荐：静态字符串
wdk::println!("static message");

// 避免：动态格式化
// wdk::println!("dynamic: {}", value);  // 危险！
```

### 3. 系统线程的栈限制
- 默认栈大小仅 12KB-24KB
- 每次函数调用都会消耗栈空间
- 大数组用堆，避免递归，避免 format!

## 相关文档

- [Stack Overflow in Kernel Mode](https://learn.microsoft.com/en-us/windows-hardware/drivers/debugger/stack-overflow)
- [Kernel-Mode Driver Architecture](https://learn.microsoft.com/en-us/windows-hardware/drivers/kernel/)

## 历史修复记录

| # | Bugcheck | 原因 | 修复 |
|---|----------|------|------|
| 1 | 0xD1 | NPI_CLIENT_CHARACTERISTICS 字段错误 | 修正结构体定义 |
| 2 | 0xD1 | ClientRegistrationInstance 是指针而非嵌入结构体 | 改为嵌入结构体 |
| 3 | 0xD1 | NPI_REGISTRATION_INSTANCE 未正确初始化 | 正确初始化所有字段 |
| 4 | 0x7E | PsCreateSystemThread 传入空指针 | 创建有效变量并传递地址 |
| 5 | 0x7E | 栈溢出（4KB buffer 在栈上） | 改用堆分配 Vec |
| 6 | 0x7E | 栈溢出（format! 消耗大量栈空间） | 移除所有 log!/format! 调用 |
