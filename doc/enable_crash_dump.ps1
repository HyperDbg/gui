# Enable Windows Crash Dump Configuration
# This script enables crash dump collection for BSOD analysis

Write-Host "=== Windows Crash Dump Configuration ===" -ForegroundColor Cyan

# Check if running as Administrator
$isAdmin = ([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
if (-not $isAdmin) {
    Write-Host "Error: This script must be run as Administrator" -ForegroundColor Red
    exit 1
}

# Registry path for crash control
$crashControlPath = 'HKLM:\SYSTEM\CurrentControlSet\Control\CrashControl'

# Display current settings
Write-Host "`nCurrent Crash Dump Settings:" -ForegroundColor Yellow
Get-ItemProperty -Path $crashControlPath | Select-Object CrashDumpEnabled, DumpFile, MinidumpDir | Format-List

# Configure crash dump settings
Write-Host "`nConfiguring crash dump settings..." -ForegroundColor Green

# Set crash dump type to Small Memory Dump (1)
# 0 = None, 1 = Small Memory Dump, 2 = Kernel Memory Dump, 3 = Complete Memory Dump
Set-ItemProperty -Path $crashControlPath -Name CrashDumpEnabled -Value 1 -Force

# Ensure minidump directory exists
$minidumpDir = "C:\Windows\Minidump"
if (-not (Test-Path $minidumpDir)) {
    New-Item -Path $minidumpDir -ItemType Directory -Force | Out-Null
    Write-Host "Created minidump directory: $minidumpDir" -ForegroundColor Green
}

# Set minidump directory
Set-ItemProperty -Path $crashControlPath -Name MinidumpDir -Value $minidumpDir -Force

# Set dump file path (only for complete dump, we're using minidump)
# Set-ItemProperty -Path $crashControlPath -Name DumpFile -Value "C:\Windows\MEMORY.DMP" -Force

# Optional: Configure additional settings for better debugging
Write-Host "`nConfiguring additional debugging settings..." -ForegroundColor Green

# Send administrative alert (1 = enabled)
Set-ItemProperty -Path $crashControlPath -Name SendAdminAlert -Value 1 -Force

# Write event to system log (1 = enabled)
Set-ItemProperty -Path $crashControlPath -Name LogEvent -Value 1 -Force

# Automatically reboot after crash (1 = enabled)
Set-ItemProperty -Path $crashControlPath -Name AutoReboot -Value 1 -Force

# Overwrite existing dump file (1 = enabled)
Set-ItemProperty -Path $crashControlPath -Name Overwrite -Value 1 -Force

# Display new settings
Write-Host "`nNew Crash Dump Settings:" -ForegroundColor Yellow
Get-ItemProperty -Path $crashControlPath | Select-Object CrashDumpEnabled, DumpFile, MinidumpDir, AutoReboot, Overwrite, LogEvent, SendAdminAlert | Format-List

Write-Host "`n=== Configuration Complete ===" -ForegroundColor Cyan
Write-Host "Crash dump collection is now enabled." -ForegroundColor Green
Write-Host "Dump files will be saved to: $minidumpDir" -ForegroundColor Green
Write-Host "Full memory dump: C:\Windows\MEMORY.DMP" -ForegroundColor Green
Write-Host "`nNote: You must restart the system for changes to take effect." -ForegroundColor Yellow
