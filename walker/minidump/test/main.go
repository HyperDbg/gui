package main

import (
	"fmt"
	"log"

	"github.com/ddkwork/HyperDbg/walker/minidump"
)

func main() {
	log.SetFlags(log.Lshortfile)

	dumpPath := `C:\Windows\Minidump\032326-3984-01.dmp`

	fmt.Println("=== Minidump Parser ===")
	fmt.Printf("Loading: %s\n\n", dumpPath)

	m, err := minidump.Parse(dumpPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
	fmt.Println()

	if exc, err := m.GetException(); err == nil {
		fmt.Println("=== Exception ===")
		fmt.Printf("  Code: 0x%08X\n", exc.ExceptionRecord.ExceptionCode)
		fmt.Printf("  Flags: 0x%08X\n", exc.ExceptionRecord.ExceptionFlags)
		fmt.Printf("  Address: 0x%016X\n", exc.ExceptionRecord.ExceptionAddress)
		fmt.Printf("  Thread ID: %d\n", exc.ThreadId)
		fmt.Printf("  RIP: 0x%016X\n", exc.ThreadContext.Rip)
		fmt.Printf("  RSP: 0x%016X\n", exc.ThreadContext.Rsp)
		fmt.Println()
	}

	if sysInfo, err := m.GetSystemInfo(); err == nil {
		fmt.Println("=== System Info ===")
		fmt.Printf("  Architecture: %d\n", sysInfo.ProcessorArchitecture)
		fmt.Printf("  Processors: %d\n", sysInfo.NumberOfProcessors)
		fmt.Printf("  Version: %d.%d.%d\n", sysInfo.MajorVersion, sysInfo.MinorVersion, sysInfo.BuildNumber)
		fmt.Printf("  Platform: %d\n", sysInfo.PlatformId)
		fmt.Println()
	}

	if threads, err := m.GetThreads(); err == nil {
		fmt.Printf("=== Threads (%d) ===\n", len(threads))
		for i, t := range threads {
			if i >= 5 {
				fmt.Printf("  ... and %d more\n", len(threads)-5)
				break
			}
			fmt.Printf("  Thread %d: ID=%d, TEB=0x%016X\n", i, t.ThreadId, t.Teb)
		}
		fmt.Println()
	}

	if modules, names, err := m.GetModules(); err == nil {
		fmt.Printf("=== Modules (%d) ===\n", len(modules))
		for i, mod := range modules {
			if i >= 10 {
				fmt.Printf("  ... and %d more\n", len(modules)-10)
				break
			}
			name := "?"
			if i < len(names) && names[i] != "" {
				name = names[i]
			}
			fmt.Printf("  %s: Base=0x%016X, Size=0x%08X\n", name, mod.BaseOfImage, mod.SizeOfImage)
		}
		fmt.Println()
	}

	if memList, err := m.GetMemory64List(); err == nil {
		fmt.Printf("=== Memory64List (%d ranges) ===\n", memList.NumberOfMemoryRanges)
		for i, r := range memList.MemoryRanges {
			if i >= 5 {
				fmt.Printf("  ... and %d more\n", len(memList.MemoryRanges)-5)
				break
			}
			fmt.Printf("  Range %d: 0x%016X - 0x%016X (size: 0x%016X)\n",
				i, r.StartOfMemoryRange, r.StartOfMemoryRange+r.DataSize, r.DataSize)
		}
	}
}
