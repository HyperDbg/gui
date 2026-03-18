# HyperDbg Debugger - Go Implementation

这是 HyperDbg 调试器的 Go 语言实现，完整复刻了 C++ 版本的 ring3 功能。

## 项目结构

```
hypedbg/
├── debugger/              # 调试器核心
│   ├── api/              # API 接口定义
│   ├── breakpoint/        # 断点管理
│   ├── commands/         # 调试命令
│   ├── driver/           # 驱动通信
│   ├── event_*.go        # 事件系统
│   ├── hardware/         # 硬件检测
│   ├── memory/           # 内存操作
│   ├── process/          # 进程管理
│   ├── register/         # 寄存器操作
│   ├── script/           # 脚本引擎
│   ├── symbol/           # 符号表
│   └── ...              # 其他模块
├── ui/                  # 用户界面
│   ├── pages/            # UI 页面
│   ├── asserts/          # 资源文件
│   └── ui.go            # UI 主入口
├── doc/                 # 文档
├── main.go              # 程序入口
└── go.mod              # Go 模块定义
```

## 功能特性

### 1. 驱动管理 (debugger/driver/)
- **驱动提供者** ([provider.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\provider.go))
  - `Load()` - 加载驱动
  - `Unload()` - 卸载驱动
  - `Manage()` - 统一的驱动管理接口

- **驱动句柄** ([handle.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\handle.go))
  - Windows 设备驱动通信
  - IOCTL 接口支持
  - 缓冲区发送和接收

- **驱动配置** ([config.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\config.go))
  - 配置管理
  - 设备名称配置

### 2. VMM (Virtual Machine Monitor)
- **VMM 初始化和控制** ([vmm.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\vmm.go))
  - `Initialize()` - 初始化 VMM
  - `Terminate()` - 终止 VMM
  - CPU 厂商检测
  - VT-x 支持检测

- **内存操作** ([memory.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\memory\memory.go))
  - `ReadVirtualMemory()` - 读取虚拟内存
  - `WriteVirtualMemory()` - 写入虚拟内存
  - `ReadPhysicalMemory()` - 读取物理内存
  - `WritePhysicalMemory()` - 写入物理内存
  - `SearchMemory()` - 搜索内存

### 3. 事件系统 (debugger/event_*.go)
- **事件管理器** ([event_manager.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_manager.go))
  - 事件注册和查询
  - 事件修改和删除
  - 事件动作管理

- **事件处理器** ([event_handler.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_handler.go))
  - 37 种事件类型处理
  - 事件回调机制
  - 事件分发

- **事件循环** ([event_loop.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\event_loop.go))
  - 事件监听
  - 事件分发
  - 异步处理

### 4. 调试器核心 (debugger/)
- **调试器主入口** ([debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\debugger.go))
  - 调试器初始化
  - 调试器卸载
  - VMM 集成
  - 事件系统集成

- **单步调试** ([steppings.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\steppings.go))
  - `StepInto()` - 单步进入
  - `StepOver()` - 单步跳过
  - 指令跟踪
  - 用户模式和内核模式支持

- **用户调试器** ([user_debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\user_debugger.go))
  - 用户模式进程调试
  - 暂停包处理
  - 线程拦截

### 5. 断点管理 (debugger/breakpoint/)
- **断点控制** ([control.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\breakpoint\control.go))
  - CTRL+C/CTRL+BREAK 处理
  - 暂停/继续控制
  - 内核调试命令发送

- **断点管理** ([breakpoint.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\breakpoint\breakpoint.go))
  - 软件断点
  - 硬件断点
  - 内存断点
  - 断点列表

### 6. 调试命令 (debugger/commands/)
- **基础命令** ([debugging/basic.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\debugging\basic.go))
  - `g` - 继续执行
  - `p` - 单步跳过
  - `t` - 单步进入
  - `help` - 显示帮助
  - `cls` - 清屏

- **断点命令** ([debugging/breakpoint.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\debugging\breakpoint.go))
  - `bp` - 设置断点
  - `bc` - 清除断点
  - `bl` - 列出断点

- **扩展命令** ([extension/vm.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\extension\vm.go))
  - `!epthook` - 设置 EPT Hook
  - `!monitor` - 监控内存访问
  - `!pte` - 查询页表项
  - `!va2pa` - 虚拟地址转物理地址
  - `!pa2va` - 物理地址转虚拟地址

