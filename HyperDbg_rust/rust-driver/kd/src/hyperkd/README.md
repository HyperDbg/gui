# HyperDbg Kernel Debugger (hyperkd)

## 概述

`hyperkd` 是 HyperDbg 的内核调试器模块，负责与虚拟化层 (hyperhv) 通信并提供调试功能接口。

## 目录结构

```
hyperkd/
├── mod.rs                    # 模块入口，定义 Vcpu、VmxContext、VMX_CONTEXT
├── hyperhv/                  # 虚拟化层实现
├── allocations.rs            # 内存分配管理
├── attaching.rs              # 进程附加功能
├── callstack.rs              # 调用栈追踪
├── commands.rs               # 调试命令处理
├── debugger.rs               # 调试器核心逻辑
├── dpc_routines.rs           # DPC 例程
├── extension_commands.rs     # 扩展命令
├── ffi.rs                    # FFI 外部接口
├── halted_broadcast.rs       # 停止广播
├── network.rs                # 网络通信
├── process.rs                # 进程管理
├── thread.rs                 # 线程管理
├── ud.rs                     # 用户态调试
└── user_access.rs            # 用户态访问
```

## 核心数据结构

### Vcpu (虚拟处理器)

```
+------------------+
|     Vcpu         |
+------------------+
| core_id: u32     |  <- CPU 核心ID
| vmxon_region     |  <- VMXON 区域物理地址
| vmcs_region      |  <- VMCS 区域物理地址
| vmm_stack        |  <- VMM 栈
| msr_bitmap       |  <- MSR 位图
| io_bitmap_a/b    |  <- I/O 位图
| has_launched     |  <- 是否已启动
| state: VmxState  |  <- VMX 状态
| exit_reason      |  <- VM-Exit 原因
| guest_rip/rsp    |  <- Guest 寄存器
+------------------+
```

### VmxState 状态机

```
                    +----------------+
                    | Uninitialized  |  未初始化
                    +-------+--------+
                            |
                            v
                    +-------+--------+
                    | VmxonEnabled   |  VMXON 已启用
                    +-------+--------+
                            |
                            v
                    +-------+--------+
                    | VmcsLoaded     |  VMCS 已加载
                    +-------+--------+
                            |
                            v
                    +-------+--------+
                    | Launched       |  已启动 (VMLAUNCH)
                    +-------+--------+
                            |
                            v
                    +-------+--------+
                    | Terminated     |  已终止 (VMXOFF)
                    +----------------+
```

## Hypervisor 调用流程

```
+-------------------+     +-------------------+     +-------------------+
|   Guest Kernel    |     |    hyperkd        |     |    hyperhv        |
+-------------------+     +-------------------+     +-------------------+
         |                        |                        |
         | hypervisor_call()      |                        |
         |----------------------->|                        |
         |                        | asm_hypervisor_call    |
         |                        |----------------------->|
         |                        |                        |
         |                        |        +---------------+
         |                        |        | VMCALL        |
         |                        |        | - 保存寄存器  |
         |                        |        | - 执行操作    |
         |                        |        | - 恢复寄存器  |
         |                        |        +---------------+
         |                        |                        |
         |                        |<-----------------------|
         |<-----------------------|                        |
         |                        |                        |
```

## 汇编调用约定

```asm
; hypervisor_call 调用
; rcx = call_number
; rdx = param1
; r8  = param2
; r9  = param3
; rax = 返回值

asm_hypervisor_call:
    push r10
    push r11
    push rbx
    mov r10, 0x48564653        ; HyperDbg 签名
    mov r11, 0x564d43616c6c    ; "VMCall" 标识
    mov rbx, 0x4e4f485950455256 ; "REVERYPHON" 验证
    vmcall
    pop rbx
    pop r11
    pop r10
    ret
```

## C++ 版本对比

### C++ hyperkd 目录结构

```
hyperkd/code/
├── assembly/
│   └── AsmDebugger.asm        # 调试器汇编
├── common/
│   ├── Common.c               # 通用功能
│   └── Synchronization.c      # 同步原语
├── debugger/
│   ├── broadcast/             # 广播机制
│   ├── commands/              # 命令处理
│   ├── communication/         # 通信
│   ├── core/                  # 核心
│   ├── events/                # 事件
│   ├── kernel-level/          # 内核级调试
│   ├── memory/                # 内存管理
│   ├── meta-events/           # 元事件
│   ├── objects/               # 对象管理
│   ├── script-engine/         # 脚本引擎
│   ├── tests/                 # 测试
│   └── user-level/            # 用户态调试
└── driver/
    ├── Driver.c               # 驱动入口
    ├── Ioctl.c                # IOCTL 处理
    └── Loader.c               # 加载器
```

