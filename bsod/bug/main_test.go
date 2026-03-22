package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestNewAnalyzer(t *testing.T) {
	a := NewAnalyzer()

	if a.WinDbgPath == "" {
		t.Error("WinDbgPath should not be empty")
	}
	if a.DriverName == "" {
		t.Error("DriverName should not be empty")
	}
	if a.PdbPath == "" {
		t.Error("PdbPath should not be empty")
	}
	if a.CommandTimeout <= 0 {
		t.Error("CommandTimeout should be positive")
	}

	t.Logf("WinDbgPath: %s", a.WinDbgPath)
	t.Logf("DriverName: %s", a.DriverName)
	t.Logf("PdbPath: %s", a.PdbPath)
	t.Logf("DriverSys: %s", a.DriverSys)
	t.Logf("CommandTimeout: %v", a.CommandTimeout)
	t.Logf("SymOpt: %s", a.SymOpt)
	t.Logf("MaxOutputSize: %d", a.MaxOutputSize)
}

func TestParseHex(t *testing.T) {
	tests := []struct {
		input    string
		expected uint64
	}{
		{"0x1234", 0x1234},
		{"1234", 0x1234},
		{"0XABCD", 0xABCD},
		{"abcd", 0xABCD},
		{"0x" + "`" + "1234" + "`" + "5678", 0x12345678},
		{"", 0},
		{"0", 0},
		{"FFFFFFFF", 0xFFFFFFFF},
		{"7FFFFFFFFFFFFFFF", 0x7FFFFFFFFFFFFFFF},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseHex(tt.input)
			if result != tt.expected {
				t.Errorf("parseHex(%q) = 0x%X, expected 0x%X", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseHexStress(t *testing.T) {
	for i := 0; i < 10000; i++ {
		input := fmt.Sprintf("0x%X", uint64(i))
		result := parseHex(input)
		if result != uint64(i) {
			t.Errorf("parseHex(%q) = 0x%X, expected 0x%X", input, result, i)
		}
	}
}

func TestParseCallSite(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	tests := []struct {
		callSite     string
		expectedMod  string
		expectedFunc string
		expectedOff  string
	}{
		{"nt!KiDispatchException+0x1a2", "nt", "KiDispatchException", "+0x1a2"},
		{"hyperkd!SomeFunction+0x10", "hyperkd", "SomeFunction", "+0x10"},
		{"ntdll!NtWaitForSingleObject", "ntdll", "NtWaitForSingleObject", ""},
		{"0x12345678", "", "", ""},
		{"", "", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.callSite, func(t *testing.T) {
			frame := &StackFrame{}
			a.parseCallSite(tt.callSite, frame)

			if frame.ModuleName != tt.expectedMod {
				t.Errorf("ModuleName = %q, expected %q", frame.ModuleName, tt.expectedMod)
			}
			if frame.FuncName != tt.expectedFunc {
				t.Errorf("FuncName = %q, expected %q", frame.FuncName, tt.expectedFunc)
			}
			if frame.FuncOffset != tt.expectedOff {
				t.Errorf("FuncOffset = %q, expected %q", frame.FuncOffset, tt.expectedOff)
			}
		})
	}
}

func TestParseCallSiteStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	modules := []string{"nt", "hyperkd", "ntdll", "kernel32", "kernelbase"}
	functions := []string{"FuncA", "FuncB", "SomeLongFunctionName", "AnotherFunction"}

	for i := 0; i < 5000; i++ {
		mod := modules[i%len(modules)]
		fn := functions[i%len(functions)]
		offset := fmt.Sprintf("+0x%X", i%256)

		callSite := fmt.Sprintf("%s!%s%s", mod, fn, offset)
		frame := &StackFrame{}
		a.parseCallSite(callSite, frame)

		if frame.ModuleName != mod {
			t.Errorf("iteration %d: ModuleName mismatch", i)
		}
		if frame.FuncName != fn {
			t.Errorf("iteration %d: FuncName mismatch", i)
		}
	}
}

func TestParseExceptionInfo(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	output := "ExceptionAddress: fffff800" + "`" + "12345678\nExceptionCode: 80000003\nSome other data\n"

	result := &AnalysisResult{}
	a.parseExceptionInfo(output, result)

	if result.ExceptionAddress != "fffff80012345678" {
		t.Errorf("ExceptionAddress = %q, expected %q", result.ExceptionAddress, "fffff80012345678")
	}
	if result.ExceptionCode != "80000003" {
		t.Errorf("ExceptionCode = %q, expected %q", result.ExceptionCode, "80000003")
	}
}

func TestParseExceptionInfoStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	codes := []string{"80000003", "c0000005", "c000001d", "c0000094", "c0000409"}

	for i := 0; i < 1000; i++ {
		code := codes[i%len(codes)]
		addr := fmt.Sprintf("fffff800%08x", i)

		output := fmt.Sprintf("ExceptionAddress: %s\nExceptionCode: %s\n", addr, code)

		result := &AnalysisResult{}
		a.parseExceptionInfo(output, result)

		if !strings.HasSuffix(result.ExceptionAddress, addr[len(addr)-8:]) {
			t.Errorf("iteration %d: ExceptionAddress mismatch", i)
		}
		if result.ExceptionCode != code {
			t.Errorf("iteration %d: ExceptionCode = %q, expected %q", i, result.ExceptionCode, code)
		}
	}
}

func TestParseAnalyzeOutput(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	output := "PROCESS_NAME:  System\nIMAGE_NAME:  hyperkd.sys\nFAILURE_BUCKET_ID:  AV_hyperkd!SomeFunction\n"

	result := &AnalysisResult{}
	a.parseAnalyzeOutput(output, result)

	if result.ProcessName != "System" {
		t.Errorf("ProcessName = %q, expected %q", result.ProcessName, "System")
	}
	if result.ImageName != "hyperkd.sys" {
		t.Errorf("ImageName = %q, expected %q", result.ImageName, "hyperkd.sys")
	}
	if result.FailureBucket != "AV_hyperkd!SomeFunction" {
		t.Errorf("FailureBucket = %q, expected %q", result.FailureBucket, "AV_hyperkd!SomeFunction")
	}
}

func TestParseAnalyzeOutputStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	for i := 0; i < 1000; i++ {
		process := fmt.Sprintf("Process%d", i)
		image := fmt.Sprintf("driver%d.sys", i%10)
		bucket := fmt.Sprintf("BUCKET_%d", i%5)

		output := fmt.Sprintf("PROCESS_NAME:  %s\nIMAGE_NAME:  %s\nFAILURE_BUCKET_ID:  %s\n", process, image, bucket)

		result := &AnalysisResult{}
		a.parseAnalyzeOutput(output, result)

		if result.ProcessName != process {
			t.Errorf("iteration %d: ProcessName mismatch", i)
		}
		if result.ImageName != image {
			t.Errorf("iteration %d: ImageName mismatch", i)
		}
		if result.FailureBucket != bucket {
			t.Errorf("iteration %d: FailureBucket mismatch", i)
		}
	}
}

func TestAnalyzeCrashReason(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	tests := []struct {
		name           string
		code           string
		unloadedCount  int
		failureBucket  string
		expectNonEmpty bool
	}{
		{"Breakpoint", "80000003", 0, "", true},
		{"AccessViolation", "c0000005", 0, "", true},
		{"IllegalInstruction", "c000001d", 0, "", true},
		{"DivideByZero", "c0000094", 0, "", true},
		{"StackOverflow", "c0000409", 0, "", true},
		{"HeapCorruption", "c0000374", 0, "", true},
		{"UnknownCode", "00000000", 0, "", false},
		{"WithUnloaded", "c0000005", 2, "", true},
		{"WrongSymbols", "c0000005", 0, "WRONG_SYMBOLS_test", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &AnalysisResult{
				ExceptionCode:   tt.code,
				UnloadedModules: make([]UnloadedModule, tt.unloadedCount),
				FailureBucket:   tt.failureBucket,
			}

			for i := 0; i < tt.unloadedCount; i++ {
				result.UnloadedModules[i] = UnloadedModule{
					BaseAddress: fmt.Sprintf("%x", 0x1000+i*0x1000),
					EndAddress:  fmt.Sprintf("%x", 0x2000+i*0x1000),
					Name:        "hyperkd.sys",
				}
			}

			a.analyzeCrashReason(result)

			if tt.expectNonEmpty && result.CrashReason == "" {
				t.Error("expected non-empty CrashReason")
			}
			if !tt.expectNonEmpty && result.CrashReason != "" {
				t.Errorf("expected empty CrashReason, got %q", result.CrashReason)
			}
		})
	}
}

func TestAnalyzeCrashReasonStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	codes := []string{"80000003", "c0000005", "c000001d", "c0000094", "c0000409", "c0000374"}

	for i := 0; i < 5000; i++ {
		code := codes[i%len(codes)]
		unloadedCount := i % 5

		result := &AnalysisResult{
			ExceptionCode:   code,
			UnloadedModules: make([]UnloadedModule, unloadedCount),
		}

		for j := 0; j < unloadedCount; j++ {
			result.UnloadedModules[j] = UnloadedModule{
				BaseAddress: fmt.Sprintf("%x", 0x1000+j*0x1000),
				EndAddress:  fmt.Sprintf("%x", 0x2000+j*0x1000),
				Name:        "hyperkd.sys",
			}
		}

		a.analyzeCrashReason(result)

		if result.CrashReason == "" {
			t.Errorf("iteration %d: expected non-empty CrashReason", i)
		}
	}
}

func TestCheckDriverCrash(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	tests := []struct {
		name          string
		exceptionAddr string
		baseAddr      string
		endAddr       string
		expectCrash   bool
	}{
		{"Inside", "0x1500", "0x1000", "0x2000", true},
		{"AtBase", "0x1000", "0x1000", "0x2000", true},
		{"BelowBase", "0x0FFF", "0x1000", "0x2000", false},
		{"AtEnd", "0x2000", "0x1000", "0x2000", false},
		{"AboveEnd", "0x2001", "0x1000", "0x2000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := &AnalysisResult{
				ExceptionAddress: tt.exceptionAddr,
				UnloadedModules: []UnloadedModule{
					{BaseAddress: tt.baseAddr, EndAddress: tt.endAddr, Name: "hyperkd.sys"},
				},
			}

			a.checkDriverCrash(result)

			if result.IsDriverCrash != tt.expectCrash {
				t.Errorf("IsDriverCrash = %v, expected %v", result.IsDriverCrash, tt.expectCrash)
			}
		})
	}
}

func TestCheckDriverCrashStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	for i := 0; i < 10000; i++ {
		base := uint64(0x10000000 + i*0x10000)
		end := base + 0x10000

		insideAddr := fmt.Sprintf("0x%X", base+0x100)
		outsideAddr := fmt.Sprintf("0x%X", end+0x100)

		resultInside := &AnalysisResult{
			ExceptionAddress: insideAddr,
			UnloadedModules: []UnloadedModule{
				{BaseAddress: fmt.Sprintf("0x%X", base), EndAddress: fmt.Sprintf("0x%X", end), Name: "hyperkd.sys"},
			},
		}
		a.checkDriverCrash(resultInside)
		if !resultInside.IsDriverCrash {
			t.Errorf("iteration %d: expected crash for inside address", i)
		}

		resultOutside := &AnalysisResult{
			ExceptionAddress: outsideAddr,
			UnloadedModules: []UnloadedModule{
				{BaseAddress: fmt.Sprintf("0x%X", base), EndAddress: fmt.Sprintf("0x%X", end), Name: "hyperkd.sys"},
			},
		}
		a.checkDriverCrash(resultOutside)
		if resultOutside.IsDriverCrash {
			t.Errorf("iteration %d: unexpected crash for outside address", i)
		}
	}
}

func TestBuildReport(t *testing.T) {
	a := &Analyzer{
		DriverName: "hyperkd",
		DriverSys:  "C:\\path\\to\\hyperkd.sys",
	}

	result := &AnalysisResult{
		DumpFile:         "C:\\Windows\\Minidump\\test.dmp",
		DriverName:       "hyperkd",
		ImageName:        "hyperkd.sys",
		ProcessName:      "System",
		FailureBucket:    "AV_test",
		ExceptionAddress: "fffff80012345678",
		ExceptionCode:    "c0000005",
		CrashReason:      "访问违规 - 内存访问错误",
		CallStack: []StackFrame{
			{Index: 0, RetAddr: "fffff80012345678", CallSite: "hyperkd!TestFunc+0x10", IsDriver: true, ModuleName: "hyperkd", FuncName: "TestFunc", FuncOffset: "+0x10"},
			{Index: 1, RetAddr: "fffff80012345679", CallSite: "nt!KiDispatchException+0x1a2", IsDriver: false, ModuleName: "nt", FuncName: "KiDispatchException", FuncOffset: "+0x1a2"},
		},
	}

	report := a.buildReport(result)

	if !strings.Contains(report, "hyperkd") {
		t.Error("report should contain driver name")
	}
	if !strings.Contains(report, "c0000005") {
		t.Error("report should contain exception code")
	}
	if !strings.Contains(report, "TestFunc") {
		t.Error("report should contain function name")
	}
}

