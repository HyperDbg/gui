package sdk

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
)

const (
	COMMUNICATION_BUFFER_SIZE     = 0x100
	TCP_END_OF_BUFFER_CHARS_COUNT = 4
)

// go test -run ^\QTestSdk\E$
func TestSdk(t *testing.T) {
	if isRunningOnGitHubActions() {
		mylog.Info("github ci windows machine not support nested vt-x virtualization,skip test")
		return
	}
	mylog.Call(func() {
		assert.True(t, VmxSupportDetection())
		assert.True(t, SetCustomDriverPathEx(SysPath))

		mylog.Call(func() {
			dragHandler = func(msg *Char) {
				if msg == nil {
					println("msg is nil,callback not be called")
				}
				goData := (*[COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT]byte)(unsafe.Pointer(&msg))
				fmt.Println("Received data:", string(goData[:]))
			}

			SetTextMessageCallback(unsafe.Pointer(reflect.ValueOf(dragHandler).Pointer()))
			//sharedBuffer := BytePointerToString((*byte)(buffer))
			//mylog.Info("SetTextMessageCallbackUsingSharedBuffer", sharedBuffer)
			// SetTextMessageCallback(Callback(reflect.ValueOf(LogCallback).Pointer()))
		})
		//go func() {
		//	for {
		//		if len(logBuffer) > 1 {
		//			println(BytePointerToString(&logBuffer[0]))
		//		}
		//	}
		//}()

		mylog.Trace("InstallVmmDriver", InstallVmmDriver())

		// ConnectLocalDebugger()
		// assert.True(t, Boolean2Bool(ConnectCurrentDebuggerUsingComPort(StringToBytePointer("127.0.0.1"), 8080)))
		// assert.True(t, Boolean2Bool(StartProcess(&[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0])))

		mylog.Trace("LoadVmm", LoadVmm())
		// assert.True(t, Boolean2Bool(ConnectRemoteDebuggerUsingNamedPipe(StringToBytePointer("\\\\.\\pipe\\HyperDbgDebug"))))

		// assert.True(t, Boolean2Bool(StartProcessWithArgs(&[]rune("path")[0], &[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0])))

		// todo:
		// start debugger
		// read memory
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
.start path C:\Windows\SysWOW64\notepad.exe
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
