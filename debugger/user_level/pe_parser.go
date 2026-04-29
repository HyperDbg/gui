package user_level

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"
	"unsafe"
)

type ImageDosHeader struct {
	E_magic    uint16
	E_cblp     uint16
	E_cp       uint16
	E_crlc     uint16
	E_cparhdr  uint16
	E_minalloc uint16
	E_maxalloc uint16
	E_ss       uint16
	E_sp       uint16
	E_csum     uint16
	E_ip       uint16
	E_cs       uint16
	E_lfarlc   uint16
	E_ovno     uint16
	E_res      [4]uint16
	E_oemid    uint16
	E_oeminfo  uint16
	E_res2     [10]uint16
	E_lfanew   int32
}

type ImageFileHeader struct {
	Machine              uint16
	NumberOfSections     uint16
	TimeDateStamp        uint32
	PointerToSymbolTable uint32
	NumberOfSymbols      uint32
	SizeOfOptionalHeader uint16
	Characteristics      uint16
}

type ImageOptionalHeader32 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	BaseOfData                  uint32
	ImageBase                   uint32
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint32
	SizeOfStackCommit           uint32
	SizeOfHeapReserve           uint32
	SizeOfHeapCommit            uint32
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]ImageDataDirectory
}

type ImageOptionalHeader64 struct {
	Magic                       uint16
	MajorLinkerVersion          uint8
	MinorLinkerVersion          uint8
	SizeOfCode                  uint32
	SizeOfInitializedData       uint32
	SizeOfUninitializedData     uint32
	AddressOfEntryPoint         uint32
	BaseOfCode                  uint32
	ImageBase                   uint64
	SectionAlignment            uint32
	FileAlignment               uint32
	MajorOperatingSystemVersion uint16
	MinorOperatingSystemVersion uint16
	MajorImageVersion           uint16
	MinorImageVersion           uint16
	MajorSubsystemVersion       uint16
	MinorSubsystemVersion       uint16
	Win32VersionValue           uint32
	SizeOfImage                 uint32
	SizeOfHeaders               uint32
	CheckSum                    uint32
	Subsystem                   uint16
	DllCharacteristics          uint16
	SizeOfStackReserve          uint64
	SizeOfStackCommit           uint64
	SizeOfHeapReserve           uint64
	SizeOfHeapCommit            uint64
	LoaderFlags                 uint32
	NumberOfRvaAndSizes         uint32
	DataDirectory               [16]ImageDataDirectory
}

type ImageDataDirectory struct {
	VirtualAddress uint32
	Isize          uint32
}

