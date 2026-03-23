package main

import (
	"fmt"
	"log"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dumpPath := `C:\Windows\Minidump\032326-3984-01.dmp`

	fmt.Println("=== Kernel Dump Module Parser ===")
	fmt.Printf("Loading: %s\n\n", dumpPath)

	kd, err := kerneldump.Parse(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer kd.Close()

	fmt.Println(kd)
	fmt.Println()

	fmt.Println("=== Physical Memory Block ===")
	pmb, err := kd.GetPhysicalMemoryBlock()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Number of runs: %d\n", pmb.NumberOfRuns)
		for i, run := range pmb.Runs {
			if i >= 10 {
				fmt.Printf("  ... and %d more\n", len(pmb.Runs)-10)
				break
			}
			fmt.Printf("  Run %d: BasePage=0x%016X, PageCount=0x%016X (0x%016X bytes)\n",
				i, run.BasePage, run.PageCount, run.PageCount*4096)
		}
	}
	fmt.Println()

	fmt.Println("=== Modules ===")
	modules, err := kd.GetModules()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found %d modules\n", len(modules))
		for i, mod := range modules {
			if i >= 20 {
				fmt.Printf("  ... and %d more\n", len(modules)-20)
				break
			}
			fmt.Printf("  %s: Base=0x%016X, Size=0x%08X\n",
				mod.GetName(), mod.DllBase, mod.SizeOfImage)
		}
	}
}
