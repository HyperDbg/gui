package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"image"
	"testing"
)

func TestName(t *testing.T) {
	decode, s, err := image.Decode(stream.NewReadFile("bmp/WINDOWS.bmp"))
	if !mylog.Error(err) {
		return
	}
	println(s)
	fmt.Println(decode.Bounds())
}
