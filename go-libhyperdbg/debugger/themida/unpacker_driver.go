// Package themida — unpacker_driver.go
//
// 驱动集成层：加载 VMM → 设置 EPT hook → 启动目标进程 → 收集回调数据
// → 检测 SetEvent VM 入口和 VM OEP。
//
// 使用 api.Debugger（与 find-oep.go 一致），支持 EptHookSymbol 符号模式。
package themida

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/api"
)

// UnpackerConfig 配置解壳器运行参数。
type UnpackerConfig struct {
	DriverPath string // hyperkd.sys 绝对路径
	ExePath    string // 目标加壳 PE 绝对路径
	LogPath    string // API Logger 输出文件路径
	RunSeconds int    // 让目标进程运行多少秒后暂停

	// DLL 路径（用于从 PE 文件中解析导出地址）
	Kernel32Path   string
	NtdllPath      string
	KernelbasePath string
}

// WLSection 记录 Themida/WinLicense 的 WL 段信息。
type WLSection struct {
	BaseAddress uint64
	RegionSize  uint64
}

// UnpackerResult 是解壳器运行的完整结果。
type UnpackerResult struct {
	LogPath         string
	APICalls        []APILogEntry
	OEPHints        []uint32
	SetEventEntry   *SetEventEntry
	IOMarkerAddress uint32
	WLSection       WLSection
	VMType          int // 0=Old CISC, 1=New CISC, 3=RISC
	VMOEP           *VMOEPResult
}

// Unpacker 封装驱动调试器的解壳会话。
type Unpacker struct {
	cfg UnpackerConfig
	dbg *api.Debugger
}

// NewUnpacker 创建解壳器实例。
func NewUnpacker(cfg UnpackerConfig) *Unpacker {
	return &Unpacker{cfg: cfg}
}

