package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/stream"
)

func TestUpdateAppModule(t *testing.T) {
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\demo\\unison"))
	cmd := "go get -x " + "github.com/ddkwork/toolbox" + "@" + stream.GetLastCommitHashLocal("D:\\workspace\\workspace\\demo\\toolbox")
	stream.RunCommand(cmd)
	stream.RunCommand("go mod tidy")
	cleaner()

	mylog.Check(os.Chdir("D:\\workspace\\workspace\\demo\\app"))
	cmd = "go get -x " + "github.com/ddkwork/unison" + "@" + stream.GetLastCommitHashLocal("D:\\workspace\\workspace\\demo\\unison")
	stream.RunCommand(cmd)
	stream.RunCommand("go mod tidy")
	cleaner()

	mylog.Check(os.Chdir("D:\\workspace\\workspace\\gui"))
	cmd = "go get -x " + "github.com/ddkwork/app" + "@" + stream.GetLastCommitHashLocal("D:\\workspace\\workspace\\demo\\app")
	stream.RunCommand(cmd)
	stream.RunCommand("go mod tidy")
	cleaner()
}

func cleaner() {
	for s := range strings.Lines(`
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .
	//go install honnef.co/go/tools/cmd/staticcheck@latest
	//staticcheck ./...
	//go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
`) {
		s = strings.TrimSpace(s)
		if strings.HasPrefix(s, "::") || strings.HasPrefix(s, "//") || s == "" {
			continue
		}
		stream.RunCommand(s)
	}
}

func TestUpdateSDKAndBin(t *testing.T) { // todo use abs path, don't copy bin dir to here
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
