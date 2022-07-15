// Copyright 2019 John Papandriopoulos.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zydis

/*
#cgo CFLAGS: -I./lib/include
#include <Zydis/Zydis.h>
*/
import "C"

import (
	"fmt"
)

// OperandType is an enum of operand types.
type OperandType int

// OperandType enum values.
const (
	// The operand is not used.
	OperandTypeInvalid OperandType = iota
	// The operand is a register operand.
	OperandTypeRegister
	// The operand is a memory operand.
	OperandTypeMemory
	// The operand is a pointer operand with a segment:offset lvalue.
	OperandTypePointer
	// The operand is an immediate operand.
	OperandTypeImmediate
)

// OperandVisibility is an enum of operand visibility types.
type OperandVisibility int

// OperandVisibility enum values.
const (
	OperandVisibilityInvalid OperandVisibility = iota
	// The operand is explicitly encoded in the instruction.
	OperandVisibilityExplicit
	// The operand is part of the opcode, but listed as an operand.
	OperandVisibilityImplicit
	// The operand is part of the opcode, and not typically listed as an operand.
	OperandVisibilityHidden
)

// OperandAction is a bitfield that represents an operand action.
type OperandAction int

// OperandAction bitfield values.
const (
	// The operand is read by the instruction.
	OperandActionRead OperandAction = 1 << iota
	// The operand is written by the instruction (must write).
	OperandActionWrite
	// The operand is conditionally read by the instruction.
	OperandActionCondRead
	// The operand is conditionally written by the instruction (may write).
	OperandActionCondWrite

	// The operand is read (must read) and written by the instruction (must
	// write).
	OperandActionReadWrite = OperandActionRead | OperandActionWrite
	// The operand is conditionally read (may read) and conditionally written
	// by the instruction (may write).
	OperandActionCondReadCondWrite = OperandActionCondRead | OperandActionCondWrite
	// The operand is read (must read) and conditionally written by the
	// instruction (may write).
	OperandActionReadCondWrite = OperandActionRead | OperandActionCondWrite
	// The operand is written (must write) and conditionally read by
	// the instruction (may read).
	OperandActionCondReadWrite = OperandActionCondRead | OperandActionWrite
	// Mask combining all reading access flags.
	OperandActionMaskRead = OperandActionRead | OperandActionCondRead
	// Mask combining all writing access flags.
	OperandActionMaskWrite = OperandActionWrite | OperandActionCondWrite
)

// OperandActions is a...?
type OperandActions uint8

// MemoryOperandType is an enum of memory operand types.
type MemoryOperandType int

// MemoryOperandType enum values.
const (
	MemoryOperandTypeInvalid MemoryOperandType = iota
	// Normal memory operand
	MemoryOperandTypeMemory
	// The memory operand is only used for address-generation.
	// No real memory-access is caused.
	MemoryOperandTypeAddressGeneration
	// A memory operand using `SIB` addressing form, where the index register
	// is not used in address calculation and scale is ignored.
	// No real memory-access is caused.
	MemoryOperandTypeMIB
)

// OperandEncoding is an enum of the operand encoding.
type OperandEncoding int

// OperandEncoding enum
const (
	OperandEncodingNone OperandEncoding = iota
	OperandEncodingModRmReg
	OperandEncodingModRmRm
	OperandEncodingOpcode
	OperandEncodingNDSNDD
	OperandEncodingIS4
	OperandEncodingMask
	OperandEncodingDisp8
	OperandEncodingDisp16
	OperandEncodingDisp32
	OperandEncodingDisp64
	OperandEncodingDisp16_32_64
	OperandEncodingDisp32_32_64
	OperandEncodingDisp16_32_32
	OperandEncodingUImm8
	OperandEncodingUImm16
	OperandEncodingUImm32
	OperandEncodingUImm64
	OperandEncodingUImm16_32_64
	OperandEncodingUImm32_32_64
	OperandEncodingUImm16_32_32
	OperandEncodingSImm8
	OperandEncodingSImm16
	OperandEncodingSImm32
	OperandEncodingSImm64
	OperandEncodingSImm16_32_64
	OperandEncodingSImm32_32_64
	OperandEncodingSImm16_32_32
	OperandEncodingJImm8
	OperandEncodingJImm16
	OperandEncodingJImm32
	OperandEncodingJImm64
	OperandEncodingJImm16_32_64
	OperandEncodingJImm32_32_64
	OperandEncodingJImm16_32_32
)

