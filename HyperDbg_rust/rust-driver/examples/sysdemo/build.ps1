# Rust 驱动构建脚本 - sysdemo

$ErrorActionPreference = "Stop"

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $SCRIPT_DIR
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$EWDK_BUILD_ENV = "$EWDK_MOUNT_LETTER\BuildEnv\SetupBuildEnv.cmd"

Write-Host "=== sysdemo 驱动构建脚本 ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "脚本目录: $SCRIPT_DIR" -ForegroundColor Green
Write-Host ""

# ============================================
# 注册表配置
# ============================================
Write-Host "配置注册表路径..." -ForegroundColor Yellow

$regPath = "HKLM:\SOFTWARE\Microsoft\Windows Kits\Installed Roots"
$regValue = Get-ItemProperty -Path $regPath -Name "KitsRoot10" -ErrorAction SilentlyContinue

if ($null -eq $regValue) {
    New-ItemProperty -Path $regPath -Name "KitsRoot10" -Value "$WDK_PATH\\"" -PropertyType String -Force | Out-Null
} else {
    Set-ItemProperty -Path $regPath -Name "KitsRoot10" -Value "$WDK_PATH\\"" -Force | Out-Null
}

$regPathWow = "HKLM:\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots"
$regValueWow = Get-ItemProperty -Path $regPathWow -Name "KitsRoot10" -ErrorAction SilentlyContinue

if ($null -eq $regValueWow) {
    New-ItemProperty -Path $regPathWow -Name "KitsRoot10" -Value "$WDK_PATH\\"" -PropertyType String -Force | Out-Null
} else {
    Set-ItemProperty -Path $regPathWow -Name "KitsRoot10" -Value "$WDK_PATH\\"" -Force | Out-Null
}

Write-Host "注册表配置完成" -ForegroundColor Green
Write-Host ""

# ============================================
# 设置 EWDK 构建环境
# ============================================
Write-Host "设置 EWDK 构建环境..." -ForegroundColor Yellow

if (-not (Test-Path $EWDK_MOUNT_LETTER)) {
    Mount-DiskImage -ImagePath $EWDK_ISO_PATH -StorageType ISO -ErrorAction SilentlyContinue
}

& $EWDK_BUILD_ENV | Out-Null

$WDK_VERSION = "10.0.28000.0"
$WDK_BIN = "$WDK_PATH\bin\$WDK_VERSION"
$MSVC_BIN = "E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\bin\HostX64\x64"

$env:WDKContentRoot = $WDK_PATH
$env:PATH = "$WDK_BIN\x64;$MSVC_BIN;$env:PATH"

Write-Host ""

# ============================================
# 配置构建环境
# ============================================
Write-Host "配置构建环境..." -ForegroundColor Yellow

$env:LIB = "$WDK_PATH\Lib\$WDK_VERSION\um\x64;$WDK_PATH\Lib\$WDK_VERSION\ucrt\x64;E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\lib\x64"
$env:RUSTFLAGS = "-C target-feature=+crt-static"
$env:CARGO_CFG_TARGET_FEATURE = "cmpxchg16b,crt-static,fxsr,sse,sse2,sse3"

Write-Host ""

# ============================================
# 编译项目
# ============================================
Write-Host "=== 编译 sysdemo ===" -ForegroundColor Cyan

cargo build --release --target x86_64-pc-windows-msvc

if ($LASTEXITCODE -eq 0) {
    Write-Host "sysdemo 编译成功！" -ForegroundColor Green
    
    $workspaceTarget = "d:\ux\examples\hypedbg\rust-driver\target\x86_64-pc-windows-msvc\release"
    $dllPath = "$workspaceTarget\sysdemo.dll"
    $sysPath = "$workspaceTarget\sysdemo.sys"
    $localSysPath = "$SCRIPT_DIR\sysdemo.sys"
    
    if (Test-Path $dllPath) {
        if (Test-Path $sysPath) {
            Remove-Item -Path $sysPath -Force -ErrorAction SilentlyContinue
        }
        Rename-Item -Path $dllPath -NewName "sysdemo.sys" -Force
        Write-Host "重命名 sysdemo.dll 为 sysdemo.sys" -ForegroundColor Yellow
        
        if (Test-Path $localSysPath) {
            Remove-Item -Path $localSysPath -Force -ErrorAction SilentlyContinue
        }
        Copy-Item -Path $sysPath -Destination $localSysPath -Force
        Write-Host "复制 sysdemo.sys 到本地目录" -ForegroundColor Yellow
    }
    
    Write-Host ""
    Write-Host "=== 构建完成 ===" -ForegroundColor Green
    Write-Host "请等待驱动签名后再进行测试..." -ForegroundColor Yellow
} else {
    Write-Error "sysdemo 编译失败"
    exit 1
}
