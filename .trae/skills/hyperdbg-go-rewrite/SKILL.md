---
name: "hyperdbg-go-rewrite"
description: "用 Go 完整生产级复刻 HyperDbg 用户层调试体系（libhyperdbg+CLI+script-engine+symbol-parser），驱动内 C 实现 Go 子集解释器执行 hook 回调。后期扩展 Go GUI 和 MCP。Invoke when 用户要求实现/修改 Go 版 libhyperdbg、驱动 Go 子集解释器、AST 二进制协议、yaegi 集成、GUI/MCP 接口预留，或涉及 ok 目录下 Go 绑定包（ddk/pe-main/pdbex/zydis/keystone 等）的使用时。"
---

# HyperDbg Go 重写方案

## 项目定位（重要）

**这是生产级完整实现，不是 demo/prototype/测试**。

- **目标**：用 Go 完整复刻 HyperDbg 用户层调试体系，与 C/C++ 版长期共存、行为一致、可对照验证
- **代码质量要求**：生产级——完整错误处理、并发安全（goroutine + sync）、context 取消传播、结构化日志、可测试性
- **API 设计原则**：所有公开 API 必须可被 **GUI 层** 和 **MCP 层** 复用，不能绑死在 CLI/REPL 上
- **后期扩展**（本次重构必须预留接口）：
  - **Go 版 GUI**：基于 `go-libhyperdbg/api/` 的公开 API，Wails/Fyne/Web 后端任选，本次不实现但 API 不能挡路
  - **MCP（Model Context Protocol）层**：把调试器能力暴露给 AI Agent，本次不实现但 API 必须可程序化调用（无 REPL 依赖、无全局状态、可并发）

**反模式（禁止）**：
- 把命令逻辑写死在 CLI 入口里（必须放 `go-libhyperdbg/` 包内，CLI 只调用）
- 用全局变量传递调试器状态（必须用结构体 + 方法，支持多实例）
- 命令输出直接 `fmt.Println`（必须走 output channel / callback，GUI 和 MCP 才能接管输出）
- 同步阻塞 API 不带 context（GUI/MCP 需要取消能力）

## 架构总览

```
┌──────────────────────────────────────────────────┐
│  消费层（本次只做 CLI，但 API 对 GUI/MCP 友好）        │
│  ├─ go-cli（REPL + 脚本模式，yaegi）                │
│  ├─ go-gui（后期，本次预留 API）                     │
│  └─ go-mcp（后期，本次预留 API）                     │
└──────────────────┬───────────────────────────────┘
                   │ 调用 go-libhyperdbg 公开 API
                   ▼
┌──────────────────────────────────────────────────┐
│  go-libhyperdbg（用户态核心库，生产级）              │
│  ├─ api/  顶层 API（无全局状态，支持多实例）           │
│  ├─ debugger/  命令实现（97 个命令）                 │
│  ├─ comm/  通信层（Named Pipe / IOCTL）             │
│  ├─ events/  事件流（hook 回调上送）                  │
│  └─ symbolparser/  符号解析                         │
└──────────────────┬───────────────────────────────┘
                   │ ① IOCTL: InstallHook(sym, astBytes)
                   │ ② 事件流（驱动→用户态，HookEvent 包）
                   │ ③ 控制指令（Continue/Break/ReadMem）
                   ▼
┌──────────────────────────────────────────────────┐
│  hyperhv + hyperkd (C，核心不动)                    │
│  ├─ EPT hook 命中 → 反序列化 Go AST → C 解释器执行    │
│  ├─ 解释器访问 ctx（寄存器快照、栈、内存）            │
│  ├─ 执行结果 → Continue/Break/ModifyReg/Log         │
│  └─ script-eval + script-engine 保留作对照（不删）    │
└──────────────────────────────────────────────────┘
```

**两套 Go 解释器**：
- 用户态：yaegi（完整 Go，跑顶层逻辑）
- 内核态：C 写的 Go **子集**解释器（跑 hook 回调）

**关键约束**：VMX-root 没有 Go runtime，内核解释器只能支持严格子集。

## Go 子集定义

### 支持

| 类别 | 语法 |
|---|---|
| 变量声明 | `var x uint64 = 0`、`x := 1` |
| 类型 | `uint8/16/32/64`、`int8/16/32/64`、`bool`、`string`（只读字面量） |
| 算术 | `+ - * / %` |
| 位运算 | `& \| ^ << >> &^` |
| 比较 | `== != < > <= >=` |
| 逻辑 | `&& \|\| !` |
| 控制流 | `if/else`、`for`（含三段式）、`break/continue/return` |
| 函数声明 | `func name(args) (rets) { ... }` |
| 简单闭包 | `func(ctx *HookCtx) { ... }`（**只捕获 ctx，不捕获外部变量**） |
| 白名单函数调用 | 见下表 |
| 固定大小数组 | `[N]uint8` |

### 不支持（VMX-root 限制）

- `goroutine` / `channel` / `select` —— 无调度器
- `interface` / type assertion —— 无类型表
- `map` —— 无 hash + 动态扩容
- `slice`（除了 `[N]uint8` 固定数组）—— 无堆分配
- `reflect` / `unsafe` —— 无运行时类型信息
- `defer` —— 无栈展开
- `string` 拼接（`+`）—— 要分配堆，用 `Printf` 替代
- `make` / `new` —— 无堆
- 闭包捕获外部变量（除 `ctx` 外）—— GC 生命周期无法管理

### 白名单函数（驱动导出，HookCtx 方法）

| 方法 | 语义 |
|---|---|
| `ctx.StackReadQword(offset uint32) uint64` | 读栈（替代 `poi(@rsp+offset)`） |
| `ctx.StackReadDword(offset uint32) uint32` | 读栈 4 字节 |
| `ctx.Reg(name string) uint64` | 读寄存器（`"rax"`/`"rip"`/`"rsp"`...） |
| `ctx.SetReg(name string, val uint64)` | 写寄存器 |
| `ctx.ReadMem(addr uint64, buf []byte)` | 读客户机内存 |
| `ctx.ReadMemQword(addr uint64) uint64` | 读 8 字节 |
| `ctx.Printf(fmt string, args ...uint64)` | 日志（只支持 `%x` `%d` `%s`） |
| `ctx.Break()` | 暂停进程，上送用户态 |
| `ctx.Continue()` | 继续执行（默认行为） |
| `ctx.GetPid() uint32` | 当前进程 PID |
| `ctx.GetTid() uint32` | 当前线程 TID |
| `ctx.GetIP() uint64` | 当前指令地址 |

## AST 二进制协议（草案）

紧凑格式，每个节点 1 字节 opcode + 数据。

```
节点 opcode:
  0x01  Literal       → 1 byte kind(u8/u16/u32/u64/i8.../bool/string) + N bytes value
  0x02  Ident         → 2 bytes name_id（索引到字符串表）
  0x03  BinaryExpr    → 1 byte op + LHS + RHS
  0x04  UnaryExpr     → 1 byte op + Operand
  0x05  CallExpr      → 2 bytes func_id + 1 byte nargs + N children
  0x06  SelectorExpr  → X + 2 bytes field_id（用于 ctx.Method）
  0x07  AssignStmt    → 1 byte op(= /= += ...) + LHS + RHS
  0x08  IfStmt        → Cond + Then + Else（Else 可为 0x00 nil）
  0x09  ForStmt       → Init + Cond + Post + Body
  0x0A  BlockStmt     → 2 bytes nstmts + N statements
  0x0B  ReturnStmt    → 1 byte nvals + N children
  0x0C  FuncDecl      → 2 bytes name_id + 1 byte nargs + N args + 1 byte nrets + N rets + Body
  0x0D  FuncLit       → 1 byte nargs + N args + 1 byte nrets + N rets + Body
  0x0E  DeclStmt      → var 声明
  0x0F  ArrayType     → 4 bytes len + 1 byte elem_kind
  0x10  CompositeLit  → Type + 2 bytes nels + N elements
  0x11  IndexExpr     → X + Index
  0x00  Nil           → 空节点（Else 缺省时用）

字符串表: 末尾附 4 bytes count + N 个 (2 bytes len + bytes)

op 编码（BinaryExpr）:
  0x01 +   0x02 -   0x03 *   0x04 /   0x05 %
  0x06 &   0x07 |   0x08 ^   0x09 <<  0x0A >>
  0x0B ==  0x0C !=  0x0D <   0x0E >   0x0F <=  0x10 >=
  0x11 &&  0x12 ||  0x13 &^
```

整个 hook 回调 AST 序列化后通常 < 1KB。

## ok 目录 Go 绑定包映射

