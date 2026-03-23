# BSOD Analysis Tools

纯Go实现的Windows蓝屏转储文件分析工具，无需依赖WinDbg。

## 功能特性

- **内核转储解析** - 解析PAGEDU64格式的内核转储文件
- **BugCheck识别** - 自动识别BSOD类型和错误代码
- **符号解析** - 结合PDB文件定位崩溃源代码位置
- **修复建议** - 根据错误类型提供修复建议

## 目录结构

```
walker/
├── kerneldump/          # 内核转储文件解析器
│   ├── kerneldump.go    # 核心解析逻辑
│   └── test/           # 测试程序
├── bsodanalyzer/        # BSOD分析器
│   ├── analyzer.go      # 分析器实现
│   └── test/           # 测试程序
└── minidump/           # 用户态minidump解析器（开发中）
```

## 使用方法

### 1. 解析内核转储文件

```go
package main

import (
    "fmt"
    "log"
    "github.com/ddkwork/HyperDbg/walker/kerneldump"
)

func main() {
    kd, err := kerneldump.Parse("C:\\Windows\\Minidump\\032326-3984-01.dmp")
    if err != nil {
        log.Fatal(err)
    }
    defer kd.Close()

    fmt.Println(kd)
    fmt.Printf("BugCheck: 0x%08X (%s)\n", kd.Header.BugCheckCode, kd.GetBugCheckName())
}
```

### 2. 完整BSOD分析（含PDB符号）

```go
package main

import (
    "fmt"
    "log"
    "github.com/ddkwork/HyperDbg/walker/bsodanalyzer"
)

func main() {
    analyzer, err := bsodanalyzer.New("C:\\Windows\\Minidump\\032326-3984-01.dmp")
    if err != nil {
        log.Fatal(err)
    }
    defer analyzer.Close()

    if err := analyzer.LoadDriverPDB("C:\\Windows\\Minidump\\hyperkd.pdb"); err != nil {
        log.Printf("Warning: Failed to load PDB: %v", err)
    }

    result, err := analyzer.Analyze()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result)
}
```

### 3. 批量分析所有转储文件

```bash
cd walker/kerneldump/test
go run analyze_all.go
```

## 支持的BugCheck类型

| 代码 | 名称 | 说明 |
|------|------|------|
| 0x00000001 | APC_INDEX_MISMATCH | APC索引不匹配 |
| 0x00000008 | IRQL_NOT_LESS_OR_EQUAL | IRQL级别错误 |
| 0x0000000A | IRQL_NOT_LESS_OR_EQUAL | IRQL级别错误 |
| 0x0000001E | KMODE_EXCEPTION_NOT_HANDLED | 内核模式异常 |
| 0x0000003B | SYSTEM_SERVICE_EXCEPTION | 系统服务异常 |
| 0x000000D1 | DRIVER_IRQL_NOT_LESS_OR_EQUAL | 驱动IRQL错误 |
| 0x000000EF | CRITICAL_PROCESS_DIED | 关键进程死亡 |
| 0x000000EA | THREAD_STUCK_IN_DEVICE_DRIVER | 驱动线程卡死 |

## 常见BSOD原因分析

### IRQL_NOT_LESS_OR_EQUAL (0x00000008 / 0x0000000A)

**原因：**
- 在DISPATCH_LEVEL或更高级别访问分页内存
- 使用无效的内存地址
- 驱动程序未正确处理IRQL级别

**修复建议：**
- 检查IRQL级别后再访问内存
- 确保所有内存正确映射
- 使用适当的同步原语
- 避免在DISPATCH_LEVEL+访问分页内存

### CRITICAL_PROCESS_DIED (0x000000EF)

**原因：**
- 驱动卸载前未清理回调
- EPT hook未在卸载前移除
- 用户模式调试会话未正确Detach
- 内核空间内存损坏

**修复建议：**
- 驱动卸载前正确清理所有回调
- 卸载前移除EPT hook
- 正确清理内核对象
- 检查内存泄漏

## 内核转储文件格式

### Header结构

```go
type Header struct {
    Signature            [8]byte  // "PAGEDU64" 或 "DUMP64"
    ValidDump            uint16
    MajorVersion         uint16
    MinorVersion         uint16
    DirectoryTableBase   uint64
    PfnDatabase          uint64
    PsLoadedModuleList   uint64
    PsActiveProcessHead  uint64
    MachineImageType     uint16
    NumberProcessors     uint8
    BugCheckCode         uint32
    BugCheckParameter1   uint64
    BugCheckParameter2   uint64
    BugCheckParameter3   uint64
    BugCheckParameter4   uint64
    // ...
}
```

### BugCheck参数解析

对于IRQL_NOT_LESS_OR_EQUAL (0x00000008)：
- Parameter 1: 包含模块ID和异常代码
  - Bits 56-63: 模块ID
  - Bits 0-7: 异常代码
  - Bits 8-55: 崩溃地址
- Parameter 2: 被访问的内存地址
- Parameter 3: IRQL级别
- Parameter 4: 保留

## 性能

- 解析速度：~1-2秒/文件（1.5MB转储文件）
- 内存占用：~10-20MB
- 支持批量分析

## 依赖

- Go 1.21+
- Windows DIA SDK（用于PDB符号解析）

## 未来计划

- [ ] 支持完整的内核转储文件（不仅仅是头）
- [ ] 解析模块列表和线程信息
- [ ] 提取调用栈
- [ ] 支持用户态minidump文件
- [ ] 集成到HyperDbg UI

## 参考

- [Microsoft Crash Dump Format](https://docs.microsoft.com/en-us/windows/win32/debug/minidump-files)
- [Bug Check Codes](https://docs.microsoft.com/en-us/windows-hardware/drivers/debugger/bug-check-code-reference2)
