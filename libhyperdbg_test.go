package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"

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
	absPath := mylog.Check2(filepath.Abs("hyperkd.sys"))
	SetCustomDriverPath(stringToBytePointer(absPath), stringToBytePointer(stream.BaseName(absPath)))
	if isGithubCI() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	mylog.Call(func() {
		assert.True(t, VmxSupportDetection())

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

func Test2(t *testing.T) {
	//Dependencies := []string{
	//	"C:\\Windows\\System32\\drivers\\hyperhv.dll",
	//	"C:\\Windows\\system32\\drivers\\hyperlog.dll",
	//	"C:\\Windows\\system32\\drivers\\kdserial.dll",
	//}
	//d := driver.NewObject("HyperdbgHypervisorDevice", "C:\\Windows\\System32\\drivers\\hyperkd.sys")
	//d.SetDependencies(Dependencies)
	//d.Load("C:\\Windows\\System32\\drivers\\hyperkd.sys")
	//d.Unload()
	//return
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
