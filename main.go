package main

import "github.com/ddkwork/hyperdbgui/ux"

//go:generate core generate
//go:generate core build -v -t android/arm64
//go:generate core build -v -t windows/amd64
//go:generate go build .
//go:generate go install .
//go:generate svg embed-image 1.png

func main() {
	ux.Run()
}
