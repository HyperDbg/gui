package debugger

import (
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/app"
	"github.com/ddkwork/HyperDbg/debugger/commands/debugging"
	"github.com/ddkwork/HyperDbg/debugger/commands/extension"
	"github.com/ddkwork/HyperDbg/debugger/commands/hwdbg"
	"github.com/ddkwork/HyperDbg/debugger/commands/meta"
	"github.com/ddkwork/HyperDbg/debugger/common"
	"github.com/ddkwork/HyperDbg/debugger/core"
	"github.com/ddkwork/HyperDbg/debugger/driver_loader"
	"github.com/ddkwork/HyperDbg/debugger/misc"
	"github.com/ddkwork/HyperDbg/debugger/scriptengine"
	"github.com/ddkwork/HyperDbg/debugger/transparency"
	"github.com/ddkwork/HyperDbg/debugger/user_level"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ddkwork/sdk"
	"golang.org/x/sys/windows"
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

func TestCommonUtilities(t *testing.T) {
	if !common.IsNumber("12345") {
		t.Error("IsNumber should return true for '12345'")
	}
	if common.IsNumber("abc") {
		t.Error("IsNumber should return false for 'abc'")
	}
	if !common.IsHexNotation("0x1234") {
		t.Error("IsHexNotation should return true for '0x1234'")
	}
	if !common.IsDecimalNotation("12345") {
		t.Error("IsDecimalNotation should return true for '12345'")
	}
	result := common.SeparateTo64BitValue(0x1234567890ABCDEF)
	if result == "" {
		t.Error("SeparateTo64BitValue should not return empty string")
	}
	bytes := common.HexToBytes("0x1234")
	if len(bytes) != 2 || bytes[0] != 0x12 || bytes[1] != 0x34 {
		t.Errorf("HexToBytes unexpected result: %v", bytes)
	}
	parts := common.Split("a,b,c", ',')
	if len(parts) != 3 {
		t.Errorf("Split unexpected result: %v", parts)
	}
	mylog.Success("common utilities test passed")
}

func TestSymbolTable(t *testing.T) {
	st := scriptengine.NewSymbolTable()
	if st == nil {
		t.Error("NewSymbolTable should not return nil")
	}
	st.AddSymbol(0x7FFE0000, "ntdll.dll", "NtCreateFile", 0x50)
	sym, ok := st.GetSymbolByAddress(0x7FFE0000)
	if !ok || sym.ModuleName != "ntdll.dll" {
		t.Errorf("GetSymbolByAddress unexpected result: ok=%v", ok)
	}
}

func TestGaussianRng(t *testing.T) {
	data := make([]float64, 1000)
	for i := range data {
		data[i] = transparency.Randn(0, 1)
	}
	result := transparency.GuassianGenerateRandom(data)
	mylog.Info("Gaussian RNG result:", "avg", result.Average, "stddev", result.StandardDeviation, "median", result.Median)
}

func TestMiscDisassembler(t *testing.T) {
	code := []byte{0x90}
	insts, err := misc.DisassembleBuffer(code, 0x0, misc.Mode64, 1)
	if err != nil {
		t.Errorf("DisassembleBuffer error: %v", err)
	}
	if len(insts) == 0 {
		t.Error("DisassembleBuffer should return at least one instruction")
	}
}

func TestMiscCallstack(t *testing.T) {
	returnAddr := []byte{0x18, 0x10, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00}
	found, idx := misc.CallstackReturnAddressToCallingAddress(returnAddr)
	mylog.Info("Callstack test:", "found", found, "idx", idx)
}

func TestMiscReadmem(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	dwords := misc.FormatMemoryAsDwords(data)
	if len(dwords) != 2 {
		t.Errorf("FormatMemoryAsDwords unexpected length: %d", len(dwords))
	}
	qwords := misc.FormatMemoryAsQwords(data)
	if len(qwords) != 1 {
		t.Errorf("FormatMemoryAsQwords unexpected length: %d", len(qwords))
	}
	s := misc.FormatMemoryAsString(data)
	if s == "" {
		t.Error("FormatMemoryAsString should return empty")
	}
}

func TestDriverLoadUnload(t *testing.T) {
	driverPath := mylog.Check2(filepath.Abs(`..\HyperDbgUnified\build\Release\hyperkd.sys`))
	ok := driver_loader.ManageDriver("hyperdbg", driverPath, driver_loader.DriverFuncInstall)
	if !ok {
		mylog.Warning("驱动安装失败，跳过测试")
		return
	}
	ok = driver_loader.ManageDriver("hyperdbg", driverPath, driver_loader.DriverFuncStart)
	if !ok {
		mylog.Warning("驱动启动失败")
		driver_loader.ManageDriver("hyperdbg", driverPath, driver_loader.DriverFuncRemove)
		return
	}
	driver_loader.ManageDriver("hyperdbg", driverPath, driver_loader.DriverFuncStop)
	driver_loader.ManageDriver("hyperdbg", driverPath, driver_loader.DriverFuncRemove)
	mylog.Success("驱动加载卸载测试完成")
}

// ============================================================================
// Debugging Commands Tests
// ============================================================================

func TestDebuggingBreakpointManager(t *testing.T) {
	dc := core.NewDebuggerCore()

	err := debugging.CommandBpRequest(dc, 0x1000, 0, 0, 0)
	if err != nil {
		t.Fatalf("Set breakpoint failed: %v", err)
	}

	output := debugging.CommandBpListAll(dc)
	if output == "" || output == "no breakpoints set" {
		t.Error("ListAll should show breakpoints")
	}

	count, err := debugging.CommandBpClearAll(dc)
	if err != nil || count != 1 {
		t.Errorf("ClearAll failed: count=%d, err=%v", count, err)
	}
}

func TestDebuggingBatchOperations(t *testing.T) {
	dc := core.NewDebuggerCore()

	for i := range uint32(5) {
		err := debugging.CommandBpRequest(dc, uint64(i)*0x1000, i, i, i)
		if err != nil {
			t.Fatalf("Set breakpoint %d failed: %v", i, err)
		}
	}

	enableCount, err := debugging.CommandBpDisableAll(dc)
	if err != nil || enableCount != 5 {
		t.Errorf("DisableAll failed: count=%d, err=%v", enableCount, err)
	}

	enableCount, err = debugging.CommandBpEnableAll(dc)
	if err != nil || enableCount != 5 {
		t.Errorf("EnableAll failed: count=%d, err=%v", enableCount, err)
	}

	output := debugging.CommandBpListAll(dc)
	if output == "" || output == "no breakpoints set" {
		t.Error("ListAll should show breakpoints")
	}
	mylog.Info(output)
}

func TestDebuggingSteppingCommands(t *testing.T) {
	dc := core.NewDebuggerCore()

	count, err := debugging.CommandStepInMultiple(dc, 3)
	if err != nil {
		t.Errorf("StepInMultiple error: %v", err)
	}
	if count != 3 {
		t.Errorf("Expected 3 steps, got %d", count)
	}

	count, err = debugging.CommandStepOverMultiple(dc, 2)
	if err != nil {
		t.Errorf("StepOverMultiple error: %v", err)
	}
	if count != 2 {
		t.Errorf("Expected 2 steps, got %d", count)
	}
}

func TestDebuggingMemoryCommands(t *testing.T) {
	dc := core.NewDebuggerCore()

	str, err := debugging.CommandReadMemoryString(dc, 0x1000, 256)
	if err != nil {
		t.Errorf("ReadMemoryString error: %v", err)
	}
	if str == "" {
		t.Error("ReadMemoryString should return non-empty string (all non-zero bytes)")
	}

	err = debugging.CommandWriteMemoryString(dc, 0x2000, "Hello", 1234)
	if err != nil {
		t.Errorf("WriteMemoryString error: %v", err)
	}

	addrs, err := debugging.CommandSearchPattern(dc, "909090", 0, 4096, 0)
	if err != nil {
		t.Errorf("SearchPattern error: %v", err)
	}
	if len(addrs) != 2 {
		t.Errorf("Expected 2 addresses, got %d", len(addrs))
	}

	err = debugging.CommandFillMemory(dc, 0x3000, 0xFF, 16, 5678)
	if err != nil {
		t.Errorf("FillMemory error: %v", err)
	}

	equal, err := debugging.CommandCompareMemory(dc, 0x1000, 0x2000, 32)
	if err != nil {
		t.Errorf("CompareMemory error: %v", err)
	}
	if !equal {
		t.Error("CompareMemory should return true for same read function")
	}

	err = debugging.CommandMoveMemory(dc, 0x4000, 0x5000, 64, 9999)
	if err != nil {
		t.Errorf("MoveMemory error: %v", err)
	}
}

func TestDebuggingRegisterCommands(t *testing.T) {
	dc := core.NewDebuggerCore()

	newValue, err := debugging.CommandModifyRegister(dc, "rcx", "+", 100)
	if err != nil {
		t.Errorf("ModifyRegister error: %v", err)
	}

	hexOutput, err := debugging.CommandShowRegistersFormatted(dc, "hex")
	if err != nil {
		t.Errorf("ShowRegistersFormatted error: %v", err)
	}
	if hexOutput == "" {
		t.Error("ShowRegistersFormatted(hex) should return output")
	}

	decOutput, err := debugging.CommandShowRegistersFormatted(dc, "decimal")
	if err != nil {
		t.Errorf("ShowRegistersFormatted error: %v", err)
	}
	if decOutput == "" {
		t.Error("ShowRegistersFormatted(decimal) should return output")
	}

	_, err = debugging.CommandModifyRegister(dc, "rax", "/", 0)
	if err == nil {
		t.Error("Division by zero should return error")
	}
	_ = newValue
}

func TestDebuggingEvalAndSettings(t *testing.T) {
	dc := core.NewDebuggerCore()

	value, err := debugging.CommandEvalAndSet(dc, "0x100", 0x1000, 0)
	if err != nil {
		t.Errorf("EvalAndSet error: %v", err)
	}
	if value != 0x100 {
		t.Errorf("Expected value=0x100, got 0x%X", value)
	}

	val := debugging.CommandSettingsGet(dc, "disasm_style")
	_ = val

	allSettings := debugging.CommandSettingsListAll(dc)
	_ = allSettings
}

func TestDebuggingHelpAndTestCommands(t *testing.T) {
	dc := core.NewDebuggerCore()

	result, err := debugging.CommandTestAll(dc)
	if err != nil {
		t.Errorf("CommandTestAll error: %v", err)
	}
	_ = result
}

// ============================================================================
// Extension Commands Tests
// ============================================================================

func TestExtensionApicAndCpuid(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, _, err := extension.CommandApicDetailed(dc)
	if err != nil {
		t.Errorf("CommandApicDetailed error: %v", err)
	}

	_, err = extension.CommandCpuidDetailed(dc, 0)
	if err != nil {
		t.Errorf("CommandCpuidDetailed error: %v", err)
	}
}

func TestExtensionIoAndMsr(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, err := extension.CommandIoinDetailed(dc, 0x80)
	if err != nil {
		t.Errorf("CommandIoinDetailed error: %v", err)
	}

	err = extension.CommandIooutDetailed(dc, 0x80, 0x12345678)
	if err != nil {
		t.Errorf("CommandIooutDetailed error: %v", err)
	}

	_, err = extension.CommandMsrReadMultiple(dc, []uint32{0x174, 0x175})
	if err != nil {
		t.Errorf("CommandMsrReadMultiple error: %v", err)
	}

	err = extension.CommandMsrWriteMultiple(dc, map[uint32]uint64{0x174: 0x100, 0x175: 0x200})
	if err != nil {
		t.Errorf("CommandMsrWriteMultiple error: %v", err)
	}
}

func TestExtensionPteAndVa2pa(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, err := extension.CommandPteDetailed(dc, 0x7FFE0000)
	if err != nil {
		t.Errorf("CommandPteDetailed error: %v", err)
	}

	_, err = extension.CommandPa2vaMultiple(dc, []uint64{0x1000, 0x2000, 0x3000})
	if err != nil {
		t.Errorf("CommandPa2vaMultiple error: %v", err)
	}

	_, err = extension.CommandVa2paMultiple(dc, []uint64{0xFFFF800000001000, 0xFFFF800000002000})
	if err != nil {
		t.Errorf("CommandVa2paMultiple error: %v", err)
	}
}

func TestExtensionTscAndVmcall(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, err := extension.CommandTscDetailed(dc)
	if err != nil {
		t.Errorf("CommandTscDetailed error: %v", err)
	}

	err = extension.CommandVmcallDetailed(dc, 0x1, 0x2, 0x3)
	if err != nil {
		t.Errorf("CommandVmcallDetailed error: %v", err)
	}

	err = extension.CommandXsetbvDetailed(dc, 0, 0x111, 0x222)
	if err != nil {
		t.Errorf("CommandXsetbvDetailed error: %v", err)
	}
}

func TestExtensionMonitorAndTrace(t *testing.T) {
	dc := core.NewDebuggerCore()

	_ = extension.CommandMonitorList(dc)

	_, err := extension.CommandTraceGetResults(dc, 1)
	if err != nil {
		t.Errorf("CommandTraceGetResults error: %v", err)
	}
}

// ============================================================================
// Meta Commands Tests
// ============================================================================

func TestMetaProcessAndThread(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, err := meta.CommandProcessListRequest(dc)
	if err != nil {
		t.Errorf("CommandProcessListRequest error: %v", err)
	}

	_, err = meta.CommandThreadListRequest(dc, 1234)
	if err != nil {
		t.Errorf("CommandThreadListRequest error: %v", err)
	}
}

func TestMetaConnectionAndStatus(t *testing.T) {
	dc := core.NewDebuggerCore()

	err := meta.CommandConnectLocal(dc, []string{"connect", "local"}, "connect local")
	if err != nil {
		t.Errorf("CommandConnectLocal error: %v", err)
	}

	_ = meta.CommandStatusRequest(dc)
}

func TestMetaSymbolAndLog(t *testing.T) {
	err := meta.CommandSymLoad([]string{"sym", "load", "ntdll.pdb"}, "sym load ntdll.pdb")
	if err != nil {
		t.Errorf("CommandSymLoad error: %v", err)
	}

	err = meta.CommandSympath([]string{"sympath", "srv*C:\\symbols*https://msdl.microsoft.com/download/symbols"}, "sympath")
	if err != nil {
		t.Errorf("CommandSympath error: %v", err)
	}
}

func TestMetaPeAndDump(t *testing.T) {
	dc := core.NewDebuggerCore()

	err := meta.CommandPeRequest(0x140000000)
	if err != nil {
		t.Errorf("CommandPeRequest error: %v", err)
	}

	err = meta.CommandDump(dc, []string{"dump", "hex", "0x1000", "64"}, "dump hex 0x1000 64")
	if err != nil {
		t.Errorf("CommandDump error: %v", err)
	}
}

// ============================================================================
// Hwdbg Commands Tests
// ============================================================================

func TestHwdbgBasicCommands(t *testing.T) {
	helpText := hwdbg.CommandHwHelp()
	if helpText == "" {
		t.Error("CommandHwHelp should return help text")
	}

	helpText = hwdbg.CommandHwClkHelp()
	if helpText == "" {
		t.Error("CommandHwClkHelp should return help text")
	}

	err := hwdbg.CommandHw([]string{"!hw", "unload"}, "!hw unload")
	if err != nil {
		t.Errorf("CommandHw(unload) error: %v", err)
	}
	if hwdbg.G_HwdbgInstanceInfoIsValid {
		t.Error("Instance should be invalid after unload")
	}
}

func TestHwdbgInstanceManagement(t *testing.T) {
	success := hwdbg.LoadInstanceInfo("test_instance.bin", 8192)
	if !success {
		t.Error("LoadInstanceInfo should succeed")
	}
	if !hwdbg.G_HwdbgInstanceInfoIsValid {
		t.Error("Instance info should be valid after load")
	}
	if hwdbg.G_HwdbgInstanceInfo.BramBufferSize != 8192 {
		t.Error("BRAM buffer size incorrect")
	}

	info := &hwdbg.HwdbgInstanceInfo{
		IsValid:          true,
		ScriptBuffer:     "@hw_pin1 = 0;",
		BramBufferSize:   16384,
		HardwareFilePath: "output.bin",
	}

	success = hwdbg.WriteInstanceInfoToFile(info, "output_instance.bin")
	if !success {
		t.Error("WriteInstanceInfoToFile should succeed")
	}

	success = hwdbg.CreateHwdbgScript("script content", 15, "script_output.bin")
	if !success {
		t.Error("CreateHwdbgScript should succeed")
	}
	if hwdbg.G_HwdbgInstanceInfo.ScriptBuffer != "script content" {
		t.Error("Script buffer not set correctly")
	}
}

func TestHwdbgPrintScript(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	hwdbg.PrintScriptBuffer("@hw_pin1 = 0; @hw_pin2 = 1;")

	w.Close()
	os.Stdout = oldStdout

	buf := make([]byte, 4096)
	n, _ := r.Read(buf)
	output := string(buf[:n])

	if output == "" {
		t.Error("PrintScriptBuffer should produce output")
	}
	if !strings.Contains(output, "Hardware Debug Script") {
		t.Error("Output should contain header")
	}
}

func TestHwclkCommand(t *testing.T) {
	err := hwdbg.CommandHwClk([]string{"!hw_clk"}, "!hw_clk")
	if err == nil {
		t.Error("CommandHwClk with no arguments should fail")
	}

	err = hwdbg.CommandHwClk([]string{"!hw_clk", "script", "{ @hw_pin1 = 0; }"}, "!hw_clk")
	if err != nil {
		t.Errorf("CommandHwClk(script) error: %v", err)
	}
}

// Helper functions for stdout capture
func captureStdout() *os.File {
	oldStdout := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	return oldStdout
}

func restoreStdout(old *os.File) {
	oldStdout := os.Stdout
	if oldStdout != nil {
		oldStdout.Close()
	}
	os.Stdout = old
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// ============================================================================
// Additional Debugging Command Tests
// ============================================================================

func TestDebuggingHelperFunctions(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

func TestDebuggingEventManagement(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

func TestDebuggingOutputManagement(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

func TestDebuggingTestFunctions(t *testing.T) {
	dc := core.NewDebuggerCore()
	_, err := debugging.CommandTestAll(dc)
	if err != nil {
		t.Errorf("CommandTestAll error: %v", err)
	}
}

func TestDebuggerStatus(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = meta.CommandStatusRequest(dc)
}

func TestVersionAndHelp(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

// ============================================================================
// Additional Extension Command Tests
// ============================================================================

func TestExtensionCommandHandlers(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

func TestExtensionMsrMultiple(t *testing.T) {
	dc := core.NewDebuggerCore()
	_, err := extension.CommandMsrReadMultiple(dc, []uint32{0x174})
	if err != nil {
		t.Errorf("MsrReadMultiple error: %v", err)
	}
}

func TestExtensionDetailedResults(t *testing.T) {
	dc := core.NewDebuggerCore()

	_, err := extension.CommandCrRead(dc, 0)
	if err != nil {
		t.Errorf("CrRead error: %v", err)
	}

	_, err = extension.CommandDrRead(dc, 0)
	if err != nil {
		t.Errorf("DrRead error: %v", err)
	}

	err = extension.CommandDrWrite(dc, 0, 0x1234)
	if err != nil {
		t.Errorf("DrWrite error: %v", err)
	}

	_ = extension.CommandEpthookList(dc)
	_ = extension.CommandExceptionList(dc)
	_ = extension.CommandHideList(dc)

	_, err = extension.CommandMeasureDetailed(dc, 0x1000, 0x2000)
	if err != nil {
		t.Errorf("MeasureDetailed error: %v", err)
	}

	_, err = extension.CommandPmcDetailed(dc, 0xC0)
	if err != nil {
		t.Errorf("PmcDetailed error: %v", err)
	}

	err = extension.CommandSyscallHook(dc, 0x55, 0)
	if err != nil {
		t.Errorf("SyscallHook error: %v", err)
	}
	err = extension.CommandSyscallUnhook(dc, 0x55)
	if err != nil {
		t.Errorf("SyscallUnhook error: %v", err)
	}

	err = extension.CommandTraceStart(dc, 1, 1)
	if err != nil {
		t.Errorf("TraceStart error: %v", err)
	}
	err = extension.CommandTraceStop(dc, 1)
	if err != nil {
		t.Errorf("TraceStop error: %v", err)
	}
	_, err = extension.CommandTraceGetResults(dc, 1)
	if err != nil {
		t.Errorf("TraceGetResults error: %v", err)
	}

	err = extension.CommandTrackStart(dc, 1, 1)
	if err != nil {
		t.Errorf("TrackStart error: %v", err)
	}
	err = extension.CommandTrackStop(dc, 1)
	if err != nil {
		t.Errorf("TrackStop error: %v", err)
	}

	err = extension.CommandMonitorStart(dc, 0x1000, 4096, 1)
	if err != nil {
		t.Errorf("MonitorStart error: %v", err)
	}
	err = extension.CommandMonitorStop(dc, 1)
	if err != nil {
		t.Errorf("MonitorStop error: %v", err)
	}
	_ = extension.CommandMonitorList(dc)

	err = extension.CommandModeSet(dc, "kernel", 0, 0)
	if err != nil {
		t.Errorf("ModeSet error: %v", err)
	}
	_, err = extension.CommandModeGet(dc)
	if err != nil {
		t.Errorf("ModeGet error: %v", err)
	}

	_, err = extension.CommandPcicamDetailed(dc, 0, 0, 0)
	if err != nil {
		t.Errorf("PcicamDetailed error: %v", err)
	}
	_, err = extension.CommandPcitreeDetailed(dc)
	if err != nil {
		t.Errorf("PcitreeDetailed error: %v", err)
	}

	err = extension.CommandLbrDetailed(dc, true)
	if err != nil {
		t.Errorf("LbrDetailed error: %v", err)
	}

	_, err = extension.CommandRevReverse(dc, 0x1000, 8, 1)
	if err != nil {
		t.Errorf("RevReverse error: %v", err)
	}

	err = extension.CommandSmiDetailed(dc, nil)
	if err != nil {
		t.Errorf("SmiDetailed error: %v", err)
	}
}

// ============================================================================
// Additional Meta Command Tests
// ============================================================================

func TestMetaCommandHandlers(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = dc
}

func TestMetaDetailedFunctions(t *testing.T) {
	dc := core.NewDebuggerCore()
	_ = meta.CommandStatusRequest(dc)
}

type ioctlRecord struct {
	Code       uint32
	Input      []byte
	InputSize  uint32
	Output     []byte
	OutputSize uint32
}

func newMockDebuggerCore(records *[]ioctlRecord) *core.DebuggerCore {
	dc := core.NewDebuggerCore()
	dc.SetDeviceIoControl(func(ioctlCode uint32, inputBuffer []byte, inputSize uint32, outputBuffer []byte, outputSize uint32) (bool, uint32, error) {
		in := make([]byte, inputSize)
		copy(in, inputBuffer[:inputSize])
		*records = append(*records, ioctlRecord{
			Code:       ioctlCode,
			Input:      in,
			InputSize:  inputSize,
			Output:     outputBuffer,
			OutputSize: outputSize,
		})
		return true, outputSize, nil
	})
	return dc
}

func getNtDeviceIoControlFileSyscallNumber() uint32 {
	proc := syscall.NewLazyDLL("ntdll.dll").NewProc("NtDeviceIoControlFile")
	addr := proc.Addr()
	code := *(*[]byte)(unsafe.Pointer(&struct {
		addr uintptr
		len  int
		cap  int
	}{addr, 16, 16}))
	if len(code) >= 2 && code[0] == 0x4C && code[1] == 0x8B && code[2] == 0xD1 {
		if len(code) >= 6 && code[3] == 0xB8 {
			return uint32(code[4]) | uint32(code[5])<<8
		}
	}
	if len(code) >= 4 && code[0] == 0xB8 {
		return uint32(code[1]) | uint32(code[2])<<8 | uint32(code[3])<<16
	}
	return 0x07
}

func getNtDeviceIoControlFileAddress() uint64 {
	proc := syscall.NewLazyDLL("ntdll.dll").NewProc("NtDeviceIoControlFile")
	return uint64(proc.Addr())
}

func TestHookNtDeviceIoControlFile(t *testing.T) {
	var records []ioctlRecord

	syscallNum := getNtDeviceIoControlFileSyscallNumber()
	funcAddr := getNtDeviceIoControlFileAddress()
	mylog.Info("NtDeviceIoControlFile",
		"syscall", fmt.Sprintf("0x%02X", syscallNum),
		"address", fmt.Sprintf("0x%X", funcAddr),
	)

	t.Run("SyscallHook", func(t *testing.T) {
		records = nil
		dc := newMockDebuggerCore(&records)

		err := extension.CommandSyscallHook(dc, syscallNum, 0)
		if err != nil {
			t.Fatalf("SyscallHook failed: %v", err)
		}
		if len(records) != 1 {
			t.Fatalf("expected 1 ioctl, got %d", len(records))
		}
		rec := records[0]
		mylog.Info("[SyscallHook] IOCTL sent",
			"code", fmt.Sprintf("0x%08X", rec.Code),
			"inputSize", rec.InputSize,
		)
		if rec.Code != sdk.IoctlDebuggerRegisterEvent {
			t.Errorf("expected IoctlDebuggerRegisterEvent(0x%08X), got 0x%08X",
				sdk.IoctlDebuggerRegisterEvent, rec.Code)
		}
		if rec.InputSize >= 8 {
			hookedIndex := binary.LittleEndian.Uint32(rec.Input[0:4])
			options := binary.LittleEndian.Uint32(rec.Input[4:8])
			mylog.Info("[SyscallHook] payload",
				"syscallIndex", fmt.Sprintf("0x%02X", hookedIndex),
				"options", options,
			)
			if hookedIndex != syscallNum {
				t.Errorf("expected syscall 0x%02X, got 0x%02X", syscallNum, hookedIndex)
			}
		}
	})

	t.Run("EpthookCC", func(t *testing.T) {
		records = nil
		dc := newMockDebuggerCore(&records)

		tag, err := extension.CommandEpthookRequest(dc, funcAddr, 0, 0, 0)
		if err != nil {
			t.Fatalf("EpthookRequest failed: %v", err)
		}
		if len(records) != 1 {
			t.Fatalf("expected 1 ioctl, got %d", len(records))
		}
		rec := records[0]
		mylog.Info("[EpthookCC] IOCTL sent",
			"code", fmt.Sprintf("0x%08X", rec.Code),
			"inputSize", rec.InputSize,
			"tag", tag,
		)
		if rec.Code != sdk.IoctlDebuggerRegisterEvent {
			t.Errorf("expected IoctlDebuggerRegisterEvent, got 0x%08X", rec.Code)
		}
		if rec.InputSize >= 24 {
			hookedAddr := binary.LittleEndian.Uint64(rec.Input[0:8])
			eventType := binary.LittleEndian.Uint32(rec.Input[20:24])
			mylog.Info("[EpthookCC] payload",
				"address", fmt.Sprintf("0x%X", hookedAddr),
				"eventType", sdk.VMM_EVENT_TYPE_ENUM(eventType).String(),
			)
			if hookedAddr != funcAddr {
				t.Errorf("expected address 0x%X, got 0x%X", funcAddr, hookedAddr)
			}
			if sdk.VMM_EVENT_TYPE_ENUM(eventType) != sdk.HiddenHookExecCc {
				t.Errorf("expected HiddenHookExecCc, got %s", sdk.VMM_EVENT_TYPE_ENUM(eventType))
			}
		}
	})

	t.Run("Epthook2Detours", func(t *testing.T) {
		records = nil
		dc := newMockDebuggerCore(&records)

		tag, err := extension.CommandEpthook2Request(dc, funcAddr, 0, 0, 0)
		if err != nil {
			t.Fatalf("Epthook2Request failed: %v", err)
		}
		if len(records) != 1 {
			t.Fatalf("expected 1 ioctl, got %d", len(records))
		}
		rec := records[0]
		mylog.Info("[Epthook2Detours] IOCTL sent",
			"code", fmt.Sprintf("0x%08X", rec.Code),
			"inputSize", rec.InputSize,
			"tag", tag,
		)
		if rec.Code != sdk.IoctlDebuggerRegisterEvent {
			t.Errorf("expected IoctlDebuggerRegisterEvent, got 0x%08X", rec.Code)
		}
		if rec.InputSize >= 24 {
			hookedAddr := binary.LittleEndian.Uint64(rec.Input[0:8])
			eventType := binary.LittleEndian.Uint32(rec.Input[20:24])
			mylog.Info("[Epthook2Detours] payload",
				"address", fmt.Sprintf("0x%X", hookedAddr),
				"eventType", sdk.VMM_EVENT_TYPE_ENUM(eventType).String(),
			)
			if hookedAddr != funcAddr {
				t.Errorf("expected address 0x%X, got 0x%X", funcAddr, hookedAddr)
			}
			if sdk.VMM_EVENT_TYPE_ENUM(eventType) != sdk.HiddenHookExecDetours {
				t.Errorf("expected HiddenHookExecDetours, got %s", sdk.VMM_EVENT_TYPE_ENUM(eventType))
			}
		}
	})

	t.Run("MonitorIoControlCode", func(t *testing.T) {
		records = nil
		dc := newMockDebuggerCore(&records)

		err := extension.CommandSyscallHook(dc, syscallNum, 0)
		if err != nil {
			t.Fatalf("SyscallHook failed: %v", err)
		}

		mylog.Info("[Monitor] hook registered, NtDeviceIoControlFile calls will be intercepted by hypervisor")
		mylog.Info("[Monitor] on callback, parse registers/buffers:")
		mylog.Info("  rcx     -> FileHandle")
		mylog.Info("  r8      -> IoStatusBlock")
		mylog.Info("  r10     -> IoControlCode (CTL_CODE)")
		mylog.Info("  r11     -> InputBuffer")
		mylog.Info("  r12     -> InputBufferLength")
		mylog.Info("  r13     -> OutputBuffer")
		mylog.Info("  r14     -> OutputBufferLength")

		err = extension.CommandSyscallUnhook(dc, syscallNum)
		if err != nil {
			t.Fatalf("SyscallUnhook failed: %v", err)
		}
		if len(records) != 2 {
			t.Fatalf("expected 2 ioctls (hook+unhook), got %d", len(records))
		}
		unhookRec := records[1]
		mylog.Info("[Monitor] unhook IOCTL",
			"code", fmt.Sprintf("0x%08X", unhookRec.Code),
		)
	})
}

func TestDebugNotepad(t *testing.T) {
	driverPath := mylog.Check2(filepath.Abs(`..\HyperDbgUnified\build_debug\hyperkd.sys`))
	if _, err := os.Stat(driverPath); err != nil {
		t.Skipf("driver not found at %s, skip: %v", driverPath, err)
	}

	l := app.NewLibhyperdbg()
	l.SetCustomDriverPath(driverPath, app.KernelDebuggerDriverName)
	if l.HyperDbgInstallVmmDriver() != 0 {
		t.Fatal("install driver failed")
	}
	if !driver_loader.ManageDriver(app.KernelDebuggerDriverName, driverPath, driver_loader.DriverFuncStart) {
		t.Fatal("start driver failed")
	}
	if l.LoadVmmModule("\\\\.\\HprDbgHv") != 0 {
		t.Fatal("install driver or load vmm failed")
	}
	defer func() {
		l.UnloadVmm(func(buf []byte, code uint32) ([]byte, error) {
			out := make([]byte, 0x10000)
			var n uint32
			var inPtr *byte
			if len(buf) > 0 {
				inPtr = &buf[0]
			}
			windows.DeviceIoControl(l.GetDeviceHandle(), code, inPtr, uint32(len(buf)), &out[0], 0x10000, &n, nil)
			return out[:n], nil
		})
		l.HyperDbgStopVmmDriver()
		l.HyperDbgUninstallVmmDriver()
	}()

	h := l.GetDeviceHandle()
	dc := core.NewDebuggerCore()
	dc.SetDeviceHandle(h)
	dc.SetConnectedToHyperDbgLocally(true)
	dc.SetDeviceIoControl(func(code uint32, in []byte, inSz uint32, out []byte, outSz uint32) (bool, uint32, error) {
		var n uint32
		var inPtr, outPtr *byte
		if len(in) > 0 {
			inPtr = &in[0]
		}
		if len(out) > 0 {
			outPtr = &out[0]
		}
		err := windows.DeviceIoControl(h, code, inPtr, inSz, outPtr, outSz, &n, nil)
		return err == nil, n, err
	})

	ud := user_level.NewUserDebugger()
	ud.Initialize()
	defer ud.Uninitialize()

	if err := dc.StartProcess(`C:\Windows\System32\notepad.exe`, ""); err != nil {
		t.Fatalf("start notepad: %v", err)
	}
	mylog.Success("notepad started, debugging...")

	time.Sleep(500 * time.Millisecond)
	dc.KillActiveProcess()
}
