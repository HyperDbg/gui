package sdk

import (
	_ "embed"
	"github.com/ddkwork/librarygo/src/caseconv"
	"github.com/ddkwork/librarygo/src/myc2go"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestC2go(t *testing.T) {
	setup := myc2go.NewSetup(myc2go.Setup{
		SetRoot: func() []string { return []string{"./HyperDbgDev"} },
		SetExt:  func() []string { return []string{".h", ".cpp", ".asm", ".txt"} },
		SetContains: func() []string {
			return []string{"hyperdbg\\miscellaneous",
				"hyperdbg\\include",
				"hyperdbg\\hprdbgctrl"}
		},
		SetContainsNot: func() []string { return []string{"build"} },
		SetSkip:        func() string { return "SCRIPT_ENGINE_KERNEL_MODE" },
		SetReplaces: func() map[string]string {
			return map[string]string{
				"sizeof(UINT32)":                 "4",
				"sizeof(DEBUGGER_REMOTE_PACKET)": "11",
			}
		},
	})
	mycheck.Assert(t).True(setup.ConvertAll())
}

func TestAll(t *testing.T) {
	cppPath := "./HyperDbgDev"
	buffer := stream.New()
	filepath.Walk(cppPath, func(path string, info fs.FileInfo, err error) error {
		switch {
		case
			strings.Contains(path, "hyperdbg\\miscellaneous"),
			strings.Contains(path, "hyperdbg\\include"),
			strings.Contains(path, "hyperdbg\\hprdbgctrl") &&
				!strings.Contains(path, "build"):
			ext := filepath.Ext(path)
			base := filepath.Base(path)
			base = base[:len(base)-len(ext)]
			if strings.Contains(base, `~`) {
				base = strings.ReplaceAll(base, `~`, `unknown`)
			}
			if strings.Contains(base, `-`) {
				base = strings.ReplaceAll(base, `-`, `_`)
			}
			if strings.Contains(base, `switch`) {
				base = strings.ReplaceAll(base, `switch`, `switchA`)
			}
			base = caseconv.ToCamel(base, false)
			base = strings.TrimRight(base, " ")
			switch ext {
			case ".h", ".cpp", ".asm", ".txt":
				buffer.Reset()
				dir := filepath.Dir(path)
				lastIndex := strings.LastIndex(dir, `\`) //todo  filepath.Base(filepath.Dir(path)))
				pkgNmae := dir[lastIndex+1:]
				if strings.Contains(pkgNmae, "-") {
					pkgNmae = strings.ReplaceAll(pkgNmae, "-", "")
				}
				if strings.Contains(pkgNmae, "~") {
					pkgNmae = strings.ReplaceAll(pkgNmae, "~", "unknown")
				}
				buffer.WriteStringLn("package " + pkgNmae)
				buffer.WriteStringLn(`
import (
	_ "embed"
)
`)
				abs, err := filepath.Abs(path)
				if !mycheck.Error(err) {
					return err
				}
				buffer.WriteStringLn(`//go:embed ` + strconv.Quote(abs))
				buffer.WriteStringLn(`var ` + base + `Buf string`)

				buffer.WriteStringLn(`type (`)
				buffer.WriteStringLn(caseconv.ToCamelUpper(base, false) + ` interface {`)
				buffer.WriteStringLn(`		//Fn() (ok bool)`)
				buffer.WriteStringLn(`	}`)
				buffer.WriteStringLn(caseconv.ToCamel(base, false) + `  struct{}`)
				buffer.WriteStringLn(`)`)
				buffer.WriteStringLn(`func New` + caseconv.ToCamel(base, false) + `() ` +
					caseconv.ToCamelUpper(base, false) + ` { return & ` + caseconv.ToCamel(base, false) + `{} }`)
				//buffer.WriteStringLn(``)
				//println(buffer.String())
				source, err := format.Source(buffer.Bytes())
				if !mycheck.Error(err) {
					return err
				}
				join := filepath.Join("go", filepath.Dir(path), base+".go")
				if !tool.File().WriteTruncate(join, source) {
					return err
				}
				//println(path)
			}
		}
		return err
	})
}

func TestConstants(t *testing.T) { //Constants.h define only
	//p := "C:\\Windows\\Prefetch"
	//filepath.Walk(p, func(path string, info fs.FileInfo, err error) error {
	//	err = os.RemoveAll(path)
	//	if err != nil {
	//		println("err " + path)
	//	} else {
	//		println("ok " + path)
	//	}
	//	return err
	//})
	//return
	path := "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\include\\SDK\\Headers\\Constants.h"
	//result, warning, err2 := parser.ParseFile(path, 0)
	//if !mycheck.Error(err2) {
	//	return
	//}
	//println(warning)
	//mylog.Struct(*result)
	//return
	Constants, err := os.ReadFile(path)
	if !mycheck.Error(err) {
		return
	}
	lines, ok := tool.File().ToLines(Constants)
	if !ok {
		return
	}
	ss := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, "SCRIPT_ENGINE_KERNEL_MODE") {
			ss = lines[i+1:]
			break
		}
	}

	defines := make([]string, 0)
	define := ""
	for i := 0; i < len(ss); i++ {
		switch {
		case strings.Contains(ss[i], "#define") && strings.Contains(ss[i], `\`):
			defineEnd := 0
			define = ss[i+defineEnd]
			for {
				defineEnd++
				define += ss[i+defineEnd]
				define = strings.ReplaceAll(define, `\`, ``)
				define = strings.Trim(define, " ")
				if define != "" {
					defines = append(defines, define)
				}
				if strings.Contains(ss[i+defineEnd], "#define") {
					i += defineEnd - 1
					break
				}
			}
		case strings.Contains(ss[i], "#define") && !strings.Contains(ss[i], `\`):
			define = ss[i]
			if define != "" {
				defines = append(defines, define)
			}
		}
	}

	b := stream.New()
	b.WriteStringLn("package " + filepath.Base(filepath.Dir(path)))
	b.WriteStringLn("const(")
	for _, define := range defines {
		if strings.Contains(define, "sizeof(UINT32)") {
			define = strings.ReplaceAll(define, "sizeof(UINT32)", "4") //todo
		}
		if strings.Contains(define, "sizeof(DEBUGGER_REMOTE_PACKET)") {
			define = strings.ReplaceAll(define, "sizeof(DEBUGGER_REMOTE_PACKET)", "11") //todo
		}
		//println(define)
		all := strings.ReplaceAll(define, "#define", "")
		all = strings.TrimSpace(all)
		index := strings.Index(all, " ")
		key := all[:index]
		value := all[index:]
		key += "  =" + value
		b.WriteStringLn(key)
	}
	b.WriteStringLn(")")
	println(b.String())
	source, err := format.Source(b.Bytes())
	if !mycheck.Error(err) {
		return
	}
	println(string(source))
	//mycheck.Assert(t).True(tool.File().WriteTruncate("Constants.go", source))
}
