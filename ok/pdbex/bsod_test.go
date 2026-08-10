package pdbex

import (
	"fmt"
	"strings"
	"testing"
)

func TestBSODAddressAnalysis(t *testing.T) {
	pdbPath := `D:\ux\examples\hypedbg\HyperDbgUnified\build\Release\Hyperkd.pdb`

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

	fmt.Println("=== 地址 1: <Unloaded_hyperkd.sys>+0x3b650 ===")
	offset1 := uint64(0x3b650)
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

	fmt.Println("=== 地址 2: <Unloaded_hyperkd.sys>+0x3b650 ===")
	offset2 := uint64(0x3b650)
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

func TestNetdemoBSOD(t *testing.T) {
	pdbPath := "D:\\ux\\examples\\hypedbg\\rust-driver\\examples\\netdemo\\netdemo.pdb"

	fmt.Println("=== netdemo BSOD 地址解析 ===")
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

	fmt.Println("=== 解析 DriverEntry+0x469 ===")
	offset := uint64(0x469)
	fmt.Printf("偏移: 0x%X\n", offset)

	funcName, ok := pdb.GetFunctionByOffset(offset)
	if ok {
		fmt.Printf("✓ 函数名称: %s\n", funcName)

		fnInfo, ok := pdb.GetFunctionInfo(funcName)
		if ok {
			fmt.Printf("  函数地址: 0x%X\n", fnInfo.Address)
			fmt.Printf("  函数大小: 0x%X\n", fnInfo.Size)
			if fnInfo.SourceFile != "" {
				fmt.Printf("  源文件: %s:%d\n", fnInfo.SourceFile, fnInfo.LineNumber)
			}
		}
	} else {
		fmt.Println("✗ 无法解析函数名称")
	}
	fmt.Println()

	fmt.Println("=== 源代码行号查询 ===")
	fileName, lineNum, ok := pdb.FindSourceLineByRVA(uint32(offset))
	if ok {
		fmt.Printf("偏移 0x%X: %s:%d\n", offset, fileName, lineNum)
	} else {
		fmt.Printf("偏移 0x%X: 无法找到源代码位置\n", offset)
	}
	fmt.Println()

	fmt.Println("=== 列出所有函数 ===")
	functions := pdb.GetAllFunctions()
	count := 0
	for name, fn := range functions {
		if count > 50 {
			fmt.Printf("... 还有 %d 个函数\n", len(functions)-count)
			break
		}
		fmt.Printf("  %s (0x%X, size: 0x%X)\n", name, fn.Address, fn.Size)
		count++
	}
	fmt.Println()

	fmt.Println("=== 搜索 driver 相关函数 ===")
	for name, fn := range functions {
		lower := strings.ToLower(name)
		if strings.Contains(lower, "driver") || strings.Contains(lower, "entry") || strings.Contains(lower, "server") || strings.Contains(lower, "thread") {
			fmt.Printf("  %s (0x%X, size: 0x%X)\n", name, fn.Address, fn.Size)
		}
	}
	fmt.Println()

	fmt.Println("=== 搜索偏移 0x469 附近的函数 ===")
	for name, fn := range functions {
		if fn.Address <= 0x469 && fn.Address+uint64(fn.Size) > 0x469 {
			fmt.Printf("✓ 找到包含偏移 0x469 的函数: %s (0x%X-0x%X)\n", name, fn.Address, fn.Address+uint64(fn.Size))
		}
	}
	fmt.Println()

	fmt.Println("=== 地址最小的 10 个函数 ===")
	type fnEntry struct {
		name string
		addr uint64
		size uint32
	}
	var sortedFuncs []fnEntry
	for name, fn := range functions {
		sortedFuncs = append(sortedFuncs, fnEntry{name, fn.Address, fn.Size})
	}
	for i := 0; i < len(sortedFuncs)-1; i++ {
		for j := i + 1; j < len(sortedFuncs); j++ {
			if sortedFuncs[j].addr < sortedFuncs[i].addr {
				sortedFuncs[i], sortedFuncs[j] = sortedFuncs[j], sortedFuncs[i]
			}
		}
	}
	for i := 0; i < 10 && i < len(sortedFuncs); i++ {
		fn := sortedFuncs[i]
		fmt.Printf("  %s (0x%X, size: 0x%X)\n", fn.name, fn.addr, fn.size)
	}
	fmt.Println()

	fmt.Println("=== 检查 0x469 是否在代码段开头 ===")
	fmt.Printf("偏移 0x469 = %d 字节\n", 0x469)
	fmt.Println("这可能是 DriverEntry 函数的开头部分，但 Rust 编译器在 release 模式下可能没有生成符号")
	fmt.Println()

	fmt.Println("=== 完成 ===")
}
