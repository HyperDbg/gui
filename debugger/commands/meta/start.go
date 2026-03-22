package meta

import (
	"fmt"
	"strings"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/golibrary/std/mylog"
)

type CommandStart struct {
	dbg debugger.UserDebugger
}

func NewCommandStart(dbg debugger.UserDebugger) *CommandStart {
	return &CommandStart{dbg: dbg}
}

func (c *CommandStart) Help() {
	mylog.Info("start : runs a user-mode process.")
	mylog.Info("syntax : \tstart [path Path (string)] [Parameters (string)]")
	mylog.Info("  path - path to executable file")
	mylog.Info("  Parameters - command line parameters (optional)")
	mylog.Info("")
	mylog.Info("examples:")
	mylog.Info("  start path c:\\windows\\system32\\notepad.exe")
	mylog.Info("  start path \"c:\\program files\\app.exe\" \"arg1\" 2 \"arg 3\"")
}

func (c *CommandStart) Execute(tokens []debugger.CommandToken, command string) error {
	if len(tokens) < 3 {
		mylog.Warning(tokens[0].Value)
		c.Help()
		return fmt.Errorf("invalid arguments")
	}

	if strings.ToLower(tokens[1].Value) != "path" {
		mylog.Warning(tokens[1].Value)
		c.Help()
		return fmt.Errorf("expected 'path' keyword")
	}

	filePath := tokens[2].Value

	var parameters []string
	for _, token := range tokens[3:] {
		parameters = append(parameters, token.Value)
	}

	commandLine := strings.Join(parameters, " ")

	return c.Start(filePath, commandLine)
}

func (c *CommandStart) Start(filePath string, commandLine string) error {
	mylog.Info(filePath, commandLine)

	(c.dbg.StartProcess(filePath))

	mylog.Info("Process started successfully")
	return nil
}
