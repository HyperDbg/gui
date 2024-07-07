package sdk

import (
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

//go:embed bin/*
var data embed.FS

var SysPath = ""

func init() {
	m := stream.ReadEmbedFileMap(data, "bin")
	sha := sha256.Sum256(m.Get("libhyperdbg.dll"))
	dir := filepath.Join(mylog.Check2(os.UserCacheDir()), "hyperdbg", "cache", base64.RawURLEncoding.EncodeToString(sha[:]))
	mylog.Check(windows.SetDllDirectory(dir))
	SysPath = filepath.Join(dir, "hyperkd.sys")
	mylog.Trace("sysPath", SysPath)
	if !stream.IsDir(dir) {
		stream.CreatDirectory(dir)
		m.Range(func(k string, v []byte) bool {
			stream.WriteTruncate(filepath.Join(dir, k), v)
			return true
		})
	}
	mylog.Check2(GengoLibrary.LoadFrom(filepath.Join(dir, "libhyperdbg.dll")))
}
