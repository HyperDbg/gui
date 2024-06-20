package sdk

import (
	"strings"
	"testing"
	"unicode"

	"github.com/ddkwork/golibrary/stream/maps"

	"github.com/ddkwork/golibrary/stream"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"

	"github.com/ddkwork/golibrary/mylog"
)

func mergeHeader() {
}

// ContainsLetter 检查字符串中是否包含字母
func ContainsLetter(s string) bool {
	if strings.HasPrefix(s, "0x") {
		return false
	}
	for _, char := range s {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

func TestBindMacros(t *testing.T) {
	mylog.Todo("handle macros func like CTL_CODE(DeviceType,Function,Method,Access) ( ((DeviceType) << 16) | ((Access) << 14) | ((Function) << 2) | (Method)) ")

	vars := stream.NewBuffer("")
	vars.WriteStringLn("package sdk")
	vars.WriteStringLn("var (")
	vars.WriteStringLn("MaxSerialPacketSize =10 * NORMAL_PAGE_SIZE") // todo need first define NORMAL_PAGE_SIZE
	vars.WriteStringLn("PAGE_SIZE = 4096")

	newVarsBody := stream.NewBuffer("")

	m := new(maps.SafeMap[string, bool])
	toLines := stream.NewBuffer("combined_headers.h").ToLines()
	for _, line := range toLines {
		m.Range(func(k string, v bool) bool { //todo bug
			if strings.HasPrefix(line, k) {
				newVarsBody.WriteStringLn(line)
				return true
			}
			return false
		})
	}

	skips := []string{ // todo 读取 combined_headers.h 保存 #define 开头的到一个map，这样原始问津的特征就完美匹配了
		"BUILD_",
		"FILE_DEVICE_UNKNOWN",
		"FILE_ANY_ACCESS",
		"FALSE",
		"TRUE",
		"_",
		"LO",
		"VOID",
		"time_t",
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
			line = strings.TrimSpace(line)
			line = strings.TrimSuffix(line, " ")
			println(line)
			if strings.Count(line, " ") == 1 {
				split := strings.Split(line, " ")
				split[1] = strings.TrimSuffix(split[1], "ull")
				split[1] = strings.TrimSuffix(split[1], "U")

				if ContainsLetter(split[1]) {
					mylog.Todo(split[1])
					continue
				}

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
