//go:build windows

package cve

import (
	"iter"

	"golang.org/x/sys/windows"
)

func GetKernelBase() ProcessModuleInformation {
	for info := range DriversModules() {
		if info.FullPathName == "ntoskrnl.exe" {
			return info
		}
	}
	panic("ntoskrnl.exe not found")
}

func DriversModules() iter.Seq[ProcessModuleInformation] {
	return func(yield func(ProcessModuleInformation) bool) {
		for module := range NtQuerySystemInformation[windows.RTL_PROCESS_MODULE_INFORMATION](windows.SystemModuleInformation) {
			if !yield(ProcessModuleInformation{
				Section:          module.Section,
				MappedBase:       module.MappedBase,
				ImageBase:        module.ImageBase,
				ImageSize:        module.ImageSize,
				Flags:            module.Flags,
				LoadOrderIndex:   module.LoadOrderIndex,
				InitOrderIndex:   module.InitOrderIndex,
				LoadCount:        module.LoadCount,
				OffsetToFileName: module.OffsetToFileName,
				FullPathName:     windows.ByteSliceToString(module.FullPathName[module.OffsetToFileName:]),
			}) {
				return
			}
		}
	}
}

type ProcessModuleInformation struct {
	Section          windows.Handle
	MappedBase       uintptr
	ImageBase        uintptr
	ImageSize        uint32
	Flags            uint32
	LoadOrderIndex   uint16
	InitOrderIndex   uint16
	LoadCount        uint16
	OffsetToFileName uint16 // todo PathLength rename it ?
	FullPathName     string
}
