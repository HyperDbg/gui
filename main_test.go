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
	mylog.FormatAllFiles()
	if !stream.IsDir("D:\\workspace\\workspace\\app") {
		return
	}
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\app"))
	session := stream.RunCommand("git log -1 --format=\"%H\"")
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\gui"))
	id := mylog.Check2(strconv.Unquote(session.Output.String()))
	mylog.Info("id", id)
	stream.RunCommand("go get github.com/ddkwork/app@" + id)
}

func TestUpdateSDKAndBin(t *testing.T) {
	if !stream.IsDir("bin") {
		return
	}
	mylog.Check(os.RemoveAll("bin/debug/SDK/Libraries"))
	mylog.Check(os.RemoveAll("sdk/bindgen/SDK"))
	mylog.CopyDir("sdk/bindgen/SDK", "bin/debug/SDK")
	filepath.Walk("bin", func(path string, info fs.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".exe", ".dll", ".sys":
			stream.WriteTruncate(filepath.Join("sdk/bin", filepath.Base(path)), stream.NewBuffer(path))
		}
		return err
	})
	mylog.Check(os.RemoveAll("bin"))
}

func TestClearTemp(t *testing.T) {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		switch ext {
		case ".json", ".txt", ".log":
			if filepath.Base(path) == "commands.json" {
				return nil
			}
			mylog.Info("clear file", path)
			mylog.Call(func() { mylog.Check(os.Remove(path)) })
		}
		return err
	})
}
