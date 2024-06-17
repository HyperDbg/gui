package control

import (
	"io/fs"
	"path/filepath"
	"regexp"
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
		if line == "" {
			continue
		}
		f.WriteStringLn(line)
	}
	f.ReplaceAll("VOID\n", "VOID ")
	stream.WriteTruncate(filename, f.String())
}
