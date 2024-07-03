package ux

import (
	"testing"

	"github.com/ddkwork/app/ms/xed"

	"github.com/ddkwork/golibrary/stream"
)

func TestParserSys(t *testing.T) {
	path := "hyperkd.sys"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "hyperkd")
}

func TestParserExe(t *testing.T) {
	path := "hyperdbg-cli.exe"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "hyperdbg-cli")
}

func TestParserDll(t *testing.T) {
	path := "msdia140.dll"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "msdia140")
}
