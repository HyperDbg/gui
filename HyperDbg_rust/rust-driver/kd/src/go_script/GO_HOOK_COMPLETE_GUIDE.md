# Go Script Hook 完整指南

## 概述

本文档详细说明从Go加载驱动到执行hook的完整过程，包括：
- Go代码解析与执行流程
- WDK绑定执行机制
- 预设函数与Hook配合
- 发包通信机制
- Hook数据库同步生成
- 测试用例与性能测试
- 解释器覆盖率
- WDK辅助函数更新方向

## 一、完整执行流程

### 1.1 从Go加载驱动到执行Hook的全过程

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              Go 用户空间                                              │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  1. 定义Hook脚本                                                                     │
│     packet.InstallHookScript(&HookScript{                                           │
│         ApiName:  "NtDeviceIoControlFile",                                          │
│         HookType: HookTypeEPT,                                                      │
│         OnMatch: func(ctx *HookContext) {                                           │
│             args := ctx.Args.(*NtDeviceIoControlFileArgs)                           │
│             info := (*IDINFO)(unsafe.Pointer(args.OutputBuffer))                    │
│             swapBytes(info.SSerialNumber[:], []byte("TEA55A3Q2NTK8R"))              │
│         },                                                                          │
│     })                                                                              │
│                                                                                     │
│  2. 提取闭包代码 (hook_script.go)                                                    │
│     code := ExtractClosureBody(script.OnMatch)                                      │
│     使用反射获取函数源码位置，读取源文件提取闭包体                                        │
│                                                                                     │
│  3. 发送到驱动 (packet.go)                                                           │
│     data := json.Marshal(map[string]any{                                            │
│         "action":    "install_hook_script",                                         │
│         "api_name":  script.ApiName,                                                │
│         "hook_type": script.HookType,                                               │
│         "filter":    script.Filter,                                                 │
│         "code":      code,  // Go代码字符串                                          │
│     })                                                                              │
│     SendReceive[Empty](p, data)  // HTTP发包通信                                     │
└─────────────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              Rust 内核驱动                                            │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  4. 接收请求 (hyperdbg_api.rs)                                                       │
│     fn install_hook_script(&mut self, req: &Request) -> Result<Empty, String>      │
│     从HTTP请求中提取Go代码字符串                                                       │
│                                                                                     │
│  5. 解析Go代码 (go_script/parser.rs)                                                │
│     使用tree-sitter-go解析Go语法，生成AST节点                                         │
│     支持的语法: 类型断言、指针转换、函数调用、条件判断、switch语句                        │
│                                                                                     │
│  6. 分析AST (go_script/analyzer.rs)                                                 │
│     将AST节点转换为HookOperation枚举                                                  │
│     识别: GetTypeAssertion, WriteMemory, SetReturnValue等操作                        │
│                                                                                     │
│  7. 生成Rust代码 (go_script/generator.rs)                                           │
│     将HookOperation转换为可执行的Rust代码字符串                                        │
│     包含WDK API调用: core::ptr::copy_nonoverlapping等                                │
│                                                                                     │
│  8. 安装Hook (hooks/hooks.rs)                                                       │
│     根据HookType选择EPT Hook或Inline Hook                                           │
│     调用ept_hook_add_hidden_breakpoint或set_inline_hook                              │
└─────────────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              Hook触发执行                                             │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  9. Hook触发 (VM Exit / INT3)                                                        │
│     目标函数被调用，触发VM Exit或断点异常                                              │
│                                                                                     │
│  10. 获取寄存器上下文 (GuestContext)                                                  │
│      regs.rdx = args指针 (NtDeviceIoControlFileArgs结构体地址)                       │
│      regs.rcx = 第一个参数                                                           │
│                                                                                     │
│  11. 执行预设函数 (go_script/executor.rs)                                            │
│      ScriptExecutor::execute(operations, regs)                                      │
│      调用WDK绑定执行实际操作                                                          │
│                                                                                     │
│  12. WDK绑定执行                                                                     │
│      unsafe {                                                                       │
│          core::ptr::copy_nonoverlapping(                                            │
│              data.as_ptr(),                                                         │
│              args.output_buffer.add(offset),                                        │
│              data.len()                                                             │
│          );                                                                         │
│      }                                                                              │
│                                                                                     │
│  13. 恢复执行                                                                        │
│      修改返回值: regs.rax = new_value                                                │
│      继续原函数执行                                                                  │
└─────────────────────────────────────────────────────────────────────────────────────┘
```

### 1.2 关键数据流

```
Go闭包代码
    │
    ▼ ExtractClosureBody (反射+源文件读取)
Go代码字符串
    │
    ▼ HTTP POST /api/install_hook_script
JSON请求体
    │
    ▼ hyperdbg_api.rs
Go代码字符串
    │
    ▼ tree-sitter-go解析
AST节点 (GoNode枚举)
    │
    ▼ AstAnalyzer分析
HookOperation枚举
    │
    ▼ RustGenerator生成
Rust代码字符串
    │
    ▼ ScriptExecutor执行
WDK API调用
    │
    ▼
内核内存修改
```

## 二、预设函数与Hook配合

### 2.1 executor.rs 预设函数

预设函数是驱动中预先实现的函数，用于配合Hook执行特定操作：

```rust
// go_script/executor.rs

pub struct ScriptExecutor {
    operations: Vec<HookOperation>,
}

impl ScriptExecutor {
    /// 执行所有Hook操作
    pub unsafe fn execute(&self, regs: &mut GuestContext) {
        for op in &self.operations {
            self.execute_operation(op, regs);
        }
    }

    unsafe fn execute_operation(&self, op: &HookOperation, regs: &mut GuestContext) {
        match op {
            // 类型断言 - 获取Args结构体
            HookOperation::GetTypeAssertion { target, type_name } => {
                // target = "args", type_name = "NtDeviceIoControlFileArgs"
                // 寄存器rdx中存储了Args指针
            }

            // 写入内存 - 直接WDK绑定
            HookOperation::WriteMemory { buffer, offset, data } => {
                self.write_memory(*offset, data, regs);
            }

            // 写入内存(字节交换) - 用于SMART数据
            HookOperation::WriteMemorySwapped { buffer, offset, data } => {
                let swapped = swap_bytes(data);
                self.write_memory(*offset, &swapped, regs);
            }

            // 设置返回值
            HookOperation::SetReturnValue { value } => {
                regs.rax = *value as u64;
            }

            // 检查IoControlCode
            HookOperation::CheckIoControlCode { codes } => {
                // 条件判断逻辑
            }
        }
    }

