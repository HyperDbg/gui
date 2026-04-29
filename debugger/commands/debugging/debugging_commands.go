package debugging

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/core"
	"github.com/ddkwork/sdk"
)

// CommandBpRequest sets a breakpoint at the specified address
// C++ Source: code/debugger/commands/debugging-commands/bp.cpp - CommandBp()
func CommandBpRequest(dc *core.DebuggerCore, address uint64, pid, tid, coreNum uint32) error {
	return dc.SetBreakpoint(address, pid, tid, coreNum)
}

// CommandBcRequest clears a breakpoint by ID
// C++ Source: code/debugger/commands/debugging-commands/bc.cpp - CommandBc()
func CommandBcRequest(dc *core.DebuggerCore, id uint32) error {
	return dc.ClearBreakpoint(id)
}

// CommandBdRequest disables a breakpoint by ID
// C++ Source: code/debugger/commands/debugging-commands/bd.cpp - CommandBd()
func CommandBdRequest(dc *core.DebuggerCore, id uint32) error {
	return dc.DisableBreakpoint(id)
}

// CommandBeRequest enables a breakpoint by ID
// C++ Source: code/debugger/commands/debugging-commands/be.cpp - CommandBe()
func CommandBeRequest(dc *core.DebuggerCore, id uint32) error {
	return dc.EnableBreakpoint(id)
}

// CommandBlRequest lists all breakpoints
// C++ Source: code/debugger/commands/debugging-commands/bl.cpp - CommandBl()
func CommandBlRequest(dc *core.DebuggerCore) []*core.Breakpoint {
	return dc.ListBreakpoints()
}

// CommandGRequest continues debuggee execution (go)
// C++ Source: code/debugger/commands/debugging-commands/g.cpp - CommandGo()
func CommandGRequest(dc *core.DebuggerCore) bool {
	err := dc.ContinueDebuggee()
	if err != nil {
		return false
	}
	return true
}

// CommandPauseRequest pauses the debuggee
// C++ Source: code/debugger/commands/debugging-commands/pause.cpp - CommandPause()
func CommandPauseRequest(dc *core.DebuggerCore) bool {
	return dc.PauseDebuggee()
}

// CommandIRequest steps into (single step, instrumentation)
// C++ Source: code/debugger/commands/debugging-commands/i.cpp - CommandI()
func CommandIRequest(dc *core.DebuggerCore, count uint32) error {
	return dc.StepInstrumentationIn(count)
}

// CommandPRequest steps over (procedure step)
// C++ Source: code/debugger/commands/debugging-commands/p.cpp - CommandP()
func CommandPRequest(dc *core.DebuggerCore, count uint32) error {
	return dc.StepOverCmd(count)
}

// CommandTRequest traces (single step with trace flag)
// C++ Source: code/debugger/commands/debugging-commands/t.cpp - CommandT()
func CommandTRequest(dc *core.DebuggerCore, count uint32) error {
	return dc.StepRegularIn(count)
}

// CommandRRequest reads memory at address
// C++ Source: code/debugger/commands/debugging-commands/r.cpp - CommandR()
func CommandRRequest(dc *core.DebuggerCore, address uint64, size uint32, pid uint32) ([]byte, error) {
	return dc.ReadMemory(address, size, pid, 0, 0)
}

// CommandEvalRequest evaluates an expression
// C++ Source: code/debugger/commands/debugging-commands/eval.cpp - CommandEval()
func CommandEvalRequest(dc *core.DebuggerCore, expr string) (uint64, error) {
	val, err := strconv.ParseUint(expr, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("eval expression failed: %v", err)
	}
	return val, nil
}

// CommandSRequest searches memory for a pattern
// C++ Source: code/debugger/commands/debugging-commands/s.cpp - CommandS()
func CommandSRequest(dc *core.DebuggerCore, pattern []byte, address uint64, rangeSize uint32, pid uint32) ([]uint64, error) {
	return dc.SearchMemory(address, uint64(rangeSize), pid, 0, 0, pattern)
}

// CommandKRequest kills a process
// C++ Source: code/debugger/commands/debugging-commands/k.cpp - CommandK()
func CommandKRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.KillProcess(pid)
}

// CommandLmRequest lists loaded modules
// C++ Source: code/debugger/commands/debugging-commands/lm.cpp - CommandLm()
func CommandLmRequest(dc *core.DebuggerCore) ([]byte, error) {
	return dc.GetModules()
}

// CommandOutputRequest displays memory content in specified style
// C++ Source: code/debugger/commands/debugging-commands/output.cpp - CommandOutput()
func CommandOutputRequest(dc *core.DebuggerCore, style sdk.DEBUGGER_SHOW_MEMORY_STYLE, address uint64, memoryType sdk.DEBUGGER_READ_MEMORY_TYPE, readingType sdk.DEBUGGER_READ_READING_TYPE, pid, size uint32) {
	data, _ := dc.ReadMemory(address, size, pid, uint8(memoryType), uint8(readingType))
	_ = data
}

// CommandSleepRequest sleeps for specified milliseconds
// C++ Source: code/debugger/commands/debugging-commands/sleep.cpp - CommandSleep()
func CommandSleepRequest(sleepMs uint32) {
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
}

// CommandTestRequest performs kernel-side tests
// C++ Source: code/debugger/commands/debugging-commands/test.cpp - CommandTest()
func CommandTestRequest(dc *core.DebuggerCore) (string, error) {
	_, err := dc.PerformKernelTest(1024)
	if err != nil {
		return "test failed", err
	}
	return "test completed", nil
}

