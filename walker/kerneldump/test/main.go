package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/ddkwork/HyperDbg/walker/kerneldump"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dumpPath := `C:\Windows\Minidump\032326-3984-01.dmp`

	fmt.Println("=== Kernel Dump Parser ===")
	fmt.Printf("Loading: %s\n\n", dumpPath)

	kd, err := kerneldump.Parse(dumpPath)
	if err != nil {
		log.Fatal(err)
	}
	defer kd.Close()

	fmt.Println(kd)
	fmt.Println()

	fmt.Println("=== Bug Check Analysis ===")
	fmt.Printf("Bug Check Code: 0x%08X\n", kd.Header.BugCheckCode)
	fmt.Printf("Bug Check Name: %s\n", kd.GetBugCheckName())
	fmt.Printf("Parameter 1: 0x%016X\n", kd.Header.BugCheckParameter1)
	fmt.Printf("Parameter 2: 0x%016X\n", kd.Header.BugCheckParameter2)
	fmt.Printf("Parameter 3: 0x%016X\n", kd.Header.BugCheckParameter3)
	fmt.Printf("Parameter 4: 0x%016X\n", kd.Header.BugCheckParameter4)
	fmt.Println()

	fmt.Println("=== System Information ===")
	fmt.Printf("Machine Type: 0x%04X", kd.Header.MachineImageType)
	if kd.Header.MachineImageType == 0x8664 {
		fmt.Printf(" (AMD64)")
	}
	fmt.Println()
	fmt.Printf("Number of Processors: %d\n", kd.Header.NumberProcessors)
	fmt.Printf("Dump Type: %d\n", kd.Header.DumpType)
	fmt.Printf("Directory Table Base: 0x%016X\n", kd.Header.DirectoryTableBase)
	fmt.Printf("PsLoadedModuleList: 0x%016X\n", kd.Header.PsLoadedModuleList)
	fmt.Printf("PsActiveProcessHead: 0x%016X\n", kd.Header.PsActiveProcessHead)
	fmt.Println()

	fmt.Println("=== Analysis Summary ===")
	fmt.Printf("File: %s\n", filepath.Base(dumpPath))
	fmt.Printf("Size: %d bytes\n", len(kd.Data))
	fmt.Printf("Signature: %s\n", string(kd.Header.Signature[:]))
	fmt.Printf("Version: %d.%d\n", kd.Header.MajorVersion, kd.Header.MinorVersion)

	if kd.Header.BugCheckCode == 0x000000EF || kd.Header.BugCheckCode == 0x000000F4 {
		fmt.Println()
		fmt.Println("!!! CRITICAL PROCESS DIED !!!")
		fmt.Println("This indicates a critical system process or driver crashed.")
		fmt.Println("Common causes:")
		fmt.Println("  - Driver callback not cleaned up before unload")
		fmt.Println("  - EPT hook not removed before driver unload")
		fmt.Println("  - Memory corruption in kernel space")
		fmt.Println("  - Improper use of kernel APIs")
	}

	if kd.Header.BugCheckCode == 0x00000008 || kd.Header.BugCheckCode == 0x0000000A {
		fmt.Println()
		fmt.Println("!!! IRQL NOT LESS OR EQUAL !!!")
		fmt.Println("This indicates a driver accessed memory at an invalid IRQL.")
		fmt.Println("Common causes:")
		fmt.Println("  - Accessing paged memory at DISPATCH_LEVEL or higher")
		fmt.Println("  - Using invalid memory addresses")
		fmt.Println("  - Driver not properly handling IRQL levels")
	}
}
