package pe

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"log"
	"os"
	"sort"

	"github.com/ddkwork/golibrary/std/mylog"

	"github.com/edsrzf/mmap-go"
)

// TODO: detect endianess instead of forcing LittleEndian

// The representation of the PEFile with some helpful abstractions
type PEFile struct {
	Filename          string
	DosHeader         *DosHeader
	NTHeader          *NTHeader
	FileHeader        *FileHeader
	OptionalHeader    *OptionalHeader32
	OptionalHeader64  *OptionalHeader64
	StringTableOffset int
	StringTable       []byte
	SymbolTable       []*Symbol
	Sections          []*SectionHeader
	ImportDescriptors []*ImportDescriptor
	ExportDirectory   *ExportDirectory
	DebugDirectories  []*DebugDirectory

	// Private Fields
	reader    *bytes.Reader
	data      mmap.MMap
	dataLen   int
	headerEnd int
}

func PE(filename string) (pe *PEFile, err error) {
	pe = new(PEFile)
	pe.Filename = filename

	// Current file offset.
	offset := 0

	handle := mylog.Check2(os.Open(pe.Filename))

	defer func() {
		_ = handle.Close()
	}()

	pe.data = mylog.Check2(mmap.Map(handle, mmap.RDONLY, 0))

	pe.dataLen = len(pe.data)
	pe.reader = bytes.NewReader(pe.data)

	pe.DosHeader = NewDosHeader(offset)
	mylog.Check(pe.parseInterface(&pe.DosHeader.ImageDosHeader, offset, pe.DosHeader.size))

	if pe.DosHeader.E_magic == IMAGE_DOSZM_SIGNATURE {
		return nil, errors.New("probably a ZM Executable (not a PE file)")
	}

	if pe.DosHeader.E_magic != IMAGE_DOS_SIGNATURE {
		return nil, errors.New("DOS Header magic not found")
	}

	if int(pe.DosHeader.E_lfanew) > pe.dataLen {
		return nil, errors.New("invalid e_lfanew value, probably not a PE file")
	}

	offset = int(pe.DosHeader.E_lfanew)

	pe.NTHeader = NewNTHeader(offset)
	mylog.Check(pe.parseInterface(&pe.NTHeader.ImageNTHeader, offset, pe.NTHeader.size))

	if (pe.NTHeader.Signature & 0xFFFF) == IMAGE_NE_SIGNATURE {
		return nil, errors.New("invalid NT Headers signature (probably a NE file)")
	} else if (pe.NTHeader.Signature & 0xFFFF) == IMAGE_LE_SIGNATURE {
		return nil, errors.New("invalid NT Headers signature (probably a LE file)")
	} else if (pe.NTHeader.Signature & 0xFFFF) == IMAGE_LX_SIGNATURE {
		return nil, errors.New("invalid NT Headers signature (probably a LX file)")
	} else if (pe.NTHeader.Signature & 0xFFFF) == IMAGE_TE_SIGNATURE {
		return nil, errors.New("invalid NT Headers signature (probably a TE file)")
	} else if pe.NTHeader.Signature != IMAGE_NT_SIGNATURE {
		return nil, errors.New("invalid NT headers signature")
	}

	offset += pe.NTHeader.size

	pe.FileHeader = NewFileHeader(offset)
	mylog.Check(pe.parseInterface(&pe.FileHeader.ImageFileHeader, offset, pe.FileHeader.size))
	SetFlags(pe.FileHeader.flags, ImageCharacteristics, uint32(pe.FileHeader.Characteristics))

	offset += pe.FileHeader.size

	pe.OptionalHeader = NewOptionalHeader32(offset)
	mylog.Check(pe.parseInterface(&pe.OptionalHeader.ImageOptionalHeader32, offset, pe.OptionalHeader.size))
	SetFlags(pe.OptionalHeader.flags, DllCharacteristics, uint32(pe.OptionalHeader.DllCharacteristics))

	if pe.OptionalHeader.Magic == IMAGE_NT_OPTIONAL_HDR64_MAGIC {
		pe.OptionalHeader64 = NewOptionalHeader64(offset)
		mylog.Check(pe.parseInterface(&pe.OptionalHeader64.ImageOptionalHeader64, offset, pe.OptionalHeader64.size))

		if pe.OptionalHeader64.Magic != IMAGE_NT_OPTIONAL_HDR64_MAGIC {
			return nil, errors.New("no optional header found - invalid PE32 or PE32+ file")
		}
		SetFlags(pe.OptionalHeader64.flags, DllCharacteristics, uint32(pe.OptionalHeader64.DllCharacteristics))
	}

	// Read the symbol table.
	if pe.FileHeader.PointerToSymbolTable != 0 {
		numberOfSymbols := int(pe.FileHeader.NumberOfSymbols)
		pe.SymbolTable = make([]*Symbol, numberOfSymbols)
		for symbolOffset, i := int(pe.FileHeader.PointerToSymbolTable), 0; i < numberOfSymbols; i++ {
			pe.SymbolTable[i] = NewSymbol(symbolOffset)
			mylog.Check(pe.parseInterface(&pe.SymbolTable[i].ImageSymbol, symbolOffset, pe.SymbolTable[i].size))

			symbolOffset += IMAGE_SIZEOF_SYMBOL
		}

		// Read the symbol string table.
		pe.StringTableOffset = int(pe.FileHeader.PointerToSymbolTable) + (numberOfSymbols * IMAGE_SIZEOF_SYMBOL)
		var stringTableSize uint32
		mylog.Check(pe.parseInterface(&stringTableSize, pe.StringTableOffset, 4))
		pe.StringTable = pe.data[pe.StringTableOffset : pe.StringTableOffset+int(stringTableSize)]

		// Fill in the symbol table names.
		for _, symbol := range pe.SymbolTable {
			if binary.LittleEndian.Uint32(symbol.ShortName[:4]) == 0 {
				n := binary.LittleEndian.Uint32(symbol.ShortName[4:])
				longName := pe.StringTable[n:]
				for i := 0; i < len(longName); i++ {
					if longName[i] == 0 {
						longName = longName[:i]
						break
					}
				}
				symbol.Name = string(longName)
			} else {
				symbol.Name = string(symbol.ShortName[:])
			}
		}
	}

	// Section data
	// MAX_ASSUMED_VALID_NUMBER_OF_RVA_AND_SIZES := 0x100
	var numRvaAndSizes uint32

	msg := "Suspicious NumberOfRvaAndSizes in the Optional Header."
	msg += "Normal values are never larger than 16, the value is: 0x%x\n"

	var dataDir map[string]*DataDirectory

	sectionOffset := offset + int(pe.FileHeader.SizeOfOptionalHeader)

	if pe.OptionalHeader64 != nil {
		if pe.OptionalHeader64.NumberOfRvaAndSizes > IMAGE_NUMBEROF_DIRECTORY_ENTRIES {
			log.Printf(msg, pe.OptionalHeader64.NumberOfRvaAndSizes)
		}
		numRvaAndSizes = pe.OptionalHeader64.NumberOfRvaAndSizes
		dataDir = pe.OptionalHeader64.DataDirs
		offset += pe.OptionalHeader64.size
	} else {
		if pe.OptionalHeader.NumberOfRvaAndSizes > IMAGE_NUMBEROF_DIRECTORY_ENTRIES {
			log.Printf(msg, pe.OptionalHeader.NumberOfRvaAndSizes)
		}
		numRvaAndSizes = pe.OptionalHeader.NumberOfRvaAndSizes
		offset += pe.OptionalHeader.size
		dataDir = pe.OptionalHeader.DataDirs
	}

	for i := uint32(0); i < numRvaAndSizes&0x7fffffff; i++ {
		if pe.dataLen-offset == 0 {
			break
		}

		dirEntry := NewDataDirectory(offset)
		mylog.Check(pe.parseInterface(&dirEntry.ImageDataDirectory, offset, dirEntry.size))

		offset += dirEntry.size

		name, ok := DirectoryEntryTypes[i]
		if !ok {
			break
		}

		dirEntry.Name = name
		dataDir[dirEntry.Name] = dirEntry

		// TODO: add skipped check at L2038
	}

	offset = mylog.Check2(pe.parseSections(sectionOffset))

	pe.calculateHeaderEnd(offset)

	if pe.OptionalHeader.AddressOfEntryPoint != 0 {
		if pe.getSectionByRva(pe.OptionalHeader.AddressOfEntryPoint) != nil {
			epOffset := pe.getOffsetFromRva(pe.OptionalHeader.AddressOfEntryPoint)
			if epOffset > pe.dataLen {
				log.Printf("AddressOfEntryPoint lies outside the file: 0x%x", pe.OptionalHeader.AddressOfEntryPoint)
			}
		} else {
			log.Printf("AddressOfEntryPoint lies outside the section boundaries: 0x%x", pe.OptionalHeader.AddressOfEntryPoint)
		}
	}

	mylog.Check(pe.parseDataDirectories())

	//offset, err = pe.parseRichHeader()
	//if err != nil {
	//	return nil, err
	//}

	return pe, nil
}