// Run 执行完整解壳流程：
//  1. 创建 api.Debugger + LoadVMM + LogOpen
//  2. 启动目标进程（挂起状态）
//  3. 用目标进程 PID 获取 32-bit DLL 基址（WOW64）
//  4. 设置 EPT hook（EptHookForProcess，限定到目标进程）
//  5. 两次 Continue（先到入口点，再继续运行）
//  6. 等待 RunSeconds 后 Pause
//  7. 卸载驱动
//
// OEP 不是在进程启动时直接找的——而是通过 hook API 收集调用日志
// （API Logger 格式："Call from: X | API: Y | NAME: Z"），然后用
// unpacker14.go 的 ParseAPILog 解析日志，从中提取 SetEvent 入口和
// IO Marker，最终推导 VM OEP。这对应 OllyDbg 脚本 L1430+ 的流程：
// 监控 VirtualAlloc 找 WL Section → 在 WL Section 中扫描 VM OEP 签名。
func (u *Unpacker) Run(ctx context.Context) (UnpackerResult, error) {
	var result UnpackerResult
	result.LogPath = u.cfg.LogPath

	// 1. 创建 api.Debugger
	dbg, err := api.New(api.WithOutput(os.Stderr))
	if err != nil {
		return result, fmt.Errorf("api.New: %w", err)
	}
	u.dbg = dbg
	defer dbg.Close()

	// LoadVMM 包含: 安装驱动 → 打开设备 → IOCTL_INIT_VMM
	if err := dbg.LoadVMM(ctx, u.cfg.DriverPath); err != nil {
		return result, fmt.Errorf("LoadVMM: %w", err)
	}
	// 注意: 不用 defer UnloadVMM，而是手动调用，确保执行顺序：
	// UnloadVMM (TERMINATE_VMX) → TerminateProcess → Close

	if u.cfg.LogPath != "" {
		if err := dbg.LogOpen(u.cfg.LogPath); err != nil {
			fmt.Fprintf(os.Stderr, "[!] LogOpen failed: %v\n", err)
		}
		defer dbg.LogClose()
	}

	// 2. 启动目标进程（挂起状态），获取 PID 和 Handle
	//    必须先启动进程才能获取 32-bit DLL 基址（WOW64 目标的 DLL
	//    基址与 64-bit 调试器进程不同，不能用 GetModuleHandleW）。
	proc, err := dbg.StartProcess(ctx, u.cfg.ExePath)
	if err != nil {
		dbg.UnloadVMM(ctx)
		return result, fmt.Errorf("StartProcess: %w", err)
	}
	pid := proc.Pid
	fmt.Fprintf(os.Stderr, "[*] Started %s (pid=%d)\n", u.cfg.ExePath, pid)

	// 3. 第一次 Continue — 运行到入口点
	//    StartProcess 以 CREATE_SUSPENDED 创建进程并 attach，内核在
	//    第一条指令处暂停。Continue 让进程运行，但 PEB monitor EPT
	//    hook 会在 PEB 被访问时再次暂停进程（此时 loader 尚未完成）。
	if err := dbg.Continue(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[!] 1st Continue (to entry point) failed: %v\n", err)
	}

	// 4. 第二次 Continue — 跳过 PEB 暂停，让进程运行到入口点之后
	//    对应 OllyDbg 脚本的 esto（运行到断点）。必须让进程实际执行
	//    Themida loader 代码，才能让 kernelbase.dll 等 DLL 的代码页
	//    被 demand-paging 映射到物理内存。如果进程在 PEB 访问时暂停，
	//    kernelbase 的代码页还未 fault in，EptHookForProcess 会因
	//    MmGetPhysicalAddress 返回 0 而失败 (0xC0000005)。
	if err := dbg.Continue(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[!] 2nd Continue (past entry point) failed: %v\n", err)
	}

	// 5. 获取 32-bit DLL 基址（在 warmup 之前计算 hook 地址）
	//    SuperRecovery.exe 是 32 位（WOW64），其 ntdll/kernelbase
	//    在 SysWOW64 目录，基址与 64 位进程不同。
	kernelbaseBase, err := getModuleBaseForWow64Target(pid, proc.Handle, "kernelbase.dll")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] getModuleBaseForWow64Target(kernelbase.dll): %v\n", err)
	}
	ntdllBase, err := getModuleBaseForWow64Target(pid, proc.Handle, "ntdll.dll")
	if err != nil {
		fmt.Fprintf(os.Stderr, "[!] getModuleBaseForWow64Target(ntdll.dll): %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "[*] 32-bit bases: kernelbase=0x%X ntdll=0x%X\n", kernelbaseBase, ntdllBase)

	// 5b. 计算 hook 地址 = 32-bit 基址 + RVA
	if ntdllBase > 0xFFFFFFFF {
		fmt.Fprintf(os.Stderr, "[!] ntdll base 0x%X is 64-bit, falling back to 32-bit base from running WOW64 process\n", ntdllBase)
		if base, e := getModuleBaseFromRunningWow64("ntdll.dll"); e == nil && base <= 0xFFFFFFFF {
			ntdllBase = base
			fmt.Fprintf(os.Stderr, "[*] 32-bit ntdll base from running WOW64: 0x%X\n", ntdllBase)
		}
	}

	var setEventAddr, virtualAllocAddr, rtlAllocateHeapAddr uint64
	if kernelbaseBase != 0 && kernelbaseBase <= 0xFFFFFFFF {
		if rva, _, e := resolveExportRVA(u.cfg.KernelbasePath, "SetEvent"); e == nil {
			setEventAddr = kernelbaseBase + rva
		} else {
			fmt.Fprintf(os.Stderr, "[!] resolveExportRVA(SetEvent): %v\n", e)
		}
		if rva, _, e := resolveExportRVA(u.cfg.KernelbasePath, "VirtualAlloc"); e == nil {
			virtualAllocAddr = kernelbaseBase + rva
		} else {
			fmt.Fprintf(os.Stderr, "[!] resolveExportRVA(VirtualAlloc): %v\n", e)
		}
	}
	if ntdllBase != 0 && ntdllBase <= 0xFFFFFFFF {
		if rva, _, e := resolveExportRVA(u.cfg.NtdllPath, "RtlAllocateHeap"); e == nil {
			rtlAllocateHeapAddr = ntdllBase + rva
		} else {
			fmt.Fprintf(os.Stderr, "[!] resolveExportRVA(RtlAllocateHeap): %v\n", e)
		}
	}
	fmt.Fprintf(os.Stderr, "[*] Hook addresses: SetEvent=0x%X VirtualAlloc=0x%X RtlAllocateHeap=0x%X\n",
		setEventAddr, virtualAllocAddr, rtlAllocateHeapAddr)

	// 6. Warmup: 让进程运行，直到 kernelbase 页面 fault in（demand-paging）
	//    Themida loader 会调用 kernelbase!SetEvent / VirtualAlloc /
	//    ntdll!RtlAllocateHeap 等 API，这些调用会 fault in 对应的
	//    代码页，使 PTE 变为 present。
	//
	//    关键：必须在 loader 调用目标 API 之前设置 EPT hook。因此
	//    warmup 循环在每次 Continue 后检查目标页是否可读，一旦全部
	//    可读就立即暂停并设置 hook。Themida loader 会多次调用
	//    VirtualAlloc / SetEvent，所以即使错过前几次调用，后续调用
	//    仍会触发 hook。
	fmt.Fprintf(os.Stderr, "[*] Warmup: running until DLL pages fault in...\n")
	warmupMax := 20 // max 20 * 300ms = 6s
	pagesReady := false
	for i := 0; i < warmupMax; i++ {
		// Resume if paused (PEB hook, exception, etc.)
		_ = dbg.Continue(ctx)
		time.Sleep(300 * time.Millisecond)

		// Check if all target pages are faulted in
		allReady := true
		for _, addr := range []uint64{setEventAddr, virtualAllocAddr, rtlAllocateHeapAddr} {
			if addr == 0 {
				continue
			}
			if e := touchPageViaReadProcessMemory(proc.Handle, addr); e != nil {
				allReady = false
				break
			}
		}
		if allReady {
			pagesReady = true
			fmt.Fprintf(os.Stderr, "[*] All target pages faulted in after %d iterations\n", i+1)
			break
		}
	}
	if !pagesReady {
		fmt.Fprintf(os.Stderr, "[!] Warmup: not all pages faulted in after %d iterations, continuing anyway\n", warmupMax)
	}

	// 7. 暂停进程，准备设置 EPT hook
	_ = dbg.Pause(ctx) // may be "already paused" — that's fine
	fmt.Fprintf(os.Stderr, "[*] Process paused, setting up hooks\n")

	// 5.5 触发页面映射（demand-paging）
	//    Windows 的 DLL 页是按需分页的（demand-paged）。如果目标进程
	//    还没访问过 kernelbase.dll/ntdll.dll 的某些代码页，PTE 是
	//    not-present，MmGetPhysicalAddress 返回 0，导致 EptHookForProcess
	//    地址验证失败（DEBUGGER_ERROR_INVALID_ADDRESS = 0xC0000005）。
	//
	//    注意：HyperDbg 的 ReadMemory（IOCTL）通过 MmGetPhysicalAddress +
	//    MmMapIoSpaceEx 读物理内存，不触发页面错误，无法把 demand-paged
	//    页映射到物理内存。HyperDbg 的 .pagein 命令在本地调试模式下也是
	//    桩函数（只打印日志不注入 #PF）。
	//
	//    必须用 Win32 ReadProcessMemory 访问目标进程的虚拟地址，触发 #PF
	//    让 OS 把页从磁盘读入物理内存，PTE 才会变成 present。
	for _, addr := range []uint64{setEventAddr, virtualAllocAddr, rtlAllocateHeapAddr} {
		if addr == 0 {
			continue
		}
		if e := touchPageViaReadProcessMemory(proc.Handle, addr); e != nil {
			fmt.Fprintf(os.Stderr, "[!] touchPage(0x%X) failed: %v\n", addr, e)
		} else {
			fmt.Fprintf(os.Stderr, "[*] Pre-touched page at 0x%X (via ReadProcessMemory)\n", addr)
		}
	}

	// 6. 设置 EPT hook（EptHookForProcess，限定到目标进程）
	//    不用全局 EptHook（ProcessId=0xFFFFFFFF），因为全局 hook 会在
	//    所有进程的 API 调用时触发，导致系统级死锁/BSOD。
	//    hook 回调生成 API Logger 格式日志：
	//    "Call from: {ret} | API: {addr} | NAME: {name}"
	//    这与 OllyDbg 脚本的 API Logger 格式一致，可用 ParseAPILog 解析。

	// RtlAllocateHeap hook
	if rtlAllocateHeapAddr != 0 {
		rahCode := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("Call from: %%x | API: %%x | NAME: RtlAllocateHeap\n", ret, 0x%X)
}`, pid, rtlAllocateHeapAddr)
		if tagRAH, err := dbg.EptHookForProcess(ctx, rtlAllocateHeapAddr, pid, rahCode); err != nil {
			fmt.Fprintf(os.Stderr, "[!] EptHookForProcess(RtlAllocateHeap) failed: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "[*] Hooked RtlAllocateHeap tag=%d\n", tagRAH)
		}
	}

	// VirtualAlloc hook（非致命）
	// Records the caller's return address via SetCtxVar so the SetEvent hook
	// (and future mem-write callback) can reference it. In the OllyDbg script
	// (L450-454) the WL section is derived from this return address via
	// gmemi MEMORYBASE/MEMORYSIZE; here we store the raw return address and
	// defer region-base resolution to user-mode post-processing.
	if virtualAllocAddr != 0 {
		vaCode := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("Call from: %%x | API: %%x | NAME: VirtualAlloc\n", ret, 0x%X)
	ctx.SetCtxVar("va_ret", ret)
}`, pid, virtualAllocAddr)
		if tagVA, err := dbg.EptHookForProcess(ctx, virtualAllocAddr, pid, vaCode); err != nil {
			fmt.Fprintf(os.Stderr, "[!] EptHookForProcess(VirtualAlloc) failed (non-fatal): %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "[*] Hooked VirtualAlloc tag=%d\n", tagVA)
		}
	}

	// SetEvent hook — detects Themida VM entry (PUSH imm32; JMP rel32 pattern
	// at the caller's return address) and logs SETEVENT_ENTRY_ADDRESS.
	// Mirrors the CHECK_OLD_WL branch of the OllyDbg script (L559-580):
	//
	//	cmp [ADDR],   68  ; PUSH imm32
	//	cmp [ADDR+5], E9  ; JMP rel32
	//	VM_PUSH = [ADDR+1]
	//	VM_JUMP = ADDR + 0xA + signed([ADDR+6])
	//
	// We read 8 bytes at ADDR and ADDR+5 via ReadMemQword, then mask:
	//
	//	q1 = [ADDR..ADDR+7]    → byte0=0x68, bytes1-4=imm32
	//	q2 = [ADDR+5..ADDR+12] → byte0=0xE9, bytes1-4=rel32
	//
	// uint64 wraparound + final & 0xFFFFFFFF handles the signed rel32.
	// If the read fails (page not mapped), ReadMemQword returns 0 and the
	// pattern check simply fails — we still log the normal Call-from line.
	if setEventAddr != 0 {
		seCode := fmt.Sprintf(`package hook
func hook(ctx *HookCtx) {
	if ctx.GetPid() != %d { return }
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("Call from: %%x | API: %%x | NAME: SetEvent\n", ret, 0x%X)
	q1 := ctx.ReadMemQword(ret)
	q2 := ctx.ReadMemQword(ret + 5)
	if (q1 & 0xFF) == 0x68 && (q2 & 0xFF) == 0xE9 {
		push := (q1 >> 8) & 0xFFFFFFFF
		disp := (q2 >> 8) & 0xFFFFFFFF
		jump := (ret + 10 + disp) & 0xFFFFFFFF
		ctx.Printf("--------------- SETEVENT_ENTRY_ADDRESS ----------------\n")
		ctx.Printf("Address: %%x | PUSH %%x | JUMP %%x\n", ret, push, jump)
	}
}`, pid, setEventAddr)
		if tagSE, err := dbg.EptHookForProcess(ctx, setEventAddr, pid, seCode); err != nil {
			fmt.Fprintf(os.Stderr, "[!] EptHookForProcess(SetEvent) failed: %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "[*] Hooked SetEvent tag=%d\n", tagSE)
		}
	}

	// 7b. Start the kernel→user message pump BEFORE resuming the debuggee.
	//     Hook callbacks fire in VMX-root and write their ctx.Printf output
	//     to a kernel ring buffer; the kernel then completes the pending
	//     IRP issued by the pump (IOCTL_REGISTER_EVENT, Type=IRP_BASED).
	//     Without a pump running, the messages pile up in the ring buffer
	//     but never reach user-mode, so test-oep.log stays empty.
	//     The pump opens a DEDICATED device handle so the blocking IRP
	//     does not stall the main IOCTL handle used by Continue/Pause.
	pump, pumpErr := dbg.StartMessagePump(ctx)
	if pumpErr != nil {
		fmt.Fprintf(os.Stderr, "[!] StartMessagePump failed: %v (hooks will fire but log will be empty)\n", pumpErr)
	} else {
		fmt.Fprintf(os.Stderr, "[*] Message pump started\n")
	}
	// Safety net: stop the pump on any return path where it wasn't stopped
	// explicitly. Stop is idempotent, so explicit calls below are no-ops
	// when the defer runs.
	defer func() {
		if pump != nil {
			pump.Stop()
		}
	}()

	// 8. 第三次 Continue — 继续运行，让 Themida loader 执行并调用 hooked API
	//    此时 EPT hook 已设置，Themida loader 调用 SetEvent/VirtualAlloc/
	//    RtlAllocateHeap 时会触发 hook 回调，输出 API Logger 格式日志。
	if err := dbg.Continue(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[!] 3rd Continue (resume with hooks) failed: %v\n", err)
	}

	// 9. 等待目标运行（Themida loader 需要 20-30 秒完成解壳）
	runSec := u.cfg.RunSeconds
	if runSec <= 0 {
		runSec = 30
	}
	fmt.Fprintf(os.Stderr, "[*] Running for %d seconds...\n", runSec)
	select {
	case <-time.After(time.Duration(runSec) * time.Second):
	case <-ctx.Done():
		// Stop the pump BEFORE UnloadVMM — Stop needs the main device
		// handle to send IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL,
		// and UnloadVMM stops the driver service which makes that handle
		// unusable.
		if pump != nil {
			pump.Stop()
		}
		dbg.UnloadVMM(ctx)
		return result, ctx.Err()
	}

	// 10. 暂停进程，检查 hook 输出
	if err := dbg.Pause(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "[!] Pause failed: %v\n", err)
	}
	fmt.Fprintf(os.Stderr, "[*] Paused. Check %s for hook output.\n", u.cfg.LogPath)

	// Stop the pump BEFORE UnloadVMM (see ctx.Done() comment above).
	if pump != nil {
		pump.Stop()
		fmt.Fprintf(os.Stderr, "[*] Message pump stopped\n")
	}

	// 清理顺序：先 UnloadVMM（含 TERMINATE_VMX），再终止进程
	dbg.UnloadVMM(ctx)
	if proc.Handle != 0 {
		syscall.TerminateProcess(syscall.Handle(proc.Handle), 1)
		fmt.Fprintf(os.Stderr, "[*] Process terminated (cleanup)\n")
	}
	proc.Close()

	return result, nil
}

