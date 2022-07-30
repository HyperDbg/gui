package cpp2go_test

import (
	"github.com/ddkwork/hyperdbgui/sdk/cpp2go"
	"path/filepath"
	"strings"
	"testing"
)

func Test2(t *testing.T) {
	p := "../HyperDbgDev/hyperdbg"
	slash := filepath.ToSlash(p)
	_, after, found := strings.Cut(slash, `/`)
	if !found {
		panic("!found")
	}
	join := filepath.Join("back", after)
	println(join)

}

func TestName(t *testing.T) {
	o := cpp2go.NewObj()
	o.Src("../HyperDbgDev/hyperdbg").
		Dsc("back").
		Ext([]string{".c", ".cpp", ".inc", ".h", ".c", ".asm"}).
		ExpandPath("miscellaneous\\constants", ".txt").
		Dsc("back").
		Back().
		Convert()
	//o.Format()
}
