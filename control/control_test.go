package control

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/golibrary/stream"
)

func TestRemoveComment(t *testing.T) {
	// t.Skip("do not run this test")
	filepath.Walk("hprdbgctrl", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		switch ext {
		case ".c", ".cpp", ".h":
			removeCommentsFromFile(path)
		}
		return err
	})
}

func removeCommentsFromFile(filename string) {
	b := stream.NewBuffer(filename)
	re := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/|//.*`)
	processedContent := re.ReplaceAllString(b.String(), "")
	lines := stream.NewBuffer(processedContent).ToLines()
	f := stream.NewBuffer("")
	src := stream.NewBuffer("")
	src.WriteStringLn("package control")
	for i, line := range lines {
		switch strings.TrimSpace(line) {
		case "":
			line = ""
		case "{":
			line = "{"
		}
		if line == "" {
			continue
		}
		if line[0] != ' ' && line[len(line)-1] == '{' {
			// VOID PrintBits(const UINT32 Size, const void * Ptr){
			before, _, found := strings.Cut(line, "(")
			if found {
				index := strings.LastIndex(before, " ") // todo get return type
				if index != -1 {
					before = before[index:]
					before = strings.TrimSpace(before)
					comment := "//" + filename + ":" + fmt.Sprint(i+1)
					signature := "func " + before + "(){ " + comment
					signature += "\n"
					signature += "}"
					mylog.Success("", signature)
					src.WriteStringLn(signature)
				}
			}
		}
		f.WriteStringLn(line)
	}
	for _, returnType := range returnTypes {
		f.ReplaceAll(returnType+"\n", returnType+" ")
	}
	f.ReplaceAll("\n{", "{")
	stream.WriteTruncate(filename, f.String())

	if filepath.Ext(filename) == ".cpp" {
		all := strings.ReplaceAll(filename, ".cpp", ".go")
		join := filepath.Join("tmp", all)
		// mylog.Warning(all)
		stream.WriteGoFile(join, src)
	}
}

var returnTypes = []string{
	"VOID",
	"BOOLEAN",
	"BOOL",
	"UINT64",
	"UINT32",
	"double",
	"BYTE",
	"HANDLE",
	"char *",
	"size_t",
	"std::vector<std::string>",
	"vector<char>",
	"HPRDBGCTRL_API bool",
	"string",
	"const vector<string>",
	"VOID *",
	"DEBUGGER_OUTPUT_SOURCE_STATUS",
	"static ZyanStatus",
	"unsigned long long",
	"int",
	"void",
	"DEBUGGER_CONDITIONAL_JUMP_STATUS",
}

func getReturnType(pre string) string {
	pre = strings.TrimSuffix(pre, " ")
	switch pre {
	case "VOID":
		return "void"
	case "BOOLEAN":
		return "bool"
	case "BOOL":
		return "bool"
	case "UINT64":
		return "uint64"
	case "UINT32":
		return "uint32"
	case "double":
		return "float64"
	case "BYTE":
		return "byte"
	case "HANDLE":
		return "handle"
	case "char *":
		return "string"
	case "size_t":
		return "int"
	case "std::vector<std::string>":
		return "[]string"
	case "vector<char>":
		return "[]byte"
	case "HPRDBGCTRL_API bool":
		return "bool"
	case "string":
		return "string"
	case "const vector<string>":
		return "[]string"
	case "VOID *":
		return "unsafe.Pointer"
	case "DEBUGGER_OUTPUT_SOURCE_STATUS":
		return "debugger.OutputSourceStatus"
	case "static ZyanStatus":
		return "zyan.Status"
	case "unsigned long long":
		return "uint64"
	case "int":
		return "int"
	case "void":
		return "void"
	case "DEBUGGER_CONDITIONAL_JUMP_STATUS":
		return "debugger.ConditionalJumpStatus"
	default:
		return ""
	}
}
