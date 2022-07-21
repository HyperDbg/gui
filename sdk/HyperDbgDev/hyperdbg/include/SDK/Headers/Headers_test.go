package Headers

import (
	_ "embed"
	"github.com/ddkwork/librarygo/src/windef"
	"testing"
)

//go:embed Constants.h
var Constants string

func TestName(t *testing.T) {
	windef.Creat(Constants)
}
