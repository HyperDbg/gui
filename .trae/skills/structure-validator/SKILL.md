---
name: "structure-validator"
description: "Validates structure layouts for kernel communication to prevent BSOD. Invoke when implementing new structures or modifying IOCTL communication code."
---

# Structure Validator for Kernel Communication

This skill ensures all structures used in kernel communication are properly validated to prevent BSOD (Blue Screen of Death) caused by memory layout mismatches.

## 📊 实时验证进度

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                        STRUCTURE VALIDATOR PROGRESS                          │
├─────────────────────────────────────────────────────────────────────────────┤
│  结构体定义总数:     54 个                                                   │
│  IOCTL 定义总数:     46 个                                                   │
│  方法总数:           1023 个                                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│  ✅ 已验证结构布局:  20/54  (37%)                                            │
│  ✅ 已实现 Validate():    20/54  (37%)                                       │
│  ✅ 已实现 ExpectedSize(): 20/54  (37%)                                      │
│  ✅ 单元测试覆盖:    20/54  (37%)                                            │
├─────────────────────────────────────────────────────────────────────────────┤
│  ⚠️  手动构造 buffer:  101 处 (高风险 BSOD)                                   │
│     - packet.go:         67 处                                               │
│     - user_debugger.go:  11 处                                               │
│     - event_handler.go:  17 处                                               │
│     - kernel_debugger.go: 3 处                                               │
│     - hardware/cpu.go:   3 处                                                │
├─────────────────────────────────────────────────────────────────────────────┤
│  ❌ 未实现接口签名:    34 个结构体                                            │
│  ❌ 未添加单元测试:    34 个结构体                                            │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 已验证结构列表 (20/54)

| # | 结构体名称 | 大小 | Validate() | ExpectedSize() | 测试 |
|---|-----------|------|:----------:|:--------------:|:----:|
| 1 | DebuggerAttachDetachUserModeProcess | 72B | ✅ | ✅ | ✅ |
| 2 | DebuggeeBpPacket | 32B | ✅ | ✅ | ✅ |
| 3 | DebuggerUdCommandAction | 40B | ✅ | ✅ | ✅ |
| 4 | DebuggerUdCommandPacket | 64B | ✅ | ✅ | ✅ |
| 5 | DebuggerReadMemory | 48B | ✅ | ✅ | ✅ |
| 6 | DebuggerVa2paAndPa2vaCommands | 32B | ✅ | ✅ | ✅ |
| 7 | DebuggerGeneralEventDetails | 24B | ✅ | ✅ | ✅ |
| 8 | DebuggerEventAction | 56B | ✅ | ✅ | ✅ |
| 9 | DebuggerModifyEventRequest | 24B | ✅ | ✅ | ✅ |
| 10 | DebuggerEventAndActionResult | 8B | ✅ | ✅ | ✅ |
| 11 | DebuggeeKdPausedPacket | 72B | ✅ | ✅ | ✅ |
| 12 | DebuggeeDetailsAndSwitchProcessPacket | 56B | ✅ | ✅ | ✅ |
| 13 | DebuggeeDetailsAndSwitchThreadPacket | 56B | ✅ | ✅ | ✅ |
| 14 | DebuggeeRegisterReadDescription | 24B | ✅ | ✅ | ✅ |
| 15 | PageTableEntries | 80B | ✅ | ✅ | ✅ |
| 16 | DebuggerReadPageTableEntriesDetails | 88B | ✅ | ✅ | ✅ |
| 17 | CoreInfo | 16B | ✅ | ✅ | ✅ |
| 18 | VersionInfo | 20B | ✅ | ✅ | ✅ |
| 19 | UserModeProcessDetails | 48B | ✅ | ✅ | ✅ |
| 20 | DebuggerSetBreakpointUserDebugger | 40B | ✅ | ✅ | ✅ |

### 未验证结构列表 (34/54)

