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

func mergeHeader() { // todo need merge func
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

func MacrosInHeader() (m *maps.SafeMap[string, bool]) {
	m = new(maps.SafeMap[string, bool])
	for _, s := range stream.NewBuffer("macros.log").ToLines() {
		for _, s2 := range stream.NewBuffer("combined_headers.h").ToLines() {
			if strings.HasPrefix(s, s2) {
				m.Set(s, true)
			}
		}
	}
	return
}

func TestBindMacros(t *testing.T) {
	mylog.Todo("handle macros func like CTL_CODE(DeviceType,Function,Method,Access) ( ((DeviceType) << 16) | ((Access) << 14) | ((Function) << 2) | (Method)) ")
	mustPrefixs := MacrosInHeader()
	mylog.Trace("number of macros: %d", mustPrefixs.Len())

	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	g.P("const (")
	g.P("MaxSerialPacketSize =10 * NORMAL_PAGE_SIZE") // todo need first define NORMAL_PAGE_SIZE
	g.P("PAGE_SIZE = 4096")
	g.P(")")

	handledVars := new(maps.SafeMap[string, bool])
	todoVars := new(maps.SafeMap[string, bool])

	for _, line := range stream.NewBuffer("macros.log").ToLines() {
		for _, sure := range mustPrefixs.Keys() {
			if strings.HasPrefix(line, sure) {
				handledVars.Set(sure, true)
			}
		}
	}

	/*
		//	line = strings.TrimPrefix(line, "#define ")
			line = strings.TrimSpace(line)
			line = strings.TrimSuffix(line, " ")
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
	*/

	g.P(")")
	stream.WriteGoFile("tmp/vars.go", g.Buffer)
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
