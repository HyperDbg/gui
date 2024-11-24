package sdk

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"github.com/ddkwork/app/ms/hardwareIndo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/bitfield"
)

func VmxSupportDetection() (ok bool) {
	hard := hardwareIndo.New()
	if !hard.CpuInfo.Get() {
		return
	}
	if hard.CpuInfo.Vendor != "GenuineIntel" {
		mylog.Check("this program is not designed to run in a non-VT-x environemnt !")
	}
	mylog.Info("", "virtualization technology is vt-x")
	// mylog.Struct("todo",hard.CpuInfo)
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

func isRunningOnGitHubActions() bool {
	return os.Getenv("GITHUB_ACTIONS") == "true"
}

func StringToBytePointer(s string) *byte {
	bytes := []byte(s)
	ptr := &bytes[0]
	return ptr
}

func BytePointerToString(ptr *byte) string {
	var bytes []byte
	for *ptr != 0 {
		bytes = append(bytes, *ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}
	return string(bytes)
}

func Boolean2Bool(b Boolean) bool {
	return int(b) == True
}

func AddCurrentDirToPath() {
	currentPath := mylog.Check2(filepath.Abs("."))
	pathEnv := os.Getenv("PATH")
	newPath := strings.Join([]string{currentPath, pathEnv}, string(os.PathListSeparator))
	mylog.Check(os.Setenv("PATH", newPath))
}
