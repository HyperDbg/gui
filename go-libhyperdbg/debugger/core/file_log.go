package core

import "os"

func init() {
	openFileForLog = func(path string) (WriteCloser, error) {
		return os.Create(path)
	}
}
