package pdbex

import (
	"fmt"
	"testing"
)

func TestBSODAddressAnalysis(t *testing.T) {
	pdbPath := "D:\\ux\\examples\\hypedbg\\doc\\cpp\\HyperDbgUnified\\build\\Release\\hyperkd.pdb"

	fmt.Println("=== BSOD 地址解析测试 ===")
	fmt.Println()

	pdb := NewPDB()
	defer pdb.Close()

	err := pdb.Open(pdbPath)
	if err != nil {
		t.Fatalf("无法打开 PDB 文件: %v", err)
	}

	fmt.Println("✓ 成功打开 PDB 文件")
	fmt.Printf("  机器类型: %s\n", pdb.GetArchitectureString())
	fmt.Println()

	fmt.Println("=== 地址 1: <Unloaded_hyperkd.sys>+0x7cc0 ===")
	offset1 := uint64(0x7cc0)
	fmt.Printf("偏移: 0x%X\n", offset1)

	funcName1, ok := pdb.GetFunctionByOffset(offset1)
	if ok {
		fmt.Printf("✓ 函数名称: %s\n", funcName1)

		fnInfo1, ok := pdb.GetFunctionInfo(funcName1)
		if ok {
			fmt.Printf("  函数地址: 0x%X\n", fnInfo1.Address)
			fmt.Printf("  函数大小: 0x%X\n", fnInfo1.Size)
			if fnInfo1.SourceFile != "" {
				fmt.Printf("  源文件: %s:%d\n", fnInfo1.SourceFile, fnInfo1.LineNumber)
			}
		}
	} else {
		fmt.Println("✗ 无法解析函数名称")
	}
	fmt.Println()

	fmt.Println("=== 地址 2: <Unloaded_hyperkd.sys>+0x3b710 ===")
	offset2 := uint64(0x3b710)
	fmt.Printf("偏移: 0x%X\n", offset2)

	funcName2, ok := pdb.GetFunctionByOffset(offset2)
	if ok {
		fmt.Printf("✓ 函数名称: %s\n", funcName2)

		fnInfo2, ok := pdb.GetFunctionInfo(funcName2)
		if ok {
			fmt.Printf("  函数地址: 0x%X\n", fnInfo2.Address)
			fmt.Printf("  函数大小: 0x%X\n", fnInfo2.Size)
			if fnInfo2.SourceFile != "" {
				fmt.Printf("  源文件: %s:%d\n", fnInfo2.SourceFile, fnInfo2.LineNumber)
			}
		}
	} else {
		fmt.Println("✗ 无法解析函数名称")
	}
	fmt.Println()

	fmt.Println("=== 源代码行号查询 ===")
	fileName1, lineNum1, ok1 := pdb.FindSourceLineByRVA(uint32(offset1))
	if ok1 {
		fmt.Printf("偏移 0x%X: %s:%d\n", offset1, fileName1, lineNum1)
	} else {
		fmt.Printf("偏移 0x%X: 无法找到源代码位置\n", offset1)
	}

	fileName2, lineNum2, ok2 := pdb.FindSourceLineByRVA(uint32(offset2))
	if ok2 {
		fmt.Printf("偏移 0x%X: %s:%d\n", offset2, fileName2, lineNum2)
	} else {
		fmt.Printf("偏移 0x%X: 无法找到源代码位置\n", offset2)
	}
	fmt.Println()

	fmt.Println("=== 完成 ===")
}
