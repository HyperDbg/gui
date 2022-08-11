package hyperdbgui

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"testing"
)

func TestName(t *testing.T) {
	run, err := cmd.Run("go get -v -d github.com/ddkwork/librarygo@main")
	if !mycheck.Error(err) {
		return
	}
	println(run)
}