| # | 结构体名称 | 需要验证 |
|---|-----------|---------|
| 1 | DebuggerReadPageTableEntriesDetails | ❌ |
| 2 | DebuggerWriteMemory | ❌ |
| 3 | EventAndActionResult | ❌ |
| 4 | FlushResult | ❌ |
| 5 | ProcessDetails | ❌ |
| 6 | ThreadDetails | ❌ |
| 7 | IDTEntry | ❌ |
| 8 | APICInfo | ❌ |
| 9 | PCIDevice | ❌ |
| 10 | PCIDeviceInfo | ❌ |
| 11 | ExportInfo | ❌ |
| 12 | LogMessage | ❌ |
| 13 | ScriptConfiguration | ❌ |
| 14 | ScriptResult | ❌ |
| 15 | ReconstructionStatus | ❌ |
| 16 | EPTHook | ❌ |
| 17 | Monitor | ❌ |
| 18 | Breakpoint | ❌ |
| 19 | StackFrame | ❌ |
| 20 | ModuleInfo | ❌ |
| 21 | ProcessInfo | ❌ |
| 22 | ThreadInfo | ❌ |
| 23 | TestResults | ❌ |
| 24 | TestDetail | ❌ |
| 25 | TestStatus | ❌ |
| 26 | Configuration | ❌ |
| 27 | SymbolInfo | ❌ |
| 28 | HWDBGInstanceInformation | ❌ |
| 29 | HWDBGScriptCapabilities | ❌ |
| 30 | HWDBGScriptBuffer | ❌ |
| 31 | HWDBGPortInformationItems | ❌ |
| 32 | PCICamInfo | ❌ |
| 33 | CommandResult | ❌ |
| 34 | CommandEvent | ❌ |

### 手动构造 Buffer 详细位置 (101处 - 高风险)

#### packet.go (67处)

| 行号 | 方法 | IOCTL |
|------|------|-------|
| 412-414 | EPTHook | IOCTL_DEBUGGER_REGISTER_EVENT |
| 450 | HookSyscall | IOCTL_DEBUGGER_REGISTER_EVENT |
| 470 | HookException | IOCTL_DEBUGGER_REGISTER_EVENT |
| 490 | HookInterrupt | IOCTL_DEBUGGER_REGISTER_EVENT |
| 510-511 | HookIO | IOCTL_DEBUGGER_REGISTER_EVENT |
| 531 | HookIOAPIC | IOCTL_DEBUGGER_REGISTER_EVENT |
| 544 | ReadMsr | IOCTL_DEBUGGER_READ_OR_WRITE_MSR |
| 560-561 | WriteMsr | IOCTL_DEBUGGER_READ_OR_WRITE_MSR |
| 582 | MeasurePerformance | IOCTL_DEBUGGER_REGISTER_EVENT |
| 606-608 | MonitorMemory | IOCTL_DEBUGGER_REGISTER_EVENT |
| 629-631 | PCICam | IOCTL_DEBUGGER_REGISTER_EVENT |
| 661 | PMC | IOCTL_DEBUGGER_REGISTER_EVENT |
| 684-687 | ReconstructMemory | IOCTL_DEBUGGER_REGISTER_EVENT |
| 715-717 | SearchMemoryPattern | IOCTL_DEBUGGER_REGISTER_EVENT |
| 749 | InstructionTrace | IOCTL_DEBUGGER_REGISTER_EVENT |
| 769-770 | TrackMemory | IOCTL_DEBUGGER_REGISTER_EVENT |
| 845-846 | PTE | IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS |
| 880-882 | VA2PA | IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS |
| 905-906 | PA2VA | IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS |
| 929 | ReadMemory | IOCTL_DEBUGGER_READ_MEMORY |
| 966 | ChangeThread | IOCTL_SWITCH_THREAD |
| 1033 | AttachProcess | IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS |
| 1100 | APIC | IOCTL_PERFORM_ACTIONS_ON_APIC |
| 1163 | PerformSMIOperation | IOCTL_PERFORM_SMI_OPERATION |
| 1183-1184 | HideDebugger | IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER |
| 1204-1205 | UnhideDebugger | IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER |
| 1224-1226 | BringPagesIn | IOCTL_DEBUGGER_BRING_PAGES_IN |
| 1246-1248 | EditMemory | IOCTL_DEBUGGER_EDIT_MEMORY |
| 1261-1264 | SearchMemory | IOCTL_DEBUGGER_SEARCH_MEMORY |
| 1277-1279 | WriteMemory | IOCTL_DEBUGGER_EDIT_MEMORY |
| 1300-1303 | SearchMemoryPattern | IOCTL_DEBUGGER_SEARCH_MEMORY |
| 1343-1344 | DetachProcess | IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS |
| 1384 | Processes | IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES |
| 1406 | Threads | IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES |

#### user_debugger.go (11处)

| 行号 | 方法 | IOCTL |
|------|------|-------|
| 1281-1283 | StepInto | IOCTL_SEND_USER_DEBUGGER_COMMANDS |
| 1362-1365 | ReadMemory | IOCTL_DEBUGGER_READ_MEMORY |
| 1415-1418 | WriteMemory | IOCTL_DEBUGGER_EDIT_MEMORY |

#### event_handler.go (17处)

| 行号 | 方法 | IOCTL |
|------|------|-------|
| 80-87 | RegisterEvent | IOCTL_DEBUGGER_REGISTER_EVENT |
| 123-131 | RegisterEventWithAction | IOCTL_DEBUGGER_REGISTER_EVENT |

#### kernel_debugger.go (3处)

| 行号 | 方法 | IOCTL |
|------|------|-------|
| 845-846 | PreallocatePools | IOCTL_RESERVE_PRE_ALLOCATED_POOLS |
| 972 | HaltCore | IOCTL_PAUSE_PACKET_RECEIVED |

#### hardware/cpu.go (3处)

| 行号 | 方法 | 说明 |
|------|------|------|
| 68-69 | CPUID | 非IOCTL，本地CPUID调用 |

## Overview

When communicating with Windows kernel drivers via IOCTL, structure layout must match exactly between Go and C/C++. Any mismatch in size, field offsets, or padding can cause system crashes.

## ⚠️ CRITICAL: Check C++ Source First

**BEFORE implementing any IOCTL communication, you MUST:**

1. **Check Go file comments** - 每个方法的注释已包含源代码路径，格式如下：
   ```
   // 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/...
   ```
2. **Check the IOCTL table** in `doc/HyperDbg_API_IOCTL_Table.md` to find the IOCTL code and handler function
3. **Read the C++ driver source code** to understand the exact structure definition
4. **Verify all fields** including Token, ProcessId, ThreadId, etc.

### 快速查找源代码

**方法一：从 Go 注释中获取（推荐）**

每个 IOCTL 相关方法的注释都包含源代码路径，例如：
```go
// StepInto 单步进入下一条指令
// IOCTL: 是, IOCTL_SEND_USER_DEBUGGER_COMMANDS (0x817)
// 源代码: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/debugging-commands/t.cpp
```

**方法二：搜索 C++ 头文件**

结构定义在 `doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/Headers/RequestStructures.h`

### Common Mistakes (DO NOT SKIP)

| Mistake | Example | Fix |
|---------|---------|-----|
| Missing Token field | `DebuggeeBpPacket` missing `ProcessDebuggingDetailToken` | Add the token field from C++ struct |
| Wrong field order | Fields not matching C++ definition | Reorder to match C++ struct exactly |
| Wrong calling sequence | ResumeThread after REMOVE_HOOKS loop | Check C++ source for correct sequence |
| Missing validation | Not checking C++ driver handler | Read driver code to understand requirements |

### Source Code Locations