func OBJ(filename string) (pe *PEFile, err error) {
	pe = new(PEFile)
	pe.Filename = filename

	offset := 0

	handle := mylog.Check2(os.Open(pe.Filename))

	defer func() {
		_ = handle.Close()
	}()

	pe.data = mylog.Check2(mmap.Map(handle, mmap.RDONLY, 0))

	pe.dataLen = len(pe.data)
	pe.reader = bytes.NewReader(pe.data)

	pe.FileHeader = NewFileHeader(offset)
	mylog.Check(pe.parseInterface(&pe.FileHeader.ImageFileHeader, offset, pe.FileHeader.size))

	// Size of the optional header, which is required for executable files but not for object files.
	// An object file should have a value of 0 here.
	if pe.FileHeader.SizeOfOptionalHeader != 0 {
		return nil, errors.New("invalid or corrupt object file - should not have an optional header size")
	}
	SetFlags(pe.FileHeader.flags, ImageCharacteristics, uint32(pe.FileHeader.Characteristics))

	// Read the symbol table.
	if pe.FileHeader.PointerToSymbolTable != 0 {
		numberOfSymbols := int(pe.FileHeader.NumberOfSymbols)
		pe.SymbolTable = make([]*Symbol, numberOfSymbols)
		for symbolOffset, i := int(pe.FileHeader.PointerToSymbolTable), 0; i < numberOfSymbols; i++ {
			pe.SymbolTable[i] = NewSymbol(symbolOffset)
			mylog.Check(pe.parseInterface(&pe.SymbolTable[i].ImageSymbol, symbolOffset, pe.SymbolTable[i].size))

			symbolOffset += IMAGE_SIZEOF_SYMBOL
		}

		// Read the symbol string table.
		pe.StringTableOffset = int(pe.FileHeader.PointerToSymbolTable) + (numberOfSymbols * IMAGE_SIZEOF_SYMBOL)
		var stringTableSize uint32
		mylog.Check(pe.parseInterface(&stringTableSize, pe.StringTableOffset, 4))
		pe.StringTable = pe.data[pe.StringTableOffset : pe.StringTableOffset+int(stringTableSize)]

		// Fill in the symbol table names.
		for _, symbol := range pe.SymbolTable {
			if binary.LittleEndian.Uint32(symbol.ShortName[:4]) == 0 {
				n := binary.LittleEndian.Uint32(symbol.ShortName[4:])
				longName := pe.StringTable[n:]
				for i := 0; i < len(longName); i++ {
					if longName[i] == 0 {
						longName = longName[:i]
						break
					}
				}
				symbol.Name = string(longName)
			} else {
				symbol.Name = string(symbol.ShortName[:])
			}
		}
	}

	offset += pe.FileHeader.size

	offset = mylog.Check2(pe.parseSections(offset))

	pe.calculateHeaderEnd(offset)

	return pe, nil
}

