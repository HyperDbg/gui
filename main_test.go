package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestUpdateAppModule(t *testing.T) {
	stream.UpdateAllLocalRep()
}

func TestUpdateSDKAndBin(t *testing.T) { // todo use abs path, don't copy bin dir to here
	if !stream.IsDir("bin") {
		return
	}
	mylog.CheckIgnore(os.RemoveAll("bin/release/SDK/Libraries"))
	mylog.CheckIgnore(os.RemoveAll("sdk/bindgen/SDK"))
	mylog.CopyDir("bin/release/SDK", "sdk/bindgen/SDK")
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
