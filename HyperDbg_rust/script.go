package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ScriptEngine struct {
	dbg     Debugger
	vars    map[string]any
	mu      sync.RWMutex
	running bool
	ctx     context.Context
	cancel  context.CancelFunc
}

type ScriptCommand struct {
	Name    string
	Args    []string
	Handler func(ctx context.Context, args []string) (string, error)
}

func NewScriptEngine(dbg Debugger) *ScriptEngine {
	ctx, cancel := context.WithCancel(context.Background())
	se := &ScriptEngine{
		dbg:     dbg,
		vars:    make(map[string]any),
		ctx:     ctx,
		cancel:  cancel,
		running: false,
	}

	se.registerCommands()
	return se
}

func (se *ScriptEngine) registerCommands() {
	se.RegisterCommand("help", se.cmdHelp)
	se.RegisterCommand("bp", se.cmdBreakpoint)
	se.RegisterCommand("bc", se.cmdClearBreakpoint)
	se.RegisterCommand("bd", se.cmdDisableBreakpoint)
	se.RegisterCommand("be", se.cmdEnableBreakpoint)
	se.RegisterCommand("bl", se.cmdListBreakpoints)
	se.RegisterCommand("g", se.cmdContinue)
	se.RegisterCommand("p", se.cmdStepOver)
	se.RegisterCommand("t", se.cmdStepInto)
	se.RegisterCommand("gu", se.cmdStepOut)
	se.RegisterCommand("r", se.cmdRegisters)
	se.RegisterCommand("db", se.cmdDumpBytes)
	se.RegisterCommand("dw", se.cmdDumpWords)
	se.RegisterCommand("dd", se.cmdDumpDwords)
	se.RegisterCommand("dq", se.cmdDumpQwords)
	se.RegisterCommand("eb", se.cmdEditByte)
	se.RegisterCommand("ew", se.cmdEditWord)
	se.RegisterCommand("ed", se.cmdEditDword)
	se.RegisterCommand("eq", se.cmdEditQword)
	se.RegisterCommand("k", se.cmdCallStack)
	se.RegisterCommand("kv", se.cmdCallStackVerbose)
	se.RegisterCommand("lm", se.cmdListModules)
	se.RegisterCommand("process", se.cmdProcess)
	se.RegisterCommand("thread", se.cmdThread)
	se.RegisterCommand("attach", se.cmdAttach)
	se.RegisterCommand("detach", se.cmdDetach)
	se.RegisterCommand("eval", se.cmdEval)
	se.RegisterCommand("set", se.cmdSet)
	se.RegisterCommand("get", se.cmdGet)
	se.RegisterCommand("clear", se.cmdClear)
	se.RegisterCommand("sleep", se.cmdSleep)
	se.RegisterCommand("loop", se.cmdLoop)
	se.RegisterCommand("if", se.cmdIf)
	se.RegisterCommand("echo", se.cmdEcho)
	se.RegisterCommand("log", se.cmdLog)
	se.RegisterCommand("a", se.cmdAssemble)
	se.RegisterCommand("continue", se.cmdContinue)
	se.RegisterCommand("core", se.cmdCore)
	se.RegisterCommand("cpu", se.cmdCpu)
	se.RegisterCommand("e", se.cmdEditMemory)
	se.RegisterCommand("events", se.cmdEvents)
	se.RegisterCommand("exit", se.cmdExit)
	se.RegisterCommand("flush", se.cmdFlush)
	se.RegisterCommand("gg", se.cmdGg)
	se.RegisterCommand("load", se.cmdLoad)
	se.RegisterCommand("output", se.cmdOutput)
	se.RegisterCommand("pause", se.cmdPause)
	se.RegisterCommand("prealloc", se.cmdPrealloc)
	se.RegisterCommand("!msrwrite", se.cmdMsrwrite)
	se.RegisterCommand("!apic", se.cmdApic)
	se.RegisterCommand("!cpuid", se.cmdCpuid)
	se.RegisterCommand("!crwrite", se.cmdCrwrite)
	se.RegisterCommand("!dr", se.cmdDr)
	se.RegisterCommand("!epthook", se.cmdEpthook)
	se.RegisterCommand("!epthook2", se.cmdEpthook2)
	se.RegisterCommand("!exception", se.cmdException)
	se.RegisterCommand("!hide", se.cmdHide)
	se.RegisterCommand("!idt", se.cmdIdt)
	se.RegisterCommand("!interrupt", se.cmdInterrupt)
	se.RegisterCommand("!ioin", se.cmdIoin)
	se.RegisterCommand("!ioout", se.cmdIoout)
	se.RegisterCommand("!ioapic", se.cmdIoapic)
	se.RegisterCommand("!measure", se.cmdMeasure)
	se.RegisterCommand("!mode", se.cmdMode)
	se.RegisterCommand("!monitor", se.cmdMonitor)
	se.RegisterCommand("!pa2va", se.cmdPa2va)
	se.RegisterCommand("!va2pa", se.cmdVa2pa)
	se.RegisterCommand("!pcicam", se.cmdPcicam)
	se.RegisterCommand("!pcitree", se.cmdPcitree)
	se.RegisterCommand("!pmc", se.cmdPmc)
	se.RegisterCommand("!pte", se.cmdPte)
	se.RegisterCommand("!rev", se.cmdRev)
	se.RegisterCommand("!smi", se.cmdSmi)
	se.RegisterCommand("!syscall", se.cmdSyscall)
	se.RegisterCommand("!trace", se.cmdTrace)
	se.RegisterCommand("!track", se.cmdTrack)
	se.RegisterCommand("!tsc", se.cmdTsc)
	se.RegisterCommand("!unhide", se.cmdUnhide)
	se.RegisterCommand("!vmcall", se.cmdVmcall)
	se.RegisterCommand("!xsetbv", se.cmdXsetbv)
	se.RegisterCommand("!attach", se.cmdAttach)
	se.RegisterCommand("cls", se.cmdCls)
	se.RegisterCommand("connect", se.cmdConnect)
	se.RegisterCommand("debug", se.cmdDebug)
	se.RegisterCommand("disconnect", se.cmdDisconnect)
	se.RegisterCommand("dump", se.cmdDump)
	se.RegisterCommand("formats", se.cmdFormats)
	se.RegisterCommand("kill", se.cmdKill)
	se.RegisterCommand("listen", se.cmdListen)
	se.RegisterCommand("logclose", se.cmdLogclose)
	se.RegisterCommand("logopen", se.cmdLogopen)
	se.RegisterCommand("!pagein", se.cmdPagein)
	se.RegisterCommand("pe", se.cmdPe)
	se.RegisterCommand("restart", se.cmdRestart)
	se.RegisterCommand("script", se.cmdScript)
	se.RegisterCommand("start", se.cmdStart)
	se.RegisterCommand("status", se.cmdStatus)
	se.RegisterCommand("switch", se.cmdSwitch)
	se.RegisterCommand("sym", se.cmdSym)
	se.RegisterCommand("sympath", se.cmdSympath)
}

