//go:build amd64

package app

// readCpuidVendorString returns the 12-byte CPUID vendor string by executing
// CPUID with EAX=0. The EBX, EDX, ECX registers (in that order) hold the
// 12 ASCII bytes of the vendor string (e.g. "GenuineIntel", "AuthenticAMD").
//
// The assembly stub lives in asm_amd64.s. It mirrors the common package's
// cpuidLeafSub but is duplicated here to avoid an import cycle.
func readCpuidVendorString() string {
	eax, ebx, ecx, edx := cpuidLeaf0()
	_ = eax // max leaf — not needed for the vendor string
	var buf [12]byte
	put32 := func(dst []byte, v uint32) {
		dst[0] = byte(v)
		dst[1] = byte(v >> 8)
		dst[2] = byte(v >> 16)
		dst[3] = byte(v >> 24)
	}
	put32(buf[0:4], ebx)
	put32(buf[4:8], edx)
	put32(buf[8:12], ecx)
	return string(buf[:])
}

// cpuidLeaf0 executes CPUID with EAX=0 and returns the four result registers.
// Implemented in asm_amd64.s.
//
//go:noescape
func cpuidLeaf0() (eax, ebx, ecx, edx uint32)
