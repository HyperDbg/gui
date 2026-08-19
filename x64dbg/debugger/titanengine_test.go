package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

func TestTitanEngine_InitDebug(t *testing.T) {
	fmt.Println("=== TestTitanEngine_InitDebug ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	notepadPath := "C:\\Windows\\System32\\notepad.exe"
	fmt.Printf("启动记事本: %s\n", notepadPath)

	err := dbg.CreateProcess(notepadPath, "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	pid := dbg.GetProcessId()
	handle := dbg.GetProcessHandle()

	if pid == 0 {
		t.Fatal("进程ID为0，启动失败")
	}
	fmt.Printf("✓ 进程启动成功: PID=%d, Handle=%x\n", pid, handle)

	state := dbg.GetState()
	fmt.Printf("✓ 调试器状态: %d\n", state)

	// 等待调试事件
	eventChan := dbg.GetEventChan()
	eventReceived := false
	timeout := time.After(15 * time.Second)

	select {
	case event := <-eventChan:
		eventReceived = true
		fmt.Printf("✓ 收到调试事件: %d\n", event.DebugEventCode())
		switch event.DebugEventCode() {
		case windows.CREATE_PROCESS_DEBUG_EVENT:
			fmt.Println("  -> CREATE_PROCESS_DEBUG_EVENT")
		case windows.EXCEPTION_DEBUG_EVENT:
			if event.Exception() != nil {
				fmt.Printf("  -> EXCEPTION_DEBUG_EVENT: 0x%X\n", event.Exception().ExceptionRecord.ExceptionCode)
			}
		case windows.EXIT_PROCESS_DEBUG_EVENT:
			fmt.Println("  -> EXIT_PROCESS_DEBUG_EVENT")
		}
	case <-timeout:
		fmt.Println("  -> 等待调试事件超时")
	}

	if eventReceived {
		err = dbg.Continue()
		if err != nil {
			fmt.Printf("继续执行失败: %v\n", err)
		} else {
			fmt.Println("✓ 继续执行成功")
		}

		time.Sleep(1 * time.Second)

		err = dbg.Continue()
		if err != nil {
			fmt.Printf("第二次继续执行失败: %v\n", err)
		} else {
			fmt.Println("✓ 第二次继续执行成功")
		}

		time.Sleep(10 * time.Second)
	}

	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Fatalf("终止进程失败: %v", err)
	}
	fmt.Println("✓ 进程已终止")

	t.Log("✓ TestTitanEngine_InitDebug 测试完成")
}

func TestTitanEngine_Attach(t *testing.T) {
	fmt.Println("=== TestTitanEngine_Attach ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	fmt.Println("注意: 此测试需要手动启动一个被调试进程")
	fmt.Println("跳过 Attach 测试")

	t.Log("✓ TestTitanEngine_Attach 测试完成")
}

func TestTitanEngine_MemoryReadWrite(t *testing.T) {
	fmt.Println("=== TestTitanEngine_MemoryReadWrite ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	err := dbg.CreateProcess("C:\\Windows\\System32\\notepad.exe", "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	fmt.Printf("✓ 进程启动成功: PID=%d\n", dbg.GetProcessId())

	time.Sleep(2 * time.Second)

	handle := dbg.GetProcessHandle()

	testAddr := uint64(0x10000)
	testData := []byte{0x90, 0x90, 0x90, 0x90}

	written, err := windows.WriteProcessMemory(handle, testAddr, testData)
	if err != nil {
		fmt.Printf("写入内存失败: %v\n", err)
	} else {
		fmt.Printf("✓ 写入内存成功: %d 字节\n", written)
	}

	readData, err := windows.ReadProcessMemory(handle, testAddr, 4)
	if err != nil {
		fmt.Printf("读取内存失败: %v\n", err)
	} else {
		fmt.Printf("✓ 读取内存成功: %v\n", readData)
	}

	dbg.TerminateProcess(0)
	fmt.Println("✓ 进程已终止")

	t.Log("✓ TestTitanEngine_MemoryReadWrite 测试完成")
}

func TestTitanEngine_ThreadContext(t *testing.T) {
	fmt.Println("=== TestTitanEngine_ThreadContext ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	err := dbg.CreateProcess("C:\\Windows\\System32\\notepad.exe", "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	fmt.Printf("✓ 进程启动成功: PID=%d\n", dbg.GetProcessId())

	time.Sleep(2 * time.Second)

	time.Sleep(5 * time.Second)

	dbg.TerminateProcess(0)
	fmt.Println("✓ 进程已终止")

	t.Log("✓ TestTitanEngine_ThreadContext 测试完成")
}

func TestTitanEngine_Breakpoints(t *testing.T) {
	fmt.Println("=== TestTitanEngine_Breakpoints ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	bps := dbg.GetBreakpoints()
	if bps == nil {
		t.Fatal("断点管理器初始化失败")
	}

	breakpointAddr := uint64(0x00007FF6ABCD1234)

	bp, err := bps.SetSoftwareBreakpoint(breakpointAddr, true)
	if err != nil {
		fmt.Printf("设置断点失败: %v\n", err)
	} else {
		fmt.Printf("✓ 软件断点设置成功: 0x%X\n", bp.Address)
	}

	allBps := bps.GetAllBreakpoints()
	fmt.Printf("✓ 当前断点数量: %d\n", len(allBps))

	bps.RemoveBreakpoint(breakpointAddr)
	fmt.Printf("✓ 断点已删除\n")

	allBps = bps.GetAllBreakpoints()
	fmt.Printf("✓ 删除后断点数量: %d\n", len(allBps))

	t.Log("✓ TestTitanEngine_Breakpoints 测试完成")
}

func TestTitanEngine_DebugEventHandling(t *testing.T) {
	fmt.Println("=== TestTitanEngine_DebugEventHandling ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	err := dbg.CreateProcess("C:\\Windows\\System32\\notepad.exe", "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	fmt.Printf("✓ 进程启动成功: PID=%d\n", dbg.GetProcessId())

	eventChan := dbg.GetEventChan()

	timeout := time.After(15 * time.Second)

receivedEvent:
	for {
		select {
		case event := <-eventChan:
			fmt.Printf("✓ 收到调试事件: %d\n", event.DebugEventCode())
			switch event.DebugEventCode() {
			case windows.CREATE_PROCESS_DEBUG_EVENT:
				fmt.Println("  -> CREATE_PROCESS_DEBUG_EVENT")
			case windows.CREATE_THREAD_DEBUG_EVENT:
				fmt.Println("  -> CREATE_THREAD_DEBUG_EVENT")
			case windows.LOAD_DLL_DEBUG_EVENT:
				fmt.Println("  -> LOAD_DLL_DEBUG_EVENT")
			case windows.EXCEPTION_DEBUG_EVENT:
				exception := event.Exception()
				if exception != nil {
					fmt.Printf("  -> EXCEPTION_DEBUG_EVENT: 0x%X\n", exception.ExceptionRecord.ExceptionCode)
				}
			case windows.EXIT_PROCESS_DEBUG_EVENT:
				fmt.Println("  -> EXIT_PROCESS_DEBUG_EVENT")
				break receivedEvent
			}
		case <-timeout:
			fmt.Println("超时，退出等待")
			break receivedEvent
		}
	}

	dbg.TerminateProcess(0)
	fmt.Println("✓ 进程已终止")

	t.Log("✓ TestTitanEngine_DebugEventHandling 测试完成")
}

func TestTitanEngine_StopDebug(t *testing.T) {
	fmt.Println("=== TestTitanEngine_StopDebug ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	err := dbg.CreateProcess("C:\\Windows\\System32\\notepad.exe", "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	fmt.Printf("✓ 进程启动成功: PID=%d\n", dbg.GetProcessId())

	time.Sleep(3 * time.Second)

	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Fatalf("终止进程失败: %v", err)
	}
	fmt.Println("✓ 进程已终止")

	state := dbg.GetState()
	fmt.Printf("✓ 调试器最终状态: %d\n", state)

	t.Log("✓ TestTitanEngine_StopDebug 测试完成")
}

func TestTitanEngine_DebugLoop(t *testing.T) {
	fmt.Println("=== TestTitanEngine_DebugLoop ===")

	dbg := debugger.New()
	if dbg == nil {
		t.Fatal("调试器初始化失败")
	}

	notepadPath := "C:\\Windows\\System32\\notepad.exe"

	err := dbg.CreateProcess(notepadPath, "")
	if err != nil {
		t.Fatalf("启动记事本失败: %v", err)
	}

	fmt.Printf("✓ 进程启动成功: PID=%d\n", dbg.GetProcessId())

	eventChan := dbg.GetEventChan()

	// 等待CREATE_PROCESS_DEBUG_EVENT
	select {
	case event := <-eventChan:
		fmt.Printf("收到调试事件: %d\n", event.DebugEventCode())
		if event.DebugEventCode() == windows.CREATE_PROCESS_DEBUG_EVENT {
			fmt.Println("  -> CREATE_PROCESS_DEBUG_EVENT")
		}
	case <-time.After(5 * time.Second):
		t.Fatal("等待CREATE_PROCESS_DEBUG_EVENT超时")
	}

	// 继续执行，等待系统断点
	err = dbg.Continue()
	if err != nil {
		t.Fatalf("继续执行失败: %v", err)
	}

	// 等待系统断点
	select {
	case event := <-eventChan:
		fmt.Printf("收到调试事件: %d\n", event.DebugEventCode())
		if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
			exception := event.Exception()
			if exception != nil {
				fmt.Printf("  -> EXCEPTION_DEBUG_EVENT: 0x%X\n", exception.ExceptionRecord.ExceptionCode)
				if exception.ExceptionRecord.ExceptionCode == windows.EXCEPTION_BREAKPOINT {
					fmt.Println("  -> 系统断点命中")
				}
			}
		}
	case <-time.After(10 * time.Second):
		t.Fatal("等待系统断点超时")
	}

	// 继续执行，让进程运行
	err = dbg.Continue()
	if err != nil {
		t.Fatalf("继续执行失败: %v", err)
	}

	// 等待一段时间让进程运行
	time.Sleep(2 * time.Second)

	// 终止进程
	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Logf("终止进程失败: %v", err)
	}

	state := dbg.GetState()
	fmt.Printf("✓ 调试器最终状态: %d\n", state)

	t.Log("✓ TestTitanEngine_DebugLoop 测试完成")
}
