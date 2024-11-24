package sdk

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/stream"

	"github.com/ddkwork/golibrary/assert"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ebitengine/purego"
)

func TestGenInterface(t *testing.T) {
	t.Skip()
	g := stream.NewGeneratedFile()
	index := 0
	g.P(`package sdk`)
	g.P(`import "unsafe"`)
	g.P(`type API interface {`)
	for _, s := range stream.ToLines("sdk.go") {
		if strings.HasPrefix(s, "func ") {
			s = strings.TrimPrefix(s, "func ")
			if strings.Contains(s, "{") && strings.Contains(s, "}") {
				before, _, found := strings.Cut(s, "{")
				if found {
					s = before
				}
			}
			s = strings.TrimSuffix(s, "{")
			if strings.Contains(s, "Anon") {
				continue
			}
			index++
			g.P(s)
		}
	}
	g.P(`}`)
	stream.WriteGoFile("api.go", g.Bytes())
}

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
	callback := purego.NewCallback(func(text *byte) int {
		fmt.Println("Received data:", BytePointerToString(text))
		return 0
	})
	SetTextMessageCallback(unsafe.Pointer(callback))
	RunCommandEx("help !monitor")

	mylog.Call(func() {
		assert.True(t, VmxSupportDetection())
		assert.True(t, SetCustomDriverPathEx(TargetFilePath))
		mylog.Trace("InstallVmmDriver", InstallVmmDriver())

		// ConnectLocalDebugger()
		// assert.True(t, Boolean2Bool(ConnectCurrentDebuggerUsingComPort(StringToBytePointer("127.0.0.1"), 8080)))
		assert.True(t, Boolean2Bool(ConnectRemoteDebuggerUsingNamedPipe(StringToBytePointer("\\\\.\\pipe\\HyperDbgPipe"), Boolean(True))))
		assert.True(t, Boolean2Bool(StartProcess(&[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0])))

		mylog.Trace("LoadVmm", LoadVmm())
		// assert.True(t, Boolean2Bool(ConnectRemoteDebuggerUsingNamedPipe(StringToBytePointer("\\\\.\\pipe\\HyperDbgDebug"))))
		// assert.True(t, Boolean2Bool(StartProcessWithArgs(&[]rune("path")[0], &[]rune("C:\\Windows\\SysWOW64\\notepad.exe")[0])))

		TargetFilePath = "../testdata/asm.exe"
		RunCommandEx("bc") //BreakpointsRemoveAll
		//get function address from symbol name
		//PrintSymbol() //? then get function address?
		//$+19     | E9 4E060000              | jmp <asm._asm1>                         |
		//$+A7     | E8 99030000              | call <asm.__allmul>                     | main.c:69
		//$+10B    | E8 70030000              | call <asm.__alldiv>                     | main.c:101
		//$+50B    | 68 34A03800              | push asm.38A034                         | 38A034:"asm1 for code3"

		//SetBreakpoint(_asm1)
		//SetBreakpoint(__allmul)
		//SetBreakpoint(__alldiv)
		//SetBreakpoint(endTrace)

		//bigNumTrace := func() {
		//	switch a.Eip() {
		//	case allmul:
		//		mul(a.PeekStack(1), a.PeekStack(2), a.PeekStack(3), a.PeekStack(4))
		//	case alldiv:
		//		div(a.PeekStack(1), a.PeekStack(2), a.PeekStack(3), a.PeekStack(4))
		//	case end: // 如何把这个条件传递到RunCommandWithCount的停止条件呢？这样就不会让程序退出了
		//	}
		//}

		for range 20 {
			RunCommandEx("start")
			//RunCommandEx("run") //todo add callback for data trace
		}

		//ReadAllRegisters()
		//ReadMemory()
		//ReadTargetRegister(REGISTER_RAX)

		// todo: interface signature need clan up

		//this api`s arg is too more
		//	SetBreakpoint(address Uint64, pid Uint32, tid Uint32, core_numer Uint32)

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

func Test_x32dbg_GetMainModuleEntry(t *testing.T) {
	//stream.RunCommand("netstat -ano | findstr :6589")
	//a := New().Connect()
	//defer a.Close()
	//// a.Restart()//restartadmin todo 重载后服务端被关闭了，还是不方便使用,解决办法是在gui命令栏内输入 restartadmin 重启服务端，这样不用关闭调试器
	//// breakpoints := a.Breakpoints()
	//// mylog.Struct("breakpoints", breakpoints)
	//a.BreakpointsRemoveAll()
	//
	////a.SetBreakpoint(a.GetMainModuleEntry())
	////$+19     | E9 4E060000              | jmp <asm._asm1>                         |
	////$+A7     | E8 99030000              | call <asm.__allmul>                     | main.c:69
	////$+10B    | E8 70030000              | call <asm.__alldiv>                     | main.c:101
	////$+50B    | 68 34A03800              | push asm.38A034                         | 38A034:"asm1 for code3"
	//singular := a.DisassemblySingular(a.MainModuleEntry() + 0x19)
	//asm1 := getAddressBySingularDisassembly[uint32](singular.Assembly)
	//mylog.Hex("asm1", asm1)
	//end := asm1 + 0x50B
	//a.SetBreakpoint(end)
	//a.SetCommentNotes(end, "end trace")
	//
	//singular = a.DisassemblySingular(asm1 + 0xa7)
	//allmul := getAddressBySingularDisassembly[uint32](singular.Assembly)
	//mylog.Hex("__allmul", allmul)
	//a.SetCommentNotes(allmul, "__allmul")
	//
	//singular = a.DisassemblySingular(asm1 + 0x10b)
	//alldiv := getAddressBySingularDisassembly[uint32](singular.Assembly)
	//mylog.Hex("__alldiv", alldiv)
	//a.SetCommentNotes(alldiv, "__alldiv")
	//
	//a.SetBreakpoint(allmul)
	//a.SetBreakpoint(alldiv)
	//
	//bigNumTrace := func() {
	//	switch a.Eip() {
	//	case allmul:
	//		mul(a.PeekStack(1), a.PeekStack(2), a.PeekStack(3), a.PeekStack(4))
	//	case alldiv:
	//		div(a.PeekStack(1), a.PeekStack(2), a.PeekStack(3), a.PeekStack(4))
	//	case end: // 如何把这个条件传递到RunCommandWithCount的停止条件呢？这样就不会让程序退出了
	//	}
	//}
	//a.RunCommandWithCount("run", 400, bigNumTrace)
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

//func TestCallback(t *testing.T) {
//	lib := mylog.Check2(syscall.LoadDLL("libhyperdbg.dll"))
//	procSetTextMessageCallback := mylog.Check2(lib.FindProc("hyperdbg_u_set_text_message_callback"))
//	procInterpreter := mylog.Check2(lib.FindProc("hyperdbg_u_interpreter"))
//	callback := syscall.NewCallback(func(text *byte) int {
//		fmt.Printf("Test in the handler | ")
//		fmt.Println("Received data:", BytePointerToString(text))
//		return 0
//	})
//	mylog.Check3(procSetTextMessageCallback.Call(callback))
//	text := append([]byte("help !monitor"), 0)
//	mylog.Check3(procInterpreter.Call(uintptr(unsafe.Pointer(&text[0]))))
//	lib.Release()
//}
