package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"

	"github.com/ddkwork/HyperDbg/bsod"
)

func main() {
	mylog.Call(func() {
		run()
	})
}

func run() {
	pdbPath := flag.String("pdb", "", "PDB 文件路径")
	driverName := flag.String("driver", "hyperkd", "驱动名称")
	dumpFile := flag.String("dump", "", "Dump 文件路径 (默认使用最新的 minidump)")
	flag.Parse()

	analyzer := bsod.NewAnalyzer()
	analyzer.SetDriverName(*driverName)

	if *pdbPath != "" {
		analyzer.SetPdbPath(*pdbPath)
	}

	var targetDump string
	if *dumpFile != "" {
		targetDump = *dumpFile
	} else {
		targetDump = analyzer.FindLatestMinidump()
	}

	var report strings.Builder
	report.WriteString("=== BSOD Dump 分析工具 ===\n")
	report.WriteString(fmt.Sprintf("Dump 文件: %s\n", targetDump))
	report.WriteString(fmt.Sprintf("驱动名称: %s\n", *driverName))
	if *pdbPath != "" {
		report.WriteString(fmt.Sprintf("PDB 路径: %s\n", *pdbPath))
	}
	report.WriteString("\n")

	result := analyzer.Analyze(targetDump)

	if *pdbPath != "" {
		driverSys := findDriverFile(*pdbPath, *driverName)
		if driverSys != "" {
			analyzer.ResolveSourceLines(result, driverSys)
			if result.CrashOffset != "" {
				analyzer.LocateCrashFunction(result, driverSys)
			}
		}
	}

	report.WriteString(result.String())

	if result.CrashOffset != "" && *pdbPath != "" {
		driverSys := findDriverFile(*pdbPath, *driverName)
		if driverSys != "" {
			report.WriteString(fmt.Sprintf("驱动文件: %s\n", driverSys))

			if result.CrashFunction != "" {
				report.WriteString("\n========================================\n")
				report.WriteString(fmt.Sprintf("崩溃函数: %s!%s\n", *driverName, result.CrashFunction))
				report.WriteString(fmt.Sprintf("偏移: %s\n", result.CrashOffset))
				report.WriteString("========================================\n")
			}
		} else {
			report.WriteString(fmt.Sprintf("找不到驱动文件 %s.sys\n", *driverName))
		}
	} else if result.IsDriverCrash {
		report.WriteString("\n检测到驱动卸载后回调问题\n")
		report.WriteString("这可能是驱动卸载时未正确清理回调导致的\n")
	}

	fmt.Print(report.String())

	bsodDir := `d:\笔记本\p\ux\examples\hypedbg\bsod`
	if !stream.IsDir(bsodDir) {
		os.MkdirAll(bsodDir, 0o755)
	}
	dumpName := filepath.Base(targetDump)
	outputPath := filepath.Join(bsodDir, dumpName+".txt")
	stream.WriteTruncate(outputPath, report.String())
	fmt.Printf("\n分析结果已保存到: %s\n", outputPath)
}

func findDriverFile(pdbPath, driverName string) string {
	directPath := filepath.Join(pdbPath, driverName+".sys")
	if stream.FileExists(directPath) {
		return directPath
	}

	matches := mylog.Check2(filepath.Glob(filepath.Join(pdbPath, "**", driverName+".sys")))
	if len(matches) > 0 {
		return matches[0]
	}

	parentDir := filepath.Dir(pdbPath)
	matches = mylog.Check2(filepath.Glob(filepath.Join(parentDir, "**", driverName+".sys")))
	if len(matches) > 0 {
		return matches[0]
	}

	return ""
}