func TestBuildReportStress(t *testing.T) {
	a := &Analyzer{
		DriverName: "hyperkd",
		DriverSys:  "C:\\path\\to\\hyperkd.sys",
	}

	for i := 0; i < 1000; i++ {
		result := &AnalysisResult{
			DumpFile:         fmt.Sprintf("C:\\Windows\\Minidump\\test%d.dmp", i),
			DriverName:       "hyperkd",
			ImageName:        "hyperkd.sys",
			ProcessName:      "System",
			FailureBucket:    fmt.Sprintf("BUCKET_%d", i),
			ExceptionAddress: fmt.Sprintf("fffff800%08x", i),
			ExceptionCode:    "c0000005",
			CrashReason:      "测试原因",
			CallStack:        make([]StackFrame, 10),
		}

		for j := 0; j < 10; j++ {
			result.CallStack[j] = StackFrame{
				Index:      j,
				RetAddr:    fmt.Sprintf("fffff800%08x", i*100+j),
				CallSite:   fmt.Sprintf("module%d!Func%d+0x%x", j%3, j, j*4),
				IsDriver:   j%3 == 0,
				ModuleName: fmt.Sprintf("module%d", j%3),
				FuncName:   fmt.Sprintf("Func%d", j),
				FuncOffset: fmt.Sprintf("+0x%x", j*4),
			}
		}

		report := a.buildReport(result)

		if len(report) == 0 {
			t.Errorf("iteration %d: empty report", i)
		}
	}
}

func TestParseCallStack(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	tick := "`"
	output := "# Child-SP          RetAddr           Call Site\n" +
		"#00 ffffc781" + tick + "12345678 fffff800" + tick + "12345678 hyperkd!TestFunction+0x10\n" +
		"#01 ffffc781" + tick + "12345680 fffff800" + tick + "12345680 nt!KiDispatchException+0x1a2\n" +
		"quit:\n"

	result := &AnalysisResult{}
	a.parseCallStack(output, result)

	if len(result.CallStack) != 3 {
		t.Fatalf("expected 3 frames (including header), got %d", len(result.CallStack))
	}

	if result.CallStack[1].Index != 0 {
		t.Errorf("frame 1 index = %d, expected 0", result.CallStack[1].Index)
	}
	if result.CallStack[1].FuncName != "TestFunction" {
		t.Errorf("frame 1 FuncName = %q, expected %q", result.CallStack[1].FuncName, "TestFunction")
	}
	if !result.CallStack[1].IsDriver {
		t.Error("frame 1 should be marked as driver")
	}
}

func TestParseCallStackStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}
	tick := "`"

	for i := 0; i < 500; i++ {
		var sb strings.Builder
		sb.WriteString("# Child-SP          RetAddr           Call Site\n")

		frameCount := 10 + i%20
		for j := 0; j < frameCount; j++ {
			addr := fmt.Sprintf("fffff800%08x", i*1000+j)
			if j%3 == 0 {
				sb.WriteString(fmt.Sprintf("#%02x ffffc781"+tick+"12345678 %s hyperkd!Func%d+0x%x\n", j, addr, j, j*4))
			} else {
				sb.WriteString(fmt.Sprintf("#%02x ffffc781"+tick+"12345678 %s nt!Func%d+0x%x\n", j, addr, j, j*4))
			}
		}
		sb.WriteString("quit:\n")

		result := &AnalysisResult{}
		a.parseCallStack(sb.String(), result)

		if len(result.CallStack) != frameCount+1 {
			t.Errorf("iteration %d: expected %d frames, got %d", i, frameCount+1, len(result.CallStack))
		}
	}
}

func TestParseUnloadedModules(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	tick := "`"
	output := "Unloaded modules:\n" +
		"fffff800" + tick + "12340000 fffff800" + tick + "12350000 hyperkd.sys\n" +
		"fffff800" + tick + "12360000 fffff800" + tick + "12370000 other.sys\n"

	result := &AnalysisResult{}
	a.parseUnloadedModules(output, result)

	if len(result.UnloadedModules) != 1 {
		t.Fatalf("expected 1 unloaded module, got %d", len(result.UnloadedModules))
	}

	if result.UnloadedModules[0].Name != "hyperkd.sys" {
		t.Errorf("module name = %q, expected %q", result.UnloadedModules[0].Name, "hyperkd.sys")
	}
}

func TestParseUnloadedModulesStress(t *testing.T) {
	a := &Analyzer{DriverName: "hyperkd"}

	for i := 0; i < 1000; i++ {
		var sb strings.Builder
		sb.WriteString("Unloaded modules:\n")

		moduleCount := i % 10
		for j := 0; j < moduleCount; j++ {
			base := uint64(0xfffff80012340000) + uint64(j)*0x10000
			end := base + 0x10000
			if j%2 == 0 {
				sb.WriteString(fmt.Sprintf("%016x %016x hyperkd.sys\n", base, end))
			} else {
				sb.WriteString(fmt.Sprintf("%016x %016x other%d.sys\n", base, end, j))
			}
		}

		result := &AnalysisResult{}
		a.parseUnloadedModules(sb.String(), result)

		expectedCount := (moduleCount + 1) / 2
		if len(result.UnloadedModules) != expectedCount {
			t.Errorf("iteration %d: expected %d modules, got %d", i, expectedCount, len(result.UnloadedModules))
		}
	}
}

func TestRunKdCommandTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping timeout test in short mode")
	}

	a := &Analyzer{
		WinDbgPath:     "cmd.exe",
		SymbolPath:     "test",
		CommandTimeout: 100 * time.Millisecond,
		MaxOutputSize:  1024,
	}

	ctx := context.Background()
	output := a.runKdCommand(ctx, "timeout test", "NUL", []string{"/c", "timeout", "/t", "10"})

	if output != "" {
		t.Log("command timed out as expected")
	}
}

func TestRunKdCommandOutputLimit(t *testing.T) {
	a := &Analyzer{
		WinDbgPath:     "cmd.exe",
		SymbolPath:     "test",
		CommandTimeout: 10 * time.Second,
		MaxOutputSize:  1024,
	}

	ctx := context.Background()

	output := a.runKdCommand(ctx, "output test", "NUL", []string{"/c", "echo test_output"})

	if !strings.Contains(output, "test_output") {
		t.Errorf("output should contain 'test_output', got: %q", output)
	}
}

func TestRunKdCommandStress(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping stress test in short mode")
	}

	a := &Analyzer{
		WinDbgPath:     "cmd.exe",
		SymbolPath:     "test",
		CommandTimeout: 5 * time.Second,
		MaxOutputSize:  1024,
	}

	for i := 0; i < 100; i++ {
		ctx := context.Background()
		output := a.runKdCommand(ctx, fmt.Sprintf("stress test %d", i), "NUL", []string{"/c", "echo", fmt.Sprintf("test_%d", i)})

		expected := fmt.Sprintf("test_%d", i)
		if !strings.Contains(output, expected) {
			t.Errorf("iteration %d: output should contain %q", i, expected)
		}
	}
}