func (p *PEFile) GetRawData() []byte {
	return p.data
}

func (p *PEFile) GetRawDataSize() int {
	return p.dataLen
}

type ByVAddr []*SectionHeader

func (a ByVAddr) Len() int {
	return len(a)
}

func (a ByVAddr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByVAddr) Less(i, j int) bool {
	return a[i].VirtualAddress < a[j].VirtualAddress
}

func (p *PEFile) parseSections(offset int) (newOffset int, err error) {
	newOffset = offset
	for i := uint32(0); i < uint32(p.FileHeader.NumberOfSections); i++ {
		section := NewSectionHeader(newOffset)
		mylog.Check(p.parseInterface(&section.ImageSectionHeader, newOffset, section.size))

		// Set raw data.
		section.RawData = make([]byte, section.SizeOfRawData)
		copy(section.RawData, p.data[section.PointerToRawData:section.PointerToRawData+section.SizeOfRawData])

		// Get relocations.
		if section.PointerToRelocations != uint32(0) {
			numberOfRelocs := int(section.NumberOfRelocations)
			section.Relocations = make([]*Relocation, numberOfRelocs)
			for relocOffset, relocIdx := int(section.PointerToRelocations), 0; relocIdx < numberOfRelocs; relocIdx++ {
				section.Relocations[relocIdx] = NewRelocation(relocOffset)
				mylog.Check(p.parseInterface(&section.Relocations[relocIdx].ImageRelocation, relocOffset, section.Relocations[relocIdx].size))
				section.Relocations[relocIdx].Symbol = p.SymbolTable[section.Relocations[relocIdx].SymbolTableIndex]

				relocOffset += IMAGE_SIZEOF_RELOCATION
			}
		}

		// TODO: More checks and error handling here from parseSections
		// L2325-2376

		SetFlags(section.flags, SectionCharacteristics, section.Characteristics)

		// Suspecious check L2383 - L2395
		p.Sections = append(p.Sections, section)

		newOffset += section.size
	}

	// Sort the sections by their VirtualAddress and add a field to each of them
	// with the VirtualAddress of the next section. This will allow to check
	// for potentially overlapping sections in badly constructed PEs.
	sort.Sort(ByVAddr(p.Sections))
	for idx, section := range p.Sections {
		if idx == len(p.Sections)-1 {
			section.nextHeaderRva = 0
		} else {
			section.nextHeaderRva = p.Sections[idx+1].VirtualAddress
		}
	}

	return newOffset, nil
}