`d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\ok\` 下的资源：

### 直接可用（无需改造）

| 包 | 路径 | 用途 |
|---|---|---|
| `pe-main` | `ok/pe-main/` | PE 文件解析（imports/exports/sections/reloc/debug）—— 用于 `.start path xxx.exe` 解析入口点、模块枚举 |
| `gjson` | `ok/gjson/` | JSON 配置解析 |
| `byteslice` | `ok/byteslice/` | 字节切片工具 |
| `pdbfetch-master` | `ok/pdbfetch-master/` | PDB 下载（Microsoft Symbol Server）—— 替代 `.sym download` |
| `keystone` | `ok/keystone/` | 汇编器（含 keystone.dll）—— 替代 `a` 命令 |
| `zydis` | `ok/zydis/` | 反汇编器（含 zydis.dll）—— 替代 `u`/`d` 命令 |
| `xed` | `ok/xed/` | Intel XED 反汇编（备选，含 xed.dll） |

### 部分可用（需适配）

| 包 | 路径 | 用途 | 适配点 |
|---|---|---|---|
| `ddk` | `ok/ddk/` | 驱动加载管理 | `Driver.Install/Start/Stop/Delete` 直接用（基于 SCM API）；**`KernelMemory` 不可用**（那是 RTCore64 漏洞驱动方案），HyperDbg 的内核读写要走自己的 IOCTL/VMCALL |
| `pdbex` | `ok/pdbex/` | PDB 符号解析 | DIA 接口可用（`dia_windows.go`、`pdbex.go`）；含 Rust 版 symbolizer 参考实现 |

### 不可用（需自行实现）

- HyperDbg VMM 通信层（IOCTL/VMCALL）—— 走 hyperhv 的 named pipe / IOCTL，不是 RTCore64
- EPT hook 事件流 —— 复用 `libhyperdbg/code/debugger/communication/forwarding.cpp` 思路，Go 重写
- libipt（Intel PT 解码）—— syscall 包装，ok 目录无

## 落地顺序（Phase A → Phase B → Phase C）

**执行原则**：先地基（驱动加载）→ 再核心（Go 解释器）→ 最后全量。每个 Phase 有单元测试验收，通过后才进下一个 Phase。

### Phase A：驱动加载地基（最小可验证，无依赖）

**目标**：Go 能加载/卸载 VMM 驱动，IOCTL 通信正常。这是所有后续工作的地基。

| 步骤 | 任务 | 文件 | 依赖 |
|---|---|---|---|
| A.1 | 移动 HyperDbgSdk.go 到 go-libhyperdbg/types/ 并改包名 | `go-libhyperdbg/types/sdk.go`（从 `ok/HyperDbgSdk.go` 移动，包名改为 `types`） | 无 |
| A.2 | 驱动加载 | `go-libhyperdbg/debugger/driverloader/loader.go`（用 `ok/ddk.Driver`） | A.1 |
| A.3 | IOCTL 通信层 | `go-libhyperdbg/debugger/comm/ioctl.go`、`device.go`（用 `types` 包打包） | A.1 |
| A.4 | Phase A 单元测试 | `go-libhyperdbg/debugger/driverloader/loader_test.go`、`comm/comm_test.go` | A.2, A.3 |

**Phase A 验收**（单元测试）：
- `TestLoadUnloadVMM`：加载 hyperhv.sys → 发 `IOCTL_INIT_VMM` → 验证驱动响应 → 发 `IOCTL_TERMINATE_VMX` → 卸载 → 服务清理干净
- `TestIoctlRoundTrip`：发送简单 IOCTL，验证返回值正确
- 通过后才进 Phase B

### Phase B：Go 子集 + 驱动解释器

**目标**：Go AST 回调能在 VMX-root 里执行，输出正确，不死机。

| 步骤 | 任务 | 文件 | 依赖 |
|---|---|---|---|
| B.1 | Go 子集 + AST 二进制协议 spec | `docs/go-subset-spec.md` | 无（可与 A 并行） |
| B.2 | go-bridge/protocol | `go-bridge/protocol/opcodes.go`、`nodes.go` | B.1 |
| B.3 | go-bridge/subset | `go-bridge/subset/validator.go` | B.1 |
| B.4 | go-bridge/ast | `go-bridge/ast/encoder.go` | B.2, B.3 |
| B.5 | 驱动 ast_decode.c | `hyperkd/code/go-interp/ast_decode.c` + `.h` | B.2 |
| B.6 | 驱动 interp.c | `hyperkd/code/go-interp/interp.c` + `.h`（MVP：if/赋值/算术/Printf/StackReadQword，cycle budget，SEH） | B.5 |
| B.7 | 驱动 whitelist.c | `hyperkd/code/go-interp/whitelist.c` + `.h`（HookCtx 方法） | B.6 |
| B.8 | 驱动 interp_stub.c | `hyperkd/code/go-interp/interp_stub.c`（EPT hook 接入，`HOOK_FLAG_GO_AST` 区分，与 script-eval 并行） | B.6, B.7 |
| B.9 | CMakeLists 集成 | `CMakeLists.txt`（go-interp/*.c 加入 hyperkd） | B.5-B.8 |
| B.10 | Phase B 单元测试 | `go-test/phase_b_test.go` + 驱动测试 | B.4, B.8, A 全通过 |

**Phase B 验收**（单元测试 + 回归）：
- `TestGoAstHookCallback`：下发简单 Go AST（`if x < 0 { printf(...) }`）→ 注册 EPT hook → hook 命中 → 驱动解释执行 → 验证日志输出正确
- `TestCycleBudget`：`for {}` 死循环在 100 万 cycle 内安全返回，不 BSOD
- `TestTypeError`：类型错误不 BSOD，返回错误日志
- **回归测试**：Phase A 的 `TestLoadUnloadVMM` 仍通过
- 通过后才进 Phase C

### Phase C：全量复刻（执行所有剩余任务）

**目标**：完整复刻 HyperDbg 用户层调试体系，与 C/C++ 版对照一致。

#### C.1 用户态核心（让 find-oep.go 跑起来）
| 步骤 | 任务 | 文件 | 依赖 |
|---|---|---|---|
| C.1.1 | debugger/core | `go-libhyperdbg/debugger/core/debugger.go`、`interpreter.go`、`break_control.go`、`steppings.go` | A.3 |
| C.1.2 | api 顶层 | `go-libhyperdbg/api/debugger.go`、`options.go`、`output.go`、`event.go` | C.1.1, B.4 |
| C.1.3 | 跑通 find-oep.go | `go-cli/examples/find-oep.go` | C.1.2, B.8 |

**验收**：`find-oep.go` 执行后 `find-oep.log` 含 `RAH ret=xxx` 行，与 C 版对照一致。

#### C.2 命令全量复刻（97 个命令）
| 步骤 | 任务 | 命令数 |
|---|---|---|
| C.2.1 | `debugger/commands/meta/` | 23（attach/cls/connect/...） |
| C.2.2 | `debugger/commands/debugging/` | 37（a/bc/bd/...） |
| C.2.3 | `debugger/commands/extension/` | 35（apic/cpuid/...） |
| C.2.4 | `debugger/commands/hwdbg/` | 2（hw、hw_clk） |

**验收**：每个命令输出与 C/C++ 版逐字节一致（`go-test/parity/`）。

#### C.3 支撑模块全量复刻
| 步骤 | 任务 | 依赖 |
|---|---|---|
| C.3.1 | `debugger/misc/`（assembler 用 ok/keystone、disassembler 用 ok/zydis、callstack、readmem、pt-helper、pci-id） | ok/keystone, ok/zydis |
| C.3.2 | `debugger/userlevel/`（pe-parser 用 ok/pe-main、ud、user-listening） | ok/pe-main |
| C.3.3 | `debugger/kernellvl/`（kd、kernel-listening） | A.3 |
| C.3.4 | `debugger/transparency/`（gaussian-rng、transparency） | 无 |
| C.3.5 | `symbolparser/`（用 ok/pdbex + ok/pdbfetch-master，含 32/64 位 PDB 区分） | ok/pdbex, ok/pdbfetch |
| C.3.6 | `app/`、`common/`、`export/`、`objects/`、`rev/`、`hwdbg/`、`ucpuid/` | 各模块 |

#### C.4 CLI + 脚本引擎 + 测试
| 步骤 | 任务 | 依赖 |
|---|---|---|
| C.4.1 | `go-cli/`（main.go、repl.go、script.go） | C.1.2, C.2.1-C.2.4 |
| C.4.2 | `debugger/scriptengine/wrapper.go`（集成 go-bridge） | B.4 |
| C.4.3 | yaegi 集成（`go-libhyperdbg/interp/`） | C.4.1 |
| C.4.4 | `go-test/parity/` 对照测试 | C.2, C.3.5 |
| C.4.5 | syscall 包装 libipt（`debugger/misc/pt/`，`//go:embed` + `windows.NewLazyDLL`，无 cgo） | libipt.dll |

#### C.5 补全 Go 子集
| 步骤 | 任务 | 依赖 |
|---|---|---|
| C.5.1 | 补全 interp.c：`for` 三段式、`return`、`ReadMem`、`SetReg`、`Break`、数组、`FuncDecl` | B.6 |
| C.5.2 | cycle budget 严格验证 | C.5.1 |

**注意**：C/C++ 代码全部保留作对照组，不清理。原 `script-eval` + `script-engine` 保留，新 Go AST 路径并行。

## 关键技术约束

### VMX-root 限制

- 无分页（Page Fault 不能处理）
- IRQL HIGH_LEVEL，不能调任何 OS 服务
- 不能堆分配（只能预分配池或栈分配）
- 不能阻塞
- 其他 vCPU 在等，执行时间敏感（防 0x101 BSOD）

### Cycle Budget

每次解释执行上限 **1,000,000 cycles**，超时强制返回 Continue。
解释器每执行一个 AST 节点 `cycle++`，超限返回。
不做好这条，`for {}` 一定 BSOD。

