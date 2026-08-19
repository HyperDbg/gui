// Assembly implementation of cpuidLeaf0() for amd64.
//
// Mirrors the layout of debugger/transparency/asm_amd64.s (cpuidLeaf) but
// specialised for leaf 0 (no argument). CPUID writes EAX/EBX/ECX/EDX — BX is
// callee-saved on amd64, so it must be saved/restored.
//
// Go assembly syntax is Plan 9 assembler; see https://go.dev/doc/asm.

#include "textflag.h"

// func cpuidLeaf0() (eax, ebx, ecx, edx uint32)
//
// Frame: $8-16 (8 bytes locals to save BX, 16 bytes returns).
// Return values (each uint32, 4 bytes, packed):
//   eax+0(FP)  uint32
//   ebx+4(FP)  uint32
//   ecx+8(FP)  uint32
//   edx+12(FP) uint32
TEXT ·cpuidLeaf0(SB), NOSPLIT, $8-16
	MOVQ BX, 0(SP)            // save callee-saved BX (clobbered by CPUID)
	XORL AX, AX               // EAX = 0 (leaf 0)
	CPUID                     // EAX/EBX/ECX/EDX = result
	MOVL AX, eax+0(FP)
	MOVL BX, ebx+4(FP)
	MOVL CX, ecx+8(FP)
	MOVL DX, edx+12(FP)
	MOVQ 0(SP), BX            // restore BX
	RET