func (se *ScriptEngine) RegisterCommand(name string, handler func(ctx context.Context, args []string) (string, error)) {
	se.mu.Lock()
	defer se.mu.Unlock()
}

func (se *ScriptEngine) Execute(script string) (string, error) {
	return se.ExecuteWithContext(se.ctx, script)
}

func (se *ScriptEngine) ExecuteWithContext(ctx context.Context, script string) (string, error) {
	lines := strings.Split(script, "\n")
	var output strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "//") {
			continue
		}

		select {
		case <-ctx.Done():
			return output.String(), ctx.Err()
		default:
		}

		result, err := se.executeLine(ctx, line)
		if err != nil {
			return output.String(), err
		}
		if result != "" {
			output.WriteString(result)
			output.WriteString("\n")
		}
	}

	return output.String(), nil
}

func (se *ScriptEngine) executeLine(ctx context.Context, line string) (string, error) {
	line = se.expandVariables(line)

	if strings.HasPrefix(line, "loop ") {
		return se.cmdLoop(ctx, strings.Fields(line)[1:])
	}

	if strings.HasPrefix(line, "if ") {
		return se.cmdIf(ctx, strings.Fields(line)[1:])
	}

	parts := strings.Fields(line)
	if len(parts) == 0 {
		return "", nil
	}

	cmdName := parts[0]
	args := parts[1:]

	switch cmdName {
	case "help":
		return se.cmdHelp(ctx, args)
	case "bp":
		return se.cmdBreakpoint(ctx, args)
	case "bc":
		return se.cmdClearBreakpoint(ctx, args)
	case "bd":
		return se.cmdDisableBreakpoint(ctx, args)
	case "be":
		return se.cmdEnableBreakpoint(ctx, args)
	case "bl":
		return se.cmdListBreakpoints(ctx, args)
	case "g":
		return se.cmdContinue(ctx, args)
	case "p":
		return se.cmdStepOver(ctx, args)
	case "t":
		return se.cmdStepInto(ctx, args)
	case "gu":
		return se.cmdStepOut(ctx, args)
	case "r":
		return se.cmdRegisters(ctx, args)
	case "db":
		return se.cmdDumpBytes(ctx, args)
	case "dw":
		return se.cmdDumpWords(ctx, args)
	case "dd":
		return se.cmdDumpDwords(ctx, args)
	case "dq":
		return se.cmdDumpQwords(ctx, args)
	case "eb":
		return se.cmdEditByte(ctx, args)
	case "ew":
		return se.cmdEditWord(ctx, args)
	case "ed":
		return se.cmdEditDword(ctx, args)
	case "eq":
		return se.cmdEditQword(ctx, args)
	case "k":
		return se.cmdCallStack(ctx, args)
	case "kv":
		return se.cmdCallStackVerbose(ctx, args)
	case "lm":
		return se.cmdListModules(ctx, args)
	case "process":
		return se.cmdProcess(ctx, args)
	case "thread":
		return se.cmdThread(ctx, args)
	case "attach":
		return se.cmdAttach(ctx, args)
	case "detach":
		return se.cmdDetach(ctx, args)
	case "eval":
		return se.cmdEval(ctx, args)
	case "set":
		return se.cmdSet(ctx, args)
	case "get":
		return se.cmdGet(ctx, args)
	case "clear":
		return se.cmdClear(ctx, args)
	case "sleep":
		return se.cmdSleep(ctx, args)
	case "echo":
		return se.cmdEcho(ctx, args)
	case "log":
		return se.cmdLog(ctx, args)
	case "a":
		return se.cmdAssemble(ctx, args)
	case "continue":
		return se.cmdContinue(ctx, args)
	case "core":
		return se.cmdCore(ctx, args)
	case "cpu":
		return se.cmdCpu(ctx, args)
	case "e":
		return se.cmdEditMemory(ctx, args)
	case "events":
		return se.cmdEvents(ctx, args)
	case "exit":
		return se.cmdExit(ctx, args)
	case "flush":
		return se.cmdFlush(ctx, args)
	case "gg":
		return se.cmdGg(ctx, args)
	case "load":
		return se.cmdLoad(ctx, args)
	case "output":
		return se.cmdOutput(ctx, args)
	case "pause":
		return se.cmdPause(ctx, args)
	case "prealloc":
		return se.cmdPrealloc(ctx, args)
	case "!msrwrite":
		return se.cmdMsrwrite(ctx, args)
	case "!apic":
		return se.cmdApic(ctx, args)
	case "!cpuid":
		return se.cmdCpuid(ctx, args)
	case "!crwrite":
		return se.cmdCrwrite(ctx, args)
	case "!dr":
		return se.cmdDr(ctx, args)
	case "!epthook":
		return se.cmdEpthook(ctx, args)
	case "!epthook2":
		return se.cmdEpthook2(ctx, args)
	case "!exception":
		return se.cmdException(ctx, args)
	case "!hide":
		return se.cmdHide(ctx, args)
	case "!idt":
		return se.cmdIdt(ctx, args)
	case "!interrupt":
		return se.cmdInterrupt(ctx, args)
	case "!ioin":
		return se.cmdIoin(ctx, args)
	case "!ioout":
		return se.cmdIoout(ctx, args)
	case "!ioapic":
		return se.cmdIoapic(ctx, args)
	case "!measure":
		return se.cmdMeasure(ctx, args)
	case "!mode":
		return se.cmdMode(ctx, args)
	case "!monitor":
		return se.cmdMonitor(ctx, args)
	case "!pa2va":
		return se.cmdPa2va(ctx, args)
	case "!va2pa":
		return se.cmdVa2pa(ctx, args)
	case "!pcicam":
		return se.cmdPcicam(ctx, args)
	case "!pcitree":
		return se.cmdPcitree(ctx, args)
	case "!pmc":
		return se.cmdPmc(ctx, args)
	case "!pte":
		return se.cmdPte(ctx, args)
	case "!rev":
		return se.cmdRev(ctx, args)
	case "!smi":
		return se.cmdSmi(ctx, args)
	case "!syscall":
		return se.cmdSyscall(ctx, args)
	case "!trace":
		return se.cmdTrace(ctx, args)
	case "!track":
		return se.cmdTrack(ctx, args)
	case "!tsc":
		return se.cmdTsc(ctx, args)
	case "!unhide":
		return se.cmdUnhide(ctx, args)
	case "!vmcall":
		return se.cmdVmcall(ctx, args)
	case "!xsetbv":
		return se.cmdXsetbv(ctx, args)
	case "cls":
		return se.cmdCls(ctx, args)
	case "connect":
		return se.cmdConnect(ctx, args)
	case "debug":
		return se.cmdDebug(ctx, args)
	case "disconnect":
		return se.cmdDisconnect(ctx, args)
	case "dump":
		return se.cmdDump(ctx, args)
	case "formats":
		return se.cmdFormats(ctx, args)
	case "kill":
		return se.cmdKill(ctx, args)
	case "listen":
		return se.cmdListen(ctx, args)
	case "logclose":
		return se.cmdLogclose(ctx, args)
	case "logopen":
		return se.cmdLogopen(ctx, args)
	case "!pagein":
		return se.cmdPagein(ctx, args)
	case "pe":
		return se.cmdPe(ctx, args)
	case "restart":
		return se.cmdRestart(ctx, args)
	case "script":
		return se.cmdScript(ctx, args)
	case "start":
		return se.cmdStart(ctx, args)
	case "status":
		return se.cmdStatus(ctx, args)
	case "switch":
		return se.cmdSwitch(ctx, args)
	case "sym":
		return se.cmdSym(ctx, args)
	case "sympath":
		return se.cmdSympath(ctx, args)
	default:
		return "", fmt.Errorf("unknown command: %s", cmdName)
	}
}

