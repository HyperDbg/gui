# Rust 驱动统一工程构建脚本

$ErrorActionPreference = "Stop"

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $SCRIPT_DIR
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$EWDK_BUILD_ENV = "$EWDK_MOUNT_LETTER\BuildEnv\SetupBuildEnv.cmd"

Write-Host "=== Rust 驱动统一工程构建脚本 ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "脚本目录: $SCRIPT_DIR" -ForegroundColor Green
Write-Host ""

# ============================================
# 注册表配置
# ============================================
Write-Host "配置注册表路径..." -ForegroundColor Yellow

# 检查并更新 WDK 安装根目录
$regPath = "HKLM:\SOFTWARE\Microsoft\Windows Kits\Installed Roots"
$regValue = Get-ItemProperty -Path $regPath -Name "KitsRoot10" -ErrorAction SilentlyContinue

if ($null -eq $regValue) {
    Write-Host "添加 WDK KitsRoot10 注册表项..." -ForegroundColor Yellow
    New-ItemProperty -Path $regPath -Name "KitsRoot10" -Value "$WDK_PATH\\"" -PropertyType String -Force | Out-Null
    Write-Host "KitsRoot10 注册表项添加成功" -ForegroundColor Green
} else {
    Write-Host "更新 WDK KitsRoot10 注册表项..." -ForegroundColor Yellow
    Set-ItemProperty -Path $regPath -Name "KitsRoot10" -Value "$WDK_PATH\\"" -Force | Out-Null
    Write-Host "KitsRoot10 注册表项更新成功" -ForegroundColor Green
}

# 检查并更新 WOW6432Node WDK 安装根目录
$regPathWow = "HKLM:\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots"
$regValueWow = Get-ItemProperty -Path $regPathWow -Name "KitsRoot10" -ErrorAction SilentlyContinue

if ($null -eq $regValueWow) {
    Write-Host "添加 WOW6432Node KitsRoot10 注册表项..." -ForegroundColor Yellow
    New-ItemProperty -Path $regPathWow -Name "KitsRoot10" -Value "$WDK_PATH\\"" -PropertyType String -Force | Out-Null
    Write-Host "WOW6432Node KitsRoot10 注册表项添加成功" -ForegroundColor Green
} else {
    Write-Host "更新 WOW6432Node KitsRoot10 注册表项..." -ForegroundColor Yellow
    Set-ItemProperty -Path $regPathWow -Name "KitsRoot10" -Value "$WDK_PATH\\"" -Force | Out-Null
    Write-Host "WOW6432Node KitsRoot10 注册表项更新成功" -ForegroundColor Green
}

# 验证注册表项
Write-Host "验证注册表配置..." -ForegroundColor Yellow
$regValue = Get-ItemProperty -Path $regPath -Name "KitsRoot10" -ErrorAction SilentlyContinue
$regValueWow = Get-ItemProperty -Path $regPathWow -Name "KitsRoot10" -ErrorAction SilentlyContinue

if ($null -ne $regValue -and $null -ne $regValueWow) {
    Write-Host "注册表配置完成" -ForegroundColor Green
} else {
    Write-Error "注册表配置失败"
    exit 1
}
Write-Host ""

# ============================================
# 设置 EWDK 构建环境
# ============================================
Write-Host "设置 EWDK 构建环境..." -ForegroundColor Yellow

# 检查 EWDK ISO 是否已挂载
if (-not (Test-Path $EWDK_MOUNT_LETTER)) {
    Write-Host "EWDK ISO 未挂载，正在挂载..." -ForegroundColor Yellow
    Mount-DiskImage -ImagePath $EWDK_ISO_PATH -StorageType ISO -ErrorAction SilentlyContinue
    if ($?) {
        Write-Host "EWDK ISO 挂载成功到 $EWDK_MOUNT_LETTER" -ForegroundColor Green
    } else {
        Write-Warning "EWDK ISO 挂载失败，尝试使用已挂载的驱动器"
    }
} else {
    Write-Host "EWDK ISO 已挂载到 $EWDK_MOUNT_LETTER" -ForegroundColor Green
}

# 检查 EWDK 构建环境脚本
if (-not (Test-Path $EWDK_BUILD_ENV)) {
    Write-Error "EWDK 构建环境脚本不存在: $EWDK_BUILD_ENV"
    exit 1
}

# 执行 EWDK 构建环境脚本
Write-Host "使用 EWDK 构建环境: $EWDK_BUILD_ENV" -ForegroundColor Yellow
& $EWDK_BUILD_ENV | Out-Null

# 设置 WDK 环境变量
$WDK_VERSION = "10.0.28000.0"
$WDK_BIN = "$WDK_PATH\bin\$WDK_VERSION"
$MSVC_BIN = "E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\bin\HostX64\x64"

Write-Host "设置 WDKContentRoot: $WDK_PATH" -ForegroundColor Yellow
$env:WDKContentRoot = $WDK_PATH

# 检测 WDK 工具链版本
if (Test-Path "$WDK_BIN\x64") {
    Write-Host "使用 WDK 版本: $WDK_VERSION" -ForegroundColor Yellow
    Write-Host "使用 WDK 工具链: $WDK_BIN\x64" -ForegroundColor Yellow
    $env:PATH = "$WDK_BIN\x64;$env:PATH"
} else {
    Write-Error "WDK 工具链不存在: $WDK_BIN\x64"
    exit 1
}

