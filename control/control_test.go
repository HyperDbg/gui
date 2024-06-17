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
					comment := "//" + filename + ":" + fmt.Sprint(i+1) // todo bug not add to every func
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
