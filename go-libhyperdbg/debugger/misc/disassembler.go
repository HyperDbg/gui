// Package misc implements the support modules under
// libhyperdbg/code/debugger/misc/: assembler (ok/keystone), disassembler
// (ok/zydis), callstack, readmem, pt-helper, pci-id.
//
// This file owns the disassembler: a thin wrapper around ok/zydis that
// mirrors the C++ ZydisDisassemble helper used by the `u` and `d` commands.
// The wrapper exists so the rest of go-libhyperdbg does not import ok/zydis
// directly (keeping the DLL-embedding concern localised to this package).
package misc

import (
	"fmt"
	"unsafe"

	"github.com/ddkwork/zydis"
)

// MachineMode selects the decoding mode (16/32/64-bit). Mirrors
// ZydisMachineMode constants from the C++ side.
type MachineMode uint8

const (
	ModeLong64   MachineMode = MachineMode(zydis.ZydisMachineModeLong64)
	ModeLegacy32 MachineMode = MachineMode(zydis.ZydisMachineModeLegacy32)
	ModeLegacy16 MachineMode = MachineMode(zydis.ZydisMachineModeLegacy16)
)

// DisasmResult is the output of Disassemble: the textual rendering plus the
// decoded length (so callers can advance the cursor).
type DisasmResult struct {
	Text    string // e.g. "mov rax, rbx"
	Length  int    // bytes consumed from the input
	Runtime uint64 // runtime address (echoed back)
}

// Disassembler wraps a zydis.Zydis instance. It is safe for concurrent use
// only from a single goroutine (the underlying zydis DLL is not guarded); a
// Debugger that needs multi-goroutine disasm should create one per consumer.
type Disassembler struct {
	z *zydis.Zydis
}

// NewDisassembler creates a Disassembler using the default zydis instance.
// The first call triggers DLL extraction (see ok/zydis/dll.go init()).
func NewDisassembler() *Disassembler {
	return &Disassembler{z: &zydis.Zydis{}}
}

// Disassemble decodes a single instruction from code at runtimeAddr.
// mode selects 16/32/64-bit decoding. Returns the text and the number of
// bytes consumed.
//
// Mirrors the C++ ZydisDisassemble(WindowsMode, DisassemblerAddress,
// buffer, ...) helper.
func (d *Disassembler) Disassemble(mode MachineMode, runtimeAddr uint64, code []byte) (DisasmResult, error) {
	if len(code) == 0 {
		return DisasmResult{}, fmt.Errorf("Disassemble: empty code buffer")
	}
	var disasm zydis.ZydisDisassembledInstruction
	status := d.z.DisassembleIntel(
		zydis.ZydisMachineMode(mode),
		runtimeAddr,
		unsafe.Pointer(&code[0]),
		uintptr(len(code)),
		&disasm,
	)
	if status != zydis.ZyanStatusSuccess {
		return DisasmResult{}, fmt.Errorf("Disassemble: zydis decode failed (status=0x%X)", status)
	}
	// disasm.Text is a fixed-size [N]int8 array; find the NUL and convert.
	n := 0
	for n < len(disasm.Text) && disasm.Text[n] != 0 {
		n++
	}
	// int8 → byte conversion (zydis uses char[] which maps to int8 in Go).
	textBytes := make([]byte, n)
	for i := 0; i < n; i++ {
		textBytes[i] = byte(disasm.Text[i])
	}
	return DisasmResult{
		Text:    string(textBytes),
		Length:  int(disasm.Info.Length),
		Runtime: runtimeAddr,
	}, nil
}

// DisassembleRange decodes every instruction in code, starting at baseAddr,
// and returns the results. Decoding stops on the first error (e.g. invalid
// opcode or running past the end of the buffer).
func (d *Disassembler) DisassembleRange(mode MachineMode, baseAddr uint64, code []byte) ([]DisasmResult, error) {
	var out []DisasmResult
	addr := baseAddr
	off := 0
	for off < len(code) {
		r, err := d.Disassemble(mode, addr, code[off:])
		if err != nil {
			return out, err
		}
		out = append(out, r)
		off += r.Length
		addr += uint64(r.Length)
	}
	return out, nil
}
