# HyperDbg C++ 命令 → IOCTL → 结构体 → Go 实现完整映射表

> 本文档基于对 C++ 源码 `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands` 的系统性分析
> 
> 生成时间: 2026-04-12

---

## 一、IOCTL 常量完整定义 (来自 Ioctls.h)

```go
// 文件: core/debugger.go

const (
    // 0x800
    IoctlRegisterEvent                            = 0x00222000
    // 0x801
    IoctlReturnIrpPendingPacketsAndDisallowIoctl   = 0x00222004
    // 0x802
    IoctlTerminateVmx                              = 0x00222008
    // 0x803
    IoctlReadMemory                                = 0x0022200C
    // 0x804
    IoctlReadWriteMsr                              = 0x00222010
    // 0x805
    IoctlReadPageTableEntriesDetails               = 0x00222014
    // 0x806
    IoctlDebuggerRegisterEvent                     = 0x00222018
    // 0x807
    IoctlAddActionToEvent                          = 0x0022201C
    // 0x808
    IoctlHideAndUnhideToTransparent                = 0x00222020
    // 0x809
    IoctlVa2paAndPa2va                             = 0x00222024
    // 0x80A
    IoctlEditMemory                                = 0x00222028
    // 0x80B
    IoctlSearchMemory                              = 0x0022202C
    // 0x80C
    IoctlModifyEvents                              = 0x00222030
    // 0x80D
    IoFlushLoggingBuffers                          = 0x00222034
    // 0x80E
    IoctlAttachDetachUserModeProcess               = 0x00222038
    // 0x80F
    IoctlPrint                                     = 0x0022203C
    // 0x810
    IoctlPrepareDebuggee                           = 0x00222040
    // 0x811
    IoctlPausePacketReceived                       = 0x00222044
    // 0x812
    IoctlSendSignalExecutionFinished               = 0x00222048
    // 0x813
    IoctlSendUsermodeMessages                      = 0x0022204C
    // 0x814
    IoctlSendGeneralBuffer                         = 0x00222050
    // 0x815
    IoctlPerformKernelSideTests                    = 0x00222054
    // 0x816
    IoctlReservePreAllocatedPools                  = 0x00222058
    // 0x817
    IoctlSendUserDebuggerCommands                  = 0x0022205C
    // 0x818
    IoctlGetDetailOfActiveThreadsAndProcesses      = 0x00222060
    // 0x819
    IoctlGetUserModeModuleDetails                  = 0x00222064
    // 0x81A
    IoctlQueryCountOfProcessesOrThreads            = 0x00222068
    // 0x81B
    IoctlGetListOfThreadsAndProcesses              = 0x0022206C
    // 0x81C
    IoctlQueryCurrentProcess                       = 0x00222070
    // 0x81D
    IoctlQueryCurrentThread                        = 0x00222074
    // 0x81E
    IoctlRequestRevMachineService                  = 0x00222078
    // 0x81F
    IoctlBringPagesIn                              = 0x0022207C
    // 0x820
    IoctlPreactivateFunctionality                  = 0x00222080
    // 0x821
    IoctlPcieEndpointEnum                          = 0x00222084
    // 0x822
    IoctlPerformActionsOnApic                      = 0x00222088
    // 0x823
    IoctlPciDevinfoEnum                           = 0x0022208C
    // 0x824
    IoctlQueryIdtEntry                             = 0x00222090
    // 0x825
    IoctlSetBreakpointUserDebugger                 = 0x00222094
    // 0x826
    IoctlPerformSmiOperation                       = 0x00222098
    // 0x827
    IoctlPerformHypertraceOperation                = 0x0022209C
)
```

---

## 二、Meta Commands 映射表

### 2.1 .attach - 附加到进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/attach.cpp` |
| **核心函数** | `UdAttachToProcess(TargetPid, Path, Arguments, CheckCallback)` |
| **IOCTL** | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80E) |
| **请求结构体** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` |
| **Action** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH` (0) |
| **关键字段** | `ProcessId`, `IsStartingNewProcess`(=TRUE if path provided), `CheckCallbackAtFirstInstruction` |

```cpp
// C++ 调用方式:
UdAttachToProcess(TargetPid, NULL, NULL, FALSE);  // attach by pid
UdAttachToProcess(NULL, Path, Arguments, FALSE);  // start new process
```

