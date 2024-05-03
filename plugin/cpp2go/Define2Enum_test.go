package cpp2go

import (
	"go/format"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestGenErrorCodes(t *testing.T) {
	//Define2Enum(EnumInfo{
	//	File:    "Constants.h",
	//	Package: "Headers",
	//	Type:    "ConstantsVar",
	//})
	//return
	Define2Enum(EnumInfo{
		File:    "ErrorCodes.h",
		Package: "Headers",
		Type:    "ErrorCodes",
	})
	return

	body := stream.NewBuffer("")
	body.WriteStringLn("package Headers")
	body.WriteStringLn("import \"fmt\"")
	body.WriteStringLn("type ErrorCodes int")
	body.WriteStringLn("const (")
	file := stream.NewBuffer("ErrorCodes.h")
	all := strings.ReplaceAll(file.String(), `\
   `, "")
	file.Reset()
	file.WriteString(all)
	lines := file.ToLines()
	// lines, ok := stream.NewBuffer("").ReadToLines("ErrorCodes.h")

	const define = "#define"
	once := sync.Once{}
	var codes []string
	for _, line := range lines {
		if strings.Contains(line, define) {
			fields := strings.Fields(line)
			codes = append(codes, fields[1])
			//line = strings.TrimPrefix(line, define)
			body.WriteString(fields[1])
			once.Do(func() {
				body.WriteString(" ErrorCodes ")
			})
			body.WriteString("=")
			body.WriteStringLn(fields[2])
		}
	}
	body.WriteStringLn(")")
	body.WriteStringLn("func (e ErrorCodes)String()string{")

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
	if !mylog.Check(err) {
		stream.WriteTruncate("ErrorCodes.h.go", body.Bytes())
		return
	}
	stream.WriteTruncate("ErrorCodes.h.go", source)
}

func TestRemoveNewline(t *testing.T) {
	code := `#define DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS \
    0xc000000d`
	all := strings.ReplaceAll(code, `\
   `, "")
	println(all)
}

type (
	EnumInfo struct {
		File    string
		Package string
		Type    string
	}
)

func Define2Enum(info EnumInfo) {
	body := stream.NewBuffer("")
	body.WriteStringLn("package " + info.Package)
	body.WriteStringLn("import \"fmt\"")
	body.WriteStringLn("type " + info.Type + " int")
	body.WriteStringLn("const (")
	file := stream.NewBuffer(info.File)
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
		body.WriteStringLn(strconv.Quote(stream.ToCamelUpper(code, false)))
	}
	body.WriteStringLn("default:")
	body.WriteStringLn("return fmt.Sprint(\"known error code \" + fmt.Sprintf(\"%d\",e))")
	body.WriteStringLn("}")
	body.WriteStringLn("}")
	mylog.Json("gen error code", body.String())
	source, err := format.Source(body.Bytes())
	if !mylog.Check(err) {
		stream.WriteTruncate(info.File+".go", body.Bytes())
		return
	}
	stream.WriteTruncate(info.File+".go", source)
}
