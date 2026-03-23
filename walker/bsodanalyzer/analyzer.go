package bsodanalyzer

import (
	"fmt"
	"path/filepath"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
	pdbex "github.com/ddkwork/HyperDbg/walker/pdbex"
)

type Analyzer struct {
	kd        *kerneldump.KernelDump
	pdb       *pdbex.PDB
	driverPDB string
}

type AnalysisResult struct {
	BugCheckCode  uint32
	BugCheckName  string
	Parameters    [4]uint64
	CrashAddress  uint64
	CrashModule   string
	CrashFunction string
	CrashOffset   uint64
	SourceFile    string
	SourceLine    uint32
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

	a.analyzeCrashAddress(result)

	if a.pdb != nil {
		a.resolveSymbols(result)
	}

	return result, nil
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
	s := fmt.Sprintf("=== BSOD Analysis Result ===\n")
	s += fmt.Sprintf("BugCheck: 0x%08X (%s)\n", r.BugCheckCode, r.BugCheckName)
	s += fmt.Sprintf("Parameters:\n")
	s += fmt.Sprintf("  1: 0x%016X\n", r.Parameters[0])
	s += fmt.Sprintf("  2: 0x%016X\n", r.Parameters[1])
	s += fmt.Sprintf("  3: 0x%016X\n", r.Parameters[2])
	s += fmt.Sprintf("  4: 0x%016X\n", r.Parameters[3])
	s += fmt.Sprintf("\nCrash Analysis:\n")
	s += fmt.Sprintf("  Address: 0x%016X\n", r.CrashAddress)
	if r.CrashModule != "" {
		s += fmt.Sprintf("  Module: %s\n", r.CrashModule)
	}
	if r.CrashFunction != "" {
		s += fmt.Sprintf("  Function: %s\n", r.CrashFunction)
	}
	if r.SourceFile != "" {
		s += fmt.Sprintf("  Source: %s:%d\n", r.SourceFile, r.SourceLine)
	}

	s += "\nRecommendations:\n"
	switch r.BugCheckCode {
	case 0x00000008, 0x0000000A:
		s += "  - Check IRQL levels before accessing memory\n"
		s += "  - Ensure all memory is properly mapped\n"
		s += "  - Use proper synchronization primitives\n"
		s += "  - Avoid accessing paged memory at DISPATCH_LEVEL+\n"
	case 0x000000EF, 0x000000F4:
		s += "  - Ensure all callbacks are unregistered before driver unload\n"
		s += "  - Remove EPT hooks before unloading\n"
		s += "  - Properly clean up kernel objects\n"
		s += "  - Check for memory leaks\n"
	default:
		s += "  - Review the crash address and parameters\n"
		s += "  - Check for memory corruption\n"
		s += "  - Verify proper use of kernel APIs\n"
	}

	return s
}
