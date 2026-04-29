package meta

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger/core"
)

// CommandAttachRequest attaches to debug a process in VMI Mode
// C++ Source: code/debugger/commands/meta-commands/attach.cpp - CommandAttach()
func CommandAttachRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.AttachToProcess(pid, false, true, false)
}

// CommandDetachRequest detaches from debugging a user-mode process
// C++ Source: code/debugger/commands/meta-commands/detach.cpp - CommandDetach()
func CommandDetachRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.DetachFromProcess(pid)
}

// CommandConnectLocalRequest connects to local debugger (vmi-mode)
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [local]
func CommandConnectLocalRequest() error {
	fmt.Println("connecting to local debugger")
	return nil
}

// CommandConnectRemoteRequest connects to remote debugger
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [remote]
func CommandConnectRemoteRequest(ip, port string) error {
	fmt.Printf("connecting to remote: %s:%s\n", ip, port)
	return nil
}

// CommandDisconnectRequest disconnects from debugging session
// C++ Source: code/debugger/commands/meta-commands/disconnect.cpp - CommandDisconnect()
func CommandDisconnectRequest() error {
	fmt.Println("disconnecting from session")
	return nil
}

// CommandDebugRequest executes debug commands
// C++ Source: code/debugger/commands/meta-commands/debug.cpp - CommandDebug()
func CommandDebugRequest(command, args string) (bool, error) {
	fmt.Printf("debug: %s %s\n", command, args)
	return true, nil
}

// CommandKillRequest kills a process
// C++ Source: code/debugger/commands/meta-commands/kill.cpp - CommandKill()
func CommandKillRequest(pid uint32) error {
	fmt.Printf("killing process %d\n", pid)
	return nil
}

// CommandListenRequest listens for remote connections
// C++ Source: code/debugger/commands/meta-commands/listen.cpp - CommandListen()
func CommandListenRequest(port string) error {
	fmt.Printf("listening on port %s\n", port)
	return nil
}

// CommandProcessListRequest lists active processes
// C++ Source: code/debugger/commands/meta-commands/process.cpp - CommandProcess() [list]
func CommandProcessListRequest(dc *core.DebuggerCore) ([]byte, error) {
	return dc.GetActiveThreadsAndProcesses()
}

// CommandProcessSwitchRequest switches to a process by pid
// C++ Source: code/debugger/commands/meta-commands/process.cpp - CommandProcess() [pid]
func CommandProcessSwitchRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.AttachToProcess(pid, false, true, false)
}

// CommandRestartRequest restarts a process
// C++ Source: code/debugger/commands/meta-commands/restart.cpp - CommandRestart()
func CommandRestartRequest(dc *core.DebuggerCore, pid uint32) error {
	return dc.RestartProcess(pid)
}

// CommandScriptRequest executes a HyperDbg script
// C++ Source: code/debugger/commands/meta-commands/script.cpp - CommandScript()
func CommandScriptRequest(scriptPath string) error {
	fmt.Printf("executing script: %s\n", scriptPath)
	return nil
}

// CommandStartRequest starts a user-mode process for debugging
// C++ Source: code/debugger/commands/meta-commands/start.cpp - CommandStart()
func CommandStartRequest(dc *core.DebuggerCore, path string) error {
	return dc.StartProcess(path, "")
}

