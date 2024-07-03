package sdk

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

//go:embed bin/*
var data embed.FS

var sysPath = ""

func init() {
	m := stream.ReadEmbedFileMap(data, "bin")
	dir := mylog.Check2(os.UserCacheDir())
	dir = filepath.Join(dir, "hyperdbg", "cache")

	mylog.CheckIgnore(os.RemoveAll(dir))
	mylog.Check(os.MkdirAll(dir, 0755))

	m.Range(func(k string, v []byte) bool {
		join := filepath.Join(dir, k)
		if k == "hyperkd.sys" {
			sysPath = join
		}
		stream.WriteTruncate(join, v)
		return true
	})
	mylog.Check(windows.SetDllDirectory(dir))
	mylog.Check2(GengoLibrary.LoadFrom(filepath.Join(dir, "libhyperdbg.dll")))
	mylog.Trace("sysPath", sysPath)
}
