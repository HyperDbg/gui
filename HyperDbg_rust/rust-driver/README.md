# HyperDbg Rust Driver

HyperDbg 的 Rust 驱动实现，使用 windows-drivers-rs 框架。

## 项目结构

```
rust-driver/
├── kd/                     # 主驱动 crate (整合所有模块)
│   ├── src/
│   │   ├── lib.rs          # 驱动入口点
│   │   ├── kd.rs           # 内核调试器
│   │   ├── ud.rs           # 用户调试器
│   │   ├── common/         # 内部通用模块
│   │   ├── logger/         # 内部日志模块
│   │   ├── net/            # 内部网络模块
│   │   ├── framework/      # 内部驱动框架
│   │   ├── disassembler/   # 反汇编器
│   │   ├── pdbex/          # PDB 解析
│   │   └── hyperkd/        # 调试器核心
│   │       └── hyperhv/    # Hypervisor 实现
│   └── build.ps1
│
├── driver-framework/       # 驱动框架库 (独立 crate)
├── logger/                 # 日志模块 (独立 crate)
├── common/                 # 通用模块 (独立 crate)
├── net/                    # 网络模块 (独立 crate)
│
└── examples/               # 示例驱动
    ├── sysdemo/            # WDM 驱动示例
    └── netdemo/            # 网络驱动示例
```

## 构建要求

- Rust 1.85+
- Windows Driver Kit (WDK)
- EWDK (Enterprise Windows Driver Kit)

## 构建步骤

```powershell
# 构建主驱动
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd
powershell -ExecutionPolicy Bypass -File build.ps1

# 构建示例驱动
cd d:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\examples\sysdemo
powershell -ExecutionPolicy Bypass -File build.ps1
```

## 许可证

GPL-3.0
