package sdk

import (
	_ "embed"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"go/format"
	"os"
	"strings"
	"testing"
)

func TestConstants(t *testing.T) { //Constants.h define only
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
	b.WriteStringLn("const(")
	for _, define := range defines {
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