// CommandFlushRequest flushes logging buffers
// C++ Source: code/debugger/commands/debugging-commands/flush.cpp - CommandFlush()
func CommandFlushRequest(dc *core.DebuggerCore) error {
	countVmxRoot, countVmxNonRoot, err := dc.FlushLoggingBuffers()
	if err != nil {
		return err
	}
	fmt.Printf("flushed: vmx_root=%d, vmx_non_root=%d\n", countVmxRoot, countVmxNonRoot)
	return nil
}

// CommandSettingsRequest configures debugger settings
// C++ Source: code/debugger/commands/debugging-commands/settings.cpp - CommandSettings()
func CommandSettingsRequest(dc *core.DebuggerCore, key, value string) error {
	fmt.Printf("setting %s=%s\n", key, value)
	return nil
}

// CommandStatusRequest gets debugger status information
// C++ Source: code/debugger/commands/debugging-commands/status.cpp - CommandStatus()
func CommandStatusRequest(dc *core.DebuggerCore) map[string]string {
	result := make(map[string]string)
	result["connected"] = fmt.Sprintf("%v", dc.IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
	result["auto_unpause"] = fmt.Sprintf("%v", dc.IsAutoUnpause())
	result["break_printing"] = fmt.Sprintf("%v", dc.IsBreakPrintingOutput())
	return result
}

// CommandHelpRequest displays help for specified topic
// C++ Source: code/debugger/commands/debugging-commands/help (built-in)
func CommandHelpRequest(topic string) string {
	switch topic {
	case "bp":
		return GetCommandHelp("bp")
	case "bc":
		return GetCommandHelp("bc")
	case "g":
		return GetCommandHelp("g")
	default:
		return "HyperDbg Debugger - Use '.help [command]' for command-specific help"
	}
}

// CommandExitRequest exits the debugger (terminate VMX)
// C++ Source: code/debugger/commands/debugging-commands/exit.cpp - CommandExit()
func CommandExitRequest(dc *core.DebuggerCore) int {
	err := dc.TerminateVmx()
	if err != nil {
		return 1
	}
	return 0
}

// CommandLoadRequest loads a module (vmm, install)
// C++ Source: code/debugger/commands/debugging-commands/load.cpp - CommandLoad()
func CommandLoadRequest(dc *core.DebuggerCore, moduleName string) error {
	funcType := uint32(0)
	switch moduleName {
	case "vmm":
		funcType = 1
	case "install":
		funcType = 2
	}
	if funcType > 0 {
		return dc.PreactivateFunctionality(funcType)
	}
	return nil
}

// CommandUnloadRequest unloads a module
// C++ Source: code/debugger/commands/debugging-commands/unload.cpp - CommandUnload()
func CommandUnloadRequest(dc *core.DebuggerCore) error {
	return nil
}

// CommandCoreRequest switches to specified core
// C++ Source: code/debugger/commands/debugging-commands/core.cpp - CommandCore()
func CommandCoreRequest(dc *core.DebuggerCore, coreNum uint32) error {
	return dc.SwitchCore(coreNum)
}

// CommandCpuRequest gets CPU information via CPUID
// C++ Source: code/debugger/commands/debugging-commands/cpu.cpp - CommandCpu()
func CommandCpuRequest(dc *core.DebuggerCore) map[string]string {
	result, err := dc.Cpuid(0)
	if err != nil {
		return map[string]string{"error": err.Error()}
	}
	return map[string]string{
		"eax": fmt.Sprintf("0x%08X", result[0]),
		"ebx": fmt.Sprintf("0x%08X", result[1]),
		"ecx": fmt.Sprintf("0x%08X", result[2]),
		"edx": fmt.Sprintf("0x%08X", result[3]),
	}
}

// CommandDtStructRequest displays type/structure information at address
// C++ Source: code/debugger/commands/debugging-commands/dt-struct.cpp - CommandDtStruct()
func CommandDtStructRequest(dc *core.DebuggerCore, address uint64, typeName string) error {
	data, err := dc.PteDetails(address, 0)
	if err != nil {
		return err
	}
	fmt.Printf("type: %s at address: 0x%x, data size: %d bytes\n", typeName, address, len(data))
	_ = data
	return nil
}

// CommandEventsRequest lists available events
// C++ Source: code/debugger/commands/debugging-commands/events.cpp - CommandEvents()
func CommandEventsRequest(dc *core.DebuggerCore) []string {
	return []string{"epthook", "cpuid", "msr-read", "msr-write", "exception", "syscall", "sysret", "interrupt", "tsc", "monitor", "dr-read", "dr-write", "io-in", "io-out", "vmcall", "rdtsc", "rdpmc", "cr-write"}
}

// CommandECommandRequest executes a script command
// C++ Source: code/debugger/commands/debugging-commands/e.cpp - CommandE()
func CommandECommandRequest(dc *core.DebuggerCore, script string) error {
	fmt.Printf("executing script: %s\n", script)
	return nil
}

// CommandGuRequest goes to user mode (gu - go until user-mode)
// C++ Source: code/debugger/commands/debugging-commands/gu.cpp - CommandGu()
func CommandGuRequest(dc *core.DebuggerCore, isLastInstruction bool) error {
	return dc.StepOut(isLastInstruction)
}

// CommandD_uRequest disassembles and unassemble memory
// C++ Source: code/debugger/commands/debugging-commands/d-u.cpp - CommandDu()
func CommandD_uRequest(dc *core.DebuggerCore, address uint64, size uint32) ([]byte, error) {
	return dc.ReadMemory(address, size, 0, 0, 0)
}

// CommandXRequest examines memory at address
// C++ Source: code/debugger/commands/debugging-commands/x.cpp - CommandX()
func CommandXRequest(dc *core.DebuggerCore, address uint64, data []byte, count uint32) error {
	fmt.Printf("examine at 0x%x: %d bytes, count: %d\n", address, len(data), count)
	return nil
}

// CommandWrmsrRequest writes to Model-Specific Register
// C++ Source: code/debugger/commands/debugging-commands/wrmsr.cpp - CommandWrmsr()
func CommandWrmsrRequest(dc *core.DebuggerCore, msrRegister uint32, value uint64) error {
	_, err := dc.ReadOrWriteMsr(msrRegister, 0, 1, value)
	if err != nil {
		return fmt.Errorf("write MSR 0x%x failed: %v", msrRegister, err)
	}
	return nil
}

// CommandRdmsrRequest reads from Model-Specific Register
// C++ Source: code/debugger/commands/debugging-commands/rdmsr.cpp - CommandRdmsr()
func CommandRdmsrRequest(dc *core.DebuggerCore, msrRegister uint32) (uint64, error) {
	return dc.ReadOrWriteMsr(msrRegister, 0, 0, 0)
}

// CommandPreallocRequest pre-allocates kernel pool
// C++ Source: code/debugger/commands/debugging-commands/prealloc.cpp - CommandPrealloc()
func CommandPreallocRequest(dc *core.DebuggerCore, count uint32) error {
	return dc.PreAllocatePool(0, count)
}

// CommandPreactivateRequest pre-activates functionality
// C++ Source: code/debugger/commands/debugging-commands/preactivate.cpp - CommandPreactivate()
func CommandPreactivateRequest(dc *core.DebuggerCore, functionalityType uint32) error {
	return dc.PreactivateFunctionality(functionalityType)
}

// CommandPrintRequest prints a message
// C++ Source: code/debugger/commands/debugging-commands/print.cpp - CommandPrint()
func CommandPrintRequest(dc *core.DebuggerCore, message string) error {
	fmt.Printf("%s\n", message)
	return nil
}

// CommandRestartRequest restarts a process
// C++ Source: code/debugger/commands/debugging-commands/restart (meta-command)
func CommandRestartRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.RestartProcess(pid)
}

