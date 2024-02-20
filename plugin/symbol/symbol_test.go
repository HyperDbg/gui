package symbol

import (
	"go/format"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestName(t *testing.T) {
	lines, ok := stream.ReadFileToLines(`ssdtTable.txt`)
	if !ok {
		return
	}

	ntdll := make([]string, 0)
	win32u := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, "win32u") {
			win32u = append(win32u, lines[i:]...)
			break
		}
		ntdll = append(ntdll, line)
	}

	type NtApiInfo struct {
		Name      string
		Number    string
		NumberHex string
	}

	fnSaveObj := func(bodyLine []string) (infos []NtApiInfo) {
		infos = make([]NtApiInfo, 0)
		for _, s := range bodyLine {
			if strings.Contains(s, "Nt") {
				// println(s)
				split := strings.Split(s, `|`)
				split[1] = strings.Trim(split[1], ` `)
				split[2] = strings.Trim(split[2], ` `)
				split[3] = strings.Trim(split[3], ` `)
				info := NtApiInfo{
					Name:      split[1],
					Number:    split[2],
					NumberHex: split[3],
				}
				infos = append(infos, info)
			}
			if strings.Contains(s, "Found") {
				// println(s)
				// println(len(infos))
			}
		}
		return
	}
	ntdlls := fnSaveObj(ntdll)
	win32us := fnSaveObj(win32u)
	fnGenGo := func(name, packageName string, obj []NtApiInfo) (ok bool) {
		buffer := stream.New("")
		buffer.WriteStringLn(`package ` + packageName)
		buffer.WriteStringLn(`type (`)
		buffer.WriteStringLn(`	Interface` + name + ` interface {`)
		for _, info := range obj {
			buffer.WriteStringLn(info.Name + `()(ok bool)`) // todo insert filter arg
		}
		buffer.WriteStringLn(`	}`)
		buffer.WriteStringLn(`	object` + name + ` struct{}`)
		buffer.WriteStringLn(`)`)
		buffer.WriteStringLn(`func New` + name + `() Interface` + name + ` { return &object` + name + `{} }`)
		goFileName := name + `.go`
		source, err2 := format.Source(buffer.Bytes())
		if !mylog.Error(err2) {
			return
		}
		return stream.WriteTruncate(goFileName, source)
	}
	fnGenGo(`Ntdll`, `symbol`, ntdlls)
	fnGenGo(`Win32u`, `symbol`, win32us)
}
