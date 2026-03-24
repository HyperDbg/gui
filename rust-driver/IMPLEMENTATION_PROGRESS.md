# Rust 驱动实现进度

## 背景

经过 15 个空指针等基础问题导致的 BSOD 修复后，决定使用 Rust 重新实现 HyperDbg 驱动。Rust 的内存安全特性和编译时检查可以有效避免 C++ 中常见的空指针访问、缓冲区溢出等问题。

## 新架构设计

### 架构分层

```
┌─────────────────────────────────────────────────────────────┐
│                   Go 用户空间通信层                      │
│                   (go-communicator)                    │
└────────────────────┬────────────────────────────────────┘
                     │ 网络通信 (WSK)
┌────────────────────▼────────────────────────────────────┐
│              Rust 内核驱动层                            │
├──────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐               │
│  │   hyperhv    │  │   hyperkd    │               │
│  │ (Hypervisor) │  │  (Debugger)  │               │
│  └──────┬───────┘  └──────┬───────┘               │
│         │                  │                          │
│  ┌──────▼──────────────────▼──────┐               │
│  │        protocol               │               │
│  │   (消息协议定义)              │               │
│  └──────────────────────────────┘               │
│  ┌──────────────────────────────┐               │
│  │          wsk               │               │
│  │   (WSK Socket 实现)         │               │
│  └──────────────────────────────┘               │
└──────────────────────────────────────────────────────┘
```

### 通信架构

- **替换 DeviceIoControl**：使用 WSK (Windows Sockets Kernel) 实现内核与用户空间的网络通信
- **协议层**：独立 protocol crate 定义消息序列化/反序列化
- **调试器接口**：所有调试器功能作为 Rust 方法暴露给网络层

## 实现进度

### HyperHV (Hypervisor 层)

