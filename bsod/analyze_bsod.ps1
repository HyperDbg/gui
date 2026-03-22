<#
.SYNOPSIS
    分析 BSOD dump 文件，定位崩溃函数和行号

.DESCRIPTION
    使用 WinDbg kd.exe 分析 minidump 文件，通过 PDB 符号文件定位崩溃位置。
    特别适用于驱动卸载后的 BSOD 分析。

.PARAMETER DumpFile
    dump 文件路径，默认为最新的 minidump

.PARAMETER PdbPath
    PDB 符号文件所在目录

.PARAMETER DriverName
    驱动名称（不含扩展名），默认为 hyperkd

.PARAMETER WinDbgPath
    WinDbg 安装路径，默认为 C:\Microsoft.WinDbg_1.2601.12001.0_x64__8wekyb3d8bbwe

.EXAMPLE
    .\analyze_bsod.ps1
    分析最新的 minidump，使用默认路径

.EXAMPLE
    .\analyze_bsod.ps1 -DumpFile "C:\Windows\Minidump\032226-4187-01.dmp" -PdbPath "D:\build\Release"
    指定 dump 文件和 PDB 路径

.EXAMPLE
    .\analyze_bsod.ps1 -DriverName "mydriver"
    分析名为 mydriver.sys 的驱动相关崩溃

.NOTES
    输出信息包括：
    - 崩溃模块和偏移
    - 具体函数名
    - 调用栈
#>

param(
    [string]$DumpFile = "",
    [string]$PdbPath = "",
    [string]$DriverName = "hyperkd",
    [string]$WinDbgPath = "C:\Microsoft.WinDbg_1.2601.12001.0_x64__8wekyb3d8bbwe"
)

$ErrorActionPreference = "Stop"

$kdExe = Join-Path $WinDbgPath "amd64\kd.exe"
if (-not (Test-Path $kdExe)) {
    $kdExe = Join-Path $WinDbgPath "x64\kd.exe"
}
if (-not (Test-Path $kdExe)) {
    Write-Error "找不到 kd.exe，请检查 WinDbgPath 参数: $WinDbgPath"
    exit 1
}

if ([string]::IsNullOrEmpty($DumpFile)) {
    $minidumpDir = "C:\Windows\Minidump"
    $DumpFile = Get-ChildItem $minidumpDir -Filter "*.dmp" | 
                Sort-Object LastWriteTime -Descending | 
                Select-Object -First 1 -ExpandProperty FullName
    if (-not $DumpFile) {
        Write-Error "在 $minidumpDir 中找不到 dump 文件"
        exit 1
    }
}

if (-not (Test-Path $DumpFile)) {
    Write-Error "Dump 文件不存在: $DumpFile"
    exit 1
}

Write-Host "=== BSOD Dump 分析工具 ===" -ForegroundColor Cyan
Write-Host "Dump 文件: $DumpFile" -ForegroundColor Yellow
Write-Host "WinDbg: $kdExe" -ForegroundColor Yellow

$symPath = "srv*C:\Symbols*https://msdl.microsoft.com/download/symbols"
if (-not [string]::IsNullOrEmpty($PdbPath)) {
    $symPath += ";$PdbPath"
    Write-Host "PDB 路径: $PdbPath" -ForegroundColor Yellow
}

Write-Host "`n[1/4] 获取已卸载模块列表..." -ForegroundColor Green
$unloadedCmd = "r; lm; q"
$unloadedOutput = & $kdExe -z $DumpFile -y $symPath -c $unloadedCmd 2>&1 | Out-String

$unloadedMatch = [regex]::Match($unloadedOutput, "Unloaded modules:\r?\n((?:.*\r?\n)+)")
if ($unloadedMatch.Success) {
    Write-Host "`n已卸载的模块:" -ForegroundColor Red
    $unloadedMatch.Groups[1].Value -split "`n" | Where-Object { $_ -match $DriverName } | ForEach-Object {
        Write-Host "  $_" -ForegroundColor Red
    }
}

