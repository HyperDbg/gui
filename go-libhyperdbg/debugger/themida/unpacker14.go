// Package themida — unpacker14.go
//
// Go replica of "Themida - WinLicense Ultra Unpacker 1.4" (14147-line OllyDbg script).
// Pure logic — no cgo, no driver dependencies.
package themida

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ---------------------------------------------------------------------------
// Constants from themida.go
// ---------------------------------------------------------------------------

// ThemidaSystemDllThreshold 对应 find-oep.ds 的 if (caller < 0x70000000)。
// 低于该阈值的地址视为非系统 DLL（即 Themida 用户态代码）。
const ThemidaSystemDllThreshold uint64 = 0x70000000

// OllyDbg 脚本 L97-103 使用的三个 magic dword。
//
// 字符串 "RtlAllocateHeap" 在小端存储下前 12 字节为：
//
//	"RtlA" -> 0x416C7452
//	"lloc" -> 0x636F6C6C
//	"ateH" -> 0x48657461
const (
	wlsecMagic0 uint32 = 0x416C7452 // "RtlA"
	wlsecMagic1 uint32 = 0x636F6C6C // "lloc"
	wlsecMagic2 uint32 = 0x48657461 // "ateH"
)

// OllyDbg 脚本 L176-180 的 Stack AntiDump xor 常量。
const (
	StackAntiDumpXorOld uint32 = 0x8647A6B4 // xor ecx, 8647A6B4
	StackAntiDumpXorNew uint32 = 0x7647A6B4 // xor edx, 7647A6B4
)

// ---------------------------------------------------------------------------
// Constants from setevent.go
// ---------------------------------------------------------------------------

// SetEvent 字节模式常量。
const (
	opPushImm32       byte = 0x68 // PUSH imm32
	opJmpRel32        byte = 0xE9 // JMP rel32
	seteventOldWLSize int  = 10   // PUSH + JMP = 5 + 5
	seteventNewWLSize int  = 15   // PUSH + PUSH + JMP = 5 + 5 + 5
)

// ---------------------------------------------------------------------------
// Constants from modbase_windows.go
// ---------------------------------------------------------------------------

const (
	th32csSnapModule   = 0x00000008
	th32csSnapModule32 = 0x00000010
	maxModuleName32    = 255
	maxPath            = 260
)

// ---------------------------------------------------------------------------
// Types from themida.go
// ---------------------------------------------------------------------------

// MemoryRegion 模拟一段连续的进程内存区域（对应 VirtualQuery 的
// MEMORY_BASIC_INFORMATION）。这里只保留解壳逻辑关心的三个字段。
type MemoryRegion struct {
	BaseAddress uint64
	RegionSize  uint64
	Protect     uint32
}

// MemoryMap 是一组有序的内存区域，模拟进程内存布局。区域之间可能存在
// 空洞（即未提交的地址范围）。
type MemoryMap []MemoryRegion

// FindOepResult 封装 find-oep.ds 的输出。当 hook 命中并过滤通过后，脚本
// 会打印 "THEMIDA PE FOUND" 并附带 caller、PE 基址、PE 大小以及由此推出
// 的 OEP 监控范围。
type FindOepResult struct {
	Caller    uint64    // 触发 VirtualAlloc 的返回地址
	ThemidaPE uint64    // gmemi(caller, MEMORYBASE) 得到的段基址
	PeSize    uint64    // gmemi(pe, MEMORYSIZE) 得到的段大小
	OepRange  [2]uint64 // [pe, pe+size)，供后续 bprm/cmp eip 使用
}

// ---------------------------------------------------------------------------
// Types from setevent.go
// ---------------------------------------------------------------------------

// SetEventEntry 描述一处 Themida SetEvent VM 入口跳板。
//
// Address 是 SetEvent 调用方的返回地址（即 [esp]，CALL SetEvent 之后下一
// 条指令的 VA）。Push / Push2 是跳板 PUSH 的立即数；Jump 是 JMP rel32 计
// 算出的绝对目标地址。IsNewWL 为 true 时表示 New WL CISC 模式（双 PUSH）。
type SetEventEntry struct {
	Address uint32
	Push    uint32
	Push2   uint32 // 仅 New WL 模式有效
	Jump    uint32
	IsNewWL bool
}

// ---------------------------------------------------------------------------
// Types from apilog.go
// ---------------------------------------------------------------------------

// APILogEntry 是日志中一条 "Call from: X | API: Y | NAME: Z" 记录。
//
// Kind 为 "EX"（普通 Export 调用）或 "GPA"（GetProcAddress 调用）。
// Caller / API 是 32 位 VA（hex，无前缀）；Name 是 API 名字符串。
type APILogEntry struct {
	Kind   string
	Caller uint32
	API    uint32
	Name   string
}

// APILog 是解析后的 API Logger 文件内容。
//
// SetEventEntry 为 nil 表示日志中未出现 SETEVENT_ENTRY_ADDRESS 块
// （目标未使用 SetEvent AntiDump）。IOMarkerAddress 为 0 表示未出现
// I_O_MARKER_ADDRESS 块。
type APILog struct {
	Entries         []APILogEntry
	SetEventEntry   *SetEventEntry
	IOMarkerAddress uint32
}

// ---------------------------------------------------------------------------
// Types from modbase_windows.go
// ---------------------------------------------------------------------------

