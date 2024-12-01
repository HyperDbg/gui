package constants

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestGenConstants(t *testing.T) {
	t.Skip("重复的value导致语法错误")
	mylog.Call(func() {
		genConstants("ioctl.txt")
		genConstants("errors.txt")
		genConstants("ntstatus.txt")
	})
}

func genConstants(fileName string) {
	m := new(maps.SafeSliceMap[string, string])
	for i, s := range stream.NewBuffer(fileName).ToLines() {
		if i == 4 {
			// break // test
		}
		split := strings.Split(s, " ")
		v := split[1]
		if fileName == "ioctl.txt" {
			v = "0x" + v
		}
		m.Set(stream.ToCamelUpper(split[0], false), v)
	}

	kind := stream.ToCamelUpper(stream.BaseName(fileName), false) + "Kind"

	g := stream.NewGeneratedFile()
	g.P("package constants")
	g.P("import \"fmt\"")
	g.P("type ", kind, " int")
	g.P("const (")
	for i, p := range m.List() {
		if i == 0 {
			g.P(p.Key, " ", kind, " = ", p.Value)
			continue
		}
		g.P(p.Key, " = ", p.Value)
	}
	g.P(")")
	g.P()

	g.P("func (k ", kind, ") String() string {")
	g.P("switch k {")
	for _, p := range m.List() {
		g.P("case ", p.Value, ":")
		g.P("return ", strconv.Quote(stream.ToCamelUpper(p.Key, false)))
	}
	g.P("default:")
	g.P("return \"unknown ", kind, " \"+fmt.Sprint(k)")
	g.P("}")
	g.P("}")
	g.P()

	g.P("func(k ", kind, ") Elements()[]", kind, " {")
	g.P("return []", kind, " {")
	for _, p := range m.List() {
		g.P(p.Key, ",")
	}
	g.P("}")
	g.P("}")
	g.P()

	stream.WriteGoFile(filepath.Join("tmp", stream.BaseName(fileName)+".go"), g.Buffer)
}