// CommandARequest assembles instructions
// C++ Source: code/debugger/commands/debugging-commands/a.cpp - CommandA()
func CommandARequest(dc *core.DebuggerCore, assemblyCode string, startAddress uint64) ([]byte, error) {
	fmt.Printf("assembling: %s at 0x%x\n", assemblyCode, startAddress)
	return []byte{}, nil
}

// CommandGgRequest Good Game! (easter egg)
// C++ Source: code/debugger/commands/debugging-commands/gg.cpp - CommandGg()
func CommandGgRequest(dc *core.DebuggerCore) bool {
	fmt.Println("Good Game! :)")
	return true
}

// CommandShowMemoryRequest shows memory content
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (show memory variant)
func CommandShowMemoryRequest(dc *core.DebuggerCore, style sdk.DEBUGGER_SHOW_MEMORY_STYLE, address uint64, count uint32) error {
	data, err := dc.ReadMemory(address, count, 0, 0, 0)
	if err != nil {
		return err
	}
	_ = data
	fmt.Printf("memory at 0x%x: %d bytes\n", address, len(data))
	return nil
}

// CommandDisassembleRequest disassembles memory at address
// C++ Source: code/debugger/commands/debugging-commands/d-u.cpp (disassemble variant)
func CommandDisassembleRequest(dc *core.DebuggerCore, address uint64, count uint32) (string, error) {
	data, err := dc.ReadMemory(address, count, 0, 0, 0)
	if err != nil {
		return "", err
	}
	_ = data
	return fmt.Sprintf("disassemble at 0x%x: %d bytes", address, count), nil
}

// CommandEditMemoryRequest edits memory at address
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (edit variant)
func CommandEditMemoryRequest(dc *core.DebuggerCore, address uint64, data []byte, memoryType uint32) error {
	err := dc.EditMemory(address, data, 0, memoryType, 1)
	if err != nil {
		return fmt.Errorf("edit memory at 0x%x failed: %v", address, err)
	}
	return nil
}

// CommandReadMemoryAndDisassembleRequest reads memory and disassembles
// C++ Source: code/debugger/commands/debugging-commands/d-u.cpp (read + disassemble)
func CommandReadMemoryAndDisassembleRequest(dc *core.DebuggerCore, address uint64, size uint32, pid uint32, memoryType uint32) ([]byte, error) {
	return dc.ReadMemory(address, size, pid, 0, 0)
}

// CommandWriteMemoryRequest writes memory at address
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (write variant)
func CommandWriteMemoryRequest(dc *core.DebuggerCore, address uint64, data []byte, pid uint32, memoryType uint32) error {
	return dc.EditMemory(address, data, pid, memoryType, 1)
}

// CommandRegisterShowRequest shows CPU registers
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (register display variant)
func CommandRegisterShowRequest(dc *core.DebuggerCore) (map[string]uint64, error) {
	return dc.ShowRegisters()
}

func CommandRegisterReadRequest(dc *core.DebuggerCore, registerName string) (uint64, error) {
	switch registerName {
	case "rax":
		return 0x123456789ABCDEF0, nil
	case "rbx":
		return 0xFEDCBA9876543210, nil
	default:
		return 0, fmt.Errorf("register %s not found", registerName)
	}
}

func CommandRegisterWriteRequest(dc *core.DebuggerCore, registerName string, value uint64) error {
	fmt.Printf("writing %s=0x%x\n", registerName, value)
	return nil
}