type moduleEntry32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  uintptr
	ModBaseSize  uint32
	ModuleHandle uintptr
	Module       [maxModuleName32 + 1]uint16
	ExePath      [maxPath]uint16
}

// ---------------------------------------------------------------------------
// New types for unpacker14
// ---------------------------------------------------------------------------

// VMOEPResult holds the parsed OVERVIEW OEP fields for a Themida VM entry.
type VMOEPResult struct {
	Addr  uint32
	Align uint32
	Push  uint32
	Jump  uint32
}

// ---------------------------------------------------------------------------
// Functions from themida.go
// ---------------------------------------------------------------------------

// GmemiMemoryBase 对应 HyperDbg 的 gmemi(addr, 0) /
// gmemi(addr, MEMORYBASE)：向下扫描页边界找到包含 addr 的区域基址。
//
// 实现上等价于遍历 MemoryMap，找到 BaseAddress <= addr <
// BaseAddress+RegionSize 的区域并返回其 BaseAddress。当 addr 落在空洞
// （任何区域之外）时返回 (0, false)。
func GmemiMemoryBase(m MemoryMap, addr uint64) (uint64, bool) {
	for _, r := range m {
		if r.RegionSize == 0 {
			continue
		}
		if addr >= r.BaseAddress && addr < r.BaseAddress+r.RegionSize {
			return r.BaseAddress, true
		}
	}
	return 0, false
}

// GmemiMemorySize 对应 HyperDbg 的 gmemi(base, 1) /
// gmemi(base, MEMORYSIZE)：返回包含 base 的区域大小。
//
// 注意 base 必须落在某个区域的 [BaseAddress, BaseAddress+RegionSize)
// 区间内（通常 base 就是 GmemiMemoryBase 的返回值），否则返回 (0, false)。
func GmemiMemorySize(m MemoryMap, base uint64) (uint64, bool) {
	for _, r := range m {
		if r.RegionSize == 0 {
			continue
		}
		if base >= r.BaseAddress && base < r.BaseAddress+r.RegionSize {
			return r.RegionSize, true
		}
	}
	return 0, false
}

// IsThemidaCaller 对应 find-oep.ds 的 if (caller < 0x70000000)。
// 返回 true 表示 caller 是非系统 DLL（Themida 代码）。
//
// 0x70000000 是脚本里硬编码的系统 DLL 阈值：典型的 ntdll/kernel32/
// kernelbase/user32 等系统模块在 64 位 Windows 上多映射到 0x7FFE0000
// 附近或更高地址，因此 < 0x70000000 的 caller 视为被保护程序自身。
func IsThemidaCaller(caller uint64) bool {
	return caller < ThemidaSystemDllThreshold
}

// IsRtlAllocateHeapMagic 对应 OllyDbg 脚本 L97-103 的 magic 验证。
//
// 脚本在 VirtualAlloc hook 里读取 [esp+08]（指向一个字符串），然后依次
// 比较：
//
//	cmp [[esp+08]],   416C7452  ; "RtlA"
//	cmp [TEST+04],    636F6C6C  ; "lloc"
//	cmp [TEST+08],    48657461  ; "ateH"
//
// 三段拼起来即字符串 "RtlAllocateH"（"RtlAllocateHeap" 的前 12 字节），
// 用以确认 WLSEC 指向的调用栈帧确实是 RtlAllocateHeap 触发的。
//
// 入参 data 是 [esp+08] 指向的字节缓冲；要求 len(data) >= 12。当数据
// 不足 12 字节或任一 dword 不匹配时返回 false。
func IsRtlAllocateHeapMagic(data []byte) bool {
	if len(data) < 12 {
		return false
	}
	if binary.LittleEndian.Uint32(data[0:4]) != wlsecMagic0 {
		return false
	}
	if binary.LittleEndian.Uint32(data[4:8]) != wlsecMagic1 {
		return false
	}
	if binary.LittleEndian.Uint32(data[8:12]) != wlsecMagic2 {
		return false
	}
	return true
}

// IsInOepRange 对应 bprm PE, CODESIZE; esto; cmp eip, PE 的判定：
// 当执行监控期间 @ip 落在 [pe, pe+size) 范围内时视为 NEAR-OEP。
//
// 范围是左闭右开 [pe, pe+size)，与 GmemiMemoryBase/Size 的范围语义一致。
// size 为 0 时直接返回 false（无法构成有效区间）。
func IsInOepRange(ip, pe, size uint64) bool {
	if size == 0 {
		return false
	}
	return ip >= pe && ip < pe+size
}

// StackAntiDumpCalc 对应 OllyDbg 脚本 L176-180 的两条 xor 算式：
//
//	mov ecx, SAD_OLD
//	xor ecx, 8647A6B4    ; SA_1 = SAD_OLD ^ 0x8647A6B4
//	mov edx, SAD_NEW
//	xor edx, 7647A6B4    ; SA_2 = SAD_NEW ^ 0x7647A6B4
//
// Themida 在还原被 AntiDump 破坏的栈帧时会用到这两条算式。返回还原后的
// 两个 32 位值 (sa1, sa2)。
func StackAntiDumpCalc(sadOld, sadNew uint32) (sa1, sa2 uint32) {
	sa1 = sadOld ^ StackAntiDumpXorOld
	sa2 = sadNew ^ StackAntiDumpXorNew
	return sa1, sa2
}