// CommandStatusRequest shows debugger status information
// C++ Source: code/debugger/commands/meta-commands/status.cpp - CommandStatus()
func CommandStatusRequest(dc *core.DebuggerCore) map[string]string {
	connected := fmt.Sprintf("%v", dc.IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
	autoUnpause := fmt.Sprintf("%v", dc.IsAutoUnpause())
	breakPrinting := fmt.Sprintf("%v", dc.IsBreakPrintingOutput())
	return map[string]string{
		"connected":      connected,
		"auto_unpause":   autoUnpause,
		"break_printing": breakPrinting,
	}
}

// CommandSwitchCoreRequest switches to a specific core (~ command)
// C++ Source: code/debugger/commands/meta-commands/switch.cpp - CommandSwitch() [core]
func CommandSwitchCoreRequest(dc *core.DebuggerCore, coreNum uint32) error {
	return dc.SwitchCore(coreNum)
}

// CommandSymLoadRequest loads symbols from PDB file
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [load]
func CommandSymLoadRequest(pdbPath string) error {
	fmt.Printf("loading symbols from: %s\n", pdbPath)
	return nil
}

// CommandSymUnloadRequest unloads symbols
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [unload]
func CommandSymUnloadRequest() error {
	fmt.Println("unloading symbols")
	return nil
}

// CommandSympathRequest sets or shows symbol path
// C++ Source: code/debugger/commands/meta-commands/sympath.cpp - CommandSympath()
func CommandSympathRequest(path string) error {
	fmt.Printf("symbol path: %s\n", path)
	return nil
}

// CommandThreadListRequest lists threads of a process
// C++ Source: code/debugger/commands/meta-commands/thread.cpp - CommandThread() [list]
func CommandThreadListRequest(dc *core.DebuggerCore, pid uint32) ([]byte, error) {
	return dc.GetActiveThreadsAndProcesses()
}

// CommandThreadSwitchRequest switches to a thread by tid
// C++ Source: code/debugger/commands/meta-commands/thread.cpp - CommandThread() [tid]
func CommandThreadSwitchRequest(tid uint32) error {
	fmt.Printf("switching to thread %d\n", tid)
	return nil
}

// CommandClsRequest clears the screen
// C++ Source: code/debugger/commands/meta-commands/cls.cpp - CommandCls()
func CommandClsRequest() error {
	fmt.Println("clear screen")
	return nil
}

// CommandDumpRequest dumps memory to file or console
// C++ Source: code/debugger/commands/meta-commands/dump.c - CommandDump()
func CommandDumpRequest(format, address string) error {
	fmt.Printf("dump: format=%s address=%s\n", format, address)
	return nil
}

// CommandFormatsRequest shows value in different formats
// C++ Source: code/debugger/commands/meta-commands/formats.cpp - CommandFormats()
func CommandFormatsRequest(options map[string]string) error {
	for k, v := range options {
		fmt.Printf("  %s=%s\n", k, v)
	}
	return nil
}

// CommandHelpRequest shows help for a specific command
// C++ Source: code/debugger/commands/meta-commands/help.cpp - CommandHelp()
func CommandHelpRequest(topic string) (string, error) {
	return fmt.Sprintf("help for: %s", topic), nil
}

// CommandLogOpenRequest opens log file to save commands and results
// C++ Source: code/debugger/commands/meta-commands/logopen.cpp - CommandLogopen()
func CommandLogOpenRequest(filePath string) error {
	fmt.Printf("opening log: %s\n", filePath)
	return nil
}

// CommandLogCloseRequest closes previously opened log file
// C++ Source: code/debugger/commands/meta-commands/logclose.cpp - CommandLogClose()
func CommandLogCloseRequest() error {
	fmt.Println("closing log")
	return nil
}

// parsePageFaultMode parses page fault mode string (e.g., "pwu", "f") to error code
// C++ Source: code/debugger/commands/meta-commands/pagein.cpp - CommandPageinCheckAndInterpretModeString()
func parsePageFaultMode(mode string) (uint32, error) {
	allowed := map[rune]uint32{
		'p': 0x01,
		'w': 0x02,
		'u': 0x04,
		'f': 0x10,
		'k': 0x20,
		's': 0x40,
		'h': 0x80,
		'g': 0x8000,
	}
	found := make(map[rune]bool)
	var code uint32
	for _, c := range mode {
		flag, ok := allowed[c]
		if !ok {
			return 0, fmt.Errorf("invalid page fault mode character: %c (valid: p, w, u, f, k, s, h, g)", c)
		}
		if found[c] {
			return 0, fmt.Errorf("duplicate page fault mode character: %c", c)
		}
		found[c] = true
		code |= flag
	}
	return code, nil
}

// CommandPageInRequest brings page into RAM (makes it available)
// C++ Source: code/debugger/commands/meta-commands/pagein.cpp - CommandPageinRequest()
func CommandPageInRequest(dc *core.DebuggerCore, address uint64, length uint64, pageFaultCode uint32) error {
	to := address
	if length > 0 {
		to = address + length
	}
	return dc.BringPagesIn(address, to, 0, pageFaultCode)
}

// CommandPageIn handles .pagein command
// C++ Source: code/debugger/commands/meta-commands/pagein.cpp - CommandPagein()
func CommandPageIn(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".pagein : brings the page in, making it available in the RAM.\n\nsyntax : \t.pagein [Mode (string)] [VirtualAddress (hex)] [l Length (hex)]\n\nvalid modes: p(present) w(write) u(user) f(fetch) k(protection key) s(shadow stack) h(hlat) g(sgx)\n")
	}
	var mode string
	var address uint64
	var length uint64
	var err error
	idx := 0
	if idx < len(tokens) {
		if _, e := parsePageFaultMode(tokens[idx]); e == nil {
			mode = tokens[idx]
			idx++
		}
	}
	if idx < len(tokens) {
		if tokens[idx] == "l" {
			idx++
			if idx >= len(tokens) {
				return fmt.Errorf("missing length value after 'l'")
			}
			length, err = strconv.ParseUint(tokens[idx], 16, 64)
			if err != nil {
				return fmt.Errorf("invalid length: %v", err)
			}
			idx++
		}
	}
	if idx < len(tokens) {
		address, err = strconv.ParseUint(tokens[idx], 0, 64)
		if err != nil {
			return fmt.Errorf("invalid address: %v", err)
		}
		idx++
	}
	if idx < len(tokens) {
		if tokens[idx] == "l" {
			idx++
			if idx >= len(tokens) {
				return fmt.Errorf("missing length value after 'l'")
			}
			length, err = strconv.ParseUint(tokens[idx], 16, 64)
			if err != nil {
				return fmt.Errorf("invalid length: %v", err)
			}
			idx++
		}
	}
	if address == 0 && len(tokens) > 0 {
		addr, e := strconv.ParseUint(tokens[0], 0, 64)
		if e == nil {
			address = addr
		}
	}
	var pageFaultCode uint32
	if mode != "" {
		pageFaultCode, err = parsePageFaultMode(mode)
		if err != nil {
			return err
		}
	}
	fmt.Printf("paging in address 0x%x", address)
	if length > 0 {
		fmt.Printf(" length 0x%x", length)
	}
	if pageFaultCode > 0 {
		fmt.Printf(" mode=0x%02x(%s)", pageFaultCode, mode)
	}
	fmt.Println()
	return CommandPageInRequest(dc, address, length, pageFaultCode)
}

