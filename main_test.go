package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestUpAppLiB(t *testing.T) {
	mylog.Check(os.Chdir("../../"))
	session := stream.RunCommand("git log -1 --format=\"%H\"")
	stream.RunCommand("go get github.com/ddkwork/app@" + mylog.Check2(strconv.Unquote(session.Output.String())))
}

func TestLoadDll(t *testing.T) {
	syscall.MustLoadDLL("libhyperdbg")
}

func TestClearAll(t *testing.T) {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		switch ext {
		case ".json", ".txt", ".pdb", ".exp", ".lib", ".cer", ".inf":
			if strings.Contains(path, "constants") {
				return err
			}
			mylog.Info("clear file", path)
			mylog.Check(os.Remove(path))
		}
		return err
	})
}

func TestClear(t *testing.T) {
	filepath.Walk("bin", func(path string, info fs.FileInfo, err error) error {
		if info != nil && info.IsDir() {
			if strings.HasPrefix(path, ".git") {
				return nil
			}
			switch path {
			case "bin\\debug\\SDK\\Examples",
				"bin\\debug\\SDK\\Libraries",
				"bin\\debug\\hyperhv",
				"bin\\debug\\hyperkd",
				"bin\\debug\\hyperlog":
				mylog.Info("clear dir", path)
				mylog.Check(os.RemoveAll(path))
				return err
			}
		}
		TestClearAll(t)
		return err
	})
	if stream.IsDir("bin/debug/SDK") {
		stream.CopyDir("sdk.gen\\SDK", "bin/debug/SDK")
		mylog.Check(os.RemoveAll("bin/debug/SDK"))
	}
	if stream.IsDir("bin/debug") {
		filepath.Walk("bin/debug", func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return err
			}
			stream.CopyFile(path, filepath.Base(path))
			mylog.Info("copy file", path)
			return err
		})
	}

	mylog.Check(os.RemoveAll("bin"))
}
