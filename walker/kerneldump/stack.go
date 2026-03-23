package kerneldump

import (
	"encoding/binary"
	"fmt"
)

type StackFrame struct {
	ChildSP  uint64
	RetAddr  uint64
	CallSite string
	FrameNum int
}

type Context struct {
	ContextFlags uint32
	Dr0          uint64
	Dr1          uint64
	Dr2          uint64
	Dr3          uint64
	Dr6          uint64
	Dr7          uint64
	FloatSave    [512]byte
	Gs           uint64
	Fs           uint64
	Es           uint64
	Ds           uint64
	Rax          uint64
	Rbx          uint64
	Rcx          uint64
	Rdx          uint64
	Rsp          uint64
	Rbp          uint64
	Rsi          uint64
	Rdi          uint64
	R8           uint64
	R9           uint64
	R10          uint64
	R11          uint64
	R12          uint64
	R13          uint64
	R14          uint64
	R15          uint64
	Rip          uint64
	Flgs         uint32
	Ss           uint16
	Cs           uint16
	DsPadding    uint16
	EspPadding   uint16
}

func (kd *KernelDump) GetContext() (*Context, error) {
	if kd.Header.ContextRecord == 0 {
		return nil, fmt.Errorf("no context record in dump")
	}

	offset := int(kd.Header.ContextRecord)
	if offset+512 > len(kd.Data) {
		return nil, fmt.Errorf("context record exceeds file size")
	}

	ctx := &Context{}
	ctx.ContextFlags = binary.LittleEndian.Uint32(kd.Data[offset : offset+4])
	ctx.Dr0 = binary.LittleEndian.Uint64(kd.Data[offset+16 : offset+24])
	ctx.Dr1 = binary.LittleEndian.Uint64(kd.Data[offset+24 : offset+32])
	ctx.Dr2 = binary.LittleEndian.Uint64(kd.Data[offset+32 : offset+40])
	ctx.Dr3 = binary.LittleEndian.Uint64(kd.Data[offset+40 : offset+48])
	ctx.Dr6 = binary.LittleEndian.Uint64(kd.Data[offset+48 : offset+56])
	ctx.Dr7 = binary.LittleEndian.Uint64(kd.Data[offset+56 : offset+64])
	copy(ctx.FloatSave[:], kd.Data[offset+64:offset+576])
	ctx.Gs = binary.LittleEndian.Uint64(kd.Data[offset+576 : offset+584])
	ctx.Fs = binary.LittleEndian.Uint64(kd.Data[offset+584 : offset+592])
	ctx.Es = binary.LittleEndian.Uint64(kd.Data[offset+592 : offset+600])
	ctx.Ds = binary.LittleEndian.Uint64(kd.Data[offset+600 : offset+608])
	ctx.Rax = binary.LittleEndian.Uint64(kd.Data[offset+608 : offset+616])
	ctx.Rbx = binary.LittleEndian.Uint64(kd.Data[offset+616 : offset+624])
	ctx.Rcx = binary.LittleEndian.Uint64(kd.Data[offset+624 : offset+632])
	ctx.Rdx = binary.LittleEndian.Uint64(kd.Data[offset+632 : offset+640])
	ctx.Rsp = binary.LittleEndian.Uint64(kd.Data[offset+640 : offset+648])
	ctx.Rbp = binary.LittleEndian.Uint64(kd.Data[offset+648 : offset+656])
	ctx.Rsi = binary.LittleEndian.Uint64(kd.Data[offset+656 : offset+664])
	ctx.Rdi = binary.LittleEndian.Uint64(kd.Data[offset+664 : offset+672])
	ctx.R8 = binary.LittleEndian.Uint64(kd.Data[offset+672 : offset+680])
	ctx.R9 = binary.LittleEndian.Uint64(kd.Data[offset+680 : offset+688])
	ctx.R10 = binary.LittleEndian.Uint64(kd.Data[offset+688 : offset+696])
	ctx.R11 = binary.LittleEndian.Uint64(kd.Data[offset+696 : offset+704])
	ctx.R12 = binary.LittleEndian.Uint64(kd.Data[offset+704 : offset+712])
	ctx.R13 = binary.LittleEndian.Uint64(kd.Data[offset+712 : offset+720])
	ctx.R14 = binary.LittleEndian.Uint64(kd.Data[offset+720 : offset+728])
	ctx.R15 = binary.LittleEndian.Uint64(kd.Data[offset+728 : offset+736])
	ctx.Rip = binary.LittleEndian.Uint64(kd.Data[offset+736 : offset+744])
	ctx.Flgs = binary.LittleEndian.Uint32(kd.Data[offset+744 : offset+748])
	ctx.Ss = binary.LittleEndian.Uint16(kd.Data[offset+748 : offset+750])
	ctx.Cs = binary.LittleEndian.Uint16(kd.Data[offset+750 : offset+752])
	ctx.DsPadding = binary.LittleEndian.Uint16(kd.Data[offset+752 : offset+754])
	ctx.EspPadding = binary.LittleEndian.Uint16(kd.Data[offset+754 : offset+756])

	return ctx, nil
}