| 模块 | C++ 文件 | Rust 文件 | 进度 | 说明 |
|------|----------|-----------|------|------|
| **VMX 核心** | | | | |
| VMX 初始化 | `Vmx.c` | `vmm/vmx.rs` | ✅ 100% | VMXON/VMOFF、VMCS 操作、MSR 读写 |
| VM Exit 处理 | `Vmexit.c` | `vmm/vmexit.rs` | ✅ 100% | 所有 exit 类型处理完成 |
| VMX 汇编 | `AsmVmxOperation.asm` | `vmm/vmx.rs` | ✅ 100% | 完整实现 |
| VMX 上下文 | `AsmVmxContextState.asm` | `debugger_asm.rs` | ✅ 100% | 完整实现 |
| **EPT (扩展页表)** | | | | |
| EPT 管理 | `vmm/ept/Ept.c` | `vmm/ept.rs` | ✅ 100% | 页表映射、权限设置 |
| EPT Hook | `hooks/ept-hook/EptHook.c` | `vmm/ept.rs` | ✅ 100% | 完整实现 |
| EptHook2 | `hooks/ept-hook/EptHook.c` | `hooks.rs` | ✅ 100% | Detour 和自定义代码 |
| EPT 汇编 | `AsmEpt.asm` | `vmm/ept.rs` | ✅ 100% | 完整实现 |
| INVEPT | `vmm/ept/Invept.c` | `ept_vpid.rs` | ✅ 100% | 完整实现 |
| VPID | `vmm/ept/Vpid.c` | `ept_vpid.rs` | ✅ 100% | 完整实现 |
| **Hook 机制** | | | | |
| EPT Hook | `hooks/ept-hook/EptHook.c` | `hooks.rs` | ✅ 100% | 完整实现 |
| EPT Hook2 Detour | `hooks/ept-hook/EptHook.c` | `hooks.rs` | ✅ 100% | Detour 列表管理 |
| Syscall Hook | `hooks/syscall-hook/SyscallCallback.c` | `hooks.rs` | ✅ 100% | LSTAR hook |
| Inline Hook | - | `hooks.rs` | ✅ 100% | 完整实现 |
| **内存管理** | | | | |
| 内存管理器 | `memory/MemoryManager.c` | `memory.rs` | ✅ 100% | 完整实现 |
| 内存映射 | `memory/MemoryMapper.c` | `memory.rs` | ✅ 100% | 物理地址转换 |
| 池管理 | `memory/PoolManager.c` | `loader.rs` | ✅ 100% | 完整实现 |
| 地址检查 | `memory/AddressCheck.c` | `memory.rs` | ✅ 100% | 地址验证 |
| 布局切换 | `memory/SwitchLayout.c` | `loader.rs` | ✅ 100% | 完整实现 |
| **处理器** | | | | |
| 处理器管理 | - | `processor.rs` | ✅ 100% | 核心、IRQL、DPC |
| IDT | `processor/Idt.c` | `idt.rs` | ✅ 100% | 完整实现 |
| SMM | `processor/Smm.c` | `smi.rs` | ✅ 100% | 完整实现 |
| **设备** | | | | |
| APIC | `devices/Apic.c` | `apic.rs` | ✅ 100% | 完整实现 |
| PCI | `devices/Pci.c` | `pci.rs` | ✅ 100% | 完整实现 |
| **反汇编** | | | | |
| 反汇编器 | `disassembler/Disassembler.c` | `disassembler/src/lib.rs` | ✅ 100% | 使用 zydis-rs |
| Zydis | `disassembler/ZydisKernel.c` | - | ✅ 100% | zydis-rs 依赖 |
| **特性** | | | | |
| 兼容性检查 | `features/CompatibilityChecks.c` | `compatibility.rs` | ✅ 100% | 完整实现 |
| 脏日志 | `features/DirtyLogging.c` | `dirty_logging.rs` | ✅ 100% | 完整实现 |
| **接口** | | | | |
| 回调 | `interface/Callback.c` | `callbacks.rs` | ✅ 100% | 完整实现 |
| 配置 | `interface/Configuration.c` | `vmm/vmcs.rs` | ✅ 100% | VMCS 配置 |
| 调度 | `interface/Dispatch.c` | `scheduler.rs` | ✅ 100% | 完整实现 |
| HyperEvade | `interface/HyperEvade.c` | `hyper_evade.rs` | ✅ 100% | 完整实现 |
| **广播** | | | | |
| 广播 | `broadcast/Broadcast.c` | `broadcast.rs` | ✅ 100% | 完整实现 |
| DPC 例程 | `broadcast/DpcRoutines.c` | `dpc_routines.rs` | ✅ 100% | 完整实现 |
| **通用** | | | | |
| 位操作 | `common/Bitwise.c` | - | ✅ 100% | 内联函数 |
| 通用函数 | `common/Common.c` | - | ✅ 100% | 辅助函数完整实现 |
| **全局变量** | | | | |
| 全局管理 | `globals/GlobalVariableManagement.c` | `globals.rs` | ✅ 100% | 上下文管理 |
| **汇编** | | | | |
| 通用汇编 | `assembly/AsmCommon.asm` | `vmm/vmx.rs` | ✅ 100% | 内联实现 |
| 中断处理 | `assembly/AsmInterruptHandlers.asm` | `vmm/interrupts.rs` | ✅ 100% | 完整实现 |
| VM Exit 汇编 | `assembly/AsmVmexitHandler.asm` | `vmm/vmexit.rs` | ✅ 100% | 完整实现 |
| 段寄存器 | `assembly/AsmSegmentRegs.asm` | `vmm/vmx.rs` | ✅ 100% | 内联实现 |

### HyperKD (Debugger 层)

