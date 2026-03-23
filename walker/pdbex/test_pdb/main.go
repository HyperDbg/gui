package main

import (
	"fmt"
	"os"
	"sort"

	pdbex "github.com/ddkwork/HyperDbg/walker/pdbex"
)

func main() {
	fmt.Println("Testing PDB...")

	pdb := pdbex.NewPDB()
	defer pdb.Close()

	pdbPath := `D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.pdb`

	err := pdb.Open(pdbPath)
	if err != nil {
		fmt.Printf("Failed to open PDB: %v\n", err)
		return
	}

	fmt.Println("PDB opened successfully!")

	functions := pdb.GetAllFunctions()
	fmt.Printf("Total functions: %d\n", len(functions))

	names := make([]string, 0, len(functions))
	for name := range functions {
		names = append(names, name)
	}
	sort.Strings(names)

	outputPath := `D:\笔记本\p\ux\examples\hypedbg\walker\pdbex\test_pdb\functions.txt`
	file, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Failed to create output file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "Total functions: %d\n\n", len(functions))
	fmt.Fprintln(file, "Function names:")
	fmt.Fprintln(file, "===============")
	for i, name := range names {
		fn := functions[name]
		fmt.Fprintf(file, "%4d. %s (Address: 0x%X, Size: %d)\n", i+1, name, fn.Address, fn.Size)
	}

	fmt.Printf("Function list written to: %s\n", outputPath)
}