### 2.2 .detach - 分离进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/detach.cpp` |
| **核心函数** | `UdDetachProcess(ProcessId, Token)` |
| **IOCTL** | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80E) |
| **请求结构体** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` |
| **Action** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH` (1) |

### 2.3 .kill - 终止进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/kill.cpp` |
| **核心函数** | `UdKillProcess(TargetPid)` |
| **IOCTL** | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80E) |
| **请求结构体** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` |
| **Action** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS` (3) |

```cpp
// C++ 完整实现:
BOOLEAN UdKillProcess(UINT32 TargetPid) {
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS KillRequest = {0};
    
    // 先尝试用户模式终止
    UdTerminateProcessByPid(TargetPid);
    
    // 如果还在，用内核模式杀
    if (UdDoesProcessExistByPid(TargetPid)) {
        KillRequest.Action = DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS;
        KillRequest.ProcessId = TargetPid;
        
        DeviceIoControl(g_DeviceHandle,
            IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
            &KillRequest,
            SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
            &KillRequest,
            SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
            &ReturnedLength, NULL);
    }
}
```

### 2.4 .start - 启动进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/start.cpp` |
| **核心函数** | `UdAttachToProcess(NULL, Path, Arguments, FALSE)` |
| **IOCTL** | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80E) |
| **请求结构体** | `DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` |
| **Action** | `ATTACH` (0), `IsStartingNewProcess`=TRUE |

### 2.5 .restart - 重启进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/restart.cpp` |
| **逻辑** | 先 kill 当前进程，再重新 start |

### 2.6 .connect - 连接调试器

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/connect.cpp` |
| **IOCTL** | **无直接IOCTL** (本地连接仅设置标志位) |
| **实现方式** | 本地: `g_IsConnectedToHyperDbgLocally = TRUE`; 远程: `RemoteConnectionConnect(Ip, Port)` |

### 2.7 .disconnect - 断开连接

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/disconnect.cpp` |
| **IOCTL** | **无直接IOCTL** |
| **实现方式** | 设置标志位 + 关闭远程连接句柄 |

### 2.8 .listen - 监听连接

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/listen.cpp` |
| **IOCTL** | **无直接IOCTL** |
| **实现方式** | `RemoteConnectionListen(Port)` |

### 2.9 .switch - 切换调试线程/进程

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/switch.cpp` |
| **核心函数** | `UdShowListActiveDebuggingProcessesAndThreads()` / `UdSetActiveDebuggingThreadByPidOrTid(Id, IsThread)` |
| **IOCTL(列表)** | `IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES` (0x818) 或 `IOCTL_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS` (0x81A) |
| **IOCTL(切换)** | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (0x80E) |
| **Action(切换)** | `SWITCH_BY_PROCESS_OR_THREAD` (6) |

### 2.10 .status - 显示状态

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/status.cpp` |
| **IOCTL** | **无IOCTL** (读取全局变量状态) |

### 2.11 .process / .process2 - 进程操作

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/process.cpp` |
| **IOCTL** | **VMI模式**: 无直接IOCTL (使用符号引擎读取 EPROCESS 链表); **Debugger模式**: 通过串口发送包 |
| **实现方式** | VMI: `ObjectShowProcessesOrThreadDetails()` / `ObjectShowProcessesOrThreadList()`; Debugger: `KdSendSwitchProcessPacketToDebuggee()` |

### 2.12 .thread / .thread2 - 线程操作

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/thread.cpp` |
| **IOCTL** | 同 .process (使用符号引擎或串口通信) |

### 2.13 .sym - 符号操作

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/sym.cpp` |
| **IOCTL** | **无IOCTL** (用户态 PDB 解析) |
| **实现方式** | `SymbolBuildAndShowSymbolTable()`, `SymbolLoadOrDownloadSymbols()`, `ScriptEngineLoadFileSymbolWrapper()` 等 |

### 2.14 .sympath - 符号路径

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/sympath.cpp` |
| **IOCTL** | **无IOCTL** (配置文件读写) |

### 2.15 .script - 脚本执行

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/script.cpp` |
| **IOCTL** | **无IOCTL** (脚本解释器本地执行) |

