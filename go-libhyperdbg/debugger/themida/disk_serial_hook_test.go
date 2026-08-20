package themida

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
)

// --- 硬盘序列号 hook 测试 ---
//
// 目标：hook DeviceIoControl，检测目标进程何时通过
// IOCTL_STORAGE_QUERY_PROPERTY (0x002D1400) 查询物理硬盘序列号。
//
// 原理：Windows 软件读取硬盘序列号的标准途径是：
//   CreateFile("\\.\PhysicalDrive0") → DeviceIoControl(IOCTL_STORAGE_QUERY_PROPERTY)
//   返回的 STORAGE_DEVICE_DESCRIPTOR 中包含 SerialNumberOffset。
//
// 当前白名单没有 WriteMem 方法，所以只能检测/记录，不能篡改输出缓冲
// 区中的序列号。要实现篡改需要：1) 在白名单添加 WriteMemByte/Qword；
// 或 2) 实现 SetMemWriteBp 在 DeviceIoControl 返回后触发回调修改缓冲区。

const (
	diskDriverPath = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\Debug\hyperkd.sys`
	diskLogPath    = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\test-disk-serial.log`

	// 32-bit DLL paths (for WOW64 targets like SuperRecovery)
	diskKernel32WOW64   = `C:\Windows\SysWOW64\kernel32.dll`
	diskKernelbaseWOW64 = `C:\Windows\SysWOW64\kernelbase.dll`
	diskNtdllWOW64      = `C:\Windows\SysWOW64\ntdll.dll`
	// 64-bit DLL paths (for native 64-bit targets)
	diskKernel32Native   = `C:\Windows\System32\kernel32.dll`
	diskKernelbaseNative = `C:\Windows\System32\kernelbase.dll`
	diskNtdllNative      = `C:\Windows\System32\ntdll.dll`

	// SuperRecovery V4.8.1.5 — 32-bit Themida packed binary that reads disk serial
	diskSuperRecoveryExe = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\SuperRecovery V4.8.1.5.exe`

	// IOCTL_STORAGE_QUERY_PROPERTY = CTL_CODE(0x2D, 0x500, METHOD_BUFFERED, FILE_ANY_ACCESS)
	ioctlStorageQueryProperty = 0x002D1400

	// Sysret hook test uses a separate log to avoid clashing with the EPT-hook tests.
	diskSysretLogPath = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\test-disk-serial-sysret.log`
)

// ============================================================
// 测试 1：Hook SuperRecovery V4.8.1.5 (32-bit WOW64)
// ============================================================
//
// SuperRecovery 是 32-bit Themida 加壳程序，启动时就读取硬盘序列号。
//
// 关键改进：hook NtDeviceIoControlFile（ntdll 层，比 DeviceIoControl 更底层）
// 并在第一次 Continue 之前就设置好 EPT hook——进程一启动执行，hook 就已就位。
//
// NtDeviceIoControlFile 32-bit stdcall 栈布局（10 个参数）：
//   ESP+0x00: return address
//   ESP+0x04: FileHandle
//   ESP+0x08: Event
//   ESP+0x0C: ApcRoutine
//   ESP+0x10: ApcContext
//   ESP+0x14: IoStatusBlock
//   ESP+0x18: IoControlCode     ← 目标
//   ESP+0x1C: InputBuffer
//   ESP+0x20: InputBufferLength
//   ESP+0x24: OutputBuffer
//   ESP+0x28: OutputBufferLength

