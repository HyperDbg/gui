package clang

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
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
			return filepath.Join(o.ToolkitPath, "clang")
		} else {
			return o.ToolkitPath
		}
	}
	return "clang"
}

func (o *Options) ClangCommand(opt ...string) ([]byte, error) {
	header := o.Sources[0]
	switched := switchEnum(stream.NewBuffer(header).String())
	switched = switchStruct(switched)
	buffer := stream.NewBuffer(`
typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;
`)
	buffer.WriteStringLn(switched)
	buffer.ReplaceAll(`#include `, `//#include `)
	stream.WriteTruncate(header, buffer.String())

	cmd := exec.Command(o.ClangPath(), opt...)
	cmd.Args = append(cmd.Args, o.AdditionalParams...)
	cmd.Args = append(cmd.Args, o.Sources...)
	buf := &bytes.Buffer{}
	e := &bytes.Buffer{}
	cmd.Stdout = buf
	cmd.Stderr = e
	defer println(e.String())
	defer println(buf.String())
	mylog.Check(cmd.Run())

	return buf.Bytes(), nil

	c := make([]string, 0)
	c = append(c, o.ClangPath())
	c = append(c, opt...)
	c = append(c, o.AdditionalParams...)
	c = append(c, o.Sources...)
	return stream.RunCommandArgs(c...).Output.Bytes(), nil
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
		if e != nil {
			return e
		}
		ast, e = ParseAST(res)
		return e
	})
	errg.Go(func() error {
		res, e := CreateLayoutMap(opt)
		if e != nil {
			return e
		}
		// debug
		// println(string(res))
		layout, e = ParseLayoutMap(res)
		return e
	})
	if mylog.Check(errg.Wait()); err != nil {
		return nil, nil, err
	}
	return ast, layout, nil
}

func Parse_(opt *Options) (ast Node, layout *LayoutMap, err error) {
	// stream.RunCommand("clang -E -dM " + opt.Sources[0] + " > macros.log")
	res := mylog.Check2(CreateLayoutMap(opt))
	stream.WriteTruncate("astLayout.log", res)
	layout = mylog.Check2(ParseLayoutMap(res))

	res = mylog.Check2(CreateAST(opt))
	stream.WriteTruncate("ast.json", res)
	ast = mylog.Check2(ParseAST(res))

	return ast, layout, nil
}

func switchEnum(src string) string {
	start := 0
	lines := strings.Split(src, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "typedef enum") {
			start = i
		}
		if start > 0 && strings.HasPrefix(line, "}") {
			line = strings.TrimPrefix(line, "}")
			line = strings.TrimSuffix(line, ";")
			enumName := strings.TrimSpace(line)
			lines[start] = "typedef enum " + enumName + "_ {"
			if strings.TrimSpace(lines[start+1]) == "{" {
				lines[start+1] = ""
			}
			start = 0
		}
	}

	actual := strings.Join(lines, "\n")
	return actual
}

func switchStruct(src string) string {
	start := 0
	lines := strings.Split(src, "\n")
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//debug
		//if strings.Has(line, "typedef struct SYMBOL") {
		//	println()
		//}

		if strings.HasPrefix(line, "typedef struct") {
			after, found := strings.CutPrefix(strings.TrimSpace(line), "typedef struct ")
			if found {
				// if after=="{" ||strings.TrimSpace(lines[start+1]) == "{"{
				if after != "{" {
					continue
				}
			}
			//if strings.Has(line, ";") || strings.Has(line, "*") || strings.Has(line, "_") {
			//	continue
			//}
			start = i
		}
		if start > 0 && strings.HasPrefix(line, "}") {
			line = strings.TrimPrefix(line, "}")
			line = strings.TrimSuffix(line, ";")
			enumName := strings.TrimSpace(line)
			lines[start] = "typedef struct " + enumName + "_ {"
			if strings.TrimSpace(lines[start+1]) == "{" {
				lines[start+1] = ""
			}
			start = 0
		}
	}

	actual := strings.Join(lines, "\n")
	return actual
}