### 闭包捕获限制

`FuncLit` 节点**只能引用 `ctx` 或函数参数**，不能引用外部变量。
用户态编译器在序列化时**静态检查**，违反则拒绝。
这是防止"用户写出 Go 子集外的代码"的第一道防线。

### 错误隔离

解释器跑在 `__try/__except`（SEH）里：
- 栈溢出 → catch，返回 Continue + log error
- 类型错误 → catch，返回 Continue + log error
- 未知节点 → catch，返回 Continue + log error
- 内存读取失败 → catch，返回 Continue + log error

**绝不能因为脚本错误 BSOD**。

### AST 缓存

同一个 hook 的 AST **反序列化一次缓存**，命中时直接跑缓存的节点树。
不要每次 hook 都反序列化（性能崩）。

### Printf 简化

驱动里实现简化版 `Printf`：
- 只支持 `%x` `%d` `%s` `%#x`
- 参数从数据栈弹出
- 输出到日志缓冲区（不直接 printf，VMX-root 不能 I/O）

## 文件路径约定

**原则**：Go 包结构与 C/C++ 目录 1:1 对应，便于对照验证。C/C++ 代码原位保留不动。

```
HyperDbgUnified/
│
├── HyperDbg/hyperdbg/           # ★ 原始 C/C++ 代码，全部保留作对照组
│   ├── libhyperdbg/             #   原用户态库（C++，不动）
│   ├── hyperdbg-cli/            #   原 CLI（C++，不动）
│   ├── script-engine/           #   原 .ds 编译器（C，不动）
│   ├── script-eval/             #   原 .ds 字节码 VM（C，双上下文，不动）
│   ├── symbol-parser/           #   原符号解析（C++，不动）
│   ├── hyperhv/                 #   VMM 驱动（C/ASM，不动）
│   ├── hyperkd/                 #   调试器驱动（C，不动）
│   │   └── code/go-interp/      #   ★ 新增：Go 子集 C 解释器
│   │       ├── ast_decode.c/.h
│   │       ├── interp.c/.h
│   │       ├── whitelist.c/.h
│   │       └── interp_stub.c    #   hook 命中接入（与 script-eval 并行）
│   └── ...
│
├── go-libhyperdbg/              # ★ Go 版 libhyperdbg（全量复刻，独立 go.mod）
│   ├── go.mod
│   ├── app/                     #   对应 libhyperdbg/code/app/
│   ├── common/                  #   对应 libhyperdbg/code/common/
│   ├── ucpuid/                  #   对应 libhyperdbg/ucpuid.cpp
│   ├── export/                  #   对应 libhyperdbg/code/export/
│   ├── objects/                 #   对应 libhyperdbg/code/objects/
│   ├── rev/                     #   对应 libhyperdbg/code/rev/
│   ├── hwdbg/                   #   对应 libhyperdbg/code/hwdbg/
│   ├── symbolparser/            #   对应 symbol-parser/（用 ok/pdbex）
│   ├── debugger/
│   │   ├── commands/
│   │   │   ├── debugging/       #   37 个调试命令（对应 debugging-commands/）
│   │   │   ├── extension/       #   35 个扩展命令（对应 extension-commands/）
│   │   │   ├── hwdbg/           #   2 个硬件调试命令
│   │   │   └── meta/            #   23 个元命令
│   │   ├── comm/                #   通信层（forwarding/namedpipe/tcp）
│   │   ├── core/                #   break-control/debugger/interpreter/steppings
│   │   ├── driverloader/        #   驱动加载（用 ok/ddk.Driver）
│   │   ├── kernellvl/           #   内核级调试
│   │   ├── misc/                #   assembler(用 ok/keystone)/disassembler(用 ok/zydis)/callstack/readmem/pt-helper/pci-id
│   │   ├── scriptengine/        #   脚本引擎包装（调用 go-bridge）
│   │   ├── tests/               #   测试
│   │   ├── transparency/        #   反检测
│   │   └── userlevel/           #   用户态调试（pe-parser 用 ok/pe-main）
│   └── api/                     #   顶层 API 汇总（Connect/Hook/StartProcess...）
│
├── go-cli/                      # ★ Go 版 hyperdbg-cli（独立）
│   ├── go.mod
│   ├── main.go                  #   对应 hyperdbg-cli.cpp
│   ├── repl.go                  #   交互式 REPL（yaegi）
│   └── script.go                #   脚本模式（执行 .go 脚本）
│
├── go-bridge/                   # ★ Go→驱动桥接（AST 编译 + 协议）
│   ├── go.mod
│   ├── ast/                     #   go/ast → 二进制 AST 编码器
│   ├── subset/                  #   Go 子集静态校验
│   └── protocol/                #   AST 二进制协议定义（Go 端）
│
├── go-test/                     # ★ Go 版测试 + 对照验证
│   ├── go.mod
│   ├── parity/                  #   对照测试用例（Go vs C/C++ 输出一致性）
│   │   └── <case>/
│   │       ├── input.txt
│   │       ├── expected_cpp.txt
│   │       └── expected_go.txt
│   └── ...
│
├── docs/
│   └── go-subset-spec.md        #   Go 子集 + AST 二进制协议 spec
│
└── ok/                          # ★ 已有 Go 绑定包（直接利用）
    ├── ddk/                     #   驱动管理（Driver 类型可用，KernelMemory 不用）
    ├── pe-main/                 #   PE 解析（直接用）
    ├── pdbex/                   #   PDB 符号（DIA 接口可用）
    ├── pdbfetch-master/         #   PDB 下载（直接用）
    ├── zydis/                   #   反汇编（直接用）
    ├── xed/                     #   反汇编备选
    ├── keystone/                #   汇编（直接用）
    ├── gjson/                   #   JSON（直接用）
    └── byteslice/               #   字节工具（直接用）
```

**Go 模块独立性**：`go-libhyperdbg`、`go-cli`、`go-bridge`、`go-test` 各自独立 `go.mod`，可单独编译。`ok/` 下的包作为依赖被引用。

**对照原则**：C/C++ 文件路径不变，Go 包路径与之平行对应。同名同语义的模块方便 diff 对照。

## find-oep.go 示例（目标形态）

```go
package main

import "hyperdbg"

func main() {
    hyperdbg.Connect("local")
    hyperdbg.LoadVMM()
    defer hyperdbg.UnloadVMM()

    proc := hyperdbg.StartProcess(`D:\...\SuperRecovery_V4.8.1.5.exe`)
    proc.WaitEntryPoint()

    hyperdbg.LogOpen(`D:\...\find-oep.log`)
    defer hyperdbg.LogClose()

    hyperdbg.EptHook("ntdll!RtlAllocateHeap", `
        func(ctx *hyperdbg.HookCtx) {
            ret := ctx.StackReadQword(0) & 0xFFFFFFFF
            ctx.Printf("RAH ret=%x", ret)
            if ret < 0x10000000 {
                ctx.Printf(" *** LOW ***")
            }
            if (ret & 0xFFFF0000) == 0x00c10000 {
                ctx.Break()
            }
        }
    `)

    hyperdbg.EptHook("kernelbase!VirtualAlloc", `
        func(ctx *hyperdbg.HookCtx) {
            ret := ctx.StackReadQword(0) & 0xFFFFFFFF
            ctx.Printf("VA ret=%x", ret)
        }
    `)

    proc.Continue()
    hyperdbg.Sleep(30 * hyperdbg.Second)
    proc.Pause()
}
```

## 决策记录

1. **保留所有 C/C++ 代码不清理**：新旧并存，Go 实现作为独立包并行存在，便于对照验证一致性。原 `libhyperdbg`、`hyperdbg-cli`、`script-engine`、`script-eval`、`symbol-parser`、`hyperdbg-test` 全部保留作对照组
2. **Go 实现是独立包，不是替换**：Go 版可独立编译运行，与 C/C++ 版对照输出验证行为一致。两套实现长期共存
3. **全量复刻用户层调试体系**：`HyperDbg/hyperdbg/` 下所有用户层模块都要 Go 实现，不能缺斤少两（见"复刻范围"章节）
4. **不 transpile Go→旧字节码**：新 Go 解释器独立实现，不复用 script-eval 的 ISA，避免被旧设计束缚
5. **不用路线 ③ 事件驱动**：用户明确选择"内核 C 解释 Go 子集"，保留内核脚本能力
6. **不用 yaegi 内核版**：Go runtime 进不了 VMX-root，C 重写 Go 子集解释器是唯一选择
7. **ok/ddk 的 KernelMemory 不用**：那是 RTCore64 漏洞驱动方案，与 HyperDbg VMM 冲突；只用 ddk.Driver 做驱动加载管理
8. **杜绝 cgo，全 syscall**：所有 DLL 调用（zydis/keystone/libipt/msdia140 等）必须用 `windows.NewLazyDLL` + `LazyProc.Call()` 模式，参考 `ok/zydis/dll.go` 和 `ok/keystone/dll.go`。DLL 用 `//go:embed` 内嵌，运行时释放到 UserCacheDir。理由：cgo 需要 C 工具链，跨平台编译麻烦，且 HyperDbg 已有 ok/ 下纯 syscall 绑定可复用

