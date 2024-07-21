package sdk

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/stretchr/testify/assert"
	"reflect"
	"syscall"
	"testing"
	"unsafe"
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
		pfn := func(msg *byte) { // msg is nil,it from TempMessage
			fmt.Println("Received data:", BytePointerToString(msg))
		}
		SetTextMessageCallback(unsafe.Pointer(reflect.ValueOf(pfn).Pointer()))
		assert.True(t, VmxSupportDetection())
		assert.True(t, SetCustomDriverPathEx(SysPath))
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

func TestCallback(t *testing.T) {
	lib, err := syscall.LoadDLL("libhyperdbg.dll")
	if err != nil {
		fmt.Printf("Failed to load libhyperdbg.dll: %v\n", err)
		return
	}

	procSetTextMessageCallback, err := lib.FindProc("hyperdbg_u_set_text_message_callback")
	if err != nil {
		fmt.Printf("Failed to find hyperdbg_u_set_text_message_callback: %v\n", err)
		lib.Release()
		return
	}

	procInterpreter, err := lib.FindProc("hyperdbg_u_interpreter")
	if err != nil {
		fmt.Printf("Failed to find hyperdbg_u_interpreter: %v\n", err)
		lib.Release()
		return
	}

	callback := syscall.NewCallback(func(text *byte) int {
		fmt.Printf("Test in the handler | ")
		fmt.Println("Received data:", BytePointerToString(text))
		return 0
	})

	_, _, _ = procSetTextMessageCallback.Call(callback)

	text := append([]byte("help !monitor"), 0)
	_, _, _ = procInterpreter.Call(uintptr(unsafe.Pointer(&text[0])))
	lib.Release()
}
