package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

func isGithubCI() bool {
	GOPATH := os.Getenv("GOPATH")
	return strings.Contains(GOPATH, "runneradmin")
}

func stringToBytePointer(s string) *byte {
	bytes := []byte(s)
	ptr := &bytes[0]
	return ptr
}

// run now will bsod
// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	// SetCustomDriverPath(stringToBytePointer("."), stringToBytePointer("hyperkd.sys"))
	if isGithubCI() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	mylog.Call(func() {
		assert.True(t, VmxSupportDetection())

		mylog.Check(os.Chdir("D:\\workspace\\workspace\\branch\\gui"))

		mylog.Trace("InstallVmmDriver", InstallVmmDriver())
		ConnectLocalDebugger()
		mylog.Trace("LoadVmm", LoadVmm())

		// todo:
		// start debugger
		// read memory(todo bug read address buffer for disassembly,see cpu.go LayoutDisassemblyTable ,oep buf seems error)
		// read registers
		// set breakpoints
		// step over/into/out
		// continue
		// read stack
		// stop debugger

		mylog.Trace("UnloadVmm", UnloadVmm())
		mylog.Trace("StopVmmDriver", StopVmmDriver())
		mylog.Trace("UninstallVmmDriver", UninstallVmmDriver())
	})
}

func TestLoadDll(t *testing.T) {
	t.Skip("unknown error")
	dll := syscall.NewLazyDLL("libhyperdbg.dll")
	proc := dll.NewProc("hyperdbg_u_set_custom_driver_path")
	mylog.Check(dll.Load())
	ret, _, callErr := proc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("driver_path"))))
	if callErr != nil && callErr.Error() != "The operation completed successfully." {
		fmt.Printf("Failed to call procedure: %v\n", callErr)
	} else {
		fmt.Printf("Procedure called successfully, return value: %v\n", ret)
	}
}

/*
.connect local
load vmm
.sympath SRV*c:\Symbols*https://msdl.microsoft.com/download/symbol
.sym load
.sym reload
.sym download
unload vmm

first in get stack on debug module
.debug remote namedpipe \\.\pipe\HyperDbgDebug
.debug prepare serial 115200 com1
u nt!IopXxxControlFile l 74f
bp nt!IopXxxControlFile
g
kq l 60
*/
