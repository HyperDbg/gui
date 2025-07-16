package testdata

import (
	"github.com/ddkwork/ddk/hardwareinfo"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream/bitfield"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func VmxSupportDetection() (ok bool) {
	hard := hardwareinfo.New()
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

func AddCurrentDirToPath() {
	currentPath := mylog.Check2(filepath.Abs("."))
	pathEnv := os.Getenv("PATH")
	newPath := strings.Join([]string{currentPath, pathEnv}, string(os.PathListSeparator))
	mylog.Check(os.Setenv("PATH", newPath))
}
