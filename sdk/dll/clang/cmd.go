package clang

import "github.com/ddkwork/golibrary/std/stream"

func CheckHeadFile(path string) {
	removeComments := "clang -E -P " + path + " > output_file.txt"
	ast2Json := "clang -fsyntax-only -nobuiltininc -Xclang -ast-dump=json " + path
	ast2Layout := "clang -fsyntax-only -nobuiltininc -emit-llvm -Xclang -fdump-record-layouts -Xclang -fdump-record-layouts-complete " + path
	stream.RunCommand(removeComments)
	stream.RunCommand(ast2Json)
	stream.RunCommand(ast2Layout)
}
