---
name: "hyperdbg-porter"
description: "完整复刻 HyperDbg C++ 代码到 Rust。当用户需要将 hyperhv/hyperkd C++ 代码完整移植到 Rust 时调用。"
---

# HyperDbg 代码复刻器

## 目标

完整复刻 HyperDbg C++ 代码到 Rust，确保：
1. 所有数据结构正确对应
2. 所有函数正确实现
3. 编译无错误
4. 驱动可以加载

## 源代码位置

- **C++ hyperhv**: `d:\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperhv`
- **C++ hyperkd**: `d:\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd`
- **Rust hyperhv**: `d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd\src\hyperkd\hyperhv`
- **Rust hyperkd**: `d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd\src\hyperkd`

## C++ 目录结构

### hyperhv 目录结构
```
hyperhv/
├── header/
│   ├── common/       # State.h, Common.h, Msr.h
│   ├── vmm/
│   │   ├── ept/      # Ept.h, Invept.h, Vpid.h
│   │   └── vmx/      # Vmx.h, Hv.h, Mtf.h, Vmexit.h
│   ├── hooks/        # Hooks.h, ExecTrap.h, ModeBasedExecHook.h
│   ├── memory/       # MemoryMapper.h, PoolManager.h
│   └── broadcast/    # Broadcast.h, DpcRoutines.h
├── code/
│   ├── vmm/
│   │   ├── ept/      # Ept.c, Invept.c, Vpid.c
│   │   └── vmx/      # Vmx.c, Hv.c, Mtf.c, Vmexit.c
│   ├── hooks/
│   │   ├── ept-hook/ # EptHook.c, ExecTrap.c
│   │   └── syscall-hook/ # EferHook.c
│   ├── memory/       # MemoryManager.c, MemoryMapper.c
│   └── broadcast/    # Broadcast.c, DpcRoutines.c
```

### hyperkd 目录结构
```
hyperkd/
├── header/
│   ├── debugger/
│   │   ├── core/     # Debugger.h, HaltedCore.h
│   │   ├── events/   # DebuggerEvents.h
│   │   └── commands/ # BreakpointCommands.h
│   └── driver/       # Driver.h, Loader.h
├── code/
│   ├── debugger/
│   │   ├── core/     # Debugger.c, HaltedCore.c
│   │   ├── events/   # DebuggerEvents.c
│   │   └── commands/ # BreakpointCommands.c
│   └── driver/       # Driver.c, Loader.c
```

## 复刻流程

### Phase 1: 核心头文件处理
按顺序处理以下头文件：

1. **State.h** - 核心状态定义
   - `VIRTUAL_MACHINE_STATE` -> `Vcpu`
   - `VMX_CONTEXT` -> `VmxContext`
   - `EPT_STATE` -> `EptState`
   - 所有枚举类型

2. **Ept.h** - EPT 结构定义
   - `EPT_PML4_ENTRY` -> `EptPml4Entry`
   - `EPT_PML3_ENTRY` -> `EptPdpEntry`
   - `EPT_PML2_ENTRY` -> `EptPdEntry`
   - `EPT_PML1_ENTRY` -> `EptPtEntry`
   - `VMM_EPT_PAGE_TABLE` -> `VmmEptPageTable`
   - `EPT_HOOKED_PAGE_DETAIL` -> `EptHookedPageDetail`

3. **Vmx.h** - VMX 结构定义
   - 所有 VMCS 字段常量
   - 所有 MSR 常量
   - 所有控制位常量

4. **Hooks.h** - 钩子结构定义
   - `EPT_HOOKS_CONTEXT` -> `EptHooksContext`
   - `MTF_TRAP_STATE` -> `MtfTrapState`
   - 所有钩子类型枚举

### Phase 2: 实现文件处理
按以下顺序处理实现文件：

1. **Ept.c** -> `ept.rs`
   - `EptCheckFeatures()`
   - `EptBuildMtrrMap()`
   - `EptGetMemoryType()`
   - `EptSplitLargePage()`
   - `EptGetPml1Entry()`
   - `EptGetPml2Entry()`
   - `EptSetPml1AndInvalidateTlb()`

