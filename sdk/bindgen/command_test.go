package bindgen

import (
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/HyperDbg/sdk"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/txt"
)

func TestCommandGenerate(t *testing.T) {
	t.Skip("not well")
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").Enum("commands", []string{
		"debugging",
		"extension",
		"hwdbg",
		"meta",
	}, nil)

	commandGenerate("DebuggingCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\debugging-commands")
	commandGenerate("ExtensionCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\extension-commands")
	commandGenerate("HwdbgCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\hwdbg-commands")
	commandGenerate("MetaCommands", "D:\\fork\\HyperDbg\\hyperdbg\\libhyperdbg\\code\\debugger\\commands\\meta-commands")
}

func commandGenerate(kindName, path string) {
	var commands []sdk.Command
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return err
		}
		name := stream.BaseName(path)
		name = strings.ReplaceAll(name, "-", "_") // not well
		name = strings.ReplaceAll(name, "~", "unknown")
		name = txt.ToCamelCase(name)
		mylog.Trace(name, path)

		command := sdk.Command{
			MethodName: name,
			Cmd:        stream.BaseName(path),
			Args:       []string{},
			Usage:      "",
			Demo:       []string{},
			DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(stream.BaseName(path)) + "))",
		}
		commands = append(commands, command)
		return err
	})

	keys := make([]string, 0)
	for _, command := range commands {
		keys = append(keys, command.MethodName) // todo this must be use Cmd,then it will return the right command
	}
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").Enum(kindName, keys, nil)
}
