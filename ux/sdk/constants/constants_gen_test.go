package constants

import (
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/orderedmap"
)

func TestGenConstants(t *testing.T) {
	mylog.Call(func() {
		genConstants("errors.txt")
		genConstants("ioctl.txt")
		genConstants("ntstatus.txt")
	})
}

func genConstants(fileName string) {
	m := orderedmap.New[string, string]()
	for _, s := range stream.NewBuffer(fileName).ToLines() {
		split := strings.Split(s, " ")
		m.Set(split[0], split[1])
	}

	kind := stream.BaseName(fileName) + "Kind"

	g := stream.NewGeneratedFile()
	g.P("package constants")
	g.P("type ", kind, " byte")
	g.P("const (")
	for i, p := range m.List() {
		if i == 0 {
			g.P(p.Key, kind, " = ", p.Value)
		}
		g.P(p.Key, " = ", p.Value)
	}
	g.P(")")
	g.P()

	g.P("func (k ", kind, ") String() string {")
	g.P("switch k {")
	for i, p := range m.List() {
		if i == 0 {
			continue
		}
		g.P("case ", p.Key, ":")
		g.P("return ", strconv.Quote(p.Value)) // todo rename
	}
	g.P("default:")
	g.P("return \"unknown\"")
	g.P("}")
	g.P()

	g.P("func (k ", kind, ") Elements() {")
	g.P("return []", kind, " {")
	for i, p := range m.List() {
		if i == 0 {
			continue
		}
		g.P(p.Key, ",")
	}
	g.P("}")
	g.P()

	stream.WriteGoFile(filepath.Join("tmp", stream.BaseName(fileName)+".go"), g.Buffer)
}
