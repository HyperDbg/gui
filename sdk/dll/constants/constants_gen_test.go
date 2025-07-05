package constants

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/std/safemap"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
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
	m := new(safemap.M[string, string])
	for s := range stream.ReadFileToLines(fileName) {
		//if i == 4 {
		//	// break // test
		//}
		split := strings.Split(s, " ")
		v := split[1]
		if fileName == "ioctl.txt" {
			v = "0x" + v
		}
		m.Set(stream.ToCamelUpper(split[0]), v)
	}

	kind := stream.ToCamelUpper(stream.BaseName(fileName)) + "Kind"

	g := stream.NewGeneratedFile()
	g.P("package constants")
	g.P("import \"fmt\"")
	g.P("type ", kind, " int")
	g.P("const (")
	//for k,v := range m.Range() {
	//	if i == 0 {
	//		g.P(p.Key, " ", kind, " = ", p.Value)
	//		continue
	//	}
	//	g.P(p.Key, " = ", p.Value)
	//}
	g.P(")")
	g.P()

	g.P("func (k ", kind, ") String() string {")
	g.P("switch k {")
	//for _, p := range m.List() {
	//	g.P("case ", p.Value, ":")
	//	g.P("return ", strconv.Quote(stream.ToCamelUpper(p.Key)))
	//}
	g.P("default:")
	g.P("return \"unknown ", kind, " \"+fmt.Sprint(k)")
	g.P("}")
	g.P("}")
	g.P()

	g.P("func(k ", kind, ") Elements()[]", kind, " {")
	g.P("return []", kind, " {")
	//for _, p := range m.List() {
	//	g.P(p.Key, ",")
	//}
	g.P("}")
	g.P("}")
	g.P()

	stream.WriteGoFile(filepath.Join("tmp", stream.BaseName(fileName)+".go"), g.Buffer)
}
