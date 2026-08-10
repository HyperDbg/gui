$dg = Get-CimInstance -ClassName Win32_DeviceGuard -Namespace root\Microsoft\Windows\DeviceGuard
Write-Host "=== VBS / Device Guard ==="
Write-Host ("VirtualizationBasedSecurityStatus: " + $dg.VirtualizationBasedSecurityStatus)
Write-Host ("SecurityServicesRunning: " + ($dg.SecurityServicesRunning -join ','))
Write-Host ("SecurityServicesConfigured: " + ($dg.SecurityServicesConfigured -join ','))
Write-Host ""
Write-Host "=== CPU VT-x ==="
$cpu = Get-CimInstance Win32_Processor
Write-Host ("VirtualizationFirmwareEnabled: " + $cpu.VirtualizationFirmwareEnabled)
Write-Host ("VMMEnable: " + $cpu.VMMEnable)
Write-Host ""
Write-Host "=== bcdedit hypervisor ==="
bcdedit /enum '{current}' | Select-String -Pattern 'hypervisor'
