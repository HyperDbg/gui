package main

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"image"
	"io/fs"
	"path/filepath"
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

func TestName2(t *testing.T) {
	filepath.Walk("asserts", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		println(path)
		return err
	})
}
