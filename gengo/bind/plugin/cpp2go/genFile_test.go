package cpp2go

import (
	"net/http"
	"testing"
)

//
//func TestGengo(t *testing.T) {
//	filepath.WalkDir(".", func(path string, info fs.DirEntry, err error) error {
//		switch filepath.Extensions(path) {
//		case ".h":
//			println(path)
//			file, _ := os.Create(path + ".go")
//			file.WriteString("package Headers\n")
//		}
//		return nil
//	})
//}

func TestTransmartByTencebt(t *testing.T) {
	handler := http.FileServer(http.Dir("/data/data/com.termux/files/home/storage/downloads/"))
	http.ListenAndServe(":7777", handler)
}