```
doc/cpp/HyperDbgUnified/HyperDbg/
├── hyperdbg/
│   ├── include/SDK/Headers/          # Structure definitions
│   │   ├── BasicTypes.h
│   │   ├── RequestStructures.h       # IOCTL request structures
│   │   └── ErrorCodes.h
│   └── libhyperdbg/code/
│       ├── debugger/commands/        # User-mode command implementations
│       └── core/debugger.cpp         # ShowErrorMessage()
└── hyperkd/
    └── code/driver/Ioctl.c           # Driver IOCTL handlers
```

## Validation Process

### Step 0: Read C++ Source Code (MANDATORY)

**首先从 Go 文件注释中获取源代码路径：**

1. 打开对应的 Go 文件（`user_debugger.go`、`kernel_debugger.go`、`packet.go`）
2. 找到目标方法的注释，查看 `源代码:` 行
3. 根据路径读取 C++ 源文件

**结构定义位置：**
- 主要结构：`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/Headers/RequestStructures.h`
- 基本类型：`doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/Headers/BasicTypes.h`

**如果需要搜索：**
```bash
# 在头文件中搜索结构定义
grep -r "typedef struct.*STRUCTURE_NAME" doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/include/SDK/Headers/
```

### Step 1: Identify Structures

Analyze all structures used in:
- `debugger/user_debugger.go`
- `debugger/kernel_debugger.go`
- `debugger/packet.go`

For each structure, identify:
- Expected size (from C/C++ source)
- Field offsets
- Padding requirements

### Step 2: Create/Update Unit Tests

**当发现新的 IOCTL 通信结构时，必须更新 `debugger/structures_test.go`：**

1. 在 `TestAll` 函数中添加新的 `t.Run` 测试块
2. 设置正确的 `ExpectedSize` 和 `ExpectedFields`
3. 从 C++ 头文件获取字段偏移量

```go
func TestAll(t *testing.T) {
    t.Run("StructureName", func(t *testing.T) {
        layout := StructLayout{
            Name:         "StructureName",
            ExpectedSize: 72,  // 从 C++ sizeof 获取
            ExpectedFields: map[string]uintptr{
                "Field1": 0,   // 从 C++ 结构偏移获取
                "Field2": 8,
                // ...
            },
        }
        verifyStruct[StructureName](t, layout)
        verifySerialize(t, layout, StructureName{
            // test values
        })
    })
}
```

**如何获取 ExpectedSize 和 ExpectedFields：**
1. 从 C++ 头文件找到结构定义
2. 使用 `sizeof(STRUCTURE)` 获取大小
3. 使用 `offsetof(STRUCTURE, field)` 获取字段偏移

### Step 3: Implement Validation Interface

**在 `debugger/structures.go` 中为每个结构实现验证接口：**

```go
type Validator interface {
    Validate() error
}

type SizeVerifier interface {
    ExpectedSize() uintptr
    ExpectedSerSize() uintptr
}
```

**实现示例：**
```go
func (s *DebuggerAttachDetachUserModeProcess) Validate() error {
    if s.ProcessId == 0 {
        return errors.New("ProcessId cannot be zero")
    }
    if s.ThreadId == 0 {
        return errors.New("ThreadId cannot be zero")
    }
    return nil
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSize() uintptr {
    return 72  // 必须与 C++ sizeof 一致
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSerSize() uintptr {
    return 72  // 序列化后大小（通常与 ExpectedSize 相同）
}
```

### Step 4: Validate Before Sending

**CRITICAL**: 每个 IOCTL 发送方法必须在发送前调用验证接口。

#### 必须的验证模式

**禁止手动构造 buffer，必须使用结构体：**

❌ 错误做法：
```go
buffer := make([]byte, 64)
binary.LittleEndian.PutUint32(buffer[0:4], uint32(ActionType))
binary.LittleEndian.PutUint64(buffer[8:16], token)
```

