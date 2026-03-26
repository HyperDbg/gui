# Rust 安装脚本（使用 MinGW/GNU 工具链，无需 VS）
# 使用方法：在 PowerShell 中以管理员身份运行此脚本
# .\install-rust-mingw.ps1

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Rust 安装脚本 (MinGW/GNU 工具链)" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# EWDK 配置（用于驱动开发）
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$EWDK_BUILD_ENV = "$EWDK_MOUNT_LETTER\BuildEnv\SetupBuildEnv.cmd"

Write-Host "EWDK 配置:" -ForegroundColor Yellow
Write-Host "  ISO: $EWDK_ISO_PATH" -ForegroundColor Gray
Write-Host "  挂载: $EWDK_MOUNT_LETTER" -ForegroundColor Gray
Write-Host "  构建环境: $EWDK_BUILD_ENV" -ForegroundColor Gray
Write-Host ""

# 检查 EWDK ISO
if (-not (Test-Path $EWDK_ISO_PATH)) {
    Write-Host "警告: 未找到 EWDK ISO: $EWDK_ISO_PATH" -ForegroundColor Yellow
    Write-Host "  驱动开发需要 EWDK 环境" -ForegroundColor Yellow
    Write-Host ""
}

# 检查必要的工具
Write-Host "[1/7] 检查必要工具..." -ForegroundColor Yellow

$llvmPath = "C:\Program Files\LLVM\bin"
$mingwPath = "C:\TDM-GCC-64\bin"
$cmakePath = "C:\Program Files\CMake\bin"

if (-not (Test-Path "$llvmPath\clang-cl.exe")) {
    Write-Host "错误: 未找到 LLVM/clang-cl.exe" -ForegroundColor Red
    Write-Host "请先安装 LLVM: https://github.com/llvm/llvm-project/releases" -ForegroundColor Red
    exit 1
}

if (-not (Test-Path "$mingwPath\gcc.exe")) {
    Write-Host "错误: 未找到 MinGW/gcc.exe" -ForegroundColor Red
    Write-Host "请先安装 TDM-GCC 或其他 MinGW 发行版" -ForegroundColor Red
    exit 1
}

if (-not (Test-Path "$cmakePath\cmake.exe")) {
    Write-Host "错误: 未找到 CMake" -ForegroundColor Red
    exit 1
}

Write-Host "  ✓ LLVM/clang-cl.exe" -ForegroundColor Green
Write-Host "  ✓ MinGW/GCC" -ForegroundColor Green
Write-Host "  ✓ CMake" -ForegroundColor Green
Write-Host ""

# 清理旧的 Rust 安装
Write-Host "[2/6] 清理旧的 Rust 安装..." -ForegroundColor Yellow

if (Test-Path "$env:USERPROFILE\.rustup") {
    Write-Host "  删除 .rustup 目录..."
    Remove-Item -Recurse -Force "$env:USERPROFILE\.rustup"
}

if (Test-Path "$env:USERPROFILE\.cargo") {
    Write-Host "  删除 .cargo 目录..."
    Remove-Item -Recurse -Force "$env:USERPROFILE\.cargo"
}

if (Test-Path "$env:USERPROFILE\.rustup.rs") {
    Write-Host "  删除 .rustup.rs 目录..."
    Remove-Item -Recurse -Force "$env:USERPROFILE\.rustup.rs"
}

Write-Host "  ✓ 清理完成" -ForegroundColor Green
Write-Host ""

# 设置环境变量
Write-Host "[3/6] 设置环境变量..." -ForegroundColor Yellow

$env:CC = "$mingwPath\gcc.exe"
$env:CXX = "$mingwPath\g++.exe"
$env:AR = "$mingwPath\ar.exe"
$env:LD = "$mingwPath\ld.exe"
$env:RANLIB = "$mingwPath\ranlib.exe"
$env:STRIP = "$mingwPath\strip.exe"

Write-Host "  CC=$env:CC" -ForegroundColor Gray
Write-Host "  CXX=$env:CXX" -ForegroundColor Gray
Write-Host "  AR=$env:AR" -ForegroundColor Gray
Write-Host "  LD=$env:LD" -ForegroundColor Gray
Write-Host ""

# 下载 Rustup
Write-Host "[4/6] 下载 Rustup..." -ForegroundColor Yellow

