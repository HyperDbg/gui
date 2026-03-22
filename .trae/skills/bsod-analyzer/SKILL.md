---
name: "bsod-analyzer"
description: "Analyzes HyperDbg BSOD crashes. Invoke when BSOD occurs or user reports system crash during debugging session."
---

# BSOD 分析器

当 HyperDbg 调试导致 BSOD 时，按以下步骤分析：

## 分析流程

### 1. 查看日志文件
读取 `d:\笔记本\p\ux\examples\hypedbg\log.log`，找到崩溃前的最后操作。

### 2. 运行 BSOD 分析脚本
执行 `d:\笔记本\p\ux\examples\hypedbg\analyze_bsod.ps1`，获取：
- 异常代码
- 异常地址
- 崩溃函数
- Bugcheck 类型

### 3. 查看测试代码
读取 `d:\笔记本\p\ux\examples\hypedbg\notepad_test.go#L12-12`，了解测试流程。

### 4. 分析源代码
根据日志中的源代码位置（如 `debugger/interfaces.go#L816-816`），定位问题代码。

## 常见 BSOD 类型

| Bugcheck | 名称 | 常见原因 |
|----------|------|----------|
| 0xEF | CRITICAL_PROCESS_DIED | 关键进程死亡，可能是回调未清理 |
| 0xD1 | DRIVER_IRQL_NOT_LESS_OR_EQUAL | 驱动访问无效内存 |
| 0x1E | KMODE_EXCEPTION_NOT_HANDLED | 内核模式异常未处理 |
| 0x3B | SYSTEM_SERVICE_EXCEPTION | 系统服务异常 |

## 检查清单

- [ ] 驱动卸载前是否正确清理所有回调
- [ ] EPT hook 是否在卸载前移除
- [ ] 用户模式调试会话是否正确 Detach
- [ ] 结构体布局是否与 C 代码匹配
- [ ] IOCTL 参数是否正确

## 相关文件

- 日志: `d:\笔记本\p\ux\examples\hypedbg\log.log`
- 分析脚本: `d:\笔记本\p\ux\examples\hypedbg\analyze_bsod.ps1`
- 测试代码: `d:\笔记本\p\ux\examples\hypedbg\notepad_test.go`
- 驱动源码: `d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd`
