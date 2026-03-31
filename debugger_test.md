# HyperDbg 用户模式调试 API 实现原理

本文档解释 HyperDbg 的 `StartProcess` 和 `KillProcess` API 的 C++ 实现原理，以及 Go/Rust 端的对应实现。

## 1. C++ 实现原理

### 1.1 .start 命令流程

**文件位置**: `hyperdbg/libhyperdbg/code/debugger/user-level/ud.cpp`

```
用户输入: .start path "c:\windows\system32\notepad.exe"
    ↓
[命令解析] 解析路径和参数
    ↓
[进程启动] 创建挂起进程 (CreateProcess + CREATE_SUSPENDED)
    ↓
[初始化] 初始化用户调试器
    ↓
[IOCTL 请求] 发送 ATTACH 请求到内核
    ↓
[内核处理] 驱动接收请求并设置断点
    ↓
[Hook 移除循环] 循环发送 REMOVE_HOOKS 请求
    ↓
[等待入口点] 等待进程到达入口点
    ↓
[恢复执行] ResumeThread 恢复进程执行
    ↓
[调试就绪] 进程在入口点暂停，等待调试命令
```

### 1.2 核心数据结构

```c
// 附加/分离用户模式进程请求
typedef struct _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS {
    UINT32 ProcessId;                          // 目标进程 ID
    UINT32 ThreadId;                           // 目标线程 ID
    BOOLEAN IsStartingNewProcess;              // 是否启动新进程
    BOOLEAN CheckCallbackAtFirstInstruction;   // 是否在第一条指令执行回调
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION Action;  // 操作类型
    UINT32 Result;                             // 返回结果
} DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS;

// 操作类型
typedef enum _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION {
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH = 0,       // 附加
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH = 1,       // 分离
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS = 2, // 移除 Hook
} DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION;
```

### 1.3 关键步骤详解

#### 步骤 1: 创建挂起进程

```cpp
UdCreateSuspendedProcess(TargetFileAddress, 】, &ProcInfo);
AttachRequest.ProcessId = ProcInfo.dwProcessId;
AttachRequest.ThreadId  = ProcInfo.dwThreadId;
```

**关键点**:
- 进程以**挂起状态**创建 (`CREATE_SUSPENDED`)
- 这样可以在调试器完全初始化后再恢复执行
- 保存进程 ID 和线程 ID 用于后续操作

#### 步骤 2: 准备附加请求

```cpp
DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS AttachRequest = {0};
AttachRequest.IsStartingNewProcess = TRUE;
AttachRequest.CheckCallbackAtFirstInstruction = RunCallbackAtTheFirstInstruction;
AttachRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH;
```

#### 步骤 3: 发送 IOCTL 请求

```cpp
Status = DeviceIoControl(
    g_DeviceHandle,
    IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,  // 0x80e
    &AttachRequest,
    SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
    &AttachRequest,
    SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
    &ReturnedLength,
    NULL
);
```

**IOCTL 代码**: `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80e)

#### 步骤 4: 移除 Hook 循环

```cpp
while (TRUE)
{
    AttachRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS;
    Status = DeviceIoControl(...);
    
    if (AttachRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL)
    {
        break; // Hook 移除成功
    }
    else if (AttachRequest.Result == DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED)
    {
        Sleep(1000); // 等待入口点到达
        continue;
    }
}
```

**关键机制**:
- 循环发送 `REMOVE_HOOKS` 请求
- 如果入口点未到达，等待 1 秒后重试
- 直到 Hook 成功移除或发生错误

#### 步骤 5: 恢复进程执行

```cpp
if (ProcInfo.hThread != NULL64_ZERO)
{
    ResumeThread(ProcInfo.hThread);
}
```

### 1.4 驱动端处理

当用户模式调试器发送 `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (0x80e)` 时：

```c
// 驱动端处理 (Ioctl.c)
case IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS:
    DebuggerAttachOrDetachToThreadRequest = 
        (PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS)Irp->AssociatedIrp.SystemBuffer;
    
    AttachingTargetProcess(DebuggerAttachOrDetachToThreadRequest);
    break;
```

#### AttachingTargetProcess 的核心操作

```c
// 1. 设置EPT钩子监控PEB
EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR EptHookDetails = {0};
EptHookDetails.StartAddress = PebAddressToMonitor;
EptHookDetails.EndAddress = PebAddressToMonitor;
EptHookDetails.SetHookForRead = TRUE;
EptHookDetails.SetHookForWrite = TRUE;
EptHookDetails.SetHookForExec = FALSE;

DebuggerEventEnableMonitorReadWriteExec(
    &EptHookDetails,
    AttachRequest->ProcessId,
    FALSE // Applied from VMX non-root mode
);

// 2. 设置等待入口点标志
g_IsWaitingForUserModeProcessEntryToBeCalled = TRUE;

