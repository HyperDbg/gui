package ui_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ddkwork/x64dbg/debugger"
	"github.com/ddkwork/x64dbg/debugger/disassembly"
	"github.com/ddkwork/x64dbg/ui"
)

func TestNotepadUIRefresh(t *testing.T) {
	ui.Run(func(dbg *debugger.Debugger) {
		go func() {
			time.Sleep(2 * time.Second)
			notepadPath := `C:\Windows\System32\notepad.exe`
			fmt.Printf("DEBUG: 自动载入记事本: %s\n", notepadPath)
			err := dbg.CreateProcess(notepadPath, "")
			if err != nil {
				fmt.Printf("CreateProcess failed: %v\n", err)
				return
			}
			fmt.Printf("DEBUG: 记事本进程创建成功\n")

			time.Sleep(2 * time.Second)

			disasm := dbg.GetDisassembly().(*disassembly.Disassembler)
			table := disasm.GetTable()
			children := table.Root().Children
			fmt.Printf("DEBUG: 反汇编表格Root节点有 %d 个子节点\n", len(children))

			if len(children) == 0 {
				fmt.Printf("ERROR: 反汇编表格没有子节点，数据未刷新\n")
			} else {
				for i, child := range children {
					if i >= 5 {
						break
					}
					data := table.GetData(child)
					fmt.Printf("  [%d] %x: %s\n", i, data.Address, data.Instruction)
				}
			}

			dbg.TerminateProcess(0)
		}()
	})
}
