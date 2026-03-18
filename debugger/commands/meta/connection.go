package meta

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
)

type CommandConnect struct {
	dbg *debugger.HyperDbg
}

func NewCommandConnect(dbg *debugger.HyperDbg) *CommandConnect {
	return &CommandConnect{dbg: dbg}
}

func (c *CommandConnect) Help() {
	slog.Info("connect : connect to debuggee.")
	slog.Info("syntax : \tconnect [connection_string]")
	slog.Info("  connection_string - connection string (e.g., 'com:port=1,baud=115200' or 'tcp:port=50000')")
}

func (c *CommandConnect) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
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
	c.dbg.SetConnectedToRemoteDebuggee(true)
	slog.Info("Connecting to", "connection", connectionString)
	slog.Info("Connected successfully")
	return nil
}

type CommandDisconnect struct {
	dbg *debugger.HyperDbg
}

func NewCommandDisconnect(dbg *debugger.HyperDbg) *CommandDisconnect {
	return &CommandDisconnect{dbg: dbg}
}

func (c *CommandDisconnect) Help() {
	slog.Info("disconnect : disconnect from debuggee.")
	slog.Info("syntax : \tdisconnect")
}

func (c *CommandDisconnect) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Disconnect()
}

func (c *CommandDisconnect) Disconnect() error {
	if !c.dbg.IsConnectedToRemoteDebuggee() {
		return fmt.Errorf("not connected to debuggee")
	}

	c.dbg.SetConnectedToRemoteDebuggee(false)
	slog.Info("Disconnected from debuggee")
	return nil
}

type CommandAttach struct {
	dbg *debugger.HyperDbg
}

func NewCommandAttach(dbg *debugger.HyperDbg) *CommandAttach {
	return &CommandAttach{dbg: dbg}
}

func (c *CommandAttach) Help() {
	slog.Info("attach : attach to a process.")
	slog.Info("syntax : \tattach [pid]")
	slog.Info("  pid - process id to attach to")
}

func (c *CommandAttach) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 2 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	pid, err := strconv.ParseUint(tokens[1].Value, 0, 32)
	if err != nil {
		return fmt.Errorf("invalid process id: %w", err)
	}

	return c.Attach(uint32(pid))
}

func (c *CommandAttach) Attach(processId uint32) error {
	activeProcess := debugger.ActiveDebuggingProcess{
		IsActive:  true,
		IsPaused:  true,
		ProcessId: processId,
	}

	c.dbg.SetActiveProcess(activeProcess)
	slog.Info("Attached to process", "pid", processId)
	return nil
}

type CommandDetach struct {
	dbg *debugger.HyperDbg
}

func NewCommandDetach(dbg *debugger.HyperDbg) *CommandDetach {
	return &CommandDetach{dbg: dbg}
}

func (c *CommandDetach) Help() {
	slog.Info("detach : detach from process.")
	slog.Info("syntax : \tdetach")
}

func (c *CommandDetach) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.Detach()
}

func (c *CommandDetach) Detach() error {
	activeProcess := c.dbg.GetActiveProcess()
	if !activeProcess.IsActive {
		return fmt.Errorf("not attached to any process")
	}

	c.dbg.SetActiveProcess(debugger.ActiveDebuggingProcess{})
	slog.Info("Detached from process")
	return nil
}

type CommandStatus struct {
	dbg *debugger.HyperDbg
}

func NewCommandStatus(dbg *debugger.HyperDbg) *CommandStatus {
	return &CommandStatus{dbg: dbg}
}

func (c *CommandStatus) Help() {
	slog.Info("status : show debugger status.")
	slog.Info("syntax : \tstatus")
}

func (c *CommandStatus) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) != 1 {
		slog.Warn("incorrect use of command", "command", tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	return c.ShowStatus()
}

func (c *CommandStatus) ShowStatus() error {
	slog.Info("Debugger Status:")
	slog.Info("State", "value", c.dbg.GetState())
	slog.Info("Connected to remote debuggee", "value", c.dbg.IsConnectedToRemoteDebuggee())
	slog.Info("Serial connected to remote debuggee", "value", c.dbg.IsSerialConnectedToRemoteDebuggee())

	activeProcess := c.dbg.GetActiveProcess()
	if activeProcess.IsActive {
		slog.Info("Active process", "pid", activeProcess.ProcessId, "paused", activeProcess.IsPaused)
	} else {
		slog.Info("Active process: none")
	}

	return nil
}
