// Package themida — unpacker14_test.go
//
// 验证 unpacker14.go 复刻 Themida - WinLicense Ultra Unpacker 1.4 的逻辑
// 与真实数据一致。
//
// 预期数据来源:
//   - 预期API Logger of - SuperRecovery V4.8.1.5.txt (L255-267)
//   - OVERVIEW - SuperRecovery V4.8.1.5.txt (L19-27)
package themida

import (
	"context"
	"encoding/binary"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	superRecoveryLogPath      = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\预期API Logger of - SuperRecovery V4.8.1.5.txt`
	superRecoveryOverviewPath = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\OVERVIEW - SuperRecovery V4.8.1.5.txt`

	// 预期 SETEVENT_ENTRY_ADDRESS (日志 L255-267)
	expectedSetEventAddr    uint32 = 0x794443
	expectedSetEventPush    uint32 = 0x29E3997A
	expectedSetEventJump    uint32 = 0x6C5F95
	expectedIOMarkerAddress uint32 = 0x6C5CD5
	expectedSetEventAPIAddr uint32 = 0x766DE910

	// 预期 VM OEP (OVERVIEW L19-27)
	expectedVMAddr  uint32 = 0x7A978B
	expectedVMAlign uint32 = 0xD695C014
	expectedVMPush  uint32 = 0x29E4826D
	expectedVMJump  uint32 = 0x6C5F95
)

// --- SetEvent 解码单元测试 (合成字节序列) ---

func TestParseSetEventEntry_OldWL(t *testing.T) {
	t.Parallel()
	var addr uint32 = expectedSetEventAddr
	var wantPush uint32 = expectedSetEventPush
	var wantJump uint32 = expectedSetEventJump
	disp := int32(wantJump - addr - uint32(seteventOldWLSize))

	code := make([]byte, seteventOldWLSize)
	code[0] = opPushImm32
	binary.LittleEndian.PutUint32(code[1:5], wantPush)
	code[5] = opJmpRel32
	binary.LittleEndian.PutUint32(code[6:10], uint32(disp))

	e, ok := ParseSetEventEntry(addr, code)
	if !ok {
		t.Fatal("ParseSetEventEntry returned false")
	}
	if e.IsNewWL {
		t.Error("IsNewWL = true, want false")
	}
	if e.Address != addr {
		t.Errorf("Address = 0x%x, want 0x%x", e.Address, addr)
	}
	if e.Push != wantPush {
		t.Errorf("Push = 0x%08X, want 0x%08X", e.Push, wantPush)
	}
	if e.Jump != wantJump {
		t.Errorf("Jump = 0x%x, want 0x%x", e.Jump, wantJump)
	}
	wantFormat := "Address: 794443 | PUSH 29E3997A | JUMP 6C5F95"
	if got := e.Format(); got != wantFormat {
		t.Errorf("Format() = %q, want %q", got, wantFormat)
	}
}

func TestParseSetEventEntry_NewWL(t *testing.T) {
	t.Parallel()
	const addr = uint32(0x00500000)
	const wantPush1 = uint32(0xAAAAAAAA)
	const wantPush2 = uint32(0xBBBBBBBB)
	const wantJump = uint32(0x00600000)
	disp := int32(wantJump - addr - uint32(seteventNewWLSize))

	code := make([]byte, seteventNewWLSize)
	code[0] = opPushImm32
	binary.LittleEndian.PutUint32(code[1:5], wantPush1)
	code[5] = opPushImm32
	binary.LittleEndian.PutUint32(code[6:10], wantPush2)
	code[10] = opJmpRel32
	binary.LittleEndian.PutUint32(code[11:15], uint32(disp))

	e, ok := ParseSetEventEntry(addr, code)
	if !ok {
		t.Fatal("ParseSetEventEntry returned false")
	}
	if !e.IsNewWL {
		t.Error("IsNewWL = false, want true")
	}
	if e.Push != wantPush1 {
		t.Errorf("Push = 0x%08X, want 0x%08X", e.Push, wantPush1)
	}
	if e.Push2 != wantPush2 {
		t.Errorf("Push2 = 0x%08X, want 0x%08X", e.Push2, wantPush2)
	}
	if e.Jump != wantJump {
		t.Errorf("Jump = 0x%x, want 0x%x", e.Jump, wantJump)
	}
}