type ImageSectionHeader struct {
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

const (
	ImageDosSignature     = 0x5A4D
	ImageNtSignature      = 0x00004550
	ImageFileMachineAmd64 = 0x8664
	ImageFileMachineI386  = 0x014C
)

type RichHeaderEntry struct {
	ProdID   uint16
	BuildID  uint16
	UseCount uint32
}

type RichHeaderInfo struct {
	Size        int
	PtrToBuffer []byte
	Entries     int
}

type RichHeader struct {
	Entries []RichHeaderEntry
}

func FindRichHeader(dosHeader *ImageDosHeader, key []byte) int {
	baseAddr := (*[unsafe.Sizeof(ImageDosHeader{})]byte)(unsafe.Pointer(dosHeader))[:]
	offset := dosHeader.E_lfanew

	for i := 0; i < int(offset)-4; i++ {
		if baseAddr[i] == 'R' &&
			baseAddr[i+1] == 'i' &&
			baseAddr[i+2] == 'c' &&
			baseAddr[i+3] == 'h' {
			copy(key, baseAddr[i+4:i+8])
			return i
		}
	}
	return 0
}

func FindRichEntries(richHeaderPtr []byte, richHeaderSize int, key []byte, info *RichHeaderInfo) {
	for i := 0; i < richHeaderSize; i += 4 {
		for x := range 4 {
			richHeaderPtr[i+x] ^= key[x]
		}
	}

	info.Size = richHeaderSize
	info.PtrToBuffer = richHeaderPtr
	info.Entries = (richHeaderSize - 16) / 8
}

func SetRichEntries(richHeaderSize int, richHeaderPtr []byte, richHeader *RichHeader) {
	richHeader.Entries = make([]RichHeaderEntry, (richHeaderSize/8)-2)

	for i := 16; i < richHeaderSize; i += 8 {
		idx := (i / 8) - 2
		if idx >= len(richHeader.Entries) {
			break
		}
		prodID := binary.LittleEndian.Uint16(richHeaderPtr[i+2 : i+4])
		buildID := binary.LittleEndian.Uint16(richHeaderPtr[i : i+2])
		useCount := binary.LittleEndian.Uint32(richHeaderPtr[i+4 : i+8])
		richHeader.Entries[idx] = RichHeaderEntry{
			ProdID:   prodID,
			BuildID:  buildID,
			UseCount: useCount,
		}
	}
}

func DecryptRichHeader(key []byte, index int, dataPtr []byte) int {
	copy(key, dataPtr[index+4:index+8])

	indexPointer := index - 4
	richHeaderSize := 0

	for {
		tmpChar := make([]byte, 4)
		copy(tmpChar, dataPtr[indexPointer:indexPointer+4])

		for i := range 4 {
			tmpChar[i] ^= key[i]
		}

		indexPointer -= 4
		richHeaderSize += 4

		if tmpChar[1] == 0x61 && tmpChar[0] == 0x44 {
			break
		}
	}

	return richHeaderSize
}

func PeHexDump(ptr []byte, size int, secAddr int) {
	buf := make([]byte, 18)
	fmt.Printf("\n\n%x: |", secAddr)

	buf[0] = ' '
	buf[16] = ' '
	buf[17] = 0
	temp := 1

	for i := 1; i <= size; i++ {
		b := ptr[i-1]
		if temp < 17 {
			if b >= 0x20 && b < 0x7f {
				buf[temp] = b
			} else {
				buf[temp] = '.'
			}
		}
		fmt.Printf("%-3.02x", b)

		if i%16 == 0 {
			fmt.Printf(" %s", string(buf))
			if i+1 <= size {
				fmt.Printf("%x: ", secAddr+16)
			}
			temp = 0
		}
		if i%4 == 0 {
			fmt.Printf("| ")
		}
		temp++
	}

	if i := size % 16; i != 0 {
		buf[temp] = 0
		for j := 0; j < (16-i)%16; j++ {
			fmt.Printf("%-3.2c", ' ')
		}
		fmt.Printf(" %s", string(buf))
	}
}

type PeSectionInfo struct {
	Name             string
	VirtualSize      uint32
	VirtualAddress   uint32
	SizeOfRawData    uint32
	PointerToRawData uint32
	Characteristics  uint32
	HexDump          string
}

func PeShowSectionInformationAndDump(addressOfFile string, sectionToShow string, is32Bit bool) ([]PeSectionInfo, error) {
	var results []PeSectionInfo

	fileHandle, err := os.Open(addressOfFile)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer fileHandle.Close()

	_, _ = fileHandle.Stat()
	data, err := os.ReadFile(addressOfFile)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	dosHeader := (*ImageDosHeader)(unsafe.Pointer(&data[0]))

	key := make([]byte, 4)
	richHeaderOffset := FindRichHeader(dosHeader, key)

	var richFound bool
	var peFileRichHeader RichHeader
	var peFileRichHeaderInfo RichHeaderInfo

	if richHeaderOffset != 0 {
		dataPtr := data[:dosHeader.E_lfanew]
		richHeaderSize := DecryptRichHeader(key, richHeaderOffset, dataPtr)
		indexPointer := richHeaderOffset - richHeaderSize

		richHeaderPtr := make([]byte, richHeaderSize)
		copy(richHeaderPtr, data[indexPointer:indexPointer+richHeaderSize])

		FindRichEntries(richHeaderPtr, richHeaderSize, key, &peFileRichHeaderInfo)
		SetRichEntries(richHeaderSize, richHeaderPtr, &peFileRichHeader)
		richFound = true
	}

	fmt.Println("\nValid Dos Exe File")
	fmt.Println("------------------")
	fmt.Println("\nDumping DOS Header Info....")
	fmt.Println("---------------------------")
	magicStr := "MZ"
	if dosHeader.E_magic != ImageDosSignature {
		magicStr = "-"
	}
	fmt.Printf("%-36s%s\n", "Magic number :", magicStr)
	fmt.Printf("%-36s%#x\n", "Bytes on last page of file :", dosHeader.E_cblp)
	fmt.Printf("%-36s%#x\n", "Pages in file :", dosHeader.E_cp)
	fmt.Printf("%-36s%#x\n", "Relocation :", dosHeader.E_crlc)
	fmt.Printf("%-36s%#x\n", "Size of header in paragraphs :", dosHeader.E_cparhdr)
	fmt.Printf("%-36s%#x\n", "Minimum extra paragraphs needed :", dosHeader.E_minalloc)
	fmt.Printf("%-36s%#x\n", "Maximum extra paragraphs needed :", dosHeader.E_maxalloc)
	fmt.Printf("%-36s%#x\n", "Initial (relative) SS value :", dosHeader.E_ss)
	fmt.Printf("%-36s%#x\n", "Initial SP value :", dosHeader.E_sp)
	fmt.Printf("%-36s%#x\n", "Checksum :", dosHeader.E_csum)
	fmt.Printf("%-36s%#x\n", "Initial IP value :", dosHeader.E_ip)
	fmt.Printf("%-36s%#x\n", "Initial (relative) CS value :", dosHeader.E_cs)
	fmt.Printf("%-36s%#x\n", "File address of relocation table :", dosHeader.E_lfarlc)
	fmt.Printf("%-36s%#x\n", "Overlay number :", dosHeader.E_ovno)
	fmt.Printf("%-36s%#x\n", "OEM identifier :", dosHeader.E_oemid)
	fmt.Printf("%-36s%#x\n", "OEM information(e_oemid specific):", dosHeader.E_oeminfo)
	fmt.Printf("%-36s%#x\n", "RVA address of PE header :", dosHeader.E_lfanew)
	fmt.Println("================================================================")

	if richFound {
		fmt.Println("\n===============================================================================")
		fmt.Println("                              RICH HEADER                                     ")
		fmt.Println("===============================================================================")
		fmt.Printf("Entries: %d\n\n", peFileRichHeaderInfo.Entries)
		fmt.Printf("%-10s %-10s %-10s\n", "Build ID", "Prod ID", "Use Count")
		fmt.Println("---------------------------------------")

		for i := 0; i < peFileRichHeaderInfo.Entries; i++ {
			fmt.Printf("0x%08X 0x%08X %10d\n",
				peFileRichHeader.Entries[i].BuildID,
				peFileRichHeader.Entries[i].ProdID,
				peFileRichHeader.Entries[i].UseCount)
		}
		fmt.Println("==============Rich Header End ==================")
	} else {
		fmt.Println("=========== Rich Header Not Found ===========")
	}

	ntHeaderOffset := dosHeader.E_lfanew

	var fileHeader ImageFileHeader
	var numberOfSections uint32

	if is32Bit {
		_ = (*ImageOptionalHeader32)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		fileHeader = *(*ImageFileHeader)(unsafe.Pointer(&data[ntHeaderOffset+4]))
		numberOfSections = uint32(fileHeader.NumberOfSections)

		if binary.LittleEndian.Uint32(data[ntHeaderOffset:ntHeaderOffset+4]) != ImageNtSignature {
			return nil, fmt.Errorf("not a valid PE32 file")
		}
		fmt.Println("\nValid PE32 file \n-------------")
	} else {
		_ = (*ImageOptionalHeader64)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		fileHeader = *(*ImageFileHeader)(unsafe.Pointer(&data[ntHeaderOffset+4]))
		numberOfSections = uint32(fileHeader.NumberOfSections)

		if binary.LittleEndian.Uint32(data[ntHeaderOffset:ntHeaderOffset+4]) != ImageNtSignature {
			return nil, fmt.Errorf("not a valid PE64 file")
		}
		fmt.Println("\nValid PE64 file \n-------------")
	}

	fmt.Println("\nDumping COFF/PE Header Info....")
	fmt.Println("--------------------------------")
	fmt.Printf("%-36s%s\n", "Signature :", "PE")

	fmt.Printf("%-36s", "Machine Architecture :")
	switch fileHeader.Machine {
	case 0x0:
		fmt.Println("All ")
	case 0x14D:
		fmt.Println("Intel i860")
	case 0x14C:
		fmt.Println("Intel i386, i486, i586")
	case 0x200:
		fmt.Println("Intel Itanium processor")
	case ImageFileMachineAmd64:
		fmt.Println("AMD x64")
	case 0x162:
		fmt.Println("MIPS R3000")
	case 0x166:
		fmt.Println("MIPS R4000")
	case 0x183:
		fmt.Println("DEC Alpha AXP")
	default:
		fmt.Println("Not Found")
	}

	fmt.Printf("%-36s", "Characteristics : ")
	if fileHeader.Characteristics&0x0002 == 0x0002 {
		fmt.Print("Executable Image, ")
	}
	if fileHeader.Characteristics&0x0020 == 0x0020 {
		fmt.Print("Application can address > 2GB, ")
	}
	if fileHeader.Characteristics&0x1000 == 0x1000 {
		fmt.Print("System file (Kernel Mode Driver), ")
	}
	if fileHeader.Characteristics&0x2000 == 0x2000 {
		fmt.Print("Dll file, ")
	}
	if fileHeader.Characteristics&0x4000 == 0x4000 {
		fmt.Print("Application runs only in Uniprocessor, ")
	}
	fmt.Println()

	fmt.Printf("%-36s%d\n", "No.sections(size) :", fileHeader.NumberOfSections)
	fmt.Printf("%-36s%d\n", "No.entries in symbol table :", fileHeader.NumberOfSymbols)
	fmt.Printf("%-36s%d\n", "Size of optional header :", fileHeader.SizeOfOptionalHeader)

	fmt.Println("\nDumping PE Optional Header Info....")
	fmt.Println("-----------------------------------")

	sectionHeaderOffset := int(ntHeaderOffset) + 24 + int(fileHeader.SizeOfOptionalHeader)

	for i := uint32(0); i < numberOfSections; i++ {
		secOffset := sectionHeaderOffset + int(i)*int(unsafe.Sizeof(ImageSectionHeader{}))
		secHeader := (*ImageSectionHeader)(unsafe.Pointer(&data[secOffset]))

		var secName strings.Builder
		for _, b := range secHeader.Name {
			if b != 0 {
				secName.WriteString(string(b))
			}
		}

		results = append(results, PeSectionInfo{
			Name:             secName.String(),
			VirtualSize:      secHeader.VirtualSize,
			VirtualAddress:   secHeader.VirtualAddress,
			SizeOfRawData:    secHeader.SizeOfRawData,
			PointerToRawData: secHeader.PointerToRawData,
			Characteristics:  secHeader.Characteristics,
		})

		if sectionToShow == "" || sectionToShow == secName.String() {
			fmt.Printf("\n--- Section [%s] ---\n", secName.String())
			fmt.Printf("Virtual Size      : %#x\n", secHeader.VirtualSize)
			fmt.Printf("Virtual Address   : %#x\n", secHeader.VirtualAddress)
			fmt.Printf("Size Of Raw Data  : %#x\n", secHeader.SizeOfRawData)
			fmt.Printf("Pointer To Raw Data: %#x\n", secHeader.PointerToRawData)
			fmt.Printf("Characteristics   : %#x\n", secHeader.Characteristics)

			rawDataStart := int(secHeader.PointerToRawData)
			rawDataEnd := rawDataStart + int(secHeader.SizeOfRawData)
			if rawDataEnd <= len(data) {
				PeHexDump(data[rawDataStart:rawDataEnd], int(secHeader.SizeOfRawData), int(secHeader.VirtualAddress))
			}
		}
	}

	return results, nil
}

func ParsePeHeaders(filePath string, is32Bit bool) (*ImageDosHeader, *ImageFileHeader, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, nil, err
	}

	dosHeader := (*ImageDosHeader)(unsafe.Pointer(&data[0]))
	if dosHeader.E_magic != ImageDosSignature {
		return nil, nil, fmt.Errorf("invalid DOS signature")
	}

	ntHeaderOffset := dosHeader.E_lfanew
	if binary.LittleEndian.Uint32(data[ntHeaderOffset:ntHeaderOffset+4]) != ImageNtSignature {
		return nil, nil, fmt.Errorf("invalid PE signature")
	}

	fileHeader := (*ImageFileHeader)(unsafe.Pointer(&data[ntHeaderOffset+4]))
	return dosHeader, fileHeader, nil
}

