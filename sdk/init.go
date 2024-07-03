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

	// mylog.CheckIgnore(os.RemoveAll(dir)) // todo test
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

	//sha := sha256.Sum256(dllData)
	//dllName := fmt.Sprintf("libhyperdbg-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	//filePath := filepath.Join(dir, dllName)
	//if !fs.FileExists(filePath) {
	//	fatal.IfErr(os.WriteFile(filePath, dllData, 0644))
	//}
}

//func init() {
//windows.SetDllDirectory("../sdk.gen/SDK/Libraries")// abs not work,need temp dir
//windows.SetDllDirectory("D:\\workspace\\workspace\\branch\\gui\\sdk.gen\\SDK\\Libraries")
//}
