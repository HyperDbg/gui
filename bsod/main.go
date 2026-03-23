package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode"

	"github.com/ddkwork/golibrary/std/mylog"
)

type Analyzer struct {
	WinDbgPath     string
	SymbolPath     string
	SourcePath     string
	DriverName     string
	PdbPath        string
	OutputDir      string
	DriverSys      string
	CommandTimeout time.Duration
	MaxOutputSize  int
	SymOpt         string
	AnalyzeCmds    []string
	ResolveCmds    func(frame *StackFrame) []string
	LocateCmds     func(crashOffset string) []string
}

type AnalysisResult struct {
	DumpFile         string
	DriverName       string
	UnloadedModules  []UnloadedModule
	ExceptionAddress string
	ExceptionCode    string
	ProcessName      string
	ImageName        string
	FailureBucket    string
	CallStack        []StackFrame
	CrashFunction    string
	CrashOffset      string
	IsDriverCrash    bool
	CrashReason      string
	RawOutput        string
	AIAnalysis       string
}

type UnloadedModule struct {
	BaseAddress string
	EndAddress  string
	Name        string
}

type StackFrame struct {
	Index      int
	ChildSP    string
	RetAddr    string
	CallSite   string
	IsDriver   bool
	ModuleName string
	FuncName   string
	FuncOffset string
	SourceFile string
	SourceLine string
}

func NewAnalyzer() *Analyzer {
	pdbPath := `D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release`
	sourcePath := `D:\笔记本\p\ux\examples\hypedbg\cpp\HyperDbgUnified\HyperDbg\hyperdbg`
	driverName := "hyperkd"

	a := &Analyzer{
		WinDbgPath:     `C:\Microsoft.WinDbg_1.2601.12001.0_x64__8wekyb3d8bbwe\amd64\kd.exe`,
		SymbolPath:     pdbPath,
		SourcePath:     sourcePath,
		DriverName:     driverName,
		PdbPath:        pdbPath,
		OutputDir:      `d:\笔记本\p\ux\examples\hypedbg\bsod`,
		CommandTimeout: 5 * time.Second,
		MaxOutputSize:  256 * 1024,
		SymOpt:         ".symopt+0x10",
	}
	a.DriverSys = a.findDriverFile()
	a.initCommands()
	return a
}

func hasChinese(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}

func (a *Analyzer) initCommands() {
	a.AnalyzeCmds = []string{
		".sympath*",
		fmt.Sprintf(".sympath %s", a.SymbolPath),
		".symopt+0x10",
		".reload /f",
		fmt.Sprintf(".srcpath %s", a.SourcePath),
		".exr -1",
		".ecxr",
		"kv 30",
		"!analyze -v",
		"q",
	}
}

func main() {
	mylog.Call(func() {
		NewAnalyzer().Run()
	})
}

func (a *Analyzer) Run() {
	targetDump := a.findLatestMinidump()
	fmt.Printf("=== BSOD Dump 分析工具 ===\n")
	fmt.Printf("Dump 文件: %s\n", targetDump)
	fmt.Printf("驱动名称: %s\n", a.DriverName)
	fmt.Printf("PDB 路径: %s\n", a.PdbPath)
	fmt.Printf("源代码路径: %s\n", a.SourcePath)
	if a.DriverSys != "" {
		fmt.Printf("驱动文件: %s\n", a.DriverSys)
	}

	if hasChinese(a.PdbPath) {
		fmt.Printf("\n[警告] PDB路径包含中文字符，WinDbg可能无法正确解析符号文件\n")
		fmt.Printf("正在复制PDB文件到minidump目录...\n")
		dumpDir := filepath.Dir(targetDump)
		a.copyPdbToDir(dumpDir)
		a.SymbolPath = dumpDir
		fmt.Printf("符号路径已更新为: %s\n", dumpDir)
	}

	fmt.Println()

	result := a.analyze(targetDump)

	report := a.buildReport(result)
	fmt.Print(report)

	outputPath := filepath.Join(a.OutputDir, filepath.Base(targetDump)+".txt")
	os.WriteFile(outputPath, []byte(report), 0o644)
	fmt.Printf("\n分析结果已保存到: %s\n", outputPath)
}