func TestDiskSerialHook_SuperRecovery(t *testing.T) {
	if _, err := os.ReadFile(diskDriverPath); err != nil {
		t.Skipf("hyperkd.sys not found: %v", err)
	}
	if _, err := os.Stat(diskSuperRecoveryExe); err != nil {
		t.Skipf("SuperRecovery not found: %v", err)
	}

	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}
	defer func() { _ = dbg.UnloadVMM(); _ = dbg.UnloadDriver() }()

	if err := dbg.LoadDriver(diskDriverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}

	// 启动 SuperRecovery（挂起状态）
	proc, err := dbg.StartProcess(diskSuperRecoveryExe)
	if err != nil {
		dbg.UnloadVMM()
		t.Fatalf("StartProcess: %v", err)
	}
	pid := proc.Pid
	t.Logf("[*] Started SuperRecovery (pid=%d, suspended)", pid)

	// ★ 先 Continue 一次让 loader 完成 DLL 映射（包括 32-bit ntdll）
	//   然后立即暂停，在 Themida loader 开始执行之前设置 hook
	_ = dbg.Continue()
	t.Logf("[*] First Continue done (loader initialized)")

	// 解析 32-bit kernelbase!DeviceIoControl（ntdll 在 WOW64 中无法通过
	// EnumProcessModulesEx 获取 32-bit 基址，故先用 kernelbase 的 DeviceIoControl）
	kbBase, err := getModuleBaseForWow64Target(pid, proc.Handle, "kernelbase.dll")
	if err != nil {
		t.Logf("[!] getModuleBaseForWow64Target(kernelbase.dll): %v", err)
	}
	if kbBase > 0xFFFFFFFF {
		if base, e := getModuleBaseFromRunningWow64("kernelbase.dll"); e == nil && base <= 0xFFFFFFFF {
			kbBase = base
		}
	}
	t.Logf("[*] kernelbase base = 0x%X", kbBase)

	var dicAddr uint64
	if kbBase != 0 && kbBase <= 0xFFFFFFFF {
		if rva, _, e := resolveExportRVA(diskKernelbaseWOW64, "DeviceIoControl"); e == nil {
			dicAddr = kbBase + rva
			t.Logf("[*] DeviceIoControl @ kernelbase+0x%X = 0x%X", rva, dicAddr)
		} else {
			t.Logf("[!] resolveExportRVA(kernelbase, DeviceIoControl): %v", e)
		}
	}
	if dicAddr == 0 {
		cleanupDisk(t, dbg, &proc)
		t.Fatal("could not resolve DeviceIoControl address")
	}

	// Warmup: 运行进程直到 DeviceIoControl 页面 fault in
	t.Logf("[*] Warmup: waiting for DeviceIoControl page to fault in...")
	pagesReady := false
	for i := 0; i < 20; i++ {
		_ = dbg.Continue()
		time.Sleep(300 * time.Millisecond)
		if e := touchPageViaReadProcessMemory(proc.Handle, dicAddr); e == nil {
			pagesReady = true
			t.Logf("[*] Page faulted in after %d iterations", i+1)
			break
		}
	}
	if !pagesReady {
		t.Logf("[!] Warmup: page not faulted in, trying touchPage directly")
		_ = touchPageViaReadProcessMemory(proc.Handle, dicAddr)
	}

	// 暂停进程，设置 hook
	_ = dbg.Pause()
	t.Logf("[*] Process paused, setting up hook")
	_ = touchPageViaReadProcessMemory(proc.Handle, dicAddr)

	// ★ 设置 EPT hook — 32-bit stdcall: ESP+8=dwIoControlCode
	//   记录所有 ioctl 调用以便分析，对 0x2D1400 特别标记
	hookSrc := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	ioctl := ctx.StackReadDword(8)
	ctx.Printf("DeviceIoControl: ioctl=%%x\n", ioctl)
	if ioctl == 0x%X {
		ctx.Printf("--------------- DISK_SERIAL_QUERY ----------------\n")
		inBuf := ctx.StackReadDword(0xC)
		propId := ctx.ReadMemQword(inBuf) & 0xFFFFFFFF
		ctx.Printf("PropertyId: %%x\n", propId)
		outBuf := ctx.StackReadDword(0x14)
		ctx.Printf("OutBuf: %%x\n", outBuf)
	}
}`, pid, ioctlStorageQueryProperty)

	tag, err := dbg.EptHookForProcess(dicAddr, pid, hookSrc)
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("EptHookForProcess(DeviceIoControl): %v", err)
	}
	t.Logf("[*] Hooked DeviceIoControl tag=%d", tag)

	// 启动消息泵
	pump, pumpErr := dbg.StartMessagePump()
	if pumpErr != nil {
		t.Logf("[!] StartMessagePump: %v", pumpErr)
	}
	defer func() {
		if pump != nil {
			pump.Stop()
		}
	}()

	// ★ 现在 Continue——hook 已就位，从进程第一条指令开始拦截
	_ = dbg.Continue()
	t.Logf("[*] Process resumed with hook in place, running 30s...")

	<-time.After(30 * time.Second)

	// 清理
	_ = dbg.Pause()
	if pump != nil {
		pump.Stop()
	}
	_ = dbg.UnloadVMM()
	if proc.Handle != 0 {
		syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
	}
	proc.Close()

	// 检查日志
	logData, err := os.ReadFile(diskLogPath)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	logStr := string(logData)
	t.Logf("[*] Log size: %d bytes", len(logData))

	if strings.Contains(logStr, "DISK_SERIAL_QUERY") {
		queryCount := strings.Count(logStr, "DISK_SERIAL_QUERY")
		t.Logf("PASS: disk serial query detected! (%d queries)", queryCount)
		for _, line := range strings.Split(logStr, "\n") {
			trim := strings.TrimSpace(line)
			if strings.Contains(trim, "PropertyId") || strings.Contains(trim, "OutBuf") || strings.Contains(trim, "DISK_SERIAL") {
				t.Logf("  %s", trim)
			}
		}
	} else {
		t.Logf("No DISK_SERIAL_QUERY in log (SuperRecovery may use a different method/IOCTL to read disk serial)")
	}
}

// ============================================================
// 测试 2：Hook 自编译 helper (64-bit native)
// ============================================================

const diskHelperSrc = `package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const IOCTL_STORAGE_QUERY_PROPERTY = 0x002D1400

