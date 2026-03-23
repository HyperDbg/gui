package bsodanalyzer

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
	pdbex "github.com/ddkwork/HyperDbg/walker/pdbex"
)

type Analyzer struct {
	kd        *kerneldump.KernelDump
	pdb       *pdbex.PDB
	driverPDB string
}

type AnalysisResult struct {
	BugCheckCode     uint32
	BugCheckName     string
	Parameters       [4]uint64
	CrashAddress     uint64
	CrashModule      string
	CrashFunction    string
	CrashOffset      uint64
	SourceFile       string
	SourceLine       uint32
	StackTrace       []kerneldump.StackFrame
	ExceptionCode    uint32
	ExceptionName    string
	ExceptionAddress uint64
	Modules          []kerneldump.Module
}

func New(dumpPath string) (*Analyzer, error) {
	kd, err := kerneldump.Parse(dumpPath)
	if err != nil {
		return nil, err
	}

	return &Analyzer{
		kd: kd,
	}, nil
}

func (a *Analyzer) LoadDriverPDB(pdbPath string) error {
	pdb := pdbex.NewPDB()
	if err := pdb.Open(pdbPath); err != nil {
		return err
	}

	a.pdb = pdb
	a.driverPDB = pdbPath
	return nil
}

func (a *Analyzer) Analyze() (*AnalysisResult, error) {
	result := &AnalysisResult{
		BugCheckCode: a.kd.Header.BugCheckCode,
		BugCheckName: a.kd.GetBugCheckName(),
		Parameters: [4]uint64{
			a.kd.Header.BugCheckParameter1,
			a.kd.Header.BugCheckParameter2,
			a.kd.Header.BugCheckParameter3,
			a.kd.Header.BugCheckParameter4,
		},
	}

	a.analyzeException(result)
	a.analyzeCrashAddress(result)
	a.analyzeStackTrace(result)
	a.analyzeModules(result)

	if a.pdb != nil {
		a.resolveSymbols(result)
	}

	return result, nil
}

func (a *Analyzer) analyzeException(result *AnalysisResult) {
	exception, err := a.kd.GetExceptionRecord()
	if err != nil {
		return
	}

	result.ExceptionCode = exception.ExceptionCode
	result.ExceptionName = a.kd.GetExceptionCodeName(exception.ExceptionCode)
	result.ExceptionAddress = exception.ExceptionAddress
}

func (a *Analyzer) GetExceptionCodeName(code uint32) string {
	codes := map[uint32]string{
		0x80000001: "EXCEPTION_ACCESS_VIOLATION",
		0x80000002: "EXCEPTION_ARRAY_BOUNDS_EXCEEDED",
		0x80000003: "EXCEPTION_BREAKPOINT",
		0x80000004: "EXCEPTION_DATATYPE_MISALIGNMENT",
		0x80000005: "EXCEPTION_FLT_DENORMAL_OPERAND",
		0x80000006: "EXCEPTION_FLT_DIVIDE_BY_ZERO",
		0x80000007: "EXCEPTION_FLT_INEXACT_RESULT",
		0x80000008: "EXCEPTION_FLT_INVALID_OPERATION",
		0x80000009: "EXCEPTION_FLT_OVERFLOW",
		0x8000000A: "EXCEPTION_FLT_STACK_CHECK",
		0x8000000B: "EXCEPTION_FLT_UNDERFLOW",
		0x8000000C: "EXCEPTION_INTEGER_DIVIDE_BY_ZERO",
		0x8000000D: "EXCEPTION_INTEGER_OVERFLOW",
		0x8000000E: "EXCEPTION_PRIV_INSTRUCTION",
		0x8000000F: "EXCEPTION_IN_PAGE_ERROR",
		0x80000010: "EXCEPTION_ILLEGAL_INSTRUCTION",
		0x80000011: "EXCEPTION_NONCONTINUABLE_EXCEPTION",
		0x80000012: "EXCEPTION_STACK_OVERFLOW",
		0x80000013: "EXCEPTION_INVALID_DISPOSITION",
		0x80000014: "EXCEPTION_GUARD_PAGE",
		0x80000015: "EXCEPTION_INVALID_HANDLE",
		0x8000001D: "EXCEPTION_POSSIBLE_DEADLOCK",
		0x8000001E: "EXCEPTION_INVALID_LOCK_SEQUENCE",
		0x80000026: "EXCEPTION_INVALID_OPCODE",
		0xC0000005: "STATUS_ACCESS_VIOLATION",
		0xC0000006: "STATUS_IN_PAGE_ERROR",
		0xC0000008: "STATUS_INVALID_HANDLE",
		0xC0000025: "STATUS_NO_MEMORY",
		0xC00000FD: "STATUS_STACK_BUFFER_OVERRUN",
	}

	if name, ok := codes[code]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_EXCEPTION_0x%08X", code)
}

func (a *Analyzer) analyzeModules(result *AnalysisResult) {
	modules, err := a.kd.GetModules()
	if err != nil {
		return
	}

	result.Modules = modules
}

func (a *Analyzer) analyzeStackTrace(result *AnalysisResult) {
	bugCheckCode, bugCheckName := a.kd.DetectBugCheckFromPhysical()
	if bugCheckCode != result.BugCheckCode && bugCheckCode != 0 {
		result.BugCheckCode = bugCheckCode
		result.BugCheckName = bugCheckName
	}

	frames, err := a.kd.GetStackTraceFromPhysical()
	if err != nil {
		return
	}

	result.StackTrace = frames
}

