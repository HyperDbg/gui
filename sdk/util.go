package sdk

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"

	"github.com/ddkwork/golibrary/stream"

	"github.com/stretchr/testify/assert"

	"github.com/ddkwork/app/ms/hardwareIndo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream/bitfield"
)

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

func isGithubCI() bool {
	return os.Getenv("GITHUB_ACTIONS") == "true"
}

func StringToBytePointer(s string) *byte {
	bytes := []byte(s)
	ptr := &bytes[0]
	return ptr
}

func SetCustomDriverPathEx(DriverFilePath string) bool {
	return SetCustomDriverPath(StringToBytePointer(DriverFilePath), StringToBytePointer(stream.BaseName(DriverFilePath))) == 0
}

//func InterpreterEx(command string) int32 { //todo decode command return status code as error string
//	code := Interpreter(StringToBytePointer(command))
//}

func AddCurrentDirToPath() {
	currentPath := mylog.Check2(filepath.Abs("."))
	pathEnv := os.Getenv("PATH")
	newPath := strings.Join([]string{currentPath, pathEnv}, string(os.PathListSeparator))
	mylog.Check(os.Setenv("PATH", newPath))
}