✅ 正确做法：
```go
packet := DebuggerUdCommandPacket{
    UdAction: DebuggerUdCommandAction{
        ActionType: DebuggerUdCommandActionTypeRegularStep,
    },
    ProcessDebuggingDetailToken: token,
    TargetThreadId:              threadId,
}

if err := packet.Validate(); err != nil {
    mylog.Warning("验证失败", err)
    return
}

if unsafe.Sizeof(packet) != packet.ExpectedSize() {
    mylog.Warning("大小不匹配")
    return
}

buffer := new(bytes.Buffer)
binary.Write(buffer, binary.LittleEndian, &packet)
```

#### Validation Checklist for Every Send Method

- [ ] 使用结构体而非手动构造 buffer
- [ ] 调用 `packet.Validate()` 检查错误
- [ ] 调用 `unsafe.Sizeof(packet) == packet.ExpectedSize()`
- [ ] 验证失败时提前返回
- [ ] 记录有意义的错误信息

#### Example: Attach Process Validation

```go
func (s *UserDebug) StartProcess(path string) {
    // ... process creation ...

    attachReq := DebuggerAttachDetachUserModeProcess{
        IsStartingNewProcess: 1,
        ProcessId:            procInfo.ProcessId,
        ThreadId:             procInfo.ThreadId,
        Action:               DebuggerAttachDetachUserModeProcessActionAttach,
    }

    // MANDATORY validation
    if err := attachReq.Validate(); err != nil {
        mylog.Warning("attach请求验证失败", err)
        return
    }

    if unsafe.Sizeof(attachReq) != attachReq.ExpectedSize() {
        mylog.Warning("attach请求大小不匹配", "got", unsafe.Sizeof(attachReq), "want", attachReq.ExpectedSize())
        return
    }

    // Safe to send
    buffer := new(bytes.Buffer)
    binary.Write(buffer, binary.LittleEndian, &attachReq)
    response := s.packeter.DriverProvider.SendReceive(buffer, IoctlDebuggerAttachDetachUserModeProcess)
    // ...
}
```

## Structure Padding Rules

### Windows x64 Alignment

| Type | Alignment | Size |
|------|-----------|------|
| uint8 | 1 | 1 |
| uint16 | 2 | 2 |
| uint32 | 4 | 4 |
| uint64 | 8 | 8 |
| bool | 1 | 1 |
| arrays | element alignment | element size * count |

### Padding Example

```go
type Example struct {
    A uint8    // offset 0, size 1
    _ [3]byte  // padding 3 bytes for alignment
    B uint32   // offset 4, size 4
    C uint64   // offset 8, size 8
    D uint8    // offset 16, size 1
    _ [7]byte  // padding 7 bytes for total size multiple of 8
}
// Total: 24 bytes
```

## Common BSOD Causes

1. **Size mismatch**: Go struct size != C struct size
2. **Offset mismatch**: Field at wrong position
3. **Missing padding**: Alignment not preserved
4. **Invalid values**: Zero pointers, invalid IDs
5. **Buffer overflow**: Output buffer too small
6. **手动构造 buffer**: 字段偏移错误导致数据错位

## Testing Checklist

- [ ] 结构大小与 C++ 定义一致
- [ ] 所有字段偏移已验证
- [ ] 序列化/反序列化测试通过
- [ ] 验证方法已实现
- [ ] IOCTL 发送前添加了验证
- [ ] 边界情况已测试（零值、最大值）

## Files to Update

当添加或修改 IOCTL 通信时，需要更新以下文件：

| 文件 | 更新内容 |
|------|----------|
| `debugger/structures.go` | 添加结构定义、Validate()、ExpectedSize() 方法 |
| `debugger/structures_test.go` | 添加单元测试验证结构布局 |
| `debugger/user_debugger.go` | 使用结构体替代手动 buffer 构造 |
| `debugger/kernel_debugger.go` | 使用结构体替代手动 buffer 构造 |
| `debugger/packet.go` | 使用结构体替代手动 buffer 构造 |
