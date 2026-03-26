# HyperDbg API 方法行为统计

## 设计思想

### 核心原则
1. **简洁性** - 不为每个方法创建请求结构体，直接使用 `map[string]interface{}` 构建 JSON
2. **可读性** - 每个方法内部清晰展示发送的 JSON 结构
3. **可生成性** - `format` 标签用于代码生成工具参考，驱动端和 Go 端可同步生成
4. **类型安全** - 使用泛型约束确保返回类型正确

### 类型定义
```go
type Empty = struct{}

type ResponseType interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64 |
    ~string | ~bool |
    []byte | *RegisterState | []ProcessInfo | []ThreadInfo | []ModuleInfo | []BreakpointInfo |
    Empty
}

type Response[T ResponseType] struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    T      `json:"data,omitempty"`
}

func SendReceive[T ResponseType](p *Packet, jsonData []byte) *Response[T]

func NewPacket() Debugger
```

### API 设计
```go
// 无返回数据的方法使用 Empty
func (p *Packet) Initialize() error {
    data := mylog.Check2(json.Marshal(map[string]string{"action": "initialize"}))
    resp := SendReceive[Empty](p, data)
    if resp == nil || !resp.Success {
        return fmt.Errorf("initialize failed: %s", resp.Message)
    }
    p.state.Store(int32(StateRunning))
    return nil
}

// 有返回数据的方法使用具体类型
func (p *Packet) ReadMemory(address uint64, size uint32) ([]byte, error) {
    data := mylog.Check2(json.Marshal(map[string]interface{}{
        "action":  "read_memory",
        "address": fmt.Sprintf("0x%x", address),
        "size":    size,
    }))
    resp := SendReceive[[]byte](p, data)
    if resp == nil || !resp.Success {
        return nil, fmt.Errorf("read memory failed: %s", resp.Message)
    }
    return resp.Data, nil
}

func (p *Packet) GetProcessList() ([]ProcessInfo, error) {
    data := mylog.Check2(json.Marshal(map[string]string{"action": "get_process_list"}))
    resp := SendReceive[[]ProcessInfo](p, data)
    if resp == nil || !resp.Success {
        return nil, fmt.Errorf("get process list failed: %s", resp.Message)
    }
    return resp.Data, nil
}
```

### format 标签规范（代码生成用）
| format 值 | 说明 | Go 格式化 | Rust 解析 |
|-----------|------|-----------|-----------|
| `hex` | 十六进制字符串 | `fmt.Sprintf("0x%x", v)` | parse with 0x prefix |
| `dec` | 十进制字符串 | `fmt.Sprintf("%d", v)` | parse as decimal |
| `int` | 直接整数 | 直接传值 | 直接解析 |

### 字段命名规范
| 字段名 | 类型 | format | 说明 |
|--------|------|--------|------|
| action | string | - | 操作类型 |
| process_id | uint32 | dec | 进程ID |
| address | uint64 | hex | 内存地址 |
| breakpoint_id | uint64 | hex | 断点ID |
| size | uint32 | int | 大小 |
| type | int | int | 类型 |
| data | []byte/struct | - | 数据 |

## 方法总览

