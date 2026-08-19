param([string]$ExePath = "d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\cmd\anti-debug\demo.v2.vmp.exe")

Add-Type @"
using System;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using System.Text;
public class WinUtil {
    [DllImport("user32.dll")]
    public static extern bool EnumWindows(EnumWindowsProc lpEnumFunc, IntPtr lParam);
    public delegate bool EnumWindowsProc(IntPtr hWnd, IntPtr lParam);
    [DllImport("user32.dll")]
    public static extern uint GetWindowThreadProcessId(IntPtr hWnd, out uint lpdwProcessId);
    [DllImport("user32.dll")]
    public static extern bool IsWindowVisible(IntPtr hWnd);
    [DllImport("user32.dll", CharSet=CharSet.Unicode)]
    public static extern int GetWindowText(IntPtr hWnd, StringBuilder lpString, int nMaxCount);
    public static List<string> VisibleWindowsForPid(uint pid) {
        var result = new List<string>();
        EnumWindows((hWnd, lp) => {
            uint wpid;
            GetWindowThreadProcessId(hWnd, out wpid);
            if (wpid != pid) return true;
            if (!IsWindowVisible(hWnd)) return true;
            var sb = new StringBuilder(512);
            if (GetWindowText(hWnd, sb, 512) > 0) result.Add(sb.ToString());
            return true;
        }, IntPtr.Zero);
        return result;
    }
}
"@

Write-Output "[*] Launching (no VMM): $ExePath"
$p = Start-Process -FilePath $ExePath -PassThru
Write-Output "[*] Started pid=$($p.Id), waiting 8s..."
for ($i = 1; $i -le 8; $i++) {
    Start-Sleep -Seconds 1
    $titles = [WinUtil]::VisibleWindowsForPid($p.Id)
    $alive = -not $p.HasExited
    $exitCode = if ($alive) { 259 } else { $p.ExitCode }
    $titlesStr = $titles -join ', '
    Write-Output "[*] t=${i}s alive=$alive exit=$exitCode wins=$($titles.Count) [$titlesStr]"
}
if (-not $p.HasExited) {
    Write-Output "[*] Killing pid=$($p.Id)"
    $p | Stop-Process -Force
} else {
    Write-Output "[*] Process already exited with code $($p.ExitCode)"
}
