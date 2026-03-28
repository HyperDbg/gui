# HyperDbg Hypervisor (hyperhv)

## 概述

`hyperhv` 是 HyperDbg 的虚拟化管理程序，负责 VMX 操作、VM-Exit 处理、EPT 管理和 Hook 机制。

## 目录结构

```
hyperhv/
├── mod.rs                    # 模块入口
├── state.rs                  # 状态定义
├── assembly/                 # 汇编代码
│   ├── mod.rs
│   ├── asm.rs                # 通用汇编函数
│   └── debugger_asm.rs       # 调试器汇编 (寄存器保存/恢复)
├── broadcast/                # 广播机制
│   ├── mod.rs
│   ├── broadcast.rs          # 多核广播
│   ├── dpc_routines.rs       # DPC 例程
│   └── halted_core.rs        # 停止核心
├── common/                   # 通用功能
│   ├── mod.rs
│   ├── bitwise.rs            # 位操作
│   └── msr.rs                # MSR 读写
├── devices/                  # 设备模拟
│   ├── mod.rs
│   ├── apic.rs               # APIC
│   └── pci.rs                # PCI
├── disassembler/             # 反汇编
│   └── mod.rs
├── features/                 # 特性
│   ├── mod.rs
│   ├── compatibility.rs      # 兼容性检查
│   ├── dirty_logging.rs      # 脏页记录
│   ├── lbr.rs                # 最后分支记录
│   ├── tracing.rs            # 追踪
│   └── transparency.rs       # 透明性
├── globals/                  # 全局状态
│   ├── mod.rs
│   └── globals.rs            # 全局变量管理
├── hooks/                    # Hook 机制
│   ├── mod.rs
│   ├── hooks.rs              # Hook 管理
│   ├── ept_hook/             # EPT Hook
│   └── syscall_hook/         # 系统调用 Hook
├── interface/                # 接口
│   ├── mod.rs
│   ├── callbacks.rs          # 回调函数
│   ├── communication.rs      # 通信
│   ├── loader.rs             # 加载器
│   └── termination.rs        # 终止
├── memory/                   # 内存管理
│   ├── mod.rs
│   ├── memory.rs             # 内存分配
│   └── switch_layout.rs      # 地址空间切换
├── mmio/                     # MMIO
│   └── mod.rs
├── processor/                # 处理器
│   ├── mod.rs
│   ├── processor.rs          # 处理器管理
│   ├── idt.rs                # IDT
│   └── interrupts.rs         # 中断处理
└── vmm/                      # 虚拟机管理
    ├── mod.rs
    ├── ept/                  # EPT
    │   ├── mod.rs
    │   ├── ept.rs            # EPT 实现
    │   └── ept_vpid.rs       # VPID
    └── vmx/                  # VMX
        ├── mod.rs
        ├── vmx.rs            # VMX 操作
        ├── vmcs.rs           # VMCS 管理
        ├── vmexit.rs         # VM-Exit 处理
        ├── vmcall.rs         # VMCALL 处理
        ├── events.rs         # 事件
        ├── mtf.rs            # MTF
        └── sync.rs           # 同步
```

## VM-Exit 处理流程

```
+-------------------+
|   Guest Mode      |
|   (VMX Non-Root)  |
+-------------------+
         |
         | VM-Exit (CPUID, VMCALL, EPT Violation, etc.)
         v
+-------------------+
| AsmVmexitHandler  |  <- Naked 函数，保存所有寄存器
+-------------------+
         |
         | push r15..rax, cr2
         | lea rcx, [rsp + 0x10]  ; Guest 寄存器指针
         | lea rdx, [rsp + 0x88]  ; Guest RSP 指针
         | lea r8, [rsp + 0x90]   ; Guest RFLAGS 指针
         | call VmxVmexitHandlerEntry
         v
+-------------------+
|VmxVmexitHandlerEntry|
+-------------------+
         |
         | 读取 exit_reason, exit_qualification
         | 调用 handle_vmexit()
         v
+-------------------+
|   handle_vmexit   |  <- 根据退出原因分发
+-------------------+
         |
         +-- 0  --> handle_exception_nmi()
         +-- 1  --> handle_external_interrupt()
         +-- 10 --> handle_cpuid()
         +-- 18 --> handle_vmcall()
         +-- 28 --> handle_cr_access()
         +-- 31 --> handle_rdmsr()
         +-- 32 --> handle_wrmsr()
         +-- 48 --> handle_ept_violation()
         +-- ...
         |
         v
+-------------------+
|  返回值处理        |
+-------------------+
         |
         | al == true  --> vmresume (继续执行)
         | al == false --> ret (退出 VMX)
         v
+-------------------+
|   Guest Mode      |
+-------------------+
```