Write-Host "`n[2/4] 获取崩溃调用栈..." -ForegroundColor Green
$stackCmd = ".exr -1; .ecxr; kv; q"
$stackOutput = & $kdExe -z $DumpFile -y $symPath -c $stackCmd 2>&1 | Out-String

$stackMatch = [regex]::Match($stackOutput, "# Child-SP\s+RetAddr.*?(?=quit:|NatVis)")
if ($stackMatch.Success) {
    Write-Host "`n调用栈:" -ForegroundColor Yellow
    $stackMatch.Value -split "`n" | ForEach-Object {
        if ($_ -match "Unloaded_$DriverName") {
            Write-Host "  $_" -ForegroundColor Red
        } else {
            Write-Host "  $_"
        }
    }
}

Write-Host "`n[3/4] 提取崩溃偏移..." -ForegroundColor Green
$offsetMatch = [regex]::Match($stackOutput, "<Unloaded_$DriverName\.sys>\+0x([0-9a-fA-F]+)")
$crashOffset = $null
$unloadedBase = $null
if ($offsetMatch.Success) {
    $crashOffset = "0x" + $offsetMatch.Groups[1].Value
    Write-Host "崩溃偏移: $crashOffset" -ForegroundColor Red
} else {
    $offsetMatch = [regex]::Match($stackOutput, "$DriverName\+0x([0-9a-fA-F]+)")
    if ($offsetMatch.Success) {
        $crashOffset = "0x" + $offsetMatch.Groups[1].Value
        Write-Host "崩溃偏移: $crashOffset" -ForegroundColor Red
    } else {
        $unloadedModuleMatch = [regex]::Matches($unloadedOutput, "([0-9a-fA-F`]+)\s+([0-9a-fA-F`]+)\s+$DriverName\.sys")
        if ($unloadedModuleMatch.Count -gt 0) {
            Write-Host "从已卸载模块地址范围计算偏移..." -ForegroundColor Yellow
            foreach ($match in $unloadedModuleMatch) {
                $baseAddr = $match.Groups[1].Value -replace '`', ''
                $endAddr = $match.Groups[2].Value -replace '`', ''
                Write-Host "  已卸载模块: 基址=0x$baseAddr, 结束=0x$endAddr" -ForegroundColor Yellow
            }
            $firstMatch = $unloadedModuleMatch[0]
            $unloadedBase = $firstMatch.Groups[1].Value -replace '`', ''
            Write-Host "  使用第一个已卸载模块基址: 0x$unloadedBase" -ForegroundColor Yellow
        }
    }
}

Write-Host "`n[4/4] 定位崩溃函数..." -ForegroundColor Green

$driverSys = $null
$pdbDir = $null
if (-not [string]::IsNullOrEmpty($PdbPath)) {
    $directPath = Join-Path $PdbPath "$DriverName.sys"
    if (Test-Path $directPath) {
        $driverSys = $directPath
        $pdbDir = $PdbPath
    } else {
        $driverSys = Get-ChildItem $PdbPath -Filter "$DriverName.sys" -Recurse -ErrorAction SilentlyContinue | 
                     Select-Object -First 1 -ExpandProperty FullName
        if ($driverSys) {
            $pdbDir = Split-Path $driverSys
        }
    }
    
    if (-not $driverSys) {
        $driverSys = Get-ChildItem (Split-Path $PdbPath) -Filter "$DriverName.sys" -Recurse -ErrorAction SilentlyContinue | 
                     Select-Object -First 1 -ExpandProperty FullName
        if ($driverSys) {
            $pdbDir = Split-Path $driverSys
        }
    }
}