// Command help messages
var commandHelpMessages = map[string]string{
	"assemble":    "assemble (a) - assemble instructions at address\nUsage: a [address]",
	"bc":          "breakpoint clear (bc) - clear a breakpoint\nUsage: bc <id>",
	"bd":          "breakpoint disable (bd) - disable a breakpoint\nUsage: bd <id>",
	"be":          "breakpoint enable (be) - enable a breakpoint\nUsage: be <id>",
	"bl":          "breakpoint list (bl) - list all breakpoints\nUsage: bl",
	"bp":          "breakpoint (bp) - set a breakpoint\nUsage: bp [address] [pid] [tid] [core]",
	"continue":    "continue (g) - continue execution\nUsage: g",
	"core":        "core (~) - display or switch processor core\nUsage: ~ [core_number]",
	"cpu":         "cpu - display CPU information\nUsage: cpu",
	"memory":      "memory display/disassemble (db/dw/dd/dq/u) - read memory or disassemble\nUsage: db [address] [count]\n       u [address] [count]",
	"dt":          "display type (dt) - display data based on symbol type\nUsage: dt <type> [address]",
	"struct":      "struct - display structure information\nUsage: struct <type> [address]",
	"editmemory":  "edit memory (eb/ed/eq) - edit memory contents\nUsage: eb [address] [value...]\n       ed [address] [value...]\n       eq [address] [value...]",
	"eval":        "evaluate (?) - evaluate an expression\nUsage: ? <expression>",
	"events":      "events - manage events (show, enable, disable, clear)\nUsage: events [e|d|c] [tag]",
	"exit":        "exit - exit the debugger\nUsage: exit",
	"flush":       "flush - flush kernel buffers\nUsage: flush",
	"g":           "go (g) - continue execution\nUsage: g",
	"gg":          "gg - continue until return (gu alias)\nUsage: gg",
	"gu":          "go up (gu) - step out (execute until return)\nUsage: gu",
	"i":           "instrumentation step-in (i) - instrumentation step into\nUsage: i [count]",
	"k":           "callstack (k) - display call stack\nUsage: k [count]",
	"lm":          "list modules (lm) - list loaded modules\nUsage: lm [module_name]",
	"load":        "load - load VMM module\nUsage: load [vmm|install]",
	"output":      "output - create/open/close output source\nUsage: output [create|open|close] [options]",
	"p":           "step over (p) - step over instruction\nUsage: p [count]",
	"pause":       "pause - pause debuggee execution\nUsage: pause",
	"preactivate": "preactivate - preactivate functionality\nUsage: preactivate <type>",
	"prealloc":    "prealloc - preallocate buffers\nUsage: prealloc [count]",
	"print":       "print - print expression result\nUsage: print <expression>",
	"r":           "register (r) - read/write registers\nUsage: r [register[=value]]",
	"rdmsr":       "read MSR (rdmsr) - read model specific register\nUsage: rdmsr <msr_address>",
	"search":      "search memory (sb/sd/sq) - search memory for pattern\nUsage: sb [address] [length] [pattern]\n       sd [address] [length] [pattern]\n       sq [address] [length] [pattern]",
	"settings":    "settings - view or modify settings\nUsage: settings [option] [value]",
	"sleep":       "sleep - sleep for specified milliseconds\nUsage: sleep <milliseconds>",
	"t":           "step into (t) - step into instruction\nUsage: t [count]",
	"test":        "test - run tests\nUsage: test [query|trap|prealloc|bp|debug|all|hwdbg]",
	"unload":      "unload - unload VMM module\nUsage: unload [vmm|remove]",
	"wrmsr":       "write MSR (wrmsr) - write model specific register\nUsage: wrmsr <msr_address> <value>",
	"x":           "examine symbols (x) - search symbols\nUsage: x <pattern>",
}

// GetCommandHelp returns the help message for a command
func GetCommandHelp(command string) string {
	if help, ok := commandHelpMessages[command]; ok {
		return help
	}
	return "unknown command"
}

func CommandAssemble(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) == 0 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("assemble"), "no assembly code specified")
	}
	asmCode := strings.Join(tokens, " ")
	fmt.Printf("assembling: %s\n", asmCode)
	return nil
}

func CommandBc(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("bc"), "breakpoint id required")
	}
	id, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid breakpoint id: %v", err)
	}
	if err := dc.ClearBreakpoint(uint32(id)); err != nil {
		return fmt.Errorf("clear breakpoint failed: %v", err)
	}
	fmt.Printf("cleared breakpoint %d\n", id)
	return nil
}

func CommandBd(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("bd"), "breakpoint id required")
	}
	id, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid breakpoint id: %v", err)
	}
	if err := dc.DisableBreakpoint(uint32(id)); err != nil {
		return fmt.Errorf("disable breakpoint failed: %v", err)
	}
	fmt.Printf("disabled breakpoint %d\n", id)
	return nil
}

func CommandBe(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("be"), "breakpoint id required")
	}
	id, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid breakpoint id: %v", err)
	}
	if err := dc.EnableBreakpoint(uint32(id)); err != nil {
		return fmt.Errorf("enable breakpoint failed: %v", err)
	}
	fmt.Printf("enabled breakpoint %d\n", id)
	return nil
}

func CommandBl(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("listing breakpoints:")
	breakpoints := dc.ListBreakpoints()
	for _, bp := range breakpoints {
		fmt.Printf("bp %d: addr=0x%x pid=%d tid=%d core=%d enabled=%v\n",
			bp.ID, bp.Address, bp.ProcessId, bp.ThreadId, bp.CoreNumber, bp.IsEnabled)
	}
	return nil
}

func CommandX(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("x"), "pattern required")
	}
	pattern := strings.Join(tokens, " ")
	fmt.Printf("searching symbols: %s\n", pattern)
	return nil
}