## 复刻范围（全量，对应 HyperDbg/hyperdbg/ 下用户层模块）

### 1. libhyperdbg 用户态核心库（全量复刻）

对应 `HyperDbg/hyperdbg/libhyperdbg/`，按 C/C++ 目录结构 1:1 映射到 Go 包：

| C/C++ 模块 | Go 包 | 内容 |
|---|---|---|
| `code/app/` | `go-libhyperdbg/app/` | dllmain、libhyperdbg 入口、messaging、packets |
| `code/common/` | `go-libhyperdbg/common/` | common、spinlock |
| `code/debugger/commands/debugging-commands/` | `go-libhyperdbg/debugger/commands/debugging/` | 37 个命令：a/bc/bd/be/bl/bp/continue/core/cpu/d-u/dt-struct/e/eval/events/exit/flush/g/gg/gu/i/k/lm/load/output/p/pause/preactivate/prealloc/print/r/rdmsr/s/settings/sleep/t/test/unload/wrmsr/x |
| `code/debugger/commands/extension-commands/` | `go-libhyperdbg/debugger/commands/extension/` | 35 个命令：apic/cpuid/crwrite/dr/epthook/epthook2/exception/hide/idt/interrupt/ioapic/ioin/ioout/lbr/lbrdump/measure/mode/monitor/msrread/msrwrite/pa2va/pcicam/pcitree/pmc/pt/pte/rev/smi/syscall-sysret/trace/track/tsc/unhide/va2pa/vmcall/xsetbv |
| `code/debugger/commands/hwdbg-commands/` | `go-libhyperdbg/debugger/commands/hwdbg/` | hw、hw_clk |
| `code/debugger/commands/meta-commands/` | `go-libhyperdbg/debugger/commands/meta/` | 23 个命令：attach/cls/connect/debug/detach/disconnect/dump/formats/help/kill/listen/logclose/logopen/pagein/pe/process/restart/script/start/status/switch/sym/sympath/thread |
| `code/debugger/communication/` | `go-libhyperdbg/debugger/comm/` | forwarding、namedpipe、remote-connection、tcpclient、tcpserver |
| `code/debugger/core/` | `go-libhyperdbg/debugger/core/` | break-control、debugger、interpreter、steppings |
| `code/debugger/driver-loader/` | `go-libhyperdbg/debugger/driverloader/` | install（用 ok/ddk.Driver） |
| `code/debugger/kernel-level/` | `go-libhyperdbg/debugger/kernellvl/` | kd、kernel-listening |
| `code/debugger/misc/` | `go-libhyperdbg/debugger/misc/` | assembler（用 ok/keystone）、callstack、disassembler（用 ok/zydis）、pci-id、pt-helper、readmem |
| `code/debugger/script-engine/` | `go-libhyperdbg/debugger/scriptengine/` | script-engine-wrapper、script-engine、symbol（用 ok/pdbex） |
| `code/debugger/tests/` | `go-libhyperdbg/debugger/tests/` | tests |
| `code/debugger/transparency/` | `go-libhyperdbg/debugger/transparency/` | gaussian-rng、transparency |
| `code/debugger/user-level/` | `go-libhyperdbg/debugger/userlevel/` | pe-parser（用 ok/pe-main）、ud、user-listening |
| `code/export/` | `go-libhyperdbg/export/` | export |
| `code/hwdbg/` | `go-libhyperdbg/hwdbg/` | hwdbg-interpreter、hwdbg-scripts |
| `code/objects/` | `go-libhyperdbg/objects/` | objects |
| `code/rev/` | `go-libhyperdbg/rev/` | rev-ctrl |
| `ucpuid.cpp` | `go-libhyperdbg/ucpuid/` | CPUID 命令处理 |

**命令总数**：37 + 35 + 2 + 23 = **97 个命令**，必须全部 Go 实现。

### 2. hyperdbg-cli（CLI 入口复刻）

对应 `HyperDbg/hyperdbg/hyperdbg-cli/hyperdbg-cli.cpp` → `go-cli/hyperdbg-cli.go`

- 命令行参数解析
- 交互式 REPL（用 yaegi 执行用户输入的 Go 代码）
- 脚本模式（执行 .go 脚本，替代 .ds）
- post-script 暂停逻辑（保留原 `getchar()` 等价行为）

### 3. script-engine（脚本编译器复刻）

对应 `HyperDbg/hyperdbg/script-engine/` → `go-bridge/scriptengine/`

- 不复刻 LALR parser（Grammar.txt + parse-table.c）
- 改为：go/ast 解析 Go 子集源码 → 二进制 AST 下发驱动
- 保留对照：原 .ds 脚本引擎不动，新 Go 脚本走新协议

### 4. script-eval 用户态部分（yaegi 替代）

对应 `HyperDbg/hyperdbg/script-eval/` 用户态那份（`SCRIPT_ENGINE_USER_MODE`）：

- 用 yaegi 在用户态执行 Go 顶层逻辑
- 用户态表达式求值（`?`、`print`、`eval`、`r` 命令中的表达式）
- 不需要复刻 SYMBOL_BUFFER 字节码执行

### 5. symbol-parser（符号解析复刻）

对应 `HyperDbg/hyperdbg/symbol-parser/` → `go-libhyperdbg/symbolparser/`

- PDB 加载/解析：用 `ok/pdbex/`（DIA 接口）
- PDB 下载：用 `ok/pdbfetch-master/`
- 符号反查（地址→符号名）、符号→地址
- 32 位/64 位 PDB 区分（WOW64 支持）

### 6. hyperdbg-test（测试复刻）

对应 `HyperDbg/hyperdbg/hyperdbg-test/` → `go-test/`

- 测试用例 Go 化
- 对照验证：同一输入下，Go 版和 C/C++ 版输出应一致

### 7. 驱动内新增 Go 子集解释器（C 实现）

对应新增 `HyperDbg/hyperdbg/hyperkd/code/go-interp/`：

- `ast_decode.c/.h` —— AST 二进制反序列化
- `interp.c/.h` —— Go 子集解释执行
- `whitelist.c/.h` —— HookCtx 白名单方法实现
- `interp_stub.c` —— hook 命中路径接入（替换原 script-eval 调用）

**注意**：原 `hyperkd/code/debugger/` 下的 EPT hook 命中路径**保留调用 script-eval**，新增并行路径调用 go-interp。两条路径用 hook 注册时的标志位区分（`HOOK_FLAG_GO_AST` vs `HOOK_FLAG_DS_BYTECODE`）。

## 对照验证策略

每个模块 Go 实现后，必须与 C/C++ 版对照验证：

1. **命令行为对照**：同一命令、同一输入，Go 版和 C/C++ 版输出应一致
2. **通信协议对照**：Go 版发出的 IOCTL/Named Pipe 消息，C/C++ 版驱动应能正确接收
3. **符号解析对照**：同一 PDB、同一地址，Go 版和 C/C++ 版解析的符号名应一致
4. **脚本执行对照**：等价的 .go 脚本（Go 版）和 .ds 脚本（C 版）应产生相同的事件序列

对照测试用例放在 `go-test/parity/` 下，每个用例包含：
- `input.txt` —— 输入命令/脚本
- `expected_cpp.txt` —— C/C++ 版输出
- `expected_go.txt` —— Go 版输出（应与 C/C++ 版一致）

## 通信协议与数据结构（Go 重写的硬约束）

### 1. IOCTL 完整列表（来源：`include/SDK/headers/Ioctls.h`）

Go 通信层必须支持以下全部 IOCTL，定义为常量：

**基础组（IOCTL_BASIC_IOCTL = 0x800 + 0x00）**：
- `IOCTL_INIT_VMM` (+0x01) —— 加载 VMM
- `IOCTL_INIT_HYPERTRACE` (+0x02) —— 加载 Hypertrace
- `IOCTL_REGISTER_EVENT` (+0x03) —— 注册事件（旧路径）
- `IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL` (+0x04)

**VMM 组（IOCTL_VMM_IOCTL = 0x800 + 0x200）**：
- `IOCTL_TERMINATE_VMX` (+0x01) —— 卸载 VMM
- `IOCTL_DEBUGGER_READ_MEMORY` (+0x02) —— 读内存
- `IOCTL_DEBUGGER_READ_OR_WRITE_MSR` (+0x03) —— 读写 MSR
- `IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS` (+0x04) —— 读 PTE
- `IOCTL_DEBUGGER_REGISTER_EVENT` (+0x05) —— 注册事件（新路径）
- `IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT` (+0x06) —— 给事件加动作（含脚本）
- `IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER` (+0x07) —— 隐藏调试器
- `IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS` (+0x08) —— 地址转换
- `IOCTL_DEBUGGER_EDIT_MEMORY` (+0x09) —— 写内存
- `IOCTL_DEBUGGER_SEARCH_MEMORY` (+0x0a) —— 搜索内存
- `IOCTL_DEBUGGER_MODIFY_EVENTS` (+0x0b) —— 修改/启用/禁用事件
- `IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS` (+0x0c) —— 刷新日志
- `IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS` (+0x0d) —— 附加/分离进程
- `IOCTL_DEBUGGER_PRINT` (+0x0e) —— 打印
- `IOCTL_PREPARE_DEBUGGEE` (+0x0f) —— 准备调试目标
- `IOCTL_PAUSE_PACKET_RECEIVED` (+0x10) —— 暂停包
- `IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED` (+0x11)
- `IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER` (+0x12)
- `IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER` (+0x13)
- `IOCTL_PERFORM_KERNEL_SIDE_TESTS` (+0x14)
- `IOCTL_RESERVE_PRE_ALLOCATED_POOLS` (+0x15)
- `IOCTL_SEND_USER_DEBUGGER_COMMANDS` (+0x16)
- `IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES` (+0x17)
- `IOCTL_GET_USER_MODE_MODULE_DETAILS` (+0x18) —— 获取模块详情

