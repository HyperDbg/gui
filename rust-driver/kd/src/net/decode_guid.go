package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

func (g GUID) String() string {
	return fmt.Sprintf("{%08X-%04X-%04X-%02X%02X-%02X%02X%02X%02X%02X%02X}",
		g.Data1, g.Data2, g.Data3,
		g.Data4[0], g.Data4[1], g.Data4[2], g.Data4[3],
		g.Data4[4], g.Data4[5], g.Data4[6], g.Data4[7])
}

func (g GUID) RustCode() string {
	return fmt.Sprintf(`pub static NPI_WSK_INTERFACE_ID: NPIID = NPIID {
    Data1: 0x%08X,
    Data2: 0x%04X,
    Data3: 0x%04X,
    Data4: [0x%02X, 0x%02X, 0x%02X, 0x%02X, 0x%02X, 0x%02X, 0x%02X, 0x%02X],
};`,
		g.Data1, g.Data2, g.Data3,
		g.Data4[0], g.Data4[1], g.Data4[2], g.Data4[3],
		g.Data4[4], g.Data4[5], g.Data4[6], g.Data4[7])
}

type ArchiveHeader struct {
	Name     [16]byte
	Date     [12]byte
	UID      [6]byte
	GID      [6]byte
	Mode     [8]byte
	Size     [10]byte
	EndMagic [2]byte
}

type COFFHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type SectionHeader struct {
	Name                 [8]byte
	VirtualSize          uint32
	VirtualAddress       uint32
	SizeOfRawData        uint32
	PointerToRawData     uint32
	PointerToRelocations uint32
	PointerToLinenumbers uint32
	NumberOfRelocations  uint16
	NumberOfLinenumbers  uint16
	Characteristics      uint32
}

type COFFSymbol struct {
	Name               [8]byte
	Value              uint32
	SectionNumber      int16
	Type               uint16
	StorageClass       uint8
	NumberOfAuxSymbols uint8
}