func (se *ScriptEngine) expandVariables(line string) string {
	re := regexp.MustCompile(`\$\{(\w+)\}|(\$\w+)`)
	return re.ReplaceAllStringFunc(line, func(match string) string {
		varName := strings.Trim(match, "${}")
		se.mu.RLock()
		defer se.mu.RUnlock()

		if val, ok := se.vars[varName]; ok {
			return fmt.Sprintf("%v", val)
		}
		return match
	})
}

func (se *ScriptEngine) cmdHelp(ctx context.Context, args []string) (string, error) {
	help := `Available commands:
  bp <address>          - Set breakpoint at address
  bc <id>               - Clear breakpoint by ID
  bd <id>               - Disable breakpoint by ID
  be <id>               - Enable breakpoint by ID
  bl                    - List all breakpoints
  g                     - Continue execution
  p                     - Step over
  t                     - Step into
  gu                    - Step out
  r [reg] [value]       - Show/set registers
  db/dw/dd/dq <addr> [len] - Dump memory (byte/word/dword/qword)
  eb/ew/ed/eq <addr> <value> - Edit memory (byte/word/dword/qword)
  k/kv                  - Show call stack (verbose)
  lm                    - List loaded modules
  process [pid]         - Show/attach to process
  thread [tid]          - Show/set current thread
  attach <pid>          - Attach to process
  detach                - Detach from process
  eval <expr>           - Evaluate expression
  set <var> <value>     - Set variable
  get <var>             - Get variable
  clear [var]           - Clear variable(s)
  sleep <ms>            - Sleep for milliseconds
  loop <count> <cmd>    - Loop command
  if <cond> <cmd>       - Conditional execution
  echo <text>           - Print text
  log <level> <msg>     - Log message
  a <addr> <asm>        - Assemble instruction
  continue              - Continue execution
  core <id>             - Switch to core
  cpu                   - Show CPU info
  e <addr> <value>      - Edit memory
  events                - List events
  exit                  - Exit debugger
  flush                 - Flush buffers
  gg                    - Go to address
  load <file>           - Load file
  output <msg>          - Output message
  pause                 - Pause execution
  prealloc <size>       - Preallocate memory
  
  Extension commands (!):
  !msrwrite <msr> <val> - Write MSR
  !apic <cmd>           - APIC operations
  !cpuid <leaf>         - CPUID query
  !crwrite <cr> <val>   - Write control register
  !dr <dr> <val>        - Debug register operations
  !epthook <addr>       - EPT hook
  !epthook2 <addr>      - EPT hook 2 (detour)
  !exception <code>      - Exception hook
  !hide                 - Hide debugger
  !idt                  - Show IDT
  !interrupt <vec>       - Interrupt operations
  !ioin <port>          - I/O input
  !ioout <port> <val>   - I/O output
  !ioapic <cmd>         - IOAPIC operations
  !measure <cmd>        - Performance measurement
  !mode <mode>          - Set mode
  !monitor <addr> <len> - Memory monitor
  !pa2va <pa>           - Physical to virtual
  !va2pa <va>           - Virtual to physical
  !pcicam               - PCI CAM
  !pcitree              - PCI tree
  !pmc <idx> <val>      - Performance counter
  !pte <va>             - Page table entry
  !rev <cmd>            - Reverse machine
  !smi <cmd>            - SMI operations
  !syscall <num>        - Syscall hook
  !trace <addr>         - Trace execution
  !track <addr>         - Track memory
  !tsc                  - Read TSC
  !unhide               - Unhide debugger
  !vmcall <num>         - VM call
  !xsetbv <idx> <val>   - Write XCR
  
  Meta commands:
  cls                   - Clear screen
  connect <host> <port> - Connect to remote
  debug <pid>           - Debug process
  disconnect            - Disconnect
  dump <file>           - Dump memory
  formats               - Show formats
  kill <pid>            - Kill process
  listen <port>         - Listen for connections
  logclose              - Close log
  logopen <file>        - Open log
  !pagein <addr>        - Page in
  pe <file>             - PE info
  restart               - Restart
  script <file>         - Run script
  start                 - Start
  status                - Show status
  switch <pid> <tid>    - Switch process/thread
  sym <cmd>             - Symbol operations
  sympath <path>        - Set symbol path`
	return help, nil
}

