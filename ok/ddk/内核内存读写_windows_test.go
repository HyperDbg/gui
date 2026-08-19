package ddk

import (
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/arch/x86/x86asm"
)

func loadRTCore64(t *testing.T) *RTCore64 {
	t.Helper()
	rt := NewRTCore64()
	if !rt.Load() {
		t.Skip("RTCore64 driver not available (requires admin privileges)")
	}
	return rt
}

func TestPatternSearch(t *testing.T) {
	memData := []byte{
		0x33, 0xC9, 0x89, 0x0D, 0xB4, 0x67, 0x92, 0x77, 0x89, 0x0D,
		0xB8, 0x67, 0x92, 0x77, 0x88, 0x08, 0x38, 0x48, 0x02, 0x74,
		0x05, 0xE8, 0x94, 0xFF, 0xFF, 0xFF, 0x33, 0xC0, 0xC3, 0x8B,
		0xFF, 0x55, 0x8B, 0xEC, 0x83, 0xE4, 0xF8,
	}

	pattern := ParsePattern("?9 ?? 0? ?? 67")
	matches := pattern.SearchMemory(memData)

	t.Logf("Pattern: ?9 ?? 0? ?? 67")
	t.Logf("Matches found at offsets: %v", matches)

	expected := []int{1}
	if len(matches) != len(expected) {
		t.Fatalf("expected %d matches, got %d", len(expected), len(matches))
	}
	for i, m := range matches {
		if m != expected[i] {
			t.Errorf("match[%d]: expected %d, got %d", i, expected[i], m)
		}
	}
}

func TestPatternSearchExact(t *testing.T) {
	memData := []byte{0x48, 0x83, 0xEC, 0x68, 0x8B, 0x84, 0x24, 0xB8, 0x00, 0x00, 0x00}
	pattern := ParsePattern("48 83 EC")
	matches := pattern.SearchMemory(memData)

	t.Logf("Pattern: 48 83 EC")
	t.Logf("Matches: %v", matches)
	if len(matches) != 1 || matches[0] != 0 {
		t.Fatalf("expected match at offset 0, got %v", matches)
	}
}

func TestPatternSearchWildcard(t *testing.T) {
	memData := []byte{0xAA, 0xBB, 0xCC, 0xDD, 0xAA, 0xBB, 0xCC, 0xDD}
	pattern := ParsePattern("AA ?? CC")
	matches := pattern.SearchMemory(memData)

	t.Logf("Pattern: AA ?? CC")
	t.Logf("Matches: %v", matches)
	if len(matches) != 2 {
		t.Fatalf("expected 2 matches, got %d", len(matches))
	}
}

func BenchmarkPatternSearch(b *testing.B) {
	memData := []byte{
		0x33, 0xC9, 0x89, 0x0D, 0xB4, 0x67, 0x92, 0x77, 0x89, 0x0D,
		0xB8, 0x67, 0x92, 0x77, 0x88, 0x08, 0x38, 0x48, 0x02, 0x74,
		0x05, 0xE8, 0x94, 0xFF, 0xFF, 0xFF, 0x33, 0xC0, 0xC3, 0x8B,
		0xFF, 0x55, 0x8B, 0xEC, 0x83, 0xE4, 0xF8,
	}
	pattern := ParsePattern("?9 ?? 0? ?? 67")
	for b.Loop() {
		pattern.SearchMemory(memData)
	}
}

func TestSSDT_DecodeByDll(t *testing.T) {
	f := NewKernelModuleFinder()
	sc := NewSysCall(f)
	apis := sc.DecodeNtApiFromDLL(`C:\Windows\System32\ntdll.dll`)
	if len(apis) == 0 {
		t.Fatal("no NT APIs found from ntdll.dll")
	}
	t.Logf("Found %d NT APIs from ntdll.dll", len(apis))
	for i := 0; i < 10 && i < len(apis); i++ {
		t.Logf("  [%d] %s index=0x%X", i, apis[i].Name, apis[i].Index)
	}

	shadowApis := sc.DecodeNtApiFromDLL(`C:\Windows\System32\win32u.dll`)
	t.Logf("Found %d NT APIs from win32u.dll", len(shadowApis))
}

func TestSSDT_DecodeByDisassembly(t *testing.T) {
	f := NewKernelModuleFinder()
	sc := NewSysCall(f)
	if !sc.DecodeByDisassembly(`C:\Windows\System32\ntoskrnl.exe`) {
		t.Fatal("DecodeByDisassembly returned false")
	}

	t.Logf("KernelBase: %s", mylog.Hex(sc.KernelBase))
	t.Logf("KeServiceDescriptorTable       RVA: 0x%X  PhysicalAddr: %s",
		sc.OffsetKeServiceDescriptorTable, mylog.Hex(sc.SSDTPhysicalAddr()))
	t.Logf("KeServiceDescriptorTableShadow RVA: 0x%X  PhysicalAddr: %s",
		sc.OffsetKeServiceDescriptorTableShadow, mylog.Hex(sc.ShadowPhysicalAddr()))
	t.Logf("KeServiceDescriptorTableFilter RVA: 0x%X  PhysicalAddr: %s",
		sc.OffsetKeServiceDescriptorTableFilter, mylog.Hex(sc.FilterPhysicalAddr()))
	t.Logf("NT APIs from ntdll.dll: %d", len(sc.KeServiceDescriptorTable))
	t.Logf("NT APIs from win32u.dll: %d", len(sc.KeServiceDescriptorTableShadow))

	table := sc.RenderTable()
	t.Logf("SSDT Table:\n%s", table)

	sc.SaveTable("ssdt_table.txt")
}