2. **EptHook.c** -> `ept_hook.rs`
   - `EptHookPerformPageHook()`
   - `EptHookCreateHookPage()`
   - `EptHookUpdateHookPage()`
   - `EptHookFindByPhysAddress()`
   - `EptHookUnHookSingleAddress()`

3. **Mtf.c** -> `mtf.rs`
   - `MtfHandleMtfVmexit()`
   - `HvSetMonitorTrapFlag()`

4. **Vmexit.c** -> `vmexit.rs`
   - `VmxVmexitHandler()`
   - `HandleEptViolation()`
   - `HandleMtf()`
   - 所有其他 VM-exit 处理函数

### Phase 3: 模块整合
1. 更新 `mod.rs` 导出所有模块
2. 处理模块间依赖
3. 解决循环依赖

## 数据结构映射规则

### 基本类型映射
| C++ 类型 | Rust 类型 |
|---------|----------|
| `UINT8` | `u8` |
| `UINT16` | `u16` |
| `UINT32` | `u32` |
| `UINT64` | `u64` |
| `SIZE_T` | `usize` |
| `BOOLEAN` | `bool` |
| `PVOID` | `*mut c_void` |
| `CHAR[]` | `[u8; N]` |

### 结构体映射
```rust
// C++
typedef struct _EPT_PML1_ENTRY {
    UINT64 ReadAccess : 1;
    UINT64 WriteAccess : 1;
    UINT64 ExecuteAccess : 1;
    UINT64 MemoryType : 3;
    UINT64 IgnorePat : 1;
    UINT64 DirtyFlag : 1;
    UINT64 PageFrameNumber : 40;
    UINT64 Reserved : 17;
} EPT_PML1_ENTRY;

// Rust
#[repr(C, packed)]
pub struct EptPtEntry {
    pub value: u64,
}

impl EptPtEntry {
    pub fn read_access(&self) -> bool { (self.value & 1) != 0 }
    pub fn set_read_access(&mut self, v: bool) {
        if v { self.value |= 1; } else { self.value &= !1; }
    }
    // ... 其他字段
}
```

## 执行步骤

对于每个要复刻的文件：

1. **读取 C++ 头文件**
   - 提取所有结构体定义
   - 提取所有常量定义
   - 提取所有函数声明

2. **读取 C++ 实现文件**
   - 理解函数实现逻辑
   - 记录依赖关系

3. **创建/更新 Rust 文件**
   - 实现所有结构体
   - 实现所有函数
   - 添加必要的 unsafe 块

4. **验证编译**
   - 运行 `cargo check`
   - 修复所有错误

## 验证清单

每个文件完成后检查：
- [ ] 所有结构体字段完整
- [ ] 所有函数签名正确
- [ ] 所有常量值正确
- [ ] 编译无错误

## 优先级文件列表

### 最高优先级（必须完成才能编译）
1. `header/common/State.h` -> `lib.rs`
2. `header/vmm/ept/Ept.h` -> `ept.rs`
3. `code/vmm/ept/Ept.c` -> `ept.rs`
4. `header/hooks/Hooks.h` -> `hooks.rs`
5. `code/hooks/ept-hook/EptHook.c` -> `ept_hook.rs`

### 高优先级（核心功能）
6. `code/vmm/vmx/Mtf.c` -> `mtf.rs`
7. `code/vmm/vmx/Vmexit.c` -> `vmexit.rs`
8. `code/vmm/vmx/Vmx.c` -> `vmx.rs`

### 中优先级（调试器功能）
9. `header/debugger/core/Debugger.h` -> `debugger.rs`
10. `code/debugger/core/Debugger.c` -> `debugger.rs`

## 错误处理

遇到问题时：
1. 检查类型是否匹配
2. 检查生命周期是否正确
3. 检查 unsafe 块是否正确
4. 检查模块导入是否正确
5. 参考 C++ 原始实现

## 输出格式

每完成一个文件，报告：
```
✓ 完成: <C++文件> -> <Rust文件>
  - 结构体: <数量>
  - 函数: <数量>
  - 常量: <数量>
  - 编译状态: <通过/失败>
```
