package HPRDBGCTRL

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/stretchr/testify/assert"
)

func TestSdk(t *testing.T) {
	sysName := "hprdbgkd.sys"
	path := filepath.Join("C:\\Windows\\System32\\drivers", sysName)
	mylog.Call(func() { // todo Exception 0xc00000fd 堆栈溢出（Stack Overflow
		mylog.Check(os.Remove(path))
		stream.CopyFile(sysName, path)

		assert.Equal(t, 1, HyperdbgUDetectVmxSupport()) // todo convert to return type as bool

		// s := make([]byte, 0, 100)
		// HyperdbgUReadVendorString(s) // todo bug
		// fmt.Println(s)

		HyperdbgUInstallVmmDriver()
		HyperdbgULoadVmm()
		// HyperdbgUstartVmm()//?

		HyperdbgUStopVmmDriver()
		HyperdbgUUnloadVmm()
		HyperdbgUUninstallVmmDriver()
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
