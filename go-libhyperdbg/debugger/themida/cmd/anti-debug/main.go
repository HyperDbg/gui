// anti-debug 是独立可执行程序，用于检测 HyperDbg VMM 是否被目标程序的反调试机制发现。
//
// 用法：
//
//	anti-debug.exe -driver <hyperkd.sys> -target <exe> [-seconds 15] [-log path]
//
// 示例：
//
//	anti-debug.exe -driver Debug\hyperkd.sys -target "c:\path\T-VMProtect.exe"
//	anti-debug.exe -driver Debug\hyperkd.sys -target "C:\path\other.exe" -seconds 30
//
// 判定逻辑：
//  1. 进程在监控期内退出（WaitForSingleObject 返回 WAIT_OBJECT_0，或
//     GetExitCodeProcess != STILL_ACTIVE=259）→ DETECTED
//  2. 出现属于目标 pid 的可疑可见窗口（标题含 debug/detect/vmprotect/...）→ DETECTED
//  3. 监控期结束进程仍存活且无可疑窗口 → NOT DETECTED
//
// HyperDbg VMM 是硬件级调试（EPT + VMX root），不使用 Win32 debug port，
// 因此 IsDebuggerPresent / CheckRemoteDebuggerPresent / NtQueryInformationProcess
// 这类 API 反调试理论上检测不到。但 VMProtect 还会用 timing（RDTSC 差值检测 VM exit
// 开销）、CPUID hypervisor bit、IDT/GDT 地址比较、异常处理（VEH/SEH）行为差异等。
// 本工具用于验证这些路径是否会被绕过。
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"golang.org/x/sys/windows"
)

const (
	stillActiveExitCode = uint32(259) // Windows STILL_ACTIVE
	waitObject0         = uint32(0)   // WAIT_OBJECT_0: 对象已 signaled（进程已退出）
	waitTimeout         = uint32(258) // WAIT_TIMEOUT: 等待超时（进程仍在运行）
)

// 可疑窗口标题关键字（小写匹配）。VMProtect/Themida 反调试触发后常见的弹窗文案。
var suspiciousTitleKeywords = []string{
	"debug", "detect", "vmprotect", "themida", "crack", "patch",
	"tamper", "sandbox", "vm ", "virtual", "cheat", "hack",
}

// user32.dll / kernel32.dll 函数（Go 标准库没有包装，直接 syscall）
var (
	user32                    = syscall.NewLazyDLL("user32.dll")
	procEnumWindows           = user32.NewProc("EnumWindows")
	procGetWindowThreadProcId = user32.NewProc("GetWindowThreadProcessId")
	procGetWindowTextW        = user32.NewProc("GetWindowTextW")
	procIsWindowVisible       = user32.NewProc("IsWindowVisible")

	kernel32                = syscall.NewLazyDLL("kernel32.dll")
	procWaitForSingleObject = kernel32.NewProc("WaitForSingleObject")
)

// waitForSingleObject 返回 WAIT_OBJECT_0(0) 表示对象已 signaled（进程已退出），
// WAIT_TIMEOUT(258) 表示超时（进程仍在运行）。
func waitForSingleObject(handle windows.Handle, ms uint32) uint32 {
	r, _, _ := procWaitForSingleObject.Call(uintptr(handle), uintptr(ms))
	return uint32(r)
}

