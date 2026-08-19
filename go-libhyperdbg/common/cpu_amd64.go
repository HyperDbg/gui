//go:build amd64

package common

// defaultCpuidReader executes the CPUID instruction directly. It is the
// production reader used when no override is installed via SetCpuidReader.
//
// The assembly implementation lives in asm_amd64.s. It mirrors the
// transparency package's cpuidLeaf but is duplicated here to avoid an import
// cycle (transparency is a higher-level package that may eventually depend
// on common).
type defaultCpuidReader struct{}

// Read executes CPUID.(EAX=func, ECX=subFunc) and returns the four result
// registers.
func (defaultCpuidReader) Read(func_, subFunc uint32) (eax, ebx, ecx, edx uint32) {
	return cpuidLeafSub(func_, subFunc)
}

// cpuidLeafSub executes CPUID with EAX=leaf, ECX=subLeaf and returns the
// four result registers. Implemented in asm_amd64.s.
//
//go:noescape
func cpuidLeafSub(leaf, subLeaf uint32) (eax, ebx, ecx, edx uint32)