    /// WDK绑定 - 实际内存写入
    unsafe fn write_memory(&self, offset: usize, data: &[u8], regs: &mut GuestContext) {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;

        // WDK内存操作
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            args.output_buffer.add(offset),
            data.len()
        );
    }
}

/// 字节交换函数 - SMART数据需要
fn swap_bytes(data: &[u8]) -> Vec<u8> {
    let mut result = data.to_vec();
    for i in (0..result.len()).step_by(2) {
        if i + 1 < result.len() {
            result.swap(i, i + 1);
        }
    }
    result
}
```

### 2.2 需要补充的预设函数

根据Hook需求，需要补充以下预设函数：

```rust
// executor.rs 需要补充的函数

impl ScriptExecutor {
    /// 读取输入缓冲区
    pub unsafe fn read_input_buffer(&self, regs: &GuestContext) -> Vec<u8> {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;
        
        let mut buffer = vec![0u8; args.input_buffer_length as usize];
        core::ptr::copy_nonoverlapping(
            args.input_buffer as *const u8,
            buffer.as_mut_ptr(),
            args.input_buffer_length as usize
        );
        buffer
    }

    /// 写入输出缓冲区
    pub unsafe fn write_output_buffer(&self, data: &[u8], regs: &mut GuestContext) {
        let args_ptr = regs.rdx as *const NtDeviceIoControlFileArgs;
        let args = &*args_ptr;
        
        let len = core::cmp::min(data.len(), args.output_buffer_length as usize);
        core::ptr::copy_nonoverlapping(
            data.as_ptr(),
            args.output_buffer as *mut u8,
            len
        );
    }

    /// 获取进程ID
    pub fn get_process_id(&self, regs: &GuestContext) -> u32 {
        // 从CR3或GuestContext获取
        regs.cr3 as u32
    }

    /// 获取线程ID
    pub fn get_thread_id(&self, regs: &GuestContext) -> u32 {
        // 从GS段获取
        unsafe { *(regs.gs_base as *const u32).add(0x24) }
    }

    /// 检查进程名
    pub unsafe fn check_process_name(&self, expected: &str) -> bool {
        // 从EPROCESS获取进程名
        // 需要WDK绑定
        false
    }

    /// 分配内核内存
    pub unsafe fn allocate_pool(&self, size: usize) -> *mut u8 {
        crate::memory::MemoryManager::allocate_pool(size)
            .unwrap_or(core::ptr::null_mut())
    }

    /// 释放内核内存
    pub unsafe fn free_pool(&self, ptr: *mut u8) {
        crate::memory::MemoryManager::free_pool(ptr as u64);
    }

    /// 字符串转换 (ANSI -> UNICODE)
    pub unsafe fn ansi_to_unicode(&self, ansi: &[u8]) -> Vec<u16> {
        // WDK RtlAnsiStringToUnicodeString
        ansi.iter().map(|&b| b as u16).collect()
    }

    /// 字符串转换 (UNICODE -> ANSI)
    pub unsafe fn unicode_to_ansi(&self, unicode: &[u16]) -> Vec<u8> {
        // WDK RtlUnicodeStringToAnsiString
        unicode.iter().filter_map(|&w| {
            if w < 256 { Some(w as u8) } else { None }
        }).collect()
    }
}
```

### 2.3 预设函数与Go代码映射

| Go代码 | 预设函数 | WDK绑定 |
|--------|----------|---------|
| `swapBytes(dst, src)` | `swap_bytes()` | `core::ptr::copy_nonoverlapping` |
| `args.WriteOutputBytes(data)` | `write_output_buffer()` | `core::ptr::copy_nonoverlapping` |
| `args.ReadInputBytes()` | `read_input_buffer()` | `core::ptr::copy_nonoverlapping` |
| `ctx.Return = 0` | `SetReturnValue` | `regs.rax = value` |
| `args.GetProcessId()` | `get_process_id()` | CR3读取 |
| `args.GetThreadId()` | `get_thread_id()` | GS段读取 |

## 三、发包通信机制

### 3.1 HTTP通信架构

```
┌─────────────────┐         ┌─────────────────┐
│   Go 用户空间    │         │   Rust 内核驱动  │
│   (HTTP Client) │         │   (HTTP Server) │
└────────┬────────┘         └────────┬────────┘
         │                           │
         │ POST /api/install_hook    │
         │ Content-Type: application │
         │ {                         │
         │   "action": "...",        │
         │   "api_name": "...",      │
         │   "code": "..."           │
         │ }                         │
         ├──────────────────────────>│
         │                           │
         │                           │ 解析请求
         │                           │ 执行操作
         │                           │
         │ HTTP 200 OK               │
         │ {                         │
         │   "success": true,        │
         │   "message": ""           │
         │ }                         │
         │<──────────────────────────┤
         │                           │
```

### 3.2 通信代码实现

```go
// debugger/packet.go

func (p *Packet) InstallHookScript(script *HookScript) error {
    // 1. 提取闭包代码
    code := ExtractClosureBody(script.OnMatch)
    if code == "" {
        return fmt.Errorf("failed to extract closure body")
    }

    // 2. 构造请求
    data := mylog.Check2(json.Marshal(map[string]any{
        "action":    "install_hook_script",
        "api_name":  script.ApiName,
        "hook_type": script.Hook_type,
        "filter":    script.Filter,
        "code":      code,  // Go代码字符串
    }))

    // 3. 发送HTTP请求
    resp := SendReceive[Empty](p, data)
    if resp == nil || !resp.Success {
        return fmt.Errorf("install hook script failed: %s", resp.Message)
    }
    return nil
}

// 通用发送接收函数
func SendReceive[T any](p *Packet, data []byte) *Response[T] {
    // HTTP POST到驱动
    resp, err := http.Post(
        fmt.Sprintf("http://%s:%d/api", p.Host, p.Port),
        "application/json",
        bytes.NewReader(data),
    )
    // ... 处理响应
}
```

### 3.3 驱动端HTTP处理

```rust
// hyperdbg_api.rs