### 2.16 .debug - 内核串口调试

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/debug.cpp` |
| **IOCTL** | **无直接IOCTL** (串口/NamedPipe 连接) |
| **实现方式** | `KdPrepareAndConnectDebugPort()` |

### 2.17 .logopen / .logclose - 日志

| 项目 | 值 |
|------|-----|
| **C++源文件** | `logopen.cpp`, `logclose.cpp` |
| **IOCTL** | **无IOCTL** (文件 I/O) |

### 2.18 .dump / !dump - 内存转储

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/dump.c` |
| **IOCTL** | `IOCTL_DEBUGGER_READ_MEMORY` (0x803) |
| **请求结构体** | `DEBUGGER_READ_MEMORY` |
| **关键字段** | `Address`, `Size`, `Pid`, `MemoryType`(VIRTUAL/PHYSICAL), `ReadingType`(KERNEL) |

### 2.19 .pe - PE解析

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/pe.cpp` |
| **IOCTL** | **无IOCTL** (本地PE解析) |

### 2.20 .pagein - 页面调入

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/pagein.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_BRING_PAGES_IN` (0x81F) |
| **请求结构体** | `DEBUGGER_PAGE_IN_REQUEST` |
| **关键字段** | `VirtualAddressFrom`, `VirtualAddressTo`, `ProcessId`, `PageFaultErrorCode` |

### 2.21 .formats - 格式化显示

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/formats.cpp` |
| **IOCTL** | **无IOCTL** (表达式求值+格式化输出) |

### 2.22 .cls - 清屏

| 项目 | 值 |
|------|-----|
| **C++源文件** | `meta-commands/cls.c` |
| **IOCTL** | **无IOCTL** (`system("cls")`) |

---

## 三、Debugging Commands 映射表

### 3.1 bp - 设置断点

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/bp.cpp` |
| **IOCTL** | `IOCTL_SET_BREAKPOINT_USER_DEBUGGER` (0x825) |
| **请求结构体** | `DEBUGGEE_BP_PACKET` |
| **关键字段** | `Address`, `Pid`, `Tid`, `Core`, `Result` |

```cpp
// C++ 实现关键代码:
BOOLEAN CommandBpPerformApplyingBreakpointOnUserDebugger(DEBUGGEE_BP_PACKET * BpPacket) {
    Status = DeviceIoControl(
        g_DeviceHandle,
        IOCTL_SET_BREAKPOINT_USER_DEBUGGER,  // 0x825
        BpPacket,
        SIZEOF_DEBUGGEE_BP_PACKET,
        BpPacket,
        SIZEOF_DEBUGGEE_BP_PACKET,
        &ReturnedLength, NULL);
}
```

### 3.2 bc - 清除断点

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/bc.cpp` |
| **IOCTL** | **通过串口发送包** (非直接IOCTL) |
| **请求结构体** | `DEBUGGEE_BP_LIST_OR_MODIFY_PACKET` |
| **Request类型** | `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR` |

### 3.3 bd - 禁用断点

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/bd.cpp` |
| **请求结构体** | `DEBUGGEE_BP_LIST_OR_MODIFY_PACKET` |
| **Request类型** | `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE` |

### 3.4 be - 启用断点

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/be.cpp` |
| **请求结构体** | `DEBUGGEE_BP_LIST_OR_MODIFY_PACKET` |
| **Request类型** | `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE` |

### 3.5 bl - 列出断点

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/bl.cpp` |
| **请求结构体** | `DEBUGGEE_BP_LIST_OR_MODIFY_PACKET` |
| **Request类型** | `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS` |

### 3.6 g - 继续执行

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/g.cpp` |
| **IOCTL(VMI远程)** | 发送 "g" 命令字符串 |
| **IOCTL(User Debugger)** | `UdContinueProcess(Token)` → 通过用户态调试器API |

### 3.7 p / pr - 单步跳过 (Step Over)

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/p.cpp` |
| **核心函数** | `SteppingStepOver()` |
| **IOCTL** | 用户态调试器单步机制 (INT3/异常) |

### 3.8 t / tr - 单步进入 (Step In)

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/t.cpp` |
| **核心函数** | `SteppingRegularStepIn()` |
| **IOCTL** | 用户态调试器单步机制 (TF标志) |

### 3.9 r - 寄存器读写

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/r.cpp` |
| **读全部寄存器结构体** | `DEBUGGEE_REGISTER_READ_DESCRIPTION` + `GUEST_REGS` + `GUEST_EXTRA_REGISTERS` |
| **读单个寄存器结构体** | `DEBUGGEE_REGISTER_READ_DESCRIPTION` { RegisterId, Value, KernelStatus } |
| **写寄存器结构体** | `DEBUGGEE_REGISTER_WRITE_DESCRIPTION` { RegisterId, Value, KernelStatus } |
| **通信方式** | 串口包 (`KdSendReadRegisterPacketToDebuggee` / `KdSendWriteRegisterPacketToDebuggee`) 或用户态调试器 (`UdSendReadRegisterToUserDebugger`) |

