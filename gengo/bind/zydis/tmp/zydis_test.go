package zydis

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

func TestGetVersion(t *testing.T) {
	mylog.Todo("panic: Failed to find ZydisCategoryGetString procedure in zydis: The specified procedure could not be found.")
	GetVersion()
}