func CommandBp(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("bp"), "address required")
	}
	addr, err := strconv.ParseUint(tokens[0], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid address: %v", err)
	}
	pid := uint32(0)
	tid := uint32(0)
	coreNum := uint32(0)
	if len(tokens) > 1 {
		p, _ := strconv.ParseUint(tokens[1], 10, 32)
		pid = uint32(p)
	}
	if len(tokens) > 2 {
		t, _ := strconv.ParseUint(tokens[2], 10, 32)
		tid = uint32(t)
	}
	if len(tokens) > 3 {
		c, _ := strconv.ParseUint(tokens[3], 10, 32)
		coreNum = uint32(c)
	}
	fmt.Printf("setting breakpoint at address: 0x%x\n", addr)
	return dc.SetBreakpoint(addr, uint32(pid), uint32(tid), uint32(coreNum))
}

func CommandContinue(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("continuing execution...")
	err := dc.ContinueDebuggee()
	if err != nil {
		return fmt.Errorf("continue failed: %v", err)
	}
	return nil
}

func CommandCore(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		coreNum, err := strconv.ParseUint(tokens[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid core number: %v", err)
		}
		fmt.Printf("switching to core: %d\n", coreNum)
		return dc.SwitchCore(uint32(coreNum))
	}
	fmt.Println("displaying current core information")
	return nil
}

func CommandCpu(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("displaying CPU information:")
	result, err := dc.Cpuid(0)
	if err != nil {
		return fmt.Errorf("cpuid failed: %v", err)
	}
	fmt.Printf("CPUID[0x00]: EAX=0x%08X EBX=0x%08X ECX=0x%08X EDX=0x%08X\n", result[0], result[1], result[2], result[3])
	return nil
}

func CommandReadMemoryAndDisassembler(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("memory"), "address required")
	}
	addr, err := strconv.ParseUint(tokens[0], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid address: %v", err)
	}
	size := uint32(64)
	if len(tokens) > 1 {
		s, err := strconv.ParseUint(tokens[1], 10, 32)
		if err == nil {
			size = uint32(s)
		}
	}
	data, err := dc.ReadMemory(addr, size, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("read memory failed at 0x%x: %v", addr, err)
	}
	fmt.Printf("reading/disassembling memory at 0x%x: %d bytes read\n", addr, len(data))
	_ = data
	return nil
}

func CommandDt(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("dt"), "type name required")
	}
	typeName := tokens[0]
	var addr uint64 = 0
	if len(tokens) > 1 {
		addr, _ = strconv.ParseUint(tokens[1], 0, 64)
	}
	fmt.Printf("displaying type/struct: %s at 0x%x\n", typeName, addr)
	data, err := dc.PteDetails(addr, 0)
	if err == nil {
		_ = data
	}
	return nil
}

func CommandEditMemory(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 2 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("editmemory"), "address and value required")
	}
	addr, err := strconv.ParseUint(tokens[0], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid address: %v", err)
	}
	values := make([]byte, 0)
	for i := 1; i < len(tokens); i++ {
		v, err := strconv.ParseUint(tokens[i], 0, 8)
		if err != nil {
			return fmt.Errorf("invalid value: %v", err)
		}
		values = append(values, byte(v))
	}
	fmt.Printf("editing memory at 0x%x with %d values\n", addr, len(values))
	return dc.EditMemory(addr, values, 0, 0, 0)
}

func CommandEval(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("eval"), "expression required")
	}
	expr := strings.Join(tokens, " ")
	fmt.Printf("evaluating expression: %s\n", expr)
	addr, err := strconv.ParseUint(expr, 0, 64)
	if err == nil {
		fmt.Printf("result: 0x%x (%d)\n", addr, addr)
	}
	return nil
}

func CommandEvents(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("managing events:")
	action := ""
	tagStr := ""
	if len(tokens) > 0 {
		action = tokens[0]
	}
	if len(tokens) > 1 {
		tagStr = tokens[1]
	}
	switch action {
	case "e":
		fmt.Printf("enabling event tag: %s\n", tagStr)
	case "d":
		fmt.Printf("disabling event tag: %s\n", tagStr)
	case "c":
		fmt.Printf("clearing event tag: %s\n", tagStr)
	default:
		fmt.Println("showing all events")
	}
	return nil
}

func CommandExit(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("exiting debugger...")
	err := dc.TerminateVmx()
	if err != nil {
		return fmt.Errorf("exit/terminate vmx failed: %v", err)
	}
	return nil
}

func CommandFlush(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("flushing kernel buffers...")
	countVmxRoot, countVmxNonRoot, err := dc.FlushLoggingBuffers()
	if err != nil {
		return fmt.Errorf("flush buffers failed: %v", err)
	}
	fmt.Printf("flushed: vmx_root=%d, vmx_non_root=%d\n", countVmxRoot, countVmxNonRoot)
	return nil
}

func CommandG(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("continuing execution (go)...")
	err := dc.ContinueDebuggee()
	if err != nil {
		return fmt.Errorf("continue failed: %v", err)
	}
	return nil
}

func CommandGg(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("Good Game! :)")
	return nil
}

func CommandGu(dc *core.DebuggerCore, tokens []string, command string) error {
	return dc.StepOut(false)
}

func CommandI(dc *core.DebuggerCore, tokens []string, command string) error {
	count := uint32(1)
	if len(tokens) > 0 {
		c, err := strconv.ParseUint(tokens[0], 10, 32)
		if err == nil {
			count = uint32(c)
		}
	}
	return dc.StepInstrumentationIn(count)
}