func TestParseSetEventEntry_RejectsBadPattern(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name string
		code []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"too_short", []byte{0x68, 0x11, 0x22, 0x33, 0x44}},
		{"no_push_opcode", []byte{0x90, 0x90, 0x90, 0x90, 0x90, 0xE9, 0x11, 0x22, 0x33, 0x44}},
		{"push_but_no_jmp", []byte{0x68, 0x11, 0x22, 0x33, 0x44, 0x90, 0x11, 0x22, 0x33, 0x44}},
		{"new_wl_missing_third_opcode", []byte{
			0x68, 0x11, 0x22, 0x33, 0x44,
			0x68, 0x55, 0x66, 0x77, 0x88,
			0x90, 0x99, 0xAA, 0xBB, 0xCC,
		}},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if _, ok := ParseSetEventEntry(0x400000, tc.code); ok {
				t.Errorf("ParseSetEventEntry returned true for bad pattern %q", tc.name)
			}
		})
	}
}

// --- API 日志解析单元测试 (合成日志文本) ---

func TestParseAPILog_Minimal(t *testing.T) {
	t.Parallel()
	const input = `---------------EX--------------------------------------
Call from: 794443 | API: 766DE910 | NAME: SetEvent
-------------------------------------------------------
*******************************************************
--------------- SETEVENT_ENTRY_ADDRESS ----------------
-------------------------------------------------------
Address: 794443 | PUSH 29E3997A | JUMP 6C5F95
-------------------------------------------------------
-------------------------------------------------------
--------------- I_O_MARKER_ADDRESS --------------------
-------------------------------------------------------
I_O_MARKER_ADDRESS VA: 6C5CD5
-------------------------------------------------------
`
	log, err := ParseAPILog([]byte(input))
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	if len(log.Entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(log.Entries))
	}
	e := log.Entries[0]
	if e.Kind != "EX" {
		t.Errorf("entry.Kind = %q, want EX", e.Kind)
	}
	if e.Caller != 0x794443 {
		t.Errorf("entry.Caller = 0x%x, want 0x794443", e.Caller)
	}
	if e.API != 0x766DE910 {
		t.Errorf("entry.API = 0x%x, want 0x766DE910", e.API)
	}
	if e.Name != "SetEvent" {
		t.Errorf("entry.Name = %q, want SetEvent", e.Name)
	}
	if log.SetEventEntry == nil {
		t.Fatal("SetEventEntry is nil")
	}
	se := *log.SetEventEntry
	if se.Address != 0x794443 || se.Push != 0x29E3997A || se.Jump != 0x6C5F95 {
		t.Errorf("SetEvent = %+v, want {0x794443 0x29E3997A 0 0x6C5F95 false}", se)
	}
	if log.IOMarkerAddress != 0x6C5CD5 {
		t.Errorf("IOMarkerAddress = 0x%x, want 0x6C5CD5", log.IOMarkerAddress)
	}
}

func TestParseAPILog_NewWLFormat(t *testing.T) {
	t.Parallel()
	const input = `Address: 5474C3 | PUSH D28AEFB | PUSH 12345678 | JUMP 478CB2
`
	log, err := ParseAPILog([]byte(input))
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	if log.SetEventEntry == nil {
		t.Fatal("SetEventEntry is nil")
	}
	se := *log.SetEventEntry
	if !se.IsNewWL {
		t.Error("IsNewWL = false, want true")
	}
	if se.Push != 0xD28AEFB {
		t.Errorf("Push = 0x%08X, want 0xD28AEFB", se.Push)
	}
	if se.Push2 != 0x12345678 {
		t.Errorf("Push2 = 0x%08X, want 0x12345678", se.Push2)
	}
}

// --- ParseOverviewOEP 单元测试 (合成 OVERVIEW 文本) ---

