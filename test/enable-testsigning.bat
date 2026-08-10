@echo off
:: Enable Windows Test Signing Mode (required for HyperDbg driver)
:: Must run as Administrator

net session >nul 2>&1
if %errorLevel% neq 0 (
    echo [!] Please run as Administrator
    pause
    exit /b 1
)

echo [*] Enabling test signing mode...
bcdedit /set testsigning on

if %errorLevel% equ 0 (
    echo [+] Test signing mode enabled successfully
    echo [*] Reboot required for changes to take effect
) else (
    echo [-] Failed to enable test signing mode
)

echo.
echo [*] Current bcdedit status:
bcdedit /enum | findstr testsigning

pause
