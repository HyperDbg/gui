package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
	"github.com/ddkwork/HyperDbg/walker/pdbex"
)

func TestPDBParsing(t *testing.T) {
	dumpFile := "C:\\Windows\\Minidump\\032326-4562-01.dmp"
	pdbDir := "D:\\笔记本\\p\\ux\\examples\\hypedbg\\bsod"

	fmt.Println("=== 测试 PDB 解析功能 ===")
	fmt.Printf("分析文件: %s\n", filepath.Base(dumpFile))
	fmt.Println()

	kd, err := kerneldump.Parse(dumpFile)
	if err != nil {
		fmt.Printf("错误: 无法解析转储文件: %v\n", err)
		os.Exit(1)
	}
	defer kd.Close()

	fmt.Println("=== 基本信息 ===")
	fmt.Printf("签名: %s\n", kd.Header.Signature)
	fmt.Printf("BugCheck: 0x%08X\n", kd.Header.BugCheckCode)
	fmt.Println()

	fmt.Println("=== BugCheck 分析 ===")
	bugCheck, bugCheckName := kd.DetectBugCheckFromStack()
	fmt.Printf("BugCheck: 0x%08X (%s)\n", bugCheck, bugCheckName)
	fmt.Println()

	fmt.Println("=== 堆栈跟踪 ===")
	frames, err := kd.GetStackTrace()
	if err != nil {
		fmt.Printf("错误: 无法获取堆栈跟踪: %v\n", err)
	} else {
		fmt.Printf("找到 %d 个堆栈帧:\n", len(frames))
		for i, frame := range frames {
			if i >= 10 {
				fmt.Printf("... 还有 %d 个帧\n", len(frames)-10)
				break
			}
			fmt.Printf("  #%d %s\n", i, frame.CallSite)
		}
	}
	fmt.Println()

	fmt.Println("=== 异常信息 ===")
	exception, err := kd.GetExceptionRecord()
	if err != nil {
		fmt.Printf("从 ContextRecord 读取异常记录失败: %v\n", err)

		exceptCode, exceptFound := kd.ExtractExceptionCodeFromStack()
		if exceptFound {
			fmt.Printf("从堆栈提取异常代码: 0x%08X (%s)\n", exceptCode, kd.GetExceptionCodeName(exceptCode))
		} else {
			fmt.Println("从堆栈也未找到异常代码")
		}
	} else {
		fmt.Printf("异常代码: 0x%08X (%s)\n", exception.ExceptionCode, kd.GetExceptionCodeName(exception.ExceptionCode))
		fmt.Printf("异常地址: 0x%016X\n", exception.ExceptionAddress)
	}
	fmt.Println()

	fmt.Println("=== 模块信息 ===")
	modules, err := kd.GetModules()
	if err != nil {
		fmt.Printf("读取模块列表失败: %v\n", err)

		modName, modBase, modSize, modFound := kd.ExtractModuleInfoFromStack()
		if modFound {
			fmt.Printf("从堆栈提取模块信息:\n")
			fmt.Printf("  模块名: %s\n", modName)
			fmt.Printf("  基地址: 0x%016X\n", modBase)
			fmt.Printf("  大小: 0x%X\n", modSize)
		} else {
			fmt.Println("从堆栈也未找到模块信息")
		}
	} else {
		fmt.Printf("找到 %d 个模块:\n", len(modules))
		for i, mod := range modules {
			if i >= 5 {
				fmt.Printf("... 还有 %d 个模块\n", len(modules)-5)
				break
			}
			fmt.Printf("  %s @ 0x%016X (0x%X)\n", mod.GetName(), mod.DllBase, mod.SizeOfImage)
		}
	}
	fmt.Println()

	fmt.Println("=== 自动崩溃地址提取 ===")

	crashAddr, crashSource, err := kd.GetCrashAddress()
	if err != nil {
		fmt.Printf("无法确定崩溃地址: %v\n", err)
		fmt.Println("尝试的方法:")
		fmt.Println("  1. 异常记录")
		fmt.Println("  2. BugCheck参数4")
		fmt.Println("  3. 堆栈帧")
		fmt.Println("  4. 物理内存扫描")
		os.Exit(1)
	}

	fmt.Printf("崩溃地址: 0x%016X (来源: %s)\n", crashAddr, crashSource)
	fmt.Println()

	fmt.Println("=== PDB 符号解析 ===")

	modules, err = kd.GetModules()
	if err == nil {
		foundModule := false
		for _, mod := range modules {
			if mod.ContainsAddress(crashAddr) {
				foundModule = true
				fmt.Printf("崩溃模块: %s @ 0x%016X\n", mod.GetName(), mod.DllBase)
				crashOffset := crashAddr - mod.DllBase
				fmt.Printf("崩溃偏移: 0x%X\n", crashOffset)

				pdbPath := filepath.Join(pdbDir, "hyperkd.pdb")
				if _, err := os.Stat(pdbPath); os.IsNotExist(err) {
					fmt.Printf("PDB 文件不存在: %s\n", pdbPath)
					fmt.Println("请确保 hyperkd.pdb 在符号目录中")
				} else {
					fmt.Printf("PDB 文件: %s\n", pdbPath)

					pdb := pdbex.NewPDB()
					err := pdb.Open(pdbPath)
					if err != nil {
						fmt.Printf("无法打开 PDB 文件: %v\n", err)
					} else {
						defer pdb.Close()

						fmt.Println("成功打开 PDB 文件")

						fmt.Printf("尝试解析偏移: 0x%X\n", crashOffset)

						funcName, ok := pdb.GetFunctionByOffset(crashOffset)
						if ok {
							fmt.Printf("✓ 函数名称: %s\n", funcName)

							fnInfo, ok := pdb.GetFunctionInfo(funcName)
							if ok {
								fmt.Printf("✓ 函数地址: 0x%X\n", fnInfo.Address)
								fmt.Printf("✓ 函数大小: 0x%X\n", fnInfo.Size)
								if fnInfo.SourceFile != "" {
									fmt.Printf("✓ 源文件: %s:%d\n", fnInfo.SourceFile, fnInfo.LineNumber)
								} else {
									fmt.Println("✗ 函数没有源文件信息")
								}
							} else {
								fmt.Println("✗ 无法获取函数详细信息")
							}
						} else {
							fmt.Println("✗ 无法解析函数名称")

							funcs := pdb.GetAllFunctions()
							fmt.Printf("PDB 中共有 %d 个函数\n", len(funcs))
							if len(funcs) > 0 {
								fmt.Println("前几个函数:")
								count := 0
								for name, fn := range funcs {
									if count >= 5 {
										break
									}
									fmt.Printf("  %s @ 0x%X\n", name, fn.Address)
									count++
								}
							}
						}
					}
				}
				break
			}
		}
		if !foundModule {
			fmt.Printf("未找到包含崩溃地址 0x%016X 的模块\n", crashAddr)
		}
	} else {
		fmt.Println("无法读取模块信息，无法确定崩溃偏移")
	}

	fmt.Println()
	fmt.Println("=== 完成 ===")
}
