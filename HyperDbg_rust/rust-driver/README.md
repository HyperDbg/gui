# HyperDbg Rust Driver

HyperDbg 的 Rust 驱动实现，使用 windows-drivers-rs 框架。

## 项目结构

```
rust-driver/
├── hyperhv/          # Hypervisor 层
├── hyperkd/          # Debugger 层
├── protocol/          # 通信协议
├── wsk/              # WSK Socket 实现
├── disassembler/      # 反汇编器
├── pdbex/            # PDB 解析
├── kd/               # KD 接口
└── sysdemo/          # 示例驱动
```

## 构建要求

- Rust 1.85+
- Windows Driver Kit (WDK)
- EWDK (Enterprise Windows Driver Kit)

## 构建步骤

```powershell
# 构建所有驱动
cd d:\ux\examples\hypedbg\rust-driver
cargo build --release

# 构建单个驱动
cargo build -p sysdemo --release
```

## 许可证

GPL-3.0