func (a *Analyzer) runKdCommand(ctx context.Context, desc, dumpFile string, commands []string) string {
	fmt.Printf("[%s] %s\n", time.Now().Format("15:04:05"), desc)

	ctx, cancel := context.WithTimeout(ctx, a.CommandTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, a.WinDbgPath,
		"-z", dumpFile,
		"-c", strings.Join(commands, "; "),
	)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("[%s] 创建管道失败: %v\n", time.Now().Format("15:04:05"), err)
		return ""
	}
	cmd.Stderr = cmd.Stdout

	start := time.Now()
	if err := cmd.Start(); err != nil {
		fmt.Printf("[%s] 启动失败: %v\n", time.Now().Format("15:04:05"), err)
		return ""
	}

	buf := make([]byte, 32*1024)
	var output strings.Builder
	output.Grow(64 * 1024)
	reachedLimit := false

	for {
		n, err := stdout.Read(buf)
		if n > 0 {
			if output.Len()+n > a.MaxOutputSize {
				remaining := a.MaxOutputSize - output.Len()
				if remaining > 0 {
					output.Write(buf[:remaining])
				}
				reachedLimit = true
				break
			}
			output.Write(buf[:n])
		}
		if err != nil {
			break
		}
	}

	if reachedLimit {
		cmd.Process.Kill()
		cmd.Wait()
		elapsed := time.Since(start)
		fmt.Printf("[%s] 达到输出上限，已终止进程 耗时: %v 输出: %d 字节\n", time.Now().Format("15:04:05"), elapsed, output.Len())
		return output.String()
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		cmd.Process.Kill()
		<-done
		elapsed := time.Since(start)
		fmt.Printf("[%s] 超时! 耗时: %v\n", time.Now().Format("15:04:05"), elapsed)
		return output.String()
	case err := <-done:
		elapsed := time.Since(start)
		if err != nil {
			fmt.Printf("[%s] 完成(有错误) 耗时: %v 输出: %d 字节\n", time.Now().Format("15:04:05"), elapsed, output.Len())
		} else {
			fmt.Printf("[%s] 完成 耗时: %v 输出: %d 字节\n", time.Now().Format("15:04:05"), elapsed, output.Len())
		}
		return output.String()
	}
}

func (a *Analyzer) analyze(dumpFile string) *AnalysisResult {
	result := &AnalysisResult{
		DumpFile:   dumpFile,
		DriverName: a.DriverName,
	}

	ctx := context.Background()
	fmt.Println("\n[步骤1] 分析 dump 文件...")

	output := a.runKdCommand(ctx, "分析dump", dumpFile, a.AnalyzeCmds)

	result.RawOutput = output

	a.parseExceptionInfo(output, result)
	a.parseUnloadedModules(output, result)
	a.parseCallStack(output, result)
	a.parseAnalyzeOutput(output, result)
	a.analyzeCrashReason(result)

	if result.ExceptionAddress != "" && len(result.UnloadedModules) > 0 {
		a.checkDriverCrash(result)
	}

	stackCommands := a.extractStackCommands(output)
	if len(stackCommands) > 0 {
		fmt.Printf("\n[步骤2] 执行动态命令...\n")
		for _, cmd := range stackCommands {
			fmt.Printf("  执行: %s\n", cmd)
			stackOutput := a.runKdCommand(ctx, "执行动态命令", dumpFile, []string{cmd, "q"})
			result.AIAnalysis += "\n" + stackOutput
		}
	}

	a.extractAIAnalysis(output, result)

	return result
}

func (a *Analyzer) findLatestMinidump() string {
	minidumpDir := `C:\Windows\Minidump`
	entries, _ := os.ReadDir(minidumpDir)

	var dmpFiles []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".dmp") {
			dmpFiles = append(dmpFiles, entry)
		}
	}

	if len(dmpFiles) == 0 {
		panic("找不到 dump 文件")
	}

	sort.Slice(dmpFiles, func(i, j int) bool {
		infoI, _ := dmpFiles[i].Info()
		infoJ, _ := dmpFiles[j].Info()
		return infoI.ModTime().After(infoJ.ModTime())
	})

	return filepath.Join(minidumpDir, dmpFiles[0].Name())
}

