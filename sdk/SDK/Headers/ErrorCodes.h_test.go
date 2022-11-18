package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/stream"
	"github.com/ddkwork/golibrary/src/stream/tool"
	"go/format"
	"strings"
	"sync"
	"testing"
)

func TestGenErrorCodes(t *testing.T) {
	body := stream.New()
	body.WriteStringLn("package Headers")
	body.WriteStringLn("type ErrorCodes int")
	body.WriteStringLn("const (")
	file := stream.NewReadFile("ErrorCodes.h")
	all := strings.ReplaceAll(file.String(), `\
   `, "")
	file.Reset()
	file.WriteString(all)
	lines, ok := file.ToLines(file.Bytes())
	//lines, ok := stream.New().ReadToLines("ErrorCodes.h")
	if !ok {
		return
	}
	const define = "#define"
	once := sync.Once{}

	for _, line := range lines {
		if strings.Contains(line, define) {
			fields := strings.Fields(line)
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

	body.WriteStringLn("}")
	mylog.Json("gen error code", body.String())
	source, err := format.Source(body.Bytes())
	if !mylog.Error(err) {
		tool.File().WriteTruncate("ErrorCodes.h.go", body.Bytes())
		return
	}
	tool.File().WriteTruncate("ErrorCodes.h.go", source)
}

func TestRemoveNewline(t *testing.T) {
	code := `#define DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS \
    0xc000000d`
	all := strings.ReplaceAll(code, `\
   `, "")
	println(all)
}