| 方法 | Action | Go形参 | 发送JSON字段 | 返回数据 | 返回类型 | 状态变更 | 前置检查 |
|------|--------|--------|-------------|---------|---------|---------|---------|
| Initialize | initialize | 无 | action | - | Empty | →Running | - |
| Terminate | terminate | 无 | action | - | Empty | →Terminated | - |
| AttachProcess | attach_process | processID uint32 | action, process_id | - | Empty | - | - |
| DetachProcess | detach_process | 无 | action | - | Empty | - | - |
| SetBreakpoint | set_breakpoint | address uint64, bpType BreakpointType | action, address, type | - | Empty | - | - |
| RemoveBreakpoint | remove_breakpoint | breakpointID uint64 | action, breakpoint_id | - | Empty | - | - |
| Continue | continue | 无 | action | - | Empty | →Running | StatePaused |
| Pause | pause | 无 | action | - | Empty | →Paused | StateRunning |
| StepInto | step_into | 无 | action | - | Empty | →Stepping | StatePaused |
| StepOver | step_over | 无 | action | - | Empty | →Stepping | StatePaused |
| StepOut | step_out | 无 | action | - | Empty | →Stepping | StatePaused |
| ReadMemory | read_memory | address uint64, size uint32 | action, address, size | []byte | []byte | - | - |
| WriteMemory | write_memory | address uint64, data []byte | action, address, data | - | Empty | - | - |
| ReadRegisters | read_registers | 无 | action | *RegisterState | *RegisterState | - | - |
| WriteRegisters | write_registers | regs *RegisterState | action, data | - | Empty | - | - |
| GetProcessList | get_process_list | 无 | action | []ProcessInfo | []ProcessInfo | - | - |
| GetThreadList | get_thread_list | processID uint32 | action, process_id | []ThreadInfo | []ThreadInfo | - | - |
| GetModuleList | get_module_list | processID uint32 | action, process_id | []ModuleInfo | []ModuleInfo | - | - |

## Go形参分类

### 无参数 (10个)
| 方法 | Action | 说明 |
|------|--------|------|
| Initialize | initialize | 初始化调试器 |
| Terminate | terminate | 终止调试器 |
| DetachProcess | detach_process | 分离进程 |
| Continue | continue | 继续执行 |
| Pause | pause | 暂停执行 |
| StepInto | step_into | 单步进入 |
| StepOver | step_over | 单步跳过 |
| StepOut | step_out | 单步跳出 |
| ReadRegisters | read_registers | 读取寄存器 |
| GetProcessList | get_process_list | 获取进程列表 |

### 单参数普通类型 (4个)
| 方法 | 参数类型 | 参数名 | JSON字段 | format |
|------|---------|--------|---------|--------|
| AttachProcess | uint32 | processID | process_id | dec |
| RemoveBreakpoint | uint64 | breakpointID | breakpoint_id | hex |
| GetThreadList | uint32 | processID | process_id | dec |
| GetModuleList | uint32 | processID | process_id | dec |

### 双参数普通类型 (2个)
| 方法 | 参数1 | 参数2 | JSON字段 | format |
|------|-------|-------|---------|--------|
| SetBreakpoint | address uint64 | bpType BreakpointType | address, type | hex, int |
| ReadMemory | address uint64 | size uint32 | address, size | hex, int |

### 普通类型+切片 (1个)
| 方法 | 参数1 | 参数2 | JSON字段 | format |
|------|-------|-------|---------|--------|
| WriteMemory | address uint64 | data []byte | address, data | hex, - |

### 单参数结构体指针 (1个)
| 方法 | 参数类型 | 参数名 | JSON字段 |
|------|---------|--------|---------|
| WriteRegisters | *RegisterState | regs | data |

## 发送数据分类

| 类型 | 字段 | 格式 | 使用方法 |
|------|------|------|---------|
| 无数据 | - | - | Initialize, Terminate, DetachProcess, Continue, Pause, Step*, ReadRegisters, GetProcessList |
| ProcessID | process_id | dec | AttachProcess, GetThreadList, GetModuleList |
| Address | address | hex | SetBreakpoint, RemoveBreakpoint, ReadMemory, WriteMemory |
| Address+Size | address, size | hex, int | SetBreakpoint, ReadMemory |
| MemoryData | data | []byte | WriteMemory |
| Registers | data | *RegisterState | WriteRegisters |

## 返回数据分类

| 类型 | 使用方法 |
|------|---------|
| Empty (仅success/message) | Initialize, Terminate, AttachProcess, DetachProcess, SetBreakpoint, RemoveBreakpoint, Continue, Pause, StepInto, StepOver, StepOut, WriteMemory, WriteRegisters |
| []byte | ReadMemory |
| *RegisterState | ReadRegisters |
| []ProcessInfo | GetProcessList |
| []ThreadInfo | GetThreadList |
| []ModuleInfo | GetModuleList |

