# gengo

Gengo is a simple and lightweight pure-Go wrapper generator for C libraries.
It uses Clang to parse the C header files and generates Go code that wraps around a shared library with high-performance calls without needing any Cgo.

## Usage

Example usage for Zydis:

```go
package main

import (
	"log"
	"os"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
)

func main() {
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
	err := pkg.Transform("zydis", &clang.Options{
		Sources: []string{"./Zydis.h"},
		AdditionalParams: []string{
			"-DZYAN_NO_LIBC",
			"-DZYAN_STATIC_ASSERT",
		},
	})
	if err != nil {
		log.Fatalf("Failed to transform: %v", err)
	}

	if err := pkg.WriteToDir("zydis"); err != nil {
		log.Fatalf("Failed to write the directory: %v", err)
	}
}
```