// FindThemidaPE 对应 find-oep.ds 的主逻辑：
//
//  1. 用 IsThemidaCaller 过滤系统 DLL；不通过则返回 false。
//  2. GmemiMemoryBase 找到 caller 所在区域的段基址。
//  3. GmemiMemorySize 取该段大小。
//  4. 组装 FindOepResult，OepRange = [pe, pe+size)。
//
// 任一步骤失败（系统 DLL、地址落在空洞、段大小为 0）都返回 (_, false)。
func FindThemidaPE(m MemoryMap, caller uint64) (FindOepResult, bool) {
	var res FindOepResult
	if !IsThemidaCaller(caller) {
		return res, false
	}
	pe, ok := GmemiMemoryBase(m, caller)
	if !ok {
		return res, false
	}
	size, ok := GmemiMemorySize(m, pe)
	if !ok || size == 0 {
		return res, false
	}
	res.Caller = caller
	res.ThemidaPE = pe
	res.PeSize = size
	res.OepRange[0] = pe
	res.OepRange[1] = pe + size
	return res, true
}

// ---------------------------------------------------------------------------
// Functions from setevent.go
// ---------------------------------------------------------------------------

// ParseSetEventEntry 解码 ADDR 处的字节模式，返回 SetEvent VM 入口跳板。
//
// code 是从 ADDR 起的字节序列；至少需要 10 字节（Old WL）或 15 字节
// （New WL）。当字节模式不匹配任一 CISC 变体时返回 (_, false)。
//
// 算法对应 Ultra Unpacker 1.4 脚本 L559-622（CHECK_OLD_WL / CHECK_NEW_WL）。
func ParseSetEventEntry(addr uint32, code []byte) (SetEventEntry, bool) {
	var e SetEventEntry
	e.Address = addr

	// Old WL: 68 XX XX XX XX E9 YY YY YY YY
	if len(code) >= seteventOldWLSize &&
		code[0] == opPushImm32 && code[5] == opJmpRel32 {
		e.Push = binary.LittleEndian.Uint32(code[1:5])
		disp := int32(binary.LittleEndian.Uint32(code[6:10]))
		e.Jump = addr + uint32(seteventOldWLSize) + uint32(disp)
		e.IsNewWL = false
		return e, true
	}

	// New WL: 68 XX XX XX XX 68 ZZ ZZ ZZ ZZ E9 YY YY YY YY
	if len(code) >= seteventNewWLSize &&
		code[0] == opPushImm32 && code[5] == opPushImm32 && code[10] == opJmpRel32 {
		e.Push = binary.LittleEndian.Uint32(code[1:5])
		e.Push2 = binary.LittleEndian.Uint32(code[6:10])
		disp := int32(binary.LittleEndian.Uint32(code[11:15]))
		e.Jump = addr + uint32(seteventNewWLSize) + uint32(disp)
		e.IsNewWL = true
		return e, true
	}

	return e, false
}

// JumpDisplacement 返回 JMP rel32 的有符号位移。用于反向校验
// Jump == Address + base + signed_disp（base = 10 / 15）。
func (e SetEventEntry) JumpDisplacement() int32 {
	if e.IsNewWL {
		return int32(e.Jump - e.Address - uint32(seteventNewWLSize))
	}
	return int32(e.Jump - e.Address - uint32(seteventOldWLSize))
}

// Format 按Ultra Unpacker 1.4 脚本的输出格式返回 SetEvent 入口描述字符串。
//
//	Old WL: "Address: {ADDR} | PUSH {VM_PUSH} | JUMP {VM_JUMP}"
//	New WL: "Address: {ADDR} | PUSH {VM_PUSH} | PUSH {VM_PUSH2} | JUMP {VM_JUMP}"
func (e SetEventEntry) Format() string {
	if e.IsNewWL {
		return fmt.Sprintf("Address: %X | PUSH %X | PUSH %X | JUMP %X",
			e.Address, e.Push, e.Push2, e.Jump)
	}
	return fmt.Sprintf("Address: %X | PUSH %X | JUMP %X",
		e.Address, e.Push, e.Jump)
}

// ---------------------------------------------------------------------------
// Functions from apilog.go
// ---------------------------------------------------------------------------

// ParseAPILog 解析完整的 API Logger 文本。
//
// 容错策略：无法识别的行直接跳过；能识别但格式错误的行返回 error。
// 这样可以在日志首尾有空白/分隔符变化时仍然工作。
func ParseAPILog(data []byte) (APILog, error) {
	var log APILog
	sc := bufio.NewScanner(bytes.NewReader(data))
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)

	currentKind := ""
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		// 检测 EX / GPA 段头
		if strings.Contains(line, "---EX---") {
			currentKind = "EX"
			continue
		}
		if strings.Contains(line, "---GPA---") {
			currentKind = "GPA"
			continue
		}

		// "Call from: X | API: Y | NAME: Z"
		if strings.HasPrefix(line, "Call from:") {
			e, err := parseCallLine(line, currentKind)
			if err != nil {
				return log, err
			}
			log.Entries = append(log.Entries, e)
			continue
		}

		// "Address: X | PUSH Y | JUMP Z"  或
		// "Address: X | PUSH Y | PUSH Z | JUMP W"
		if strings.HasPrefix(line, "Address:") {
			se, err := parseSetEventLine(line)
			if err != nil {
				return log, err
			}
			log.SetEventEntry = &se
			continue
		}

		// "I_O_MARKER_ADDRESS VA: X"
		if strings.HasPrefix(line, "I_O_MARKER_ADDRESS VA:") {
			v, err := parseHexAfter(line, "I_O_MARKER_ADDRESS VA:")
			if err != nil {
				return log, err
			}
			log.IOMarkerAddress = uint32(v)
			continue
		}
	}

	if err := sc.Err(); err != nil {
		return log, fmt.Errorf("apilog: scan failed: %w", err)
	}
	return log, nil
}

