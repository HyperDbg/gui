package xed_test

import (
	_ "embed"
	"testing"
)

//go:embed asm.exe
var data []byte

func TestDecode32(t *testing.T) {
	x := xed.New(data).Decode32()
	println(x.IntelSyntaxAsm.String())
}

func TestDecode64(t *testing.T) {
	// Decode64(tt.args.data)
}
