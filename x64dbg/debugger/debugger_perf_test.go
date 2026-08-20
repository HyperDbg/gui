package debugger_test

import (
	"testing"
	"time"

	"github.com/ddkwork/ddk/xed"
	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/disassembly"
	"github.com/saferwall/pe"
)

func TestDebuggerLoadPerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	totalStart := time.Now()

	t.Log("=== Phase 1: CreateProcess ===")
	start := time.Now()
	err := dbg.CreateProcess(exePath, "")
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}
	t.Logf("CreateProcess took: %v", time.Since(start))

	t.Log("=== Phase 2: Wait for process to initialize ===")
	start = time.Now()
	for range 50 {
		time.Sleep(100 * time.Millisecond)
		if dbg.GetProcessHandle() != 0 {
			break
		}
	}
	t.Logf("Process initialization took: %v", time.Since(start))

	t.Log("=== Phase 3: UpdateAllPages ===")
	start = time.Now()
	err = dbg.UpdateAllPages()
	if err != nil {
		t.Fatalf("UpdateAllPages failed: %v", err)
	}
	t.Logf("UpdateAllPages took: %v", time.Since(start))

	disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
	table := disasm.GetTable()
	children := table.Root().Children
	t.Logf("反汇编表格Root节点有 %d 个子节点", len(children))

	t.Logf("Total time: %v", time.Since(totalStart))

	dbg.TerminateProcess(0)
}

func TestDisassemblyPerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	err := dbg.CreateProcess(exePath, "")
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	start := time.Now()
	disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
	err = disasm.Update()
	if err != nil {
		t.Fatalf("disasm.Update failed: %v", err)
	}
	t.Logf("disasm.Update took: %v", time.Since(start))

	table := disasm.GetTable()
	children := table.Root().Children
	t.Logf("反汇编表格Root节点有 %d 个子节点", len(children))

	dbg.TerminateProcess(0)
}

func TestPEParsePerformance(t *testing.T) {
	exePath := `C:\Windows\System32\notepad.exe`

	start := time.Now()
	f := xed.ParserPe(exePath)
	t.Logf("ParserPe took: %v", time.Since(start))

	optionalHeader := f.NtHeader.OptionalHeader
	switch o := optionalHeader.(type) {
	case pe.ImageOptionalHeader64:
		t.Logf("ImageBase: %x, EntryPoint: %x", o.ImageBase, o.AddressOfEntryPoint)
	case pe.ImageOptionalHeader32:
		t.Logf("ImageBase: %x, EntryPoint: %x", o.ImageBase, o.AddressOfEntryPoint)
	}
}

func TestDisassemblyDecodePerformance(t *testing.T) {
	buffer := make([]byte, 4096)
	for i := range buffer {
		buffer[i] = byte(i % 256)
	}

	start := time.Now()
	x := xed.New(buffer)
	x.SetBaseAddress(0x140000000)
	x.SetIsFilterModel(false)
	x.Decode64()
	t.Logf("Decode64 took: %v", time.Since(start))
	t.Logf("Decoded %d instructions", len(x.AsmObjects))
}

func TestComponentUpdatePerformance(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`

	err := dbg.CreateProcess(exePath, "")
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	time.Sleep(500 * time.Millisecond)

	components := []struct {
		name string
		fn   func() error
	}{
		{"disassembly", func() error {
			disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
			return disasm.Update()
		}},
		{"registers", func() error {
			return dbg.GetRegisters().Update()
		}},
		{"threads", func() error {
			return dbg.GetThreads().Update()
		}},
		{"memory", func() error {
			return dbg.GetMemory().Update()
		}},
		{"symbols", func() error {
			return dbg.GetSymbols().Update()
		}},
		{"breakpoints", func() error {
			return dbg.GetBreakpoints().Update()
		}},
		{"stack", func() error {
			return dbg.GetStack().Update()
		}},
		{"seh", func() error {
			return dbg.GetSEH().Update()
		}},
		{"trace", func() error {
			return dbg.GetTrace().Update()
		}},
		{"scylla", func() error {
			return dbg.GetScylla().Update()
		}},
		{"labels", func() error {
			return dbg.GetLabels().Update()
		}},
		{"comments", func() error {
			return dbg.GetComments().Update()
		}},
		{"functions", func() error {
			return dbg.GetFunctions().Update()
		}},
		{"xrefs", func() error {
			return dbg.GetXrefs().Update()
		}},
		{"types", func() error {
			return dbg.GetTypes().Update()
		}},
		{"watches", func() error {
			return dbg.GetWatches().Update()
		}},
		{"graphs", func() error {
			return dbg.GetGraphs().Update()
		}},
		{"exceptions", func() error {
			return dbg.GetExceptions().Update()
		}},
		{"bookmarks", func() error {
			return dbg.GetBookmarks().Update()
		}},
		{"loops", func() error {
			return dbg.GetLoops().Update()
		}},
		{"peview", func() error {
			return dbg.GetPeView().Update()
		}},
		{"ark", func() error {
			return dbg.GetArk().Update()
		}},
		{"imm", func() error {
			return dbg.GetImm().Update()
		}},
	}

	for _, comp := range components {
		start := time.Now()
		err := comp.fn()
		elapsed := time.Since(start)
		if err != nil {
			t.Logf("%s: ERROR %v (%v)", comp.name, err, elapsed)
		} else {
			t.Logf("%s: %v", comp.name, elapsed)
		}
	}

	dbg.TerminateProcess(0)
}