## JSON 请求示例

### 无数据请求
```json
{"action": "initialize"}
{"action": "terminate"}
{"action": "detach_process"}
{"action": "continue"}
{"action": "pause"}
{"action": "step_into"}
{"action": "step_over"}
{"action": "step_out"}
{"action": "read_registers"}
{"action": "get_process_list"}
```

### 单参数请求
```json
{"action": "attach_process", "process_id": "1234"}
{"action": "remove_breakpoint", "breakpoint_id": "0x7ffe12345678"}
{"action": "get_thread_list", "process_id": "1234"}
{"action": "get_module_list", "process_id": "1234"}
```

### 双参数请求
```json
{"action": "set_breakpoint", "address": "0x7ffe12345678", "type": 1}
{"action": "read_memory", "address": "0x7ffe12345678", "size": 64}
```

### 普通类型+切片请求
```json
{"action": "write_memory", "address": "0x7ffe12345678", "data": [144, 144, 144]}
```

### 结构体请求
```json
{
  "action": "write_registers",
  "data": {
    "RAX": 1,
    "RBX": 2,
    "RCX": 3,
    "RDX": 4,
    "RSI": 5,
    "RDI": 6,
    "RBP": 7,
    "RSP": 8,
    "R8": 9,
    "R9": 10,
    "R10": 11,
    "R11": 12,
    "R12": 13,
    "R13": 14,
    "R14": 15,
    "R15": 16,
    "RIP": 17,
    "RFLAGS": 18
  }
}
```

## JSON 响应示例

### 无数据响应 (Empty)
```json
{"success": true, "message": "OK"}
{"success": false, "message": "error description"}
```

### []byte 响应
```json
{"success": true, "message": "OK", "data": [144, 144, 144]}
```

### *RegisterState 响应
```json
{
  "success": true,
  "message": "OK",
  "data": {
    "RAX": 1,
    "RBX": 2,
    "RCX": 3,
    "RDX": 4,
    "RSI": 5,
    "RDI": 6,
    "RBP": 7,
    "RSP": 8,
    "R8": 9,
    "R9": 10,
    "R10": 11,
    "R11": 12,
    "R12": 13,
    "R13": 14,
    "R14": 15,
    "R15": 16,
    "RIP": 17,
    "RFLAGS": 18
  }
}
```

### []ProcessInfo 响应
```json
{
  "success": true,
  "message": "OK",
  "data": [
    {"process_id": 1, "name": "System", "base_address": "0x0", "size": 0},
    {"process_id": 4, "name": "System", "base_address": "0x0", "size": 0}
  ]
}
```

### []ThreadInfo 响应
```json
{
  "success": true,
  "message": "OK",
  "data": [
    {"thread_id": 1, "process_id": 1, "start_address": "0x0", "priority": 0}
  ]
}
```

### []ModuleInfo 响应
```json
{
  "success": true,
  "message": "OK",
  "data": [
    {"name": "ntoskrnl.exe", "base_address": "0x0", "size": 0, "path": "C:\\Windows\\System32\\ntoskrnl.exe"}
  ]
}
```

## 统计汇总

| 分类 | 数量 | 占比 |
|------|------|------|
| 无参数 | 10 | 55.6% |
| 单参数普通类型 | 4 | 22.2% |
| 双参数普通类型 | 2 | 11.1% |
| 普通类型+切片 | 1 | 5.6% |
| 单参数结构体 | 1 | 5.6% |
| **总计** | **18** | **100%** |

## 返回类型统计

| 返回类型 | 数量 | 占比 |
|---------|------|------|
| Empty | 13 | 72.2% |
| []byte | 1 | 5.6% |
| *RegisterState | 1 | 5.6% |
| []ProcessInfo | 1 | 5.6% |
| []ThreadInfo | 1 | 5.6% |
| []ModuleInfo | 1 | 5.6% |
| **总计** | **18** | **100%** |