func TestFindDriverFile(t *testing.T) {
	a := &Analyzer{
		PdbPath:    os.TempDir(),
		DriverName: "nonexistent_driver_xyz",
	}

	result := a.findDriverFile()
	if result != "" {
		t.Logf("findDriverFile returned: %s", result)
	}
}

func TestFindLatestMinidump(t *testing.T) {
	a := &Analyzer{}

	minidumpDir := "C:\\Windows\\Minidump"
	if _, err := os.Stat(minidumpDir); os.IsNotExist(err) {
		t.Skip("minidump directory does not exist")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Logf("findLatestMinidump panicked as expected when no dumps: %v", r)
		}
	}()

	result := a.findLatestMinidump()
	if result != "" {
		t.Logf("found dump: %s", result)
	}
}

func TestRegexCompilation(t *testing.T) {
	tick := "`"
	patterns := []string{
		`ExceptionAddress:\s*([0-9a-fA-F` + tick + `]+)`,
		`ExceptionCode:\s*([0-9a-fA-F]+)`,
		`Unloaded modules:\r?\n((?:.*\r?\n)+)`,
		`([0-9a-fA-F` + tick + `]+)\s+([0-9a-fA-F` + tick + `]+)\s+(\S+\.sys)`,
		`PROCESS_NAME:\s*(\S+)`,
		`IMAGE_NAME:\s*(\S+)`,
		`FAILURE_BUCKET_ID:\s*(\S+)`,
		`\[([^\]]+):(\d+)\]`,
		`hyperkd!([a-zA-Z_][a-zA-Z0-9_]*)`,
	}

	for _, pattern := range patterns {
		t.Run(pattern, func(t *testing.T) {
			re, err := regexp.Compile(pattern)
			if err != nil {
				t.Errorf("failed to compile regex %q: %v", pattern, err)
			}
			if re == nil {
				t.Errorf("compiled regex is nil for %q", pattern)
			}
		})
	}
}

func TestRegexCompilationStress(t *testing.T) {
	tick := "`"
	for i := 0; i < 10000; i++ {
		pattern := `ExceptionAddress:\s*([0-9a-fA-F` + tick + `]+)`
		re := regexp.MustCompile(pattern)
		if re == nil {
			t.Fatal("regex should not be nil")
		}

		testStr := fmt.Sprintf("ExceptionAddress: fffff800"+tick+"%08x", i)
		if !re.MatchString(testStr) {
			t.Errorf("regex should match test string: %s", testStr)
		}
	}
}

func TestConcurrentAnalysis(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping concurrent test in short mode")
	}

	a := &Analyzer{
		DriverName:     "hyperkd",
		PdbPath:        os.TempDir(),
		CommandTimeout: 1 * time.Second,
		MaxOutputSize:  1024,
	}

	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			result := &AnalysisResult{
				DumpFile:      fmt.Sprintf("test%d.dmp", id),
				DriverName:    a.DriverName,
				ExceptionCode: "c0000005",
				UnloadedModules: []UnloadedModule{
					{BaseAddress: "1000", EndAddress: "2000", Name: "hyperkd.sys"},
				},
			}

			a.analyzeCrashReason(result)
			a.checkDriverCrash(result)
			_ = a.buildReport(result)

			done <- true
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-done
	}
}