if ($crashOffset -and $driverSys -and (Test-Path $driverSys)) {
    Write-Host "驱动文件: $driverSys" -ForegroundColor Yellow
    Write-Host "PDB 目录: $pdbDir" -ForegroundColor Yellow
    
    $funcCmd = "lm; ln ${DriverName}+$crashOffset; q"
    $funcOutput = & $kdExe -z $driverSys -y $pdbDir -c $funcCmd 2>&1 | Out-String
    
    $funcMatch = [regex]::Match($funcOutput, "\(${DriverName}!([^)\s]+)")
    if ($funcMatch.Success) {
        $funcName = $funcMatch.Groups[1].Value
        Write-Host "`n========================================" -ForegroundColor Cyan
        Write-Host "崩溃函数: ${DriverName}!$funcName" -ForegroundColor Red
        Write-Host "偏移: $crashOffset" -ForegroundColor Red
        Write-Host "========================================" -ForegroundColor Cyan
    } else {
        $funcMatch = [regex]::Match($funcOutput, "Exact matches:\s*\r?\n\s*${DriverName}!([^\s]+)")
        if ($funcMatch.Success) {
            $funcName = $funcMatch.Groups[1].Value
            Write-Host "`n========================================" -ForegroundColor Cyan
            Write-Host "崩溃函数: ${DriverName}!$funcName" -ForegroundColor Red
            Write-Host "偏移: $crashOffset" -ForegroundColor Red
            Write-Host "========================================" -ForegroundColor Cyan
        }
    }
} elseif ($unloadedBase -and $driverSys -and (Test-Path $driverSys)) {
    Write-Host "驱动文件: $driverSys" -ForegroundColor Yellow
    Write-Host "分析已卸载模块回调问题..." -ForegroundColor Yellow
    
    $exceptionAddrMatch = [regex]::Match($stackOutput, "ExceptionAddress:\s*([0-9a-fA-F`]+)")
    if ($exceptionAddrMatch.Success) {
        $exceptionAddr = $exceptionAddrMatch.Groups[1].Value -replace '`', ''
        Write-Host "异常地址: 0x$exceptionAddr" -ForegroundColor Yellow
        
        $baseInt = [convert]::ToInt64($unloadedBase, 16)
        $exceptionInt = [convert]::ToInt64($exceptionAddr, 16)
        
        if ($exceptionInt -ge $baseInt -and $exceptionInt -lt ($baseInt + 0xD0000)) {
            $calculatedOffset = "0x{0:X}" -f ($exceptionInt - $baseInt)
            Write-Host "异常地址在已卸载模块范围内!" -ForegroundColor Red
            Write-Host "计算偏移: $calculatedOffset" -ForegroundColor Red
            
            $pdbDir = Split-Path $driverSys
            $funcCmd = "lm; ln ${DriverName}+$calculatedOffset; q"
            $funcOutput = & $kdExe -z $driverSys -y $pdbDir -c $funcCmd 2>&1 | Out-String
            
            $funcMatch = [regex]::Match($funcOutput, "\(${DriverName}!([^)\s]+)")
            if ($funcMatch.Success) {
                $funcName = $funcMatch.Groups[1].Value
                Write-Host "`n========================================" -ForegroundColor Cyan
                Write-Host "崩溃函数: ${DriverName}!$funcName" -ForegroundColor Red
                Write-Host "偏移: $calculatedOffset" -ForegroundColor Red
                Write-Host "========================================" -ForegroundColor Cyan
            }
        } else {
            Write-Host "异常地址不在已卸载模块范围内" -ForegroundColor Yellow
            Write-Host "这可能是驱动卸载后回调未清理导致的间接崩溃" -ForegroundColor Yellow
        }
    }
} else {
    Write-Warning "未找到崩溃偏移或驱动文件，跳过函数定位"
}

Write-Host "`n[详细分析输出]" -ForegroundColor Green
$analyzeCmd = "!analyze -v; q"
& $kdExe -z $DumpFile -y $symPath -c $analyzeCmd 2>&1 | 
    Select-String -Pattern "MODULE_NAME|IMAGE_NAME|PROCESS_NAME|FAILURE_BUCKET_ID|STACK_TEXT|ExceptionAddress|ExceptionCode|$DriverName" -Context 0,1

Write-Host "`n=== 分析完成 ===" -ForegroundColor Cyan
