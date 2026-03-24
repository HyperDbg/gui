# 使用 EWDK 构建 Rust 驱动
# 使用方法：在 PowerShell 中以管理员身份运行此脚本
# .\build-driver-ewdk.ps1

$ErrorActionPreference = "Stop"

# EWDK 配置
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$EWDK_BUILD_ENV = "$EWDK_MOUNT_LETTER\BuildEnv\SetupBuildEnv.cmd"

# 获取脚本所在目录
$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Rust 驱动构建脚本 (EWDK 环境)" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "脚本目录: $SCRIPT_DIR" -ForegroundColor Green
Write-Host ""

# 检查 Rust
Write-Host "[1/5] 检查 Rust..." -ForegroundColor Yellow

try {
    $rustcVersion = & rustc --version 2>&1
    Write-Host "  ✓ $rustcVersion" -ForegroundColor Green
} catch {
    Write-Host "  错误: 未找到 Rust，请先运行 install-rust-mingw.ps1" -ForegroundColor Red
    exit 1
}

try {
    $cargoVersion = & cargo --version 2>&1
    Write-Host "  ✓ $cargoVersion" -ForegroundColor Green
} catch {
    Write-Host "  错误: 未找到 Cargo" -ForegroundColor Red
    exit 1
}

Write-Host ""

# 检查 EWDK ISO
Write-Host "[2/5] 检查 EWDK..." -ForegroundColor Yellow

if (-not (Test-Path $EWDK_ISO_PATH)) {
    Write-Host "  错误: 未找到 EWDK ISO: $EWDK_ISO_PATH" -ForegroundColor Red
    exit 1
}

Write-Host "  ✓ EWDK ISO 存在" -ForegroundColor Green
Write-Host "  路径: $EWDK_ISO_PATH" -ForegroundColor Gray
Write-Host ""

# 检查是否已挂载
Write-Host "[3/5] 检查 EWDK 挂载..." -ForegroundColor Yellow

$mounted = $false
try {
    $volume = Get-Volume -DriveLetter $EWDK_MOUNT_LETTER.Replace(":", "") -ErrorAction SilentlyContinue
    if ($volume -and $volume.DriveType -eq "CD-ROM") {
        Write-Host "  ✓ EWDK 已挂载到 $EWDK_MOUNT_LETTER" -ForegroundColor Green
        $mounted = $true
    }
} catch {}

if (-not $mounted) {
    Write-Host "  正在挂载 EWDK ISO..." -ForegroundColor Yellow
    try {
        $diskImage = Mount-DiskImage -ImagePath $EWDK_ISO_PATH -PassThru
        $driveLetter = ($diskImage | Get-Volume).DriveLetter
        Write-Host "  ✓ EWDK 已挂载到 ${driveLetter}:" -ForegroundColor Green
    } catch {
        Write-Host "  错误: 挂载失败 - $_" -ForegroundColor Red
        exit 1
    }
}

Write-Host ""

# 检查构建环境脚本
Write-Host "[4/5] 检查构建环境..." -ForegroundColor Yellow

if (-not (Test-Path $EWDK_BUILD_ENV)) {
    Write-Host "  错误: 未找到 EWDK 构建环境: $EWDK_BUILD_ENV" -ForegroundColor Red
    exit 1
}

Write-Host "  ✓ 构建环境: $EWDK_BUILD_ENV" -ForegroundColor Green
Write-Host ""

# 构建驱动
Write-Host "[5/5] 构建驱动..." -ForegroundColor Yellow

Push-Location $SCRIPT_DIR

$buildConfig = "Release"
if ($args -contains "--debug") {
    $buildConfig = "Debug"
}

Write-Host "  配置: $buildConfig" -ForegroundColor Gray
Write-Host "  目录: $SCRIPT_DIR" -ForegroundColor Gray
Write-Host ""

Write-Host "  执行: EWDK 环境中运行 cargo build" -ForegroundColor Gray

$buildCmd = "cmd /c `"$EWDK_BUILD_ENV` && cargo build --$($buildConfig.ToLower())"

try {
    Invoke-Expression $buildCmd
    if ($LASTEXITCODE -ne 0) {
        throw "构建失败"
    }
    Write-Host ""
    Write-Host "  ✓ 构建成功" -ForegroundColor Green
} catch {
    Write-Host ""
    Write-Host "  错误: 构建失败 - $_" -ForegroundColor Red
    Pop-Location
    exit 1
}

Pop-Location

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "构建完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 显示生成的文件
$targetDir = Join-Path $SCRIPT_DIR "target\x86_64-pc-windows-gnu\$buildConfig"
$sysFiles = Get-ChildItem -Path "$targetDir\*.sys" -ErrorAction SilentlyContinue
$pdbFiles = Get-ChildItem -Path "$targetDir\*.pdb" -ErrorAction SilentlyContinue

if ($sysFiles) {
    Write-Host "生成的驱动文件:" -ForegroundColor Yellow
    $sysFiles | ForEach-Object {
        $size = [math]::Round($_.Length / 1KB, 2)
        Write-Host "  $($_.Name) ($size KB)" -ForegroundColor White
    }
}

if ($pdbFiles) {
    Write-Host "生成的符号文件:" -ForegroundColor Yellow
    $pdbFiles | ForEach-Object {
        $size = [math]::Round($_.Length / 1KB, 2)
        Write-Host "  $($_.Name) ($size KB)" -ForegroundColor White
    }
}

Write-Host ""
Write-Host "输出目录: $targetDir" -ForegroundColor Cyan
Write-Host ""
