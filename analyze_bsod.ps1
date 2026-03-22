param(
    [string]$DumpFile = "",
    [string]$DriverName = "hyperkd",
    [string]$WinDbgPath = "C:\Microsoft.WinDbg_1.2601.12001.0_x64__8wekyb3d8bbwe"
)

$ErrorActionPreference = "Stop"

$kdExe = Join-Path $WinDbgPath "amd64\kd.exe"
if (-not (Test-Path $kdExe)) {
    $kdExe = Join-Path $WinDbgPath "x64\kd.exe"
}

if ([string]::IsNullOrEmpty($DumpFile)) {
    $DumpFile = Get-ChildItem "C:\Windows\Minidump" -Filter "*.dmp" | 
                Sort-Object LastWriteTime -Descending | 
                Select-Object -First 1 -ExpandProperty FullName
}

$dumpDir = Split-Path $DumpFile
$buildPath = "D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release"
$bsodDir = "d:\笔记本\p\ux\examples\hypedbg\bsod"

Write-Host "=== BSOD 分析 ===" -ForegroundColor Cyan
Write-Host "Dump: $DumpFile" -ForegroundColor Yellow

Copy-Item "$buildPath\$DriverName.pdb" "$dumpDir\$DriverName.pdb" -Force -ErrorAction SilentlyContinue
Copy-Item "$buildPath\$DriverName.sys" "$dumpDir\$DriverName.sys" -Force -ErrorAction SilentlyContinue

$env:NT_SYMBOL_PATH = $dumpDir
$env:_NT_DEBUGGER_EXTENSION_PATH = ""

$output = & $kdExe -z $DumpFile -y $dumpDir -c "lm; .exr -1; k; q" 2>&1 | Out-String

$exceptionCode = ""
$exceptionAddr = ""
$unloadedBase = ""
$unloadedEnd = ""
$stackOffset = ""

if ($output -match "ExceptionCode:\s*([0-9a-fA-FxX]+)") { $exceptionCode = $matches[1] }
if ($output -match "ExceptionAddress:\s*([0-9a-fA-F`]+)") { $exceptionAddr = $matches[1] -replace '`', '' }

$driverPattern = '([0-9a-fA-F`]+)\s+([0-9a-fA-F`]+)\s+' + $DriverName + '\.sys'
$stackPattern = '<Unloaded_' + $DriverName + '\.sys>\+0x([0-9a-fA-F]+)'

$lines = $output -split "`n"
foreach ($line in $lines) {
    if ($line -match $driverPattern) {
        $unloadedBase = $matches[1] -replace '`', ''
        $unloadedEnd = $matches[2] -replace '`', ''
        break
    }
    if ($line -match $stackPattern) {
        $stackOffset = "0x" + $matches[1]
    }
}

Write-Host "`n异常代码: $exceptionCode" -ForegroundColor Red
Write-Host "异常地址: 0x$exceptionAddr" -ForegroundColor Red
if ($unloadedBase) { Write-Host "已卸载模块: 0x$unloadedBase - 0x$unloadedEnd" -ForegroundColor Yellow }

$crashOffset = $stackOffset
if (-not $crashOffset -and $exceptionAddr -and $unloadedBase) {
    $e = [convert]::ToInt64($exceptionAddr, 16)
    $b = [convert]::ToInt64($unloadedBase, 16)
    $end = [convert]::ToInt64($unloadedEnd, 16)
    if ($e -ge $b -and $e -lt $end) {
        $crashOffset = "0x{0:X}" -f ($e - $b)
    }
}

if ($crashOffset) {
    $lnOut = & $kdExe -z "$dumpDir\$DriverName.sys" -y $dumpDir -c ".reload /f $DriverName.sys; ln $DriverName+$crashOffset; q" 2>&1 | Out-String
    
    $crashFunc = ""
    $crashFile = ""
    $crashLine = ""
    
    $funcPattern = $DriverName + '!([^\s\(\)]+)'
    if ($lnOut -match $funcPattern) { $crashFunc = $matches[1] }
    if ($lnOut -match '\[([^\]]+\.cpp|\.c|\.h)[\s:@]+(\d+)') {
        $crashFile = $matches[1]
        $crashLine = $matches[2]
    }
    
    Write-Host "`n崩溃偏移: $crashOffset" -ForegroundColor Yellow
    if ($crashFunc) {
        Write-Host "崩溃函数: $DriverName!$crashFunc" -ForegroundColor Red
        if ($crashFile) { Write-Host "源文件: $crashFile : $crashLine" -ForegroundColor Red }
    }
} elseif ($unloadedBase) {
    Write-Host "`n驱动已卸载，执行详细分析..." -ForegroundColor Yellow
    
    $analyzeOut = & $kdExe -z $DumpFile -y $dumpDir -c "!analyze -v; q" 2>&1 | Out-String
    
    $bugcheck = ""
    $image = ""
    if ($analyzeOut -match "BUGCHECK_CODE:\s*(\S+)") { $bugcheck = $matches[1] }
    if ($analyzeOut -match "IMAGE_NAME:\s*(\S+)") { $image = $matches[1] }
    
    Write-Host "Bugcheck: $bugcheck" -ForegroundColor Red
    Write-Host "崩溃模块: $image" -ForegroundColor Red
    
    $analyzeLines = $analyzeOut -split "`n"
    foreach ($line in $analyzeLines) {
        if ($line -match $stackPattern) {
            $stackOffset = "0x" + $matches[1]
            Write-Host "调用栈偏移: $stackOffset" -ForegroundColor Yellow
            
            $lnOut = & $kdExe -z "$dumpDir\$DriverName.sys" -y $dumpDir -c ".reload /f $DriverName.sys; ln $DriverName+$stackOffset; q" 2>&1 | Out-String
            
            $crashFunc = ""
            $crashFile = ""
            $crashLine = ""
            
            $funcPattern = $DriverName + '!([^\s\(\)]+)'
            if ($lnOut -match $funcPattern) { $crashFunc = $matches[1] }
            if ($lnOut -match '\[([^\]]+\.cpp|\.c|\.h)[\s:@]+(\d+)') {
                $crashFile = $matches[1]
                $crashLine = $matches[2]
            }
            
            if ($crashFunc) {
                Write-Host "崩溃函数: $DriverName!$crashFunc" -ForegroundColor Red
                if ($crashFile) { Write-Host "源文件: $crashFile : $crashLine" -ForegroundColor Red }
            }
            break
        }
    }
    
    if (-not $stackOffset) {
        Write-Host "`n调用栈中未找到 $DriverName 偏移信息" -ForegroundColor Yellow
        Write-Host "可能是驱动卸载后回调未清理导致的间接崩溃" -ForegroundColor Yellow
    }
} else {
    Write-Host "`n未找到驱动信息" -ForegroundColor Yellow
}

Remove-Item "$bsodDir\debug_output.txt" -ErrorAction SilentlyContinue
Remove-Item "$bsodDir\lm_output.txt" -ErrorAction SilentlyContinue
Remove-Item "$bsodDir\analyze_output.txt" -ErrorAction SilentlyContinue
Remove-Item "$bsodDir\test_regex.ps1" -ErrorAction SilentlyContinue

Write-Host "`n=== 完成 ===" -ForegroundColor Cyan
