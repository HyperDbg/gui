package debugger_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/windows"
)

func TestSingleStepExecution(t *testing.T) {
	fmt.Println("=== TestSingleStepExecution ===")

	// 使用context设置超时，确保测试在20秒内完成
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

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

	// 等待第一次异常（系统断点）
	eventChan := dbg.GetEventChan()
	if eventChan == nil {
		t.Fatal("eventChan is nil")
	}
	fmt.Printf("✓ eventChan is not nil\n")

	eventReceived := false
	timeout := time.After(5 * time.Second)

	select {
	case event := <-eventChan:
		eventReceived = true
		fmt.Printf("✓ 收到调试事件: %d\n", event.DebugEventCode())
		fmt.Println("✓ 准备执行单步")
	case <-timeout:
		t.Fatal("等待调试事件超时")
	case <-ctx.Done():
		t.Fatal("测试超时（20秒）")
	}

	if !eventReceived {
		t.Fatal("没有收到调试事件")
	}

	// 执行5次单步步过
	for i := 1; i <= 5; i++ {
		fmt.Printf("\n=== 执行第 %d 次单步 ===\n", i)

		// 执行单步
		err = dbg.StepInto()
		if err != nil {
			t.Fatalf("执行单步失败: %v", err)
		}
		fmt.Printf("✓ 第 %d 次单步执行成功\n", i)

		// 等待单步异常
		timeout = time.After(10 * time.Second)
		for {
			select {
			case event := <-eventChan:
				fmt.Printf("✓ 收到调试事件: %d\n", event.DebugEventCode())
				if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
					if event.Exception() != nil {
						fmt.Printf("  -> EXCEPTION_DEBUG_EVENT: 0x%X\n", event.Exception().ExceptionRecord.ExceptionCode)
						if event.Exception().ExceptionRecord.ExceptionCode == windows.EXCEPTION_SINGLE_STEP {
							fmt.Printf("  -> 单步异常，第 %d 次单步成功\n", i)
							break
						} else if event.Exception().ExceptionRecord.ExceptionCode == windows.EXCEPTION_BREAKPOINT {
							// 断点异常，调试器会自动处理 pendingStepInto 和 skipBreakpointOnce
							// 继续等待单步异常
							fmt.Printf("  -> 断点异常，继续等待单步异常\n")
							timeout = time.After(10 * time.Second)
							continue
						} else {
							// 其他异常
							fmt.Printf("  -> 其他异常(0x%X)，继续等待\n", event.Exception().ExceptionRecord.ExceptionCode)
							timeout = time.After(3 * time.Second)
							continue
						}
					}
				} else {
					// 跳过非异常事件，继续等待
					fmt.Printf("  -> 跳过非异常事件: %d，继续等待\n", event.DebugEventCode())
					timeout = time.After(3 * time.Second)
					continue
				}
			case <-timeout:
				t.Fatalf("等待单步异常超时，第 %d 次单步失败", i)
			case <-ctx.Done():
				t.Fatalf("测试超时（20秒），第 %d 次单步失败", i)
			}
			break
		}
	}

	// 终止进程
	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Fatalf("终止进程失败: %v", err)
	}
	fmt.Println("\n✓ 进程已终止")

	t.Log("✓ TestSingleStepExecution 测试完成")
}

