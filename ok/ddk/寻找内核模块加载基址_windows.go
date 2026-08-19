package ddk

import (
	"encoding/binary"
	"fmt"
	"iter"
	"strings"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/saferwall/pe"
	"golang.org/x/arch/x86/x86asm"
	"golang.org/x/sys/windows"
)

type KernelModuleInfo struct {
	Name      string
	ImageBase uint64
	ImageSize uint32
	FullPath  string
}

type KernelModuleFinder struct {
	privilege *Privilege
}

func NewKernelModuleFinder() *KernelModuleFinder {
	p := NewPrivilege()
	p.Debug()
	p.LoadDriver()
	return &KernelModuleFinder{privilege: p}
}

func (f *KernelModuleFinder) Modules() ([]KernelModuleInfo, error) {
	for size := uint32(128 * 1024); ; {
		buf := make([]byte, size)
		e := windows.NtQuerySystemInformation(
			windows.SystemModuleInformation,
			unsafe.Pointer(&buf[0]),
			size,
			&size,
		)
		switch e {
		case windows.STATUS_INFO_LENGTH_MISMATCH:
			continue
		case nil:
			mods := byteslice.ToStruct[windows.RTL_PROCESS_MODULES](buf)
			modules := unsafe.Slice(&mods.Modules[0], int(mods.NumberOfModules))
			result := make([]KernelModuleInfo, len(modules))
			for i, mod := range modules {
				name := byteslice.ToString(mod.FullPathName[mod.OffsetToFileName:])
				result[i] = KernelModuleInfo{
					Name:      name,
					ImageBase: uint64(mod.ImageBase),
					ImageSize: mod.ImageSize,
					FullPath:  byteslice.ToString(mod.FullPathName[:]),
				}
			}
			return result, nil
		default:
			return nil, fmt.Errorf("NtQuerySystemInformation failed: %v", e)
		}
	}
}

func (f *KernelModuleFinder) ModuleByName(name string) KernelModuleInfo {
	name = strings.ToLower(name)
	modules := mylog.Check2(f.Modules())

	for _, mod := range modules {
		if strings.Contains(strings.ToLower(mod.Name), name) {
			return mod
		}
	}
	panic(fmt.Errorf("kernel module %q not found in loaded modules", name))
}

func (f *KernelModuleFinder) ModuleBaseByName(name string) uint64 {
	return f.ModuleByName(name).ImageBase
}

func (f *KernelModuleFinder) FindExportedSymbolAddress(moduleName string, symbolName string) uint64 {
	base := f.ModuleBaseByName(moduleName)
	rva := mylog.Check2(f.FindExportedSymbolRVA(moduleName, symbolName))
	addr := base + uint64(rva)
	mylog.Success("exported symbol found: " + symbolName + " at " + mylog.Hex(addr))
	return addr
}

func (f *KernelModuleFinder) FindExportedSymbolRVA(moduleName string, symbolName string) (uint32, error) {
	path := mylog.Check2(f.resolveModulePath(moduleName))
	peFile := mylog.Check2(pe.New(path, &pe.Options{}))
	defer func() { mylog.Check(peFile.Close()) }()
	mylog.Check(peFile.Parse())
	for _, fn := range peFile.Export.Functions {
		if fn.Name == symbolName {
			mylog.Info("export found: " + symbolName + " RVA=0x" + fmt.Sprintf("%06X", fn.FunctionRVA))
			return fn.FunctionRVA, nil
		}
	}
	return 0, fmt.Errorf("export %q not found in %s", symbolName, path)
}

func (f *KernelModuleFinder) FindNonExportedSymbolRVA(moduleName string, entryExportName string) iter.Seq2[string, x86asm.Inst] {
	return func(yield func(string, x86asm.Inst) bool) {
		path := mylog.Check2(f.resolveModulePath(moduleName))
		peFile := mylog.Check2(pe.New(path, &pe.Options{}))
		defer func() { mylog.Check(peFile.Close()) }()
		mylog.Check(peFile.Parse())
		var entryRVA uint32
		for _, fn := range peFile.Export.Functions {
			if fn.Name == entryExportName {
				entryRVA = fn.FunctionRVA
				break
			}
		}
		if entryRVA == 0 {
			return
		}
		mylog.Info("entry export: " + entryExportName + " RVA=0x" + fmt.Sprintf("%06X", entryRVA))

		for _, sec := range peFile.Sections {
			if !sec.Contains(entryRVA, peFile) {
				continue
			}
			data := sec.Data(0, 0, peFile)
			funcOff := int(entryRVA - sec.Header.VirtualAddress)

			const maxDecodeSize = 500
			end := min(funcOff+maxDecodeSize, len(data))
			code := data[funcOff:end]

			for line, inst := range DisassembleSeq2(code, int(entryRVA), nil) {
				if !yield(line, inst) {
					return
				}
			}
			return
		}
	}
}

func (f *KernelModuleFinder) TraceCallTarget(callerRVA uint32, callInstOffset int, inst x86asm.Inst) uint32 {
	if inst.Op != x86asm.CALL || inst.Len < 5 {
		return 0
	}
	for _, arg := range inst.Args {
		if rel, ok := arg.(x86asm.Rel); ok {
			return callerRVA + uint32(callInstOffset) + uint32(inst.Len) + uint32(rel)
		}
	}
	return 0
}

func (f *KernelModuleFinder) TraceLeaTarget(baseRVA uint32, leaInstOffset int, inst x86asm.Inst, code []byte) uint32 {
	if inst.Op != x86asm.LEA || inst.PCRel <= 0 {
		return 0
	}
	disp := binary.LittleEndian.Uint32(code[inst.PCRelOff : inst.PCRelOff+inst.PCRel])
	return baseRVA + uint32(leaInstOffset) + uint32(inst.Len) + disp
}

func (f *KernelModuleFinder) TraceMovRipRelative(baseRVA uint32, movInstOffset int, inst x86asm.Inst, code []byte) uint32 {
	if inst.Op != x86asm.MOV || inst.PCRel <= 0 {
		return 0
	}
	disp := binary.LittleEndian.Uint32(code[inst.PCRelOff : inst.PCRelOff+inst.PCRel])
	return baseRVA + uint32(movInstOffset) + uint32(inst.Len) + disp
}

func (f *KernelModuleFinder) resolveModulePath(moduleName string) (string, error) {
	moduleName = strings.ToLower(moduleName)
	modules := mylog.Check2(f.Modules())

	for _, mod := range modules {
		if strings.Contains(strings.ToLower(mod.Name), moduleName) {
			name := mod.Name
			if strings.Contains(name, `\`) {
				parts := strings.Split(name, `\`)
				name = parts[len(parts)-1]
			}
			return `C:\Windows\System32\` + name, nil
		}
	}
	return "", fmt.Errorf("kernel module %q not found", moduleName)
}
