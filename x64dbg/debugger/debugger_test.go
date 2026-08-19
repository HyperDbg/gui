package debugger_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/disassembly"
)

func TestDebuggerCreateProcess(t *testing.T) {
	dbg := debugger.New()

	exePath := `C:\Windows\System32\notepad.exe`
	err := dbg.CreateProcess(exePath, "")
	if err != nil {
		t.Fatalf("CreateProcess failed: %v", err)
	}

	time.Sleep(1 * time.Second)

	fmt.Printf("DEBUG: processHandle=%x, processId=%d\n", dbg.GetProcessHandle(), dbg.GetProcessId())
	fmt.Printf("DEBUG: baseAddress=%x, entryPoint=%x\n", dbg.GetBaseAddress(), dbg.GetEntryPoint())

	err = dbg.UpdateAllPages()
	if err != nil {
		t.Fatalf("UpdateAllPages failed: %v", err)
	}

	disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
	table := disasm.GetTable()
	children := table.Root().Children
	fmt.Printf("DEBUG: 反汇编表格Root节点有 %d 个子节点\n", len(children))

	if len(children) == 0 {
		t.Error("反汇编表格没有子节点，数据未刷新")
	}

	for i, child := range children {
		if i >= 10 {
			break
		}
		data := table.GetData(child)
		fmt.Printf("  [%d] %x: %s\n", i, data.Address, data.Instruction)
	}

	dbg.TerminateProcess(0)
}
