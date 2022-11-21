package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/caseconv"
	"github.com/ddkwork/golibrary/src/stream"
	"github.com/ddkwork/golibrary/src/stream/tool"
	"go/format"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestGen_CTL_CODE(t *testing.T) {
	Define2CtlCode(CtlCodeInfo{
		File:    "Ioctls.h",
		Package: "Headers",
		Type:    "IoctlsKind",
	})
}

type (
	CtlCodeInfo struct {
		File    string
		Package string
		Type    string
	}
)

func Define2CtlCode(info CtlCodeInfo) {
	body := stream.New()
	body.WriteStringLn("package " + info.Package)
	body.WriteStringLn("import \"fmt\"")
	body.WriteStringLn("type " + info.Type + " int")
	body.WriteStringLn("const (")
	file := stream.NewReadFile(info.File)
	all := strings.ReplaceAll(file.String(), `\
   `, "")
	file.Reset()
	file.WriteString(all)
	lines, ok := file.ToLines(file.Bytes())
	if !ok {
		return
	}
	const define = "#define"
	once := sync.Once{}
	var codes []string
	for _, line := range lines {
		if strings.Contains(line, define) {
			fields := strings.Fields(line)
			codes = append(codes, fields[1])
			body.WriteString(fields[1])
			once.Do(func() {
				body.WriteString("\t" + info.Type + "\t")
			})
			body.WriteString("=")
			body.WriteStringLn(strings.Join(fields[2:], " "))
		}
	}
	body.WriteStringLn(")")
	body.WriteStringLn("func (e " + info.Type + ")String()string{")

	body.WriteStringLn("switch e {")
	for _, code := range codes {
		body.WriteString("case ")
		body.WriteString(code)
		body.WriteStringLn(":")
		body.WriteString("return ")
		body.WriteStringLn(strconv.Quote(caseconv.ToCamelUpper(code, false)))
	}
	body.WriteStringLn("default:")
	body.WriteStringLn("return fmt.Sprint(\"known error code \" + fmt.Sprintf(\"%d\",e))")
	body.WriteStringLn("}")
	body.WriteStringLn("}")
	mylog.Json("gen error code", body.String())
	source, err := format.Source(body.Bytes())
	if !mylog.Error(err) {
		tool.File().WriteTruncate(info.File+".go", body.Bytes())
		return
	}
	tool.File().WriteTruncate(info.File+".go", source)
}