func main() {
	libPath := `E:\Program Files\Windows Kits\10\Lib\10.0.28000.0\um\x64\Uuid.Lib`

	data, err := os.ReadFile(libPath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Printf("=== 解析 uuid.lib 中的 NPI_WSK_INTERFACE_ID GUID ===\n\n")
	fmt.Printf("文件大小: %d bytes\n\n", len(data))

	magic := string(data[0:8])
	fmt.Printf("Archive Magic: %q\n", magic)

	if magic != "!<arch>\n" {
		fmt.Println("Not a valid archive file")
		return
	}

	target := "NPI_WSK_INTERFACE_ID"
	found := false

	offset := 8
	for offset < len(data) {
		if offset+60 > len(data) {
			break
		}

		var header ArchiveHeader
		copy(header.Name[:], data[offset:offset+16])
		copy(header.Size[:], data[offset+48:offset+58])

		sizeStr := string(bytes.TrimRight(header.Size[:], " \x00"))
		var memberSize int
		fmt.Sscanf(sizeStr, "%d", &memberSize)

		memberStart := offset + 60
		memberData := data[memberStart : memberStart+memberSize]

		if len(memberData) >= 20 {
			r := bytes.NewReader(memberData)
			var coffHeader COFFHeader
			if err := binary.Read(r, binary.LittleEndian, &coffHeader); err == nil {

				if coffHeader.PointerToSymbolTable > 0 && coffHeader.NumberOfSymbols > 0 {

					sections := make([]SectionHeader, coffHeader.NumberOfSections)
					for i := 0; i < int(coffHeader.NumberOfSections); i++ {
						if err := binary.Read(r, binary.LittleEndian, &sections[i]); err != nil {
							break
						}
					}

					stringTableOffset := coffHeader.PointerToSymbolTable + 18*uint32(coffHeader.NumberOfSymbols)

					r.Seek(int64(coffHeader.PointerToSymbolTable), 0)

					for i := 0; i < int(coffHeader.NumberOfSymbols); i++ {
						var sym COFFSymbol
						if err := binary.Read(r, binary.LittleEndian, &sym); err != nil {
							break
						}

						var name string
						if sym.Name[0] == 0 && sym.Name[1] == 0 {
							stringOffset := binary.LittleEndian.Uint32(sym.Name[4:8])
							if int(stringTableOffset+stringOffset) < len(memberData) {
								end := bytes.IndexByte(memberData[stringTableOffset+stringOffset:], 0)
								if end >= 0 {
									name = string(memberData[stringTableOffset+stringOffset : stringTableOffset+stringOffset+uint32(end)])
								}
							}
						} else {
							name = string(bytes.TrimRight(sym.Name[:], "\x00"))
						}

						if name == target {
							found = true
							fmt.Printf("\n找到符号: %s\n", name)
							fmt.Printf("  Member Offset: 0x%X\n", memberStart)
							fmt.Printf("  Symbol Index: %d\n", i)
							fmt.Printf("  Value: 0x%X\n", sym.Value)
							fmt.Printf("  Section Number: %d\n", sym.SectionNumber)
							fmt.Printf("  Storage Class: %d\n\n", sym.StorageClass)

							if sym.SectionNumber > 0 && sym.SectionNumber <= int16(len(sections)) {
								section := sections[sym.SectionNumber-1]
								sectionName := string(bytes.TrimRight(section.Name[:], "\x00"))
								fmt.Printf("所在节:\n")
								fmt.Printf("  Name: %s\n", sectionName)
								fmt.Printf("  PointerToRawData: 0x%X\n", section.PointerToRawData)
								fmt.Printf("  SizeOfRawData: %d\n\n", section.SizeOfRawData)

								guidOffset := int(section.PointerToRawData + sym.Value)
								fmt.Printf("GUID 数据偏移: 0x%X\n\n", guidOffset)

								if guidOffset+16 <= len(memberData) {
									var guid GUID
									buf := bytes.NewReader(memberData[guidOffset : guidOffset+16])
									binary.Read(buf, binary.LittleEndian, &guid)

									fmt.Printf("=== 解析结果 ===\n\n")
									fmt.Printf("GUID 结构:\n")
									fmt.Printf("  Data1: 0x%08X\n", guid.Data1)
									fmt.Printf("  Data2: 0x%04X\n", guid.Data2)
									fmt.Printf("  Data3: 0x%04X\n", guid.Data3)
									fmt.Printf("  Data4: % X\n\n", guid.Data4)

									fmt.Printf("GUID 字符串:\n")
									fmt.Printf("  %s\n\n", guid.String())

									fmt.Printf("Rust 代码:\n")
									fmt.Printf("%s\n", guid.RustCode())
								}
							}
							break
						}

						for j := 0; j < int(sym.NumberOfAuxSymbols); j++ {
							r.Seek(18, 1)
							i++
						}
					}
				}
			}
		}

		if found {
			break
		}

		offset = memberStart + memberSize
		if offset%2 == 1 {
			offset++
		}
	}

	if !found {
		fmt.Println("\n未通过 COFF 解析找到符号，显示手动验证结果:")
		fmt.Println("\n从 dumpbin /all 输出中提取的原始数据:")
		fmt.Println("RAW DATA #6")
		fmt.Println("  00000000: 03 E8 27 22 8B 8D D4 11 AB AD 00 90 27 71 9E 09")

		fmt.Println("\n解析为 GUID (Little Endian):")
		fmt.Println("  Data1: 03 E8 27 22 -> 0x2227E803")
		fmt.Println("  Data2: 8B 8D -> 0x8D8B")
		fmt.Println("  Data3: D4 11 -> 0x11D4")
		fmt.Println("  Data4: AB AD 00 90 27 71 9E 09")

		fmt.Println("\n最终结果:")
		fmt.Println("  NPI_WSK_INTERFACE_ID = {2227E803-8D8B-11D4-ABAD-009027719E09}")
	}
}
