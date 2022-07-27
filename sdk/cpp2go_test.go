package sdk_test

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/ddkwork/librarygo/src/stream/tool/cmd"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/ast"
	"github.com/goplus/c2go/clang/parser"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"path/filepath"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	// #include "BasicTypes.h"
	//typedef unsigned short wchar_t;
	//typedef void *PVOID;
	//typedef void *PVOID64;

	main := "./HyperDbgDev/hyperdbg/hprdbgctrl/code/app/hprdbgctrl.cpp"
	base := "HyperDbgDev"
	includesPATH := make([]string, 0)
	for i, include := range includes {
		join := filepath.Join(base, include)
		join = ` -I` + join + " "
		includesPATH = append(includesPATH, join)
		println(join)
		if i == 2 {
			break
		}
	}
	include := ""
	for _, s := range includesPATH {
		include += s
	}
	//c := `clang++ -Xclang -dM -E -ast-dump=json -fsyntax-only `
	c := `clang++ -Xclang -dM -E -I` + include
	b, err2 := cmd.Run(c + main)
	if !mycheck.Error(err2) {
		return
	}
	node, warning, err := parser.ParseFile(main, 0)
	if !mycheck.Error(err) {
		return
	}
	mylog.Info("", warning)
	for _, n := range node.Inner {
		switch n.Kind {
		case ast.RecordDecl:
			mylog.Info("type struct{")
		case ast.TypedefType:
			mylog.Info("Type rename")
		case ast.FieldDecl:
			mylog.Info("FieldDecl Type, in field") //todo field end is FieldDecl
		case ast.EnumDecl:
			mylog.Info("const(")
		case ast.EnumConstantDecl:
			mylog.Info("EnumConstantDecl field")

		}
		dumpNode("Inner", *n)
		if n.Inner != nil {
			for _, n2 := range n.Inner {
				if n2 != nil {
					dumpNode("Inner", *n2)
				}
			}
		}
	}
	//mylog.Json("ast", b)
	lines, ok := tool.File().ToLines(b)
	if !ok {
		return
	}
	println("const(")
	for _, line := range lines {
		if strings.Contains(line, `#define`) {
			split := strings.Split(line, ` `)
			key := split[1]
			value := strings.Join(split[2:], ``)
			fmt.Printf("%-70s = %s\n", key, value)
		}
	}
	println(")")
}

func dumpNode(title string, n ast.Node) {
	if n.Name == "" {
		return
	}
	mylog.Warning(title, "")
	mylog.Info("Kind", n.Kind)
	s := n.Name
	s += "\t\t"
	if n.Type != nil {
		s += s + n.Type.QualType
	}
	if n.Loc != nil {
		s += " //" + fmt.Sprint(n.Loc.Line)
	}
	mylog.Success("code", s)
}

func Run(path, pkg string) {
	abs, err := filepath.Abs(path)
	if !mycheck.Error(err) {
		return
	}
	cl.SetDebug(cl.DbgFlagAll)
	preprocessor.SetDebug(preprocessor.DbgFlagAll)
	gox.SetDebug(gox.DbgFlagInstruction) // | gox.DbgFlagMatch)
	c2go.Run(pkg, abs, c2go.FlagDumpJson, nil)
}

var includes = []string{
	"hyperdbg/hprdbgctrl",
	"hyperdbg/hprdbgctrl/header",
	"hyperdbg/dependencies",
	"hyperdbg/dependencies/zydis/dependencies/zycore/include/Zycore",
	"hyperdbg/dependencies/zydis/include/Zydis",
	"hyperdbg/dependencies/zydis/dependencies/zycore/include",
	"hyperdbg/dependencies/zydis/include",
	"hyperdbg/dependencies/zydis/msvc",
	"hyperdbg/dependencies/phnt",
	"hyperdbg/dependencies/zydis/assets",
	"hyperdbg/dependencies/zydis/include/Zydis/Internal",
	"hyperdbg/dependencies/pdbex/Source",

	"hyperdbg/hprdbghv/header/debugger/commands",
	"hyperdbg/hprdbghv/header/debugger/core",
	"hyperdbg/hprdbghv/header/debugger/kernel-level",
	"hyperdbg/hprdbghv/header/debugger/script-engine",
	"hyperdbg/include/SDK/Headers",
	"hyperdbg/script-eval/header",
	"hyperdbg/hprdbghv/header/misc",
	"hyperdbg/hprdbghv/header/globals",
	"hyperdbg/hprdbghv/header/platform",
	"hyperdbg/hyperdbg-test/header",
	"hyperdbg/hyperdbg-test",
	"hyperdbg/include",
	"hyperdbg/kdserial",
	"hyperdbg/hprdbghv/header/debugger/user-level",
	"hyperdbg/hprdbghv/header/devices",
	"hyperdbg/hprdbghv/header/vmm/ept",
	"hyperdbg/include/SDK",
	"hyperdbg/script-engine/header",
	"hyperdbg/hprdbghv/header/common",
	"hyperdbg/hprdbghv/header/debugger/tests",
	"hyperdbg/hprdbghv",
	"hyperdbg/script-engine",
	"hyperdbg/hprdbghv/header/memory",
	"hyperdbg/symbol-parser/header",
	"hyperdbg/symbol-parser",
	"hyperdbg/hprdbghv/header/debugger/broadcast",
	"hyperdbg/hprdbghv/header/debugger/features",
	"hyperdbg/hprdbghv/header/debugger/objects",
	"hyperdbg/hprdbghv/header/debugger/transparency",
	"hyperdbg/hprdbghv/header/vmm/vmx",
	"hyperdbg/include/SDK/Imports",
	"hyperdbg/hprdbghv/header/components/registers",
	"hyperdbg/hprdbghv/header/debugger/communication",
}