### 3.10 db/dc/dd/dq !db/!dc/!dd/!dq - 读内存

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/d-u.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_READ_MEMORY` (0x803) |
| **请求结构体** | `DEBUGGER_READ_MEMORY` |
| **关键字段** | `Address`, `Size`, `Pid`, `MemoryType`(VIRTUAL=0/PHYSICAL=1), `ReadingType`(KERNEL=0/VMX_ROOT=1), `AddressMode`(64BIT=1) |
| **显示风格** | `DEBUGGER_SHOW_COMMAND_DB`(5)/DC(6)/DD(7)/DQ(8)/DISASSEMBLE64(2)/DISASSEMBLE32(3)/DUMP(9) |

```cpp
// C++ 实现关键代码:
void HyperDbgShowMemoryOrDisassemble(ShowCommand Type, UINT64 Address, 
    MemoryType MemType, ReadingType ReadType, UINT32 Pid, UINT32 Length, PVOID Buffer) {
    DEBUGGER_READ_MEMORY ReadMemReq = {0};
    ReadMemReq.Pid = Pid;
    ReadMemReq.Address = Address;
    ReadMemReq.Size = Length;
    ReadMemReq.MemoryType = MemType;
    ReadMemReq.ReadingType = ReadType;
    // ... 发送 IOCTL_READ_MEMORY ...
}
```

### 3.11 u/u64/u2/u32 !u/!u64/!u2/!u32 - 反汇编

| 项目 | 值 |
|------|-----|
| **C++源文件** | 同 d-u.cpp |
| **IOCTL** | `IOCTL_DEBUGGER_READ_MEMORY` (0x803) + 反汇编引擎 |

### 3.12 eb/ed/eq !eb/!ed/!eq - 写内存

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/e.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_EDIT_MEMORY` (0x80A) |
| **请求结构体** | `DEBUGGER_EDIT_MEMORY` + 数据缓冲区 |
| **关键字段** | `ProcessId`, `Address`, `MemoryType`(EDIT_VIRTUAL/EDIT_PHYSICAL), `ByteSize`(EDIT_BYTE/DWORD/QWORD), `CountOf64Chunks`, `FinalStructureSize` |

```cpp
// C++ 结构布局:
typedef struct _DEBUGGER_EDIT_MEMORY {
    UINT32 Result;
    UINT64 Address;
    UINT32 ProcessId;
    DEBUGGER_EDIT_MEMORY_TYPE MemoryType;     // 0=VIRTUAL, 1=PHYSICAL
    DEBUGGER_EDIT_MEMORY_BYTE_SIZE ByteSize;  // 0=BYTE, 1=DWORD, 2=QWORD
    UINT32 CountOf64Chunks;
    UINT32 FinalStructureSize;
} DEBUGGER_EDIT_MEMORY;
// 后跟 CountOf64Chunks * UINT64 数据
```

### 3.13 sb/sd/sq !sb/!sd/!sq - 搜索内存

| 项目 | 值 |
|------|-----|
| **C++源文件** | `debugging-commands/s.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_SEARCH_MEMORY` (0x80B) |
| **请求结构体** | `DEBUGGER_SEARCH_MEMORY` + 搜索模式缓冲区 |
| **关键字段** | `Address`, `Length`, `ProcessId`, `MemoryType`(SEARCH_VIRTUAL/PHYSICAL/PHYS_FROM_VIRT), `ByteSize`(SEARCH_BYTE/DWORD/QWORD), `CountOf64Chunks`, `FinalStructureSize` |

```cpp
// C++ 结构布局:
typedef struct _DEBUGGER_SEARCH_MEMORY {
    UINT64 Address;
    UINT64 Length;
    UINT32 ProcessId;
    DEBUGGER_SEARCH_MEMORY_TYPE MemoryType;
    DEBUGGER_SEARCH_MEMORY_BYTE_SIZE ByteSize;
    UINT32 CountOf64Chunks;
    UINT32 FinalStructureSize;
} DEBUGGER_SEARCH_MEMORY;
// 后跟 CountOf64Chunks * UINT64 搜索模式
```