func TestSSDT_WithKernelMemory(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)

	f := NewKernelModuleFinder()
	sc := NewSysCall(f)
	if !sc.DecodeByDisassembly(`C:\Windows\System32\ntoskrnl.exe`) {
		t.Fatal("DecodeByDisassembly returned false")
	}

	table := sc.RenderTableWithKernelAddrs(km)
	t.Logf("SSDT Table with kernel addresses:\n%s", table)

	sc.SaveTable("ssdt_table_live.txt")
}

func TestRTCore64_ReadKernelExportedFunction(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	addr := f.FindExportedSymbolAddress("ntoskrnl.exe", "NtDeviceIoControlFile")

	t.Logf("NtDeviceIoControlFile address: %s", mylog.Hex(addr))

	code := km.ReadCode(addr, 256)

	t.Logf("read %d bytes from kernel", len(code))

	disasm := DisassembleToString(code, addr)
	t.Logf("NtDeviceIoControlFile disassembly:\n%s", disasm)
}

func TestRTCore64_ReadKernelNonExported_IopXxxControlFile(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	var targetRVA uint32
	var prev x86asm.Inst
	for _, inst := range f.FindNonExportedSymbolRVA("ntoskrnl.exe", "NtDeviceIoControlFile") {
		if prev.Op == x86asm.CALL && inst.Op == x86asm.ADD {
			for _, arg := range prev.Args {
				if rel, ok := arg.(x86asm.Rel); ok {
					targetRVA = uint32(rel)
					break
				}
			}
			if targetRVA != 0 {
				break
			}
		}
		prev = inst
	}

	base := f.ModuleBaseByName("ntoskrnl.exe")
	addr := base + uint64(targetRVA)
	t.Logf("IopXxxControlFile address: %s", mylog.Hex(addr))

	code := km.ReadCode(addr, 512)

	disasm := DisassembleToString(code, addr)
	t.Logf("IopXxxControlFile disassembly:\n%s", disasm)
}

func TestRTCore64_ParsePESectionsFromKernel(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	base := f.ModuleBaseByName("ntoskrnl.exe")

	t.Logf("ntoskrnl.exe base: %s", mylog.Hex(base))

	sections := km.ParsePESections(base)

	t.Logf("ntoskrnl.exe has %d sections:", len(sections))
	for _, sec := range sections {
		t.Logf("  %-10s VA=0x%08X VSize=0x%08X RawAddr=0x%08X RawSize=0x%08X Char=0x%08X",
			sec.Name, sec.VirtualAddr, sec.VirtualSize, sec.RawAddr, sec.RawSize, sec.Characteristics)
	}
}

func TestRTCore64_MultipleKernelModules(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	modules := []struct {
		name      string
		exportSym string
	}{
		{"ntoskrnl.exe", "NtClose"},
		{"ci.dll", "CiInitialize"},
		{"hal.dll", "KeQueryPerformanceCounter"},
	}

	for _, m := range modules {
		t.Run(m.name, func(t *testing.T) {
			addr := f.FindExportedSymbolAddress(m.name, m.exportSym)

			t.Logf("%s!%s address: %s", m.name, m.exportSym, mylog.Hex(addr))

			code := km.ReadCode(addr, 64)

			disasm := DisassembleToString(code, addr)
			t.Logf("%s!%s disassembly:\n%s", m.name, m.exportSym, disasm)
		})
	}
}

func TestRTCore64_DisassembleKernelAPI(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	addr := f.FindExportedSymbolAddress("ntoskrnl.exe", "NtClose")

	t.Logf("NtClose address: %s", mylog.Hex(addr))

	result := km.DisassembleToString(addr, 128)
	t.Logf("NtClose disassembly:\n%s", result)
}

func TestRTCore64_ReadCiDllCiInitialize(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	addr := f.FindExportedSymbolAddress("ci.dll", "CiInitialize")

	t.Logf("CiInitialize address: %s", mylog.Hex(addr))

	code := km.ReadCode(addr, 256)

	disasm := DisassembleToString(code, addr)
	t.Logf("CiInitialize disassembly:\n%s", disasm)
}

func TestRTCore64_FindTextSection(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()

	base := f.ModuleBaseByName("ntoskrnl.exe")

	textVA, textSize := mylog.Check3(km.FindTextSection(base))

	t.Logf("ntoskrnl.exe .text section: VA=%s Size=0x%X", mylog.Hex(textVA), textSize)
}