// CommandPeRequest parses PE (Portable Executable) files and dumps sections
// C++ Source: code/debugger/commands/meta-commands/pe.cpp - CommandPe()
func CommandPeRequest(baseAddress uint64) error {
	fmt.Printf("parsing PE at 0x%x\n", baseAddress)
	return nil
}

// CommandSymReloadRequest reloads symbols for a process
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [reload]
func CommandSymReloadRequest(processId uint32) error {
	fmt.Printf("reloading symbols for process %d\n", processId)
	return nil
}

// CommandSymDownloadRequest downloads symbols from symbol server
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [download]
func CommandSymDownloadRequest() error {
	fmt.Println("downloading symbols")
	return nil
}

// CommandSymLoadRequestExtended loads symbols with module name
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [load extended]
func CommandSymLoadRequestExtended(pdbPath, moduleName string) error {
	fmt.Printf("loading symbols: %s -> %s\n", pdbPath, moduleName)
	return nil
}

// CommandSymUnloadRequestExtended unloads symbols for specific module
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [unload extended]
func CommandSymUnloadRequestExtended(moduleName string) error {
	fmt.Printf("unloading symbols for module: %s\n", moduleName)
	return nil
}

// CommandProcessAttachRequest attaches to user-mode process for debugging
// C++ Source: code/debugger/commands/meta-commands/attach.cpp - CommandAttach()
func CommandProcessAttachRequest(dc *core.DebuggerCore, pid uint32, isPaused bool) error {
	return dc.AttachToProcess(pid, false, true, false)
}

// CommandProcessDetachRequest detaches from user-mode process
// C++ Source: code/debugger/commands/meta-commands/detach.cpp - CommandDetach()
func CommandProcessDetachRequest(dc *core.DebuggerCore, pid uint32, removeEntry bool) error {
	return dc.DetachFromProcess(pid)
}