---

## 四、Extension Commands 映射表 (!开头)

### 4.1 !pte - 页表查询

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/pte.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS` (0x805) |
| **请求结构体** | `DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS` |
| **关键字段** | `VirtualAddress`, `ProcessId` |
| **返回字段** | `Pml4eVirtualAddress/Value`, `PdpteVirtualAddress/Value`, `PdeVirtualAddress/Value`, `PteVirtualAddress/Value`, `KernelStatus` |

```cpp
typedef struct _DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS {
    UINT64 VirtualAddress;
    UINT32 ProcessId;
    UINT64 Pml4eVirtualAddress;
    UINT64 Pml4eValue;
    UINT64 PdpteVirtualAddress;
    UINT64 PdpteValue;
    UINT64 PdeVirtualAddress;
    UINT64 PdeValue;
    UINT64 PteVirtualAddress;
    UINT64 PteValue;
    UINT32 KernelStatus;
} DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS;
```

### 4.2 !va2pa - 虚拟地址转物理地址

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/va2pa.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS` (0x809) |
| **请求结构体** | `DEBUGGER_VA2PA_AND_PA2VA_COMMANDS` |
| **关键字段** | `VirtualAddress`, `ProcessId`, `IsVirtual2Physical`=TRUE |
| **返回字段** | `PhysicalAddress`, `KernelStatus` |

### 4.3 !pa2va - 物理地址转虚拟地址

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/pa2va.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS` (0x809) |
| **请求结构体** | `DEBUGGER_VA2PA_AND_PA2VA_COMMANDS` |
| **关键字段** | `PhysicalAddress`, `ProcessId`, `IsVirtual2Physical`=FALSE |
| **返回字段** | `VirtualAddress`, `KernelStatus` |

```cpp
typedef struct _DEBUGGER_VA2PA_AND_PA2VA_COMMANDS {
    UINT64 VirtualAddress;
    UINT64 PhysicalAddress;
    UINT32 ProcessId;
    BOOLEAN IsVirtual2Physical;
    UINT32 KernelStatus;
} DEBUGGER_VA2PA_AND_PA2VA_COMMANDS;
```

### 4.4 !cpuid - CPUID事件

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/cpuid.cpp` |
| **IOCTL** | `IOCTL_DEBUGGER_REGISTER_EVENT` (0x806) + `IOCTL_ADD_ACTION_TO_EVENT` (0x807) |
| **事件类型** | `CPUID_INSTRUCTION_EXECUTION` |
| **实现方式** | `InterpretGeneralEventAndActionsFields()` → `SendEventToKernel()` → `RegisterActionToEvent()` |

### 4.5 !msrread - MSR读取事件

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/msrread.cpp` |
| **IOCTL** | 同上 (事件注册机制) |
| **事件类型** | `RDMSR_INSTRUCTION_EXECUTION` |

### 4.6 !dr - 调试寄存器访问事件

| 项目 | 值 |
|------|-----|
| **C++源文件** | `extension-commands/dr.cpp` |
| **IOCTL** | 同上 (事件注册机制) |
| **事件类型** | `DEBUG_REGISTERS_ACCESSED` |

### 4.7 !msrwrite - MSR写入事件

同 !msrread 模式，事件类型 `WRMSR_INSTRUCTION_EXECUTION`

### 4.8 !rdmsr / !wrmsr - 直接读写MSR

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_DEBUGGER_READ_OR_WRITE_MSR` (0x804) |
| **请求结构体** | `DEBUGGER_READ_AND_WRITE_ON_MSR` |
| **关键字段** | `Msr`(实际是32位), `CoreNumber`(0xFFFFFFFF=ALL), `ActionType`(READ=0/WRITE=1), `Value` |

```cpp
typedef struct _DEBUGGER_READ_AND_WRITE_ON_MSR {
    UINT64 Msr;          // MSR 地址 (32-bit value stored in 64-bit)
    UINT32 CoreNumber;   // 核心号 (0xFFFFFFFF = all cores)
    DEBUGGER_MSR_ACTION_TYPE ActionType;  // READ=0, WRITE=1
    UINT64 Value;        // 写入值 (仅WRITE时有效)
} DEBUGGER_READ_AND_WRITE_ON_MSR;
```

### 4.9 !crwrite - 写CR寄存器