func (a *Analyzer) findDriverFile() string {
	directPath := filepath.Join(a.PdbPath, a.DriverName+".sys")
	if _, err := os.Stat(directPath); err == nil {
		return directPath
	}
	if matches, _ := filepath.Glob(filepath.Join(a.PdbPath, "**", a.DriverName+".sys")); len(matches) > 0 {
		return matches[0]
	}
	return ""
}

func (a *Analyzer) copyPdbToDir(targetDir string) {
	pdbFiles, err := filepath.Glob(filepath.Join(a.PdbPath, "*.pdb"))
	if err != nil || len(pdbFiles) == 0 {
		return
	}

	for _, pdbFile := range pdbFiles {
		destPath := filepath.Join(targetDir, filepath.Base(pdbFile))
		if _, err := os.Stat(destPath); err == nil {
			continue
		}

		src, err := os.ReadFile(pdbFile)
		if err != nil {
			continue
		}

		if err := os.WriteFile(destPath, src, 0o644); err == nil {
			fmt.Printf("已复制 PDB: %s -> %s\n", filepath.Base(pdbFile), targetDir)
		}
	}
}

func (a *Analyzer) parseExceptionInfo(output string, result *AnalysisResult) {
	if match := regexp.MustCompile(`ExceptionAddress:\s*([0-9a-fA-F` + "`" + `]+)`).FindStringSubmatch(output); match != nil {
		result.ExceptionAddress = strings.ReplaceAll(match[1], "`", "")
	}
	if match := regexp.MustCompile(`ExceptionCode:\s*([0-9a-fA-F]+)`).FindStringSubmatch(output); match != nil {
		result.ExceptionCode = match[1]
	}
}

func (a *Analyzer) parseUnloadedModules(output string, result *AnalysisResult) {
	unloadedMatch := regexp.MustCompile(`Unloaded modules:\r?\n((?:.*\r?\n)+)`).FindStringSubmatch(output)
	if unloadedMatch == nil {
		return
	}

	moduleRegex := regexp.MustCompile(`([0-9a-fA-F` + "`" + `]+)\s+([0-9a-fA-F` + "`" + `]+)\s+(\S+\.sys)`)
	for _, match := range moduleRegex.FindAllStringSubmatch(unloadedMatch[1], -1) {
		if strings.Contains(strings.ToLower(match[3]), strings.ToLower(a.DriverName)) {
			result.UnloadedModules = append(result.UnloadedModules, UnloadedModule{
				BaseAddress: strings.ReplaceAll(match[1], "`", ""),
				EndAddress:  strings.ReplaceAll(match[2], "`", ""),
				Name:        match[3],
			})
		}
	}
}

func (a *Analyzer) parseCallStack(output string, result *AnalysisResult) {
	startIdx := strings.Index(output, "# Child-SP")
	if startIdx == -1 {
		return
	}

	endIdx := strings.Index(output[startIdx:], "quit:")
	if endIdx == -1 {
		endIdx = len(output) - startIdx
	}

	for line := range strings.SplitSeq(output[startIdx:startIdx+endIdx], "\n") {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 4 {
			continue
		}

		frame := StackFrame{}
		fmt.Sscanf(parts[0], "#%d", &frame.Index)
		frame.ChildSP = strings.ReplaceAll(parts[1], "`", "")
		frame.RetAddr = strings.ReplaceAll(parts[2], "`", "")

		callSiteStart := strings.Index(line, parts[2])
		if callSiteStart != -1 {
			frame.CallSite = strings.TrimSpace(line[callSiteStart+len(parts[2]):])
			a.parseCallSite(frame.CallSite, &frame)
		}

		if strings.Contains(frame.CallSite, "Unloaded_"+a.DriverName) ||
			strings.Contains(frame.CallSite, a.DriverName+"!") {
			frame.IsDriver = true
			result.IsDriverCrash = true
		}

		result.CallStack = append(result.CallStack, frame)
	}
}