**Go 实现要点**：
- 用 `golang.org/x/sys/windows.DeviceIoControl` 发送
- IOCTL code 用 `CTL_CODE` 宏等价计算（FILE_DEVICE_UNKNOWN=0x22, METHOD_BUFFERED=0）
- 数据结构按 C 原始布局 1:1 映射到 Go（注意对齐和 packing）

### 2. 事件结构（来源：`include/SDK/headers/Events.h` + `RequestStructures.h`）

Go 必须复刻以下结构（用于 `IOCTL_DEBUGGER_REGISTER_EVENT` + `IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT`）：

```go
// 对应 DEBUGGER_EVENT_ACTION_TYPE_ENUM
type EventActionType uint32
const (
    ActionRunScriptCode   EventActionType = ...  // 执行脚本（旧 .ds 字节码 / 新 Go AST）
    ActionRunCustomCode   EventActionType = ...
    ActionBreak           EventActionType = ...
    ActionLog             EventActionType = ...
    ActionClear           EventActionType = ...
)

// 对应 DEBUGGER_EVENT_OPTIONS
type EventOptions struct {
    ...
}

// 对应 DEBUGGER_EVENT_ACTION（含 ActionType + 脚本配置）
type EventAction struct {
    ActionType           EventActionType
    ScriptConfiguration  EventActionRunScriptConfiguration
    ...
}

// 对应 DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
type EventActionRunScriptConfiguration struct {
    ScriptCode      []byte  // 旧：.ds 字节码；新：Go AST 二进制
    ScriptCodeSize  uint32
    ...
}

// 对应 DEBUGGER_EVENT_REQUEST_BUFFER
type EventRequestBuffer struct {
    EventRequest EventRequest
    ...
}
```

**关键 API（libhyperdbg/code/debugger/core/debugger.cpp）**：
- `SendEventToKernel(PDEBUGGER_GENERAL_EVENT_DETAIL Event, UINT32 EventBufferLength)` —— 通过 `IOCTL_DEBUGGER_REGISTER_EVENT` 发送事件注册
- `RegisterActionToEvent(PDEBUGGER_GENERAL_EVENT_DETAIL Event, PDEBUGGER_GENERAL_ACTION Action, UINT32 ActionBufferLength)` —— 通过 `IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT` 给事件加动作

Go 重写时这两个函数对应 `Debugger.SendEvent` 和 `Debugger.AddAction`，是 hook 注册的核心。

### 3. 命令注册表（来源：`libhyperdbg/code/debugger/core/interpreter.cpp`）

C/C++ 用 `std::map<string, COMMAND_PROPERTIES>` 注册命令：

```cpp
g_CommandsList["g"]  = {&CommandG, &CommandGHelp, DEBUGGER_COMMAND_G_ATTRIBUTES};
g_CommandsList["go"] = {&CommandG, &CommandGHelp, DEBUGGER_COMMAND_G_ATTRIBUTES};  // 别名
g_CommandsList[".connect"] = {&CommandConnect, ...};
g_CommandsList["connect"]  = {&CommandConnect, ...};  // 别名
```

`COMMAND_PROPERTIES` 结构含：命令函数指针、帮助函数指针、属性（可见性、命令类型）。

**Go 重写方式**：
```go
type CommandSpec struct {
    Handler func(ctx context.Context, d *Debugger, args []string) error
    Help    func(d *Debugger) error
    Attrs   CommandAttributes  // 可见性、类型
}

var commandRegistry = map[string]CommandSpec{
    "g":      {Debugger.CmdContinue, Debugger.CmdContinueHelp, ...},
    "go":     {Debugger.CmdContinue, Debugger.CmdContinueHelp, ...},  // 别名
    ".connect": {Debugger.CmdConnect, Debugger.CmdConnectHelp, ...},
    "connect":  {Debugger.CmdConnect, Debugger.CmdConnectHelp, ...},
    // ... 97 个命令 + 别名
}
```

**命令别名机制**：同一 Handler 注册多个命令字符串（如 `g`/`go`、`.connect`/`connect`、`.help`/`help`/`.hh`/`!help`）。Go 重写必须保留全部别名。

### 4. Named Pipe 通信（来源：`communication/namedpipe.cpp`）

- 服务器端：`CreateNamedPipeA` + `ConnectNamedPipe`
- 客户端：`CreateFileA` + `ReadFile`/`WriteFile`
- 管道名约定：`\\.\pipe\HyperDbgDebugPipe`（具体名待确认）

**Go 实现**：用 `github.com/Microsoft/go-winio` 的 `CreatePipe`/`DialPipe`。

### 5. Forwarding 机制（来源：`communication/forwarding.cpp`）

`ForwardingPerformEventForwarding` 把事件结果转发到 Named Pipe / 文件 / TCP。
Go 重写要支持这三种转发目标，事件结构对齐 C 版。

## ok 目录 Go 绑定包精确 API

### HyperDbgSdk.go（核心资产，类型层基础）

**文件**：`ok/HyperDbgSdk.go`（>128KB，单文件）
**包名**：`hyperdbgsdk`
**来源**：HyperDbgSdk.h 完整 Go 翻译（用 c2go 类工具生成）

**包含内容**：
- 所有 C 基础类型映射：`UINT64=uint64`、`PVOID=uintptr`、`HANDLE=uintptr`、`BOOL=int32`、`BOOLEAN=bool`、`DWORD=uint32` 等
- 所有指针类型别名：`PDEBUGGER_GENERAL_EVENT_DETAIL = *DEBUGGER_GENERAL_EVENT_DETAIL` 等
- **所有核心结构体 Go 定义**（关键字段对齐用 `_ [N]byte` 填充，保证 C ABI 兼容）：
  - `GUEST_REGS`（16 个通用寄存器 Rax-R15）
  - `GUEST_XMM_REGS`、`GUEST_EXTRA_REGISTERS`（含 CS/DS/FS/GS/RFLAGS/RIP）
  - `CR3_TYPE`、`VMX_SEGMENT_SELECTOR`
  - `EPT_HOOKS_CONTEXT`、`EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR`、`EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2`
  - `DEBUGGER_EVENT_OPTIONS`（6 个 OptionalParam）
  - `DEBUGGER_GENERAL_EVENT_DETAIL`（事件注册：EventType/Options/ConditionBufferSize/Tag）
  - `DEBUGGER_GENERAL_ACTION`（动作：ActionType/ScriptBufferSize/ScriptBufferPointer）
  - `DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION`（ScriptBuffer/ScriptLength）
  - `DEBUGGER_EVENT_REQUEST_BUFFER`、`DEBUGGER_EVENT_REQUEST_CUSTOM_CODE`
  - `DEBUGGER_EVENT_AND_ACTION_RESULT`
  - `DEBUGGER_READ_MEMORY`、`DEBUGGER_EDIT_MEMORY`、`DEBUGGER_SEARCH_MEMORY`
  - `DEBUGGER_READ_AND_WRITE_ON_MSR`、`DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS`
  - `DEBUGGER_VA2PA_AND_PA2VA_COMMANDS`
  - `DEBUGGER_PREPARE_DEBUGGEE`、`DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS`
  - `DEBUGGER_MODIFY_EVENTS`、`DEBUGGER_SHORT_CIRCUITING_EVENT`
  - `DEBUGGER_PREALLOC_COMMAND`、`DEBUGGER_PREACTIVATE_COMMAND`
  - `DEBUGGER_FLUSH_LOGGING_BUFFERS`、`DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE`
  - `DEBUGGEE_KD_PAUSED_PACKET`、`DEBUGGEE_UD_PAUSED_PACKET`
  - `DEBUGGER_CALLSTACK_REQUEST`、`DEBUGGER_SINGLE_CALLSTACK_FRAME`
  - `DEBUGGER_APIC_REQUEST`、`LAPIC_PAGE`、`IO_APIC_ENTRY_PACKETS`
  - `PCI_DEV`、`PORTABLE_PCI_CONFIG_SPACE_HEADER` 等 PCI 结构
  - `SYSCALL_CALLBACK_CONTEXT_PARAMS`、`DIRECT_VMCALL_PARAMETERS`
  - `REGISTER_NOTIFY_BUFFER`、`SCRIPT_ENGINE_GENERAL_REGISTERS`
  - `REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST`
  - `HYPERTRACE_LBR_OPERATION_PACKETS`、`SMI_OPERATION_PACKETS`
- **所有枚举类型**（含 `String()` 方法）：
  - `VMM_EVENT_TYPE_ENUM`、`DEBUGGER_EVENT_ACTION_TYPE_ENUM`
  - `VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE`
  - `DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE`
  - `DEBUGGER_REMOTE_PACKET_TYPE`、`DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION`
  - `DEBUGGEE_PAUSING_REASON`
  - `VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE`