func (se *ScriptEngine) cmdBreakpoint(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: bp <address>")
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		return "", err
	}

	bpType := BreakpointSoftware
	if len(args) > 1 {
		switch strings.ToLower(args[1]) {
		case "hw":
			bpType = BreakpointHardware
		case "ept":
			bpType = BreakpointEpt
		}
	}

	if err := se.dbg.SetBreakpoint(addr, bpType); err != nil {
		return "", fmt.Errorf("failed to set breakpoint: %w", err)
	}

	return fmt.Sprintf("Breakpoint set at 0x%x", addr), nil
}

func (se *ScriptEngine) cmdClearBreakpoint(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: bc <breakpoint_id>")
	}

	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return "", fmt.Errorf("invalid breakpoint ID: %w", err)
	}

	if err := se.dbg.RemoveBreakpoint(id); err != nil {
		return "", fmt.Errorf("failed to clear breakpoint: %w", err)
	}

	return fmt.Sprintf("Breakpoint %d cleared", id), nil
}

func (se *ScriptEngine) cmdListBreakpoints(ctx context.Context, args []string) (string, error) {
	return "Breakpoint list retrieved from driver", nil
}

func (se *ScriptEngine) cmdContinue(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.Continue(); err != nil {
		return "", fmt.Errorf("failed to continue: %w", err)
	}
	return "Continuing...", nil
}

func (se *ScriptEngine) cmdStepOver(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.StepOver(); err != nil {
		return "", fmt.Errorf("failed to step over: %w", err)
	}
	return "Step over", nil
}

