package bsod

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

type AnalysisResult struct {
	DumpFile         string
	DriverName       string
	UnloadedModules  []UnloadedModule
	ExceptionAddress string
	ExceptionCode    string
	ExceptionDesc    string
	ProcessName      string
	ImageName        string
	FailureBucket    string
	CallStack        []StackFrame
	CrashFunction    string
	CrashOffset      string
	IsDriverCrash    bool
	CrashReason      string
	RawOutput        string
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
	Args       []string
	IsDriver   bool
	ModuleName string
	FuncName   string
	FuncOffset string
	SourceFile string
	SourceLine string
}

type Analyzer struct {
	WinDbgPath string
	SymbolPath string
	DriverName string
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		WinDbgPath: `C:\Microsoft.WinDbg_1.2601.12001.0_x64__8wekyb3d8bbwe\amd64\kd.exe`,
		SymbolPath: `srv*C:\Symbols*https://msdl.microsoft.com/download/symbols`,
		DriverName: "hyperkd",
	}
}

func (a *Analyzer) SetDriverName(name string) {
	a.DriverName = name
}

func (a *Analyzer) SetPdbPath(path string) {
	if a.SymbolPath != "" {
		a.SymbolPath += ";" + path
	} else {
		a.SymbolPath = path
	}
}

func (a *Analyzer) FindLatestMinidump() string {
	minidumpDir := `C:\Windows\Minidump`
	entries := mylog.Check2(os.ReadDir(minidumpDir))

	var dmpFiles []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(strings.ToLower(entry.Name()), ".dmp") {
			dmpFiles = append(dmpFiles, entry)
		}
	}

	if len(dmpFiles) == 0 {
		mylog.Check("在 " + minidumpDir + " 中找不到 dump 文件")
	}

	sort.Slice(dmpFiles, func(i, j int) bool {
		infoI, _ := dmpFiles[i].Info()
		infoJ, _ := dmpFiles[j].Info()
		return infoI.ModTime().After(infoJ.ModTime())
	})

	return filepath.Join(minidumpDir, dmpFiles[0].Name())
}

func (a *Analyzer) runKdCommand(dumpFile, command string) string {
	args := []string{a.WinDbgPath, "-z", dumpFile, "-y", a.SymbolPath, "-c", command}
	output := stream.RunCommandSilent(args...)
	return output.String()
}

func (a *Analyzer) Analyze(dumpFile string) *AnalysisResult {
	result := &AnalysisResult{
		DumpFile:   dumpFile,
		DriverName: a.DriverName,
	}

	output := a.runKdCommand(dumpFile, "!analyze -v; r; lm; .exr -1; kv; q")
	result.RawOutput = output

	a.parseUnloadedModules(output, result)
	a.parseExceptionInfo(output, result)
	a.parseCallStack(output, result)
	a.parseAnalyzeOutput(output, result)
	a.analyzeCrashReason(result)

	if result.ExceptionAddress != "" && len(result.UnloadedModules) > 0 {
		a.checkDriverCrash(result)
	}

	return result
}

func (a *Analyzer) parseAnalyzeOutput(output string, result *AnalysisResult) {
	processRegex := regexp.MustCompile(`PROCESS_NAME:\s*(\S+)`)
	if match := processRegex.FindStringSubmatch(output); match != nil {
		result.ProcessName = match[1]
	}

	imageRegex := regexp.MustCompile(`IMAGE_NAME:\s*(\S+)`)
	if match := imageRegex.FindStringSubmatch(output); match != nil {
		result.ImageName = match[1]
	}

	bucketRegex := regexp.MustCompile(`FAILURE_BUCKET_ID:\s*(\S+)`)
	if match := bucketRegex.FindStringSubmatch(output); match != nil {
		result.FailureBucket = match[1]
	}
}