func (a *Analyzer) parseCallSite(callSite string, frame *StackFrame) {
	if !strings.Contains(callSite, "!") {
		return
	}

	parts := strings.SplitN(callSite, "!", 2)
	frame.ModuleName = strings.TrimSpace(parts[0])
	if len(parts) > 1 {
		rest := strings.TrimSpace(parts[1])
		if plusIdx := strings.Index(rest, "+"); plusIdx != -1 {
			frame.FuncName = strings.TrimSpace(rest[:plusIdx])
			frame.FuncOffset = rest[plusIdx:]
		} else {
			frame.FuncName = rest
		}
	}
}

func (a *Analyzer) parseAnalyzeOutput(output string, result *AnalysisResult) {
	if match := regexp.MustCompile(`PROCESS_NAME:\s*(\S+)`).FindStringSubmatch(output); match != nil {
		result.ProcessName = match[1]
	}
	if match := regexp.MustCompile(`IMAGE_NAME:\s*(\S+)`).FindStringSubmatch(output); match != nil {
		result.ImageName = match[1]
	}
	if match := regexp.MustCompile(`FAILURE_BUCKET_ID:\s*(\S+)`).FindStringSubmatch(output); match != nil {
		result.FailureBucket = match[1]
	}
}

func (a *Analyzer) extractStackCommands(output string) []string {
	var commands []string

	if match := regexp.MustCompile(`STACK_COMMAND:\s*(.+)`).FindStringSubmatch(output); match != nil {
		cmd := strings.TrimSpace(match[1])
		if cmd != "" {
			commands = append(commands, cmd)
		}
	}

	return commands
}

func (a *Analyzer) extractAIAnalysis(output string, result *AnalysisResult) {
	analyzeIdx := strings.Index(output, "Bugcheck Analysis")
	if analyzeIdx == -1 {
		analyzeIdx = strings.Index(output, "DRIVER_IRQL_NOT_LESS_OR_EQUAL")
		if analyzeIdx == -1 {
			analyzeIdx = strings.Index(output, "PAGE_FAULT_IN_NONPAGED_AREA")
		}
		if analyzeIdx == -1 {
			analyzeIdx = strings.Index(output, "KERNEL_MODE_EXCEPTION_NOT_HANDLED")
		}
		if analyzeIdx == -1 {
			analyzeIdx = strings.Index(output, "SYSTEM_SERVICE_EXCEPTION")
		}
	}

	if analyzeIdx != -1 {
		result.AIAnalysis = output[analyzeIdx:]
	} else {
		result.AIAnalysis = output[len(output)/2:]
	}
}

func (a *Analyzer) analyzeCrashReason(result *AnalysisResult) {
	var reasons []string

	switch result.ExceptionCode {
	case "80000003":
		reasons = append(reasons, "断点异常 - 可能是调试断点触发或代码中的int 3指令")
	case "c0000005":
		reasons = append(reasons, "访问违规 - 内存访问错误")
	case "c000001d":
		reasons = append(reasons, "非法指令")
	case "c0000094":
		reasons = append(reasons, "整数除零")
	case "c0000409":
		reasons = append(reasons, "栈缓冲区溢出")
	case "c0000374":
		reasons = append(reasons, "堆损坏")
	}

	if len(result.UnloadedModules) > 0 {
		reasons = append(reasons, fmt.Sprintf("检测到 %d 个已卸载的 %s.sys 实例", len(result.UnloadedModules), a.DriverName))
		reasons = append(reasons, "驱动被多次加载/卸载，可能存在回调未清理问题")
	}

	if result.FailureBucket != "" && strings.Contains(result.FailureBucket, "WRONG_SYMBOLS") {
		reasons = append(reasons, "符号文件不匹配")
	}

	if len(reasons) > 0 {
		result.CrashReason = strings.Join(reasons, "\n")
	}
}

func (a *Analyzer) checkDriverCrash(result *AnalysisResult) {
	for _, mod := range result.UnloadedModules {
		baseAddr := parseHex(mod.BaseAddress)
		endAddr := parseHex(mod.EndAddress)
		exceptionAddr := parseHex(result.ExceptionAddress)

		if exceptionAddr >= baseAddr && exceptionAddr < endAddr {
			result.CrashOffset = fmt.Sprintf("0x%X", exceptionAddr-baseAddr)
			result.IsDriverCrash = true
			break
		}
	}
}

