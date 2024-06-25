package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ddkwork/app/ms/driverTool/driver"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/stretchr/testify/assert"
)

func TestSdk(t *testing.T) {
	sysName := "hyperkd.sys"
	d := driver.New()
	d.Load(sysName)
	d.Unload()

	return
	path := filepath.Join("C:\\Windows\\System32\\drivers", sysName)
	mylog.Call(func() {
		mylog.CheckIgnore(os.Remove(path))
		stream.CopyFile(sysName, path)

		// assert.Equal(t, 1, HyperdbgUDetectVmxSupport()) // todo convert to return type as bool
		// fmt.Println(HyperdbgUDetectVmxSupport())
		assert.True(t, VmxSupportDetection())
		mylog.Trace("InstallVmmDriver", InstallVmmDriver()) // not working ? return 1
		ConnectLocalDebugger()                              // not working,need return status string for check error
		mylog.Trace("LoadVmm", LoadVmm())                   // not working

		// time.Sleep(1 * time.Second) // wait for debugger to connect
		// todo:
		// start debugger
		// read memory(? address buffer for disassembly)
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
