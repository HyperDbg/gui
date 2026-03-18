package hardware

import (
	"encoding/binary"
	"fmt"
	"log/slog"
	"syscall"
	"unsafe"
)

type CpuidResult struct {
	EAX uint32
	EBX uint32
	ECX uint32
	EDX uint32
}

var (
	modkernel32                   = syscall.NewLazyDLL("kernel32.dll")
	procIsProcessorFeaturePresent = modkernel32.NewProc("IsProcessorFeaturePresent")
	procGetNativeSystemInfo       = modkernel32.NewProc("GetNativeSystemInfo")
)

type SYSTEM_INFO struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
}

func GetNativeSystemInfo() *SYSTEM_INFO {
	var sysInfo SYSTEM_INFO
	procGetNativeSystemInfo.Call(uintptr(unsafe.Pointer(&sysInfo)))
	return &sysInfo
}

func cpuid(eaxArg uint32, ecxArg uint32) (eax, ebx, ecx, edx uint32)

func init() {
	syscall.NewLazyDLL("kernel32.dll").NewProc("IsProcessorFeaturePresent").Call(0)
}

func Cpuid(eax uint32, ecx uint32) CpuidResult {
	eax, ebx, ecx, edx := cpuid(eax, ecx)
	return CpuidResult{
		EAX: eax,
		EBX: ebx,
		ECX: ecx,
		EDX: edx,
	}
}

func CpuReadVendorString() string {
	result := Cpuid(0, 0)

	ebx := make([]byte, 4)
	edx := make([]byte, 4)
	ecx := make([]byte, 4)

	binary.LittleEndian.PutUint32(ebx, result.EBX)
	binary.LittleEndian.PutUint32(edx, result.EDX)
	binary.LittleEndian.PutUint32(ecx, result.ECX)

	vendor := string(ebx) + string(edx) + string(ecx)

	return vendor
}

func VmxSupportDetection() bool {
	result := Cpuid(1, 0)

	vmxSupported := (result.ECX & (1 << 5)) != 0

	if vmxSupported {
		slog.Info("VMX operation is supported by your processor")
	} else {
		slog.Error("VMX operation is not supported by your processor")
	}

	return vmxSupported
}

func CheckCpuVendor() error {
	vendor := CpuReadVendorString()
	slog.Info("current processor vendor is", "vendor", vendor)

	if vendor == "GenuineIntel" {
		slog.Info("virtualization technology is vt-x")
		return nil
	} else if vendor == "AuthenticAMD" {
		slog.Info("virtualization technology is AMD-V")
		return nil
	} else {
		return fmt.Errorf("this program is not designed to run in a non-VT-x environment")
	}
}

func IsProcessorFeaturePresent(feature uint32) bool {
	ret, _, _ := procIsProcessorFeaturePresent.Call(uintptr(feature))
	return ret != 0
}
