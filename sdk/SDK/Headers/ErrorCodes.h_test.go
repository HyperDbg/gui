package Headers

import (
	"github.com/ddkwork/golibrary/src/stream"
	"strings"
	"testing"
)

func TestGenErrorCodes(t *testing.T) {
	lines, ok := stream.New().ReadToLines("ErrorCodes.h")
	if !ok {
		return
	}
	for _, line := range lines {
		if strings.Contains(line, "#define") {
			println(line)
		}
	}
}
