package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ddkwork/HyperDbg/walker/bsodanalyzer"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dumpPath := `C:\Windows\Minidump\032326-3984-01.dmp`
	driverPDB := `C:\Windows\Minidump\hyperkd.pdb`

	fmt.Println("=== BSOD Analyzer ===")
	fmt.Printf("Dump: %s\n", filepath.Base(dumpPath))
	fmt.Printf("Driver PDB: %s\n\n", filepath.Base(driverPDB))

	analyzer, err := bsodanalyzer.New(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer analyzer.Close()

	if _, err := filepath.Glob(driverPDB); err == nil {
		fmt.Println("Loading driver PDB...")
		if err := analyzer.LoadDriverPDB(driverPDB); err != nil {
			fmt.Printf("Warning: Failed to load PDB: %v\n", err)
			fmt.Println("Continuing without symbol resolution...\n")
		} else {
			fmt.Println("PDB loaded successfully\n")
		}
	} else {
		fmt.Println("Driver PDB not found, continuing without symbol resolution\n")
	}

	result, err := analyzer.Analyze()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
