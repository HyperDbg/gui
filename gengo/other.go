package main

import (
	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
	"testing"
)

func TestDemoDll(t *testing.T) {
	pkg := gengo.NewPackage("demo",
		gengo.WithRemovePrefix(
			"Zydis_", "Zyan_", "Zycore_",
			"Zydis", "Zyan", "Zycore",
		),
		gengo.WithInferredMethods([]gengo.MethodInferenceRule{
			{Name: "ZydisDecoder", Receiver: "Decoder"},
			{Name: "ZydisEncoder", Receiver: "EncoderRequest"},
			{Name: "ZydisFormatterBuffer", Receiver: "FormatterBuffer"},
			{Name: "ZydisFormatter", Receiver: "ZydisFormatter *"},
			{Name: "ZyanVector", Receiver: "Vector"},
			{Name: "ZyanStringView", Receiver: "StringView"},
			{Name: "ZyanString", Receiver: "String"},
			{Name: "ZydisRegister", Receiver: "Register"},
			{Name: "ZydisMnemonic", Receiver: "Mnemonic"},
			{Name: "ZydisISASet", Receiver: "ISASet"},
			{Name: "ZydisISAExt", Receiver: "ISAExt"},
			{Name: "ZydisCategory", Receiver: "Category"},
		}),
		gengo.WithForcedSynthetic(
			"ZydisShortString_",
			"struct ZydisShortString_",
		),
	)
	path := "D:\\workspace\\workspace\\branch\\pkg\\cpp2go\\cgo\\tt\\library.hpp"
	mylog.Check(pkg.Transform("demo", &clang.Options{
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
			"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km",
			"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km\\crt",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\winrt",
			//"-IC:\\Program Files\\Microsoft Visual Studio\\2022\\Enterprise\\VC\\Tools\\MSVC\\14.40.33807\\include",

			//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbghv",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl\\header",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\include",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies\\phnt",
		},
	}))

	mylog.Check(pkg.WriteToDir("demo"))
}

func TestSdk(t *testing.T) {
	pkg := gengo.NewPackage("sdk",
		gengo.WithRemovePrefix(
			"Zydis_", "Zyan_", "Zycore_",
			"Zydis", "Zyan", "Zycore",
		),
		gengo.WithInferredMethods([]gengo.MethodInferenceRule{
			{Name: "ZydisDecoder", Receiver: "Decoder"},
			{Name: "ZydisEncoder", Receiver: "EncoderRequest"},
			{Name: "ZydisFormatterBuffer", Receiver: "FormatterBuffer"},
			{Name: "ZydisFormatter", Receiver: "ZydisFormatter *"},
			{Name: "ZyanVector", Receiver: "Vector"},
			{Name: "ZyanStringView", Receiver: "StringView"},
			{Name: "ZyanString", Receiver: "String"},
			{Name: "ZydisRegister", Receiver: "Register"},
			{Name: "ZydisMnemonic", Receiver: "Mnemonic"},
			{Name: "ZydisISASet", Receiver: "ISASet"},
			{Name: "ZydisISAExt", Receiver: "ISAExt"},
			{Name: "ZydisCategory", Receiver: "Category"},
		}),
		gengo.WithForcedSynthetic(
			"ZydisShortString_",
			"struct ZydisShortString_",
		),
	)
	path := "D:\\fork\\HyperDbg\\hyperdbg\\build\\bin\\debug\\SDK\\Headers\\BasicTypes.h"
	path = "D:\\fork\\HyperDbg\\hyperdbg\\hprdbghv\\code\\vmm\\ept\\Ept.c"
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
			"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km",
			"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\km\\crt",

			//"-IC:\\Program Files (x86)\\Windows Kits\\10\\Include\\10.0.26100.0\\winrt",
			//"-IC:\\Program Files\\Microsoft Visual Studio\\2022\\Enterprise\\VC\\Tools\\MSVC\\14.40.33807\\include",

			//"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbghv",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\hprdbgctrl\\header",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\include",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies",
			"-ID:\\fork\\HyperDbg\\hyperdbg\\dependencies\\phnt",
		},
	}))

	mylog.Check(pkg.WriteToDir("sdk"))
}
