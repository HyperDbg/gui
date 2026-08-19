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
| 项目根（CMake 顶层目录） | `d:\ux\examples\ewdk\debuger\SimpleHv\` |
| .sys 输出 | `<项目根>\Debug\SimpleHv.sys` |
| .pdb 符号  | `<项目根>\Debug\SimpleHv.pdb` |

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

## 输出格式

分析完成后输出：
1. **错误码** + **异常码**
2. **完整调用栈**
3. **根因定位**（函数名 + 源码行号）
4. **修复内容**（改了哪个文件的哪行）