- 所有 `const` 块（IOCTL 之外的常量）

**Go 重写用途**：
- 直接复制到 `go-libhyperdbg/types/`（或 `go-libhyperdbg/sdk/`），作为所有 IOCTL 通信、事件注册、hook 配置的类型基础
- **不需要手动映射 C 结构体到 Go**——所有结构体已现成，字段对齐已正确
- 通信层（T2.1）直接 `import "go-libhyperdbg/types"` 即可

**使用示例**：
```go
import "go-libhyperdbg/types"

// 注册 EPT hook 事件
event := types.DEBUGGER_GENERAL_EVENT_DETAIL{
    EventType: types.VMM_EVENT_TYPE_ENUM_EPT_HOOK,
    ProcessId: processId,
    Options: types.DEBUGGER_EVENT_OPTIONS{
        OptionalParam1: hookAddr,
    },
}

// 动作（含 Go AST 脚本）
action := types.DEBUGGER_GENERAL_ACTION{
    ActionType:       types.DEBUGGER_EVENT_ACTION_TYPE_ENUM_RUN_SCRIPT_CODE,
    ScriptBufferSize: uint32(len(astBytes)),
}
```

**注意事项**：
- 文件 >128KB，需切分（按结构体类别拆为多个 .go 文件）或保留单文件
- 包名 `hyperdbgsdk` 建议改为 `types` 或 `sdk`
- 字段 `_ [N]byte` 填充不可删除，否则破坏 C ABI 对齐
- 部分 `uintptr` 字段在 Go 里要小心 GC（指针藏入 uintptr）

### ddk（部分可用）

**导入路径**：`github.com/ddkwork/golibrary/std/ddk`（需确认，看 go.mod）

**可用 API**：
```go
type Driver struct {
    Path string
    Name string
}
func NewDriver(driverPath string) *Driver
func (d *Driver) Install() bool    // SCM CreateService
func (d *Driver) Remove() bool     // SCM DeleteService
func (d *Driver) Start() bool      // SCM StartService
func (d *Driver) Stop() bool       // SCM ControlService(STOP)
```

**不可用 API**（RTCore64 漏洞驱动方案，与 HyperDbg VMM 冲突）：
```go
type RTCore64 struct { ... }       // 不用
type KernelMemory struct { ... }   // 不用
```

**依赖**：`github.com/ddkwork/golibrary/byteslice`、`github.com/ddkwork/golibrary/std/mylog`、`golang.org/x/sys/windows`

**Go 重写用途**：`go-libhyperdbg/debugger/driverloader/` 复用 `Driver` 类型加载 hyperhv.sys / hyperkd.sys。

### pdbex（可用，syscall + COM，无 cgo）

**导入路径**：`github.com/ddkwork/golibrary/pdbex`（需确认）

**API**：
```go
type PDB struct { ... }
func NewPDB() *PDB
func (p *PDB) Open(path string) error      // 通过 DIA (msdia140.dll) 解析
func (p *PDB) Close()
func (p *PDB) IsOpened() bool
func (p *PDB) GetPath() string
func (p *PDB) GetMachineType() uint16      // 区分 32/64 位
func (p *PDB) GetSymbolByName(name string) (*Symbol, bool)
// 还有 GetFunctionByName, GetFunctionByAddr 等
```

**依赖**：msdia140.dll，通过 `syscall.NewLazyDLL("ole32.dll")` + `CoCreateInstance` 调用 DIA COM 接口（参考 `ok/pdbex/dia_windows.go`），无 cgo

**Go 重写用途**：`go-libhyperdbg/symbolparser/` 复用 `PDB` 类型，对齐 C 版 `SymLoadFileSymbol` / `SymConvertNameToAddress`。

### zydis（可用，syscall，无 cgo）

**导入路径**：`github.com/ddkwork/golibrary/zydis`

**API**：完整类型别名（`ZyanU8`/`ZydisMnemonic`/`ZydisDecodedInstruction` 等）+ `windows.NewLazyDLL` 加载 `zydis.dll`（用 `//go:embed` 内嵌 DLL，运行时释放到 UserCacheDir）。提供 `Decode` / `Format` / `RegisterGetString` 等函数，模式参考 `ok/zydis/dll.go`。

**Go 重写用途**：`go-libhyperdbg/debugger/misc/disassembler/` 复用，对齐 C 版 `ZydisDisassemble`。

### keystone（可用，syscall，无 cgo）

**导入路径**：`github.com/ddkwork/golibrary/keystone`

**API**：`Ks_engine` 类型 + `Ks_arch` 枚举（KsArchX86 等）+ `windows.NewLazyDLL` 加载 `keystone.dll`（用 `//go:embed` 内嵌，运行时释放），模式参考 `ok/keystone/dll.go`。提供汇编函数。

**Go 重写用途**：`go-libhyperdbg/debugger/misc/assembler/` 复用，对齐 C 版 `Assemble` 函数。

### pe-main（可用，纯 Go）

**导入路径**：`github.com/saferwall/pe`（看 go.mod）

**API**：`File` 类型，支持 imports/exports/sections/reloc/resource/debug/exception/delayimports/boundimports/tls/dotnet。

**Go 重写用途**：`go-libhyperdbg/debugger/userlevel/pe-parser/` 复用，对齐 C 版 `PeParse`。

### pdbfetch-master（可用，纯 Go）

**导入路径**：`github.com/pdbfetch/pe`（看 go.mod）

**API**：从 Microsoft Symbol Server 下载 PDB，含 PE 解析提取 PDB GUID。

**Go 重写用途**：`go-libhyperdbg/symbolparser/` 的 PDB 下载部分，对齐 C 版 `SymbolPdbDownload`。

### gjson（可用，纯 Go）

**Go 重写用途**：配置文件解析、命令输出 JSON 化（GUI/MCP 需要）。

### byteslice（可用，纯 Go）

**Go 重写用途**：`byteslice.FromStruct` / `PtrFromAnySlice` 用于 IOCTL 数据包序列化。

## API 设计规范（GUI/MCP 友好，强制）

### 核心原则

所有公开 API 必须满足以下原则，**否则后期 GUI/MCP 无法复用**：

1. **无全局状态**：调试器状态封装在 `Debugger` 结构体里，支持多实例（GUI 可同时调试多个目标，MCP 可并发服务多个 Agent）
2. **context 传播**：所有可能阻塞的 API 第一个参数是 `ctx context.Context`，支持取消和超时
3. **输出抽象**：命令输出不直接 `fmt.Println`，走 `Output` 接口（CLI 接 stdout，GUI 接 widget，MCP 接 JSON channel）
4. **错误返回**：用 `error` 不用 panic，错误信息结构化（含错误码、上下文）
5. **并发安全**：`Debugger` 方法可被多 goroutine 调用，内部用 sync 保护共享状态
6. **事件流**：hook 事件通过 channel 上送，消费者可选（CLI 可不消费，GUI/MCP 必须消费）

### 顶层 API 范式

```go
package api

// Debugger 是调试器核心，无全局状态，支持多实例
type Debugger struct {
    // 私有字段
}

// New 创建调试器实例
func New(opts ...Option) (*Debugger, error)

// Output 接口：CLI/GUI/MCP 各自实现
type Output interface {
    Write(p []byte) (int, error)
    WriteErr(p []byte) (int, error)
    Printf(format string, args ...any) error
}

// Event 是 hook 命中事件
type Event struct {
    Type      EventType
    HookID    uint64
    Regs      Registers
    Stack     []byte
    Timestamp uint64
}

// Option 配置选项
type Option func(*Debugger) error

// 核心方法（都带 ctx）
func (d *Debugger) Connect(ctx context.Context, target string) error
func (d *Debugger) LoadVMM(ctx context.Context) error
func (d *Debugger) UnloadVMM(ctx context.Context) error
func (d *Debugger) StartProcess(ctx context.Context, path string) (*Process, error)
func (d *Debugger) Continue(ctx context.Context) error
func (d *Debugger) Pause(ctx context.Context) error
func (d *Debugger) EptHook(ctx context.Context, symbol string, callbackSrc string) (hookID uint64, err error)
func (d *Debugger) RemoveHook(ctx context.Context, hookID uint64) error
func (d *Debugger) ReadMem(ctx context.Context, addr uint64, buf []byte) error
func (d *Debugger) Events() <-chan Event  // 事件流，消费者可选

// 命令执行（统一入口，GUI/MCP 可调用）
func (d *Debugger) Exec(ctx context.Context, cmd string) error
```

### CLI / GUI / MCP 复用方式

```go
// CLI 用法
dbg, _ := api.New(api.WithOutput(os.Stdout))
dbg.Connect(ctx, "local")
dbg.Exec(ctx, "lm")  // 直接执行命令字符串

// GUI 用法（后期）
dbg, _ := api.New(api.WithOutput(guiTextWriter))
events := dbg.Events()
go func() {
    for ev := range events {
        guiShowHookHit(ev)  // GUI 刷新 hook 命中显示
    }
}()

// MCP 用法（后期）
dbg, _ := api.New(api.WithOutput(mcpJsonWriter))
// MCP tool: "list_modules" -> dbg.Exec(ctx, "lm")
// MCP tool: "set_hook" -> dbg.EptHook(ctx, sym, cb)
// MCP tool: "read_mem" -> dbg.ReadMem(ctx, addr, buf)
```

