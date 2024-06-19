package sdk

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestName(t *testing.T) {
	mylog.Todo("test bind bitset")
	//typedef struct _CR3_TYPE
	//{
	//    union
	//    {
	//        UINT64 Flags;
	//
	//        struct
	//        {
	//            UINT64 Pcid : 12;
	//            UINT64 PageFrameNumber : 36;
	//            UINT64 Reserved1 : 12;
	//            UINT64 Reserved_2 : 3;
	//            UINT64 PcidInvalidate : 1;
	//        } Fields;
	//    };
	//} CR3_TYPE, *PCR3_TYPE;
}

func TestBindAll(t *testing.T) {
	mylog.Warning("cpp stl not supported")
	root := "../../../bin/debug"
	root = "D:\\workspace\\workspace\\branch\\gui\\bin\\debug\\SDK\\Imports"
	filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".h" {
			if strings.Contains(path, "Examples") {
				return err
			}
			mylog.Trace("binding", path)
			mylog.Call(func() { bindOne(path) })
		}
		return err
	})
}

func bindOne(path string) {
	// todo "需要实现处理多个dll导出函数的头文件问题，"
	// "是像zydis一样合并头文件还是修改gengo支持的方案好?不确定，都需要尝试一下,"
	// "问题是输出文件是一个而不是多个"
	pkg := gengo.NewPackage("HPRDBGCTRL",
		gengo.WithRemovePrefix(
		//"Zydis_", "Zyan_", "Zycore_",
		//"Zydis", "Zyan", "Zycore",
		),
		gengo.WithInferredMethods([]gengo.MethodInferenceRule{
			//{Name: "ZydisDecoder", Receiver: "Decoder"},
		}),
		gengo.WithForcedSynthetic(
		//"ZydisShortString_",
		//"struct ZydisShortString_",
		),
	)
	mylog.Check(pkg.Transform("HPRDBGCTRL", &clang.Options{
		Sources: []string{path},
		AdditionalParams: []string{
			//"-DZYAN_NO_LIBC",
			//"-DZYAN_STATIC_ASSERT",
			//"-DZYDIS_STATIC_BUILD",
			"-DHYPERDBG_HPRDBGCTRL",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\ucrt",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\um",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km\\crt",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\winrt",
			//"-IC:\\Program Files\\Microsoft Visual Studio\\2022\\Enterprise\\VC\\Tools\\MSVC\\14.40.33807\\include",

			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbghv",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl\\header",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\include",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies\\phnt",
		},
	}))
	// mylog.Check(pkg.WriteToDir("../../../bin/debug"))
	pkg.Fprint(func(path_ string) (io.WriteCloser, error) {
		return os.Create(path + ".go")
	})
}
