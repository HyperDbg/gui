package aa

import (
	"github.com/ddkwork/librarygo/src/stream/tool"
	"testing"
)

func TestName(t *testing.T) {
	tool.File().ToLines()
	
	
	
	
	s := `clang++ -v -c -fno-color-diagnostics -Xclang -ast-dump `
	s += `-I../../hprdbgctrl `
	s += `hprdbgctrl.cpp `
	s += `> output.txt`
	println(s)
	tool.File().WriteTruncate("../clang.sh", s)
}