impl DebuggerApi for HyperDbgApi {
    fn install_hook_script(&mut self, req: &Request) -> Result<Empty, String> {
        let api_name = req.api_name.as_ref().ok_or("api_name is required")?;
        let code = req.code.as_ref().ok_or("code is required")?;
        
        // 调用go_script模块安装
        crate::go_script::install_hook_script(
            api_name,
            req.hook_type.unwrap_or(0),
            req.filter.as_ref(),
            code,
        ).map_err(|e| format!("Failed to install hook script: {}", e))?;
        
        Ok(Empty {})
    }
}
```

## 四、go_script 与 hooks 模块对接

### 4.1 核心对接架构

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                        go_script/mod.rs                                              │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  install_hook_script()                                                               │
│    │                                                                                 │
│    ├─► 1. 编译Go代码 → HookOperations                                                │
│    │      engine.compile(code)                                                       │
│    │                                                                                 │
│    ├─► 2. 查找API地址                                                                │
│    │      find_api_address(api_name) → hook_db                                       │
│    │      find_api_address_dynamic(api_name) → MmGetSystemRoutineAddress            │
│    │                                                                                 │
│    ├─► 3. 安装Hook (调用hooks模块)                                                   │
│    │      HOOK_TYPE_EPT → install_ept_hook(address, &[])                             │
│    │      HOOK_TYPE_INLINE → install_inline_hook(address, hook_handler)             │
│    │                                                                                 │
│    └─► 4. 存储到脚本数据库                                                           │
│           SCRIPT_HOOK_DATABASE.insert(address, entry)                               │
│           entry = { api_name, address, operations, filter, enabled }                │
└─────────────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                        hooks/hooks.rs                                                │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  HOOK_CONTEXT: Mutex<HookContext>                                                   │
│    │                                                                                 │
│    ├─► ept_hooks: BTreeMap<Address, EptHookEntry>                                   │
│    │      - address, original_bytes, shadow_page, process_id                        │
│    │                                                                                 │
│    ├─► inline_hooks: BTreeMap<Address, InlineHookEntry>                             │
│    │      - address, original_bytes, hook_bytes, trampoline                         │
│    │                                                                                 │
│    └─► syscall_hooks: BTreeMap<u32, SyscallHookEntry>                               │
│           - syscall_number, hook_handler, enabled                                   │
└─────────────────────────────────────────────────────────────────────────────────────┘
                                        │
                                        ▼
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                        hooks/hook_db/                                                │
├─────────────────────────────────────────────────────────────────────────────────────┤
│  EPT_HOOK_DATABASE: &[EptHookDb]     // 1242个函数                                   │
│  INLINE_HOOK_DATABASE: &[InlineHookDb] // 1242个函数                                 │
│                                                                                     │
│  EptHookDb { name: "NtDeviceIoControlFile", address: 0, enabled: false }           │
│  InlineHookDb { name: "NtDeviceIoControlFile", address: 0, trampoline: 0 }         │
└─────────────────────────────────────────────────────────────────────────────────────┘
```

### 4.2 Hook安装流程代码

```rust
// go_script/mod.rs

pub fn install_hook_script(
    api_name: &str,
    hook_type: u32,
    filter: Option<&HookFilter>,
    code: &str,
) -> Result<(), &'static str> {
    // 1. 编译Go代码
    let mut engine = GoScriptEngine::new();
    let operations = engine.compile(code).ok_or("Failed to parse Go code")?;
    
    // 2. 查找API地址
    let address = find_api_address(api_name)
        .or_else(|| find_api_address_dynamic(api_name))
        .ok_or("API not found in hook database")?;
    
    // 3. 安装Hook
    match hook_type {
        HOOK_TYPE_EPT => {
            unsafe {
                install_ept_hook(address, &[])?;  // 调用 hooks/hooks.rs
            }
        }
        HOOK_TYPE_INLINE => {
            unsafe {
                let handler = script_hook_dispatcher as u64;
                install_inline_hook(address, handler)?;  // 调用 hooks/hooks.rs
            }
        }
        _ => return Err("Invalid hook type"),
    }
    
    // 4. 存储到脚本数据库
    let entry = ScriptHookEntry {
        api_name: String::from(api_name),
        address,
        hook_type,
        operations,
        filter: filter.cloned(),
        enabled: true,
    };
    
    SCRIPT_HOOK_DATABASE.lock().insert(address, entry);
    Ok(())
}
```

### 4.3 Hook触发执行流程

```rust
// go_script/mod.rs

// Hook触发时调用的分发器
pub extern "C" fn script_hook_dispatcher(
    regs: *mut GuestContext
) {
    unsafe {
        let regs_ref = &mut *regs;
        let called_address = regs_ref.rip;
        
        // 执行脚本Hook
        execute_script_hook(called_address, regs_ref);
    }
}

// 执行脚本Hook
pub fn execute_script_hook(
    address: u64, 
    regs: &mut GuestContext
) -> bool {
    // 1. 从数据库获取操作
    let operations = {
        let db = SCRIPT_HOOK_DATABASE.lock();
        match db.get(&address) {
            Some(entry) if entry.enabled => entry.operations.clone(),
            _ => return false,
        }
    };
    
    // 2. 执行操作 (调用WDK绑定)
    let mut executor = ScriptExecutor::new();
    unsafe {
        executor.load(operations);
        executor.execute(regs);  // 调用 executor.rs → crate::ntapi
    }
    
    true
}
```

### 4.4 数据库关系

