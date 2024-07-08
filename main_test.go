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

func TestOpenWith(t *testing.T) {
	generateRegistryFile()
}

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

// todo update sdk/gen/SDK dir and remove libraries dir
func TestClear(t *testing.T) {
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
		case ".json", ".txt":
			if filepath.Base(path) == "sina2.json" {
				return nil
			}
			mylog.Info("clear file", path)
			mylog.Check(os.Remove(path))
		}
		return err
	})
}
