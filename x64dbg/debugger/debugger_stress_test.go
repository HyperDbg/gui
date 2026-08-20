package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
)

func TestStressRapidRestart(t *testing.T) {
	dbg := debugger.New()
	exePath := `C:\Windows\System32\notepad.exe`

	for i := range 10 {
		fmt.Printf("\n=== Iteration %d ===\n", i+1)

		start := time.Now()
		err := dbg.CreateProcess(exePath, "")
		if err != nil {
			t.Fatalf("Iteration %d: CreateProcess failed: %v", i+1, err)
		}
		fmt.Printf("CreateProcess took: %v\n", time.Since(start))

		for range 50 {
			if dbg.GetProcessHandle() != 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() == 0 {
			t.Fatalf("Iteration %d: Process handle is 0", i+1)
		}

		start = time.Now()
		err = dbg.TerminateProcess(0)
		if err != nil {
			t.Fatalf("Iteration %d: TerminateProcess failed: %v", i+1, err)
		}
		fmt.Printf("TerminateProcess took: %v\n", time.Since(start))

		for range 50 {
			if dbg.GetProcessHandle() == 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() != 0 {
			t.Fatalf("Iteration %d: Process handle should be 0 after terminate", i+1)
		}
	}
}

func TestStressRapidDetach(t *testing.T) {
	dbg := debugger.New()
	exePath := `C:\Windows\System32\notepad.exe`

	for i := range 10 {
		fmt.Printf("\n=== Iteration %d ===\n", i+1)

		start := time.Now()
		err := dbg.CreateProcess(exePath, "")
		if err != nil {
			t.Fatalf("Iteration %d: CreateProcess failed: %v", i+1, err)
		}
		fmt.Printf("CreateProcess took: %v\n", time.Since(start))

		for range 100 {
			if dbg.GetProcessHandle() != 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() == 0 {
			t.Fatalf("Iteration %d: Process handle is 0", i+1)
		}

		start = time.Now()
		err = dbg.Detach()
		if err != nil {
			t.Fatalf("Iteration %d: Detach failed: %v", i+1, err)
		}
		fmt.Printf("Detach took: %v\n", time.Since(start))

		for range 100 {
			if dbg.GetProcessHandle() == 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() != 0 {
			t.Fatalf("Iteration %d: Process handle should be 0 after detach, got: %x", i+1, dbg.GetProcessHandle())
		}
	}
}

func TestStressMixedOperations(t *testing.T) {
	dbg := debugger.New()
	exePath := `C:\Windows\System32\notepad.exe`

	for i := range 10 {
		fmt.Printf("\n=== Iteration %d ===\n", i+1)

		start := time.Now()
		err := dbg.CreateProcess(exePath, "")
		if err != nil {
			t.Fatalf("Iteration %d: CreateProcess failed: %v", i+1, err)
		}
		fmt.Printf("CreateProcess took: %v\n", time.Since(start))

		for range 50 {
			if dbg.GetProcessHandle() != 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() == 0 {
			t.Fatalf("Iteration %d: Process handle is 0", i+1)
		}

		if i%2 == 0 {
			start = time.Now()
			err = dbg.TerminateProcess(0)
			fmt.Printf("TerminateProcess took: %v\n", time.Since(start))
		} else {
			start = time.Now()
			err = dbg.Detach()
			fmt.Printf("Detach took: %v\n", time.Since(start))
		}

		if err != nil {
			t.Fatalf("Iteration %d: Operation failed: %v", i+1, err)
		}

		for range 50 {
			if dbg.GetProcessHandle() == 0 {
				break
			}
			time.Sleep(10 * time.Millisecond)
		}

		if dbg.GetProcessHandle() != 0 {
			t.Fatalf("Iteration %d: Process handle should be 0", i+1)
		}
	}
}

func TestStressNoWait(t *testing.T) {
	dbg := debugger.New()
	exePath := `C:\Windows\System32\notepad.exe`

	for i := range 5 {
		fmt.Printf("\n=== Iteration %d (no wait) ===\n", i+1)

		start := time.Now()
		err := dbg.CreateProcess(exePath, "")
		if err != nil {
			t.Fatalf("Iteration %d: CreateProcess failed: %v", i+1, err)
		}
		fmt.Printf("CreateProcess took: %v\n", time.Since(start))

		time.Sleep(5 * time.Millisecond)

		start = time.Now()
		err = dbg.TerminateProcess(0)
		if err != nil {
			fmt.Printf("TerminateProcess error (expected): %v\n", err)
		}
		fmt.Printf("TerminateProcess took: %v\n", time.Since(start))

		time.Sleep(20 * time.Millisecond)
	}
}