func (p *PEFile) parseInterface(iface any, offset, size int) (err error) {
	mylog.Check2(p.reader.Seek(int64(offset), io.SeekStart))

	mylog.Check(binary.Read(p.reader, binary.LittleEndian, iface))

	return nil
}

func (p *PEFile) parseDataDirectories() error {
	var dataDirs map[string]*DataDirectory

	funcMap := map[string]any{
		"IMAGE_DIRECTORY_ENTRY_IMPORT": p.parseImportDirectory,
		"IMAGE_DIRECTORY_ENTRY_EXPORT": p.parseExportDirectory,
		// TODO at a later time
		//"IMAGE_DIRECTORY_ENTRY_RESOURCE": p.parseResourcesDirectory,
		"IMAGE_DIRECTORY_ENTRY_DEBUG": p.parseDebugDirectory,
		//"IMAGE_DIRECTORY_ENTRY_BASERELOC": p.parseRelocationsDirectory,
		//"IMAGE_DIRECTORY_ENTRY_TLS": p.parseDirectoryTls,
		//"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG": p.parseDirectoryLoadConfig,
		//"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT": p.parseDelayImportDirectory,
		//"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT": p.parseDirectoryBoundImports,
	}

	if p.OptionalHeader64 != nil {
		dataDirs = p.OptionalHeader64.DataDirs
	} else {
		dataDirs = p.OptionalHeader.DataDirs
	}

	for name, dirEntry := range dataDirs {
		if dirEntry.VirtualAddress > 0 {
			parser, ok := funcMap[name]
			if !ok {
				continue
			}
			mylog.Check(parser.(func(uint32, uint32) error)(dirEntry.VirtualAddress, dirEntry.Size))

		}
	}

	return nil
}

func (p *PEFile) getSectionByRva(rva uint32) *SectionHeader {
	for _, section := range p.Sections {
		var size uint32
		adjustedPointer := p.adjustFileAlignment(section.PointerToRawData, p.getFileAlignment())
		if uint32(p.dataLen)-adjustedPointer < section.SizeOfRawData {
			size = section.Misc_VirtualSize_PhysicalAddress
		} else {
			size = MaxUInt32(section.SizeOfRawData, section.Misc_VirtualSize_PhysicalAddress)
		}

		vaddr := p.adjustSectionAlignment(section.VirtualAddress, p.getSectionAlignment(), p.getFileAlignment())

		if section.nextHeaderRva != 0 && section.nextHeaderRva > section.VirtualAddress && vaddr+size > section.nextHeaderRva {
			size = section.nextHeaderRva - vaddr
		}

		if vaddr <= rva && rva < (vaddr+size) {
			return section
		}
	}

	return nil
}