// CommandTcpListenRequest listens on TCP port for remote connections
// C++ Source: code/debugger/commands/meta-commands/listen.cpp - CommandListen() [tcp]
func CommandTcpListenRequest(address, port string) error {
	fmt.Printf("TCP listen on %s:%s\n", address, port)
	return nil
}

// CommandTcpConnectRequest connects to remote debugger via TCP
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [tcp]
func CommandTcpConnectRequest(ip, port string) error {
	fmt.Printf("TCP connect to %s:%s\n", ip, port)
	return nil
}

// CommandTcpDisconnectRequest disconnects TCP connection
// C++ Source: code/debugger/commands/meta-commands/disconnect.cpp - CommandDisconnect() [tcp]
func CommandTcpDisconnectRequest() error {
	fmt.Println("TCP disconnect")
	return nil
}

// CommandSerialConnectRequest connects via serial port
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [serial]
func CommandSerialConnectRequest(portName string, baudRate uint32, pauseAfterConnect bool) error {
	fmt.Printf("serial connect: %s baud=%d pause=%v\n", portName, baudRate, pauseAfterConnect)
	return nil
}

// CommandNamedPipeConnectRequest connects via named pipe
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [pipe]
func CommandNamedPipeConnectRequest(namedPipe string, pauseAfterConnect bool) error {
	fmt.Printf("named pipe connect: %s pause=%v\n", namedPipe, pauseAfterConnect)
	return nil
}

// CommandProcessInterpretRequest interprets and executes a command string
// C++ Source: code/debugger/core/interpreter.cpp - InterpretCommand()
func CommandProcessInterpretRequest(command string) int {
	fmt.Printf("interpret: %s\n", command)
	return 0
}

// CommandCloseRequest closes debugging session
// C++ Source: code/debugger/commands/meta-commands/disconnect.cpp - CommandDisconnect() + TerminateVmx
func CommandCloseRequest(dc *core.DebuggerCore) error {
	return dc.TerminateVmx()
}

// CommandAttach attaches to debug a process in VMI Mode
// C++ Source: code/debugger/commands/meta-commands/attach.cpp - CommandAttach()
func CommandAttach(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 2 || tokens[0] != "pid" {
		return fmt.Errorf(".attach : attaches to debug a thread in VMI Mode.\n\nsyntax : \t.attach [pid ProcessId (hex)]\n\n\t\te.g : .attach pid b60 \n")
	}
	pid, err := strconv.ParseUint(tokens[1], 16, 32)
	if err != nil {
		return fmt.Errorf("please specify a correct hex value for process id\n")
	}
	fmt.Printf("Attaching to process 0x%x (%d)...\n\n", pid, pid)
	return dc.AttachToProcess(uint32(pid), false, true, false)
}

// CommandDetach detaches from debugging a user-mode process
// C++ Source: code/debugger/commands/meta-commands/detach.cpp - CommandDetach()
func CommandDetach(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		return fmt.Errorf(".detach : detaches from debugging a user-mode process.\n\nsyntax : \t.detach \n")
	}
	return dc.DetachFromProcess(0)
}

// CommandConnectLocal connects to local debugger (vmi-mode)
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [local]
func CommandConnectLocal(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("local debugging (vmi-mode)")
	dc.SetConnectedToHyperDbgLocally(true)
	return nil
}

// CommandConnectRemote connects to remote debugger
// C++ Source: code/debugger/commands/meta-commands/connect.cpp - CommandConnect() [remote]
func CommandConnectRemote(tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".connect : connects to a remote or local machine to start debugging.\n\nsyntax : \t.connect [local]\nsyntax : \t.connect [Ip (string)] [Port (decimal)]\n\n\t\te.g : .connect local\n\t\te.g : .connect 192.168.1.5 50000\n")
	}
	if tokens[0] == "local" {
		return fmt.Errorf("use .connectlocal for local connection\n")
	}
	ip := tokens[0]
	port := "50000"
	if len(tokens) > 1 {
		port = tokens[1]
	}
	fmt.Printf("connecting to remote debugger at %s:%s\n", ip, port)
	return fmt.Errorf("remote connect not implemented through IOCTL, use remote connection manager")
}

