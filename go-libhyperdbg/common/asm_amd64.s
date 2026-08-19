// Assembly implementation of cpuidLeafSub() for amd64.
//
// Mirrors the layout of debugger/transparency/asm_amd64.s (cpuidLeaf) but
// takes an additional sub-leaf argument to support the
// CommonCpuidInstruction(func, subFunc, ...) signature.
//
// Go assembly syntax is Plan 9 assembler; see https://go.dev/doc/asm.
//
// Register conventions in Go amd64 asm:
//   - AX, CX, DX, R8-R11 are caller-saved (scratch) — safe to clobber.
//   - BX, BP, R12-R15 are callee-saved — must save/restore if used.
//   - CPUID writes EAX/EBX/ECX/EDX — BX is callee-saved, so it must be saved.

#include "textflag.h"

// func cpuidLeafSub(leaf, subLeaf uint32) (eax, ebx, ecx, edx uint32)
//
// Frame: $8-24 (8 bytes locals to save BX, 24 bytes args+returns).
// Each uint32 argument/return occupies a 4-byte slot:
//   Args:   leaf+0(FP)     uint32
//           subLeaf+4(FP)  uint32
//   Return: eax+8(FP)   uint32
//           ebx+12(FP)  uint32
//           ecx+16(FP)  uint32
//           edx+20(FP)  uint32
TEXT ·cpuidLeafSub(SB), NOSPLIT, $8-24
	MOVQ BX, 0(SP)            // save callee-saved BX (clobbered by CPUID)
	MOVL leaf+0(FP), AX       // EAX = leaf
	MOVL subLeaf+4(FP), CX    // ECX = subLeaf
	CPUID                     // EAX/EBX/ECX/EDX = result
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	MOVQ 0(SP), BX            // restore BX
	RET
