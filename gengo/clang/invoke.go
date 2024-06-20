package clang

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/stream"

	"github.com/ddkwork/golibrary/mylog"
	"golang.org/x/sync/errgroup"
)

type Options struct {
	ToolkitPath      string
	AdditionalParams []string
	Sources          []string
}

func (o *Options) ClangPath() string {
	if o.ToolkitPath != "" {
		if stat, e := os.Stat(o.ToolkitPath); e == nil && stat.IsDir() {
			return filepath.Join(o.ToolkitPath, "clang++")
		} else {
			return o.ToolkitPath
		}
	}
	return "clang++"
}

func (o *Options) ClangCommand(opt ...string) ([]byte, error) {
	cmd := exec.Command(o.ClangPath(), opt...)
	cmd.Args = append(cmd.Args, o.AdditionalParams...)
	cmd.Args = append(cmd.Args, o.Sources...)
	mylog.Trace("commands", strings.Join(cmd.Args, " "))
	Stdout := &bytes.Buffer{}
	Stderr := &bytes.Buffer{}
	cmd.Stdout = Stdout
	cmd.Stderr = Stderr
	mylog.CheckIgnore(cmd.Run())
	mylog.Check(Stderr.Bytes())
	return Stdout.Bytes(), nil
}

func CreateAST(opt *Options) ([]byte, error) {
	return opt.ClangCommand(
		"-fsyntax-only",
		"-nobuiltininc",
		"-Xclang",
		"-ast-dump=json",
	)
}

func CreateLayoutMap(opt *Options) ([]byte, error) {
	return opt.ClangCommand(
		"-fsyntax-only",
		"-nobuiltininc",
		"-emit-llvm",
		"-Xclang",
		"-fdump-record-layouts",
		"-Xclang",
		"-fdump-record-layouts-complete",
	)
}

func Parse(opt *Options) (ast Node, layout *LayoutMap, err error) {
	errg := &errgroup.Group{}
	errg.Go(func() error {
		res, e := CreateAST(opt)
		stream.WriteTruncate("ast.json", res)
		if e != nil {
			return e
		}
		ast, e = ParseAST(res)
		return e
	})
	errg.Go(func() error {
		res, e := CreateLayoutMap(opt)
		stream.WriteTruncate("astLayout.log", res)
		if e != nil {
			return e
		}
		layout, e = ParseLayoutMap(res)
		return e
	})
	mylog.Check(errg.Wait())
	stream.RunCommand("clang++ -E -dM " + opt.Sources[0] + " > macros.log")
	return ast, layout, nil
}