## VM-Exit 原因处理表

| 退出原因 | 名称 | Rust 实现 | C++ 实现 | 状态 |
|----------|------|-----------|----------|------|
| 0 | EXCEPTION_NMI | `handle_exception_nmi` | `DispatchEventException` | ✅ |
| 1 | EXTERNAL_INTERRUPT | `handle_external_interrupt` | `DispatchEventExternalInterrupts` | ✅ |
| 3 | INIT_SIGNAL | `handle_init_signal` | - | ✅ |
| 4 | STARTUP_IPI | `handle_startup_ipi` | - | ✅ |
| 7 | INTERRUPT_WINDOW | `handle_interrupt_window` | `IdtEmulationHandleInterruptWindowExiting` | ✅ |
| 8 | NMI_WINDOW | `handle_nmi_window` | `IdtEmulationHandleNmiWindowExiting` | ✅ |
| 10 | CPUID | `handle_cpuid` | `DispatchEventCpuid` | ✅ |
| 12 | HLT | `handle_hlt` | - | ✅ |
| 18 | VMCALL | `handle_vmcall` | `DispatchEventVmcall` | ✅ |
| 28 | MOV_CR | `handle_cr_access` | `DispatchEventMovToFromControlRegisters` | ✅ |
| 29 | MOV_DR | `handle_dr_access` | `DispatchEventMovToFromDebugRegisters` | ✅ |
| 30 | IO_INSTRUCTION | `handle_io_instruction` | `DispatchEventIO` | ✅ |
| 31 | RDMSR | `handle_rdmsr` | `DispatchEventRdmsr` | ✅ |
| 32 | WRMSR | `handle_wrmsr` | `DispatchEventWrmsr` | ✅ |
| 37 | MTF | `handle_mtf` | `MtfHandleVmexit` | ✅ |
| 48 | EPT_VIOLATION | `handle_ept_violation` | `EptHandleEptViolation` | ✅ |
| 49 | EPT_MISCONFIG | `handle_ept_misconfig` | `EptHandleMisconfiguration` | ✅ |
| 52 | INVPCID | `handle_invpcid` | - | ✅ |

## VMCALL 处理

```
+-------------------+
|   handle_vmcall   |
+-------------------+
         |
         | 读取 rax = vmcall_number
         v
+-------------------+
|  VmcallNumber 枚举 |
+-------------------+
         |
         +-- 0x01 --> VMXOFF (关闭虚拟化)
         +-- 0x02 --> Test (测试调用)
         +-- 0x03 --> ReadGuestMemory
         +-- 0x04 --> SetBreakpoint
         +-- 0x05 --> ClearBreakpoint
         +-- 0x06 --> EnableDebugExits
         +-- 0x0C --> ChangeMsrBitmapRead
         +-- 0x0D --> ChangeMsrBitmapWrite
         +-- 0x12 --> EnableRdtscExiting
         +-- 0x1A --> EnableExternalInterruptExiting
         +-- 0x1C --> EnableSyscallHookEfer
         +-- 0x1E --> EnableMovToCr3Exiting
         +-- ...
```

## EPT (扩展页表) 架构

```
+-------------------+
|    EPT PML4       |  <- 4 级页表
+-------------------+
         |
         | 512 entries
         v
+-------------------+
|    EPT PDPT       |  <- 页目录指针表
+-------------------+
         |
         | 512 entries
         v
+-------------------+
|    EPT PD         |  <- 页目录 (2MB 大页)
+-------------------+
         |
         | 512 entries
         v
+-------------------+
|    EPT PT         |  <- 页表 (4KB 页)
+-------------------+

EPT 权限位:
+------+------+--------+----------+
| Read | Write| Execute | User-Exec|
+------+------+--------+----------+
| bit 0| bit 1| bit 2   | bit 10   |
+------+------+--------+----------+

内存类型:
+------------------+-------+
| 类型             | 值    |
+------------------+-------+
| UC (Uncacheable) | 0     |
| WC (Write-Comb.) | 1     |
| WT (Write-Thru)  | 4     |
| WP (Write-Prot.) | 5     |
| WB (Write-Back)  | 6     |
+------------------+-------+
```