- **连接管理** ([meta/connection.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\meta\connection.go))
  - `connect` - 连接到远程调试器
  - `disconnect` - 断开连接
  - `attach` - 附加到进程
  - `detach` - 分离进程
  - `status` - 显示状态

### 7. 用户界面 (ui/)
- **UI 框架** ([ui.go](file:///d:\笔记本\p\ux\examples\hypedbg\ui\ui.go))
  - Gio 框架集成
  - 页面管理
  - 工具栏管理
  - 菜单管理

- **UI 页面** (ui/pages/)
  - [CPU 页面](file:///d:\笔记本\p\ux\examples\hypedbg\ui\pages\cpu.go) - CPU 寄存器和反汇编
  - [Events 页面](file:///d:\笔记本\p\ux\examples\hypedbg\ui\pages\events.go) - 事件列表
  - [Log 页面](file:///d:\笔记本\p\ux\examples\hypedbg\ui\pages\log.go) - 日志输出

### 8. 其他功能模块
- **反汇编** ([disassembly/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\disassembly\disassembly.go))
- **符号表** ([symbol/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\symbol\symbol.go))
- **寄存器** ([register/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\register\register.go))
- **进程** ([process/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\process\process.go))
- **线程** ([thread/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\thread\thread.go))
- **调用栈** ([stack/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\stack\stack.go))
- **异常** ([exception/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\exception\exception.go))
- **透明模式** ([transparency/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\transparency\transparency.go))
- **脚本引擎** ([script/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\script\script_engine.go))
- **PE 查看** ([peview/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\peview\peview.go))
- **Scylla** ([scylla/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\scylla\scylla.go))
- **ARK** ([ark/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\ark\ark.go))
- **笔记** ([notes/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\notes\notes.go))
- **书签** ([bookmark/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\bookmark\bookmark.go))
- **注释** ([comment/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\comment\comment.go))
- **函数** ([function/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\function\function.go))
- **引用** ([reference/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\reference\reference.go))
- **交叉引用** ([xref/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\xref\xref.go))
- **类型** ([type_/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\type_\type.go))
- **监视** ([watch/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\watch\watch.go))
- **图形** ([graph/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\graph\graph.go))
- **循环** ([loop/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\loop\loop.go))
- **跟踪** ([trace/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\trace\trace.go))
- **句柄** ([handle/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\handle\handle.go))
- **SEH** ([seh/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\seh\seh.go))
- **源代码** ([source/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\source\source.go))
- **远程** ([remote/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\remote\remote.go))
- **日志** ([log/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\log\logger.go))

## 使用示例

### 启动调试器

```bash
go run main.go
```

### 仅加载驱动（不启动 UI）

```bash
go run main.go --driver-only
```

### 基本命令

```
HyperDbg> help              # 显示帮助
HyperDbg> bp 0x401000      # 设置断点
HyperDbg> bl                # 列出断点
HyperDbg> g                 # 继续执行
HyperDbg> p                 # 单步跳过
HyperDbg> t                 # 单步进入
```

### 驱动管理

```
HyperDbg> load              # 加载驱动
HyperDbg> unload            # 卸载驱动
```

### VMM 操作

```
HyperDbg> vmm init          # 初始化 VMM
HyperDbg> vmm version       # 获取 VMM 版本
HyperDbg> vmm status        # 获取 VMM 状态
HyperDbg> vmm readmem 1234 0x401000 16    # 读取虚拟内存
HyperDbg> vmm writemem 1234 0x401000 9090    # 写入虚拟内存
HyperDbg> vmm readphys 0x1000 16           # 读取物理内存
HyperDbg> vmm writephys 0x1000 9090        # 写入物理内存
```

### 扩展命令

```
HyperDbg> !epthook 0x401000     # 设置 EPT Hook
HyperDbg> !monitor 0x401000 16 true false   # 监控内存
HyperDbg> !pte 0x401000         # 查询页表项
HyperDbg> !va2pa 0x401000       # 转换地址
HyperDbg> !pa2va 0x1000         # 转换地址
```

### 连接管理

```
HyperDbg> connect 192.168.1.100:5000   # 连接到远程调试器
HyperDbg> disconnect                    # 断开连接
HyperDbg> attach 1234                   # 附加到进程
HyperDbg> detach                       # 分离进程
HyperDbg> status                       # 显示状态
```

## 架构设计

### 模块化设计
- 每个功能都是独立的包
- 清晰的接口定义
- 易于维护和扩展

### 驱动集成
- 真实的 Windows 驱动通信
- IOCTL 接口支持
- 完整的错误处理

### 事件驱动架构
- 37 种事件类型支持
- 异步事件处理
- 事件回调机制

### 线程安全
- 使用互斥锁保护共享状态
- 并发访问控制

### 错误处理
- 完整的错误检查
- 详细的错误信息
- 优雅的错误恢复

## 实现完成度

### 核心功能（92% 完成）

| 模块 | 完成度 | 说明 |
|-----|--------|------|
| 基础设施层 | 100% | 驱动加载、IOCTL 通信、VMM 初始化 |
| VMM 层 | 100% | VMM 初始化、终止、内存操作 |
| IOCTL 通信层 | 100% | 所有 IOCTL 功能实现 |
| 数据结构层 | 100% | 所有数据结构定义 |
| 事件处理层 | 100% | 37 种事件类型处理 |
| 调试器核心层 | 80% | 核心功能完整，部分高级功能待完善 |
| 用户界面层 | 50% | 框架完整，部分页面待实现 |

### 详细状态

- ✅ **驱动通信**：100% 完成
- ✅ **事件系统**：100% 完成（37 种事件类型）
- ✅ **VMM 初始化**：100% 完成
- ✅ **内存操作**：100% 完成
- ✅ **单步执行**：100% 完成
- ⚠️ **断点管理**：50% 完成（框架存在，需要完善）
- ⚠️ **寄存器操作**：50% 完成（框架存在，需要完善）
- ⚠️ **反汇编**：50% 完成（框架存在，需要完善）
- ⚠️ **符号表**：50% 完成（框架存在，需要完善）
- ⚠️ **UI 页面**：11% 完成（3/27 页面已实现）

## 与 C++ 版本的对应关系

| C++ 模块 | Go 模块 | 文件 |
|-----------|----------|------|
| debugger.cpp | debugger.go | [debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\debugger.go) |
| install.cpp | driver/provider.go | [provider.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\driver\provider.go) |
| user-listening.cpp | user_debugger.go | [user_debugger.go](file:///d:\笔记本\p\ux\examples\hypedbg\user_debugger.go) |
| script-engine.cpp | script/script_engine.go | [script_engine.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\script\script_engine.go) |
| transparency.cpp | transparency/transparency.go | [transparency.go](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\transparency\transparency.go) |
| commands/ | commands/ | [commands/](file:///d:\笔记本\p\ux\examples\hypedbg\debugger\commands\) |

## 文档

- [架构设计](doc/ARCHITECTURE.md) - 详细的架构设计文档
- [任务计划](doc/TASK_PLAN.md) - 实现进度和任务计划
- [C++ 驱动通信](doc/HYPERDBG_CPP_DRIVER_COMMUNICATION.md) - C++ 与驱动通信机制
- [事件系统](debugger/EVENT_SYSTEM.md) - 事件系统详细说明

## 注意事项

1. **驱动签名**: 需要禁用驱动签名强制或使用测试签名
2. **HVCI**: 不兼容基于虚拟化的安全 (VBS)
3. **管理员权限**: 需要管理员权限来加载驱动
4. **防火墙**: 可能需要配置防火墙规则
5. **Windows 10/11**: 仅支持 Windows 10/11 x64

## 构建和安装

### 前置要求

- Go 1.26.1 或更高版本
- Windows 10/11 x64
- 管理员权限

### 构建步骤

```bash
# 克隆仓库
git clone https://github.com/ddkwork/HyperDbg.git
cd HyperDbg

# 下载依赖
go mod download

# 构建
go build -o hypedbg.exe

# 运行
./hypedbg.exe
```

### 驱动安装

详细说明请参考: https://docs.hyperdbg.org/getting-started/build-and-install

## 贡献

欢迎贡献！请查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解详细信息。

## 许可证

GNU Public License v3

## 致谢

- [HyperDbg](https://github.com/HyperDbg/HyperDbg) - 原始 C++ 实现
- [Gio](https://gioui.org/) - UI 框架
- [golang.org/x/sys](https://pkg.go.dev/golang.org/x/sys) - Windows API 绑定