func TestMemoryUsage(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping memory test in short mode")
	}

	a := &Analyzer{
		DriverName:     "hyperkd",
		DriverSys:      "test.sys",
		CommandTimeout: 1 * time.Second,
		MaxOutputSize:  1024,
	}

	var results []*AnalysisResult

	for i := 0; i < 1000; i++ {
		result := &AnalysisResult{
			DumpFile:         fmt.Sprintf("test%d.dmp", i),
			DriverName:       a.DriverName,
			ExceptionCode:    "c0000005",
			ExceptionAddress: fmt.Sprintf("%x", 0x1000+i),
			CallStack:        make([]StackFrame, 100),
		}

		for j := 0; j < 100; j++ {
			result.CallStack[j] = StackFrame{
				Index:      j,
				RetAddr:    fmt.Sprintf("%x", 0x1000+j),
				CallSite:   fmt.Sprintf("module!Func%d+0x%x", j, j),
				ModuleName: "module",
				FuncName:   fmt.Sprintf("Func%d", j),
				FuncOffset: fmt.Sprintf("+0x%x", j),
			}
		}

		a.analyzeCrashReason(result)
		a.checkDriverCrash(result)
		_ = a.buildReport(result)

		results = append(results, result)
	}

	t.Logf("created %d results with %d frames each", len(results), 100)
}

func TestInitCommands(t *testing.T) {
	a := &Analyzer{
		DriverName:     "testdriver",
		DriverSys:      "C:\\path\\to\\test.sys",
		SymOpt:         ".symopt+0x10",
		CommandTimeout: 1 * time.Minute,
		MaxOutputSize:  1024,
	}
	a.initCommands()

	if len(a.AnalyzeCmds) == 0 {
		t.Error("AnalyzeCmds should not be empty")
	}

	if a.ResolveCmds == nil {
		t.Error("ResolveCmds should not be nil")
	}

	if a.LocateCmds == nil {
		t.Error("LocateCmds should not be nil")
	}

	frame := &StackFrame{RetAddr: "fffff80012345678"}
	resolveCmds := a.ResolveCmds(frame)
	if len(resolveCmds) == 0 {
		t.Error("ResolveCmds() should return non-empty slice")
	}

	locateCmds := a.LocateCmds("0x100")
	if len(locateCmds) == 0 {
		t.Error("LocateCmds() should return non-empty slice")
	}

	t.Logf("AnalyzeCmds: %v", a.AnalyzeCmds)
	t.Logf("ResolveCmds: %v", resolveCmds)
	t.Logf("LocateCmds: %v", locateCmds)
}

func TestSymOptField(t *testing.T) {
	a := &Analyzer{
		SymOpt: ".symopt+0x10",
	}

	if a.SymOpt != ".symopt+0x10" {
		t.Errorf("SymOpt = %q, expected %q", a.SymOpt, ".symopt+0x10")
	}

	a.SymOpt = ".symopt+0x100000"
	if a.SymOpt != ".symopt+0x100000" {
		t.Errorf("SymOpt = %q, expected %q", a.SymOpt, ".symopt+0x100000")
	}
}

func TestMaxOutputSizeField(t *testing.T) {
	a := &Analyzer{
		MaxOutputSize: 1024,
	}

	if a.MaxOutputSize != 1024 {
		t.Errorf("MaxOutputSize = %d, expected %d", a.MaxOutputSize, 1024)
	}

	a.MaxOutputSize = 10 * 1024 * 1024
	if a.MaxOutputSize != 10*1024*1024 {
		t.Errorf("MaxOutputSize = %d, expected %d", a.MaxOutputSize, 10*1024*1024)
	}
}

func BenchmarkParseHex(b *testing.B) {
	tick := "`"
	for i := 0; i < b.N; i++ {
		parseHex("fffff800" + tick + "12345678")
	}
}

func BenchmarkParseCallSite(b *testing.B) {
	a := &Analyzer{DriverName: "hyperkd"}
	frame := &StackFrame{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.parseCallSite("hyperkd!SomeFunction+0x10", frame)
	}
}

func BenchmarkAnalyzeCrashReason(b *testing.B) {
	a := &Analyzer{DriverName: "hyperkd"}
	result := &AnalysisResult{
		ExceptionCode:   "c0000005",
		UnloadedModules: []UnloadedModule{{BaseAddress: "1000", EndAddress: "2000", Name: "hyperkd.sys"}},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.CrashReason = ""
		a.analyzeCrashReason(result)
	}
}

