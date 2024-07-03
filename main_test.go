package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestUpdateAppModule(t *testing.T) {
	if !stream.IsDir("D:\\workspace\\workspace\\app") {
		return
	}
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\app"))
	session := stream.RunCommand("git log -1 --format=\"%H\"")
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\branch\\gui"))
	id := mylog.Check2(strconv.Unquote(session.Output.String()))
	mylog.Info("id", id)
	stream.RunCommand("go get github.com/ddkwork/app@" + id)
}

func TestClear(t *testing.T) {
	src := "bin/debug/SDK"
	dst := "sdk.gen/SDK"
	sysSrc := "bin/debug/hyperkd.sys"
	sysDst := "sdk.gen/SDK/Libraries/hyperkd.sys"

	mylog.Check(os.RemoveAll(dst))
	stream.CopyDir(dst, src)
	stream.WriteTruncate(sysDst, stream.NewBuffer(sysSrc))
	mylog.Check(os.RemoveAll("bin"))
}

func TestClearTemp(t *testing.T) {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		switch ext {
		case ".json", ".txt":
			mylog.Info("clear file", path)
			mylog.Check(os.Remove(path))
		}
		return err
	})
}