func (a *Analyzer) analyzeCrashReason(result *AnalysisResult) {
	var reasons []string

	exceptionCode := result.ExceptionCode
	if exceptionCode != "" {
		switch exceptionCode {
		case "80000003":
			reasons = append(reasons, "断点异常(Breakpoint Exception) - 可能是调试断点触发或代码中的int 3指令")
		case "c0000005":
			reasons = append(reasons, "访问违规(Access Violation) - 内存访问错误，可能是空指针或无效地址")
		case "c000001d":
			reasons = append(reasons, "非法指令(Illegal Instruction) - 执行了无效的CPU指令")
		case "c0000094":
			reasons = append(reasons, "整数除零(Integer Division by Zero)")
		case "c0000095":
			reasons = append(reasons, "整数溢出(Integer Overflow)")
		case "c0000096":
			reasons = append(reasons, "特权指令(Privileged Instruction) - 用户态执行了内核指令")
		case "c0000409":
			reasons = append(reasons, "栈缓冲区溢出(Stack Buffer Overflow) - 可能是安全漏洞利用")
		case "c0000374":
			reasons = append(reasons, "堆损坏(Heap Corruption)")
		default:
			reasons = append(reasons, fmt.Sprintf("异常代码: %s", exceptionCode))
		}
	}

	if len(result.UnloadedModules) > 0 {
		reasons = append(reasons, fmt.Sprintf("检测到 %d 个已卸载的 %s.sys 实例", len(result.UnloadedModules), a.DriverName))
		reasons = append(reasons, "驱动被多次加载/卸载，可能存在回调未清理问题")
	}

	if result.FailureBucket != "" {
		if strings.Contains(result.FailureBucket, "WRONG_SYMBOLS") {
			reasons = append(reasons, "符号文件不匹配，需要正确的PDB文件进行详细分析")
		}
		if strings.Contains(result.FailureBucket, "AV") {
			reasons = append(reasons, "访问违规导致的崩溃")
		}
	}

	if result.ImageName != "" && strings.Contains(result.ImageName, "wrong.symbols") {
		reasons = append(reasons, "内核符号不匹配，请更新Windows符号文件")
	}

	for _, frame := range result.CallStack {
		if strings.Contains(frame.CallSite, "nt") && strings.Contains(frame.CallSite, "Ki") {
			reasons = append(reasons, "崩溃发生在内核异常处理路径")
			break
		}
	}

	if len(reasons) > 0 {
		result.CrashReason = strings.Join(reasons, "\n")
	} else {
		result.CrashReason = "无法确定具体崩溃原因，请检查原始输出"
	}
}

func (a *Analyzer) parseUnloadedModules(output string, result *AnalysisResult) {
	unloadedRegex := regexp.MustCompile(`Unloaded modules:\r?\n((?:.*\r?\n)+)`)
	unloadedMatch := unloadedRegex.FindStringSubmatch(output)
	if unloadedMatch == nil {
		return
	}

	moduleRegex := regexp.MustCompile(`([0-9a-fA-F` + "`" + `]+)\s+([0-9a-fA-F` + "`" + `]+)\s+(\S+\.sys)`)
	matches := moduleRegex.FindAllStringSubmatch(unloadedMatch[1], -1)

	for _, match := range matches {
		if strings.Contains(strings.ToLower(match[3]), strings.ToLower(a.DriverName)) {
			result.UnloadedModules = append(result.UnloadedModules, UnloadedModule{
				BaseAddress: strings.ReplaceAll(match[1], "`", ""),
				EndAddress:  strings.ReplaceAll(match[2], "`", ""),
				Name:        match[3],
			})
		}
	}
}

func (a *Analyzer) parseExceptionInfo(output string, result *AnalysisResult) {
	addrRegex := regexp.MustCompile(`ExceptionAddress:\s*([0-9a-fA-F` + "`" + `]+)`)
	if match := addrRegex.FindStringSubmatch(output); match != nil {
		result.ExceptionAddress = strings.ReplaceAll(match[1], "`", "")
	}

	codeRegex := regexp.MustCompile(`ExceptionCode:\s*([0-9a-fA-F]+)`)
	if match := codeRegex.FindStringSubmatch(output); match != nil {
		result.ExceptionCode = match[1]
	}
}

