# HyperDbg Rust Driver

HyperDbg 的 Rust 驱动实现，使用 windows-drivers-rs 框架。

## 项目结构

```
rust-driver/
└── kd/                     # 主驱动 crate
    ├── Cargo.toml
    ├── build.ps1           # 构建脚本 (EWDK)
    ├── build.rs            # 构建脚本 (Rust)
    └── src/
        ├── lib.rs          # 驱动入口
        ├── kd.rs           # 内核调试器
        ├── ud.rs           # 用户调试器
        ├── hyperdbg_api.rs # API 实现
        │
        ├── generated/      # 自动生成目录
        │   ├── mod.rs
        │   ├── common.rs
        │   ├── register.rs
        │   ├── event_*.rs  # 事件类型 (22个文件)
        │   ├── router.rs   # API 路由器
        │   ├── emit.rs     # 事件发射器
        │   ├── ioctl.rs    # IOCTL 代码
        │   ├── hook_db.rs  # Hook 数据库
        │   ├── hook_types.rs
        │   ├── hook_args.rs
        │   ├── ept_hook.rs
        │   ├── inline_hook.rs
        │   ├── ntddk.rs    # WDK 函数导出 (1251个)
        │   ├── types.rs    # WDK 类型 (7317个)
        │   └── constants.rs # WDK 常量 (5348个)
        │
        ├── go_script/      # Go 脚本引擎
        │   ├── parser/     # Go 解析器 (10个文件)
        │   ├── analyzer.rs
        │   ├── executor.rs
        │   └── generator.rs
        │
        ├── hyperkd/        # 调试器核心
        │   ├── hyperhv/    # Hypervisor 实现
        │   │   ├── vmm/    # VMX/EPT 核心
        │   │   │   ├── ept/ # EPT 管理
        │   │   │   └── vmx/ # VMX 操作
        │   │   ├── hooks/  # Hook 机制
        │   │   ├── memory/ # 内存管理
        │   │   ├── processor/ # 处理器管理
        │   │   └── interface/ # 接口层
        │   ├── debugger.rs
        │   ├── commands.rs
        │   └── network.rs
        │
        ├── net/            # WSK HTTP Server
        │   ├── http.rs
        │   ├── json.rs
        │   └── util.rs
        │
        ├── framework/      # 驱动框架
        │   ├── device.rs
        │   ├── ioctl.rs
        │   └── utils.rs
        │
        ├── logger/         # 日志模块
        │   └── debug_logger.rs
        │
        ├── disassembler/   # 反汇编器
        └── pdbex/          # PDB 解析器
```

## 构建要求

- Rust 1.85+
- Windows Driver Kit (WDK)
- EWDK (Enterprise Windows Driver Kit)

## 构建步骤

```powershell
# 编译主驱动 (使用 EWDK)
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd
powershell -ExecutionPolicy Bypass -File build.ps1

# 检查编译 (不生成驱动)
cargo check
```

## 代码生成

当修改 Go 端的类型定义时，需要重新生成 Rust 代码：

```bash
cd d:\ux\examples\hypedbg\HyperDbg_rust\cmd\rustgen
go run .
```

## 架构

```
┌─────────────────────────────────────────────────────────────┐
│                   Go 用户空间通信层                          │
│                   (packet.go + event_handlers.go)            │
└────────────────────┬────────────────────────────────────────┘
                     │ HTTP/JSON (WSK Socket :50080)
┌────────────────────▼────────────────────────────────────────┐
│              Rust 内核驱动层                                 │
├──────────────────────────────────────────────────────────────┤
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐       │
│  │   hyperhv    │  │   hyperkd    │  │   pdbex      │       │
│  │ (Hypervisor) │  │  (Debugger)  │  │ (PDB Parser) │       │
│  └──────────────┘  └──────────────┘  └──────────────┘       │
│         │                 │                  │               │
│  ┌──────▼──────────────────▼──────────────────▼─────────┐   │
│  │        generated (types + handlers + hooks)           │   │
│  │   - event_*.rs    : 事件类型定义                       │   │
│  │   - router.rs     : API 路由器                        │   │
│  │   - emit.rs       : 事件发射器                        │   │
│  │   - hook_db.rs    : Hook 数据库                       │   │
│  └───────────────────────────────────────────────────────┘   │
│  ┌───────────────────────────────────────────────────────┐   │
│  │                 net (WSK HTTP Server)                 │   │
│  │   - HTTP 服务器监听 50080 端口                         │   │
│  │   - JSON 序列化/反序列化                               │   │
│  └───────────────────────────────────────────────────────┘   │
└──────────────────────────────────────────────────────────────┘
```

## 许可证

GPL-3.0