func (p *PEFile) getSectionByOffset(offset int) *SectionHeader {
	for _, section := range p.Sections {
		if section.PointerToRawData == 0 {
			continue
		}

		adjustedPointer := p.adjustFileAlignment(section.PointerToRawData, p.getFileAlignment())
		if int(adjustedPointer) <= offset && offset < int(adjustedPointer+section.SizeOfRawData) {
			return section
		}
	}

	return nil
}

func (p *PEFile) getRvaFromOffset(offset int) uint32 {
	section := p.getSectionByOffset(offset)
	minAddr := ^uint32(0)
	if section == nil {

		if len(p.Sections) == 0 {
			return uint32(offset)
		}

		for _, section := range p.Sections {
			vaddr := p.adjustSectionAlignment(section.VirtualAddress, p.getSectionAlignment(), p.getFileAlignment())
			if vaddr < minAddr {
				minAddr = vaddr
			}
		}

		// Assume that offset lies within the headers
		// The case illustrating this behavior can be found at:
		// http://corkami.blogspot.com/2010/01/hey-hey-hey-whats-in-your-head.html
		// where the import table is not contained by any section
		// hence the RVA needs to be resolved to a raw offset
		if offset < int(minAddr) {
			return uint32(offset)
		}

		log.Println("data at Offset cannot be fetched - corrupt header?")
		return ^uint32(0)
	}

	sectionAlignment := p.adjustSectionAlignment(section.VirtualAddress, p.getSectionAlignment(), p.getFileAlignment())
	fileAlignment := p.adjustFileAlignment(section.PointerToRawData, p.getFileAlignment())
	return uint32(offset - (int(fileAlignment) + int(sectionAlignment)))
}

func (p *PEFile) getOffsetFromRva(rva uint32) int {
	section := p.getSectionByRva(rva)
	if section == nil {
		if int(rva) < p.dataLen {
			return int(rva)
		}
		log.Println("data at RVA cannot be fetched - corrupt header?")
		return ^0
	}
	sectionAlignment := p.adjustSectionAlignment(section.VirtualAddress, p.getSectionAlignment(), p.getFileAlignment())
	fileAlignment := p.adjustFileAlignment(section.PointerToRawData, p.getFileAlignment())
	return int(rva - sectionAlignment + fileAlignment)
}

func (p *PEFile) getFileAlignment() uint32 {
	// For executable images, this must be a multiple of FileAlignment from the optional header.
	// For object files, the value should be aligned on a four-byte boundary for best performance.
	if p.OptionalHeader != nil {
		return p.OptionalHeader.FileAlignment
	}
	return uint32(4)
}

func (p *PEFile) getSectionAlignment() uint32 {
	if p.OptionalHeader != nil {
		return p.OptionalHeader.SectionAlignment
	}
	return uint32(4)
}

// According to http://corkami.blogspot.com/2010/01/parce-que-la-planche-aura-brule.html
// if PointerToRawData is less that 512 it's rounded to zero. Loading the test file
// in a debugger it's easy to verify that the PointerToRawData value of 1 is rounded
// to zero. Hence we reproduce the behavior
//
// According to the document:
// [ Microsoft Portable Executable and Common Object File Format Specification ]
// "The alignment factor (in bytes) that is used to align the raw data of sections in
//
//	the image file. The value should be a power of 2 between 512 and 64 K, inclusive.
//	The default is 512. If the SectionAlignment is less than the architecture's page
//	size, then FileAlignment must match SectionAlignment."
//
// The following is a hard-coded constant if the Windows loader
func (p *PEFile) adjustFileAlignment(pointer, fileAlignment uint32) uint32 {
	if fileAlignment > IMAGE_FILE_ALIGNMENT_HARDCODED_VALUE {
		// If it's not a power of two, report it:
		if !PowerOfTwo(fileAlignment) {
			log.Printf("if FileAlignment > 512 it should be a power of 2: %x", fileAlignment)
		}
	}

	if fileAlignment < IMAGE_FILE_ALIGNMENT_HARDCODED_VALUE {
		return pointer
	}

	return (pointer / IMAGE_FILE_ALIGNMENT_HARDCODED_VALUE) * IMAGE_FILE_ALIGNMENT_HARDCODED_VALUE
}