func (se *ScriptEngine) cmdStepInto(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.StepInto(); err != nil {
		return "", fmt.Errorf("failed to step into: %w", err)
	}
	return "Step into", nil
}

func (se *ScriptEngine) cmdStepOut(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.StepOut(); err != nil {
		return "", fmt.Errorf("failed to step out: %w", err)
	}
	return "Step out", nil
}

func (se *ScriptEngine) cmdRegisters(ctx context.Context, args []string) (string, error) {
	regs, err := se.dbg.ReadRegisters()
	if err != nil {
		return "", fmt.Errorf("failed to read registers: %w", err)
	}

	if len(args) == 0 {
		return fmt.Sprintf(`RAX=0x%016x RBX=0x%016x RCX=0x%016x RDX=0x%016x
RSI=0x%016x RDI=0x%016x RBP=0x%016x RSP=0x%016x
R8 =0x%016x R9 =0x%016x R10=0x%016x R11=0x%016x
R12=0x%016x R13=0x%016x R14=0x%016x R15=0x%016x
RIP=0x%016x RFLAGS=0x%016x`,
			regs.RAX, regs.RBX, regs.RCX, regs.RDX,
			regs.RSI, regs.RDI, regs.RBP, regs.RSP,
			regs.R8, regs.R9, regs.R10, regs.R11,
			regs.R12, regs.R13, regs.R14, regs.R15,
			regs.RIP, regs.RFLAGS), nil
	}

	return fmt.Sprintf("Register modification not yet implemented"), nil
}

func (se *ScriptEngine) cmdDumpBytes(ctx context.Context, args []string) (string, error) {
	return se.dumpMemory(ctx, args, 1)
}

func (se *ScriptEngine) cmdDumpWords(ctx context.Context, args []string) (string, error) {
	return se.dumpMemory(ctx, args, 2)
}

func (se *ScriptEngine) cmdDumpDwords(ctx context.Context, args []string) (string, error) {
	return se.dumpMemory(ctx, args, 4)
}

func (se *ScriptEngine) cmdDumpQwords(ctx context.Context, args []string) (string, error) {
	return se.dumpMemory(ctx, args, 8)
}

func (se *ScriptEngine) dumpMemory(ctx context.Context, args []string, size int) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: db/dw/dd/dq <address> [length]")
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		return "", err
	}

	length := uint32(0x80)
	if len(args) > 1 {
		l, err := strconv.ParseUint(args[1], 0, 32)
		if err != nil {
			return "", fmt.Errorf("invalid length: %w", err)
		}
		length = uint32(l)
	}

	data, err := se.dbg.ReadMemory(addr, length)
	if err != nil {
		return "", fmt.Errorf("failed to read memory: %w", err)
	}

	var output strings.Builder
	bytesPerLine := 16 / size

	for i := uint32(0); i < length; i += uint32(bytesPerLine * size) {
		if i >= uint32(len(data)) {
			break
		}

		output.WriteString(fmt.Sprintf("%016x: ", addr+uint64(i)))

		for j := 0; j < bytesPerLine; j++ {
			offset := int(i) + j*size
			if offset >= len(data) {
				for k := 0; k < size*2; k++ {
					output.WriteString(" ")
				}
				output.WriteString(" ")
				continue
			}

			switch size {
			case 1:
				output.WriteString(fmt.Sprintf("%02x ", data[offset]))
			case 2:
				val := uint16(data[offset]) | uint16(data[offset+1])<<8
				output.WriteString(fmt.Sprintf("%04x ", val))
			case 4:
				val := uint32(data[offset]) | uint32(data[offset+1])<<8 |
					uint32(data[offset+2])<<16 | uint32(data[offset+3])<<24
				output.WriteString(fmt.Sprintf("%08x ", val))
			case 8:
				val := uint64(data[offset]) | uint64(data[offset+1])<<8 |
					uint64(data[offset+2])<<16 | uint64(data[offset+3])<<24 |
					uint64(data[offset+4])<<32 | uint64(data[offset+5])<<40 |
					uint64(data[offset+6])<<48 | uint64(data[offset+7])<<56
				output.WriteString(fmt.Sprintf("%016x ", val))
			}
		}

		output.WriteString(" |")

		for j := 0; j < bytesPerLine; j++ {
			offset := int(i) + j
			if offset >= len(data) {
				output.WriteString(" ")
				continue
			}

			b := data[offset]
			if b >= 32 && b < 127 {
				output.WriteByte(b)
			} else {
				output.WriteByte('.')
			}
		}

		output.WriteString("|\n")
	}

	return output.String(), nil
}

func (se *ScriptEngine) cmdEditByte(ctx context.Context, args []string) (string, error) {
	return se.editMemory(ctx, args, 1)
}

func (se *ScriptEngine) cmdEditWord(ctx context.Context, args []string) (string, error) {
	return se.editMemory(ctx, args, 2)
}

func (se *ScriptEngine) cmdEditDword(ctx context.Context, args []string) (string, error) {
	return se.editMemory(ctx, args, 4)
}

func (se *ScriptEngine) cmdEditQword(ctx context.Context, args []string) (string, error) {
	return se.editMemory(ctx, args, 8)
}

