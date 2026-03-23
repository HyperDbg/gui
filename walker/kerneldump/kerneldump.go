package kerneldump

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	SignaturePagedu64 = "PAGEDU64"
	SignatureFull     = "DUMP64"
)

type Header struct {
	Signature             [8]byte
	ValidDump             uint16
	MajorVersion          uint16
	MinorVersion          uint16
	DirectoryTableBase    uint64
	PfnDatabase           uint64
	PsLoadedModuleList    uint64
	PsActiveProcessHead   uint64
	MachineImageType      uint16
	NumberProcessors      uint8
	BugCheckCode          uint32
	BugCheckParameter1    uint64
	BugCheckParameter2    uint64
	BugCheckParameter3    uint64
	BugCheckParameter4    uint64
	Reserved0             [8]byte
	KdDebuggerDataBlock   uint64
	PhysicalMemoryBlock   uint64
	PhysicalMemoryBlock64 uint64
	ContextRecord         uint64
	DumpType              uint32
	RequiredDumpSpace     uint32
	Reserved1             [4]byte
	SystemTime            uint64
	Reserved2             [8]byte
	FeatureFlags          uint64
	RequiredDumpSpace2    uint64
}

type KernelDump struct {
	Header Header
	File   *os.File
	Data   []byte
}

func Parse(path string) (*KernelDump, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	kd := &KernelDump{
		File: file,
		Data: data,
	}

	if len(data) < 512 {
		file.Close()
		return nil, fmt.Errorf("file too small")
	}

	kd.Header.Signature = [8]byte{data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7]}
	kd.Header.ValidDump = binary.LittleEndian.Uint16(data[8:10])
	kd.Header.MajorVersion = binary.LittleEndian.Uint16(data[10:12])
	kd.Header.MinorVersion = binary.LittleEndian.Uint16(data[12:14])
	kd.Header.DirectoryTableBase = binary.LittleEndian.Uint64(data[16:24])
	kd.Header.PfnDatabase = binary.LittleEndian.Uint64(data[24:32])
	kd.Header.PsLoadedModuleList = binary.LittleEndian.Uint64(data[32:40])
	kd.Header.PsActiveProcessHead = binary.LittleEndian.Uint64(data[40:48])
	kd.Header.MachineImageType = binary.LittleEndian.Uint16(data[48:50])
	kd.Header.NumberProcessors = data[50]
	kd.Header.BugCheckCode = binary.LittleEndian.Uint32(data[52:56])
	kd.Header.BugCheckParameter1 = binary.LittleEndian.Uint64(data[56:64])
	kd.Header.BugCheckParameter2 = binary.LittleEndian.Uint64(data[64:72])
	kd.Header.BugCheckParameter3 = binary.LittleEndian.Uint64(data[72:80])
	kd.Header.BugCheckParameter4 = binary.LittleEndian.Uint64(data[80:88])
	kd.Header.KdDebuggerDataBlock = binary.LittleEndian.Uint64(data[96:104])
	kd.Header.PhysicalMemoryBlock = binary.LittleEndian.Uint64(data[104:112])
	kd.Header.PhysicalMemoryBlock64 = binary.LittleEndian.Uint64(data[112:120])
	kd.Header.ContextRecord = binary.LittleEndian.Uint64(data[120:128])
	kd.Header.DumpType = binary.LittleEndian.Uint32(data[128:132])
	kd.Header.SystemTime = binary.LittleEndian.Uint64(data[136:144])
	kd.Header.FeatureFlags = binary.LittleEndian.Uint64(data[152:160])

	sig := string(kd.Header.Signature[:])
	if sig != SignaturePagedu64 && sig != SignatureFull {
		file.Close()
		return nil, fmt.Errorf("invalid signature: %s", sig)
	}

	return kd, nil
}

func (kd *KernelDump) Close() error {
	if kd.File != nil {
		return kd.File.Close()
	}
	return nil
}

func (kd *KernelDump) String() string {
	return fmt.Sprintf("KernelDump:\n  Signature: %s\n  BugCheck: 0x%08X\n  Parameters: 0x%016X 0x%016X 0x%016X 0x%016X\n  DumpType: %d",
		string(kd.Header.Signature[:]),
		kd.Header.BugCheckCode,
		kd.Header.BugCheckParameter1,
		kd.Header.BugCheckParameter2,
		kd.Header.BugCheckParameter3,
		kd.Header.BugCheckParameter4,
		kd.Header.DumpType,
	)
}

func (kd *KernelDump) GetBugCheckName() string {
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

	if name, ok := codes[kd.Header.BugCheckCode]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_0x%08X", kd.Header.BugCheckCode)
}
