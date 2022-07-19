package sdk_test

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"os"
	"path/filepath"
	"strings"
	"testing"
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

func TestName(t *testing.T) {
	files, err := filepath.Glob("*.h")
	if !mycheck.Error(err) {
		return
	}
	for _, file := range files {
		println(file)
	}
	return
	lines, ok := fineToLines(`D:\codespace\workspace\src\cppkit\driver\HyperDbgDev\hyperdbg\include\SDK\Headers\Datatypes.h`)
	if !ok {
		return
	}
	for _, line := range lines {
		println(line)
		if strings.Contains(line, `typedef enum `) {

		}
		if strings.Contains(line, `typedef struct `) {

		}
	}
}