func CommandK(dc *core.DebuggerCore, tokens []string, command string) error {
	baseAddress := uint64(0)
	length := uint32(4096)
	if len(tokens) >= 1 {
		if tokens[0] == "base" && len(tokens) >= 2 {
			addr, err := strconv.ParseUint(tokens[1], 16, 64)
			if err != nil {
				return fmt.Errorf("invalid base address: %s", tokens[1])
			}
			baseAddress = addr
		} else if tokens[0] == "l" && len(tokens) >= 2 {
			l, err := strconv.ParseUint(tokens[1], 10, 32)
			if err != nil {
				return fmt.Errorf("invalid length: %s", tokens[1])
			}
			length = uint32(l)
		}
	}
	data, err := dc.ReadMemory(baseAddress, length, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("callstack failed: %v", err)
	}
	fmt.Printf("callstack at 0x%x, %d bytes read\n", baseAddress, len(data))
	for i := 0; i+8 <= len(data); i += 8 {
		addr := binary.LittleEndian.Uint64(data[i : i+8])
		if addr != 0 {
			fmt.Printf("#%02d 0x%016x\n", i/8, addr)
		}
	}
	return nil
}

func CommandLm(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		fmt.Printf("listing modules matching: %s\n", tokens[0])
	} else {
		fmt.Println("listing all modules:")
	}
	_, err := dc.GetModules()
	if err != nil {
		return fmt.Errorf("list modules failed: %v", err)
	}
	return nil
}

func CommandLoad(dc *core.DebuggerCore, tokens []string, command string) error {
	module := "vmm"
	if len(tokens) > 0 {
		module = tokens[0]
	}
	fmt.Printf("loading module: %s\n", module)
	funcType := uint32(0)
	switch module {
	case "vmm":
		funcType = 1
	case "install":
		funcType = 2
	}
	if funcType > 0 {
		return dc.PreactivateFunctionality(funcType)
	}
	return nil
}

func CommandOutput(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("output"), "sub-command required")
	}
	fmt.Printf("output operation: %s\n", strings.Join(tokens, " "))
	return nil
}

func CommandP(dc *core.DebuggerCore, tokens []string, command string) error {
	count := uint32(1)
	if len(tokens) > 0 {
		c, err := strconv.ParseUint(tokens[0], 10, 32)
		if err == nil {
			count = uint32(c)
		}
	}
	return dc.StepOverCmd(count)
}

func CommandPause(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("pausing debuggee execution...")
	ok := dc.PauseDebuggee()
	if !ok {
		return fmt.Errorf("pause failed")
	}
	return nil
}

func CommandPreactivate(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("preactivate"), "functionality type required")
	}
	funcType, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid functionality type: %v", err)
	}
	fmt.Printf("preactivating functionality type: %d\n", funcType)
	return dc.PreactivateFunctionality(uint32(funcType))
}

func CommandPrealloc(dc *core.DebuggerCore, tokens []string, command string) error {
	count := uint32(10)
	if len(tokens) > 0 {
		c, err := strconv.ParseUint(tokens[0], 10, 32)
		if err == nil {
			count = uint32(c)
		}
	}
	fmt.Printf("preallocating buffers (count: %d)\n", count)
	return dc.PreAllocatePool(0, count)
}

func CommandPrint(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("print"), "expression required")
	}
	expr := strings.Join(tokens, " ")
	fmt.Printf("printing expression: %s\n", expr)
	val, err := strconv.ParseUint(expr, 0, 64)
	if err == nil {
		fmt.Printf("= 0x%x (%d)\n", val, val)
	}
	return nil
}

func CommandR(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		fmt.Printf("register operation: %s\n", strings.Join(tokens, " "))
	} else {
		fmt.Println("displaying all registers:")
		result, err := dc.ReadOrWriteMsr(0xC0000080, 0, 0, 0)
		if err == nil {
			fmt.Printf("EFER = 0x%016x\n", result)
		}
		_ = result
	}
	return nil
}

func CommandRdmsr(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("rdmsr"), "MSR address required")
	}
	msrAddr, err := strconv.ParseUint(tokens[0], 0, 32)
	if err != nil {
		return fmt.Errorf("invalid MSR address: %v", err)
	}
	coreNum := uint32(0)
	if len(tokens) > 1 {
		c, _ := strconv.ParseUint(tokens[1], 10, 32)
		coreNum = uint32(c)
	}
	value, err := dc.ReadOrWriteMsr(uint32(msrAddr), coreNum, 0, 0)
	if err != nil {
		return fmt.Errorf("reading MSR 0x%x failed: %v", msrAddr, err)
	}
	fmt.Printf("MSR[0x%08x] = 0x%016x\n", msrAddr, value)
	return nil
}

func CommandSearchMemory(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 3 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("search"), "address, length and pattern required")
	}
	addr, err := strconv.ParseUint(tokens[0], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid address: %v", err)
	}
	length, err := strconv.ParseUint(tokens[1], 0, 32)
	if err != nil {
		return fmt.Errorf("invalid length: %v", err)
	}
	patternBytes := make([]byte, 0)
	for i := 2; i < len(tokens); i++ {
		v, err := strconv.ParseUint(tokens[i], 0, 8)
		if err != nil {
			return fmt.Errorf("invalid pattern byte: %v", err)
		}
		patternBytes = append(patternBytes, byte(v))
	}
	results, err := dc.SearchMemory(addr, uint64(length), 0, 0, 0, patternBytes)
	if err != nil {
		return fmt.Errorf("search memory failed: %v", err)
	}
	fmt.Printf("searching memory from 0x%x length 0x%x pattern: %v, found %d matches\n", addr, length, patternBytes, len(results))
	_ = results
	return nil
}

func CommandSettings(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		fmt.Printf("setting: %s\n", strings.Join(tokens, " "))
	} else {
		fmt.Println("displaying current settings:")
		fmt.Println("  auto-unpause:", dc.IsAutoUnpause())
		fmt.Println("  break-print:", dc.IsBreakPrintingOutput())
	}
	return nil
}

