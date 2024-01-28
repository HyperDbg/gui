package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"image"
	"testing"
)

func TestName(t *testing.T) {
	stream.Ico2PngAll("SND/ICO")
	stream.Png2SvgAll("SND/ICO")
	stream.Png2SvgAll("SND/png")
	return
	path := "D:\\workspace\\workspace\\gui\\BITMAP\\MODULES.bmp"
	//path="bmp/WINDOWS.bmp"
	decode, s, err := image.Decode(stream.NewReadFile(path))
	if !mylog.Error(err) {
		return
	}
	println(s)
	fmt.Println(decode.Bounds())
}
