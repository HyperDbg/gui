package cpp2go_test

import (
	"github.com/ddkwork/hyperdbgui/sdk/cpp2go"
	"github.com/ddkwork/librarygo/src/myc2go"
	"github.com/ddkwork/librarygo/src/mycheck"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"
)

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

func TestName(t *testing.T) {
	o := cpp2go.NewObj()
	o.Src("../HyperDbgDev/hyperdbg").
		Dsc("back").
		ExpandPath("miscellaneous\\constants", ".txt").
		Dsc("back").
		Back().
		Convert()
	//o.Format()
}

func TestC2go(t *testing.T) { //todo rewrite lib pkg
	setup := myc2go.NewSetup(myc2go.Setup{
		SetRoot: func() []string {
			return []string{
				//"./HyperDbgDev/hyperdbg/include/SDK/Headers",
				"./HyperDbgDev",
			}
		},
		SetExt: func() []string { return []string{".h", ".cpp", ".asm", ".txt"} },
		SetContains: func() []string {
			return []string{"hyperdbg\\miscellaneous",
				"hyperdbg\\include",
				"hyperdbg\\hprdbgctrl"}
		},
		SetContainsNot: func() []string { return []string{"build"} },
		SetSkip:        func() string { return "SCRIPT_ENGINE_KERNEL_MODE" },
		SetReplaces: func() map[string]string {
			return map[string]string{
				"sizeof(UINT32)":                 "4",
				"sizeof(DEBUGGER_REMOTE_PACKET)": "11",
			}
		},
		SetBasicTypes: func() string { return `Headers` },
	})
	mycheck.Assert(t).True(setup.ConvertAll())
}
