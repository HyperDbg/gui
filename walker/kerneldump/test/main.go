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

	detectedBugCheck, detectedBugCheckName := kd.DetectBugCheckFromPhysical()

	if detectedBugCheck != 0 && detectedBugCheck != kd.Header.BugCheckCode {
		fmt.Println("!!! WARNING: Header BugCheck may be inaccurate !!!")
		fmt.Printf("!!! Header BugCheck: 0x%08X (%s) - IGNORED\n", kd.Header.BugCheckCode, kd.GetBugCheckName())
		fmt.Printf("!!! Stack BugCheck: 0x%08X (%s) - USED\n", detectedBugCheck, detectedBugCheckName)
		fmt.Println()
		fmt.Println("Reason: Real BugCheck detected from stack trace, not from dump header")
		fmt.Println()
	}

	bugCheckCode := detectedBugCheck
	if bugCheckCode == 0 {
		bugCheckCode = kd.Header.BugCheckCode
		fmt.Println("Note: Using header BugCheck (no stack detection)")
		fmt.Println()
	}

	fmt.Printf("Bug Check Code: 0x%08X\n", bugCheckCode)
	fmt.Printf("Bug Check Name: %s\n", kd.GetBugCheckNameFromCode(bugCheckCode))

	if detectedBugCheck != 0 && detectedBugCheck != kd.Header.BugCheckCode {
		fmt.Println("!!! WARNING: Parameters may be inaccurate (from header) !!!")
	}

	fmt.Printf("Parameter 1: 0x%016X\n", kd.Header.BugCheckParameter1)
	fmt.Printf("Parameter 2: 0x%016X\n", kd.Header.BugCheckParameter2)
	fmt.Printf("Parameter 3: 0x%016X\n", kd.Header.BugCheckParameter3)
	fmt.Printf("Parameter 4: 0x%016X\n", kd.Header.BugCheckParameter4)
	fmt.Println()

	fmt.Println("=== Exception Information ===")
	exception, err := kd.GetExceptionRecord()
	if err != nil {
		fmt.Printf("No exception record: %v\n", err)
	} else {
		fmt.Printf("Exception Code: 0x%08X (%s)\n", exception.ExceptionCode, kd.GetExceptionCodeName(exception.ExceptionCode))
		fmt.Printf("Exception Address: 0x%016X\n", exception.ExceptionAddress)
		fmt.Printf("Exception Flags: 0x%08X\n", exception.ExceptionFlags)
		fmt.Printf("Number Parameters: %d\n", exception.NumberParameters)
		if exception.NumberParameters > 0 {
			fmt.Println("Exception Information:")
			for i := uint32(0); i < exception.NumberParameters && i < 15; i++ {
				fmt.Printf("  [%d] 0x%016X\n", i, exception.ExceptionInformation[i])
			}
		}
	}
	fmt.Println()

	fmt.Println("=== Module Information ===")
	modules, err := kd.GetModules()
	if err != nil {
		fmt.Printf("Error reading modules: %v\n", err)
	} else {
		fmt.Printf("Total modules: %d\n", len(modules))
		if len(modules) > 0 {
			fmt.Println("Loaded modules:")
			for i, mod := range modules {
				if i >= 10 {
					fmt.Printf("  ... and %d more\n", len(modules)-10)
					break
				}
				fmt.Printf("  [%d] %s\n", i, mod.GetName())
				fmt.Printf("       Base: 0x%016X, Size: 0x%08X\n", mod.DllBase, mod.SizeOfImage)
			}
		}
	}
	fmt.Println()

	fmt.Println("=== Stack Trace ===")
	frames, err := kd.GetStackTraceFromPhysical()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		if len(frames) > 0 {
			for _, frame := range frames {
				fmt.Printf("#%d %s\n", frame.FrameNum, frame.CallSite)
			}
		} else {
			fmt.Println("No stack frames found")
		}
	}
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
	fmt.Println()

	if bugCheckCode == 0x000000EF || bugCheckCode == 0x000000F4 {
		fmt.Println("!!! CRITICAL PROCESS DIED !!!")
		fmt.Println("This indicates a critical system process or driver crashed.")
		fmt.Println("Common causes:")
		fmt.Println("  - Driver callback not cleaned up before unload")
		fmt.Println("  - EPT hook not removed before driver unload")
		fmt.Println("  - Memory corruption in kernel space")
		fmt.Println("  - Improper use of kernel APIs")
	}

	if bugCheckCode == 0x00000008 || bugCheckCode == 0x0000000A {
		fmt.Println("!!! IRQL NOT LESS OR EQUAL !!!")
		fmt.Println("This indicates a driver accessed memory at an invalid IRQL.")
		fmt.Println("Common causes:")
		fmt.Println("  - Accessing paged memory at DISPATCH_LEVEL or higher")
		fmt.Println("  - Using invalid memory addresses")
		fmt.Println("  - Driver not properly handling IRQL levels")
	}
}
