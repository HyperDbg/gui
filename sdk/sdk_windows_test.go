package sdk_test

import (
	_ "embed"
	"github.com/ddkwork/hyperdbgui/sdk"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestLoadVmm(t *testing.T) {
	s := sdk.New()
	s.LoadVmm()
}

func TestXmake(t *testing.T) {
	targets := stream.New()
	targets.WriteStringLn("add_rules(\"mode.debug\", \"mode.release\")\n")
	filepath.Walk("./HyperDbgDev", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".vcxproj" {
			if !strings.Contains(path, "dependencies") {
				s := stream.New()
				s.WriteStringLn("--" + filepath.Dir(path))
				s.WriteStringLn("target(" + strconv.Quote(filepath.Base(filepath.Dir(path))) + ")")
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
	tool.File().WriteTruncate("./HyperDbgDev/xmake.lua", targets.String())
}