func (a *Analyzer) buildReport(result *AnalysisResult) string {
	var sb strings.Builder

	sb.WriteString("=== BSOD Dump 分析结果 ===\n")
	sb.WriteString(fmt.Sprintf("Dump 文件: %s\n", result.DumpFile))
	sb.WriteString(fmt.Sprintf("驱动名称: %s\n", a.DriverName))
	if result.ImageName != "" {
		sb.WriteString(fmt.Sprintf("崩溃映像: %s\n", result.ImageName))
	}
	if result.ProcessName != "" {
		sb.WriteString(fmt.Sprintf("进程名称: %s\n", result.ProcessName))
	}
	if result.FailureBucket != "" {
		sb.WriteString(fmt.Sprintf("失败桶ID: %s\n", result.FailureBucket))
	}

	if len(result.UnloadedModules) > 0 {
		sb.WriteString("\n已卸载的模块:\n")
		for _, mod := range result.UnloadedModules {
			sb.WriteString(fmt.Sprintf("  %s - %s  %s\n", mod.BaseAddress, mod.EndAddress, mod.Name))
		}
	}

	if result.ExceptionAddress != "" {
		sb.WriteString(fmt.Sprintf("\n异常地址: %s\n", result.ExceptionAddress))
	}
	if result.ExceptionCode != "" {
		sb.WriteString(fmt.Sprintf("异常代码: %s\n", result.ExceptionCode))
	}

	if result.CrashReason != "" {
		sb.WriteString("\n========================================\n")
		sb.WriteString("崩溃原因分析:\n")
		sb.WriteString("========================================\n")
		sb.WriteString(result.CrashReason)
		sb.WriteString("\n")
	}

	if len(result.CallStack) > 0 {
		sb.WriteString("\n调用栈:\n")
		sb.WriteString("========================================\n")
		for _, frame := range result.CallStack {
			marker := "  "
			if frame.IsDriver {
				marker = "* "
			}
			sb.WriteString(fmt.Sprintf("%s#%d\n", marker, frame.Index))
			if frame.ModuleName != "" {
				sb.WriteString(fmt.Sprintf("    模块: %s\n", frame.ModuleName))
			}
			if frame.FuncName != "" {
				sb.WriteString(fmt.Sprintf("    函数: %s\n", frame.FuncName))
			}
			if frame.FuncOffset != "" {
				sb.WriteString(fmt.Sprintf("    偏移: %s\n", frame.FuncOffset))
			}
			if frame.SourceFile != "" {
				sb.WriteString(fmt.Sprintf("    源文件: %s:%s\n", frame.SourceFile, frame.SourceLine))
			}
			sb.WriteString(fmt.Sprintf("    地址: %s\n", frame.RetAddr))
		}
	}

	if result.CrashFunction != "" {
		sb.WriteString("\n========================================\n")
		sb.WriteString(fmt.Sprintf("崩溃函数: %s!%s\n", a.DriverName, result.CrashFunction))
		sb.WriteString(fmt.Sprintf("偏移: %s\n", result.CrashOffset))
		sb.WriteString("========================================\n")
	} else if result.IsDriverCrash {
		sb.WriteString("\n检测到驱动卸载后回调问题\n")
	}

	if a.DriverSys != "" {
		sb.WriteString(fmt.Sprintf("\n驱动文件: %s\n", a.DriverSys))
	}

	if result.AIAnalysis != "" {
		sb.WriteString("\n========================================\n")
		sb.WriteString("AI分析部分 (原始输出):\n")
		sb.WriteString("========================================\n")
		sb.WriteString(result.AIAnalysis)
		sb.WriteString("\n========================================\n")
	}

	return sb.String()
}

func parseHex(s string) uint64 {
	s = strings.ReplaceAll(s, "`", "")
	s = strings.TrimPrefix(strings.TrimPrefix(s, "0x"), "0X")
	var val uint64
	fmt.Sscanf(s, "%x", &val)
	return val
}