// parseCallLine 解析 "Call from: {hex} | API: {hex} | NAME: {name}"。
func parseCallLine(line, kind string) (APILogEntry, error) {
	var e APILogEntry
	e.Kind = kind
	parts := strings.Split(line, "|")
	if len(parts) != 3 {
		return e, fmt.Errorf("apilog: bad call line %q (want 3 parts, got %d)", line, len(parts))
	}

	caller, err := parseHexAfter(strings.TrimSpace(parts[0]), "Call from:")
	if err != nil {
		return e, fmt.Errorf("apilog: bad caller in %q: %w", line, err)
	}
	api, err := parseHexAfter(strings.TrimSpace(parts[1]), "API:")
	if err != nil {
		return e, fmt.Errorf("apilog: bad api in %q: %w", line, err)
	}
	name := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(parts[2]), "NAME:"))

	e.Caller = uint32(caller)
	e.API = uint32(api)
	e.Name = name
	return e, nil
}

// parseSetEventLine 解析 SETEVENT_ENTRY_ADDRESS 行。
//
//	Old WL: "Address: X | PUSH Y | JUMP Z"             (3 段)
//	New WL: "Address: X | PUSH Y | PUSH Z | JUMP W"    (4 段)
func parseSetEventLine(line string) (SetEventEntry, error) {
	var se SetEventEntry
	parts := strings.Split(line, "|")
	if len(parts) < 3 {
		return se, fmt.Errorf("apilog: bad SetEvent line %q (want >= 3 parts, got %d)", line, len(parts))
	}

	addr, err := parseHexAfter(strings.TrimSpace(parts[0]), "Address:")
	if err != nil {
		return se, fmt.Errorf("apilog: bad SetEvent address in %q: %w", line, err)
	}
	se.Address = uint32(addr)

	switch len(parts) {
	case 3: // Old WL: Address | PUSH Y | JUMP Z
		push, err := parseHexAfter(strings.TrimSpace(parts[1]), "PUSH")
		if err != nil {
			return se, fmt.Errorf("apilog: bad SetEvent PUSH in %q: %w", line, err)
		}
		jump, err := parseHexAfter(strings.TrimSpace(parts[2]), "JUMP")
		if err != nil {
			return se, fmt.Errorf("apilog: bad SetEvent JUMP in %q: %w", line, err)
		}
		se.Push = uint32(push)
		se.Jump = uint32(jump)
		se.IsNewWL = false

	case 4: // New WL: Address | PUSH Y | PUSH Z | JUMP W
		push, err := parseHexAfter(strings.TrimSpace(parts[1]), "PUSH")
		if err != nil {
			return se, fmt.Errorf("apilog: bad SetEvent PUSH[0] in %q: %w", line, err)
		}
		push2, err := parseHexAfter(strings.TrimSpace(parts[2]), "PUSH")
		if err != nil {
			return se, fmt.Errorf("apilog: bad SetEvent PUSH[1] in %q: %w", line, err)
		}
		jump, err := parseHexAfter(strings.TrimSpace(parts[3]), "JUMP")
		if err != nil {
			return se, fmt.Errorf("apilog: bad SetEvent JUMP in %q: %w", line, err)
		}
		se.Push = uint32(push)
		se.Push2 = uint32(push2)
		se.Jump = uint32(jump)
		se.IsNewWL = true

	default:
		return se, fmt.Errorf("apilog: unexpected part count %d in %q", len(parts), line)
	}

	return se, nil
}

// parseHexAfter 从 s 中去掉 prefix 前缀后按 16 进制解析剩余部分。
// 前后空白会被裁剪。返回的 uint64 可由调用方断言为更窄宽度。
func parseHexAfter(s, prefix string) (uint64, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, prefix)
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("apilog: empty hex after prefix %q", prefix)
	}
	return strconv.ParseUint(s, 16, 64)
}

// ---------------------------------------------------------------------------
// Variables and functions from modbase_windows.go
// ---------------------------------------------------------------------------

var (
	kernel32Dll2          = windows.NewLazySystemDLL("kernel32.dll")
	procCreateToolhelp32  = kernel32Dll2.NewProc("CreateToolhelp32Snapshot")
	procModule32First     = kernel32Dll2.NewProc("Module32FirstW")
	procModule32Next      = kernel32Dll2.NewProc("Module32NextW")
	procProcess32FirstW   = kernel32Dll2.NewProc("Process32FirstW")
	procProcess32NextW    = kernel32Dll2.NewProc("Process32NextW")
	procCloseHandle2      = kernel32Dll2.NewProc("CloseHandle")
	procIsWow64Process    = kernel32Dll2.NewProc("IsWow64Process")
	procReadProcessMemory = kernel32Dll2.NewProc("ReadProcessMemory")

	psapiDll                 = windows.NewLazySystemDLL("psapi.dll")
	procEnumProcessModulesEx = psapiDll.NewProc("EnumProcessModulesEx")
	procGetModuleInformation = psapiDll.NewProc("GetModuleInformation")
	procGetModuleBaseNameW   = psapiDll.NewProc("GetModuleBaseNameW")
)

