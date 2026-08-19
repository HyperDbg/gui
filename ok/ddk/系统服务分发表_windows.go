package ddk

import (
	"encoding/binary"
	"fmt"
	"sort"
	"strings"

	"github.com/aquasecurity/table"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/saferwall/pe"
	"golang.org/x/arch/x86/x86asm"
)

type NtApi struct {
	Name  string
	Index uint32
}

type SysCall struct {
	finder                               *KernelModuleFinder
	KernelBase                           uint64
	OffsetKeServiceDescriptorTable       uint32
	OffsetKeServiceDescriptorTableShadow uint32
	OffsetKeServiceDescriptorTableFilter uint32
	KeServiceDescriptorTable             []NtApi
	KeServiceDescriptorTableShadow       []NtApi
}

func NewSysCall(finder *KernelModuleFinder) *SysCall {
	return &SysCall{
		finder:                         finder,
		KeServiceDescriptorTable:       make([]NtApi, 0),
		KeServiceDescriptorTableShadow: make([]NtApi, 0),
	}
}

func (s *SysCall) SSDTPhysicalAddr() uint64 {
	return s.KernelBase + uint64(s.OffsetKeServiceDescriptorTable)
}

func (s *SysCall) ShadowPhysicalAddr() uint64 {
	return s.KernelBase + uint64(s.OffsetKeServiceDescriptorTableShadow)
}

func (s *SysCall) FilterPhysicalAddr() uint64 {
	return s.KernelBase + uint64(s.OffsetKeServiceDescriptorTableFilter)
}

func (s *SysCall) DecodeNtApiFromDLL(filename string) []NtApi {
	f := mylog.Check2(pe.New(filename, &pe.Options{}))

	mylog.Check(f.Parse())
	defer func() { mylog.Check(f.Close()) }()

	var apis []NtApi
	ntCount := 0
	for _, entry := range f.Export.Functions {
		if entry.Name == "NtGetTickCount" {
			continue
		}
		if strings.HasPrefix(entry.Name, "Ntdll") {
			continue
		}
		if !strings.HasPrefix(entry.Name, "Nt") {
			continue
		}
		ntCount++

		data := mylog.Check2(f.GetData(entry.FunctionRVA, 32))

		var index uint32
		found := false
		for _, inst := range DisassembleSeq2(data, 0, nil) {
			if inst.Op == x86asm.MOV {
				for i, arg := range inst.Args {
					if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.EAX {
						if i+1 < len(inst.Args) {
							if imm, ok := inst.Args[i+1].(x86asm.Imm); ok && imm >= 0 && imm <= 4096 {
								index = uint32(imm)
								found = true
								break
							}
						}
					}
				}
			}
			if found {
				break
			}
			if inst.Op == x86asm.SYSCALL || inst.Op == x86asm.RET {
				break
			}
		}
		if !found {
			continue
		}

		if strings.Contains(filename, "win32") {
			if index == 0 {
				continue
			}
			if index >= 4096 {
				index -= 4096
			}
		}

		apis = append(apis, NtApi{Name: entry.Name, Index: index})
	}

	sort.Slice(apis, func(i, j int) bool {
		return apis[i].Index < apis[j].Index
	})
	return apis
}

func (s *SysCall) DecodeByDisassembly(ntoskrnlPath string) bool {
	f := mylog.Check2(pe.New(ntoskrnlPath, &pe.Options{}))
	mylog.Check(f.Parse())
	defer func() { mylog.Check(f.Close()) }()

	var zwDevIoCtrlRVA uint32
	for _, entry := range f.Export.Functions {
		if entry.Name == "ZwDeviceIoControlFile" {
			zwDevIoCtrlRVA = entry.FunctionRVA
			break
		}
	}
	if zwDevIoCtrlRVA == 0 {
		mylog.Warning("ZwDeviceIoControlFile not found in exports")
		return false
	}

	data := mylog.Check2(f.GetData(zwDevIoCtrlRVA, 500))

	kiServiceInternalRVA, syscallNum := s.findKiServiceInternal(data, zwDevIoCtrlRVA)
	if kiServiceInternalRVA == 0 {
		mylog.Warning("KiServiceInternal not found")
		return false
	}
	mylog.Info("KiServiceInternal RVA", "rva", fmt.Sprintf("0x%X", kiServiceInternalRVA), "syscallNum", fmt.Sprintf("0x%X", syscallNum))

	kiServiceInternalData := mylog.Check2(f.GetData(kiServiceInternalRVA, 500))

	kiSystemServiceStartRVA := s.findKiSystemServiceStart(kiServiceInternalData, kiServiceInternalRVA)
	if kiSystemServiceStartRVA == 0 {
		mylog.Warning("KiSystemServiceStart not found")
		return false
	}
	mylog.Info("KiSystemServiceStart RVA", "rva", fmt.Sprintf("0x%X", kiSystemServiceStartRVA))

	kiSystemServiceStartData := mylog.Check2(f.GetData(uint32(kiSystemServiceStartRVA), 500))

	offsetSSDT, offsetShadow, offsetFilter := s.findSSDTOffsets(kiSystemServiceStartData, uint32(kiSystemServiceStartRVA))
	if offsetSSDT == 0 {
		mylog.Warning("SSDT offsets not found")
		return false
	}

	if s.finder != nil {
		s.KernelBase = s.finder.ModuleBaseByName("ntoskrnl.exe")
	}

	s.OffsetKeServiceDescriptorTable = offsetSSDT
	s.OffsetKeServiceDescriptorTableShadow = offsetShadow
	s.OffsetKeServiceDescriptorTableFilter = offsetFilter

	s.KeServiceDescriptorTable = s.DecodeNtApiFromDLL(`C:\Windows\System32\ntdll.dll`)
	s.KeServiceDescriptorTableShadow = s.DecodeNtApiFromDLL(`C:\Windows\System32\win32u.dll`)

	return true
}

