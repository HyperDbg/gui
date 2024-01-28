package sdk_test

import (
	_ "embed"
	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/cmd"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestXmake(t *testing.T) {
	targets := stream.New("")
	targets.WriteStringLn("add_rules(\"mode.debug\", \"mode.release\")\n")
	filepath.Walk("./HyperDbgDev", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".vcxproj" {
			if !strings.Contains(path, "dependencies") {
				project := filepath.Base(filepath.Dir(path))
				//println(project)
				//cl /analyze /d1Aprintast *.cpp > 1.ast
				ast := cmd.MakeArg("/analyze", "/d1Aprintast", ">", project+".ast")
				println(ast)

				s := stream.New("")
				s.WriteStringLn("--" + filepath.Dir(path))
				s.WriteStringLn("target(" + strconv.Quote(project) + ")")
				if strings.Contains(path, "hyperdbg-cli") {
					s.WriteStringLn("set_kind(\"binary\")")
				} else {
					s.WriteStringLn("set_kind(\"static\")")
				}
				s.WriteStringLn("add_headerfiles()")
				s.WriteStringLn("add_files()")
				s.WriteStringLn("add_includedirs(\"\", { public = true })")
				s.WriteStringLn("add_deps()")
				targets.WriteStringLn(s.String())
			}
		}
		return err
	})
	//stream.WriteTruncate("./HyperDbgDev/xmake.lua", targets.String())
}
