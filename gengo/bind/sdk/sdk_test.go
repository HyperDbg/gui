package sdk

import (
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestSdk(t *testing.T) {
	mylog.Todo("需要实现处理多个dll导出函数的头文件问题，" +
		"是像zydis一样合并头文件还是修改gengo支持的方案好?不确定，都需要尝试一下," +
		"问题是输出文件是一个而不是多个")
	pkg := gengo.NewPackage("sdk",
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
	path := "../../../bin/debug/SDK/Imports/HyperDbgVmmImports.h"
	mylog.Check(pkg.Transform("sdk", &clang.Options{
		Sources: []string{path},
		// Sources: []string{"./Zydis.h"},
		AdditionalParams: []string{
			//"-D__int64 long long                             ",
			//"-D__iamcu__                                     ",
			//"-D__int32 int                                   ",
			//"-DNTDDI_WIN7 0x06010000                         ",
			//"-D__forceinline __attribute__((always_inline))  ",
			//"-D_AMD64_                                       ",
			//"-D_M_AMD64                                      ",
			//"-D__unaligned                                   ",
			//"-D_MSC_FULL_VER 192930133                       ",
			//"-DWIN32_LEAN_AND_MEAN                           ",
			//"-DUSE__NATIVE_PHNT_HEADERS                      ",
			//"-DZYAN_NO_LIBC",
			//"-DZYAN_NO_LIBC",
			//"-DZYAN_STATIC_ASSERT",
			//"-DZYDIS_STATIC_BUILD",
			//"-IC:\\Users\\Admin\\Desktop\\zydis\\include",
			//"-IC:\\Users\\Admin\\Desktop\\zydis\\dependencies\\zycore\\include",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\shared",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\ucrt",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\um",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km",
			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km\\crt",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\winrt",
			//"-IC:\\Program Files\\Microsoft Visual Studio\\2022\\Enterprise\\VC\\Tools\\MSVC\\14.40.33807\\include",

			//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl",
			//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbghv",
			//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl\\header",
			//"-ID:\\fork\\HyperDbg\\hyperdbg\\include",
			//"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies",
			//"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies\\phnt",
		},
	}))
	mylog.Check(pkg.WriteToDir("../../../bin/debug"))
}
