package debugging

import (
	"fmt"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
)

type CommandG struct {
	dbg debugger.UserDebugger
}

func NewCommandG(dbg debugger.UserDebugger) *CommandG {
	return &CommandG{dbg: dbg}
}

func (c *CommandG) Help() {
	mylog.Info("g : continues debuggee or continues processing kernel messages.")
	mylog.Info("syntax : \tg")
}

func (c *CommandG) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Request()
}

func (c *CommandG) Request() error {
	if c.dbg.IsConnected() {
		return c.RemoteConnectionSendCommand("g")
	} else if c.dbg.GetActiveDebuggingProcess().IsActive {
		activeProcess := c.dbg.GetActiveDebuggingProcess()
		if activeProcess.IsPaused {
			c.UdContinueProcess(activeProcess.ProcessId)
			activeProcess.IsPaused = false
			c.dbg.SetActiveDebuggingProcess(activeProcess)
		} else {
			mylog.Warning("the target process is already running")
		}
	}

	return nil
}

func (c *CommandG) RemoteConnectionSendCommand(cmd string) error {
	mylog.Info(cmd)
	return nil
}

func (c *CommandG) UdContinueProcess(processId uint32) error {
	mylog.Info(processId)
	return nil
}

type CommandP struct {
	dbg debugger.UserDebugger
}

func NewCommandP(dbg debugger.UserDebugger) *CommandP {
	return &CommandP{dbg: dbg}
}

func (c *CommandP) Help() {
	mylog.Info("p : step over (execute one instruction, stepping over function calls).")
	mylog.Info("syntax : \tp")
}

func (c *CommandP) Execute(tokens []debugger.CommandToken, command string) {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		mylog.Check("invalid arguments")
	}

	c.dbg.StepOver(0)
}

type CommandT struct {
	dbg debugger.UserDebugger
}

func NewCommandT(dbg debugger.UserDebugger) *CommandT {
	return &CommandT{dbg: dbg}
}

func (c *CommandT) Help() {
	mylog.Info("t : step in (execute one instruction, stepping into function calls).")
	mylog.Info("syntax : \tt")
}

func (c *CommandT) Execute(tokens []debugger.CommandToken, command string) {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		mylog.Check("invalid arguments")
	}

	c.dbg.StepInto()
}

type CommandHelp struct {
	dbg debugger.Packeter
}

func NewCommandHelp(dbg debugger.Packeter) *CommandHelp {
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
	mylog.Info("HyperDbg Commands:")
	mylog.Info("Debugging Commands:")
	mylog.Info("  g          - Continue execution")
	mylog.Info("  p          - Step over")
	mylog.Info("  t          - Step in")
	mylog.Info("  bp         - Set breakpoint")
	mylog.Info("  bc         - Clear breakpoint")
	mylog.Info("  bl         - List breakpoints")
	mylog.Info("  r          - Read/modify registers")
	mylog.Info("  k          - Show call stack")
	mylog.Info("Meta Commands:")
	mylog.Info("  help       - Show help")
	mylog.Info("  cls        - Clear screen")
	mylog.Info("  connect    - Connect to debuggee")
	mylog.Info("  disconnect - Disconnect from debuggee")
	mylog.Info("  attach     - Attach to process")
	mylog.Info("  detach     - Detach from process")
	mylog.Info("Extension Commands:")
	mylog.Info("  !epthook   - Set EPT hook")
	mylog.Info("  !monitor   - Monitor memory")
	mylog.Info("  !pte       - Query PTE")
	mylog.Info("  !va2pa     - Virtual to physical address")
	mylog.Info("  !pa2va     - Physical to virtual address")
}

func (c *CommandHelp) ShowCommandHelp(cmdName string) {
	mylog.Info(cmdName)
}

type CommandCls struct {
	dbg debugger.Packeter
}

func NewCommandCls(dbg debugger.Packeter) *CommandCls {
	return &CommandCls{dbg: dbg}
}

func (c *CommandCls) Help() {
	mylog.Info("cls : clear screen.")
	mylog.Info("syntax : \tcls")
}

func (c *CommandCls) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	mylog.Info("\033[2J")
	return nil
}
