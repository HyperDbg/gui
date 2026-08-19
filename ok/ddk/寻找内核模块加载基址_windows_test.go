package ddk

import (
	"fmt"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/arch/x86/x86asm"
)

func TestQueryKernelModules(t *testing.T) {
	f := NewKernelModuleFinder()
	modules := mylog.Check2(f.Modules())
	mylog.Struct(modules)
}

func TestFindCiDllBase(t *testing.T) {
	f := NewKernelModuleFinder()
	mod := f.ModuleByName("ci.dll")
	mylog.Struct(mod)
}

func TestFindNtoskrnlBase(t *testing.T) {
	f := NewKernelModuleFinder()
	base := f.ModuleBaseByName("ntoskrnl.exe")
	mylog.Hex(base)
}

func TestFindExportedSymbolRVA_Ntoskrnl(t *testing.T) {
	f := NewKernelModuleFinder()
	rva := mylog.Check2(f.FindExportedSymbolRVA("ntoskrnl.exe", "NtDeviceIoControlFile"))

	t.Logf("NtDeviceIoControlFile RVA: 0x%06X", rva)
	if rva == 0 {
		t.Fatal("RVA should not be 0")
	}
}

func TestFindExportedSymbolRVA_CiDll(t *testing.T) {
	f := NewKernelModuleFinder()
	rva := mylog.Check2(f.FindExportedSymbolRVA("ci.dll", "CiInitialize"))

	t.Logf("CiInitialize RVA: 0x%06X", rva)
	if rva == 0 {
		t.Fatal("RVA should not be 0")
	}
}

func TestFindExportedSymbolAddress_Ntoskrnl(t *testing.T) {
	f := NewKernelModuleFinder()
	addr := f.FindExportedSymbolAddress("ntoskrnl.exe", "NtDeviceIoControlFile")
	mylog.Hex(addr)
}

// 0xFFFFF807E96F5E89 e812000000       call 0xfffff807e96f5ea0
func TestFindNonExportedSymbol_IopXxxControlFile(t *testing.T) {
	f := NewKernelModuleFinder()
	var targetRVA uint32
	for _, inst := range f.FindNonExportedSymbolRVA("ntoskrnl.exe", "NtDeviceIoControlFile") {
		if inst.Op != x86asm.CALL {
			continue
		}
		for _, arg := range inst.Args {
			if rel, ok := arg.(x86asm.Rel); ok {
				targetRVA = uint32(rel)
				break
			}
		}
		if targetRVA != 0 {
			break
		}
	}

	base := f.ModuleBaseByName("ntoskrnl.exe")
	addr := base + uint64(targetRVA)
	t.Logf("IopXxxControlFile address: %s", mylog.Hex(addr))
}

func TestFindNonExportedSymbol_MiGetPteAddress(t *testing.T) {
	f := NewKernelModuleFinder()
	var targetRVA uint32
	var prev x86asm.Inst
	count := 0
	for _, inst := range f.FindNonExportedSymbolRVA("ntoskrnl.exe", "MmFreeNonCachedMemory") {
		if prev.Op == x86asm.SUB && inst.Op == x86asm.MOV {
			for _, arg := range inst.Args {
				if rel, ok := arg.(x86asm.Rel); ok {
					targetRVA = uint32(rel)
					break
				}
			}
			if targetRVA != 0 {
				break
			}
		}
		count++
		if count > 100 {
			break
		}
		prev = inst
	}

	base := f.ModuleBaseByName("ntoskrnl.exe")
	addr := base + uint64(targetRVA)
	t.Logf("MiGetPteAddress address: %s", mylog.Hex(addr))
}

func TestFindNonExportedSymbol_KeServiceDescriptorTable(t *testing.T) {
	f := NewKernelModuleFinder()
	var targetRVA uint32
	for _, inst := range f.FindNonExportedSymbolRVA("ntoskrnl.exe", "ZwDeviceIoControlFile") {
		if inst.Op != x86asm.MOV {
			continue
		}
		var sysCallNum uint32
		for _, arg := range inst.Args {
			if imm, ok := arg.(x86asm.Imm); ok {
				sysCallNum = uint32(imm)
				break
			}
		}
		if sysCallNum != 0x7 {
			continue
		}
		for _, arg := range inst.Args {
			if rel, ok := arg.(x86asm.Rel); ok {
				targetRVA = uint32(rel)
				break
			}
		}
		if targetRVA != 0 {
			break
		}
	}

	base := f.ModuleBaseByName("ntoskrnl.exe")
	addr := base + uint64(targetRVA)
	t.Logf("KiServiceInternal address: %s", mylog.Hex(addr))
}

func TestAllKernelModuleBases(t *testing.T) {
	targetModules := []string{
		"ntoskrnl.exe",
		"ci.dll",
		"hal.dll",
		"ndis.sys",
		"ntfs.sys",
		"fltmgr.sys",
		"ksecdd.sys",
		"clipsp.sys",
		"cng.sys",
		"pcw.sys",
	}
	f := NewKernelModuleFinder()
	modules := mylog.Check2(f.Modules())

	found := map[string]uint64{}
	for _, mod := range modules {
		for _, target := range targetModules {
			if containsIgnoreCase(mod.Name, target) {
				found[target] = mod.ImageBase
			}
		}
	}
	for _, target := range targetModules {
		base, ok := found[target]
		if ok {
			t.Logf("  %-20s Base=%s", target, mylog.Hex(base))
		} else {
			t.Logf("  %-20s NOT FOUND", target)
		}
	}
}

func containsIgnoreCase(s, substr string) bool {
	s = fmt.Sprintf("%s", s)
	for i := 0; i <= len(s)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			c1 := s[i+j]
			c2 := substr[j]
			if c1 >= 'A' && c1 <= 'Z' {
				c1 += 32
			}
			if c2 >= 'A' && c2 <= 'Z' {
				c2 += 32
			}
			if c1 != c2 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
