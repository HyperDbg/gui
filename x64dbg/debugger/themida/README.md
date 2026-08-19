# Themida/WinLicense 脱壳器

这是一个基于 HyperDbg 调试器的 Themida 和 WinLicense 脱壳器，将 ODBGScript 脚本翻译成了 Go 语言实现。

## 功能特性

- ✅ **VM 版本检测**：自动检测旧版、新版和 RISC VM
- ✅ **API 跟踪**：记录所有 API 调用
- ✅ **SetEvent 查找**：自动查找 SetEvent 入口点
- ✅ **I/O 标记查找**：查找 VM I/O 标记地址
- ✅ **PE 文件重建**：转储进程内存并重建 PE 文件
- ✅ **XBundler 支持**：自动处理 XBundler 加密的 DLL
- ✅ **HWID 绕过**：支持 HWID 检测和绕过
- ✅ **IAT 修复**：自动修复导入表

## 文件结构

```
themida/
├── unpacker.go          # 主脱壳器实现
├── api_logger.go       # API 调用日志记录器
├── pe_rebuilder.go     # PE 文件重建器
├── unpacker_test.go    # 单元测试
├── example.go          # 使用示例
└── README.md          # 本文件
```

## 快速开始

### 基本使用

```go
package main

import (
    "github.com/ddkwork/x64dbg/debugger"
    "github.com/ddkwork/x64dbg/debugger/themida"
)

func main() {
    // 创建调试器
    dbg := debugger.NewDebugger()
    if err := dbg.Attach("protected.exe"); err != nil {
        log.Fatal(err)
    }
    defer dbg.Detach()

    // 配置脱壳器
    config := themida.Config{
        TryIATPatch:       true,
        XBundlerAuto:      true,
        ARImpRecPath:       "api_log.txt",
    }

    // 创建并启动脱壳器
    unpacker := themida.NewThemidaUnpacker(dbg, config)
    if err := unpacker.Start(); err != nil {
        log.Fatal(err)
    }
    defer unpacker.Stop()

    // 获取脱壳信息
    fmt.Printf("VM 版本: %v\n", unpacker.GetVMVersion())
    fmt.Printf("SetEvent 地址: 0x%X\n", unpacker.GetSetEventAddress())
    fmt.Printf("I/O 标记地址: 0x%X\n", unpacker.GetIOMarkerAddress())

    // 转储进程
    unpacker.DumpProcess("unpacked.exe")
}
```

### 命令行使用

```bash
# 编译示例程序
go build -o themida_unpacker example.go

# 运行脱壳器
./themida_unpacker C:\\path\\to\\protected.exe
```

## 配置选项

### Config 结构

```go
type Config struct {
    SetEventUserData    bool    // 是否使用 SetEvent 用户数据
    CheckHWID          bool    // 是否检查 HWID
    BypassHWIDSimple   bool    // 是否使用简单 HWID 绕过
    TryIATPatch       bool    // 是否尝试 IAT 修复
    AllocSize          uint32  // RISC VM 分配大小（默认：0x200000）
    AllocSizePEADS     uint32  // PE_ADS 分配大小（默认：0x30000）
    XBundlerAuto      bool    // 是否自动处理 XBundler
    UseMessageHWBP     bool    // 是否使用消息硬件断点
    ARImpRecPath       string  // API 日志文件路径
    SetEventEntryAddr  uint32  // SetEvent 入口地址（0 表示自动查找）
    IOMarkerAddress    uint32  // I/O 标记地址（0 表示自动查找）
    SecLocation       uint32  // 段位置
}
```

## VM 版本

脱壳器支持三种 VM 版本：

### 1. 旧版 VM (VMOld)

- 签名：`68 ???????? E9 ??????? FF 68 ???????? E9 ??????? FF`
- 特征：单 PUSH + 单 JUMP 模式
- 处理方法：查找旧版 VM 入口点

### 2. 新版 VM (VMNew)

- 签名：`68 ???????? 68 ???????? E9 ??????? FF 68 ???????? 68 ???????? E9 ??????? FF`
- 特征：双 PUSH + 单 JUMP 模式
- 处理方法：查找新版 VM 入口点

