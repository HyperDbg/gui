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
	fmt.Printf("First 256 bytes:\n")

	for i := 0; i < 256 && i < len(data); i++ {
		fmt.Printf("%02X ", data[i])
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}

	fmt.Println("\n\nHeader analysis:")
	fmt.Printf("Signature (0-7): %s\n", string(data[0:8]))
	fmt.Printf("Signature hex: %02X %02X %02X %02X %02X %02X %02X %02X\n",
		data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7])

	if len(data) >= 16 {
		fmt.Printf("ValidDump: 0x%04X\n", binary.LittleEndian.Uint16(data[8:10]))
		fmt.Printf("MajorVersion: 0x%04X\n", binary.LittleEndian.Uint16(data[10:12]))
		fmt.Printf("MinorVersion: 0x%04X\n", binary.LittleEndian.Uint16(data[12:14]))
		fmt.Printf("DirectoryTableBase: 0x%016X\n", binary.LittleEndian.Uint64(data[16:24]))
		fmt.Printf("PfnDatabase: 0x%016X\n", binary.LittleEndian.Uint64(data[24:32]))
		fmt.Printf("PsLoadedModuleList: 0x%016X\n", binary.LittleEndian.Uint64(data[32:40]))
		fmt.Printf("PsActiveProcessHead: 0x%016X\n", binary.LittleEndian.Uint64(data[40:48]))
		fmt.Printf("MachineImageType: 0x%04X\n", binary.LittleEndian.Uint16(data[48:50]))
		fmt.Printf("NumberProcessors: %d\n", data[50])
		fmt.Printf("BugCheckCode: 0x%08X\n", binary.LittleEndian.Uint32(data[52:56]))
		fmt.Printf("BugCheckParameter1: 0x%016X\n", binary.LittleEndian.Uint64(data[56:64]))
		fmt.Printf("BugCheckParameter2: 0x%016X\n", binary.LittleEndian.Uint64(data[64:72]))
		fmt.Printf("BugCheckParameter3: 0x%016X\n", binary.LittleEndian.Uint64(data[72:80]))
		fmt.Printf("BugCheckParameter4: 0x%016X\n", binary.LittleEndian.Uint64(data[80:88]))
	}
}
