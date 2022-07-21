package Headers

import (
	_ "embed"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/windef"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

//go:embed Constants.h
var Constants string

func TestName(t *testing.T) {
	paths := make([]string, 0)
	BasicTypes := ""
	includes := "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include"
	filepath.Walk(includes, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) == ".h" {
			if filepath.Base(path) != "BasicTypes.h" {
				paths = append(paths, path)
				println(path)
			} else {
				BasicTypes = path
			}
		}
		return err
	})

	paths_ := make([]string, 0)
	paths_ = append(paths_, BasicTypes)
	paths_ = append(paths_, paths...)
	buffer := stream.New()
	for _, path := range paths_ {
		b, err := os.ReadFile(path)
		if !mycheck.Error(err) {
			return
		}
		buffer.WriteBytesLn(b)
	}
	windef.Creat(buffer.String(), false)
}