func (a *Analyzer) parseCallStack(output string, result *AnalysisResult) {
	startIdx := strings.Index(output, "# Child-SP")
	if startIdx == -1 {
		return
	}

	endIdx := strings.Index(output[startIdx:], "quit:")
	if endIdx == -1 {
		endIdx = strings.Index(output[startIdx:], "NatVis")
	}
	if endIdx == -1 {
		endIdx = len(output) - startIdx
	}

	stackMatch := output[startIdx : startIdx+endIdx]

	lines := strings.SplitSeq(stackMatch, "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || !strings.HasPrefix(line, "#") {
			continue
		}

		frame := StackFrame{}

		parts := strings.Fields(line)
		if len(parts) < 4 {
			continue
		}

		fmt.Sscanf(parts[0], "#%d", &frame.Index)
		frame.ChildSP = strings.ReplaceAll(parts[1], "`", "")
		frame.RetAddr = strings.ReplaceAll(parts[2], "`", "")

		_, after, ok := strings.Cut(line, ":")
		if ok {
			afterColon := strings.TrimSpace(after)

			argsEnd := strings.LastIndex(afterColon, ":")
			if argsEnd != -1 {
				callSite := strings.TrimSpace(afterColon[argsEnd+1:])
				frame.CallSite = callSite

				if before, after, ok := strings.Cut(callSite, "!"); ok {
					frame.ModuleName = before
					rest := after

					if plusIdx := strings.Index(rest, "+"); plusIdx != -1 {
						frame.FuncName = rest[:plusIdx]
						frame.FuncOffset = rest[plusIdx:]
					} else {
						frame.FuncName = rest
					}
				} else if strings.Contains(callSite, "+") {
					plusIdx := strings.Index(callSite, "+")
					frame.ModuleName = callSite[:plusIdx]
					frame.FuncOffset = callSite[plusIdx:]
				} else {
					frame.ModuleName = callSite
				}
			}
		}

		if strings.Contains(frame.CallSite, "Unloaded_"+a.DriverName) ||
			strings.Contains(frame.CallSite, a.DriverName) {
			frame.IsDriver = true
			result.IsDriverCrash = true
		}

		result.CallStack = append(result.CallStack, frame)
	}
}

func (a *Analyzer) checkDriverCrash(result *AnalysisResult) {
	if len(result.UnloadedModules) == 0 || result.ExceptionAddress == "" {
		return
	}

	for _, mod := range result.UnloadedModules {
		baseAddr := parseHex(mod.BaseAddress)
		endAddr := parseHex(mod.EndAddress)
		exceptionAddr := parseHex(result.ExceptionAddress)

		if exceptionAddr >= baseAddr && exceptionAddr < endAddr {
			offset := exceptionAddr - baseAddr
			result.CrashOffset = fmt.Sprintf("0x%X", offset)
			result.IsDriverCrash = true
			break
		}
	}
}

func (a *Analyzer) LocateCrashFunction(result *AnalysisResult, driverPath string) {
	if result.CrashOffset == "" {
		mylog.Warning("没有找到崩溃偏移")
		return
	}

	cmd := fmt.Sprintf("lm; ln %s+%s; q", a.DriverName, result.CrashOffset)
	output := a.runKdCommand(driverPath, cmd)

	funcRegex := regexp.MustCompile(`\(` + regexp.QuoteMeta(a.DriverName) + `!([^)\s]+)`)
	if match := funcRegex.FindStringSubmatch(output); match != nil {
		result.CrashFunction = match[1]
		return
	}

	exactRegex := regexp.MustCompile(`Exact matches:\s*\r?\n\s*` + regexp.QuoteMeta(a.DriverName) + `!([^\s]+)`)
	if match := exactRegex.FindStringSubmatch(output); match != nil {
		result.CrashFunction = match[1]
		return
	}

	mylog.Warning("无法定位崩溃函数")
}