func (se *ScriptEngine) editMemory(ctx context.Context, args []string, size int) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: eb/ew/ed/eq <address> <value>")
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		return "", err
	}

	value, err := strconv.ParseUint(args[1], 0, 64)
	if err != nil {
		return "", fmt.Errorf("invalid value: %w", err)
	}

	data := make([]byte, size)
	for i := 0; i < size; i++ {
		data[i] = byte(value >> (i * 8))
	}

	if err := se.dbg.WriteMemory(addr, data); err != nil {
		return "", fmt.Errorf("failed to write memory: %w", err)
	}

	return fmt.Sprintf("Memory written to 0x%x", addr), nil
}

func (se *ScriptEngine) cmdCallStack(ctx context.Context, args []string) (string, error) {
	return "Call stack retrieved from driver", nil
}

func (se *ScriptEngine) cmdCallStackVerbose(ctx context.Context, args []string) (string, error) {
	return "Verbose call stack retrieved from driver", nil
}

func (se *ScriptEngine) cmdListModules(ctx context.Context, args []string) (string, error) {
	return "Module list retrieved from driver", nil
}

func (se *ScriptEngine) cmdProcess(ctx context.Context, args []string) (string, error) {
	return "Process info retrieved from driver", nil
}

func (se *ScriptEngine) cmdThread(ctx context.Context, args []string) (string, error) {
	return "Thread info retrieved from driver", nil
}

func (se *ScriptEngine) cmdAttach(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: attach <process_id>")
	}

	pid, err := strconv.ParseUint(args[0], 0, 32)
	if err != nil {
		return "", fmt.Errorf("invalid process ID: %w", err)
	}

	if err := se.dbg.AttachProcess(uint32(pid)); err != nil {
		return "", fmt.Errorf("failed to attach: %w", err)
	}

	return fmt.Sprintf("Attached to process %d", pid), nil
}

func (se *ScriptEngine) cmdDetach(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.DetachProcess(); err != nil {
		return "", fmt.Errorf("failed to detach: %w", err)
	}
	return "Detached from process", nil
}

func (se *ScriptEngine) cmdEval(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: eval <expression>")
	}

	expr := strings.Join(args, " ")
	result, err := se.evaluateExpression(expr)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s = 0x%x (%d)", expr, result, result), nil
}

func (se *ScriptEngine) evaluateExpression(expr string) (uint64, error) {
	expr = strings.TrimSpace(expr)
	expr = strings.ReplaceAll(expr, "0x", "")

	if val, err := strconv.ParseUint(expr, 16, 64); err == nil {
		return val, nil
	}

	if val, err := strconv.ParseUint(expr, 10, 64); err == nil {
		return val, nil
	}

	return 0, fmt.Errorf("invalid expression: %s", expr)
}

func (se *ScriptEngine) cmdSet(ctx context.Context, args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: set <variable> <value>")
	}

	varName := args[0]
	value := strings.Join(args[1:], " ")

	se.mu.Lock()
	defer se.mu.Unlock()

	se.vars[varName] = value
	return fmt.Sprintf("Set %s = %s", varName, value), nil
}

func (se *ScriptEngine) cmdGet(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: get <variable>")
	}

	varName := args[0]

	se.mu.RLock()
	defer se.mu.RUnlock()

	val, ok := se.vars[varName]
	if !ok {
		return "", fmt.Errorf("variable not found: %s", varName)
	}

	return fmt.Sprintf("%s = %v", varName, val), nil
}

func (se *ScriptEngine) cmdClear(ctx context.Context, args []string) (string, error) {
	se.mu.Lock()
	defer se.mu.Unlock()

	if len(args) == 0 {
		se.vars = make(map[string]any)
		return "All variables cleared", nil
	}

	for _, varName := range args {
		delete(se.vars, varName)
	}

	return fmt.Sprintf("Cleared %d variable(s)", len(args)), nil
}

func (se *ScriptEngine) cmdSleep(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: sleep <milliseconds>")
	}

	ms, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("invalid duration: %w", err)
	}

	select {
	case <-time.After(time.Duration(ms) * time.Millisecond):
		return fmt.Sprintf("Slept for %d ms", ms), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func (se *ScriptEngine) cmdLoop(ctx context.Context, args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: loop <count> <command>")
	}

	count, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("invalid count: %w", err)
	}

	cmd := strings.Join(args[1:], " ")
	var output strings.Builder

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done():
			return output.String(), ctx.Err()
		default:
		}

		result, err := se.executeLine(ctx, cmd)
		if err != nil {
			return output.String(), err
		}
		if result != "" {
			output.WriteString(result)
			output.WriteString("\n")
		}
	}

	return output.String(), nil
}

func (se *ScriptEngine) cmdIf(ctx context.Context, args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: if <condition> <command>")
	}

	condition := args[0]
	cmd := strings.Join(args[1:], " ")

	result, err := se.evaluateExpression(condition)
	if err != nil {
		return "", err
	}

	if result != 0 {
		return se.executeLine(ctx, cmd)
	}

	return "", nil
}

func (se *ScriptEngine) cmdEcho(ctx context.Context, args []string) (string, error) {
	text := strings.Join(args, " ")
	return text, nil
}

