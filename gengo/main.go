package main

import (
	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func main() {
	mylog.Call(run)
}

func run() {
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
		Sources: []string{"amalgamated-dist\\Zydis.h"},
		// Sources: []string{"./Zydis.h"},
		AdditionalParams: []string{
			"-DZYAN_NO_LIBC",
			"-DZYAN_STATIC_ASSERT",
			//"-IC:\\Users\\Admin\\Desktop\\zydis\\dependencies\\zycore\\include",
			//"-IC:\\Users\\Admin\\Desktop\\zydis\\include",
		},
	}))

	mylog.Check(pkg.WriteToDir("zydis"))
}
