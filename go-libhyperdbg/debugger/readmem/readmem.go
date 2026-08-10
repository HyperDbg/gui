// Package readmem implements HyperDbgReadMemory, the IOCTL-backed memory
// reader/writer used by the `u`, `d`, `e`, `r`, and `s` commands (and
// indirectly by anything that needs to peek at debuggee memory).
//
// Mirrors libhyperdbg/code/debugger/misc/readmem.cpp. The C++ version has two
// backends: a local one (DeviceIoControl to hyperhv/hyperkd) and a remote one
// (serial/TCP packet). The Go version currently implements the local backend
// via comm.Device; the remote backend is added when the kernellvl module
// lands (Phase C.3.3).
package readmem

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
	"github.com/hyperdbg/go-libhyperdbg/types"
)

// DebuggerOperationWasSuccessful mirrors DEBUGGER_OPERATION_WAS_SUCCESSFUL
// (HyperDbg/hyperdbg/include/SDK/headers/ErrorCodes.h): the driver sets
// KernelStatus to this value when an IOCTL handler completed without error.
const DebuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// ReadMemory issues IOCTL_DEBUGGER_READ_MEMORY to read Size bytes at Address
// from the debuggee. The caller must supply an open Device (typically owned
// by core.Debugger).
//
// Returns the bytes read and the address-mode the driver probed (32 or 64
// when GetAddressMode was requested, 0 otherwise).
func ReadMemory(ctx context.Context, dev *comm.Device, addr uint64, pid uint32, size uint32, memType types.DEBUGGER_READ_MEMORY_TYPE, readingType types.DEBUGGER_READ_READING_TYPE, getAddrMode bool) ([]byte, types.DEBUGGER_READ_MEMORY_ADDRESS_MODE, error) {
	if dev == nil {
		return nil, 0, fmt.Errorf("ReadMemory: nil device (not connected?)")
	}
	if size == 0 {
		return nil, 0, nil
	}

	const pktSize = unsafe.Sizeof(types.DEBUGGER_READ_MEMORY{})
	pkt := types.DEBUGGER_READ_MEMORY{
		Pid:            pid,
		Address:        addr,
		Size:           size,
		GetAddressMode: getAddrMode,
		MemoryType:     memType,
		ReadingType:    readingType,
	}
	inBuf := (*[pktSize]byte)(unsafe.Pointer(&pkt))[:]

	// Output buffer: the response struct plus the read bytes.
	outBuf := make([]byte, int(pktSize)+int(size))

	returned, err := dev.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_READ_MEMORY, inBuf, outBuf)
	if err != nil {
		return nil, 0, fmt.Errorf("ReadMemory: IOCTL failed: %w", err)
	}
	if int(returned) < int(pktSize) {
		return nil, 0, fmt.Errorf("ReadMemory: short response (%d < %d)", returned, pktSize)
	}

	// Re-interpret the first pktSize bytes as the response.
	resp := (*types.DEBUGGER_READ_MEMORY)(unsafe.Pointer(&outBuf[0]))
	if resp.KernelStatus != DebuggerOperationWasSuccessful {
		return nil, 0, fmt.Errorf("ReadMemory: driver returned status 0x%08x", resp.KernelStatus)
	}

	// The driver's normal (non-vmx-root) read path does NOT set
	// resp.ReturnLength — only DebuggerCommandReadMemoryVmxRoot does. The
	// authoritative byte count is the IOCTL's bytesReturned minus the header
	// struct, which works for both paths. Fall back to resp.ReturnLength only
	// when the IOCTL returned nothing beyond the header (defensive).
	n := returned - uint32(pktSize)
	if rl := resp.ReturnLength; rl > 0 && rl < n {
		n = rl
	}
	if uint32(n) > size {
		n = size
	}
	payload := make([]byte, n)
	copy(payload, outBuf[pktSize:pktSize+uintptr(n)])
	return payload, resp.AddressMode, nil
}

// WriteMemory issues IOCTL_DEBUGGER_EDIT_MEMORY to write data at Address.
//
// The driver expects data packed in 64-bit chunks, so the caller's data is
// zero-padded to the next multiple of 8 bytes. Returns the in-chunk count
// the driver reports as written.
//
// NOTE: the full chunking logic from the C side is intentionally elided here;
// this is the minimal path used by the `e` command. The Phase C.3.6 work
// extends it to support per-byte and per-dword writes.
func WriteMemory(ctx context.Context, dev *comm.Device, addr uint64, pid uint32, data []byte, memType types.DEBUGGER_EDIT_MEMORY_TYPE, readingType types.DEBUGGER_READ_READING_TYPE) (uint32, error) {
	if dev == nil {
		return 0, fmt.Errorf("WriteMemory: nil device")
	}
	if len(data) == 0 {
		return 0, nil
	}

	// Pad to 8-byte chunks (the driver always reads/writes qwords).
	chunks := (uint32(len(data)) + 7) / 8
	const hdrSize = unsafe.Sizeof(types.DEBUGGER_EDIT_MEMORY{})
	pkt := types.DEBUGGER_EDIT_MEMORY{
		Address:            addr,
		ProcessId:          pid,
		MemoryType:         memType,
		CountOf64Chunks:    chunks,
		FinalStructureSize: uint32(hdrSize) + uint32(chunks*8),
	}
	inBuf := make([]byte, int(hdrSize)+int(chunks*8))
	src := (*[hdrSize]byte)(unsafe.Pointer(&pkt))[:]
	copy(inBuf[:hdrSize], src)
	copy(inBuf[hdrSize:], data)

	outBuf := make([]byte, hdrSize)
	if _, err := dev.Ioctl(ctx, comm.IOCTL_CODE_DEBUGGER_EDIT_MEMORY, inBuf, outBuf); err != nil {
		return 0, fmt.Errorf("WriteMemory: IOCTL failed: %w", err)
	}
	resp := (*types.DEBUGGER_EDIT_MEMORY)(unsafe.Pointer(&outBuf[0]))
	return resp.Result, nil
}