### 3. RISC VM (VMRISC)

- 签名：无固定签名
- 特征：使用 RISC 指令集
- 处理方法：单步执行查找 I/O 标记

## API 日志

脱壳器会记录所有 API 调用到日志文件：

```
=== Themida/WinLicense Unpacker Log ===
Started: 2024-01-01 12:00:00
=======================================================

[2024-01-01 12:00:01.000] Call from: 0x00401000 | API: 0x7FFE5FD87FB0 | NAME: SetEvent
[2024-01-01 12:00:01.100] Call from: 0x00401005 | API: 0x7FFE5FD87650 | NAME: VirtualAlloc
[2024-01-01 12:00:01.200] Call from: 0x0040100A | API: 0x7FFE5FD893C0 | NAME: GetProcAddress

-------------------------------------------------------
--------------- SETEVENT_ENTRY_ADDRESS ----------------
Address: 0x00401000 | PUSH: 0x12345678 | JUMP: 0x00402000
-------------------------------------------------------

-------------------------------------------------------
--------------- I_O_MARKER_ADDRESS --------------------
Address: 0x00403000 | Section Location: 0x00400000 | I_O_MARKER_ADDRESS RVA: 0x00003000
-------------------------------------------------------
```

## 测试

运行单元测试：

```bash
# 运行所有测试
go test ./themida/...

# 运行特定测试
go test ./themida/... -run TestThemidaUnpacker_APILogger

# 运行集成测试
go test ./themida/... -v
```

## 高级用法

### 自定义 VM 入口点

如果你已经知道 SetEvent 入口点，可以手动指定：

```go
config.SetEventEntryAddr = 0x00401000
```

### HWID 绕过

使用简单 HWID 绕过方法：

```go
config.BypassHWIDSimple = true
config.CheckHWID = true
```

### IAT 修复

自动修复导入表：

```go
config.TryIATPatch = true
```

### XBundler 处理

自动处理 XBundler 加密的 DLL：

```go
config.XBundlerAuto = true
config.AllocSizePEADS = 0x30000
```

## 故障排除

### 问题：无法检测到 VM 版本

**解决方案**：
1. 确保目标程序确实使用 Themida/WinLicense 加密
2. 检查段名称是否包含 `.themida`、`.winlice`、`.wlsec` 等
3. 增加读取的内存大小

### 问题：SetEvent 断点未触发

**解决方案**：
1. 检查是否正确获取了 SetEvent 地址
2. 确保程序已经运行到调用 SetEvent 的位置
3. 尝试使用硬件断点代替软件断点

### 问题：I/O 标记未找到

**解决方案**：
1. 确认 VM 版本检测正确
2. 增加搜索范围
3. 手动指定 I/O 标记地址

### 问题：转储的文件无法运行

**解决方案**：
1. 检查 IAT 是否正确修复
2. 确认 PE 文件头是否正确重建
3. 使用外部工具（如 PE Explorer）验证 PE 结构

## 注意事项

1. **法律警告**：仅用于学习和研究目的，不要用于非法用途
2. **备份**：在脱壳前备份原始文件
3. **测试环境**：建议在虚拟机中进行测试
4. **兼容性**：支持 Windows XP SP2/SP3 和 Windows 7 32 位
5. **VMWare**：如果使用 VMWare，需要修改 .vmx 文件以绕过检测

## 参考资料

- 原始 ODBGScript：`Themida - Winlicense Ultra Unpacker 1.4 杨开银.txt`
- HyperDbg 调试器：https://github.com/ddkwork/x64dbg
- Themida 官方网站：https://www.oreans.com/

## 许可证

本项目仅供学习和研究使用。

## 作者

基于 LCF-AT 的 ODBGScript 脚本翻译而来。

## 更新日志

### v1.0.0 (2024-01-01)

- ✅ 初始版本发布
- ✅ 支持 VM 版本检测
- ✅ 支持 API 跟踪
- ✅ 支持 SetEvent 查找
- ✅ 支持 I/O 标记查找
- ✅ 支持 PE 文件重建
- ✅ 支持 XBundler 处理
- ✅ 支持 HWID 绕过
- ✅ 支持 IAT 修复