package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	path := `C:\Windows\Minidump\032326-3984-01.dmp`

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("File size: %d bytes\n\n", len(data))

	fmt.Println("=== Header Analysis ===")
	fmt.Printf("Signature: %s\n", string(data[0:8]))
	fmt.Printf("ValidDump: 0x%04X\n", binary.LittleEndian.Uint16(data[8:10]))
	fmt.Printf("Version: 0x%04X\n", binary.LittleEndian.Uint16(data[10:12]))
	fmt.Printf("DirectoryTableBase: 0x%016X\n", binary.LittleEndian.Uint64(data[16:24]))
	fmt.Printf("PsLoadedModuleList: 0x%016X\n", binary.LittleEndian.Uint64(data[32:40]))
	fmt.Printf("ContextRecord: 0x%016X\n", binary.LittleEndian.Uint64(data[120:128]))
	fmt.Printf("DumpType: 0x%08X\n", binary.LittleEndian.Uint32(data[128:132]))

	fmt.Println("\n=== Context Record (at 0x120) ===")
	ctxOffset := 0x120
	fmt.Printf("ContextFlags: 0x%08X\n", binary.LittleEndian.Uint32(data[ctxOffset:ctxOffset+4]))
	fmt.Printf("Rax: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+48:ctxOffset+56]))
	fmt.Printf("Rbx: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+56:ctxOffset+64]))
	fmt.Printf("Rcx: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+64:ctxOffset+72]))
	fmt.Printf("Rdx: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+72:ctxOffset+80]))
	fmt.Printf("Rsp: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+80:ctxOffset+88]))
	fmt.Printf("Rbp: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+88:ctxOffset+96]))
	fmt.Printf("Rsi: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+96:ctxOffset+104]))
	fmt.Printf("Rdi: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+104:ctxOffset+112]))
	fmt.Printf("Rip: 0x%016X\n", binary.LittleEndian.Uint64(data[ctxOffset+120:ctxOffset+128]))

	fmt.Println("\n=== Physical Memory Block ===")
	physMemOffset := 0x160
	physMemRva := binary.LittleEndian.Uint64(data[physMemOffset:physMemOffset+8])
	fmt.Printf("PhysicalMemoryBlock RVA: 0x%016X\n", physMemRva)

	if physMemRva > 0 && physMemRva < uint64(len(data)) {
		numRuns := binary.LittleEndian.Uint64(data[physMemRva:physMemRva+8])
		fmt.Printf("Number of runs: %d\n", numRuns)

		for i := uint64(0); i < numRuns && i < 10; i++ {
			runOffset := physMemRva + 16 + i*24
			basePage := binary.LittleEndian.Uint64(data[runOffset:runOffset+8])
			pageCount := binary.LittleEndian.Uint64(data[runOffset+8:runOffset+16])
			fmt.Printf("  Run %d: BasePage=0x%016X, PageCount=0x%016X\n", i, basePage, pageCount)
		}
	}

	fmt.Println("\n=== First 512 bytes hex ===")
	for i := 0; i < 512 && i < len(data); i++ {
		fmt.Printf("%02X ", data[i])
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
}
