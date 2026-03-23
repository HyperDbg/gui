package main

import (
	"fmt"
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
}
