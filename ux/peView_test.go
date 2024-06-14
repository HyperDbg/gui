package ux

import (
	"ms/xed"
	"testing"

	"github.com/ddkwork/golibrary/stream"
)

func TestParserSys(t *testing.T) {
	path := "../bin/debug/hprdbgkd.sys"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "hprdbgkd")
}

func TestParserExe(t *testing.T) {
	path := "../bin/debug/hyperdbg-cli.exe"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "hyperdbg-cli")
}

func TestParserDll(t *testing.T) {
	path := "../bin/debug/msdia140.dll"
	file := xed.ParserPe(path)
	stream.MarshalJsonToFile(file, "msdia140")
}