// 3. 附加到目标进程
// 保存进程信息，设置调试上下文
```

---

## 2. Go 端实现

### 2.1 接口定义

```go
// debugger/interfaces.go
type Debugger interface {
    StartProcess(exePath string) (uint32, error)
    KillProcess(processID uint32) error
    // ...
}
```

### 2.2 实现

```go
// debugger/packet.go
func (p *Packet) StartProcess(exePath string) (uint32, error) {
    // 1. 创建挂起进程
    const CREATE_SUSPENDED = 0x00000004
    var si StartupInfo
    var pi ProcessInformation
    
    ret, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("CreateProcessW").Call(
        0,
        uintptr(unsafe.Pointer(cmdLinePtr)),
        0, 0, 0,
        CREATE_SUSPENDED,
        0, 0,
        uintptr(unsafe.Pointer(&si)),
        uintptr(unsafe.Pointer(&pi)),
    )
    
    // 2. 发送 IOCTL 请求到驱动
    data := json.Marshal(map[string]any{
        "action":     "start_process",
        "process_id": pi.ProcessID,
        "thread_id":  pi.ThreadID,
        "exe_path":   exePath,
    })
    resp := SendReceive[Empty](p, data)
    
    // 3. 恢复进程执行
    ResumeThread(pi.ThreadHandle)
    
    return pi.ProcessID, nil
}
```

---

## 3. Rust 端实现

### 3.1 Request 结构体

```rust
// generated/router.rs
pub struct Request {
    pub action: String,
    pub process_id: Option<u32>,
    pub thread_id: Option<u32>,
    pub exe_path: Option<String>,
    // ...
}
```

### 3.2 API 实现

```rust
// hyperdbg_api.rs
impl DebuggerApi for HyperDbgApi {
    fn start_process(&mut self, req: &Request) -> Result<Empty, String> {
        let process_id = req.process_id.ok_or("process_id required")?;
        let mut ud = self.user_debugger.lock();
        ud.initialize()?;
        ud.start_process(process_id)?;
        self.current_process_id = Some(process_id);
        Ok(Empty {})
    }

    fn kill_process(&mut self, req: &Request) -> Result<Empty, String> {
        let process_id = req.process_id.ok_or("process_id required")?;
        let mut ud = self.user_debugger.lock();
        ud.kill_process(process_id)?;
        if self.current_process_id == Some(process_id) {
            self.current_process_id = None;
        }
        Ok(Empty {})
    }
}
```

### 3.3 UserDebugger 实现

```rust
// ud.rs
impl UserDebugger {
    pub fn start_process(&mut self, process_id: ProcessId) -> Result<u64, UdError> {
        if !self.active {
            return Err(UdError::NotActive);
        }

        let token = self.generate_debugging_token();
        let process_details = ProcessDebuggingDetails::new(process_id, token);
        self.process_debugging_details.push(Arc::new(Mutex::new(process_details)));

        Ok(token)
    }

    pub fn kill_process(&mut self, process_id: ProcessId) -> Result<(), UdError> {
        if !self.active {
            return Err(UdError::NotActive);
        }

        let pos = self.process_debugging_details
            .iter()
            .position(|p| p.lock().process_id == process_id)
            .ok_or(UdError::ProcessNotFound(process_id))?;

        self.process_debugging_details.remove(pos);
        Ok(())
    }
}
```

---

## 4. 单元测试说明

### 4.1 TestRustDriverHTTP

```go
func TestRustDriverHTTP(t *testing.T) {
    // 1. 安装并启动驱动
    drv := driver.NewWithOptions(driverPath, "hyperdbg", "\\\\.\\hyperdbg", true)
    drv.Install()
    drv.Start()
    defer drv.Stop()
    defer drv.Uninstall()

    // 2. 连接到驱动 HTTP 服务
    conn := debugger.NewPacket("http://127.0.0.1:50080")
    conn.Connect()
    defer conn.Disconnect()

    // 3. 启动目标进程 (使用 API 而非 exec.Command)
    notepadPID, err := conn.StartProcess("notepad.exe")
    defer conn.KillProcess(notepadPID)

    // 4. 加载 VMM
    conn.LoadVmm()

    // 5. 附加到进程
    conn.AttachProcess(notepadPID)

    // 6. 设置断点
    conn.SetBreakpoint(0x7FFE0000, debugger.BreakpointSoftware)

    // 7. 继续执行
    conn.Continue()
}
```

### 4.2 与 C++ 实现的对比

| 方面 | C++ 实现 | Go 实现 |
|------|----------|---------|
| 进程创建 | `CreateProcessW` + `CREATE_SUSPENDED` | `CreateProcessW` + `CREATE_SUSPENDED` |
| IOCTL 通信 | `DeviceIoControl` | HTTP JSON API |
| Hook 移除循环 | 循环发送 `REMOVE_HOOKS` | 驱动内部处理 |
| 恢复执行 | `ResumeThread` | `ResumeThread` |
| 进程终止 | `TerminateProcess` | `KillProcess` API |

---

## 5. 关键特性

1. **挂起启动**: 进程以挂起状态创建，确保调试器完全初始化
2. **Hook 管理**: 通过 IOCTL 与内核驱动通信，管理断点和 Hook
3. **入口点同步**: 等待进程到达入口点后再恢复执行
4. **参数支持**: 支持传递命令行参数给目标进程
5. **错误处理**: 完善的错误检查和消息提示

---

## 6. IOCTL 代码表

| IOCTL | 代码 | 说明 |
|-------|------|------|
| `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` | 0x80e | 附加/分离用户态进程 |
| `IOCTL_KILL_PROCESS` | 0x82a | 杀死进程 |
| `IOCTL_RESTART_PROCESS` | 0x82b | 重启进程 |
| `IOCTL_SEND_USER_DEBUGGER_COMMANDS` | 0x817 | 发送用户调试器命令 |
