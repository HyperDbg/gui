package aa

import (
	"github.com/ddkwork/librarygo/src/stream/tool"
	"testing"
)

func TestName(t *testing.T) {
	tool.File().ToLines()
	// c2go hprdbgctrl.cpp
	//clang++ -c -fno-color-diagnostics -Xclang -ast-dump hprdbgctrl.cpp > output.txt -Irun/media/ddk/F8F6179A4491F617/codespace/workspace/src/gui/sdk/HyperDbgDev/hyperdbg/hprdbgctrl
	//clang++ -c -fno-color-diagnostics -Xclang -ast-dump hprdbgctrl.cpp > output.txt
	//
	s := `clang++ -v -c -fno-color-diagnostics -Xclang -ast-dump `
	s += `-I../../hprdbgctrl `
	s += `hprdbgctrl.cpp `
	s += `> output.txt`
	println(s)
	tool.File().WriteTruncate("../clang.sh", s)
}
