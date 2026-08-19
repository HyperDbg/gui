// Package misc — pt_helper.go
//
// Provides the user-mode side of Intel Processor Trace (PT) decoding. The
// actual decoder is libipt (loaded via syscall in debugger/misc/pt/ipt.go,
// Phase C.4.5); this file owns the higher-level helpers that capture the
// .text section of a target process and feed it to the decoder as the image
// callback.
//
// Mirrors libhyperdbg/code/debugger/misc/pt-helper.cpp.
package misc

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// ImageSymbolContext mirrors the C IMAGE_SYMBOL_CONTEXT: it caches the .text
// section bytes of a process so the libipt image-read callback can satisfy
// PT decode queries without re-reading memory on every callback.
type ImageSymbolContext struct {
	CodeBase uint64 // base address of the captured .text section
	CodeSize uint64 // size in bytes
	Code     []byte // the actual bytes
}

// PtHelperReadImage is the image-read callback invoked by libipt during PT
// decoding. It copies Count bytes starting at virtual address Ip into Buffer.
// Returns the number of bytes copied, or -1 when Ip is outside the captured
// .text range (mapped to libipt's -pte_nomap on the C side).
func PtHelperReadImage(buf []byte, ip uint64, ctx *ImageSymbolContext) int {
	if ctx == nil || ctx.Code == nil || ip < ctx.CodeBase || ip >= ctx.CodeBase+ctx.CodeSize {
		return -1
	}
	available := ctx.CodeBase + ctx.CodeSize - ip
	count := uint64(len(buf))
	if count > available {
		count = available
	}
	copy(buf[:count], ctx.Code[ip-ctx.CodeBase:])
	return int(count)
}