func (a *Analyzer) analyzeCrashAddress(result *AnalysisResult) {
	param1 := a.kd.Header.BugCheckParameter1

	moduleID := (param1 >> 56) & 0xFF
	exceptionCode := param1 & 0xFF

	result.CrashAddress = param1 & 0x00FFFFFFFFFFFFFF

	if moduleID != 0 {
		result.CrashModule = fmt.Sprintf("Module_0x%02X", moduleID)
	}

	if exceptionCode != 0 {
		result.CrashFunction = fmt.Sprintf("ExceptionCode_0x%02X", exceptionCode)
	}
}

func (a *Analyzer) resolveSymbols(result *AnalysisResult) {
	if a.pdb == nil {
		return
	}

	crashRVA := uint32(result.CrashAddress & 0xFFFFFFFF)

	funcName, ok := a.pdb.GetFunctionByOffset(uint64(crashRVA))
	if ok {
		result.CrashFunction = funcName
	}

	fileName, lineNumber, ok := a.pdb.FindSourceLineByRVA(crashRVA)
	if ok {
		result.SourceFile = filepath.Base(fileName)
		result.SourceLine = lineNumber
	}
}

func (a *Analyzer) Close() {
	if a.pdb != nil {
		a.pdb.Close()
	}
	a.kd.Close()
}

func (r *AnalysisResult) String() string {
	var s strings.Builder
	s.WriteString(fmt.Sprintf("=== BSOD Analysis Result ===\n"))
	s.WriteString(fmt.Sprintf("BugCheck: 0x%08X (%s)\n", r.BugCheckCode, r.BugCheckName))
	s.WriteString(fmt.Sprintf("Parameters:\n"))
	s.WriteString(fmt.Sprintf("  1: 0x%016X\n", r.Parameters[0]))
	s.WriteString(fmt.Sprintf("  2: 0x%016X\n", r.Parameters[1]))
	s.WriteString(fmt.Sprintf("  3: 0x%016X\n", r.Parameters[2]))
	s.WriteString(fmt.Sprintf("  4: 0x%016X\n", r.Parameters[3]))
	s.WriteString(fmt.Sprintf("\nCrash Analysis:\n"))
	s.WriteString(fmt.Sprintf("  Address: 0x%016X\n", r.CrashAddress))
	if r.CrashModule != "" {
		s.WriteString(fmt.Sprintf("  Module: %s\n", r.CrashModule))
	}
	if r.CrashFunction != "" {
		s.WriteString(fmt.Sprintf("  Function: %s\n", r.CrashFunction))
	}
	if r.SourceFile != "" {
		s.WriteString(fmt.Sprintf("  Source: %s:%d\n", r.SourceFile, r.SourceLine))
	}

	if r.ExceptionCode != 0 {
		s.WriteString(fmt.Sprintf("\nException Information:\n"))
		s.WriteString(fmt.Sprintf("  Code: 0x%08X (%s)\n", r.ExceptionCode, r.ExceptionName))
		s.WriteString(fmt.Sprintf("  Address: 0x%016X\n", r.ExceptionAddress))
	}

	if len(r.Modules) > 0 {
		s.WriteString(fmt.Sprintf("\nModule Information:\n"))
		s.WriteString(fmt.Sprintf("  Total modules: %d\n", len(r.Modules)))
		for i, mod := range r.Modules {
			if i >= 10 {
				s.WriteString(fmt.Sprintf("  ... and %d more\n", len(r.Modules)-10))
				break
			}
			s.WriteString(fmt.Sprintf("  [%d] %s\n", i, mod.GetName()))
			s.WriteString(fmt.Sprintf("       Base: 0x%016X, Size: 0x%08X\n", mod.DllBase, mod.SizeOfImage))
		}
	}

	if len(r.StackTrace) > 0 {
		s.WriteString(fmt.Sprintf("\nCall Stack:\n"))
		s.WriteString(fmt.Sprintf("========================================\n"))
		for i, frame := range r.StackTrace {
			if i >= 10 {
				s.WriteString(fmt.Sprintf("  ... and %d more\n", len(r.StackTrace)-10))
				break
			}
			s.WriteString(fmt.Sprintf("  #%d %s\n", frame.FrameNum, frame.CallSite))
		}
	}

	s.WriteString("\nRecommendations:\n")
	switch r.BugCheckCode {
	case 0x00000008, 0x0000000A:
		s.WriteString("  - Check IRQL levels before accessing memory\n")
		s.WriteString("  - Ensure all memory is properly mapped\n")
		s.WriteString("  - Use proper synchronization primitives\n")
		s.WriteString("  - Avoid accessing paged memory at DISPATCH_LEVEL+\n")
	case 0x000000EF, 0x000000F4:
		s.WriteString("  - Ensure all callbacks are unregistered before driver unload\n")
		s.WriteString("  - Remove EPT hooks before unloading\n")
		s.WriteString("  - Properly clean up kernel objects\n")
		s.WriteString("  - Check for memory leaks\n")
	default:
		s.WriteString("  - Review the crash address and parameters\n")
		s.WriteString("  - Check for memory corruption\n")
		s.WriteString("  - Verify proper use of kernel APIs\n")
	}

	return s.String()
}
