package zydis

import (
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestGetVersion(t *testing.T) {
	mylog.Todo("panic: Failed to find ZydisCategoryGetString procedure in zydis: The specified procedure could not be found.")
	GetVersion()
}

func TestZydis(t *testing.T) {
	pkg := gengo.NewPackage("zydis",
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
	mylog.Check(pkg.Transform("zydis", &clang.Options{
		Sources: []string{"amalgamated-dist/Zydis.h"},
		AdditionalParams: []string{
			"-DZYAN_NO_LIBC",
			"-DZYAN_STATIC_ASSERT",
		},
	}))
	mylog.Check(pkg.WriteToDir("."))
}