func (s *SysCall) findKiServiceInternal(data []byte, baseRVA uint32) (rva uint32, syscallNum uint32) {
	var prev x86asm.Inst
	var offset uint32
	for _, inst := range DisassembleSeq2(data, int(baseRVA), nil) {
		if inst.Op == x86asm.INT {
			break
		}
		if prev.Op == x86asm.MOV && inst.Op == x86asm.JMP {
			for i2, arg := range prev.Args {
				if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.EAX {
					if i2+1 < len(prev.Args) {
						if imm, ok := prev.Args[i2+1].(x86asm.Imm); ok {
							syscallNum = uint32(imm)
						}
					}
				}
			}
			for _, arg := range inst.Args {
				if rel, ok := arg.(x86asm.Rel); ok {
					return baseRVA + offset + uint32(rel), syscallNum
				}
			}
		}
		offset += uint32(prev.Len)
		prev = inst
	}
	return 0, 0
}

func (s *SysCall) findKiSystemServiceStart(data []byte, baseRVA uint32) uint64 {
	var prev x86asm.Inst
	var offset uint32
	for _, inst := range DisassembleSeq2(data, int(baseRVA), nil) {
		if inst.Op == x86asm.INT {
			break
		}
		if prev.Op == x86asm.LEA && inst.Op == x86asm.JMP {
			for _, arg := range prev.Args {
				if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.R11 {
					for _, arg2 := range prev.Args {
						if mem, ok := arg2.(x86asm.Mem); ok && mem.Base == x86asm.RIP {
							totalOffset := offset + uint32(inst.Len)
							return uint64(baseRVA) + uint64(totalOffset) + uint64(mem.Disp)
						}
					}
				}
			}
		}
		if inst.Op == x86asm.JMP {
			for _, arg := range inst.Args {
				if rel, ok := arg.(x86asm.Rel); ok {
					totalOffset := offset + uint32(inst.Len)
					return uint64(baseRVA) + uint64(totalOffset) + uint64(rel)
				}
			}
		}
		offset += uint32(prev.Len)
		prev = inst
	}
	return 0
}

func (s *SysCall) findSSDTOffsets(data []byte, baseRVA uint32) (ssdt, shadow, filter uint32) {
	const windowSize = 7
	var window [windowSize]x86asm.Inst
	var size [windowSize]uint32
	wi := 0
	totalSize := uint32(0)

	for _, inst := range DisassembleSeq2(data, int(baseRVA), nil) {
		if inst.Op == x86asm.INT {
			break
		}
		window[wi] = inst
		size[wi] = totalSize
		totalSize += uint32(inst.Len)
		wi = (wi + 1) % windowSize

		if wi != 0 {
			continue
		}

		inst := window[0]
		next := window[1]
		next1 := window[2]
		next2 := window[3]
		next3 := window[4]
		next4 := window[5]
		next5 := window[6]

		if inst.Op == x86asm.LEA &&
			next.Op == x86asm.LEA &&
			next1.Op == x86asm.TEST &&
			next2.Op == x86asm.JE &&
			next3.Op == x86asm.TEST &&
			next4.Op == x86asm.JE &&
			next5.Op == x86asm.LEA {

			for i2, arg := range inst.Args {
				if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.R10 {
					if i2+1 < len(inst.Args) {
						if mem, ok := inst.Args[i2+1].(x86asm.Mem); ok {
							ssdt = baseRVA + uint32(mem.Disp) + size[0]
						}
					}
				}
			}

			for i2, arg := range next.Args {
				if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.R11 {
					if i2+1 < len(next.Args) {
						if mem, ok := next.Args[i2+1].(x86asm.Mem); ok {
							size2 := size[0] + uint32(next.Len)
							shadow = baseRVA + uint32(mem.Disp) + size2
						}
					}
				}
			}

			for i2, arg := range next5.Args {
				if reg, ok := arg.(x86asm.Reg); ok && reg == x86asm.R11 {
					if i2+1 < len(next5.Args) {
						if mem, ok := next5.Args[i2+1].(x86asm.Mem); ok {
							size3 := size[0] + uint32(next.Len) + uint32(next1.Len) + uint32(next2.Len) + uint32(next3.Len) + uint32(next4.Len) + uint32(next5.Len)
							filter = baseRVA + uint32(mem.Disp) + size3
						}
					}
				}
			}
			return
		}
	}
	return 0, 0, 0
}