func (a *Analyzer) ResolveSourceLines(result *AnalysisResult, driverPath string) {
	if len(result.CallStack) == 0 {
		return
	}

	for i := range result.CallStack {
		frame := &result.CallStack[i]
		if frame.RetAddr == "" || frame.RetAddr == "0000000000000000" {
			continue
		}

		cmd := fmt.Sprintf("ln %s; q", frame.RetAddr)
		output := a.runKdCommand(driverPath, cmd)
		if output == "" {
			continue
		}

		sourceRegex := regexp.MustCompile(`\[(.+):(\d+)\]`)
		if match := sourceRegex.FindStringSubmatch(output); match != nil {
			frame.SourceFile = match[1]
			frame.SourceLine = match[2]
		}

		funcRegex := regexp.MustCompile(`([a-zA-Z_][a-zA-Z0-9_]*)\s*\(`)
		if match := funcRegex.FindStringSubmatch(output); match != nil {
			if frame.FuncName == "" || frame.FuncName == "0xef" {
				frame.FuncName = match[1]
			}
		}
	}
}

func parseHex(s string) uint64 {
	s = strings.ReplaceAll(s, "`", "")
	var val uint64
	fmt.Sscanf(s, "%x", &val)
	return val
}

func (r *AnalysisResult) String() string {
	var sb strings.Builder

	sb.WriteString("=== BSOD Dump 分析结果 ===\n")
	sb.WriteString(fmt.Sprintf("Dump 文件: %s\n", r.DumpFile))
	sb.WriteString(fmt.Sprintf("驱动名称: %s\n", r.DriverName))

	if r.ImageName != "" {
		sb.WriteString(fmt.Sprintf("崩溃映像: %s\n", r.ImageName))
	}
	if r.ProcessName != "" {
		sb.WriteString(fmt.Sprintf("进程名称: %s\n", r.ProcessName))
	}
	if r.FailureBucket != "" {
		sb.WriteString(fmt.Sprintf("失败桶ID: %s\n", r.FailureBucket))
	}

	if len(r.UnloadedModules) > 0 {
		sb.WriteString("\n已卸载的模块:\n")
		for _, mod := range r.UnloadedModules {
			sb.WriteString(fmt.Sprintf("  %s - %s  %s\n", mod.BaseAddress, mod.EndAddress, mod.Name))
		}
	}

	if r.ExceptionAddress != "" {
		sb.WriteString(fmt.Sprintf("\n异常地址: %s\n", r.ExceptionAddress))
	}
	if r.ExceptionCode != "" {
		sb.WriteString(fmt.Sprintf("异常代码: %s\n", r.ExceptionCode))
	}

	if r.CrashReason != "" {
		sb.WriteString("\n========================================\n")
		sb.WriteString("崩溃原因分析:\n")
		sb.WriteString("========================================\n")
		sb.WriteString(r.CrashReason)
		sb.WriteString("\n")
	}

	if len(r.CallStack) > 0 {
		sb.WriteString("\n调用栈:\n")
		sb.WriteString("========================================\n")
		for _, frame := range r.CallStack {
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
			sb.WriteString(fmt.Sprintf("    地址: %s -> %s\n", frame.ChildSP, frame.RetAddr))
			sb.WriteString("\n")
		}
	}

	if r.CrashFunction != "" {
		sb.WriteString("\n========================================\n")
		sb.WriteString(fmt.Sprintf("崩溃函数: %s!%s\n", r.DriverName, r.CrashFunction))
		sb.WriteString(fmt.Sprintf("偏移: %s\n", r.CrashOffset))
		sb.WriteString("========================================\n")
	} else if r.IsDriverCrash {
		sb.WriteString("\n检测到驱动卸载后回调问题\n")
		sb.WriteString("这可能是驱动卸载时未正确清理回调导致的\n")
	}

	return sb.String()
}
