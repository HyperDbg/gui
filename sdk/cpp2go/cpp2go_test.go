package cpp2go_test

import (
	"github.com/ddkwork/hyperdbgui/sdk/cpp2go"
	"github.com/ddkwork/librarygo/src/mycheck"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"
)

func TestCpp2Go(t *testing.T) {
	o := cpp2go.NewObj()
	o.Src("../HyperDbgDev/hyperdbg").
		Dst("back").
		ExpandPath("miscellaneous\\constants", ".txt").
		Dst("back").
		Back().
		Convert()
	//o.Format()
}

func TestUnsafe(t *testing.T) {
	sizeofUINT32 := unsafe.Sizeof(uint32(0))
	mycheck.Assert(t).Equal(uint32(sizeofUINT32), uint32(4))
}
func Test3(t *testing.T) {
	dir := filepath.Dir("back\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Constants.h.back")
	pkgName := filepath.Base(dir)
	println(pkgName)
}

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
