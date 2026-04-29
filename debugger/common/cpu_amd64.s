#include "textflag.h"

TEXT ·CpuId(SB), NOSPLIT, $0-24
	MOVL eaxInput+0(FP), AX
	CPUID
	MOVL AX, eax+8(FP)
	MOVL BX, ebx+12(FP)
	MOVL CX, ecx+16(FP)
	MOVL DX, edx+20(FP)
	RET

TEXT ·RdTsc(SB), NOSPLIT, $0-8
	RDTSC
	MOVL AX, retLo+0(FP)
	MOVL DX, retHi+4(FP)
	RET