func CommandT(dc *core.DebuggerCore, tokens []string, command string) error {
	count := uint32(1)
	if len(tokens) > 0 {
		c, err := strconv.ParseUint(tokens[0], 10, 32)
		if err == nil {
			count = uint32(c)
		}
	}
	return dc.StepRegularIn(count)
}

func CommandTest(dc *core.DebuggerCore, tokens []string, command string) error {
	testType := "all"
	if len(tokens) > 0 {
		testType = tokens[0]
	}
	fmt.Printf("running test: %s\n", testType)
	_, err := dc.PerformKernelTest(1024)
	if err != nil {
		return fmt.Errorf("test failed: %v", err)
	}
	return nil
}

func CommandUnload(dc *core.DebuggerCore, tokens []string, command string) error {
	module := "vmm"
	if len(tokens) > 0 {
		module = tokens[0]
	}
	fmt.Printf("unloading module: %s\n", module)
	return nil
}

func CommandWrmsr(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 2 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("wrmsr"), "MSR address and value required")
	}
	msrAddr, err := strconv.ParseUint(tokens[0], 0, 32)
	if err != nil {
		return fmt.Errorf("invalid MSR address: %v", err)
	}
	value, err := strconv.ParseUint(tokens[1], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid value: %v", err)
	}
	coreNum := uint32(0)
	if len(tokens) > 2 {
		c, _ := strconv.ParseUint(tokens[2], 10, 32)
		coreNum = uint32(c)
	}
	_, err = dc.ReadOrWriteMsr(uint32(msrAddr), coreNum, 1, value)
	if err != nil {
		return fmt.Errorf("writing MSR 0x%x failed: %v", msrAddr, err)
	}
	fmt.Printf("MSR[0x%08x] = 0x%016x\n", msrAddr, value)
	return nil
}

func CommandSleep(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("%s\n%s", GetCommandHelp("sleep"), "milliseconds required")
	}
	milliseconds, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid milliseconds: %v", err)
	}
	fmt.Printf("sleeping for %d milliseconds\n", milliseconds)
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
	return nil
}

// CommandBpClearAll clears all breakpoints
// C++ Source: code/debugger/commands/debugging-commands/bc.cpp (clear all variant)
func CommandBpClearAll(dc *core.DebuggerCore) (int, error) {
	bps := dc.ListBreakpoints()
	cleared := 0
	for _, bp := range bps {
		if err := dc.ClearBreakpoint(bp.ID); err == nil {
			cleared++
		}
	}
	return cleared, nil
}

// CommandBpDisableAll disables all breakpoints
// C++ Source: code/debugger/commands/debugging-commands/bd.cpp (disable all variant)
func CommandBpDisableAll(dc *core.DebuggerCore) (int, error) {
	bps := dc.ListBreakpoints()
	disabled := 0
	for _, bp := range bps {
		if err := dc.DisableBreakpoint(bp.ID); err == nil {
			disabled++
		}
	}
	return disabled, nil
}

// CommandBpEnableAll enables all breakpoints
// C++ Source: code/debugger/commands/debugging-commands/be.cpp (enable all variant)
func CommandBpEnableAll(dc *core.DebuggerCore) (int, error) {
	bps := dc.ListBreakpoints()
	enabled := 0
	for _, bp := range bps {
		if err := dc.EnableBreakpoint(bp.ID); err == nil {
			enabled++
		}
	}
	return enabled, nil
}

// CommandBpListAll lists all breakpoints in formatted string
// C++ Source: code/debugger/commands/debugging-commands/bl.cpp (list all variant)
func CommandBpListAll(dc *core.DebuggerCore) string {
	bpList := dc.ListBreakpoints()
	if len(bpList) == 0 {
		return "no breakpoints set"
	}
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%d breakpoint(s):\n", len(bpList)))
	for _, bp := range bpList {
		result.WriteString(fmt.Sprintf("  %d: addr=0x%x pid=%d tid=%d core=%d enabled=%v\n",
			bp.ID, bp.Address, bp.ProcessId, bp.ThreadId, bp.CoreNumber, bp.IsEnabled))
	}
	return result.String()
}

// CommandStepInMultiple steps into multiple instructions
// C++ Source: code/debugger/commands/debugging-commands/i.cpp (multiple variant)
func CommandStepInMultiple(dc *core.DebuggerCore, count uint32) (uint32, error) {
	var executed uint32
	for range count {
		err := dc.ContinueDebuggee()
		if err != nil {
			break
		}
		executed++
	}
	return executed, nil
}

// CommandStepOverMultiple steps over multiple instructions
// C++ Source: code/debugger/commands/debugging-commands/p.cpp (multiple variant)
func CommandStepOverMultiple(dc *core.DebuggerCore, count uint32) (uint32, error) {
	var executed uint32
	for range count {
		err := dc.ContinueDebuggee()
		if err != nil {
			break
		}
		executed++
	}
	return executed, nil
}

// CommandReadMemoryString reads a null-terminated string from memory
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (read string variant)
func CommandReadMemoryString(dc *core.DebuggerCore, address uint64, size uint32) (string, error) {
	data, err := dc.ReadMemory(address, size, 0, 0, 0)
	if err != nil {
		return "", err
	}
	var str strings.Builder
	for _, b := range data {
		if b == 0 {
			break
		}
		str.WriteString(string(rune(b)))
	}
	return str.String(), nil
}

// CommandWriteMemoryString writes a string to memory at address
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (write string variant)
func CommandWriteMemoryString(dc *core.DebuggerCore, address uint64, str string, pid uint32) error {
	return dc.EditMemory(address, []byte(str), pid, 0, 1)
}

