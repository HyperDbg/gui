//go:build !amd64

package common

// defaultCpuidReader is a no-op reader for non-amd64 platforms. It returns
// zeros so the Common helpers degrade gracefully on architectures that do
// not implement CPUID (e.g. arm64). Tests on those platforms must install a
// fake reader via SetCpuidReader.
type defaultCpuidReader struct{}

// Read returns zeros on non-amd64 platforms.
func (defaultCpuidReader) Read(func_, subFunc uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}