### 功能对比表

| 功能 | C++ 实现 | Rust 实现 | 状态 |
|------|----------|-----------|------|
| Vcpu 管理 | `VIRTUAL_MACHINE_STATE` | `Vcpu` | ✅ 已实现 |
| VMX 上下文 | `g_GuestState[]` | `VMX_CONTEXT` | ✅ 已实现 |
| Hypervisor 调用 | `AsmVmcall` | `asm_hypervisor_call` | ✅ 已实现 |
| 内存分配 | `Memory.c` | `allocations.rs` | ⚠️ 部分实现 |
| 进程附加 | `Attaching.c` | `attaching.rs` | ⚠️ 部分实现 |
| 调用栈 | `Callstack.c` | `callstack.rs` | ⚠️ 部分实现 |
| 调试命令 | `DebuggerCommands.c` | `commands.rs` | ⚠️ 部分实现 |
| DPC 例程 | `DpcRoutines.c` | `dpc_routines.rs` | ⚠️ 部分实现 |
| 串口通信 | `SerialConnection.c` | `serial_connection.rs` | ⚠️ 部分实现 |
| 脚本引擎 | `ScriptEngine.c` | - | ❌ 未实现 |
| 用户态调试 | `Ud.c` | `ud.rs` | ⚠️ 部分实现 |

## 缺失功能清单

### 高优先级

1. **脚本引擎** (`ScriptEngine.c`)
   - 脚本解析和执行
   - 表达式求值

2. **完整的事件系统**
   - 断点事件
   - 系统调用事件
   - 内存访问事件

3. **同步原语** (`Synchronization.c`)
   - 自旋锁优化
   - 等待队列

### 中优先级

4. **IOCTL 接口** (`Ioctl.c`)
   - 用户态通信
   - 命令分发

5. **广播机制** (`HaltedBroadcast.c`)
   - 多核同步
   - 停止/恢复

6. **扩展命令** (`ExtensionCommands.c`)
   - 自定义命令支持

### 低优先级

7. **内核测试** (`KernelTests.c`)
   - 单元测试框架

8. **线程持有器** (`ThreadHolder.c`)
   - 线程状态管理

## 架构图

```
+----------------------------------------------------------+
|                    User Mode Application                  |
+----------------------------------------------------------+
                            |
                            | IOCTL
                            v
+----------------------------------------------------------+
|                      hyperkd (Kernel)                     |
+----------------------------------------------------------+
|  +------------+  +------------+  +------------+           |
|  | commands   |  | debugger   |  | process    |           |
|  +------------+  +------------+  +------------+           |
|  +------------+  +------------+  +------------+           |
|  | callstack  |  | allocations|  | thread     |           |
|  +------------+  +------------+  +------------+           |
+----------------------------------------------------------+
                            |
                            | VMCALL
                            v
+----------------------------------------------------------+
|                      hyperhv (VMX Root)                   |
+----------------------------------------------------------+
|  +------------+  +------------+  +------------+           |
|  | vmexit     |  | ept        |  | hooks      |           |
|  +------------+  +------------+  +------------+           |
|  +------------+  +------------+  +------------+           |
|  | vmx        |  | vmcall     |  | callbacks  |           |
|  +------------+  +------------+  +------------+           |
+----------------------------------------------------------+
```

## 使用示例

```rust
// 初始化 VMX
unsafe {
    let mut context = VMX_CONTEXT.lock();
    context.vcpus = Some(Box::new([Vcpu::new(0)]));
    context.state = VmxState::Uninitialized;
}

// 调用 hypervisor
unsafe {
    let result = hypervisor_call(0x2, 0, 0, 0);
    if result == 0x12345678 {
        log_info!("Hypervisor test successful!");
    }
}

// 读取 Guest 内存
unsafe {
    let mut buffer = [0u8; 64];
    if read_guest_memory(0x1000, buffer.as_mut_ptr(), 64) {
        log_info!("Read successful!");
    }
}
```