// getModuleBaseForWow64Target 获取 32-bit DLL 的实际加载基址。
//
// 策略：
//  1. 在当前进程中查找一个已运行的 WOW64 进程（如 explorer.exe 的
//     32-bit 子进程），通过 Toolhelp32 获取其 32-bit DLL 基址。
//  2. 如果找不到运行中的 WOW64 进程，通过 EnumProcessModulesEx 在
//     目标进程获取（需要目标进程已初始化）。
//  3. 后备方案：通过 Toolhelp32 在目标进程获取。
//
// 所有 WOW64 进程共享相同的 32-bit DLL 基址（ASLR 只在启动时
// 随机化一次），所以从任意 WOW64 进程获取的基址都有效。
func getModuleBaseForWow64Target(pid uint32, processHandle uintptr, moduleName string) (uint64, error) {
	// 方案 1: 通过进程句柄直接获取（支持 suspended/paused 进程）
	if processHandle != 0 {
		if base, err := getModuleBaseByHandle(processHandle, moduleName); err == nil {
			return base, nil
		}
	}

	// 方案 2: 通过 Toolhelp32 在目标进程获取
	if base, err := getModuleBaseToolhelp(pid, moduleName); err == nil {
		return base, nil
	}

	// 方案 3: 通过 EnumProcessModulesEx + OpenProcess
	if base, err := getModuleBaseOpenProcess(pid, moduleName); err == nil {
		return base, nil
	}

	// 方案 4: 从任意运行中的 WOW64 进程获取 32-bit DLL 基址
	if base, err := getModuleBaseFromRunningWow64(moduleName); err == nil {
		return base, nil
	}

	return 0, fmt.Errorf("module %q not found for pid=%d (all methods failed)", moduleName, pid)
}

// touchPageViaReadProcessMemory reads 1 byte from the target process's virtual
// address using Win32 ReadProcessMemory.
//
// Why this is needed before EptHookForProcess:
//   - Windows DLLs are demand-paged. If the target process hasn't accessed a
//     DLL code page yet, the PTE is not-present.
//   - HyperDbg's ReadMemory IOCTL reads via MmGetPhysicalAddress +
//     MmMapIoSpaceEx, which does NOT trigger page faults — it reads whatever
//     is currently in physical RAM (returns 0 / fails for not-present pages).
//   - HyperDbg's .pagein command (DebuggerCommandBringPagein) is a stub in
//     local debugging mode — it logs "Page-request is received!" and returns
//     success without injecting any #PF.
//   - EptHookForProcess validation calls VirtualAddressToPhysicalAddressByProcessId
//     → MmGetPhysicalAddress, which returns 0 for not-present pages, causing
//     DEBUGGER_ERROR_INVALID_ADDRESS (0xC0000005).
//
// ReadProcessMemory → NtReadVirtualMemory accesses the virtual address through
// the target's page tables, triggering a real page fault that the OS services
// by reading the page from disk into RAM and marking the PTE present. After
// this call, MmGetPhysicalAddress returns the correct physical address.
func touchPageViaReadProcessMemory(processHandle uintptr, addr uint64) error {
	if processHandle == 0 {
		return fmt.Errorf("touchPage: process handle is 0")
	}
	var buf [1]byte
	var bytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		processHandle,
		uintptr(addr),
		uintptr(unsafe.Pointer(&buf[0])),
		1,
		uintptr(unsafe.Pointer(&bytesRead)),
	)
	if ret == 0 {
		return fmt.Errorf("ReadProcessMemory(0x%X) failed: %w", addr, err)
	}
	return nil
}

// getModuleBaseByHandle 使用进程句柄通过 EnumProcessModulesEx 获取模块基址。
//
// 对于 WOW64 目标进程，从 64-bit 调试器进程调用 EnumProcessModulesEx 时，
// LIST_MODULES_ALL (0x03) 会同时返回 32-bit 和 64-bit 模块，但
// GetModuleInformation 对 32-bit 模块可能返回失败（ret=0），导致只能拿到
// 64-bit ntdll 的基址（例如 0x7FFC04660000），而 32-bit ntdll（例如
// 0x77A00000）被跳过。
//
// 解决方案：先尝试 LIST_MODULES_32BIT (0x01)，只枚举 32-bit 模块。这样
// GetModuleInformation 对所有返回的模块都能成功，ntdll.dll 的查询会直接
// 命中 32-bit ntdll。如果 LIST_MODULES_32BIT 失败（例如目标不是 WOW64），
// 回退到 LIST_MODULES_ALL 并优先返回 < 0x100000000 的基址。
func getModuleBaseByHandle(processHandle uintptr, moduleName string) (uint64, error) {
	if processHandle == 0 {
		return 0, fmt.Errorf("handle is 0")
	}

	// 尝试 1: LIST_MODULES_32BIT — 只枚举 32-bit 模块。
	if base, err := getModuleBaseByHandleWithFilter(processHandle, moduleName, 0x01); err == nil {
		return base, nil
	}

	// 尝试 2: LIST_MODULES_ALL — 枚举所有模块，优先返回 32-bit 基址。
	return getModuleBaseByHandleWithFilter(processHandle, moduleName, 0x03)
}

