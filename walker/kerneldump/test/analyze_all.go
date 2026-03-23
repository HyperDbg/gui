package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dumpDir := `C:\Windows\Minidump`

	entries, err := filepath.Glob(filepath.Join(dumpDir, "*.dmp"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("=== Analyzing %d dump files ===\n\n", len(entries))

	for i, dumpPath := range entries {
		fmt.Printf("[%d] %s\n", i+1, filepath.Base(dumpPath))

		kd, err := kerneldump.Parse(dumpPath)
		if err != nil {
			fmt.Printf("  ERROR: %v\n\n", err)
			continue
		}
		defer kd.Close()

		fmt.Printf("  BugCheck: 0x%08X (%s)\n", kd.Header.BugCheckCode, kd.GetBugCheckName())
		fmt.Printf("  Params: 0x%016X 0x%016X 0x%016X 0x%016X\n",
			kd.Header.BugCheckParameter1,
			kd.Header.BugCheckParameter2,
			kd.Header.BugCheckParameter3,
			kd.Header.BugCheckParameter4)
		fmt.Printf("  Version: %d.%d\n", kd.Header.MajorVersion, kd.Header.MinorVersion)
		fmt.Printf("  Size: %d bytes\n", len(kd.Data))
		fmt.Println()
	}
}