func TestParseOverviewOEP_Synthetic(t *testing.T) {
	t.Parallel()
	input := `************************************************************
VM OEP Address found! - Is in use!

VM ADDR: 7A978B 
VM ALIGN: D695C014 
VM PUSH: 29E4826D 
VM JUMP: 6C5F95
********************
`
	oep, err := ParseOverviewOEP([]byte(input))
	if err != nil {
		t.Fatalf("ParseOverviewOEP failed: %v", err)
	}
	if oep.Addr != 0x7A978B {
		t.Errorf("Addr = 0x%x, want 0x7A978B", oep.Addr)
	}
	if oep.Align != 0xD695C014 {
		t.Errorf("Align = 0x%x, want 0xD695C014", oep.Align)
	}
	if oep.Push != 0x29E4826D {
		t.Errorf("Push = 0x%x, want 0x29E4826D", oep.Push)
	}
	if oep.Jump != 0x6C5F95 {
		t.Errorf("Jump = 0x%x, want 0x6C5F95", oep.Jump)
	}
}

// --- StackAntiDump 算法测试 ---

func TestComputeStackAntiDump(t *testing.T) {
	t.Parallel()
	// 用脚本 L1375-1390 的公式: old = sad ^ 0x8647A6B4, new = (sad-8) ^ 0x7647A6B4
	sad := uint32(0xDEADBEEF)
	oldCalc := sad ^ 0x8647A6B4
	newCalc := (sad - 8) ^ 0x7647A6B4

	gotOld, gotNew := ComputeStackAntiDump(sad)
	if gotOld != oldCalc {
		t.Errorf("oldCalc = 0x%x, want 0x%x", gotOld, oldCalc)
	}
	if gotNew != newCalc {
		t.Errorf("newCalc = 0x%x, want 0x%x", gotNew, newCalc)
	}
}

// --- DetectVMType 测试 (合成 WLSEC 字节) ---

func TestDetectVMType_OldWL(t *testing.T) {
	t.Parallel()
	// 构造 Old WL CISC 签名: 68 XX XX XX XX E9 YY YY YY YY FF (重复 3 次)
	sig := make([]byte, 0)
	for i := 0; i < 3; i++ {
		sig = append(sig, 0x68, 0x11, 0x22, 0x33, 0x44) // PUSH
		sig = append(sig, 0xE9, 0x55, 0x66, 0x77, 0x88) // JMP
		sig = append(sig, 0xFF)                         // separator
	}
	if got := DetectVMType(sig); got != 0 {
		t.Errorf("DetectVMType = %d, want 0 (Old WL CISC)", got)
	}
}

func TestDetectVMType_NewWL(t *testing.T) {
	t.Parallel()
	// 构造 New WL CISC 签名: 68 XX 68 XX E9 YY FF (重复 2 次)
	sig := make([]byte, 0)
	for i := 0; i < 2; i++ {
		sig = append(sig, 0x68, 0x11, 0x22, 0x33, 0x44) // PUSH1
		sig = append(sig, 0x68, 0x55, 0x66, 0x77, 0x88) // PUSH2
		sig = append(sig, 0xE9, 0xAA, 0xBB, 0xCC, 0xDD) // JMP
		sig = append(sig, 0xFF)                         // separator
	}
	if got := DetectVMType(sig); got != 1 {
		t.Errorf("DetectVMType = %d, want 1 (New WL CISC)", got)
	}
}

func TestDetectVMType_RISC(t *testing.T) {
	t.Parallel()
	// 无匹配签名 → RISC
	if got := DetectVMType([]byte{0x90, 0x90, 0x90}); got != 3 {
		t.Errorf("DetectVMType = %d, want 3 (RISC)", got)
	}
}

// --- FormatVMOEP 测试 ---

func TestFormatVMOEP(t *testing.T) {
	t.Parallel()
	got := FormatVMOEP(expectedVMAddr, expectedVMAlign, expectedVMPush, expectedVMJump)
	want := "VM ADDR: 7A978B\nVM ALIGN: D695C014\nVM PUSH: 29E4826D\nVM JUMP: 6C5F95"
	if got != want {
		t.Errorf("FormatVMOEP = %q, want %q", got, want)
	}
}

// --- gmemi / IsThemidaCaller / IsInOepRange 测试 ---

var testMap = MemoryMap{
	{BaseAddress: 0x10000, RegionSize: 0x1000},
	{BaseAddress: 0x40000, RegionSize: 0x20000},
	{BaseAddress: 0x70000000, RegionSize: 0x100000},
}

func TestGmemiMemoryBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		addr uint64
		want uint64
		ok   bool
	}{
		{0x10100, 0x10000, true},
		{0x41234, 0x40000, true},
		{0x50000, 0x40000, true},
		{0x30000, 0, false}, // hole
		{0x70001000, 0x70000000, true},
	}
	for _, tc := range tests {
		base, ok := GmemiMemoryBase(testMap, tc.addr)
		if ok != tc.ok || base != tc.want {
			t.Errorf("GmemiMemoryBase(0x%x) = (0x%x, %v), want (0x%x, %v)", tc.addr, base, ok, tc.want, tc.ok)
		}
	}
}

func TestGmemiMemorySize(t *testing.T) {
	t.Parallel()
	tests := []struct {
		base uint64
		want uint64
		ok   bool
	}{
		{0x10000, 0x1000, true},
		{0x40000, 0x20000, true},
		{0x70000000, 0x100000, true},
		{0x99999, 0, false},
	}
	for _, tc := range tests {
		size, ok := GmemiMemorySize(testMap, tc.base)
		if ok != tc.ok || size != tc.want {
			t.Errorf("GmemiMemorySize(0x%x) = (0x%x, %v), want (0x%x, %v)", tc.base, size, ok, tc.want, tc.ok)
		}
	}
}

func TestIsThemidaCaller(t *testing.T) {
	t.Parallel()
	if IsThemidaCaller(0x794443) != true {
		t.Error("0x794443 should be Themida caller (< 0x70000000)")
	}
	if IsThemidaCaller(0x766DE910) != false {
		t.Error("0x766DE910 should NOT be Themida caller (>= 0x70000000)")
	}
}

func TestIsInOepRange(t *testing.T) {
	t.Parallel()
	if !IsInOepRange(0x400FFE, 0x400000, 0x1000) {
		t.Error("0x400FFE should be in range [0x400000, 0x401000)")
	}
	if IsInOepRange(0x401000, 0x400000, 0x1000) {
		t.Error("0x401000 should NOT be in range [0x400000, 0x401000) (exclusive upper bound)")
	}
	if IsInOepRange(0x401000, 0x400000, 0) {
		t.Error("zero size should always return false")
	}
}

func TestIsRtlAllocateHeapMagic(t *testing.T) {
	t.Parallel()
	magic := []byte("RtlAllocateHeap")
	if !IsRtlAllocateHeapMagic(magic) {
		t.Error("'RtlAllocateHeap' should match magic")
	}
	if IsRtlAllocateHeapMagic([]byte("SomethingElse")) {
		t.Error("'SomethingElse' should NOT match magic")
	}
}

func TestStackAntiDumpCalc(t *testing.T) {
	t.Parallel()
	sa1, sa2 := StackAntiDumpCalc(0x12345678, 0xABCDEF01)
	if sa1 != 0x12345678^0x8647A6B4 {
		t.Errorf("sa1 = 0x%x, want 0x%x", sa1, 0x12345678^0x8647A6B4)
	}
	if sa2 != 0xABCDEF01^0x7647A6B4 {
		t.Errorf("sa2 = 0x%x, want 0x%x", sa2, 0xABCDEF01^0x7647A6B4)
	}
}

// --- 真实数据集成测试 (SuperRecovery V4.8.1.5) ---