### 命令实现规范

每个命令（97 个）的实现必须：
- 是 `Debugger` 的方法，不是包级函数
- 第一个参数 `ctx context.Context`
- 通过 `d.output` 写输出，不直接 `fmt`
- 返回 `error`，不 panic
- 可被 `Exec(ctx, cmdString)` 统一调度，也可被 GUI/MCP 直接方法调用

```go
// 示例：lm 命令（列出模块）
func (d *Debugger) CmdListModules(ctx context.Context, args []string) error {
    modules, err := d.kd.ListModules(ctx)
    if err != nil {
        return err
    }
    for _, m := range modules {
        d.output.Printf("%-32s %016x %08x\n", m.Name, m.Base, m.Size)
    }
    return nil
}
```

## 未来扩展（本次预留接口，不实现）

### Go 版 GUI

- 基于 `go-libhyperdbg/api/` 公开 API
- 候选框架：Wails（Web 后端）/ Fyne（原生）/ 自绘
- 关键依赖：API 无全局状态、输出可重定向、事件流可消费
- GUI 特有能力：hook 命中可视化、寄存器实时刷新、内存视图、反汇编视图
- **本次必须保证**：API 不绑死 CLI，GUI 能拿到所有必要数据（events、regs、mem）

### MCP（Model Context Protocol）层

把 HyperDbg 调试器能力暴露给 AI Agent：

- **Tools**（可调用函数）：
  - `connect(target)` / `load_vmm()` / `start_process(path)`
  - `list_modules()` / `list_threads()` / `read_memory(addr, len)`
  - `set_hook(symbol, callback)` / `remove_hook(id)`
  - `continue()` / `pause()` / `step()`
  - `disasm(addr, count)` / `asm(code)`
- **Resources**（可读数据）：
  - 寄存器快照、内存区域、模块列表、符号表
- **Prompts**（预设场景）：
  - "find OEP of packed exe" → 自动执行 find-oep.go 流程
  - "analyze crash at addr" → 自动反汇编 + 符号化栈
- **关键依赖**：API 可程序化调用（无 REPL 依赖）、并发安全、输出可序列化为 JSON

**本次必须保证**：
- `Debugger` 结构体可被 MCP server 持有
- 所有命令可通过方法调用（不必走 `Exec` 字符串解析）
- 事件流可被 MCP server 转为 JSON 通知

## 反模式复查清单

每次提交 Go 代码前检查：

- [ ] 命令逻辑在 `go-libhyperdbg/` 包内，不在 `go-cli/`
- [ ] `Debugger` 方法带 `ctx context.Context`
- [ ] 输出走 `d.output`，不直接 `fmt.Println`
- [ ] 无包级全局变量（除常量）
- [ ] 错误用 `error` 返回，不 panic
- [ ] hook 注册返回 ID，不靠全局状态匹配
- [ ] 事件流是 channel，消费者可选
- [ ] API 可被 GUI/MCP 直接调用（无需经过 CLI REPL）

## 实现细节任务清单（给用户检查）

本章节是落地执行的总纲，每项任务标注：编号、文件路径、依赖、验收标准。

### 阶段 0：地基（无依赖，必须先做）

#### T0.1 编写 Go 子集 + AST 二进制协议 spec
- **文件**：`docs/go-subset-spec.md`
- **依赖**：无
- **内容**：
  - Go 子集语法定义（支持/不支持清单）
  - AST 节点二进制格式（opcode 表）
  - 字符串表编码
  - 闭包捕获限制规则
  - cycle budget 协议
- **验收**：spec 文档可独立指导 Go 编译器和 C 解释器两端实现

#### T0.2 go-bridge/protocol 包
- **文件**：`go-bridge/protocol/opcodes.go`、`nodes.go`、`errors.go`
- **依赖**：T0.1
- **内容**：AST 节点 opcode 常量、Go 结构体定义、编码/解码接口
- **验收**：`go test ./protocol` 通过，覆盖所有节点类型编解码

#### T0.3 go-bridge/subset 包
- **文件**：`go-bridge/subset/validator.go`、`checker.go`
- **依赖**：T0.1
- **内容**：go/ast 遍历器，拒绝子集外语法（goroutine/channel/interface/map/slice/reflect/defer/外部捕获）
- **验收**：合法子集代码通过，非法代码报错并指出位置

#### T0.4 go-bridge/ast 包
- **文件**：`go-bridge/ast/encoder.go`、`parser.go`、`cache.go`
- **依赖**：T0.2、T0.3
- **内容**：`go/ast.Parse(src)` → 子集校验 → 二进制 AST 序列化
- **验收**：`Encode(callbackSrc)` 返回 `[]byte`，可被 C 反序列化

### 阶段 1：内核侧（驱动 Go 解释器）

#### T1.1 hyperkd/code/go-interp/ast_decode.c
- **文件**：`HyperDbg/hyperdbg/hyperkd/code/go-interp/ast_decode.c` + `.h`
- **依赖**：T0.2
- **内容**：二进制 AST → C 结构体节点树，严格长度校验
- **验收**：单元测试覆盖所有节点类型，畸形输入不 crash

#### T1.2 hyperkd/code/go-interp/interp.c
- **文件**：`HyperDbg/hyperdbg/hyperkd/code/go-interp/interp.c` + `.h`
- **依赖**：T1.1
- **内容**：Go 子集解释执行（MVP：if/赋值/算术/Printf/StackReadQword），cycle budget，SEH 包裹
- **验收**：简单脚本正确执行，`for {}` 超时不 BSOD，类型错误不 BSOD

#### T1.3 hyperkd/code/go-interp/whitelist.c
- **文件**：`HyperDbg/hyperdbg/hyperkd/code/go-interp/whitelist.c` + `.h`
- **依赖**：T1.2
- **内容**：HookCtx 方法实现（StackReadQword/Reg/ReadMem/Printf/Break/Continue 等）
- **验收**：每个方法在 VMX-root 下正确读写客户机状态

#### T1.4 hyperkd/code/go-interp/interp_stub.c
- **文件**：`HyperDbg/hyperdbg/hyperkd/code/go-interp/interp_stub.c`
- **依赖**：T1.2、T1.3
- **内容**：EPT hook 命中路径接入，用 `HOOK_FLAG_GO_AST` 区分新路径，与原 script-eval 并行
- **验收**：同一 hook 可注册 .ds 字节码或 Go AST，两条路径独立工作

