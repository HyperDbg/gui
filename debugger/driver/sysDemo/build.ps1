# HyperDbg 统一工程构建脚本

$ErrorActionPreference = "Stop"

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
$BUILD_DIR = Join-Path $SCRIPT_DIR "build"
$CONFIG = "Release"
$WDK_PATH = "E:\Program Files\Windows Kits\10"
$EWDK_BUILD_ENV = "E:\BuildEnv\SetupBuildEnv.cmd"

Write-Host "=== HyperDbg 统一工程构建脚本 ===" -ForegroundColor Cyan
Write-Host ""
Write-Host "脚本目录: $SCRIPT_DIR" -ForegroundColor Green
Write-Host ""

# 设置WDK环境变量
$env:WDKContentRoot = $WDK_PATH
Write-Host "设置WDKContentRoot: $env:WDKContentRoot" -ForegroundColor Green
Write-Host ""

# 检查EWDK构建环境脚本是否存在
if (-not (Test-Path $EWDK_BUILD_ENV)) {
    Write-Host "错误: 未找到EWDK构建环境脚本: $EWDK_BUILD_ENV" -ForegroundColor Red
    Read-Host "按任意键退出"
    exit 1
}

Write-Host "使用EWDK构建环境: $EWDK_BUILD_ENV" -ForegroundColor Green
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
Write-Host "输出目录: $BUILD_DIR\$CONFIG" -ForegroundColor Yellow
Write-Host ""

# 显示生成的文件
$sysFiles = Get-ChildItem -Path "$BUILD_DIR\$CONFIG\*.sys" -ErrorAction SilentlyContinue
$pdbFiles = Get-ChildItem -Path "$BUILD_DIR\$CONFIG\*.pdb" -ErrorAction SilentlyContinue

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