func GetEntryPoint(filePath string, is32Bit bool) (uint64, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	dosHeader := (*ImageDosHeader)(unsafe.Pointer(&data[0]))
	ntHeaderOffset := dosHeader.E_lfanew

	if is32Bit {
		optHeader := (*ImageOptionalHeader32)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		return uint64(optHeader.AddressOfEntryPoint), nil
	} else {
		optHeader := (*ImageOptionalHeader64)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		return uint64(optHeader.AddressOfEntryPoint), nil
	}
}

func GetImageBase(filePath string, is32Bit bool) (uint64, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return 0, err
	}

	dosHeader := (*ImageDosHeader)(unsafe.Pointer(&data[0]))
	ntHeaderOffset := dosHeader.E_lfanew

	if is32Bit {
		optHeader := (*ImageOptionalHeader32)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		return uint64(optHeader.ImageBase), nil
	} else {
		optHeader := (*ImageOptionalHeader64)(unsafe.Pointer(&data[ntHeaderOffset+24]))
		return optHeader.ImageBase, nil
	}
}

func IsPeFileValid(filePath string) (bool, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return false, err
	}

	if len(data) < int(unsafe.Sizeof(ImageDosHeader{})) {
		return false, fmt.Errorf("file too small")
	}

	dosHeader := (*ImageDosHeader)(unsafe.Pointer(&data[0]))
	if dosHeader.E_magic != ImageDosSignature {
		return false, nil
	}

	ntHeaderOffset := dosHeader.E_lfanew
	if int(ntHeaderOffset) >= len(data)-4 {
		return false, fmt.Errorf("NT header offset out of bounds")
	}

	if binary.LittleEndian.Uint32(data[ntHeaderOffset:ntHeaderOffset+4]) != ImageNtSignature {
		return false, nil
	}

	return true, nil
}
