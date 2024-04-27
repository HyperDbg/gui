package Headers

import (
	"go/format"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestGen_CTL_CODE(t *testing.T) {
	Define2CtlCode(CtlCodeInfo{
		File:    "Ioctls.h",
		Package: "Headers",
		Other: `
func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return ((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method)
}

const (
	FILE_DEVICE_UNKNOWN = windef.FILE_DEVICE_UNKNOWN
	METHOD_BUFFERED     = windef.METHOD_BUFFERED
	FILE_ANY_ACCESS     = windef.FILE_ANY_ACCESS
)
`,
		imports: []string{
			"fmt",
			"github.com/ddkwork/golibrary/src/cpp2go/delete/myc2go/windef",
		},
		Type:     map[string]string{"IoctlsKind": "uint32"},
		TypeInto: false,
	})
}

type (
	CtlCodeInfo struct {
		File     string
		Package  string
		Other    string
		imports  []string
		Type     map[string]string
		TypeInto bool
	}
)

func Define2CtlCode(info CtlCodeInfo) {
	body := stream.New("")
	body.WriteStringLn("package " + info.Package)

	body.WriteStringLn("import (")
	for _, s := range info.imports {
		body.WriteStringLn(strconv.Quote(s))
	}
	body.WriteStringLn(")")
	body.WriteStringLn(info.Other)
	body.WriteString("type ")
	typeKind := ""
	for k, v := range info.Type {
		typeKind = k
		body.WriteString(k)
		body.Indent(1)
		body.WriteStringLn(v)

	}
	body.WriteStringLn("var (")
	file := stream.NewReadFile(info.File)
	all := strings.ReplaceAll(file.String(), `\
   `, "")
	file.Reset()
	file.WriteString(all)
	lines := file.ToLines()

	const define = "#define"
	once := sync.Once{}
	var codes []string
	for _, line := range lines {
		if strings.Contains(line, define) {
			fields := strings.Fields(line)
			codes = append(codes, fields[1])
			body.WriteString(fields[1])
			if info.TypeInto {
				once.Do(func() {
					body.WriteString("\t" + typeKind + "\t")
				})
			}
			body.WriteString("=")
			body.WriteString(typeKind)
			body.WriteString("(")
			body.WriteString(strings.Join(fields[2:], " "))
			body.WriteStringLn(")")
		}
	}
	body.WriteStringLn(")")
	body.WriteStringLn("func (e " + typeKind + ")String()string{")

	body.WriteStringLn("switch e {")
	for _, code := range codes {
		body.WriteString("case ")
		body.WriteString(code)
		body.WriteStringLn(":")
		body.WriteString("return ")
		body.WriteStringLn(strconv.Quote(stream.ToCamelUpper(code, false)))
	}
	body.WriteStringLn("default:")
	body.WriteStringLn("return fmt.Sprint(\"known error code \" + fmt.Sprintf(\"%d\",e))")
	body.WriteStringLn("}")
	body.WriteStringLn("}")
	mylog.Json("gen error code", body.String())
	source, err := format.Source(body.Bytes())
	if !mylog.Error(err) {
		stream.WriteTruncate(info.File+".go", body.Bytes())
		return
	}
	stream.WriteTruncate(info.File+".go", source)
}
