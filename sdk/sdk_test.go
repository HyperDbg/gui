package sdk

import (
	"testing"

	"github.com/ddkwork/app/ms/driverTool/driver"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	if isRunningOnGitHubActions() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	mylog.Call(func() {
		assert.True(t, VmxSupportDetection())
		assert.True(t, SetCustomDriverPathEx(SysPath))

		// SetTextMessageCallback()

		mylog.Trace("InstallVmmDriver", InstallVmmDriver())

		// ConnectLocalDebugger()
		// assert.True(t, Boolean2Bool(ConnectCurrentDebuggerUsingComPort(StringToBytePointer("127.0.0.1"), 8080)))
		assert.True(t, Boolean2Bool(StartProcess(&[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0])))
		// assert.True(t, Boolean2Bool(StartProcessWithArgs(&[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0], &[]rune("xxoo")[0])))

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

func Test2(t *testing.T) {
	t.Skip()
	if isRunningOnGitHubActions() {
		mylog.Info("github ci windows not support vt-x nested virtualization,skip test")
		return
	}
	//Dependencies := []string{
	//	"hyperhv.dll",
	//	"hyperlog.dll",
	//	"kdserial.dll",
	//}
	assert.True(t, SetCustomDriverPathEx(SysPath))
	d := driver.NewObject("HyperdbgHypervisorDevice", SysPath)
	// d.SetDependencies(Dependencies)
	d.Load(SysPath)
	d.Unload()
}
