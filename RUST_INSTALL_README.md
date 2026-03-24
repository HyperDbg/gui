# Rust 安装脚本（MinGW/GNU 工具链）

这个脚本用于在 Windows 上安装 Rust，使用 MinGW/GNU 工具链，无需安装 Visual Studio。

## 前置要求

1. **LLVM/Clang** - 用于编译 Rust 标准库
   - 下载: https://github.com/llvm/llvm-project/releases
   - 安装到: `C:\Program Files\LLVM`

2. **MinGW/GCC** - GNU 工具链
   - 推荐: TDM-GCC (https://jmeubank.github.io/tdm-gcc/)
   - 安装到: `C:\TDM-GCC-64`

3. **CMake** - 构建工具
   - 下载: https://cmake.org/download/
   - 安装到: `C:\Program Files\CMake`

4. **EWDK** - Enterprise Windows Driver Kit（驱动开发需要）
   - ISO 路径: `D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso`
   - 挂载盘符: `E:`
   - 构建环境: `E:\BuildEnv\SetupBuildEnv.cmd`

## 使用方法

### 1. 安装 Rust

#### PowerShell 脚本

```powershell
# 以管理员身份运行
.\install-rust-mingw.ps1
```

### 2. 构建 Rust 驱动

使用 EWDK 环境构建驱动：

```powershell
# 以管理员身份运行
.\build-driver-ewdk.ps1

# 或者构建 Debug 版本
.\build-driver-ewdk.ps1 --debug
```

## 安装内容

- **Rust 工具链**: stable 版本
- **目标平台**: x86_64-pc-windows-gnu
- **Cargo 配置**: 自动配置 MinGW 链接器

## 安装后配置

### 1. 添加到 PATH

将以下路径添加到系统环境变量 PATH：

```
%USERPROFILE%\.cargo\bin
```

### 2. 验证安装

```powershell
rustc --version
cargo --version
```

### 3. 编译 Rust 驱动

对于 Windows 驱动开发，需要使用 **EWDK 环境**：

1. 打开 "Enterprise WDK Command Prompt"
2. 在该环境中运行:
   ```powershell
   cargo build --release
   ```

## 故障排除

### 问题: rustup 提示需要 VS

**解决**: 使用 GNU 工具链
```powershell
rustup target add x86_64-pc-windows-gnu
rustup default stable-x86_64-pc-windows-gnu
```

### 问题: 链接错误

**解决**: 检查 Cargo 配置
```powershell
cat $env:USERPROFILE\.cargo\config.toml
```

确保包含:
```toml
[target.x86_64-pc-windows-gnu]
linker = "x86_64-w64-mingw32-gcc"
ar = "x86_64-w64-mingw32-gcc-ar"
```

### 问题: 编译驱动时找不到 WDK

**解决**: 使用 EWDK 命令提示符

```cmd
# 打开 EWDK 命令提示符
# 然后运行
cargo build
```

## 卸载 Rust

```powershell
rustup self uninstall
```

## 相关资源

- Rust 官网: https://www.rust-lang.org/
- Rustup 文档: https://rust-lang.github.io/rustup/
- Windows Drivers Rust: https://github.com/microsoft/windows-drivers-rs