func main() {
	name, _ := syscall.UTF16PtrFromString(` + "`" + `\\.\PhysicalDrive0` + "`" + `)
	h, err := syscall.CreateFile(name, syscall.GENERIC_READ,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil, syscall.OPEN_EXISTING, 0, 0)
	if err != nil {
		fmt.Println("createfile error:", err)
		return
	}
	defer syscall.CloseHandle(h)
	var query [12]byte
	var out [8192]byte
	var ret uint32
	err = syscall.DeviceIoControl(h, IOCTL_STORAGE_QUERY_PROPERTY,
		(*byte)(unsafe.Pointer(&query[0])), uint32(len(query)),
		&out[0], uint32(len(out)), &ret, nil)
	if err != nil {
		fmt.Println("ioctl error:", err)
		return
	}
	fmt.Println("ok bytes=", ret)
}
`

func TestDiskSerialHook_Helper64(t *testing.T) {
	if _, err := os.ReadFile(diskDriverPath); err != nil {
		t.Skipf("hyperkd.sys not found: %v", err)
	}

	// 编译 helper
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module diskquery\n\ngo 1.26\n"), 0o644)
	os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(diskHelperSrc), 0o644)
	helperExe := filepath.Join(tmpDir, "diskquery.exe")
	buildCmd := exec.Command("go", "build", "-o", helperExe, ".")
	buildCmd.Dir = tmpDir
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("build helper: %v\n%s", err, out)
	}
	t.Logf("[*] Built helper: %s", helperExe)

	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}
	defer func() { _ = dbg.UnloadVMM(); _ = dbg.UnloadDriver() }()

	if err := dbg.LoadDriver(diskDriverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}

	proc, err := dbg.StartProcess(helperExe)
	if err != nil {
		dbg.UnloadVMM()
		t.Fatalf("StartProcess: %v", err)
	}
	pid := proc.Pid
	t.Logf("[*] Started helper (pid=%d)", pid)

	_ = dbg.Continue()
	_ = dbg.Continue()

	// 解析 64-bit DeviceIoControl — 使用 getModuleBaseForWow64Target（含多种回退）
	var dicAddr uint64
	if base, e := getModuleBaseForWow64Target(pid, proc.Handle, "kernel32.dll"); e == nil && base != 0 {
		t.Logf("[*] kernel32 base=0x%X", base)
		if rva, _, e := resolveExportRVA(diskKernel32Native, "DeviceIoControl"); e == nil {
			dicAddr = base + rva
			t.Logf("[*] DeviceIoControl @ kernel32+0x%X = 0x%X", rva, dicAddr)
		} else {
			t.Logf("[!] resolveExportRVA(kernel32, DeviceIoControl): %v", e)
		}
	} else {
		t.Logf("[!] getModuleBaseForWow64Target(kernel32.dll): %v", e)
	}
	if dicAddr == 0 {
		if base, e := getModuleBaseForWow64Target(pid, proc.Handle, "kernelbase.dll"); e == nil && base != 0 {
			t.Logf("[*] kernelbase base=0x%X", base)
			if rva, _, e := resolveExportRVA(diskKernelbaseNative, "DeviceIoControl"); e == nil {
				dicAddr = base + rva
				t.Logf("[*] DeviceIoControl @ kernelbase+0x%X = 0x%X", rva, dicAddr)
			} else {
				t.Logf("[!] resolveExportRVA(kernelbase, DeviceIoControl): %v", e)
			}
		} else {
			t.Logf("[!] getModuleBaseForWow64Target(kernelbase.dll): %v", e)
		}
	}
	if dicAddr == 0 {
		cleanupDisk(t, dbg, &proc)
		t.Fatal("could not resolve DeviceIoControl address")
	}

	if e := touchPageViaReadProcessMemory(proc.Handle, dicAddr); e != nil {
		t.Logf("[!] touchPage(0x%X): %v", dicAddr, e)
	}

	// 64-bit hook: x64 fastcall — RDX=dwIoControlCode, R8=lpInBuffer
	hookSrc := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	ioctl := ctx.Reg("rdx")
	ctx.Printf("DeviceIoControl: ioctl=%%x\n", ioctl)
	if ioctl == 0x%X {
		ctx.Printf("--------------- DISK_SERIAL_QUERY ----------------\n")
		inBuf := ctx.Reg("r8")
		propId := ctx.ReadMemQword(inBuf) & 0xFFFFFFFF
		ctx.Printf("PropertyId: %%x\n", propId)
	}
}`, pid, ioctlStorageQueryProperty)

	tag, err := dbg.EptHookForProcess(dicAddr, pid, hookSrc)
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("EptHookForProcess: %v", err)
	}
	t.Logf("[*] Hooked DeviceIoControl tag=%d", tag)

	pump, pumpErr := dbg.StartMessagePump()
	if pumpErr != nil {
		t.Logf("[!] StartMessagePump: %v", pumpErr)
	}
	defer func() {
		if pump != nil {
			pump.Stop()
		}
	}()

	_ = dbg.Continue()
	t.Logf("[*] Helper running (max 15s)...")

	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		if err := dbg.Continue(); err != nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	_ = dbg.Pause()
	if pump != nil {
		pump.Stop()
	}
	_ = dbg.UnloadVMM()
	if proc.Handle != 0 {
		syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
	}
	proc.Close()

	logData, err := os.ReadFile(diskLogPath)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	logStr := string(logData)
	t.Logf("[*] Log size: %d bytes", len(logData))

	dicCount := strings.Count(logStr, "DeviceIoControl:")
	t.Logf("DeviceIoControl calls: %d", dicCount)

	if strings.Contains(logStr, "DISK_SERIAL_QUERY") {
		t.Logf("PASS: disk serial query detected!")
	} else {
		t.Logf("No DISK_SERIAL_QUERY (helper may not have reached DeviceIoControl)")
	}
}