func main() {
	var (
		driverPath   = flag.String("driver", `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\Debug\hyperkd.sys`, "path to hyperkd.sys")
		targetExe    = flag.String("target", `demo.v2.vmp.exe`, "target exe to test (path or name in cwd)")
		runSeconds   = flag.Int("seconds", 15, "seconds to monitor the target")
		logPath      = flag.String("log", "anti-debug.log", "path to log output")
		expectWindow = flag.String("expect-window", "", "window title keyword the target should show when NOT debugged (e.g. 'Demo'); if set and never appears within runSeconds -> DETECTED")
	)
	flag.Parse()

	if _, err := os.Stat(*driverPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: driver not found: %v\n", err)
		os.Exit(2)
	}
	if _, err := os.Stat(*targetExe); err != nil {
		fmt.Fprintf(os.Stderr, "error: target not found: %v\n", err)
		os.Exit(2)
	}

	fmt.Printf("[*] Target  : %s\n", *targetExe)
	fmt.Printf("[*] Driver  : %s\n", *driverPath)
	fmt.Printf("[*] Run sec : %d\n", *runSeconds)
	fmt.Printf("[*] Log     : %s\n", *logPath)
	if *expectWindow != "" {
		fmt.Printf("[*] Expect  : window containing %q (absent => DETECTED)\n", *expectWindow)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: api.New: %v\n", err)
		os.Exit(1)
	}
	defer dbg.Close()

	if err := dbg.LoadVMM(ctx, *driverPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: LoadVMM: %v\n", err)
		os.Exit(1)
	}
	if err := dbg.LogOpen(*logPath); err != nil {
		fmt.Printf("[!] LogOpen failed: %v\n", err)
	}
	defer dbg.LogClose()

	proc, err := dbg.StartProcess(ctx, *targetExe)
	if err != nil {
		_ = dbg.UnloadVMM(ctx)
		fmt.Fprintf(os.Stderr, "error: StartProcess: %v\n", err)
		os.Exit(1)
	}
	pid := proc.Pid
	fmt.Printf("[*] Started pid=%d\n", pid)

	// 第一次 Continue：从第一条指令运行到入口点
	if e := dbg.Continue(ctx); e != nil {
		fmt.Printf("[!] first Continue: %v\n", e)
	}
	// 第二次 Continue：从入口点继续执行（让反调试代码有机会运行）
	if e := dbg.Continue(ctx); e != nil {
		fmt.Printf("[!] second Continue: %v\n", e)
	}
	fmt.Printf("[*] Process running, monitoring %ds...\n", *runSeconds)

	detected := false
	detectedReason := ""
	deadline := time.Now().Add(time.Duration(*runSeconds) * time.Second)
	tick := 0
	expectWindowLower := strings.ToLower(*expectWindow)
	expectWindowSeen := false

	for time.Now().Before(deadline) {
		select {
		case <-ctx.Done():
			detected = true
			detectedReason = "interrupted by signal"
			break
		case <-time.After(1 * time.Second):
		}
		if detected {
			break
		}
		tick++

		// 1) 进程是否已退出 —— 用 WaitForSingleObject 做真实存活检查。
		//    GetExitCodeProcess 在 VMM 调试状态下可能仍返回 STILL_ACTIVE(259)
		//    即使进程已退出（句柄被调试器持有），WaitForSingleObject 更可靠。
		waitResult := waitForSingleObject(windows.Handle(proc.Handle), 0)
		var exitCode uint32
		if e := windows.GetExitCodeProcess(windows.Handle(proc.Handle), &exitCode); e != nil {
			fmt.Printf("[!] GetExitCodeProcess: %v\n", e)
			exitCode = 0xDEAD
		}
		alive := waitResult == waitTimeout

		// 2) 可疑窗口枚举
		titles := enumVisibleWindowsForPid(pid)

		// 每秒打印详细状态：alive/exit/wait/wins/titles
		// 截断 titles 显示避免刷屏
		titlesShort := titles
		if len(titlesShort) > 3 {
			titlesShort = append(titlesShort[:3], "...")
		}
		fmt.Printf("[*] t=%ds alive=%v wait=%d exit=%d wins=%d %v\n",
			tick, alive, waitResult, exitCode, len(titles), titlesShort)

		if !alive {
			// waitResult == WAIT_OBJECT_0(0)：进程已退出
			detected = true
			detectedReason = fmt.Sprintf("process exited at t=%ds (WaitForSingleObject=WAIT_OBJECT_0, GetExitCodeProcess=%d)", tick, exitCode)
			break
		}
		if exitCode != stillActiveExitCode {
			detected = true
			detectedReason = fmt.Sprintf("process exited prematurely at t=%ds (exitCode=%d, wait=%d)", tick, exitCode, waitResult)
			break
		}

		for _, title := range titles {
			lower := strings.ToLower(title)
			// 跟踪期望窗口是否出现
			if expectWindowLower != "" && strings.Contains(lower, expectWindowLower) {
				expectWindowSeen = true
			}
			for _, kw := range suspiciousTitleKeywords {
				if strings.Contains(lower, kw) {
					detected = true
					detectedReason = fmt.Sprintf("suspicious window at t=%ds: %q (matched %q)", tick, title, kw)
					break
				}
			}
			if detected {
				break
			}
		}
		if detected {
			break
		}
	}

	// 3) 监控期结束，进程仍存活但期望窗口从未出现 → DETECTED。
	//    这覆盖 VMP 检测到 VMM 后不退出也不弹窗（挂起/卡在解壳）的情形。
	if !detected && expectWindowLower != "" && !expectWindowSeen {
		detected = true
		detectedReason = fmt.Sprintf("expected window %q never appeared within %ds (process alive but no window — anti-debug likely blocked GUI init)", *expectWindow, *runSeconds)
	}

	// 清理（先 Pause 再 Unload，避免 VM exit 时进程仍在运行触发 BSOD）
	_ = dbg.Pause(ctx)
	_ = dbg.UnloadVMM(ctx)
	if proc.Handle != 0 {
		_ = syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
	}
	proc.Close()

	fmt.Println()
	if detected {
		fmt.Printf("[!] RESULT: DETECTED — %s\n", detectedReason)
		fmt.Printf("[!] HyperDbg VMM was detected by: %s\n", *targetExe)
		fmt.Printf("[!] Note: check %s for kernel-side traces\n", *logPath)
		os.Exit(1)
	}
	fmt.Printf("[+] RESULT: NOT DETECTED — target ran %d seconds without anti-debug trigger\n", *runSeconds)
	fmt.Printf("[+] HyperDbg VMM survived anti-debug checks of: %s\n", *targetExe)
}

// enumVisibleWindowsForPid 枚举所有属于 pid 的可见窗口，返回它们的标题。
func enumVisibleWindowsForPid(pid uint32) []string {
	var titles []string
	cb := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
		var wpid uint32
		_, _, _ = procGetWindowThreadProcId.Call(hwnd, uintptr(unsafe.Pointer(&wpid)))
		if wpid != pid {
			return 1
		}
		vis, _, _ := procIsWindowVisible.Call(hwnd)
		if vis == 0 {
			return 1
		}
		buf := make([]uint16, 512)
		n, _, _ := procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
		if n > 0 {
			titles = append(titles, syscall.UTF16ToString(buf[:n]))
		}
		return 1
	})
	_, _, _ = procEnumWindows.Call(cb, 0)
	return titles
}