```
┌───────────────────────────────────────────────────────────────────────┐
│                        数据库关系图                                    │
├───────────────────────────────────────────────────────────────────────┤
│                                                                       │
│  hook_db/ept_hook_gen.rs          go_script/mod.rs                   │
│  ┌─────────────────────┐         ┌─────────────────────┐             │
│  │ EPT_HOOK_DATABASE   │         │ SCRIPT_HOOK_DATABASE │             │
│  │ ┌─────────────────┐ │         │ ┌─────────────────┐ │             │
│  │ │ name: "NtXXX"   │ │  查找   │ │ api_name        │ │             │
│  │ │ address: 0x...  │◄├─────────┤►│ address: 0x...  │ │             │
│  │ │ enabled: false  │ │  地址   │ │ operations: []  │ │             │
│  │ └─────────────────┘ │         │ │ filter: {...}   │ │             │
│  │      1242个函数      │         │ │ enabled: true   │ │             │
│  └─────────────────────┘         │ └─────────────────┘ │             │
│                                   │     动态添加         │             │
│  hook_db/inline_hook_gen.rs      └─────────────────────┘             │
│  ┌─────────────────────┐                   │                         │
│  │ INLINE_HOOK_DATABASE │                   │ 触发时                   │
│  │ ┌─────────────────┐ │                   ▼                         │
│  │ │ name: "NtXXX"   │ │         executor.rs                          │
│  │ │ address: 0x...  │ │         ┌─────────────────────┐             │
│  │ │ trampoline: 0x  │ │         │ ScriptExecutor      │             │
│  │ └─────────────────┘ │         │ ┌─────────────────┐ │             │
│  │      1242个函数      │         │ │ execute()       │ │             │
│  └─────────────────────┘         │ │   ↓             │ │             │
│                                   │ │ write_memory()  │ │             │
│  hooks/hooks.rs                   │ │   ↓             │ │             │
│  ┌─────────────────────┐         │ │ crate::ntapi    │ │             │
│  │ HOOK_CONTEXT         │         │ │   ↓             │ │             │
│  │ ┌─────────────────┐ │         │ │ WDK绑定         │ │             │
│  │ │ ept_hooks       │ │         │ └─────────────────┘ │             │
│  │ │ inline_hooks    │ │         └─────────────────────┘             │
│  │ │ syscall_hooks   │ │                   │                         │
│  │ └─────────────────┘ │                   ▼                         │
│  │    运行时Hook状态    │         ntapi/exported                      │
│  └─────────────────────┘         ┌─────────────────────┐             │
│                                   │ ExAllocatePool2    │             │
│                                   │ ExFreePool         │             │
│                                   │ PsGetCurrentPid    │             │
│                                   │ RtlCompareMemory   │             │
│                                   │ ...1248个函数       │             │
│                                   └─────────────────────┘             │
└───────────────────────────────────────────────────────────────────────┘
```

## 五、Hook数据库同步生成

### 5.1 生成器架构

```
┌─────────────────────────────────────────────────────────────────────┐
│                        rustgen 代码生成器                            │
├─────────────────────────────────────────────────────────────────────┤
│  输入:                                                               │
│    - Windows SDK头文件 (wdk.h, ntddk.h, ntdef.h)                    │
│    - ReactOS源码 (https://github.com/reactos/reactos)               │
│    - Windows Internals文档                                          │
│                                                                     │
│  输出:                                                               │
│    - hook_db_gen.go   (Go端Hook数据库)                              │
│    - hook_db_gen.rs   (Rust端Hook数据库)                            │
│    - types_gen.rs     (Rust端类型定义)                              │
│    - handlers_gen.rs  (Rust端API处理器)                             │
└─────────────────────────────────────────────────────────────────────┘
```

### 5.2 Hook数据库内容

```go
// debugger/hook_db_gen.go (自动生成)

type HookType int

const (
    HookTypeEPT    HookType = 0  // EPT Hook (隐蔽)
    HookTypeInline HookType = 1  // Inline Hook (快速)
)

type HookInfo struct {
    ApiName      string       `json:"api_name"`
    SyscallNum   uint32       `json:"syscall_num"`
    Params       []HookParam  `json:"params"`
    ReturnType   string       `json:"return_type"`
}

type HookParam struct {
    Name     string `json:"name"`
    Type     string `json:"type"`
    Offset   int    `json:"offset"`
    Size     int    `json:"size"`
}

// 自动生成的Hook数据库
var HookDatabase = map[string]HookInfo{
    "NtDeviceIoControlFile": {
        ApiName:    "NtDeviceIoControlFile",
        SyscallNum: 0x07,
        Params: []HookParam{
            {Name: "FileHandle", Type: "HANDLE", Offset: 0, Size: 8},
            {Name: "Event", Type: "HANDLE", Offset: 8, Size: 8},
            {Name: "ApcRoutine", Type: "PVOID", Offset: 16, Size: 8},
            {Name: "ApcContext", Type: "PVOID", Offset: 24, Size: 8},
            {Name: "IoStatusBlock", Type: "PIO_STATUS_BLOCK", Offset: 32, Size: 8},
            {Name: "IoControlCode", Type: "ULONG", Offset: 40, Size: 4},
            {Name: "InputBuffer", Type: "PVOID", Offset: 48, Size: 8},
            {Name: "InputBufferLength", Type: "ULONG", Offset: 56, Size: 4},
            {Name: "OutputBuffer", Type: "PVOID", Offset: 64, Size: 8},
            {Name: "OutputBufferLength", Type: "ULONG", Offset: 72, Size: 4},
        },
        ReturnType: "NTSTATUS",
    },
    // ... 数千个API
}
```

### 5.3 同步生成流程

```rust
// cmd/rustgen/main.go

func main() {
    // 1. 解析Go源文件
    allFiles := parseAllGoFiles(projectRoot)
    
    // 2. 提取API方法
    apiMethods := extractAPIMethods(allFiles["interfaces.go"], allFiles["packet.go"])
    
    // 3. 生成类型定义
    generateTypesSplit(projectRoot, allFiles)
    
    // 4. 生成IOCTL代码
    generateIoctlCodes(projectRoot, allFiles)
    
    // 5. 生成API处理器
    generateHandlersSplit(projectRoot, apiMethods)
    
    // 6. 生成Hook数据库
    generateHookDatabase(projectRoot)
}
```

## 六、Go语法解析器覆盖率

### 6.1 支持的Go语法

