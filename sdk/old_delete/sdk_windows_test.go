package sdk_test

import (
	_ "embed"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"
	"github.com/ddkwork/golibrary/stream/cmd"
)

func TestXmake(t *testing.T) {
	targets := stream.NewBuffer("")
	targets.WriteStringLn("add_rules(\"mode.debug\", \"mode.release\")\n")
	filepath.WalkDir("./HyperDbgDev", func(path string, info fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		if ext == ".vcxproj" {
			if !strings.Contains(path, "dependencies") {
				project := filepath.Base(filepath.Dir(path))
				// println(project)
				// cl /analyze /d1Aprintast *.cpp > 1.ast
				ast := cmd.MakeArg("/analyze", "/d1Aprintast", ">", project+".ast")
				println(ast)

				s := stream.NewBuffer("")
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
	// stream.WriteTruncate("./HyperDbgDev/xmake.lua", targets.String())
}
