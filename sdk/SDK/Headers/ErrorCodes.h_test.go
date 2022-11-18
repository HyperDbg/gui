package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/stream"
	"strings"
	"testing"
)

func TestGenErrorCodes(t *testing.T) {
	body := stream.New()
	body.WriteStringLn("package Headers")
	body.WriteStringLn("type ErrorCodes int")
	body.WriteStringLn("const (")
	lines, ok := stream.New().ReadToLines("ErrorCodes.h")
	if !ok {
		return
	}
	const define = "#define"
	for _, line := range lines {
		if strings.Contains(line, define) {
			fields := strings.Fields(line)
			//line = strings.TrimPrefix(line, define)
			body.WriteString(fields[1])
			if fields[2] == "0xFFFFFFFF" {
				body.WriteString(" ErrorCodes ")
			}
			body.WriteString("=")
			body.WriteStringLn(fields[2])
		}
	}
	body.WriteStringLn(")")
	mylog.Json("gen error code", body.String())
}
