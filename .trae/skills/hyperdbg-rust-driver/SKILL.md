---
name: "hyperdbg-rust-driver"
description: "HyperDbg Rust 驱动开发完整指南。包含 C++ 代码复刻、Rust 驱动开发、类型同步、编译验证等。在需要开发或维护 HyperDbg_rust 驱动时调用。"
---

# HyperDbg Rust 驱动开发完整指南

## ⚠️ 重要警告

**严禁直接使用 `cargo build` 命令！**

- 本项目使用 **EWDK (Enterprise Windows Driver Kit)** 进行编译
- 只能使用 PowerShell 构建脚本进行编译
- 当用户要求执行某个 `.ps1` 脚本时，**必须直接执行该脚本**

## ⚠️ 类型同步警告

**修改数据结构时必须更新生成器！**

| 修改类型 | 需要运行的命令 |
|----------|---------------|
| 修改 `types.go` | `cd cmd/rustgen && go run main.go` |
| 修改 `interfaces.go` | `cd cmd/rustgen && go run main.go` |
| 修改 `packet.go` | `cd cmd/rustgen && go run main.go` |

## 源代码位置

| 组件 | C++ 源码 | Rust 实现 |
|------|---------|----------|
| hyperhv | `doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperhv` | `HyperDbg_rust/rust-driver/kd/src/hyperkd/hyperhv` |
| hyperkd | `doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/hyperkd` | `HyperDbg_rust/rust-driver/kd/src/hyperkd` |

## 文件结构

```
rust-driver/
├── driver-framework/       # 驱动框架库 (独立 crate)
├── logger/                 # 日志模块 (独立 crate)
├── common/                 # 通用模块（自动生成）
│   ├── types_gen/          # 生成的类型
│   └── handlers_gen/       # 生成的处理器
├── net/                    # 网络模块 (独立 crate)
└── kd/                     # 主驱动 crate
    └── src/
        ├── hyperkd/        # 调试器核心
        │   └── hyperhv/    # Hypervisor 实现
        ├── disassembler/   # 反汇编器
        └── logger/         # 内部日志模块
```

---

# 第一部分：C++ 代码复刻

## 复刻流程

### Phase 1: 核心头文件处理

1. **State.h** - 核心状态定义
   - `VIRTUAL_MACHINE_STATE` -> `Vcpu`
   - `VMX_CONTEXT` -> `VmxContext`
   - `EPT_STATE` -> `EptState`

2. **Ept.h** - EPT 结构定义
   - `EPT_PML4_ENTRY` -> `EptPml4Entry`
   - `EPT_PML3_ENTRY` -> `EptPdpEntry`
   - `EPT_PML2_ENTRY` -> `EptPdEntry`
   - `EPT_PML1_ENTRY` -> `EptPtEntry`

3. **Vmx.h** - VMX 结构定义
   - 所有 VMCS 字段常量
   - 所有 MSR 常量
   - 所有控制位常量

### Phase 2: 实现文件处理

1. **Ept.c** -> `ept.rs`
2. **EptHook.c** -> `ept_hook.rs`
3. **Mtf.c** -> `mtf.rs`
4. **Vmexit.c** -> `vmexit.rs`

## 数据结构映射规则

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

### 结构体映射示例

```rust
// C++
typedef struct _EPT_PML1_ENTRY {
    UINT64 ReadAccess : 1;
    UINT64 WriteAccess : 1;
    UINT64 ExecuteAccess : 1;
    UINT64 MemoryType : 3;
    UINT64 PageFrameNumber : 40;
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
}
```

---

# 第二部分：WDK 绑定

## 🔴 强制阅读 WDK-SYS 绑定文件

**写任何模块之前，必须先阅读以下文件：**

```
rust-driver/out/ntddk.rs      # 所有内核函数
rust-driver/out/types.rs      # 所有类型定义
rust-driver/out/constants.rs  # 所有常量
```

## WDK-SYS 关键函数签名

### IRP 相关函数

```rust
pub fn IoAllocateIrp(StackSize: CCHAR, ChargeQuota: BOOLEAN) -> PIRP;
pub fn IoFreeIrp(Irp: PIRP);
pub fn IoAllocateMdl(VirtualAddress: PVOID, Length: ULONG, ...) -> PMDL;
pub fn IoFreeMdl(Mdl: PMDL);
```

### 事件相关函数

```rust
pub fn KeInitializeEvent(Event: PRKEVENT, Type: EVENT_TYPE, State: BOOLEAN);
pub fn KeSetEvent(Event: PRKEVENT, Increment: KPRIORITY, Wait: BOOLEAN) -> LONG;
pub fn KeWaitForSingleObject(Object: PVOID, WaitReason: KWAIT_REASON, ...) -> NTSTATUS;
```

## 枚举类型使用

**枚举类型是 `i32` 类型别名，必须通过模块路径访问常量！**

```rust
// 错误用法 ❌
use wdk_sys::EVENT_TYPE;
EVENT_TYPE::SynchronizationEvent  // 编译错误！

// 正确用法 ✅
use wdk_sys::_EVENT_TYPE;
_EVENT_TYPE::SynchronizationEvent  // 正确！
```

---

# 第三部分：Rust 生成器同步