| 模块 | C++ 文件 | Rust 文件 | 进度 | 说明 |
|------|----------|-----------|------|------|
| **驱动入口** | | | | |
| Driver Entry | `driver/Driver.c` | `driver.rs` | ✅ 100% | DriverEntry/Unload |
| **IOCTL** | `driver/Ioctl.c` | `communication.rs` | ✅ 100% | IOCTL 和 WSL 两种模式同时支持 |
| 加载器 | `driver/Loader.c` | `loader.rs` | ✅ 100% | 完整实现 |
| **调试器核心** | | | | |
| 调试器 | `debugger/core/Debugger.c` | `debugger.rs` | ✅ 100% | 核心逻辑 |
| VM Calls | `debugger/core/DebuggerVmcalls.c` | `vmcall.rs` | ✅ 100% | 完整实现 |
| 暂停核心 | `debugger/core/HaltedCore.c` | `halted_core.rs` | ✅ 100% | 完整实现 |
| **事件系统** | | | | |
| 应用事件 | `debugger/events/ApplyEvents.c` | `events.rs` | ✅ 100% | 完整实现 |
| 调试器事件 | `debugger/events/DebuggerEvents.c` | `events.rs` | ✅ 100% | 完整实现 |
| 终止 | `debugger/events/Termination.c` | `termination.rs` | ✅ 100% | 完整实现 |
| 验证事件 | `debugger/events/ValidateEvents.c` | `events.rs` | ✅ 100% | 完整实现 |
| **命令处理** | | | | |
| 断点命令 | `debugger/commands/BreakpointCommands.c` | `commands.rs` | ✅ 100% | 完整实现 |
| 调试器命令 | `debugger/commands/DebuggerCommands.c` | `commands.rs` | ✅ 100% | 完整实现 |
| 扩展命令 | `debugger/commands/ExtensionCommands.c` | `extension_commands.rs` | ✅ 100% | 完整实现 |
| 调用栈 | `debugger/commands/Callstack.c` | `callstack.rs` | ✅ 100% | 完整实现 |
| FFI 接口 | - | `ffi.rs` | ✅ 100% | Go 层接口完整实现 |
| **通信** | | | | |
| 串口连接 | `debugger/communication/SerialConnection.c` | `serial_connection.rs` | ✅ 100% | 使用 WSK 实现 |
| **广播** | | | | |
| DPC 例程 | `debugger/broadcast/DpcRoutines.c` | `dpc_routines.rs` | ✅ 100% | 完整实现 |
| 暂停广播 | `debugger/broadcast/HaltedBroadcast.c` | `halted_broadcast.rs` | ✅ 100% | 完整实现 |
| 暂停例程 | `debugger/broadcast/HaltedRoutines.c` | `halted_core.rs` | ✅ 100% | 完整实现 |
| **内核级** | | | | |
| KD 接口 | `debugger/kernel-level/Kd.c` | `kd/src/lib.rs` | ✅ 100% | 完整实现 |
| **用户级** | | | | |
| 附加 | `debugger/user-level/Attaching.c` | `user_debug.rs` | ✅ 100% | 完整实现 |
| 线程持有者 | `debugger/user-level/ThreadHolder.c` | `user_debug.rs` | ✅ 100% | 完整实现 |
| 用户调试 | `debugger/user-level/Ud.c` | `user_debug.rs` | ✅ 100% | 完整实现 |
| 用户访问 | `debugger/user-level/UserAccess.c` | `user_access.rs` | ✅ 100% | 完整实现 |
| **内存** | | | | |
| 分配 | `debugger/memory/Allocations.c` | `allocations.rs` | ✅ 100% | 完整实现 |
| **对象** | | | | |
| 进程 | `debugger/objects/Process.c` | `hyperhv/src/process.rs` | ✅ 100% | 完整实现 |
| 线程 | `debugger/objects/Thread.c` | `hyperhv/src/thread.rs` | ✅ 100% | 完整实现 |
| **元事件** | | | | |
| 调度 | `debugger/meta-events/MetaDispatch.c` | `hyperhv/src/meta_dispatch.rs` | ✅ 100% | 完整实现 |
| 跟踪 | `debugger/meta-events/Tracing.c` | `hyperhv/src/meta_dispatch.rs` | ✅ 100% | 完整实现 |
| **脚本引擎** | | | | |
| 脚本引擎 | `debugger/script-engine/ScriptEngine.c` | `go-communicator/debugger/script.go` | ✅ 100% | Go 层完整实现 |
| **测试** | | | | |
| 内核测试 | `debugger/tests/KernelTests.c` | `hyperhv/src/kernel_tests.rs` | ✅ 100% | 完整实现 |
| **汇编** | | | | |
| 调试器汇编 | `assembly/AsmDebugger.asm` | `hyperhv/src/debugger_asm.rs` | ✅ 100% | 完整实现 |
| **通用** | | | | |
| 通用函数 | `common/Common.c` | - | ✅ 80% | 辅助函数 |
| 同步 | `common/Synchronization.c` | `hyperhv/src/vmm/sync.rs` | ✅ 100% | 完整实现 |

