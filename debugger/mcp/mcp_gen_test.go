package mcp

import (
	"os"
	"os/exec"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
)

func TestGenerateMCP(t *testing.T) {
	cmd := exec.Command("go", "run", "github.com/ddkwork/HyperDbg/debugger/cmd/generate")
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	mylog.Check(cmd.Run())
}
