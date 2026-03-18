package debugging

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
)

type CommandG struct {
	dbg *debugger.HyperDbg
}

func NewCommandG(dbg *debugger.HyperDbg) *CommandG {
	return &CommandG{dbg: dbg}
}

func (c *CommandG) Help() {
	slog.Info("g : continues debuggee or continues processing kernel messages.")
	slog.Info("syntax : \tg")
}

func (c *CommandG) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Request()
}

func (c *CommandG) Request() error {
	c.dbg.SetBreakPrintingOutput(false)

	if c.dbg.IsConnectedToRemoteDebuggee() {
		return c.RemoteConnectionSendCommand("g")
	} else if c.dbg.GetActiveProcess().IsActive {
		activeProcess := c.dbg.GetActiveProcess()
		if activeProcess.IsPaused {
			c.UdContinueProcess(activeProcess.ProcessId)
			activeProcess.IsPaused = false
			c.dbg.SetActiveProcess(activeProcess)
		} else {
			slog.Warn("the target process is already running")
		}
	}

	return nil
}

func (c *CommandG) RemoteConnectionSendCommand(cmd string) error {
	slog.Info("Sending command to remote debuggee", "command", cmd)
	return nil
}

func (c *CommandG) UdContinueProcess(processId uint32) error {
	slog.Info("Continuing process", "pid", processId)
	return nil
}

type CommandP struct {
	dbg *debugger.HyperDbg
}

func NewCommandP(dbg *debugger.HyperDbg) *CommandP {
	return &CommandP{dbg: dbg}
}

func (c *CommandP) Help() {
	slog.Info("p : step over (execute one instruction, stepping over function calls).")
	slog.Info("syntax : \tp")
}

func (c *CommandP) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.dbg.PerformStepOver()
}

type CommandT struct {
	dbg *debugger.HyperDbg
}

func NewCommandT(dbg *debugger.HyperDbg) *CommandT {
	return &CommandT{dbg: dbg}
}

func (c *CommandT) Help() {
	slog.Info("t : step in (execute one instruction, stepping into function calls).")
	slog.Info("syntax : \tt")
}

func (c *CommandT) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.dbg.PerformStepIn()
}

type CommandHelp struct {
	dbg *debugger.HyperDbg
}

func NewCommandHelp(dbg *debugger.HyperDbg) *CommandHelp {
	return &CommandHelp{dbg: dbg}
}

func (c *CommandHelp) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) == 1 {
		c.ShowGeneralHelp()
		return nil
	}

	cmdName := strings.ToLower(tokens[1].Value)
	c.ShowCommandHelp(cmdName)
	return nil
}

func (c *CommandHelp) ShowGeneralHelp() {
	slog.Info("HyperDbg Commands:")
	slog.Info("Debugging Commands:")
	slog.Info("  g          - Continue execution")
	slog.Info("  p          - Step over")
	slog.Info("  t          - Step in")
	slog.Info("  bp         - Set breakpoint")
	slog.Info("  bc         - Clear breakpoint")
	slog.Info("  bl         - List breakpoints")
	slog.Info("  r          - Read/modify registers")
	slog.Info("  k          - Show call stack")
	slog.Info("Meta Commands:")
	slog.Info("  help       - Show help")
	slog.Info("  cls        - Clear screen")
	slog.Info("  connect    - Connect to debuggee")
	slog.Info("  disconnect - Disconnect from debuggee")
	slog.Info("  attach     - Attach to process")
	slog.Info("  detach     - Detach from process")
	slog.Info("Extension Commands:")
	slog.Info("  !epthook   - Set EPT hook")
	slog.Info("  !monitor   - Monitor memory")
	slog.Info("  !pte       - Query PTE")
	slog.Info("  !va2pa     - Virtual to physical address")
	slog.Info("  !pa2va     - Physical to virtual address")
}

func (c *CommandHelp) ShowCommandHelp(cmdName string) {
	slog.Info("Help for command", "command", cmdName)
}

type CommandCls struct {
	dbg *debugger.HyperDbg
}

func NewCommandCls(dbg *debugger.HyperDbg) *CommandCls {
	return &CommandCls{dbg: dbg}
}

func (c *CommandCls) Help() {
	slog.Info("cls : clear screen.")
	slog.Info("syntax : \tcls")
}

func (c *CommandCls) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	slog.Info("\033[2J")
	return nil
}
