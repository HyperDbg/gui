package sdk_test

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"os"
)

func fineToLines(name string) (lines []string, ok bool) {
	b, err := os.ReadFile(name)
	if !mycheck.Error(err) {
		return
	}
	return tool.File().ToLines(b)
}

type (
	eventType struct {
		names []string
	}
	eventObject struct {
	}
)