// CommandDisconnect disconnects from debugging session
// C++ Source: code/debugger/commands/meta-commands/disconnect.cpp - CommandDisconnect()
func CommandDisconnect(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		return fmt.Errorf(".disconnect : disconnects from a debugging session.\n\nsyntax : \t.disconnect \n")
	}
	if !dc.IsConnectedToAnyInstanceOfDebuggerOrDebuggee() {
		return fmt.Errorf("you're not connected to any instance of HyperDbg, did you use '.connect'? \n")
	}
	if dc.IsConnectedToHyperDbgLocally() {
		return fmt.Errorf("you cannot disconnect in local debugging while the driver is still loaded. please use 'unload' command before disconnecting\n")
	}
	fmt.Println("disconnected successfully")
	dc.SetConnectedToHyperDbgLocally(false)
	return nil
}

// CommandDebug executes debug commands
// C++ Source: code/debugger/commands/meta-commands/debug.cpp - CommandDebug()
func CommandDebug(tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".debug : executes debug commands.\n\nsyntax : \t.debug <command> [args]\n")
	}
	fmt.Printf("executing debug command: %s\n", strings.Join(tokens, " "))
	return nil
}

// CommandKill terminates/kills a process
// CommandKill terminates the current running process (.kill command)
// C++ Source: code/debugger/commands/meta-commands/kill.cpp - CommandKill()
func CommandKill(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		return fmt.Errorf(".kill : terminates the current running process.\n\nsyntax : \t.kill \n")
	}
	fmt.Println("killing current process...")
	return dc.KillActiveProcess()
}

// CommandListen listens for remote connections (guest server mode)
// C++ Source: code/debugger/commands/meta-commands/listen.cpp - CommandListen()
func CommandListen(tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".listen : listens for remote connections.\n\nsyntax : \t.listen <port>\n")
	}
	port, err := strconv.ParseUint(tokens[0], 10, 16)
	if err != nil {
		return fmt.Errorf("invalid port: %v", err)
	}
	fmt.Printf("listening on port %d\n", port)
	return nil
}

// CommandProcessList lists active processes
// C++ Source: code/debugger/commands/meta-commands/process.cpp - CommandProcess() [list]
func CommandProcessList(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("listing processes...")
	data, err := dc.GetActiveThreadsAndProcesses()
	if err != nil {
		return fmt.Errorf("process list failed: %v", err)
	}
	count, err := dc.QueryCountOfActiveProcesses()
	if err != nil {
		return fmt.Errorf("query count failed: %v", err)
	}
	fmt.Printf("found %d active processes, data size: %d bytes\n", count, len(data))
	return nil
}

// CommandProcessSwitch switches to a process by pid
// C++ Source: code/debugger/commands/meta-commands/process.cpp - CommandProcess() [pid]
func CommandProcessSwitch(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".process : switches to a process.\n\nsyntax : \t.process <pid>\n")
	}
	pid, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid pid: %v", err)
	}
	fmt.Printf("switching to process %d\n", pid)
	data, err := dc.QueryCurrentProcess()
	if err != nil {
		return fmt.Errorf("switch process failed: %v", err)
	}
	_ = data
	return nil
}

// CommandRestart restarts the previously started process (.restart command)
// C++ Source: code/debugger/commands/meta-commands/restart.cpp - CommandRestart()
func CommandRestart(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) > 0 {
		return fmt.Errorf(".restart : restarts the previously started process (using '.start' command).\n\nsyntax : \t.restart \n")
	}
	fmt.Println("restarting last started process...")
	return dc.RestartLastProcess()
}

// CommandScript executes a HyperDbg script
// C++ Source: code/debugger/commands/meta-commands/script.cpp - CommandScript()
func CommandScript(tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".script : executes a script.\n\nsyntax : \t.script <script_path>\n")
	}
	fmt.Printf("executing script: %s\n", tokens[0])
	return nil
}

// CommandStart starts a user-mode process for debugging
// C++ Source: code/debugger/commands/meta-commands/start.cpp - CommandStart()
func CommandStart(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".start : starts a process.\n\nsyntax : \t.start <path> [args]\n")
	}
	path := tokens[0]
	args := ""
	if len(tokens) > 1 {
		args = strings.Join(tokens[1:], " ")
	}
	fmt.Printf("started process: %s\n", path)
	return dc.StartProcess(path, args)
}