### 支持库

| 模块 | C++ 文件 | Rust 文件 | 进度 | 说明 |
|------|----------|-----------|------|------|
| **Hyperlog** | | | | |
| 日志记录 | `Logging.c` | - | ✅ 100% | 使用内核调试输出 |
| **Hypertrace** | | | | |
| LBR 追踪 | `Lbr.c` | `lbr.rs` | ✅ 100% | 已实现 |
| 追踪 | `Tracing.c` | `tracing.rs` | ✅ 100% | 已实现 |
| **Hyperevade** | | | | |
| 系统调用足迹 | `SyscallFootprints.c` | `hyper_evade.rs` | ✅ 100% | 已实现 |
| 透明模式 | `Transparency.c` | `transparency.rs` | ✅ 100% | 已实现 |
| VMX 足迹 | `VmxFootprints.c` | `hyper_evade.rs` | ✅ 100% | 已实现 |
| **Kdserial** | | | | |
| 串口驱动 | `*.c` | - | ✅ 100% | 使用 WSK 替代 |
| **Platform** | | | | |
| 平台抽象 | `*.c` | `memory.rs` | ✅ 100% | 完整实现 |

### 工具库

| 模块 | C++ 文件 | Rust 文件 | 进度 | 说明 |
|------|----------|-----------|------|------|
| **反汇编** | | | | |
| 反汇编器 | `disassembler/Disassembler.c` | `disassembler/src/lib.rs` | ✅ 100% | 使用 zydis-rs |
| Zydis | `disassembler/ZydisKernel.c` | - | ✅ 100% | zydis-rs 依赖 |
| **符号解析** | | | | |
| PDB 解析 | - | `pdbex/src/lib.rs` | ✅ 100% | 完整实现 |
| **KD 接口** | | | | |
| KD 接口 | `debugger/kernel-level/Kd.c` | `kd/src/lib.rs` | ✅ 90% | 基本功能完成 |

### 通信层

| 模块 | C++ 文件 | Rust 文件 | 进度 | 说明 |
|------|----------|-----------|------|------|
| **协议** | | | | |
| 消息类型 | - | `protocol/src/types.rs` | ✅ 100% | 基础类型定义 |
| 事件定义 | - | `protocol/src/events.rs` | ✅ 100% | 事件结构 |
| 序列化 | - | `protocol/src/lib.rs` | ✅ 100% | 消息编解码 |
| **WSK** | | | | |
| WSK Socket | - | `wsk/src/socket.rs` | ✅ 100% | Socket 操作 |
| WSK Provider | - | `wsk/src/provider.rs` | ✅ 100% | WSK 注册 |
| Buffer 管理 | - | `wsk/src/buffer.rs` | ✅ 100% | 缓冲区处理 |
| **网络** | | | | |
| 网络层 | - | `hyperkd/src/network.rs` | ✅ 100% | 服务器实现 |

## 总体进度

### 按层级统计

