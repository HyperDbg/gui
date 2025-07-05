package ui

import (
	"fmt"
	"image"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestName(t *testing.T) {
	path := "D:\\workspace\\workspace\\HyperDbg\\BITMAP\\MODULES.bmp"
	if !stream.IsFilePath(path) {
		return
	}
	// path="bmp/WINDOWS.bmp"
	decode, s := mylog.Check3(image.Decode(stream.NewBuffer(path)))
	println(s)
	fmt.Println(decode.Bounds())
}

func TestName2(t *testing.T) {
	filepath.WalkDir("asserts", func(path string, info fs.DirEntry, err error) error {
		if info.IsDir() {
			return nil
		}
		mylog.Check(os.Rename(path, strings.ToLower(path)))
		return err
	})
}