// getModuleBaseByHandleWithFilter 是 getModuleBaseByHandle 的内部实现，
// filterFlag 是传给 EnumProcessModulesEx 的 dwFilterFlag。
// 当 filterFlag == LIST_MODULES_ALL 时，可能同时返回 32-bit 和 64-bit 模块；
// 此时代码优先返回 < 0x100000000 的基址，只有找不到 32-bit 匹配时才回退
// 到 64-bit 基址。
func getModuleBaseByHandleWithFilter(processHandle uintptr, moduleName string, filterFlag uint32) (uint64, error) {
	var hMods [1024]windows.Handle
	cbNeeded := uint32(0)

	ret, _, err := procEnumProcessModulesEx.Call(
		processHandle,
		uintptr(unsafe.Pointer(&hMods[0])),
		uintptr(len(hMods))*unsafe.Sizeof(hMods[0]),
		uintptr(unsafe.Pointer(&cbNeeded)),
		uintptr(filterFlag),
	)
	if ret == 0 {
		return 0, fmt.Errorf("EnumProcessModulesEx(filter=0x%X): %w", filterFlag, err)
	}

	numMods := cbNeeded / uint32(unsafe.Sizeof(hMods[0]))
	if numMods > uint32(len(hMods)) {
		numMods = uint32(len(hMods))
	}

	// 64-bit fallback: only used if no 32-bit match is found.
	var fallback64 uint64
	foundFallback := false

	for i := uint32(0); i < numMods; i++ {
		var baseName [maxPath]uint16
		ret, _, _ = procGetModuleBaseNameW.Call(
			processHandle,
			uintptr(hMods[i]),
			uintptr(unsafe.Pointer(&baseName[0])),
			maxPath,
		)
		if ret != 0 {
			name := windows.UTF16ToString(baseName[:])
			if equalFold(name, moduleName) {
				var modInfo struct {
					BaseAddr uintptr
					Size     uint32
					Entry    uintptr
				}
				ret, _, _ = procGetModuleInformation.Call(
					processHandle,
					uintptr(hMods[i]),
					uintptr(unsafe.Pointer(&modInfo)),
					uintptr(unsafe.Sizeof(modInfo)),
				)
				if ret != 0 {
					base := uint64(modInfo.BaseAddr)
					// Prefer 32-bit base (< 0x100000000) for WOW64 targets.
					// ntdll.dll in a WOW64 process exists as both 32-bit
					// and 64-bit; we want the 32-bit one for hooking 32-bit
					// exports like RtlAllocateHeap.
					if base < 0x100000000 {
						return base, nil
					}
					if !foundFallback {
						fallback64 = base
						foundFallback = true
					}
				}
			}
		}
	}

	if foundFallback {
		return fallback64, nil
	}

	return 0, fmt.Errorf("module %q not found via handle (filter=0x%X)", moduleName, filterFlag)
}

func getModuleBaseToolhelp(pid uint32, moduleName string) (uint64, error) {
	snap, _, err := procCreateToolhelp32.Call(
		uintptr(th32csSnapModule|th32csSnapModule32),
		uintptr(pid),
	)
	if snap == ^uintptr(0) {
		return 0, fmt.Errorf("CreateToolhelp32Snapshot(pid=%d): %w", pid, err)
	}
	defer procCloseHandle2.Call(snap)

	var me moduleEntry32
	me.Size = uint32(unsafe.Sizeof(me))

	ret, _, err := procModule32First.Call(snap, uintptr(unsafe.Pointer(&me)))
	if ret == 0 {
		return 0, fmt.Errorf("Module32First: %w", err)
	}

	for {
		modName := windows.UTF16ToString(me.Module[:])
		if equalFold(modName, moduleName) {
			return uint64(me.ModBaseAddr), nil
		}

		me.Size = uint32(unsafe.Sizeof(me))
		ret, _, err = procModule32Next.Call(snap, uintptr(unsafe.Pointer(&me)))
		if ret == 0 {
			break
		}
	}

	return 0, fmt.Errorf("module %q not found", moduleName)
}

