package control

import (
	"io/fs"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"
)

func TestRemoveComment(t *testing.T) {
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
	for _, line := range lines {
		switch strings.TrimSpace(line) {
		case "":
			line = ""
		case "{":
			line = "{"
		}
		if line == "" {
			continue
		}
		f.WriteStringLn(line)
	}
	f.ReplaceAll("VOID\n", "VOID ")
	f.ReplaceAll("BOOLEAN\n", "BOOLEAN ")
	f.ReplaceAll("BOOL\n", "BOOL ")
	f.ReplaceAll("UINT64\n", "UINT64 ")
	f.ReplaceAll("UINT32\n", "UINT32 ")
	f.ReplaceAll("double\n", "double ")
	f.ReplaceAll("BYTE\n", "BYTE ")
	f.ReplaceAll("HANDLE\n", "HANDLE ")
	f.ReplaceAll("char *\n", "char * ")
	f.ReplaceAll("size_t\n", "size_t ")
	f.ReplaceAll("std::vector<std::string>\n", "std::vector<std::string> ")
	f.ReplaceAll("vector<char>\n", "vector<char> ")
	f.ReplaceAll("HPRDBGCTRL_API bool\n", "HPRDBGCTRL_API bool ")
	f.ReplaceAll("string\n", "string ")
	f.ReplaceAll("const vector<string>\n", "const vector<string> ")
	f.ReplaceAll("VOID *\n", "VOID * ")
	f.ReplaceAll("DEBUGGER_OUTPUT_SOURCE_STATUS\n", "DEBUGGER_OUTPUT_SOURCE_STATUS ")
	f.ReplaceAll("static ZyanStatus\n", "static ZyanStatus ")
	f.ReplaceAll("unsigned long long\n", "unsigned long long ")
	f.ReplaceAll("int\n", "int ")
	f.ReplaceAll("void\n", "void ")
	f.ReplaceAll("\n{", "{")
	f.ReplaceAll("DEBUGGER_CONDITIONAL_JUMP_STATUS\n", "DEBUGGER_CONDITIONAL_JUMP_STATUS ")

	stream.WriteTruncate(filename, f.String())
}