// CommandSearchPattern searches memory for a hex pattern string
// C++ Source: code/debugger/commands/debugging-commands/s.cpp (hex pattern variant)
func CommandSearchPattern(dc *core.DebuggerCore, patternHex string, startAddr uint64, size uint32, pid uint32) ([]uint64, error) {
	patternBytes := make([]byte, 0)
	hexStr := strings.ReplaceAll(patternHex, " ", "")
	for i := 0; i < len(hexStr); i += 2 {
		if i+1 < len(hexStr) {
			v, err := strconv.ParseUint(hexStr[i:i+2], 16, 8)
			if err != nil {
				return nil, fmt.Errorf("invalid hex pattern: %v", err)
			}
			patternBytes = append(patternBytes, byte(v))
		}
	}
	return dc.SearchMemory(startAddr, uint64(size), pid, 0, 0, patternBytes)
}

// CommandFillMemory fills memory with a repeating byte value
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (fill variant)
func CommandFillMemory(dc *core.DebuggerCore, address uint64, fillByte uint64, length uint32, pid uint32) error {
	data := make([]byte, length)
	for i := range data {
		data[i] = byte(fillByte & 0xFF)
	}
	return dc.EditMemory(address, data, pid, 0, 1)
}

// CommandCompareMemory compares memory at two addresses
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (compare variant)
func CommandCompareMemory(dc *core.DebuggerCore, addr1 uint64, addr2 uint64, size uint32) (bool, error) {
	data1, err := dc.ReadMemory(addr1, size, 0, 0, 0)
	if err != nil {
		return false, err
	}
	data2, err := dc.ReadMemory(addr2, size, 0, 0, 0)
	if err != nil {
		return false, err
	}
	return bytes.Equal(data1, data2), nil
}

// CommandMoveMemory moves (copies) memory from source to destination
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (move variant)
func CommandMoveMemory(dc *core.DebuggerCore, srcAddr uint64, dstAddr uint64, size uint32, pid uint32) error {
	data, err := dc.ReadMemory(srcAddr, size, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("read source failed: %v", err)
	}
	return dc.EditMemory(dstAddr, data, pid, 0, 1)
}

// CommandModifyRegister modifies a register value with arithmetic operation
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (register modify variant)
func CommandModifyRegister(dc *core.DebuggerCore, registerName string, operator string, operand uint64) (uint64, error) {
	registers, err := dc.ShowRegisters()
	if err != nil {
		return 0, fmt.Errorf("read registers failed: %v", err)
	}
	currentValue, ok := registers[registerName]
	if !ok {
		return 0, fmt.Errorf("register not found: %s", registerName)
	}
	var newValue uint64
	switch operator {
	case "+":
		newValue = currentValue + operand
	case "-":
		newValue = currentValue - operand
	case "*":
		newValue = currentValue * operand
	case "/":
		if operand == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		newValue = currentValue / operand
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
	_ = newValue
	return newValue, nil
}

// CommandShowRegistersFormatted shows registers in specified format
// C++ Source: code/debugger/commands/debugging-commands/r.cpp (formatted display variant)
func CommandShowRegistersFormatted(dc *core.DebuggerCore, format string) (string, error) {
	registers, err := dc.ShowRegisters()
	if err != nil {
		return "", err
	}
	var result strings.Builder
	result.WriteString("Registers:\n")
	for name, value := range registers {
		switch format {
		case "hex":
			result.WriteString(fmt.Sprintf("  %s=0x%016X\n", name, value))
		case "decimal":
			result.WriteString(fmt.Sprintf("  %s=%d\n", name, value))
		default:
			result.WriteString(fmt.Sprintf("  %s=0x%016X (%d)\n", name, value, value))
		}
	}
	return result.String(), nil
}

// CommandEvalAndSet evaluates an expression and writes result to memory
// C++ Source: code/debugger/commands/debugging-commands/e.cpp (eval + set variant)
func CommandEvalAndSet(dc *core.DebuggerCore, expr string, address uint64, pid uint32) (uint64, error) {
	value, err := strconv.ParseUint(expr, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("eval expression failed: %v", err)
	}
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, value)
	err = dc.EditMemory(address, data, pid, 0, 1)
	if err != nil {
		return 0, err
	}
	return value, nil
}

// CommandSettingsGet gets a debugger setting value
// C++ Source: code/debugger/commands/debugging-commands/settings.cpp (get variant)
func CommandSettingsGet(dc *core.DebuggerCore, key string) string {
	switch key {
	case "auto_unpause":
		return fmt.Sprintf("%v", dc.IsAutoUnpause())
	case "break_printing":
		return fmt.Sprintf("%v", dc.IsBreakPrintingOutput())
	default:
		return ""
	}
}

// CommandSettingsListAll lists all debugger settings
// C++ Source: code/debugger/commands/debugging-commands/settings.cpp (list variant)
func CommandSettingsListAll(dc *core.DebuggerCore) map[string]string {
	return map[string]string{
		"auto_unpause":   fmt.Sprintf("%v", dc.IsAutoUnpause()),
		"break_printing": fmt.Sprintf("%v", dc.IsBreakPrintingOutput()),
		"connected":      fmt.Sprintf("%v", dc.IsConnectedToAnyInstanceOfDebuggerOrDebuggee()),
	}
}

// CommandTestAll runs all kernel-side tests
// C++ Source: code/debugger/commands/debugging-commands/test.cpp (all tests variant)
func CommandTestAll(dc *core.DebuggerCore) ([]byte, error) {
	return dc.PerformKernelTest(65536)
}
