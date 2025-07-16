package main

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
	"github.com/ollama/ollama/api"
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

func TestAi(t *testing.T) {
	os.RemoveAll("tmp")
	client := mylog.Check2(api.ClientFromEnvironment())
	w := sync.WaitGroup{}
	m := sync.Mutex{}
	filepath.Walk("D:\\workspace\\Up\\HyperDbg", func(path string, info fs.FileInfo, err error) error {
		switch filepath.Ext(path) {
		case ".c", ".h", ".cpp", ".hpp", ".cc", ".hh", ".cxx", ".hxx", ".c++", ".h++", ".C", ".H", ".CPP", ".HPP", ".CC", ".HH", ".CXX", ".HXX", ".C++", ".H++":
			w.Go(func() {
				m.Lock()
				defer m.Unlock()
				b := mylog.Check2(os.ReadFile(path))
				req := &api.GenerateRequest{
					Model:  "qwen2.5-coder:0.5b",
					Prompt: "完整翻译成go代码，不要任何注释，因为我直接把你的返回结果写入go文件了\n" + string(b),
					Stream: new(bool),
				}
				mylog.Check(client.Generate(context.Background(), req, func(response api.GenerateResponse) error {
					// Only print the response here; GenerateResponse has a number of other
					// interesting fields you want to examine.
					//fmt.Println(response.Response)
					_, after, found := strings.Cut(path, "workspace\\Up")
					if found {
						join := filepath.Join("tmp", after)
						join += ".go"
						mylog.Success("generate file", path)
						replace := strings.NewReplacer(
							"```go", ``,
							"```", ``,
						).Replace(response.Response)
						stream.WriteTruncate(join, replace)
					}
					return nil
				}))
			})
		}
		return nil
	})
	w.Wait()
}
