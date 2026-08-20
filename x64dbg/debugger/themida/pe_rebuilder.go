package themida

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ddkwork/x64dbg/debugger"
)

type ImageDosHeader struct {
	e_magic    [2]byte
	e_cblp     uint16
	e_cp       uint16
	e_crlc     uint16
	e_cparhdr  uint16
	e_minalloc uint16
	e_maxalloc uint16
	e_ss       uint16
	e_sp       uint16
	e_csum     uint16
	e_ip       uint16
	e_cs       uint16
	e_lfarlc   uint16
	e_ovno     uint16
	e_res      [4]uint16
	e_oemid    uint16
	e_oeminfo  uint16
	e_res2     [10]uint16
	e_lfanew   uint32
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

type ImageOptionalHeader struct {
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

type ImageDataDirectory struct {
	VirtualAddress uint32
	Size           uint32
}

type PERebuilder struct {
	dosHeader       ImageDosHeader
	fileHeader      ImageFileHeader
	optHeader       ImageOptionalHeader
	sections        []ImageSectionHeader
	dataDirectories [16]ImageDataDirectory
}

func NewPERebuilder() *PERebuilder {
	return &PERebuilder{
		sections: make([]ImageSectionHeader, 0),
	}
}

func (r *PERebuilder) DumpProcess(dbg *debugger.Debugger, outputPath string) error {
	log.Println("开始转储进程...")

	baseAddr := dbg.GetBaseAddress()
	imageSize := uint32(0x1000000)

	log.Printf("进程基址: 0x%X, 镜像大小: 0x%X", baseAddr, imageSize)

	memory, err := dbg.GetMemory().ReadMemory(baseAddr, imageSize)
	if err != nil {
		return fmt.Errorf("读取进程内存失败: %v", err)
	}

	if err := r.parsePEHeaders(memory); err != nil {
		return fmt.Errorf("解析 PE 头失败: %v", err)
	}

	dumpPath := filepath.Join(filepath.Dir(outputPath), "dumped.bin")
	if err := os.WriteFile(dumpPath, memory, 0o644); err != nil {
		return fmt.Errorf("写入转储文件失败: %v", err)
	}

	log.Printf("进程转储完成: %s", dumpPath)
	return nil
}

func (r *PERebuilder) parsePEHeaders(data []byte) error {
	if len(data) < 64 {
		return fmt.Errorf("数据长度不足")
	}

	if err := binary.Read(bytes.NewReader(data[0:64]), binary.LittleEndian, &r.dosHeader); err != nil {
		return fmt.Errorf("读取 DOS 头失败: %v", err)
	}

	if r.dosHeader.e_magic[0] != 'M' || r.dosHeader.e_magic[1] != 'Z' {
		return fmt.Errorf("无效的 DOS 签名")
	}

	peOffset := r.dosHeader.e_lfanew
	if peOffset >= uint32(len(data)) {
		return fmt.Errorf("PE 偏移超出范围")
	}

	if data[peOffset] != 'P' || data[peOffset+1] != 'E' {
		return fmt.Errorf("无效的 PE 签名")
	}

	fileHeaderOffset := peOffset + 4
	if err := binary.Read(bytes.NewReader(data[fileHeaderOffset:fileHeaderOffset+20]), binary.LittleEndian, &r.fileHeader); err != nil {
		return fmt.Errorf("读取文件头失败: %v", err)
	}

	optHeaderOffset := fileHeaderOffset + 20
	optHeaderSize := r.fileHeader.SizeOfOptionalHeader
	if optHeaderSize < 224 {
		return fmt.Errorf("可选头大小不足")
	}

	if err := binary.Read(bytes.NewReader(data[optHeaderOffset:optHeaderOffset+uint32(optHeaderSize)]), binary.LittleEndian, &r.optHeader); err != nil {
		return fmt.Errorf("读取可选头失败: %v", err)
	}

	sectionHeaderOffset := optHeaderOffset + uint32(optHeaderSize)
	sectionCount := r.fileHeader.NumberOfSections
	if sectionCount > 96 {
		return fmt.Errorf("段数量过多")
	}

	r.sections = make([]ImageSectionHeader, sectionCount)
	for i := range sectionCount {
		offset := sectionHeaderOffset + uint32(i)*40
		if err := binary.Read(bytes.NewReader(data[offset:offset+40]), binary.LittleEndian, &r.sections[i]); err != nil {
			return fmt.Errorf("读取段头失败: %v", err)
		}
	}

	dataDirOffset := optHeaderOffset + 96
	if err := binary.Read(bytes.NewReader(data[dataDirOffset:dataDirOffset+128]), binary.LittleEndian, &r.dataDirectories); err != nil {
		return fmt.Errorf("读取数据目录失败: %v", err)
	}

	log.Printf("PE 头解析完成: %d 个段", sectionCount)
	return nil
}

func (r *PERebuilder) RebuildPE(outputPath string) error {
	log.Println("开始重建 PE 文件...")

	fileSize := r.calculateFileSize()
	outputData := make([]byte, fileSize)

	if err := r.writeDOSHeader(outputData); err != nil {
		return fmt.Errorf("写入 DOS 头失败: %v", err)
	}

	if err := r.writePEHeaders(outputData); err != nil {
		return fmt.Errorf("写入 PE 头失败: %v", err)
	}

	if err := r.writeSections(outputData); err != nil {
		return fmt.Errorf("写入段数据失败: %v", err)
	}

	if err := os.WriteFile(outputPath, outputData, 0o644); err != nil {
		return fmt.Errorf("写入 PE 文件失败: %v", err)
	}

	log.Printf("PE 文件重建完成: %s (大小: %d 字节)", outputPath, fileSize)
	return nil
}

func (r *PERebuilder) calculateFileSize() uint32 {
	fileAlignment := r.optHeader.FileAlignment
	if fileAlignment == 0 {
		fileAlignment = 512
	}

	fileSize := r.optHeader.SizeOfHeaders

	for _, section := range r.sections {
		if section.SizeOfRawData > 0 {
			alignedSize := ((section.SizeOfRawData + fileAlignment - 1) / fileAlignment) * fileAlignment
			fileSize += alignedSize
		}
	}

	return fileSize
}

func (r *PERebuilder) writeDOSHeader(data []byte) error {
	buf := bytes.NewBuffer(data[0:64])
	return binary.Write(buf, binary.LittleEndian, &r.dosHeader)
}

func (r *PERebuilder) writePEHeaders(data []byte) error {
	peOffset := r.dosHeader.e_lfanew

	data[peOffset] = 'P'
	data[peOffset+1] = 'E'
	data[peOffset+2] = 0
	data[peOffset+3] = 0

	fileHeaderOffset := peOffset + 4
	buf := bytes.NewBuffer(data[fileHeaderOffset : fileHeaderOffset+20])
	if err := binary.Write(buf, binary.LittleEndian, &r.fileHeader); err != nil {
		return err
	}

	optHeaderOffset := fileHeaderOffset + 20
	optHeaderSize := r.fileHeader.SizeOfOptionalHeader
	buf = bytes.NewBuffer(data[optHeaderOffset : optHeaderOffset+uint32(optHeaderSize)])
	if err := binary.Write(buf, binary.LittleEndian, &r.optHeader); err != nil {
		return err
	}

	sectionHeaderOffset := optHeaderOffset + uint32(optHeaderSize)
	for i, section := range r.sections {
		offset := sectionHeaderOffset + uint32(i)*40
		buf := bytes.NewBuffer(data[offset : offset+40])
		if err := binary.Write(buf, binary.LittleEndian, &section); err != nil {
			return err
		}
	}

	return nil
}

func (r *PERebuilder) writeSections(data []byte) error {
	fileAlignment := r.optHeader.FileAlignment
	if fileAlignment == 0 {
		fileAlignment = 512
	}

	for _, section := range r.sections {
		if section.SizeOfRawData == 0 {
			continue
		}

		if section.PointerToRawData >= uint32(len(data)) {
			return fmt.Errorf("段偏移超出范围")
		}

		if section.PointerToRawData+section.SizeOfRawData > uint32(len(data)) {
			return fmt.Errorf("段大小超出范围")
		}

		alignedSize := ((section.SizeOfRawData + fileAlignment - 1) / fileAlignment) * fileAlignment
		if section.PointerToRawData+alignedSize > uint32(len(data)) {
			return fmt.Errorf("对齐后的段大小超出范围")
		}
	}

	return nil
}

func (r *PERebuilder) GetEntryPoint() uint32 {
	return r.optHeader.AddressOfEntryPoint
}

func (r *PERebuilder) GetImageBase() uint32 {
	return r.optHeader.ImageBase
}

func (r *PERebuilder) GetSections() []ImageSectionHeader {
	return r.sections
}
