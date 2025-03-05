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
	paths := []string{
		"D:\\workspace\\workspace\\demo\\toolbox",
		"D:\\workspace\\workspace\\demo\\unison",
		"D:\\workspace\\workspace\\demo\\app",
	}
	for _, path := range paths {
		mylog.Check(os.Chdir(path))
		UpdateDependencies() // 执行一次提交一次，直到所有依赖都更新完毕，需要多次提交
	}
	mylog.Check(os.Chdir("D:\\workspace\\workspace\\gui"))
	UpdateDependencies()
}

func UpdateDependencies() {
	for s := range strings.Lines(`
     //go get -x gioui.org@main
	 //go get -x gioui.org/cmd@main
	 //go get -x gioui.org/example@main
	 //go get -x gioui.org/x@main
	 //go get -x github.com/oligo/gvcode@main
	 go get -x github.com/ddkwork/golibrary@master
	 //go get -x github.com/ddkwork/ux@master
	 go get -x github.com/google/go-cmp@master
	 go get -x github.com/ddkwork/app@master
	 go get -x github.com/ddkwork/toolbox@master
	 go get -x github.com/ddkwork/unison@master
	 //go get -x github.com/ebitengine/purego@main
	 //go get -x github.com/saferwall/pe@main
	 ::go get -u -x all
	 go mod tidy

	//go install mvdan.cc/gofumpt@latest
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