func (kd *KernelDump) GetStackTrace() ([]StackFrame, error) {
	ctx, err := kd.GetContext()
	if err != nil {
		return nil, err
	}

	var frames []StackFrame

	currentSP := ctx.Rsp
	frameNum := 0

	for range 50 {
		stackData, err := kd.readVirtualMemory(currentSP, 16)
		if err != nil {
			break
		}

		retAddr := binary.LittleEndian.Uint64(stackData[0:8])
		nextSP := binary.LittleEndian.Uint64(stackData[8:16])

		if retAddr == 0 || retAddr == 0xFFFFFFFFFFFFFFFF {
			break
		}

		frame := StackFrame{
			ChildSP:  currentSP,
			RetAddr:  retAddr,
			FrameNum: frameNum,
		}

		modName, funcName, offset := kd.resolveSymbol(retAddr)
		if modName != "" {
			frame.CallSite = fmt.Sprintf("%s!%s+0x%X", modName, funcName, offset)
		} else {
			frame.CallSite = fmt.Sprintf("0x%016X", retAddr)
		}

		frames = append(frames, frame)

		if nextSP <= currentSP {
			break
		}

		currentSP = nextSP
		frameNum++
	}

	return frames, nil
}

func (kd *KernelDump) resolveSymbol(addr uint64) (string, string, uint64) {
	modules, err := kd.GetModules()
	if err != nil {
		return "", "", 0
	}

	for _, mod := range modules {
		if mod.ContainsAddress(addr) {
			name := mod.GetName()
			offset := addr - mod.DllBase
			return name, "", offset
		}
	}

	return "", "", 0
}

func (kd *KernelDump) DetectBugCheckFromStack() (uint32, string) {
	frames, err := kd.GetStackTrace()
	if err != nil {
		return kd.Header.BugCheckCode, ""
	}

	for _, frame := range frames {
		if frame.RetAddr >= 0 && frame.RetAddr <= 0xFFFFFFFF {
			bugCheck := uint32(frame.RetAddr)
			if bugCheck != 0 && bugCheck != 0x8 {
				return bugCheck, kd.GetBugCheckNameFromCode(bugCheck)
			}
		}
	}

	return kd.Header.BugCheckCode, ""
}

func (kd *KernelDump) GetBugCheckNameFromCode(code uint32) string {
	codes := map[uint32]string{
		0x00000001: "APC_INDEX_MISMATCH",
		0x00000003: "INVALID_PROCESS_ATTACH_ATTEMPT",
		0x00000004: "INVALID_DATA_ACCESS_TRAP",
		0x00000008: "IRQL_NOT_LESS_OR_EQUAL",
		0x0000000A: "IRQL_NOT_LESS_OR_EQUAL",
		0x0000001A: "MEMORY_MANAGEMENT",
		0x0000001E: "KMODE_EXCEPTION_NOT_HANDLED",
		0x00000024: "NTFS_FILE_SYSTEM",
		0x0000002E: "DATA_BUS_ERROR",
		0x0000003B: "SYSTEM_SERVICE_EXCEPTION",
		0x0000003D: "INTERRUPT_EXCEPTION_NOT_HANDLED",
		0x0000003F: "NO_MORE_SYSTEM_PTES",
		0x00000044: "MULTIPLE_IRP_COMPLETE_REQUESTS",
		0x0000004D: "NO_PAGES_AVAILABLE",
		0x00000050: "PAGE_FAULT_IN_NONPAGED_AREA",
		0x00000051: "REGISTRY_ERROR",
		0x00000074: "BAD_SYSTEM_CONFIG_INFO",
		0x0000007B: "INACCESSIBLE_BOOT_DEVICE",
		0x0000007E: "SYSTEM_THREAD_EXCEPTION_NOT_HANDLED",
		0x0000007F: "UNEXPECTED_KERNEL_MODE_TRAP",
		0x0000008E: "KERNEL_MODE_EXCEPTION_NOT_HANDLED",
		0x0000009F: "DRIVER_POWER_STATE_FAILURE",
		0x000000A5: "ACPI_BIOS_ERROR",
		0x000000BE: "ATTEMPTED_WRITE_TO_READONLY_MEMORY",
		0x000000C2: "BAD_POOL_CALLER",
		0x000000C5: "DRIVER_CORRUPTED_EXPOOL",
		0x000000D1: "DRIVER_IRQL_NOT_LESS_OR_EQUAL",
		0x000000EA: "THREAD_STUCK_IN_DEVICE_DRIVER",
		0x000000ED: "UNMOUNTABLE_BOOT_VOLUME",
		0x000000EF: "CRITICAL_PROCESS_DIED",
		0x00000109: "CRITICAL_STRUCTURE_CORRUPTION",
		0x00000119: "VIDEO_SCHEDULER_INTERNAL_ERROR",
		0x00000133: "DPC_WATCHDOG_VIOLATION",
		0x00000139: "KERNEL_SECURITY_CHECK_FAILURE",
		0x00000144: "BUGCODE_USB_DRIVER",
		0x00000193: "BAD_OBJECT_HEADER",
	}

	if name, ok := codes[code]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_0x%08X", code)
}