当 Go 端的类型定义或接口发生变化时，需要同步更新 Rust 生成器生成的代码。

## 同步步骤

1. **运行 Rust 生成器**
   ```bash
   cd HyperDbg_rust/cmd/rustgen
   go run main.go
   ```

2. **检查 git diff**
   ```bash
   cd HyperDbg_rust/rust-driver/kd
   git diff --stat
   git diff
   ```

3. **编译 Rust 驱动验证**
   ```bash
   cd HyperDbg_rust/rust-driver/kd
   cargo build
   ```

4. **修复生成器（如有警告或错误）**
   - 如果出现 "unused import" 警告，在生成器中添加 `#![allow(unused_imports)]`
   - 如果出现类型找不到错误，检查是否需要 `use super::*;`
   - 生成器文件：`HyperDbg_rust/cmd/rustgen/main.go`

## 生成器关键代码位置

| 函数 | 作用 |
|------|------|
| `generateTypesSplit()` | 类型生成 |
| `generateTypesMod()` | 模块生成 |
| `generateResponseTypes()` | 响应类型 |
| `generateHandlersSplit()` | 处理器生成 |

## 常见问题

| 问题 | 解决方案 |
|------|---------|
| 未使用的导入警告 | 添加 `#![allow(unused_imports)]` 属性 |
| 类型找不到错误 | 确保 `use super::*;` 存在 |
| mod.rs 中的警告 | 移除 mod.rs 中不需要的导入 |

---

# 第四部分：WSK 网络编程

## 参考项目

| 项目 | 路径 | 说明 |
|------|------|------|
| KSOCKET | `todo/参考价值/KSOCKET-master/` | Berkeley Socket API 风格的 WSK 封装 |
| afdmjhk | `todo/参考价值/afdmjhk-master/WSK/` | WSK 示例代码 |
| DriverUtil | `todo/参考价值/代码+课件/DriverUtil/` | EPT Hook 和 WSK Socket 实现 |

## WSK 核心实现模式

```rust
// 异步上下文结构
struct Context {
    event: KEVENT,  // 使用 KEVENT 类型 (24字节)
    irp: PIRP,
}

// 完成回调 - 必须返回 STATUS_MORE_PROCESSING_REQUIRED
unsafe extern "C" fn wsk_completion_routine(
    _device: PDEVICE_OBJECT,
    _irp: PIRP,
    context: PVOID,
) -> NTSTATUS {
    if !context.is_null() {
        KeSetEvent(context as PRKEVENT, 0, false as BOOLEAN);
    }
    STATUS_MORE_PROCESSING_REQUIRED  // 关键！
}

// 初始化异步上下文
impl Context {
    fn new() -> Result<Self, SocketError> {
        unsafe {
            let irp = IoAllocateIrp(1, false as BOOLEAN);
            let mut ctx = Context {
                event: core::mem::zeroed(),
                irp,
            };
            KeInitializeEvent(
                &raw mut ctx.event as PRKEVENT, 
                _EVENT_TYPE::SynchronizationEvent, 
                false as BOOLEAN
            );
            Ok(ctx)
        }
    }

    fn wait_for_completion(&mut self) -> NTSTATUS {
        unsafe {
            let _ = KeWaitForSingleObject(
                &raw mut self.event as PVOID,
                _KWAIT_REASON::Executive,
                KERNEL_MODE,
                false as BOOLEAN,
                ptr::null::<LARGE_INTEGER>() as PLARGE_INTEGER,
            );
            (*self.irp).IoStatus.__bindgen_anon_1.Status
        }
    }
}
```

## WSK Socket 类型标志

```c
#define WSK_FLAG_BASIC_SOCKET        0x00000000  // 基本套接字
#define WSK_FLAG_LISTEN_SOCKET       0x00000001  // 监听套接字 (TCP Server)
#define WSK_FLAG_DATAGRAM_SOCKET     0x00000002  // 数据报套接字 (UDP)
#define WSK_FLAG_CONNECTION_SOCKET   0x00000004  // 连接套接字 (TCP Client)
```

---

# 第五部分：验证清单

## 每个文件完成后检查

- [ ] 所有结构体字段完整
- [ ] 所有函数签名正确
- [ ] 所有常量值正确
- [ ] 编译无错误
- [ ] 使用 WDK 绑定而非手动声明

## 禁止修改的文件

- `rust-driver/net/src/logger.rs` - net 模块内的日志宏定义
- `rust-driver/logger/src/lib.rs` - 独立 logger crate 的日志宏定义

## 禁止查看的未验证模块

- `hyperhv` - Hypervisor 核心
- `hyperkd` - 内核调试器
- `disassembler` - 反汇编器
- `pdbex` - PDB 解析

---

# 第六部分：错误处理

遇到问题时：

1. 检查类型是否匹配
2. 检查生命周期是否正确
3. 检查 unsafe 块是否正确
4. 检查模块导入是否正确
5. 参考 C++ 原始实现
6. 检查 WDK 绑定文件中的函数签名

## 输出格式

每完成一个文件，报告：
```
✓ 完成: <C++文件> -> <Rust文件>
  - 结构体: <数量>
  - 函数: <数量>
  - 常量: <数量>
  - 编译状态: <通过/失败>
```
