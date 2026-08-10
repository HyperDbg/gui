package disassembly_test

import (
	"fmt"
	"testing"

	"github.com/ddkwork/x64dbg/debugger/disassembly"
)

func TestDisassemblerUpdate(t *testing.T) {
	d := disassembly.New().(*disassembly.Disassembler)

	exePath := `C:\Windows\System32\notepad.exe`
	d.SetUpdateParams(0, 0, exePath)

	err := d.Update()
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	table := d.GetTable()
	children := table.Root().Children
	fmt.Printf("DEBUG: 表格Root节点有 %d 个子节点\n", len(children))

	if len(children) == 0 {
		t.Error("表格没有子节点，数据未刷新")
	}

	for i, child := range children {
		if i >= 10 {
			break
		}
		data := table.GetData(child)
		fmt.Printf("  [%d] %x: %s\n", i, data.Address, data.Instruction)
	}
}