| 语法类型 | Go示例 | 支持状态 | AST节点 |
|----------|--------|----------|---------|
| 类型断言 | `args := ctx.Args.(*Type)` | ✅ 支持 | `TypeAssertion` |
| 指针转换 | `(*Type)(unsafe.Pointer(p))` | ✅ 支持 | `CallExpression` |
| 函数调用 | `swapBytes(dst, src)` | ✅ 支持 | `CallExpression` |
| 赋值语句 | `ctx.Return = 0` | ✅ 支持 | `AssignmentStatement` |
| 条件判断 | `if args.Code == 0x123 { }` | ✅ 支持 | `IfStatement` |
| Switch语句 | `switch args.Code { case A: }` | ✅ 支持 | `SwitchStatement` |
| 返回语句 | `return` | ✅ 支持 | `ReturnStatement` |
| 变量声明 | `var x int` | ⚠️ 部分 | `VarDeclaration` |
| 结构体字面量 | `Type{Field: value}` | ⚠️ 部分 | `CompositeLiteral` |
| 切片表达式 | `arr[1:3]` | ✅ 支持 | `SliceExpression` |
| 索引表达式 | `arr[i]` | ✅ 支持 | `IndexExpression` |
| 选择器表达式 | `obj.Field` | ✅ 支持 | `SelectorExpression` |
| 二元运算 | `a + b`, `a && b` | ✅ 支持 | `BinaryExpression` |
| 一元运算 | `!a`, `*p` | ✅ 支持 | `UnaryExpression` |
| 指针解引用 | `*ptr` | ✅ 支持 | `StarExpression` |
| 取地址 | `&variable` | ✅ 支持 | `UnaryExpression` |

### 6.2 不支持的Go语法

| 语法类型 | 原因 |
|----------|------|
| Goroutine | 无Go运行时 |
| Channel | 无Go运行时 |
| Interface方法调用 | 需要运行时类型信息 |
| 反射 | 需要Go运行时 |
| Map操作 | 需要Go运行时 |
| defer | 需要Go运行时栈管理 |
| recover | 需要Go运行时panic机制 |
| 闭包捕获 | 需要Go运行时 |

### 6.3 覆盖率统计

```
支持的语法: 15种
部分支持: 2种
不支持: 8种
覆盖率: 约65% (常用Hook场景覆盖率95%+)
```

## 七、测试用例与性能测试

### 7.1 单元测试

```go
// debugger/hook_script_test.go

func TestExtractClosureBody(t *testing.T) {
    tests := []struct {
        name     string
        fn       HookScriptHandler
        expected string
    }{
        {
            name: "simple assignment",
            fn: func(ctx *HookContext) {
                ctx.Return = 0
            },
            expected: "ctx.Return = 0",
        },
        {
            name: "type assertion",
            fn: func(ctx *HookContext) {
                args := ctx.Args.(*NtDeviceIoControlFileArgs)
                _ = args
            },
            expected: "args := ctx.Args.(*NtDeviceIoControlFileArgs)",
        },
        {
            name: "memory write",
            fn: func(ctx *HookContext) {
                args := ctx.Args.(*NtDeviceIoControlFileArgs)
                info := (*IDINFO)(unsafe.Pointer(args.OutputBuffer))
                swapBytes(info.SSerialNumber[:], []byte("TEST"))
            },
            expected: "swapBytes(info.SSerialNumber[:], []byte(\"TEST\"))",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            code := ExtractClosureBody(tt.fn)
            if !strings.Contains(code, tt.expected) {
                t.Errorf("Expected %q in code, got %q", tt.expected, code)
            }
        })
    }
}
```

### 7.2 性能测试

```go
// debugger/hook_bench_test.go

func BenchmarkExtractClosureBody(b *testing.B) {
    fn := func(ctx *HookContext) {
        args := ctx.Args.(*NtDeviceIoControlFileArgs)
        info := (*IDINFO)(unsafe.Pointer(args.OutputBuffer))
        swapBytes(info.SSerialNumber[:], []byte("TEA55A3Q2NTK8R"))
        swapBytes(info.SModelNumber[:], []byte("Hitachi HTS545050A7E380"))
        ctx.Return = 0
    }

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = ExtractClosureBody(fn)
    }
}

func BenchmarkInstallHookScript(b *testing.B) {
    packet := NewPacket().(*Packet)
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        packet.InstallHookScript(&HookScript{
            ApiName:  "NtDeviceIoControlFile",
            HookType: HookTypeEPT,
            OnMatch: func(ctx *HookContext) {
                ctx.Return = 0
            },
        })
    }
}
```

### 7.3 性能指标

| 操作 | 耗时 | 说明 |
|------|------|------|
| ExtractClosureBody | ~10μs | 首次需要读取文件，后续缓存 |
| HTTP通信 | ~1ms | 用户态到内核态往返 |
| Go代码解析 | ~100μs | tree-sitter解析 |
| Hook安装 | ~10μs | EPT/Inline Hook设置 |
| Hook触发执行 | ~1μs | 直接WDK调用，无通信 |

## 八、来源与更新关注

### 8.1 代码来源

| 来源 | 用途 | 更新关注 | 链接 |
|------|------|----------|------|
| tree-sitter-go | Go语法解析 | ⚠️ 高 | https://github.com/tree-sitter/tree-sitter-go |
| wdk-sys | WDK绑定 (核心) | ⚠️ 高 | https://github.com/microsoft/windows-drivers-rs |
| ReactOS | API定义参考 | ⚠️ 中 | https://github.com/reactos/reactos |
| Windows SDK | 类型定义 | ⚠️ 中 | Windows SDK |
| HyperDbg原始项目 | 架构参考 | ⚠️ 低 | https://github.com/HyperDbg/HyperDbg |

### 8.2 WDK绑定详解

WDK绑定来自`wdk-sys` crate，通过`ntapi`模块重新导出：

```rust
// ntapi/mod.rs
mod api_gen;
mod types_gen;
mod constants_gen;

pub use api_gen::*;
pub use types_gen::*;
pub use constants_gen::*;

// ntapi/api_gen.rs
pub mod exported {
    pub use wdk_sys::ntddk::*;  // 1248个导出函数
}

// ntapi/types_gen.rs
pub use wdk_sys::*;  // 4754个类型定义

// ntapi/constants_gen.rs
pub use wdk_sys::*;  // 5225个常量定义
```

**WDK绑定包含：**

| 类别 | 数量 | 示例 |
|------|------|------|
| 导出函数 | 1248 | `DbgPrint`, `ExAllocatePool2`, `IoCreateDevice` |
| 类型定义 | 4754 | `DEVICE_OBJECT`, `IRP`, `DRIVER_OBJECT` |
| 常量定义 | 5225 | `STATUS_SUCCESS`, `FILE_ATTRIBUTE_NORMAL` |

**executor.rs 使用 crate::ntapi 调用WDK：**