#### T1.5 CMakeLists 集成
- **文件**：`CMakeLists.txt`（根）
- **依赖**：T1.1-T1.4
- **内容**：把 go-interp/*.c 加入 hyperkd target
- **验收**：`build.cmd` 成功编译，201/201

### 阶段 2：用户态核心（让 find-oep.go 跑起来）

#### T2.0 复制 HyperDbgSdk.go 建立类型层
- **文件**：`go-libhyperdbg/types/`（从 `ok/HyperDbgSdk.go` 复制并切分）
- **依赖**：无
- **内容**：
  - 复制 `ok/HyperDbgSdk.go` 到 `go-libhyperdbg/types/sdk.go`（或按类别切分为 `basics.go`/`regs.go`/`events.go`/`memory.go`/`pci.go` 等）
  - 包名改为 `types`
  - 验证字段对齐与 C 版一致（用 `unsafe.Sizeof` 对照 C 的 `sizeof`）
- **验收**：`go build ./types` 通过，结构体大小与 C 版一致

#### T2.1 go-libhyperdbg/debugger/comm 包
- **文件**：`go-libhyperdbg/debugger/comm/ioctl.go`、`namedpipe.go`、`device.go`
- **依赖**：T2.0（用 `types` 包的结构体）
- **内容**：`DeviceIoControl` 封装 + Named Pipe 客户端（go-winio）+ 所有 IOCTL 常量
- **验收**：能打开 `\\.\HyperDbgDevice`，发送 `IOCTL_INIT_VMM`（用 `types` 包结构体打包），驱动响应正确

#### T2.2 go-libhyperdbg/debugger/driverloader 包
- **文件**：`go-libhyperdbg/debugger/driverloader/loader.go`
- **依赖**：`ok/ddk.Driver`
- **内容**：用 `ddk.Driver` 加载/卸载 hyperhv.sys / hyperkd.sys
- **验收**：`loader.Load("hyperhv.sys")` 后驱动服务运行，`loader.Unload()` 后清理

#### T2.3 go-libhyperdbg/debugger/core 包
- **文件**：`go-libhyperdbg/debugger/core/debugger.go`、`interpreter.go`、`break_control.go`、`steppings.go`
- **依赖**：T2.1
- **内容**：`Debugger` 结构体 + `Exec(ctx, cmd)` 命令分发 + 状态机
- **验收**：能执行 `connect local`、`load vmm`、`g`、`pause` 等基本命令

#### T2.4 go-libhyperdbg/api 包
- **文件**：`go-libhyperdbg/api/debugger.go`、`options.go`、`output.go`、`event.go`
- **依赖**：T2.3、T0.4
- **内容**：顶层 API（`New/Connect/LoadVMM/StartProcess/Continue/Pause/EptHook/Events/ReadMem`）+ `Output` 接口 + `Option` 模式
- **验收**：`api.New(WithOutput(os.Stdout))` 创建实例，方法可调用

#### T2.5 跑通 find-oep.go
- **文件**：`go-cli/examples/find-oep.go`
- **依赖**：T2.4、T1.4
- **内容**：用 Go API 编写等价于 `find-oep.ds` 的脚本
- **验收**：脚本执行后 `find-oep.log` 包含 `RAH ret=xxx` 行，与 C 版输出对照一致

### 阶段 3：命令全量复刻（97 个命令）

每个命令对应一个 Go 文件，文件名与 C++ 一致（如 `g.cpp` → `g.go`）。每完成一组与 C/C++ 版对照。

#### T3.1 meta 命令（23 个）
- **文件**：`go-libhyperdbg/debugger/commands/meta/*.go`
- **命令**：attach/cls/connect/debug/detach/disconnect/dump/formats/help/kill/listen/logclose/logopen/pagein/pe/process/restart/script/start/status/switch/sym/sympath/thread
- **依赖**：T2.3
- **验收**：每个命令的输出与 C 版逐字节一致（`go-test/parity/meta/`）

#### T3.2 debugging 命令（37 个）
- **文件**：`go-libhyperdbg/debugger/commands/debugging/*.go`
- **命令**：a/bc/bd/be/bl/bp/continue/core/cpu/d-u/dt-struct/e/eval/events/exit/flush/g/gg/gu/i/k/lm/load/output/p/pause/preactivate/prealloc/print/r/rdmsr/s/settings/sleep/t/test/unload/wrmsr/x
- **依赖**：T2.3、T3.5（misc 反汇编/汇编）
- **验收**：同上

#### T3.3 extension 命令（35 个）
- **文件**：`go-libhyperdbg/debugger/commands/extension/*.go`
- **命令**：apic/cpuid/crwrite/dr/epthook/epthook2/exception/hide/idt/interrupt/ioapic/ioin/ioout/lbr/lbrdump/measure/mode/monitor/msrread/msrwrite/pa2va/pcicam/pcitree/pmc/pt/pte/rev/smi/syscall-sysret/trace/track/tsc/unhide/va2pa/vmcall/xsetbv
- **依赖**：T2.3
- **验收**：同上

#### T3.4 hwdbg 命令（2 个）
- **文件**：`go-libhyperdbg/debugger/commands/hwdbg/*.go`
- **命令**：hw、hw_clk
- **依赖**：T2.3
- **验收**：同上

#### T3.5 misc 模块
- **文件**：`go-libhyperdbg/debugger/misc/assembler.go`、`disassembler.go`、`callstack.go`、`readmem.go`、`pt_helper.go`、`pci_id.go`
- **依赖**：`ok/keystone`、`ok/zydis`
- **验收**：汇编/反汇编输出与 C 版一致

### 阶段 4：支撑模块全量复刻

#### T4.1 userlevel 模块
- **文件**：`go-libhyperdbg/debugger/userlevel/pe_parser.go`、`ud.go`、`user_listening.go`
- **依赖**：`ok/pe-main`
- **验收**：PE 解析结果与 C 版一致

#### T4.2 kernellvl 模块
- **文件**：`go-libhyperdbg/debugger/kernellvl/kd.go`、`kernel_listening.go`
- **依赖**：T2.1
- **验收**：内核调试监听行为与 C 版一致

#### T4.3 transparency 模块
- **文件**：`go-libhyperdbg/debugger/transparency/gaussian_rng.go`、`transparency.go`
- **依赖**：无
- **验收**：反检测行为与 C 版一致

#### T4.4 symbolparser 模块
- **文件**：`go-libhyperdbg/symbolparser/parser.go`、`module.go`、`download.go`
- **依赖**：`ok/pdbex`、`ok/pdbfetch-master`
- **内容**：对齐 C 版 `SymLoadFileSymbol` / `SymConvertNameToAddress` / `SymConvertAddressToName` / `SymbolPdbDownload`，含 32/64 位 PDB 区分
- **验收**：同一 PDB、同一地址，解析符号名与 C 版一致

#### T4.5 其余模块
- **文件**：`go-libhyperdbg/app/`、`common/`、`export/`、`objects/`、`rev/`、`hwdbg/`、`ucpuid/`
- **依赖**：各模块
- **验收**：功能对齐 C 版

### 阶段 5：CLI + 脚本引擎 + 测试

#### T5.1 go-cli 入口
- **文件**：`go-cli/main.go`、`repl.go`、`script.go`
- **依赖**：T2.4、T3.1-T3.4
- **内容**：命令行参数 + REPL + 脚本模式 + post-script 暂停
- **验收**：`go-cli.exe` 行为对齐 `hyperdbg-cli.exe`

#### T5.2 scriptengine 包装
- **文件**：`go-libhyperdbg/debugger/scriptengine/wrapper.go`
- **依赖**：T0.4
- **内容**：把 `go-bridge/ast` 集成到 `Debugger.EptHook`，自动解析 Go 回调下发驱动
- **验收**：`EptHook(sym, goCallbackSrc)` 端到端工作

#### T5.3 yaegi 集成
- **文件**：`go-libhyperdbg/interp/yaegi.go`、`exports.go`
- **依赖**：T5.1
- **内容**：yaegi 解释器加载 Go 脚本，导出 `api` 包符号
- **验收**：`go-cli --script find-oep.go` 执行成功

#### T5.4 对照测试
- **文件**：`go-test/parity/*`
- **依赖**：T3.1-T3.4、T4.4
- **内容**：每个命令的 input/expected_cpp/expected_go 三件套
- **验收**：所有用例 Go 输出 == C++ 输出

#### T5.5 syscall 包装 libipt
- **文件**：`go-libhyperdbg/debugger/misc/pt/ipt.go`、`go-libhyperdbg/debugger/misc/pt/dll.go`
- **依赖**：libipt.dll（用 `//go:embed` 内嵌 + `windows.NewLazyDLL` 加载，参考 `ok/zydis/dll.go` 模式）
- **内容**：Intel PT 解码，对齐 C 版 `pt-helper.cpp`，无 cgo
- **验收**：PT 解码输出与 C 版一致

### 阶段 6：补全 Go 子集

#### T6.1 补全解释器
- **文件**：`hyperkd/code/go-interp/interp.c`（扩展）
- **依赖**：T1.2
- **内容**：补全 `for` 三段式、`return`、`ReadMem`、`SetReg`、`Break`、数组、`FuncDecl`
- **验收**：复杂脚本（含循环、多 return 路径）正确执行

#### T6.2 cycle budget 严格验证
- **文件**：`go-test/cycle-budget-test.go` + 对应驱动测试
- **依赖**：T6.1
- **内容**：各种循环场景的 cycle 上限测试
- **验收**：所有死循环场景都在 100 万 cycle 内安全返回，不 BSOD

### 任务依赖图（Phase 结构）

```
Phase A（驱动加载地基，无依赖）
  A.1 (types ← 移动 HyperDbgSdk.go)
  ├─ A.2 (driverloader ← ok/ddk)
  └─ A.3 (comm/IOCTL)
        └─ A.4 (单元测试) ← A.2, A.3
              │
              ▼ [Phase A 验收通过]
Phase B（Go 子集 + 驱动解释器）
  B.1 (spec) ── 可与 Phase A 并行
  ├─ B.2 (protocol) ← B.1
  ├─ B.3 (subset) ← B.1
  └─ B.4 (ast encoder) ← B.2, B.3
  │
  B.5 (ast_decode.c) ← B.2
  └─ B.6 (interp.c) ← B.5
        ├─ B.7 (whitelist.c) ← B.6
        └─ B.8 (interp_stub.c) ← B.6, B.7
              └─ B.9 (CMake) ← B.5-B.8
                    └─ B.10 (单元测试 + 回归) ← B.4, B.8, Phase A 通过
                          │
                          ▼ [Phase B 验收通过]
Phase C（全量复刻）
  C.1.1 (core) ← A.3
  └─ C.1.2 (api) ← C.1.1, B.4
        └─ C.1.3 (find-oep.go) ← C.1.2, B.8
              │
              ├─ C.2.1-C.2.4 (97 命令) ← C.1.1
              │    └─ C.3.1-C.3.6 (支撑模块) ← ok/* 包
              │         └─ C.4.1-C.4.5 (CLI/yaegi/test/libipt)
              │              └─ C.5.1-C.5.2 (补全子集) ← B.6
              │
              └─ [每个 C 步骤都有 parity 对照测试]
```

### 总工作量估算

| Phase | 步骤数 | Go 代码 | C 代码 | 备注 |
|---|---|---|---|---|
| Phase A（驱动加载地基） | 4 | ~1500 行（含类型层复制） | 0 | 最小可验证 |
| Phase B（Go 子集 + 驱动解释器） | 10 | ~1500 行 | ~4500 行 | spec + go-bridge + 驱动 C 解释器 |
| Phase C（全量复刻） | 18 | ~19000 行 | ~1500 行（补全 interp） | 97 命令 + 支撑 + CLI + 测试 |
| **总计** | **32** | **~22000 行 Go** | **~6000 行 C** | |
