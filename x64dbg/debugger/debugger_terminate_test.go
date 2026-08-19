package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
)

func TestTerminatePerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	fmt.Println("=== Phase 1: CreateProcess ===")
	start := time.Now()
	err := dbg.CreateProcess(exePath, "")
	fmt.Printf("CreateProcess took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	fmt.Println("=== Phase 2: Wait for process handle ===")
	start = time.Now()
	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("Wait for process handle took: %v\n", time.Since(start))

	if dbg.GetProcessHandle() == 0 {
		t.Fatal("Process handle is still 0 after waiting")
	}

	fmt.Println("=== Phase 3: UpdateAllPages ===")
	start = time.Now()
	err = dbg.UpdateAllPages()
	fmt.Printf("UpdateAllPages took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("UpdateAllPages failed: %v", err)
	}

	fmt.Println("=== Phase 4: TerminateProcess ===")
	start = time.Now()
	err = dbg.TerminateProcess(0)
	fmt.Printf("TerminateProcess took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("TerminateProcess failed: %v", err)
	}

	fmt.Printf("=== Process handle after terminate: %x ===\n", dbg.GetProcessHandle())
}

func TestRestartPerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	fmt.Println("=== First CreateProcess ===")
	start := time.Now()
	err := dbg.CreateProcess(exePath, "")
	fmt.Printf("First CreateProcess took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println("=== TerminateProcess ===")
	start = time.Now()
	err = dbg.TerminateProcess(0)
	fmt.Printf("TerminateProcess took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("TerminateProcess failed: %v", err)
	}

	fmt.Println("=== Second CreateProcess (Restart) ===")
	start = time.Now()
	err = dbg.CreateProcess(exePath, "")
	fmt.Printf("Second CreateProcess took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("Second CreateProcess failed: %v", err)
	}

	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	dbg.TerminateProcess(0)
}