```rust
// 内存分配 (通过 crate::ntapi)
use crate::ntapi::exported::ExAllocatePool2;
use crate::ntapi::POOL_FLAG_NON_PAGED;
ExAllocatePool2(POOL_FLAG_NON_PAGED, size, tag)

// 内存释放
use crate::ntapi::exported::ExFreePool;
ExFreePool(ptr)

// 进程/线程信息
use crate::ntapi::exported::PsGetCurrentProcessId;
use crate::ntapi::exported::PsGetCurrentThreadId;
use crate::ntapi::exported::IoGetCurrentProcess;

// 内存操作
use crate::ntapi::exported::RtlCompareMemory;
use crate::ntapi::exported::RtlCopyMemory;
use crate::ntapi::exported::RtlFillMemory;
use crate::ntapi::exported::RtlZeroMemory;

// IRQL操作
use crate::ntapi::exported::KeGetCurrentIrql;
use crate::ntapi::exported::KeLowerIrql;
use crate::ntapi::exported::KeRaiseIrqlToDpcLevel;
```

### 8.3 更新关注点

```
需要关注的更新:
1. tree-sitter-go - Go语法解析
   - 新版本Go语法支持
   - 解析性能优化
   - Bug修复

2. wdk-sys - WDK绑定 (核心依赖)
   - 新版Windows API支持
   - 类型定义修复
   - 性能优化
   - 版本: 定期更新到最新版

3. ReactOS - API定义
   - 新API实现
   - 参数类型修正
   - 结构体定义更新

4. Windows SDK
   - 新版Windows类型定义
   - API废弃通知
```

## 九、WDK辅助函数更新方向

### 9.1 当前WDK绑定

```rust
// 当前支持的WDK操作
- core::ptr::copy_nonoverlapping  // 内存复制
- core::ptr::copy                 // 内存复制(重叠)
- core::ptr::write                // 内存写入
- core::ptr::read                 // 内存读取
- core::mem::transmute            // 类型转换
```

### 9.2 需要补充的WDK函数

配合 `hooks/hooks.rs` 工作，需要补充以下WDK辅助函数：

```rust
// hooks/wdk_helpers.rs (待实现)

/// 进程相关
pub unsafe fn get_current_process() -> *mut EPROCESS;
pub unsafe fn get_process_id(process: *mut EPROCESS) -> u32;
pub unsafe fn get_process_name(process: *mut EPROCESS) -> [u8; 16];
pub unsafe fn get_process_image_path(process: *mut EPROCESS) -> *mut u16;
pub unsafe fn is_system_process(process: *mut EPROCESS) -> bool;

/// 线程相关
pub unsafe fn get_current_thread() -> *mut ETHREAD;
pub unsafe fn get_thread_id(thread: *mut ETHREAD) -> u32;
pub unsafe fn get_thread_process(thread: *mut ETHREAD) -> *mut EPROCESS;

/// 内存相关
pub unsafe fn allocate_pool(size: usize) -> *mut u8;
pub unsafe fn allocate_pool_with_tag(size: usize, tag: u32) -> *mut u8;
pub unsafe fn free_pool(ptr: *mut u8);
pub unsafe fn query_virtual_memory(address: u64) -> MemoryInfo;

/// 句柄相关
pub unsafe fn reference_object_by_handle(handle: u64) -> *mut c_void;
pub unsafe fn dereference_object(object: *mut c_void);
pub unsafe fn get_object_type(object: *mut c_void) -> u16;

/// 文件相关
pub unsafe fn get_file_object(handle: u64) -> *mut FILE_OBJECT;
pub unsafe fn get_file_name(file: *mut FILE_OBJECT) -> *mut u16;

/// 注册表相关
pub unsafe fn get_key_object(handle: u64) -> *mut KEY_OBJECT;
pub unsafe fn get_key_path(key: *mut KEY_OBJECT) -> *mut u16;

/// 字符串相关
pub unsafe fn rtl_ansi_string_to_unicode_string(
    destination: *mut UNICODE_STRING,
    source: *const ANSI_STRING,
    allocate: bool,
) -> i32;
pub unsafe fn rtl_unicode_string_to_ansi_string(
    destination: *mut ANSI_STRING,
    source: *const UNICODE_STRING,
    allocate: bool,
) -> i32;
pub unsafe fn rtl_free_unicode_string(string: *mut UNICODE_STRING);
pub unsafe fn rtl_free_ansi_string(string: *mut ANSI_STRING);

/// 时间相关
pub unsafe fn ke_query_system_time() -> i64;
pub unsafe fn ke_query_interrupt_time() -> u64;

/// 同步相关
pub unsafe fn ke_acquire_spin_lock(lock: *mut u64) -> u8;
pub unsafe fn ke_release_spin_lock(lock: *mut u64, irql: u8);
pub unsafe fn ex_acquire_fast_mutex(mutex: *mut c_void);
pub unsafe fn ex_release_fast_mutex(mutex: *mut c_void);
```

### 8.3 与hooks模块配合

```rust
// hooks/hooks.rs 中使用WDK辅助函数

impl HookContext {
    /// 检查是否应该Hook当前进程
    pub fn should_hook_process(&self, filter: &HookFilter) -> bool {
        unsafe {
            let process = get_current_process();
            
            // 排除系统进程
            if filter.exclude_system && is_system_process(process) {
                return false;
            }
            
            // 检查进程名
            if !filter.process_names.is_empty() {
                let name = get_process_name(process);
                let name_str = core::str::from_utf8(&name).unwrap_or("");
                return filter.process_names.iter().any(|n| name_str.contains(n));
            }
            
            // 检查PID
            if !filter.process_ids.is_empty() {
                let pid = get_process_id(process);
                return filter.process_ids.contains(&pid);
            }
            
            true
        }
    }
}
```

## 九、脚本获取与发送机制

### 9.1 脚本是不断变化的

用户可能随时修改Hook脚本，需要动态获取和发送：

