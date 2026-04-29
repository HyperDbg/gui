---
name: "hyperdbg-consistency-checker"
description: "Checks Go implementation consistency with C++ HyperDbg. Invoke when comparing Go test code with C++ Libhyperdbg implementation or debugging IOCTL structure mismatches."
---

# HyperDbg Go/C++ Consistency Checker

This skill helps verify that Go implementations are consistent with the original C++ HyperDbg codebase.

## When to Invoke

- User wants to compare Go test code with C++ implementation
- Debugging IOCTL structure mismatches
- Verifying struct layouts between Go and C++
- Checking if Go implementation follows C++ patterns
- Investigating BSOD crashes related to driver communication
- Fixing compilation errors after SDK bindgen regeneration
- Resolving type/field name differences between old code and new SDK

## Directory Structure Mapping

| C++ 目录 | Go 目录 | 完成度 |
|---------|---------|--------|
| `code/app/` | `debugger/app/` | 40% |
| `code/debugger/kernel-level/` | `debugger/kernel_level/` | ✅ 超集 |
| `code/debugger/user-level/` | `debugger/user_level/` | ✅ 超集 |
| `code/debugger/core/` | `debugger/core/` | ✅ 超集 |
| `code/debugger/communication/` | `debugger/communication/` | ✅ 超集 |
| `code/debugger/misc/` | `debugger/misc/` | ✅ 超集 |
| `code/debugger/script-engine/` | `debugger/script_engine/` | 78% |
| `code/debugger/transparency/` | `debugger/transparency/` | ✅ 超集 |
| `code/debugger/driver-loader/` | `debugger/driver_loader/` | 83% |
| `code/common/` | `debugger/common/` | 61% |
| `code/export/` | `debugger/export/` | 19% |
| `code/debugger/commands/` | `debugger/commands/` | 40% |
| `code/objects/` | `debugger/objects/` | ✅ 超集 |
| `code/rev/` | `debugger/rev/` | ✅ 超集 |
| — | `debugger/pdbex/` | Go独有 |

## Key Files to Compare

### Go Implementation
- `debugger/app/libhyperdbg.go` - 应用层入口
- `debugger/kernel_level/kd.go` - 内核调试器
- `debugger/kernel_level/kernel_listening.go` - 内核监听
- `debugger/user_level/ud.go` - 用户态调试器
- `debugger/user_level/user_listening.go` - 用户态监听
- `debugger/user_level/pe_parser.go` - PE解析器
- `debugger/core/debugger.go` - 核心调试逻辑
- `debugger/core/interpreter.go` - 命令解释器
- `debugger/core/steppings.go` - 单步执行
- `debugger/core/break_control.go` - 断点控制
- `debugger/communication/namedpipe.go` - 命名管道
- `debugger/communication/remote_connection.go` - 远程连接
- `debugger/communication/tcp.go` - TCP通信
- `debugger/communication/forwarding.go` - 事件转发
- `debugger/misc/assembler.go` - 汇编器
- `debugger/misc/disassembler.go` - 反汇编器
- `debugger/misc/callstack.go` - 调用栈
- `debugger/misc/pci_id.go` - PCI ID数据库
- `debugger/misc/readmem.go` - 内存读取
- `debugger/script_engine/script_engine.go` - 脚本引擎
- `debugger/script_engine/script_engine_wrapper.go` - 脚本引擎包装器
- `debugger/script_engine/symbol.go` - 符号表
- `debugger/transparency/transparency.go` - 透明模式
- `debugger/transparency/gaussian_rng.go` - 高斯随机数
- `debugger/driver_loader/install.go` - 驱动安装
- `debugger/common/common.go` - 通用工具
- `debugger/common/cpu.go` - CPU相关
- `debugger/common/list_spinlock.go` - 链表和自旋锁
- `debugger/export/types.go` - 导出类型
- `debugger/export/common_export.go` - 导出通用函数
- `debugger/commands/debugging/debugging_commands.go` - 调试命令
- `debugger/commands/extension/extension_commands.go` - 扩展命令
- `debugger/commands/meta/meta_commands.go` - 元命令
- `debugger/objects/objects.go` - 对象管理
- `debugger/rev/rev_ctrl.go` - 逆向控制
- `debugger/pdbex/` - PDB解析器(Go独有)
- `debugger/debugger_test.go` - 测试代码

