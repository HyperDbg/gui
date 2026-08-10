package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
)

func TestDetachPerformance(t *testing.T) {
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

	fmt.Println("=== Phase 4: Detach ===")
	start = time.Now()
	err = dbg.Detach()
	fmt.Printf("Detach took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("Detach failed: %v", err)
	}

	fmt.Println("=== Phase 5: Wait for detach to complete ===")
	start = time.Now()
	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() == 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Printf("Wait for detach took: %v\n", time.Since(start))

	fmt.Printf("=== Process handle after detach: %x ===\n", dbg.GetProcessHandle())
	if dbg.GetProcessHandle() != 0 {
		t.Fatal("Process handle should be 0 after detach")
	}
}
