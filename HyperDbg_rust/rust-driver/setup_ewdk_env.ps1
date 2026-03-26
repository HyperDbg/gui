# EWDK 环境配置脚本 - 设置系统环境变量
# 需要以管理员权限运行

$ErrorActionPreference = "Stop"

Write-Host "=== EWDK 环境配置脚本 ===" -ForegroundColor Cyan
Write-Host ""

# 检查管理员权限
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $isAdmin) {
    Write-Error "请以管理员权限运行此脚本"
    exit 1
}

# EWDK 配置
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$WDK_VERSION = "10.0.28000.0"
$MSVC_PATH = "E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207"

# 挂载 EWDK
Write-Host "挂载 EWDK ISO..." -ForegroundColor Yellow
if (-not (Test-Path $EWDK_MOUNT_LETTER)) {
    Mount-DiskImage -ImagePath $EWDK_ISO_PATH -StorageType ISO
    Write-Host "EWDK 已挂载到 $EWDK_MOUNT_LETTER" -ForegroundColor Green
} else {
    Write-Host "EWDK 已经挂载" -ForegroundColor Green
}

# 设置 WDK 注册表
Write-Host "配置 WDK 注册表..." -ForegroundColor Yellow

$regPath = "HKLM:\SOFTWARE\Microsoft\Windows Kits\Installed Roots"
if (-not (Test-Path $regPath)) {
    New-Item -Path $regPath -Force | Out-Null
}
Set-ItemProperty -Path $regPath -Name "KitsRoot10" -Value "$WDK_PATH\" -Force

$regPathWow = "HKLM:\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots"
if (-not (Test-Path $regPathWow)) {
    New-Item -Path $regPathWow -Force | Out-Null
}
Set-ItemProperty -Path $regPathWow -Name "KitsRoot10" -Value "$WDK_PATH\" -Force

Write-Host "WDK 注册表配置完成" -ForegroundColor Green

# 设置系统环境变量
Write-Host "配置系统环境变量..." -ForegroundColor Yellow

# PATH
$wdkBin = "$WDK_PATH\bin\$WDK_VERSION\x64"
$msvcBin = "$MSVC_PATH\bin\HostX64\x64"
$ewdkBin = "$EWDK_MOUNT_LETTER\BuildEnv\Scripts"

$currentPath = [Environment]::GetEnvironmentVariable("Path", "Machine")
$pathParts = $currentPath -split ";"

$pathsToAdd = @($wdkBin, $msvcBin, $ewdkBin)
foreach ($p in $pathsToAdd) {
    if ($pathParts -notcontains $p) {
        $currentPath = "$p;$currentPath"
    }
}
[Environment]::SetEnvironmentVariable("Path", $currentPath, "Machine")

# LIB
$libPaths = @(
    "$WDK_PATH\Lib\$WDK_VERSION\um\x64",
    "$WDK_PATH\Lib\$WDK_VERSION\ucrt\x64",
    "$MSVC_PATH\lib\x64"
)
$libValue = $libPaths -join ";"
[Environment]::SetEnvironmentVariable("LIB", $libValue, "Machine")

# INCLUDE
$includePaths = @(
    "$WDK_PATH\Include\$WDK_VERSION\um",
    "$WDK_PATH\Include\$WDK_VERSION\ucrt",
    "$MSVC_PATH\include"
)
$includeValue = $includePaths -join ";"
[Environment]::SetEnvironmentVariable("INCLUDE", $includeValue, "Machine")

# RUSTFLAGS
[Environment]::SetEnvironmentVariable("RUSTFLAGS", "-C target-feature=+crt-static", "Machine")

# WDKContentRoot
[Environment]::SetEnvironmentVariable("WDKContentRoot", $WDK_PATH, "Machine")

Write-Host "系统环境变量配置完成" -ForegroundColor Green

# 创建开机自动挂载任务
Write-Host "创建开机自动挂载 EWDK 任务..." -ForegroundColor Yellow

$taskName = "MountEWDK"
$taskExists = Get-ScheduledTask -TaskName $taskName -ErrorAction SilentlyContinue

if ($taskExists) {
    Unregister-ScheduledTask -TaskName $taskName -Confirm:$false
}

$action = New-ScheduledTaskAction -Execute "PowerShell.exe" -Argument "-NoProfile -ExecutionPolicy Bypass -Command `"Mount-DiskImage -ImagePath '$EWDK_ISO_PATH' -StorageType ISO`""
$trigger = New-ScheduledTaskTrigger -AtStartup
$principal = New-ScheduledTaskPrincipal -UserId "SYSTEM" -LogonType ServiceAccount -RunLevel Highest
$settings = New-ScheduledTaskSettingsSet -AllowStartIfOnBatteries -DontStopIfGoingOnBatteries -StartWhenAvailable

Register-ScheduledTask -TaskName $taskName -Action $action -Trigger $trigger -Principal $principal -Settings $settings -Force | Out-Null

Write-Host "开机自动挂载任务已创建" -ForegroundColor Green

# 更新当前会话的环境变量
Write-Host "更新当前会话环境变量..." -ForegroundColor Yellow

$env:Path = "$wdkBin;$msvcBin;$ewdkBin;$env:Path"
$env:LIB = $libValue
$env:INCLUDE = $includeValue
$env:RUSTFLAGS = "-C target-feature=+crt-static"
$env:WDKContentRoot = $WDK_PATH

Write-Host ""
Write-Host "=== 配置完成 ===" -ForegroundColor Green
Write-Host ""
Write-Host "已配置的环境变量:" -ForegroundColor Cyan
Write-Host "  PATH        : 已添加 WDK/MSVC/EWDK 路径" -ForegroundColor White
Write-Host "  LIB         : $libValue" -ForegroundColor White
Write-Host "  INCLUDE     : $includeValue" -ForegroundColor White
Write-Host "  RUSTFLAGS   : -C target-feature=+crt-static" -ForegroundColor White
Write-Host "  WDKContentRoot : $WDK_PATH" -ForegroundColor White
Write-Host ""
Write-Host "开机自动挂载 EWDK 任务已创建" -ForegroundColor Yellow
Write-Host "现在可以直接使用 cargo build/test 等命令" -ForegroundColor Green