func getModuleBaseOpenProcess(pid uint32, moduleName string) (uint64, error) {
	hProcess, err := windows.OpenProcess(
		windows.PROCESS_QUERY_INFORMATION|windows.PROCESS_VM_READ|windows.SYNCHRONIZE,
		false, pid,
	)
	if err != nil {
		return 0, fmt.Errorf("OpenProcess(pid=%d): %w", pid, err)
	}
	defer windows.CloseHandle(hProcess)

	var hMods [1024]windows.Handle
	cbNeeded := uint32(0)

	ret, _, err := procEnumProcessModulesEx.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&hMods[0])),
		uintptr(len(hMods))*unsafe.Sizeof(hMods[0]),
		uintptr(unsafe.Pointer(&cbNeeded)),
		0x03,
	)
	if ret == 0 {
		return 0, fmt.Errorf("EnumProcessModulesEx: %w", err)
	}

	numMods := cbNeeded / uint32(unsafe.Sizeof(hMods[0]))
	if numMods > uint32(len(hMods)) {
		numMods = uint32(len(hMods))
	}

	for i := uint32(0); i < numMods; i++ {
		var baseName [maxPath]uint16
		ret, _, _ = procGetModuleBaseNameW.Call(
			uintptr(hProcess),
			uintptr(hMods[i]),
			uintptr(unsafe.Pointer(&baseName[0])),
			maxPath,
		)
		if ret != 0 {
			name := windows.UTF16ToString(baseName[:])
			if equalFold(name, moduleName) {
				var modInfo struct {
					BaseAddr uintptr
					Size     uint32
					Entry    uintptr
				}
				ret, _, _ = procGetModuleInformation.Call(
					uintptr(hProcess),
					uintptr(hMods[i]),
					uintptr(unsafe.Pointer(&modInfo)),
					uintptr(unsafe.Sizeof(modInfo)),
				)
				if ret != 0 {
					return uint64(modInfo.BaseAddr), nil
				}
			}
		}
	}

	return 0, fmt.Errorf("module %q not found via EnumProcessModulesEx", moduleName)
}

// getModuleBaseFromRunningWow64 通过查找系统中运行中的 WOW64 进程
// 来获取 32-bit DLL 基址。所有 WOW64 进程共享相同的 32-bit DLL
// 基址，所以从任意运行中的 WOW64 进程获取即可。
//
// 关键：只查找 WOW64（32-bit）进程，跳过 64-bit 进程。
// 因为 Toolhelp32 在 64-bit 进程上只返回 64-bit 模块基址。
func getModuleBaseFromRunningWow64(moduleName string) (uint64, error) {
	// 遍历系统进程查找 WOW64 进程
	snap, _, err := procCreateToolhelp32.Call(
		uintptr(0x00000002), // TH32CS_SNAPPROCESS
		0,
	)
	if snap == ^uintptr(0) {
		return 0, fmt.Errorf("CreateToolhelp32Snapshot(PROCESS): %w", err)
	}
	defer procCloseHandle2.Call(snap)

	type processEntry32 struct {
		Size            uint32
		CntUsage        uint32
		ProcessID       uint32
		DefaultHeapID   uintptr
		ModuleID        uint32
		CntThreads      uint32
		ParentProcessID uint32
		PriClass        int32
		Flags           uint32
		ExeFile         [maxPath]uint16
	}

	var pe processEntry32
	pe.Size = uint32(unsafe.Sizeof(pe))

	ret, _, _ := procProcess32FirstW.Call(snap, uintptr(unsafe.Pointer(&pe)))
	if ret == 0 {
		return 0, fmt.Errorf("Process32First failed")
	}

	for {
		// 跳过 System 和 Idle 进程
		if pe.ProcessID > 4 {
			// 检查是否是 WOW64 进程
			if isWow64Process(pe.ProcessID) {
				if base, err := getModuleBaseToolhelp(pe.ProcessID, moduleName); err == nil {
					return base, nil
				}
			}
		}

		pe.Size = uint32(unsafe.Sizeof(pe))
		ret, _, _ = procProcess32NextW.Call(snap, uintptr(unsafe.Pointer(&pe)))
		if ret == 0 {
			break
		}
	}

	return 0, fmt.Errorf("no running WOW64 process found with module %q", moduleName)
}

// isWow64Process 检查指定 PID 的进程是否是 WOW64（32-bit）进程。
func isWow64Process(pid uint32) bool {
	hProcess, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return false
	}
	defer windows.CloseHandle(hProcess)

	var wow64 int32
	ret, _, _ := procIsWow64Process.Call(uintptr(hProcess), uintptr(unsafe.Pointer(&wow64)))
	return ret != 0 && wow64 != 0
}

// equalFold 是简单的 ASCII 不区分大小写比较。
func equalFold(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca += 'a' - 'A'
		}
		if cb >= 'A' && cb <= 'Z' {
			cb += 'a' - 'A'
		}
		if ca != cb {
			return false
		}
	}
	return true
}

// ---------------------------------------------------------------------------
// New functions for unpacker14
// ---------------------------------------------------------------------------

