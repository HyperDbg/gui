// Package main — diagnose 32-bit kernelbase.dll export RVAs and check
// whether SetEvent / VirtualAlloc are forwarders (which would give a
// wrong code address for EPT hooking).
package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func le16(d []byte, o uint32) uint16 { return binary.LittleEndian.Uint16(d[o:]) }
func le32(d []byte, o uint32) uint32 { return binary.LittleEndian.Uint32(d[o:]) }

func rvaToOff(data []byte, coffOff uint32, numSec uint16, sizeOpt uint16, rva uint32) (uint32, bool) {
	secOff := coffOff + 20 + uint32(sizeOpt)
	for i := uint16(0); i < numSec; i++ {
		off := secOff + uint32(i)*40
		if int(off)+40 > len(data) {
			continue
		}
		vSize := le32(data, off+8)
		vAddr := le32(data, off+12)
		rawOff := le32(data, off+20)
		if rva >= vAddr && rva < vAddr+vSize {
			return rawOff + (rva - vAddr), true
		}
	}
	return 0, false
}

func main() {
	pePath := `C:\Windows\SysWOW64\kernelbase.dll`
	data, err := os.ReadFile(pePath)
	if err != nil {
		fmt.Printf("read %s: %v\n", pePath, err)
		return
	}
	if data[0] != 'M' || data[1] != 'Z' {
		fmt.Println("not a PE file")
		return
	}
	e_lfanew := le32(data, 0x3C)
	if data[e_lfanew] != 'P' || data[e_lfanew+1] != 'E' {
		fmt.Println("bad PE signature")
		return
	}
	coffOff := e_lfanew + 4
	numSec := le16(data, coffOff+2)
	sizeOpt := le16(data, coffOff+16)
	optOff := coffOff + 20
	magic := le16(data, optOff)
	fmt.Printf("PE magic: 0x%X (%s)\n", magic, map[uint16]string{0x10B: "PE32", 0x20B: "PE32+"}[magic])

	var imageBase, sizeOfImage, exportRVA, exportSize uint32
	if magic == 0x20B {
		fmt.Println("ERROR: this is a 64-bit DLL, expected 32-bit")
		return
	} else {
		imageBase = le32(data, optOff+28)
		sizeOfImage = le32(data, optOff+56)
		exportRVA = le32(data, optOff+96)
		exportSize = le32(data, optOff+100)
	}
	fmt.Printf("ImageBase: 0x%X  SizeOfImage: 0x%X (%d bytes)\n", imageBase, sizeOfImage, sizeOfImage)
	fmt.Printf("Export Directory: RVA=0x%X Size=0x%X\n", exportRVA, exportSize)

	// Section table
	secOff := coffOff + 20 + uint32(sizeOpt)
	fmt.Println("\nSections:")
	for i := uint16(0); i < numSec; i++ {
		off := secOff + uint32(i)*40
		name := string(data[off : off+8])
		vSize := le32(data, off+8)
		vAddr := le32(data, off+12)
		rawSize := le32(data, off+16)
		rawOff := le32(data, off+20)
		chars := le32(data, off+36)
		fmt.Printf("  %-8s VAddr=0x%08X VSize=0x%08X RawOff=0x%08X RawSize=0x%08X Chars=0x%08X\n",
			name, vAddr, vSize, rawOff, rawSize, chars)
	}

	expOff, ok := rvaToOff(data, coffOff, numSec, sizeOpt, exportRVA)
	if !ok {
		fmt.Println("export dir not found in sections")
		return
	}
	numNames := le32(data, expOff+24)
	addrFuncs := le32(data, expOff+28)
	addrNames := le32(data, expOff+32)
	addrOrds := le32(data, expOff+36)

	namesOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, addrNames)
	ordsOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, addrOrds)
	funcsOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, addrFuncs)

	fmt.Printf("\nExports: %d named functions\n", numNames)
	fmt.Println("Export directory range: [0x%X, 0x%X) — RVAs in this range are FORWARDERS",
		exportRVA, exportRVA+exportSize)

	targets := []string{"SetEvent", "VirtualAlloc", "RtlAllocateHeap"}
	for _, target := range targets {
		found := false
		for i := uint32(0); i < numNames; i++ {
			nameRVA := le32(data, namesOff+i*4)
			nOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, nameRVA)
			end := nOff
			for end < uint32(len(data)) && data[end] != 0 {
				end++
			}
			name := string(data[nOff:end])
			if name != target {
				continue
			}
			ord := le16(data, ordsOff+i*2)
			funcRVA := le32(data, funcsOff+uint32(ord)*4)
			isForwarder := funcRVA >= exportRVA && funcRVA < exportRVA+exportSize
			fmt.Printf("\n%s: ordinal=%d funcRVA=0x%X", name, ord, funcRVA)
			if isForwarder {
				// Read forwarder string
				fOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, funcRVA)
				fEnd := fOff
				for fEnd < uint32(len(data)) && data[fEnd] != 0 {
					fEnd++
				}
				fwdStr := string(data[fOff:fEnd])
				fmt.Printf("  FORWARDER -> %s", fwdStr)
			} else {
				// Read first bytes at the function
				fOff, _ := rvaToOff(data, coffOff, numSec, sizeOpt, funcRVA)
				if fOff+8 <= uint32(len(data)) {
					fmt.Printf("  bytes: %02X %02X %02X %02X %02X %02X %02X %02X",
						data[fOff], data[fOff+1], data[fOff+2], data[fOff+3],
						data[fOff+4], data[fOff+5], data[fOff+6], data[fOff+7])
				}
			}
			fmt.Println()
			found = true
			break
		}
		if !found {
			fmt.Printf("\n%s: NOT FOUND in %s\n", target, pePath)
		}
	}
}
