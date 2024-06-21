package clang

import (
	"os"
	"path/filepath"

	"github.com/ddkwork/golibrary/stream"

	"github.com/ddkwork/golibrary/mylog"
)

type Options struct {
	ToolkitPath      string
	AdditionalParams []string
	Sources          []string
}

func (o *Options) ClangPath() string {
	if o.ToolkitPath != "" {
		if stat, e := os.Stat(o.ToolkitPath); e == nil && stat.IsDir() {
			return filepath.Join(o.ToolkitPath, "clang")
		} else {
			return o.ToolkitPath
		}
	}
	return "clang"
}

func (o *Options) ClangCommand(opt ...string) ([]byte, error) {
	// cmd := exec.Command(o.ClangPath(), opt...)
	// cmd.Args = append(cmd.Args, o.AdditionalParams...)
	// cmd.Args = append(cmd.Args, o.Sources...)
	c := make([]string, 0)
	c = append(c, o.ClangPath())
	c = append(c, opt...)
	c = append(c, o.AdditionalParams...)
	c = append(c, o.Sources...)
	return stream.RunCommandArgs(c...).Output.Bytes(), nil

	// cmd.Args = append(cmd.Args, "2>&1")
	// mylog.Trace("commands", strings.Join(cmd.Args, " "))
	// Stdout := &bytes.Buffer{}
	// Stderr := &bytes.Buffer{}
	// cmd.Stdout = Stdout
	// cmd.Stderr = Stderr
	// mylog.CheckIgnore(cmd.Run())
	// mylog.Check(Stderr.Bytes())
	// return Stdout.Bytes(), nil
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
	stream.RunCommand("clang -E -dM " + opt.Sources[0] + " > macros.log") // 2>&1
	// errg := &errgroup.Group{}

	// errg.Go(func() error {

	//return nil
	//})
	//mylog.Check(errg.Wait())
	//errg.Go(func() error {
	res := mylog.Check2(CreateAST(opt))
	stream.WriteTruncate("ast.json", res)
	ast = mylog.Check2(ParseAST(res)) //must ok

	res = mylog.Check2(CreateLayoutMap(opt))
	stream.WriteTruncate("astLayout.log", res)
	layout = mylog.Check2(ParseLayoutMap(res))

	//return nil
	//})

	// mylog.Check(errg.Wait())
	return ast, layout, nil
}
