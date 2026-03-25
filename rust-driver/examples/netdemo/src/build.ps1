# HyperDbg 统一工程构建脚本

$ErrorActionPreference = "Stop"

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
$BUILD_DIR = Join-Path $SCRIPT_DIR "build"
$CONFIG = "Release"
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$EWDK_ISO_PATH = "D:\Admin\vs2022VC\EWDK_br_release_28000_251103-1709.iso"
$EWDK_MOUNT_LETTER = "E:"
$EWDK_BUILD_ENV = "$EWDK_MOUNT_LETTER\BuildEnv\SetupBuildEnv.cmd"

Write-Host "=== HyperDbg 统一工程构建脚本 ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "脚本目录: $SCRIPT_DIR" -ForegroundColor Green
Write-Host ""

# 检查ISO文件是否存在
if (-not (Test-Path $EWDK_ISO_PATH)) {
    Write-Host "错误: 未找到EWDK ISO文件: $EWDK_ISO_PATH" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

# 检查是否已挂载
$mounted = $false
try {
    $volume = Get-Volume -DriveLetter $EWDK_MOUNT_LETTER.Replace(":", "") -ErrorAction SilentlyContinue
    if ($volume -and $volume.DriveType -eq "CD-ROM") {
        Write-Host "EWDK ISO 已挂载到 $EWDK_MOUNT_LETTER" -ForegroundColor Green
        $mounted = $true
    }
} catch {}

if (-not $mounted) {
    Write-Host "正在挂载 EWDK ISO..." -ForegroundColor Yellow
    $diskImage = Mount-DiskImage -ImagePath $EWDK_ISO_PATH -PassThru
    $driveLetter = ($diskImage | Get-Volume).DriveLetter
    Write-Host "EWDK ISO 已挂载到 ${driveLetter}:" -ForegroundColor Green
}

# 检查构建环境脚本是否存在
if (-not (Test-Path $EWDK_BUILD_ENV)) {
    Write-Host "错误: 未找到EWDK构建环境脚本: $EWDK_BUILD_ENV" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

Write-Host "使用EWDK构建环境: $EWDK_BUILD_ENV" -ForegroundColor Green
Write-Host ""

# 设置WDK环境变量
$env:WDKContentRoot = $WDK_PATH
Write-Host "设置WDKContentRoot: $env:WDKContentRoot" -ForegroundColor Green
Write-Host ""

# 清理并创建构建目录
if (Test-Path $BUILD_DIR) {
    Write-Host "清理构建目录: $BUILD_DIR" -ForegroundColor Yellow
    Remove-Item -Path $BUILD_DIR -Recurse -Force
}
Write-Host "创建构建目录: $BUILD_DIR" -ForegroundColor Yellow
New-Item -ItemType Directory -Path $BUILD_DIR | Out-Null

# 生成项目
Write-Host "=== 生成项目 ===" -ForegroundColor Cyan
Push-Location $BUILD_DIR
try {
    cmd /c "`"$EWDK_BUILD_ENV`" && cmake `"$SCRIPT_DIR`""
    if ($LASTEXITCODE -ne 0) {
        throw "生成项目失败"
    }
} catch {
    Write-Host "=== 生成项目失败 ===" -ForegroundColor Red
    Pop-Location
    Read-Host "按任意键退出"
    exit 1
}

Write-Host ""
Write-Host "=== 编译项目 ===" -ForegroundColor Cyan
try {
    cmd /c "`"$EWDK_BUILD_ENV`" && cmake --build . --config $CONFIG"
    if ($LASTEXITCODE -ne 0) {
        throw "编译失败"
    }
} catch {
    Write-Host "=== 编译失败 ===" -ForegroundColor Red
    Pop-Location
    Read-Host "按任意键退出"
    exit 1
}

Pop-Location

Write-Host ""
Write-Host "=== 构建完成 ===" -ForegroundColor Green
Write-Host ""

# 复制生成的文件到源代码目录
$sysFile = Join-Path "$BUILD_DIR\$CONFIG" "wsksample.sys"
$pdbFile = Join-Path "$BUILD_DIR\$CONFIG" "wsksample.pdb"

if (Test-Path $sysFile) {
    Copy-Item -Path $sysFile -Destination $SCRIPT_DIR -Force
    Write-Host "复制: wsksample.sys -> $SCRIPT_DIR" -ForegroundColor Yellow
}
if (Test-Path $pdbFile) {
    Copy-Item -Path $pdbFile -Destination $SCRIPT_DIR -Force
    Write-Host "复制: wsksample.pdb -> $SCRIPT_DIR" -ForegroundColor Yellow
}

# 删除构建目录
if (Test-Path $BUILD_DIR) {
    Write-Host "清理构建目录: $BUILD_DIR" -ForegroundColor Yellow
    Remove-Item -Path $BUILD_DIR -Recurse -Force
}

Write-Host ""
Write-Host "输出目录: $SCRIPT_DIR" -ForegroundColor Yellow

# 显示生成的文件
$sysFiles = Get-ChildItem -Path "$SCRIPT_DIR\*.sys" -ErrorAction SilentlyContinue
$pdbFiles = Get-ChildItem -Path "$SCRIPT_DIR\*.pdb" -ErrorAction SilentlyContinue

if ($sysFiles) {
    Write-Host "生成的驱动文件:" -ForegroundColor Yellow
    $sysFiles | ForEach-Object { Write-Host "  $($_.Name)" }
}

if ($pdbFiles) {
    Write-Host "生成的符号文件:" -ForegroundColor Yellow
    $pdbFiles | ForEach-Object { Write-Host "  $($_.Name)" }
}

Write-Host ""
Read-Host "按任意键退出"