事件类型: `WRITE_CR_REGISTER`

### 4.10 !ioin / !ioout - I/O端口读写

事件类型: `INSTRUCTION_IO_IN` / `INSTRUCTION_IO_OUT`

### 4.11 !epthook / !epthook2 - EPT Hook

事件类型: `HOOK_EPT_EXECUTION` / `HOOK_EPT_EXECUTION_DETOURS`

### 4.12 !monitor - 监控特定地址读写

事件类型: `MONITOR_READ_WRITE` / `MONITOR_EXECUTE`

### 4.13 !trace - 跟踪

事件类型: `TRACE_INSTRUCTIONS`

### 4.14 !hide / !unhide - 隐藏/显示调试器

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT` (0x808) |
| **请求结构体** | `DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE` |

### 4.15 !idt - IDT查询

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_QUERY_IDT_ENTRY` (0x824) |

### 4.16 !apic - APIC操作

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_PERFORM_ACTIONS_ON_APIC` (0x822) |

### 4.17 !pcitree / !pcicam - PCI枚举

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_PCIE_ENDPOINT_ENUM` (0x821) / `IOCTL_PCIDEVINFO_ENUM` (0x823) |

### 4.18 !rev - 反转内存

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_REQUEST_REV_MACHINE_SERVICE` (0x81E) |
| **请求结构体** | `REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST` |

### 4.19 !tsc - 时间戳计数器

事件类型相关

### 4.20 !pmc - 性能监控计数器

事件类型相关

### 4.21 !lbr - 最后分支记录

事件类型相关

### 4.22 !exception - 异常事件

事件类型: `EXCEPTION_OCCURRED`

### 4.23 !interrupt - 中断事件

事件类型: `INTERRUPT_OCCURRED`

### 4.24 !syscall-sysret - 系统调用事件

事件类型: `SYSCALL_SYSRET_EXECUTED`

### 4.25 !mode - VMI模式切换

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_PREACTIVATE_FUNCTIONALITY` (0x820) |
| **请求结构体** | `DEBUGGER_PREACTIVATE_COMMAND` { Type=MODE } |

### 4.26 !vmcall - VMCALL事件

事件类型: `VMCALL_INSTRUCTION_EXECUTED`

### 4.27 !xsetbv - XSETBV事件

事件类型: `XSETBV_INSTRUCTION_EXECUTED`

### 4.28 !smi - SMI操作

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_PERFORM_SMI_OPERATION` (0x826) |

### 4.29 !ioapic - IOAPIC操作

事件类型或IOCTL

### 4.30 !measure - 测量操作

事件类型相关

### 4.31 !hypertrace - HyperTrace操作

| 项目 | 值 |
|------|-----|
| **IOCTL** | `IOCTL_PERFORM_HYPERTRACE_OPERATION` (0x827) |

---

## 五、核心数据结构 Go 定义参考

### 5.1 DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS (最常用)

```go
type AttachDetachUserModeProcess struct {
    IsStartingNewProcess           bool    // offset 0   (1 byte)
    _                              byte    // padding
    ProcessId                      uint32  // offset 4
    ThreadId                       uint32  // offset 8
    CheckCallbackAtFirstInstruction bool    // offset 12  (1 byte)
    _                              byte    // padding
    Is32Bit                        bool    // offset 14  (1 byte)
    _                              byte    // padding
    Rip                            uint64  // offset 16
    InstructionBytesOnRip         [15]byte // offset 24 (MAXIMUM_INSTR_SIZE)
    SizeOfInstruction             uint32  // offset 39
    IsPaused                       bool    // offset 43  (1 byte)
    _                              [3]byte // padding to align
    Action                         uint32  // offset 47 → 实际 offset 48 (aligned)
    CountOfActiveDebuggingThreads  uint32  // offset 52
    Token                          uint64  // offset 56
    Result                         uint64  // offset 64
    // Total size: 72 bytes
}