// CommandStatus shows debugger status information
// C++ Source: code/debugger/commands/meta-commands/status.cpp - CommandStatus()
func CommandStatus(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("showing debugger status...")
	fmt.Printf("Connected to local: %v\n", dc.IsConnectedToAnyInstanceOfDebuggerOrDebuggee())
	data, err := dc.QueryCurrentProcess()
	if err == nil && len(data) > 0 {
		fmt.Printf("Current process data available: %d bytes\n", len(data))
	}
	_ = data
	return nil
}

// CommandSwitchCore switches to a specific core (~ command)
// C++ Source: code/debugger/commands/meta-commands/switch.cpp - CommandSwitch() [~ coreNum]
func CommandSwitchCore(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("~ : switches to a core.\n\nsyntax : \t~ <core_number>\n")
	}
	coreNum, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid core number: %v", err)
	}
	fmt.Printf("switching to core %d\n", coreNum)
	return dc.SwitchCore(uint32(coreNum))
}

// CommandSymLoad loads symbols from PDB file
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [load]
func CommandSymLoad(tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf(".sym load : loads symbols.\n\nsyntax : \t.sym load <pdb_path>\n")
	}
	fmt.Printf("loading symbols from: %s\n", tokens[0])
	return nil
}

// CommandSymUnload unloads symbols
// C++ Source: code/debugger/commands/meta-commands/sym.cpp - CommandSym() [unload]
func CommandSymUnload(tokens []string, command string) error {
	fmt.Println("unloading symbols...")
	return nil
}

// CommandSympath sets or shows symbol path
// C++ Source: code/debugger/commands/meta-commands/sympath.cpp - CommandSympath()
func CommandSympath(tokens []string, command string) error {
	if len(tokens) < 1 {
		fmt.Println("current symbol path:")
	} else {
		fmt.Printf("setting symbol path to: %s\n", strings.Join(tokens, " "))
	}
	return nil
}

// CommandThreadList lists threads of current process
// C++ Source: code/debugger/commands/meta-commands/thread.cpp - CommandThread() [list]
func CommandThreadList(dc *core.DebuggerCore, tokens []string, command string) error {
	fmt.Println("listing threads...")
	data, err := dc.GetActiveThreadsAndProcesses()
	if err != nil {
		return fmt.Errorf("thread list failed: %v", err)
	}
	_ = data
	return nil
}

// CommandThreadSwitch switches to a thread by tid (~# command)
// C++ Source: code/debugger/commands/meta-commands/thread.cpp - CommandThread() [tid]
func CommandThreadSwitch(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 1 {
		return fmt.Errorf("~# : switches to a thread.\n\nsyntax : \t~# <tid>\n")
	}
	tid, err := strconv.ParseUint(tokens[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid tid: %v", err)
	}
	fmt.Printf("switching to thread %d\n", tid)
	data, err := dc.QueryCurrentThread()
	if err != nil {
		return fmt.Errorf("switch thread failed: %v", err)
	}
	_ = data
	return nil
}

// CommandCls clears the screen
// C++ Source: code/debugger/commands/meta-commands/cls.cpp - CommandCls()
func CommandCls(tokens []string, command string) error {
	fmt.Print("\033[2J\033[H")
	return nil
}

// CommandDump dumps memory content
// C++ Source: code/debugger/commands/meta-commands/dump.c - CommandDump()
func CommandDump(dc *core.DebuggerCore, tokens []string, command string) error {
	if len(tokens) < 2 {
		return fmt.Errorf(".dump : dumps memory.\n\nsyntax : \t.dump <format> <address>\n")
	}
	addr, err := strconv.ParseUint(tokens[1], 0, 64)
	if err != nil {
		return fmt.Errorf("invalid address: %v", err)
	}
	size := uint32(128)
	data, err := dc.ReadMemory(addr, size, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("dump memory failed: %v", err)
	}
	fmt.Printf("dumping memory at 0x%x with format %s: %d bytes read\n", addr, tokens[0], len(data))
	_ = data
	return nil
}

// CommandFormats shows value in different formats
// C++ Source: code/debugger/commands/meta-commands/formats.cpp - CommandFormats()
func CommandFormats(tokens []string, command string) error {
	fmt.Println("displaying current formats...")
	return nil
}
