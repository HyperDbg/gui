package sdk

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"
	"unicode"

	"github.com/ddkwork/golibrary/stream/orderedmap"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/maps"
	"github.com/stretchr/testify/assert"
)

const bugfix = `
typedef unsigned short wchar_t;
typedef int bool ;
#define PVOID void*
#define HANDLE void*
#define MAX_PATH 260
typedef unsigned __int64   SIZE_T;
typedef unsigned __int64   time_t;

typedef struct _LIST_ENTRY {
  struct _LIST_ENTRY *Flink;
  struct _LIST_ENTRY *Blink;
} LIST_ENTRY, *PLIST_ENTRY, PRLIST_ENTRY;
`

func TestMergeHeader(t *testing.T) {
	BasicTypes := ""
	Headers := orderedmap.New[string, bool]()
	Modules := orderedmap.New[string, bool]()
	Imports := orderedmap.New[string, bool]()

	g := stream.NewGeneratedFile()
	filepath.Walk("../../../bin", func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(path, "Examples") {
			return err
		}
		switch stream.BaseName(path) {
		case "Assertions", "HyperDbgSdk":
			return err
		}
		if filepath.Ext(path) == ".h" {
			println(path)
			switch {
			case stream.BaseName(path) == "BasicTypes":
				BasicTypes = path
			case strings.Contains(path, "Headers"):
				Headers.Set(path, true)
			case strings.Contains(path, "Imports"):
				Imports.Set(path, true)
			case strings.Contains(path, "Modules"):
				Modules.Set(path, true)
			}
		}
		return err
	})

	// todo merge step
	// bugfix
	// BasicTypes
	// Headers
	// Modules
	// Imports

	g.P("//bugfix.h")
	g.P(bugfix)
	g.P()
	mylog.Trace("merge", "bugfix.h")

	fnDo := func(path string) {
		g.P("//" + path)
		g.P(stream.NewBuffer(path)) // todo remove #pragma once ?
		g.P()
		mylog.Trace("merge", path)
	}

	fnDo(BasicTypes)

	for _, s := range Headers.Keys() {
		fnDo(s)
	}
	for _, s := range Modules.Keys() {
		fnDo(s)
	}
	for _, s := range Imports.Keys() {
		fnDo(s)
	}

	stream.WriteBinaryFile("merged_headers.h", g.Buffer)
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
		m.Set(s, true)
	}

	m2 := new(maps.SafeMap[string, bool])
	// for _, s := range stream.NewBuffer("combined_headers.h").ToLines() {
	for _, s := range stream.NewBuffer("merged_headers.h").ToLines() {
		if strings.HasPrefix(s, "#define ") {
			m2.Set(s, true)
		}
	}
	println(m.Len())
	println(m2.Len())

	for _, s := range m.Keys() {
		if !m2.HasPrefix(s) { // todo bug
			m.Delete(s)
			mylog.Trace("delete macro", s)
		}
	}
	println(m.Len())
	println(m2.Len())

	return
}

func TestBindMacros(t *testing.T) {
	t.Skip("not working now")
	mylog.Todo("handle macros func like CTL_CODE(DeviceType,Function,Method,Access) ( ((DeviceType) << 16) | ((Access) << 14) | ((Function) << 2) | (Method)) ")
	mustPrefixs := MacrosInHeader()
	return
	mylog.Trace("number of macros", mustPrefixs.Len())

	g := stream.NewGeneratedFile()
	g.P("package sdk")
	g.P()
	g.P("const (")
	g.P("MaxSerialPacketSize =10 * NORMAL_PAGE_SIZE") // todo need first define NORMAL_PAGE_SIZE
	g.P("PAGE_SIZE = 4096")

	allVars := new(maps.SafeMap[string, bool])

	for _, line := range stream.NewBuffer("macros.log").ToLines() {
		for _, sure := range mustPrefixs.Keys() {
			if strings.HasPrefix(line, sure) {
				allVars.Set(line, true)
				mylog.Trace("found macro", line)
			}
		}
	}

	return

	assert.Equal(t, mustPrefixs.Len(), allVars.Len())

	allVars.Range(func(k string, v bool) bool {
		line := strings.TrimPrefix(k, "#define ")
		line = strings.TrimSpace(line)
		if strings.Count(line, " ") == 1 {
			split := strings.Split(line, " ")
			split[1] = strings.TrimSuffix(split[1], "ull")
			split[1] = strings.TrimSuffix(split[1], "U")
			if ContainsLetter(split[1]) {
				return true
			}
			g.P(split[0] + "=" + split[1])
			allVars.Delete(k)
		}
		return true
	})
	g.P(")")
	stream.WriteGoFile("tmp/vars.go", g.Buffer)

	allVars.Range(func(k string, v bool) bool {
		mylog.Todo(k + " need handle")
		return true
	})
}

func TestBindSdk(t *testing.T) {
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
			// Sources:          []string{"combined_headers.h"},
			Sources:          []string{"merged_headers.h"},
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
