package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/disassembly"
)

func TestDetailedPerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	totalStart := time.Now()

	t.Log("=== Phase 1: CreateProcess ===")
	start := time.Now()
	err := dbg.CreateProcess(exePath, "")
	t.Logf("CreateProcess took: %v", time.Since(start))
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	t.Log("=== Phase 2: Wait for process handle ===")
	start = time.Now()
	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	t.Logf("Wait for process handle took: %v", time.Since(start))

	if dbg.GetProcessHandle() == 0 {
		t.Fatal("Process handle is still 0 after waiting")
	}

	t.Log("=== Phase 3: UpdateAllPages ===")
	start = time.Now()
	err = dbg.UpdateAllPages()
	t.Logf("UpdateAllPages took: %v", time.Since(start))
	if err != nil {
		t.Fatalf("UpdateAllPages failed: %v", err)
	}

	t.Logf("=== Total time: %v ===", time.Since(totalStart))

	dbg.TerminateProcess(0)
}

func TestUpdateAllPagesDetailed(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	err := dbg.CreateProcess(exePath, "")
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	for i := 0; i < 100; i++ {
		if dbg.GetProcessHandle() != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	if dbg.GetProcessHandle() == 0 {
		t.Fatal("Process handle is still 0")
	}

	t.Log("=== Measuring each component in UpdateAllPages ===")

	t.Log("-- disassembly --")
	start := time.Now()
	disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
	err = disasm.Update()
	t.Logf("  disassembly.Update took: %v", time.Since(start))
	if err != nil {
		t.Logf("  disassembly.Update error: %v", err)
	}

	t.Log("-- registers --")
	start = time.Now()
	err = dbg.GetRegisters().Update()
	t.Logf("  registers.Update took: %v", time.Since(start))

	t.Log("-- threads --")
	start = time.Now()
	err = dbg.GetThreads().Update()
	t.Logf("  threads.Update took: %v", time.Since(start))

	t.Log("-- memory --")
	start = time.Now()
	err = dbg.GetMemory().Update()
	t.Logf("  memory.Update took: %v", time.Since(start))

	t.Log("-- symbols --")
	start = time.Now()
	err = dbg.GetSymbols().Update()
	t.Logf("  symbols.Update took: %v", time.Since(start))

	dbg.TerminateProcess(0)
}

func TestEventLoopPerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	totalStart := time.Now()

	fmt.Println("=== Starting CreateProcess ===")
	start := time.Now()
	err := dbg.CreateProcess(exePath, "")
	fmt.Printf("CreateProcess returned after: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	fmt.Println("=== Waiting for process to be ready ===")
	start = time.Now()
	for i := 0; i < 200; i++ {
		if dbg.GetProcessHandle() != 0 {
			fmt.Printf("Process handle ready after: %v (iteration %d)\n", time.Since(start), i)
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	if dbg.GetProcessHandle() == 0 {
		t.Fatal("Process handle is still 0 after 2 seconds")
	}

	fmt.Println("=== Calling UpdateAllPages ===")
	start = time.Now()
	err = dbg.UpdateAllPages()
	fmt.Printf("UpdateAllPages took: %v\n", time.Since(start))
	if err != nil {
		t.Fatalf("UpdateAllPages failed: %v", err)
	}

	fmt.Printf("=== Total test time: %v ===\n", time.Since(totalStart))

	dbg.TerminateProcess(0)
}