## Hook 机制

### EPT Hook 流程

```
+-------------------+
|  ept_hook_add     |
+-------------------+
         |
         | 1. 分配影子页面
         | 2. 复制原始页面内容
         | 3. 修改影子页面 (写入 Hook)
         | 4. 设置 EPT 权限
         v
+-------------------+
|  EPT 配置         |
+-------------------+
         |
         | 原始页面: R-X (读执行)
         | 影子页面: R-W (读写)
         v
+-------------------+
|  执行流程         |
+-------------------+
         |
         | 执行: 原始页面 (R-X)
         | 写入: 影子页面 (R-W)
         | 读取: 原始页面 (R-X)
         v
+-------------------+
|  EPT Violation    |
+-------------------+
         |
         | 检测写入操作
         | 重定向到影子页面
         v
```

### 系统调用 Hook

```
+-------------------+
|  Syscall Hook     |
+-------------------+
         |
         | 保存原始 IA32_LSTAR
         | 写入新的处理函数地址
         v
+-------------------+
|  Hook Handler     |
+-------------------+
         |
         | 1. 保存寄存器
         | 2. 检查系统调用号
         | 3. 执行自定义逻辑
         | 4. 调用原始处理函数
         | 5. 恢复寄存器
         v
```

## C++ 版本对比

### C++ hyperhv 目录结构

```
hyperhv/code/
├── assembly/
│   ├── AsmCommon.asm           # 通用汇编
│   ├── AsmEpt.asm              # EPT 汇编
│   ├── AsmHooks.asm            # Hook 汇编
│   ├── AsmInterruptHandlers.asm # 中断处理
│   ├── AsmSegmentRegs.asm      # 段寄存器
│   ├── AsmVmexitHandler.asm    # VM-Exit 处理
│   ├── AsmVmxContextState.asm  # VMX 上下文
│   └── AsmVmxOperation.asm     # VMX 操作
├── broadcast/
│   ├── Broadcast.c             # 广播
│   └── DpcRoutines.c           # DPC
├── common/
│   ├── Bitwise.c               # 位操作
│   ├── Common.c                # 通用
│   └── UnloadDll.c             # 卸载
├── components/
│   └── registers/
│       └── DebugRegisters.c    # 调试寄存器
├── devices/
│   ├── Apic.c                  # APIC
│   └── Pci.c                   # PCI
├── disassembler/
│   ├── Disassembler.c          # 反汇编
│   └── ZydisKernel.c           # Zydis
├── features/
│   ├── CompatibilityChecks.c   # 兼容性检查
│   └── DirtyLogging.c          # 脏页记录
├── globals/
│   └── GlobalVariableManagement.c # 全局变量
├── hooks/
│   ├── ept-hook/
│   │   ├── EptHook.c           # EPT Hook
│   │   ├── ExecTrap.c          # 执行陷阱
│   │   └── ModeBasedExecHook.c # 模式执行 Hook
│   └── syscall-hook/
│       ├── EferHook.c          # EFER Hook
│       └── SyscallCallback.c   # 系统调用回调
├── interface/
│   ├── Callback.c              # 回调
│   ├── Configuration.c         # 配置
│   ├── DirectVmcall.c          # 直接 VMCALL
│   ├── Dispatch.c              # 分发
│   ├── Export.c                # 导出
│   └── HyperEvade.c            # 逃避检测
├── memory/
│   ├── AddressCheck.c          # 地址检查
│   ├── Conversion.c            # 地址转换
│   ├── Layout.c                # 内存布局
│   ├── MemoryManager.c         # 内存管理
│   ├── MemoryMapper.c          # 内存映射
│   ├── PoolManager.c           # 内存池
│   ├── Segmentation.c          # 分段
│   └── SwitchLayout.c          # 切换布局
├── mmio/
│   └── MmioShadowing.c         # MMIO 影子
├── processor/
│   ├── Idt.c                   # IDT
│   └── Smm.c                   # SMM
└── vmm/
    ├── ept/
    │   ├── Ept.c               # EPT
    │   ├── Invept.c            # INVEPT
    │   └── Vpid.c              # VPID
    └── vmx/
        ├── Counters.c          # 计数器
        ├── CrossVmexits.c      # 跨 VM-Exit
        ├── Events.c            # 事件
        ├── Hv.c                # Hypervisor
        ├── IdtEmulation.c      # IDT 模拟
        ├── IoHandler.c         # I/O 处理
        ├── ManageRegs.c        # 寄存器管理
        ├── MsrHandlers.c       # MSR 处理
        ├── Mtf.c               # MTF
        ├── ProtectedHv.c       # 保护
        ├── Vmcall.c            # VMCALL
        ├── Vmexit.c            # VM-Exit
        ├── Vmx.c               # VMX
        ├── VmxBroadcast.c      # 广播
        ├── VmxMechanisms.c     # 机制
        └── VmxRegions.c        # 区域
```

