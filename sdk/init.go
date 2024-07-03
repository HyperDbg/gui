package sdk

import (
	"crypto/sha256"
	"embed"
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
	"github.com/richardwilkes/toolbox/fatal"
	"github.com/richardwilkes/toolbox/xio/fs"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
)

//go:embed Libraries/*
var data embed.FS

func init() {
	m := stream.ReadEmbedFileMap(data, "Libraries")
	dir := mylog.Check2(os.UserCacheDir())
	dir = filepath.Join(dir, "hyperdbg", "cache")
	mylog.Check(os.MkdirAll(dir, 0755))

	dllData := m.Get("libhyperdbg.dll")
	GengoLibrary.LoadEmbed(dllData)

	mylog.Check(windows.SetDllDirectory(dir))

	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("libhyperdbg-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if !fs.FileExists(filePath) {
		fatal.IfErr(os.WriteFile(filePath, dllData, 0644))
	}
}

func init() {
	//windows.SetDllDirectory("../sdk.gen/SDK/Libraries")// abs not work,need temp dir
	//SetDllDirectory("D:\\workspace\\workspace\\branch\\gui\\sdk.gen\\SDK\\Libraries")
}
