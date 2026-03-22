package meta

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
)

type CommandConnect struct {
	dbg debugger.Packeter
}

func NewCommandConnect(dbg debugger.Packeter) *CommandConnect {
	return &CommandConnect{dbg: dbg}
}

func (c *CommandConnect) Help() {
	mylog.Info("connect : connect to debuggee.")
	mylog.Info("syntax : \tconnect [connection_string]")
	mylog.Info("  connection_string - connection string (e.g., 'com:port=1,baud=115200' or 'tcp:port=50000')")
}

func (c *CommandConnect) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	var args []string
	for _, token := range tokens[1:] {
		args = append(args, token.Value)
	}
	connectionString := strings.Join(args, " ")
	return c.Connect(connectionString)
}

func (c *CommandConnect) Connect(connectionString string) error {
	mylog.Info(connectionString)

	parts := strings.Split(connectionString, ",")
	if len(parts) == 0 {
		return fmt.Errorf("invalid connection string")
	}

	connectionType := strings.TrimSpace(parts[0])
	if connectionType == "tcp" || connectionType == "com" {
		c.dbg.ConnectSerial(connectionString, 115200)
	}

	mylog.Info("Connected successfully")
	return nil
}

type CommandDisconnect struct {
	dbg debugger.Packeter
}

func NewCommandDisconnect(dbg debugger.Packeter) *CommandDisconnect {
	return &CommandDisconnect{dbg: dbg}
}

func (c *CommandDisconnect) Help() {
	mylog.Info("disconnect : disconnect from debuggee.")
	mylog.Info("syntax : \tdisconnect")
}

func (c *CommandDisconnect) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Disconnect()
}

func (c *CommandDisconnect) Disconnect() error {
	if !c.dbg.IsConnected() {
		return fmt.Errorf("not connected to debuggee")
	}

	mylog.Info("Disconnected from debuggee")
	return nil
}

type CommandAttach struct {
	dbg debugger.UserDebugger
}

func NewCommandAttach(dbg debugger.UserDebugger) *CommandAttach {
	return &CommandAttach{dbg: dbg}
}

func (c *CommandAttach) Help() {
	mylog.Info("attach : attach to a process.")
	mylog.Info("syntax : \tattach [pid]")
	mylog.Info("  pid - process id to attach to")
}

func (c *CommandAttach) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	pid := mylog.Check2(strconv.ParseUint(tokens[1].Value, 0, 32))

	return c.Attach(uint32(pid))
}

func (c *CommandAttach) Attach(processId uint32) error {
	activeProcess := debugger.UserActiveDebuggingProcess{
		IsActive:  true,
		IsPaused:  true,
		ProcessId: processId,
	}

	c.dbg.SetActiveDebuggingProcess(activeProcess)
	mylog.Info(processId)
	return nil
}

type CommandDetach struct {
	dbg debugger.UserDebugger
}

func NewCommandDetach(dbg debugger.UserDebugger) *CommandDetach {
	return &CommandDetach{dbg: dbg}
}

func (c *CommandDetach) Help() {
	mylog.Info("detach : detach from process.")
	mylog.Info("syntax : \tdetach")
}

func (c *CommandDetach) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Detach()
}

func (c *CommandDetach) Detach() error {
	activeProcess := c.dbg.GetActiveDebuggingProcess()
	if !activeProcess.IsActive {
		return fmt.Errorf("not attached to any process")
	}

	c.dbg.SetActiveDebuggingProcess(debugger.UserActiveDebuggingProcess{})
	mylog.Info("Detached from process")
	return nil
}

type CommandStatus struct {
	dbg debugger.Packeter
}

func NewCommandStatus(dbg debugger.Packeter) *CommandStatus {
	return &CommandStatus{dbg: dbg}
}

func (c *CommandStatus) Help() {
	mylog.Info("status : show debugger status.")
	mylog.Info("syntax : \tstatus")
}

func (c *CommandStatus) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.ShowStatus()
}

func (c *CommandStatus) ShowStatus() error {
	mylog.Info("Debugger Status:")
	mylog.Info(c.dbg.IsConnected())

	if userDbg, ok := c.dbg.(debugger.UserDebugger); ok {
		activeProcess := userDbg.GetActiveDebuggingProcess()
		if activeProcess.IsActive {
			mylog.Info(activeProcess.ProcessId, activeProcess.IsPaused)
		} else {
			mylog.Info("Active process: none")
		}
	}

	return nil
}
