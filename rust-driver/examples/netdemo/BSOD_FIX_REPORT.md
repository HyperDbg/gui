# BSOD 修复报告 #7

## 问题信息

**Bugcheck Code**: 0xCE (DRIVER_UNLOADED_WITHOUT_CANCELLING_PENDING_OPERATIONS)
**Faulting Module**: netdemo.sys (已卸载)
**Faulting IP**: `<Unloaded_netdemo.sys>+0x10a5`
**时间**: 2026-03-25 (第七次崩溃）

## 问题分析

### 错误堆栈
```
nt!KeBugCheckEx
  -> nt!memset+0x28dad
    -> nt!MmProbeAndLockPages+0x3d80
      -> nt!setjmpex+0x4598
        -> <Unloaded_netdemo.sys>+0x10a5  <-- 驱动已卸载，但代码仍在执行！
```

### 根本原因

**驱动卸载时没有等待系统线程结束！**

`driver_unload` 设置 `SERVER_RUNNING = false` 后立即返回，但系统线程仍在运行。当线程继续执行时，驱动代码已被卸载，导致访问无效内存。

## 修复方案

### 第一次修复尝试（失败）

使用 `KeWaitForSingleObject` 等待线程句柄：

```rust
// 错误 ❌ - 线程句柄不是同步对象
let _ = KeWaitForSingleObject(
    SERVER_THREAD_HANDLE as PVOID,
    0x02,  // Executive
    0,     // KernelMode
    false,
    &timeout as *const i64 as PLARGE_INTEGER,
);
```

**结果**: Bugcheck 0xA (IRQL_NOT_LESS_OR_EQUAL)
- `KeWaitForSingleObject` 不能直接等待线程句柄
- 线程句柄不是可等待的同步对象

### 最终修复方案

使用轮询等待 + `PsTerminateSystemThread`：

```rust
#[inline(never)]
unsafe extern "C" fn server_thread_entry(_ctx: PVOID) {
    while SERVER_RUNNING {
        let mut interval: i64 = -10000000;
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    PsTerminateSystemThread(STATUS_SUCCESS);  // 正确终止线程
}

#[inline(never)]
pub unsafe extern "C" fn driver_unload(_driver: PDRIVER_OBJECT) {
    SERVER_RUNNING = false;
    
    // 轮询等待线程检测到退出标志
    let mut interval: i64 = -1000000;
    for _ in 0..100 {
        KeDelayExecutionThread(0, false, &mut interval);
    }
    
    if !SERVER_THREAD_HANDLE.is_null() {
        let _ = ZwClose(SERVER_THREAD_HANDLE);
        SERVER_THREAD_HANDLE = core::ptr::null_mut();
    }
}
```

## 验证结果

```
=== Step 1: Load Driver ===
驱动安装成功
驱动启动成功

=== Step 2: TCP Client Test ===
Dial failed: dial tcp 127.0.0.1:50080: connectex: No connection could be made...
(WSK 功能待启用)

=== Step 3: Unload Driver ===
驱动停止成功
驱动卸载成功
```

**驱动加载/卸载成功！** 无 BSOD。

## 经验教训

### 1. 线程同步的正确方式

| 方法 | 适用场景 | 注意事项 |
|------|----------|----------|
| `KeWaitForSingleObject` | 等待 Event, Semaphore, Mutex | 不能等待线程句柄 |
| `KeWaitForMultipleObjects` | 等待多个同步对象 | 同上 |
| 轮询 + `PsTerminateSystemThread` | 系统线程退出 | 简单可靠 |
| `KeSetEvent` + Event 对象 | 精确同步 | 需要额外 Event |

### 2. 驱动卸载的正确流程

```
1. 设置退出标志 (SERVER_RUNNING = false)
2. 等待线程检测到标志并退出
3. 调用 PsTerminateSystemThread (在线程内部)
4. 关闭线程句柄 (ZwClose)
5. 清理其他资源
6. 返回
```

### 3. Bugcheck 0xCE 的含义

`DRIVER_UNLOADED_WITHOUT_CANCELLING_PENDING_OPERATIONS` 表示：
- 驱动已卸载
- 但还有未完成的操作（如定时器、工作线程、回调等）
- 这些操作试图执行已卸载的代码

## 待解决问题

WSK 网络功能因栈溢出问题暂时禁用，需要：
1. 分析 `wsk` crate 的栈使用情况
2. 考虑使用 `WSK_NO_WAIT` 替代 `WSK_INFINITE_WAIT`
3. 优化 JSON 序列化/反序列化的栈使用

## 历史修复记录

| # | Bugcheck | 原因 | 修复 |
|---|----------|------|------|
| 1 | 0xD1 | NPI_CLIENT_CHARACTERISTICS 字段错误 | 修正结构体定义 |
| 2 | 0xD1 | ClientRegistrationInstance 是指针而非嵌入结构体 | 改为嵌入结构体 |
| 3 | 0xD1 | NPI_REGISTRATION_INSTANCE 未正确初始化 | 正确初始化所有字段 |
| 4 | 0x7E | PsCreateSystemThread 传入空指针 | 创建有效变量并传递地址 |
| 5 | 0x7E | 栈溢出（4KB buffer 在栈上） | 改用堆分配 Vec |
| 6 | 0x7E | 栈溢出（format! 消耗大量栈空间） | 移除所有 log!/format! 调用 |
| 7 | 0xCE | 驱动卸载时线程仍在运行 | 轮询等待 + PsTerminateSystemThread |
| 8 | 0xA | KeWaitForSingleObject 等待线程句柄失败 | 改用轮询等待 |