// DecodedOperand is an operand in a DecodedInstruction.
type DecodedOperand struct {
	instr *DecodedInstruction
	zdo   *C.ZydisDecodedOperand

	// Operand-ID.
	ID uint8
	// Type of the operand.
	Type OperandType
	// Visibility of the operand.
	Visibility OperandVisibility
	// Operand actions.
	Actions OperandActions
	// Operand encoding.
	Encoding OperandEncoding
	// Logical size of the operand (in bits).
	Size uint16
	// Element type.
	ElementType
	// Element size.
	ElementSize
	// Number of elements.
	ElementCount uint16
	// Extended info for register operands.
	Reg struct {
		Value Register
	}
	// Extended info for memory operands.
	Mem struct {
		// Type of the memory operand.
		Type MemoryOperandType
		// Segment register.
		Segment Register
		// Base register.
		Base Register
		// Index register.
		Index Register
		// Scale factor.
		Scale uint8
		// Extended info for memory-operands with displacement.
		Disp struct {
			// Signals if the displacement value is used.
			HasDisplacement bool
			// The displacement value.
			Value int64
		}
	}
	// Extended info for pointer-operands.
	Ptr struct {
		Segment uint16
		Offset  uint32
	}
	// Extended info for immediate-operands.
	Imm struct {
		// Signals if the immediate value is signed.
		IsSigned bool
		// Signals if the immediate value contains a relative offset.
		// You can use `ZydisCalcAbsoluteAddress` to determine the absolute
		// address value.
		IsRelative bool
		// The immediate value.
		Value struct {
			Unsigned uint64
			Signed   int64
		}
	}
}

// AbsoluteAddress returns the absolute address value for the given instruction operand.
//
// You should use this function in the following cases:
// - `IMM` operands with relative address (e.g. `JMP`, `CALL`, ...)
// - `MEM` operands with `RIP`/`EIP`-relative address (e.g. `MOV RAX, [RIP+0x12345678]`)
// - `MEM` operands with absolute address (e.g. `MOV RAX, [0x12345678]`)
//   - The displacement needs to get truncated and zero extended
func (do *DecodedOperand) AbsoluteAddress(runtimeAddress uint64) (uint64, error) {
	var resultAddress C.ZyanU64
	ret := C.ZydisCalcAbsoluteAddress(
		do.instr.zdi,
		do.zdo,
		C.ZyanU64(runtimeAddress),
		&resultAddress,
	)
	if zyanFailure(ret) {
		return 0, fmt.Errorf("zydis: failed to compute absolute address: 0x%x", ret)
	}
	return uint64(resultAddress), nil
}

// RegisterContext is a mapping from register to its value
type RegisterContext map[Register]uint64

// AbsoluteAddressWithContext is like AbsoluteAddress, but takes some register
// context values, to allow calculation of addresses depending on runtime
// register values.
//
// Note that `IP/EIP/RIP` from the register-context will be ignored in favor
// of the passed runtime-address.
func (do *DecodedOperand) AbsoluteAddressWithContext(runtimeAddress uint64, context RegisterContext) (uint64, error) {
	var rctx C.ZydisRegisterContext
	for reg, val := range context {
		rctx.values[reg] = C.ZyanU64(val)
	}
	var resultAddress C.ZyanU64
	ret := C.ZydisCalcAbsoluteAddressEx(
		do.instr.zdi,
		do.zdo,
		C.ZyanU64(runtimeAddress),
		&rctx,
		&resultAddress,
	)
	if zyanFailure(ret) {
		return 0, fmt.Errorf("zydis: failed to compute absolute address: 0x%x", ret)
	}
	return uint64(resultAddress), nil
}
