// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

/*
#cgo CFLAGS: -I./lib/include
#include <Zydis/Zydis.h>

const ZyanStatus statusSuccess = ZYAN_STATUS_SUCCESS;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Decoder performs decoding of instruction bytes to a machine interpretable
// struct.
type Decoder struct {
	zdecoder C.ZydisDecoder
}

// NewDecoder returns a new instruction decoder.
func NewDecoder(mm MachineMode, aw AddressWidth) *Decoder {
	d := &Decoder{}
	C.ZydisDecoderInit(&d.zdecoder, C.ZydisMachineMode(mm), C.ZydisAddressWidth(aw))
	return d
}

// EnableMode enables/disables the given decoder mode.
func (d *Decoder) EnableMode(m DecoderMode, en bool) error {
	enValue := 0
	if en {
		enValue = 1
	}
	ret := C.ZydisDecoderEnableMode(&d.zdecoder, C.ZydisDecoderMode(m), C.ZyanBool(enValue))
	if zyanFailure(ret) {
		return fmt.Errorf("zydis: error enabling mode: 0x%x", ret)
	}
	return nil
}

// Decode decodes the instruction in the given input buffer.
func (d *Decoder) Decode(buffer []byte) (*DecodedInstruction, error) {
	var zinst C.ZydisDecodedInstruction
	ret := C.ZydisDecoderDecodeBuffer(
		&d.zdecoder,
		unsafe.Pointer(&buffer[0]),
		C.ZyanUSize(len(buffer)),
		&zinst,
	)
	if zyanFailure(ret) {
		return nil, fmt.Errorf("zydis: error decoding buffer: 0x%x", ret)
	}
	// Convert ZydisDecodedInstruction to a DecodedInstruction
	return newInstructionFromZydis(&zinst), nil
}

func zyanFailure(zs C.ZyanStatus) bool {
	return zs != C.statusSuccess
}

// DecoderMode is an enum of decoder modes.
type DecoderMode int

// DecoderMode enum values.
const (
	// Enables minimal instruction decoding without semantic analysis.
	//
	// This mode provides access to the mnemonic, the instruction-length, the
	// effective operand-size, the effective address-width, some attributes
	// (e.g. `ZYDIS_ATTRIB_IS_RELATIVE`) and all of the information in the
	// `raw` field of the `ZydisDecodedInstruction` struct.
	//
	// Operands, most attributes and other specific information (like `AVX`
	// info) are not accessible in this mode.
	//
	// This mode is NOT enabled by default.
	DecoderModeMinimal DecoderMode = iota

	// Enables the `AMD`-branch mode.
	//
	// Intel ignores the operand-size override-prefix (`0x66`) for all branches with 32-bit
	// immediates and forces the operand-size of the instruction to 64-bit in 64-bit mode.
	// In `AMD`-branch mode `0x66` is not ignored and changes the operand-size and the size of the
	// immediate to 16-bit.
	//
	// This mode is NOT enabled by default.
	DecoderModeAMDBranches

	// Enables `KNC` compatibility-mode.
	//
	// `KNC` and `KNL+` chips are sharing opcodes and encodings for some mask-related instructions.
	// Enable this mode to use the old `KNC` specifications (different mnemonics, operands, ..).
	//
	// This mode is NOT enabled by default.
	DecoderModeKNC

	// Enables the `MPX` mode.
	//
	// The `MPX` isa-extension reuses (overrides) some of the widenop instruction opcodes.
	//
	// This mode is enabled by default.
	DecoderModeMPX

	// Enables the `CET` mode.
	//
	// The `CET` isa-extension reuses (overrides) some of the widenop instruction opcodes.
	//
	// This mode is enabled by default.
	DecoderModeCET

	// Enables the `LZCNT` mode.
	//
	// The `LZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes.
	//
	// This mode is enabled by default.
	DecoderModeLZCNT

	// Enables the `TZCNT` mode.
	//
	// The `TZCNT` isa-extension reuses (overrides) some of the widenop instruction opcodes.
	//
	// This mode is enabled by default.
	DecoderModeTZCNT

	// Enables the `WBNOINVD` mode.
	//
	// The `WBINVD` instruction is interpreted as `WBNOINVD` on ICL chips, if a `F3` prefix is
	// used.
	//
	// This mode is disabled by default.
	DecoderModeWBNOINVD

	// Enables the `CLDEMOTE` mode.
	//
	// The `CLDEMOTE` isa-extension reuses (overrides) some of the widenop instruction opcodes.
	//
	// This mode is enabled by default.
	DecoderModeCLDEMOTE
)