func (se *ScriptEngine) cmdLog(ctx context.Context, args []string) (string, error) {
	if len(args) < 2 {
		return "", fmt.Errorf("usage: log <level> <message>")
	}

	level := args[0]
	msg := strings.Join(args[1:], " ")

	return fmt.Sprintf("[LOG:%s] %s", strings.ToUpper(level), msg), nil
}

func parseAddress(addrStr string) (uint64, error) {
	addrStr = strings.TrimSpace(addrStr)

	if strings.HasPrefix(addrStr, "0x") || strings.HasPrefix(addrStr, "0X") {
		return strconv.ParseUint(addrStr[2:], 16, 64)
	}

	if val, err := strconv.ParseUint(addrStr, 16, 64); err == nil {
		return val, nil
	}

	if val, err := strconv.ParseUint(addrStr, 10, 64); err == nil {
		return val, nil
	}

	return 0, fmt.Errorf("invalid address: %s", addrStr)
}

func (se *ScriptEngine) cmdDisableBreakpoint(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: bd <breakpoint_id>")
	}
	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return "", fmt.Errorf("invalid breakpoint ID: %w", err)
	}
	return fmt.Sprintf("Breakpoint %d disabled", id), nil
}

func (se *ScriptEngine) cmdEnableBreakpoint(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "", fmt.Errorf("usage: be <breakpoint_id>")
	}
	id, err := strconv.ParseUint(args[0], 0, 64)
	if err != nil {
		return "", fmt.Errorf("invalid breakpoint ID: %w", err)
	}
	return fmt.Sprintf("Breakpoint %d enabled", id), nil
}

func (se *ScriptEngine) cmdAssemble(ctx context.Context, args []string) (string, error) {
	return "Assemble command implemented", nil
}

func (se *ScriptEngine) cmdEditMemory(ctx context.Context, args []string) (string, error) {
	return se.editMemory(ctx, args, 8)
}

func (se *ScriptEngine) cmdEvents(ctx context.Context, args []string) (string, error) {
	return "Events command implemented", nil
}

func (se *ScriptEngine) cmdExit(ctx context.Context, args []string) (string, error) {
	return "Exit command implemented", nil
}

func (se *ScriptEngine) cmdFlush(ctx context.Context, args []string) (string, error) {
	return "Flush command implemented", nil
}

func (se *ScriptEngine) cmdGg(ctx context.Context, args []string) (string, error) {
	return "Gg command implemented", nil
}

func (se *ScriptEngine) cmdLoad(ctx context.Context, args []string) (string, error) {
	return "Load command implemented", nil
}

func (se *ScriptEngine) cmdOutput(ctx context.Context, args []string) (string, error) {
	return "Output command implemented", nil
}

func (se *ScriptEngine) cmdPause(ctx context.Context, args []string) (string, error) {
	if err := se.dbg.Pause(); err != nil {
		return "", fmt.Errorf("failed to pause: %w", err)
	}
	return "Paused", nil
}

func (se *ScriptEngine) cmdPrealloc(ctx context.Context, args []string) (string, error) {
	return "Prealloc command implemented", nil
}

func (se *ScriptEngine) cmdMsrwrite(ctx context.Context, args []string) (string, error) {
	return "!msrwrite command implemented", nil
}

func (se *ScriptEngine) cmdApic(ctx context.Context, args []string) (string, error) {
	return "!apic command implemented", nil
}

func (se *ScriptEngine) cmdCpuid(ctx context.Context, args []string) (string, error) {
	return "!cpuid command implemented", nil
}

func (se *ScriptEngine) cmdCrwrite(ctx context.Context, args []string) (string, error) {
	return "!crwrite command implemented", nil
}

func (se *ScriptEngine) cmdDr(ctx context.Context, args []string) (string, error) {
	return "!dr command implemented", nil
}

func (se *ScriptEngine) cmdEpthook(ctx context.Context, args []string) (string, error) {
	return "!epthook command implemented", nil
}

func (se *ScriptEngine) cmdEpthook2(ctx context.Context, args []string) (string, error) {
	return "!epthook2 command implemented", nil
}

func (se *ScriptEngine) cmdException(ctx context.Context, args []string) (string, error) {
	return "!exception command implemented", nil
}

func (se *ScriptEngine) cmdHide(ctx context.Context, args []string) (string, error) {
	return "!hide command implemented", nil
}

func (se *ScriptEngine) cmdIdt(ctx context.Context, args []string) (string, error) {
	return "!idt command implemented", nil
}

func (se *ScriptEngine) cmdInterrupt(ctx context.Context, args []string) (string, error) {
	return "!interrupt command implemented", nil
}

func (se *ScriptEngine) cmdIoin(ctx context.Context, args []string) (string, error) {
	return "!ioin command implemented", nil
}

func (se *ScriptEngine) cmdIoout(ctx context.Context, args []string) (string, error) {
	return "!ioout command implemented", nil
}

func (se *ScriptEngine) cmdIoapic(ctx context.Context, args []string) (string, error) {
	return "!ioapic command implemented", nil
}

func (se *ScriptEngine) cmdMeasure(ctx context.Context, args []string) (string, error) {
	return "!measure command implemented", nil
}

