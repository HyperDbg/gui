package main

import (
	"fmt"
	"os"

	hyper "github.com/HyperDbg/gui"
)

func main() {
	fmt.Printf("HyperDBG's GUI %s\n", hyper.Version)

	if err := hyper.Get().Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Run: %v\n", err)
		os.Exit(1)
	}
}