### C++ Implementation
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/app/libhyperdbg.cpp` - 应用层入口
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/kernel-level/kd.cpp` - 内核调试器
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/kernel-level/kernel-listening.cpp` - 内核监听
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/user-level/ud.cpp` - 用户态调试器
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/user-level/user-listening.cpp` - 用户态监听
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/user-level/pe-parser.cpp` - PE解析器
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/debugger.cpp` - 核心调试逻辑
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/interpreter.cpp` - 命令解释器
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/steppings.cpp` - 单步执行
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/core/break-control.cpp` - 断点控制
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/communication/` - 通信模块
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/misc/` - 辅助模块
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/script-engine/` - 脚本引擎
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/transparency/` - 透明模式
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/driver-loader/install.cpp` - 驱动安装
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/common/common.cpp` - 通用工具
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/export/export.cpp` - 导出函数
- `HyperDbgUnified/HyperDbg/hyperdbg/libhyperdbg/code/debugger/commands/` - 命令模块
- `HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/code/driver/Ioctl.c` - 内核IOCTL处理

## Bindgen SDK Generator (2026-04-16 ✅ 已完成)

### 架构概述

Bindgen 是一个基于 `modernc.org/cc/v4` C 解析器的自动代码生成器，将 HyperDbg C 头文件转换为 Go 类型绑定。

**核心文件：** `debugger/bindgen/generate_test.go`

**输入：** `debugger/bindgen/SDK/headers/*.h` (13个头文件)

**输出：** `debugger/sdk/*.go` (12个Go绑定文件)

### BindgenConfig 多目标配置

```go
type BindgenConfig struct {
    Name        string
    HeadersDir  string
    OutputDir   string
    PackageName string
    HeaderOrder []string
}
```

**当前激活配置：**
- `sdk` — 输入: `SDK/headers/`, 输出: `../sdk/`, 包名: `sdk`

**待激活配置（已注释）：**
- `header` — 输入: `HyperDbgUnified/.../header/`, 输出: `../header/`, 包名: `header`（因C++语法问题暂未启用）

### 头文件 → Go文件映射

| C 头文件 | Go 输出文件 | 内容 |
|---------|------------|------|
| `BasicTypes.h` | `BasicTypes.go` | 基础类型别名 (UINT64→uint64等) |
| `Constants.h` | `Constants.go` | 常量宏 + 预制结构体 (BUFFER_HEADER等) |
| `ErrorCodes.h` | `ErrorCodes.go` | 错误码枚举 + String() |
| `Connection.h` | `Connection.go` | 远程通信结构体 + 枚举 |
| `DataTypes.h` | `DataTypes.go` | 调试器数据类型 |
| `Events.h` | `Events.go` | 事件注册/修改结构体 |
| `HardwareDebugger.h` | `HardwareDebugger.go` | hwdbg脚本能力结构体 |
| `Ioctls.h` | `Ioctls.go` | IOCTL控制码常量 |
| `Pcie.h` | `Pcie.go` | PCI设备信息结构体 |
| `RequestStructures.h` | `RequestStructures.go` | 请求/响应结构体 (最大文件) |
| `ScriptEngineCommonDefinitions.h` | `ScriptEngineCommonDefinitions.go` | 脚本引擎定义 |
| `Symbols.h` | `Symbols.go` | 符号表结构体 |
| `Assertions.h` | _(内联到其他文件)_ | 断言宏 |

### 核心生成机制

#### 1. 全局去重 (Global Deduplication)

使用 `safemap.M` 有序Map确保跨头文件的类型唯一性：

```go
type Result struct {
    Structs  *safemap.M[string, structInfo]
    Enums    *safemap.M[string, enumInfo]
    Typedefs *safemap.M[string, typedefInfo]
    Macros   *safemap.M[string, macroConstInfo]
    FnMacros *safemap.M[string, macroConstInfo]
}
```

#### 2. N-ary Tree 结构体解析

嵌套匿名结构体通过递归提取为独立类型：

```go
func generateStructFields(t *cc.StructType, structGoName string, isPacked bool) (fields, methods string, innerTypes []structInfo)
func generateStructFieldsInnerTypes(t *cc.StructType, structGoName string, forcePacked bool) []structInfo
```

- 有tag的嵌套结构体 → 提取为独立 `innerType`
- 递归处理多层嵌套
- 生成 getter/setter 方法访问嵌套字段

#### 3. 预制结构体 (Prebuilt Structs)

SDK头文件中引用但未定义的外部类型，手动注入：

```go
type prebuiltStruct struct {
    goName, cName, source, originalSource string
    fields                                string
    lineNo                                int
}
```

**已注入的预制结构体：**

| Go类型名 | C类型名 | 原始来源 | 输出文件 |
|---------|---------|---------|---------|
| `BUFFER_HEADER` | `_BUFFER_HEADER` | `hyperlog/header/Logging.h:57` | `Constants.go` |
| `DEBUGGER_EVENT_ACTION` | `_DEBUGGER_EVENT_ACTION` | `hyperkd/.../Debugger.h:79` | `Constants.go` |
| `DEBUGGER_EVENT` | `_DEBUGGER_EVENT` | `hyperkd/.../Debugger.h:107` | `Constants.go` |

- `source` 字段：分配到哪个输出文件（如 `Constants.h`）
- `originalSource` 字段：实际来源路径（用于注释追踪）

#### 4. 统一类型映射

```go
func mapPredefinedType(kind cc.Kind, fallback string) string
```

| C 类型 | Go 类型 |
|--------|---------|
| `BOOLEAN` | `bool` |
| `WCHAR` | `rune` |
| `time_t` | `time.Duration` |
| `char` | `int8` |
| `unsigned char` | `byte` |
| `size_t` | `uintptr` |

#### 5. 枚举生成策略

- **连续值从0开始** → 使用 `iota`
- **连续值从N开始** → 显式赋值（如 `= 1`, `= 2`, ...）
- **非连续值** → 逐个显式赋值
- 每个枚举自动生成 `String()` 方法

### 运行 Bindgen

```powershell
cd c:\Users\Admin\Desktop\hypedbg\debugger
go test -run TestGenerate -v -count=1 ./bindgen/
```

### 验证编译

```powershell
cd c:\Users\Admin\Desktop\hypedbg\debugger
go test -run TestFixError -v -count=1
```

## gen → sdk 迁移 (2026-04-16 ✅ 已完成)

### 迁移概述

将所有 Go 代码从旧的 `bindgen/gen` 包迁移到新生成的 `sdk` 包。

### 已迁移文件

| 文件 | 变更 |
|------|------|
| `kernel_level/kd.go` | `gen.` → `sdk.` + 类型修复 |
| `kernel_level/kernel_listening.go` | `gen.` → `sdk.` + bool/int8修复 |
| `user_level/ud.go` | `gen.` → `sdk.` + bool修复 |
| `user_level/user_listening.go` | `gen.` → `sdk.` |
| `core/debugger.go` | `gen.` → `sdk.` + Ioctl名修复 |
| `scriptengine/script_engine.go` | `gen.` → `sdk.` |
| `scriptengine/script_engine_wrapper.go` | `gen.` → `sdk.` |
| `commands/debugging/debugging_commands.go` | `gen.` → `sdk.` |
| `commands/extension/extension_commands.go` | `gen.` → `sdk.` + PCIE→PCI修复 |
| `misc/readmem.go` | `gen.` → `sdk.` |

### 已删除

- `debugger/bindgen/gen/` — 整个目录已删除

## SDK 差异修复模式 (2026-04-16)

当 `TestFixError` 编译失败时，遵循以下修复流程：

### 核心原则：**以绑定的SDK为准**

名字不同就去检查C++头文件和SDK生成的Go文件，理解正确的名称。

### 常见差异类型与修复方法

#### 1. 常量名差异

**问题：** 旧代码使用的常量名与SDK生成的不一致。

**修复方法：** 检查 `bindgen/SDK/headers/` 对应头文件，找到正确的C枚举名，然后查看SDK生成的Go常量名。

| 旧名称 | SDK正确名称 | 来源头文件 |
|--------|-----------|-----------|
| `DebuggerRemotePacketTypeRequest` | `DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot` | `Connection.h:166` |
| `IoctlAddActionToEvent` | `IoctlDebuggerAddActionToEvent` | `Ioctls.h` |
| `PCIE_ENDPOINT_DETAILS` | `PCI_DEV_MINIMAL` | `Pcie.h` |
| `DebuggerBreakpointModificationRequest*` | `DebuggeeBreakpointModificationRequest*` | `RequestStructures.h` |

#### 2. 结构体字段名差异

**问题：** 结构体字段名与SDK生成的不一致。

**修复方法：** 查看 `sdk/*.go` 中对应结构体的字段定义。

| 旧字段名 | SDK正确字段名 | 结构体 |
|---------|-------------|--------|
| `RequestedActionOfPacket` | `RequestedActionOfThePacket` | `DEBUGGER_REMOTE_PACKET` |

#### 3. 类型不存在

**问题：** 旧代码引用了SDK中不存在的类型。

**修复方法：**
1. 搜索 `bindgen/SDK/headers/` 找到C头文件中的对应定义
2. 搜索 `sdk/*.go` 找到SDK生成的实际类型名
3. 如果确实不存在，考虑添加 `prebuiltStruct` 或使用匿名结构体

| 旧类型 | SDK实际类型 | 说明 |
|--------|-----------|------|
| `DEBUGGER_SEARCH_LOGGING_BUFFERS` | `DEBUGGEE_RESULT_OF_SEARCH_PACKET` | 搜索结果包 |
| `SYMBOL_UPDATE_PACKET` | 匿名struct | 临时方案 |

#### 4. 字段不存在

**问题：** 旧代码访问了结构体中不存在的字段。

**修复方法：** 检查C头文件确认字段是否真的存在。如果C头文件也没有该字段，则删除引用。

| 结构体 | 不存在的字段 | 说明 |
|--------|------------|------|
| `DEBUGGER_EVENT_AND_ACTION_RESULT` | `Tag` | C头文件(Events.h:427)确认只有IsSuccessful和Error |

#### 5. bool vs uint8 类型不匹配

**问题：** `BOOLEAN` 现在映射为 `bool` 而非 `uint8`。

**修复方法：**
- `0` → `false`, `1` → `true`
- `== 0` → `!`, `!= 0` → 直接bool
- 移除 `func() uint8 { if x { return 1 }; return 0 }()` 闭包
- `uint8(x)` → 直接使用bool值

#### 6. int8 vs byte 类型不匹配

**问题：** `char` 映射为 `int8`，但 `[]byte` 与 `[]int8` 不能直接 copy。

**修复方法：**
```go
// []byte → []int8
for i := range srcBytes {
    dstInt8[i] = int8(srcBytes[i])
}

// []int8 → string
byteSlice := make([]byte, len(int8Slice))
for i, v := range int8Slice {
    byteSlice[i] = byte(v)
}
s := string(byteSlice)
```

#### 7. 枚举类型转换

**问题：** 从 `binary.LittleEndian.Uint32()` 读取的 `uint32` 需要显式转换为枚举类型。

**修复方法：**
```go
// 错误
packet.TypeOfAction = binary.LittleEndian.Uint32(payload[8:12])

// 正确
packet.TypeOfAction = sdk.DEBUGGER_MODIFY_EVENTS_TYPE(binary.LittleEndian.Uint32(payload[8:12]))
```

#### 8. 枚举值从非0开始

**问题：** 生成器默认使用 `iota`（从0开始），但C枚举值可能从1或其他值开始。

**修复方法：** 手动修正生成的枚举常量值，匹配C头文件定义。

**示例：** `DEBUGGER_REMOTE_PACKET_TYPE` 在 Connection.h:160 中从1开始：
```go
// 错误 (iota从0开始)
DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot DEBUGGER_REMOTE_PACKET_TYPE = iota

// 正确 (显式赋值)
DebuggerRemotePacketTypeDebuggerToDebuggeeExecuteOnVmxRoot DEBUGGER_REMOTE_PACKET_TYPE = 1
```

### 修复流程

1. 运行 `go test -run TestFixError -v -count=1`
2. 读取错误信息，定位文件和行号
3. 检查 `sdk/*.go` 中对应类型的实际定义
4. 检查 `bindgen/SDK/headers/*.h` 中C头文件的原始定义
5. 按照上述模式修复
6. 重新运行测试验证

### ShowErrorMessage 迁移

旧代码使用 `ShowErrorMessage(errorCode)`，SDK中没有此函数。替换为：

```go
// 旧
ShowErrorMessage(errorCode)

// 新
fmt.Println(sdk.DebuggerErrorCode(errorCode).String())
```

## Common Issues to Check

### 1. Struct Layout Mismatches

Compare struct definitions between Go and C++, ensuring:
- Field order matches exactly
- Padding is explicit in Go (`[N]byte` padding fields)
- Total struct size matches (`unsafe.Sizeof` in Go vs `sizeof` in C++)

### 2. IOCTL Code Consistency

Verify IOCTL codes match between Go and C++:

| Go Constant | C++ Define | Value |
|-------------|------------|-------|
| `IoctlDebuggerAttachDetachUserModeProcess` | `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` | 0x80e |
| `IoctlTerminateVmx` | `IOCTL_TERMINATE_VMX` | 0x802 |
| `IoctlReturnIrpPendingPacketsAndDisallowIoctl` | `IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL` | 0x801 |

**注意：** 所有IOCTL常量现在由bindgen自动从 `Ioctls.h` 生成到 `sdk/Ioctls.go`，以SDK生成的值为准。

### 3. Package Naming Convention

C++ uses hyphens in directory names, Go uses underscores:
- `kernel-level/` → `kernel_level/` (package: `kernel_level`)
- `user-level/` → `user_level/` (package: `user_level`)
- `script-engine/` → `script_engine/` (package: `scriptengine`)
- `driver-loader/` → `driver_loader/` (package: `driver_loader`)

### 4. BSOD Analysis

For Bugcheck 0xD1 (DRIVER_IRQL_NOT_LESS_OR_EQUAL):
- Check if VMX state is corrupted during unload
- Verify IOCTL sequence is correct
- Ensure Token is properly passed for all operations
- Check if event loop is stopped before VMM unload

## Completion Progress

See `COMPLETION_PROGRESS.md` for detailed function-by-function comparison.

### Priority Areas Needing Work

1. **export/** — 57→11 functions (19% complete)
2. **commands/** — 279→111 functions (40% complete) ✅ **核心命令已实现IOCTL通信**
3. **app/** — 15→6 functions (40% complete)
4. **common/** — 38→23 functions (61% complete)
5. **script-engine/** — 41→32 functions (78% complete)

## Historical Implementation Records

### Core IOCTL Communication Layer (2026-04-12)

**File:** `debugger/core/debugger.go`

已完成以下核心 IOCTL 通信方法的实现：

| 方法 | IOCTL 常量 | 功能 |
|------|-----------|------|
| `AttachToProcess()` | `IoctlDebuggerAttachDetachUserModeProcess` | 附加/分离进程 |
| `DetachFromProcess()` | `IoctlDebuggerAttachDetachUserModeProcess` | 分离进程 |
| `SetBreakpoint()` | `IoctlDebugSetOrRemoveBreakpoints` | 设置/删除断点 |
| `ContinueDebuggee()` | `IoctlContinuePacketReceived` | 继续执行 |
| `PauseDebuggee()` | `IoctlPausePacketReceived` | 暂停执行 |
| `ReadMemory()` | `IoctlReadMemory` | 读内存 |
| `WriteMemory()` | `IoctlWriteMemory` | 写内存 |
| `SearchMemory()` | `IoctlSearchMemory` | 搜索内存 |
| `BringPagesIn()` | `IoctlVmmReadPhysicalMemory` | 页面调入 |
| `PerformActionsOnApic()` | `IoctlApicActions` | APIC操作 |
| `QueryIdtEntry()` | `IoctlQueryIdtEntry` | IDT查询 |
| `PerformSmiOperation()` | `IoctlSmiOperations` | SMI操作 |

### 回调函数消除规则 (2026-04-12)

**装逼回调函数模式（必须禁止）：**

❌ **错误模式 - 回调传递：**
```go
func CommandXxxRequest(callbackFunc func(...) ..., params ...) ... {
    return callbackFunc(params...)
}
```

✅ **正确模式 - 直接IOCTL通信：**
```go
func CommandXxxRequest(dc *core.DebuggerCore, params ...) ... {
    input := make([]byte, size)
    binary.LittleEndian.PutUint32(input[0:4], param)
    success, _, err := dc.SendIoctlRequest(core.DebuggerIoctlRequest{
        Code:    core.IoctlXxx,
        Input:   input,
        Output:  output,
        InSize:  size,
        OutSize: size,
    })
    if !success || err != nil {
        return fmt.Errorf("operation failed: %v", err)
    }
    return nil
}
```

**关键原则：**
1. **所有命令函数必须接受 `*core.DebuggerCore` 作为第一个参数**
2. **绝对禁止使用回调函数作为参数**
3. **每个命令必须直接调用IOCTL或核心方法**
4. **错误处理必须返回具体的error信息，不能简单转发**

### 断点命令修复规则 (2026-04-12)

**断点修改结构体 (`DEBUGGEE_BP_LIST_OR_MODIFY_PACKET`)：**

| 字段 | 类型 | 偏移 | 大小 |
|------|------|------|------|
| `BreakpointId` | `UINT64` | 0 | 8 |
| `Request` | `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST` | 8 | 4 |
| `Result` | `UINT32` | 12 | 4 |

**请求类型常量：**

| 常量 | 值 | 描述 |
|------|-----|------|
| `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS` | 0 | 列出所有断点 |
| `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE` | 1 | 启用断点 |
| `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE` | 2 | 禁用断点 |
| `DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR` | 3 | 清除断点 |

## 任务完成状态 (2026-04-16)

### ✅ 已完成

1. **Bindgen SDK 生成器** — 基于 `modernc.org/cc/v4` 的C头文件→Go绑定自动生成器
   - 全局去重、N-ary树结构体解析、预制结构体、统一类型映射
   - 12个Go绑定文件从13个C头文件自动生成
   - 运行: `go test -run TestGenerate -v -count=1 ./bindgen/`

2. **gen → sdk 迁移** — 所有10个文件从 `bindgen/gen` 迁移到 `sdk` 包
   - 100+ 代码引用从 `gen.` 替换为 `sdk.`
   - `bindgen/gen/` 目录已删除
   - 所有类型差异已修复，`TestFixError` 通过

3. **核心模块重构** — `core/debugger.go` 使用SDK生成的结构体和枚举
   - IOCTL控制码从SDK调用
   - 所有通信发包逻辑已实现

4. **命令模块重构** — `commands/` 下的所有文件使用SDK绑定
   - 消除了回调函数模式
   - 直接调用core方法

### 🔄 进行中

1. **header 包绑定** — `libhyperdbg/header` 的C++语法过滤
   - `isCppLine()` 和 `preprocessHeader()` 已实现
   - 因C++语法复杂性暂未启用（配置已注释）

### 📋 待完成

1. **export/** — 57→11 functions (19% complete)
2. **app/** — 15→6 functions (40% complete)
3. **common/** — 38→23 functions (61% complete)
4. **script-engine/** — 41→32 functions (78% complete)
5. **枚举值修正** — 生成器对非0起始枚举默认使用iota，需手动修正或改进生成器

## Verification Steps

1. **Compare struct sizes:**
   ```go
   fmt.Printf("Struct size: %d\n", unsafe.Sizeof(SomeStruct{}))
   ```

2. **Serialize and compare bytes:**
   - Serialize Go struct to bytes
   - Compare with C++ struct serialized bytes
   - Check for padding/alignment issues

3. **Trace IOCTL calls:**
   - Log all IOCTL calls in Go
   - Compare with expected C++ behavior
   - Verify parameters match

4. **Check kernel logs:**
   - Use WinDbg to trace kernel execution
   - Compare with expected code path

5. **Run bindgen and verify:**
   ```powershell
   # 重新生成SDK绑定
   go test -run TestGenerate -v -count=1 ./bindgen/
   # 验证编译通过
   go test -run TestFixError -v -count=1
   ```

## Type Mapping Rules (2026-04-13)

### BOOLEAN Type Mapping

**C++ `BOOLEAN` (Windows type) → Go `bool`**

| C++ 类型 | Go 类型 | 说明 |
|---------|---------|------|
| `BOOLEAN` (1 byte) | `bool` (1 byte) | 直接映射，无需转换 |
| `BOOL` (4 byte, int32) | `int32` | 保持 int32 类型 |
| `_Bool` / `bool` (C99) | `bool` | Go 原生 bool |

**重要规则：**
1. **C++ 的 `BOOLEAN` 是 1 字节的 unsigned char 类型**，与 Go 的 `bool` 大小相同（都是 1 字节）
2. **在结构体中使用 `bool` 时必须添加显式填充**：如果字段后需要对齐到 4 字节边界，需要添加 `[3]byte` 填充
3. **不要使用 `uint8` 替代 `bool`**：Go 的 `bool` 在内存布局上与 C++ 的 `BOOLEAN` 完全兼容

**示例：**

```go
// C++ 结构体:
// typedef struct _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS {
//     BOOLEAN IsStartingNewProcess;      // offset 0, size 1
//     // 3 bytes padding
//     UINT32 ProcessId;                  // offset 4, size 4
// } DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS;

// Go 实现:
type DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS struct {
    IsStartingNewProcess bool    // 1 byte
    _                    [3]byte // 显式填充
    ProcessId            uint32  // 4 bytes
}
```

### Struct Padding Guidelines

1. **始终使用显式填充字段** (`[N]byte`) 而不是依赖 Go 编译器自动填充
2. **填充大小计算**：
   - 如果下一个字段是 `uint32`/`int32`，当前字段是 1 字节类型（如 `bool`），则需要填充 3 字节
   - 如果下一个字段是 `uint64`/`int64`，当前字段是 4 字节类型，则需要填充 4 字节
3. **保持字段顺序与 C++ 一致**：不要重新排列字段以优化填充

### Source Comments Standard

所有 SDK 包中的结构体和枚举都必须包含 Source 注释：

```go
// Source: HyperDbg/libhyperdbg/code/debugger/RequestStructures.h -> _STRUCT_NAME
type STRUCT_NAME struct {
    ...
}

// Source: HyperDbg/libhyperdbg/code/debugger/Events.h:160 -> _ENUM_NAME
type ENUM_NAME uint32

const (
    EnumValue1 ENUM_NAME = 1
    EnumValue2 ENUM_NAME = 2
)
```

**注释格式要求：**
- 格式：`// Source: <C++头文件路径> -> <C++结构体/枚举名>` 或 `// Source: <C++头文件路径>:行号 -> <C++结构体/枚举名>`
- 路径相对于 HyperDbg 源码根目录
- 使用下划线前缀的 C++ 名称（如 `_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS`）
- 预制结构体额外包含 `originalSource` 字段追踪实际来源

### char → int8 映射

**C++ `char` → Go `int8`**（非 `byte`/`uint8`）

C 的 `char` 是有符号的，映射为 Go 的 `int8`。当需要与 `[]byte` 交互时，需要显式转换：

```go
// int8数组 → string
byteSlice := make([]byte, len(int8Arr))
for i, v := range int8Arr {
    byteSlice[i] = byte(v)
}
s := string(byteSlice)

// []byte → int8数组
for i := range srcBytes {
    dstInt8[i] = int8(srcBytes[i])
}
```