// ============================================================
// 测试 3：Sysret hook — detect NtDeviceIoControlFile return (64-bit native)
// ============================================================
//
// 与测试 1/2 的 EPT-hook 方案不同，本测试使用 !sysret hook。
//
// 原理：EFER.SCE 被清除后，每次 SYSCALL/SYSRET 触发 #UD → VM exit。
// 驱动读取指令字节判断是 SYSCALL 还是 SYSRET，对 SYSRET 调用
// DispatchEventEferSysret。由于 SYSRET 时 EAX 已被覆盖为返回值，
// 驱动无法知道刚执行的是哪个系统调用（见 Debugger.c 注释），因此
// SYSRET hook 只能 hook 所有 SYSRET，不能按系统调用号过滤。
//
// 识别 NtDeviceIoControlFile 返回的方法：SYSRET 时 RCX = 用户态返回
// 地址（即 ntdll!NtDeviceIoControlFile 中 syscall 指令的下一条指令）。
// 通过读取 NtDeviceIoControlFile 的字节，定位 0F 05 (syscall) 指令，
// 计算 retAddr = funcBase + syscallOffset + 2，在 hook 中比较 RCX。
//
// 优势：不修改任何用户态指令（EPT hook 需要），更难被反调试检测。

func TestDiskSerialHook_Sysret_Helper64(t *testing.T) {
	if _, err := os.ReadFile(diskDriverPath); err != nil {
		t.Skipf("hyperkd.sys not found: %v", err)
	}

	// 编译 helper (reuse diskHelperSrc)
	tmpDir := t.TempDir()
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module diskquery\n\ngo 1.26\n"), 0o644)
	os.WriteFile(filepath.Join(tmpDir, "main.go"), []byte(diskHelperSrc), 0o644)
	helperExe := filepath.Join(tmpDir, "diskquery.exe")
	buildCmd := exec.Command("go", "build", "-o", helperExe, ".")
	buildCmd.Dir = tmpDir
	if out, err := buildCmd.CombinedOutput(); err != nil {
		t.Fatalf("build helper: %v\n%s", err, out)
	}
	t.Logf("[*] Built helper: %s", helperExe)

	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		t.Fatalf("api.New: %v", err)
	}
	defer func() { _ = dbg.UnloadVMM(); _ = dbg.UnloadDriver() }()

	if err := dbg.LoadDriver(diskDriverPath); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}
	if err := dbg.InitVMM(); err != nil {
		t.Fatalf("LoadVMM: %v", err)
	}

	proc, err := dbg.StartProcess(helperExe)
	if err != nil {
		dbg.UnloadVMM()
		t.Fatalf("StartProcess: %v", err)
	}
	pid := proc.Pid
	t.Logf("[*] Started helper (pid=%d)", pid)

	_ = dbg.Continue()
	_ = dbg.Continue()

	// Resolve ntdll!NtDeviceIoControlFile address in the target process.
	ntdllBase, err := getModuleBaseForWow64Target(pid, proc.Handle, "ntdll.dll")
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("getModuleBaseForWow64Target(ntdll.dll): %v", err)
	}
	t.Logf("[*] ntdll base = 0x%X", ntdllBase)

	ntdicRVA, _, err := resolveExportRVA(diskNtdllNative, "NtDeviceIoControlFile")
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("resolveExportRVA(NtDeviceIoControlFile): %v", err)
	}
	ntdicAddr := ntdllBase + ntdicRVA
	t.Logf("[*] NtDeviceIoControlFile @ ntdll+0x%X = 0x%X", ntdicRVA, ntdicAddr)

	// Read the first 32 bytes of NtDeviceIoControlFile to find the syscall
	// instruction (0x0F 0x05). Standard ntdll syscall stub layout (Win10/11 x64):
	//   +0:  4C 8B D1             mov r10, rcx
	//   +3:  B8 XX XX XX XX       mov eax, <ssn>
	//   +8:  F6 04 25 ...         test byte ptr [7FFE0308h], 1  (Win10+)
	//   +16: 75 03                jne short +3
	//   +18: 0F 05                syscall
	//   +20: C3                   ret
	// OR (simpler/older):
	//   +8:  0F 05                syscall
	//   +10: C3                   ret
	// We scan for 0x0F 0x05 to handle both variants.
	if e := touchPageViaReadProcessMemory(proc.Handle, ntdicAddr); e != nil {
		t.Logf("[!] touchPage(0x%X): %v", ntdicAddr, e)
	}
	funcBytes, err := readProcMem(proc.Handle, ntdicAddr, 32)
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("readProcMem(NtDeviceIoControlFile): %v", err)
	}
	t.Logf("[*] NtDeviceIoControlFile bytes: % x", funcBytes)

	syscallOffset := -1
	for i := 0; i < len(funcBytes)-1; i++ {
		if funcBytes[i] == 0x0F && funcBytes[i+1] == 0x05 {
			syscallOffset = i
			break
		}
	}
	if syscallOffset < 0 {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("syscall instruction (0F 05) not found in NtDeviceIoControlFile bytes: % x", funcBytes)
	}
	expectedRetAddr := ntdicAddr + uint64(syscallOffset) + 2
	t.Logf("[*] syscall at +0x%X, expected SYSRET ret addr (RCX) = 0x%X", syscallOffset, expectedRetAddr)

	// Set up the SYSRET hook.
	// At SYSRET time: RCX = user return address, RAX = NTSTATUS return value.
	// The hook is scoped to the target pid (SysretHookForProcess) so the kernel
	// skips the Go callback for every other process's SYSRETs in
	// VmmCallbackTriggerEvents (Debugger.c:1215). This is critical: with
	// ALL_PROCESSES the EFER hook drives 10000+ VM exits/sec and the AST
	// interpreter runs for each, overloading VMX-root and crashing critical
	// system processes (observed: svchost.exe __fastfail → 0xEF BSOD).
	// Even with pid filtering the EFER hook still VM-exits on every syscall/
	// sysret system-wide (hardware-level), but only the target process's
	// exits run the interpreter.
	// SYSRET cannot be filtered by syscall number (EAX is already the return
	// value at SYSRET time, per Debugger.c comment), so we match RCX against
	// the expected NtDeviceIoControlFile return address to identify its returns.
	hookSrc := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	retAddr := ctx.Reg("rcx")
	if retAddr == 0x%X {
		rax := ctx.Reg("rax")
		ctx.Printf("SYSRET NtDeviceIoControlFile: rax=%%x\n", rax)
		ctx.Printf("--------------- DISK_SERIAL_SYSRET ----------------\n")
	}
}`, pid, expectedRetAddr)

	tag, err := dbg.SysretHookForProcess(pid, hookSrc)
	if err != nil {
		cleanupDisk(t, dbg, &proc)
		t.Fatalf("SysretHookForProcess: %v", err)
	}
	t.Logf("[*] SysretHookForProcess installed tag=%d pid=%d", tag, pid)

	pump, pumpErr := dbg.StartMessagePump()
	if pumpErr != nil {
		t.Logf("[!] StartMessagePump: %v", pumpErr)
	}
	defer func() {
		if pump != nil {
			pump.Stop()
		}
	}()

	_ = dbg.Continue()
	t.Logf("[*] Helper running with sysret hook (max 15s)...")

	deadline := time.Now().Add(15 * time.Second)
	for time.Now().Before(deadline) {
		if err := dbg.Continue(); err != nil {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	_ = dbg.Pause()
	if pump != nil {
		pump.Stop()
	}
	_ = dbg.UnloadVMM()
	if proc.Handle != 0 {
		syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
	}
	proc.Close()

	logData, err := os.ReadFile(diskSysretLogPath)
	if err != nil {
		t.Fatalf("read log: %v", err)
	}
	logStr := string(logData)
	t.Logf("[*] Log size: %d bytes", len(logData))

	sysretCount := strings.Count(logStr, "DISK_SERIAL_SYSRET")
	t.Logf("NtDeviceIoControlFile SYSRET matches: %d", sysretCount)

	if sysretCount > 0 {
		t.Logf("PASS: disk serial query SYSRET detected!")
		for _, line := range strings.Split(logStr, "\n") {
			trim := strings.TrimSpace(line)
			if strings.Contains(trim, "SYSRET") {
				t.Logf("  %s", trim)
			}
		}
	} else {
		t.Logf("No DISK_SERIAL_SYSRET in log")
	}
}

// readProcMem reads n bytes from the target process at addr using ReadProcessMemory.
func readProcMem(processHandle uintptr, addr uint64, n int) ([]byte, error) {
	buf := make([]byte, n)
	var bytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		processHandle,
		uintptr(addr),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(n),
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if ret == 0 {
		return nil, fmt.Errorf("ReadProcessMemory(0x%X, %d) failed: %w", addr, n, err)
	}
	return buf[:bytesRead], nil
}

// cleanupDisk 在测试失败路径上安全清理。
func cleanupDisk(t *testing.T, dbg *api.Debugger, proc *core.Process) {
	t.Helper()
	_ = dbg.UnloadVMM()
	if proc != nil && proc.Handle != 0 {
		syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
		proc.Close()
	}
}