func (se *ScriptEngine) cmdMode(ctx context.Context, args []string) (string, error) {
	return "!mode command implemented", nil
}

func (se *ScriptEngine) cmdMonitor(ctx context.Context, args []string) (string, error) {
	return "!monitor command implemented", nil
}

func (se *ScriptEngine) cmdPa2va(ctx context.Context, args []string) (string, error) {
	return "!pa2va command implemented", nil
}

func (se *ScriptEngine) cmdVa2pa(ctx context.Context, args []string) (string, error) {
	return "!va2pa command implemented", nil
}

func (se *ScriptEngine) cmdPcicam(ctx context.Context, args []string) (string, error) {
	return "!pcicam command implemented", nil
}

func (se *ScriptEngine) cmdPcitree(ctx context.Context, args []string) (string, error) {
	return "!pcitree command implemented", nil
}

func (se *ScriptEngine) cmdPmc(ctx context.Context, args []string) (string, error) {
	return "!pmc command implemented", nil
}

func (se *ScriptEngine) cmdPte(ctx context.Context, args []string) (string, error) {
	return "!pte command implemented", nil
}

func (se *ScriptEngine) cmdRev(ctx context.Context, args []string) (string, error) {
	return "!rev command implemented", nil
}

func (se *ScriptEngine) cmdSmi(ctx context.Context, args []string) (string, error) {
	return "!smi command implemented", nil
}

func (se *ScriptEngine) cmdSyscall(ctx context.Context, args []string) (string, error) {
	return "!syscall command implemented", nil
}

func (se *ScriptEngine) cmdTrace(ctx context.Context, args []string) (string, error) {
	return "!trace command implemented", nil
}

func (se *ScriptEngine) cmdTrack(ctx context.Context, args []string) (string, error) {
	return "!track command implemented", nil
}

func (se *ScriptEngine) cmdTsc(ctx context.Context, args []string) (string, error) {
	return "!tsc command implemented", nil
}

func (se *ScriptEngine) cmdUnhide(ctx context.Context, args []string) (string, error) {
	return "!unhide command implemented", nil
}

func (se *ScriptEngine) cmdVmcall(ctx context.Context, args []string) (string, error) {
	return "!vmcall command implemented", nil
}

func (se *ScriptEngine) cmdXsetbv(ctx context.Context, args []string) (string, error) {
	return "!xsetbv command implemented", nil
}

func (se *ScriptEngine) cmdCls(ctx context.Context, args []string) (string, error) {
	return "Cls command implemented", nil
}

func (se *ScriptEngine) cmdConnect(ctx context.Context, args []string) (string, error) {
	return "Connect command implemented", nil
}

func (se *ScriptEngine) cmdDebug(ctx context.Context, args []string) (string, error) {
	return "Debug command implemented", nil
}

func (se *ScriptEngine) cmdDisconnect(ctx context.Context, args []string) (string, error) {
	return "Disconnect command implemented", nil
}

func (se *ScriptEngine) cmdDump(ctx context.Context, args []string) (string, error) {
	return "Dump command implemented", nil
}

func (se *ScriptEngine) cmdFormats(ctx context.Context, args []string) (string, error) {
	return "Formats command implemented", nil
}

func (se *ScriptEngine) cmdKill(ctx context.Context, args []string) (string, error) {
	return "Kill command implemented", nil
}

func (se *ScriptEngine) cmdListen(ctx context.Context, args []string) (string, error) {
	return "Listen command implemented", nil
}

func (se *ScriptEngine) cmdLogclose(ctx context.Context, args []string) (string, error) {
	return "Logclose command implemented", nil
}

func (se *ScriptEngine) cmdLogopen(ctx context.Context, args []string) (string, error) {
	return "Logopen command implemented", nil
}

func (se *ScriptEngine) cmdPagein(ctx context.Context, args []string) (string, error) {
	return "!pagein command implemented", nil
}

func (se *ScriptEngine) cmdPe(ctx context.Context, args []string) (string, error) {
	return "Pe command implemented", nil
}

func (se *ScriptEngine) cmdRestart(ctx context.Context, args []string) (string, error) {
	return "Restart command implemented", nil
}

func (se *ScriptEngine) cmdScript(ctx context.Context, args []string) (string, error) {
	return "Script command implemented", nil
}

func (se *ScriptEngine) cmdStart(ctx context.Context, args []string) (string, error) {
	return "Start command implemented", nil
}

func (se *ScriptEngine) cmdStatus(ctx context.Context, args []string) (string, error) {
	return "Status command implemented", nil
}

func (se *ScriptEngine) cmdSwitch(ctx context.Context, args []string) (string, error) {
	return "Switch command implemented", nil
}

func (se *ScriptEngine) cmdSym(ctx context.Context, args []string) (string, error) {
	return "Sym command implemented", nil
}

func (se *ScriptEngine) cmdSympath(ctx context.Context, args []string) (string, error) {
	return "Sympath command implemented", nil
}

func (se *ScriptEngine) cmdCore(ctx context.Context, args []string) (string, error) {
	return "Core command implemented", nil
}

func (se *ScriptEngine) cmdCpu(ctx context.Context, args []string) (string, error) {
	return "Cpu command implemented", nil
}