const (
    ActionAttach       = 0
    ActionDetach       = 1
    ActionRemoveHooks  = 2
    ActionKillProcess  = 3
    ActionContinue     = 4
    ActionPause        = 5
    ActionSwitch       = 6
    ActionQueryCount   = 7
)
```

### 5.2 DEBUGGER_READ_MEMORY

```go
type DebuggerReadMemory struct {
    Pid            uint32  // offset 0
    Address        uint64  // offset 4 (aligned)
    Size           uint32  // offset 12
    GetAddressMode bool    // offset 16
    _              byte    // padding
    AddressMode    uint8   // offset 18 (0=32bit, 1=64bit)
    _              byte    // padding
    MemoryType     uint8   // offset 20 (0=PHYSICAL, 1=VIRTUAL)
    _              byte    // padding
    ReadingType    uint8   // offset 22 (0=KERNEL, 1=VMX_ROOT)
    ReturnLength   uint32  // offset 24 (aligned)
    KernelStatus   uint32  // offset 28
    // Total: 32 bytes, 后跟实际数据缓冲区
}
```

### 5.3 DEBUGGER_EDIT_MEMORY

```go
type DebuggerEditMemory struct {
    Result             uint32  // offset 0
    Address            uint64  // offset 4 (aligned)
    ProcessId          uint32  // offset 12
    MemoryType         uint32  // offset 16 (0=EDIT_VIRTUAL, 1=EDIT_PHYSICAL)
    ByteSize           uint32  // offset 20 (0=BYTE, 1=DWORD, 2=QWORD)
    CountOf64Chunks    uint32  // offset 24
    FinalStructureSize uint32  // offset 28
    // Total: 32 bytes, 后跟 CountOf64Chunks * uint64 数据
}
```

### 5.4 DEBUGGER_SEARCH_MEMORY

```go
type DebuggerSearchMemory struct {
    Address            uint64  // offset 0
    Length             uint64  // offset 8
    ProcessId          uint32  // offset 16
    MemoryType         uint32  // offset 20 (0=VIRTUAL, 1=PHYSICAL, 2=PHYS_FROM_VIRT)
    ByteSize           uint32  // offset 24 (0=BYTE, 1=DWORD, 2=QWORD)
    CountOf64Chunks    uint32  // offset 28
    FinalStructureSize uint32  // offset 32
    // Total: 36 bytes (aligned to 40?), 后跟搜索模式
}
```

### 5.5 DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS

```go
type DebuggerPageTableEntriesDetails struct {
    VirtualAddress      uint64 // offset 0
    ProcessId           uint32 // offset 8
    _                   uint32 // padding
    Pml4eVirtualAddress uint64 // offset 16
    Pml4eValue          uint64 // offset 24
    PdpteVirtualAddress uint64 // offset 32
    PdpteValue          uint64 // offset 40
    PdeVirtualAddress   uint64 // offset 48
    PdeValue            uint64 // offset 56
    PteVirtualAddress   uint64 // offset 64
    PteValue            uint64 // offset 72
    KernelStatus        uint32 // offset 80
    // Total: 84 bytes
}
```

### 5.6 DEBUGGER_VA2PA_AND_PA2VA_COMMANDS

```go
type DebuggerVa2paPa2vaCommands struct {
    VirtualAddress     uint64 // offset 0
    PhysicalAddress    uint64 // offset 8
    ProcessId          uint32 // offset 16
    IsVirtual2Physical bool   // offset 20
    _                  byte   // padding
    KernelStatus       uint32 // offset 24
    // Total: 28 bytes (aligned to 32?)
}
```

### 5.7 DEBUGGER_READ_AND_WRITE_ON_MSR

```go
type DebuggerReadWriteOnMsr struct {
    Msr        uint64 // offset 0 (实际32位值存于低32位)
    CoreNumber uint32 // offset 8
    ActionType uint32 // offset 12 (0=READ, 1=WRITE)
    Value      uint64 // offset 16 (仅WRITE时有效)
    // Total: 24 bytes
}
```

### 5.8 DEBUGGER_PAGE_IN_REQUEST

```go
type DebuggerPageInRequest struct {
    VirtualAddressFrom  uint64 // offset 0
    VirtualAddressTo    uint64 // offset 8
    ProcessId           uint32 // offset 16
    PageFaultErrorCode  uint32 // offset 20
    KernelStatus        uint32 // offset 24
    // Total: 28 bytes
}
```

### 5.9 DEBUGGEE_BP_PACKET (用户调试器断点)

```go
type DebuggeeBpPacket struct {
    Address uint64 // offset 0
    Core    uint32 // offset 8
    Pid     uint32 // offset 12
    Tid     uint32 // offset 16
    Result  uint64 // offset 20 (aligned)
    // Total: 28 bytes
}
```

---

## 六、Go 实现注意事项

### 6.1 结构体内存对齐
- C++ 结构体在 Windows 上默认按自然边界对齐
- Go 结构体也需要注意字段顺序和对齐
- 使用 `unsafe.Offsetof` 验证偏移量是否正确

### 6.2 IOCTL 缓冲区方式
- 所有 IOCTL 都使用 `METHOD_BUFFERED` 方式
- 输入和输出可以共享同一块缓冲区（很多C++代码就是这么做的）

### 6.3 两种调试模式
- **VMI Mode**: 直接与内核驱动通信 (DeviceIoControl)
- **Debugger Mode**: 通过串口/NamedPipe 与远程 debuggee 通信

### 6.4 事件类命令的特殊性
- `!cpuid`, `!msrread`, `!dr`, `!epthook` 等都是**事件注册**机制
- 使用两步: 先 `IOCTL_DEBUGGER_REGISTER_EVENT` 注册事件, 再 `IOCTL_ADD_ACTION_TO_EVENT` 添加动作
- 这些命令不是简单的请求-响应模式

### 6.5 无IOCTL的命令
以下命令不涉及驱动通信，纯用户态操作：
- `.connect` / `.disconnect` / `.listen` - 网络连接管理
- `.status` - 状态查询
- `.sym` / `.sympath` - 符号引擎
- `.script` - 脚本解释器
- `.formats` - 表达式求值
- `.cls` - 清屏
- `.pe` - PE解析
- `.help` - 帮助
- `.logopen` / `.logclose` - 日志

---

## 七、命令分类总结

| 分类 | 命令数量 | 主要IOCTL |
|------|---------|-----------|
| Meta Commands | 22个 | 0x80E (Attach/Detach), 0x81F (PageIn), 0x803 (ReadMem/Dump) |
| Debugging Commands | 13个 | 0x803 (ReadMem), 0x80A (EditMem), 0x80B (Search), 0x825 (BP) |
| Extension Commands | 31个 | 0x805 (PTE), 0x809 (VA2PA), 0x804 (MSR), 0x806/0x807 (Events) |

---

## 八、待实现的Go方法清单 (core/debugger.go)

基于以上分析，以下是需要实现的核心方法：

### 进程/线程管理 (IOCTL 0x80E)
- [ ] `AttachToProcess(pid uint32, path string, checkCallback bool) error`
- [ ] `DetachFromProcess(pid uint32, token uint64) error`
- [ ] `KillProcess(pid uint32) error`
- [ ] `StartProcess(path, args string) (uint32, error)`
- [ ] `ContinueProcess(pid uint32, token uint64) error`
- [ ] `PauseProcess(pid uint32, token uint64) error`
- [ ] `SwitchToProcessOrThread(pid, tid uint32, rip uint64, instr []byte, instrSize uint32, isPaused bool) error`
- [ ] `QueryActiveDebuggingThreadCount() (uint32, error)`
- [ ] `ListActiveProcesses() ([]ProcessInfo, error)`

### 内存操作
- [ ] `ReadMemory(address, size uint64, pid uint32, memType, readType uint8) ([]byte, error)` — IOCTL 0x803
- [ ] `EditMemory(address uint64, pid uint32, memType, byteSize uint8, data []byte) error` — IOCTL 0x80A
- [ ] `SearchMemory(address, length uint64, pid uint32, memType, byteSize uint8, pattern []byte) ([]uint64, error)` — IOCTL 0x80B

### 断点 (IOCTL 0x825)
- [ ] `SetBreakpoint(address uint64, pid, tid, core uint32) error`

### 页表/地址转换
- [ ] `GetPteDetails(virtualAddr uint64, pid uint32) (*PageTableEntryDetails, error)` — IOCTL 0x805
- [ ] `Va2pa(virtualAddr uint64, pid uint32) (uint64, error)` — IOCTL 0x809
- [ ] `Pa2va(physicalAddr uint64, pid uint32) (uint64, error)` — IOCTL 0x809

### MSR (IOCTL 0x804)
- [ ] `ReadMsr(msr, core uint32) (uint64, error)`
- [ ] `WriteMsr(msr, core uint32, value uint64) error`

### 其他
- [ ] `BringPagesIn(from, to uint64, pid uint32, pageFaultCode uint32) error` — IOCTL 0x81F
- [ ] `FlushBuffers() (countVmRoot, countVmNonRoot uint32, error)` — IOCTL 0x80D
