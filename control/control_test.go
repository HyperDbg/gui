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
			if ext == ".cpp" {
				all := strings.ReplaceAll(path, ".cpp", ".go")
				join := filepath.Join("tmp", all)
				// mylog.Warning(all)
				name := "package main"
				stream.WriteGoFile(join, name)
			}
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
		comment := "//" + filename + ":" + fmt.Sprint(i+1)
		if line[0] != ' ' && line[len(line)-1] == '{' {
			// VOID PrintBits(const UINT32 Size, const void * Ptr){

			before, found := strings.CutSuffix(line, "(")
			if found {
				after, found2 := strings.CutPrefix(before, " ")
				if found2 {
					signature := "func " + after + "(){ " + comment
					signature += "\n"
					signature += "}"
					mylog.Success("", signature)
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