func TestStepOverExecution(t *testing.T) {
	fmt.Println("=== TestStepOverExecution ===")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

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

	eventChan := dbg.GetEventChan()

	waitForBreakpoint := func(timeoutDuration time.Duration, description string) (*windows.DebugEvent, error) {
		timeout := time.After(timeoutDuration)
		for {
			select {
			case event := <-eventChan:
				fmt.Printf("  收到事件: %d\n", event.DebugEventCode())
				if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
					if event.Exception() != nil {
						code := event.Exception().ExceptionRecord.ExceptionCode
						fmt.Printf("  异常代码: 0x%X\n", code)
						if code == windows.EXCEPTION_BREAKPOINT {
							return event, nil
						}
					}
				}
			case <-timeout:
				return nil, fmt.Errorf("等待%s超时", description)
			case <-ctx.Done():
				return nil, fmt.Errorf("测试超时")
			}
		}
	}

	fmt.Println("\n=== 继续执行，等待系统断点 ===")
	err = dbg.Continue()
	if err != nil {
		t.Fatalf("继续执行失败: %v", err)
	}
	fmt.Println("✓ 已发送继续执行命令")

	fmt.Println("\n=== 等待系统断点 ===")
	event, err := waitForBreakpoint(10*time.Second, "系统断点")
	if err != nil {
		t.Fatalf("等待系统断点失败: %v", err)
	}

	addr := event.Exception().ExceptionRecord.ExceptionAddress
	fmt.Printf("✓ 系统断点命中: 地址=0x%X\n", addr)

	waitForStepOrBreakpoint := func(timeoutDuration time.Duration, description string) (*windows.DebugEvent, error) {
		timeout := time.After(timeoutDuration)
		for {
			select {
			case event := <-eventChan:
				fmt.Printf("  收到事件: %d\n", event.DebugEventCode())
				if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
					if event.Exception() != nil {
						code := event.Exception().ExceptionRecord.ExceptionCode
						fmt.Printf("  异常代码: 0x%X\n", code)
						if code == windows.EXCEPTION_BREAKPOINT || code == windows.EXCEPTION_SINGLE_STEP {
							return event, nil
						}
					}
				}
			case <-timeout:
				return nil, fmt.Errorf("等待%s超时", description)
			case <-ctx.Done():
				return nil, fmt.Errorf("测试超时")
			}
		}
	}

	// 执行5次步过
	for i := 1; i <= 5; i++ {
		fmt.Printf("\n=== 执行第 %d 次步过 ===\n", i)

		// 执行步过
		err = dbg.StepOver()
		if err != nil {
			t.Fatalf("执行步过失败: %v", err)
		}
		fmt.Printf("✓ 第 %d 次步过执行成功\n", i)

		// 等待单步异常或断点异常
		event, err := waitForStepOrBreakpoint(3*time.Second, fmt.Sprintf("第%d次步过", i))
		if err != nil {
			t.Fatalf("等待异常失败: %v", err)
		}

		code := event.Exception().ExceptionRecord.ExceptionCode
		if code == windows.EXCEPTION_SINGLE_STEP {
			fmt.Printf("  -> 单步异常，第 %d 次步过成功\n", i)
		} else if code == windows.EXCEPTION_BREAKPOINT {
			fmt.Printf("  -> 断点异常，第 %d 次步过成功\n", i)
		}
	}

	// 终止进程
	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Fatalf("终止进程失败: %v", err)
	}
	fmt.Println("\n✓ 进程已终止")

	t.Log("✓ TestStepOverExecution 测试完成")
}

