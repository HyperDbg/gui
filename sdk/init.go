package sdk

import (
	"embed"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"golang.org/x/sys/windows"
)

//go:embed Libraries/*
var data embed.FS

var sysPath = ""

func init() {
	m := stream.ReadEmbedFileMap(data, "Libraries")
	dir := mylog.Check2(os.UserCacheDir())
	dir = filepath.Join(dir, "hyperdbg", "cache")

	mylog.CheckIgnore(os.RemoveAll(dir)) // todo test
	mylog.Check(os.MkdirAll(dir, 0755))

	dllData := m.Get("libhyperdbg.dll")
	mylog.Check2(GengoLibrary.LoadEmbed(dllData))

	SetCustomDriverPath(stringToBytePointer(dir), stringToBytePointer(stream.BaseName(dir)))

	m.Range(func(k string, v []byte) bool { // copy sys files to cache dir
		if k == "hyperkd.sys" {
			sysPath = filepath.Join(dir, k)
		}
		stream.WriteTruncate(filepath.Join(dir, k), v)
		return true
	})

	mylog.Check(windows.SetDllDirectory(dir)) // todo what another dep dll names ?

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