func TestRealSuperRecovery_APILog(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile(superRecoveryLogPath)
	if err != nil {
		t.Skipf("API log not found: %v", err)
	}

	log, err := ParseAPILog(data)
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	if len(log.Entries) == 0 {
		t.Fatal("no entries parsed from API log")
	}
	t.Logf("parsed %d API call entries", len(log.Entries))

	// 验证 SETEVENT_ENTRY_ADDRESS
	if log.SetEventEntry == nil {
		t.Fatal("SetEventEntry is nil")
	}
	se := *log.SetEventEntry
	if se.Address != expectedSetEventAddr {
		t.Errorf("SetEvent Address = 0x%x, want 0x%x", se.Address, expectedSetEventAddr)
	}
	if se.Push != expectedSetEventPush {
		t.Errorf("SetEvent Push = 0x%08X, want 0x%08X", se.Push, expectedSetEventPush)
	}
	if se.Jump != expectedSetEventJump {
		t.Errorf("SetEvent Jump = 0x%x, want 0x%x", se.Jump, expectedSetEventJump)
	}
	if se.IsNewWL {
		t.Error("SetEvent IsNewWL = true, want false (Old WL CISC)")
	}
	t.Logf("SETEVENT_ENTRY_ADDRESS: %s", se.Format())

	// 验证 I_O_MARKER_ADDRESS
	if log.IOMarkerAddress != expectedIOMarkerAddress {
		t.Errorf("IOMarkerAddress = 0x%x, want 0x%x", log.IOMarkerAddress, expectedIOMarkerAddress)
	}
	t.Logf("I_O_MARKER_ADDRESS VA: %X", log.IOMarkerAddress)

	// 算法一致性: JUMP == ADDR + 10 + signed(int32) disp
	disp := se.JumpDisplacement()
	computedJump := se.Address + uint32(seteventOldWLSize) + uint32(disp)
	if computedJump != se.Jump {
		t.Errorf("algorithm: ADDR + 10 + disp = 0x%x, want JUMP = 0x%x", computedJump, se.Jump)
	}
	t.Logf("algorithm: ADDR + 10 + disp(0x%08X) = 0x%x == JUMP", disp, computedJump)

	// Format 与日志原文一致
	wantFormat := "Address: 794443 | PUSH 29E3997A | JUMP 6C5F95"
	if got := se.Format(); got != wantFormat {
		t.Errorf("Format() = %q, want %q", got, wantFormat)
	}

	// SetEvent 调用方在日志流水中能找到
	found := false
	for _, e := range log.Entries {
		if e.Caller == se.Address && e.Name == "SetEvent" {
			found = true
			if e.API != expectedSetEventAPIAddr {
				t.Errorf("SetEvent API = 0x%08X, want 0x%08X", e.API, expectedSetEventAPIAddr)
			}
			break
		}
	}
	if !found {
		t.Errorf("SetEvent caller 0x%x not found in entries", se.Address)
	}

	// I_O_MARKER 与 JUMP 距离合理 (< 4KB)
	if log.IOMarkerAddress > se.Jump {
		t.Errorf("I_O_MARKER 0x%x > JUMP 0x%x", log.IOMarkerAddress, se.Jump)
	}
	distance := se.Jump - log.IOMarkerAddress
	if distance > 0x1000 {
		t.Errorf("I_O_MARKER -> JUMP distance = 0x%x, want <= 0x1000", distance)
	}
	t.Logf("I_O_MARKER -> JUMP distance = 0x%x (%d bytes)", distance, distance)
}

func TestRealSuperRecovery_CallerFilter(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile(superRecoveryLogPath)
	if err != nil {
		t.Skipf("API log not found: %v", err)
	}
	log, err := ParseAPILog(data)
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	for _, e := range log.Entries {
		if !IsThemidaCaller(uint64(e.Caller)) {
			t.Errorf("caller 0x%x (NAME=%s) should pass IsThemidaCaller", e.Caller, e.Name)
		}
		if IsThemidaCaller(uint64(e.API)) {
			t.Errorf("API 0x%x (NAME=%s) should be filtered by IsThemidaCaller", e.API, e.Name)
		}
	}
}

func TestRealSuperRecovery_SetEventEntryReconstruction(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile(superRecoveryLogPath)
	if err != nil {
		t.Skipf("API log not found: %v", err)
	}
	log, err := ParseAPILog(data)
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	if log.SetEventEntry == nil {
		t.Fatal("SetEventEntry is nil")
	}
	expected := *log.SetEventEntry

	disp := expected.JumpDisplacement()
	code := make([]byte, seteventOldWLSize)
	code[0] = opPushImm32
	binary.LittleEndian.PutUint32(code[1:5], expected.Push)
	code[5] = opJmpRel32
	binary.LittleEndian.PutUint32(code[6:10], uint32(disp))

	got, ok := ParseSetEventEntry(expected.Address, code)
	if !ok {
		t.Fatal("ParseSetEventEntry returned false for reconstructed bytes")
	}
	if got.Address != expected.Address {
		t.Errorf("Address: got 0x%x, want 0x%x", got.Address, expected.Address)
	}
	if got.Push != expected.Push {
		t.Errorf("Push: got 0x%08X, want 0x%08X", got.Push, expected.Push)
	}
	if got.Jump != expected.Jump {
		t.Errorf("Jump: got 0x%x, want 0x%x", got.Jump, expected.Jump)
	}
	t.Logf("reconstruction OK: %s", got.Format())
}

