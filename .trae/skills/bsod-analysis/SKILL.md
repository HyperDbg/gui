---
name: "bsod-analysis"
description: "Analyzes Windows kernel crash dumps (.dmp) with kd.exe and PDB symbols, identifies root cause, applies minimal source fix. Invoke when user provides .sys, .pdb, .dmp files for driver crash analysis."
---

# BSOD 驱动崩溃分析技能

## 固定环境配置

| 项目 | 路径 |
|------|------|
| kd.exe | `C:\Program Files\WindowsApps\Microsoft.WinDbg_1.2606.22001.0_x64__8wekyb3d8bbwe\amd64\kd.exe` |
| dumpbin.exe | `E:\Program Files\Microsoft Visual Studio\18\BuildTools\VC\Tools\MSVC\14.50.35717\bin\Hostx64\x64\dumpbin.exe` |
| 系统符号路径 | `C:\ProgramData\Dbg\sym` |
| 符号服务器 | `srv*`（自动 fallback） |
| DMP 文件 | `C:\Windows\MEMORY.DMP`（每次 BSOD 覆盖，路径不变） |

## 当前项目快捷配置

对于当前正在开发的项目，默认 PDB 和源代码路径已在下方列出。
如果是其他项目，用户会额外说明路径。

| 项目 | 路径 |
|------|------|
| 项目根（CMake 顶层目录） | `d:\ux\HyperDbgUnified\` |
| 源码根 | `d:\ux\HyperDbgUnified\HyperDbg\hyperdbg\` |
| .sys 输出 | `D:\ux\HyperDbgUnified\Debug\hyperkd.sys`（hyperhv 静态链入 hyperkd，栈帧符号统一在 hyperkd） |
| .pdb 符号  | `D:\ux\HyperDbgUnified\Debug\hyperkd.pdb` |

注意：本项目的崩溃栈里 `hyperkd!Vmx*` 系列函数实际来自 hyperhv 源码（静态链接合并），
定位源码行时用 PDB 记录的文件路径，源文件在 `hyperhv\code\` 下。

## 标准分析命令

分析时三要素：**符号路径（PDB）** + **源代码路径** + **dump 文件**。

```powershell
$env:_NT_SOURCE_PATH = "<项目根>"
& "C:\Program Files\WindowsApps\Microsoft.WinDbg_1.2606.22001.0_x64__8wekyb3d8bbwe\amd64\kd.exe" -z "C:\Windows\MEMORY.DMP" -y "C:\ProgramData\Dbg\sym;<项目根>\Debug" -c ".lines;!analyze -v;q"
```

`_NT_SOURCE_PATH` 环境变量让 kd 能根据 PDB 中记录的相对路径找到 `.cpp`/`.h` 源代码文件，
这样 `!analyze -v` 就能直接输出 `FAULTING_SOURCE_LINE` 对应的代码上下文，
`uf` 反汇编时也能 inline 显示源码行。

## 分析步骤

### 1. 加载 dump 获取初步信息

```powershell
$env:_NT_SOURCE_PATH = "<项目根>"
& "C:\Program Files\WindowsApps\Microsoft.WinDbg_1.2606.22001.0_x64__8wekyb3d8bbwe\amd64\kd.exe" -z "C:\Windows\MEMORY.DMP" -y "C:\ProgramData\Dbg\sym;<项目根>\Debug" -c ".lines;!analyze -v;q"
```

### 2. 提取关键信息

从 `!analyze -v` 输出提取：
- **BUGCHECK_CODE** — 如 `0x7E`, `0xD1`, `0x3B`, `0x101`
- **IMAGE_NAME** — 哪个模块崩溃
- **STACK_TEXT** — 完整调用栈
- **FAULTING_SOURCE_LINE** — 源码位置（需 PDB + 源码路径配合）
- **FAILURE_BUCKET_ID** — 微软错误桶

### 3. 深度分析

根据栈回溯，用反汇编定位精确崩溃点：

```powershell
# 反汇编关键函数（_NT_SOURCE_PATH 已设时自动显示源码行）
uf <module>!<function>

# 查看完整栈
knL

# 查看特定栈帧的源码上下文
.lines; .srcpath
```

如果 PDB 路径正确但无法加载镜像，用 `.reload /f <module>` 强制加载符号。

### 4. 修复原则

- **最小修复**：只删除/修改直接导致崩溃的代码行，不删日志、不重构、不改其他文件
- 只改源码，不改构建配置或项目文件

## 通用 BSOD 模式对照

| BUGCHECK_CODE | 常见根因 |
|---------------|----------|
| 0x7E | 异常未处理，如 assert 失败 → int 3、空指针、除零 |
| 0xD1 | IRQL 违规，驱动在错误 IRQL 访问分页内存 |
| 0x3B | APC_LEVEL 下 PagedPool 分配 |
| 0x19 | Bad Pool Header，tag 不匹配、double free |
| 0x101 | CLOCK_WATCHDOG_TIMEOUT，中断被禁用导致 IPI 无响应 |
| 0x1AA | EXCEPTION_ON_INVALID_STACK：VMX 根模式（VM-exit 处理器、VmmStack）上调用了阻塞型 Windows API（如 `KeGenericCallDpc`），调度器在 VmmStack 上切换线程污染 VM-exit 上下文，vmresume 触发 #UD。特征：故障 CONTEXT 的 RSP 落在 VmmStack、EFlags.IF=1、Dr0/Dr7=0 |

## PDB 与 .sys 匹配验证（dumpbin）

PDB 找不到/不匹配时，先用 dumpbin 看 .sys 的 debug 目录记录的 PDB GUID+Age 和路径：

```powershell
& "E:\Program Files\Microsoft Visual Studio\18\BuildTools\VC\Tools\MSVC\14.50.35717\bin\Hostx64\x64\dumpbin.exe" /headers "<sys路径>" | Select-String -Pattern "pdb|Debug" -Context 0,2
```

输出中的 `Format: RSDS, {GUID}, Age, 文件名` 就是该 .sys 编译时期望的 PDB；
对比 `.pdb` 文件名与修改时间是否是同一次编译产物（同时刻生成即匹配）。

## VMX 项目专项检查清单

崩溃栈涉及 `hyperkd!VmxPerformVmresume` / `vmresume` #UD 时，重点核查：

1. 故障 CONTEXT 的 RSP 是否落在某核 VmmStack（`dt hyperkd!_VIRTUAL_MACHINE_STATE poi(hyperkd!g_GuestState)+i*sizeof(...)` 看 VmmStack 字段）
2. `ExitReason` / `ExitQualification` / `LastVmexitRip`（崩溃核最后一次 VM-exit 的原因与位置）
3. `g_ProcessDebuggingDetailsListHead` 走链找 `USERMODE_DEBUGGING_THREAD_DETAILS.UdAction[]`（VMX 根模式里未消费完的用户调试命令）
4. 代码路径上是否存在任何"会阻塞/调度"的 Windows API（KeGenericCallDpc / KeWaitForSingleObject / ExAllocatePool with PoolQuota 等）被 VMX 根模式调用

## 输出格式

分析完成后输出：
1. **错误码** + **异常码**
2. **完整调用栈**
3. **根因定位**（函数名 + 源码行号）
4. **修复内容**（改了哪个文件的哪行）