### 功能对比表

| 模块 | C++ 文件 | Rust 文件 | 状态 |
|------|----------|-----------|------|
| **汇编** |
| VM-Exit Handler | `AsmVmexitHandler.asm` | `vmexit.rs` (naked_asm!) | ✅ |
| VMX 操作 | `AsmVmxOperation.asm` | `vmx.rs` (asm!) | ✅ |
| EPT | `AsmEpt.asm` | `ept.rs` | ✅ |
| **VMX** |
| VMX 初始化 | `Vmx.c` | `vmx.rs` | ✅ |
| VMCS 管理 | `VmxRegions.c` | `vmcs.rs` | ✅ |
| VM-Exit | `Vmexit.c` | `vmexit.rs` | ✅ |
| VMCALL | `Vmcall.c` | `vmcall.rs` | ✅ |
| MSR 处理 | `MsrHandlers.c` | `vmexit.rs` (handle_rdmsr/wrmsr) | ✅ |
| I/O 处理 | `IoHandler.c` | `vmexit.rs` (handle_io_instruction) | ✅ |
| IDT 模拟 | `IdtEmulation.c` | `idt.rs` | ⚠️ 部分 |
| **EPT** |
| EPT 管理 | `Ept.c` | `ept.rs` | ✅ |
| INVEPT | `Invept.c` | `ept.rs` | ✅ |
| VPID | `Vpid.c` | `ept_vpid.rs` | ✅ |
| **Hook** |
| EPT Hook | `EptHook.c` | `ept_hook.rs` | ✅ |
| 执行陷阱 | `ExecTrap.c` | - | ❌ 未实现 |
| 模式执行 Hook | `ModeBasedExecHook.c` | - | ❌ 未实现 |
| EFER Hook | `EferHook.c` | `syscall_hook/mod.rs` | ⚠️ 部分 |
| **内存** |
| 内存管理 | `MemoryManager.c` | `memory.rs` | ⚠️ 部分 |
| 内存映射 | `MemoryMapper.c` | - | ❌ 未实现 |
| 内存池 | `PoolManager.c` | - | ❌ 未实现 |
| 地址转换 | `Conversion.c` | - | ❌ 未实现 |
| **其他** |
| 广播 | `Broadcast.c` | `broadcast.rs` | ⚠️ 部分 |
| 调试寄存器 | `DebugRegisters.c` | - | ❌ 未实现 |
| SMM | `Smm.c` | - | ❌ 未实现 |
| MMIO 影子 | `MmioShadowing.c` | - | ❌ 未实现 |
| 逃避检测 | `HyperEvade.c` | `hyper_evade.rs` | ⚠️ 部分 |

## 缺失功能清单

### 高优先级

1. **内存映射器** (`MemoryMapper.c`)
   - GPA 到 HPA 转换
   - 内存区域管理

2. **内存池管理** (`PoolManager.c`)
   - 高效内存分配
   - 内存池缓存

3. **执行陷阱** (`ExecTrap.c`)
   - 隐藏断点执行
   - 代码执行追踪

4. **模式执行 Hook** (`ModeBasedExecHook.c`)
   - 用户/内核模式区分执行

### 中优先级

