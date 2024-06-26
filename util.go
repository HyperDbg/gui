package main

import (
	"fmt"
	"syscall"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"github.com/ddkwork/app/ms/hardwareIndo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream/bitfield"
)

func SetDllDirectory(path string) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	setDllDirectory := kernel32.NewProc("SetDllDirectoryW")
	utf16Ptr := mylog.Check2(syscall.UTF16PtrFromString(path))
	mylog.Check3(setDllDirectory.Call(uintptr(unsafe.Pointer(utf16Ptr))))
}

func AddDllDirectory(path string) (uintptr, error) {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	addDllDirectory := kernel32.NewProc("AddDllDirectory")
	utf16Ptr := mylog.Check2(syscall.UTF16PtrFromString(path))

	// 调用 AddDllDirectory 函数
	handle, _ := mylog.Check3(addDllDirectory.Call(uintptr(unsafe.Pointer(utf16Ptr))))

	fmt.Println("Successfully added DLL directory")
	return handle, nil
}

func LOWORD(l uint32) uint16 { return uint16(l) }
func LOBYTE(l uint32) uint8  { return byte(l) }
func HIWORD(l uint32) uint16 { return uint16(l >> 16) }
func HIBYTE(l uint32) uint8  { return byte(l >> 24) }

func TestSizeof(t *testing.T) {
	// assert.Equal(t, 11, binary.Size(DEBUGGER_REMOTE_PACKET{}))
}

func TestHIBYTE(t *testing.T) {
	v := uint32(0x11223344)
	assert.Equal(t, byte(0x11), HIBYTE(v))
	assert.Equal(t, uint16(0x1122), HIWORD(v))
	assert.Equal(t, byte(0x44), LOBYTE(v))
	assert.Equal(t, uint16(0x3344), LOWORD(v))
}

func VmxSupportDetection() (ok bool) {
	hard := hardwareIndo.New()
	if !hard.CpuInfo.Get() {
		return
	}
	if hard.CpuInfo.Vendor != "GenuineIntel" {
		mylog.Check("this program is not designed to run in a non-VT-x environemnt !")
	}
	mylog.Info("", "virtualization technology is vt-x")
	// mylog.Struct(hard.CpuInfo)
	field := bitfield.NewFromUint32(hard.CpuInfo.Cpu1.Ecx)
	if !field.Test(5) {
		mylog.Check("vmx operation is not supported by your processor")
	}
	mylog.Info("", "vmx operation is supported by your processor")
	return true
}

func DeviceName() string { return "HyperdbgHypervisorDevice" }
func LinkName() (*uint16, error) {
	return syscall.UTF16PtrFromString("\\\\\\\\.\\\\" + DeviceName())
}