// --- 真实数据 VM OEP 验证 (OVERVIEW - SuperRecovery V4.8.1.5.txt) ---

func TestRealSuperRecovery_VMOEP(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile(superRecoveryOverviewPath)
	if err != nil {
		t.Skipf("OVERVIEW not found: %v", err)
	}

	oep, err := ParseOverviewOEP(data)
	if err != nil {
		t.Fatalf("ParseOverviewOEP failed: %v", err)
	}

	if oep.Addr != expectedVMAddr {
		t.Errorf("VM ADDR = 0x%x, want 0x%x", oep.Addr, expectedVMAddr)
	}
	if oep.Align != expectedVMAlign {
		t.Errorf("VM ALIGN = 0x%x, want 0x%x", oep.Align, expectedVMAlign)
	}
	if oep.Push != expectedVMPush {
		t.Errorf("VM PUSH = 0x%x, want 0x%x", oep.Push, expectedVMPush)
	}
	if oep.Jump != expectedVMJump {
		t.Errorf("VM JUMP = 0x%x, want 0x%x", oep.Jump, expectedVMJump)
	}
	t.Logf("VM OEP: ADDR=%X ALIGN=%X PUSH=%X JUMP=%X", oep.Addr, oep.Align, oep.Push, oep.Jump)

	want := "VM ADDR: 7A978B\nVM ALIGN: D695C014\nVM PUSH: 29E4826D\nVM JUMP: 6C5F95"
	if got := FormatVMOEP(oep.Addr, oep.Align, oep.Push, oep.Jump); got != want {
		t.Errorf("FormatVMOEP = %q, want %q", got, want)
	}

	if oep.Jump != expectedSetEventJump {
		t.Errorf("VM JUMP (0x%x) != SetEvent JUMP (0x%x) — expected same VM entry target", oep.Jump, expectedSetEventJump)
	} else {
		t.Logf("VM JUMP == SetEvent JUMP == 0x%x (consistent VM entry target)", oep.Jump)
	}
}

// --- VM OEP 与 SetEvent 交叉验证 ---

func TestRealSuperRecovery_SetEventAndVMOEPCrossCheck(t *testing.T) {
	t.Parallel()
	logData, err := os.ReadFile(superRecoveryLogPath)
	if err != nil {
		t.Skipf("API log not found: %v", err)
	}
	overviewData, err := os.ReadFile(superRecoveryOverviewPath)
	if err != nil {
		t.Skipf("OVERVIEW not found: %v", err)
	}

	log, err := ParseAPILog(logData)
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	oep, err := ParseOverviewOEP(overviewData)
	if err != nil {
		t.Fatalf("ParseOverviewOEP failed: %v", err)
	}

	if log.SetEventEntry != nil {
		if oep.Jump != log.SetEventEntry.Jump {
			t.Errorf("cross-check: VM JUMP (0x%x) != SetEvent JUMP (0x%x)", oep.Jump, log.SetEventEntry.Jump)
		} else {
			t.Logf("cross-check: VM JUMP == SetEvent JUMP == 0x%x", oep.Jump)
		}
	}

	if log.SetEventEntry != nil {
		if oep.Push != log.SetEventEntry.Push {
			t.Logf("cross-check: VM PUSH (0x%x) != SetEvent PUSH (0x%x) (expected: different push values)", oep.Push, log.SetEventEntry.Push)
		}
	}
}

// --- 集成测试：加载驱动 + 运行 SuperRecovery + 收集 hook 日志 ---
//
// 这是真正的端到端测试：加载 hyperkd.sys 驱动 → 启动 SuperRecovery.exe
// → 用 EptHookForProcess 设置 32-bit API hook → 收集 API Logger 格式日志
// → 用 ParseAPILog 解析 → 验证 SetEvent 入口与 OVERVIEW 预期值一致。
//
// OEP 不是在进程启动时找的——而是通过 hook API 收集调用日志，然后分析日志
// 推导出来的。这对应 OllyDbg 脚本 L1430+ 的流程。
//
// 需要管理员权限（加载驱动）。

