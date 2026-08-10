package ddk

import (
	"encoding/binary"
	"fmt"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/arch/x86/x86asm"
)

type KernelMemory struct {
	rt *RTCore64
}

func NewKernelMemory(rt *RTCore64) *KernelMemory {
	return &KernelMemory{rt: rt}
}

func (k *KernelMemory) RTCore64() *RTCore64 {
	return k.rt
}

func (k *KernelMemory) ReadMemory(addr uint64, buf []byte) {
	k.rt.ReadMemory(addr, buf)
}

func (k *KernelMemory) ReadUint64(addr uint64) uint64 {
	return k.rt.ReadUint64(addr)
}

func (k *KernelMemory) ReadUint32(addr uint64) uint32 {
	return k.rt.ReadUint32(addr)
}

func (k *KernelMemory) ReadUint16(addr uint64) uint16 {
	return k.rt.ReadUint16(addr)
}

func (k *KernelMemory) WriteUint32(addr uint64, val uint32) {
	k.rt.WriteUint32(addr, val)
}

type PESection struct {
	Name            string
	VirtualAddr     uint32
	VirtualSize     uint32
	RawAddr         uint32
	RawSize         uint32
	Characteristics uint32
}

func (k *KernelMemory) ParsePESections(moduleBase uint64) []PESection {
	eLfanew := k.ReadUint32(moduleBase + 0x3C)

	ntHdr := moduleBase + uint64(eLfanew)
	sig := k.ReadUint32(ntHdr)
	if sig != 0x4550 {
		mylog.Check(fmt.Errorf("invalid PE signature at moduleBase+0x%X", eLfanew))
	}

	fileHdrOffset := ntHdr + 4
	numSections := k.ReadUint16(fileHdrOffset + 0x06)

	optHdrSize := k.ReadUint16(fileHdrOffset + 0x14)

	firstSec := fileHdrOffset + 20 + uint64(optHdrSize)

	var sections []PESection
	for i := range numSections {
		secOff := firstSec + uint64(i)*40
		raw := make([]byte, 40)
		k.ReadMemory(moduleBase+secOff, raw)
		name := byteslice.ToString(raw[:8])
		sections = append(sections, PESection{
			Name:            name,
			VirtualAddr:     binary.LittleEndian.Uint32(raw[12:]),
			VirtualSize:     binary.LittleEndian.Uint32(raw[8:]),
			RawAddr:         binary.LittleEndian.Uint32(raw[20:]),
			RawSize:         binary.LittleEndian.Uint32(raw[16:]),
			Characteristics: binary.LittleEndian.Uint32(raw[36:]),
		})
	}
	return sections
}

func (k *KernelMemory) FindTextSection(moduleBase uint64) (va uint64, size uint64, err error) {
	sections := k.ParsePESections(moduleBase)

	for _, sec := range sections {
		if sec.Name == ".text" && sec.VirtualSize > 0 {
			textVA := moduleBase + uint64(sec.VirtualAddr)
			mylog.Info(".text section found", "va", fmt.Sprintf("0x%X", textVA), "size", fmt.Sprintf("0x%X", sec.VirtualSize))
			return textVA, uint64(sec.VirtualSize), nil
		}
	}
	return 0, 0, fmt.Errorf(".text section not found")
}

func (k *KernelMemory) ReadCode(addr uint64, size int) []byte {
	buf := make([]byte, size)
	k.ReadMemory(addr, buf)
	return buf
}

func (k *KernelMemory) Disassemble(addr uint64, size int) []x86asm.Inst {
	code := k.ReadCode(addr, size)
	return Disassemble(code)
}

func (k *KernelMemory) DisassembleToString(addr uint64, size int) string {
	code := k.ReadCode(addr, size)
	return DisassembleToString(code, addr)
}
