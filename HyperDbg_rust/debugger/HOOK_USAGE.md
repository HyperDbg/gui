# Hook Database 使用指南

## 概述

Hook数据库提供了对Windows NT API的hook能力，支持两种模式：
- **EPT Hook**: 使用Extended Page Table实现，隐蔽但较慢
- **Inline Hook**: 直接修改函数入口，快速但可被检测

## 架构

```
┌─────────────────┐                    ┌─────────────────┐
│    Go用户态      │                    │    驱动内核      │
└─────────────────┘                    └─────────────────┘
        │                                      │
        │  1. InstallHook (IOCTL)              │
        │ ─────────────────────────────────────>│
        │                                      │ 设置EPT Hook
        │                                      │
        │  2. 事件循环                          │
        │ <─────────────────────────────────────│
        │     (Hook触发，驱动暂停线程)           │
        │                                      │
        │  3. 读取参数                          │
        │ ─────────────────────────────────────>│
        │ <─────────────────────────────────────│
        │     (返回Args数据)                    │
        │                                      │
        │  4. 执行用户回调                      │
        │                                      │
        │  5. 返回响应                          │
        │ ─────────────────────────────────────>│
        │                                      │ 继续执行/修改参数
```

## 使用示例

### 1. Hook NtDeviceIoControlFile

```go
package main

import (
    "encoding/json"
    "fmt"
    "time"
    "github.com/ddkwork/HyperDbg_rust/debugger"
)

func main() {
    packet := debugger.NewPacket().(*debugger.Packet)

    // 安装Hook
    err := packet.InstallHook(&debugger.HookConfig{
        ApiName:  "NtDeviceIoControlFile",
        HookType: debugger.HookTypeEPT,
        Filter: &debugger.HookFilter{
            ProcessNames: []string{"target.exe"}, // 只hook目标进程
        },
        Handler: func(event *debugger.HookEvent) (*debugger.HookEventResponse, error) {
            // 解析参数
            var args debugger.NtDeviceIoControlFileArgs
            if err := json.Unmarshal(event.Args, &args); err != nil {
                return nil, err
            }

            // 过滤特定的IoControlCode
            if args.IoControlCode == 0x12345678 {
                fmt.Printf("[%s] PID=%d IoControlCode=0x%X\n",
                    event.ProcessName, event.ProcessId, args.IoControlCode)
                fmt.Printf("  InputBuffer: 0x%X (size=%d)\n",
                    args.InputBuffer, args.InputBufferLength)
                fmt.Printf("  OutputBuffer: 0x%X (size=%d)\n",
                    args.OutputBuffer, args.OutputBufferLength)

                // 可选：修改参数
                // args.InputBuffer = newBuffer
                // return &debugger.HookEventResponse{
                //     EventId:      event.EventId,
                //     Action:       "continue",
                //     ModifyArgs:   true,
                //     ModifiedArgs: args,
                // }, nil
            }

            return &debugger.HookEventResponse{
                EventId: event.EventId,
                Action:  "pass",
            }, nil
        },
    })
    if err != nil {
        fmt.Printf("Failed to install hook: %v\n", err)
        return
    }

    fmt.Println("Hook installed, waiting for events...")

    // 事件循环
    for {
        event := packet.WaitForHookEvent(5 * time.Second)
        if event == nil {
            continue
        }

        // 获取并执行handler
        handler := getHandler(event.ApiName)
        if handler == nil {
            packet.RespondHookEvent(&debugger.HookEventResponse{
                EventId: event.EventId,
                Action:  "pass",
            })
            continue
        }

        resp, err := handler(event)
        if err != nil {
            resp = &debugger.HookEventResponse{
                EventId: event.EventId,
                Action:  "pass",
            }
        }

        packet.RespondHookEvent(resp)
    }
}
```

### 2. 自动生成的Args结构体

每个API都会自动生成对应的Args结构体：

```go
// 自动生成
type NtDeviceIoControlFileArgs struct {
    FileHandle        uintptr `json:"FileHandle"`
    Event             uintptr `json:"Event"`
    ApcRoutine        uintptr `json:"ApcRoutine"`
    ApcContext        uintptr `json:"ApcContext"`
    IoStatusBlock     uintptr `json:"IoStatusBlock"`
    IoControlCode     uint32  `json:"IoControlCode"`
    InputBuffer       uintptr `json:"InputBuffer"`
    InputBufferLength uint32  `json:"InputBufferLength"`
    OutputBuffer      uintptr `json:"OutputBuffer"`
    OutputBufferLength uint32 `json:"OutputBufferLength"`
}

type NtCreateFileArgs struct {
    FileHandle           uintptr `json:"FileHandle"`
    DesiredAccess        uint32  `json:"DesiredAccess"`
    ObjectAttributes     uintptr `json:"ObjectAttributes"`
    IoStatusBlock        uintptr `json:"IoStatusBlock"`
    AllocationSize       uint64  `json:"AllocationSize"`
    FileAttributes       uint32  `json:"FileAttributes"`
    ShareAccess          uint32  `json:"ShareAccess"`
    CreateDisposition    uint32  `json:"CreateDisposition"`
    CreateOptions        uint32  `json:"CreateOptions"`
    EaBuffer             uintptr `json:"EaBuffer"`
    EaLength             uint32  `json:"EaLength"`
}
```

## 进程过滤

避免hook系统进程导致崩溃：

```go
Filter: &debugger.HookFilter{
    // 按进程名过滤
    ProcessNames: []string{"target.exe", "notepad.exe"},
    
    // 或按PID过滤
    ProcessIds: []uint32{1234, 5678},
    
    // 排除系统进程 (PID <= 4)
    ExcludeSystem: true,
}
```

## HookEventResponse 动作

| Action | 说明 |
|--------|------|
| `pass` | 继续执行，不做任何修改 |
| `continue` | 继续执行，可能修改参数 |
| `skip` | 跳过原始调用，直接返回指定值 |

## 类型映射

| Rust类型 | Go类型 |
|----------|--------|
| NTSTATUS | int32 |
| HANDLE | uintptr |
| PVOID | uintptr |
| ULONG | uint32 |
| PULONG | *uint32 |
| BOOLEAN | bool |
| LARGE_INTEGER | int64 |
| PCSTR/PCWSTR | uintptr |

## 性能考虑

1. **EPT Hook**: 每次触发需要VM Exit，性能开销较大，但隐蔽性好
2. **Inline Hook**: 直接修改函数入口，速度快但可被检测
3. **进程过滤**: 建议始终设置过滤器，避免hook系统进程

## 文件说明

- `hook_db_gen.go`: 自动生成的Hook数据库，包含所有NT API的参数信息
- `hook_demo_test.go`: Hook使用示例
- `types.go`: 包含MessageType定义
