package HyperDbgSdk

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestGengo(t *testing.T) {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".h":
			println(path)
			file, _ := os.Create(path + ".go")
			file.WriteString("package HyperDbgSdk\n")
		}
		return nil
	})
}