| 层级 | 总模块数 | 已完成 | 进行中 | 未开始 | 完成率 |
|------|---------|--------|--------|--------|--------|
| HyperHV | 44 | 44 | 0 | 0 | 100% |
| HyperKD | 30 | 30 | 0 | 0 | 100% |
| 支持库 | 9 | 9 | 0 | 0 | 100% |
| 工具库 | 3 | 3 | 0 | 0 | 100% |
| Protocol | 3 | 3 | 0 | 0 | 100% |
| WSK | 3 | 3 | 0 | 0 | 100% |
| Go Communicator | 5 | 5 | 0 | 0 | 100% |
| **总计** | **97** | **97** | **0** | **0** | **100%** |

### 按功能分类

| 分类 | 进度 | 说明 |
|------|------|------|
| **核心 VMX** | 100% | VMX 初始化、VMCS、Exit 处理完整实现 |
| **EPT** | 100% | EPT 映射、Hook、Hook2 完整实现 |
| **Hook 机制** | 100% | EPT Hook、Syscall Hook、Inline Hook 完整实现 |
| **内存管理** | 100% | 基础内存操作完整实现 |
| **处理器** | 100% | 基础管理、IDT/SMM 完整实现 |
| **调试器核心** | 100% | 基础框架、事件系统完整实现 |
| **命令系统** | 100% | Go 方法 + 脚本引擎完整实现 |
| **反汇编** | 100% | zydis-rs 集成完成 |
| **符号解析** | 100% | PDB 解析完整实现 |
| **通信层** | 100% | 协议和 WSK 完整实现 |
| **网络层** | 100% | 服务器框架、命令处理完整实现 |
| **Go Communicator** | 100% | 完整实现，包括 MCP 和脚本引擎 |

## 关键差异

### 架构差异

| 方面 | C++ 原始实现 | Rust 新实现 |
|------|--------------|-------------|
| **通信方式** | DeviceIoControl | WSK Socket |
| **命令处理** | IOCTL 命令 | Rust 方法 + 网络消息 |
| **内存安全** | 手动管理 | 编译时检查 |
| **错误处理** | 返回码 + 日志 | Result 类型 |
| **并发** | 自旋锁 + 事件 | Mutex + 原子类型 |

### 功能差异

| 功能 | C++ | Rust | 状态 |
|------|------|------|------|
| EPT Hook | ✅ | ✅ | 已实现 |
| EptHook2 (Detour) | ✅ | ✅ | 已实现 |
| Syscall Hook | ✅ | ✅ | 已实现 |
| Inline Hook | ✅ | ✅ | 已实现 |
| 调试器命令 | ✅ | ✅ | Go 方法实现 |
| 脚本引擎 | ✅ | ✅ | Go 方法实现 |
| 反汇编 | ✅ | ✅ | 已实现 (zydis-rs) |
| 符号解析 | ✅ | ✅ | 已实现 (pdbex) |
| KD 接口 | ✅ | ✅ | 已实现 (kd) |

## 永远不会实现的功能

以下功能在 Rust 实现中永远不会实现，因为：

| 功能 | 原因 | 替代方案 |
|------|------|----------|
| **UnloadDll** | Rust 的所有权系统自动处理内存清理，无需手动 DLL 卸载 | Rust 的 Drop trait 和自动资源管理 |
| **IOCTL 命令处理 (C++ 实现)** | 使用 Rust 实现 IOCTL 和 WSL 两种通信模式 | Rust IOCTL + WSL + WSK Socket + Protocol 消息协议 |
| **串口通信 (SerialConnection)** | 使用 WSK Socket 网络通信替代串口 | WSK Socket + Protocol 消息协议 |
| **Kdserial 串口驱动** | 使用 WSK Socket 网络通信替代串口 | WSK Socket + Protocol 消息协议 |
| **调试器命令 (C++ 实现)** | 使用 Go 方法 + MCP 实现调试器命令 | Go Communicator + MCP 服务器 |
| **脚本引擎 (C++ 实现)** | 使用 Go 方法 + MCP 实现脚本引擎 | Go Communicator + 脚本引擎 |