5. **调试寄存器管理** (`DebugRegisters.c`)
   - DR0-DR3 管理
   - 硬件断点支持

6. **IDT 模拟** (`IdtEmulation.c`)
   - 中断注入
   - 异常处理

7. **地址转换** (`Conversion.c`)
   - 虚拟地址转换
   - 物理地址映射

### 低优先级

8. **SMM 支持** (`Smm.c`)
   - SMI 处理
   - SMM 状态管理

9. **MMIO 影子** (`MmioShadowing.c`)
   - MMIO 拦截
   - 设备模拟

10. **跨 VM-Exit** (`CrossVmexits.c`)
    - 跨核心同步
    - 状态共享

## 汇编实现对比

### C++ 版本 (AsmVmexitHandler.asm)

```asm
AsmVmexitHandler PROC
    push 0                      ; 对齐
    pushfq                      ; 保存 RFLAGS
    
    sub rsp, 0110h              ; XMM 寄存器空间
    
    movaps xmmword ptr [rsp+000h], xmm0
    movaps xmmword ptr [rsp+010h], xmm1
    ; ... 保存 xmm2-xmm5
    
    stmxcsr dword ptr [rsp+0100h]
    
    push r15
    push r14
    ; ... 保存 r13-rax
    
    mov rcx, rsp                ; 参数: PGUEST_REGS
    sub rsp, 020h               ; 影子空间
    call VmxVmexitHandler
    add rsp, 020h
    
    cmp al, 1                   ; 检查是否 VMXOFF
    je AsmVmxoffHandler
    
    ; 恢复寄存器
    pop rax
    ; ... 恢复所有寄存器
    
    jmp VmxVmresume
    
AsmVmxoffHandler:
    ; 恢复寄存器并退出
    add rsp, 0118h
    ; ... 获取返回地址
    ret
AsmVmexitHandler ENDP
```

### Rust 版本 (vmexit.rs)

```rust
#[no_mangle]
#[unsafe(naked)]
pub unsafe extern "C" fn AsmVmexitHandler() -> ! {
    naked_asm!(
        "push r15",
        "push r14",
        // ... 保存所有寄存器
        
        "mov rax, cr2",
        "push rax",
        
        "lea rcx, [rsp + 0x10]",   // Guest 寄存器
        "lea rdx, [rsp + 0x88]",   // Guest RSP
        "lea r8, [rsp + 0x90]",    // Guest RFLAGS
        
        "sub rsp, 0x28",
        "call {vmexit_handler}",
        "add rsp, 0x28",
        
        // 恢复寄存器
        
        "test al, al",
        "jz 2f",                   // VMXOFF
        
        // vmresume 路径
        "vmresume",
        "jmp 3f",
        
        "2:",                      // VMXOFF 路径
        "add rsp, 16",
        "xor rax, rax",
        "ret",
        
        "3:",
        "int3",
        vmexit_handler = sym VmxVmexitHandlerEntry,
    );
}
```

## 关键差异

1. **寄存器保存方式**
   - C++ 使用 `.asm` 文件，完全手动管理
   - Rust 使用 `naked_asm!` 宏，更安全但需要更多注意

2. **XMM 寄存器**
   - C++ 保存 xmm0-xmm5
   - Rust 版本当前未保存 XMM (需要添加)

3. **错误处理**
   - C++ 使用 NTSTATUS 返回值
   - Rust 使用 `Result<T, VmxError>` 枚举

4. **内存管理**
   - C++ 使用 Windows 内核池
   - Rust 使用 `alloc` crate 和自定义分配器

## 使用示例

```rust
// 初始化 VMX
unsafe {
    if !check_vmx_support() {
        return Err(VmxError::VmxNotSupported);
    }
    
    enable_vmx_operation();
    vmxon(vmxon_region_pa)?;
    
    vmclear(vmcs_region_pa)?;
    vmptrld(vmcs_region_pa)?;
    
    setup_vmcs(vcpu, host_rip, host_rsp)?;
    
    vmlaunch()?;
}

// 处理 VM-Exit
unsafe fn handle_vmexit(...) -> bool {
    match exit_reason {
        10 => handle_cpuid(guest_rsp, guest_rip),
        18 => handle_vmcall(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        // ...
        _ => true,
    }
}
```
