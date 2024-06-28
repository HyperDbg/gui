package main

import (
	"os"
	"strings"
	"testing"

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

// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	// SetDllDirectory(".")
	SetCustomDriverPath(stringToBytePointer("."))
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

/*
2024-06-28 20:11:28    Trace ->  --------- title --------- │ ------------------ info ------------------ //runtime.doInit1+0xec C:/Program Files/Go/src/runtime/proc.go:7176
panic: Failed to find hyperdbg_u_set_custom_driver_path procedure in libhyperdbg: The specified procedure could not be found.

goroutine 1 [running, locked to thread]:
github.com/ddkwork/golibrary/mylog.check[...](0x2?)
	D:/workspace/workspace/branch/golibrary/mylog/check.go:216 +0x186
github.com/ddkwork/golibrary/mylog.Check2[...](0x0, {0x7ff7424de2c0, 0xc0002bd230?})
	D:/workspace/workspace/branch/golibrary/mylog/check.go:27 +0x45
github.com/ddkwork/app/bindgen/bindlib.windll.Lookup({0xc0005e1dc8?}, {0x7ff741eafc6c?, 0xc0003ba918?})
	D:/workspace/workspace/app/bindgen/bindlib/lib_win32.go:19 +0x2d
github.com/ddkwork/app/bindgen/bindlib.(*Proc).addrSlow(0xc000353bc0)
	D:/workspace/workspace/app/bindgen/bindlib/proc.go:23 +0x5f
github.com/ddkwork/app/bindgen/bindlib.(*Proc).Addr(...)
	D:/workspace/workspace/app/bindgen/bindlib/proc.go:36
github.com/ddkwork/app/bindgen/bindlib.(*Library).ImportNow(0x7ff741d2db80, {0x7ff741eafc6c, 0x21})
	D:/workspace/workspace/app/bindgen/bindlib/lib.go:80 +0x7f
github.com/ddkwork/hyperdbgui.init.0()
	D:/workspace/workspace/branch/gui/libhyperdbg.go:1541 +0x1d0

Process finished with the exit code 1


*/

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