// PtHelperCaptureImage reads the .text section of process `proc` (a Win32
// process handle) into ctx. The section base/end are returned via textStart
// and textEnd on success.
//
// Mirrors the C PtHelperCaptureImage: walks the PEB→Ldr→InMemoryOrderModuleList
// to find the main image, then locates its .text section via the section
// headers.
func PtHelperCaptureImage(proc windows.Handle) (textStart, textEnd uint64, ctx *ImageSymbolContext, err error) {
	// 1. Get the PEB address via NtQueryInformationProcess(ProcessBasicInformation).
	ntdll := windows.NewLazyDLL("ntdll.dll")
	procQip := ntdll.NewProc("NtQueryInformationProcess")

	type processBasicInformation struct {
		ExitStatus                   uintptr
		PebBaseAddress               uintptr
		AffinityMask                 uintptr
		BasePriority                 uintptr
		UniqueProcessId              uintptr
		InheritedFromUniqueProcessId uintptr
	}
	var pbi processBasicInformation
	var retLen uint32
	r1, _, e := procQip.Call(
		uintptr(proc),
		uintptr(0), // ProcessBasicInformation
		uintptr(unsafe.Pointer(&pbi)),
		unsafe.Sizeof(pbi),
		uintptr(unsafe.Pointer(&retLen)),
	)
	if r1 != 0 {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: NtQueryInformationProcess failed: %v (r1=%v)", e, r1)
	}
	if pbi.PebBaseAddress == 0 {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: PEB is null")
	}

	// 2. Read the ImageBaseAddress from the PEB. The PEB layout on x64 has
	// ImageBaseAddress at offset 0x10; on x86 at offset 0x08. We assume x64
	// (HyperDbg targets 64-bit Windows for VMX).
	var imageBase uint64
	imageBaseOffset := uintptr(0x10) // x64 PEB.ImageBaseAddress
	var nRead uintptr
	err = windows.ReadProcessMemory(proc, pbi.PebBaseAddress+imageBaseOffset, (*byte)(unsafe.Pointer(&imageBase)), unsafe.Sizeof(imageBase), &nRead)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: ReadProcessMemory(PEB.ImageBaseAddress) failed: %w", err)
	}
	if nRead != unsafe.Sizeof(imageBase) {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: short read for ImageBaseAddress (%d/%d)", nRead, unsafe.Sizeof(imageBase))
	}

	// 3. Read the PE headers from the remote process to find .text.
	const hdrBufSize = 0x1000 // one page is enough for DOS+NT+section headers
	hdrBuf := make([]byte, hdrBufSize)
	nRead = 0
	err = windows.ReadProcessMemory(proc, uintptr(imageBase), &hdrBuf[0], hdrBufSize, &nRead)
	if err != nil {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: ReadProcessMemory(PE headers) failed: %w", err)
	}
	n := int(nRead)
	if n < 0x400 {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: short read for PE headers (%d)", n)
	}

	// Parse DOS header to find the PE header offset.
	e_lfanew := uint32(hdrBuf[0x3C]) | uint32(hdrBuf[0x3D])<<8 | uint32(hdrBuf[0x3E])<<16 | uint32(hdrBuf[0x3F])<<24
	if e_lfanew+0x108 > uint32(n) {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: PE header out of buffer (e_lfanew=%d)", e_lfanew)
	}
	// Verify PE\0\0 signature.
	if string(hdrBuf[e_lfanew:e_lfanew+4]) != "PE\x00\x00" {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: bad PE signature")
	}
	// COFF header: 20 bytes; NumberOfSections at offset e_lfanew+6, SizeOfOptionalHeader at e_lfanew+20.
	numSections := uint16(hdrBuf[e_lfanew+6]) | uint16(hdrBuf[e_lfanew+7])<<8
	sizeOfOptHdr := uint16(hdrBuf[e_lfanew+20]) | uint16(hdrBuf[e_lfanew+21])<<8
	// Section table starts at e_lfanew + 24 + sizeOfOptHdr.
	secTblOff := e_lfanew + 24 + uint32(sizeOfOptHdr)
	if secTblOff+uint32(numSections)*40 > uint32(n) {
		return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: section table out of buffer")
	}

	// Find the .text section (the first section whose name starts with ".text").
	for i := uint16(0); i < numSections; i++ {
		off := secTblOff + uint32(i)*40
		// IMAGE_SECTION_HEADER.Name is 8 bytes, NUL-padded. Match ".text".
		if string(hdrBuf[off:off+5]) != ".text" {
			continue
		}
		// IMAGE_SECTION_HEADER layout (40 bytes):
		//   Name[8]            +0
		//   VirtualSize        +8  (uint32)
		//   VirtualAddress     +12 (uint32)
		//   SizeOfRawData      +16
		//   PointerToRawData   +20
		//   ...
		virtualSize := uint32(hdrBuf[off+8]) | uint32(hdrBuf[off+9])<<8 | uint32(hdrBuf[off+10])<<16 | uint32(hdrBuf[off+11])<<24
		virtualAddr := uint32(hdrBuf[off+12]) | uint32(hdrBuf[off+13])<<8 | uint32(hdrBuf[off+14])<<16 | uint32(hdrBuf[off+15])<<24
		textStart = imageBase + uint64(virtualAddr)
		textEnd = textStart + uint64(virtualSize)
		// 4. Read the .text bytes from the remote process.
		ctx = &ImageSymbolContext{
			CodeBase: textStart,
			CodeSize: uint64(virtualSize),
			Code:     make([]byte, virtualSize),
		}
		read := 0
		var nRead uintptr
		for read < int(virtualSize) {
			want := int(virtualSize) - read
			if want > 0x10000 {
				want = 0x10000
			}
			nRead = 0
			err := windows.ReadProcessMemory(proc, uintptr(textStart)+uintptr(read), &ctx.Code[read], uintptr(want), &nRead)
			if err != nil {
				return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: ReadProcessMemory(.text @+%d) failed: %w", read, err)
			}
			if nRead == 0 {
				return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: zero-length read at +0x%x", read)
			}
			read += int(nRead)
		}
		return textStart, textEnd, ctx, nil
	}

	return 0, 0, nil, fmt.Errorf("PtHelperCaptureImage: no .text section found")
}
