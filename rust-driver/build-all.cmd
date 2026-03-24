@echo off
setlocal enabledelayedexpansion

echo Building all Rust driver workspace...
echo.

rem ============================================
rem Registry Configuration
rem ============================================
echo Configuring registry paths...

rem Check and update WDK installation root
reg query "HKLM\SOFTWARE\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 >nul 2>&1
if errorlevel 1 (
    echo Adding WDK KitsRoot10 registry entry...
    reg add "HKLM\SOFTWARE\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /t REG_SZ /d "E:\Program Files\Windows Kits\10\" /f
    if errorlevel 1 (
        echo ERROR: Failed to add KitsRoot10 registry entry
        exit /b 1
    )
    echo KitsRoot10 registry entry added successfully
) else (
    echo Removing old WDK KitsRoot10 registry entry...
    reg delete "HKLM\SOFTWARE\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /f >nul 2>&1
    echo Adding new WDK KitsRoot10 registry entry...
    reg add "HKLM\SOFTWARE\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /t REG_SZ /d "E:\Program Files\Windows Kits\10\" /f
    if errorlevel 1 (
        echo ERROR: Failed to add KitsRoot10 registry entry
        exit /b 1
    )
    echo KitsRoot10 registry entry updated successfully
)

rem Check and update WOW6432Node WDK installation root
reg query "HKLM\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 >nul 2>&1
if errorlevel 1 (
    echo Adding WOW6432Node KitsRoot10 registry entry...
    reg add "HKLM\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /t REG_SZ /d "E:\Program Files\Windows Kits\10\" /f
    if errorlevel 1 (
        echo WARNING: Failed to add WOW6432Node KitsRoot10 registry entry
    ) else (
        echo WOW6432Node KitsRoot10 registry entry added successfully
    )
) else (
    echo Removing old WOW6432Node KitsRoot10 registry entry...
    reg delete "HKLM\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /f >nul 2>&1
    echo Adding new WOW6432Node KitsRoot10 registry entry...
    reg add "HKLM\SOFTWARE\WOW6432Node\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10 /t REG_SZ /d "E:\Program Files\Windows Kits\10\" /f
    if errorlevel 1 (
        echo WARNING: Failed to add WOW6432Node KitsRoot10 registry entry
    ) else (
        echo WOW6432Node KitsRoot10 registry entry updated successfully
    )
)

rem Verify registry entries
echo Verifying registry configuration...
reg query "HKLM\SOFTWARE\Microsoft\Windows Kits\Installed Roots" /v KitsRoot10
if errorlevel 1 (
    echo ERROR: KitsRoot10 registry verification failed
    exit /b 1
)

echo Registry configuration completed successfully
echo.

rem ============================================
rem EWDK Environment Setup
rem ============================================
echo Setting up EWDK build environment...
call "E:\BuildEnv\SetupBuildEnv.cmd" amd64

if errorlevel 1 (
    echo ERROR: Failed to setup EWDK build environment
    exit /b 1
)

echo EWDK build environment setup complete
echo.

rem ============================================
rem Build Environment Configuration
rem ============================================

rem Add MSVC linker to PATH
set "PATH=E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\bin\HostX64\x64;%PATH%"

rem Add user-mode library paths for build scripts
rem Build scripts need kernel32.lib, user32.lib, etc.
set "LIB=E:\Program Files\Windows Kits\10\Lib\10.0.28000.0\um\x64;E:\Program Files\Windows Kits\10\Lib\10.0.28000.0\ucrt\x64;%LIB%"

rem Enable static CRT for Windows drivers
set "RUSTFLAGS=-C target-feature=+crt-static"

cd /d "%~dp0"

rem ============================================
rem Build Workspace
rem ============================================
echo Building hyperkd driver (includes hyperhv)...
cd hyperkd
cargo build --release --target x86_64-pc-windows-msvc

if errorlevel 1 (
    echo ERROR: hyperkd build failed
    cd ..
    exit /b 1
)

cd ..
echo hyperkd build completed successfully
echo.

rem ============================================
rem Rename Driver Binary
rem ============================================
echo === Renaming driver binary ===

rem Rename hyperkd driver (.dll -> .sys)
cd hyperkd
if exist "target\x86_64-pc-windows-msvc\release\hyperkd.dll" (
    echo Renaming hyperkd.dll to hyperkd.sys
    ren "target\x86_64-pc-windows-msvc\release\hyperkd.dll" "hyperkd.sys"
)
cd ..

echo.
echo === Build completed ===
echo.
echo Driver binary:
echo - hyperkd.sys (includes hyperhv)