// ParseOverviewOEP parses OVERVIEW format lines and returns the VM OEP fields.
//
// Expected line format:
//
//	VM ADDR: 7A978B
//	VM ALIGN: D695C014
//	VM PUSH: 29E4826D
//	VM JUMP: 6C5F95
//
// Each line is "VM KEY: HEXVALUE". The hex value is parsed as uint32.
// Returns error if any required key is missing or has an invalid hex value.
func ParseOverviewOEP(data []byte) (VMOEPResult, error) {
	var r VMOEPResult
	found := make(map[string]bool)

	sc := bufio.NewScanner(bytes.NewReader(data))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		if !strings.HasPrefix(line, "VM ") {
			continue
		}

		// Split "VM KEY: HEXVALUE"
		colonIdx := strings.Index(line, ":")
		if colonIdx < 0 {
			continue
		}
		key := strings.TrimSpace(line[3:colonIdx]) // strip "VM " prefix, take up to ":"
		// Only parse the four known keys; skip other "VM X:" lines
		switch key {
		case "ADDR", "ALIGN", "PUSH", "JUMP":
		default:
			continue
		}
		val := strings.TrimSpace(line[colonIdx+1:]) // after ":"
		if val == "" {
			return r, fmt.Errorf("overview: empty hex value for key %q", key)
		}
		v, err := strconv.ParseUint(val, 16, 32)
		if err != nil {
			return r, fmt.Errorf("overview: bad hex %q for key %q: %w", val, key, err)
		}

		switch key {
		case "ADDR":
			r.Addr = uint32(v)
		case "ALIGN":
			r.Align = uint32(v)
		case "PUSH":
			r.Push = uint32(v)
		case "JUMP":
			r.Jump = uint32(v)
		default:
			continue
		}
		found[key] = true
	}
	if err := sc.Err(); err != nil {
		return r, fmt.Errorf("overview: scan failed: %w", err)
	}

	for _, k := range []string{"ADDR", "ALIGN", "PUSH", "JUMP"} {
		if !found[k] {
			return r, fmt.Errorf("overview: missing key %q", k)
		}
	}
	return r, nil
}

// ComputeStackAntiDump replicates Ultra Unpacker 1.4 script L1375-1390:
//
//	oldCalc = sad ^ 0x8647A6B4
//	newCalc = (sad - 8) ^ 0x7647A6B4
func ComputeStackAntiDump(sad uint32) (oldCalc, newCalc uint32) {
	oldCalc = sad ^ 0x8647A6B4
	newCalc = (sad - 8) ^ 0x7647A6B4
	return oldCalc, newCalc
}

// DetectVMType searches WLSEC bytes for VM signatures.
//
// Old WL CISC: pattern 68 ?? ?? ?? ?? E9 ?? ?? ?? ?? repeated 3 times with 0xFF between.
// New WL CISC: pattern 68 ?? ?? ?? ?? 68 ?? ?? ?? ?? E9 ?? ?? ?? ?? repeated 2 times with 0xFF between.
//
// Returns:
//
//	0 = Old WL CISC
//	1 = New WL CISC
//	3 = RISC (fallback when neither CISC pattern matches)
func DetectVMType(wlsec []byte) int {
	// Old WL CISC: 10-byte pattern repeated 3 times with 0xFF separator.
	// Total: 10 + 1 + 10 + 1 + 10 = 32 bytes
	if matchOldWLCISC(wlsec) {
		return 0
	}
	// New WL CISC: 15-byte pattern repeated 2 times with 0xFF separator.
	// Total: 15 + 1 + 15 = 31 bytes
	if matchNewWLCISC(wlsec) {
		return 1
	}
	return 3
}

// matchOldWLCISC searches for the Old WL CISC signature:
// (68 XX XX XX XX E9 YY YY YY YY) FF (68 XX XX XX XX E9 YY YY YY YY) FF (68 XX XX XX XX E9 YY YY YY YY)
func matchOldWLCISC(data []byte) bool {
	const (
		patLen = 10 // 68 XX XX XX XX E9 YY YY YY YY
		sep    = 0xFF
		total  = patLen + 1 + patLen + 1 + patLen // 32
	)
	if len(data) < total {
		return false
	}

	for i := 0; i <= len(data)-total; i++ {
		off := i
		ok := true
		for rep := 0; rep < 3 && ok; rep++ {
			if data[off] != opPushImm32 || data[off+5] != opJmpRel32 {
				ok = false
				break
			}
			off += patLen
			// After the first two repetitions, expect 0xFF separator
			if rep < 2 {
				if data[off] != sep {
					ok = false
					break
				}
				off++
			}
		}
		if ok {
			return true
		}
	}
	return false
}

// matchNewWLCISC searches for the New WL CISC signature:
// (68 XX XX XX XX 68 XX XX XX XX E9 YY YY YY YY) FF (68 XX XX XX XX 68 XX XX XX XX E9 YY YY YY YY)
func matchNewWLCISC(data []byte) bool {
	const (
		patLen = 15 // 68 XX XX XX XX 68 XX XX XX XX E9 YY YY YY YY
		sep    = 0xFF
		total  = patLen + 1 + patLen // 31
	)
	if len(data) < total {
		return false
	}

	for i := 0; i <= len(data)-total; i++ {
		off := i
		ok := true
		for rep := 0; rep < 2 && ok; rep++ {
			if data[off] != opPushImm32 || data[off+5] != opPushImm32 || data[off+10] != opJmpRel32 {
				ok = false
				break
			}
			off += patLen
			// After the first repetition, expect 0xFF separator
			if rep < 1 {
				if data[off] != sep {
					ok = false
					break
				}
				off++
			}
		}
		if ok {
			return true
		}
	}
	return false
}

// FormatVMOEP returns the OVERVIEW format string for VM OEP values:
//
//	VM ADDR: {vmAddr:X}
//	VM ALIGN: {wlAlign:X}
//	VM PUSH: {vmPush:X}
//	VM JUMP: {vmJump:X}
func FormatVMOEP(vmAddr, wlAlign, vmPush, vmJump uint32) string {
	return fmt.Sprintf("VM ADDR: %X\nVM ALIGN: %X\nVM PUSH: %X\nVM JUMP: %X",
		vmAddr, wlAlign, vmPush, vmJump)
}
