package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	// t.Skip("github ci windows not support vt-x nested virtualization")
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

/*
=== RUN   TestSdk
2024-06-28 14:42:27     Info ->                            │ virtualization technology is vt-x //github.com/ddkwork/hyperdbgui.VmxSupportDetection+0xbd D:/workspace/workspace/branch/gui/util.go:47
2024-06-28 14:42:27   Struct ->                            │  //github.com/ddkwork/hyperdbgui.VmxSupportDetection+0xf8 D:/workspace/workspace/branch/gui/util.go:48
{
 "Cpu0": {
  "Eax": 27,
  "Ebx": 1970169159,
  "Ecx": 1818588270,
  "Edx": 1231384169
 },
 "Cpu1": {
  "Eax": 526017,
  "Ebx": 1050624,
  "Ecx": 2147154879,
  "Edx": 3219913727
 },
 "Vendor": "GenuineIntel",
 "ProcessorBrandString": "11th Gen Intel(R) Core(TM) i7-1160G7 @ 1.20GHz\u0000\u0000"
}
2024-06-28 14:42:27     Info ->                            │ vmx operation is supported by your processor //github.com/ddkwork/hyperdbgui.VmxSupportDetection+0x135 D:/workspace/workspace/branch/gui/util.go:53
2024-06-28 14:42:27    Trace ->           InstallVmmDriver │ 1 //github.com/ddkwork/hyperdbgui.TestSdk.func1+0x45 D:/workspace/workspace/branch/gui/libhyperdbg_test.go:15
2024-06-28 14:42:27    Trace ->                    LoadVmm │ 1 //github.com/ddkwork/hyperdbgui.TestSdk.func1+0xb1 D:/workspace/workspace/branch/gui/libhyperdbg_test.go:17
2024-06-28 14:42:27    Trace ->                  UnloadVmm │ 1 //github.com/ddkwork/hyperdbgui.TestSdk.func1+0x100 D:/workspace/workspace/branch/gui/libhyperdbg_test.go:29
2024-06-28 14:42:27    Trace ->              StopVmmDriver │ 1 //github.com/ddkwork/hyperdbgui.TestSdk.func1+0x146 D:/workspace/workspace/branch/gui/libhyperdbg_test.go:30
2024-06-28 14:42:27    Trace ->         UninstallVmmDriver │ 1 //github.com/ddkwork/hyperdbgui.TestSdk.func1+0x18c D:/workspace/workspace/branch/gui/libhyperdbg_test.go:31
--- PASS: TestSdk (0.00s)
PASS
err, target file is not loaded
current processor vendor is : GenuineIntel
virtualization technology is vt-x
vmx operation is supported by your processor
err, CreateFile failed (2)
handle of the driver not found, probably the driver is not loaded. Did you use 'load' command?
err, OpenService failed (424)
err, OpenService failed (424)

Process finished with the exit code 0

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
