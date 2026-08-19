// Assembly implementations of rdtsc() and cpuidLeaf() for amd64.
//
// These mirror the inline assembly in the C++ libhyperdbg (CpuReadTsc and
// CpuCpuId in hardware/Cpu.cpp) that the transparency measurement code uses.
// Go assembly syntax is Plan 9 assembler; see https://go.dev/doc/asm.
//
// Register conventions in Go amd64 asm:
//   - AX, CX, DX, R8-R11 are caller-saved (scratch) — safe to clobber.
//   - BX, BP, R12-R15 are callee-saved — must save/restore if used.
//   - RDTSC writes EDX:EAX (zero-extending the upper halves of RDX/RAX).
//   - CPUID writes EAX/EBX/ECX/EDX — BX is callee-saved, so it must be saved.

#include "textflag.h"

// func rdtsc() uint64
//
// Frame: $0-8 (no locals, 8 bytes return).
TEXT ·rdtsc(SB), NOSPLIT, $0-8
	RDTSC               // EDX:EAX = TSC
	SHLQ $32, DX        // DX = upper 32 bits << 32
	ADDQ DX, AX         // AX = full 64-bit TSC
	MOVQ AX, ret+0(FP)
	RET

// func cpuidLeaf(leaf uint32) (eax, ebx, ecx, edx uint32)
//
// Frame: $8-24 (8 bytes locals to save BX, 24 bytes args+returns).
// On amd64 each argument occupies a register-sized slot (8 bytes) even when
// the type is smaller, so:
//   Args:   leaf+0(FP)  uint32  (8-byte slot, value in low 4 bytes)
//   Return: eax+8(FP)   uint32
//           ebx+12(FP)  uint32
//           ecx+16(FP)  uint32
//           edx+20(FP)  uint32
TEXT ·cpuidLeaf(SB), NOSPLIT, $8-24
	MOVQ BX, 0(SP)          // save callee-saved BX (clobbered by CPUID)
	MOVL leaf+0(FP), AX     // EAX = leaf
	CPUID                   // EAX/EBX/ECX/EDX = result
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	MOVQ 0(SP), BX          // restore BX
	RET