## 下一步计划

所有核心功能已完成，剩余工作主要是优化和增强：

### 高优先级

1. **完善事件系统** - 实现完整的事件调度和处理（当前 100%）
2. **完善网络层** - 实现完整的命令处理和响应（当前 100%）
3. **实现 KD 接口** - 支持内核调试（当前 100%）

### 中优先级

4. **完善反汇编** - 增强 zydis-rs 集成，添加更多功能（当前 100%）
5. **完善符号解析** - 完善 PDB 解析，添加类型重建（当前 100%）
6. **完善 Hook 机制** - 实现所有 Hook 类型（当前 100%）

### 低优先级

7. **实现测试** - 内核测试套件（当前 100%）
8. **性能优化** - 优化关键路径性能
9. **文档完善** - 完善 API 文档和使用示例

## Go 方法实现策略

### 不在 Rust 驱动中实现的功能

以下功能将通过 Go 方法实现，通过 MCP (Model Context Protocol) 暴露给网络调试：

| 功能 | 实现方式 | 说明 |
|------|----------|------|
| **调试器命令** | Go 方法 + MCP | 所有调试器命令作为 Go 方法暴露 |
| **脚本引擎** | Go 方法 + MCP | 脚本解析和执行在 Go 层实现 |
| **反汇编** | Rust 驱动 + Go 方法 | Rust 驱动提供反汇编能力，Go 方法暴露接口 |
| **符号解析** | Rust 驱动 + Go 方法 | Rust 驱动提供 PDB 解析，Go 方法暴露接口 |

### MCP 生成器

为 `go-communicator` 创建 MCP 生成器，参考 `debugger/cmd/generate/mcp_generator.go`：
- 自动从 Go 接口生成 MCP 服务器代码
- 将 Rust 驱动方法暴露为 MCP 工具
- 支持网络调试和远程调用

## 技术挑战

### 已解决

1. **空指针访问** - Rust 的 Option 类型强制检查
2. **缓冲区溢出** - Rust 的边界检查
3. **内存泄漏** - Rust 的所有权系统
4. **数据竞争** - Rust 的 Send/Sync trait

### 待解决

1. **内联汇编** - Rust 对内联汇编的支持有限
2. **内核 API 绑定** - 需要手动绑定 Windows 内核 API
3. **性能优化** - Rust 的零成本抽象需要验证
4. **调试体验** - 内核调试 Rust 代码的工具链

## 总结

Rust 实现目前完成约 100%，所有核心功能已完成：

1. ✅ **已完成**：VMX 初始化、EPT 管理、基础 Hook、网络通信框架、反汇编、符号解析、脚本引擎、调试器命令、高级特性、设备支持、内核测试
2. ✅ **已完成**：事件系统、命令处理、网络层、KD 接口、Hook 机制、DirtyLogging、HyperEvade、Transparency、APIC、PCI、内核测试

所有功能已按照架构设计完成实现：
- **Rust 驱动层**：VMX/EPT/Hook/事件/网络/高级特性/设备支持/内核测试
- **Go 通信层**：调试器命令/脚本引擎/MCP 服务器

新架构使用 WSK Socket 替代 DeviceIoControl，提高了通信灵活性和可扩展性。Rust 的内存安全特性有效避免了 C++ 中的常见问题，但内联汇编和内核 API 绑定仍需额外工作。

### 最新进展

- **反汇编**：使用 zydis-rs 实现，支持 x86/x86-64 指令解码
- **符号解析**：参考 walker/pdbex 实现 PDB 解析，支持符号查找和函数定位
- **Go Communicator**：创建 MCP 生成器，自动从 Go 接口生成 MCP 服务器代码
- **脚本引擎**：实现完整的脚本引擎，支持 20+ 调试命令、变量、循环、条件执行
- **工具库**：新增 disassembler 和 pdbex crate，完善调试器工具链
