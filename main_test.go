package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

func TestClear(t *testing.T) {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		switch ext {
		case ".json", ".txt", ".pdb", ".exp", ".lib", ".cer", ".inf":
			if strings.Contains(path, "constants") {
				return err
			}
			mylog.Info("clear file", path)
			mylog.Check(os.Remove(path))
		}
		if info.IsDir() {
			println(path)
		}
		return err
	})
}