const (
	integDriverPath     = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\Debug\hyperkd.sys`
	integExePath        = `d:\ux\examples\ewdk\debuger\Magicmida\SuperRecovery.exe`
	integKernelbasePath = `C:\Windows\SysWOW64\kernelbase.dll`
	integNtdllPath      = `C:\Windows\SysWOW64\ntdll.dll`
	integLogPath        = `d:\ux\examples\ewdk\tt\vt\good\todo\HyperDbgUnified\go-libhyperdbg\debugger\themida\test-oep.log`
)

func TestUnpacker_SuperRecovery_DriverIntegration(t *testing.T) {
	// 检查前置条件
	if _, err := os.ReadFile(integDriverPath); err != nil {
		t.Skipf("hyperkd.sys not found: %v", err)
	}
	if _, err := os.ReadFile(integExePath); err != nil {
		t.Skipf("SuperRecovery.exe not found: %v", err)
	}

	cfg := UnpackerConfig{
		DriverPath:     integDriverPath,
		ExePath:        integExePath,
		LogPath:        integLogPath,
		RunSeconds:     30,
		KernelbasePath: integKernelbasePath,
		NtdllPath:      integNtdllPath,
	}

	u := NewUnpacker(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	result, err := u.Run(ctx)
	if err != nil {
		t.Fatalf("Unpacker.Run failed: %v", err)
	}

	// 读取生成的日志
	logData, err := os.ReadFile(result.LogPath)
	if err != nil {
		t.Fatalf("read log %s: %v", result.LogPath, err)
	}

	logStr := string(logData)
	if len(logData) == 0 {
		t.Fatal("log file is empty — hooks did not fire (32-bit address mismatch?)")
	}

	// 检查日志是否有 API Logger 格式的行
	if !strings.Contains(logStr, "Call from:") {
		t.Fatalf("log does not contain API Logger format lines.\nLog content:\n%s", logStr)
	}

	// 统计 hook 触发次数
	vaCount := strings.Count(logStr, "NAME: VirtualAlloc")
	rahCount := strings.Count(logStr, "NAME: RtlAllocateHeap")
	seCount := strings.Count(logStr, "NAME: SetEvent")
	t.Logf("Hook hits: VirtualAlloc=%d RtlAllocateHeap=%d SetEvent=%d", vaCount, rahCount, seCount)

	// 用 ParseAPILog 解析日志
	log, err := ParseAPILog(logData)
	if err != nil {
		t.Fatalf("ParseAPILog failed: %v", err)
	}
	t.Logf("Parsed %d API call entries", len(log.Entries))

	// 验证 SetEvent 入口（如果检测到）
	if log.SetEventEntry != nil {
		t.Logf("SetEvent entry: Address=0x%X PUSH=0x%X JUMP=0x%X",
			log.SetEventEntry.Address, log.SetEventEntry.Push, log.SetEventEntry.Jump)
		if log.SetEventEntry.Jump == expectedVMJump {
			t.Logf("PASS: SetEvent JUMP (0x%X) == VM JUMP (0x%X)", log.SetEventEntry.Jump, expectedVMJump)
		} else {
			t.Errorf("SetEvent JUMP (0x%X) != expected VM JUMP (0x%X)", log.SetEventEntry.Jump, expectedVMJump)
		}
	} else {
		t.Logf("SetEvent entry not detected in log (Themida may not have reached SetEvent yet)")
	}

	// 验证 IO Marker（如果检测到）
	if log.IOMarkerAddress != 0 {
		t.Logf("IO Marker: 0x%X", log.IOMarkerAddress)
	} else {
		t.Logf("IO Marker not detected in log")
	}

	// 检查是否有 near-OEP 的 caller（地址 < 0x70000000）
	for _, e := range log.Entries {
		if e.Caller < 0x70000000 {
			t.Logf("Near-OEP caller: 0x%X (%s)", e.Caller, e.Name)
		}
	}
}