func rvaToOff(data []byte, coffOff uint32, numSec uint16, sizeOpt uint16, rva uint32) (uint32, bool) {
	secOff := coffOff + 20 + uint32(sizeOpt)
	for i := uint16(0); i < numSec; i++ {
		off := secOff + uint32(i)*40
		if int(off)+40 > len(data) {
			continue
		}
		vSize := le32(data, off+8)
		vAddr := le32(data, off+12)
		rawOff := le32(data, off+20)
		if rva >= vAddr && rva < vAddr+vSize {
			return rawOff + (rva - vAddr), true
		}
	}
	return 0, false
}

func le16(d []byte, off uint32) uint16 { return uint16(d[off]) | uint16(d[off+1])<<8 }
func le32(d []byte, off uint32) uint32 {
	return uint32(d[off]) | uint32(d[off+1])<<8 | uint32(d[off+2])<<16 | uint32(d[off+3])<<24
}

func le64(d []byte, off uint32) uint64 {
	return uint64(le32(d, off)) | uint64(le32(d, off+4))<<32
}

var (
	kernel32   = syscall.NewLazyDLL("kernel32.dll")
	getModProc = kernel32.NewProc("GetModuleHandleW")
)

// resolveRuntimeExport 从当前进程的运行时 DLL 基址 + PE 文件中的 RVA 计算实际导出地址。
// 这解决了 ASLR 导致的地址偏移问题。
func resolveRuntimeExport(pePath, funcName string) (uint64, error) {
	// 从 PE 文件解析 RVA
	rva, peBase, err := resolveExportRVA(pePath, funcName)
	if err != nil {
		return 0, err
	}

	// 获取 DLL 在当前进程中的实际基址
	dllName := dllBaseName(pePath)
	basePtr, _, _ := getModProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dllName))))
	if basePtr == 0 {
		// DLL 未加载到当前进程，回退到 PE 文件中的 ImageBase
		return peBase + rva, nil
	}

	return uint64(basePtr) + rva, nil
}