# 设置 MSVC 链接器
if (Test-Path $MSVC_BIN) {
    Write-Host "使用 MSVC 链接器: $MSVC_BIN" -ForegroundColor Yellow
    $env:PATH = "$MSVC_BIN;$env:PATH"
} else {
    Write-Warning "MSVC 链接器不存在: $MSVC_BIN"
}

Write-Host ""

# ============================================
# 配置构建环境
# ============================================
Write-Host "配置构建环境..." -ForegroundColor Yellow

$env:LIB = "$WDK_PATH\Lib\$WDK_VERSION\um\x64;$WDK_PATH\Lib\$WDK_VERSION\ucrt\x64"
$env:UCRT = "$WDK_PATH\Lib\$WDK_VERSION\ucrt\x64"
$env:MSVC_LIB = "E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\lib\x64"

Write-Host "设置 LIB: $env:LIB" -ForegroundColor Yellow
Write-Host "设置 UCRT: $env:UCRT" -ForegroundColor Yellow
Write-Host "设置 MSVC LIB: $env:MSVC_LIB" -ForegroundColor Yellow

# 设置 Rust 构建标志
$env:RUSTFLAGS = "-C target-feature=+crt-static"
$env:CARGO_CFG_TARGET_FEATURE = "cmpxchg16b,crt-static,fxsr,sse,sse2,sse3"

Write-Host "设置 RUSTFLAGS: $env:RUSTFLAGS" -ForegroundColor Yellow
Write-Host "设置 CARGO_CFG_TARGET_FEATURE: $env:CARGO_CFG_TARGET_FEATURE" -ForegroundColor Yellow

# 设置 MSVC 库路径
$env:LIB = "$env:LIB;$env:MSVC_LIB"
Write-Host "设置完整 LIB: $env:LIB" -ForegroundColor Yellow
Write-Host ""

# ============================================
# 自动检测并编译所有项目
# ============================================
Write-Host "=== 检测项目 ===" -ForegroundColor Cyan

# 获取所有包含 Cargo.toml 的子目录（排除特定目录）
$excludeDirs = @(".git", "target", ".cargo")
$projects = Get-ChildItem -Directory | Where-Object {
    -not ($excludeDirs -contains $_.Name) -and (Test-Path "$($_.FullName)\Cargo.toml")
}

if ($projects.Count -eq 0) {
    Write-Warning "未找到任何包含 Cargo.toml 的项目目录"
    exit 0
}

Write-Host "找到 $($projects.Count) 个项目:" -ForegroundColor Green
foreach ($project in $projects) {
    Write-Host "  - $($project.Name)" -ForegroundColor White
}
Write-Host ""

# ============================================
# 编译所有项目
# ============================================
Write-Host "=== 编译项目 ===" -ForegroundColor Cyan

$successCount = 0
$failCount = 0

foreach ($project in $projects) {
    Write-Host ""
    Write-Host "=== 编译 $($project.Name) ===" -ForegroundColor Cyan
    Push-Location $project.Name
    try {
        cargo build --release --target x86_64-pc-windows-msvc
        if ($LASTEXITCODE -eq 0) {
            Write-Host "$($project.Name) 编译成功！" -ForegroundColor Green
            $successCount++
        } else {
            Write-Error "$($project.Name) 编译失败"
            $failCount++
        }
    } catch {
        Write-Error "$($project.Name) 编译失败: $_"
        $failCount++
    }
    Pop-Location
}

Write-Host ""
Write-Host "=== 编译完成 ===" -ForegroundColor Green
Write-Host "成功: $successCount, 失败: $failCount" -ForegroundColor $(if ($failCount -eq 0) { "Green" } else { "Yellow" })
Write-Host ""

# ============================================
# 自动重命名驱动二进制文件
# ============================================
Write-Host "=== 重命名驱动二进制文件 ===" -ForegroundColor Cyan

foreach ($project in $projects) {
    $cargoTomlPath = "$($project.FullName)\Cargo.toml"
    if (Test-Path $cargoTomlPath) {
        $cargoToml = Get-Content $cargoTomlPath -Raw
        if ($cargoToml -match 'name\s*=\s*"([^"]+)"') {
            $packageName = $matches[1]
            $dllPath = "$SCRIPT_DIR\target\x86_64-pc-windows-msvc\release\$packageName.dll"
            $sysPath = "$SCRIPT_DIR\target\x86_64-pc-windows-msvc\release\$packageName.sys"

            if (Test-Path $dllPath) {
                if (Test-Path $sysPath) {
                    try {
                        Remove-Item -Path $sysPath -Force -ErrorAction Stop
                    } catch {
                        Write-Host "无法删除 $packageName.sys，可能正在被使用，跳过重命名" -ForegroundColor Yellow
                        continue
                    }
                }
                Write-Host "重命名 $packageName.dll 为 $packageName.sys" -ForegroundColor Yellow
                Rename-Item -Path $dllPath -NewName "$packageName.sys" -Force
            }
        }
    }
}

Write-Host ""
Write-Host "=== 构建完成 ===" -ForegroundColor Green
Write-Host ""
Write-Host "请等待驱动签名后再进行测试..." -ForegroundColor Yellow
Write-Host ""
