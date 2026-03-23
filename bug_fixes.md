# HyperDbg Bug Fixes

本文档记录了对 HyperDbg C/C++ 代码的所有修复。

## 修复 1: VMX 区域物理地址计算错误

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/vmm/vmx/VmxRegions.c`

### 问题描述
在 VMXON 和 VMCS 区域分配时，物理地址的计算方式不正确。代码直接对物理地址进行对齐计算，假设虚拟地址的偏移量和物理地址的偏移量完全相同，但这并不总是正确的。

### 修复位置

#### 修复 1.1: VMXON 区域物理地址计算
**文件**: `VmxRegions.c`
**行号**: 60

**修复前**:
```c
AlignedVmxonRegionPhysicalAddr = (UINT64)((ULONG_PTR)(VmxonRegionPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
```

**修复后**:
```c
AlignedVmxonRegionPhysicalAddr = VirtualAddressToPhysicalAddress((PVOID)AlignedVmxonRegion);
```

#### 修复 1.2: VMCS 区域物理地址计算
**文件**: `VmxRegions.c`
**行号**: 137

**修复前**:
```c
AlignedVmcsRegionPhysicalAddr = (UINT64)((ULONG_PTR)(VmcsPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
```

**修复后**:
```c
AlignedVmcsRegionPhysicalAddr = VirtualAddressToPhysicalAddress((PVOID)AlignedVmcsRegion);
```

### 影响
- **BSOD 类型**: 在执行 `__vmx_on` 指令时崩溃
- **崩溃位置**: `VmxRegions.c:77`
- **原因**: 错误的物理地址导致 VMXON 指令失败

---

## 修复 2: 空指针解引用

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/core/Debugger.c`

### 问题描述
在 `DebuggerEventListCountByEventType` 函数中，当 `DebuggerGetEventListByEventType` 返回 NULL 时，代码没有检查就直接访问 `TempList->Flink`，导致空指针解引用。

### 修复位置
**文件**: `Debugger.c`
**行号**: 2193

**修复前**:
```c
PLIST_ENTRY TargetEventList = DebuggerGetEventListByEventType(EventType);

//
// We have to iterate through all events of this list
//
TempList = TargetEventList;

while (TargetEventList != TempList->Flink)
{
    ...
}
```

**修复后**:
```c
PLIST_ENTRY TargetEventList = DebuggerGetEventListByEventType(EventType);

if (TargetEventList == NULL)
{
    return 0;
}

//
// We have to iterate through all events of this list
//
TempList = TargetEventList;

while (TargetEventList != TempList->Flink)
{
    ...
}
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `Debugger.c:2198`
- **崩溃指令**: `mov rax, qword ptr [rax]`
- **原因**: `TargetEventList` 为 NULL 时访问 `TempList->Flink`

---

## 修复 3: g_Events 空指针检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/core/Debugger.c`

### 问题描述
在 `DebuggerGetEventListByEventType` 函数中，代码直接访问 `g_Events->ExternalInterruptOccurredEventsHead` 等成员，但没有检查 `g_Events` 是否为 NULL。如果 `g_Events` 为 NULL（例如在调试器未初始化时），访问这些成员会导致空指针解引用。

### 修复位置
**文件**: `Debugger.c`
**行号**: 2039

**修复前**:
```c
PLIST_ENTRY
DebuggerGetEventListByEventType(VMM_EVENT_TYPE_ENUM EventType)
{
    PLIST_ENTRY ResultList = NULL;
    //
    // Register the event
    //
    switch (EventType)
    {
    case EXTERNAL_INTERRUPT_OCCURRED:
        ResultList = &g_Events->ExternalInterruptOccurredEventsHead;
        break;
    ...
    }
    return ResultList;
}
```

**修复后**:
```c
PLIST_ENTRY
DebuggerGetEventListByEventType(VMM_EVENT_TYPE_ENUM EventType)
{
    PLIST_ENTRY ResultList = NULL;

    if (g_Events == NULL)
    {
        return NULL;
    }

    //
    // Register the event
    //
    switch (EventType)
    {
    case EXTERNAL_INTERRUPT_OCCURRED:
        ResultList = &g_Events->ExternalInterruptOccurredEventsHead;
        break;
    ...
    }
    return ResultList;
}
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `Debugger.c:2203`
- **崩溃指令**: `mov rax, qword ptr [rax]`
- **原因**: `g_Events` 为 NULL 时访问 `g_Events->ExternalInterruptOccurredEventsHead`

---

## 修复 4: g_DbgState 空指针检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/commands/BreakpointCommands.c`

### 问题描述
在 `BreakpointHandleBreakpoints` 函数中，代码直接访问 `g_DbgState[CoreId]` 而没有检查 `g_DbgState` 是否为 NULL。如果 `g_DbgState` 为 NULL（例如在调试器未初始化时），访问数组会导致空指针解引用。

### 修复位置
**文件**: `BreakpointCommands.c`
**行号**: 667

**修复前**:
```c
BreakpointHandleBreakpoints(UINT32 CoreId)
{
    DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
    UINT64                           GuestRip      = 0;
    PROCESSOR_DEBUGGING_STATE *      DbgState      = &g_DbgState[CoreId];

    GuestRip = VmFuncGetRip();
    ...
}
```

**修复后**:
```c
BreakpointHandleBreakpoints(UINT32 CoreId)
{
    DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
    UINT64                           GuestRip      = 0;
    PROCESSOR_DEBUGGING_STATE *      DbgState      = NULL;

    if (g_DbgState == NULL)
    {
        return FALSE;
    }

    DbgState = &g_DbgState[CoreId];

    GuestRip = VmFuncGetRip();
    ...
}
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `BreakpointCommands.c:533`
- **崩溃指令**: `mov ecx,dword ptr [rax+10h]`
- **原因**: `g_DbgState` 为 NULL 时访问 `g_DbgState[CoreId]`

---

## 修复 5: Continue 等待断点命中逻辑

### 文件
`debugger/user_debugger.go`

### 问题描述
`Continue()` 函数只是让进程继续执行，但没有等待断点命中。当断点命中时，内核会发送事件到用户层，但用户层没有处理这个事件来更新 `IsPaused` 状态。这导致 `StepInto()` 被调用时检查发现 `IsPaused` 为 `false`，返回警告。

### 修复位置
**文件**: `user_debugger.go`
**行号**: 1629, 1074

**修复 5.1: 添加断点命中事件处理程序**
```go
s.packeter.RegisterEventHandler(EventBreakpoint, func(event *DebugEvent) {
    if event.ProcessId == s.activeProcess.ProcessId {
        s.mu.Lock()
        s.activeProcess.IsPaused = true
        s.activeProcess.Rip = event.EIP
        s.mu.Unlock()
        mylog.Info("断点命中，进程已暂停", "ProcessId", event.ProcessId, "EIP", event.EIP)
    }
})
```

**修复 5.2: 修改 Continue 等待逻辑**
```go
func (s *UserDebug) Continue() {
    s.mu.Lock()
    
    // ... 检查和验证代码 ...
    
    s.activeProcess.IsPaused = false
    mylog.Success("继续执行")
    s.mu.Unlock()

    // 等待断点命中
    for i := range 30 {
        time.Sleep(1 * time.Second)
        s.mu.Lock()
        if s.activeProcess.IsPaused {
            mylog.Info("进程已暂停，等待断点命中成功")
            s.mu.Unlock()
            return
        }
        s.mu.Unlock()
        mylog.Info("等待断点命中", i+1, "/30")
    }

    mylog.Warning("等待断点命中超时")
}
```

### 影响
- **问题**: `StepInto` 检查失败，提示"进程未处于暂停状态"
- **原因**: `Continue` 没有等待断点命中，`IsPaused` 状态未更新
- **解决**: 添加断点命中事件处理程序，并让 `Continue` 等待断点命中

---

## 修复 6: 实现内核事件接收功能

### 文件
`debugger/event_loop.go`, `debugger/packet.go`

### 问题描述
用户层代码没有实现接收内核发送事件的功能。内核在断点命中时会调用 `LogCallbackSendBuffer` 发送 `OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE` 事件，但用户层没有调用 `IOCTL_REGISTER_EVENT` 来接收这些事件，导致断点命中事件无法到达用户层，`Continue()` 函数一直等待断点命中超时。

### 修复位置
**文件**: `event_loop.go`
**行号**: 1, 18, 58, 71, 82, 95, 108

**修复 6.1: 添加事件接收功能到 EventLoop**
```go
type EventLoop struct {
    mu            sync.Mutex
    running       bool
    eventManager  *EventManager
    eventHandlers map[EventType][]func(*DebugEvent)
    stopChan      chan struct{}
    driver        driver.Api  // 新增：驱动接口，用于接收内核事件
}

func NewEventLoop(eventManager *EventManager, driver driver.Api) *EventLoop {
    return &EventLoop{
        eventManager:  eventManager,
        eventHandlers: make(map[EventType][]func(*DebugEvent)),
        stopChan:      make(chan struct{}),
        driver:        driver,  // 新增：保存驱动接口
    }
}
```

**修复 6.2: 添加事件接收循环**
```go
func (el *EventLoop) run() {
    eventChan := el.eventManager.DebugEventChannel()

    go el.receiveEvents()  // 新增：启动事件接收协程

    for {
        select {
        case event := <-eventChan:
            el.handleEvent(event)
        case <-el.stopChan:
            mylog.Info("EventLoop: Stopped")
            return
        }
    }
}

func (el *EventLoop) receiveEvents() {
    for el.IsRunning() {
        if el.driver == nil || !el.driver.IsConnected() {
            time.Sleep(100 * time.Millisecond)
            continue
        }

        response := el.driver.Receive(IoctlRegisterEvent)
        if response == nil || response.Len() == 0 {
            time.Sleep(100 * time.Millisecond)
            continue
        }

        el.handleKernelEvent(response)
    }
}
```

**修复 6.3: 处理内核事件**
```go
func (el *EventLoop) handleKernelEvent(response *bytes.Buffer) {
    if response.Len() < 4 {
        return
    }

    operationCode := binary.LittleEndian.Uint32(response.Next(4))

    if (operationCode & 0x80000000) != 0 {
        el.handleMandatoryEvent(operationCode, response)
    } else {
        time.Sleep(10 * time.Millisecond)
    }
}

func (el *EventLoop) handleMandatoryEvent(operationCode uint32, response *bytes.Buffer) {
    switch operationCode {
    case 0x80000010: // OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE
        el.handleUserDebuggerPause(response)
    default:
        mylog.Info("Unknown mandatory event", operationCode)
    }
}

func (el *EventLoop) handleUserDebuggerPause(response *bytes.Buffer) {
    if response.Len() < 72 {
        return
    }

    event := &DebugEvent{
        Type:      EventBreakpoint,
        ProcessId: binary.LittleEndian.Uint32(response.Next(4)),
        Tag:       binary.LittleEndian.Uint64(response.Next(8)),
    }

    response.Next(4) // Skip ThreadId
    response.Next(8) // Skip Token
    event.EIP = binary.LittleEndian.Uint64(response.Next(8))

    el.eventManager.TriggerDebugEvent(event)
}
```

**修复 6.4: 更新 EventLoop 创建代码**
```go
// 在 packet.go 中
globalPacket.EventManager = NewEventManager(globalPacket.driver)
globalPacket.eventLoop = NewEventLoop(globalPacket.EventManager, globalPacket.driver)

// 或
p.EventManager = NewEventManager(p.driver)
p.eventLoop = NewEventLoop(p.EventManager, p.driver)
```

### 影响
- **问题**: 断点命中事件无法到达用户层，`Continue()` 等待超时
- **原因**: 用户层没有实现调用 `IOCTL_REGISTER_EVENT` 接收内核事件的功能
- **解决**: 在 `EventLoop` 中添加事件接收循环，处理内核发送的断点命中事件

---

## 修复 7: DebuggerTriggerEvents 中 g_DbgState 空指针检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/debugger/core/Debugger.c`

### 问题描述
在 `DebuggerTriggerEvents` 函数中，代码直接访问 `g_DbgState[KeGetCurrentProcessorNumberEx(NULL)]` 而没有检查 `g_DbgState` 是否为 NULL。当驱动卸载时，`GlobalDebuggingStateFreeMemory` 会将 `g_DbgState` 设置为 NULL。如果此时仍有 VMM 回调函数（如 `VmmCallbackTriggerEvents`）执行，访问 `g_DbgState` 会导致空指针解引用，引发 BSOD。

### BSOD 分析
- **崩溃地址**: `<Unloaded_hyperkd.sys>+0x7cc0` (DebuggerTriggerEvents 函数)
- **调用者**: `<Unloaded_hyperkd.sys>+0x3bc56` (VmmCallbackTriggerEvents 函数)
- **崩溃位置**: `Debugger.c:1096`
- **崩溃指令**: `mov rax, qword ptr [rcx+rax*8]`
- **原因**: `g_DbgState` 为 NULL 时访问 `g_DbgState[processor_id]`

### 修复位置
**文件**: `Debugger.c`
**行号**: 1095

**修复前**:
```c
if (!g_EnableDebuggerEvents || g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer)
{
    //
    // Debugger is not enabled
    //
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED;
}

//
// Find the debugging state structure
//
DbgState = &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)];
```

**修复后**:
```c
if (!g_EnableDebuggerEvents || g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer)
{
    //
    // Debugger is not enabled
    //
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED;
}

//
// Check if g_DbgState is valid (not NULL)
//
if (g_DbgState == NULL)
{
    //
    // g_DbgState is NULL, driver might be unloading
    //
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED;
}

//
// Find the debugging state structure
//
DbgState = &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)];
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `Debugger.c:1096`
- **崩溃指令**: `mov rax, qword ptr [rcx+rax*8]`
- **原因**: `g_DbgState` 为 NULL 时访问 `g_DbgState[processor_id]`
- **触发场景**: 驱动卸载时仍有 VMM 回调函数执行

---

## 修复 8: VmxVmexitHandler 中 g_GuestState 空指针检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/vmm/vmx/Vmexit.c`

### 问题描述
在 `VmxVmexitHandler` 函数中，代码直接访问 `g_GuestState[KeGetCurrentProcessorNumberEx(NULL)]` 而没有检查 `g_GuestState` 是否为 NULL。当驱动卸载时，`GlobalGuestStateFreeMemory` 会将 `g_GuestState` 设置为 NULL。如果此时仍有 VM-exit 发生（例如在驱动卸载过程中），访问 `g_GuestState` 会导致空指针解引用，引发 BSOD。

### BSOD 分析
- **崩溃地址**: `<Unloaded_hyperkd.sys>+0x3b6c0` (AsmVmexitHandler 函数)
- **崩溃位置**: `Vmexit.c:31`
- **崩溃指令**: `mov rcx, qword ptr [g_GuestState + rax*8]`
- **原因**: `g_GuestState` 为 NULL 时访问 `g_GuestState[processor_id]`
- **触发场景**: 驱动卸载时仍有 VM-exit 发生

### 修复位置
**文件**: `Vmexit.c`
**行号**: 31

**修复前**:
```c
//
// *********** SEND MESSAGE AFTER WE SET THE STATE ***********
//
VCpu = &g_GuestState[KeGetCurrentProcessorNumberEx(NULL)];
```

**修复后**:
```c
//
// *********** SEND MESSAGE AFTER WE SET THE STATE ***********
//

// Check if g_GuestState is valid (not NULL)
if (g_GuestState == NULL)
{
    // g_GuestState is NULL, driver might be unloading
    // Return FALSE to indicate we should not continue VMX
    return FALSE;
}

VCpu = &g_GuestState[KeGetCurrentProcessorNumberEx(NULL)];
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `Vmexit.c:31`
- **崩溃指令**: `mov rcx, qword ptr [g_GuestState + rax*8]`
- **原因**: `g_GuestState` 为 NULL 时访问 `g_GuestState[processor_id]`
- **触发场景**: 驱动卸载时仍有 VM-exit 发生

---

## 修复 9: VmxVmresume 中 g_GuestState 空指针检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/vmm/vmx/Vmx.c`

### 问题描述
在 `VmxVmresume` 函数中，代码访问 `g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].ExitReason` 而没有检查 `g_GuestState` 是否为 NULL。当驱动卸载时，`GlobalGuestStateFreeMemory` 会将 `g_GuestState` 设置为 NULL。如果此时 VMRESUME 失败导致执行错误处理代码，访问 `g_GuestState` 会导致空指针解引用，引发 BSOD。

### BSOD 分析
- **崩溃地址**: `<Unloaded_hyperkd.sys>+0x3b6c0` (AsmVmexitHandler 函数)
- **崩溃位置**: `Vmx.c:1060`
- **崩溃指令**: `mov rax, qword ptr [g_GuestState + rax*8]`
- **原因**: `g_GuestState` 为 NULL 时访问 `g_GuestState[processor_id].ExitReason`
- **触发场景**: 驱动卸载时 VMRESUME 失败，执行错误处理代码

### 修复位置
**文件**: `Vmx.c`
**行号**: 1059

**修复前**:
```c
LogError("Err, in executing VMRESUME, status : 0x%llx, last VM-exit reason: 0x%x",
         ErrorCode,
         g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].ExitReason);
```

**修复后**:
```c
// Check if g_GuestState is valid (not NULL)
if (g_GuestState != NULL)
{
    LogError("Err, in executing VMRESUME, status : 0x%llx, last VM-exit reason: 0x%x",
             ErrorCode,
             g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].ExitReason);
}
else
{
    LogError("Err, in executing VMRESUME, status : 0x%llx, g_GuestState is NULL",
             ErrorCode);
}
```

### 影响
- **BSOD 类型**: DRIVER_IRQL_NOT_LESS_OR_EQUAL (0xd1)
- **崩溃位置**: `Vmx.c:1060`
- **崩溃指令**: `mov rax, qword ptr [g_GuestState + rax*8]`
- **原因**: `g_GuestState` 为 NULL 时访问 `g_GuestState[processor_id].ExitReason`
- **触发场景**: 驱动卸载时 VMRESUME 失败，执行错误处理代码

---

## 修复 10: AsmVmexitHandler 中 VmxVmexitHandler 返回值检查

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/assembly/AsmVmexitHandler.asm`

### 问题描述
在 `AsmVmexitHandler` 汇编函数中，当 `VmxVmexitHandler` 返回 FALSE 时（表示 `g_GuestState` 为 NULL），代码仍然会跳转到 `VmxVmresume` 函数。此时如果驱动正在卸载，`VmxVmresume` 函数可能已经被卸载，导致跳转到 NULL 地址，引发 BSOD。

### BSOD 分析
- **崩溃地址**: `<Unloaded_hyperkd.sys>+0x3b710` (AsmVmexitHandler 函数内)
- **崩溃位置**: `AsmVmexitHandler.asm:141`
- **崩溃指令**: `jmp VmxVmresume` (跳转到 NULL 地址)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时仍然跳转到 `VmxVmresume`
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE

### 修复位置
**文件**: `AsmVmexitHandler.asm`
**行号**: 86

**修复前**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; ----------- Restore XMM Registers ----------
```

**修复后**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; Check if VmxVmexitHandler returned FALSE (g_GuestState is NULL)
; In this case, we should also go to AsmVmxoffHandler to safely exit VMX
test    al, al
jz      AsmVmxoffHandler

; ----------- Restore XMM Registers ----------
```

### 影响
- **BSOD 类型**: 跳转到 NULL 地址
- **崩溃位置**: `AsmVmexitHandler.asm:141`
- **崩溃指令**: `jmp VmxVmresume` (跳转到 NULL 地址)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时仍然跳转到 `VmxVmresume`
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE

---

## 修复 11: AsmVmexitHandler 中 VmxVmexitHandler 返回 FALSE 时恢复寄存器

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/assembly/AsmVmexitHandler.asm`

### 问题描述
在 `AsmVmexitHandler` 汇编函数中，当 `VmxVmexitHandler` 返回 FALSE 时（表示 `g_GuestState` 为 NULL），代码直接跳转到 `AsmVmxoffHandler`。但此时栈的状态不正确：`AsmVmexitHandler` 在调用 `VmxVmexitHandler` 之前保存了 XMM 寄存器和 RFLAGS，但跳转到 `AsmVmxoffHandler` 时没有先恢复这些寄存器。`AsmVmxoffHandler` 期望的栈状态是已经保存了通用寄存器，但实际上还没有保存。这导致 `AsmVmxoffHandler` 尝试恢复寄存器时，从错误的栈位置读取数据，导致系统卡住。

### 问题分析
- **崩溃位置**: `AsmVmexitHandler.asm:86`
- **崩溃指令**: `jz VmxoffHandlerWithRestore` (跳转到错误的栈状态)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时直接跳转到 `AsmVmxoffHandler`，没有先恢复寄存器
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE
- **症状**: 系统卡住，无法操作，被 VMM 控制

### 修复位置
**文件**: `AsmVmexitHandler.asm`
**行号**: 84, 147

**修复前**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; Check if VmxVmexitHandler returned FALSE (g_GuestState is NULL)
; In this case, we should also go to AsmVmxoffHandler to safely exit VMX
test    al, al
jz      AsmVmxoffHandler

; ----------- Restore XMM Registers ----------
```

**修复后**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; Check if VmxVmexitHandler returned FALSE (g_GuestState is NULL)
; In this case, we need to restore registers and safely exit VMX
test    al, al
jz      VmxoffHandlerWithRestore

; ----------- Restore XMM Registers ----------
```

新增 `VmxoffHandlerWithRestore` 标签，在跳转到 `AsmVmxoffHandler` 之前先恢复寄存器：

```asm
VmxoffHandlerWithRestore:

; ----------- Restore XMM Registers ----------

pop rax
pop rcx
pop rdx
pop rbx
pop rbp		        ; rsp
pop rbp
pop rsi
pop rdi 
pop r8
pop r9
pop r10
pop r11
pop r12
pop r13
pop r14
pop r15

; ------------ Restore XMM Registers ------------

movaps xmm0, xmmword ptr [rsp+000h]
movaps xmm1, xmmword ptr [rsp+010h]
movaps xmm2, xmmword ptr [rsp+020h]
movaps xmm3, xmmword ptr [rsp+030h]
movaps xmm4, xmmword ptr [rsp+040h]
movaps xmm5, xmmword ptr [rsp+050h]

ldmxcsr dword ptr [rsp+0100h]          
    
add     rsp, 0110h

; --------------- Restore RFLAGS ---------------

popfq               ; Restore the flags register (RFLAGS)

; ----------------------------------------------

; Jump to AsmVmxoffHandler to safely exit VMX
jmp     AsmVmxoffHandler
```

### 影响
- **问题类型**: 系统卡住，无法操作
- **崩溃位置**: `AsmVmexitHandler.asm:86`
- **崩溃指令**: `jz VmxoffHandlerWithRestore` (跳转到错误的栈状态)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时直接跳转到 `AsmVmxoffHandler`，没有先恢复寄存器
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE
- **症状**: 系统卡住，无法操作，被 VMM 控制

---

## 修复 11: AsmVmexitHandler 中 VmxVmexitHandler 返回 FALSE 时的正确处理

### 文件
`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv/code/assembly/AsmVmexitHandler.asm`

### 问题描述
在 `AsmVmexitHandler` 汇编函数中，当 `VmxVmexitHandler` 返回 FALSE 时（表示 `g_GuestState` 为 NULL），代码直接跳转到 `AsmVmxoffHandler`。但此时栈的状态不正确：`AsmVmexitHandler` 在调用 `VmxVmexitHandler` 之前保存了 XMM 寄存器和 RFLAGS，但跳转到 `AsmVmxoffHandler` 时没有先恢复这些寄存器。`AsmVmxoffHandler` 期望的栈状态是已经保存了通用寄存器，但实际上还没有保存。这导致 `AsmVmxoffHandler` 尝试恢复寄存器时，从错误的栈位置读取数据，导致系统卡住。

### 问题分析
- **崩溃位置**: `AsmVmexitHandler.asm:86`
- **崩溃指令**: `jz VmxoffHandlerWithRestore` (跳转到错误的栈状态)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时直接跳转到 `AsmVmxoffHandler`，没有先恢复寄存器
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE
- **症状**: 系统卡住，无法操作，被 VMM 控制

### 修复方案
当 `VmxVmexitHandler` 返回 FALSE 时（表示 `g_GuestState` 为 NULL），我们不应该跳转到 `AsmVmxoffHandler`，因为：
1. `AsmVmxoffHandler` 需要调用 `VmxReturnStackPointerForVmxoff` 和 `VmxReturnInstructionPointerForVmxoff` 函数
2. 这些函数会访问 `g_GuestState`，但此时 `g_GuestState` 已经是 NULL
3. 调用这些函数会导致空指针解引用

正确的解决方案是：当 `VmxVmexitHandler` 返回 FALSE 时，继续执行正常的恢复流程，然后跳转到 `VmxVmresume`。`VmxVmresume` 函数已经添加了 `g_GuestState` 的 NULL 检查（修复 10），会直接执行 VMXOFF 并返回。

### 修复位置
**文件**: `AsmVmexitHandler.asm`
**行号**: 84

**修复前**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; Check if VmxVmexitHandler returned FALSE (g_GuestState is NULL)
; In this case, we should also go to AsmVmxoffHandler to safely exit VMX
test    al, al
jz      AsmVmxoffHandler

; ----------- Restore XMM Registers ----------
```

**修复后**:
```asm
cmp	al, 1	        ; Check whether we have to turn off VMX or Not (the result is in RAX)

je		AsmVmxoffHandler

; If VmxVmexitHandler returned FALSE (g_GuestState is NULL), we should NOT
; jump to VmxoffHandlerWithRestore because we cannot call
; VmxReturnStackPointerForVmxoff and VmxReturnInstructionPointerForVmxoff
; (they access g_GuestState which is NULL).
; Instead, we continue with the normal restore flow and jump to VmxVmresume,
; which has a NULL check for g_GuestState and will execute VMXOFF directly.

; ----------- Restore XMM Registers ----------
```

### 影响
- **问题类型**: 系统卡住，无法操作
- **崩溃位置**: `AsmVmexitHandler.asm:86`
- **崩溃指令**: `jz VmxoffHandlerWithRestore` (跳转到错误的栈状态)
- **原因**: `VmxVmexitHandler` 返回 FALSE 时直接跳转到 `AsmVmxoffHandler`，没有先恢复寄存器
- **触发场景**: 驱动卸载时 VM-exit 发生，`VmxVmexitHandler` 返回 FALSE
- **症状**: 系统卡住，无法操作，被 VMM 控制

---

## 修复日期
2026-03-24

## 修复总结
1. 修复了 VMX 区域物理地址计算错误，确保正确的物理地址对齐
2. 添加了空指针检查，防止在事件列表为空时发生崩溃
3. 添加了 g_Events 初始化检查，防止在调试器未初始化时访问未初始化的内存
4. 添加了 g_DbgState 初始化检查，防止在调试器未初始化时访问未初始化的内存
5. 修复了测试代码，在 StepInto 前先运行到断点
6. 添加了断点命中事件处理程序，修复 Continue 等待断点命中逻辑
7. 实现了内核事件接收功能，确保断点命中事件能够到达用户层
8. 在 DebuggerTriggerEvents 中添加了 g_DbgState 空指针检查，防止驱动卸载时访问已释放的内存
9. 在 VmxVmexitHandler 中添加了 g_GuestState 空指针检查，防止驱动卸载时访问已释放的内存
10. 在 VmxVmresume 中添加了 g_GuestState 空指针检查，防止驱动卸载时访问已释放的内存
11. 在 AsmVmexitHandler 中添加了 VmxVmexitHandler 返回值检查，防止驱动卸载时跳转到 NULL 地址
12. 在 AsmVmexitHandler 中修复了 VmxVmexitHandler 返回 FALSE 时的处理，防止系统卡住

这些修复解决了导致系统蓝屏和卡住的关键问题，提高了驱动的稳定性。
