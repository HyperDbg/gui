//go:build amd64

package transparency

// rdtsc reads the Time Stamp Counter. Mirrors C++ CpuReadTsc() used by
// TransparentModeRdtscDiffVmexit and TransparentModeRdtscVmexitTracing.
//
// The implementation is in asm_amd64.s. It executes the RDTSC instruction
// (EDX:EAX = TSC) and combines the two 32-bit halves into a single uint64.
//
// RDTSC is a non-privileged instruction on x86-64; it does not trigger a
// VM-exit on hypervisors that do not intercept TSC reads, which is exactly
// what the transparency measurement code relies on.
//
//go:noescape
func rdtsc() uint64

// cpuidLeaf executes CPUID with the given leaf and returns the four result
// registers (EAX, EBX, ECX, EDX). Mirrors C++ CpuCpuId(int result[4], int leaf).
//
// CPUID is a non-privileged instruction, but it unconditionally causes a
// VM-exit when a hypervisor is present — this is the behaviour the
// transparency measurement exploits to detect hypervisors.
//
// The implementation is in asm_amd64.s.
//
//go:noescape
func cpuidLeaf(leaf uint32) (eax, ebx, ecx, edx uint32)

// cpuidLeaf0 is a convenience wrapper for cpuidLeaf(0), the leaf used by the
// transparency measurement code (TransparentModeRdtscDiffVmexit calls
// CpuCpuId(result, 0) to force a VM-exit). Returning a [4]uint32 matches the
// C++ `INT CpuidResult[4]` shape.
func cpuidLeaf0() [4]uint32 {
	a, b, c, d := cpuidLeaf(0)
	return [4]uint32{a, b, c, d}
}