func (s *SysCall) RenderTable() string {
	buf := stream.NewBuffer("")

	buf.WriteStringLn("Windows System Service Dispatch Table (SSDT)")
	buf.WriteStringLn(fmt.Sprintf("KernelBase: 0x%X", s.KernelBase))
	buf.WriteStringLn("")

	buf.WriteStringLn("=== KeServiceDescriptorTable ===")
	buf.WriteStringLn(fmt.Sprintf("  RVA: 0x%X  PhysicalAddr: 0x%X", s.OffsetKeServiceDescriptorTable, s.SSDTPhysicalAddr()))
	t := table.New(buf)
	t.SetRowLines(false)
	t.SetHeaders("Id", "Name", "Index", "PhysicalAddr")
	for i, api := range s.KeServiceDescriptorTable {
		t.AddRow(
			fmt.Sprint(i+1),
			api.Name,
			fmt.Sprintf("0x%X", api.Index),
			fmt.Sprintf("0x%X", s.KernelBase),
		)
	}
	t.Render()

	buf.WriteStringLn("")
	buf.WriteStringLn("=== KeServiceDescriptorTableShadow ===")
	buf.WriteStringLn(fmt.Sprintf("  RVA: 0x%X  PhysicalAddr: 0x%X", s.OffsetKeServiceDescriptorTableShadow, s.ShadowPhysicalAddr()))
	t2 := table.New(buf)
	t2.SetRowLines(false)
	t2.SetHeaders("Id", "Name", "Index")
	for i, api := range s.KeServiceDescriptorTableShadow {
		t2.AddRow(
			fmt.Sprint(i+1),
			api.Name,
			fmt.Sprintf("0x%X", api.Index),
		)
	}
	t2.Render()

	buf.WriteStringLn("")
	buf.WriteStringLn("=== KeServiceDescriptorTableFilter ===")
	buf.WriteStringLn(fmt.Sprintf("  RVA: 0x%X  PhysicalAddr: 0x%X", s.OffsetKeServiceDescriptorTableFilter, s.FilterPhysicalAddr()))

	return buf.String()
}

func (s *SysCall) SaveTable(filename string) {
	content := s.RenderTable()
	stream.WriteTruncate(filename, content)
	mylog.Success("SSDT table saved", "file", filename)
}

func (s *SysCall) ReadSSDTFromKernel(km *KernelMemory) ([]uint64, error) {
	if s.OffsetKeServiceDescriptorTable == 0 {
		return nil, fmt.Errorf("SSDT offset not resolved")
	}
	ssdtAddr := s.SSDTPhysicalAddr()
	numEntries := len(s.KeServiceDescriptorTable)
	if numEntries == 0 {
		numEntries = 512
	}

	buf := make([]byte, numEntries*4)
	km.ReadMemory(ssdtAddr, buf)

	entries := make([]uint64, numEntries)
	for i := 0; i < numEntries; i++ {
		rva := binary.LittleEndian.Uint32(buf[i*4 : (i+1)*4])
		entries[i] = s.KernelBase + uint64(rva)
	}
	return entries, nil
}

func (s *SysCall) RenderTableWithKernelAddrs(km *KernelMemory) string {
	entries := mylog.Check2(s.ReadSSDTFromKernel(km))

	buf := stream.NewBuffer("")
	buf.WriteStringLn("Windows System Service Dispatch Table (SSDT) - Live Kernel Addresses")
	buf.WriteStringLn(fmt.Sprintf("KernelBase: 0x%X", s.KernelBase))
	buf.WriteStringLn("")

	buf.WriteStringLn("=== KeServiceDescriptorTable ===")
	buf.WriteStringLn(fmt.Sprintf("  RVA: 0x%X  PhysicalAddr: 0x%X", s.OffsetKeServiceDescriptorTable, s.SSDTPhysicalAddr()))
	t := table.New(buf)
	t.SetRowLines(false)
	t.SetHeaders("Id", "Name", "Index", "KernelAddr")
	for i, api := range s.KeServiceDescriptorTable {
		kernelAddr := ""
		if i < len(entries) {
			kernelAddr = fmt.Sprintf("0x%X", entries[i])
		}
		t.AddRow(
			fmt.Sprint(i+1),
			api.Name,
			fmt.Sprintf("0x%X", api.Index),
			kernelAddr,
		)
	}
	t.Render()

	return buf.String()
}