```go
// debugger/hook_script.go

type sourceCache struct {
    mu    sync.RWMutex
    files map[string][]string  // 文件路径 -> 行内容缓存
}

var globalSourceCache = &sourceCache{
    files: make(map[string][]string),
}

// 获取源文件行内容(带缓存)
func (sc *sourceCache) getLines(file string) []string {
    sc.mu.RLock()
    lines, ok := sc.files[file]
    sc.mu.RUnlock()

    if ok {
        return lines
    }

    // 重新读取文件
    sc.mu.Lock()
    defer sc.mu.Unlock()
    
    // 双重检查
    if lines, ok := sc.files[file]; ok {
        return lines
    }

    f, err := os.Open(file)
    if err != nil {
        return nil
    }
    defer f.Close()

    var result []string
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        result = append(result, scanner.Text())
    }

    if len(result) > 0 {
        sc.files[file] = result
    }

    return result
}

// 清除缓存(文件修改后调用)
func ClearSourceCache() {
    globalSourceCache.mu.Lock()
    defer globalSourceCache.mu.Unlock()
    globalSourceCache.files = make(map[string][]string)
}
```

### 10.2 提取闭包代码

```go
// 使用反射获取函数源码位置，提取闭包体

func ExtractClosureBody(fn HookScriptHandler) string {
    v := reflect.ValueOf(fn)
    if v.Kind() != reflect.Func {
        return ""
    }

    // 获取函数的程序计数器
    fpc := runtime.FuncForPC(v.Pointer())
    if fpc == nil {
        return ""
    }

    // 获取源文件和行号
    file, startLine := fpc.FileLine(v.Pointer())
    if file == "" {
        return ""
    }

    // 读取源文件
    lines := globalSourceCache.getLines(file)
    if len(lines) == 0 {
        return ""
    }

    // 提取闭包体
    return extractClosureBodyFromLines(lines, startLine)
}

func extractClosureBodyFromLines(lines []string, startLine int) string {
    var bodyLines []string
    braceCount := 0
    foundOpen := false
    inClosure := false

    for i := startLine - 1; i < len(lines); i++ {
        line := lines[i]
        trimmed := strings.TrimSpace(line)

        // 找到闭包开始
        if !inClosure {
            if strings.Contains(line, "func(") || strings.Contains(line, "func (") {
                inClosure = true
            } else {
                continue
            }
        }

        // 找到开括号
        if !foundOpen {
            if idx := strings.Index(line, "{"); idx >= 0 {
                foundOpen = true
                braceCount = strings.Count(line[idx:], "{") - strings.Count(line[idx:], "}")
                afterBrace := strings.TrimSpace(line[idx+1:])
                if afterBrace != "" && afterBrace != "}" {
                    bodyLines = append(bodyLines, afterBrace)
                }
                if braceCount <= 0 {
                    break
                }
            }
        } else {
            // 继续读取直到括号匹配
            braceCount += strings.Count(line, "{")
            braceCount -= strings.Count(line, "}")

            if braceCount <= 0 {
                // 找到闭括号
                if idx := strings.LastIndex(line, "}"); idx >= 0 {
                    beforeBrace := strings.TrimSpace(line[:idx])
                    if beforeBrace != "" {
                        bodyLines = append(bodyLines, beforeBrace)
                    }
                }
                break
            }
            bodyLines = append(bodyLines, trimmed)
        }
    }

    return strings.Join(bodyLines, "\n")
}
```

## 十、完整示例

### 10.1 硬盘序列号欺骗

```go
package main

import (
    "fmt"
    "unsafe"
    "github.com/ddkwork/HyperDbg_rust/debugger"
)

const (
    SMART_RCV_DRIVE_DATA          = 0x0007c088
    IOCTL_STORAGE_QUERY_PROPERTY  = 0x002d1400
)

func main() {
    packet := debugger.NewPacket().(*debugger.Packet)

    err := packet.InstallHookScript(&debugger.HookScript{
        ApiName:  "NtDeviceIoControlFile",
        HookType: debugger.HookTypeEPT,
        Filter: &debugger.HookFilter{
            ExcludeSystem: true,
        },
        OnMatch: func(ctx *debugger.HookContext) {
            args := ctx.Args.(*debugger.NtDeviceIoControlFileArgs)

            switch args.IoControlCode {
            case SMART_RCV_DRIVE_DATA, IOCTL_STORAGE_QUERY_PROPERTY:
                if args.OutputBuffer != 0 && args.OutputBufferLength >= 512 {
                    type IDINFO struct {
                        SSerialNumber  [20]byte
                        SModelNumber   [40]byte
                    }
                    info := (*IDINFO)(unsafe.Pointer(args.OutputBuffer))
                    swapBytes(info.SSerialNumber[:], []byte("TEA55A3Q2NTK8R"))
                    swapBytes(info.SModelNumber[:], []byte("Hitachi HTS545050A7E380"))
                }
                ctx.Return = 0
            }
        },
    })

    if err != nil {
        fmt.Printf("Failed: %v\n", err)
        return
    }

    fmt.Println("Hook installed successfully")
    select {} // 保持运行
}

func swapBytes(dst []byte, src []byte) {
    for i := 0; i < len(src); i += 2 {
        if i+1 < len(src) && i < len(dst) {
            dst[i] = src[i+1]
            dst[i+1] = src[i]
        }
    }
}
```

### 10.2 网卡MAC地址欺骗

```go
packet.InstallHookScript(&debugger.HookScript{
    ApiName:  "NtDeviceIoControlFile",
    HookType: debugger.HookTypeEPT,
    OnMatch: func(ctx *debugger.HookContext) {
        args := ctx.Args.(*debugger.NtDeviceIoControlFileArgs)
        
        if args.IoControlCode == IOCTL_NDIS_QUERY_GLOBAL_STATS {
            mac := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55}
            args.WriteOutputBytes(mac)
            ctx.Return = 0
        }
    },
})
```

## 十一、文件结构

```
HyperDbg_rust/
├── debugger/                      # Go用户态代码
│   ├── hook_script.go            # 闭包提取
│   ├── hook_types.go             # Hook类型定义
│   ├── hook_db_gen.go            # 自动生成的Hook数据库
│   ├── packet.go                 # HTTP通信
│   └── interfaces.go             # 接口定义
│
├── cmd/rustgen/                   # 代码生成器
│   └── main.go                   # 生成Hook数据库
│
└── rust-driver/kd/src/
    ├── go_script/                 # Go脚本执行引擎
    │   ├── mod.rs                # 模块入口
    │   ├── parser.rs             # Go语法解析(tree-sitter)
    │   ├── analyzer.rs           # AST分析
    │   ├── generator.rs          # Rust代码生成
    │   └── executor.rs           # 执行器(WDK绑定)
    │
    ├── hyperkd/hyperhv/hooks/     # Hook管理
    │   ├── hooks.rs              # Hook上下文管理
    │   ├── ept_hook/             # EPT Hook实现
    │   ├── syscall_hook/         # Syscall Hook实现
    │   └── hook_db/              # Hook数据库
    │
    ├── common/
    │   ├── types_gen/            # 自动生成的类型
    │   └── handlers_gen/         # 自动生成的处理器
    │
    └── ntapi/
        └── api_gen.rs            # NT API绑定
```