// dllBaseName 从完整路径中提取 DLL 文件名（如 "C:\Windows\System32\ntdll.dll" → "ntdll.dll"）
func dllBaseName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '\\' || path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}

// resolveExportRVA 只解析 PE 文件中的 RVA 和 ImageBase（不返回绝对地址）
func resolveExportRVA(pePath, funcName string) (rva uint64, imageBase uint64, err error) {
	data, err := os.ReadFile(pePath)
	if err != nil {
		return 0, 0, fmt.Errorf("read %s: %w", pePath, err)
	}
	if len(data) < 64 || data[0] != 'M' || data[1] != 'Z' {
		return 0, 0, fmt.Errorf("%s: not a PE file", pePath)
	}
	e_lfanew := le32(data, 0x3C)
	if int(e_lfanew)+4 > len(data) || data[e_lfanew] != 'P' || data[e_lfanew+1] != 'E' {
		return 0, 0, fmt.Errorf("%s: bad PE signature", pePath)
	}
	coffOff := e_lfanew + 4
	numSec := le16(data, coffOff+2)
	sizeOpt := le16(data, coffOff+16)
	optOff := coffOff + 20
	magic := le16(data, optOff)
	var exportRVA uint32
	if magic == 0x20B { // PE32+
		imageBase = le64(data, optOff+24)
		exportRVA = le32(data, optOff+112)
	} else { // PE32
		imageBase = uint64(le32(data, optOff+28))
		exportRVA = le32(data, optOff+96)
	}
	if exportRVA == 0 {
		return 0, 0, fmt.Errorf("%s: no export directory", pePath)
	}
	expOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, exportRVA)
	if !ok || int(expOff)+40 > len(data) {
		return 0, 0, fmt.Errorf("%s: export dir truncated", pePath)
	}
	numNames := le32(data, expOff+24)
	addrFuncs := le32(data, expOff+28)
	addrNames := le32(data, expOff+32)
	addrOrds := le32(data, expOff+36)
	namesOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, addrNames)
	if !ok {
		return 0, 0, fmt.Errorf("%s: name pointers bad", pePath)
	}
	ordsOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, addrOrds)
	if !ok {
		return 0, 0, fmt.Errorf("%s: ordinal pointers bad", pePath)
	}
	funcsOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, addrFuncs)
	if !ok {
		return 0, 0, fmt.Errorf("%s: function pointers bad", pePath)
	}
	for i := uint32(0); i < numNames; i++ {
		nameRVA := le32(data, namesOff+i*4)
		nOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, nameRVA)
		if !ok {
			continue
		}
		end := nOff
		for end < uint32(len(data)) && data[end] != 0 {
			end++
		}
		if string(data[nOff:end]) != funcName {
			continue
		}
		ord := le16(data, ordsOff+i*2)
		funcRVA := le32(data, funcsOff+uint32(ord)*4)
		return uint64(funcRVA), imageBase, nil
	}
	return 0, 0, fmt.Errorf("%s: export %s not found", pePath, funcName)
}
