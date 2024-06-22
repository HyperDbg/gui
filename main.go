package main

import "github.com/ddkwork/hyperdbgui/ux"

//go:generate go build .
//go:generate go install .

func main() {
	ux.Run()
}