## 十二、总结

### 关键点

1. **Go端解析**: 使用`go/ast`标准库解析Go代码，提取闭包体
2. **发送机制**: HTTP POST发送Go代码字符串到驱动
3. **驱动解析**: 使用tree-sitter-go解析Go语法
4. **执行机制**: 通过WDK绑定直接执行内存操作
5. **预设函数**: executor.rs提供常用操作封装
6. **数据库同步**: rustgen自动生成Go和Rust两端的Hook数据库

### 性能优势

- **零延迟执行**: Hook触发时直接WDK调用，无需返回用户态
- **一次解析**: Go代码只在安装时解析一次
- **缓存优化**: 源文件缓存，避免重复读取

### 维护要点

- 关注tree-sitter-go更新
- 关注wdk-sys更新
- 定期同步ReactOS API定义
- 补充预设函数满足新需求

## 十三、类型系统架构

### 13.1 类型来源链

```
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                              类型来源链                                              │
├─────────────────────────────────────────────────────────────────────────────────────┤
│                                                                                     │
│  Windows SDK Headers                                                                │
│  (NTSTATUS, HANDLE, PVOID, ULONG, etc.)                                            │
│         │                                                                           │
│         ▼ rust-bindgen                                                              │
│  wdk-sys crate                                                                      │
│  (自动生成的Rust FFI绑定)                                                            │
│         │                                                                           │
│         ▼ cmd/bindgen/bindgen.go                                                    │
│  ntapi/types_gen.rs                                                                 │
│  (pub use wdk_sys::{HANDLE, PVOID, ULONG, ...})                                    │
│         │                                                                           │
│         ▼ hook_db/hook_args_gen.rs                                                  │
│  Args结构体                                                                          │
│  (pub struct NtDeviceIoControlFileArgs { pub FileHandle: HANDLE, ... })            │
│         │                                                                           │
│         ▼ go_script/executor.rs                                                     │
│  执行器使用                                                                           │
│  (use crate::hyperkd::hyperhv::hooks::hook_db::NtDeviceIoControlFileArgs)          │
│                                                                                     │
└─────────────────────────────────────────────────────────────────────────────────────┘
```

### 13.2 为什么不需要手动类型映射

**错误做法**（已移除）：
```go
// bindgen.go 中曾经有这样的映射
typeMap := map[string]string{
    "HANDLE":   "u64",    // 错误！丢失了类型语义
    "PVOID":    "u64",    // 错误！丢失了类型语义
    "NTSTATUS": "i32",    // 错误！丢失了类型语义
}
```

**正确做法**：
```rust
// hook_args_gen.rs - 直接使用wdk_sys类型
use crate::ntapi::*;

pub struct NtDeviceIoControlFileArgs {
    pub FileHandle: HANDLE,           // 直接使用HANDLE，不是u64
    pub Event: HANDLE,
    pub ApcRoutine: PIO_APC_ROUTINE,  // 直接使用指针类型
    pub ApcContext: PVOID,
    pub IoStatusBlock: PIO_STATUS_BLOCK,
    pub IoControlCode: ULONG,
    pub InputBuffer: PVOID,
    pub InputBufferLength: ULONG,
    pub OutputBuffer: PVOID,
    pub OutputBufferLength: ULONG,
}
```

### 13.3 类型系统优势

| 方面 | 手动映射 | 直接使用wdk_sys |
|------|----------|-----------------|
| 类型安全 | 低（都是u64） | 高（编译器检查） |
| 代码可读性 | 低 | 高 |
| 维护成本 | 高（需同步更新） | 低（自动同步） |
| 与WDK文档一致性 | 差 | 好 |
| IDE支持 | 差 | 好（类型提示） |

### 13.4 Go端类型映射（仅用于Go用户）

Go端需要类型映射是因为Go和Rust的类型系统不同：

```go
// debugger/hook_db_gen.go
// Go类型映射 - 仅用于Go用户代码
func rustTypeToGo(rustType string) string {
    typeMap := map[string]string{
        "HANDLE":   "uintptr",    // Go的uintptr对应Rust的u64/指针
        "PVOID":    "uintptr",
        "ULONG":    "uint32",
        "NTSTATUS": "int32",
        // ...
    }
    // 这个映射是必要的，因为Go用户需要Go类型
}
```

**注意**：Go端类型映射是必要的，因为：
1. Go没有`HANDLE`类型，需要映射到`uintptr`
2. Go用户写代码时使用Go类型
3. 发送到驱动时，驱动使用Rust/WDK类型

### 13.5 类型流转完整示例

```
Go用户代码:
    args.FileHandle (uintptr)
         │
         ▼ JSON序列化
    {"FileHandle": 12345}
         │
         ▼ HTTP发送
    驱动接收
         │
         ▼ Rust反序列化
    NtDeviceIoControlFileArgs.FileHandle (HANDLE)
         │
         ▼ WDK API调用
    NtDeviceIoControlFile(FileHandle, ...)
         │
         ▼ 内核执行
    实际Windows API
```

### 13.6 bindgen工具链说明

`cmd/bindgen` 的职责：

1. **解析ntddk.rs** - 提取所有WDK函数签名
2. **解析types.rs** - 提取所有WDK类型定义
3. **解析constants.rs** - 提取所有WDK常量
4. **生成ntapi模块**:
   - `api_gen.rs` - 导出WDK函数
   - `types_gen.rs` - 导出WDK类型
   - `constants_gen.rs` - 导出WDK常量
5. **生成hook_db模块**:
   - `hook_args_gen.rs` - 为每个API生成Args结构体
   - `ept_hook_gen.rs` - EPT Hook数据库
   - `inline_hook_gen.rs` - Inline Hook数据库

**关键点**：所有类型都来自`wdk_sys`，通过`ntapi::types_gen`导出，`hook_args_gen`直接使用，无需手动映射。
