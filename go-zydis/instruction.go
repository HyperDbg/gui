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
	"unsafe"
)

const (
	maxInstructionLength = 16
	maxOperandCount      = 10
)

// DecodedInstruction represents a processor instruction.
type DecodedInstruction struct {
	zdi *C.ZydisDecodedInstruction

	// The machine mode used to decode this instruction.
	MachineMode MachineMode
	// The instruction-mnemonic.
	Mnemonic Mnemonic
	// The length of the decoded instruction.
	Length uint8
	// The instruction-encoding.
	Encoding InstructionEncoding
	// Opcode-map.
	OpcodeMap OpcodeMap
	// Instruction opcode.
	Opcode uint8
	// Stack width.
	StackWidth uint8
	// Effective operand width.
	OperandWidth uint8
	// Effective address width.
	AddressWidth uint8
	// Detailed info for all instruction operands.
	//
	// Explicit operands are guaranteed to be in the front and ordered as they are printed
	// by the formatter in Intel mode. No assumptions can be made about the order of hidden
	// operands, except that they always located behind the explicit operands.
	Operands []DecodedOperand
	// Instruction attributes.
	Attributes InstructionAttributes
	// Information about accessed CPU flags.
	AccessedFlags []struct {
		Action CPUFlagAction
	}
	// Extended info for `AVX` instructions.
	AVX struct {
		// The `AVX` vector-length.
		VectorLength uint16
		// Info about the embedded writemask-register (`AVX-512` and `KNC` only).
		Mask struct {
			// The masking mode.
			Mode MaskMode
			// The mask register.
			Reg Register
		}
		// Contains info about the `AVX` broadcast.
		Broadcast struct {
			// Signals if the broadcast is a static broadcast.
			IsStatic bool
			// The `AVX` broadcast-mode.
			Mode BroadcastMode
		}
		// Contains info about the `AVX` rounding.
		Rounding struct {
			// The `AVX` rounding-mode.
			Mode RoundingMode
		}
		// Contains info about the `AVX` register-swizzle (`KNC` only).
		Swizzle struct {
			// The `AVX` register-swizzle mode.
			Mode SwizzleMode
		}
		// Contains info about the `AVX` data-conversion (`KNC` only).
		Conversion struct {
			// The `AVX` data-conversion mode.
			Mode ConversionMode
		}
		// Signals if the `SAE` (suppress-all-exceptions) functionality is
		// enabled for the instruction.
		HasSAE bool
		// Signals if the instruction has a memory-eviction-hint (`KNC` only).
		HasEvictionHint bool
	}
	// Meta info.
	Meta struct {
		// The instruction category.
		Category InstructionCategory
		// The ISA-set.
		ISASet ISASet
		// The ISA-set extension.
		ISAExt ISAExt
		// The branch type.
		BranchType BranchType
		// The exception class.
		ExceptionClass ExceptionClass
	}
	// Detailed info about different instruction-parts like `ModRM`, `SIB` or
	// encoding-prefixes.
	Raw struct {
		// Detailed info about the legacy prefixes (including `REX`).
		Prefixes []struct {
			// The prefix type.
			Type PrefixType
			// The prefix byte.
			Value uint8
		}
		// Detailed info about the `REX` prefix.
		REX struct {
			// 64-bit operand-size promotion.
			W uint8
			// Extension of the `ModRM.reg` field.
			R uint8
			// Extension of the `SIB.index` field.
			X uint8
			// Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field.
			B uint8
			// The offset of the effective `REX` byte, relative to the beginning of the
			// instruction, in bytes.
			//
			// This offset always points to the "effective" `REX` prefix (the one
			// closest to the instruction opcode), if multiple `REX` prefixes are
			// present.
			//
			// Note that the `REX` byte can be the first byte of the instruction,
			// which would lead to an offset of `0`. Please refer to the instruction
			// attributes to check for the presence of the `REX` prefix.
			Offset uint8
		}
		// Detailed info about the `XOP` prefix.
		XOP struct {
			// Extension of the `ModRM.reg` field (inverted).
			R uint8
			// Extension of the `SIB.index` field (inverted).
			X uint8
			// Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field (inverted).
			B uint8
			// Opcode-map specifier.
			MMMMM uint8
			// 64-bit operand-size promotion or opcode-extension.
			W uint8
			// `NDS`/`NDD` (non-destructive-source/destination) register specifier
			// (inverted).
			VVVV uint8
			// Vector-length specifier.
			L uint8
			// Compressed legacy prefix.
			PP uint8
			// The offset of the first xop byte, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
		// Detailed info about the `VEX` prefix.
		VEX struct {
			// Extension of the `ModRM.reg` field (inverted).
			R uint8
			// Extension of the `SIB.index` field (inverted).
			X uint8
			// Extension of the `ModRM.rm`, `SIB.base`, or `opcode.reg` field
			// (inverted).
			B uint8
			// Opcode-map specifier.
			MMMMM uint8
			// 64-bit operand-size promotion or opcode-extension.
			W uint8
			// `NDS`/`NDD` (non-destructive-source/destination) register specifier
			// (inverted).
			VVVV uint8
			// Vector-length specifier.
			L uint8
			// Compressed legacy prefix.
			PP uint8
			// The offset of the first `VEX` byte, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
			// The size of the `VEX` prefix, in bytes.
			Size uint8
		}
		// Detailed info about the `EVEX` prefix.
		EVEX struct {
			// Extension of the `ModRM.reg` field (inverted).
			R uint8
			// Extension of the `SIB.index/vidx` field (inverted).
			X uint8
			// Extension of the `ModRM.rm` or `SIB.base` field (inverted).
			B uint8
			// High-16 register specifier modifier (inverted).
			R2 uint8
			// Opcode-map specifier.
			MM uint8
			// 64-bit operand-size promotion or opcode-extension.
			W uint8
			// `NDS`/`NDD` (non-destructive-source/destination) register specifier
			// (inverted).
			VVVV uint8
			// Compressed legacy prefix.
			PP uint8
			// Zeroing/Merging.
			Z uint8
			// Vector-length specifier or rounding-control (most significant bit).
			L2 uint8
			// Vector-length specifier or rounding-control (least significant bit).
			L uint8
			// Broadcast/RC/SAE context.
			BLower uint8
			// High-16 `NDS`/`VIDX` register specifier.
			V2 uint8
			// Embedded opmask register specifier.
			AAA uint8
			// The offset of the first evex byte, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
		// Detailed info about the `MVEX` prefix.
		MVEX struct {
			// Extension of the `ModRM.reg` field (inverted).
			R uint8
			// Extension of the `SIB.index/vidx` field (inverted).
			X uint8
			// Extension of the `ModRM.rm` or `SIB.base` field (inverted).
			B uint8
			// High-16 register specifier modifier (inverted).
			R2 uint8
			// Opcode-map specifier.
			MMMM uint8
			// 64-bit operand-size promotion or opcode-extension.
			W uint8
			// `NDS`/`NDD` (non-destructive-source/destination) register specifier
			// (inverted).
			VVVV uint8
			// Compressed legacy prefix.
			PP uint8
			// Non-temporal/eviction hint.
			E uint8
			// Swizzle/broadcast/up-convert/down-convert/static-rounding controls.
			SSS uint8
			// High-16 `NDS`/`VIDX` register specifier.
			V2 uint8
			// Embedded opmask register specifier.
			KKK uint8
			// The offset of the first mvex byte, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
		// Detailed info about the `ModRM` byte.
		ModRM struct {
			// The addressing mode.
			Mod uint8
			// Register specifier or opcode-extension.
			Reg uint8
			// Register specifier or opcode-extension.
			Rm uint8
			// The offset of the `ModRM` byte, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
		// Detailed info about the `SIB` byte.
		SIB struct {
			// The scale factor.
			Scale uint8
			// The index-register specifier.
			Index uint8
			// The base-register specifier.
			Base uint8
			// The offset of the `SIB` byte, relative to the beginning of the instruction,
			// in bytes.
			Offset uint8
		}
		// Detailed info about displacement-bytes.
		Disp struct {
			// The displacement value
			Value int64
			// The physical displacement size, in bits.
			Size uint8
			// The offset of the displacement data, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
		// Detailed info about immediate-bytes.
		Imm [2]struct {
			// Signals, if the immediate value is signed.
			IsSigned bool
			// Signals, if the immediate value contains a relative offset. You can use
			// `ZydisCalcAbsoluteAddress` to determine the absolute address value.
			IsRelative bool
			// The immediate value.
			Value struct {
				Unsigned uint64
				Signed   int64
			}
			// The physical immediate size, in bits.
			Size uint8
			// The offset of the immediate data, relative to the beginning of the
			// instruction, in bytes.
			Offset uint8
		}
	}
}

// InstructionEncoding is an enum of instruction encodings.
type InstructionEncoding int

// InstructionEncoding enum values.
const (
	InstructionEncodingLegacy = iota
	InstructionEncoding3DNOW
	InstructionEncodingXOP
	InstructionEncodingVEX
	InstructionEncodingEVEX
	InstructionEncodingMVEX
)

// InstructionAttributes is a bitfield of instruction attributes.
type InstructionAttributes uint64

// InstructionAttributes bitfield values
const (
	// The instruction has the `ModRM` byte.
	InstructionAttributesHasModRM InstructionAttributes = 1 << iota
	// The instruction has the `SIB` byte.
	InstructionAttributesHasSIB
	// The instruction has the `REX` prefix.
	InstructionAttributesHasREX
	// The instruction has the `XOP` prefix.
	InstructionAttributesHasXOP
	// The instruction has the `VEX` prefix.
	InstructionAttributesHasVEX
	// The instruction has the `EVEX` prefix.
	InstructionAttributesHasEVEX
	// The instruction has the `MVEX` prefix.
	InstructionAttributesHasMVEX
	// The instruction has one or more operands with position-relative offsets.
	InstructionAttributesIsRelative
	// The instruction is privileged.
	InstructionAttributesIsPrivileged
	// The instruction accepts the `LOCK` prefix (`0xF0`).
	InstructionAttributesAcceptsLock
	// The instruction accepts the `REP` prefix (`0xF3`).
	InstructionAttributesAcceptsREP
	// The instruction accepts the `REPE`/`REPZ` prefix (`0xF3`).
	InstructionAttributesAcceptsREPE
	// The instruction accepts the `REPE`/`REPZ` prefix (`0xF3`).
	InstructionAttributesAcceptsREPZ InstructionAttributes = 1 << (iota - 1)
	// The instruction accepts the `REPNE`/`REPNZ` prefix (`0xF2`).
	InstructionAttributesAcceptsREPNE
	// The instruction accepts the `REPNE`/`REPNZ` prefix (`0xF2`).
	InstructionAttributesAcceptsREPNZ InstructionAttributes = 1 << (iota - 1)
	// The instruction accepts the `BND` prefix (`0xF2`).
	InstructionAttributesAcceptsBND
	// The instruction accepts the `XACQUIRE` prefix (`0xF2`).
	InstructionAttributesAcceptsXAcquire
	// The instruction accepts the `XRELEASE` prefix (`0xF3`).
	InstructionAttributesAcceptsXRelease
	// The instruction accepts the `XACQUIRE`/`XRELEASE` prefixes (`0xF2`, `0xF3`)
	// without the `LOCK` prefix (`0x0F`).
	InstructionAttributesAcceptsHLEWithoutLock
	// The instruction accepts branch hints (0x2E, 0x3E).
	InstructionAttributesAcceptsBranchHints
	// The instruction accepts segment prefixes (`0x2E`, `0x36`, `0x3E`, `0x26`,
	// `0x64`, `0x65`).
	InstructionAttributesAcceptsSegment
	// The instruction has the `LOCK` prefix (`0xF0`).
	InstructionAttributesHasLock
	// The instruction has the `REP` prefix (`0xF3`).
	InstructionAttributesHasREP
	// The instruction has the `REPE`/`REPZ` prefix (`0xF3`).
	InstructionAttributesHasREPE
	// The instruction has the `REPE`/`REPZ` prefix (`0xF3`).
	InstructionAttributesHasREPZ
	// The instruction has the `REPNE`/`REPNZ` prefix (`0xF2`).
	InstructionAttributesHasREPNZ
	// The instruction has the `BND` prefix (`0xF2`).
	InstructionAttributesHasBND
	// The instruction has the `XACQUIRE` prefix (`0xF2`).
	InstructionAttributesHasXAcquire
	// The instruction has the `XRELEASE` prefix (`0xF3`).
	InstructionAttributesHasXRelease
	// The instruction has the branch-not-taken hint (`0x2E`).
	InstructionAttributesHasBranchNotTaken
	// The instruction has the branch-taken hint (`0x3E`).
	InstructionAttributesHasBranchTaken
	// The instruction has the `CS` segment modifier (`0x2E`).
	InstructionAttributesHasSegmentCS
	// The instruction has the `SS` segment modifier (`0x36`).
	InstructionAttributesHasSegmentSS
	// The instruction has the `DS` segment modifier (`0x3E`).
	InstructionAttributesHasSegmentDS
	// The instruction has the `ES` segment modifier (`0x26`).
	InstructionAttributesHasSegmentES
	// The instruction has the `FS` segment modifier (`0x64`).
	InstructionAttributesHasSegmentFS
	// The instruction has the `GS` segment modifier (`0x65`).
	InstructionAttributesHasSegmentGS
	// The instruction has the operand-size override prefix (`0x66`).
	InstructionAttributesHasOperandSize
	// The instruction has the address-size override prefix (`0x67`).
	InstructionAttributesHasAddressSize
	// The instruction may conditionally read the general CPU state.
	InstructionAttributesCPUFlagAccess
	// The instruction may conditionally write the general CPU state.
	InstructionAttributesCPUStateCR
	// The instruction may conditionally read the FPU state (X87, MMX).
	InstructionAttributesCPUStateCW
	// The instruction may conditionally write the FPU state (X87, MMX).
	InstructionAttributesFPUStateCR
	// The instruction may conditionally read the XMM state (AVX, AVX2, AVX-512).
	InstructionAttributesXMMStateCR
	// The instruction may conditionally write the XMM state (AVX, AVX2, AVX-512).
	InstructionAttributesXMMStateCW

	// The instruction has a segment modifier.
	InstructionAttributesHasSegment InstructionAttributes = InstructionAttributesHasSegmentCS |
		InstructionAttributesHasSegmentSS |
		InstructionAttributesHasSegmentDS |
		InstructionAttributesHasSegmentES |
		InstructionAttributesHasSegmentFS |
		InstructionAttributesHasSegmentGS
)

// MaskMode is an enum of AVX mask modes.
type MaskMode int

// MaskMode enum values.
const (
	MaskModeInvalid MaskMode = iota
	// Masking is disabled for the current instruction (`K0` register is used).
	MaskModeDisabled
	// The embedded mask register is used as a merge-mask.
	MaskModeMerging
	// The embedded mask register is used as a zero-mask.
	MaskModeZeroing
	// The embedded mask register is used as a control-mask (element selector).
	MaskModeControl
	// The embedded mask register is used as a zeroing control-mask (element selector).
	MaskModeControlZeroing
)

// BroadcastMode is an enum of AVX broadcast modes.
type BroadcastMode int

// BroadcastMode enum values.
const (
	BroadcastModeInvalid BroadcastMode = iota
	BroadcastMode1To2
	BroadcastMode1To4
	BroadcastMode1to8
	BroadcastMode1to16
	BroadcastMode1to32
	BroadcastMode1to64
	BroadcastMode2to4
	BroadcastMode2to8
	BroadcastMode2to16
	BroadcastMode4to8
	BroadcastMode4to16
	BroadcastMode8to16
)

// RoundingMode is an enum of AVX rounding modes.
type RoundingMode int

// RoundingMode enum values.
const (
	RoundingModeInvalid RoundingMode = iota
	// Round to nearest.
	RoundingModeRN
	// Round down.
	RoundingModeRD
	// Round up.
	RoundingModeRU
	// Round towards zero.
	RoundingModeRZ
)

// SwizzleMode is an enum of KNC swizzle modes.
type SwizzleMode int

// SwizzleMode enum values.
const (
	SwizzleModeInvalid SwizzleMode = iota
	SwizzleModeDCBA
	SwizzleModeCDAB
	SwizzleModeBADC
	SwizzleModeDACB
	SwizzleModeAAAA
	SwizzleModeBBBB
	SwizzleModeCCCC
	SwizzleModeDDDD
)

// ConversionMode is an enum of KNC conversion modes.
type ConversionMode int

// ConversionMode enum values.
const (
	ConversionModeInvalid ConversionMode = iota
	ConversionModeFloat16
	ConversionModeSInt8
	ConversionModeUInt8
	ConversionModeSInt16
	ConversionModeUInt16
)

// PrefixType is an enum of prefix types.
type PrefixType int

// PrefixType enum values.
const (
	// The prefix is ignored by the instruction.
	// This applies to all prefixes that are not accepted by the instruction
	// in general or the ones that are overwritten by a prefix of the same
	// group closer to the instruction opcode.
	PrefixTypeIgnored PrefixType = iota
	// The prefix is effectively used by the instruction.
	PrefixTypeEffective
	// The prefix is used as a mandatory prefix, interpreted as an opcode
	// extension and has no further effect on the instruction.
	PrefixTypeMandatory
)

// InstructionCategory is an enum of instruction categories.
type InstructionCategory int

//go:generate stringer -type=InstructionCategory -linecomment

// InstructionCategory enum values
const (
	InstructionCategoryInvalid             InstructionCategory = iota // INVALID
	InstructionCategoryADOX_ADCX                                      // ADOX_ADCX
	InstructionCategoryAES                                            // AES
	InstructionCategoryAMD3DNOW                                       // AMD3DNOW
	InstructionCategoryAMX_TILE                                       // AMX_TILE
	InstructionCategoryAVX                                            // AVX
	InstructionCategoryAVX2                                           // AVX2
	InstructionCategoryAVX2GATHER                                     // AVX2GATHER
	InstructionCategoryAVX512                                         // AVX512
	InstructionCategoryAVX512_4FMAPS                                  // AVX512_4FMAPS
	InstructionCategoryAVX512_4VNNIW                                  // AVX512_4VNNIW
	InstructionCategoryAVX512_BITALG                                  // AVX512_BITALG
	InstructionCategoryAVX512_VBMI                                    // AVX512_VBMI
	InstructionCategoryAVX512_VP2INTERSECT                            // AVX512_VP2INTERSECT
	InstructionCategoryBINARY                                         // BINARY
	InstructionCategoryBITBYTE                                        // BITBYTE
	InstructionCategoryBLEND                                          // BLEND
	InstructionCategoryBMI1                                           // BMI1
	InstructionCategoryBMI2                                           // BMI2
	InstructionCategoryBROADCAST                                      // BROADCAST
	InstructionCategoryCALL                                           // CALL
	InstructionCategoryCET                                            // CET
	InstructionCategoryCLDEMOTE                                       // CLDEMOTE
	InstructionCategoryCLFLUSHOPT                                     // CLFLUSHOPT
	InstructionCategoryCLWB                                           // CLWB
	InstructionCategoryCLZERO                                         // CLZERO
	InstructionCategoryCMOV                                           // CMOV
	InstructionCategoryCOMPRESS                                       // COMPRESS
	InstructionCategoryCOND_BR                                        // COND_BR
	InstructionCategoryCONFLICT                                       // CONFLICT
	InstructionCategoryCONVERT                                        // CONVERT
	InstructionCategoryDATAXFER                                       // DATAXFER
	InstructionCategoryDECIMAL                                        // DECIMAL
	InstructionCategoryENQCMD                                         // ENQCMD
	InstructionCategoryEXPAND                                         // EXPAND
	InstructionCategoryFCMOV                                          // FCMOV
	InstructionCategoryFLAGOP                                         // FLAGOP
	InstructionCategoryFMA4                                           // FMA4
	InstructionCategoryGATHER                                         // GATHER
	InstructionCategoryGFNI                                           // GFNI
	InstructionCategoryIFMA                                           // IFMA
	InstructionCategoryINTERRUPT                                      // INTERRUPT
	InstructionCategoryIO                                             // IO
	InstructionCategoryIOSTRINGOP                                     // IOSTRINGOP
	InstructionCategoryKMASK                                          // KMASK
	InstructionCategoryKNC                                            // KNC
	InstructionCategoryKNCMASK                                        // KNCMASK
	InstructionCategoryKNCSCALAR                                      // KNCSCALAR
	InstructionCategoryLOGICAL                                        // LOGICAL
	InstructionCategoryLOGICAL_FP                                     // LOGICAL_FP
	InstructionCategoryLZCNT                                          // LZCNT
	InstructionCategoryMISC                                           // MISC
	InstructionCategoryMMX                                            // MMX
	InstructionCategoryMOVDIR                                         // MOVDIR
	InstructionCategoryMPX                                            // MPX
	InstructionCategoryNOP                                            // NOP
	InstructionCategoryPADLOCK                                        // PADLOCK
	InstructionCategoryPCLMULQDQ                                      // PCLMULQDQ
	InstructionCategoryPCONFIG                                        // PCONFIG
	InstructionCategoryPKU                                            // PKU
	InstructionCategoryPOP                                            // POP
	InstructionCategoryPREFETCH                                       // PREFETCH
	InstructionCategoryPREFETCHWT1                                    // PREFETCHWT1
	InstructionCategoryPT                                             // PT
	InstructionCategoryPUSH                                           // PUSH
	InstructionCategoryRDPID                                          // RDPID
	InstructionCategoryRDPRU                                          // RDPRU
	InstructionCategoryRDRAND                                         // RDRAND
	InstructionCategoryRDSEED                                         // RDSEED
	InstructionCategoryRDWRFSGS                                       // RDWRFSGS
	InstructionCategoryRET                                            // RET
	InstructionCategoryROTATE                                         // ROTATE
	InstructionCategorySCATTER                                        // SCATTER
	InstructionCategorySEGOP                                          // SEGOP
	InstructionCategorySEMAPHORE                                      // SEMAPHORE
	InstructionCategorySERIALIZE                                      // SERIALIZE
	InstructionCategorySETCC                                          // SETCC
	InstructionCategorySGX                                            // SGX
	InstructionCategorySHA                                            // SHA
	InstructionCategorySHIFT                                          // SHIFT
	InstructionCategorySMAP                                           // SMAP
	InstructionCategorySSE                                            // SSE
	InstructionCategorySTRINGOP                                       // STRINGOP
	InstructionCategorySTTNI                                          // STTNI
	InstructionCategorySYSCALL                                        // SYSCALL
	InstructionCategorySYSRET                                         // SYSRET
	InstructionCategorySYSTEM                                         // SYSTEM
	InstructionCategoryTBM                                            // TBM
	InstructionCategoryTSX_LDTRK                                      // TSX_LDTRK
	InstructionCategoryUFMA                                           // UFMA
	InstructionCategoryUNCOND_BR                                      // UNCOND_BR
	InstructionCategoryVAES                                           // VAES
	InstructionCategoryVBMI2                                          // VBMI2
	InstructionCategoryVFMA                                           // VFMA
	InstructionCategoryVPCLMULQDQ                                     // VPCLMULQDQ
	InstructionCategoryVTX                                            // VTX
	InstructionCategoryWAITPKG                                        // WAITPKG
	InstructionCategoryWIDENOP                                        // WIDENOP
	InstructionCategoryX87_ALU                                        // X87_ALU
	InstructionCategoryXOP                                            // XOP
	InstructionCategoryXSAVE                                          // XSAVE
	InstructionCategoryXSAVEOPT                                       // XSAVEOPT
)

// BranchType is an enum of instruction branch types.
type BranchType int

// BranchType enum values.
const (
	// The instruction is not a branch instruction.
	BranchTypeNone BranchType = iota
	// The instruction is a short (8-bit) branch instruction.
	BranchTypeShort
	// The instruction is a near (16-bit or 32-bit) branch instruction.
	BranchTypeNear
	// The instruction is a far (intersegment) branch instruction.
	BranchTypeFar
)

// ExceptionClass is an enum of exception classes.
type ExceptionClass int

//go:generate stringer -type=ExceptionClass -linecomment

// ExceptionClass enum values.
const (
	ExceptionClassNone    ExceptionClass = iota
	ExceptionClassSSE1                   // SSE1
	ExceptionClassSSE2                   // SSE2
	ExceptionClassSSE3                   // SSE3
	ExceptionClassSSE4                   // SSE4
	ExceptionClassSSE5                   // SSE5
	ExceptionClassSSE7                   // SSE7
	ExceptionClassAVX1                   // AVX1
	ExceptionClassAVX2                   // AVX2
	ExceptionClassAVX3                   // AVX3
	ExceptionClassAVX4                   // AVX4
	ExceptionClassAVX5                   // AVX5
	ExceptionClassAVX6                   // AVX6
	ExceptionClassAVX7                   // AVX7
	ExceptionClassAVX8                   // AVX8
	ExceptionClassAVX11                  // AVX11
	ExceptionClassAVX12                  // AVX12
	ExceptionClassE1                     // E1
	ExceptionClassE1NF                   // E1NF
	ExceptionClassE2                     // E2
	ExceptionClassE2NF                   // E2NF
	ExceptionClassE3                     // E3
	ExceptionClassE3NF                   // E3NF
	ExceptionClassE4                     // E4
	ExceptionClassE4NF                   // E4NF
	ExceptionClassE5                     // E5
	ExceptionClassE5NF                   // E5NF
	ExceptionClassE6                     // E6
	ExceptionClassE6NF                   // E6NF
	ExceptionClassE7NM                   // E7NM
	ExceptionClassE7NM128                // E7NM128
	ExceptionClassE9NF                   // E9NF
	ExceptionClassE10                    // E10
	ExceptionClassE10NF                  // E10NF
	ExceptionClassE11                    // E11
	ExceptionClassE11NF                  // E11NF
	ExceptionClassE12                    // E12
	ExceptionClassE12NP                  // E12NP
	ExceptionClassK20                    // K20
	ExceptionClassK21                    // K21
	ExceptionClassAMXE1                  // AMXE1
	ExceptionClassAMXE2                  // AMXE2
	ExceptionClassAMXE3                  // AMXE3
	ExceptionClassAMXE4                  // AMXE4
	ExceptionClassAMXE5                  // AMXE5
	ExceptionClassAMXE6                  // AMXE6
)

func newInstructionFromZydis(zins *C.ZydisDecodedInstruction) *DecodedInstruction {
	ins := &DecodedInstruction{}

	ins.zdi = zins
	ins.MachineMode = MachineMode(zins.machine_mode)
	ins.Mnemonic = Mnemonic(zins.mnemonic)
	ins.Length = uint8(zins.length)
	ins.Encoding = InstructionEncoding(zins.encoding)
	ins.OpcodeMap = OpcodeMap(zins.opcode_map)
	ins.Opcode = uint8(zins.opcode)
	ins.StackWidth = uint8(zins.stack_width)
	ins.OperandWidth = uint8(zins.operand_width)
	ins.AddressWidth = uint8(zins.address_width)
	ins.Operands = make([]DecodedOperand, int(zins.operand_count))

	for i := 0; i < len(ins.Operands); i++ {
		op := &ins.Operands[i]
		zop := &zins.operands[i]

		op.instr = ins
		op.zdo = zop

		op.ID = uint8(zop.id)
		op.Type = OperandType(zop._type)
		op.Visibility = OperandVisibility(zop.visibility)
		op.Actions = OperandActions(zop.actions)
		op.Encoding = OperandEncoding(zop.encoding)
		op.Size = uint16(zop.size)

		op.ElementType = ElementType(zop.element_type)
		op.ElementSize = ElementSize(zop.element_size)
		op.ElementCount = uint16(zop.element_count)

		op.Reg.Value = Register(zop.reg.value)

		op.Mem.Type = MemoryOperandType(zop.mem._type)
		op.Mem.Segment = Register(zop.mem.segment)
		op.Mem.Base = Register(zop.mem.base)
		op.Mem.Index = Register(zop.mem.index)
		op.Mem.Scale = uint8(zop.mem.scale)
		op.Mem.Disp.HasDisplacement = zop.mem.disp.has_displacement != 0
		op.Mem.Disp.Value = int64(zop.mem.disp.value)

		op.Ptr.Segment = uint16(zop.ptr.segment)
		op.Ptr.Offset = uint32(zop.ptr.offset)

		op.Imm.IsSigned = zop.imm.is_signed != 0
		op.Imm.IsRelative = zop.imm.is_relative != 0
		op.Imm.Value.Unsigned = uint64(*(*C.ZyanU64)(unsafe.Pointer(&zop.imm.value[0])))
		op.Imm.Value.Signed = int64(*(*C.ZyanI64)(unsafe.Pointer(&zop.imm.value[0])))
	}
	ins.Attributes = InstructionAttributes(zins.attributes)

	for i := 0; i < C.ZYDIS_CPUFLAG_MAX_VALUE+1; i++ {
		action := zins.accessed_flags[i].action
		if action != C.ZYDIS_CPUFLAG_ACTION_NONE {
			ins.AccessedFlags = append(ins.AccessedFlags, struct {
				Action CPUFlagAction
			}{
				Action: CPUFlagAction(action),
			})
		}
	}

	ins.AVX.VectorLength = uint16(zins.avx.vector_length)
	ins.AVX.Mask.Mode = MaskMode(zins.avx.mask.mode)
	ins.AVX.Mask.Reg = Register(zins.avx.mask.reg)
	ins.AVX.Broadcast.IsStatic = zins.avx.broadcast.is_static != 0
	ins.AVX.Broadcast.Mode = BroadcastMode(zins.avx.broadcast.mode)
	ins.AVX.Rounding.Mode = RoundingMode(zins.avx.rounding.mode)
	ins.AVX.Swizzle.Mode = SwizzleMode(zins.avx.swizzle.mode)
	ins.AVX.Conversion.Mode = ConversionMode(zins.avx.conversion.mode)
	ins.AVX.HasSAE = zins.avx.has_sae != 0
	ins.AVX.HasEvictionHint = zins.avx.has_eviction_hint != 0

	ins.Meta.Category = InstructionCategory(zins.meta.category)
	ins.Meta.ISASet = ISASet(zins.meta.isa_set)
	ins.Meta.ISAExt = ISAExt(zins.meta.isa_ext)
	ins.Meta.BranchType = BranchType(zins.meta.branch_type)
	ins.Meta.ExceptionClass = ExceptionClass(zins.meta.exception_class)

	for i := 0; i < int(zins.raw.prefix_count); i++ {
		zp := &zins.raw.prefixes[i]
		ins.Raw.Prefixes = append(ins.Raw.Prefixes, struct {
			Type  PrefixType
			Value uint8
		}{
			Type:  PrefixType(zp._type),
			Value: uint8(zp.value),
		})
	}

	ins.Raw.REX.W = uint8(zins.raw.rex.W)
	ins.Raw.REX.R = uint8(zins.raw.rex.R)
	ins.Raw.REX.X = uint8(zins.raw.rex.X)
	ins.Raw.REX.B = uint8(zins.raw.rex.B)
	ins.Raw.REX.Offset = uint8(zins.raw.rex.offset)

	ins.Raw.XOP.R = uint8(zins.raw.xop.R)
	ins.Raw.XOP.X = uint8(zins.raw.xop.X)
	ins.Raw.XOP.B = uint8(zins.raw.xop.B)
	ins.Raw.XOP.MMMMM = uint8(zins.raw.xop.m_mmmm)
	ins.Raw.XOP.W = uint8(zins.raw.xop.W)
	ins.Raw.XOP.VVVV = uint8(zins.raw.xop.vvvv)
	ins.Raw.XOP.L = uint8(zins.raw.xop.L)
	ins.Raw.XOP.PP = uint8(zins.raw.xop.pp)
	ins.Raw.XOP.Offset = uint8(zins.raw.xop.offset)

	ins.Raw.VEX.R = uint8(zins.raw.vex.R)
	ins.Raw.VEX.X = uint8(zins.raw.vex.X)
	ins.Raw.VEX.B = uint8(zins.raw.vex.B)
	ins.Raw.VEX.MMMMM = uint8(zins.raw.vex.m_mmmm)
	ins.Raw.VEX.W = uint8(zins.raw.vex.W)
	ins.Raw.VEX.VVVV = uint8(zins.raw.vex.vvvv)
	ins.Raw.VEX.L = uint8(zins.raw.vex.L)
	ins.Raw.VEX.PP = uint8(zins.raw.vex.pp)
	ins.Raw.VEX.Offset = uint8(zins.raw.vex.offset)
	ins.Raw.VEX.Size = uint8(zins.raw.vex.size)

	ins.Raw.EVEX.R = uint8(zins.raw.evex.R)
	ins.Raw.EVEX.X = uint8(zins.raw.evex.X)
	ins.Raw.EVEX.B = uint8(zins.raw.evex.B)
	ins.Raw.EVEX.R2 = uint8(zins.raw.evex.R2)
	ins.Raw.EVEX.MM = uint8(zins.raw.evex.mm)
	ins.Raw.EVEX.W = uint8(zins.raw.evex.W)
	ins.Raw.EVEX.VVVV = uint8(zins.raw.evex.vvvv)
	ins.Raw.EVEX.PP = uint8(zins.raw.evex.pp)
	ins.Raw.EVEX.Z = uint8(zins.raw.evex.z)
	ins.Raw.EVEX.L2 = uint8(zins.raw.evex.L2)
	ins.Raw.EVEX.L = uint8(zins.raw.evex.L)
	ins.Raw.EVEX.BLower = uint8(zins.raw.evex.b)
	ins.Raw.EVEX.V2 = uint8(zins.raw.evex.V2)
	ins.Raw.EVEX.AAA = uint8(zins.raw.evex.aaa)
	ins.Raw.EVEX.Offset = uint8(zins.raw.evex.offset)

	ins.Raw.MVEX.R = uint8(zins.raw.mvex.R)
	ins.Raw.MVEX.X = uint8(zins.raw.mvex.X)
	ins.Raw.MVEX.B = uint8(zins.raw.mvex.B)
	ins.Raw.MVEX.R2 = uint8(zins.raw.mvex.R2)
	ins.Raw.MVEX.MMMM = uint8(zins.raw.mvex.mmmm)
	ins.Raw.MVEX.W = uint8(zins.raw.mvex.W)
	ins.Raw.MVEX.VVVV = uint8(zins.raw.mvex.vvvv)
	ins.Raw.MVEX.PP = uint8(zins.raw.mvex.pp)
	ins.Raw.MVEX.E = uint8(zins.raw.mvex.E)
	ins.Raw.MVEX.SSS = uint8(zins.raw.mvex.SSS)
	ins.Raw.MVEX.V2 = uint8(zins.raw.mvex.V2)
	ins.Raw.MVEX.KKK = uint8(zins.raw.mvex.kkk)
	ins.Raw.MVEX.Offset = uint8(zins.raw.mvex.offset)

	ins.Raw.ModRM.Mod = uint8(zins.raw.modrm.mod)
	ins.Raw.ModRM.Reg = uint8(zins.raw.modrm.reg)
	ins.Raw.ModRM.Rm = uint8(zins.raw.modrm.rm)
	ins.Raw.ModRM.Offset = uint8(zins.raw.modrm.offset)

	ins.Raw.SIB.Scale = uint8(zins.raw.sib.scale)
	ins.Raw.SIB.Index = uint8(zins.raw.sib.index)
	ins.Raw.SIB.Base = uint8(zins.raw.sib.base)
	ins.Raw.SIB.Offset = uint8(zins.raw.sib.offset)

	ins.Raw.Disp.Value = int64(zins.raw.disp.value)
	ins.Raw.Disp.Size = uint8(zins.raw.disp.size)
	ins.Raw.Disp.Offset = uint8(zins.raw.disp.offset)

	for i := 0; i < 2; i++ {
		imm := &ins.Raw.Imm[i]
		zimm := &zins.raw.imm[i]
		imm.IsSigned = zimm.is_signed != 0
		imm.IsRelative = zimm.is_relative != 0
		imm.Value.Unsigned = uint64(*(*C.ZyanU64)(unsafe.Pointer(&zimm.value[0])))
		imm.Value.Signed = int64(*(*C.ZyanI64)(unsafe.Pointer(&zimm.value[0])))
		imm.Size = uint8(zimm.size)
		imm.Offset = uint8(zimm.offset)
	}

	return ins
}

// InstructionSegment is an enum of instruction segments.
type InstructionSegment int

// InstructionSegment enum values.
const (
	InstructionSegmentNone InstructionSegment = iota
	// The legacy prefixes (including ignored `REX` prefixes).
	InstructionSegmentPrefixes
	// The effective `REX` prefix byte.
	InstructionSegmentREX
	// The `XOP` prefix bytes.
	InstructionSegmentXOP
	// The `VEX` prefix bytes.
	InstructionSegmentVEX
	// The `EVEX` prefix bytes.
	InstructionSegmentEVEX
	// The `MVEX` prefix bytes.
	InstructionSegmentMVEX
	// The opcode bytes.
	InstructionSegmentOpCode
	// The `ModRM` byte.
	InstructionSegmentModRM
	// The `SIB` byte.
	InstructionSegmentSIB
	// The displacement bytes.
	InstructionSegmentDisplacement
	// The immediate bytes.
	InstructionSegmentImmediate
)

// Segment provides the logical breakdown of an instruction.
type Segment struct {
	// The type of the segment.
	Type InstructionSegment
	// The offset of the segment relative to the start of the instruction, in bytes.
	Offset uint8
	// The size of the segment, in bytes.
	Size uint8
}

// AccessedFlagsByAction returns a bitfield for the current instruction that
// has the n-th bit set, corresponding to n == CPUFlagAction, for CPU-flags
// that match the given action.
func (di *DecodedInstruction) AccessedFlagsByAction(action CPUFlagAction) (uint32, error) {
	var zflags C.ZydisCPUFlags
	ret := C.ZydisGetAccessedFlagsByAction(
		di.zdi,
		C.ZydisCPUFlagAction(action),
		&zflags,
	)
	if zyanFailure(ret) {
		return 0, fmt.Errorf("zydis: error accessing flags by action: 0x%x", ret)
	}
	return uint32(zflags), nil
}

// AccessedFlagsRead returns a bitfield for the current instruction that has
// the n-th bit set, corresponding to n == CPUFlagAction, for CPU-flags that
// are read (tested).
func (di *DecodedInstruction) AccessedFlagsRead() (uint32, error) {
	var zflags C.ZydisCPUFlags
	ret := C.ZydisGetAccessedFlagsRead(
		di.zdi,
		&zflags,
	)
	if zyanFailure(ret) {
		return 0, fmt.Errorf("zydis: error accessing flags read: 0x%x", ret)
	}
	return uint32(zflags), nil
}

// AccessedFlagsWritten returns a bitfield for the current instruction that has
// the n-th bit set, corresponding to n == CPUFlagAction, for CPU-flags that
// are written (modified, undefined).
func (di *DecodedInstruction) AccessedFlagsWritten() (uint32, error) {
	var zflags C.ZydisCPUFlags
	ret := C.ZydisGetAccessedFlagsWritten(
		di.zdi,
		&zflags,
	)
	if zyanFailure(ret) {
		return 0, fmt.Errorf("zydis: error accessing flags written: 0x%x", ret)
	}
	return uint32(zflags), nil
}

// Segments returns offsets and sizes of all logical instruction
// segments (e.g. `OPCODE`, `MODRM`, ...).
func (di *DecodedInstruction) Segments() ([]Segment, error) {
	var zsegments C.ZydisInstructionSegments
	ret := C.ZydisGetInstructionSegments(di.zdi, &zsegments)
	if zyanFailure(ret) {
		return nil, fmt.Errorf("zydis: error getting instruction segments: 0x%x", ret)
	}
	n := int(zsegments.count)
	segments := make([]Segment, n)
	for i := 0; i < n; i++ {
		s := &segments[i]
		zs := &zsegments.segments[i]

		s.Type = InstructionSegment(zs._type)
		s.Offset = uint8(zs.offset)
		s.Size = uint8(zs.size)
	}
	return segments, nil
}
