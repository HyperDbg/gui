package kerneldump

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type Module struct {
	InLoadOrderLinks           [16]byte
	InMemoryOrderLinks         [16]byte
	InInitializationOrderLinks [16]byte
	DllBase                    uint64
	EntryPoint                 uint64
	SizeOfImage                uint32
	FullDllName                [256]byte
	BaseDllName                [64]byte
}

func (kd *KernelDump) GetModules() ([]Module, error) {
	if kd.Header.PsLoadedModuleList == 0 {
		return nil, fmt.Errorf("PsLoadedModuleList not available")
	}

	var modules []Module

	listHead := kd.Header.PsLoadedModuleList

	current := listHead
	visited := make(map[uint64]bool)

	for i := 0; i < 1000; i++ {
		if visited[current] {
			break
		}
		visited[current] = true

		moduleData, err := kd.readVirtualMemory(current, 0x150)
		if err != nil {
			break
		}

		var mod Module
		mod.InLoadOrderLinks = [16]byte(moduleData[0:16])
		mod.DllBase = binary.LittleEndian.Uint64(moduleData[24:32])
		mod.EntryPoint = binary.LittleEndian.Uint64(moduleData[32:40])
		mod.SizeOfImage = binary.LittleEndian.Uint32(moduleData[40:44])
		mod.FullDllName = [256]byte(moduleData[72:328])
		mod.BaseDllName = [64]byte(moduleData[328:392])

		if mod.DllBase != 0 && mod.SizeOfImage > 0 {
			modules = append(modules, mod)
		}

		flink := binary.LittleEndian.Uint64(moduleData[0:8])
		if flink == listHead || flink == 0 {
			break
		}
		current = flink
	}

	return modules, nil
}

func (kd *KernelDump) readVirtualMemory(virtualAddr uint64, size uint64) ([]byte, error) {
	if kd.Header.DirectoryTableBase == 0 {
		return nil, fmt.Errorf("no page table available")
	}

	physicalAddr, err := kd.virtualToPhysical(virtualAddr)
	if err != nil {
		return nil, err
	}

	return kd.ReadMemory(physicalAddr, size)
}

func (kd *KernelDump) virtualToPhysical(virtualAddr uint64) (uint64, error) {
	if virtualAddr >= 0xFFFF800000000000 {
		return virtualAddr - 0xFFFF800000000000, nil
	}
	if virtualAddr >= 0xFFFF000000000000 {
		return virtualAddr - 0xFFFF000000000000, nil
	}
	return 0, fmt.Errorf("unsupported virtual address: 0x%016X", virtualAddr)
}

func (mod *Module) GetName() string {
	name := string(mod.BaseDllName[:])
	end := strings.IndexByte(name, 0)
	if end >= 0 {
		name = name[:end]
	}
	return name
}

func (mod *Module) GetFullName() string {
	name := string(mod.FullDllName[:])
	end := strings.IndexByte(name, 0)
	if end >= 0 {
		name = name[:end]
	}
	return name
}

func (mod *Module) ContainsAddress(addr uint64) bool {
	return addr >= mod.DllBase && addr < mod.DllBase+uint64(mod.SizeOfImage)
}
