# Unit Tests - EWDK Environment

$ErrorActionPreference = "Stop"

$SCRIPT_DIR = Split-Path -Parent $MyInvocation.MyCommand.Path
Set-Location $SCRIPT_DIR

$WDK_PATH = "E:\Program Files\Windows Kits\10"
$WDK_VERSION = "10.0.28000.0"
$MSVC_PATH = "E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207"

Write-Host "=== Unit Tests ===" -ForegroundColor Cyan

# 设置环境变量
$env:Path = "$WDK_PATH\bin\$WDK_VERSION\x64;$MSVC_PATH\bin\HostX64\x64;$env:Path"
$env:LIB = "$WDK_PATH\Lib\$WDK_VERSION\um\x64;$WDK_PATH\Lib\$WDK_VERSION\ucrt\x64;$MSVC_PATH\lib\x64"
$env:INCLUDE = "$WDK_PATH\Include\$WDK_VERSION\um;$WDK_PATH\Include\$WDK_VERSION\ucrt;$MSVC_PATH\include"

Set-Location tests
cargo test -- --nocapture
$exitCode = $LASTEXITCODE
Set-Location ..

exit $exitCode