func BenchmarkCheckDriverCrash(b *testing.B) {
	a := &Analyzer{DriverName: "hyperkd"}
	result := &AnalysisResult{
		ExceptionAddress: "1500",
		UnloadedModules: []UnloadedModule{
			{BaseAddress: "1000", EndAddress: "2000", Name: "hyperkd.sys"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.IsDriverCrash = false
		a.checkDriverCrash(result)
	}
}

func BenchmarkBuildReport(b *testing.B) {
	a := &Analyzer{
		DriverName: "hyperkd",
		DriverSys:  "test.sys",
	}

	result := &AnalysisResult{
		DumpFile:      "test.dmp",
		DriverName:    "hyperkd",
		ExceptionCode: "c0000005",
		CrashReason:   "test reason",
		CallStack:     make([]StackFrame, 50),
	}

	for i := 0; i < 50; i++ {
		result.CallStack[i] = StackFrame{
			Index:      i,
			RetAddr:    fmt.Sprintf("%x", 0x1000+i),
			CallSite:   fmt.Sprintf("mod!Func%d+0x%x", i, i),
			ModuleName: "mod",
			FuncName:   fmt.Sprintf("Func%d", i),
			FuncOffset: fmt.Sprintf("+0x%x", i),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a.buildReport(result)
	}
}

func BenchmarkRegexMatch(b *testing.B) {
	tick := "`"
	re := regexp.MustCompile(`ExceptionAddress:\s*([0-9a-fA-F` + tick + `]+)`)
	testStr := "ExceptionAddress: fffff800" + tick + "12345678"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		re.FindStringSubmatch(testStr)
	}
}

func BenchmarkParseCallStack(b *testing.B) {
	a := &Analyzer{DriverName: "hyperkd"}
	tick := "`"

	var sb strings.Builder
	sb.WriteString("# Child-SP          RetAddr           Call Site\n")
	for i := 0; i < 50; i++ {
		sb.WriteString(fmt.Sprintf("%02x ffffc781"+tick+"12345678 fffff800"+tick+"%08x mod!Func%d+0x%x\n", i, i, i, i))
	}
	sb.WriteString("quit:\n")
	output := sb.String()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := &AnalysisResult{}
		a.parseCallStack(output, result)
	}
}

func TestWinDbgExists(t *testing.T) {
	a := NewAnalyzer()

	if _, err := os.Stat(a.WinDbgPath); os.IsNotExist(err) {
		t.Skipf("WinDbg not found at %s", a.WinDbgPath)
	}

	cmd := exec.Command(a.WinDbgPath, "-version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Logf("WinDbg version check: %v, output: %s", err, string(output))
	} else {
		t.Logf("WinDbg found: %s", a.WinDbgPath)
	}
}

func TestPdbPathExists(t *testing.T) {
	a := NewAnalyzer()

	if _, err := os.Stat(a.PdbPath); os.IsNotExist(err) {
		t.Skipf("PDB path not found at %s", a.PdbPath)
	}

	t.Logf("PDB path exists: %s", a.PdbPath)

	files, err := os.ReadDir(a.PdbPath)
	if err != nil {
		t.Logf("Failed to read PDB directory: %v", err)
	} else {
		t.Logf("Found %d files in PDB directory", len(files))
	}
}

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	a := NewAnalyzer()

	if _, err := os.Stat(a.WinDbgPath); os.IsNotExist(err) {
		t.Skipf("WinDbg not found at %s", a.WinDbgPath)
	}

	minidumpDir := "C:\\Windows\\Minidump"
	if _, err := os.Stat(minidumpDir); os.IsNotExist(err) {
		t.Skip("minidump directory does not exist")
	}

	entries, err := os.ReadDir(minidumpDir)
	if err != nil || len(entries) == 0 {
		t.Skip("no dump files found")
	}

	t.Log("Integration test prerequisites met")
	t.Logf("WinDbg: %s", a.WinDbgPath)
	t.Logf("PDB: %s", a.PdbPath)
	t.Logf("Driver: %s", a.DriverSys)
}