// According to the document:
// [ Microsoft Portable Executable and Common Object File Format Specification ]
// "The alignment (in bytes) of sections when they are loaded into memory. It must be
//
//	greater than or equal to FileAlignment. The default is the page size for the
//	architecture."
func (p *PEFile) adjustSectionAlignment(pointer, sectionAlignment, fileAlignment uint32) uint32 {
	if fileAlignment < IMAGE_FILE_ALIGNMENT_HARDCODED_VALUE {
		if fileAlignment != sectionAlignment {
			log.Printf("if FileAlignment(%x) < 512 it should equal SectionAlignment(%x)", fileAlignment, sectionAlignment)
		}
	}

	if int(sectionAlignment) < os.Getpagesize() { // page size
		sectionAlignment = fileAlignment
	} else if sectionAlignment < 0x80 {
		// 512 is the minimum valid FileAlignment according to the documentation
		// although ntoskrnl.exe has an alignment of 0x80 in some Windows versions
		sectionAlignment = 0x80
	}

	if sectionAlignment != 0 && (pointer%sectionAlignment) != 0 {
		return sectionAlignment * (pointer / sectionAlignment)
	}

	return pointer
}

func (p *PEFile) getDataBounds(rva, length uint32) (start, size int) {
	var offset, end int

	section := p.getSectionByRva(rva)

	if length > 0 {
		end = int(rva + length)
	} else {
		end = p.dataLen
	}

	if section == nil {
		if int(rva) < p.headerEnd {
			end = MinInt(end, p.headerEnd)
		}
		// Before we give up we check whether the file might
		// contain the data anyway. There are cases of PE files
		// without sections that rely on windows loading the first
		// 8291 bytes into memory and assume the data will be there
		// A functional file with these characteristics is:
		// MD5: 0008892cdfbc3bda5ce047c565e52295
		// SHA-1: c7116b9ff950f86af256defb95b5d4859d4752a9
		if int(rva) < p.dataLen {
			return int(rva), end
		}

		return ^0, ^0
	}
	pointer := p.adjustFileAlignment(section.PointerToRawData, p.getFileAlignment())
	vaddr := p.adjustSectionAlignment(section.VirtualAddress, p.getSectionAlignment(), p.getFileAlignment())

	if rva == 0 {
		offset = int(pointer)
	} else {
		offset = int((rva - vaddr) + pointer)
	}

	if length != 0 {
		end = offset + int(length)
	} else {
		end = offset + int(section.SizeOfRawData)
	}

	if end > int(pointer+section.SizeOfRawData) {
		end = int(section.PointerToRawData) + int(section.SizeOfRawData)
	}

	return offset, end
}

// Get an ASCII string from within the data at an RVA considering section
func (p *PEFile) getStringAtRva(rva uint32) []byte {
	start, _ := p.getDataBounds(rva, 0)
	return p.getStringFromData(start)
}

// Get an ASCII string from within the data.
func (p *PEFile) getStringFromData(offset int) []byte {
	if offset > p.dataLen {
		return []byte{}
	}

	end := offset
	for end < p.dataLen {
		if p.data[end] == 0 {
			break
		}
		end += 1
	}

	return p.data[offset:end]
}

// OC Patch:
// There could be a problem if there are no raw data sections
// greater than 0
// fc91013eb72529da005110a3403541b6 example
// Should this throw an exception in the minimum header offset
// can't be found?
func (p *PEFile) calculateHeaderEnd(offset int) {
	var rawDataPointers []uint32

	for _, section := range p.Sections {
		prd := section.PointerToRawData
		if prd > uint32(0) {
			rawDataPointers = append(rawDataPointers, p.adjustFileAlignment(prd, p.getFileAlignment()))
		}
	}

	minSectionOffset := 0
	if len(rawDataPointers) > 0 {
		minSectionOffset = int(rawDataPointers[0])
		for _, pointer := range rawDataPointers {
			if int(pointer) < minSectionOffset {
				minSectionOffset = int(pointer)
			}
		}
	}

	if minSectionOffset == 0 || minSectionOffset < offset {
		p.headerEnd = offset
	} else {
		p.headerEnd = minSectionOffset
	}
}