$rustupUrl = "https://win.rustup.rs/x86_64"
$rustupPath = "$env:TEMP\rustup-init.exe"

if (Test-Path $rustupPath) {
    Write-Host "  使用已下载的 rustup-init.exe" -ForegroundColor Gray
} else {
    Write-Host "  从 $rustupUrl 下载..." -ForegroundColor Gray
    try {
        Invoke-WebRequest -Uri $rustupUrl -OutFile $rustupPath -UseBasicParsing
        Write-Host "  ✓ 下载完成" -ForegroundColor Green
    } catch {
        Write-Host "  错误: 下载失败 - $_" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""

# 安装 Rust（使用 GNU 工具链）
Write-Host "[5/7] 安装 Rust..." -ForegroundColor Yellow

$installArgs = @(
    "--default-toolchain", "stable",
    "--profile", "minimal",
    "--target", "x86_64-pc-windows-gnu",
    "-y"
)

Write-Host "  执行: rustup-init.exe $installArgs" -ForegroundColor Gray

& $rustupPath @installArgs

if ($LASTEXITCODE -ne 0) {
    Write-Host "  错误: Rust 安装失败" -ForegroundColor Red
    exit 1
}

Write-Host "  ✓ Rust 安装完成" -ForegroundColor Green
Write-Host ""

# 配置 Cargo
Write-Host "[6/7] 配置 Cargo..." -ForegroundColor Yellow

$cargoConfigPath = "$env:USERPROFILE\.cargo\config.toml"
$cargoConfigDir = Split-Path -Parent $cargoConfigPath

if (-not (Test-Path $cargoConfigDir)) {
    New-Item -ItemType Directory -Path $cargoConfigDir -Force | Out-Null
}

$cargoConfig = @"
[target.x86_64-pc-windows-gnu]
linker = "x86_64-w64-mingw32-gcc"
ar = "x86_64-w64-mingw32-gcc-ar"

[build]
target = "x86_64-pc-windows-gnu"

[target.'cfg(target_arch = "x86_64")']
rustflags = ["-C", "link-arg=-fuse-ld=lld"]
"@

Set-Content -Path $cargoConfigPath -Value $cargoConfig -Encoding UTF8

Write-Host "  ✓ Cargo 配置完成" -ForegroundColor Green
Write-Host "  配置文件: $cargoConfigPath" -ForegroundColor Gray
Write-Host ""

# 添加到 PATH
Write-Host "[7/7] 配置系统 PATH..." -ForegroundColor Yellow

$cargoBin = "$env:USERPROFILE\.cargo\bin"
$env:PATH = "$cargoBin;$env:PATH"

Write-Host "  请手动将以下路径添加到系统 PATH:" -ForegroundColor Cyan
Write-Host "  $cargoBin" -ForegroundColor White
Write-Host ""

# 验证安装
Write-Host "验证安装..." -ForegroundColor Yellow

& "$cargoBin\rustc.exe" --version
& "$cargoBin\cargo.exe" --version

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "安装完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "下一步操作:" -ForegroundColor Yellow
Write-Host "1. 重启终端或重新登录" -ForegroundColor White
Write-Host "2. 运行: rustc --version" -ForegroundColor White
Write-Host "3. 运行: cargo --version" -ForegroundColor White
Write-Host ""

# EWDK 配置说明
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "驱动开发（需要 EWDK）" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

Write-Host "EWDK ISO 路径: $EWDK_ISO_PATH" -ForegroundColor Yellow
Write-Host "挂载盘符: $EWDK_MOUNT_LETTER" -ForegroundColor Yellow
Write-Host "构建环境: $EWDK_BUILD_ENV" -ForegroundColor Yellow
Write-Host ""

Write-Host "使用 EWDK 构建驱动:" -ForegroundColor Yellow
Write-Host "1. 挂载 EWDK ISO:" -ForegroundColor White
Write-Host "   Mount-DiskImage -ImagePath '$EWDK_ISO_PATH'" -ForegroundColor Gray
Write-Host ""
Write-Host "2. 运行通用构建脚本:" -ForegroundColor White
Write-Host "   .\rust-driver\examples\build-all.ps1" -ForegroundColor Gray
Write-Host ""
