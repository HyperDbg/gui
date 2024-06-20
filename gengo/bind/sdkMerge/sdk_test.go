package sdk

import (
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"

	"github.com/ddkwork/golibrary/mylog"
)

func mergeHeader() {
}

func TestBindMacros(t *testing.T) {
	mylog.Todo("handle macros func like CTL_CODE(DeviceType,Function,Method,Access) ( ((DeviceType) << 16) | ((Access) << 14) | ((Function) << 2) | (Method)) ")

	vars := stream.NewBuffer("")
	vars.WriteStringLn("package sdk")
	vars.WriteStringLn("var (")

	skips := []string{
		"BUILD_",
		"FILE_DEVICE_UNKNOWN",
		"FILE_ANY_ACCESS",
		"FALSE",
		"TRUE",
		"_",
		//"",
		//"",
	}
	lines := stream.NewBuffer("macros.log").ToLines()
	for _, line := range lines {
		line = strings.TrimPrefix(line, "#define ")
		//line = strings.TrimSpace(line)
		stop := false
		for _, skip := range skips {
			if strings.HasPrefix(line, skip) {
				stop = true
				continue
			}
		}
		if !stop {
			println(line)
			if strings.Count(line, " ") == 1 {
				split := strings.Split(line, " ")
				vars.WriteStringLn(split[0] + "=" + split[1])
			}
		}
	}
	vars.WriteStringLn(")")
	stream.WriteGoFile("tmp/vars.go", vars)
}

func TestBind(t *testing.T) {
	mylog.SetDebug(false)
	mylog.Call(func() {
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
			Sources:          []string{"combined_headers.h"},
			AdditionalParams: []string{
				//"-DZYAN_NO_LIBC",
				//"-DZYAN_STATIC_ASSERT",
				//"-DZYDIS_STATIC_BUILD",
				//"-DHYPERDBG_HPRDBGCTRL",

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
		mylog.Check(pkg.WriteToDir("C:\\Users\\Admin\\Desktop\\New folder"))
	})
}
