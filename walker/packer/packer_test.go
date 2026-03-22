package packer

import (
	"strings"
	"testing"
)

// D:\clone\VMProtect-devirtualization-main
// D:\workspace\hv\unlicense\unlicense-py3.11-x86
func TestDumpPe(t *testing.T) {
	CheckPacker("D:\\clone\\unlicense\\unlicense-py3.11-x86\\SuperRecovery.exe")
}

func TestDecodeSysCall(t *testing.T) {
	// println(strings.Has(".themida\n", "themida-winlicense"))
	after, found := strings.CutPrefix(".themida", ".")
	if !found {
		return
	}
	println(strings.Contains("themida-winlicense", after))
}
