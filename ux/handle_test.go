package ux

import (
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestName(t *testing.T) {
	stream.Ico2PngAll("asserts")
	stream.Png2SvgAll("asserts")
	return
	path := "D:\\workspace\\workspace\\gui\\BITMAP\\MODULES.bmp"
	// path="bmp/WINDOWS.bmp"
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
		mylog.Error(os.Rename(path, strings.ToLower(path)))
		return err
	})
}
