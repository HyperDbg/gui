package bindgen

import (
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/HyperDbg/sdk"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestUnmarshalCommandJson(t *testing.T) {
	t.Skip("Because the current json does not provide command parameters, this unit test should not be executed, and will override the manually entered parameters")
	var generated []sdk.Command
	mylog.Check(json.Unmarshal(stream.NewBuffer("commands.json").Bytes(), &generated))
	// mylog.Struct("todo",generated)

	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	m := map[string]string{
		"HiddenHookEPTBreakpoints": "EptHook",
		"HiddenHookEPTDetours":     "EptHook2",
		"MonitorIDTEntries":        "IdtEntriesMonitor",
		"RunsAUser_modeProcess":    "StartProcess",
		"MakePageAvailableInRAM":   "PageAvailableInRam",
		"DetectIOInstructionsIn":   "IoInDetect",
		"DetectIOInstructionsOut":  "IoOutDetect",
	}
	for _, s := range generated {
		for k, v := range m {
			if k == s.FullName {
				println(k)
				s.FullName = v
				delete(m, k)
				break
			}
		}
		s.Description = strings.TrimSuffix(s.Description, "\n")
		g.P("//", s.FullName)
		g.P("//", "Description:", s.Description)
		g.P("//", "Syntax:")
		for _, syntax := range s.Syntax {
			g.P("//", syntax)
		}
		if len(s.Examples) > 0 {
			g.P("//", "Examples:")
			for _, example := range s.Examples {
				g.P("//", example)
			}
		}
		if len(s.Notes) > 0 {
			g.P("//", "Notes:")
			for _, note := range s.Notes {
				g.P("//", note)
			}
		}
		g.P("func ", s.FullName, "()(status string) {")
		g.P("return ", "RunCommandEx(", strconv.Quote(s.Name), ") ")
		g.P("}")
		g.P()
	}
	stream.WriteGoFile("../commandWrapper.go", g.Bytes())
}

func TestCommandGenerate(t *testing.T) {
	t.Skip("not well")
	m := new(safemap.M[string, string])
	m.Set("debugging", "debugging")
	m.Set("extension", "extension")
	m.Set("hwdbg", "hwdbg")
	m.Set("meta", "meta")
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").EnumTypes("commands", m)

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
		name = stream.ToCamelUpper(name)
		mylog.Trace(name, path)

		command := sdk.Command{
			Name:        "",
			Description: "",
			Syntax:      nil,
			Examples:    nil,
			Notes:       nil,
			FullName:    "",
		}
		commands = append(commands, command)
		return err
	})

	m := new(safemap.M[string, string])
	for _, command := range commands {
		m.Set(command.FullName, command.FullName) // todo this must be use Cmd,then it will return the right command
	}
	stream.NewGeneratedFile().SetPackageName("sdk").SetFilePath("../").EnumTypes(kindName, m)
}
