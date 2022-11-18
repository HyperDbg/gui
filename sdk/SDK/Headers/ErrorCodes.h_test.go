package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/src/stream"
	"strings"
	"sync"
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
	mylog.Json("gen error code", body.String())
}
