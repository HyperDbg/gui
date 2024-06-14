package ux

import (
	"testing"

	"github.com/ddkwork/golibrary/stream"
)

func TestParserSys(t *testing.T) {
	path := "bin/debug/hprdbgkd.sys"
	stream.NewBuffer(path)
}

func TestParserExe(t *testing.T) {
	path := "bin/debug/hyperdbg-cli.exe"
	stream.NewBuffer(path)
}

func TestParserDll(t *testing.T) {
	path := "bin/debug/msdia140.dll"
	stream.NewBuffer(path)
}