func TestSingleStepAtOEP(t *testing.T) {
	fmt.Println("=== TestSingleStepAtOEP ===")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

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
	if pid == 0 {
		t.Fatal("进程ID为0，启动失败")
	}
	fmt.Printf("✓ 进程启动成功: PID=%d\n", pid)

	eventChan := dbg.GetEventChan()
	if eventChan == nil {
		t.Fatal("eventChan is nil")
	}

	entryPoint := dbg.GetEntryPoint()
	baseAddress := dbg.GetBaseAddress()
	fmt.Printf("✓ EntryPoint: 0x%X, BaseAddress: 0x%X\n", entryPoint, baseAddress)

	if entryPoint == 0 {
		t.Fatal("入口点为0，无法设置断点")
	}

	oep := entryPoint
	fmt.Printf("✓ OEP: 0x%X\n", oep)

	fmt.Println("\n=== 等待调试器进入暂停状态 ===")
	for range 100 {
		time.Sleep(50 * time.Millisecond)
		state := dbg.GetState()
		if state == debugger.StatePaused {
			fmt.Printf("✓ 调试器已进入暂停状态\n")
			break
		}
	}

	state := dbg.GetState()
	fmt.Printf("当前调试器状态: %d\n", state)

	fmt.Println("\n=== 继续执行，等待系统断点 ===")
	err = dbg.Continue()
	if err != nil {
		t.Fatalf("继续执行失败: %v", err)
	}
	fmt.Println("✓ 已发送继续执行命令")

	waitForBreakpoint := func(timeoutDuration time.Duration, description string) (*windows.DebugEvent, error) {
		timeout := time.After(timeoutDuration)
		for {
			select {
			case event := <-eventChan:
				fmt.Printf("  收到事件: %d\n", event.DebugEventCode())
				if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
					if event.Exception() != nil {
						code := event.Exception().ExceptionRecord.ExceptionCode
						fmt.Printf("  异常代码: 0x%X\n", code)
						if code == windows.EXCEPTION_BREAKPOINT {
							return event, nil
						}
					}
				}
			case <-timeout:
				return nil, fmt.Errorf("等待%s超时", description)
			case <-ctx.Done():
				return nil, fmt.Errorf("测试超时")
			}
		}
	}

	fmt.Println("\n=== 等待系统断点 ===")
	event, err := waitForBreakpoint(10*time.Second, "系统断点")
	if err != nil {
		t.Fatalf("等待系统断点失败: %v", err)
	}

	addr := event.Exception().ExceptionRecord.ExceptionAddress
	fmt.Printf("✓ 系统断点命中: 地址=0x%X\n", addr)

	waitForBreakpointOrSingleStep := func(timeoutDuration time.Duration, description string) (*windows.DebugEvent, error) {
		timeout := time.After(timeoutDuration)
		for {
			select {
			case event := <-eventChan:
				fmt.Printf("  收到事件: %d\n", event.DebugEventCode())
				if event.DebugEventCode() == windows.EXCEPTION_DEBUG_EVENT {
					if event.Exception() != nil {
						code := event.Exception().ExceptionRecord.ExceptionCode
						fmt.Printf("  异常代码: 0x%X\n", code)
						if code == windows.EXCEPTION_BREAKPOINT || code == windows.EXCEPTION_SINGLE_STEP {
							return event, nil
						}
					}
				}
			case <-timeout:
				return nil, fmt.Errorf("等待%s超时", description)
			case <-ctx.Done():
				return nil, fmt.Errorf("测试超时")
			}
		}
	}

	fmt.Println("\n=== 开始执行8次单步 ===")
	for i := 1; i <= 8; i++ {
		fmt.Printf("\n--- 第 %d 次单步 ---\n", i)

		err = dbg.StepInto()
		if err != nil {
			t.Fatalf("第%d次单步执行失败: %v", i, err)
		}
		fmt.Printf("✓ 第%d次单步命令发送成功\n", i)

		event, err := waitForBreakpointOrSingleStep(5*time.Second, "单步异常")
		if err != nil {
			t.Fatalf("第%d次单步等待异常失败: %v", i, err)
		}

		if event.Exception() != nil {
			code := event.Exception().ExceptionRecord.ExceptionCode
			if code == windows.EXCEPTION_SINGLE_STEP {
				fmt.Printf("✓ 第%d次单步成功 (EXCEPTION_SINGLE_STEP)\n", i)
			} else if code == windows.EXCEPTION_BREAKPOINT {
				fmt.Printf("✓ 第%d次单步成功 (遇到断点)\n", i)
			} else {
				fmt.Printf("✓ 第%d次单步成功 (异常代码: 0x%X)\n", i, code)
			}
		}

		state := dbg.GetState()
		fmt.Printf("  调试器状态: %d\n", state)
	}

	fmt.Println("\n=== 8次单步执行完成，停止进程 ===")
	err = dbg.TerminateProcess(0)
	if err != nil {
		t.Fatalf("终止进程失败: %v", err)
	}
	fmt.Println("✓ 进程已终止")

	t.Log("✓ TestSingleStepAtOEP 测试完成")
}
